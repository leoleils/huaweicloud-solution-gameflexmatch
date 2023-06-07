#!/bin/bash

version={version} 								    # appgateway的版本号

# AASS的地址与监听端口(端口默认9091)
export AASS_ADDR={aass_host}:{aass_port}

# Appgateway的数据库相关配置，默认端口3306
export DATABASE_ADDR={mysql_host}:{mysql_port}		# RDS地址与端口
export DATABASE_NAME={mysql_appgateway_database}	# RDS数据库名
export DATABASE_PASSWORD={mysql_password}			# GCM 加密后的RDS密码
export DATABASE_USER={mysql_username}				# RDS用户名

# influxDB的相关配置
export INFLUX_ADDR={influx_host}:{influx_port}		# influxDB的地址与端口
export INFLUX_PASSWORD={influx_password}			# GCM 加密后的influxDB密码
export INFLUX_DBNAME={influx_database}				# 与aass使用相同的数据库
export INFLUX_USER={influx_username}				# influxDB的用户名

# redis配置相关
export REDIS_ADDR={redis_host}:{redis_port}			# redis的地址与端口
export REDIS_PASSWORD={redis_password}				# GCM加密后的redis密码
export REDIS_DATABASE=12							# appgateway使用的redis的数据库

# 本地传输加密相关
export GCM_KEY=**************		# 24位的GCM key
export GCM_NONCE=***************	# 16位的GCM nonce

# 日志相关
export LOG_ROTATE_SIZE=1024			# 单个日志文件存储的最大尺寸(M)
export LOG_BACKUP_COUNT=100			# 单个日志文件存储的最大数量
export LOG_MAX_AGE=7				# 单个日志文件存储的最大时间(天)
export LOG_LEVEL="info"	

# 数据清理配置
export CLEAN_UP=14					# 备份表中的数据清理天数
export BACKUP=3						# 运行表中的异常数据备份到备份表中的时间

# 其他配置
export GATEWAY_ADDR=0.0.0.0:60003
export AUXPROXY_IP_TYPE=publicIP	# appgateway与auxproxy通信的ip类型 publicIP/privateIP
export DEPLOY_MODE=multi-instances	# 多实例multi-instances或单实例singleton


function start_service(){
	cd /home/appgateway/bin/
	./appgateway-${version}
}


function main(){
	start_service
}

main
