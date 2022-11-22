#!/bin/bash

function getJsonValuesByAwk() {
    awk -v json="$1" -v key="$2" -v defaultValue="$3" 'BEGIN{
        foundKeyCount = 0
        while (length(json) > 0) {
            # pos = index(json, "\""key"\""); ## 这行更快一些，但是如果有value是字符串，且刚好与要查找的key相同，会被误认为是key而导致值获取错误
            pos = match(json, "\""key"\"[ \\t]*?:[ \\t]*");
            if (pos == 0) {if (foundKeyCount == 0) {print defaultValue;} exit 0;}

            ++foundKeyCount;
            start = 0; stop = 0; layer = 0;
            for (i = pos + length(key) + 1; i <= length(json); ++i) {
                lastChar = substr(json, i - 1, 1)
                currChar = substr(json, i, 1)

                if (start <= 0) {
                    if (lastChar == ":") {
                        start = currChar == " " ? i + 1: i;
                        if (currChar == "{" || currChar == "[") {
                            layer = 1;
                        }
                    }
                } else {
                    if (currChar == "{" || currChar == "[") {
                        ++layer;
                    }
                    if (currChar == "}" || currChar == "]") {
                        --layer;
                    }
                    if ((currChar == "," || currChar == "}" || currChar == "]") && layer <= 0) {
                        stop = currChar == "," ? i : i + 1 + layer;
                        break;
                    }
                }
            }

            if (start <= 0 || stop <= 0 || start > length(json) || stop > length(json) || start >= stop) {
                if (foundKeyCount == 0) {print defaultValue;} exit 0;
            } else {
                print substr(json, start, stop - start);
            }

            json = substr(json, stop + 1, length(json) - stop)
        }
    }'
}


echo "here is auxproxy service"

# 不断重试
gateway_addr=""
while [ "$gateway_addr" = "" ]; do
	echo "start try to fetch user_data"
	user_data=$(curl http://169.254.169.254/openstack/latest/user_data -s)
	#user_data='{"fleet_id":"fleet-1","gateway_address":" 100.93.18.230:60003","scaling_group_id":"4c11ead6-1f61-44be-8a46-abb8351acdb8"}'
	gateway_addr=$(getJsonValuesByAwk "$user_data" "gateway_address" | sed 's/\"//g')
	sleep 2s
done

echo "fetch user data $user_data, and parse json to get gateway_addr $gateway_addr, start auxproxy"

#/etc/auxproxy/auxproxy -auxproxy-address 0.0.0.0:10000 -grpc-address localhost:10001 -gateway-address $gateway_addr
#/etc/auxproxy/auxproxy -auxproxy-address 0.0.0.0:60001 -grpc-address 0.0.0.0:60002 -gateway-address $gateway_addr -scaling-group-id edd39e1e-5623-4ce7-baaf-7f39c0669098
/etc/auxproxy/auxproxy -auxproxy-address 0.0.0.0:60001 -grpc-address 0.0.0.0:60002 -cloud-platform-address http://169.254.169.254
