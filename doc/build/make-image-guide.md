# make-image-guide
## 功能说明：

`FleetMange服务器组件`提供了应用包的全生命周期管理能力，并且基于应用包提供了镜像打包能力，用于`MetaSpace`平台扩容实例，应用包镜像主要集成了`Auxproxy`服务组件以及希望执行的服务端应用，同时集成了云日志服务LTS与云监控CES的插件以对应用状态进行掌握。



## 使用说明：

1.将`Auxproxy`服务组件压缩包（zip或rar）、镜像打包环境构建脚本[image_env.sh](../../doc/build/image_env.sh)添加至管理账号的OBS桶中

2.将应用压缩包（zip或rar）添加至资源账号的OBS桶中

3.应用镜像中，应用默认在路径/local/app/{创建的应用包名}下进行解压，请根据该路径调整应用启动脚本，或修改环境构建脚本中的应用下载目标路径

4.`FleetMange服务器组件`启动脚本[fleetmanage_run.sh](../../doc/build/fleetmanager/fleetmanager_run.sh)中添加参数，具体为DEFAULT_SCRIPT_PATH（环境构建脚本的OBS路径）、DEFAULT_AUXPROXY_PATH（`Auxproxy`服务组件压缩包OBS路径），其余镜像创建相关参数也可以在启动脚本中修改

