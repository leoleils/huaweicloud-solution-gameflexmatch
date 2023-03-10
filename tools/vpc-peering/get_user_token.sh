#!/bin/bash


# -d: domain name,即华为云账号名
# -u: user name,即IAM用户名
# -p: password

domain_name=""
user_name=""
password=""

while getopts "d:u:p:" optname
do
	case "$optname" in
		"d")
			domain_name=$OPTARG
			;;
		"u")
			user_name=$OPTARG
			;;
		"p")
			password=$OPTARG
			;;
		"?")
			echo "Unknown option $OPTARG"
			;;
		*)
			echo "Unknown error while processing options"
			;;
	esac
done


iam_url="https://iam.cn-east-3.myhuaweicloud.com/v3/auth/tokens"
region="cn-east-3"
body="{\"auth\":{\"identity\":{\"methods\":[\"password\"],\"password\":{\"user\":{\"domain\":{\"name\":\"${domain_name}\"},\"name\":\"${user_name}\",\"password\":\"${password}\"}}},\"scope\":{\"project\":{\"name\":\"${region}\"}}}}"
resp=$(curl -k -s -i -H "Content-Type: application/json" -H "X-Auth-Token: ${user_token}" -X POST -d "${body}" ${iam_url}| grep -E 'X-Subject-Token')
echo ${resp}
