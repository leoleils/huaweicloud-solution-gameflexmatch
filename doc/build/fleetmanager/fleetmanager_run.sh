#!/bin/bash

version={version}

# FleetManager的数据库相关配置 
export MYSQL_ADDRESS={mysql_host}:{mysql_port}
export MYSQL_DB_NAME={mysql_fleetmanager_database}
export MYSQL_PASSWORD={mysql_pwd}
export MYSQL_USER={mysql_user}
export MYSQL_CHARSET=utf8

# 服务账号的相关配置
export SERVICE_AK={ak}
export SERVICE_SK={sk}
export SERVICE_DOMAIN_ID={managerUser_domain_id}

# 资源配置
export DEFAULT_FLEET_BANDWIDTH=1
export DEFAULT_FLEET_MAX_SESSION_NUM_PER_PROCESS=1
export DEFAULT_FLEET_NEW_SESSION_NUM_PER_CREATOR=2
export DEFAULT_FLEET_POLICY_PERIOD=3
export DEFAULT_FLEET_PROTECT_POLICY=TIME_LIMIT_PROTECTION
export DEFAULT_FLEET_PROTECT_TIME_LIMIT=5
export DEFAULT_FLEET_REGION=cn-north-4
export DEFAULT_FLEET_SESSION_TIMEOUT_SECONDS=600
export DEFAULT_FLEET_SPECIFICATION=scase.standard.4u8g
export FLEET_DISK_TYPE=SYS
export FLEET_DISK_SIZE=40
export FLEET_VOLUME_TYPE=SAS
export FLEET_EIP_SHARE_TYPE=PER
export ENTERPRISE_PROJECT={enterprise_project}

# 服务其他配置
export AASS_ENABLE_HMAC=false
export APPGATEWAY_ENABLE_HMAC=false
export ENABLE_HTTP=flase
export ENABLE_HTTPS=true
export ENABLE_TOKEN_CHECK=false
export HTTPS_CERT_FILE=/home/tlsSecret/tls.crt
export HTTPS_KEY_FILE=/home/tlsSecret/tls.key
export CONFIG_FILE=/home/fleetmanager/configmap/service_config.json
export FLEET_QUOTA=100
export LOG_ROTATE_SIZE=50
export LOG_BACKUP_COUNT=7
export MAX_PROCESS_NUM_PER_FLEET=50
export REGION={region}
export SUPPORT_REGIONS={region}
export WEB_HTTPS_ADDR=0.0.0.0
export WEB_HTTPS_PORT=31002

# Fleetmanager Redis相关配置
export REDIS_ADDRESS={redis_host}:{redis_port}
export REDIS_PASSWORD={redis_password}
export REDIS_MAX_CONN={redis_max_connection}

# 登录功能相关配置
export SESSION_LIFETIME=43200
export JWTKEY={jwt_token_generate_key}
export JWT_TOKEN_LIFETIME=7200

# 是否开启登录验证功能
export ENABLE_AUTHORIZED=false

function start_service(){
	cd /home/bin/
	./fleetmanager-${version}
}

function main(){
	start_service
}

main
