# cipher加解密脚本使用方法
## 简介
该脚本集成了RSA与GCM加解密的方法
+ GCM对称加解密需要16/24/32位的key与16位的Nonce，用于后台敏感数据的加解密
+ RSA非对称加解密需要生成公钥与私钥，生成方式见[README.md](/README.md)部署指南中第4步证书准备第2节，用于网络传输中的敏感数据加解密

## 使用方式
+ 配置`go`语言环境变量，运行: `go env -w GOOS=linux/winodws`，用于生成`linux`/`windows`下的可执行文件
+ 使用`go build ./main.go`生成`windows`/`linux`可执行文件后，输入执行参数后可执行加解密操作

## 参数介绍
|     字段名     |            描述            |                     参数要求                     |
| :------------: | :------------------------: | :----------------------------------------------: |
|     -mode      |        加密或者解密        |                  encode/decode                   |
|      -str      | 执行加密或解密的明文或密文 |                        --                        |
|      -key      |       GCM加解密的key       | GCM加解密必填，RSA加解密可不填，长度必须16/24/32 |
|     -nonce     |      GCM加解密的nonce      |    GCM加解密必填，RSA加解密可不填，长度为16位    |
| -cipher-method |       加解密得到方法       |                     GCM/RSA                      |
| -private-cert  |      RSA加解密的私钥       |              linux/windows格式路径               |
|  -public-cert  |      RSA加解密的公钥       |       基于私钥生成，linux/windows格式路径        |











