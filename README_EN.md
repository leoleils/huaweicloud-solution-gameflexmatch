# huaweicloud-solution-GameFlexMatch #

Language: [中文](/README.md) | **ENGLISH**

# Brief Introduction #

`GameFlexMatch`is a service hosting solution that consists of four service components (`Fleetmanager`/`AppGateway`/`AASS`/`AuxProxy`), which can implement application hosting, elastic scaling of resources required by hosting applications, resource scheduling and management of application processes, and gray release of applications. Multi-region deployment enables users to access the nearest network, reducing latency and cross-region DR of service resources. 
It helps developers quickly build a stable and low-latency multiplayer game deployment environment and saves a lot of O&M costs.`Unreal Engine`,`Unity`,`C#`,`C++`And also the`gRPC`any language supported`server`Deploy and run the framework.

# Logical Architecture #

![Image](doc/img/architecture-EN.jpg)	

The GameFlexMatch platform consists of five service components:

 *  [FleetManager](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-fleetmanager): Globally deploys and manages application processes, supports configuration of dynamic deployment policies, optimizes application distribution based on costs or latency, configures auto scaling policies, manages server sessions, client sessions, and application packages, and supports gray release of server applications.
 *  [AppGateway](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-appgateway): Manages application processes, sessions, and client connections.`AuxProxy`Obtain application process information through communication, and make decision on scheduling process resources.
 *  [AASS](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-aass): Manages and executes AS groups and policies, monitors application resources on the server, and implements elastic resource scaling.
 *  [AuxProxy](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-auxproxy): Automatically starts the instance after capacity expansion, which creates application processes, reports process status, and communicates with application processes.
 *  [Console](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-console): O&M platform, used for monitoring`GameFlexMatch`Running status and O&M management of the`GameFlexMatch`of the`fleet`2. Application packages and user information

## Content
```
build                   # Project build directory
doc                     # Document directory
    |-- api             # Management plane API document
    |-- deployment      # Solution deployment operation document
    |-- developer       # Application access development document
    |-- user-guide      # User operation guide
    |-- version         # Version update records
sdk                     #App access SDK and demo
src                     #Source code directory of the project component
tools                   #Auxiliary tools for application deployment and usage
```

## Deployment Guide ##

 *  The solution runs on Huawei cloud services. Resources involved include ECS, IMS, OBS, VPC, and CES. MySQL
    ，Redis and InfluxDB are used. 
 *  For details about front-end and back-end deployment, see.[doc/deployment/deployment-guide-EN.md](doc/deployment/deployment-guide-EN.md)
 *  For details about the parameters required during the deployment, see.[doc/build/param-annotation.md](/build/param-annotation-EN.md)

## Usage Guide ##

 *  For details about the console user guide, see [/doc/user-guide/gameflexmatch-user-guide-EN.md](doc/user-guide/gameflexmatch-user-guide-EN.md)

## Application Access Guide ##

 *  Apps can be hosted on GameFlexMatch in gRPC mode. For details about the interface and access process, see [doc/developer/developer_guide_EN.md](doc/developer/developer_guide_EN.md)
 *  For details about the example of application hosting access, see.`/demo`

## Reference ##

 *  API Reference Document [doc/api/FleetManager.yaml](doc/api/FleetManager.yaml)

## Additional information ##

 *  For details about the version update history, see.`doc/version/`
 *  Platform FAQ Reference [doc/user-guide/help/help.md](doc/user-guide/help/help.md)

## Auxiliary Tools ##

1.  Encryption tool: provides tools for encrypting and decrypting sensitive data using GCM and RSA. For details, see.`/tools/cipher`
2.  Peer-to-peer connection tool: provides two`VPC`For details about how to create a VPC peering connection, see.`/tools/vpc-peering`

