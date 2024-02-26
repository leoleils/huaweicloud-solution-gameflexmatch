package common

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
	"os"
	"gopkg.in/yaml.v3"
)

const DefaultConfPath = "./conf"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const (
	FleetmanagerJwtKey = "SessionJwtKey"
	CipherGcmKey = "GCMKey"
	CipherGcmNonce = "GCMNonce"
	CipherRsaPublicKey = "RSAPublicKey"
	CipherRsaPrivateKey = "RSAPrivateKey"
	CipherTlsKey= "TlsKey"
	CipherTlsCsr= "TlsCsr"
	CipherTlsCrt= "TlsCrt"
)

// 配置文件内块相关
const (
	SectionCloud = "cloud"
	SectionCipher = "cipher"
	SectionFleetManager = "fleetmanager"
	SectionAppGateway = "appgateway"
	SectionAASS = "aass"
	SectionInfluxDB = "influxdb"
)

// 配置文件内字段相关
const (
	KeypairName = "KeypairName"
	KeyPairScope = "KeyPairScope"
	ObsBucketName = "ObsBucketName"
	EnterpriseProjectId = "EnterpriseProjectId"
	LtsAgencyName = "LtsAgencyName"
	DomianId = "DomainId"
	CloudRegion = "Region"
)

// 命令行标志位相关
const (
	FlagService = "service"
	FlagConfig = "config"
	FlagMode = "mode"
	FlagMethod = "method"
	FlagText = "text"
	FlagCreate = "create"
	FlagAuthConf = "init-conf"
	FlagAk = "ak"
	FlagSk = "sk"
	FlagObsName = "obs-bucket-name"
	FlagCreateOBSBucket = "create-obs-bucket"
	FlagKeyPairName = "keypair-name"
	FlagCreateKeyPair = "create-keypair"
	FlagUploadAuxproxy = "upload-auxproxy"
	FlagAuxproxyPath = "auxproxy-path"
	FlagRegion = "region"
	FlagUploadFile = "upload-file"
	FlagCreateLtsAgency = "create-lts-agency"
	FlagLtsAgencyName = "lts-agency-name"
	FlagServiceName = "service"
	FlagStartBin = "start-bin"
	FlagStartConfig = "start-config"
	FlagStartLog = "start-log"
)

// 华为云环境信息相关
const (
	EndpointFormat = "https://%s.%s.%s"
	UlanqabEndpoint = "ulanqab.huawei.com"
	MyhuaweicloudEndpoint = "myhuaweicloud.com"
)

const (
	ServiceFleetManager = "fleetmanager"
	ServiceAppGateway = "appgateway"
	ServiceAASS = "aass"
	ServiceConsole = "console"
)

// 其他
const (
	// TODO: 需要修改为真实的AuxProxy路径
	AuxproxyPath = "/gameflexmatch/conf/init.conf"
	ObsAuxproxyName = "auxproxy.zip"
	LtsOpDomainName = "op_svc_lts"

)

var RootPath = ""

func GetConfKey(path string, section string, key string) (string, error) {
	// 该部分解析的是ini文件
	// file, err := ini.Load(path)
	// if err != nil {
	// 	fmt.Println("init conf path err: ", err)
	// }
	// res := file.Section(section).Key(key).String()
	// return res, nil

	// 该部分解析的是yaml文件
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read yaml file: %s", path)
	}
	var conf interface{}
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return "", fmt.Errorf("failed to unmarshal yaml: %s", path)
	}
	sectionValue, ok := conf.(map[string]interface{})[section].(map[string]interface{})
	if ok {
		return sectionValue[key].(string), nil
	}
	return "", fmt.Errorf("failed to get %s -> %s from %s", section, key, path)
}

func SedSprintf(key string, value string) string {
	return fmt.Sprintf("s,%s:.*$,%s: %s,g", key, key, value)
}

func ExecCmd(cmd *exec.Cmd) error {
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("exec cmd: %s error: %+v", cmd.String(), err)
		return err
	}
	fmt.Println(string(output))
	return nil
}

func StringWithCharset(length int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
	  	b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func MkDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsExist(err) {
		fmt.Printf("the dir: %s exists", dirPath)
		return nil
	}
	if err := os.Mkdir(dirPath, os.ModeDir); err != nil {
		fmt.Printf("create dir %s failed %+v", dirPath, err)
		return err
	}
	fmt.Printf("success to create dir: %s", dirPath)
	return nil
}