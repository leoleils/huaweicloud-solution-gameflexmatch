package action

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"solutionbuild/gameflexmatch/common"
	"solutionbuild/gameflexmatch/model"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type InstallAction struct {
	c *cli.Context
}

func NewInstallAction(c *cli.Context) *InstallAction {
	return &InstallAction{
		c: c,
	}
}

func (a *InstallAction) install(initConf string, service string, param map[string]string) error {
	absPath, err := filepath.Abs(initConf)
	if err != nil {
		return err
	}
	targetPath := filepath.Dir(absPath)
	fmt.Printf("Start to install %s...\n", service)
	fmt.Printf("Config: %s, target path: %s\n", absPath, targetPath)
	if err := a.InitConfig(absPath, targetPath, service); err != nil {
		return err
	}
	fmt.Printf("Configure deamon process for %s\n", service)
	paramS := map[string]string{
		"service": service,
		"bin":     fmt.Sprintf("%s/bin/%s", common.RootPath, service),
		"conf":    fmt.Sprintf("%s/%s.yaml", targetPath, service),
		"log":     param["log"],
	}
	if err := a.touchStartScript(paramS); err != nil {
		return err
	}
	paramI := map[string]string{
		"service": service,
		"bin":     fmt.Sprintf("%s/bin/%s-start.sh", common.RootPath, service),
	}
	if err := a.touchSystemd(paramI); err != nil {
		return err
	}
	fmt.Printf("Enable %s systemd service\n", service)
	if err := common.ExecCmd(exec.Command("systemctl", "enable", service)); err != nil {
		return err
	}
	fmt.Printf("Succeed to install %s service\n", service)
	fmt.Printf("If you want to start the service on terminal, can choose to exec:\n")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%s\n", paramI["bin"])
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	return nil
}

func (a *InstallAction) InitConfig(confPath string, targetPath string, service string) error {
	confdata, err := os.ReadFile(confPath)
	if err != nil {
		return err
	}
	conf := make(map[string]interface{})
	if err := yaml.Unmarshal(confdata, &conf); err != nil {
		return err
	}
	targetPath = fmt.Sprintf("%s/%s.yaml", targetPath, service)
	cipher := make(map[string]string)
	for k, v := range conf["cipher"].(map[string]interface{}) {
		cipher[k] = v.(string)
	}
	if !path.IsAbs(cipher["RSAPublicKey"]) {
		cipher["RSAPublicKey"] = path.Join(common.RootPath, cipher["RSAPublicKey"])
	}
	if !path.IsAbs(cipher["RSAPrivateKey"]) {
		cipher["RSAPrivateKey"] = path.Join(common.RootPath, cipher["RSAPrivateKey"])
	}
	if !path.IsAbs(cipher["TlsKey"]) {
		cipher["TlsKey"] = path.Join(common.RootPath, cipher["TlsKey"])
	}
	if !path.IsAbs(cipher["TlsCsr"]) {
		cipher["TlsCsr"] = path.Join(common.RootPath, cipher["TlsCsr"])
	}
	if !path.IsAbs(cipher["TlsCrt"]) {
		cipher["TlsCrt"] = path.Join(common.RootPath, cipher["TlsCrt"])
	}
	mapCipher := make(map[string]interface{})
	for k, v := range cipher {
		mapCipher[k] = v
	}

	conf["cipher"] = mapCipher
	switch service {
	case "fleetmanager":
		if err := a.generateFleetmanagerConfig(conf, targetPath); err != nil {
			return err
		}
		return nil
	case "appgateway":
		if err := a.generateAppgatewayConfig(conf, targetPath); err != nil {
			return err
		}
		return nil
	case "aass":
		if err := a.generateAASSConfig(conf, targetPath); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("you should choose one mode, one of 'fleetmanager'/'appgateway'/'aass'")
	}
}

