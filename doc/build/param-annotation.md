# 环境部署字段说明
敏感信息加密方法详见：[tools/gcm_dec_enc/gcm_dec_enc.go](/tools/gcm_dec_enc/gcm_dec_enc.go)
|             字段             |                            定义                            |               作用               |        参数值示例        |
| :--------------------------: | :--------------------------------------------------------: | :------------------------------: | :----------------------: |
|           version            |                        组件版本后缀                        |           便于版本迭代           |            --            |
|          mysql_host          |                         数据库地址                         |                --                |            --            |
|          mysql_port          |                         数据库端口                         |                --                |            --            |
| mysql_fleetmanager_database  |                    fleetmanager数据库名                    |     存储fleetmanager必要数据     |            --            |
|  mysql_appgateway_database   |                     appgateway数据库名                     |      存储appgateway必要数据      |            --            |
|     mysql_aass_database      |                        aass数据库名                        |         存储aass必要数据         |            --            |
|          mysql_user          |                        数据库用户名                        |                --                |            --            |
|          mysql_pwd           |                         数据库密码                         |                --                |       加密后的密码       |
|         influx_host          |                       influxdb的地址                       |                --                |            --            |
|         influx_port          |                       influxdb的端口                       |                --                |            --            |
|       influx_database        |            influx的数据库(aass与apgateway相同)             | 用于动态计算获取弹性伸缩所需信息 |            --            |
|       influx_username        |                      influxdb的用户名                      |                --                |            --            |
|          influx_pwd          |                       influxdb的密码                       |                --                |       加密后的密码       |
|              ak              |                    加密后的华为云租户ak                    |          访问华为云资源          |        加密后的ak        |
|              sk              |                    加密后的华为云租户sk                    |          访问华为云资源          |        加密后的sk        |
|    managerUser_domain_id     |               管理账号的华为云租户domain_id                |                --                |            --            |
|      enterprise_project      |                         企业项目id                         |       管控资源所在企业项目       |            --            |
|            region            |                          数据中心                          |      部署资源所在的数据中心      |        cn-north-4        |
|           endpoint           |                          终端节点                          |         云资源所在的节点         | 一般为myhuaweicloud.com  |
|          aass_host           |                  aass组件所在服务器的地址                  |                --                |            --            |
|          aass_port           |                       aass的业务端口                       |                --                |        一般为9091        |
|       appgateway_host        |                 appgateway所在服务器的地址                 |                --                |            --            |
|       appgateway_port        |                    appgateway的业务端口                    |                --                |       一般为60003        |
|      fleetmanager_port       |                   fleetmanager的业务端口                   |                --                |       一般为31002        |
|          fleet_cidr          |                新建fleet所申请资源的子网段                 |                --                |            --            |
|        specification         |                       具体的实例规格                       |                --                |      如s6.xlarge.2       |
| internal_inbound_permissions | 弹性申请资源的安全组配置，默认必须放通appgateway的业务端口 |                --                |            --            |
|          dns_config          |                  在某个region下的dns配置                   |         便于连接业务集群         | 若为cn-north-4，则为示例 |


