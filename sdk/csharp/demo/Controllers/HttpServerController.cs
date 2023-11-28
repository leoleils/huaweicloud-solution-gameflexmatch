using System;
using Grpc.Core;
using Newtonsoft.Json;
using Microsoft.AspNetCore.Mvc;
using CSharpDemo.Models;
using AuxproxyService;
using ProcessService;

/* Http请求的接口，供游戏服务端测试调用*/

namespace CSharpDemo.Controllers
{
    [Route("/gse/")]
    [Produces("application/json")]
    public class HttpServerController : ControllerBase
    {
        private Logs logger
        {
            get
            {
                return new Logs();
            }
        }

        private string getOkDefaultResponse()
        {
            return JsonConvert.SerializeObject(new GseHttpResponse());
        }
        private string getclientSessionResponse(AuxproxyService.DescribeClientSessionsResponse clientSessions)
        {
            return JsonConvert.SerializeObject(new GseHttpResponse(0, "SUCCESS", clientSessions));
        }

        private string getBadResponse(string msg)
        {
            return JsonConvert.SerializeObject(new GseHttpResponse(400, msg, null));
        }

        // 玩家登录
        [Route("login")]
        public IActionResult Login(string clientSessionId)
        {
            if(clientSessionId == null || clientSessionId == "")
            {
                return BadRequest(getBadResponse($"clientSessionId can't be empty"));
            }
            logger.Println($"client {clientSessionId} Login");

            AuxProxyResponse resp = new AuxProxyResponse();
            try
            {
                resp = GseManager.AcceptClientSession(clientSessionId);
            } 
            catch (RpcException ex)
            {
                string errMsg = $"AcceptclientSession failed, error: {ex.Status.Detail}";
                logger.Println(errMsg);
                return BadRequest(getBadResponse(errMsg));
            }
            logger.Println($"AcceptclientSession succeed, resp: {resp}");
            return Ok(getOkDefaultResponse());
        }

        // 玩家退出
        [Route("logout")]
        public IActionResult LoginOut(string clientSessionId)
        {
            if(clientSessionId == null || clientSessionId == "")
            {
                return BadRequest(getBadResponse($"clientSessionId can't be empty"));
            }
            logger.Println($"client {clientSessionId} LoginOut");
            
            AuxProxyResponse resp = new AuxProxyResponse();
            try
            {
                resp = GseManager.RemoveClientSession(clientSessionId);
            }
            catch (RpcException ex)
            {
                string errMsg = $"RemoveclientSession failed, error: {ex.Status.Detail}";
                logger.Println(errMsg);
                return BadRequest(getBadResponse(errMsg));
            }

            logger.Println($"RemoveclientSession succeed, resp: {resp}");
            return Ok(getOkDefaultResponse());
        }

        // 结束游戏
        [Route("terminate-game-server-session")]
        public IActionResult TerminateSession()
        {
            logger.Println($"Start to TerminateSession");

            AuxProxyResponse resp = new AuxProxyResponse();
            try
            {
                resp = GseManager.TerminateServerSession();
            }
            catch (RpcException ex)
            {
                string errMsg = $"TerminateSession failed, error: {ex.Status.Detail}";
                logger.Println(errMsg);
                return BadRequest(getBadResponse(errMsg));
            }

            logger.Println($"TerminateSession succeed, resp: {resp}");
            return Ok(getOkDefaultResponse());
        }

        // 结束进程
        [Route("end-process")]
        public IActionResult EndProcess()
        {
            logger.Println($"Start to TerminateSession");

            AuxProxyResponse resp = new AuxProxyResponse();
            try
            {
                resp = GseManager.ProcessEnding();
            } 
            catch (RpcException ex)
            {
                string errMsg = $"ProcessEnding failed, error: {ex.Status.Detail}";
                logger.Println(errMsg);
                return BadRequest(getBadResponse(errMsg));
            }

            logger.Println($"ProcessEnding succeed, resp: {resp}");
            return Ok(getOkDefaultResponse());
        } 

        // 查询玩家会话列表
        [Route("describe-client-sessions")]
        public IActionResult DescribeclientSessions(string gameServerSessionId, string clientId, string clientSessionId, string clientSessionStatusFilter, string nextToken, string limit)
        {
            logger.Println($"Start to DescribeclientSessions, GameServerSessionId: {gameServerSessionId}, clientId: {clientId}, clientSessionId: {clientSessionId}, clientSessionStatusFilter: {clientSessionStatusFilter}, NextToken: {nextToken}, Limit: {limit}");
            int limitInt = 10;
            if(limit != null)
            {
                limitInt = Int32.Parse(limit);
            }
            var resp = new DescribeClientSessionsResponse();
            try
            {
                resp = GseManager.DescribeClientSessions(gameServerSessionId, clientId, clientSessionId, 
                clientSessionStatusFilter, nextToken, limitInt);
            }
            catch (RpcException ex)
            {
                string errMsg = $"DescribeclientSessions failed, error: {ex.Status.Detail}";
                logger.Println(errMsg);
                return BadRequest(getBadResponse(errMsg));
            } 
            logger.Println($"DescribeclientSessions succeed, resp: {resp}");
            return Ok(getclientSessionResponse(resp));
        }

        // 修改玩家会话创建策略
        [Route("update-client-session-policy")]
        public IActionResult UpdateclientSessionCreationPolicy(string newclientSessionCreationPolicy)
        {
            logger.Println($"Start to UpdateclientSessionCreationPolicy, NewclientSessionCreationPolicy: {newclientSessionCreationPolicy}");
            
            AuxProxyResponse resp = new AuxProxyResponse();
            try
            {
                resp = GseManager.UpdateClientSessionCreationPolicy(newclientSessionCreationPolicy);
            }
            catch (RpcException ex)
            {
                string errMsg = $"UpdateclientSessionCreationPolicy failed, error: {ex.Status.Detail}";
                logger.Println(errMsg);
                return BadRequest(getBadResponse(errMsg));
            } 

            logger.Println($"UpdateclientSessionCreationPolicy succeed, resp: {resp}");
            return Ok(getOkDefaultResponse());
        }

        // 设置进程健康状态
        [Route("set-process-health-status")]
        public IActionResult SetHealthStatus(string healthStatus)
        {
            logger.Println($"Start to SetHealthStatus, Status: {healthStatus}");
            int hStatus = 0;
            if(healthStatus != null)
            {
                hStatus = Int32.Parse(healthStatus);
            }
            if(hStatus == 0) 
            {
		        GseManager.HealthStatus = false;
	        } 
            else 
            {
		        GseManager.HealthStatus = true;
	        }
            logger.Println($"SetHealthStatus succeed, HealthStatus: {GseManager.HealthStatus}");
            return Ok(getOkDefaultResponse());
        }
    }
}

