# C++ GRPC demo

## 环境

OS：Windows/Linux

c++：ISO C++14

第三方库：grpc，protobuf

其他：cmake >= 3.15.0, Visual Studio 2019

提示：windows环境可使用vcpkg快速安装和编译grpc相关的环境



## Demo

demo目录

```
demo
 |-build  // cmake build directory and demo project
 |-src    // demo source code
 |-proto  // sdk and protobuf file
```



1. 安装好grpc，protobuf和cmake环境后，将protoc加入系统环境变量，例：“D:\grpc\vcpkg\packages\protobuf_x64-windows\tools\protobuf”

2. 进入demo根目录/proto，使用protoc和grpc_cpp代码生成工具生成grpc和 proto的sdk文件

   ```powershell
   cd /proto
   
   protoc --proto_path=. --cpp_out=. process_grpc_service.proto
   protoc --proto_path=. --grpc_out=. --plugin=protoc-gen-grpc="D:\grpc\vcpkg\packages\grpc_x64-windows\tools\grpc\grpc_cpp_plugin.exe" process_grpc_service.proto
   
   protoc --proto_path=. --cpp_out=. auxproxy_grpc_service.proto
   protoc --proto_path=. --grpc_out=. --plugin=protoc-gen-grpc="D:\grpc\vcpkg\packages\grpc_x64-windows\tools\grpc\grpc_cpp_plugin.exe" auxproxy_grpc_service.proto
   ```

3. 进入根目录/build路径，使用cmake编译为VS 项目

   ```powershell
   cd /build
   cmake -G "Visual Studio 16 2019" ../ -DCMAKE_TOOLCHAIN_FILE=D:/grpc/vcpkg/scripts/buildsystems/vcpkg.cmake
   cmake --build .
   ```

4. 运行build目录下的VS项目文件 game-example.sln

5. 右键项目->重新生成，将在项目路径下的Debug目录生成server.exe和client.exe

6. 运行server.exe，再运行client.exe

   

7. server端输出

   ```
   $ ./server.exe
   Server listening on 0.0.0.0:60002
   [ProcessReady] get process ready, time Tue Nov 14 11:15:50 2023
   
   client port: 8080
   client grpc port: 50051
   client pid: 123
   [ActivateServerSession] get activate session
   session id: session_id
   max client: 1
   call back to localhost:50051
   send on health check
   health check success response: 1
   health status is: 1
   on start health check success response: 1
   start on start server session
   [ActivateServerSession] get activate session
   session id: session_id_cZFXENx4s8tcvX9PXH7xqytqjVHK9adQ
   max client: 1
   on start server session success response ok
   start OnProcessTerminate
   on start terminate success response: 1
   [TerminateServerSession] terminated server session: session id
   [ProcessEnding] terminated process : 123
   ```

   client端输出

   ```
   $ ./client.exe
   Server listening on localhost:50051
   process running
   [RunProcessReady] send process ready
   [ProcessReady] process ready.
   [RunProcessReady] process ready: 1
   --------------- process ready --------------
   process running[ActivateServerSession] session
   session_id active now
   [RunActivateServerSession] success activate serversession
   [RunActivateServerSession] game session_id activated
   -------  gaming running 1 second  ----------
   [OnHealthCheck] get health check request
   receive session:
   server session id: session_id_cZFXENx4s8tcvX9PXH7xqytqjVHK9adQ
   fleetn id: fleet id
   name: name
   maxClients: 1
   joinable: 1
   port: 0
   ipAddress: 0.0.0.0
   sessionData: session data
   activing session, please wait ...
   process running
   -------  gaming running 2 second  ----------
   session active, player can enter/start game session
   [ActivateServerSession] session session_id_cZFXENx4s8tcvX9PXH7xqytqjVHK9adQ active now
   [OnHealthCheck] get process terminate
   [OnHealthCheck] process terminate time 10
   process running
   -------  gaming running 3 second  ----------
   [OnHealthCheck] process terminated
   -------  gaming running 4 second  ----------
   -------  gaming running 5 second  ----------
   [RunActivateServerSession] game session_id finish
   [TerminateServerSession] terminated Client session.
   [RunActivateServerSession] game session id terminate success
   [ProcessEnding] process 123prepare to ending
   [ProcessEnding] ending process success: 123
   [RunActivateServerSession] game 123 terminate success
   ```

   

