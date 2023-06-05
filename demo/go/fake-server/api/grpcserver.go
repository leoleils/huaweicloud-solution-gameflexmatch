package api

import (
	"context"
	"fake-server/config"
	"fake-server/grpcsdk"
	"fake-server/gsemanager"
	"fake-server/logger"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

// func GetRpcService() grpcsdk.GameServerGrpcSdkServiceServer {
func GetRpcService() *rpcService {
	once.Do(func() {
		rpcServerIns = new(rpcService)
		rpcServerIns.healthStatus = true
	})

	return rpcServerIns
}

func (s *rpcService) StartGrpcServer() error {
	listen, err := net.Listen("tcp", "localhost:")
	if err != nil {
		logger.Logger.Errorf("grpc fail to listen", zap.Error(err))
		return fmt.Errorf("grpc fail to listen")
	}

	addr := listen.Addr().String()
	portStr := strings.Split(addr, ":")[1]
	s.grpcPort, err = strconv.Atoi(portStr)
	if err != nil {
		logger.Logger.Errorf("grpc fail to get port", zap.Error(err))
		return fmt.Errorf("grpc fail to get port")
	}

	logger.Logger.Infof("grpc listen port is", zap.Int("port", s.grpcPort))

	grpcServer := grpc.NewServer()
	grpcsdk.RegisterProcessGrpcSdkServiceServer(grpcServer, s)

	reflection.Register(grpcServer)

	logger.Logger.Infof("start grpc server success")
	go grpcServer.Serve(listen)
	return nil
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

	logger.Logger.Infof("OnHealthCheck status: " + strconv.FormatBool(s.healthStatus))
	return resp, nil
}

func (s *rpcService) OnStartServerSession(ctx context.Context, req *grpcsdk.StartServerSessionRequest) (*grpcsdk.ProcessResponse,
	error) {
	logger.Logger.Infof("OnStartGameServerSession called, req:" + req.String())

	gseManager := gsemanager.GetGseManager()
	err := gseManager.SetGameServerSession(req.ServerSession)
	if err != nil {
		return nil, err
	}
	err = gseManager.ActivateGameServerSession(req.ServerSession.ServerSessionId, req.GetServerSession().MaxClients)
	if err != nil {
		logger.Logger.Errorf("activate game session err: %+v", err)
		return nil, err
	}
	go gseManager.HandleGameSession(req.ServerSession.ServerSessionId)
	resp := new(grpcsdk.ProcessResponse)

	return resp, nil
}

func (s *rpcService) OnProcessTerminate(ctx context.Context, req *grpcsdk.ProcessTerminateRequest) (*grpcsdk.ProcessResponse, error) {
	logger.Logger.Infof("OnProcessTerminate called, req:" + req.String())
	gseManager := gsemanager.GetGseManager()
	gseManager.HandingTerminatingProcess(req.TerminationTime)
	resp := new(grpcsdk.ProcessResponse)
	return resp, nil
}

func (s *rpcService) HandleReadyProcess() {
	processRunMinute := config.GlobalConfig.ProcessRunMinute 
	runSeconds := processRunMinute * 60 + rand.Intn(100) - 50
	logger.Logger.Infof("process will run %d seconds", runSeconds)
	timechannel := time.After(time.Duration(runSeconds) * time.Second)
	select {
	case <-timechannel:
		logger.Logger.Infof("start to terminated becase run %d seconds", runSeconds)
		gseManager := gsemanager.GetGseManager()
		// 把状态置为不健康，不接受新的会话
		s.SetHealthStatus(false)
		// 休眠一段时间，保证该进程上不会再被分配会话后再进行会话的接管
		time.Sleep(35 * time.Second)
		logger.Logger.Infof("start to terminated becase run %d seconds after sleep 35 seconds", runSeconds)
		go gseManager.HandingTerminatingProcess(30)
	}
}