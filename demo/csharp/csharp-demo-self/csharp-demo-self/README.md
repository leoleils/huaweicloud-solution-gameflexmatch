# 安装 .Net Core 3.1 SDK

该项目依赖 .Net Core SDK，在编译运行代码前需要先安装 .Net Core 3.1 SDK。以 CentOS 操作系统为例，版本不得低于 CentOS 7 或 CentOS 8。

1. 添加签名密钥
```
sudo rpm -Uvh https://packages.microsoft.com/config/centos/7/packages-microsoft-prod.rpm
```

2. 安装 .NET Core SDK 
```
sudo yum install dotnet-sdk-3.1
```

# 编译运行

1. 下载代码，在csharp-demo目录下，执行如下指令即可自动编译并运行服务
```
dotnet run
```
程序正确编译运行后，会在 csharp-demo/obj/Debug/netcoreapp3.1 文件夹生成项目依赖的库、二进制文件以及 proto 文件编译后的 .cs 文件。**注意**，proto 文件的引入是在 csharp-demo/csharpdemo.csproj 中：
```
<Protobuf Include="..\proto\csharp-demo\GameServerGrpcSdkService.proto" Link="GameServerGrpcSdkService.proto"/>
<Protobuf Include="..\proto\csharp-demo\GseGrpcSdkService.proto" Link="GseGrpcSdkService.proto" />
```
项目依赖于 proto/csharp-demo 文件夹中的 GameServerGrpcSdkService.proto 和 GseGrpcSdkService.proto 两个 proto 文件。

# 打包生成包

1. 生成可执行文件及依赖
```
dotnet publish -c Release -r linux-x64 --self-contained true
```
以上会在 csharp-demo/bin/Release/netcoreapp3.1/linux-x64 目录下生成打包生成包所需要的所有依赖文件，其中即包含运行该服务的可执行文件**csharpdemo**。

2. 拷贝前置脚本install.sh
```
chmod u+x install.sh
cp install.sh bin/Release/netcoreapp3.1/linux-x64
```

3. 打包 GSE 生成包
```
cd csharp-demo/bin/Release/netcoreapp3.1/linux-x64
zip -r csharpdemo.zip *
```
打包好的 csharpdemo.zip 即 GSE 需要的生成包。创建 Fleet 前需要上传该生成包，**启动路径填写csharpdemo**。
