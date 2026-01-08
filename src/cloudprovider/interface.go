// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package cloudprovider 云厂商抽象层 - 核心接口定义
package cloudprovider

import "context"

// CloudProvider 云厂商顶级接口
type CloudProvider interface {
	// Name 获取厂商名称
	Name() string
	// Initialize 初始化Provider
	Initialize(config *ProviderConfig) error
	// Compute 获取计算服务
	Compute() ComputeService
	// Network 获取网络服务
	Network() NetworkService
	// Scaling 获取弹性伸缩服务
	Scaling() ScalingService
	// Storage 获取存储服务
	Storage() StorageService
	// Identity 获取身份认证服务
	Identity() IdentityService
	// DNS 获取DNS服务
	DNS() DNSService
	// Monitor 获取监控服务
	Monitor() MonitorService
	// Image 获取镜像服务
	Image() ImageService
}

// ComputeService 计算服务接口 (华为云ECS / 阿里云ECS / AWS EC2)
type ComputeService interface {
	// CreateInstance 创建虚拟机实例
	CreateInstance(ctx context.Context, req *CreateInstanceRequest) (*CreateInstanceResponse, error)
	// DeleteInstance 删除虚拟机实例
	DeleteInstance(ctx context.Context, instanceId string) error
	// BatchDeleteInstances 批量删除虚拟机实例
	BatchDeleteInstances(ctx context.Context, instanceIds []string) (string, error)
	// GetInstance 获取虚拟机实例信息
	GetInstance(ctx context.Context, instanceId string) (*Instance, error)
	// ListInstances 列出虚拟机实例
	ListInstances(ctx context.Context, req *ListInstancesRequest) (*ListInstancesResponse, error)
	// StartInstance 启动虚拟机实例
	StartInstance(ctx context.Context, instanceId string) error
	// StopInstance 停止虚拟机实例
	StopInstance(ctx context.Context, instanceId string) error
	// BatchStopInstances 批量停止虚拟机实例
	BatchStopInstances(ctx context.Context, instanceIds []string) error
	// ListAvailabilityZones 列出可用区
	ListAvailabilityZones(ctx context.Context) ([]string, error)
	// ListFlavors 列出规格
	ListFlavors(ctx context.Context) (map[string][]string, error)
	// GetJob 获取异步任务状态
	GetJob(ctx context.Context, jobId string) (*JobInfo, error)
	// UpdateInstanceTags 更新实例标签
	UpdateInstanceTags(ctx context.Context, instanceId string, tags []Tag) error
	// ParseInstanceIPs 解析实例IP地址
	ParseInstanceIPs(ctx context.Context, instance *Instance, ipType string) ([]string, error)
}

// NetworkService 网络服务接口 (华为云VPC / 阿里云VPC / AWS VPC)
type NetworkService interface {
	// VPC操作
	CreateVpc(ctx context.Context, req *CreateVpcRequest) (string, error)
	GetVpc(ctx context.Context, vpcId string) (*Vpc, error)
	GetVpcByName(ctx context.Context, name string) (*Vpc, error)
	ListVpcs(ctx context.Context) ([]Vpc, error)
	DeleteVpc(ctx context.Context, vpcId string) error
	WaitVpcReady(ctx context.Context, vpcId string) error
	WaitVpcDeleted(ctx context.Context, vpcId string) error

	// 子网操作
	CreateSubnet(ctx context.Context, req *CreateSubnetRequest) (*CreateSubnetResponse, error)
	GetSubnet(ctx context.Context, subnetId string) (*Subnet, error)
	ListSubnets(ctx context.Context, vpcId string) ([]Subnet, error)
	DeleteSubnet(ctx context.Context, vpcId, subnetId string) error
	WaitSubnetReady(ctx context.Context, subnetId string) error
	WaitSubnetDeleted(ctx context.Context, subnetId string) error
	CheckSubnetConflict(ctx context.Context, vpcId, cidr string) (bool, error)

	// 安全组操作
	CreateSecurityGroup(ctx context.Context, req *CreateSecurityGroupRequest) (string, error)
	GetSecurityGroup(ctx context.Context, name string) (string, error)
	GetSecurityGroupById(ctx context.Context, id string) (*SecurityGroup, error)
	ListSecurityGroups(ctx context.Context) ([]SecurityGroup, error)
	// 安全组规则操作
	CreateSecurityGroupRule(ctx context.Context, req *CreateSecurityGroupRuleRequest) (string, error)
	DeleteSecurityGroupRule(ctx context.Context, ruleId string) error

	// EIP/弹性公网IP操作
	CreateEip(ctx context.Context, req *CreateEipRequest) (string, error)
	DeleteEip(ctx context.Context, eipId string) error
	// BindEipToInstance 将EIP绑定到实例 (阿里云需要单独调用)
	BindEipToInstance(ctx context.Context, eipId, instanceId string) error
	// UnbindEipFromInstance 解绑EIP
	UnbindEipFromInstance(ctx context.Context, eipId, instanceId string) error
	CreateBandwidth(ctx context.Context, req *CreateBandwidthRequest) (string, error)
	ListBandwidths(ctx context.Context) ([]Bandwidth, error)
}

