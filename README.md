# huaweicloud-solution-metaspace
---
## 简介

`MetaSpace`是一个服务端应用托管平台，包含四个服务组件(`Fleetmanager`/`AppGateway`/`AASS`/`AuxProxy`)，可以实现应用的托管、托管应用所需资源的弹性伸缩、应用进程的资源调度管理、应用的灰度发布，多`region`部署时可以实现用户的就近接入，减少时延，以及服务资源的跨地域容灾。可以帮助开发者快速构建稳定、低延时的多人游戏的部署环境，并节省大量的运维成本，支持`Unreal`、`Unity`引擎，`C#`、`C++`以及`gRPC`支持的任何语言的`server`框架部署和运行。


## 逻辑架构
<img src="/img/architecture.jpg" width="80%">

MetaSpace平台由五个服务组件组成：
+ [FleetManager](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-metaspace-fleetmanager): 负责应用进程的全局化动态部署及管理，支持配置动态部署策略，基于成本或时延优化应用分布，负责弹性伸缩策略的配置和服务端会话、客户端会话与应用包的管理，服务端应用的灰度发布等
+ [AppGateway](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-metaspace-appgateway): 负责应用进程、会话与客户端连接的管理，通过与`AuxProxy`通信获得应用进程信息，决策进程资源的调度
+ [AASS](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-metaspace-aass): 负责弹性伸缩组和弹性伸缩策略的管理与执行，以及服务端应用资源的监控，调用华为云`AS`(弹性伸缩服务)实现资源的弹性伸缩
+ [AuxProxy](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-metaspace-auxproxy): 在扩容出的实例中自动拉起，负责应用进程的创建、进程状态的上报以及应用进程的通信
+ [Console](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-metaspace-console): 运维平台，用于监控`metaspace`的运行状态，以及运维管理`metaspace`的`fleet`、应用包与用户信息等

