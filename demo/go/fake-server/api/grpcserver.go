package api

import (
	"context"
	"fake-server/grpcsdk"
	"fake-server/gsemanager"
	"fake-server/logger"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
	"sync"
)

var (
	rpcServerIns *rpcService
	once         sync.Once
)

type rpcService struct {
	healthStatus bool
	grpcPort     int
	grpcsdk.UnsafeProcessGrpcSdkServiceServer
}

//func GetRpcService() grpcsdk.GameServerGrpcSdkServiceServer {
func GetRpcService() *rpcService {
	once.Do(func() {
		rpcServerIns = new(rpcService)
		rpcServerIns.healthStatus = true
	})

	return rpcServerIns
}

func (s *rpcService) StartGrpcServer() {
	listen, err := net.Listen("tcp", "localhost:")
	if err != nil {
		logger.Fatal("grpc fail to listen", zap.Error(err))
	}

	addr := listen.Addr().String()
	portStr := strings.Split(addr, ":")[1]
	s.grpcPort, err = strconv.Atoi(portStr)
	if err != nil {
		logger.Fatal("grpc fail to get port", zap.Error(err))
	}

	logger.Info("grpc listen port is", zap.Int("port", s.grpcPort))

	grpcServer := grpc.NewServer()
	grpcsdk.RegisterProcessGrpcSdkServiceServer(grpcServer, s)
	logger.Info("start grpc server success")
	go grpcServer.Serve(listen)
}

func (s *rpcService) GetGrpcPort() int {
	return s.grpcPort
}

func (s *rpcService) SetHealthStatus(healthStatus bool) {
	s.healthStatus = healthStatus
}

func (s *rpcService) OnHealthCheck(ctx context.Context, req *grpcsdk.HealthCheckRequest) (*grpcsdk.HealthCheckResponse, error) {
	resp := &grpcsdk.HealthCheckResponse{
		HealthStatus: s.healthStatus,
	}

	logger.Info("OnHealthCheck status: " + strconv.FormatBool(s.healthStatus))
	return resp, nil
}

func (s *rpcService) OnStartServerSession(ctx context.Context, req *grpcsdk.StartServerSessionRequest) (*grpcsdk.ProcessResponse,
	error) {
	logger.Info("OnStartGameServerSession called, req:" + req.String())

	gseManager := gsemanager.GetGseManager()
	gseManager.SetGameServerSession(req.ServerSession)
	gseManager.ActivateGameServerSession(req.ServerSession.ServerSessionId, req.GetServerSession().MaxClients)

	resp := new(grpcsdk.ProcessResponse)

	return resp, nil
}

func (s *rpcService) OnProcessTerminate(ctx context.Context, req *grpcsdk.ProcessTerminateRequest) (*grpcsdk.ProcessResponse, error) {
	logger.Info("OnProcessTerminate called, req:" + req.String())

	gseManager := gsemanager.GetGseManager()
	gseManager.SetTerminationTime(req.TerminationTime)
	// 结束client session
	fmt.Println("start to remove all player session")
	logger.Info("start to remove all player session")
	gseManager.RemoveAllPlayerSession()

	//结束游戏会话
	fmt.Println("start to terminate game session")
	logger.Info("start to terminate game session")
	gseManager.TerminateGameServerSession()

	// 进程退出
	fmt.Println("start to end process")
	logger.Info("start to end process")
	gseManager.ProcessEnding()

	resp := new(grpcsdk.ProcessResponse)
	return resp, nil
}