func (a *InstallAction) generateFleetmanagerConfig(conf map[string]interface{}, targetPath string) error {
	var fmm = &model.FleetManager{}
	// init mysql config
	fleetM := conf[common.SectionFleetManager].(map[string]interface{})
	// influx := conf[common.SectionInfluxDB]
	cipher := conf[common.SectionCipher].(map[string]interface{})
	// mysql
	fmm.MysqlConfig = map[string]interface{}{
		"mysqlUser":               a.interface2String(fleetM["MysqlUser"]),
		"mysqlDBName":             a.interface2String(fleetM["MysqlDBName"]),
		"mysqlPassword":           a.interface2String(fleetM["MysqlPassword"]),
		"mysqlAddress":            a.interface2String(fleetM["MysqlAddress"]),
		"mysqlCharset":            a.interface2String(fleetM["MysqlCharset"]),
		"serverSessionBackupDays": fleetM["ServerSessionBackupDays"],
	}
	// redis
	fmm.RedisConfig = map[string]interface{}{
		"redisAddress":  a.interface2String(fleetM["RedisAddress"]),
		"redisDb":       fleetM["RedisDB"],
		"redisPassword": a.interface2String(fleetM["RedisPassword"]),
	}
	// auth
	fmm.AuthConfig = map[string]interface{}{
		"gcmKey":        a.interface2String(cipher["GCMKey"]),
		"gcmNonce":      a.interface2String(cipher["GCMNonce"]),
		"rsaPublicKey":  a.interface2String(cipher["RSAPublicKey"]),
		"rsaPrivateKey": a.interface2String(cipher["RSAPrivateKey"]),
	}
	// log
	fmm.LogConfig = map[string]interface{}{
		"logRotateSize":  fleetM["LogRotateSize"],
		"logBackupCount": fleetM["LogBackupCount"],
		"logMaxAge":      fleetM["LogMaxAge"],
	}
	// env
	workflowPath := path.Join(common.RootPath, a.interface2String(fleetM["WorkflowPath"]))
	fmm.EnvConfig = map[string]interface{}{
		"supportRegions":       fleetM["SupportRegions"],
		"defaultLoginPassword": a.interface2String(fleetM["DefaultLoginPassword"]),
		"workflowPath":         workflowPath+"/",
		"supportPublicImage":   fleetM["SupportPublicImage"],
		"supportDockerBased":   fleetM["SupportDockerImage"],
		"eipType":              fleetM["EipType"],
		"dnsConfig":            fleetM["DnsConfig"],
		"ltsIps":               fleetM["LtsIps"],
	}
	fmm.FleetConfig = map[string]interface{}{
		"defaultRegion":           a.interface2String(fleetM["DefaultRegion"]),
		"protectPolicy":           a.interface2String(fleetM["DefaultFleetProtectPolicy"]),
		"protectTimeLimit":        fleetM["DefaultProtectTimeLimit"],
		"sessionTimeoutSeconds":   fleetM["DefaultSessionTimeoutSeconds"],
		"maxSessionNumPerProcess": fleetM["DefaultMaxSessionNumPerProcess"],
		"processNumPerFleet":      fleetM["DefaultProcessNumPerInstance"],
		"bandwidth":               fleetM["DefaultBandwidth"],
		"bandwidthChargingMode":   a.interface2String(fleetM["DefaultBandwidthChargingMode"]),
		"diskSize":                fleetM["DefaultDiskSize"],
		"volumeType":              fleetM["DefaultVolumeType"],
		"fleetCidr":               fleetM["FleetCidr"],
	}

	fmm.InternalInboundPermissions = fleetM["InternalInboundPermissions"].([]interface{})

	fmm.BuildConfig = map[string]interface{}{
		"imageRef":         a.interface2String(fleetM["BuildImageRef"]),
		"imageFlavor":      a.interface2String(fleetM["BuildFlavor"]),
		"buildBandwidth":   fleetM["BuildBandwidth"],
		"scriptPath":       a.interface2String(fleetM["BuildScriptPath"]),
		"dockerScriptPath": a.interface2String(fleetM["BuildDockerScriptPath"]),
		"auxproxyPath":     a.interface2String(fleetM["AuxproxyPath"]),
		"imageDiskSize":    fleetM["ImageDiskSize"],
	}

	fmm.GroupConfig = map[string]interface{}{
		"scalingInCoolDownInterval": fleetM["DefaultScalingInCoolDownInterval"],
		"groupMaxSize":              fleetM["DefaultGroupMaxSize"],
		"groupMinSize":              fleetM["DefaultGroupMinSize"],
		"groupDesiredSize":          fleetM["DefaultGroupDesiredSize"],
	}
	serviceEndpoint := fleetM["ServiceEndpoint"].(map[string]interface{})
	fmm.ServiceEndpoint = map[string]interface{}{
		"iamService": serviceEndpoint["IamService"],
		"aass":       serviceEndpoint["AASS"],
		"appGateway": serviceEndpoint["AppGateway"],
		"vpcService": serviceEndpoint["VpcService"],
		"obsService": serviceEndpoint["ObsService"],
		"imsService": serviceEndpoint["ImsService"],
		"ecsService": serviceEndpoint["EcsService"],
		"swrService": serviceEndpoint["SwrService"],
		"cciService": serviceEndpoint["CciService"],
		"evsService": serviceEndpoint["EvsService"],
		"cesService": serviceEndpoint["CesService"],
		"dnsService": serviceEndpoint["DnsService"],
		"epsService": serviceEndpoint["EpsService"],
	}

	fmm.SessionConfig = map[string]interface{}{
		"jwtKey":                 a.interface2String(fleetM["SessionJwtKey"]),
		"jwtTokenLifeTimeSecond": fleetM["SessionJwtTokenLifeTimeSecond"],
	}

	fmm.HttpsConfig = map[string]interface{}{
		"httpsAddress":  fleetM["HttpsAddress"],
		"httpsPort":     fleetM["HttpsPort"],
		"httpsCertFile": cipher["TlsCrt"],
		"httpsKeyFile":  cipher["TlsKey"],
	}

	fmm.WorkNodeConfig = map[string]interface{}{
		"takeOverTaskIntervalSeconds":  fleetM["TakeOverTaskIntervalSeconds"],
		"heartBeatTaskIntervalSeconds": fleetM["HeartBeatTaskIntervalSeconds"],
		"deadCheckTaskIntervalSeconds": fleetM["DeadCheckTaskIntervalSeconds"],
		"maxDeadMinutes":               fleetM["MaxDeadMinutes"],
	}

	if err := a.writeToYaml(fmm, targetPath); err != nil {
		return err
	}
	return nil
}

