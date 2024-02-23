# Game Flex Match Developer Interconnection Guide #

## Terminology: ##

1.  Tenant: indicates the enterprise user who uses Huawei cloud services.
2.  User: indicates the user of the tenant application.
3.  Game Flex Match service: indicates the application hosting service platform provided by HUAWEI CLOUD.
4.  The tenant management side, which represents the applications, such as game lobbies, in which the tenant interacts with the GameFlexMatch service.
5.  Tenant-hosted application, denotes an application, such as a gaming application, that is hosted by a tenant on the Game Flex Match service platform.
<font color=red>Note: All interfaces involved in the client session process in this document are reserved interfaces.</font>

## Interface interaction ##

Interface interactions are categorized into management layer interfaces and application layer interfaces.

1.  Management interface: Interaction between the tenant management plane and the Game Flex Match service through RESTful APIs.
2.  Application interface: The overall flow of the tenant hosted application interacting with the GameFlexMatch service (via the integration SDK) is as follows.	

## Interface authentication ##

Authentication checks are required before calling all interfaces.

1.  Construct a login request where the password is RSA encrypted cipher, the encrypted public key needs to be consistent with the private key of the backend deployed by fleetmanager, the encryption script can be found in `/tools/cipher`. You can exec: `./sac-gfm cipher --mode encode --method rsa --text {登录密码}`

`POST URL: /v1/user/login` `Request Body`\:

```
{
    "username": "admin",
    "password": "cipher-password"
}
```

`Response Body`\:

```
{
    "username": "admin",
    "id": "63da950a-b3ea-11ed-bb27-fa163**********",
    "Auth-Token": "eyJhbGciO*************eXrA",
    "Activation": 1,
    "UserType": 9,
    "total_res_count": 1
}
```

2.  Add the `Auth-Token` field in the `Response` and its value to the `header` of the request body to access the `Game Flex Match` interface service normally.

## Management layer interface ##

### Description of the interface between the tenant management surface and the GameFlexMatch service. ###

See the `API` interface documentation for detailed interface information, and some description of the interface below:

1.  **CreateFleet Create**
Create a `fleet`, where you can specify the boot path and parameters after booting via "`process_configuration`" (these are parameters that are available to every VM in this `fleet`)
2.  **ShowFleet**
The state of the `fleet` is `Activating`, you need to query the state of the `fleet` through `ShowFleet`, only after the state changes to "`active`" it means that the `fleet` is available, and then you can create the `Server Session`.
3.  **UpdateFleetInstanceCapacity**
After `Fleet` is created, a virtual machine will be started by default to start the process, you need to wait for the state of `Fleet` to be `active`, through this interface you can modify the maximum, minimum and expected value of the virtual machine for the specified `Fleet`.
4.  **UpdateFleet**
This interface focuses on modifying `Fleet` properties, such as turning on elastic scaling, modifying instance labels, session protection policies and protection duration, etc.
5.  **CreateScalingPolicy**
After `Fleet` is created, elastic scaling is not activated by default, you need to wait for the state of `Fleet` to be `active`, call `updateFleet` to activate elastic scaling, and then you can configure your own elasticity thresholds through `CreateScalingPolicy`.
6.  **CreateServerSession**
A `server session` can be created for a `Fleet` via `CreateServerSession`, and if the `Fleet` has an available process, that process will receive the parameters of this interface. 
If the specified Fleet has no currently available VMs to allocate, the interface will return a failure, and the business can wait for some time before retrying if auto scaling policies are enabled.
7.  **ShowServerSession**
You can use `ShowServerSession` to get the connection information of the specified `ServerSession`, or you can create a `client session` for the `Server Session` to get the connection information, the `GameFlexMatch` platform manages the connection through the `client session`. The `GameFlexMatch` platform carefully manages connections through `client sessions`, which is recommended. If the `ServerSession` is not `active`, the access address will be hidden.
8.  **CreateClientSession`(reserved interface)`**
You can create a `client session` for a specified `Server Session` by `CreateClientSession`, according to this `client session`.Different players can connect to the server, for example, in a game, `10` players in a game, the game board is the `server session`, and the `client session` represents different player, these users use the same server connection address to connect to the hosting server; after the `client session` is created, it will have a validity period of `60s`, and the `client session` will become unavailable after this period of time.
9. **DeleteFleet**  
When finished, this interface is provided for the user to do a final cleanup, which cleans up all resources for all specified `fleet`

