# metaspace 前端测试部署流程

## 安装前端代码

1. 安装 git

```
yum -y install git
```

2. 安装 nodejs

```
yum -y install nodejs
```

3. 安装 npm

```
curl -sL https://rpm.nodesource.com/setup_14.x | bash -
```

4. 安装 vue

```
npm install -g @vue/cli
```

5. 安装 vite

```
npm init vite@latest
```

6. 拉取前端代码

```
cd /usr/local
git clone -b master-dev https://codehub-dg-g.huawei.com/PublicCloudSolution/huaweicloud-solution-metaspace-console.git
cd huaweicloud-solution-metaspace-console
```

7. 修改密码加密公钥配置文件 src/api/crypto.ts

```
export function encryptedData(data: string) {
    const encryptor = new JSEncrypt()
    let key = `XXXXXXXXX` // 设置公钥
    encryptor.setPublicKey(key)
    return encryptor.encrypt(data).toString()
}
```

8. 在根目录下运行 npm install 命令，安装项目所需要的依赖

```
npm install
```

9. 在根目录下运行 npm run dev 命令，运行项目

```
npm run dev
```

## 打包部署

10. 在根目录下运行 npm run build 命令，将项目编译打包至根目录的 dist 文件夹下。

```
npm run build
```

11. 安装编译工具及库文件

```
yum -y install make zlib zlib-devel gcc-c++ libtool  openssl openssl-devel
```

12. 下载 PCRE 安装包

```
cd /usr/local/src/ || exit
wget https://documentation-samples.obs.cn-north-4.myhuaweicloud.com/solution-as-code-publicbucket/solution-as-code-moudle/build-a-digital-assets-platform-based-on-MetaTown/open-source-software/pcre-8.35.tar.gz
tar zxvf pcre-8.35.tar.gz
```

13. 编译安装

```
cd pcre-8.35 || exit
./configure
make && make install
```

14. 安装 Nginx

```
cd /usr/local/src/ || exit
wget https://documentation-samples.obs.cn-north-4.myhuaweicloud.com/solution-as-code-publicbucket/solution-as-code-moudle/build-a-digital-assets-platform-based-on-MetaTown/open-source-software/nginx-1.7.8.tar.gz
tar zxvf nginx-1.7.8.tar.gz
cd nginx-1.7.8 || exit
```

15. 编译安装到/usr/local/webserver/nginx

```
./configure --prefix=/usr/local/webserver/nginx --with-http_stub_status_module --with-http_ssl_module --with-pcre=/usr/local/src/pcre-8.35
make
make install
```

16. 替换 /usr/local/webserver/nginx/conf/nginx.conf 为以下内容

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

12. 把 dist 目录下的所有文件都复制到 nginx 网站根目录 /usr/local/webserver/nginx/html 下
13. 在 /usr/local/webserver/nginx/sbin 目录下运行./nginx，启动 nginx

```
./nginx
```

# centos7 + x86 ecs 配置乌兰 203yum 镜像源

```
sed -i '$i nameserver 10.189.32.59' /etc/resolv.conf;
nmcli con mod 'System eth0' ipv4.dns "10.129.0.220 10.98.48.39";
nmcli con up 'System eth0';
sed -i '$a 7.223.219.40  mirrors.tools.huawei.com' /etc/hosts;
sed -i '$a 7.223.219.58  cmc-cd-mirror.rnd.huawei.com' /etc/hosts;
if [ "`cat /etc/yum.conf | grep sslverify=`" != "" ]; then sed -i 's/sslverify=.*/sslverify=false/g' /etc/yum.conf; else sed -i '$a sslverify=false' /etc/yum.conf; fi
mkdir -p /etc/repoback;
mv -f /etc/yum.repos.d/* /etc/repoback;
rm -rf /etc/yum.repos.d/*;
wget -q -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.tools.huawei.com/repository/conf/CentOS-7-anon.repo;
yum clean all;
yum makecache;
```
