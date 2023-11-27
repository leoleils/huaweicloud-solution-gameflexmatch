# Brief Introduction #

`GameFlexMatch`is a service hosting solution that consists of four service components (`Fleetmanager`/`AppGateway`/`AASS`/`AuxProxy`), which can implement application hosting, elastic scaling of resources required by hosting applications, resource scheduling and management of application processes, and gray release of applications.`region`During deployment, users can access the network nearby, reducing latency and implementing cross-region DR for service resources. It helps developers quickly build a stable and low-latency multiplayer game deployment environment and saves a lot of O&M costs.`Unreal`,`Unity`The engine,`C#`,`C++`And also the`gRPC`any language supported`server`Deploy and run the framework.

# Logical Architecture #

The GameFlexMatch platform consists of five service components:

 *  [FleetManager](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-fleetmanager)	\: Globally deploys and manages application processes, supports configuration of dynamic deployment policies, optimizes application distribution based on costs or latency, configures auto scaling policies, manages server sessions, client sessions, and application packages, and supports gray release of server applications.
 *  [AppGateway](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-appgateway)	\: Manages application processes, sessions, and client connections.`AuxProxy`Obtain application process information through communication, and schedule process resources.
 *  [AASS](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-aass)	\: manages and executes AS groups and AS policies, monitors application resources on the server, and invokes HUAWEI CLOUD.`AS`(Auto Scaling) Implements elastic resource scaling.
 *  [AuxProxy](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-auxproxy)	\: Automatically starts the instance after capacity expansion, and creates application processes, reports process status, and communicates with application processes.
 *  [Console](https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-console)	\: O&M platform, used for monitoring`GameFlexMatch`Running status and O&M management of the`GameFlexMatch`of the`fleet`2. Application packages and user information

# Creating Resources #

### Creating an ECS ###

Purchase four ECSs, deploy the GameFlexMatch service component on each ECS, and deploy the management plane resources in the same VPC. The following test specifications are available. Select the specifications as required.

|        | Resource Specifications | Service component | Listening Port | Whether EIP must be configured |
| ------ | ----------------------- | ----------------- | -------------- | ------------------------------ |
| ECS-01 | 2vCPUs/4GB              | appgateway        | 60003          | Y                              |
| ECS-02 | 2vCPUs/4GB              | aass              | 9091           | Y                              |
| ECS-03 | 2vCPUs/4GB              | fleetmanager      | 31002          | Y                              |
| ECS-04 | 2vCPUs/4GB              | console           | 80             | Y                              |

Add the VPC network segment of the ECS in the inbound direction to the security group rule so that service components can communicate with each other.

### Creating a database ###

Purchase RDS for HUAWEI CLOUD, select the specifications as required, and use the default port number 3306 to create three databases: appgateway, aass, and fleetmanager.

### Creating InfluxDB ###

Purchase InfluxDB, enable the SSL connection, and use the default certificate to create a database (aass/appgateway) for service components. Aass and appgateway share the same InfluxDB database.

### Creating a Redis ###

Purchase the Redis service component and configure the access port and password. The three service components share the same Redis database.

### Creating InfluxDB ###

Purchase GaussDB (for Influx), enable the SSL connection, and use the default certificate to create a database (aass/appgateway) for the service component. Aass and appgateway share the same InfluxDB database.

### Creating a Key Pair ###

On the Key Pair Management page of DEW, create a private key pair (only for the IAM account) or an account key pair (recommended) for logging in to the AS instance.

# Service installation #

## Prerequisite ##

 *  Operating system: CentOS 7.6 64-bit
 *  Minimum configuration: 2 vCPUs and 4 GB memory
 *  Install git.
    
    ```
    yum -y install git
    ```

## Front-end components ##

### Environment installation ###

#### Installing the Node.js ####

```
wget https://nodejs.org/download/release/v16.13.1/node-v16.13.1-linux-x64.tar.gz
tar xf node-v16.13.1-linux-x64.tar.gz
mv node-v16.13.1-linux-x64 /usr/local/
```

Creating and Modifying the Node.sh File

```
vi /etc/profile.d/node.sh
```

Setting Environment Variables