## Application Layer Interface ##

### Description of the API for tenant-hosted applications to interact with the GameFlexMatch service ###

1.  The following table lists the callback APIs of hosted applications (interfaces invoked by the GameFlexMatch service). The three API hosting services need to be implemented based on their own service logic.

|       API Name       | API Description                                                                                                                              |
|:--------------------:|:-------------------------------------------------------------------------------------------------------------------------------------------- |
|    OnHealthCheck     | Returns the health status of the hosted service, only applications with a health status of true are assigned to the server session. |
| OnStartServerSession | Receive the server session creation information, the business needs to execute the server session creation logic here.                   |
|  OnProcessTermiante  | Receive information about service shutdown, where the business needs to shut down the process correctly to reclaim resources.         |

2.  A list of the main APIs (GameFlexMatch interfaces called by the hosted application) for the hosted application is as follows.

|                 API Name                  | API Description                                                                                                                                            |
|:-----------------------------------------:|:---------------------------------------------------------------------------------------------------------------------------------------------------------- |
|               ProcessReady                | Register the hosted application process information to notify the GameFlexMatch service process that it has been started.                  |
|           ActivateServerSession           | Activating a server session indicates that the corresponding server session has been created and can be used for subsequent processes.                        |
|        AcceptClientSession`(Reserved interface)`        | The client session has been successfully connected.                                                                                                        |
|        RemoveClientSession`(Reserved interface)`        | Terminate the client session.                                                                                                                              |
|      DescribeClientSessions`(Reserved interface)`       | Get information about all client sessions for a given server session.                                                                               |
| UpdateClientSessionCreationPolicy`(Reserved interface)` | Updates the client session creation policy for the specified server session, mainly whether to allow new client session access. |
|        TermianteGameServerSession         | The server session ends, terminating the responding server session.                                                                               |
|              PrcocessEnding               | Process recycling is complete and the process can be shut down normally.                                                                                |

## Hosted Application Integration SDK Tutorial ##

Prerequisites: tutorials on installing dependencies (grpc, protobuf versions, etc.) to be added.

1.  Available demos
    Currently provides demos in c#, go, c++ (see: `demo/go`, `demo/csharp`, `demo/cpp`).
2.  After the hosted service has started and confirmed that it is ready to provide services, the hosted service process needs to call the ProcessReady API to notify the Game Flex Match platform that it has started and is ready to be assigned a server session.The GameFlexMatch platform receives this notification, and then it will set the state of the process to After receiving the notification, the GameFlexMatch platform will set the state of the process to Activating, at which point the process is not yet available for server session allocation.
    `C++ Example`
    ```c++
    bool ProcessReady(const std::vector<std::string>& log_paths,
                      int32_t client_port, int32_t grpc_port, int32_t pid) {
        ProcessReadyRequest *request;
        for (const auto& path : log_paths) {
            request->add_logpathstoupload(path);
        }
        request->set_clientport(client_port);
        request->set_grpcport(grpc_port);
        request->set_pid(pid);
    
        AuxProxyResponse response;
        ClientContext context;
    
        Status status = stub_->ProcessReady(&context, request, &response);
        if (status.ok()) {
            //Processing success
            return true;
        } else {
            //Handling Failures
            return false;
        }
    }
    ```
    
    `C# Example`
    ```c#
    public static AuxProxyResponse ProcessReady(string[] logPath, int clientPort, int grpcPort)
    {
        logger.Println($"Getting process ready, LogPath: {logPath}, ClientPort: {clientPort}, GrpcPort: {grpcPort}");
        var req = new ProcessReadyRequest{
            ClientPort = clientPort,
            GrpcPort = grpcPort,
            //pid indicates the ID of the current process.
            Pid = Process.GetCurrentProcess().Id,
        };
        req.LogPathsToUpload.Add(logPath);         //The repeated type is read-only after the PB is parsed. You need to add the repeated type.           
        return GrpcClient.ScaseClient.ProcessReady(req, meta);
    }
    ```
