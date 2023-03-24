# fake-server
用于对接Metaspace的模拟托管应用，默认启动路径为:`/local/app/fake-server/fake-server.sh`，默认启动端口为`1025-60000`，默认日志路径为`/local/app/fake-server/log`

## 使用方式
+ 编译与配置
```sh
    # 设置编译的可执行文件的操作系统
    go env -w GOOS=linux
    # 配置go代理
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://repo.huaweicloud.com/repository/goproxy/
    go env -w GONOSUMDB=*
    # 生成可执行二进制文件fake-server
    cd ./huaweicloud-solution-metaspace/demo/go/fake-server
    go mod tidy
    go build
    # 配置启动脚本
    vim ./fake-server.sh
    # 按需修改启动端口范围与日志路径
```
+ 配置完成后即可使用