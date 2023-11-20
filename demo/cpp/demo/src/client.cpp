/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

#include <iostream>
#include <memory>
#include <string>
#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>
#include <grpcpp/ext/proto_server_reflection_plugin.h>
#include <thread>
#include "help.h"

#include"../proto/auxproxy_grpc_service.grpc.pb.h"
#include"../proto/process_grpc_service.grpc.pb.h"
using std::cout;
using std::endl;
using std::string;

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using auxproxyService::AuxProxyResponse;
using auxproxyService::Error;
using auxproxyService::ScaseGrpcSdkService;
using auxproxyService::ProcessReadyRequest;
using auxproxyService::ActivateServerSessionRequest;
using auxproxyService::TerminateServerSessionRequest;
using auxproxyService::ProcessEndingRequest;

 // when using process grpc server, the role is server.
using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using processService::ProcessGrpcSdkService;
using processService::HealthCheckRequest;
using processService::HealthCheckResponse;
using processService::StartServerSessionRequest;
using processService::ServerSession;
using processService::ProcessResponse;
using processService::SessionProperty;
using processService::ProcessTerminateRequest;


string GlobalSessionId;
class ScaseClient
{
public:
	ScaseClient(std::shared_ptr<Channel> channel)
		: stub_(ScaseGrpcSdkService::NewStub(channel)) {}
	
	string sessionId;
	string getSessionId() {
		return sessionId;
	}
	bool ProcessReady(const std::vector<std::string>& log_paths, int32_t client_port, int32_t grpc_port, int32_t pid)
	{
		std::cout << "[ProcessReady] process ready." << std::endl;
		ProcessReadyRequest request;
		for (const auto& path : log_paths)
		{
			request.add_logpathstoupload(path);
		}
		request.set_clientport(client_port);
		request.set_grpcport(grpc_port);
		request.set_pid(pid);

		AuxProxyResponse response;
		ClientContext context;

		Status status = stub_->ProcessReady(&context, request, &response);
		if (status.ok())
		{
			// handle success
			std::cout << "[ProcessReady] process ready success." << std::endl;
			return true;
		}
		else
		{
			// handle failure
			std::cout << "[ProcessReady] process ready failed." << std::endl;
			return false;
		}
	}


	bool ActivateServerSession(const string& session_id, int32_t max_clients) {
		ActivateServerSessionRequest request;
		request.set_serversessionid(session_id);
		request.set_maxclients(max_clients);
		AuxProxyResponse response;
		ClientContext context;

		Status status = stub_->ActivateServerSession(&context, request, &response);
		if (status.ok()) {
			this->sessionId = sessionId;
			GlobalSessionId = sessionId;
			cout << "[ActivateServerSession] session " << session_id << " active now " << endl;
			return true;
		}
		cout << status.error_message() << endl;
		return false;
	}

	bool TerminateServerSession(const string sessionId)
	{
		std::cout << "[TerminateServerSession] terminated Client session." << std::endl;

		TerminateServerSessionRequest request;
		request.set_serversessionid(sessionId);

		AuxProxyResponse response;
		ClientContext context;

		Status status = stub_->TerminateServerSession(&context, request, &response);
		if (status.ok())
		{
			// handle success
			std::cout << "[TerminateServerSession] success to terminate server session." << std::endl;
			return true;
		}
		else
		{
			// handle failure
			std::cout << "[TerminateServerSession] failed to terminate server session." << std::endl;
			return false;
		}
	}

	bool ProcessEnding(const int32_t pid) {
		cout << "[ProcessEnding] process " << pid << "prepare to ending" << endl;
		//TerminateprocessRe
		AuxProxyResponse response;
		ClientContext context;
		ProcessEndingRequest request;
		request.set_pid(pid);
		Status status = stub_->ProcessEnding(&context, request, &response);
		if (status.ok()) {
			cout << "[ProcessEnding] ending process success: " << pid << endl;
			return true;
		}
		else {
			cout << "[ProcessEnding] ending process failed:" << pid << endl;
			return false;
		}
	}
	
private:
	std::unique_ptr<ScaseGrpcSdkService::Stub> stub_;
};


class ProcessServiceImp final : public ProcessGrpcSdkService::Service {
public:
	bool existFlag = false;
	string auxproxyIP = "localhost:60002";

	string serverSessionId;
	string fleetId;
	string name;
	int32_t maxClients;
	string ipAddress;
	Status OnStartServerSession(ServerContext* context, const StartServerSessionRequest* request, ProcessResponse* response) {
		ServerSession ss = request->serversession();
		cout << "receive session:" << endl;
		this->serverSessionId = ss.serversessionid();
		cout << "server session id: " << this->serverSessionId << endl;
		this->fleetId = ss.fleetid();
		cout << "fleetn id: " << fleetId << endl;
		this->name = ss.name();
		cout << "name: " << name << endl;
		this->maxClients = ss.maxclients();
		cout << "maxClients: " << maxClients << endl;
		bool joinable = ss.joinable();
		cout << "joinable: " << joinable << endl;
		auto sp = ss.sessionproperties();
		for (auto element : sp) {
			cout << "properity key: " << element.key() << endl;
			cout << "properity value: " << element.value() << endl;
		}
		int32_t port = ss.port();
		cout << "port: " << port << endl;
		this->ipAddress = ss.ipaddress();
		cout << "ipAddress: " << ipAddress << endl;
		string sessionData = ss.sessiondata();
		cout << "sessionData: " << sessionData << endl;
		cout << "activing session, please wait ..." << endl;

		// activate session
		Sleep(1000);

		cout << "session active, player can enter/start game session" << endl;

		// call back
		std::string target_str = this->auxproxyIP;
		ScaseClient client(grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials()));
		bool ActivateServerSessionResponse = client.ActivateServerSession(this->serverSessionId, this->maxClients);
		if (ActivateServerSessionResponse) {
			return Status::OK;
		}

