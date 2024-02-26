package model

type FleetManager struct {
	MysqlConfig                map[string]interface{} `yaml:"mysqlConfig"`
	RedisConfig                map[string]interface{} `yaml:"redisConfig"`
	AuthConfig                 map[string]interface{} `yaml:"authConfig"`
	LogConfig                  map[string]interface{} `yaml:"logConfig"`
	EnvConfig                  map[string]interface{} `yaml:"envConfig"`
	FleetConfig                map[string]interface{} `yaml:"fleetConfig"`
	InternalInboundPermissions []interface{}          `yaml:"internalInboundPermissions"`
	BuildConfig                map[string]interface{} `yaml:"buildConfig"`
	GroupConfig                map[string]interface{} `yaml:"groupConfig"`
	ServiceEndpoint            map[string]interface{} `yaml:"serviceEndpoint"`
	SessionConfig              map[string]interface{} `yaml:"sessionConfig"`
	HttpsConfig                map[string]interface{} `yaml:"httpsConfig"`
	WorkNodeConfig             map[string]interface{} `yaml:"workNodeConfig"`
}

type AppGateway struct {
	MysqlConfig    map[string]interface{} `yaml:"mysqlConfig"`
	InfluxDBConfig map[string]interface{} `yaml:"influxDBConfig"`
	RedisConfig    map[string]interface{} `yaml:"redisConfig"`
	CleanConfig    map[string]interface{} `yaml:"cleanConfig"`
	EnvConfig      map[string]interface{} `yaml:"envConfig"`
	HttpsConfig    map[string]interface{} `yaml:"httpsConfig"`
	LogConfig      map[string]interface{} `yaml:"logConfig"`
	AuthConfig     map[string]interface{} `yaml:"authConfig"`
}

type AASS struct {
	MysqlConfig        map[string]interface{} `yaml:"mysqlConfig"`
	RedisConfig        map[string]interface{} `yaml:"redisConfig"`
	AuthConfig         map[string]interface{} `yaml:"authConfig"`
	LogConfig          map[string]interface{} `yaml:"logConfig"`
	ServiceEndpoint    map[string]interface{} `yaml:"serviceEndpoint"`
	CloudEndpoint      map[string]interface{} `yaml:"cloudEndpoint"`
	InfluxDBConfig     map[string]interface{} `yaml:"influxDBConfig"`
	EnvConfig          map[string]interface{} `yaml:"envConfig"`
	HttpsConfig        map[string]interface{} `yaml:"httpsConfig"`
	WorkNodeConfig     map[string]interface{} `yaml:"workNodeConfig"`
	ScalingGroupConfig map[string]interface{} `yaml:"scalingGroupConfig"`
}