3.  When the GameFlexMatch platform receives the ProcessReady notification, it will call the onHealthCheck of the process to confirm that the process has entered the ready state, and then set the state of the process to Active.
    `C++ Example`
    ```c++
    //The object must inherit ProcessGrpcSdkService::Service.
    class ProcessGrpcSdkServiceImpl final : public ProcessGrpcSdkService::Service
    //Receiving the health check request
    Status OnHealthCheck(ServerContext* context, const HealthCheckRequest* request,
                         HealthCheckResponse* response) override {
        response->set_healthstatus(this->healthStatus);
        return Status::OK;
    }
    ```
    
    `C# Example`
    ```c#
    //Object needs to inherit ProcessGrpcSdkService.ProcessGrpcSdkServiceBase
    public class ServerSdk : ProcessGrpcSdkService.ProcessGrpcSdkServiceBase
    
    public override Task<HealthCheckResponse> OnHealthCheck(HealthCheckRequest request, ServerCallContext context)
    {
        logger.Println($"OnHealthCheck, HealthStatus: {ScaseManager.HealthStatus}");
        return Task.FromResult(new HealthCheckResponse{
            HealthStatus = ScaseManager.HealthStatus
        });
    }
    ```
4.  The tenant management side can create a server session by calling CreateServerSession API, and the server session will be bound to a hosted application process of the specified fleet. when Game Flex Match platform receives the request to create a server session, it will asynchronously call the When GameFlexMatch platform receives the server session creation request, it will asynchronously call the onStartServerSession API to notify the hosting application process and set the server session to "Activating" state.
    `C++ Example`
    ```c++
    //Receiving a Game Session
    Status OnStartServerSession(ServerContext* context, const StartServerSessionRequest* request,ProcessResponse* response) override { 
        //Save Game Session
        scaseClient.SetGameServerSession(request->serversession());
        //Activate session
        Status status = scaseClient.ActivateServerSession(request->serversession().serversessionid(), request->serversession().maxclients());
        response->Clear();
        return status;
    }
    ```
    
    `C# Example`
    ```c#
    public override Task<ProcessResponse> OnStartServerSession(StartServerSessionRequest request, ServerCallContext context)
    {
        logger.Println($"OnStartServerSession, request: {request}");
        ScaseManager.SetServerSession(request.ServerSession);
        return Task.FromResult(new ProcessResponse());
    }
    ```
5.  After receiving the onStartServerSession message, the hosting application process needs to process its own services. After all the services are processed, the hosting application process needs to invoke the ActivateServerSession API to notify the GameFlexMatch platform that the current server session has been activated. GameFlexMatch sets the server session status to active.
    `C++ Example`
    ```c++
    //Report server session activation
    Status ActivateServerSession(const std::string& session_id, int32_t max_clients) {
        ActivateServerSessionRequest request;
        request.set_serversessionid(session_id);
        request.set_maxclients(max_clients);
    
        AuxProxyResponse response;
        ClientContext context;
    
        Status status = stub_->ActivateServerSession(&context, request, &response);
        return status;
    }
    ```
    

    `C# Example`
    ```c#
    public static AuxProxyResponse ActivateServerSession(string serverSessionId, int maxClients)
    {
        logger.Println($"Activating game server session, ServerSessionId: {serverSessionId}, MaxClients: {maxClients}");
        var req = new ActivateServerSessionRequest{
            ServerSessionId = serverSessionId,
            MaxClients = maxClients,
        };  
        return GrpcClient.ScaseClient.ActivateServerSession(req, meta);
    }
    ```