## 组织结构

    |-- huaweicloud-solution-metaspace
        |-- demo
        |   |-- go                  -- go语言的服务SDK Demo
        |-- doc
        |   |-- api                 -- api技术文档
        |   |-- build               -- 测试环境部署指导
        |   |-- developer           -- 开发者对接指南(c#)
        |-- img                     
        |-- tools                   -- 使用脚本工具
        |   |-- cipher              -- 敏感数据的加密解密工具
        |   |-- vpc-peering         -- vpc之间创建对等连接工具
## 服务端SDK-Demo
1. 提供了`go`语言的服务端`SDK Demo`，可以与`Metaspace`平台进行无缝对接
2. 提供了`C#`语言的开发者对接指南，详见 [doc/developer](/doc/developer/developer_guide.md)

## API技术文档
提供了[fleetmanager](/doc/api/FleetManager.yaml)/[appgateway](/doc/api/AppGateway.yaml)/[aass](/doc/api/AASS.yaml)的API技术文档的`Yaml`文件，可以在[swagger](https://editor.swagger.io/)中导入查看，参考文档详见：`/doc/api/`

## 部署指南
1. **准备工作**
   + 准备华为云资源
      + 管理账号与资源账号：管理账号用于`MetaSpace`的管理面服务组件的管理与执行，资源账号用于计算资源的申请，账号架构详见[doc/build/user-management.md](/doc/build/user-management.md)
      
      + 创建委托资源账号委托给管理账号：
         - 创建委托可以参考链接[创建委托（委托方操作）](https://support.huaweicloud.com/intl/zh-cn/usermanual-iam/iam_06_0002.html)，并将资源账号委托给管理账号；
         - 授予该委托`DEW KeypairFullAccess`权限
         - 新建委托策略权限，增加委托权限策略，委托权限`json`视图见[doc/build/agency_region.json](/doc/build/agency_region.json)与[doc/build/agency_global.json](/doc/build/agency_global.json)，由于区域级委托与全局级委托不能同时配置，该步骤需要为委托新建两种权限并关联
         
      + (可选)若想使用LTS配置日志转存，需在委托资源账号下创建LTS委托，给ECS以安装ICagent：
        
           委托配置流程见[创建icagent委托](https://support.huaweicloud.com/usermanual-lts/lts_03_0002.html)。
           
   + 准备管理面资源，并将管理面资源部署在同一`VPC`下，以下测试规格，具体规格按需选择：
           
           | 购买账号 |      资源类型       |  资源规格  | 数量  |
           | :------: | :-----------------: | :--------: | :---: |
           | 资源账号 |         ECS         | 2vCPUs/4GB |   3   |
           | 资源账号 |         RDS         | 2vCPUs/4GB |   3   |
           | 资源账号 | GaussDB(for Influx) | 2vCPUs/4GB |   1   |
        
           RDS可以按需选择单机或主备节点
           influxdb选择集群(默认3节点)

      + 在华为云`ECS`云服务中购买`3`台`ECS`，分别部署`MetaSpace`服务组件，并放通安全组相关端口：
           |        |   服务组件   | 监听端口 | 是否必须配置EIP |
           | :----: | :----------: | :------: | :-------------: |
           | ECS-01 |  appgateway  |  60003   |        Y        |
           | ECS-02 |     aass     |   9091   |        N        |
           | ECS-03 | fleetmanager |  31002   |        Y        |
           入方向至少需要保障`60003`端口和`31002`端口开放
      
      + 准备`RDS`数据库，默认端口为`3306`，依次为三个服务组件(`appgateway`/`aass`/`fleetmanager`)创建数据库，创建用户并授予**读写权限**
      + 创建`GaussDB(for Influx)`：选择购买`InfluxDB`，并开启`SSL`安全连接，使用默认证书即可，为服务组件创建数据库(`aass`/`appgateway`)，`aass`与`appgateway`共用一个`influxDB`的数据库
      + [管理租户](/doc/build/user-management.md)创建`AK`与`SK`，参考链接[管理IAM用户访问密匙](https://support.huaweicloud.com/usermanual-iam/iam_02_0003.html)，用于访问资源租户
      + [资源租户](/doc/build/user-management.md)新建密匙对，用于弹性伸缩实例的密匙验证登录

2. **环境依赖**:
   + go1.16及以上版本
3. **文件编译**：
   + 将源码下载到本地，编译`linux`可执行的二进制文件，**以下步骤中{version}中的变量需按具体情况更改**
    ```sh
        # 设置编译的可执行文件的操作系统
        go env -w GOOS=linux
        # 1. fleetmanager
        cd ~/huaweicloud-solution-metaspace-fleetmanager
        go build ./main.go
        # 修改文件名
        mv main fleetmanager-{version}

        # 2. appgateway
        cd ~/huaweicloud-solution-metaspace-appgateway
        go build ./cmd/application_gateway.go
        # 修改文件名
        mv application_gateway appgateway-{version}

        # 3. aass
        cd ~/huaweicloud-solution-metaspace-aass
        go build ./cmd/application-auto-scaling-service/application_auto_scaling_service.go
        # 修改文件名
        mv application_auto_scaling_service aass-{version}

        # 4. auxproxy
        cd ~/huaweicloud-solution-metaspace-auxproxy
        go build ./cmd/auxproxy.go

    ```
4. **证书准备**
   在`linux`系统下通过`openssl`获取自签名证书，可在任一台`ECS`下操作，三个服务组件使用相同的自签名证书：
   + 获取`https`签名证书
    ```sh
        # 1. 创建tlsSecret文件夹
        mkdir -p /home/tlsSecret
        cd /home/tlsSecret
        # 2. 生成私钥tls.key
        openssl genrsa -out tls.key 3072
        # 3. 使用私钥生成csr，并查看
        openssl req -new -key tls.key -out tls.csr
        # 上述步骤会要求输入以下信息，可按实际情况填写，如
        # There are quite a few fields but you can leave some blank
        # For some fields there will be a default value,
        # If you enter '.', the field will be left blank.

        # Country Name (2 letter code) [XX]:China
        # State or Province Name (full name) []:GuangDong
        # Locality Name (eg, city) [Default City]:ShenZhen
        # Organization Name (eg, company) [Default Company Ltd]:Huawei
        # Organizational Unit Name (eg, section) []:Cloud
        # Common Name (eg, your name or your server's hostname) []:MetaSpace
        # Email Address []:metaspace@huawei.com

        # Please enter the following 'extra' attributes
        # to be sent with your certificate request
        # A challenge password []:metaspace@123
        # An optional company name []:Huawei

        openssl req -in tls.csr -text
        # 4. 生成自签名证书 tls.crt, 并查看
        openssl x509 -req -days 365 -in tls.csr -signkey tls.key -out tls.crt
        openssl x509 -in tls.crt -text
    ```
   + 获取网络传输的RSA非对称加密的公钥与私钥，用户敏感数据的加密与解密
    ```sh
        # 1. 创建RSA私钥，长度可以为1024，也可以为2048
        cd /home/tlsSecret
        openssl genrsa -out rsa_private.pem 2048
        # 2. 在私钥的基础上生成公钥
        openssl rsa -in rsa_private.pem -pubout -out rsa_public.pem
    ```
5. **服务组件安装**
   + 安装`appgateway`服务组件

    ```sh
        # 1. 登录ECS-01，新建/home/tlsSecret，并上传已生成的tls.crt与tls.key，以及RSA非对称加密的公钥与私钥rsa_private.pem，rsa_public.pem
        mkdir -p /home/tlsSecret
        # 2. 创建文件夹/home/appgateway/conf/hmac，
        mkdir -p /home/appgateway/conf/hmac
        # 上传client_hmac_conf.json、server_hmac_conf.json上传至hmac文件夹，样例见: 
        # doc/build/appgateway/client_hmac_conf.json
        # doc/build/appgateway/server_hmac_conf.json

        # 3. 新建bin目录，上传可执行二进制文件与启动脚本并修改权限
        mkdir -p /home/appgateway/bin
        cd /home/appgateway/bin
        # 上传生成的appgateway-{version}二进制文件与启动脚本appgateway_run.sh，并修改相关配置(样例见：doc/build/appgateway/appgateway_run.sh)
        # 修改权限
        chmod 750 appgateway-{version}
        chmod 750 appgateway_run.sh

        # 4. 配置完成后运行启动脚本
        ./appgateway_run.sh

        # 5. 验证是否执行成功
        ps -aux | grep appgateway
        
        # 6. 关掉进程后配置开机自启动
        # 在/etc/systemd/system下新建appgateway.service
        # appgateway.service 样例见 /doc/build/appgateway
        # 启动appgateway.service保证进程自动拉起
        systemctl enable appgateway.service
        systemctl start appgateway.service

        # 7. 验证是否成功
        ps -aux | grep appgateway
        # 8. 可以正常运行则为部署成功

    ```

   + 安装`aass`服务组件

    ```sh
        # 1. 登录ECS-02，并创建tlsSecret文件夹
        mkdir -p /home/tlsSecret
        # 2. 将已生成的tls.crt与tls.key上传至tlsSecret文件夹中
        # 3. 创建文件夹/home/aass
        mkdir -p /home/aass/configmap
        # 4. 将server_hmac_conf.json与service_config.json上传至configmap中,修改相关配置，样例见：
        # doc/build/aass/server_hmac_conf.json
        # doc/build/aass/service_config.json

        # 5. 创建文件夹/home/aass/bin
        mkdir -p /home/aass/bin
        # 6. 将aass的二进制可执行文件aass-{version}与执行脚本aass_run.sh上传至bin目录下，修改相关配置，并修改文件权限
        cd /home/aass/bin
        chmod 750 aass-{version}
        chmod 750 aass_run.sh

        # 7. 执行并验证是否执行成功
        sh ./aass_run.sh
        ps -aux | grep aass

        # 8. Ctrl+c 关掉进程后配置开机自启动
        # 在/etc/systemd/system下新建aass.service
        # aass.service 样例见 /doc/build/aass
        # 启动aass.service保证进程自动拉起
        systemctl enable aass.service
        systemctl start aass.service

        # 9. 验证是否成功
        ps -aux | grep aass
        # 10. 可以正常运行则为部署成功

    ```

   + 安装`fleetmanager`服务组件

    ```sh
        # 1. 登录ECS-03，创建文件夹tlsSecret并上传已生成的tls.crt与tls.key文件
        mkdir -p /home/tlsSecret
        # 2. 创建/home/configmap文件夹
        mkdir -p /home/fleetmanager/configmap
        # 3. 上传server_config.json到configmap下，并修改相关配置
        # 样例见 doc/build/fleetmanager/server_config.json
        
        # 4. 创建文件夹/home/fleetmanager/bin/conf/workflow
        mkdir -p /home/fleetmanager/bin/conf/workflow

        # 5. 上传create_fleet_workflow.json、delete_fleet_workflow.json以及create_build_image_workflow.json，详见
        # doc/build/fleetmanager/create_fleet_workflow.json
        # doc/build/fleetmanager/delete_fleet_workflow.json
        # doc/build/fleetmanager/create_build_image_workflow.json
        
        # 6. 上传fleetmanager的二进制可执行文件fleetmanager-{version}
        # 与启动脚本fleetmanager_run.sh上传至bin文件夹，修改相关配置与文件权限
        cd /home/fleetmanager/bin
        chmod 750 fleetmanager-{version}
        chmod 750 fleetmanager_run.sh

        # 7. 启动脚本并验证是否成功
        sh ./fleetmanager_run.sh
        ps -aux | grep fleetmanager

        # 8. 关掉进程后配置开机自启动
        # 在/etc/systemd/system下新建fleetmanager.service
        # fleetmanager.service 样例见 /doc/build/fleetmanager
        # 启动fleetmanager.service保证进程自动拉起
        systemctl enable fleetmanager.service
        systemctl start fleetmanager.service

        # 9. 验证是否成功
        ps -aux | grep fleetmanager
        # 10. 可以正常运行则为部署成功
    ```
6. **其他说明**：
   + 前端部署指导详见 [doc/build/console.md](/doc/build/console.md)
   + 应用镜像制作详见 [doc/build/make-image-guide.md](/doc/build/make-image-guide.md)
   + 应用的资源数据导入详见 [doc/build/user-data-import.md](/doc/build/user-data-import.md)
   + 部署过程中必要的参数注解详见 [doc/build/param-annotation.md](/doc/build/param-annotation.md)
   + 平台用户管理模块使用详见 [doc/build/user-management.md](/doc/build/user-management.md)
   + console平台的用户指南详见`doc/user-guide`

## 日志导出功能
Metaspace平台可以借助华为云LTS服务，实现服务组件以及托管应用的日志转存功能
详细过程步骤请[参考链接](https://support.huaweicloud.com/usermanual-lts/lts_04_1031.html)

## 辅助工具
1. 加密工具：提供了GCM与RSA加解密敏感数据的工具，详见`/tools/cipher`
2. 对等连接工具：提供了两个`VPC`创建对等连接的工具，详见`/tools/vpc-peering`