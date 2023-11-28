// Copyright 2020 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fleetmanager_build/gfm1125/huaweicloud-solution-metaspace/sdk/go/demo/api"
	"fleetmanager_build/gfm1125/huaweicloud-solution-metaspace/sdk/go/demo/config"
	"fleetmanager_build/gfm1125/huaweicloud-solution-metaspace/sdk/go/demo/gsemanager"
	"fleetmanager_build/gfm1125/huaweicloud-solution-metaspace/sdk/go/demo/logger"
	"os"
	"os/signal"
	"syscall"
)

func startGrpcServer() (int, error) {
	// 启动grpc server，监听agent回调
	grpcServer := api.GetRpcService()
	err := grpcServer.StartGrpcServer()
	if err != nil {
		return -1, err
	}
	grpcPort := grpcServer.GetGrpcPort()

	// 返回 grpc port
	return grpcPort, nil
}

// main intercepts the log file of the SuperTuxKart gameserver and uses it
// to determine if the game server is ready or not.
const (
	// DefaultlogPath	= "D://Desktop//workdir//gameflexmatch//gameflexmatch-dev//fake-server//log"
	DefaultlogPath                 = "/local/app/fake-server/log"
	DefaultHttpStartPort           = 1025
	DefaultHttpEndPort             = 60001
	DefaultProcessExitedMinute     = 5 // min
	DefaultProcessRunMinute        = 10
	DefaultGameSessionRetainMinute = 1
	DefaultMaxGameSessionCount     = 50
)

func main() {
	// 启动Grpc Server
	flag.StringVar(&config.GlobalConfig.LogPath, "log-path", DefaultlogPath, "log path")
	flag.StringVar(&config.GlobalConfig.Ak, "ak", "", "log path")
	flag.StringVar(&config.GlobalConfig.Sk, "sk", "", "log path")
	flag.IntVar(&config.GlobalConfig.HttpStartPort, "start-port", DefaultHttpStartPort, "http server start port")
	flag.IntVar(&config.GlobalConfig.HttpEndPort, "end-port", DefaultHttpEndPort, "http server end port")
	flag.IntVar(&config.GlobalConfig.GameSessionRetainMinute, "session-retain-time", DefaultGameSessionRetainMinute, "game session retain minute")
	flag.IntVar(&config.GlobalConfig.ProcessRunMinute, "process-run-time", DefaultProcessRunMinute, "process run time minute")
	flag.IntVar(&config.GlobalConfig.MaxGameSessionCount, "max-game-session-count", DefaultMaxGameSessionCount, "max game session count")
	flag.Parse()

	logger.Logger, _ = logger.Init()
	logger.Logger.Infof("[fake server init]config: %+v", config.GlobalConfig)
	grpcPort, err := startGrpcServer()
	if err != nil {
		logger.Logger.Errorf("start grpc server failed: %+v", err)
		os.Exit(0)
	}
	logger.Logger.Infof("[fake server init] grpc port %d", grpcPort)

	if config.GlobalConfig.HttpEndPort <= config.GlobalConfig.HttpStartPort {
		logger.Logger.Errorf("http end port should lagger then http start port")
		os.Exit(0)
	}
	logger.Logger.Infof("http port will be [%d, %d)", config.GlobalConfig.HttpStartPort, config.GlobalConfig.HttpEndPort)
	httpServer := api.NewHttpProcess()
	err = httpServer.StartHttpServer()
	if err != nil {
		logger.Logger.Errorf("start http server failed: %+v", err)
		os.Exit(0)
	}
	gseManager := gsemanager.GetGseManager()
	gseManager.ProcessReady(nil, int32(httpServer.GetHttpPort()), int32(grpcPort))

	grpcService := api.GetRpcService()
	go grpcService.HandleReadyProcess()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-sigChan:
		os.Exit(1)
	}
}
