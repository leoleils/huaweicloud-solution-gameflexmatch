- [1. Game Flex Match 编译指导](#1-game-flex-match-编译指导)
  - [1.1. FleetManager](#11-fleetmanager)
  - [1.2. AppGateway](#12-appgateway)
  - [1.3. AASS](#13-aass)
  - [1.4. Auxproxy](#14-auxproxy)
  - [1.5. Console](#15-console)

# 1. Game Flex Match 编译指导
若你需要从源码进行编译安装，或者在本地进行自定义开发后再进行使用，可以参考该文档：
在安装部署之前，你需要先准备一些环境信息，以下操作基于`Centos`操作系统：
+ 后端编译所需环境：
    安装`git`：
    ```sh
    yum install -y git
    ```
    配置`go`语言环境：
    ```sh
    # 配置go语言环境 
    yum install -y golang 
    go env -w GOOS=linux
    # 配置go语言代理
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://repo.huaweicloud.com/repository/goproxy/
    go env -w GONOSUMDB=*
    ```
## 1.1. FleetManager
+ 编译
```shell
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-fleetmanager.git
cd huaweicloud-solution-gameflexmatch-fleetmanager
go mod tidy
go build -o fleetmanager ./main.go
```
+ 将编译后的二进制包，移动至`./game-flex-match_release/bin`目录下
## 1.2. AppGateway
+ 编译
```shell
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-appgateway.git
cd huaweicloud-solution-gameflexmatch-appgateway
go mod tidy
go build -o appgateway ./cmd/application_gateway.go
```
+ 将编译后的二进制包，移动至`./game-flex-match_release/bin`目录下
## 1.3. AASS
+ 编译
```shell
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-aass.git
cd huaweicloud-solution-gameflexmatch-aass
go mod tidy
go build -o aass ./cmd/application-auto-scaling-service/application_auto_scaling_service.go
```
+ 将编译后的二进制包，移动至`./game-flex-match_release/bin`目录下
## 1.4. Auxproxy
+ 编译
```shell
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-auxproxy.git
cd huaweicloud-solution-gameflexmatch-auxproxy
go mod tidy
go build -o auxproxy ./cmd/auxproxy.go
```
+ 解压`game-flex-match_release/bin/auxproxy.zip`，将编译后的二进制包，移动至`game-flex-match_release/bin/auxproxy`，并重新压缩成`auxproxy.zip`
    `cd ./game-flex-match_release/bin/auxproxy`
    `zip -r -o ../auxproxy.zip ./*`
+ 将压缩后的`auxproxy.zip`上传至`OBS`桶中

## 1.5. Console
+ 配置编译环境：
  `yum install -y nodejs`
  `yum install -y npm`
+ 编译
```sh
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-console.git
cd huaweicloud-solution-gameflexmatch-console
```
+ 将`game-flex-match_release/conf/public.pem`公钥文件复制到`huaweicloud-solution-gameflexmatch-console/cert/public.pem`
+ 在`huaweicloud-solution-gameflexmatch-console`目录下执行`npm install` ，用于安装必要依赖
+ 执行: `npm run build` 进行编译
+ 在目录下会得到编译后的`huaweicloud-solution-gameflexmatch-console/dist`前端应用