package gsemanager

import (
	"context"
	"fake-server/grpcsdk"
	"fake-server/logger"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"strconv"
	"sync"
)

var (
	gseManagerIns *gsemanager
	once          sync.Once
)

const (
	localhost = "127.0.0.1"
	//agentPort = 10001
	agentPort = 60002
)

type gsemanager struct {
	pid                 string
	gameServerSession   *grpcsdk.ServerSession
	playerSessionIDList []string
	terminationTime     int64
	rpcClient           grpcsdk.ScaseGrpcSdkServiceClient
}

func GetGseManagerByPid(pid int) *gsemanager {
	once.Do(func() {
		gseManagerIns = &gsemanager{
			pid: strconv.Itoa(pid),
		}

		url := fmt.Sprintf("%s:%d", localhost, agentPort)

		conn, err := grpc.DialContext(context.Background(), url, grpc.WithInsecure())
		if err != nil {
			logger.Fatal("dail to gse fail", zap.String("url", url), zap.Error(err))
		}

		gseManagerIns.rpcClient = grpcsdk.NewScaseGrpcSdkServiceClient(conn)
	})

	return gseManagerIns
}

func GetGseManager() *gsemanager {
	once.Do(func() {
		gseManagerIns = &gsemanager{
			pid:                 strconv.Itoa(os.Getpid()),
			playerSessionIDList: make([]string, 0),
		}

		url := fmt.Sprintf("%s:%d", localhost, agentPort)

		conn, err := grpc.DialContext(context.Background(), url, grpc.WithInsecure())
		if err != nil {
			logger.Fatal("dail to gse fail", zap.String("url", url), zap.Error(err))
		}

		gseManagerIns.rpcClient = grpcsdk.NewScaseGrpcSdkServiceClient(conn)
	})

	return gseManagerIns
}

func (g *gsemanager) SetGameServerSession(gameserversession *grpcsdk.ServerSession) {
	g.gameServerSession = gameserversession
}

func (g *gsemanager) AddPlayerSessionID(playerSessionID string) {
	g.playerSessionIDList = append(g.playerSessionIDList, playerSessionID)
}

func (g *gsemanager) RemovePlayerSessionID(playerSessionID string) {
	for i, id := range g.playerSessionIDList {
		if id == playerSessionID {
			g.playerSessionIDList = append(g.playerSessionIDList[:i], g.playerSessionIDList[i+1:]...)
			return
		}
	}
}

func (g *gsemanager) RemoveAllPlayerSession() {
	fmt.Printf("player sessionID: %+v", g.playerSessionIDList)
	for _, id := range g.playerSessionIDList {
		_, err := g.RemovePlayerSession(id)
		if err != nil {
			logger.Info("RemoveAllPlayerSession failed for ", zap.String("palyersesionid", id), zap.Error(err))
		}
	}
}

func (g *gsemanager) SetTerminationTime(terminationTime int64) {
	g.terminationTime = terminationTime
}

func (g *gsemanager) getContext() context.Context {
	requestId := uuid.NewV4().String()
	ctx := metadata.AppendToOutgoingContext(context.Background(), "pid", g.pid)
	return metadata.AppendToOutgoingContext(ctx, "requestId", requestId)
}

// 1. ProcessReady
func (g *gsemanager) ProcessReady(logPath []string, clientPort int32, grpcPort int32) error {
	logger.Info("start to processready", zap.Any("logPath", logPath), zap.Int32("clientPort", clientPort),
		zap.Int32("grpcPort", grpcPort))
	pid, _ := strconv.ParseInt(g.pid, 10, 32)
	req := &grpcsdk.ProcessReadyRequest{
		LogPathsToUpload: logPath,
		ClientPort:       clientPort,
		GrpcPort:         grpcPort,
		Pid:              int32(pid),
	}

	_, err := g.rpcClient.ProcessReady(g.getContext(), req)
	if err != nil {
		logger.Info("ProcessReady fail", zap.Error(err))
		return err
	}

	logger.Info("ProcessReady success")
	return nil
}

// 2. ActivateGameServerSession
func (g *gsemanager) ActivateGameServerSession(gameServerSessionId string, maxPlayers int32) error {
	logger.Info("start to ActivateGameServerSession", zap.String("gameServerSessionId", gameServerSessionId),
		zap.Int32("maxPlayers", maxPlayers))
	req := &grpcsdk.ActivateServerSessionRequest{
		ServerSessionId: gameServerSessionId,
		MaxClients:      maxPlayers,
	}

	_, err := g.rpcClient.ActivateServerSession(g.getContext(), req)
	if err != nil {
		logger.Error("ActivateGameServerSession fail", zap.Error(err))
		return err
	}

	logger.Info("ActivateGameServerSession success")
	return nil
}