func (a *InstallAction) generateAppgatewayConfig(conf map[string]interface{}, targetPath string) error {
	var agm = &model.AppGateway{}
	// init mysql config
	appG := conf[common.SectionAppGateway].(map[string]interface{})
	influx := conf[common.SectionInfluxDB].(map[string]interface{})
	cipher := conf[common.SectionCipher].(map[string]interface{})
	agm.MysqlConfig = map[string]interface{}{
		"mysqlDBName":   a.interface2String(appG["MysqlDBName"]),
		"mysqlUser":     a.interface2String(appG["MysqlUser"]),
		"mysqlPassword": a.interface2String(appG["MysqlPassword"]),
		"mysqlAddress":  a.interface2String(appG["MysqlAddress"]),
	}

	agm.InfluxDBConfig = map[string]interface{}{
		"influxDBUser":     a.interface2String(influx["InfluxDBUser"]),
		"influxDBPassword": a.interface2String(influx["InfluxDBPassword"]),
		"influxDBAddress":  a.interface2String(influx["influxDBAddress"]),
		"influxDBName":     a.interface2String(influx["inflxuDBName"]),
	}

	agm.RedisConfig = map[string]interface{}{
		"redisAddress":  a.interface2String(appG["RedisAddress"]),
		"redisDb":       appG["RedisDB"],
		"redisPassword": a.interface2String(appG["RedisPassword"]),
	}

	agm.CleanConfig = map[string]interface{}{
		"cleanStrategy": a.interface2String(appG["CleanStrategy"]),
		"cleanupDays":   appG["CleanupDays"],
		"backupDays":    appG["BackupDays"],
	}

	agm.EnvConfig = map[string]interface{}{
		"aassAddress":    a.interface2String(appG["AASSAddress"]),
		"deployModel":    a.interface2String(appG["DeployModel"]),
		"auxproxyIPType": a.interface2String(appG["AuxproxyIpType"]),
	}
	agm.HttpsConfig = map[string]interface{}{
		"httpsAddress":  appG["HttpsAddress"],
		"httpsPort":     appG["HttpsPort"],
		"httpsCertFile": cipher["TlsCrt"],
		"httpsKeyFile":  cipher["TlsKey"],
	}

	agm.LogConfig = map[string]interface{}{
		"logLevel":       appG["LogLevel"],
		"logRotateSize":  appG["LogRotateSize"],
		"logBackupCount": appG["LogBackupCount"],
		"logMaxAge":      appG["LogMaxAge"],
		"logCompress":    appG["LogCompress"],
	}
	agm.AuthConfig = map[string]interface{}{
		"gcmKey":   a.interface2String(cipher["GCMKey"]),
		"gcmNonce": a.interface2String(cipher["GCMNonce"]),
	}
	if err := a.writeToYaml(agm, targetPath); err != nil {
		return err
	}
	return nil
}

