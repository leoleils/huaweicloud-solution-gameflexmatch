#! /bin/bash
export FleetManagerAddr=**************      # fleetmanager地址
export FleetManagerPort=***********         # fleetmanager端口
export ProjectId=************************   # 资源账户项目ID
export AliasId=**************************** # 别名id
export BatchCount=50                        # 每秒创建的会话数量
export SleepSecond=1                        # sleep的时间间隔
export UserName=admin                       # console的登录用户
export Password=**************************  # console的登录密码(加密后)
export RunSleep=60                          # 运行*s自行终止一次

function start_service(){
        cd /home
        ./auto-create-session -fleetmanager-addr $FleetManagerAddr -fleetmanager-port $FleetManagerPort -project-id $ProjectId -alias-id $AliasId -batch-count $BatchCount -sleep-second $SleepSecond -username $UserName -password $Password -process-run-sleep-second $RunSleep
}


function main(){
        start_service
}

main