# metaspace 前端测试部署流程

操作系统：CentOS 7.6 64bit

## 本地运行前端代码

1. 安装 git

```
yum -y install git
```

2. 安装 nodejs

```
wget https://nodejs.org/download/release/v16.13.1/node-v16.13.1-linux-x64.tar.gz
tar xf node-v16.13.1-linux-x64.tar.gz
mv node-v16.13.1-linux-x64 /usr/local/
```

3. 修改 node.sh 文件

```
vi /etc/profile.d/node.sh
```

4. 设置环境变量

```
export NODE_HOME=/usr/local/node-v16.13.1-linux-x64
export PATH=${NODE_HOME}/bin:$PATH
```

5. 执行脚本使环境变量生效

```
chmod +x /etc/profile.d/node.sh
source /etc/profile.d/node.sh
```

6. 安装 vue

```
npm install -g @vue/cli
```

7. 拉取前端代码

```
cd /usr/local
git clone -b master-dev https://codehub-dg-g.huawei.com/PublicCloudSolution/huaweicloud-solution-metaspace-console.git
cd huaweicloud-solution-metaspace-console
```

8. 修改密码加密公钥配置文件 src/api/crypto.ts

```
export function encryptedData(data: string) {
    const encryptor = new JSEncrypt()
    let key = `XXXXXXXXX` // 设置公钥
    encryptor.setPublicKey(key)
    return encryptor.encrypt(data).toString()
}
```

9. 在根目录下运行 npm install 命令，安装项目所需要的依赖

```
npm install
```

10. 在根目录下安装 vue 国际化插件

```
npm install --save vue-i18n@next
```

11. 在根目录下运行 npm run dev 命令，运行项目

```
npm run dev
```

## 打包部署

12. 在根目录下运行 npm run build 命令，将项目编译打包至根目录的 dist 文件夹下。

```
npm run build
```

13. 安装编译工具及库文件

```
yum -y install make zlib zlib-devel gcc-c++ libtool  openssl openssl-devel
```

14. 下载 PCRE 安装包

```
cd /usr/local/src/ || exit
wget http://downloads.sourceforge.net/project/pcre/pcre/8.35/pcre-8.35.tar.gz
tar zxvf pcre-8.35.tar.gz
```

15. 编译安装

```
cd pcre-8.35 || exit
./configure
make && make install
```

16. 安装 Nginx

```
cd /usr/local/src/ || exit
wget http://nginx.org/download/nginx-1.7.8.tar.gz
tar zxvf nginx-1.7.8.tar.gz
cd nginx-1.7.8 || exit
```

17. 编译安装到/usr/local/webserver/nginx

```
./configure --prefix=/usr/local/webserver/nginx --with-http_stub_status_module --with-http_ssl_module --with-pcre=/usr/local/src/pcre-8.35
make
make install
```

18. 替换 /usr/local/webserver/nginx/conf/nginx.conf 为以下内容

```
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  200;
    client_max_body_size 10m;
    server {
        listen       80;
        server_name  localhost;

        location /api/ {
            # 修改以下IP为后端服务器IP，把 /api 路径下的请求转发给真正的后端服务器
            proxy_pass https://127.0.0.1:8080/;
        }

        location / {
                root html;
                try_files $uri /index.html;  # try_files：检查文件； $uri：监测的文件路径； /index.html：文件不存在重定向的新路径
                index index.html;
        }

        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}

```

19. 把 dist 目录下的所有文件都复制到 nginx 网站根目录 /usr/local/webserver/nginx/html 下
20. 在 /usr/local/webserver/nginx/sbin 目录下运行./nginx，启动 nginx

```
./nginx
```