6.  `(Reserved interface)`Tenant management surface through the CreateClientSession interface to the user to obtain connection information, user access to use client session access, hosting surplus public process will call AcceptClientSession to notify the GameFlexMatch platform the current client session has been accessed, the GameFlexMatch platform will be set to change the client session status to active; if the client session is not accessed after 60 seconds, the status will change to timeout. GameFlexMatch platform will set the client session status to active; if the client session is not accessed after 60 seconds of creation, the status will be changed to timeout, and the client session can no longer be used.
    `C++ Example`
    ```c++
    //Report Client Session Access
    bool AcceptClientSession(const std::string& session_id, const std::string&client_id) { 
        AcceptClientSessionRequest request;
        request.set_serversessionid(session_id);
        request.set_clientsessionid(client_id);
    
        AuxProxyResponse response;
        ClientContext context;
    
        Status status = stub_->AcceptClientSession(&context, request, &response);
        if (status.ok()) {
            //Processing success
            return true;
        } else {
            //Handling Failures
            return false;
        }
    }
    ```
    
    `C# Example`
    ```c#
    public static AuxProxyResponse AcceptClientSession(string ClientSessionId)
    {
        logger.Println($"Accepting Client session, ClientSessionId: {ClientSessionId}");
        var req = new AcceptClientSessionRequest{
            ServerSessionId = serverSession.ServerSessionId,
            ClientSessionId = ClientSessionId,
        };            
        return GrpcClient.ScaseClient.AcceptClientSession(req, meta);
    }
    ```
7.  `(Reserved interface)`After a player disconnects, the hosting application process needs to call the RemoveClientSession API to remove the player, and the GameFlexMatch platform will set the corresponding client session to the "complete" state and reclaim the quota.
    `C++ Example`
    ```c++
    //Report Client Session Away
    bool RemoveClientSession(const std::string& session_id, const std::string& client_id) {
        RemoveClientSessionRequest request;
        request.set_serversessionid(session_id);
        request.set_clientsessionid(client_id);
    
        AuxProxyResponse response;
        ClientContext context;
    
        Status status = stub_->RemoveClientSession(&context, request, &response);
        if (status.ok()) {
            //Processing success
            return true;
        } else {
            //Handling Failures
            return false;
        }
    }
    ```
    

    `C# Example`
    ```c#
    public static AuxProxyResponse RemoveClientSession(string ClientSessionId)
    {
        logger.Println($"Removing Client session, ClientSessionId: {ClientSessionId}");
        var req = new RemoveClientSessionRequest{
            ServerSessionId = serverSession.ServerSessionId,
            ClientSessionId = ClientSessionId,
        };            
        return GrpcClient.ScaseClient.RemoveClientSession(req, meta);
    }
    ```
8.  After a server session has ended, the hosting application process needs to call the TermianteServerSession API to notify the GameFlexMatch platform to set the server session state to terminated.
    `C++ Example`
    ```c++
    //Report Server Session End
    bool TerminateServerSession(const std::string& session_id) {
        TerminateServerSessionRequest request;
        request.set_serversessionid(session_id);
    
        AuxProxyResponse response;
        ClientContext context;
    
        Status status = stub_->TerminateServerSession(&context, request, &response);
        if (status.ok()) {
            return true;
        } else {
            return false;
        }
    }
    ```
    

    `C# Example`
    ```c#
    public static AuxProxyResponse TerminateServerSession()
    {
        logger.Println($"Terminating game server session, ServerSessionId: {serverSession.ServerSessionId}");
        var req = new TerminateServerSessionRequest{
            ServerSessionId = serverSession.ServerSessionId
        };            
        return GrpcClient.ScaseClient.TerminateServerSession(req, meta);
    }
    ```
