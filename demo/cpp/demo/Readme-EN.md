# C++ GRPC demo

## Environment

OS：Windows/Linux

c++：ISO C++14

library：grpc，protobuf, Visual Studio 2019

other：cmake >= 3.15.0



Note：In the Windows environment, you can use vcpkg to quickly install and compile grpc-related environments.

## Demo

demo directory

```
demo
 |-build  // cmake build directory and demo project
 |-src    // demo source code
 |-proto  // sdk and protobuf file
```



1. After the grpc, protobuf, and cmake environments are installed, add protoc to the system environment variable, for example, D:\grpc\vcpkg\packages\protobuf_x64-windows\tools\protobuf.

2. Go to the demo root directory /proto and use the protoc and grpc_cpp code generation tools to generate the grpc and proto SDK files.

   ```powershell
cd /proto
   
   protoc --proto_path=. --cpp_out=. process_grpc_service.proto
   protoc --proto_path=. --grpc_out=. --plugin=protoc-gen-grpc="D:\grpc\vcpkg\packages\grpc_x64-windows\tools\grpc\grpc_cpp_plugin.exe" process_grpc_service.proto
   
   protoc --proto_path=. --cpp_out=. auxproxy_grpc_service.proto
   protoc --proto_path=. --grpc_out=. --plugin=protoc-gen-grpc="D:\grpc\vcpkg\packages\grpc_x64-windows\tools\grpc\grpc_cpp_plugin.exe" auxproxy_grpc_service.proto
   ```
   
3. Go to the root directory /build and compile the Visual Studio project by cmake.

   ```powershell
cd /build
   cmake -G "Visual Studio 16 2019" ../ -DCMAKE_TOOLCHAIN_FILE=D:/grpc/vcpkg/scripts/buildsystems/vcpkg.cmake
   cmake --build .
   ```
   
4. Run the VS project file game-example.sln in the build directory.

5. Right-click the project and choose Rebuild from the shortcut menu. The server.exe and client.exe files are generated in the Debug directory in the project path.

6. Run server.exe, then client.exe

7. Output on the server

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
   
   Output on the client

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
   
   