func (a *InstallAction) generateAASSConfig(conf map[string]interface{}, targetPath string) error {
	var asm = &model.AASS{}
	// init mysql config
	aass := conf[common.SectionAASS].(map[string]interface{})
	influx := conf[common.SectionInfluxDB].(map[string]interface{})
	cipher := conf[common.SectionCipher].(map[string]interface{})
	asm.MysqlConfig = map[string]interface{}{
		"mysqlDBName":     a.interface2String(aass["MysqlDBName"]),
		"mysqlUser":       a.interface2String(aass["MysqlUser"]),
		"mysqldbPassword": a.interface2String(aass["MysqlDBPassword"]),
		"mysqlAddress":    a.interface2String(aass["MysqlAddress"]),
		"mysqlCharset":    a.interface2String(aass["MysqlCharset"]),
	}
	asm.RedisConfig = map[string]interface{}{
		"redisAddress":  a.interface2String(aass["RedisAddress"]),
		"redisDb":       aass["RedisDB"],
		"redisPassword": a.interface2String(aass["RedisPassword"]),
	}

	asm.AuthConfig = map[string]interface{}{
		"gcmKey":   a.interface2String(cipher["GCMKey"]),
		"gcmNonce": a.interface2String(cipher["GCMNonce"]),
	}

	asm.LogConfig = map[string]interface{}{
		"logRotateSize":  aass["LogRotateSize"],
		"logBackupCount": aass["LogBackupCount"],
		"logMaxAge":      aass["LogMaxAge"],
	}

	asm.ServiceEndpoint = map[string]interface{}{
		"appgatewayAddress": aass["AppgatewayAddress"],
	}

	asm.CloudEndpoint = map[string]interface{}{
		"cloudClientRegion":      a.interface2String(aass["CloudClientRegion"]),
		"cloudClientIamEndpoint": a.interface2String(aass["CloudClientIamEndpoint"]),
		"cloudClientAsEndpoint":  a.interface2String(aass["CloudClientAsEndpoint"]),
		"cloudClientEcsEndpoint": a.interface2String(aass["CloudClientEcsEndpoint"]),
		"cloudClientLtsEndpoint": a.interface2String(aass["CloudClientLtsEndpoint"]),
		"cloudClientSmnEndpoint": a.interface2String(aass["CloudClientSmnEndpoint"]),
		"cloudClientPodEndpoint": a.interface2String(aass["CloudClientPodEndpoint"]),
		"cloudClientDNSEndpoint": a.interface2String(aass["CloudClientDNSEndpoint"]),
		"cloudPlatformAddr":      a.interface2String(aass["CloudPlatformAddr"]),
	}

	asm.InfluxDBConfig = map[string]interface{}{
		"influxDBUser":     a.interface2String(influx["InfluxDBUser"]),
		"influxDBPassword": a.interface2String(influx["InfluxDBPassword"]),
		"influxDBAddress":  a.interface2String(influx["influxDBAddress"]),
		"influxDBName":     a.interface2String(influx["inflxuDBName"]),
	}

	asm.EnvConfig = map[string]interface{}{
		"connectToAuxproxyIPType": a.interface2String(aass["ConnectToAuxproxyIPType"]),
		"monitorDuration":         a.interface2String(aass["MonitorDuration"]),
		"cleanUpDayBefore":        aass["CleanUpDayBefore"],
		"cleanUpPeriodHour":       aass["CleanUpPeriodHour"],
		"healthCheckInterval":     aass["HealthCheckInterval"],
		"hostProtect":             aass["HostProtect"],
		"hostProtectValue":        a.interface2String(aass["HostProtectValue"]),
		"batchCreatePodNum":       aass["BatchCreatePodNum"],
	}

	asm.HttpsConfig = map[string]interface{}{
		"httpsAddress":  aass["HttpsAddress"],
		"httpsPort":     aass["HttpsPort"],
		"httpsCertFile": cipher["TlsCrt"],
		"httpsKeyFile":  cipher["TlsKey"],
	}
	asm.WorkNodeConfig = map[string]interface{}{
		"takeOverTaskIntervalSeconds":  aass["TakeOverTaskIntervalSeconds"],
		"heartBeatTaskIntervalSeconds": aass["HeartBeatTaskIntervalSeconds"],
		"deadCheckTaskIntervalSeconds": aass["DeadCheckTaskIntervalSeconds"],
		"maxDeadMinutes":               aass["MaxDeadMinutes"],
	}
	asm.ScalingGroupConfig = map[string]interface{}{
		"instanceMaximumLimit":  aass["InstanceMaximumLimit"],
		"bandwidthMaximumLimit": aass["BandwidthMaximumLimit"],
		"supportedVolumeTypes":  aass["SupportedVolumeTypes"],
		"bandwidthChargingMode": aass["BandwidthChargingMode"],
	}

	if err := a.writeToYaml(asm, targetPath); err != nil {
		return err
	}
	return nil

}

