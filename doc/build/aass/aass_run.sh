#!/bin/bash

version={version}

# influxDB的相关配置
export INFLUX_ADDRESS=https://{influx_host}:{influx_port}
export INFLUX_PASSWORD={influx_pwd}
export INFLUX_DATABASE={influx_aass_database}
export INFLUX_USER={influx_username}

# AASS的数据库相关配置
export MYSQL_ADDRESS={mysql_host}:{mysql_port}
export MYSQL_DB_NAME={mysql_aass_database}
export MYSQL_PASSWORD={mysql_pwd}
export MYSQL_USER={mysql_username}
export MYSQL_CHARSET=utf8

# 服务账号的相关配置
export SERVICE_AK={ak}
export SERVICE_SK={sk}
export SERVICE_DOMAIN_ID={managerUser_domain_id}

# 其他配置
# endpoint是各云服务在该region下的终端节点
# 如myhuaweicloud.com/ulanqab.huawei.com
export CLOUD_CLIENT_REGION={region}
export CLOUD_CLIENT_AS_ENDPOINT=https://as.{region}.{endpoint}
export CLOUD_CLIENT_ECS_ENDPOINT=https://ecs.{region}.{endpoint}
export CLOUD_CLIENT_IAM_ENDPOINT=https://iam.{region}.{endpoint}
export CONFIG_FILE=/home/aass/configmap/service_config.json
export HTTPS_CERT_FILE=/home/tlsSecret/tls.crt
export HTTPS_KEY_FILE=/home/tlsSecret/tls.key
export SERVER_HMAC_CONF_FILE=/home/aass/configmap/server_hmac_conf.json
export HTTPS_LISTEN_ADDR=0.0.0.0


function start_service(){
	cd /home/bin/
	nohup ./aass-${version} &
}

function main(){
	start_service
}

main


