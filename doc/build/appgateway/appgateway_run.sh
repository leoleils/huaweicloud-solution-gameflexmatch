#!/bin/bash

version={version} 

# AASS的地址与监听端口(默认9091).
export AASS_ADDR={aass_host}:{aass_port}

# Appgateway的数据库相关配置，默认端口3306
export DATABASE_ADDR={mysql_host}
export DATABASE_NAME={mysql_appgateway_database}
export DATABASE_PASSWORD={mysql_password}
export DATABASE_USER={mysql_username}

# influxDB的相关配置
export INFLUX_ADDR={influx_host}:{influx_port}
export INFLUX_PASSWORD={influx_password}
export INFLUX_DBNAME={influx_appgateway_database}
export INFLUX_USER={influx_username}

# 其他配置
export GATEWAY_ADDR=0.0.0.0:60003
export AUXPROXY_IP_MODE=public

function start_service(){
	cd /home/appgateway/bin/
	nohup ./appgateway-${version} -gateway-addrress $GATEWAY_ADDR -database-address $DATABASE_ADDR -database-name $DATABASE_NAME -database-user-name $DATABASE_USER -database-password $DATABASE_PASSWORD --aass-address $AASS_ADDR --auxproxy-ip-mode $AUXPROXY_IP_MODE -influx-address $INFLUX_ADDR -influx-username $INFLUX_USER -influx-password $INFLUX_PASSWORD -influx-dbname $INFLUX_DBNAME &

}


function main(){
	start_service
}

main
