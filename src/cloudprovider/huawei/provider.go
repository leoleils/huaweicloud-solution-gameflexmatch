// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package huawei 华为云Provider实现
package huawei

import (
	"net/http"

	"scase.io/cloudprovider"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	as "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/as/v1"
	ces "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1"
	dns "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	evs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2"
	ims "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
)

const defaultSdkRetries = 3

func init() {
	cloudprovider.RegisterProvider(cloudprovider.ProviderHuawei, func() cloudprovider.CloudProvider {
		return &HuaweiProvider{}
	})
}

// HuaweiProvider 华为云Provider实现
type HuaweiProvider struct {
	config     *cloudprovider.ProviderConfig
	credential *basic.Credentials
	httpConfig *config.HttpConfig

	computeService  *HuaweiComputeService
	networkService  *HuaweiNetworkService
	scalingService  *HuaweiScalingService
	storageService  *HuaweiStorageService
	identityService *HuaweiIdentityService
	dnsService      *HuaweiDNSService
	monitorService  *HuaweiMonitorService
	imageService    *HuaweiImageService
}

// Name 获取厂商名称
func (p *HuaweiProvider) Name() string {
	return cloudprovider.ProviderHuawei
}

// Initialize 初始化Provider
func (p *HuaweiProvider) Initialize(cfg *cloudprovider.ProviderConfig) error {
	p.config = cfg

	// 初始化认证
	p.credential = basic.NewCredentialsBuilder().
		WithAk(cfg.AccessKey).
		WithSk(cfg.SecretKey).
		WithProjectId(cfg.ProjectId).
		Build()

	// 初始化HTTP配置
	p.httpConfig = config.DefaultHttpConfig().
		WithIgnoreSSLVerification(true).
		WithHttpHandler(httphandler.NewHttpHandler().
			AddRequestHandler(requestHandler).
			AddResponseHandler(responseHandler)).
		WithRetries(defaultSdkRetries)

	// 初始化各子服务
	p.computeService = &HuaweiComputeService{
		provider: p,
		client:   p.newEcsClient(),
	}
	p.networkService = &HuaweiNetworkService{
		provider:  p,
		vpcClient: p.newVpcClient(),
	}
	p.scalingService = &HuaweiScalingService{
		provider: p,
		client:   p.newAsClient(),
	}
	p.storageService = &HuaweiStorageService{
		provider:  p,
		evsClient: p.newEvsClient(),
	}
	p.identityService = &HuaweiIdentityService{
		provider: p,
	}
	p.dnsService = &HuaweiDNSService{
		provider: p,
		client:   p.newDnsClient(),
	}
	p.monitorService = &HuaweiMonitorService{
		provider: p,
		client:   p.newCesClient(),
	}
	p.imageService = &HuaweiImageService{
		provider: p,
		client:   p.newImsClient(),
	}

	return nil
}

// Compute 获取计算服务
func (p *HuaweiProvider) Compute() cloudprovider.ComputeService {
	return p.computeService
}

// Network 获取网络服务
func (p *HuaweiProvider) Network() cloudprovider.NetworkService {
	return p.networkService
}

// Scaling 获取弹性伸缩服务
func (p *HuaweiProvider) Scaling() cloudprovider.ScalingService {
	return p.scalingService
}

// Storage 获取存储服务
func (p *HuaweiProvider) Storage() cloudprovider.StorageService {
	return p.storageService
}

// Identity 获取身份认证服务
func (p *HuaweiProvider) Identity() cloudprovider.IdentityService {
	return p.identityService
}

// DNS 获取DNS服务
func (p *HuaweiProvider) DNS() cloudprovider.DNSService {
	return p.dnsService
}

// Monitor 获取监控服务
func (p *HuaweiProvider) Monitor() cloudprovider.MonitorService {
	return p.monitorService
}

// Image 获取镜像服务
func (p *HuaweiProvider) Image() cloudprovider.ImageService {
	return p.imageService
}

// getEndpoint 获取服务端点
func (p *HuaweiProvider) getEndpoint(serviceName string) string {
	if p.config.Endpoints != nil {
		if endpoint, ok := p.config.Endpoints[serviceName]; ok {
			return endpoint
		}
	}
	return ""
}

// newEcsClient 创建ECS客户端
func (p *HuaweiProvider) newEcsClient() *ecs.EcsClient {
	return ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithEndpoint(p.getEndpoint("ecs")).
			WithCredential(p.credential).
			WithHttpConfig(p.httpConfig).
			Build(),
	)
}

// newVpcClient 创建VPC客户端
func (p *HuaweiProvider) newVpcClient() *vpc.VpcClient {
	return vpc.NewVpcClient(
		vpc.VpcClientBuilder().
			WithEndpoint(p.getEndpoint("vpc")).
			WithCredential(p.credential).
			WithHttpConfig(p.httpConfig).
			Build(),
	)
}

// newAsClient 创建AS客户端
func (p *HuaweiProvider) newAsClient() *as.AsClient {
	return as.NewAsClient(
		as.AsClientBuilder().
			WithEndpoint(p.getEndpoint("as")).
			WithCredential(p.credential).
			WithHttpConfig(p.httpConfig).
			Build(),
	)
}

// newEvsClient 创建EVS客户端
func (p *HuaweiProvider) newEvsClient() *evs.EvsClient {
	return evs.NewEvsClient(
		evs.EvsClientBuilder().
			WithEndpoint(p.getEndpoint("evs")).
			WithCredential(p.credential).
			WithHttpConfig(p.httpConfig).
			Build(),
	)
}

// newDnsClient 创建DNS客户端
func (p *HuaweiProvider) newDnsClient() *dns.DnsClient {
	return dns.NewDnsClient(
		dns.DnsClientBuilder().
			WithEndpoint(p.getEndpoint("dns")).
			WithCredential(p.credential).
			WithHttpConfig(p.httpConfig).
			Build(),
	)
}

// newCesClient 创建CES客户端
func (p *HuaweiProvider) newCesClient() *ces.CesClient {
	return ces.NewCesClient(
		ces.CesClientBuilder().
			WithEndpoint(p.getEndpoint("ces")).
			WithCredential(p.credential).
			WithHttpConfig(p.httpConfig).
			Build(),
	)
}

// newImsClient 创建IMS客户端
func (p *HuaweiProvider) newImsClient() *ims.ImsClient {
	return ims.NewImsClient(
		ims.ImsClientBuilder().
			WithEndpoint(p.getEndpoint("ims")).
			WithCredential(p.credential).
			WithHttpConfig(p.httpConfig).
			Build(),
	)
}

// requestHandler 请求处理器
func requestHandler(request http.Request) {
	// 可以在这里添加日志记录
}

// responseHandler 响应处理器
func responseHandler(response http.Response) {
	// 可以在这里添加日志记录
}
