- [1. sac-gfm脚本介绍](#1-sac-gfm脚本介绍)
  - [1.1. 功能介绍](#11-功能介绍)
  - [1.2. 命令与参数介绍](#12-命令与参数介绍)

# 1. sac-gfm脚本介绍
## 1.1. 功能介绍
`sac-gfm`是一个帮助你部署与管理`GFM`的二进制文件，主要的功能如下：

- [X] 生成加解密的信息与相关的秘钥文件，并配置到配置文件中
- [X] 对配置文件进行分析处理，为每个服务组件生成对应的配置文件，并生成启动脚本，配置
- [X] 对GFM进行管理，支持启动、停止、重启各个服务组件

## 1.2. 命令与参数介绍
`sac-gfm`的使用方式为：`sac-gfm [global options] command [command options] [arguments...]`

你也可以直接运行`sac-gfm`或`sac-gfm -h`查看命令帮助：
```sh
$ cd ./gfm
$ ./sac-gfm -h

the work dir is: /home/gfm
NAME:
   sac-gfm - Usages: game-flex-match run script

USAGE:
   sac-gfm [global options] command [command options] [arguments...]

VERSION:
   v1.0

COMMANDS:
   cipher   cipher for key info, can help you encode or decode key info, and generate key file for gcm/rsa/tls/jwt
   install  install service info, can help you to install fleetmanager/appgateway/aass
   start    start a service, support: [fleetmanager, appgateway, aass]
   stop     stop a service, support: [fleetmanager, appgateway, aass]
   restart  restart a service, support: [fleetmanager, appgateway, aass]
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```
|  命令   | 说明                                                                                                   |
| :-----: | :----------------------------------------------------------------------------------------------------- |
| cipher  | 生成加解密的秘钥，并可以帮你加密或解密敏感信息                                                         |
| install | 根据你已配置好的配置文件，自动帮你生成各个服务组件对应的配置文件，并会帮你生成启动脚本，配置开机自启动 |
|  start  | 帮你启动一个服务                                                                                       |
|  stop   | 帮你停止一个服务                                                                                       |
| restart | 帮你重启一个服务                                                                                       |
|  help   | 输出帮助文档                                                                                           |

**各个命令的使用方式以及样例如下**
+ `cipher`命令：`./sac-gfm cipher [command options] [arguments...]`，命令支持如下：
  |    命令     |                参数                 |      默认值      | 说明                                                                                                                                                                                   |
  | :---------: | :---------------------------------: | :--------------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
  |   --mode    |   string: oneof [decode, encode]    |      encode      | 用于为敏感信息加密或解密                                                                                                                                                               |
  |  --method   |      string: oneof [gcm, rsa]       |       gcm        | 为敏感信息加密或者解密的方法                                                                                                                                                           |
  |   --text    |               string                |                  | 用于加密或者解密的字符                                                                                                                                                                 |
  | --init-conf |                path                 | ./conf/init.yaml | 配置文件路径，当你需要生成秘钥信息时会自动将生成的信息同步至该配置文件中                                                                                                               |
  |  --create   | string: one of [gcm, rsa, tls, jwt] |                  | 当参数为过程`gcm`或`jwt`时，会将生成的敏感信息同步至`--init-conf`对应的配置文件中，当为`rsa`或`tls`时，会在`--init-conf`对应的配置文件同级目录中生成配置文件，且会同步至对应配置文件中 |
  | --help, -h  |                                     |                  | 生成帮助信息                                                                                                                                                                           |
    
    可生成的帮助信息如下
    ```sh
    $ ./sac-gfm cipher -h

    the work dir is: /home/gfm
    NAME:
    sac-gfm cipher - cipher for key info, can help you encode or decode key info, and generate key file for gcm/rsa/tls/jwt

    USAGE:
    sac-gfm cipher [command options] [arguments...]

    OPTIONS:
    --mode value       mode for cipher, options: [decode | encode], default: encode (default: "encode")
    --method value     method for cipher, options: [gcm | rsa], default: gcm (default: "gcm")
    --text value       text to decode or encode for cipher, will use the text to cipher if it had
    --init-conf value  init conf using for cipher
    --create value     create gcm key and nonce, create rsa key file, create tls key file, options: [gcm, rsa, tls, jwt]
    --help, -h         show help
    ```
+ `install`命令：`sac-gfm install [command options] [arguments...]`
    |    命令     |                      参数                      |        默认值         | 说明                                 |
    | :---------: | :--------------------------------------------: | :-------------------: | :----------------------------------- |
    |  --service  | string: oneof [fleetmanager, appgateway, aass] |                       | 需要执行安装命令服务组件             |
    | --init-conf |                      path                      |   ./conf/init.yaml    | 配置文件路径，基于该路径进行安装配置 |
    | --start-log |                      path                      | /local/log/{service}/ | 各个服务组件运行时的日志路径         |
    | --help, -h  |                                                |                       | 生成帮助信息                         |

    可生成的帮助信息如下
    ```sh
    $ ./sac-gfm install -h
    the work dir is: /home/gfm
    NAME:
    sac-gfm install - install service info, can help you to install fleetmanager/appgateway/aass

    USAGE:
    sac-gfm install [command options] [arguments...]

    OPTIONS:
    --service value    service name, Options: [fleetmanager, appgateway, aass]
    --init-conf value  init conf using for cipher
    --start-log value  service log path for configuring deamon process
    --help, -h         show help
    ```
+ `start`,`stop`,`restart`命令：
  以start命令为例：你可使用以下命令进行启动fleetmanager服务组件：
  ./sac-gfm start --service fleetmaanger

  可生成的帮助信息如下：
    ```sh
    $ ./sac-gfm start -h
    the work dir is: /home/gfm
    NAME:
    sac-gfm start - start a service, support: [fleetmanager, appgateway, aass]

    USAGE:
    sac-gfm start [command options] [arguments...]

    OPTIONS:
    --service value  service name, Options: [fleetmanager, appgateway, aass]
    --help, -h       show help
    ```