		return Status::CANCELLED;
	}

	Status OnHealthCheck(ServerContext* context, const HealthCheckRequest* request, HealthCheckResponse* response) {
		cout << "[OnHealthCheck] get health check request" << endl;
		bool healthStatus = true;
		response->set_healthstatus(healthStatus);


		return Status::OK;
	}

	Status OnProcessTerminate(ServerContext* context, const ProcessTerminateRequest* request, ProcessResponse* response) {
		cout << "[OnHealthCheck] get process terminate " << endl;
		int64_t MaxTerminatedTime = request->terminationtime();
		cout << "[OnHealthCheck] process terminate time " << MaxTerminatedTime << endl;
		Sleep(MaxTerminatedTime);
		existFlag = true;
		cout << "[OnHealthCheck] process terminated" << endl;
		return Status::OK;
	}
	bool getExist() {
		return existFlag;
	}
};

void RunProcessReady(ScaseClient& greeter,int32_t grpc_port,int32_t client_port,int32_t pid) {
	std::vector<std::string> logPath;

	logPath.push_back("/root/log");
	std::cout << "[RunProcessReady] send process ready " << std::endl;
	bool ready = greeter.ProcessReady(logPath, client_port, grpc_port, pid);
	std::cout << "[RunProcessReady] process ready: " << ready << std::endl;
}

void RunActivateServerSession(ScaseClient& greeter, string sessionId, int32_t maxClient) {
	bool ActivateServerSessionResponse = greeter.ActivateServerSession(sessionId, maxClient);
	if (ActivateServerSessionResponse) {
		cout << "[RunActivateServerSession] success activate serversession" << endl;
	}
	else {
		cout << "[RunActivateServerSession] activate server session failed" << endl;
	}
	printf("[RunActivateServerSession] game %s activated \n", sessionId.c_str());


	for (int i = 1; i <= 5; i++) {
		cout << "-------  gaming running " << i << " second  ----------" << endl;
		Sleep(1000);
	}
	printf("[RunActivateServerSession] game %s finish \n", sessionId.c_str());
}

void RunTerminateServerSession(ScaseClient& greeter, string sessionId) {
	bool TerminateServerSessionResp = greeter.TerminateServerSession(sessionId);
	if (TerminateServerSessionResp) {
		printf("[RunActivateServerSession] game %s terminate success \n", sessionId.c_str());
	}
	else {
		printf("[RunActivateServerSession] game %s terminate failed \n", sessionId.c_str());
	}
}

void RunProcessEnding(ScaseClient& greeter, int32_t pid) {
	bool TerminateServerSessionResp = greeter.ProcessEnding(pid);
	if (TerminateServerSessionResp) {
		printf("[RunActivateServerSession] game %d terminate success \n", pid);
	}
	else {
		printf("[RunActivateServerSession] game %d terminate failed \n", pid);
	}
	
}


void RunServer(int port) {
	string server_address = "localhost:" + to_string(port);

	ProcessServiceImp service;

	grpc::EnableDefaultHealthCheckService(true);
	grpc::reflection::InitProtoReflectionServerBuilderPlugin();
	ServerBuilder builder;
	builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());

	builder.RegisterService(&service);

	std::unique_ptr<Server> server(builder.BuildAndStart());
	std::cout << "Server listening on " << server_address << std::endl;
	//server->Wait();
	while (!service.getExist()) {
		cout << "process running" << endl;
		std::this_thread::sleep_for(std::chrono::seconds(1));
	}
	server->Shutdown();
}

void RunClient() {
	// auxproxy client port 60002
	std::string target_str = "localhost:60002";
	ScaseClient greeter(
		grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials()));

	int32_t client_port = 8080;
	int32_t pid = 123;
	int32_t grpc_prot = 50051;

	// for auxproxy call back
	std::thread callback(&RunServer, grpc_prot);


	Sleep(1000);
	// 1. send process ready to auxproxy
	RunProcessReady(greeter, grpc_prot,client_port,pid);
	cout << "--------------- process ready --------------" << endl;
	// 2.wait "onStartServerSession" to write session id;
	string sessionId = "session_id";
	int32_t maxClients = 1;
	// 3. start game session
	RunActivateServerSession(greeter, sessionId, maxClients);
	// 4. run terminate server session
	RunTerminateServerSession(greeter, "session id");
	// 5. terminate process
	RunProcessEnding(greeter, pid);
	
	callback.join();

}



int main(int argc, char** argv)
{
	
	RunClient();

	return 0;
}
