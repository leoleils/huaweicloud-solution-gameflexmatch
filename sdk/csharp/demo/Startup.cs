using System;
using Grpc.Core;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using CSharpDemo.Api;
using CSharpDemo.Models;
using AuxproxyService;
using ProcessService;

namespace CSharpDemo
{
    public class Startup
    {
        public Logs logger
        {
            get
            {
                return new Logs();
            }
        }

        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            // services.AddMvc();
            services.AddGrpc();
            // 将Controller的核心服务注册到容器中
            services.AddControllers();

             // 进程启动后，通知agent ProcessReady
            AuxproxyService.AuxProxyResponse resp = new AuxproxyService.AuxProxyResponse();
            try
            {
                resp = GseManager.ProcessReady(Logs.LogPaths, Program.ClientPort, Program.GrpcPort);
            } 
            catch (RpcException ex)
            {
                logger.Println($"Process ready failed, error: {ex.Status.Detail}");
                return;
            }
            logger.Println($"Process ready succeed, resp: {resp}");
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            // 重定向
            //app.UseHttpsRedirection();

            // 将EndpointRoutingMiddleware中间件注册到http管道
            app.UseRouting();

            // 将AuthorizationMiddleware中间件注册到http管道
            app.UseAuthorization();

            // 将EndpointMiddleware中间件注册到http管道
            app.UseEndpoints(endpoints =>
            {
                endpoints.MapGrpcService<ServerSdk>();
                // 将所有Controller和Action转换为EndPoint放到路由中间件的配置对象RouteOptions中
                endpoints.MapControllers();
            });
        }
    }
}