```
export NODE_HOME=/usr/local/node-v16.13.1-linux-x64
export PATH=${NODE_HOME}/bin:$PATH
```

Run the script for the environment variables to take effect.

```
chmod +x /etc/profile.d/node.sh
source /etc/profile.d/node.sh
```

#### Installing the vue ####

```
npm install -g @vue/cli
```

#### Installing the Compilation Tool and Library Files ####

```
yum -y install make zlib zlib-devel gcc-c++ libtool  openssl openssl-devel
```

Downloading the PCRE Installation Package

```
cd /usr/local/src/ || exit
wget http://downloads.sourceforge.net/project/pcre/pcre/8.35/pcre-8.35.tar.gz
tar zxvf pcre-8.35.tar.gz
```

Compilation and installation

```
cd pcre-8.35 || exit
./configure
make && make install
```

#### Installing the Nginx ####

```
cd /usr/local/src/ || exit
wget http://nginx.org/download/nginx-1.7.8.tar.gz
tar zxvf nginx-1.7.8.tar.gz
cd nginx-1.7.8 || exit
```

Compilation and installation to /usr/local/webserver/nginx

```
./configure --prefix=/usr/local/webserver/nginx --with-http_stub_status_module --with-http_ssl_module --with-pcre=/usr/local/src/pcre-8.35
make
make install
```

Replace /usr/local/webserver/nginx/conf/nginx.conf with the following information and proxy_pass with the IP address of the backend server.

```
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  300;
    client_max_body_size 6g;
    server {
        listen       80;
        server_name  localhost;

        location /api/ {
            #Change the following IP address to the IP address of the backend server and forward the requests in the /api directory to the actual backend server.
            proxy_pass https://127.0. 0.1:8080 /;
        }

        location / {
                root html;
                try_files $uri /index.html;  #try_files: Check files. $uri: indicates the path of the monitored file. /index.html: The file does not have a new redirection path.
                index index.html;
        }

        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}
```

### Packaged deployment ###

Pull front-end code

```
cd /usr/local
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-console.git
cd huaweicloud-solution-gameflexmatch-console
```

Change the password encryption public key configuration file code root directory /src/api/crypto.ts in line 42

```
export function encryptedData(data: string) {
    const encryptor = new JSEncrypt()
    let key = `XXXXXXXXX` //Set the public key.
    encryptor.setPublicKey(key)
    return encryptor.encrypt(data).toString()
}
```

Run the npm install command in the root directory of the code to install the dependencies required by the project.

```
npm install
```

Install the vue internationalization plug-in in the root directory.

```
npm install --save vue-i18n@next
```

Run the npm run build command in the root directory of the code to compile and package the project to the dist folder in the root directory.

```
npm run build
```

Copy all files in the dist directory to the /usr/local/webserver/nginx/html root directory of the Nginx website.

Configure the Nginx to automatically start upon system startup and create the nginx.service file in the /lib/systemd/system/ directory.

```
vi /lib/systemd/system/nginx.service
```

Add the following information to the file:

```
[Unit]
Description=nginx service
After=network.target
[Service]
Type=forking
ExecStart=/usr/local/webserver/nginx/sbin/nginx
ExecReload=/usr/local/webserver/nginx/sbin/nginx -s reload
ExecStop=/usr/local/webserver/nginx/sbin/nginx -s quit
PrivateTmp=true
[Install]
WantedBy=multi-user.target
```

Setting the execute permission on a file

```
chmod a+x /lib/systemd/system/nginx.service
```

Set the automatic startup.

```
systemctl enable nginx.service
```

Starting the Nginx Service

```
systemctl start nginx.service
```

Enter the IP address bound to the ECS in the address box of the browser to access the GameFlexMatch frontend page, or enter http://\{ipv4\}:80 in the address box.

## Back-end components ##

### Environment installation ###

 *  Version: Go1.16 or later
 *  Install go.

```
yum install golang
#Set the operating system for the compiled executable file
go env -w GOOS=linux
#Configuring the Go Proxy
go env -w GO111MODULE=on
go env -w GOPROXY=https://repo.huaweicloud.com/repository/goproxy/
go env -w GONOSUMDB=*
```

### File compilation ###

Download the source code to the local computer and compile the source code.`linux`This is an executable binary file. In the following steps, the variables in \{version\} need to be changed as required.

