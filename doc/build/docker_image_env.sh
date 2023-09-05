#!bin/bash
APP_URL=$1
AUX_URL=$2
APP_FILE=$3
AUX_FILE=$4
APP_NAME=$5
USER=$6
GROUP=$7
PWD=$8
SWR_LOGIN=$9
SWR_ORGANIZATION=${10}
DAEMON_SPEEDUP_URL=${11}
IMAGE_NAME=${12}
IMAGE_VERSION=${13}
DOCKER_SYSTEM=${14}
REGION=${15}


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
mkdir -p /etc/docker-build/${APP_NAME}
cd /etc/docker-build/${APP_NAME}
wget ${APP_URL} -O ${APP_FILE}
if [ "${APP_FILE##*.}" = "zip" ]; then
	unzip -o ${APP_FILE}
fi
if [ "${APP_FILE##*.}" = "rar" ]; then
	unrar e  ${APP_FILE}
fi
chmod 750 *

# download auxproxy
mkdir -p /etc/docker-build/auxproxy
cd /etc/docker-build/auxproxy
wget ${AUX_URL} -O ${AUX_FILE}
if [ "${AUX_FILE##*.}" = "zip" ]; then
	unzip -o ${AUX_FILE}
fi
if [ "${AUX_FILE##*.}" = "rar" ]; then
	unrar e  ${AUX_FILE}
fi
chmod 750 *


mkdir -p /etc/docker-build/auxproxy/security
cd /etc/docker-build/auxproxy/security

openssl genrsa -out tls.key 3072
openssl req -new -key tls.key -out tls.csr -subj "/OU=metaspace/"
openssl x509 -req -days 365 -in tls.csr -signkey tls.key -out tls.crt

mkdir -p /etc/docker-build/auxproxy/logs

cat > /etc/docker-build/supervisord.conf <<- EOF
[supervisord]
nodaemon=true

[program:auxproxy]
command=/bin/bash -c "exec /etc/auxproxy/auxproxy -auxproxy-address 0.0.0.0:60001 -grpc-address 0.0.0.0:60002 -cloud-platform-address http://169.254.169.254"
user = ${USER}
group = ${GROUP}
autostart = true
autorestart = true
startsecs = 30

EOF


#download docker
yum -y install docker

#register docker
systemctl enable docker

cat > /etc/docker/daemon.json <<- EOF
{
    "insecure-registries": ["https://${DAEMON_SPEEDUP_URL}"]
}
EOF

#restart docker
systemctl daemon-reload
systemctl restart docker

#build Dockerfile
cat > /etc/docker-build/Dockerfile <<- EOF
FROM ${DOCKER_SYSTEM}
RUN mkdir -p /etc/auxproxy \
&& mkdir -p /local/app/${APP_NAME} \
&& mkdir -p /etc/supervisor \
&& yum -y install supervisor
COPY auxproxy /etc/auxproxy
COPY ${APP_NAME} /local/app/${APP_NAME}
COPY supervisord.conf /etc/supervisor
RUN mkdir -p /var/log/supervisor \
&& useradd ${USER} || true \
&& groupadd ${GROUP} || true \
&& usermod -g ${GROUP} ${USER} || true \
&& echo ${USER}:${PWD} | chpasswd \
&& chown -R ${USER}:${GROUP} /local/app/${APP_NAME} \
&& chown -R ${USER}:${GROUP} /etc/auxproxy

ENTRYPOINT ["/usr/bin/supervisord","-c","/etc/supervisor/supervisord.conf"]

EOF

#build image
${SWR_LOGIN} swr.${REGION}.myhuaweicloud.com
docker build -t ${IMAGE_NAME}:${IMAGE_VERSION} /etc/docker-build/

#upload SWR
docker tag ${IMAGE_NAME}:${IMAGE_VERSION} swr.${REGION}.myhuaweicloud.com/${SWR_ORGANIZATION}/${IMAGE_NAME}:${IMAGE_VERSION}
docker push swr.${REGION}.myhuaweicloud.com/${SWR_ORGANIZATION}/${IMAGE_NAME}:${IMAGE_VERSION}

shutdown