9.  The GameFlexMatch platform calls onProcessTermainte to notify the hosted application process to perform resource reclamation and shut down the process  (Note, the action does not change the session state).if it wants to shut down the hosted application process (e.g., if a tenant deletes a fleet)
    `C++ Example`
    ```c++
    //End game progress
    Status OnProcessTerminate(ServerContext* context, const ProcessTerminateRequest* request,ProcessResponse* response) override {
        int64_t terminationTime = request->terminationtime();
        scaseClient.HandingTerminatingProcess(terminationTime);  //Taking over the Processes to Be Cleared
        response->Clear();
        return Status::OK;
    }
    ```
    

    `C# Example`
    ```c#
    public override Task<ProcessResponse> OnProcessTerminate(ProcessTerminateRequest request, ServerCallContext context)
    {
        logger.Println($"OnProcessTerminate, request: {request}");
        //Set the process end time.
        ScaseManager.SetTerminationTime(request.TerminationTime);
        //Terminate Game Server Session
        ScaseManager.TerminateServerSession();
        //Process exit
        ScaseManager.ProcessEnding();
        return Task.FromResult(new ProcessResponse());
    }
    ```
10. The hosted application process needs to call ProcessEnding before shutting down to notify the Game Flex Match platform to set the state of its own process object to Termianted.
    `C++ Example`
    ```c++
    //End of reporting process
    bool ProcessEnding(int32_t pid) {
    
        ProcessEndingRequest request;
        request.set_pid(pid);
    
        AuxProxyResponse response;
        ClientContext context;
    
        Status status = stub_->ProcessEnding(&context, request, &response);
        if (status.ok()) {
            //Processing success
            return true;
        } else {
            //Handling Failures
            return false;
        }
    }
    ```
    
    `C# Example`
    ```c#
    public static AuxProxyResponse ProcessEnding()
    {
        logger.Println($"Process ending, pid: {pid}");
        var req = new ProcessEndingRequest();            
        return GrpcClient.ScaseClient.ProcessEnding(req, meta);
    }
    ```
11. `(Reserved interface)`The hosted application process can call DescribeClientSessions to get all client session information for a specified server session based on business needs.
    `C++ Example`
    ```c++
    //Obtaining Client Session Information
    bool DescribeClientSessions(const std::string& session_id, const std::string& client_id,
                                const std::string& status_filter, const std::string& next_token,
                                int32_t limit) {
        DescribeClientSessionsRequest request;
        request.set_serversessionid(session_id);
        request.set_clientid(client_id);
        request.set_clientsessionid(client_id);
        request.set_clientsessionstatusfilter(status_filter);
        request.set_nexttoken(next_token);
        request.set_limit(limit);
    
        DescribeClientSessionsResponse response;
        ClientContext context;
    
        Status status = stub_->DescribeClientSessions(&context, request, &response);
        if (status.ok()) {
            //Processing success
            return true;
        } else {
            //Handling Failures
            return false;
        }
    }
    ```
    
    `C# Example`
    ```c#
    public static DescribeClientSessionsResponse DescribeClientSessions(
        string ServerSessionId, string ClientId, string ClientSessionId, string ClientSessionStatusFilter, string nextToken, int limit)
    {
        logger.Println($"Describing Client session, ServerSessionId: {ServerSessionId}, \
                        ClientId: {ClientId}, ClientSessionId: {ClientSessionId}, \
                        ClientSessionStatusFilter: {ClientSessionStatusFilter}, \
                        NextToken: {nextToken}, Limit: {limit}");
        var req = new DescribeClientSessionsRequest{
            ServerSessionId = ServerSessionId,
            ClientId = ClientId,
            ClientSessionId = ClientSessionId,
            ClientSessionStatusFilter = ClientSessionStatusFilter,
            NextToken = nextToken,
            Limit = limit,
        };            
        return GrpcClient.ScaseClient.DescribeClientSessions(req, meta);
    }
    ```