func (a *InstallAction) writeToYaml(data interface{}, targetPath string) error {
	dataW, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	if err = os.WriteFile(targetPath, dataW, 0644); err != nil {
		return err
	}
	fmt.Printf("succeed to generate config file into: %s\n", targetPath)
	return nil
}

func (a *InstallAction) interface2String(source interface{}) string {
	if source == nil {
		return ""
	}
	res, ok := source.(string)
	if !ok {
		return ""
	}
	return res
}

func (a *InstallAction) touchSystemd(param map[string]string) error {
	// TODO: 补充官网链接
	service := param["service"]
	bin_path := param["bin"]
	content := fmt.Sprintf(
`[Unit]
Description=Game Flex Match is an open-source project for game battle server cluster on Huawei Cloud.
Documentation=""
After=network-online.target

[Service]
User=root
Group=root
LimitNOFILE=65536
ExecStart=%s
KillMode=control-group
StandardOutput=null
Restart=on-failure
RestartSec=30

[Install]
WantedBy=multi-user.target
Alias=%s
`, bin_path, service)
	file, err := os.Create(fmt.Sprintf("/etc/systemd/system/%s.service", service))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	fmt.Printf("Systemd service file for %s created successfully\n", service)
	return nil
}

func (a *InstallAction) touchStartScript(param map[string]string) error {
	service := param["service"]
	bin_path := param["bin"]
	conf_path := param["conf"]
	log_path := param["log"]
	content := fmt.Sprintf(
		`#!/bin/bash
%s -config-path %s -log-path %s
`, bin_path, conf_path, log_path)
	file, err := os.Create(fmt.Sprintf("%s/bin/%s-start.sh", common.RootPath, service))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	if err := os.Chmod(fmt.Sprintf("%s/bin/%s-start.sh", common.RootPath, service), 750); err != nil {
		return err
	}
	fmt.Printf("start script for %s created successfully\n", service)
	return nil
}
