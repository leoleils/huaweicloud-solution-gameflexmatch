1. 运行create_vpc_peering.sh或delete_vpc_peering.sh前，需运行get_user_token.sh获取X-Subject-Token的值
2. get_user_token.sh运行命令示例：./get_user_token.sh -d {华为云账号名} -u {IAM用户名} -p {密码}
   【注】IAM获取token的API中，密码信息是明文。建议通过postman或华为云API Explorer等工具获取token。
             IAM获取token的API详情：https://support.huaweicloud.com/intl/zh-cn/api-iam/iam_30_0001.html
             华为云API Explorer链接：https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=KeystoneCreateUserTokenByPasswordAndMfa
3. 运行create_vpc_peering.sh前需配置：①部署固定资源的VPC ID（即结算服等ECS的VPC ID）②部署固定资源的路由地址CIDR；③账号上海一的project id
4. create_vpc_peering.sh运行命令示例：./create_vpc_peering.sh -f {fleet id} -t {X-Subject-Token}
5. delete_vpc_peering.sh运行命令示例：./delete_vpc_peering.sh -f {fleet id} -t {X-Subject-Token}