12. `(Reserved interface)`The hosted application process can call the UpdateClientSessionCreationPolicy interface to update the client session creation policy (whether to allow new users to access the current server session) based on business needs.
    `C++ Example`
    ```c++
    //Update Client Session Join Policy
    bool UpdateClientSessionCreationPolicy(const std::string& session_id,const std::string& new_policy) {
    
        UpdateClientSessionCreationPolicyRequest request;
        request.set_serversessionid(session_id);
        request.set_newclientsessioncreationpolicy(new_policy);
    
        AuxProxyResponse response;
        ClientContext context;
    
        Status status = stub_->UpdateClientSessionCreationPolicy(&context, request, &response);
        if (status.ok()) {
            //Processing success
            return true;
        } else {
            //Handling Failures
            return false;
        }
    }
    ```
    
    `C# Example`
    ```c#
    public static AuxProxyResponse UpdateClientSessionCreationPolicy(string newPolicy)
    {
        logger.Println($"Updating Client session creation policy, newPolicy: {newPolicy}");
        var req = new UpdateClientSessionCreationPolicyRequest{
            ServerSessionId = serverSession.ServerSessionId,
            NewClientSessionCreationPolicy = newPolicy,
        };            
        return GrpcClient.ScaseClient.UpdateClientSessionCreationPolicy(req, meta);
    }
    ```
13. Start the grpc service.
    `C++ Example`
    ```c++
    void RunServer() {
        srand(time(0)); //Initializing a random number seed
        int port = rand() % 58977 + 1025; //Generates random integers between 1025 and 60001.
        std::string server_address("0.0.0.0:" + std::to_string(port));
        //Creating a Service Builder
        ServerBuilder builder;
        //Adding a Listening Address
        builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
        //Registration service implementation
        builder.RegisterService(&service);
        //Build Services
        std::unique_ptr<Server> server(builder.BuildAndStart());
        std::cout << "Server listening on " << server_address << std::endl;
        //Blocking wait-to-end signal
        server->Wait();
    }
    ```
    
    `C# Example`
    ```c#
    public class Program
    {
        public static int ClientPort = PortServer.GenerateRandomPort(2000, 6000);
        public static int GrpcPort = PortServer.GenerateRandomPort(6001, 10000);
    
        public static void Main(string[] args)
        {
            CreateHostBuilder(args).Build().Run();
        }
    
        public static IHostBuilder CreateHostBuilder(string[] args) =>{
            Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    webBuilder.ConfigureKestrel(options =>
                    {
                        //gRPC Port (Setup a HTTP/2 endpoint without TLS.)
                        options.ListenAnyIP(GrpcPort, o => o.Protocols = 
                            HttpProtocols.Http2);
    
                        //HTTP Port
                        options.ListenAnyIP(ClientPort);
                    });
    
                    webBuilder.UseStartup<Startup>();
                });
        }
    }
    ```
14. Connect to the Game Flex Match platform's grpc service
    `C++ Example` 
    ```c++
    void RunClient() {
        scaseClient.pid = to_string((int)getpid());
        std::vector<std::string> log_paths_to_upload = {"/path/to/log1", "/path/to/log2"};
        int32_t client_port = 8080;
        int32_t pid = 12345;
        bool response = scaseClient.ProcessReady(log_paths_to_upload, client_port, grpc_port, pid);
        std::cout << "Response: " << response << std::endl;
        if (response) {
            //Invoking succeeded.
        } else {
            //Invoking failed.
        }
    }
    ```
    
    `C# Example`
    ```c#
    public class GrpcClient
    {
        //Port 60002 is the start port of the GameFlexMatch grpc service.
        private static string agentAdress = "127.0.0.1:60002";
    
        public static ProcessGrpcSdkService.ProcessGrpcSdkServiceClient ProcessServerClient
        {
            get
            {
                Channel channel = new Channel(agentAdress, ChannelCredentials.Insecure);
                return new ProcessGrpcSdkService.ProcessGrpcSdkServiceClient(channel);
            }
        }
    
        public static ScaseGrpcSdkService.ScaseGrpcSdkServiceClient ScaseClient
        {
            get
            {
                Channel channel = new Channel(agentAdress, ChannelCredentials.Insecure);
                return new ScaseGrpcSdkService.ScaseGrpcSdkServiceClient(channel);
            }
        }
    }
    ```

