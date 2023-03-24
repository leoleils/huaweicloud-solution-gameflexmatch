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
	"fake-server/api"
	"fake-server/config"
	"fake-server/gsemanager"
	"fake-server/logger"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func startGrpcServer() int {
	// 启动grpc server，监听agent回调
	grpcServer := api.GetRpcService()
	grpcServer.StartGrpcServer()
	grpcPort := grpcServer.GetGrpcPort()

	// 返回 grpc port
	return grpcPort
}

// main intercepts the log file of the SuperTuxKart gameserver and uses it
// to determine if the game server is ready or not.
const (
	// DefaultlogPath	= "D://Desktop//workdir//MetaSpace//MetaSpace-dev//fake-server//log"
	DefaultlogPath	= "/local/app/fake-server/log"
	DefaultHttpStartPort = 1025
	DefaultHttpEndPort = 60001
)

func main() {
	// 启动Grpc Server
	flag.StringVar(&config.GlobalConfig.LogPath, "log-path", DefaultlogPath, "log path")
	flag.IntVar(&config.GlobalConfig.HttpStartPort, "start-port", DefaultHttpStartPort, "http server start port")
	flag.IntVar(&config.GlobalConfig.HttpEndPort, "end-port", DefaultHttpEndPort, "http server end port")
	flag.Parse()
	
	logger.Logger, _ = logger.Init()
	grpcPort := startGrpcServer()

	gseManager := gsemanager.GetGseManager()

	if config.GlobalConfig.HttpEndPort <= config.GlobalConfig.HttpStartPort {
		logger.Logger.Errorf("http end port should lagger then http start port")
		os.Exit(0)
	}
	logger.Logger.Infof("http port will be [%d, %d)", config.GlobalConfig.HttpStartPort, config.GlobalConfig.HttpEndPort)
	httpServer := api.NewHttpProcess()
	httpServer.StartHttpServer()
	gseManager.ProcessReady(nil, int32(httpServer.GetHttpPort()), int32(grpcPort))
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-sigChan:
		os.Exit(1)
	}
}
