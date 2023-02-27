#!/bin/bash

# -t: 用户Token，获取方式参考https://support.huaweicloud.com/api-iam/iam_30_0001.html
# -f: fleet id

# target_vpc_id： 部署固定资源的VPC，即结算服等ECS的vpc id
# target_vpc_route_destination：部署固定资源的路由地址CIDR，即结算服等ECS的路由地址CIDR
# user_project_id：用户账号所在region的project id，即账号miniwan在上海一的project id
# vpc_endpoint： vpc服务在上海一的服务域名
target_vpc_id=""
target_vpc_route_destination=""
user_project_id=""
vpc_endpoint="https://vpc.cn-east-3.myhuaweicloud.com"


while getopts "t:f:" optname
do
        case "$optname" in
                "t")
                        user_token=$OPTARG
                        ;;
                "f")
                        fleet_id=$OPTARG
                        ;;
                "?")
                        echo "Unknown option $OPTARG"
                        ;;
                *)
                        echo "Unknown error while processing options"
                        ;;
        esac
done


function sentGetRequest(){
	curl -k -s -H "Content-Type: application/json" -H "X-Auth-Token: $1" $2
}


function sentPostRequest(){
        curl -k -s -H "Content-Type: application/json" -H "X-Auth-Token: $1"  -X POST -d "$3" $2
}


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



# list vpc id of fleet
list_vpcs_url="${vpc_endpoint}/v3/${user_project_id}/vpc/vpcs?name=${fleet_id}"
echo "[list vpc of fleet url] ${list_vpcs_url}"
list_vpcs_resp=$(sentGetRequest ${user_token} ${list_vpcs_url})
echo "[list vpc of fleet response] ${list_vpcs_resp}"
fleet_vpc_id=$(getJsonValuesByAwk "$list_vpcs_resp" "id" "" | sed 's/\"//g')


# list subnets of fleet vpc
list_subnets_url="${vpc_endpoint}/v1/${user_project_id}/subnets?vpc_id=${fleet_vpc_id}"
echo "[list subnets of fleet vpc url] ${list_subnets_url}"
list_subnets_resp=$(sentGetRequest ${user_token} ${list_subnets_url})
echo "[list subnets of fleet vpc response] ${list_subnets_resp}"
subnet_cidr=$(getJsonValuesByAwk "$list_subnets_resp" "cidr" "" | sed 's/\"//g')


# create vpc peering
create_peering_url="${vpc_endpoint}/v2.0/vpc/peerings"
create_peering_body="{\"peering\":{\"name\":\"${fleet_id}\",\"request_vpc_info\":{\"vpc_id\":\"${fleet_vpc_id}\"},\"accept_vpc_info\":{\"vpc_id\":\"${target_vpc_id}\"}}}"
echo "[create peering request body] ${create_peering_body}"
create_peering_resp=$(sentPostRequest ${user_token} ${create_peering_url} ${create_peering_body})
echo "[create peering response] ${create_peering_resp}"
peering_id=$(getJsonValuesByAwk "$create_peering_resp" "id" "" | sed 's/\"//g')


# create vpc route for fleet vpc
create_vpc_route_url="${vpc_endpoint}/v2.0/vpc/routes"
create_fleet_vpc_route_body="{\"route\":{\"destination\":\"${target_vpc_route_destination}\",\"nexthop\":\"${peering_id}\",\"type\":\"peering\",\"vpc_id\":\"${fleet_vpc_id}\"}}"
create_fleet_vpc_route_resp=$(sentPostRequest ${user_token} ${create_vpc_route_url} ${create_fleet_vpc_route_body})

# create vpc route for target vpc
create_target_vpc_route_body="{\"route\":{\"destination\":\"${subnet_cidr}\",\"nexthop\":\"${peering_id}\",\"type\":\"peering\",\"vpc_id\":\"${target_vpc_id}\"}}"
create_target_vpc_route_resp=$(sentPostRequest ${user_token} ${create_vpc_route_url} ${create_target_vpc_route_body})

