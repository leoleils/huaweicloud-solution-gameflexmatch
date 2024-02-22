# Introduction to Configuration Files

GFM generates configuration files required by each service component based on the `conf/init.yaml` configuration file. This section describes the parameters in the configuration files and how to configure the parameters.)

**Modify the parameters in red based on the site requirements. Others you can keep default value**

## 参数介绍
<style type="text/css">
.tg  {border-collapse:collapse;border-color:#ccc;border-spacing:0;}
.tg td{background-color:#fff;border-color:#ccc;border-style:solid;border-width:1px;color:#333;
  font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;word-break:normal;}
.tg th{background-color:#f0f0f0;border-color:#ccc;border-style:solid;border-width:1px;color:#333;
  font-family:Arial, sans-serif;font-size:14px;font-weight:normal;overflow:hidden;padding:10px 5px;word-break:normal;}
.tg .tg-cly1{text-align:left;vertical-align:middle}
.tg .tg-x3ds{background-color:#f9f9f9;color:#00F;text-align:left;text-decoration:underline;vertical-align:middle}
.tg .tg-m64s{background-color:#F2F2F2;font-weight:bold;text-align:center;vertical-align:middle}
.tg .tg-ym56{background-color:#f9f9f9;color:#F00;text-align:left;vertical-align:middle}
.tg .tg-6vn3{background-color:#F2F2F2;color:#F00;text-align:left;vertical-align:middle}
.tg .tg-fapl{color:#F00;text-align:left;vertical-align:middle}
.tg .tg-8v0r{background-color:#F2F2F2;color:#00F;text-align:left;text-decoration:underline;vertical-align:top}
.tg .tg-61g0{background-color:#F2F2F2;text-align:left;vertical-align:middle}
.tg .tg-ukly{background-color:#D9D9D9;font-weight:bold;text-align:center;vertical-align:middle}
.tg .tg-ogp5{background-color:#f9f9f9;font-weight:bold;text-align:center;vertical-align:middle}
.tg .tg-yjjc{background-color:#f9f9f9;text-align:left;vertical-align:middle}
.tg .tg-wu5t{background-color:#F2F2F2;font-weight:bold;text-align:center;vertical-align:middle}
.tg .tg-poke{background-color:#F2F2F2;color:#F00;text-align:left;vertical-align:middle}
.tg .tg-5fx9{background-color:#F2F2F2;text-align:left;vertical-align:middle}
.tg .tg-l010{background-color:#F2F2F2;color:#00F;text-align:left;text-decoration:underline;vertical-align:middle}
.tg .tg-0vn4{background-color:#F2F2F2;text-align:left;vertical-align:top}
</style>
<table class="tg">
<thead>
  <tr>
    <th class="tg-ukly">Section(模块)</th>
    <th class="tg-ukly">Parameter(参数)</th>
    <th class="tg-ukly">Description(描述)</th>
    <th class="tg-ukly">How to Configurate(如何配置)</th>
    <th class="tg-ukly">Example(示例)</th>
    <th class="tg-ukly">其他(Others)</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td class="tg-ogp5" rowspan="7">cipher</td>
    <td class="tg-yjjc">GCMKey</td>
    <td class="tg-yjjc">24 characters used for encrypting and decrypting sensitive data such as passwords</td>
    <td class="tg-yjjc" rowspan="7">It has been generated for you by default. You can regenerate it if necessary. </td>
    <td class="tg-yjjc">Fi1xOkgXvevHlTR5SIHHXAb2</td>
    <td class="tg-yjjc" rowspan="2">If you need to regenerate, you can exec: ./sac-gfm cipher --create gcm --init-conf conf/init.yaml</td>
  </tr>
  <tr>
    <td class="tg-cly1">GCMNonce</td>
    <td class="tg-cly1">16 characters used for encrypting and decrypting sensitive data such as passwords</td>
    <td class="tg-cly1">dlbaGmj9O32CFbTs</td>
  </tr>
  <tr>
    <td class="tg-yjjc">RSAPublicKey</td>
    <td class="tg-yjjc" rowspan="2">Encryption for password transmission between the console and the service with RSA</td>
    <td class="tg-yjjc">./conf/public.key</td>
    <td class="tg-yjjc" rowspan="2">If you need to regenerate, you can exec: ./sac-gfm cipher --create rsa --init-conf conf/init.yaml. The key must be the same as that on the console. The default public key has been compiled with the frontend software package. If you regenerate the key, you must recompile the frontend software package. For details about how to compile the key, you can refer to: .</td>
  </tr>
  <tr>
    <td class="tg-cly1">RSAPrivateKey</td>
    <td class="tg-cly1">./conf/private.key</td>
  </tr>
  <tr>
    <td class="tg-yjjc">TlsKey</td>
    <td class="tg-yjjc" rowspan="3">From signature certificate for HTTPS protocol</td>
    <td class="tg-yjjc">./conf/tls.key</td>
    <td class="tg-yjjc" rowspan="3">If you need to regenerate, you can exec: ./sac-gfm cipher --create tls --init-conf conf/init.yaml</td>
  </tr>
  <tr>
    <td class="tg-cly1">TlsCsr</td>
    <td class="tg-cly1">./conf/tls.key</td>
  </tr>
  <tr>
    <td class="tg-yjjc">TlsCrt</td>
    <td class="tg-yjjc">./conf/tls.key</td>
  </tr>
  <tr>
    <td class="tg-wu5t" rowspan="65">fleetmanager</td>
    <td class="tg-poke">MysqlAddress</td>
    <td class="tg-61g0">mysql address used by fleetmanager</td>
    <td class="tg-61g0">Ensure that the fleetmanager service can connect to the database.</td>
    <td class="tg-61g0">127.0.0.1:3306</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">MysqlUser</td>
    <td class="tg-5fx9">mysql user used by fleetmanager</td>
    <td class="tg-5fx9">Generally, the user is root.</td>
    <td class="tg-5fx9">root</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">MysqlDBName</td>
    <td class="tg-61g0">database name used by fleetmanager</td>
    <td class="tg-61g0">It is not created by default. Create it in the database in advance.</td>
    <td class="tg-61g0">fleetmanager</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">MysqlPassword</td>
    <td class="tg-5fx9">mysql password used by fleetmanager</td>
    <td class="tg-5fx9">Database password encrypted using GCM</td>
    <td class="tg-8v0r">　</td>
    <td class="tg-5fx9">You can run the following command to encrypt the password: /sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
  </tr>
  <tr>
    <td class="tg-61g0">MysqlCharset</td>
    <td class="tg-61g0">Database coding format</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">utf8</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">RedisAddress</td>
    <td class="tg-5fx9">redis address used by fleetmanager</td>
    <td class="tg-5fx9">Ensure that the fleetmanager service can connect to the redis.</td>
    <td class="tg-5fx9">127.0.0.1:6379</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">RedisPassword</td>
    <td class="tg-61g0">redis password used by fleetmanager</td>
    <td class="tg-61g0">redis password encrypted using GCM</td>
    <td class="tg-61g0">　</td>
    <td class="tg-61g0">You can run the following command to encrypt the password: /sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
  </tr>
  <tr>
    <td class="tg-5fx9">RedisDB</td>
    <td class="tg-5fx9">Default Redis partition</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">1</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">LogRotateSize</td>
    <td class="tg-61g0">Size of a log split, in MB.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">1024</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">LogBackupCount</td>
    <td class="tg-5fx9">Maximum number of logs that can be retained. If the number of logs exceeds the maximum, the logs are deleted by time.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">100</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">LogMaxAge</td>
    <td class="tg-61g0">Maximum number of days for storing logs.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">7</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">SupportRegions</td>
    <td class="tg-5fx9">If multiple regions are deployed, you can enter multiple values. Otherwise, only one value is required.</td>
    <td class="tg-5fx9">You can refer: <br> &nbsp;&nbsp;China: https://developer.huaweicloud.com/endpoint<br> &nbsp;&nbsp;Intl: https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
    <td class="tg-5fx9">cn-north-4</td>
    <td class="tg-5fx9">This configuration item needs to be configured as a YAML list.</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultLoginPassword</td>
    <td class="tg-61g0">It is used only when you log in to the system for the first time. After the initialization, the password will be changed to a new one.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-l010">Gfm@2024</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">WorkflowPath</td>
    <td class="tg-5fx9">Path of the workflow configuration file</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">./conf/workflow/</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">SupportPublicImage</td>
    <td class="tg-61g0">List of public ECS images that have been verified. <br> &nbsp;&nbsp;If other images are available, you can change them to the ones you want.</td>
    <td class="tg-61g0">You can customize the region and retain the default values for other parameters.</td>
    <td class="tg-61g0">SupportPublicImage: <br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      image: "CentOS 7.2 64bit,CentOS 7.9 64bit,CentOS 8.1 64bit,Ubuntu 22.04 server 64bit,CentOS 7.6 64bit for Tenant 20210525"</td>
    <td class="tg-61g0">This configuration item needs to be configured as a YAML list.</td>
  </tr>
  <tr>
    <td class="tg-6vn3">SupportDockerImage</td>
    <td class="tg-5fx9">List of Docker images of verified pods. <br> &nbsp;&nbsp;If you are sure that other images are available, you can change them to the ones you want.</td>
    <td class="tg-5fx9">You can customize the region and retain the default values for other parameters.</td>
    <td class="tg-5fx9">SupportDockerImage: <br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      dockerOS: "swr.cn-north-7.myhuaweicloud.com/game/centos-super:v1,centos:7.6.1810,centos:7.2.1511"</td>
    <td class="tg-5fx9">This configuration item needs to be configured as a YAML list.</td>
  </tr>
  <tr>
    <td class="tg-poke">EipType</td>
    <td class="tg-61g0">The EIP type supported by each region may be different. <br> &nbsp;&nbsp;You need to manually configure the EIP type after querying the EIP type on the cloud service.</td>
    <td class="tg-61g0">You can refer to: <br> &nbsp;&nbsp;Table 4 in China:  https://support.huaweicloud.com/api-eip/eip_api_0001.html<br> &nbsp;&nbsp;Table 3 in Intl: https://support.huaweicloud.com/intl/en-us/api-eip/eip_api_0001.html<br> &nbsp;&nbsp;If you have more, should Use commas to separate, for example: "eip-type1,eip-type2"</td>
    <td class="tg-61g0">EipType:<br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      supportEipType: 5_bgp,5_sbgp</td>
    <td class="tg-61g0">This configuration item needs to be configured as a YAML list.</td>
  </tr>
  <tr>
    <td class="tg-6vn3">DnsConfig</td>
    <td class="tg-5fx9">Huawei Cloud Private DNS Server Addresses.<br> &nbsp;&nbsp;</td>
    <td class="tg-5fx9">You can get it from Table 1 in:<br> &nbsp;&nbsp;China: https://support.huaweicloud.com/dns_faq/dns_faq_002.html<br> &nbsp;&nbsp;Intl: https://support.huaweicloud.com/intl/en-us/dns_faq/dns_faq_002.html</td>
    <td class="tg-5fx9">DnsConfig: <br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      dnsServer: 100.125.1.250,100.125.129.250</td>
    <td class="tg-5fx9">This configuration item needs to be configured as a YAML list.</td>
  </tr>
  <tr>
    <td class="tg-poke">LtsIps</td>
    <td class="tg-61g0">This configuration item is used if you want to use LTS to manage logs.<br> &nbsp;&nbsp;This configuration item is used to automatically install ICAgent on HUAWEI CLOUD hosts for log collection. </td>
    <td class="tg-61g0">You need to query the configuration address.<br> &nbsp;&nbsp;You can log in to the LTS service console, Select the region to be deployed, <br> &nbsp;&nbsp;Select: Host Management -&gt; Install ICAgent.<br> &nbsp;&nbsp;You can get it from the step 2 by accessip in command given.</td>
    <td class="tg-61g0">LtsIps:<br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      ltsIp: 100.125.12.150</td>
    <td class="tg-61g0">This configuration item needs to be configured as a YAML list.</td>
  </tr>
  <tr>
    <td class="tg-6vn3">DefaultRegion</td>
    <td class="tg-5fx9">If the region information is not provided when creating fleet, the value will be used.</td>
    <td class="tg-5fx9">　</td>
    <td class="tg-5fx9">cn-north-4</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultFleetProtectPolicy</td>
    <td class="tg-61g0">Default protection policy of the fleet. <br> &nbsp;&nbsp;This configuration item will be used during scale-in.<br> &nbsp;&nbsp;Support: [TIME_LIMIT_PROTECTION, NO_PROTECTION]<br> &nbsp;&nbsp;NO_PROTECTION maybe affect the player who is playing the game, not suggest<br> &nbsp;&nbsp;TIME_LIMIT_PROTECTION: You will give a protection period. <br> &nbsp;&nbsp;After the scale-in is triggered, the machine will not be forcibly reclaimed during the protection period.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">TIME_LIMIT_PROTECTION</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultProtectTimeLimit</td>
    <td class="tg-5fx9">Protection time during Fleet scale-in, in minute.<br> &nbsp;&nbsp;This parameter is valid only when Protection Policy is set to TIME_LIMIT_PROTECTION.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">30</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultSessionTimeoutSeconds</td>
    <td class="tg-61g0">When a session is created, <br> &nbsp;&nbsp;if the session is not activated after the time specified by this parameter, in seconds.<br> &nbsp;&nbsp;the session will set to error.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">60</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultMaxSessionNumPerProcess</td>
    <td class="tg-5fx9">Configure the default number of session to be started per process.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">50</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultProcessNumPerInstance</td>
    <td class="tg-61g0">Configure the default number of Process to be started per instance.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">50</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultBandwidth</td>
    <td class="tg-5fx9">Default bandwidth</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">5</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultBandwidthChargingMode</td>
    <td class="tg-61g0">Default bandwidth charge mode<br> &nbsp;&nbsp;Support: [bandwidth, traffic]</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">traffic</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultDiskSize</td>
    <td class="tg-5fx9">Default disk size</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">40</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultVolumeType</td>
    <td class="tg-61g0">Default Volume Type</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">SATA</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultEipShareType</td>
    <td class="tg-5fx9">Specifies the bandwidth sharing type.<br> &nbsp;&nbsp;Share type enumerated value: <br> &nbsp;&nbsp;PER, indicating exclusive use. WHOLE: indicates sharing.<br> &nbsp;&nbsp;Support: [PER, WHOLE]<br> &nbsp;&nbsp;Default: PER</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">PER</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">FleetCidr</td>
    <td class="tg-61g0">This configuration item is required when a VPC needs to be created during fleet creation.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">FleetCidr:<br> &nbsp;&nbsp;    from: 10.100.0.0<br> &nbsp;&nbsp;    to: 10.240.0.0<br> &nbsp;&nbsp;    vpcNetMask: 19<br> &nbsp;&nbsp;    subnetNetMask: 20</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">InternalInboundPermissions</td>
    <td class="tg-5fx9">A security group is created by default when a fleet is created. <br> &nbsp;&nbsp;The following lists the security group rules that need to be configured by default.</td>
    <td class="tg-5fx9">The ipRange configuration item needs to be set to the IP addresses of all AppGateway nodes. <br> &nbsp;&nbsp;For example, if there are two AppGateway nodes (10.100.0.1 and 10.100.0.2), you need to set ipRange to the values as the example value. <br> &nbsp;&nbsp;You can keep the default settings for other parameters.</td>
    <td class="tg-5fx9">InternalInboundPermissions:<br> &nbsp;&nbsp;    - protocol: "TCP"<br> &nbsp;&nbsp;      ipRange: 10.100.0.1/32<br> &nbsp;&nbsp;      fromPort: 60001<br> &nbsp;&nbsp;      toPort: 60001<br> &nbsp;&nbsp;    - protocol: "TCP"<br> &nbsp;&nbsp;      ipRange: 10.100.0.1/32<br> &nbsp;&nbsp;      fromPort: 60001<br> &nbsp;&nbsp;      toPort: 60001</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BuildImageRef</td>
    <td class="tg-61g0">Default OS used when creating an image</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">CentOS 7.2 64bit</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">BuildFlavor</td>
    <td class="tg-5fx9">Default flavor used during image creation</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">s6.large.2</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BuildBandwidth</td>
    <td class="tg-61g0">Default bandwidth used during image creation</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">100</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">BuildScriptPath</td>
    <td class="tg-5fx9">Script for creating an ECS image</td>
    <td class="tg-5fx9">Upload the script for packaging the image to the OBS bucket. <br> &nbsp;&nbsp;For example, if the OBS bucket is gfm, set this parameter to a value similar to the example value.</td>
    <td class="tg-5fx9">gfm/image_env.sh</td>
    <td class="tg-5fx9">Upload the bin/image_env.sh script to the gfm bucket of OBS.</td>
  </tr>
  <tr>
    <td class="tg-poke">BuildDockerScriptPath</td>
    <td class="tg-61g0">Script used for creating POD images</td>
    <td class="tg-61g0">Upload the script for packaging the docker image to the OBS bucket. <br> &nbsp;&nbsp;For example, if the OBS bucket is gfm, set this parameter to a value similar to the example value.</td>
    <td class="tg-61g0">gfm/docker_image_env.sh</td>
    <td class="tg-61g0">Upload the bin/docker_image_env.sh script to the gfm bucket of OBS.</td>
  </tr>
  <tr>
    <td class="tg-6vn3">AuxproxyPath</td>
    <td class="tg-5fx9">Application package of the auxproxy service</td>
    <td class="tg-5fx9">Upload the application package to the OBS bucket. <br> &nbsp;&nbsp;For example, if the OBS bucket is gfm, set this parameter to a value similar to the example value.</td>
    <td class="tg-5fx9">gfm/auxproxy.zip</td>
    <td class="tg-5fx9">Upload the bin/auxproxy.zip to the gfm bucket of OBS.</td>
  </tr>
  <tr>
    <td class="tg-61g0">ImageDiskSize</td>
    <td class="tg-61g0">Disk size of image michine</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">40</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">ServerSessionBackupDays</td>
    <td class="tg-5fx9">Expiration time of historical data, in days.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">1</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultScalingInCoolDownInterval</td>
    <td class="tg-61g0">Scale-in interval, in minutes</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">10</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultGroupMaxSize</td>
    <td class="tg-5fx9" rowspan="3">Maximum, minimum, and desire number of instances in an AS group</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">1</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultGroupMinSize</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">1</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultGroupDesiredSize</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">1</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">ServiceEndpoint.AASS</td>
    <td class="tg-61g0">Address used by the AASS to provide services for external systems. FleetManager can access the AASS through this address.</td>
    <td class="tg-61g0">If you have configured the ELB service for AASS, configurate the IP address of the ELB service.</td>
    <td class="tg-61g0">https://127.0.0.1:9091</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">ServiceEndpoint.AppGateway</td>
    <td class="tg-5fx9">Address used by the Appgateway to provide services for external systems. <br> &nbsp;&nbsp;FleetManager can access the Appgateway through this address.</td>
    <td class="tg-5fx9">If you have configured the ELB service for appgateway, configurate the IP address of the ELB service.</td>
    <td class="tg-5fx9">https://127.0.0.1:60003</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">ServiceEndpoint.IamService</td>
    <td class="tg-61g0" rowspan="11">Service Endpoint about HuaweiCloud Service</td>
    <td class="tg-0vn4" rowspan="11"> you can get those from: <br> &nbsp;&nbsp;China: https://developer.huaweicloud.com/endpoint<br> &nbsp;&nbsp;Intl: https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
    <td class="tg-61g0">https://iam.cn-north-4.myhuaweicloud.com</td>
    <td class="tg-61g0" rowspan="11">The example value omits the region. In practice, the region should be set to:<br> &nbsp;&nbsp;region: cn-north-4<br> &nbsp;&nbsp;endPoint:*****</td>
  </tr>
  <tr>
    <td class="tg-6vn3">ServiceEndpoint.VpcService</td>
    <td class="tg-5fx9">https://vpc.cn-north-4.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-poke">ServiceEndpoint.ObsService</td>
    <td class="tg-61g0">https://obs.cn-north-4.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-6vn3">ServiceEndpoint.ImsService</td>
    <td class="tg-5fx9">https://ims.cn-north-4.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-poke">ServiceEndpoint.EcsService</td>
    <td class="tg-61g0">https://ecs.cn-north-4.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-6vn3">ServiceEndpoint.SwrService</td>
    <td class="tg-5fx9">https://swr.cn-north-4.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-poke">ServiceEndpoint.CciService</td>
    <td class="tg-61g0">https://cci.cn-north-4.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-6vn3">ServiceEndpoint.EvsService</td>
    <td class="tg-5fx9">https://evs.cn-north-4.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-poke">ServiceEndpoint.CesService</td>
    <td class="tg-61g0">https://ces.cn-north-4.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-6vn3">ServiceEndpoint.DnsService</td>
    <td class="tg-5fx9">https://dns.cn-north-4.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-poke">ServiceEndpoint.EpsService</td>
    <td class="tg-61g0">https://eps.myhuaweicloud.com</td>
  </tr>
  <tr>
    <td class="tg-6vn3">SessionJwtKey</td>
    <td class="tg-5fx9">Similar to the key used to generate or verify the login session,</td>
    <td class="tg-5fx9">you can choose a random character string of 24 or 32 bits. It is generated by default.</td>
    <td class="tg-5fx9">fFMH9nZgM3rpzp2os6RSbIicAeTazBV0I7ZZwFWnIQJFMHVP</td>
    <td class="tg-5fx9">If you need to regenerate, you can exec: ./sac-gfm cipher --create jwt --init-conf conf/init.yaml</td>
  </tr>
  <tr>
    <td class="tg-61g0">SessionJwtTokenLifeTimeSecond</td>
    <td class="tg-61g0">Timeout interval of the login status token.<br> &nbsp;&nbsp;If the timeout interval is exceeded, you need to log in again</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">7200</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">TakeOverTaskIntervalSeconds</td>
    <td class="tg-5fx9">Interval for taking over asynchronous tasks, in seconds. <br> &nbsp;&nbsp;This parameter is used when a node is faulty.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">60</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HttpsAddress</td>
    <td class="tg-61g0">bind address</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">0.0.0.0</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">HttpsPort</td>
    <td class="tg-5fx9">bind port</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">31002</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HeartBeatTaskIntervalSeconds</td>
    <td class="tg-61g0">Interval for reporting node health status, in seconds, <br> &nbsp;&nbsp;which is used to determine whether a node is normal.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">60</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DeadCheckTaskIntervalSeconds</td>
    <td class="tg-5fx9">Interval for monitoring zombie nodes, in seconds</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">60</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">MaxDeadMinutes</td>
    <td class="tg-61g0">Interval for determining a zombie node, in mintues.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">3</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-ogp5" rowspan="20">appgateway</td>
    <td class="tg-ym56">AASSAddress</td>
    <td class="tg-yjjc">AASS address for the AppGateway to access the AASS</td>
    <td class="tg-yjjc">If you have configured the ELB service for AASS, configurate the IP address of the ELB service.</td>
    <td class="tg-yjjc">127.0.0.1:9091</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-fapl">MysqlAddress</td>
    <td class="tg-cly1">mysql address used by appgateway</td>
    <td class="tg-cly1">Ensure that the appgateway service can connect to the database.</td>
    <td class="tg-cly1">127.0.0.1:3306</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">MysqlUser</td>
    <td class="tg-yjjc">mysql user used by appgateway</td>
    <td class="tg-yjjc">Generally, the user is root.</td>
    <td class="tg-yjjc">root</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-fapl">MysqlDBName</td>
    <td class="tg-cly1">database name used by appgateway</td>
    <td class="tg-cly1">It is not created by default. Create it in the database in advance.</td>
    <td class="tg-cly1">appgateway</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-ym56">MysqlPassword</td>
    <td class="tg-yjjc">mysql password used by appgateway</td>
    <td class="tg-yjjc">Database password encrypted using GCM</td>
    <td class="tg-x3ds"></td>
    <td class="tg-yjjc">You can run the following command to encrypt the password: /sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
  </tr>
  <tr>
    <td class="tg-fapl">RedisAddress</td>
    <td class="tg-cly1">redis address used by appgateway</td>
    <td class="tg-cly1">Ensure that the appgateway service can connect to the redis.</td>
    <td class="tg-cly1">127.0.0.1:6379</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-ym56">RedisPassword</td>
    <td class="tg-yjjc">redis password used by appgateway</td>
    <td class="tg-yjjc">redis password encrypted using GCM</td>
    <td class="tg-yjjc"></td>
    <td class="tg-yjjc">You can run the following command to encrypt the password: /sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
  </tr>
  <tr>
    <td class="tg-cly1">RedisDB</td>
    <td class="tg-cly1">Default Redis partition</td>
    <td class="tg-cly1">You can keep the default</td>
    <td class="tg-cly1">3</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">CleanStrategy</td>
    <td class="tg-yjjc">clean data in Database or not<br> &nbsp;&nbsp;Support: [off, on]</td>
    <td class="tg-yjjc">You can keep the default</td>
    <td class="tg-yjjc">on</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">CleanupDays</td>
    <td class="tg-cly1">Delete data in process and session backup table from database, Default 14 days</td>
    <td class="tg-cly1">You can keep the default</td>
    <td class="tg-cly1">14</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">BackupDays</td>
    <td class="tg-yjjc">Backup data in process and session runtime table from database, Default 3 days</td>
    <td class="tg-yjjc">You can keep the default</td>
    <td class="tg-yjjc">3</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">HttpsAddress</td>
    <td class="tg-cly1">bind address</td>
    <td class="tg-cly1">You can keep the default</td>
    <td class="tg-cly1">0.0.0.0</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">HttpsPort</td>
    <td class="tg-yjjc">bind port</td>
    <td class="tg-yjjc">You can keep the default</td>
    <td class="tg-yjjc">60003</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">DeployModel</td>
    <td class="tg-cly1">Appgateway startup mode<br> &nbsp;&nbsp;Support: [singleton, multi-instances]</td>
    <td class="tg-cly1">You can keep the default</td>
    <td class="tg-cly1">multi-instances</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">AuxproxyIpType</td>
    <td class="tg-yjjc">Configure the Appgateway to connect to the AuxProxy using a public or private IP address.<br> &nbsp;&nbsp;Support: [publicIP, privateIP]</td>
    <td class="tg-yjjc">You can keep the default</td>
    <td class="tg-yjjc">publicIP</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">LogLevel</td>
    <td class="tg-cly1">Indicates whether to print debug logs.<br> &nbsp;&nbsp;Support: [info, debug]</td>
    <td class="tg-cly1">You can keep the default</td>
    <td class="tg-cly1">info</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">LogRotateSize</td>
    <td class="tg-yjjc">Size of a log split, in MB.</td>
    <td class="tg-yjjc">You can keep the default</td>
    <td class="tg-yjjc">1024</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">LogBackupCount</td>
    <td class="tg-cly1">Maximum number of logs that can be retained. <br> &nbsp;&nbsp;If the number of logs exceeds the maximum, the logs are deleted by time.</td>
    <td class="tg-cly1">You can keep the default</td>
    <td class="tg-cly1">100</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">LogMaxAge</td>
    <td class="tg-yjjc">Maximum number of days for storing logs</td>
    <td class="tg-yjjc">You can keep the default</td>
    <td class="tg-yjjc">7</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">LogCompress</td>
    <td class="tg-cly1">Indicates whether logs are automatically compressed.<br> &nbsp;&nbsp;NOTE: Some logs may be damaged after log compression. <br> &nbsp;&nbsp;Therefore, you are advised not to compress logs.</td>
    <td class="tg-cly1">You can keep the default</td>
    <td class="tg-cly1">FALSE</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-m64s" rowspan="38">aass</td>
    <td class="tg-6vn3">AppgatewayAddress</td>
    <td class="tg-5fx9">AppGateway address for the AASS to access the AppGateway</td>
    <td class="tg-5fx9">If you have configured the ELB service for AppGateway, configurate the IP address of the ELB service.</td>
    <td class="tg-5fx9">127.0.0.1:60003</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">MysqlAddress</td>
    <td class="tg-61g0">mysql address used by aass</td>
    <td class="tg-61g0">Ensure that the aass service can connect to the database.</td>
    <td class="tg-61g0">127.0.0.1:3306</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">MysqlUser</td>
    <td class="tg-5fx9">mysql user used by aass</td>
    <td class="tg-5fx9">Generally, the user is root.</td>
    <td class="tg-5fx9">root</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">MysqlDBName</td>
    <td class="tg-61g0">database name used by aass</td>
    <td class="tg-61g0">It is not created by default. Create it in the database in advance.</td>
    <td class="tg-61g0">aass</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">MysqlPassword</td>
    <td class="tg-5fx9">mysql password used by aass</td>
    <td class="tg-5fx9">Database password encrypted using GCM</td>
    <td class="tg-8v0r">　</td>
    <td class="tg-5fx9">You can run the following command to encrypt the password: /sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
  </tr>
  <tr>
    <td class="tg-61g0">MysqlCharset</td>
    <td class="tg-61g0">Database coding format</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">utf8</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">RedisAddress</td>
    <td class="tg-5fx9">redis address used by aass</td>
    <td class="tg-5fx9">Ensure that the aass service can connect to the redis.</td>
    <td class="tg-5fx9">127.0.0.1:6379</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">RedisPassword</td>
    <td class="tg-61g0">redis password used by aass</td>
    <td class="tg-61g0">redis password encrypted using GCM</td>
    <td class="tg-61g0">　</td>
    <td class="tg-61g0">You can run the following command to encrypt the password: /sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
  </tr>
  <tr>
    <td class="tg-5fx9">RedisDB</td>
    <td class="tg-5fx9">Default Redis partition</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">2</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">LogRotateSize</td>
    <td class="tg-61g0">Size of a log split, in MB.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">1024</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">LogBackupCount</td>
    <td class="tg-5fx9">Maximum number of logs that can be retained. <br> &nbsp;&nbsp;If the number of logs exceeds the maximum, the logs are deleted by time.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">100</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">LogMaxAge</td>
    <td class="tg-61g0">Maximum number of days for storing logs</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">7</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">CloudClientRegion</td>
    <td class="tg-5fx9" rowspan="8">Region where the AASS service is deployed and Service Endpoint about HuaweiCloud Service</td>
    <td class="tg-5fx9" rowspan="8">you can get those from: <br> &nbsp;&nbsp;China: https://developer.huaweicloud.com/endpoint<br> &nbsp;&nbsp;Intl: https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
    <td class="tg-5fx9">cn-north-4</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">CloudClientIamEndpoint</td>
    <td class="tg-61g0">https://iam.cn-north-4.myhuaweicloud.com</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">CloudClientAsEndpoint</td>
    <td class="tg-5fx9">https://as.cn-north-4.myhuaweicloud.com</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">CloudClientEcsEndpoint</td>
    <td class="tg-61g0">https://ecs.cn-north-4.myhuaweicloud.com</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">CloudClientLtsEndpoint</td>
    <td class="tg-5fx9">https://lts.cn-north-4.myhuaweicloud.com</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">CloudClientSmnEndpoint</td>
    <td class="tg-61g0">https://smn.cn-north-4.myhuaweicloud.com</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">CloudClientPodEndpoint</td>
    <td class="tg-5fx9">https://cci.cn-north-4.myhuaweicloud.com</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">CloudClientDNSEndpoint</td>
    <td class="tg-61g0">https://dns.cn-north-4.myhuaweicloud.com</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">CloudPlatformAddr</td>
    <td class="tg-5fx9">ECS metadata includes basic information of an ECS on the cloud platform, such as the ECS ID, hostname, and network information. </td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-8v0r"><a href="http://169.254.169.254/">http://169.254.169.254</a></td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">MonitorDuration</td>
    <td class="tg-61g0">Interval for monitoring asynchronous autoscaling tasks. <br> &nbsp;&nbsp;You can retain the default value.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">10s</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">CleanUpDayBefore</td>
    <td class="tg-5fx9">Expiration time of data in the data table. <br> &nbsp;&nbsp;Data will be deleted after the expiration time, in days.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">14</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">CleanUpPeriodHour</td>
    <td class="tg-61g0">Interval for starting an asynchronous cleanup task, in hours</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">1</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">HealthCheckInterval</td>
    <td class="tg-5fx9">Interval between obtaining the distributed lock and updating the status of the AASS instance, in seconds.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">30</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HostProtect</td>
    <td class="tg-61g0">Whether to enable host protection.<br> &nbsp;&nbsp;Support: [true, false]</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">TRUE</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">HostProtectValue</td>
    <td class="tg-5fx9">Specifies the host protection type if host protection is enabled.<br> &nbsp;&nbsp;You can refer to: <br> &nbsp;&nbsp;China: https://support.huaweicloud.com/api-ecs/zh-cn_topic_0167957246.html#ZH-CN_TOPIC_0167957246__table2373623012315<br> &nbsp;&nbsp;Intl: Not valid for the time being by 2024-01-22.<br> &nbsp;&nbsp;Support: ["ces", "hss", "hss,hss-ent"]<br> &nbsp;&nbsp;<br> &nbsp;&nbsp;</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">hss</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BatchCreatePodNum</td>
    <td class="tg-61g0">Due to the restriction of the APIG during pod creation, <br> &nbsp;&nbsp;to avoid excessive invoking times, <br> &nbsp;&nbsp;if the number of times exceeds the value of this parameter, <br> &nbsp;&nbsp;the pod is hibernated for 10 seconds. <br> &nbsp;&nbsp;You can retain the default value.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">50</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">HttpsAddress</td>
    <td class="tg-5fx9">bind address</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">0.0.0.0</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HttpsPort</td>
    <td class="tg-61g0">bind port</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">9091</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">TakeOverTaskIntervalSeconds</td>
    <td class="tg-5fx9">Interval for taking over asynchronous tasks, in seconds. <br> &nbsp;&nbsp;This parameter is used when a node is faulty.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">60</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HeartBeatTaskIntervalSeconds</td>
    <td class="tg-61g0">Interval for reporting node health status, in seconds, <br> &nbsp;&nbsp;which is used to determine whether a node is normal.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">60</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DeadCheckTaskIntervalSeconds</td>
    <td class="tg-5fx9">Interval for monitoring zombie nodes, in seconds</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">60</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">MaxDeadMinutes</td>
    <td class="tg-61g0">Interval for determining a zombie node, in mintues.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">3</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">InstanceMaximumLimit</td>
    <td class="tg-5fx9">Specifies the maximum number of instances in an AS group.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">500</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BandwidthMaximumLimit</td>
    <td class="tg-61g0">Maximum network bandwidth used when creating a machine.</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">300</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">SupportedVolumeTypes</td>
    <td class="tg-5fx9">Supported volume types.</td>
    <td class="tg-5fx9">You can keep the default</td>
    <td class="tg-5fx9">SATA;SAS;GPSSD;SSD;ESSD</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BandwidthChargingMode</td>
    <td class="tg-61g0">Bandwidth charging mode.<br> &nbsp;&nbsp;Support: [traffic, bandwidth]</td>
    <td class="tg-61g0">You can keep the default</td>
    <td class="tg-61g0">traffic</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-ogp5" rowspan="4">influxdb</td>
    <td class="tg-yjjc">InfluxDBUser</td>
    <td class="tg-yjjc">influxdb user used by aass and appgateway</td>
    <td class="tg-yjjc">You can keep the default</td>
    <td class="tg-yjjc">root</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-fapl">influxDBAddress</td>
    <td class="tg-cly1">influxdb address used by aass and appgateway</td>
    <td class="tg-cly1">Ensure that the aass and appgateway service can connect to the database.</td>
    <td class="tg-cly1">127.0.0.1:8086</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-ym56">InfluxDBPassword</td>
    <td class="tg-yjjc">influxdb passworld used by aass and appgateway
</td>
    <td class="tg-yjjc">Database password encrypted using GCM</td>
    <td class="tg-yjjc"></td>
    <td class="tg-yjjc">You can run the following command to encrypt the password: /sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
  </tr>
  <tr>
    <td class="tg-cly1">inflxuDBName</td>
    <td class="tg-cly1">The AppGateway service is the same as the AASS service. </td>
    <td class="tg-cly1">You can keep the default</td>
    <td class="tg-cly1">autoscaling</td>
    <td class="tg-cly1"></td>
  </tr>
</tbody>
</table>
 




### 1. 解压二进制文件，进入文件目录

```sh
    # TODO: 增加解压命令
    cd /home/gfm
```
应用包内文件目录如下：

```lua
/home/gfm
    |── bin
        |—— appgateway      -- appgateway应用
        |—— aass            -- aass应用
        |—— fleetmanager    -- fleetmanager应用
        |—— auxproxy        -- auxproxy应用
        |—— console         -- console应用
    |—— conf
        |—— init.yaml       -- 配置文件，根据该配置文件生成各个服务组件的配置文件
    |—— sac-gfm             -- 部署辅助脚本

```
### 2. 生成加密所需信息
+ 生成`gcm`加密的`key`与`nonce`，用于后端敏感数据如密码等的存储加密解密，并会自动同步修改`./conf/init.yaml`文件中`cipher`部分的`GCMKey`与`GCMNonce`参数
```sh
    ./sac-gfm cipher --create gcm
    # the work dir is: /home/gfm
    # init conf: ./conf/init.yaml
    # new gcm key: RcFTbhaiY0Hfd1Cg3cdVAjgQ
    # new gcm nonce: kw1uziBJx391fB0Y
```
+  生成用于登录会话的`token`的生成秘钥，并会同步修改`./conf/init.yaml`文件中`fleetmanager`部分的`SessionJwtKey`参数
```sh
    ./sac-gfm cipher --create jwt
    # the work dir is: /home/gfm
    # init conf: ./conf/init.yaml
    # new jwt key: fFMH9nZgM3rpzp2os6RSbIicAeTazBV0I7ZZwFWnIQJFMHVP
```

+  生成`rsa`加密的公钥与私钥，用于控制台与`gfm`服务交互时的敏感信息如密码等的加密与解密, 并会同步修改`./conf/init.yaml`文件中`cipher`部分的`RSAPublicKey`与`RSAPrivateKey`参数
```sh
    ./sac-gfm cipher --create rsa
    # the work dir is: /home/gfm
    # init conf: ./conf/init.yaml
    # succeed to generate new rsa private key and public key into /home/gfm/conf
```

+  生成用于`https`协议的证书文件，并会同步修改`./conf/init.yaml`文件中`cipher`部分的`TlsKey`、`TlsCsr`与`TlsCrt`
```sh
    ./sac-gfm cipher --create tls
    # the work dir is: /home/gfm
    # init conf: ./conf/init.yaml
    # succeed to generate new tls.key / tls.csr / tls.crt into /home/gfm/conf
```

### 3. 修改fleetmanager服务所需的配置文件
#### 配置`mysql`信息
  准备一个`mysql`数据库，并在`mysql`内创建一个名为`fleetmanager`的`DB`，创建具有读写权限的账号与密码，配置修改`./conf/init.yaml`文件中的以下参数：
  ```yaml
    fleetmanager: 
        MysqlAddress: 127.0.0.1:3306
        MysqlUser: root
        MysqlDBName: fleetmanager
        MysqlPassword: {加密后的mysql数据库密码}
  ```
  **其中**：`MysqlPassword`参数需要配置使用`gcm`加密后的信息，加密方式与参考如下，默认使用`./conf/init.yaml`内的`GCMKey`与`GCMNonce`，假设数据库密码为：Sacdb@1234
  ```sh
    ./sac-gfm cipher --mode encode --method gcm --text {数据库密码}
    # the work dir is: /home/gfm
    # init conf: ./conf/init.yaml
    # mode: encode method: gcm 
    # cipher encode use gcm
    # gcm key: knPU2r9gNvmU4dAhh9zHnTiR
    # gcm nonce: YJkVk7cNF9cMRjJL
    # cipher encode result: 4NqAxVUCD282kMQOroRahZSi5JioKOT3EQs for Sacdb@1234
  ```
#### 配置`redis`信息
  准备一个redis缓存数据库，配置修改`./conf/init.yaml`文件中的以下参数
  ```yaml
    fleetmanager: 
        RedisAddress: 127.0.0.1:6379
        RedisPassword: {加密后的redis密码}
        RedisDB: 1
  ```
  **其中**：`RedisPassword`参数需要使用`gcm`加密，加密方式请参考fleetmanager的[数据库密码加密](#配置mysql信息)操作

#### 配置HuaweiCloud信息：
  确定将GFM应用部署在华为云上的哪一个region，以部署在中国站北京四为例，配置修改`./conf/init.yaml`文件中的以下参数
  其中region信息，可以参考：
  中国站: https://developer.huaweicloud.com/endpoint
  国际站: https://developer.huaweicloud.com/intl/en-us/endpoint?all
  ```yaml
    fleetmanager: 
        # If multiple regions are deployed, you can enter multiple values. 
        # Otherwise, only one value is required.
        # You can refer: 
        # China: https://developer.huaweicloud.com/endpoint
        # Intl: https://developer.huaweicloud.com/intl/en-us/endpoint?all
        SupportRegions:
            - cn-north-4
        # List of public ECS images that have been verified. 
        # If other images are available, you can change them to the ones you want.
        SupportPublicImage:
            - region: cn-north-4
              image: "CentOS 7.2 64bit,CentOS 7.9 64bit,CentOS 8.1 64bit,Ubuntu 22.04 server 64bit,CentOS 7.6 64bit for Tenant 20210525"
        # List of Docker images of verified pods. 
        # If you are sure that other images are available, you can change them to the ones you want.
        SupportDockerImage:
            - region: cn-north-4
              dockerOS: "swr.cn-north-7.myhuaweicloud.com/game/centos-super:v1,centos:7.6.1810,centos:7.2.1511"
        # The EIP type supported by each region may be different. 
        # You need to manually configure the EIP type after querying the EIP type on the cloud service.
        # You can refer to: 
        # Table 4 in China:  https://support.huaweicloud.com/api-eip/eip_api_0001.html
        # Table 3 in Intl: https://support.huaweicloud.com/intl/en-us/api-eip/eip_api_0001.html
        EipType:
            - region: cn-north-4
              supportEipType: 5_bgp
        # Huawei Cloud Private DNS Server Addresses.
        # You can get it from Table 1 in:
        # China: https://support.huaweicloud.com/dns_faq/dns_faq_002.html
        # Intl: https://support.huaweicloud.com/intl/en-us/dns_faq/dns_faq_002.html
        DnsConfig:  
            - region: cn-north-4
              dnsServer: "100.125.1.250,100.125.129.250"
        # This configuration item is used if you want to use LTS to manage logs.
        # This configuration item is used to automatically install ICAgent on HUAWEI CLOUD hosts for log collection. 
        # You need to query the configuration address.
        # You can log in to the LTS service console, Select the region to be deployed, 
        # Select: Host Management -> Install ICAgent.
        # You can get it from the step 2 by accessip in command given.
        LtsIps:
            - region: cn-north-4
              ltsIp: "100.125.12.150"
        # If the region information is not provided when creating fleet, the value will be used.
        DefaultRegion: cn-north-4
        # Service Endpoint about HuaweiCloud Service, you can get those from: 
        # China: https://developer.huaweicloud.com/endpoint
        # Intl: https://developer.huaweicloud.com/intl/en-us/endpoint?all
        ServiceEndpoint:
            IamService:
              - region: cn-north-4
                endPoint: https://iam.cn-north-4.myhuaweicloud.com
            # Address used by the AASS to provide services for external systems. 
            # FleetManager can access the AASS through this address.
            # TODO: modify
            AASS: 
              - region: cn-north-4
                endPoint: {aass ipAddress}:9091
            # Address used by the Appgateway to provide services for external systems. 
            # FleetManager can access the AASS through this address.
            AppGateway:
              - region: cn-north-4
                endPoint: {appgateway ipAddress}:60003
            VpcService:
              - region: cn-north-4
                endPoint: https://vpc.cn-north-4.myhuaweicloud.com
            ObsService:
              - region: cn-north-4
                endPoint: https://obs.cn-north-4.myhuaweicloud.com
            ImsService:
              - region: cn-north-4
                endPoint: https://ims.cn-north-4.myhuaweicloud.com
            EcsService:
              - region: cn-north-4
                endPoint: https://ecs.cn-north-4.myhuaweicloud.com
            SwrService:
              - region: cn-north-4
                endPoint: https://swr.cn-north-4.myhuaweicloud.com
            CciService:
              - region: cn-north-4
                endPoint: https://cci.cn-north-4.myhuaweicloud.com
            EvsService:
              - region: cn-north-4
                endPoint: https://evs.cn-north-4.myhuaweicloud.com
            CesService:
              - region: cn-north-4
                endPoint: https://ces.cn-north-4.myhuaweicloud.com
            DnsService:
            - region: cn-north-4
              endpoint: "https://dns.cn-north-4.myhuaweicloud.com"
            EpsService:
            - region: cn-north-4
              endpoint: "https://eps.cn-north-4.myhuaweicloud.com"
  ```
#### 配置应用包的配置信息
在gfm正常使用之前，您需要创建或选择已有的obs桶，需要将以下文件上传至OBS桶中，建议OBS桶配置在与GFM部署的region一致，以obs桶gameflexmatch为例，该桶位于北京四区域：
+ ./bin/image_env.sh
+ ./bin/docker_image_env.sh
+ ./bin/auxproxy.zip:
  
确定上传后的路径，并配置修改`./conf/init.yaml`文件中的以下参数:
```yaml
    fleetmanager: 
        BuildScriptPath: gameflexmatch/image_env.sh
        BuildDockerScriptPath: gameflexmatch/docker_image_env.sh
        AuxproxyPath: gameflexmatch/auxproxy.zip
        ProfilesStorageRegion: cn-north-4
```
#### 配置默认的安全组规则信息：
该配置项用于配置fleet内的fleet集群所需的必要安全组入方向规则，以保证服务可以正常使用
配置修改`./conf/init.yaml`文件中的以下参数:

```yaml
    InternalInboundPermissions:
        - protocol: "TCP"
          ipRange: "{appgateway public/private ip}/32"
          fromPort: 60001
          toPort: 60001
        - protocol: "TCP"
          ipRange: "{aass public/private ip}/32"
          fromPort: 60001
          toPort: 60001
```
注意：
1. 需要配置appgateway与aass所有节点的信息
2. 安全组的入方向规则的ipRange的IP类型根据appgateway与aass访问auxproxy的IP类型来决定

至此，fleetmanager服务组件的必需参数已配置完成

### 4. 修改aass服务所需的配置文件
#### 配置`mysql`信息
  准备一个`mysql`数据库，并在`mysql`内创建一个名为`aass`的`DB`，创建具有读写权限的账号与密码，配置修改`./conf/init.yaml`文件中的以下参数：
  ```yaml
    aass: 
        MysqlAddress: 127.0.0.1:3306
        MysqlUser: root
        MysqlDBName: aass
        MysqlPassword: {加密后的mysql数据库密码}
  ```
  **其中**：`MysqlPassword`参数需要使用`gcm`加密，加密方式请参考fleetmanager的[数据库密码加密](#配置mysql信息)操作
#### 配置`redis`信息
  准备一个redis缓存数据库，配置修改`./conf/init.yaml`文件中的以下参数
  ```yaml
    aass: 
        RedisAddress: 127.0.0.1:6379
        RedisPassword: {加密后的redis密码}
        RedisDB: 2
  ```
  **其中**：`RedisPassword`参数需要使用`gcm`加密，加密方式请参考fleetmanager的[数据库密码加密](#配置mysql信息)操作
#### 配置`influxDB`信息
准备一个influxdb时序数据库，配置修改`./conf/init.yaml`文件中的以下参数
```yaml
    influxdb: 
        InfluxDBUser: 127.0.0.1:6379
        influxDBAddress: {加密后的redis密码}
        InfluxDBPassword: 
        inflxuDBName: autoscaling
```
**其中**：
1. `RedisPassword`参数需要使用`gcm`加密，加密方式请参考fleetmanager的[数据库密码加密](#配置mysql信息)操作
2. aass与appgateway共用一个influxDB数据库

#### 配置华为云的区域信息
根据aass服务部署的region，以北京四(cn-north-4)为例，配置修改`./conf/init.yaml`文件中的以下参数：
```yaml
    aass:
        CloudClientRegion: cn-north-4
        CloudClientIamEndpoint: https://iam.cn-north-4.myhuaweicloud.com
        CloudClientAsEndpoint: https://as.cn-north-4.myhuaweicloud.com
        CloudClientEcsEndpoint: https://ecs.cn-north-4.myhuaweicloud.com
        CloudClientLtsEndpoint: https://lts.cn-north-4.myhuaweicloud.com
        CloudClientSmnEndpoint: https://smn.cn-north-4.myhuaweicloud.com
        CloudClientPodEndpoint: https://pod.cn-north-4.myhuaweicloud.com
        CloudClientDNSEndpoint: https://dns.cn-north-4.myhuaweicloud.com
```
至此，aass服务组件的必需参数已配置完成

### 4. 修改appgateway服务所需的配置文件
#### 配置`mysql`信息
  准备一个`mysql`数据库，并在`mysql`内创建一个名为`aass`的`DB`，创建具有读写权限的账号与密码，配置修改`./conf/init.yaml`文件中的以下参数：
  ```yaml
    appgateway: 
        MysqlAddress: 127.0.0.1:3306
        MysqlUser: root
        MysqlDBName: appgateway
        MysqlPassword: {加密后的mysql数据库密码}
  ```
  **其中**：`MysqlPassword`参数需要使用`gcm`加密，加密方式请参考fleetmanager的[数据库密码加密](#配置mysql信息)操作
#### 配置`redis`信息
  准备一个redis缓存数据库，配置修改`./conf/init.yaml`文件中的以下参数
  ```yaml
    appgateway: 
        RedisAddress: 127.0.0.1:6379
        RedisPassword: {加密后的redis密码}
        RedisDB: 2
  ```
  **其中**：`RedisPassword`参数需要使用`gcm`加密，加密方式请参考fleetmanager的[数据库密码加密](#配置mysql信息)操作
#### 配置`influxDB`信息
已在aass的配置步骤中进行配置，该步骤无需操作

至此，已完成所有服务配置文件的基础配置



