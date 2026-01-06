// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package alibaba 阿里云Provider实现
package alibaba

import (
	"scase.io/cloudprovider"

	openapiv1 "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiv2 "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
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
	p.config = cfg

	// 初始化ECS客户端
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
			return endpoint
		}
	}
	// 默认端点格式
	return serviceName + "." + p.config.Region + ".aliyuncs.com"
}

// newEcsClient 创建ECS客户端
func (p *AlibabaProvider) newEcsClient() (*ecs.Client, error) {
	config := &openapiv2.Config{
		AccessKeyId:     tea.String(p.config.AccessKey),
		AccessKeySecret: tea.String(p.config.SecretKey),
		Endpoint:        tea.String(p.getEndpoint("ecs")),
	}
	return ecs.NewClient(config)
}

// newEssClient 创建ESS客户端
func (p *AlibabaProvider) newEssClient() (*ess.Client, error) {
	config := &openapiv1.Config{
		AccessKeyId:     tea.String(p.config.AccessKey),
		AccessKeySecret: tea.String(p.config.SecretKey),
		Endpoint:        tea.String(p.getEndpoint("ess")),
	}
	return ess.NewClient(config)
}

// newVpcClient 创建VPC客户端
func (p *AlibabaProvider) newVpcClient() (*vpc.Client, error) {
	config := &openapiv1.Config{
		AccessKeyId:     tea.String(p.config.AccessKey),
		AccessKeySecret: tea.String(p.config.SecretKey),
		Endpoint:        tea.String(p.getEndpoint("vpc")),
	}
	return vpc.NewClient(config)
}
