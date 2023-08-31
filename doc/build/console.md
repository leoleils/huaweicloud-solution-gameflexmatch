# GameFlexMatch 前端测试部署流程

操作系统：CentOS 7.6 64bit

## 准备工作

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

3. 新建并修改 node.sh 文件

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
git clone -b master-dev https://gitee.com/HuaweiCloudDeveloper/huaweicloud-solution-gameflexmatch-console.git
cd huaweicloud-solution-gameflexmatch-console
```

8. 修改密码加密公钥配置文件 代码根目录/src/api/crypto.ts 第 42 行

```
export function encryptedData(data: string) {
    const encryptor = new JSEncrypt()
    let key = `XXXXXXXXX` // 设置公钥
    encryptor.setPublicKey(key)
    return encryptor.encrypt(data).toString()
}
```

9. 在代码根目录下运行 npm install 命令，安装项目所需要的依赖

```
npm install
```

10. 在根目录下安装 vue 国际化插件

```
npm install --save vue-i18n@next
```

## 本地运行前端代码

11. 在 env/.env.development 文件里配置后端 IP 和端口

```
VITE_APP_INTERFACE_URL="https://127.0.0.0:8080"
```

12. 在根目录下运行 npm run dev 命令，运行项目

```
npm run dev
```

## 打包部署

13. 在代码根目录下运行 npm run build 命令，将项目编译打包至根目录的 dist 文件夹下。

```
npm run build
```

14. 安装编译工具及库文件

```
yum -y install make zlib zlib-devel gcc-c++ libtool  openssl openssl-devel
```

15. 下载 PCRE 安装包

```
cd /usr/local/src/ || exit
wget http://downloads.sourceforge.net/project/pcre/pcre/8.35/pcre-8.35.tar.gz
tar zxvf pcre-8.35.tar.gz
```

16. 编译安装

```
cd pcre-8.35 || exit
./configure
make && make install
```

17. 安装 Nginx

```
cd /usr/local/src/ || exit
wget http://nginx.org/download/nginx-1.7.8.tar.gz
tar zxvf nginx-1.7.8.tar.gz
cd nginx-1.7.8 || exit
```

18. 编译安装到/usr/local/webserver/nginx

```
./configure --prefix=/usr/local/webserver/nginx --with-http_stub_status_module --with-http_ssl_module --with-pcre=/usr/local/src/pcre-8.35
make
make install
```

19. 替换 /usr/local/webserver/nginx/conf/nginx.conf 为以下内容

```
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  300;
    client_max_body_size 6g;
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

20. 把 dist 目录下的所有文件都复制到 nginx 网站根目录 /usr/local/webserver/nginx/html 下
21. 配置 nginx 开机自启动，在/lib/systemd/system/目录下创建 nginx.service 文件

```
vi /lib/systemd/system/nginx.service
```

22. 在该文件中添加如下内容

```
[Unit]
Description=nginx service
After=network.target
[Service]
Type=forking
ExecStart=/usr/local/webserver/nginx/sbin/nginx
ExecReload=/usr/local/webserver/nginx/sbin/nginx -s reload
ExecStop=/usr/local/webserver/nginx/sbin/nginx -s quit
PrivateTmp=true
[Install]
WantedBy=multi-user.target
```

23. 设置文件的执行权限

```
chmod a+x /lib/systemd/system/nginx.service
```

24. 设置开机自启动

```
systemctl enable nginx.service
```

25. 启动nginx服务 

```
systemctl start nginx.service
```

26. 浏览器输入该ECS绑定的ip地址即可访问GameFlexMatch前端界面，或输入: http://{ipv4}:80
