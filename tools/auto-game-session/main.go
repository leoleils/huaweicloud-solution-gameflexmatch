package main

import (
	"flag"
	"os"
	"sync"
	"time"
)

func MockGameSession() {
	token := ""
	ticker := time.NewTicker(time.Duration(GlobalConfig.CreateGameSessionSleepSecond) * time.Second)
	for {
		<-ticker.C
		if token == "" {
			token = Login()
		}
		go BatchCreateGameSession(GlobalConfig.BatchCreateGameSession, token)
	}
}

func BatchCreateGameSession(batch int, token string) {
	wg := &sync.WaitGroup{}
	wg.Add(batch)
	for i := 0; i < batch; i++ {
		go CreateAndQueryGameSession(token, wg)
	}
	wg.Wait()
}

func CreateAndQueryGameSession(token string, wg *sync.WaitGroup) {
	sessionId, fleetId, err := CreateGameSession(token)
	if err != nil {
		Logger.Errorf("Create game session err: %+v\n", err)
		wg.Done()
		return
	}
	Logger.Infof("Success to send game session %s", sessionId)
	count := 0
	for {
		sessionState, err := QueryGameSession(sessionId, fleetId, token)
		count += 1
		if err != nil {
			Logger.Errorf("failed to query game session")
			wg.Done()
			return
		}
		if sessionState == GameSessionActiveState {
			Logger.Infof("success to active game session %s after query %d times", sessionId, count)
			wg.Done()
			return
		} else if sessionState == GameSessionErrorState {
			Logger.Infof("failed to active game session %s after query %d times", sessionId, count)
			wg.Done()
			return
		}
	}
}

func main() {
	flag.StringVar(&GlobalConfig.FleetManagerAddr, "fleetmanager-addr", "100.85.113.239", "fleetmanager-addr")
	flag.StringVar(&GlobalConfig.FleetManagerPort, "fleetmanager-port", "31002", "fleetmanager-port")
	flag.StringVar(&GlobalConfig.ProjectId, "project-id", "", "project-id")
	flag.StringVar(&GlobalConfig.AliasId, "alias-id", "", "alias-id")
	flag.StringVar(&GlobalConfig.FleetId, "fleet-id", "", "fleet-id")
	flag.StringVar(&GlobalConfig.LoginName, "username", DefaultLoginName, "uername")
	flag.StringVar(&GlobalConfig.LoginPassword, "password", DefaultLoginPassword, "password")
	flag.IntVar(&GlobalConfig.BatchCreateGameSession, "batch-count", DefaultBatchCreateGameSession, "batch create game session")
	flag.IntVar(&GlobalConfig.CreateGameSessionSleepSecond, "sleep-second", DefaultCreateGameSessionSleepSecond, "sleep second")
	flag.IntVar(&GlobalConfig.ProcessRunSleepSeconds, "process-run-sleep-second", DefaultProcessRunSleepSeconds, "process-run-sleep-second")

	flag.Parse()

	// 日志初始化
	Logger, _ = Init()
	go MockGameSession()
	time.Sleep(time.Duration(GlobalConfig.ProcessRunSleepSeconds) * time.Second)
	Logger.Infof("time to exited")
	os.Exit(1)
}
