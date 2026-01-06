// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package cloudprovider 云厂商抽象层 - 统一数据模型
package cloudprovider

// ProviderConfig 云厂商配置
type ProviderConfig struct {
	ProviderName string            // 云厂商名称: huawei, alibaba
	Region       string            // 区域
	ProjectId    string            // 项目ID (华为云) / RegionId (阿里云)
	AccessKey    string            // AK
	SecretKey    string            // SK
	Endpoints    map[string]string // 各服务端点
	ExtraConfig  map[string]string // 扩展配置
}

// Tag 标签
type Tag struct {
	Key   string
	Value string
}

// Instance 虚拟机实例
type Instance struct {
	Id               string
	Name             string
	Status           string
	PrivateIPs       []string
	PublicIPs        []string
	FlavorId         string
	ImageId          string
	VpcId            string
	SubnetId         string
	SecurityGroupIds []string
	AvailabilityZone string
	Tags             []Tag
	CreatedAt        string
	RawData          interface{} // 原始云厂商返回数据
}

// CreateInstanceRequest 创建虚拟机请求
type CreateInstanceRequest struct {
	Name                string
	ImageId             string
	FlavorId            string
	VpcId               string
	SubnetId            string
	SecurityGroupIds    []string
	AvailabilityZone    string
	KeyName             string
	UserData            string
	Count               int
	Tags                []Tag
	Disks               []Disk
	Eip                 *EipConfig
	AgencyName          string
	EnterpriseProjectId string
}

// CreateInstanceResponse 创建虚拟机响应
type CreateInstanceResponse struct {
	JobId       string
	InstanceIds []string
}

// ListInstancesRequest 列出虚拟机请求
type ListInstancesRequest struct {
	InstanceIds []string
	Name        string
	Status      string
	VpcId       string
	Offset      int
	Limit       int
	Tags        []Tag
}

// ListInstancesResponse 列出虚拟机响应
type ListInstancesResponse struct {
	Instances  []Instance
	TotalCount int
}

// Disk 磁盘配置
type Disk struct {
	DiskType   string // SYS, DATA
	VolumeType string // SATA, SAS, SSD, GPSSD, ESSD
	Size       int32
}

// EipConfig 弹性公网IP配置
type EipConfig struct {
	IpType        string
	BandwidthSize int32
	ShareType     string // PER, WHOLE
	ChargeMode    string // bandwidth, traffic
	BandwidthId   string
}

// JobInfo 异步任务信息
type JobInfo struct {
	Id       string
	Status   string // SUCCESS, FAIL, RUNNING
	Progress int
	SubJobs  []SubJobInfo
}

// SubJobInfo 子任务信息
type SubJobInfo struct {
	Id       string
	Status   string
	ServerId string
}

// Vpc VPC
type Vpc struct {
	Id                  string
	Name                string
	Cidr                string
	Status              string
	EnterpriseProjectId string
}

// CreateVpcRequest 创建VPC请求
type CreateVpcRequest struct {
	Name                string
	Cidr                string
	EnterpriseProjectId string
}

// Subnet 子网
type Subnet struct {
	Id               string
	Name             string
	Cidr             string
	VpcId            string
	GatewayIp        string
	AvailabilityZone string
	Status           string
	NeutronSubnetId  string // 华为云特有
	VSwitchId        string // 阿里云特有
}

// CreateSubnetRequest 创建子网请求
type CreateSubnetRequest struct {
	Name         string
	Cidr         string
	VpcId        string
	GatewayIp    string
	ZoneId       string
	PrimaryDns   string
	SecondaryDns string
	DnsList      []string
}

// CreateSubnetResponse 创建子网响应
type CreateSubnetResponse struct {
	SubnetId        string
	NeutronSubnetId string // 华为云特有
	VSwitchId       string // 阿里云特有
}

// SecurityGroup 安全组
type SecurityGroup struct {
	Id   string
	Name string
}

// CreateSecurityGroupRequest 创建安全组请求
type CreateSecurityGroupRequest struct {
	Name                string
	VpcId               string // 阿里云需要
	EnterpriseProjectId string
}

// CreateEipRequest 创建弹性公网IP请求
type CreateEipRequest struct {
	Name                string
	BandwidthSize       int32
	BandwidthType       string
	BandwidthMode       string
	EnterpriseProjectId string
}

// CreateBandwidthRequest 创建带宽请求
type CreateBandwidthRequest struct {
	Name          string
	Size          int32
	BandwidthType string
}

// Bandwidth 带宽
type Bandwidth struct {
	Id   string
	Name string
	Size int32
}