// ScalingService 弹性伸缩服务接口 (华为云AS / 阿里云ESS / AWS Auto Scaling)
type ScalingService interface {
	// 伸缩配置
	CreateScalingConfig(ctx context.Context, req *CreateScalingConfigRequest) (string, error)
	DeleteScalingConfig(ctx context.Context, configId string) error

	// 伸缩组
	CreateScalingGroup(ctx context.Context, req *CreateScalingGroupRequest) (string, error)
	DeleteScalingGroup(ctx context.Context, groupId string, forceDelete bool) error
	EnableScalingGroup(ctx context.Context, groupId string) error
	DisableScalingGroup(ctx context.Context, groupId string) error
	GetScalingGroup(ctx context.Context, groupId string) (*ScalingGroup, error)
	UpdateScalingGroup(ctx context.Context, groupId string, req *UpdateScalingGroupRequest) error
	WaitScalingGroupStable(ctx context.Context, groupId string) error

	// 伸缩实例
	ListScalingInstances(ctx context.Context, groupId string) ([]string, error)
	BatchRemoveScalingInstances(ctx context.Context, groupId string, instanceIds []string, deleteInstance bool) error

	// 标签
	UpdateScalingGroupTags(ctx context.Context, groupId string, tags []Tag) error
	CreateScalingGroupTags(ctx context.Context, groupId string, tags []Tag) error
	DeleteScalingGroupTags(ctx context.Context, groupId string, tags []Tag) error
}

// StorageService 存储服务接口 (华为云OBS/EVS / 阿里云OSS/EBS / AWS S3/EBS)
type StorageService interface {
	// 对象存储
	CreateBucket(ctx context.Context, bucketName string) error
	DeleteBucket(ctx context.Context, bucketName string) error
	HeadBucket(ctx context.Context, bucketName string) (bool, error)
	DeleteObject(ctx context.Context, bucketName, objectKey string) error
	GetObjectMetadata(ctx context.Context, bucketName, objectKey string) (*ObjectMetadata, error)
	UploadObject(ctx context.Context, bucketName, objectKey string, data []byte) error
	// CreateSignedUrl 创建签名URL用于下载对象
	CreateSignedUrl(ctx context.Context, bucketName, objectKey string, expireSeconds int64) (string, error)
	// ListBuckets 列出所有存储桶
	ListBuckets(ctx context.Context) ([]string, error)

	// 块存储
	ListVolumeTypes(ctx context.Context) ([]VolumeType, error)
	GetAvailableZonesByVolumeType(ctx context.Context, volumeType string, azList []string) ([]string, error)
}

// IdentityService 身份认证服务接口 (华为云IAM / 阿里云RAM / AWS IAM)
type IdentityService interface {
	// GetToken 获取Token
	GetToken(ctx context.Context) (string, error)
	// ValidateCredentials 验证凭证
	ValidateCredentials(ctx context.Context) error
	// ListProjects 列出项目
	ListProjects(ctx context.Context) ([]Project, error)
}

// DNSService DNS服务接口
type DNSService interface {
	// GetZoneId 获取域名Zone ID
	GetZoneId(ctx context.Context, domainName string) (string, error)
	// CreateRecordSet 创建DNS记录
	CreateRecordSet(ctx context.Context, zoneId, recordName, recordType string, records []string) (string, error)
	// DeleteRecordSet 删除DNS记录
	DeleteRecordSet(ctx context.Context, zoneId, recordSetId string) error
	// BatchDeleteRecordSets 批量删除DNS记录
	BatchDeleteRecordSets(ctx context.Context, zoneId string, recordSetIds []string) error
}

// MonitorService 监控服务接口 (华为云CES / 阿里云CloudMonitor / AWS CloudWatch)
type MonitorService interface {
	// GetMetrics 获取监控指标
	GetMetrics(ctx context.Context, req *GetMetricsRequest) (*MetricsResponse, error)
	// CreateAlarm 创建告警
	CreateAlarm(ctx context.Context, req *CreateAlarmRequest) (string, error)
	// DeleteAlarm 删除告警
	DeleteAlarm(ctx context.Context, alarmId string) error
}

// ImageService 镜像服务接口 (华为云IMS / 阿里云ECS Image / AWS AMI)
type ImageService interface {
	// GetPublicImageId 获取公共镜像ID
	GetPublicImageId(ctx context.Context, imageName string) (string, error)
	// ListImages 列出镜像
	ListImages(ctx context.Context, imageType string) ([]Image, error)
	// GetImageById 根据ID获取镜像详情
	GetImageById(ctx context.Context, imageId string) (*Image, error)
	// CreateImage 从ECS实例创建镜像
	CreateImage(ctx context.Context, instanceId, imageName string) (string, error)
	// WaitImageReady 等待镜像创建完成并返回镜像ID
	WaitImageReady(ctx context.Context, taskId string) (string, error)
	// DeleteImage 删除镜像
	DeleteImage(ctx context.Context, imageId string) error
}
