using System;
using Grpc.Core;
using System.Threading.Tasks;
using ProcessService;
using AuxproxyService;
using CSharpDemo.Models;

/* gRPC接口，供Agent调用 */

namespace CSharpDemo.Api
{
    public class ServerSdk : ProcessGrpcSdkService.ProcessGrpcSdkServiceBase
    {
        private static Logs logger
        {
            get
            {
                return new Logs();
            }
        }

        public override Task<HealthCheckResponse> OnHealthCheck(HealthCheckRequest request, ServerCallContext context)
        {
            logger.Println($"OnHealthCheck, HealthStatus: {GseManager.HealthStatus}");
            return Task.FromResult(new HealthCheckResponse{
                HealthStatus = GseManager.HealthStatus
            });
        }

        public override Task<ProcessResponse> OnStartServerSession(StartServerSessionRequest request, ServerCallContext context)
        {
            logger.Println($"OnStartServerSession, request: {request}");
            GseManager.SetServerSession(request.ServerSession);
            var resp = GseManager.ActivateServerSession(request.ServerSession.ServerSessionId, request.ServerSession.MaxClients);
            // return Task.FromResult(resp);
            return Task.FromResult(new ProcessResponse());
        }

        public override Task<ProcessResponse> OnProcessTerminate(ProcessTerminateRequest request, ServerCallContext context)
        {
            logger.Println($"OnProcessTerminate, request: {request}");
            // 设置进程终止时间
            GseManager.SetTerminationTime(request.TerminationTime);
            // 终止游戏服务器会话
            GseManager.TerminateServerSession();
            // 进程退出
            GseManager.ProcessEnding();
            return Task.FromResult(new ProcessResponse());
        }
    }
}