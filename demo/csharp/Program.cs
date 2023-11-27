using System;
using System.Net;
using Microsoft.Extensions.Hosting;
using Microsoft.AspNetCore.Hosting;
using Grpc.Core;
using Microsoft.AspNetCore.Server.Kestrel.Core;
using CSharpDemo.Models;

namespace CSharpDemo
{
    public class Program
    {
        public static int ClientPort = PortServer.GenerateRandomPort(2000, 6000);
        public static int GrpcPort = PortServer.GenerateRandomPort(6001, 10000);

        public static void Main(string[] args)
        {
            CreateHostBuilder(args).Build().Run();
        }

        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    webBuilder.ConfigureKestrel(options =>
                    {
                        // gRPC Port (Setup a HTTP/2 endpoint without TLS.)
                        options.ListenAnyIP(GrpcPort, o => o.Protocols = 
                            HttpProtocols.Http2);

                        // HTTP Port
                        options.ListenAnyIP(ClientPort);
                    });

                    webBuilder.UseStartup<Startup>();
                });
    }
}
