#!/bin/bash

# -t: 用户Token，获取方式参考https://support.huaweicloud.com/api-iam/iam_30_0001.html
# -f: fleet id

# 参数说明：
# vpc_endpoint：vpc服务在上海一的服务域名
vpc_endpoint="https://vpc.cn-east-3.myhuaweicloud.com"


while getopts "f:t:" optname
do
        case "$optname" in
                "f")
                        fleet_id=$OPTARG
                        ;;
                "t")
                        user_token=$OPTARG
                        ;;
                "?")
                        echo "Unknown option $OPTARG"
                        ;;
                *)
                        echo "Unknown error while processing options"
                        ;;
        esac
done


function getJsonValuesByAwk(){
	awk -v json="$1" -v key="$2" -v defaultValue="$3" 'BEGIN{
		foundKeyCount = 0
		while (length(json) > 0 ){
			pos = match(json,"\""key"\"[\\t]*?:[\\t]*");
			if (pos == 0) {if (foundKeyCount == 0) {print defaultValue;} exit 0;}

			++foundKeyCount;
			start = 0; stop = 0; layer = 0;
			for (i = pos + length(key) +1; i <= length(json); ++i) {
				lastChar = substr(json,i - 1,1)
				currChar = substr(json,i,1)

				if (start <= 0) {
					if (lastChar == ":") {
						start = currChar == " "? i+1:i;
						if (currChar == "{" || currChar == "[") {
							layer = 1;
						}
					}
				} else {
					if (currChar == "{" || currChar == "[") {
						++layer;
					}
					if (currChar == "}" || currChar == "]"){
						--layer;
					}
					if ((currChar == ","|| currChar == "}" || currChar == "]") && layer <= 0) {
						stop = currChar == ","?i:i+1+layer;
						break;
					}
				}
			}

			if ( start <= 0 || stop <= 0 || start > length(json) || stop > length(json) || start >= stop) {
				if ( foundKeyCount == 0) { print defaultValue;} exit 0;
			} else {
				print substr(json, start, stop-start);
			}

			json = substr(json, stop+1, length(json)-stop)
		}
	}'
}


# get vcp peering id
list_peerings_url="${vpc_endpoint}/v2.0/vpc/peerings?name=${fleet_id}"
echo "[list peerings url] ${list_peerings_url}"
list_peerings_resp=$(curl -k -s -H "Content-Type: application/json" -H "X-Auth-Token: ${user_token}" ${list_peerings_url})
echo "[list peerings response] ${list_peerings_resp}"
peering_id=$(getJsonValuesByAwk "${list_peerings_resp}" "id" "" | sed 's/\"//g')


# delete vpc peering
delete_peering_url="${vpc_endpoint}/v2.0/vpc/peerings/${peering_id}"
echo "[delete peering url] ${delete_peering_url}"
curl -k -s -w "delete peering status: %{http_code}" -H "Content-Type: application/json" -H "X-Auth-Token: ${user_token}" -X DELETE ${delete_peering_url}
