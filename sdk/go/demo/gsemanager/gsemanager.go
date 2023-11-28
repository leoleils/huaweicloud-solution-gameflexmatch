package gsemanager

import (
	"context"
	"fleetmanager_build/gfm1125/huaweicloud-solution-metaspace/sdk/go/demo/config"
	"fleetmanager_build/gfm1125/huaweicloud-solution-metaspace/sdk/go/demo/grpcsdk"
	"fleetmanager_build/gfm1125/huaweicloud-solution-metaspace/sdk/go/demo/logger"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	pid            string
	gameSessionMux sync.Mutex
	gameSessions   map[string]*grpcsdk.ServerSession

	playSessionMux  sync.Mutex
	playerSessions  map[string][]string
	terminationTime int64
	rpcClient       grpcsdk.ScaseGrpcSdkServiceClient
}

func GetGseManagerByPid(pid int) *gsemanager {
	once.Do(func() {
		gseManagerIns = &gsemanager{
			pid: strconv.Itoa(pid),
		}

		url := fmt.Sprintf("%s:%d", localhost, agentPort)

		conn, err := grpc.DialContext(context.Background(), url, grpc.WithInsecure())
		if err != nil {
			logger.Logger.Errorf("dail to gse fail", zap.String("url", url), zap.Error(err))
		}

		gseManagerIns.rpcClient = grpcsdk.NewScaseGrpcSdkServiceClient(conn)
	})

	return gseManagerIns
}

func GetGseManager() *gsemanager {
	once.Do(func() {
		gseManagerIns = &gsemanager{
			pid:            strconv.Itoa(os.Getpid()),
			playerSessions: make(map[string][]string, 0),
			gameSessions:   make(map[string]*grpcsdk.ServerSession, 0),
		}

		url := fmt.Sprintf("%s:%d", localhost, agentPort)

		conn, err := grpc.DialContext(context.Background(), url, grpc.WithInsecure())
		if err != nil {
			logger.Logger.Errorf("dail to gse fail", zap.String("url", url), zap.Error(err))
		}

		gseManagerIns.rpcClient = grpcsdk.NewScaseGrpcSdkServiceClient(conn)
	})

	return gseManagerIns
}

// 设置会话
func (g *gsemanager) SetGameServerSession(gameserversession *grpcsdk.ServerSession) error {
	g.gameSessionMux.Lock()
	defer g.gameSessionMux.Unlock()
	if len(g.gameSessions) > config.GlobalConfig.MaxGameSessionCount {
		return fmt.Errorf("set game session %s err because game session count cant lagger than %d, now: %d",
			gameserversession.ServerSessionId, config.GlobalConfig.MaxGameSessionCount, len(g.gameSessions))
	}
	g.gameSessions[gameserversession.ServerSessionId] = gameserversession
	return nil
}

func (g *gsemanager) RemoveAllPlayerSession(gameSessionId string) {
	g.playSessionMux.Lock()
	if _, ok := g.playerSessions[gameSessionId]; !ok {
		logger.Logger.Infof("game session %s have no play sessions", gameSessionId)
		g.playSessionMux.Unlock()
		return
	}
	playSessionIds := g.playerSessions[gameSessionId]
	g.playSessionMux.Unlock()
	for _, id := range playSessionIds {
		_, err := g.RemovePlayerSession(id, gameSessionId)
		if err != nil {
			logger.Logger.Infof("RemoveAllPlayerSession failed for ", zap.String("palyersesionid", id), zap.Error(err))
		}
	}
}