```
#1. fleetmanager
cd ~/huaweicloud-solution-gameflexmatch-fleetmanager
#Downloading the Dependency Package
go mod tidy
go build ./main.go
#Modify File Name
mv main fleetmanager-{version}

#2. appgateway
cd ~/huaweicloud-solution-gameflexmatch-appgateway
go mod tidy
go build ./cmd/application_gateway.go
#Modify File Name
mv application_gateway appgateway-{version}

#3. aass
cd ~/huaweicloud-solution-gameflexmatch-aass
go mod tidy
go build ./cmd/application-auto-scaling-service/application_auto_scaling_service.go
#Modify File Name
mv application_auto_scaling_service aass-{version}

#4. auxproxy
cd ~/huaweicloud-solution-gameflexmatch-auxproxy
go mod tidy
go build ./cmd/auxproxy.go
```

### **Certificate Preparation** ###

In the`linux`Passed in the system`openssl`Obtained from signature certificate and can be used on any server.`ECS`The three service components use the same self signature certificate as follows:

 *  Get`https`signature certificate

```
#1. Create the tlsSecret folder.
mkdir -p /home/tlsSecret
cd /home/tlsSecret
#2. Generate a private key tls.key.
openssl genrsa -out tls.key 3072
#3. Use the private key to generate a csr and check the csr.
openssl req -new -key tls.key -out tls.csr
#The preceding steps require the following information, which can be filled in based on the actual situation, such as
#There are quite a few fields but you can leave some blank
#For some fields there will be a default value,
#If you enter'. ', the field will be left blank.

#Country Name (2 letter code) [XX]:China
#State or Province Name (full name) []:GuangDong
#Locality Name (eg, city) [Default City]:ShenZhen
#Organization Name (eg, company) [Default Company Ltd]:Huawei
#Organizational Unit Name (eg, section) []:Cloud
#Common Name (eg, your name or your server's hostname) []:GameFlexMatch
#Email Address []:GameFlexMatch@huawei.com

#Please enter the following'extra' attributes
#to be sent with your certificate request
#A challenge password []:********
#An optional company name []:Huawei

openssl req -in tls.csr -text
#4. Generated from signature certificate tls.crt and viewed
openssl x509 -req -days 365 -in tls.csr -signkey tls.key -out tls.crt
openssl x509 -in tls.crt -text
```

 *  Obtain the RSA asymmetrically encrypted public and private keys transmitted over the network, and encrypt and decrypt sensitive user data such as login passwords and cloud resources.

```
#1. Create an RSA private key with a length of 1024 or 2048 characters.
cd /home/tlsSecret
openssl genrsa -out rsa_private.pem 2048
#2. Generate a public key based on the private key.
openssl rsa -in rsa_private.pem -pubout -out rsa_public.pem
```

### **Service component installation** ###

#### Installing the AppGateway Service Component ####

```
#1. Log in to ECS-01, create /home/tlsSecret, and upload the generated tls.crt and tls.key files as well as the public and private keys rsa_private.pem and rsa_public.pem for asymmetric RSA encryption.
mkdir -p /home/tlsSecret
#2. Create the /home/appgateway/conf/hmac folder,
mkdir -p /home/appgateway/conf/hmac
#Upload the client_hmac_conf.json and server_hmac_conf.json files to the hmac folder. For details, see. 

#doc/build/appgateway/client_hmac_conf.json
#doc/build/appgateway/server_hmac_conf.json

#3. Create the bin directory, upload the executable binary file and startup script, and modify the permission.
mkdir -p /home/appgateway/bin
cd /home/appgateway/bin
#Upload the generated binary file appgateway-{version} and the startup script appgateway_run.sh, and modify related configurations. For details, see doc/build/appgateway/appgateway_run.sh.
#Modifying Permission
chmod 750 appgateway-{version}
chmod 750 appgateway_run.sh

#4. Run the startup script after the configuration is complete.
./appgateway_run.sh

#5. Check whether the execution is successful.
ps -aux | grep appgateway

#6. Disable the process and configure automatic startup.
#Create appgateway.service in /etc/systemd/system.
#For an example of appgateway.service, see /doc/build/appgateway.
#Start appgateway.service to ensure that the process starts automatically.
systemctl enable appgateway.service
systemctl start appgateway.service

#7. Check whether the test is successful.
ps -aux | grep appgateway
#8. If the system runs properly, the deployment is successful.
```

