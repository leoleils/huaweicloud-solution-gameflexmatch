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
	"fake-server/gsemanager"
	"os"
	"os/signal"
	"syscall"
	"C"
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
func main() {
	// 启动Grpc Server
	grpcPort := startGrpcServer()

	gseManager := gsemanager.GetGseManager()

	gseManager.ProcessReady(nil, 1111, int32(grpcPort))

	httpServer := api.NewHttpProcess()
	httpServer.StartHttpServer()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-sigChan:
		os.Exit(1)
	}
	defer close(sigChan)
}