// ScalingGroup 伸缩组
type ScalingGroup struct {
	Id                    string
	Name                  string
	Status                string
	CurrentInstanceNumber int32
	DesireInstanceNumber  int32
	MinInstanceNumber     int32
	MaxInstanceNumber     int32
	ConfigId              string
	VpcId                 string
	SubnetId              string
}

// CreateScalingConfigRequest 创建伸缩配置请求
type CreateScalingConfigRequest struct {
	Name           string
	ImageId        string
	FlavorIds      []string
	Disks          []Disk
	Eip            *EipConfig
	KeyName        string
	SecurityGroups []string
	UserData       string
}

// CreateScalingGroupRequest 创建伸缩组请求
type CreateScalingGroupRequest struct {
	Name                 string
	ConfigId             string
	VpcId                string
	SubnetId             string
	MinInstanceNumber    int32
	MaxInstanceNumber    int32
	DesireInstanceNumber int32
	EnterpriseProjectId  string
	IamAgencyName        string
	DeletePublicIp       bool
	DeleteVolume         bool
}

// UpdateScalingGroupRequest 更新伸缩组请求
type UpdateScalingGroupRequest struct {
	DesireInstanceNumber *int32
	MinInstanceNumber    *int32
	MaxInstanceNumber    *int32
}

// ObjectMetadata 对象元数据
type ObjectMetadata struct {
	ContentLength int64
	ContentType   string
	ETag          string
	LastModified  string
}

// VolumeType 卷类型
type VolumeType struct {
	Id             string
	Name           string
	AvailableZones []string
	IsPublic       bool
}

// Project 项目
type Project struct {
	Id       string
	Name     string
	DomainId string
}

// GetMetricsRequest 获取监控指标请求
type GetMetricsRequest struct {
	Namespace  string
	MetricName string
	Dimensions []MetricDimension
	StartTime  int64
	EndTime    int64
	Period     int
}

// MetricDimension 监控维度
type MetricDimension struct {
	Name  string
	Value string
}

// MetricsResponse 监控指标响应
type MetricsResponse struct {
	Datapoints []Datapoint
}

// Datapoint 监控数据点
type Datapoint struct {
	Timestamp int64
	Average   float64
	Max       float64
	Min       float64
	Sum       float64
}

// CreateAlarmRequest 创建告警请求
type CreateAlarmRequest struct {
	Name       string
	Namespace  string
	MetricName string
	Threshold  float64
	Period     int
}

// Image 镜像
type Image struct {
	Id       string
	Name     string
	Status   string
	Platform string
	OsType   string
}

// 常量定义
const (
	// ProviderHuawei 华为云
	ProviderHuawei = "huawei"
	// ProviderAlibaba 阿里云
	ProviderAlibaba = "alibaba"

	// DiskTypeSys 系统盘
	DiskTypeSys = "SYS"
	// DiskTypeData 数据盘
	DiskTypeData = "DATA"

	// VolumeTypeSATA SATA
	VolumeTypeSATA = "SATA"
	// VolumeTypeSAS SAS
	VolumeTypeSAS = "SAS"
	// VolumeTypeSSD SSD
	VolumeTypeSSD = "SSD"
	// VolumeTypeGPSSD GPSSD
	VolumeTypeGPSSD = "GPSSD"
	// VolumeTypeESSD ESSD
	VolumeTypeESSD = "ESSD"

	// ShareTypePer 独享带宽
	ShareTypePer = "PER"
	// ShareTypeWhole 共享带宽
	ShareTypeWhole = "WHOLE"

	// ChargeModeTraffic 按流量计费
	ChargeModeTraffic = "traffic"
	// ChargeModeBandwidth 按带宽计费
	ChargeModeBandwidth = "bandwidth"

	// IPTypePublic 公网IP
	IPTypePublic = "publicIP"
	// IPTypePrivate 内网IP
	IPTypePrivate = "privateIP"

	// InstanceStatusRunning 运行中
	InstanceStatusRunning = "RUNNING"
	// InstanceStatusStopped 已停止
	InstanceStatusStopped = "STOPPED"
	// InstanceStatusShutoff 已关机 (华为云)
	InstanceStatusShutoff = "SHUTOFF"
	// InstanceStatusDeleted 已删除
	InstanceStatusDeleted = "DELETED"

	// JobStatusSuccess 成功
	JobStatusSuccess = "SUCCESS"
	// JobStatusFail 失败
	JobStatusFail = "FAIL"
	// JobStatusRunning 运行中
	JobStatusRunning = "RUNNING"
)
