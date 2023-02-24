#!/bin/bash

version={version} 

# AASS的地址与监听端口(默认9091).
export AASS_ADDR={aass_host}:{aass_port}

# Appgateway的数据库相关配置，默认端口3306
export DATABASE_ADDR={mysql_host}
export DATABASE_NAME={mysql_appgateway_database}
export DATABASE_PASSWORD={mysql_password}			# GCM 加密后的RDS密码
export DATABASE_USER={mysql_username}

# influxDB的相关配置
export INFLUX_ADDR={influx_host}:{influx_port}
export INFLUX_PASSWORD={influx_password}			# GCM 加密后的influxDB密码
export INFLUX_DBNAME={influx_database}				# 与aass使用相同的数据库
export INFLUX_USER={influx_username}

# 本地传输加密相关
export GCM_KEY=**************		# 24位的GCM key
export GCM_NONCE=***************	# 16位的GCM nonce

# 日志相关
export LOG_ROTATE_SIZE=100			# 单个日志文件存储的最大尺寸(M)
export LOG_BACKUP_COUNT=100			# 单个日志文件存储的最大数量
export LOG_MAX_AGE=7				# 单个日志文件存储的最大时间(天)	

# 其他配置
export GATEWAY_ADDR=0.0.0.0:60003

function start_service(){
	cd /home/appgateway/bin/
	./appgateway-${version} -gateway-addrress $GATEWAY_ADDR -database-address $DATABASE_ADDR -database-name $DATABASE_NAME -database-user-name $DATABASE_USER -database-password $DATABASE_PASSWORD --aass-address $AASS_ADDR -influx-address $INFLUX_ADDR -influx-username $INFLUX_USER -influx-password $INFLUX_PASSWORD -influx-dbname $INFLUX_DBNAME -gcm-key $GCM_KEY -gcm-nonce $GCM_NONCE -log-rotate-size $LOG_ROTATE_SIZE -log-backup-count $LOG_BACKUP_COUNT -log-max-age $LOG_MAX_AGE
}


function main(){
	start_service
}

main
