# vpc_peer使用方法
1. 运行`create_vpc_peering.sh`或`delete_vpc_peering.sh`前，需运行`get_user_token.sh`获取`X-Subject-Token`的值
2. `get_user_token.sh`运行命令示例：
   ```sh
   ./get_user_token.sh -d {华为云账号名} -u {IAM用户名} -p {密码}
   ```
   **注意：** `IAM`获取`token`的`API`中，密码信息是明文。建议通过`postman`或华为云`API Explorer`等工具获取`token`。
   + [IAM获取token的API详情](https://support.huaweicloud.com/intl/zh-cn/api-iam/iam_30_0001.html)
   + [华为云API Explorer链接](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=KeystoneCreateUserTokenByPasswordAndMfa)
3. 运行`create_vpc_peering.sh`前需配置：
   + 部署固定资源的`VPC ID`
   + 部署固定资源的路由地址`CIDR`；
   + 账号在预定region的`project id`
4. `create_vpc_peering.sh`运行命令示例：
   
```sh
./create_vpc_peering.sh -f {fleet id} -t {X-Subject-Token}
```

6. `delete_vpc_peering.sh`运行命令示例：
   
```sh
./delete_vpc_peering.sh -f {fleet id} -t {X-Subject-Token}
```