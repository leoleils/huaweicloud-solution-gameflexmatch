#bin/bash
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
chmod 750 *

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
chown -R ${USER}:${GROUP} /etc/auxproxy
chmod 750 *
mv auxproxy.service /etc/systemd/system

mkdir -p /etc/auxproxy/security
cd /etc/auxproxy/security
openssl genrsa -out tls.key 3072
openssl req -new -key tls.key -out tls.csr -subj "/OU=gameflexmatch/"
openssl x509 -req -days 365 -in tls.csr -signkey tls.key -out tls.crt

mkdir -p /etc/auxproxy/logs

systemctl enable auxproxy.service

# download LTS-ICAgent
mkdir -p /etc/lts
cd /etc/lts

curl http://icagent-${REGION}.obs.${REGION}.myhuaweicloud.com/ICAgent_linux/apm_agent_install.sh > apm_agent_install.sh && REGION=${REGION} bash apm_agent_install.sh -accessip ${LTS_IP} -obsdomain obs.${REGION}.myhuaweicloud.com;

# download CES-Agent
mkdir -p /etc/ces
cd /etc/ces && curl -k -O https://obs.${REGION}.myhuaweicloud.com/uniagent-${REGION}/script/agent_install.sh && bash agent_install.sh

shutdown

