#!/bin/bash
# 阿里云适配版本的VM镜像构建脚本
REGION=$1
APP_URL=$2
AUX_URL=$3
LTS_IP=$4
APP_FILE=$5
AUX_FILE=$6
APP_NAME=$7
USER=$8
GROUP=$9
PWD=${10}

# 安装必要的解压工具
if command -v yum &> /dev/null; then
    yum install -y unzip wget
elif command -v apt-get &> /dev/null; then
    apt-get update && apt-get install -y unzip wget
fi

# download rarlab
mkdir /tmp/rar
cd /tmp/rar
wget --no-check-certificate https://www.rarlab.com/rar/rarlinux-x64-612.tar.gz
tar -xzvf rarlinux-x64-612.tar.gz
cd rar
make

useradd ${USER}
groupadd ${GROUP}
usermod -g ${GROUP} ${USER}
echo ${USER}:${PWD} | chpasswd

# download app
mkdir -p /local/app/${APP_NAME}
cd /local/app/${APP_NAME}
wget ${APP_URL} -O ${APP_FILE}
if [ "${APP_FILE##*.}" = "zip" ]; then
	unzip -o ${APP_FILE}
fi
if [ "${APP_FILE##*.}" = "rar" ]; then
	unrar e  ${APP_FILE}
fi
chown -R ${USER}:${GROUP} /local/app/${APP_NAME}
chmod -R 750 *

# download auxproxy
mkdir -p /etc/auxproxy
cd /etc/auxproxy
wget ${AUX_URL} -O ${AUX_FILE}
if [ "${AUX_FILE##*.}" = "zip" ]; then
	unzip -o ${AUX_FILE}
fi
if [ "${AUX_FILE##*.}" = "rar" ]; then
	unrar e  ${AUX_FILE}
fi
# 处理压缩包内有子目录的情况，将文件移动到 /etc/auxproxy/ 目录下
if [ -d "/etc/auxproxy/auxproxy_pkg" ]; then
	mv /etc/auxproxy/auxproxy_pkg/* /etc/auxproxy/
	rmdir /etc/auxproxy/auxproxy_pkg
fi
chown -R ${USER}:${GROUP} /etc/auxproxy
chmod -R 750 *

cat > /etc/systemd/system/auxproxy.service <<- EOF
[Unit]
Description=auxproxy.service

[Service]
Type=simple
User=${USER}
Group=${GROUP}
ExecStart=/etc/auxproxy/auxproxy-start.sh
Restart=always
RestartSec=30
TimeoutSec=0
KillMode=process
StandardOutput=null

[Install]
WantedBy=multi-user.target

EOF

mkdir -p /etc/auxproxy/security
cd /etc/auxproxy/security
openssl genrsa -out tls.key 3072
openssl req -new -key tls.key -out tls.csr -subj "/OU=gameflexmatch/"
openssl x509 -req -days 365 -in tls.csr -signkey tls.key -out tls.crt

chown -R ${USER}:${GROUP} /etc/auxproxy/security
chmod -R 750 *

systemctl enable auxproxy.service

# ============================================
# 阿里云日志服务和云监控（可选安装）
# ============================================
# 阿里云日志服务Logtail安装（如需要）
# mkdir -p /etc/logtail
# cd /etc/logtail
# wget http://logtail-release-${REGION}.oss-${REGION}.aliyuncs.com/linux64/logtail.sh -O logtail.sh
# chmod 755 logtail.sh && ./logtail.sh install ${REGION}

# 阿里云云监控插件安装（如需要）
# ARGUS_VERSION=3.5.8
# wget -O /tmp/cloudmonitor.tar.gz "http://cms-agent-${REGION}.oss-${REGION}-internal.aliyuncs.com/cms-go-agent/${ARGUS_VERSION}/cms-go-agent-${ARGUS_VERSION}.linux-amd64.tar.gz"
# tar -xzf /tmp/cloudmonitor.tar.gz -C /usr/local/
# /usr/local/cms-go-agent/install.sh

shutdown
