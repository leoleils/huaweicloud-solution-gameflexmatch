using System;
using Grpc.Core;
using ProcessService;
using AuxproxyService;

namespace CSharpDemo.Models
{
    public class GrpcClient
    {
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
}