// 3. AcceptPlayerSession
func (g *gsemanager) AcceptPlayerSession(playerSessionId string) (*grpcsdk.AuxProxyResponse, error) {

	logger.Info("start to AcceptPlayerSession", zap.String("playerSessionId", playerSessionId))
	req := &grpcsdk.AcceptClientSessionRequest{
		ServerSessionId: g.gameServerSession.ServerSessionId,
		ClientSessionId: playerSessionId,
	}

	return g.rpcClient.AcceptClientSession(g.getContext(), req)
}

// 4. RemovePlayerSession
func (g *gsemanager) RemovePlayerSession(playerSessionId string) (*grpcsdk.AuxProxyResponse, error) {
	logger.Info("start to RemovePlayerSession", zap.String("playerSessionId", playerSessionId))
	req := &grpcsdk.RemoveClientSessionRequest{
		ServerSessionId: g.gameServerSession.GetServerSessionId(),
		ClientSessionId: playerSessionId,
	}

	return g.rpcClient.RemoveClientSession(g.getContext(), req)
}

// 5. TerminateGameServerSession
func (g *gsemanager) TerminateGameServerSession() (*grpcsdk.AuxProxyResponse, error) {
	if g.gameServerSession == nil || g.gameServerSession.ServerSessionId == "" {
		logger.Info("gameServerSession is nil or server session id is empty, skip TerminateGameServerSession")
		return nil, nil
	}
	logger.Info("start to TerminateGameServerSession", zap.String("serverSessionID", g.gameServerSession.ServerSessionId))
	req := &grpcsdk.TerminateServerSessionRequest{
		ServerSessionId: g.gameServerSession.ServerSessionId,
	}

	return g.rpcClient.TerminateServerSession(g.getContext(), req)
}

// 6. ProcessEnding
func (g *gsemanager) ProcessEnding() (*grpcsdk.AuxProxyResponse, error) {
	logger.Info("start to ProcessEnding")
	pid, _ := strconv.ParseInt(g.pid, 10, 32)
	req := &grpcsdk.ProcessEndingRequest{
		Pid: int32(pid),
	}

	return g.rpcClient.ProcessEnding(g.getContext(), req)
}

// 7. DescribePlayerSessions
func (g *gsemanager) DescribePlayerSessions(gameServerSessionId, playerId, playerSessionId, playerSessionStatusFilter, nextToken string,
	limit int32) (*grpcsdk.DescribeClientSessionsResponse, error) {
	logger.Info("start to DescribePlayerSessions", zap.String("gameServerSessionId", gameServerSessionId),
		zap.String("playerId", playerId), zap.String("playerSessionId", playerSessionId),
		zap.String("playerSessionStatusFilter", playerSessionStatusFilter), zap.String("nextToken", nextToken),
		zap.Int32("limit", limit))

	req := &grpcsdk.DescribeClientSessionsRequest{
		ServerSessionId:           gameServerSessionId,
		ClientId:                  playerId,
		ClientSessionId:           playerSessionId,
		ClientSessionStatusFilter: playerSessionStatusFilter,
		NextToken:                 nextToken,
		Limit:                     limit,
	}

	return g.rpcClient.DescribeClientSessions(g.getContext(), req)
}

// 8. UpdatePlayerSessionCreationPolicy
func (g *gsemanager) UpdatePlayerSessionCreationPolicy(newPolicy string) (*grpcsdk.AuxProxyResponse, error) {
	logger.Info("start to UpdatePlayerSessionCreationPolicy", zap.String("newPolicy", newPolicy))
	req := &grpcsdk.UpdateClientSessionCreationPolicyRequest{
		ServerSessionId:                g.gameServerSession.ServerSessionId,
		NewClientSessionCreationPolicy: newPolicy,
	}

	return g.rpcClient.UpdateClientSessionCreationPolicy(g.getContext(), req)
}

// 9.ReportCustomData
func (g *gsemanager) ReportCustomData(currentCustomCount, maxCustomCount int32) (*grpcsdk.AuxProxyResponse, error) {
	logger.Info("start to UpdatePlayerSessionCreationPolicy", zap.Int32("currentCustomCount", currentCustomCount),
		zap.Int32("maxCustomCount", maxCustomCount))

	return &grpcsdk.AuxProxyResponse{}, nil
}

func (g *gsemanager) DescribeGameServerSession() string {
	return fmt.Sprintf("game server session id %s", g.gameServerSession.ServerSessionId)
}
