# make-image-guide
应用镜像用于`MetaSpace`平台扩容实例，主要集成了`Auxproxy`服务组件以及希望执行的服务端应用
1. 在华为云控制台上购买1台ECS，选择密匙对登录
2. 上传应用安装包到指定路径，并修改可执行文件的权限

```sh
    # 1. 建立文件夹
    mkdir -p /local/app/{app-name}
    cd /local/app/{app-name}
    # 2. 上传应用包的二进制可执行文件到该路径下如application，并修改权限
    chmod 750 application
    # 3. 上传应用的启动脚本到同一目录下如application.sh，并修改权限
    chmod 750 application.sh
```

3. 上传`auxproxy`安装文件
   
```sh
    # 1. 上传以下文件并修改执行权限到指定文件夹
    mkdir -p /etc/auxproxy
    cd /etc/auxproxy
    # auxproxy-{version}是auxproxy组件编译后的二进制可执行文件
    # auxproxy-cleanup.sh是退出auxproxy组件的脚本文件
    # auxproxy-start.sh是启动auxproxy组件的脚本文件
    # 样例见/doc/build/auxproxy
    chmod 750 auxproxy-{version}
    chmod 750 auxproxy-cleanup.sh
    chmod 750 auxproxy-start.sh

    # 2. 新建security文件夹，生成密匙文件
    mkdir -p /etc/auxproxy/security
    cd /etc/auxproxy/security
    openssl genrsa -out tls.key 3072
    openssl req -new -key tls.key -out tls.csr
    openssl x509 -req -days 365 -in tls.csr -signkey tls.key -out tls.crt
    
    # 3. 新建logs文件夹
    mkdir -p /etc/auxproxy/logs

    # 4. 上传json文件
    # 上传client_hmac_conf.json与server_hmac_conf.json文件，
    # 样例见 /doc/build/auxproxy/security

    # 5. 在/etc/systemed/system下新建auxproxy.service
    # auxproxy.service 样例见 /doc/build/auxproxy
    # 启动auxproxy.service保证镜像自动拉起
    systemctl enable auxproxy.service
    systemctl start auxproxy.service

    # 验证是否配置成功：
    ps -aux | grep auxproxy
```

4. 将该ECS系统盘打包成Image，可参考[链接](https://support.huaweicloud.com/usermanual-ims/ims_01_0202.html)
5. 镜像打包完成