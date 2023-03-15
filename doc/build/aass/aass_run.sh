#!/bin/bash

version={version}							# aass的版本名

# influxDB的相关配置
export INFLUX_ADDRESS=https://{influx_host}:{influx_port}	# influxDB的地址与端口
export INFLUX_PASSWORD={influx_pwd}			# 加密后的influxDB密码
export INFLUX_DATABASE={influx_database}	# influxDB数据库，与appgateway配置相同
export INFLUX_USER={influx_username}

# AASS的数据库相关配置
export MYSQL_ADDRESS={mysql_host}:{mysql_port}		# 数据库的地址与端口
export MYSQL_DB_NAME={mysql_aass_database}
export MYSQL_PASSWORD={mysql_pwd} 			# 加密后的mysql密码
export MYSQL_USER={mysql_username}
export MYSQL_CHARSET=utf8

# 服务账号的相关配置
export SERVICE_AK={ak}						# 加密后的管理用户ak
export SERVICE_SK={sk}						# 加密后的管理用户sk
export SERVICE_DOMAIN_ID={managerUser_domain_id}

# 其他配置
# endpoint是各云服务在该region下的终端节点
# 如myhuaweicloud.com/ulanqab.huawei.com
export CLOUD_CLIENT_REGION={region}
export CLOUD_CLIENT_AS_ENDPOINT=https://as.{region}.{endpoint}
export CLOUD_CLIENT_ECS_ENDPOINT=https://ecs.{region}.{endpoint}
export CLOUD_CLIENT_IAM_ENDPOINT=https://iam.{region}.{endpoint}
export CLOUD_CLIENT_LTS_ENDPOINT=https://lts.{region}.{endpoint}
export CONFIG_FILE=/home/aass/configmap/service_config.json
export HTTPS_CERT_FILE=/home/tlsSecret/tls.crt
export HTTPS_KEY_FILE=/home/tlsSecret/tls.key
export SERVER_HMAC_CONF_FILE=/home/aass/configmap/server_hmac_conf.json
export HTTPS_LISTEN_ADDR=0.0.0.0

export GCM_KEY=******************     # GCM加解密所用的24位key
export GCM_NONCE=**************		# GCM加解密所用的16位Nonce

# 日志配置
export LOG_ROTATE_SIZE=100				# 单个日志文件的最大尺寸(M)		
export LOG_BACKUP_COUNT=100				# 日志文件的的最大数量
export LOG_MAX_AGE=7					# 单个日志文件存储的最大时间

function start_service(){
	cd /home/bin/
	./aass-${version}
}

function main(){
	start_service
}

main