// 启动一个定时器，纳管服务端会话，一段时间后自动结束
func (g *gsemanager) HandleGameSession(gameSessionId string) {
	gameSessionRetainMinute := config.GlobalConfig.GameSessionRetainMinute
	logger.Logger.Infof("[clean game session] the game session will retain %d minute", gameSessionRetainMinute)
	timeChannel := time.After(time.Duration(gameSessionRetainMinute) * time.Minute)
	select {
	case <-timeChannel:
		logger.Logger.Infof("terminate game session %s after %d minutes", gameSessionId, gameSessionRetainMinute)
		resp, err := g.TerminateGameServerSession(gameSessionId)
		if err != nil {
			logger.Logger.Errorf("terminate game session %s err %+v, resp: %+v", gameSessionId, err, resp)
			return
		}
		logger.Logger.Infof("success to terminate game session %s", gameSessionId)
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
	logger.Logger.Infof("start to processready", zap.Any("logPath", logPath), zap.Int32("clientPort", clientPort),
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
		logger.Logger.Infof("ProcessReady fail", zap.Error(err))
		return err
	}

	logger.Logger.Infof("ProcessReady success")
	return nil
}

// 2. ActivateGameServerSession
func (g *gsemanager) ActivateGameServerSession(gameServerSessionId string, maxPlayers int32) error {
	logger.Logger.Infof("start to ActivateGameServerSession", zap.String("gameServerSessionId", gameServerSessionId),
		zap.Int32("maxPlayers", maxPlayers))
	req := &grpcsdk.ActivateServerSessionRequest{
		ServerSessionId: gameServerSessionId,
		MaxClients:      maxPlayers,
	}

	_, err := g.rpcClient.ActivateServerSession(g.getContext(), req)
	if err != nil {
		logger.Logger.Errorf("ActivateGameServerSession fail", zap.Error(err))
		return err
	}

	logger.Logger.Infof("ActivateGameServerSession success")
	return nil
}

// 3. AcceptPlayerSession
func (g *gsemanager) AcceptPlayerSession(playerSessionId string, gameSessionId string) (*grpcsdk.AuxProxyResponse, error) {
	logger.Logger.Infof("start to AcceptPlayerSession", zap.String("playerSessionId", playerSessionId), zap.String("gameSessionId", gameSessionId))
	g.gameSessionMux.Lock()
	if _, ok := g.gameSessions[gameSessionId]; !ok {
		logger.Logger.Infof("game session %s not found", gameSessionId)
		return nil, fmt.Errorf("game session %s not found", gameSessionId)
	}
	g.gameSessionMux.Unlock()
	g.playSessionMux.Lock()
	req := &grpcsdk.AcceptClientSessionRequest{
		ServerSessionId: gameSessionId,
		ClientSessionId: playerSessionId,
	}
	if _, ok := g.playerSessions[gameSessionId]; !ok {
		playSessiondIds := make([]string, 0)
		playSessiondIds = append(playSessiondIds, playerSessionId)
		g.playerSessions[gameSessionId] = playSessiondIds
	} else {
		playsessiondIds := g.playerSessions[gameSessionId]
		playsessiondIds = append(playsessiondIds, playerSessionId)
		g.playerSessions[gameSessionId] = playsessiondIds
	}
	g.playSessionMux.Unlock()

	return g.rpcClient.AcceptClientSession(g.getContext(), req)
}

// 4. RemovePlayerSession
func (g *gsemanager) RemovePlayerSession(playerSessionId string, gameSessionId string) (*grpcsdk.AuxProxyResponse, error) {
	g.playSessionMux.Lock()
	idx := -1
	if _, ok := g.playerSessions[gameSessionId]; !ok {
		return nil, fmt.Errorf("game session %s not found", gameSessionId)
	}
	for i, pyid := range g.playerSessions[gameSessionId] {
		if pyid == playerSessionId {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, fmt.Errorf("play session %s not found", playerSessionId)
	}
	g.playerSessions[gameSessionId] = append(g.playerSessions[gameSessionId][:idx], g.playerSessions[gameSessionId][idx+1:]...)
	g.playSessionMux.Unlock()
	logger.Logger.Infof("Remove play session %s success on game session %s", playerSessionId, gameSessionId)
	req := &grpcsdk.RemoveClientSessionRequest{
		ServerSessionId: gameSessionId,
		ClientSessionId: playerSessionId,
	}

	return g.rpcClient.RemoveClientSession(g.getContext(), req)
}

func (g *gsemanager) TerminateAllGameServerSession() error {
	g.gameSessionMux.Lock()
	gameSessions := g.gameSessions
	g.gameSessionMux.Unlock()
	logger.Logger.Infof("start terminate all game session")
	for key, _ := range gameSessions {
		resp, err := g.TerminateGameServerSession(key)
		if err != nil {
			logger.Logger.Errorf("terminated game session %s failed, err %+v, resp: %+v", key, err, resp)
			return err
		}
		logger.Logger.Infof("success to terminate game session %s", key)
	}
	return nil
}

// 5. TerminateGameServerSession
func (g *gsemanager) TerminateGameServerSession(gameSessionId string) (*grpcsdk.AuxProxyResponse, error) {
	g.gameSessionMux.Lock()
	if _, ok := g.gameSessions[gameSessionId]; !ok {
		logger.Logger.Infof("gameServerSession is nil or server session id is empty, skip TerminateGameServerSession")
		g.gameSessionMux.Unlock()
		return nil, nil
	}
	logger.Logger.Infof("start to TerminateGameServerSession", zap.String("serverSessionID", gameSessionId))
	delete(g.gameSessions, gameSessionId)
	g.gameSessionMux.Unlock()
	logger.Logger.Infof("start to clear PlayerServerSession", zap.String("serverSessionID", gameSessionId))
	g.RemoveAllPlayerSession(gameSessionId)
	logger.Logger.Infof("success to clear PlayerServerSession", zap.String("serverSessionID", gameSessionId))
	req := &grpcsdk.TerminateServerSessionRequest{
		ServerSessionId: gameSessionId,
	}

	return g.rpcClient.TerminateServerSession(g.getContext(), req)
}

// 6. ProcessEnding
func (g *gsemanager) ProcessEnding() (*grpcsdk.AuxProxyResponse, error) {
	logger.Logger.Infof("start to ProcessEnding")
	pid, _ := strconv.ParseInt(g.pid, 10, 32)
	req := &grpcsdk.ProcessEndingRequest{
		Pid: int32(pid),
	}

	return g.rpcClient.ProcessEnding(g.getContext(), req)
}

// 7. DescribePlayerSessions
func (g *gsemanager) DescribePlayerSessions(gameServerSessionId, playerId, playerSessionId, playerSessionStatusFilter, nextToken string,
	limit int32) (*grpcsdk.DescribeClientSessionsResponse, error) {
	logger.Logger.Infof("start to DescribePlayerSessions", zap.String("gameServerSessionId", gameServerSessionId),
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
func (g *gsemanager) UpdatePlayerSessionCreationPolicy(newPolicy string, gameSessionId string) (*grpcsdk.AuxProxyResponse, error) {
	logger.Logger.Infof("start to UpdatePlayerSessionCreationPolicy", zap.String("newPolicy", newPolicy))
	req := &grpcsdk.UpdateClientSessionCreationPolicyRequest{
		ServerSessionId:                gameSessionId,
		NewClientSessionCreationPolicy: newPolicy,
	}

	return g.rpcClient.UpdateClientSessionCreationPolicy(g.getContext(), req)
}

// 9.ReportCustomData
func (g *gsemanager) ReportCustomData(currentCustomCount, maxCustomCount int32) (*grpcsdk.AuxProxyResponse, error) {
	logger.Logger.Infof("start to UpdatePlayerSessionCreationPolicy", zap.Int32("currentCustomCount", currentCustomCount),
		zap.Int32("maxCustomCount", maxCustomCount))

	return &grpcsdk.AuxProxyResponse{}, nil
}

// 接管待清理的进程
func (g *gsemanager) HandingTerminatingProcess(terminationTime int64) {
	g.SetTerminationTime(terminationTime)
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		g.gameSessionMux.Lock()
		if len(g.gameSessions) > 0 {
			keys := make([]string, 0, len(g.gameSessions))
			for key := range g.gameSessions {
				keys = append(keys, key)
			}
			logger.Logger.Infof("[clean game session] the process still exist game session: %d, game session ids: %+v", len(g.gameSessions), keys)
			g.gameSessionMux.Unlock()
			continue
		}
		g.gameSessionMux.Unlock()
		logger.Logger.Infof("[clean game session] the process have no game session, end the process")
		ticker.Stop()
		_, err := g.ProcessEnding()
		if err != nil {
			logger.Logger.Errorf("[clean game session] process ending failed, %+v", err)
		}
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
}
