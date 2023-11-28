using System;
using Grpc.Core;
using System.Diagnostics;
using ProcessService;
using AuxproxyService;

namespace CSharpDemo.Models
{
    public static class GseManager
    {
        private static long terminationTime;
        private static ServerSession serverSession;
        private static Logs logger
        {
            get
            {
                return new Logs();
            }
        }
        
        private static string pid
        {
            get
            {
                return Process.GetCurrentProcess().Id.ToString();
            }
        }

        private static Metadata meta
        {
            get
            {
                return new Metadata{
                    {"pid", pid},
                    {"requestId", Guid.NewGuid().ToString("n")},
                };
            }
        }

        private static bool healthStatus = true;
        public static bool HealthStatus
        {
            get
            {
                return healthStatus;
            }
            set
            {
                healthStatus = value;
            }
        }
        
        public static void SetTerminationTime(long stime)
        {
            terminationTime = stime;
        }

        public static void SetServerSession(ServerSession session)
        {
            serverSession = session;
        }

        public static AuxProxyResponse ProcessReady(string[] logPath, int clientPort, int grpcPort)
        {
            logger.Println($"Getting process ready, LogPath: {logPath}, ClientPort: {clientPort}, GrpcPort: {grpcPort}");
            var req = new ProcessReadyRequest{
                ClientPort = clientPort,
                GrpcPort = grpcPort,
		Pid = Process.GetCurrentProcess().Id,
            };
            req.LogPathsToUpload.Add(logPath);         //repeated类型解析pb后，是只读类型，需要Add加入           
            return GrpcClient.ScaseClient.ProcessReady(req, meta);
        }

        public static AuxProxyResponse ActivateServerSession(string serverSessionId, int maxClients)
        {
            logger.Println($"Activating game server session, ServerSessionId: {serverSessionId}, MaxClients: {maxClients}");
            var req = new ActivateServerSessionRequest{
                ServerSessionId = serverSessionId,
                MaxClients = maxClients,
            };  
            return GrpcClient.ScaseClient.ActivateServerSession(req, meta);
        }

        public static AuxProxyResponse AcceptClientSession(string ClientSessionId)
        {
            logger.Println($"Accepting Client session, ClientSessionId: {ClientSessionId}");
            var req = new AcceptClientSessionRequest{
                ServerSessionId = serverSession.ServerSessionId,
                ClientSessionId = ClientSessionId,
            };            
            return GrpcClient.ScaseClient.AcceptClientSession(req, meta);
        }

        public static AuxProxyResponse RemoveClientSession(string ClientSessionId)
        {
            logger.Println($"Removing Client session, ClientSessionId: {ClientSessionId}");
            var req = new RemoveClientSessionRequest{
                ServerSessionId = serverSession.ServerSessionId,
                ClientSessionId = ClientSessionId,
            };            
            return GrpcClient.ScaseClient.RemoveClientSession(req, meta);
        }

        public static AuxProxyResponse TerminateServerSession()
        {
            logger.Println($"Terminating game server session, ServerSessionId: {serverSession.ServerSessionId}");
            var req = new TerminateServerSessionRequest{
                ServerSessionId = serverSession.ServerSessionId
            };            
            return GrpcClient.ScaseClient.TerminateServerSession(req, meta);
        }

        public static AuxProxyResponse ProcessEnding()
        {
            logger.Println($"Process ending, pid: {pid}");
            var req = new ProcessEndingRequest();            
            return GrpcClient.ScaseClient.ProcessEnding(req, meta);
        }

        public static DescribeClientSessionsResponse DescribeClientSessions(string ServerSessionId, string ClientId, string ClientSessionId, string ClientSessionStatusFilter, string nextToken, int limit)
        {
            logger.Println($"Describing Client session, ServerSessionId: {ServerSessionId}, ClientId: {ClientId}, ClientSessionId: {ClientSessionId}, ClientSessionStatusFilter: {ClientSessionStatusFilter}, NextToken: {nextToken}, Limit: {limit}");
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

        public static AuxProxyResponse UpdateClientSessionCreationPolicy(string newPolicy)
        {
            logger.Println($"Updating Client session creation policy, newPolicy: {newPolicy}");
            var req = new UpdateClientSessionCreationPolicyRequest{
                ServerSessionId = serverSession.ServerSessionId,
		        NewClientSessionCreationPolicy = newPolicy,
            };            
            return GrpcClient.ScaseClient.UpdateClientSessionCreationPolicy(req, meta);
        }
    }
}
