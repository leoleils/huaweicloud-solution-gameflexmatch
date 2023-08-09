#!/bin/bash

version={version}										# fleetmanager的版本

# FleetManager的数据库相关配置 
export MYSQL_ADDRESS={mysql_host}:{mysql_port}			# 数据库地址与端口
export MYSQL_DB_NAME={mysql_fleetmanager_database}		# 数据库名
export MYSQL_PASSWORD={mysql_pwd}						# 加密后的数据库用户密码
export MYSQL_USER={mysql_user}							# 数据库用户名
export MYSQL_CHARSET=utf8								# 数据库中数据存储格式

# Fleetmanager Redis相关配置
export REDIS_ADDRESS={redis_host}:{redis_port}			# redis地址与端口
export REDIS_PASSWORD={redis_password}					# 加密后的redis的密码
export REDIS_DB=10										# fleetmanager使用redis的数据库

# 服务账号的相关配置
export SERVICE_AK={ak}									# 管理账号使用GCM加密后的ak
export SERVICE_SK={sk}									# 管理账号使用GCM加密后的sk
export SERVICE_DOMAIN_ID={managerUser_domain_id}		# 管理账号的domain_id

# 资源配置
export DEFAULT_FLEET_BANDWIDTH=1						# 默认带宽
export DEFAULT_FLEET_MAX_SESSION_NUM_PER_PROCESS=1		# 单个进程默认最大会话数量
export MAX_PROCESS_NUM_PER_FLEET=50						# 每个instance的进程数量最大限制
export DEFAULT_FLEET_PROTECT_POLICY=TIME_LIMIT_PROTECTION	# 默认的会话保护策略
export DEFAULT_FLEET_PROTECT_TIME_LIMIT=5				# 默认的关于会话保护时间
export DEFAULT_FLEET_SESSION_TIMEOUT_SECONDS=600		# 默认会话激活超时时间
export DEFAULT_FLEET_SPECIFICATION=scase.standard.4u8g	# 默认的实例规格
export DEFAULT_SCALING_INTERVAL=10						# 默认的弹性伸缩时间间隔
export FLEET_DISK_TYPE=SYS								# fleet磁盘类型
export FLEET_DISK_SIZE=40								# fleet磁盘尺寸(G)
export FLEET_VOLUME_TYPE=SAS							# fleet卷类型
export FLEET_EIP_SHARE_TYPE=PER							# fleet的eip类型
export ENTERPRISE_PROJECT={enterprise_project}			# 默认企业项目

# 签名秘钥校验
export ENABLE_TOKEN_CHECK=false							# 是否开启华为云IAM token校验
export HTTPS_CERT_FILE=/home/tlsSecret/tls.crt			# https签名文件路径
export HTTPS_KEY_FILE=/home/tlsSecret/tls.key			# https秘钥路径
export RSA_PUBLIC_FILE=/home/tlsSecret/rsa_public.pem		# 网络敏感信息加密传输的RSA公钥
export RSA_PRIVATE_FILE=/home/tlsSecret/rsa_private.pem		# 网络敏感信息加密传输的RSA私钥
export GCM_KEY=**************							# 本地数据加密的GCM 24位key
export GCM_NONCE=********************					# 本地数据加密的GCM 16位Nonce

# 应用包打包相关
export DEFAULT_IMAGE_REF="CentOS 7.2 64bit" # 用于镜像打包的ECS默认操作系统
export DEFAULT_SCRIPT_PATH=gamebounce/image_env.sh # 镜像环境配置脚本 OBS路径 桶名/对象名
export DEFAULT_AUXPROXY_PATH=gamebounce/auxproxy.zip # AuxProxy组件压缩包 OBS路径 用于镜像环境配置 桶名/对象名
export IMAGE_DISK_SIZE=40 # 用于镜像打包的ECS默认磁盘大小
export DEFAULT_IMAGE_FLAVOR=s6.large.2 # 用于镜像打包的ECS默认规格
export PROFILE_STORAGE_REGION={region} # 管理账号存放auxproxy及部署脚本的区域
export DEFAULT_BUILD_BANDWIDTH=10 # 用于镜像打包的ECS绑定的EIP默认带宽

# 服务其他配置
export AASS_ENABLE_HMAC=false							# aass是否开启hmac验证
export APPGATEWAY_ENABLE_HMAC=false						# appgateway是否开启hmac验证
export ENABLE_HTTP=flase								# 是否http传输(暂不支持)
export ENABLE_HTTPS=true								# 是否https传输
export CONFIG_FILE=/home/fleetmanager/configmap/service_config.json	# 配置文件路径
export FLEET_QUOTA=100									# fleet的数据限制
export LOG_ROTATE_SIZE=1024								# 每个log文件的尺寸(MB)
export LOG_BACKUP_COUNT=100								# 每个日志文件储存的最大数量
export LOG_MAX_AGE=7									# 每个日志存储的最大时间					
export REGION={region}									# fleet的region，默认region会保证与该region相同
export SUPPORT_REGIONS={region}							# 支持的region
export WEB_HTTPS_ADDR=0.0.0.0							# fleetmanager的启动地址
export WEB_HTTPS_PORT=31002								# fleetmanager的启动端口

# 登录功能相关配置
export JWTKEY={jwt_token_generate_key}					# 登录会话token生成的key
export JWT_TOKEN_LIFETIME=7200							# 登录会话token的有效时间(s)						
export DEFAULT_LOGIN_PASSWORD={default_login_password}  # 登录默认密码设置

export WORKFLOW_PATH=/home/fleetmanager/bin/conf/workflow/ # workflow文件所在文件夹路径，注意以/结尾
export FLEET_SERVER_SESSION_BACKUP_DAYS=1				# 数据表自动清理的配置(day)

function start_service(){
	cd /home/fleetmanager/bin/
	./fleetmanager-${version}
}

function main(){
	start_service
}

main
