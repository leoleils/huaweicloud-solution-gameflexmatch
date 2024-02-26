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
#include <thread>
#include <grpcpp/ext/proto_server_reflection_plugin.h>
#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>
#include "help.h"
#include"../proto/process_grpc_service.grpc.pb.h"
#include"../proto/auxproxy_grpc_service.grpc.pb.h"
using std::cout;
using std::endl;
using std::string;
using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;

using auxproxyService::AuxProxyResponse;
using auxproxyService::Error;
using auxproxyService::ScaseGrpcSdkService;
using auxproxyService::ProcessReadyRequest;
using auxproxyService::ActivateServerSessionRequest;
using auxproxyService::TerminateServerSessionRequest;
using auxproxyService::ProcessEndingRequest;



using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using processService::ProcessGrpcSdkService;
using processService::ServerSession;
using processService::SessionProperty;
using processService::ProcessResponse;
using processService::StartServerSessionRequest;
using processService::HealthCheckRequest;
using processService::HealthCheckResponse;
using processService::ProcessTerminateRequest;


// client call back
class ProcessServiceClient {

public:
	ProcessServiceClient(std::shared_ptr<Channel> channel)
		: stub_(ProcessGrpcSdkService::NewStub(channel)) {}

	bool OnHealthCheck() {
		cout << "send on health check " << endl;
		HealthCheckRequest request;
		HealthCheckResponse response;
		ClientContext context;
		Status status = stub_->OnHealthCheck(&context, request, &response);
		if (status.ok()) {
			cout << "health check success response: " << response.healthstatus() << endl;
		}
		else {
			cout << "health check failed" << endl;
		}
		bool healthStatus = response.healthstatus();
		cout << "health status is: " << healthStatus << endl;
		return healthStatus;
	}

	string OnStartServerSession() {
		cout << "start on start server session " << endl;
		StartServerSessionRequest request;
		ProcessResponse response;
		ClientContext context;
		ServerSession* ss = new ServerSession();
		string serverSessionId = "session_id_" + random_string(32);
		string fleetId = "fleet id";
		string name = "name";
		int32_t maxClients = 1;
		bool joinable = true;
		int32_t port = 8080;
		string ipAddress = "0.0.0.0";
		string sessionData = "session data";
		request.set_allocated_serversession(ss);
		SessionProperty sessionProperties;
		sessionProperties.set_key("key");
		sessionProperties.set_value("value");

		ss->set_serversessionid(serverSessionId);
		ss->set_fleetid(fleetId);
		ss->set_name(name);
		ss->set_maxclients(maxClients);
		ss->set_joinable(joinable);
		ss->set_sessiondata(sessionData);
		ss->set_ipaddress(ipAddress);
		ss->set_sessiondata(sessionData);

		Status status = stub_->OnStartServerSession(&context, request, &response);

		if (status.ok()) {
			return "ok";
		}
		cout << "error: " << status.error_message() << endl;
		return "failed";
	}
	bool OnProcessTerminate(int64_t terminatedTime) {
		cout << "start OnProcessTerminate " << endl;
		ProcessTerminateRequest request;
		ProcessResponse response;
		ClientContext context;

		request.set_terminationtime(terminatedTime);
		Status status = stub_->OnProcessTerminate(&context, request, &response);

		if (status.ok()) {
			return true;
		}
		else {
			cout << status.error_message() << endl;
			return false;
		}
		return true;
	}


private:
	std::unique_ptr<ProcessGrpcSdkService::Stub> stub_;
};

// Logic and data behind the server's behavior.
class ScaseServiceImpl final : public ScaseGrpcSdkService::Service
{
public:
	void RunClient(int32_t grpc_port) {
		string target_str = "localhost:" + std::to_string(grpc_port);
		ProcessServiceClient psc(
			grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials()));
		cout << "call back to " << target_str << endl;
		Sleep(1000);
		bool responseHealth = psc.OnHealthCheck();
		if (responseHealth) {
			cout << "on start health check success response: " << responseHealth << endl;
		}
		else {
			cout << "on start health check success false: " << responseHealth << endl;
		}

		string responseStartServer = psc.OnStartServerSession();
		cout << "on start server session success response " << responseStartServer << endl;

		int64_t terminatedTime = 10;
		bool response = psc.OnProcessTerminate(terminatedTime);
		if (response) {
			cout << "on start terminate success response: " << response << endl;
		}
		else {
			cout << "on start terminate success failed: " << response << endl;
		}
	}
	Status ProcessReady(ServerContext* context, const ProcessReadyRequest* request, AuxProxyResponse* response)
	{
		std::time_t t = std::time(nullptr);
		std::tm* local_time = std::localtime(&t);
		cout << "[ProcessReady] get process ready, time " << std::asctime(local_time) << endl;
		std::cout <<"client port: " << request->clientport() << std::endl;
		std::cout << "client grpc port: " << request->grpcport() << std::endl;
		std::cout << "client pid: " << request->pid() << std::endl;
		int32_t target_grpc_port = request->grpcport();
		std::thread callbackThread(&ScaseServiceImpl::RunClient,this,target_grpc_port);
		callbackThread.detach();
		return Status::OK;
	}
	Status ActivateServerSession(ServerContext* context, const ActivateServerSessionRequest* request, AuxProxyResponse* response) 
	{
		cout << "[ActivateServerSession] get activate session " << endl;
		std::cout << "session id: " << request->serversessionid() << std::endl;
		cout << "max client: " << request->maxclients() << endl;
		return Status::OK;
	}

	Status TerminateServerSession(ServerContext* context, const TerminateServerSessionRequest* request, AuxProxyResponse* response) {
		cout << "[TerminateServerSession] terminated server session: " << request->serversessionid() << endl;

		return Status::OK;
	}

	Status ProcessEnding(ServerContext* context, const ProcessEndingRequest* request, AuxProxyResponse* response) {
		cout << "[ProcessEnding] terminated process : " << request->pid() << endl;
		
		return Status::OK;
	}
};


void RunOnStartServerSession(ProcessServiceClient& psc) {

	string response = psc.OnStartServerSession();
	cout << "on start server session success response " << response << endl;

}

void RunOnHealthCheck(ProcessServiceClient& psc) {

	bool response = psc.OnHealthCheck();
	if (response) {
		cout << "on start health check success response: " << response << endl;
	}
	else {
		cout << "on start health check success false: " << response << endl;
	}


}

void RunProcessTerminate(ProcessServiceClient& psc) {
	int64_t terminatedTime = 10;
	bool response = psc.OnProcessTerminate(terminatedTime);
	if (response) {
		cout << "on start terminate success response: " << response << endl;
	}
	else {
		cout << "on start terminate success failed: " << response << endl;
	}
}

void RunServer()
{
	std::string server_address("0.0.0.0:60002");
	ScaseServiceImpl service;

	grpc::EnableDefaultHealthCheckService(true);
	grpc::reflection::InitProtoReflectionServerBuilderPlugin();
	ServerBuilder builder;
	builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());

	builder.RegisterService(&service);

	std::unique_ptr<Server> server(builder.BuildAndStart());
	std::cout << "Server listening on " << server_address << std::endl;
	
	server->Wait();
	cout << "waiting to shutdown" << endl;
	server->Shutdown();
}

void RunClient(int32_t target_port) {
	string target_str = "0.0.0.0:" + std::to_string(target_port);
	ProcessServiceClient psc(
		grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials()));

	RunOnStartServerSession(psc);
	RunOnHealthCheck(psc);
	RunProcessTerminate(psc);
}

int main(int argc, char** argv)
{
	std::thread th_Server(&RunServer);

	th_Server.join();
	return 0;
}
