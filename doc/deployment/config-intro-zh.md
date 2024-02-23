# Introduction to Configuration Files
GFM基于`conf/init.yaml`配置文件生成各个服务组件所需的配置文件，这里将对配置文件中的参数以及如何配置做出说明。

**其中：表中红色的参数需要根据实际情况修改，其他可以保持默认**

## 参数介绍
<table style="border-collapse:collapse;border-color:#ccc;border-spacing:0" class="tg">
  <thead>
    <tr>
      <th style="background-color:#D9D9D9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">Section(模块)</th>
      <th style="background-color:#D9D9D9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">Parameter(参数)</th>
      <th style="background-color:#D9D9D9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">Description(描述)</th>
      <th style="background-color:#D9D9D9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">How to Configurate(如何配置)</th>
      <th style="background-color:#D9D9D9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">Example(示例)</th>
      <th style="background-color:#D9D9D9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">其他(Others)</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="7">cipher</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">GCMKey</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">密码等敏感数据的加解密秘钥，长度为24个字符</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="7">已经默认为您生成。如果有必要，可以重新生成它。</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Fi1xOkgXvevHlTR5SIHHXAb2</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="2">如果需要重新生成，可以执行：./sac-gfm cipher --create gcm --init-conf conf/init.yaml</td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">GCMNonce</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">密码等敏感数据加解密偏移，长度16个字符</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">dlbaGmj9O32CFbTs</td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RSAPublicKey</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="2">控制台与服务之间的密码传输加密秘钥的路径</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/public.key</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="2">如果需要重新生成，可以执行exec: ./sac-gfm cipher --create rsa --init-conf conf/init.yaml，密钥需要和Console中的一致。默认公钥已经随前端软件包一起编译。如果重新生成密钥，则需要重新编译前端软件包。关于Console的编译方法，可以参考：<a href="build.zhmd#console">Console编译指导</a></td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RSAPrivateKey</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/private.key</td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TlsKey</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="3">HTTPS协议的自签名证书</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/tls.key</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="3">如果需要重新生成，可以执行：./sac-gfm cipher --create tls --init-conf conf/init.yaml</td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TlsCsr</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/tls.key</td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TlsCrt</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/tls.key</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="65">fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlAddress</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fleetmanager使用的mysql地址</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">确保fleetmanager服务可以连接到数据库。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:3306</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlUser</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fleetmanager使用的mysql用户</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">一般情况下，用户为root。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">root</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlDBName</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fleetmanager使用的数据库名</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">默认情况下不会创建它，需要提前在数据库中创建。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlPassword</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql使用的密码</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">GCM加密后的数据库密码</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#00F;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;text-decoration:underline;vertical-align:top;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">可以使用如下命令加密：./sac-gfm cipher --mode encode --method gcm --text {数据库密码}</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlCharset</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">数据库编码格式</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">utf8</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisAddress</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fleetmanager使用的redis地址</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">确保fleetmanager可以连接到redis。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:6379</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisPassword</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fleetmanager使用的redis密码</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis密码gcm加密</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">可以使用如下命令加密：./sac-gfm cipher --mode encode --method gcm --text {数据库密码}</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisDB</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Redis默认分区</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogRotateSize</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志分割的大小，单位为MB。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1024</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogBackupCount</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志最大保留条数。如果日志超过最大条数，则按时间删除日志。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">100</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogMaxAge</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志保存的最大天数。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">7</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportRegions</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">如果部署了多个区域，可以输入多个值。否则，只需要一个值。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以参考：
        <br> 中国站：https://developer.huaweicloud.com/endpoint
        <br> 国际站：https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">cn-north-4</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">该配置项需要配置为YAML列表类型。</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultLoginPassword</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">仅在首次登录系统时使用。初始化完成后，密码会被修改为新的密码。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#00F;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;text-decoration:underline;vertical-align:middle;word-break:normal">Gfm@2024</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">WorkflowPath</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">工作流配置文件路径</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/workflow/</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportPublicImage</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">已校验的公有云服务器镜像列表。如果有其他镜像，可以将其更改为你想要的。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">用户可以自定义区域，其他参数保持默认值。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportPublicImage:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      image: "CentOS 7.2 64bit,CentOS 7.9 64bit,CentOS 8.1 64bit,Ubuntu 22.04 server 64bit,CentOS 7.6 64bit for Tenant 20210525"</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">该配置项需要配置为YAML列表类型。</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportDockerImage</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">已验证Pod的Docker镜像列表。
      <br> &nbsp;&nbsp; 如果你确定还有其他的镜像，你可以将它们更改为你想要的。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">用户可以自定义区域，其他参数保持默认值。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportDockerImage:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      dockerOS: "swr.cn-north-7.myhuaweicloud.com/game/centos-super:v1,centos:7.6.1810,centos:7.2.1511"</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">该配置项需要配置为YAML列表类型。</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">EipType</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">每个区域支持的弹性IP类型可能不同。
        <br> &nbsp;&nbsp;在云服务上查询到弹性IP类型后，需要手动配置弹性IP类型。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以参考：
        <br> 中国区表4：https://support.huaweicloud.com/api-eip/eip_api_0001.html
        <br> 国际站表3:https://support.huaweicloud.com/intl/en-us/api-eip/eip_api_0001.html
        <br> 如果有更多，应该用逗号隔开，例如：“eip-type1,eip-type2”</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">EipType:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      supportEipType: 5_bgp,5_sbgp</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">该配置项需要配置为YAML列表类型。</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DnsConfig</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">华为云私有DNS服务器地址。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以从表1中获取：
        <br> &nbsp;&nbsp;中国站：https://support.huaweicloud.com/dns_faq/dns_faq_002.html
        <br> &nbsp;&nbsp;国际站：https://support.huaweicloud.com/intl/en-us/dns_faq/dns_faq_002.html</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DnsConfig:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      dnsServer: 100.125.1.250,100.125.129.250</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">该配置项需要配置为YAML列表类型类型。</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LtsIps</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">当您需要使用日志服务管理日志时，需要使用此配置项。
        <br> &nbsp;&nbsp;该配置项用于在华为云主机上自动安装ICAgent，用于日志采集。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">需要查询配置地址。
        <br> &nbsp;&nbsp;您可以登录云日志服务控制台，选择需要部署的区域。
        <br> &nbsp;&nbsp;选择：主机管理->安装ICAgent。
        <br> &nbsp;&nbsp;你可以通过在给定的命令第2步中的accessip获得它。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LtsIps:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      ltsIp: 100.125.12.150</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">该配置项需要配置为YAML列表类型。</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultRegion</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">如果创建fleet时没有提供区域信息，则使用该参数值。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">cn-north-4</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultFleetProtectPolicy</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">车队的默认保护策略。
        <br> &nbsp;&nbsp;该配置项在缩容时会使用。
        <br> &nbsp;&nbsp;支持：[TIME_LIMIT_PROTECTION, NO_PROTECTION]
        <br> &nbsp;&nbsp;NO_PROTECTION可能会影响正在玩游戏的玩家，不建议
        <br> &nbsp;&nbsp;TIME_LIMIT_PROTECTION在保护期内，机器不会被强制回收。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TIME_LIMIT_PROTECTION</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultProtectTimeLimit</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fleet缩容保护时间，单位为分钟。
        <br> &nbsp;&nbsp;仅当本命令中的Protection Policy参数为TIME_LIMIT_PROTECTION时，该参数有效。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">30</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultSessionTimeoutSeconds</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">创建会话时，如果会话在该时间后还没有被激活，则会话将设置为错误。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultMaxSessionNumPerProcess</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">配置每个进程默认启动的会话数。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">50</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultProcessNumPerInstance</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">配置每个实例默认启动的进程数。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">50</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultBandwidth</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">默认带宽</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">5</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultBandwidthChargingMode</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">默认带宽计费方式</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">traffic</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultDiskSize</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">默认磁盘大小</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">40</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultVolumeType</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">默认卷类型</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SATA</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultEipShareType</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">带宽共享类型。
        <br> &nbsp;&nbsp;共享类型枚举值：
        <br> &nbsp;&nbsp;PER，表示独占使用。WHOLE：表示共享。
        <br> &nbsp;&nbsp;支持：[每，全部]</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">PER</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">FleetCidr</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">当创建车队时需要创建VPC时，需要配置该配置项。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">FleetCidr:
        <br> &nbsp;&nbsp;    from: 10.100.0.0
        <br> &nbsp;&nbsp;    to: 10.240.0.0
        <br> &nbsp;&nbsp;    vpcNetMask: 19
        <br> &nbsp;&nbsp;    subnetNetMask: 20</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InternalInboundPermissions</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">创建fleet时，默认会创建一个安全组。
        <br> &nbsp;&nbsp;下面列出了默认需要配置的安全组规则。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ipRange需要配置为所有AppGateway节点的IP地址。
        <br> &nbsp;&nbsp;例如，AppGateway有两个节点10.100.0.1和10.100.0.2，则需要将ipRange的值设置为示例值。
        <br> &nbsp;&nbsp;其他参数可以保持默认配置。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InternalInboundPermissions:
        <br> &nbsp;&nbsp;    - protocol: "TCP"
        <br> &nbsp;&nbsp;      ipRange: 10.100.0.1/32
        <br> &nbsp;&nbsp;      fromPort: 60001
        <br> &nbsp;&nbsp;      toPort: 60001
        <br> &nbsp;&nbsp;    - protocol: "TCP"
        <br> &nbsp;&nbsp;      ipRange: 10.100.0.1/32
        <br> &nbsp;&nbsp;      fromPort: 60001
        <br> &nbsp;&nbsp;      toPort: 60001</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildImageRef</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">制作镜像时使用的默认操作系统</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CentOS 7.2 64bit</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildFlavor</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">创建镜像时使用的默认规格</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">s6.large.2</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildBandwidth</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">创建镜像时的默认带宽</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">100</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildScriptPath</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">创建ECS镜像脚本</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">上传镜像打包脚本到OBS桶中。
        <br> &nbsp;&nbsp;例如，OBS桶为gfm，则配置为类似示例的值。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">gfm/image_env.sh</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">将bin/image_env.sh脚本上传到OBS桶中。</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildDockerScriptPath</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">制作POD镜像脚本</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">将打包docker镜像的脚本上传到OBS桶中。
        <br> &nbsp;&nbsp;例如，OBS桶为gfm，则配置为类似示例的值。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">gfm/docker_image_env.sh</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">将bin/docker_image_env.sh脚本上传到OBS桶中。</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AuxproxyPath</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">auxproxy服务应用包</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">上传应用包到OBS桶中。
        <br> &nbsp;&nbsp;例如，OBS桶为gfm，则配置为类似示例的值。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">gfm/auxproxy.zip</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">将bin/auxproxy.zip上传到OBS桶中。</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ImageDiskSize</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">镜像磁盘大小</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">40</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServerSessionBackupDays</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">历史数据失效时间，单位为天。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultScalingInCoolDownInterval</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">缩容周期，单位为分钟。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">10</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultGroupMaxSize</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="3">伸缩组的最大实例数、最小实例数和期望实例数</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultGroupMinSize</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultGroupDesiredSize</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.AASS</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AASS对外提供服务的地址。FleetManager通过该地址可以访问AASS。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">如果已为AASS配置了ELB服务，请配置ELB服务的IP地址。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://127.0.0.1:9091</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.AppGateway</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">appgateway对外提供服务的地址。
        <br> &nbsp;&nbsp;FleetManager可以通过该地址访问Appgateway。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">如果appgateway配置了ELB服务，请配置ELB服务的IP地址。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://127.0.0.1:60003</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.IamService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="11">华为云服务的服务端点</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="11">您可以从：
        <br> &nbsp;&nbsp;中国站：https://developer.huaweicloud.com/endpoint
        <br> &nbsp;&nbsp;国际站：https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://iam.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="11">示例值省略了区域。在实际应用中，区域应设置为：
        <br> &nbsp;&nbsp;区域：cn-north-4
        <br> &nbsp;&nbsp;终端节点：*****</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.VpcService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://vpc.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.ObsService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://obs.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.ImsService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://ims.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.EcsService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://ecs.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.SwrService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://swr.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.CciService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://cci.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.EvsService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://evs.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.CesService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://ces.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.DnsService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://dns.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.EpsService</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://eps.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SessionJwtKey</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">用于生成或验证登录会话的密钥。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">可以选择24位或32位的随机字符串。默认生成。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fFMH9nZgM3rpzp2os6RSbIicAeTazBV0I7ZZwFWnIQJFMHVP</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">如果需要重新生成，可以执行：./sac-gfm cipher --create jwt --init-conf conf/init.yaml</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SessionJwtTokenLifeTimeSecond</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">登录状态令牌的超时时间。
        <br> &nbsp;&nbsp;如果超过超时时间，需要重新登录</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">7200</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TakeOverTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">接管异步任务的时间间隔，单位为秒。
        <br> &nbsp;&nbsp;节点故障时使用。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsAddress</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">监听地址</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">0.0.0.0</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsPort</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">监听端口</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">31002</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HeartBeatTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">节点健康状态上报周期，单位为秒。
        <br> &nbsp;&nbsp;用于判断节点是否正常。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DeadCheckTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">监控僵尸节点的时间间隔，单位为秒。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MaxDeadMinutes</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">判断僵尸节点的时间间隔，单位为分钟。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">3</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="20">appgateway</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AASSAddress</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AppGateway访问AASS的地址</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">如果已为AASS配置了ELB服务，请配置ELB服务的IP地址。</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:9091</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlAddress</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">appgateway使用的mysql地址</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">确保appgateway服务可以连接到数据库。</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:3306</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlUser</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">appgateway使用的mysql用户</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">一般情况下，用户为root。</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">root</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlDBName</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">appgateway使用的数据库名</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">默认情况下不会创建它，需要提前在数据库中创建。</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">appgateway</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlPassword</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">appgateway使用的mysql密码</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">GCM加密后的数据库密码</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#00F;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;text-decoration:underline;vertical-align:middle;word-break:normal">
      </td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">可以使用如下命令加密：./sac-gfm cipher --mode encode --method gcm --text {数据库密码}</td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisAddress</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">appgateway使用的redis地址</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">确保appgateway服务可以连接到redis。</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:6379</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisPassword</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">appgateway使用的redis密码</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis密码gcm加密</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">可以使用如下命令加密：./sac-gfm cipher --mode encode --method gcm --text {数据库密码}</td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisDB</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Redis默认分区</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">3</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CleanStrategy</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">是否清理数据库中的数据
        <br> &nbsp;&nbsp;支持：[on，off]</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">on</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CleanupDays</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">从数据库中删除进程中的数据和会话备份表，默认14天</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">14</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BackupDays</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">从数据库备份进程中的数据和会话运行时表，默认为3天</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">3</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsAddress</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">绑定地址</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">0.0.0.0</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsPort</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">绑定端口</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60003</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DeployModel</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Appgateway启动方式
        <br> &nbsp;&nbsp;支持：【单实例，多实例】</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">multi-instances</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AuxproxyIpType</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">配置Appgateway通过公网IP或私网IP连接AuxProxy。
        <br> &nbsp;&nbsp;支持：[publicIP，privateIP]</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">publicIP</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogLevel</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">是否打印调试日志。
        <br> &nbsp;&nbsp;支持：[info，debug] </td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">info</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogRotateSize</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志分割的大小，单位为MB。</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1024</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogBackupCount</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志最大保留条数。
        <br> &nbsp;&nbsp;如果日志超过最大条数，则按时间删除日志。</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">100</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogMaxAge</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志最大保存天数</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">7</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogCompress</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志是否自动压缩。
        <br> &nbsp;&nbsp;注意：日志压缩后可能会损坏部分日志。
        <br> &nbsp;&nbsp;因此，建议不要压缩日志。</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">False</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="38">aass</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AppgatewayAddress</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AASS访问AppGateway的AppGateway地址</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">如果AppGateway配置了ELB服务，请配置ELB服务的IP地址。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:60003</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlAddress</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass使用的mysql地址</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">确保aass服务可以连接到数据库。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:3306</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlUser</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass使用的mysql用户</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">一般情况下，用户为root。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">root</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlDBName</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass使用的数据库名</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">默认情况下不会创建它。提前在数据库中创建。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlPassword</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass使用的mysql密码</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">GCM加密后的数据库密码</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#00F;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;text-decoration:underline;vertical-align:top;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">可以使用如下命令加密：./sac-gfm cipher --mode encode --method gcm --text {数据库密码}</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlCharset</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">数据库编码格式</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">utf8</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisAddress</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass使用的redis地址</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">确保aass服务可以连接到redis。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:6379</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisPassword</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass使用的redis密码</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis密码gcm加密</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">可以使用如下命令加密：./sac-gfm cipher --mode encode --method gcm --text {数据库密码}</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisDB</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Redis默认分区</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">2</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogRotateSize</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志分割的大小，单位为MB。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1024</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogBackupCount</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志最大保留条数。
        <br> &nbsp;&nbsp;如果日志超过最大条数，则按时间删除日志。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">100</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogMaxAge</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">日志最大保存天数</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">7</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientRegion</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="8">AASS服务所在的Region和华为云服务的Service Endpoint</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="8">您可以从：
        <br> &nbsp;&nbsp;中国站：https://developer.huaweicloud.com/endpoint
        <br> &nbsp;&nbsp;国际站：https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">cn-north-4</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientIamEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://iam.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientAsEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://as.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientEcsEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://ecs.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientLtsEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://lts.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientSmnEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://smn.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientPodEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://cci.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientDNSEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://dns.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudPlatformAddr</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">弹性云服务器元数据包括弹性云服务器在云平台上的基本信息，如弹性云服务器ID、主机名、网络信息等。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#00F;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;text-decoration:underline;vertical-align:top;word-break:normal">http://169.254.169.254
      </td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MonitorDuration</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">异步弹性伸缩任务监控周期。
        <br> &nbsp;&nbsp;可以保持默认值。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">10s</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CleanUpDayBefore</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">数据表中数据的失效时间。
        <br> &nbsp;&nbsp;过期后数据会被删除，单位为天。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">14</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CleanUpPeriodHour</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">启动异步清理任务的时间间隔，单位为小时。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HealthCheckInterval</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">获取分布式锁和更新AASS实例状态的时间间隔，单位为秒。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">30</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HostProtect</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">是否启用主机保护。
        <br> &nbsp;&nbsp;支持：【true,false】</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TRUE</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HostProtectValue</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">启用主机保护时，主机保护类型。
        <br> &nbsp;&nbsp;您可以参考：
        <br>  &nbsp;中国站：https://support.huaweicloud.com/api-ecs/zh-cn_topic_0167957246.html#ZH-CN_TOPIC_0167957246__table2373623012315
        <br> &nbsp;&nbsp;国际站：2024-01-22之前暂时无效。
        <br> &nbsp;&nbsp;支持：["ces", "hss", "hss,hss-ent"]</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">hss</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BatchCreatePodNum</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">由于创建Pod时API网关的限制。
        <br> &nbsp;&nbsp;为了避免过多的调用次数，
        <br> &nbsp;&nbsp;如果次数超过该参数的值，
        <br> &nbsp;&nbsp;Pod休眠10秒。
        <br> &nbsp;&nbsp;可以保持默认值。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">50</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsAddress</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">监听地址</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">0.0.0.0</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsPort</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">监听端口</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">9091</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TakeOverTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">接管异步任务的时间间隔，单位为秒。
        <br> &nbsp;&nbsp;节点故障时使用。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HeartBeatTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">节点健康状态上报周期，单位为秒。
        <br> &nbsp;&nbsp;用于判断节点是否正常。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DeadCheckTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">监控僵尸节点的时间间隔，单位为秒。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MaxDeadMinutes</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">判断僵尸节点的时间间隔，单位为分钟。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">3</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InstanceMaximumLimit</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">伸缩组的最大实例数。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">500</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BandwidthMaximumLimit</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">创建机器时使用的最大网络带宽。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">300</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportedVolumeTypes</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">支持的卷类型。</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SATA;SAS;GPSSD;SSD;ESSD</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BandwidthChargingMode</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">带宽计费方式。
        <br> &nbsp;&nbsp;支持：[traffic、bandwidth]</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">traffic</td>
      <td style="background-color:#F2F2F2;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="4">influxdb</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InfluxDBUser</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass和appgateway使用的influxdb用户</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">root</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">influxDBAddress</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass和appgateway使用的influxdb地址</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">确保aass和appgateway服务能够连接到数据库。</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:8086</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InfluxDBPassword</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass和appgateway使用的influxdb密码</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">GCM加密后的数据库密码</td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
      <td style="background-color:#f9f9f9;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">可以使用如下命令加密：./sac-gfm cipher --mode encode --method gcm --text {数据库密码}</td>
    </tr>
    <tr>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">inflxuDBName</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AppGateway服务与AASS服务相同</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">您可以保留默认值</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">autoscaling</td>
      <td style="background-color:#fff;border-color:inherit;border-style:solid;border-width:1px;color:#333;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
  </tbody>
</table>