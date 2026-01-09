// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package alibaba 阿里云Provider实现
package alibaba

import (
	"fmt"
	"os"
	"strings"
	"time"

	"scase.io/cloudprovider"

	openapiv1 "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiv2 "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v7/client"
	ess "github.com/alibabacloud-go/ess-20220222/client"
	"github.com/alibabacloud-go/tea/tea"
	vpc "github.com/alibabacloud-go/vpc-20160428/v2/client"
)

func init() {
	cloudprovider.RegisterProvider(cloudprovider.ProviderAlibaba, func() cloudprovider.CloudProvider {
		return &AlibabaProvider{}
	})
}

// AlibabaProvider 阿里云Provider实现
type AlibabaProvider struct {
	config *cloudprovider.ProviderConfig

	ecsClient *ecs.Client
	essClient *ess.Client
	vpcClient *vpc.Client

	computeService  *AlibabaComputeService
	networkService  *AlibabaNetworkService
	scalingService  *AlibabaScalingService
	storageService  *AlibabaStorageService
	identityService *AlibabaIdentityService
	dnsService      *AlibabaDNSService
	monitorService  *AlibabaMonitorService
	imageService    *AlibabaImageService
}

// Name 获取厂商名称
func (p *AlibabaProvider) Name() string {
	return cloudprovider.ProviderAlibaba
}

// Initialize 初始化Provider
func (p *AlibabaProvider) Initialize(cfg *cloudprovider.ProviderConfig) error {
	debugLog("Initialize called: provider=%s, region=%s, projectId=%s", cfg.ProviderName, cfg.Region, cfg.ProjectId)
	p.config = cfg

	// 初始化ECS客户端
	debugLog("Calling newEcsClient...")
	ecsClient, err := p.newEcsClient()
	if err != nil {
		return err
	}
	p.ecsClient = ecsClient

	// 初始化ESS客户端
	essClient, err := p.newEssClient()
	if err != nil {
		return err
	}
	p.essClient = essClient

	// 初始化VPC客户端
	vpcClient, err := p.newVpcClient()
	if err != nil {
		return err
	}
	p.vpcClient = vpcClient

	// 初始化各子服务
	p.computeService = &AlibabaComputeService{
		provider: p,
		client:   p.ecsClient,
	}
	p.networkService = &AlibabaNetworkService{
		provider: p,
		client:   p.vpcClient,
	}
	p.scalingService = &AlibabaScalingService{
		provider: p,
		client:   p.essClient,
	}
	p.storageService = &AlibabaStorageService{
		provider: p,
	}
	p.identityService = &AlibabaIdentityService{
		provider: p,
	}
	p.dnsService = &AlibabaDNSService{
		provider: p,
	}
	p.monitorService = &AlibabaMonitorService{
		provider: p,
	}
	p.imageService = &AlibabaImageService{
		provider: p,
		client:   p.ecsClient,
	}

	return nil
}

// Compute 获取计算服务
func (p *AlibabaProvider) Compute() cloudprovider.ComputeService {
	return p.computeService
}

// Network 获取网络服务
func (p *AlibabaProvider) Network() cloudprovider.NetworkService {
	return p.networkService
}

// Scaling 获取弹性伸缩服务
func (p *AlibabaProvider) Scaling() cloudprovider.ScalingService {
	return p.scalingService
}

// Storage 获取存储服务
func (p *AlibabaProvider) Storage() cloudprovider.StorageService {
	return p.storageService
}

// Identity 获取身份认证服务
func (p *AlibabaProvider) Identity() cloudprovider.IdentityService {
	return p.identityService
}

// DNS 获取DNS服务
func (p *AlibabaProvider) DNS() cloudprovider.DNSService {
	return p.dnsService
}

// Monitor 获取监控服务
func (p *AlibabaProvider) Monitor() cloudprovider.MonitorService {
	return p.monitorService
}

// Image 获取镜像服务
func (p *AlibabaProvider) Image() cloudprovider.ImageService {
	return p.imageService
}

// getEndpoint 获取服务端点
func (p *AlibabaProvider) getEndpoint(serviceName string) string {
	if p.config.Endpoints != nil {
		if endpoint, ok := p.config.Endpoints[serviceName]; ok {
			// 去除 https:// 或 http:// 前缀，因为阿里云 SDK 会自动添加
			endpoint = strings.TrimPrefix(endpoint, "https://")
			endpoint = strings.TrimPrefix(endpoint, "http://")
			return endpoint
		}
	}
	// 默认端点格式
	return serviceName + "." + p.config.Region + ".aliyuncs.com"
}

// debugLog 写入调试日志到文件
func debugLog(format string, args ...interface{}) {
	f, err := os.OpenFile("/tmp/alibaba_provider_debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	msg := fmt.Sprintf("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	msg += fmt.Sprintf(format, args...)
	msg += "\n"
	f.WriteString(msg)
}

// newEcsClient 创建ECS客户端
// 根据阿里云官方示例直接使用 AccessKeyId 和 AccessKeySecret
func (p *AlibabaProvider) newEcsClient() (*ecs.Client, error) {
	// 构建 Endpoint: ecs.{region}.aliyuncs.com
	endpoint := fmt.Sprintf("ecs.%s.aliyuncs.com", p.config.Region)

	// 调试日志 - 写入文件
	debugLog("Creating ECS client: endpoint=%s, region=%s, ak=%s...", endpoint, p.config.Region, p.config.AccessKey[:8])

	config := &openapiv2.Config{
		AccessKeyId:     tea.String(p.config.AccessKey),
		AccessKeySecret: tea.String(p.config.SecretKey),
		Endpoint:        tea.String(endpoint),
	}

	debugLog("Config: AK=%s, Endpoint=%s", *config.AccessKeyId, *config.Endpoint)

	client, err := ecs.NewClient(config)
	if err != nil {
		debugLog("ECS client creation failed: %v", err)
		return nil, err
	}

	debugLog("ECS client created successfully")
	return client, nil
}

// newEssClient 创建ESS客户端
// 注意：不要手动设置 Endpoint，让 SDK 根据 RegionId 自动选择正确的 endpoint
func (p *AlibabaProvider) newEssClient() (*ess.Client, error) {
	config := &openapiv1.Config{
		AccessKeyId:     tea.String(p.config.AccessKey),
		AccessKeySecret: tea.String(p.config.SecretKey),
		RegionId:        tea.String(p.config.Region),
	}
	return ess.NewClient(config)
}

// newVpcClient 创建VPC客户端
// 根据阿里云最佳实践，显式设置 HTTPS 协议和 SignatureAlgorithm
func (p *AlibabaProvider) newVpcClient() (*vpc.Client, error) {
	config := &openapiv2.Config{
		AccessKeyId:        tea.String(p.config.AccessKey),
		AccessKeySecret:    tea.String(p.config.SecretKey),
		RegionId:           tea.String(p.config.Region),
		Protocol:           tea.String("HTTPS"),
		SignatureAlgorithm: tea.String("v2"),
	}
	return vpc.NewClient(config)
}
