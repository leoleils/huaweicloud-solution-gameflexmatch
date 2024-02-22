# Introduction to Configuration Files
GFM基于`conf/init.yaml`配置文件生成各个服务组件所需的配置文件，这里将对配置文件中的参数以及如何配置做出说明。

**其中：表中红色的参数需要根据实际情况修改，其他可以保持默认**

## 参数介绍
```html
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
    <td class="tg-yjjc">密码等敏感数据的加解密长度为24个字符</td>
    <td class="tg-yjjc" rowspan="7">已经默认为您生成。如果有必要，可以重新生成它。</td>
    <td class="tg-yjjc">Fi1xOkgXvevHlTR5SIHHXAb2</td>
    <td class="tg-yjjc" rowspan="2">如果需要重新生成，可以执行：./sac-gfm cipher --create gcm --init-conf conf/init.yaml</td>
  </tr>
  <tr>
    <td class="tg-cly1">GCMNonce</td>
    <td class="tg-cly1">密码等敏感数据加解密的16个字符</td>
    <td class="tg-cly1">dlbaGmj9O32CFbTs</td>
  </tr>
  <tr>
    <td class="tg-yjjc">RSAPublicKey</td>
    <td class="tg-yjjc" rowspan="2">控制台与服务之间的密码传输加密(RSA)</td>
    <td class="tg-yjjc">./conf/public.key</td>
    <td class="tg-yjjc" rowspan="2">如果需要重新生成，可以执行exec: ./sac-gfm cipher --create rsa --init-conf conf/init.yaml，密钥需要和控制台中的一致。默认公钥已经随前端软件包一起编译。如果重新生成密钥，则需要重新编译前端软件包。关于密钥的编译方法，可以参考：。</td>
  </tr>
  <tr>
    <td class="tg-cly1">RSAPrivateKey</td>
    <td class="tg-cly1">./conf/private.key</td>
  </tr>
  <tr>
    <td class="tg-yjjc">TlsKey</td>
    <td class="tg-yjjc" rowspan="3">HTTPS协议的自签名证书</td>
    <td class="tg-yjjc">./conf/tls.key</td>
    <td class="tg-yjjc" rowspan="3">如果需要重新生成，可以执行：./sac-gfm cipher --create tls --init-conf conf/init.yaml</td>
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
    <td class="tg-61g0">fleetmanager使用的mysql地址</td>
    <td class="tg-61g0">确保手机服务可以连接到数据库。</td>
    <td class="tg-61g0">127.0.0.1:3306</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">MysqlUser</td>
    <td class="tg-5fx9">fleetmanager使用的mysql用户</td>
    <td class="tg-5fx9">一般情况下，用户为root。</td>
    <td class="tg-5fx9">root</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">MysqlDBName</td>
    <td class="tg-61g0">fleetmanager使用的数据库名</td>
    <td class="tg-61g0">默认情况下不会创建它，需要提前在数据库中创建。</td>
    <td class="tg-61g0">fleetmanager</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">MysqlPassword</td>
    <td class="tg-5fx9">mysql使用的密码</td>
    <td class="tg-5fx9">GCM加密后的数据库密码</td>
    <td class="tg-8v0r">　</td>
    <td class="tg-5fx9">可以使用如下命令加密：/sac-gfm cipher --mode encode --method gcm --text{数据库密码}</td>
  </tr>
  <tr>
    <td class="tg-61g0">MysqlCharset</td>
    <td class="tg-61g0">数据库编码格式</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">utf8</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">RedisAddress</td>
    <td class="tg-5fx9">fleetmanager使用的redis地址</td>
    <td class="tg-5fx9">确保fleetmanager可以连接到redis。</td>
    <td class="tg-5fx9">127.0.0.1:6379</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">RedisPassword</td>
    <td class="tg-61g0">fleetmanager使用的redis密码</td>
    <td class="tg-61g0">redis密码gcm加密</td>
    <td class="tg-61g0">　</td>
    <td class="tg-61g0">可以使用如下命令加密：/sac-gfm cipher --mode encode --method gcm --text{数据库密码}</td>
  </tr>
  <tr>
    <td class="tg-5fx9">RedisDB</td>
    <td class="tg-5fx9">Redis默认分区</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">1</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">LogRotateSize</td>
    <td class="tg-61g0">日志分割的大小，单位为MB。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">1024</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">LogBackupCount</td>
    <td class="tg-5fx9">日志最大保留条数。如果日志超过最大条数，则按时间删除日志。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">100</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">LogMaxAge</td>
    <td class="tg-61g0">日志保存的最大天数。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">7</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">SupportRegions</td>
    <td class="tg-5fx9">如果部署了多个区域，可以输入多个值。否则，只需要一个值。</td>
    <td class="tg-5fx9">您可以参考：<br> &nbsp;&nbsp;中国站：https://developer.huaweicloud.com/endpoint<br> &nbsp;&nbsp;国际站：https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
    <td class="tg-5fx9">cn-north-4</td>
    <td class="tg-5fx9">该配置项需要配置为YAML列表。</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultLoginPassword</td>
    <td class="tg-61g0">仅在首次登录系统时使用。初始化完成后，密码会被修改为新的密码。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-l010">Gfm@2024</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">WorkflowPath</td>
    <td class="tg-5fx9">工作流配置文件路径</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">./conf/workflow/</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">SupportPublicImage</td>
    <td class="tg-61g0">已校验的公有云服务器镜像列表。<br> &nbsp;&nbsp;如果有其他镜像，可以将其更改为你想要的。</td>
    <td class="tg-61g0">用户可以自定义区域，其他参数保持默认值。</td>
    <td class="tg-61g0">SupportPublicImage: <br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      image: "CentOS 7.2 64bit,CentOS 7.9 64bit,CentOS 8.1 64bit,Ubuntu 22.04 server 64bit,CentOS 7.6 64bit for Tenant 20210525"</td>
    <td class="tg-61g0">该配置项需要配置为YAML列表。</td>
  </tr>
  <tr>
    <td class="tg-6vn3">SupportDockerImage</td>
    <td class="tg-5fx9">已验证Pod的Docker镜像列表。<br> &nbsp;&nbsp;如果你确定还有其他的镜像，你可以将它们更改为你想要的。</td>
    <td class="tg-5fx9">用户可以自定义区域，其他参数保持默认值。</td>
    <td class="tg-5fx9">SupportDockerImage: <br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      dockerOS: "swr.cn-north-7.myhuaweicloud.com/game/centos-super:v1,centos:7.6.1810,centos:7.2.1511"</td>
    <td class="tg-5fx9">该配置项需要配置为YAML列表。</td>
  </tr>
  <tr>
    <td class="tg-poke">EipType</td>
    <td class="tg-61g0">每个区域支持的弹性IP类型可能不同。<br> &nbsp;&nbsp;在云服务上查询到弹性IP类型后，需要手动配置弹性IP类型。</td>
    <td class="tg-61g0">您可以参考：<br> &nbsp;&nbsp;中国区表4：https://support.huaweicloud.com/api-eip/eip_api_0001.html<br>  &nbsp;国际站表3:https://support.huaweicloud.com/intl/en-us/api-eip/eip_api_0001.html<br> &nbsp;&nbsp;如果有更多，应该用逗号隔开，例如：“eip-type1,eip-type2”</td>
    <td class="tg-61g0">EipType:<br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      supportEipType: 5_bgp,5_sbgp</td>
    <td class="tg-61g0">该配置项需要配置为YAML列表。</td>
  </tr>
  <tr>
    <td class="tg-6vn3">DnsConfig</td>
    <td class="tg-5fx9">华为云私有DNS服务器地址。</td>
    <td class="tg-5fx9">您可以从表1中获取：<br> &nbsp;&nbsp;中国站：https://support.huaweicloud.com/dns_faq/dns_faq_002.html<br> &nbsp;&nbsp;国际站：https://support.huaweicloud.com/intl/en-us/dns_faq/dns_faq_002.html</td>
    <td class="tg-5fx9">DnsConfig: <br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      dnsServer: 100.125.1.250,100.125.129.250</td>
    <td class="tg-5fx9">该配置项需要配置为YAML列表。</td>
  </tr>
  <tr>
    <td class="tg-poke">LtsIps</td>
    <td class="tg-61g0">当您需要使用日志服务管理日志时，需要使用此配置项。<br> &nbsp;&nbsp;该配置项用于在华为云主机上自动安装ICAgent，用于日志采集。</td>
    <td class="tg-61g0">需要查询配置地址。<br> &nbsp;&nbsp;您可以登录云日志服务控制台，选择需要部署的区域。<br> &nbsp;&nbsp;选择：主机管理-&gt;安装ICAgent。<br> &nbsp;&nbsp;你可以通过在给定的命令中的accessip从第2步中获得它。</td>
    <td class="tg-61g0">LtsIps:<br> &nbsp;&nbsp;    - region: cn-north-4<br> &nbsp;&nbsp;      ltsIp: 100.125.12.150</td>
    <td class="tg-61g0">该配置项需要配置为YAML列表。</td>
  </tr>
  <tr>
    <td class="tg-6vn3">DefaultRegion</td>
    <td class="tg-5fx9">如果创建车队时没有提供区域信息，则使用该参数值。</td>
    <td class="tg-5fx9">　</td>
    <td class="tg-5fx9">cn-north-4</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultFleetProtectPolicy</td>
    <td class="tg-61g0">车队的默认保护策略。<br> &nbsp;&nbsp;该配置项在缩容时会使用。<br> &nbsp;&nbsp;支持：[TIME_LIMIT_PROTECTION, NO_PROTECTION]<br> &nbsp;&nbsp;NO_PROTECTION可能会影响正在玩游戏的玩家，不建议<br> &nbsp;&nbsp;TIME_LIMIT_PROTECTION在保护期内，机器不会被强制回收。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">TIME_LIMIT_PROTECTION</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultProtectTimeLimit</td>
    <td class="tg-5fx9">车队缩容保护时间，单位为分钟。<br> &nbsp;&nbsp;仅当本命令中的Protection Policy参数为TIME_LIMIT_PROTECTION时，该参数有效。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">30</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultSessionTimeoutSeconds</td>
    <td class="tg-61g0">创建会话时，如果会话在该时间后还没有被激活，则会话将设置为错误。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">60</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultMaxSessionNumPerProcess</td>
    <td class="tg-5fx9">配置每个进程默认启动的会话数。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">50</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultProcessNumPerInstance</td>
    <td class="tg-61g0">配置每个实例默认启动的进程数。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">50</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultBandwidth</td>
    <td class="tg-5fx9">默认带宽</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">5</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultBandwidthChargingMode</td>
    <td class="tg-61g0">默认带宽计费方式</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">traffic</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultDiskSize</td>
    <td class="tg-5fx9">默认磁盘大小</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">40</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultVolumeType</td>
    <td class="tg-61g0">默认卷类型</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">SATA</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultEipShareType</td>
    <td class="tg-5fx9">带宽共享类型。<br> &nbsp;&nbsp;共享类型枚举值：<br> &nbsp;&nbsp;PER，表示独占使用。WHOLE：表示共享。<br> &nbsp;&nbsp;支持：[每，全部]<br> &nbsp;&nbsp;默认值：PER</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">PER</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">FleetCidr</td>
    <td class="tg-61g0">当创建车队时需要创建VPC时，需要配置该配置项。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">FleetCidr:<br> &nbsp;&nbsp;    from: 10.100.0.0<br> &nbsp;&nbsp;    to: 10.240.0.0<br> &nbsp;&nbsp;    vpcNetMask: 19<br> &nbsp;&nbsp;    subnetNetMask: 20</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">InternalInboundPermissions</td>
    <td class="tg-5fx9">创建fleet时，默认会创建一个安全组。<br> &nbsp;&nbsp;下面列出了默认需要配置的安全组规则。</td>
    <td class="tg-5fx9">ipRange需要配置为所有AppGateway节点的IP地址。<br> &nbsp;&nbsp;例如，AppGateway有两个节点10.100.0.1和10.100.0.2，则需要将ipRange的值设置为示例值。<br> &nbsp;&nbsp;其他参数可以保持默认配置。</td>
    <td class="tg-5fx9">InternalInboundPermissions:<br> &nbsp;&nbsp;    - protocol: "TCP"<br> &nbsp;&nbsp;      ipRange: 10.100.0.1/32<br> &nbsp;&nbsp;      fromPort: 60001<br> &nbsp;&nbsp;      toPort: 60001<br> &nbsp;&nbsp;    - protocol: "TCP"<br> &nbsp;&nbsp;      ipRange: 10.100.0.1/32<br> &nbsp;&nbsp;      fromPort: 60001<br> &nbsp;&nbsp;      toPort: 60001</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BuildImageRef</td>
    <td class="tg-61g0">制作镜像时使用的默认操作系统</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">CentOS 7.2 64bit</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">BuildFlavor</td>
    <td class="tg-5fx9">创建镜像时使用的默认规格</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">s6.large.2</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BuildBandwidth</td>
    <td class="tg-61g0">创建镜像时的默认带宽</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">100</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">BuildScriptPath</td>
    <td class="tg-5fx9">创建ECS镜像脚本</td>
    <td class="tg-5fx9">上传镜像打包脚本到OBS桶中。<br> &nbsp;&nbsp;例如，OBS桶为gfm，则配置为类似示例的值。</td>
    <td class="tg-5fx9">gfm/image_env.sh</td>
    <td class="tg-5fx9">将bin/image_env.sh脚本上传到OBS的gfm桶中。</td>
  </tr>
  <tr>
    <td class="tg-poke">BuildDockerScriptPath</td>
    <td class="tg-61g0">制作POD镜像脚本</td>
    <td class="tg-61g0">将打包docker镜像的脚本上传到OBS桶中。<br> &nbsp;&nbsp;例如，OBS桶为gfm，则配置为类似示例的值。</td>
    <td class="tg-61g0">gfm/docker_image_env.sh</td>
    <td class="tg-61g0">将bin/docker_image_env.sh脚本上传到OBS的gfm桶中。</td>
  </tr>
  <tr>
    <td class="tg-6vn3">AuxproxyPath</td>
    <td class="tg-5fx9">auxproxy服务应用包</td>
    <td class="tg-5fx9">上传应用包到OBS桶中。<br> &nbsp;&nbsp;例如，OBS桶为gfm，则配置为类似示例的值。</td>
    <td class="tg-5fx9">gfm/auxproxy.zip</td>
    <td class="tg-5fx9">将bin/auxproxy.zip上传到OBS的gfm桶中。</td>
  </tr>
  <tr>
    <td class="tg-61g0">ImageDiskSize</td>
    <td class="tg-61g0">镜像磁盘大小</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">40</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">ServerSessionBackupDays</td>
    <td class="tg-5fx9">历史数据失效时间，单位为天。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">1</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultScalingInCoolDownInterval</td>
    <td class="tg-61g0">缩容周期，单位为分钟。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">10</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultGroupMaxSize</td>
    <td class="tg-5fx9" rowspan="3">伸缩组的最大实例数、最小实例数和期望实例数</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">1</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">DefaultGroupMinSize</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">1</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DefaultGroupDesiredSize</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">1</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">ServiceEndpoint.AASS</td>
    <td class="tg-61g0">AASS对外提供服务的地址。FleetManager通过该地址可以访问AASS。</td>
    <td class="tg-61g0">如果已为AASS配置了ELB服务，请配置ELB服务的IP地址。</td>
    <td class="tg-61g0">https://127.0.0.1:9091</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">ServiceEndpoint.AppGateway</td>
    <td class="tg-5fx9">Appgateway对外提供服务的地址。<br> &nbsp;&nbsp;FleetManager可以通过该地址访问Appgateway。</td>
    <td class="tg-5fx9">如果appgateway配置了ELB服务，请配置ELB服务的IP地址。</td>
    <td class="tg-5fx9">https://127.0.0.1:60003</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">ServiceEndpoint.IamService</td>
    <td class="tg-61g0" rowspan="11">华为云服务的服务端点</td>
    <td class="tg-61g0" rowspan="11">您可以从：<br> &nbsp;&nbsp;中国：https://developer.huaweicloud.com/endpoint<br> &nbsp;&nbsp;国际地址：https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
    <td class="tg-61g0">https://iam.cn-north-4.myhuaweicloud.com</td>
    <td class="tg-61g0" rowspan="11">示例值省略了区域。在实际应用中，区域应设置为：<br> &nbsp;&nbsp;区域：cn-north-4<br> &nbsp;&nbsp;端点：*****</td>
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
    <td class="tg-5fx9">类似于用于生成或验证登录会话的密钥。</td>
    <td class="tg-5fx9">可以选择24位或32位的随机字符串。默认生成。</td>
    <td class="tg-5fx9">fFMH9nZgM3rpzp2os6RSbIicAeTazBV0I7ZZwFWnIQJFMHVP</td>
    <td class="tg-5fx9">如果需要重新生成，可以执行：./sac-gfm cipher --create jwt --init-conf conf/init.yaml</td>
  </tr>
  <tr>
    <td class="tg-61g0">SessionJwtTokenLifeTimeSecond</td>
    <td class="tg-61g0">登录状态令牌的超时时间。<br> &nbsp;&nbsp;如果超过超时时间，需要重新登录</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">7200</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">TakeOverTaskIntervalSeconds</td>
    <td class="tg-5fx9">接管异步任务的时间间隔，单位为秒。<br> &nbsp;&nbsp;节点故障时使用。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">60</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HttpsAddress</td>
    <td class="tg-61g0">监听地址</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">0.0.0.0</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">HttpsPort</td>
    <td class="tg-5fx9">监听端口</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">31002</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HeartBeatTaskIntervalSeconds</td>
    <td class="tg-61g0">节点健康状态上报周期，单位为秒。<br> &nbsp;&nbsp;用于判断节点是否正常。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">60</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DeadCheckTaskIntervalSeconds</td>
    <td class="tg-5fx9">监控僵尸节点的时间间隔，单位为秒。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">60</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">MaxDeadMinutes</td>
    <td class="tg-61g0">判断僵尸节点的时间间隔，单位为分钟。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">3</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-ogp5" rowspan="20">appgateway</td>
    <td class="tg-ym56">AASSAddress</td>
    <td class="tg-yjjc">AppGateway访问AASS的地址</td>
    <td class="tg-yjjc">如果已为AASS配置了ELB服务，请配置ELB服务的IP地址。</td>
    <td class="tg-yjjc">127.0.0.1:9091</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-fapl">MysqlAddress</td>
    <td class="tg-cly1">appgateway使用的mysql地址</td>
    <td class="tg-cly1">确保appgateway服务可以连接到数据库。</td>
    <td class="tg-cly1">127.0.0.1:3306</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">MysqlUser</td>
    <td class="tg-yjjc">appgateway使用的mysql用户</td>
    <td class="tg-yjjc">一般情况下，用户为root。</td>
    <td class="tg-yjjc">root</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-fapl">MysqlDBName</td>
    <td class="tg-cly1">appgateway使用的数据库名</td>
    <td class="tg-cly1">默认情况下不会创建它，需要提前在数据库中创建。</td>
    <td class="tg-cly1">appgateway</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-ym56">MysqlPassword</td>
    <td class="tg-yjjc">appgateway使用的mysql密码</td>
    <td class="tg-yjjc">GCM加密后的数据库密码</td>
    <td class="tg-x3ds"></td>
    <td class="tg-yjjc">可以使用如下命令加密：/sac-gfm cipher --mode encode --method gcm --text{数据库密码}</td>
  </tr>
  <tr>
    <td class="tg-fapl">RedisAddress</td>
    <td class="tg-cly1">appgateway使用的redis地址</td>
    <td class="tg-cly1">确保appgateway服务可以连接到redis。</td>
    <td class="tg-cly1">127.0.0.1:6379</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-ym56">RedisPassword</td>
    <td class="tg-yjjc">appgateway使用的redis密码</td>
    <td class="tg-yjjc">redis密码gcm加密</td>
    <td class="tg-yjjc"></td>
    <td class="tg-yjjc">可以使用如下命令加密：/sac-gfm cipher --mode encode --method gcm --text{数据库密码}</td>
  </tr>
  <tr>
    <td class="tg-cly1">RedisDB</td>
    <td class="tg-cly1">Redis默认分区</td>
    <td class="tg-cly1">您可以保留默认值</td>
    <td class="tg-cly1">3</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">CleanStrategy</td>
    <td class="tg-yjjc">是否清理数据库中的数据<br> &nbsp;&nbsp;支持：【关，开】</td>
    <td class="tg-yjjc">您可以保留默认值</td>
    <td class="tg-yjjc">on</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">CleanupDays</td>
    <td class="tg-cly1">从数据库中删除进程中的数据和会话备份表，默认14天</td>
    <td class="tg-cly1">您可以保留默认值</td>
    <td class="tg-cly1">14</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">BackupDays</td>
    <td class="tg-yjjc">从数据库备份进程中的数据和会话运行时表，默认为3天</td>
    <td class="tg-yjjc">您可以保留默认值</td>
    <td class="tg-yjjc">3</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">HttpsAddress</td>
    <td class="tg-cly1">绑定地址</td>
    <td class="tg-cly1">您可以保留默认值</td>
    <td class="tg-cly1">0.0.0.0</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">HttpsPort</td>
    <td class="tg-yjjc">绑定端口</td>
    <td class="tg-yjjc">您可以保留默认值</td>
    <td class="tg-yjjc">60003</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">DeployModel</td>
    <td class="tg-cly1">Appgateway启动方式<br> &nbsp;&nbsp;支持：【单实例，多实例】</td>
    <td class="tg-cly1">您可以保留默认值</td>
    <td class="tg-cly1">multi-instances</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">AuxproxyIpType</td>
    <td class="tg-yjjc">配置Appgateway通过公网IP或私网IP连接AuxProxy。<br> &nbsp;&nbsp;支持：【公网IP，私网IP】</td>
    <td class="tg-yjjc">您可以保留默认值</td>
    <td class="tg-yjjc">publicIP</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">LogLevel</td>
    <td class="tg-cly1">是否打印调试日志。<br> &nbsp;&nbsp;支持：【信息，调试】</td>
    <td class="tg-cly1">您可以保留默认值</td>
    <td class="tg-cly1">info</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">LogRotateSize</td>
    <td class="tg-yjjc">日志分割的大小，单位为MB。</td>
    <td class="tg-yjjc">您可以保留默认值</td>
    <td class="tg-yjjc">1024</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">LogBackupCount</td>
    <td class="tg-cly1">日志最大保留条数。<br> &nbsp;&nbsp;如果日志超过最大条数，则按时间删除日志。</td>
    <td class="tg-cly1">您可以保留默认值</td>
    <td class="tg-cly1">100</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-yjjc">LogMaxAge</td>
    <td class="tg-yjjc">日志最大保存天数</td>
    <td class="tg-yjjc">您可以保留默认值</td>
    <td class="tg-yjjc">7</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-cly1">LogCompress</td>
    <td class="tg-cly1">日志是否自动压缩。<br> &nbsp;&nbsp;注意：日志压缩后可能会损坏部分日志。<br> &nbsp;&nbsp;因此，建议不要压缩日志。</td>
    <td class="tg-cly1">您可以保留默认值</td>
    <td class="tg-cly1">FALSE</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-m64s" rowspan="38">aass</td>
    <td class="tg-6vn3">AppgatewayAddress</td>
    <td class="tg-5fx9">AASS访问AppGateway的AppGateway地址</td>
    <td class="tg-5fx9">如果AppGateway配置了ELB服务，请配置ELB服务的IP地址。</td>
    <td class="tg-5fx9">127.0.0.1:60003</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">MysqlAddress</td>
    <td class="tg-61g0">aass使用的mysql地址</td>
    <td class="tg-61g0">确保aass服务可以连接到数据库。</td>
    <td class="tg-61g0">127.0.0.1:3306</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">MysqlUser</td>
    <td class="tg-5fx9">aass使用的mysql用户</td>
    <td class="tg-5fx9">一般情况下，用户为root。</td>
    <td class="tg-5fx9">root</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">MysqlDBName</td>
    <td class="tg-61g0">aass使用的数据库名</td>
    <td class="tg-61g0">默认情况下不会创建它。提前在数据库中创建。</td>
    <td class="tg-61g0">aass</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">MysqlPassword</td>
    <td class="tg-5fx9">aass使用的mysql密码</td>
    <td class="tg-5fx9">GCM加密后的数据库密码</td>
    <td class="tg-8v0r">　</td>
    <td class="tg-5fx9">可以使用如下命令加密：/sac-gfm cipher --mode encode --method gcm --text{数据库密码}</td>
  </tr>
  <tr>
    <td class="tg-61g0">MysqlCharset</td>
    <td class="tg-61g0">数据库编码格式</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">utf8</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">RedisAddress</td>
    <td class="tg-5fx9">aass使用的redis地址</td>
    <td class="tg-5fx9">确保aass服务可以连接到redis。</td>
    <td class="tg-5fx9">127.0.0.1:6379</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-poke">RedisPassword</td>
    <td class="tg-61g0">aass使用的redis密码</td>
    <td class="tg-61g0">redis密码gcm加密</td>
    <td class="tg-61g0">　</td>
    <td class="tg-61g0">可以使用如下命令加密：/sac-gfm cipher --mode encode --method gcm --text{数据库密码}</td>
  </tr>
  <tr>
    <td class="tg-5fx9">RedisDB</td>
    <td class="tg-5fx9">Redis默认分区</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">2</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">LogRotateSize</td>
    <td class="tg-61g0">日志分割的大小，单位为MB。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">1024</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">LogBackupCount</td>
    <td class="tg-5fx9">日志最大保留条数。<br> &nbsp;&nbsp;如果日志超过最大条数，则按时间删除日志。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">100</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">LogMaxAge</td>
    <td class="tg-61g0">日志最大保存天数</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">7</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-6vn3">CloudClientRegion</td>
    <td class="tg-5fx9" rowspan="8">AASS服务所在的Region和华为云服务的Service Endpoint</td>
    <td class="tg-5fx9" rowspan="8">您可以从：<br> &nbsp;&nbsp;中国：https://developer.huaweicloud.com/endpoint<br> &nbsp;&nbsp;国际地址：https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
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
    <td class="tg-5fx9">弹性云服务器元数据包括弹性云服务器在云平台上的基本信息，如弹性云服务器ID、主机名、网络信息等。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-8v0r"><a href="http://169.254.169.254/">http://169.254.169.254</a></td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">MonitorDuration</td>
    <td class="tg-61g0">异步弹性伸缩任务监控周期。<br> &nbsp;&nbsp;可以保持默认值。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">10s</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">CleanUpDayBefore</td>
    <td class="tg-5fx9">数据表中数据的失效时间。<br> &nbsp;&nbsp;过期后数据会被删除，单位为天。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">14</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">CleanUpPeriodHour</td>
    <td class="tg-61g0">启动异步清理任务的时间间隔，单位为小时。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">1</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">HealthCheckInterval</td>
    <td class="tg-5fx9">获取分布式锁和更新AASS实例状态的时间间隔，单位为秒。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">30</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HostProtect</td>
    <td class="tg-61g0">是否启用主机保护。<br> &nbsp;&nbsp;支持：【true,false】</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">TRUE</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">HostProtectValue</td>
    <td class="tg-5fx9">启用主机保护时，主机保护类型。<br> &nbsp;&nbsp;您可以参考：<br>  &nbsp;中国：https://support.huaweicloud.com/api-ecs/zh-cn_topic_0167957246.html#ZH-CN_TOPIC_0167957246__table2373623012315<br> &nbsp;&nbsp;国际：2024-01-22之前暂时无效。<br> &nbsp;&nbsp;支持：["ces", "hss", "hss,hss-ent"]</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">hss</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BatchCreatePodNum</td>
    <td class="tg-61g0">由于创建Pod时API网关的限制。<br> &nbsp;&nbsp;为了避免过多的调用次数，<br> &nbsp;&nbsp;如果次数超过该参数的值，<br> &nbsp;&nbsp;Pod休眠10秒。<br> &nbsp;&nbsp;可以保持默认值。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">50</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">HttpsAddress</td>
    <td class="tg-5fx9">监听地址</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">0.0.0.0</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HttpsPort</td>
    <td class="tg-61g0">监听端口</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">9091</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">TakeOverTaskIntervalSeconds</td>
    <td class="tg-5fx9">接管异步任务的时间间隔，单位为秒。<br> &nbsp;&nbsp;节点故障时使用。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">60</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">HeartBeatTaskIntervalSeconds</td>
    <td class="tg-61g0">节点健康状态上报周期，单位为秒。<br> &nbsp;&nbsp;用于判断节点是否正常。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">60</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">DeadCheckTaskIntervalSeconds</td>
    <td class="tg-5fx9">监控僵尸节点的时间间隔，单位为秒。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">60</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">MaxDeadMinutes</td>
    <td class="tg-61g0">判断僵尸节点的时间间隔，单位为分钟。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">3</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">InstanceMaximumLimit</td>
    <td class="tg-5fx9">伸缩组的最大实例数。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">500</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BandwidthMaximumLimit</td>
    <td class="tg-61g0">创建机器时使用的最大网络带宽。</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">300</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-5fx9">SupportedVolumeTypes</td>
    <td class="tg-5fx9">支持的卷类型。</td>
    <td class="tg-5fx9">您可以保留默认值</td>
    <td class="tg-5fx9">SATA;SAS;GPSSD;SSD;ESSD</td>
    <td class="tg-5fx9">　</td>
  </tr>
  <tr>
    <td class="tg-61g0">BandwidthChargingMode</td>
    <td class="tg-61g0">带宽计费方式。<br> &nbsp;&nbsp;支持：【流量、带宽】</td>
    <td class="tg-61g0">您可以保留默认值</td>
    <td class="tg-61g0">traffic</td>
    <td class="tg-61g0">　</td>
  </tr>
  <tr>
    <td class="tg-ogp5" rowspan="4">influxdb</td>
    <td class="tg-yjjc">InfluxDBUser</td>
    <td class="tg-yjjc">aass和appgateway使用的influxdb用户</td>
    <td class="tg-yjjc">您可以保留默认值</td>
    <td class="tg-yjjc">root</td>
    <td class="tg-yjjc"></td>
  </tr>
  <tr>
    <td class="tg-fapl">influxDBAddress</td>
    <td class="tg-cly1">aass和appgateway使用的influxdb地址</td>
    <td class="tg-cly1">确保aass和appgateway服务能够连接到数据库。</td>
    <td class="tg-cly1">127.0.0.1:8086</td>
    <td class="tg-cly1"></td>
  </tr>
  <tr>
    <td class="tg-ym56">InfluxDBPassword</td>
    <td class="tg-yjjc">aass和appgateway使用的influxdb密码</td>
    <td class="tg-yjjc">GCM加密后的数据库密码</td>
    <td class="tg-yjjc"></td>
    <td class="tg-yjjc">可以使用如下命令加密：/sac-gfm cipher --mode encode --method gcm --text{数据库密码}</td>
  </tr>
  <tr>
    <td class="tg-cly1">inflxuDBName</td>
    <td class="tg-cly1">AppGateway服务与AASS服务相同</td>
    <td class="tg-cly1">您可以保留默认值</td>
    <td class="tg-cly1">autoscaling</td>
    <td class="tg-cly1"></td>
  </tr>
</tbody>
</table>
```