# Introduction to Configuration Files

GFM generates configuration files required by each service component based on the `conf/init.yaml` configuration file. This section describes the parameters in the configuration files and how to configure the parameters.

**Modify the parameters in red based on the site requirements. Others you can keep default value**

## 参数介绍

<table style="border-collapse:collapse;border-spacing:0" class="tg">
  <thead>
    <tr>
      <th style="background-color:#D9D9D9;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">Section(模块)</th>
      <th style="background-color:#D9D9D9;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">Parameter(参数)</th>
      <th style="background-color:#D9D9D9;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">Description(描述)</th>
      <th style="background-color:#D9D9D9;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">How to Configurate(如何配置)</th>
      <th style="background-color:#D9D9D9;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">Example(示例)</th>
      <th style="background-color:#D9D9D9;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal">其他(Others)</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="7">cipher</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">GCMKey</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">24 characters used for encrypting and decrypting sensitive data such as passwords</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="7">It has been generated for you by default. You can regenerate it if necessary. </td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Fi1xOkgXvevHlTR5SIHHXAb2</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="2">If you need to regenerate, you can exec: ./sac-gfm cipher --create gcm --init-conf conf/init.yaml</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">GCMNonce</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">16 characters used for encrypting and decrypting sensitive data such as passwords</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">dlbaGmj9O32CFbTs</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RSAPublicKey</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="2">Encryption for password transmission between the console and the service with RSA</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/public.key</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="2">If you need to regenerate, you can exec: ./sac-gfm cipher --create rsa --init-conf conf/init.yaml. The key must be the same as that on the console. The default public key has been compiled with the frontend software package. If you regenerate the key, you must recompile the frontend software package. For details about how to compile the key, you can refer to: <a href="build-en.md#console">Console Build Instructions.</a></td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RSAPrivateKey</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/private.key</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TlsKey</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="3">From signature certificate for HTTPS protocol</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/tls.key</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="3">If you need to regenerate, you can exec: ./sac-gfm cipher --create tls --init-conf conf/init.yaml</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TlsCsr</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/tls.key</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TlsCrt</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/tls.key</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="65">fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlAddress</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql address used by fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Ensure that the fleetmanager service can connect to the database.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:3306</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlUser</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql user used by fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Generally, the user is root.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">root</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlDBName</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">database name used by fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">It is not created by default. Create it in the database in advance.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlPassword</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql password used by fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Database password encrypted using GCM</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#00F;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;text-decoration:underline;vertical-align:top;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can run the following command to encrypt the password: ./sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlCharset</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Database coding format</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">utf8</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisAddress</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis address used by fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Ensure that the fleetmanager service can connect to the redis.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:6379</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisPassword</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis password used by fleetmanager</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis password encrypted using GCM</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can run the following command to encrypt the password: ./sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisDB</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default Redis partition</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogRotateSize</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Size of a log split, in MB.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1024</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogBackupCount</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Maximum number of logs that can be retained. If the number of logs exceeds the maximum, the logs are deleted by time.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">100</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogMaxAge</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Maximum number of days for storing logs.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">7</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportRegions</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">If multiple regions are deployed, you can enter multiple values. Otherwise, only one value is required.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can refer:
        <br> &nbsp;&nbsp;China: https://developer.huaweicloud.com/endpoint
        <br> &nbsp;&nbsp;Intl: https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">cn-north-4</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">This configuration item needs to be configured as a YAML list.</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultLoginPassword</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">It is used only when you log in to the system for the first time. After the initialization, the password will be changed to a new one.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Gfm@2024</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">WorkflowPath</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Path of the workflow configuration file</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">./conf/workflow/</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportPublicImage</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">List of public ECS images that have been verified.
        <br> &nbsp;&nbsp;If other images are available, you can change them to the ones you want.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can customize the region and retain the default values for other parameters.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportPublicImage:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      image: "CentOS 7.2 64bit,CentOS 7.9 64bit,CentOS 8.1 64bit,Ubuntu 22.04 server 64bit,CentOS 7.6 64bit for Tenant 20210525"</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">This configuration item needs to be configured as a YAML list.</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportDockerImage</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">List of Docker images of verified pods.
        <br> &nbsp;&nbsp;If you are sure that other images are available, you can change them to the ones you want.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can customize the region and retain the default values for other parameters.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportDockerImage:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      dockerOS: "swr.cn-north-7.myhuaweicloud.com/game/centos-super:v1,centos:7.6.1810,centos:7.2.1511"</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">This configuration item needs to be configured as a YAML list.</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">EipType</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">The EIP type supported by each region may be different.
        <br> &nbsp;&nbsp;You need to manually configure the EIP type after querying the EIP type on the cloud service.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can refer to:
        <br> &nbsp;&nbsp;Table 4 in China:  https://support.huaweicloud.com/api-eip/eip_api_0001.html
        <br> &nbsp;&nbsp;Table 3 in Intl: https://support.huaweicloud.com/intl/en-us/api-eip/eip_api_0001.html
        <br> &nbsp;&nbsp;If you have more, should Use commas to separate, for example: "eip-type1,eip-type2"</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">EipType:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      supportEipType: 5_bgp,5_sbgp</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">This configuration item needs to be configured as a YAML list.</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DnsConfig</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Huawei Cloud Private DNS Server Addresses.
        <br> &nbsp;&nbsp;</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can get it from Table 1 in:
        <br> &nbsp;&nbsp;China: https://support.huaweicloud.com/dns_faq/dns_faq_002.html
        <br> &nbsp;&nbsp;Intl: https://support.huaweicloud.com/intl/en-us/dns_faq/dns_faq_002.html</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DnsConfig:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      dnsServer: 100.125.1.250,100.125.129.250</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">This configuration item needs to be configured as a YAML list.</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LtsIps</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">This configuration item is used if you want to use LTS to manage logs.
        <br> &nbsp;&nbsp;This configuration item is used to automatically install ICAgent on HUAWEI CLOUD hosts for log collection. </td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You need to query the configuration address.
        <br> &nbsp;&nbsp;You can log in to the LTS service console, Select the region to be deployed,
        <br> &nbsp;&nbsp;Select: Host Management -&gt; Install ICAgent.
        <br> &nbsp;&nbsp;You can get it from the step 2 by accessip in command given.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LtsIps:
        <br> &nbsp;&nbsp;    - region: cn-north-4
        <br> &nbsp;&nbsp;      ltsIp: 100.125.12.150</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">This configuration item needs to be configured as a YAML list.</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultRegion</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">If the region information is not provided when creating fleet, the value will be used.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">cn-north-4</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultFleetProtectPolicy</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default protection policy of the fleet.
        <br> &nbsp;&nbsp;This configuration item will be used during scale-in.
        <br> &nbsp;&nbsp;Support: [TIME_LIMIT_PROTECTION, NO_PROTECTION]
        <br> &nbsp;&nbsp;NO_PROTECTION maybe affect the player who is playing the game, not suggest
        <br> &nbsp;&nbsp;TIME_LIMIT_PROTECTION: You will give a protection period.
        <br> &nbsp;&nbsp;After the scale-in is triggered, the machine will not be forcibly reclaimed during the protection period.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TIME_LIMIT_PROTECTION</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultProtectTimeLimit</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Protection time during Fleet scale-in, in minute.
        <br> &nbsp;&nbsp;This parameter is valid only when Protection Policy is set to TIME_LIMIT_PROTECTION.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">30</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultSessionTimeoutSeconds</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">When a session is created, 
        <br> &nbsp;&nbsp;if the session is not activated after the time specified by this parameter, in seconds.
        <br> &nbsp;&nbsp;the session will set to error.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultMaxSessionNumPerProcess</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Configure the default number of session to be started per process.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">50</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultProcessNumPerInstance</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Configure the default number of Process to be started per instance.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">50</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultBandwidth</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default bandwidth</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">5</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultBandwidthChargingMode</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default bandwidth charge mode
        <br> &nbsp;&nbsp;Support: [bandwidth, traffic]</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">traffic</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultDiskSize</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default disk size</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">40</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultVolumeType</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default Volume Type</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SATA</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultEipShareType</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Specifies the bandwidth sharing type.
        <br> &nbsp;&nbsp;Share type enumerated value:
        <br> &nbsp;&nbsp;PER, indicating exclusive use. WHOLE: indicates sharing.
        <br> &nbsp;&nbsp;Support: [PER, WHOLE]
        <br> &nbsp;&nbsp;Default: PER</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">PER</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">FleetCidr</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">This configuration item is required when a VPC needs to be created during fleet creation.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">FleetCidr:
        <br> &nbsp;&nbsp;    from: 10.100.0.0
        <br> &nbsp;&nbsp;    to: 10.240.0.0
        <br> &nbsp;&nbsp;    vpcNetMask: 19
        <br> &nbsp;&nbsp;    subnetNetMask: 20</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InternalInboundPermissions</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">A security group is created by default when a fleet is created.
        <br> &nbsp;&nbsp;The following lists the security group rules that need to be configured by default.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">The ipRange configuration item needs to be set to the IP addresses of all AppGateway nodes.
        <br> &nbsp;&nbsp;For example, if there are two AppGateway nodes (10.100.0.1 and 10.100.0.2), you need to set ipRange to the values as the example value.
        <br> &nbsp;&nbsp;You can keep the default settings for other parameters.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InternalInboundPermissions:
        <br> &nbsp;&nbsp;    - protocol: "TCP"
        <br> &nbsp;&nbsp;      ipRange: 10.100.0.1/32
        <br> &nbsp;&nbsp;      fromPort: 60001
        <br> &nbsp;&nbsp;      toPort: 60001
        <br> &nbsp;&nbsp;    - protocol: "TCP"
        <br> &nbsp;&nbsp;      ipRange: 10.100.0.1/32
        <br> &nbsp;&nbsp;      fromPort: 60001
        <br> &nbsp;&nbsp;      toPort: 60001</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildImageRef</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default OS used when creating an image</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CentOS 7.2 64bit</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildFlavor</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default flavor used during image creation</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">s6.large.2</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildBandwidth</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default bandwidth used during image creation</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">100</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildScriptPath</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Script for creating an ECS image</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Upload the script for packaging the image to the OBS bucket.
        <br> &nbsp;&nbsp;For example, if the OBS bucket is gfm, set this parameter to a value similar to the example value.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">gfm/image_env.sh</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Upload the bin/image_env.sh script to the gfm bucket of OBS.</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BuildDockerScriptPath</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Script used for creating POD images</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Upload the script for packaging the docker image to the OBS bucket.
        <br> &nbsp;&nbsp;For example, if the OBS bucket is gfm, set this parameter to a value similar to the example value.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">gfm/docker_image_env.sh</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Upload the bin/docker_image_env.sh script to the gfm bucket of OBS.</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AuxproxyPath</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Application package of the auxproxy service</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Upload the application package to the OBS bucket.
        <br> &nbsp;&nbsp;For example, if the OBS bucket is gfm, set this parameter to a value similar to the example value.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">gfm/auxproxy.zip</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Upload the bin/auxproxy.zip to the gfm bucket of OBS.</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ImageDiskSize</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Disk size of image michine</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">40</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServerSessionBackupDays</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Expiration time of historical data, in days.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultScalingInCoolDownInterval</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Scale-in interval, in minutes</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">10</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultGroupMaxSize</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="3">Maximum, minimum, and desire number of instances in an AS group</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultGroupMinSize</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DefaultGroupDesiredSize</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.AASS</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Address used by the AASS to provide services for external systems. FleetManager can access the AASS through this address.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">If you have configured the ELB service for AASS, configurate the IP address of the ELB service.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://127.0.0.1:9091</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.AppGateway</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Address used by the Appgateway to provide services for external systems.
        <br> &nbsp;&nbsp;FleetManager can access the Appgateway through this address.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">If you have configured the ELB service for appgateway, configurate the IP address of the ELB service.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://127.0.0.1:60003</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.IamService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="11">Service Endpoint about HuaweiCloud Service</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:top;word-break:normal" rowspan="11"> you can get those from:
        <br> &nbsp;&nbsp;China: https://developer.huaweicloud.com/endpoint
        <br> &nbsp;&nbsp;Intl: https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://iam.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="11">The example value omits the region. In practice, the region should be set to:
        <br> &nbsp;&nbsp;region: cn-north-4
        <br> &nbsp;&nbsp;endPoint:*****</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.VpcService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://vpc.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.ObsService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://obs.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.ImsService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://ims.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.EcsService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://ecs.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.SwrService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://swr.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.CciService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://cci.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.EvsService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://evs.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.CesService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://ces.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.DnsService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://dns.cn-north-4.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ServiceEndpoint.EpsService</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://eps.myhuaweicloud.com</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SessionJwtKey</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Similar to the key used to generate or verify the login session,</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">you can choose a random character string of 24 or 32 bits. It is generated by default.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">fFMH9nZgM3rpzp2os6RSbIicAeTazBV0I7ZZwFWnIQJFMHVP</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">If you need to regenerate, you can exec: ./sac-gfm cipher --create jwt --init-conf conf/init.yaml</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SessionJwtTokenLifeTimeSecond</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Timeout interval of the login status token.
        <br> &nbsp;&nbsp;If the timeout interval is exceeded, you need to log in again</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">7200</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TakeOverTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for taking over asynchronous tasks, in seconds.
        <br> &nbsp;&nbsp;This parameter is used when a node is faulty.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsAddress</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">bind address</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">0.0.0.0</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsPort</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">bind port</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">31002</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HeartBeatTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for reporting node health status, in seconds,
        <br> &nbsp;&nbsp;which is used to determine whether a node is normal.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DeadCheckTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for monitoring zombie nodes, in seconds</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MaxDeadMinutes</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for determining a zombie node, in mintues.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">3</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="20">appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AASSAddress</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AASS address for the AppGateway to access the AASS</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">If you have configured the ELB service for AASS, configurate the IP address of the ELB service.</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:9091</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlAddress</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql address used by appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Ensure that the appgateway service can connect to the database.</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:3306</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlUser</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql user used by appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Generally, the user is root.</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">root</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlDBName</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">database name used by appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">It is not created by default. Create it in the database in advance.</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlPassword</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql password used by appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Database password encrypted using GCM</td>
      <td style="border-color:black;border-style:solid;border-width:1px;color:#00F;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;text-decoration:underline;vertical-align:middle;word-break:normal">
      </td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can run the following command to encrypt the password: ./sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisAddress</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis address used by appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Ensure that the appgateway service can connect to the redis.</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:6379</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisPassword</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis password used by appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis password encrypted using GCM</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can run the following command to encrypt the password: ./sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisDB</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default Redis partition</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">3</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CleanStrategy</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">clean data in Database or not
        <br> &nbsp;&nbsp;Support: [off, on]</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">on</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CleanupDays</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Delete data in process and session backup table from database, Default 14 days</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">14</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BackupDays</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Backup data in process and session runtime table from database, Default 3 days</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">3</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsAddress</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">bind address</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">0.0.0.0</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsPort</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">bind port</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60003</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DeployModel</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Appgateway startup mode
        <br> &nbsp;&nbsp;Support: [singleton, multi-instances]</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">multi-instances</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AuxproxyIpType</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Configure the Appgateway to connect to the AuxProxy using a public or private IP address.
        <br> &nbsp;&nbsp;Support: [publicIP, privateIP]</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">publicIP</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogLevel</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Indicates whether to print debug logs.
        <br> &nbsp;&nbsp;Support: [info, debug]</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">info</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogRotateSize</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Size of a log split, in MB.</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1024</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogBackupCount</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Maximum number of logs that can be retained.
        <br> &nbsp;&nbsp;If the number of logs exceeds the maximum, the logs are deleted by time.</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">100</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogMaxAge</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Maximum number of days for storing logs</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">7</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogCompress</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Indicates whether logs are automatically compressed.
        <br> &nbsp;&nbsp;NOTE: Some logs may be damaged after log compression.
        <br> &nbsp;&nbsp;Therefore, you are advised not to compress logs.</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">FALSE</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="38">aass</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AppgatewayAddress</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">AppGateway address for the AASS to access the AppGateway</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">If you have configured the ELB service for AppGateway, configurate the IP address of the ELB service.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:60003</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlAddress</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql address used by aass</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Ensure that the aass service can connect to the database.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:3306</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlUser</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql user used by aass</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Generally, the user is root.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">root</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlDBName</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">database name used by aass</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">It is not created by default. Create it in the database in advance.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">aass</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlPassword</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">mysql password used by aass</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Database password encrypted using GCM</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#00F;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;text-decoration:underline;vertical-align:top;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can run the following command to encrypt the password: ./sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MysqlCharset</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Database coding format</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">utf8</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisAddress</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis address used by aass</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Ensure that the aass service can connect to the redis.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:6379</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisPassword</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis password used by aass</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">redis password encrypted using GCM</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can run the following command to encrypt the password: ./sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">RedisDB</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Default Redis partition</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">2</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogRotateSize</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Size of a log split, in MB.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1024</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogBackupCount</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Maximum number of logs that can be retained.
        <br> &nbsp;&nbsp;If the number of logs exceeds the maximum, the logs are deleted by time.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">100</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">LogMaxAge</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Maximum number of days for storing logs</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">7</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientRegion</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="8">Region where the AASS service is deployed and Service Endpoint about HuaweiCloud Service</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal" rowspan="8">you can get those from:
        <br> &nbsp;&nbsp;China: https://developer.huaweicloud.com/endpoint
        <br> &nbsp;&nbsp;Intl: https://developer.huaweicloud.com/intl/en-us/endpoint?all</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">cn-north-4</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientIamEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://iam.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientAsEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://as.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientEcsEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://ecs.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientLtsEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://lts.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientSmnEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://smn.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientPodEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://cci.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudClientDNSEndpoint</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">https://dns.cn-north-4.myhuaweicloud.com</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CloudPlatformAddr</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">ECS metadata includes basic information of an ECS on the cloud platform, such as the ECS ID, hostname, and network information. </td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">http://169.254.169.254</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MonitorDuration</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for monitoring asynchronous autoscaling tasks.
        <br> &nbsp;&nbsp;You can retain the default value.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">10s</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CleanUpDayBefore</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Expiration time of data in the data table.
        <br> &nbsp;&nbsp;Data will be deleted after the expiration time, in days.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">14</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">CleanUpPeriodHour</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for starting an asynchronous cleanup task, in hours</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">1</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HealthCheckInterval</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval between obtaining the distributed lock and updating the status of the AASS instance, in seconds.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">30</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HostProtect</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Whether to enable host protection.
        <br> &nbsp;&nbsp;Support: [true, false]</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TRUE</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HostProtectValue</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Specifies the host protection type if host protection is enabled.
        <br> &nbsp;&nbsp;You can refer to:
        <br> &nbsp;&nbsp;China: https://support.huaweicloud.com/api-ecs/zh-cn_topic_0167957246.html#ZH-CN_TOPIC_0167957246__table2373623012315
        <br> &nbsp;&nbsp;Intl: Not valid for the time being by 2024-01-22.
        <br> &nbsp;&nbsp;Support: ["ces", "hss", "hss,hss-ent"]
        <br> &nbsp;&nbsp;
        <br> &nbsp;&nbsp;</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">hss</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BatchCreatePodNum</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Due to the restriction of the APIG during pod creation,
        <br> &nbsp;&nbsp;to avoid excessive invoking times,
        <br> &nbsp;&nbsp;if the number of times exceeds the value of this parameter,
        <br> &nbsp;&nbsp;the pod is hibernated for 10 seconds.
        <br> &nbsp;&nbsp;You can retain the default value.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">50</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsAddress</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">bind address</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">0.0.0.0</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HttpsPort</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">bind port</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">9091</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">TakeOverTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for taking over asynchronous tasks, in seconds.
        <br> &nbsp;&nbsp;This parameter is used when a node is faulty.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">HeartBeatTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for reporting node health status, in seconds,
        <br> &nbsp;&nbsp;which is used to determine whether a node is normal.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">DeadCheckTaskIntervalSeconds</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for monitoring zombie nodes, in seconds</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">60</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">MaxDeadMinutes</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Interval for determining a zombie node, in mintues.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">3</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InstanceMaximumLimit</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Specifies the maximum number of instances in an AS group.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">500</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BandwidthMaximumLimit</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Maximum network bandwidth used when creating a machine.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">300</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SupportedVolumeTypes</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Supported volume types.</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">SATA;SAS;GPSSD;SSD;ESSD</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">BandwidthChargingMode</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Bandwidth charging mode.
        <br> &nbsp;&nbsp;Support: [traffic, bandwidth]</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">traffic</td>
      <td style="background-color:#F2F2F2;border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">　</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;font-weight:bold;overflow:hidden;padding:10px 5px;text-align:center;vertical-align:middle;word-break:normal" rowspan="4">influxdb</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InfluxDBUser</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">influxdb user used by aass and appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">root</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">influxDBAddress</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">influxdb address used by aass and appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Ensure that the aass and appgateway service can connect to the database.</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">127.0.0.1:8086</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;color:#F00;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">InfluxDBPassword</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">influxdb passworld used by aass and appgateway</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">Database password encrypted using GCM</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can run the following command to encrypt the password: ./sac-gfm cipher --mode encode --method gcm --text {Database Password}</td>
    </tr>
    <tr>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">inflxuDBName</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">The AppGateway service is the same as the AASS service. </td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">You can keep the default</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">autoscaling</td>
      <td style="border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;overflow:hidden;padding:10px 5px;text-align:left;vertical-align:middle;word-break:normal">
      </td>
    </tr>
  </tbody>
</table>