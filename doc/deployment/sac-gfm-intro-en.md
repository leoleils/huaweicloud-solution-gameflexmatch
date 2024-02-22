# Introduction of the sac-gfm Script

## Function Introduction

`sac-gfm` is a binary file that helps you deploy and manage the GFM. Its main functions are as follows:

- [X] Generate key of the encryption and decryption and related key files, and configure the files in the configuration file.
- [X] Analyze and process the configuration files, generate the corresponding configuration files for each service component, generate the startup script, and enabling automatic startup
- [X] Manages the GFM and supports starting, stopping, and restarting service components.

## Command and Parameter Introduction

The usage of the is as follows: `sac-gfm [global options] command [command options] [arguments...]`

You can also just run `sac-gfm` or the `sac-gfm -h` to view the command help information.

```
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

| commanded | Description                                                                                                                                                                                 |
| :-------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
|  cipher   | Generates encryption and decryption keys, and can help you encrypt or decrypt sensitive information.                                                                                        |
|  install  | Based on the configured configuration file, the system automatically generates configuration files for each service component and generates startup scripts to configure automatic startup. |
|   start   | To help you start a service                                                                                                                                                                 |
|   stop    | To help you stop a service                                                                                                                                                                  |
|  restart  | Help you restart a service                                                                                                                                                                  |
|   help    | Output help documents.                                                                                                                                                                      |

**The usage and example of each command are as follows:**

+ cipher command: `./sac-gfm cipher [command options] [arguments...]`, the following commands are supported:

|  commanded  |              Parameters               |  Default value   | Description                                                                                                                                                                                                                                                                                                                                                                      |
| :---------: | :-----------------------------------: | :--------------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|   --mode    |   string: oneof \[decode, encode\]    |      encode      | Encrypts or decrypts sensitive information.                                                                                                                                                                                                                                                                                                                                      |
|  --method   |      string: oneof \[gcm, rsa\]       |       gcm        | Methods for encrypting or decrypting sensitive information                                                                                                                                                                                                                                                                                                                       |
|   --text    |                string                 |                  | Characters used for encryption or decryption                                                                                                                                                                                                                                                                                                                                     |
| --init-conf |                 path                  | ./conf/init.yaml | Path of the configuration file. When the key information needs to be generated, the generated information is automatically synchronized to the configuration file.                                                                                                                                                                                                               |
|  --create   | string: one of [gcm, rsa, tls, jwt] |                  | If the parameter is set to `gcm` or `jwt`, the generated sensitive information is synchronized to the configuration file corresponding to `--init-conf`. If the parameter is set to `rsa` or `tls`, A configuration file is generated in the directory where the configuration file corresponding to `--init-conf` is located and is synchronized to the corresponding configuration file. |
| --help, -h  |                                       |                  | Generate Help Information                                                                                                                                                                                                                                                                                                                                                        |

The following help information can be generated:
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

+ install command:`sac-gfm install [command options] [arguments...]`

|    命令     |                      参数                      |        默认值         | 说明                                                                       |
| :---------: | :--------------------------------------------: | :-------------------: | :------------------------------------------------------------------------- |
|  --service  | string: oneof [fleetmanager, appgateway, aass] |                       | The command service component needs to be installed.                       |
| --init-conf |                      path                      |   ./conf/init.yaml    | Configuration file path, which is used for installation and configuration. |
| --start-log |                      path                      | /local/log/{service}/ | Log path of each service component                                         |
| --help, -h  |                                                |                       | Generating help information                                                |


   The following help information can be generated:
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

+ start, stop, restart command: Use the start command as an example. You can run the following command to start the FleetManager service component: ./sac-gfm start --service fleetmaanger
    
   The following help information can be generated:
    
    ```
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

