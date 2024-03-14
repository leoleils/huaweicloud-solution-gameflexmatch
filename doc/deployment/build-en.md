- [1. Game Flex Match Build Instructions](#1-game-flex-match-build-instructions)
  - [1.1. FleetManager](#11-fleetmanager)
  - [1.2. AppGateway](#12-appgateway)
  - [1.3. AASS](#13-aass)
  - [1.4. Auxproxy](#14-auxproxy)
  - [1.5. Console](#15-console)

# 1. Game Flex Match Build Instructions
If you need to compile and install from the source code, or use it after local custom development, you can refer to this document:
Before the installation and deployment, you need to prepare some environment information. The following operations are based on the `Centos` operating system:
+ Environment required for backend compilation:
    Install `git`:
    ```sh
    yum install -y git
    ```
    Configure the `go` locale:
    ```sh
    # Configure the Go language environment. 
    yum install -y golang 
    go env -w GOOS=linux
    # Configure the Go language proxy.
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://repo.huaweicloud.com/repository/goproxy/
    go env -w GONOSUMDB=*
    ```
## 1.1. FleetManager
+ Compile
```Shell
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-fleetmanager.git
cd huaweicloud-solution-gameflexmatch-fleetmanager
go mod neat
go build -o fleetmanager. /main.go
```
+ Move the compiled binary package `./fleetmanager` to: `/home/gfm/bin`
## 1.2. AppGateway
+ Compile
```shell
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-appgateway.git
cd huaweicloud-solution-gameflexmatch-appgateway
go mod tidy
go build -o appgateway ./cmd/application_gateway.go
```
+ Move the compiled binary package `./appgateway` to: `/home/gfm/bin`
## 1.3. AASS
+ Compile
```shell
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-aass.git
cd huaweicloud-solution-gameflexmatch-aass
go mod tidy
go build -o aass ./cmd/application-auto-scaling-service/application_auto_scaling_service.go
```
+ Move the compiled binary package: `./aass` to: `/home/gfm/bin`
## 1.4. Auxproxy
+ Compile
```shell
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-auxproxy.git
cd huaweicloud-solution-gameflexmatch-auxproxy
go mod tidy
go build -o auxproxy ./cmd/auxproxy.go
```
+ Decompress `/home/gfm/bin/auxproxy.zip`, move the compiled binary package to `/home/gfm/bin/auxproxy`, and compress it into `auxproxy.zip`.
    `cd /home/gfm/bin/auxproxy`
    `zip -r -o ../auxproxy.zip  ./*`
+ Upload the compressed `auxproxy.zip` to the `OBS` bucket.

## 1.5. Console
+ Configuring the Compilation Environment
```shell
  yum install -y nodejs
  yum install -y npm
```
+ Compile
```sh
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-console.git
cd huaweicloud-solution-gameflexmatch-console
```
+ Copy `/home/gfm/conf/public.pem` public key file to `huaweicloud-solution-gameflexmatch-console/cert/public.pem`
+ Run `npm install` in the `huaweicloud-solution-gameflexmatch-console` directory to install necessary dependencies.
+ Execute: `npm install --save vue-i18n@next`
+ Execute: `npm run build` to compile
+ You will get the compiled `huaweicloud-solution-gameflexmatch-console/dist` front-end application in the directory
+ Move `huaweicloud-solution-gameflexmatch-console/dist` to the `/home/gfm/bin` directory.