#### Installing the ASS Service Components ####

```
#1. Log in to ECS-02 and create the tlsSecret folder.
mkdir -p /home/tlsSecret
#2. Upload the generated tls.crt and tls.key files to the tlsSecret folder.
#3. Create the /home/aass folder.
mkdir -p /home/aass/configmap
#4. Upload the server_hmac_conf.json and service_config.json files to the configmap directory and modify related configurations. For details, see.
#doc/build/aass/server_hmac_conf.json
#doc/build/aass/service_config.json

#5. Create the /home/aass/bin folder.
mkdir -p /home/aass/bin
#6. Upload the binary executable file aass-{version} and the executable script aass_run.sh of the Aass to the bin directory, modify related configurations, and modify the file permission.
cd /home/aass/bin
chmod 750 aass-{version}
chmod 750 aass_run.sh

#7. Run the command and check whether the command is successfully executed.
sh ./aass_run.sh
ps -aux | grep aass

#8. Press Ctrl+C to stop the process and configure automatic startup.
#Create aass.service in /etc/systemd/system.
#For an example of aass.service, see /doc/build/aass.
#Starting the aass.service Process to Be Automatically Started
systemctl enable aass.service
systemctl start aass.service

#9. Check whether the test is successful.
ps -aux | grep aass
#10. If the system runs properly, the deployment is successful.
```

#### Installing the FleetManager Service Component ####

```
#1. Log in to ECS-03, create the tlsSecret folder, and upload the generated tls.crt and tls.key files.
mkdir -p /home/tlsSecret
#2. Create the /home/configmap folder.
mkdir -p /home/fleetmanager/configmap
#3. Upload the server_config.json file to the configmap directory and modify related configurations.
#For details, see doc/build/fleetmanager/server_config.json.

#4. Create folder /home/fleetmanager/bin/conf/workflow
mkdir -p /home/fleetmanager/bin/conf/workflow

#5. Upload create_fleet_workflow.json, delete_fleet_workflow.json, and create_build_image_workflow.json. For details, see.
#doc/build/fleetmanager/create_fleet_workflow.json
#doc/build/fleetmanager/delete_fleet_workflow.json
#doc/build/fleetmanager/create_build_image_workflow.json

#6. Upload the binary executable file fleetmanager-{version} of the fleetmanager.
#Upload the startup script fleetmanager_run.sh to the bin folder and modify related configuration and file permission.
cd /home/fleetmanager/bin
chmod 750 fleetmanager-{version}
chmod 750 fleetmanager_run.sh

#7. Start the script and verify that the script is successful.
sh ./fleetmanager_run.sh
ps -aux | grep fleetmanager

#8. Disable the process and configure automatic startup.
#Create fleetmanager.service in /etc/systemd/system.
#For an example of fleetmanager.service, see /doc/build/fleetmanager.
#Start fleetmanager.service to ensure that the process is automatically started.
systemctl enable fleetmanager.service
systemctl start fleetmanager.service

#9. Check whether the test is successful.
ps -aux | grep fleetmanager
#10. If the system runs properly, the deployment is successful.
```

## **Other instructions** ##

 *  For details about PU deployment, see.[doc/build/console.md](/doc/build/console.md)	
 *  For details about how to create an application image, see.[doc/build/make-image-guide.md](/doc/build/make-image-guide.md)	
 *  For details about the parameters required during deployment, see.[doc/build/param-annotation.md](/doc/build/param-annotation.md)	
 *  For details about how to use the platform user management module, see.[doc/build/user-management.md](/doc/build/user-management.md)	
 *  For details about the console user guide, see.`/doc/user-guide`
 *  For details about the version update history, see.`doc/version/`

## Auxiliary Tools ##

1.  Encryption tool: provides tools for encrypting and decrypting sensitive data using GCM and RSA. For details, see.`/tools/cipher`
2.  Peer Connection Tools: provides two`VPC`For details about how to create a VPC peering connection, see.`/tools/vpc-peering`

