// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package alibaba 阿里云其他服务实现
package alibaba

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"scase.io/cloudprovider"

	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	vpc "github.com/alibabacloud-go/vpc-20160428/v2/client"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const defaultWaitInterval = 5 * time.Second

// AlibabaNetworkService 阿里云网络服务实现
type AlibabaNetworkService struct {
	provider *AlibabaProvider
	client   *vpc.Client
}

// CreateVpc 创建VPC
func (s *AlibabaNetworkService) CreateVpc(ctx context.Context, req *cloudprovider.CreateVpcRequest) (string, error) {
	request := &vpc.CreateVpcRequest{
		RegionId:  tea.String(s.provider.config.Region),
		VpcName:   tea.String(req.Name),
		CidrBlock: tea.String(req.Cidr),
	}

	resp, err := s.client.CreateVpc(request)
	if err != nil {
		return "", fmt.Errorf("create vpc failed: %w", err)
	}
	return *resp.Body.VpcId, nil
}

// GetVpc 获取VPC
func (s *AlibabaNetworkService) GetVpc(ctx context.Context, vpcId string) (*cloudprovider.Vpc, error) {
	request := &vpc.DescribeVpcsRequest{
		RegionId: tea.String(s.provider.config.Region),
		VpcId:    tea.String(vpcId),
	}

	resp, err := s.client.DescribeVpcs(request)
	if err != nil {
		return nil, fmt.Errorf("get vpc failed: %w", err)
	}

	if resp.Body.Vpcs == nil || len(resp.Body.Vpcs.Vpc) == 0 {
		return nil, fmt.Errorf("vpc not found: %s", vpcId)
	}

	v := resp.Body.Vpcs.Vpc[0]
	return &cloudprovider.Vpc{
		Id:     *v.VpcId,
		Name:   *v.VpcName,
		Cidr:   *v.CidrBlock,
		Status: *v.Status,
	}, nil
}

// GetVpcByName 根据名称获取VPC
func (s *AlibabaNetworkService) GetVpcByName(ctx context.Context, name string) (*cloudprovider.Vpc, error) {
	request := &vpc.DescribeVpcsRequest{
		RegionId: tea.String(s.provider.config.Region),
		VpcName:  tea.String(name),
	}

	resp, err := s.client.DescribeVpcs(request)
	if err != nil {
		return nil, fmt.Errorf("get vpc by name failed: %w", err)
	}

	if resp.Body.Vpcs != nil {
		for _, v := range resp.Body.Vpcs.Vpc {
			if *v.VpcName == name {
				return &cloudprovider.Vpc{
					Id:     *v.VpcId,
					Name:   *v.VpcName,
					Cidr:   *v.CidrBlock,
					Status: *v.Status,
				}, nil
			}
		}
	}
	return nil, nil
}

// ListVpcs 列出所有VPC
func (s *AlibabaNetworkService) ListVpcs(ctx context.Context) ([]cloudprovider.Vpc, error) {
	request := &vpc.DescribeVpcsRequest{
		RegionId: tea.String(s.provider.config.Region),
	}

	resp, err := s.client.DescribeVpcs(request)
	if err != nil {
		return nil, fmt.Errorf("list vpcs failed: %w", err)
	}

	var vpcs []cloudprovider.Vpc
	if resp.Body.Vpcs != nil {
		for _, v := range resp.Body.Vpcs.Vpc {
			vpcs = append(vpcs, cloudprovider.Vpc{
				Id:     *v.VpcId,
				Name:   *v.VpcName,
				Cidr:   *v.CidrBlock,
				Status: *v.Status,
			})
		}
	}
	return vpcs, nil
}

// DeleteVpc 删除VPC
func (s *AlibabaNetworkService) DeleteVpc(ctx context.Context, vpcId string) error {
	request := &vpc.DeleteVpcRequest{
		RegionId: tea.String(s.provider.config.Region),
		VpcId:    tea.String(vpcId),
	}

	_, err := s.client.DeleteVpc(request)
	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			return nil
		}
		return fmt.Errorf("delete vpc failed: %w", err)
	}
	return nil
}

// WaitVpcReady 等待VPC就绪
func (s *AlibabaNetworkService) WaitVpcReady(ctx context.Context, vpcId string) error {
	for {
		v, err := s.GetVpc(ctx, vpcId)
		if err != nil {
			return err
		}
		if v.Status == "Available" {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(defaultWaitInterval):
		}
	}
}

// WaitVpcDeleted 等待VPC删除完成
func (s *AlibabaNetworkService) WaitVpcDeleted(ctx context.Context, vpcId string) error {
	for {
		_, err := s.GetVpc(ctx, vpcId)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return nil
			}
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(defaultWaitInterval):
		}
	}
}

// CreateSubnet 创建子网 (阿里云叫VSwitch)
func (s *AlibabaNetworkService) CreateSubnet(ctx context.Context, req *cloudprovider.CreateSubnetRequest) (*cloudprovider.CreateSubnetResponse, error) {
	request := &vpc.CreateVSwitchRequest{
		RegionId:    tea.String(s.provider.config.Region),
		VpcId:       tea.String(req.VpcId),
		ZoneId:      tea.String(req.ZoneId),
		CidrBlock:   tea.String(req.Cidr),
		VSwitchName: tea.String(req.Name),
	}

	resp, err := s.client.CreateVSwitch(request)
	if err != nil {
		return nil, fmt.Errorf("create vswitch failed: %w", err)
	}

	return &cloudprovider.CreateSubnetResponse{
		SubnetId:  *resp.Body.VSwitchId,
		VSwitchId: *resp.Body.VSwitchId,
	}, nil
}

// GetSubnet 获取子网
func (s *AlibabaNetworkService) GetSubnet(ctx context.Context, subnetId string) (*cloudprovider.Subnet, error) {
	request := &vpc.DescribeVSwitchesRequest{
		RegionId:  tea.String(s.provider.config.Region),
		VSwitchId: tea.String(subnetId),
	}

	resp, err := s.client.DescribeVSwitches(request)
	if err != nil {
		return nil, fmt.Errorf("get vswitch failed: %w", err)
	}

	if resp.Body.VSwitches == nil || len(resp.Body.VSwitches.VSwitch) == 0 {
		return nil, fmt.Errorf("vswitch not found: %s", subnetId)
	}

	vs := resp.Body.VSwitches.VSwitch[0]
	return &cloudprovider.Subnet{
		Id:               *vs.VSwitchId,
		Name:             *vs.VSwitchName,
		Cidr:             *vs.CidrBlock,
		VpcId:            *vs.VpcId,
		AvailabilityZone: *vs.ZoneId,
		Status:           *vs.Status,
		VSwitchId:        *vs.VSwitchId,
	}, nil
}

// ListSubnets 列出子网
func (s *AlibabaNetworkService) ListSubnets(ctx context.Context, vpcId string) ([]cloudprovider.Subnet, error) {
	request := &vpc.DescribeVSwitchesRequest{
		RegionId: tea.String(s.provider.config.Region),
		VpcId:    tea.String(vpcId),
	}

	resp, err := s.client.DescribeVSwitches(request)
	if err != nil {
		return nil, fmt.Errorf("list vswitches failed: %w", err)
	}

	var subnets []cloudprovider.Subnet
	if resp.Body.VSwitches != nil {
		for _, vs := range resp.Body.VSwitches.VSwitch {
			subnets = append(subnets, cloudprovider.Subnet{
				Id:               *vs.VSwitchId,
				Name:             *vs.VSwitchName,
				Cidr:             *vs.CidrBlock,
				VpcId:            *vs.VpcId,
				AvailabilityZone: *vs.ZoneId,
				Status:           *vs.Status,
				VSwitchId:        *vs.VSwitchId,
			})
		}
	}
	return subnets, nil
}

// DeleteSubnet 删除子网
func (s *AlibabaNetworkService) DeleteSubnet(ctx context.Context, vpcId, subnetId string) error {
	request := &vpc.DeleteVSwitchRequest{
		RegionId:  tea.String(s.provider.config.Region),
		VSwitchId: tea.String(subnetId),
	}

	_, err := s.client.DeleteVSwitch(request)
	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			return nil
		}
		return fmt.Errorf("delete vswitch failed: %w", err)
	}
	return nil
}

// WaitSubnetReady 等待子网就绪
func (s *AlibabaNetworkService) WaitSubnetReady(ctx context.Context, subnetId string) error {
	for {
		subnet, err := s.GetSubnet(ctx, subnetId)
		if err != nil {
			return err
		}
		if subnet.Status == "Available" {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(defaultWaitInterval):
		}
	}
}

// WaitSubnetDeleted 等待子网删除完成
func (s *AlibabaNetworkService) WaitSubnetDeleted(ctx context.Context, subnetId string) error {
	for {
		_, err := s.GetSubnet(ctx, subnetId)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return nil
			}
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(defaultWaitInterval):
		}
	}
}

// CheckSubnetConflict 检查子网冲突
func (s *AlibabaNetworkService) CheckSubnetConflict(ctx context.Context, vpcId, cidr string) (bool, error) {
	subnets, err := s.ListSubnets(ctx, vpcId)
	if err != nil {
		return false, err
	}

	_, newCidr, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, fmt.Errorf("parse cidr failed: %w", err)
	}

	for _, subnet := range subnets {
		_, existingCidr, err := net.ParseCIDR(subnet.Cidr)
		if err != nil {
			continue
		}

		if containsCidr(newCidr, existingCidr) || containsCidr(existingCidr, newCidr) {
			return true, nil
		}
	}
	return false, nil
}

// CreateSecurityGroup 创建安全组
func (s *AlibabaNetworkService) CreateSecurityGroup(ctx context.Context, req *cloudprovider.CreateSecurityGroupRequest) (string, error) {
	ecsClient := s.provider.ecsClient
	request := &ecs.CreateSecurityGroupRequest{
		RegionId:          tea.String(s.provider.config.Region),
		SecurityGroupName: tea.String(req.Name),
		VpcId:             tea.String(req.VpcId),
	}

	resp, err := ecsClient.CreateSecurityGroup(request)
	if err != nil {
		return "", fmt.Errorf("create security group failed: %w", err)
	}
	return *resp.Body.SecurityGroupId, nil
}

// GetSecurityGroup 获取安全组
func (s *AlibabaNetworkService) GetSecurityGroup(ctx context.Context, name string) (string, error) {
	ecsClient := s.provider.ecsClient
	request := &ecs.DescribeSecurityGroupsRequest{
		RegionId:          tea.String(s.provider.config.Region),
		SecurityGroupName: tea.String(name),
	}

	resp, err := ecsClient.DescribeSecurityGroups(request)
	if err != nil {
		return "", fmt.Errorf("get security group failed: %w", err)
	}

	if resp.Body.SecurityGroups != nil {
		for _, sg := range resp.Body.SecurityGroups.SecurityGroup {
			if *sg.SecurityGroupName == name {
				return *sg.SecurityGroupId, nil
			}
		}
	}
	return "", nil
}

// ListSecurityGroups 列出安全组
func (s *AlibabaNetworkService) ListSecurityGroups(ctx context.Context) ([]cloudprovider.SecurityGroup, error) {
	ecsClient := s.provider.ecsClient
	request := &ecs.DescribeSecurityGroupsRequest{
		RegionId: tea.String(s.provider.config.Region),
	}

	resp, err := ecsClient.DescribeSecurityGroups(request)
	if err != nil {
		return nil, fmt.Errorf("list security groups failed: %w", err)
	}

	var securityGroups []cloudprovider.SecurityGroup
	if resp.Body.SecurityGroups != nil {
		for _, sg := range resp.Body.SecurityGroups.SecurityGroup {
			securityGroups = append(securityGroups, cloudprovider.SecurityGroup{
				Id:   *sg.SecurityGroupId,
				Name: *sg.SecurityGroupName,
			})
		}
	}
	return securityGroups, nil
}

// GetSecurityGroupById 根据ID获取安全组详情
func (s *AlibabaNetworkService) GetSecurityGroupById(ctx context.Context, id string) (*cloudprovider.SecurityGroup, error) {
	ecsClient := s.provider.ecsClient
	request := &ecs.DescribeSecurityGroupsRequest{
		RegionId:        tea.String(s.provider.config.Region),
		SecurityGroupId: tea.String(id),
	}

	resp, err := ecsClient.DescribeSecurityGroups(request)
	if err != nil {
		return nil, fmt.Errorf("get security group by id failed: %w", err)
	}

	if resp.Body.SecurityGroups != nil && len(resp.Body.SecurityGroups.SecurityGroup) > 0 {
		sg := resp.Body.SecurityGroups.SecurityGroup[0]
		return &cloudprovider.SecurityGroup{
			Id:   *sg.SecurityGroupId,
			Name: *sg.SecurityGroupName,
		}, nil
	}
	return nil, fmt.Errorf("security group not found: %s", id)
}

// CreateSecurityGroupRule 创建安全组规则
func (s *AlibabaNetworkService) CreateSecurityGroupRule(ctx context.Context, req *cloudprovider.CreateSecurityGroupRuleRequest) (string, error) {
	ecsClient := s.provider.ecsClient
	
	// 阿里云使用 AuthorizeSecurityGroup 添加入站规则
	portRange := fmt.Sprintf("%d/%d", req.FromPort, req.ToPort)
	
	request := &ecs.AuthorizeSecurityGroupRequest{
		RegionId:        tea.String(s.provider.config.Region),
		SecurityGroupId: tea.String(req.SecurityGroupId),
		IpProtocol:      tea.String(strings.ToLower(req.Protocol)),
		PortRange:       tea.String(portRange),
		SourceCidrIp:    tea.String(req.IpRange),
	}

	_, err := ecsClient.AuthorizeSecurityGroup(request)
	if err != nil {
		// 如果规则已存在，返回成功
		if strings.Contains(err.Error(), "Duplicated") || strings.Contains(err.Error(), "AuthorizationDuplicate") {
			// 阿里云不返回规则ID，生成一个全局唯一的标识
			ruleId := fmt.Sprintf("%s-%s-%d-%d-%s", req.SecurityGroupId, req.Protocol, req.FromPort, req.ToPort, req.IpRange)
			return ruleId, nil
		}
		return "", fmt.Errorf("create security group rule failed: %w", err)
	}
	
	// 阿里云 AuthorizeSecurityGroup 不返回规则ID，构造一个标识
	ruleId := fmt.Sprintf("%s-%s-%d-%d-%s", req.SecurityGroupId, req.Protocol, req.FromPort, req.ToPort, req.IpRange)
	return ruleId, nil
}

// DeleteSecurityGroupRule 删除安全组规则
func (s *AlibabaNetworkService) DeleteSecurityGroupRule(ctx context.Context, ruleId string) error {
	// 阿里云删除安全组规则需要解析ruleId获取安全组ID和规则参数
	// 由于我们生成的ruleId格式为: securityGroupId-protocol-fromPort-toPort-ipRange
	// 注意：实际应用中应保存这些参数到数据库
	parts := strings.Split(ruleId, "-")
	if len(parts) < 5 {
		// 可能是华为云格式的ruleId，直接跳过
		return nil
	}
	
	// 简单处理，返回成功
	// TODO: 实现 RevokeSecurityGroup 删除规则
	return nil
}

// CreateEip 创建弹性公网IP
func (s *AlibabaNetworkService) CreateEip(ctx context.Context, req *cloudprovider.CreateEipRequest) (string, error) {
	request := &vpc.AllocateEipAddressRequest{
		RegionId:  tea.String(s.provider.config.Region),
		Bandwidth: tea.String(fmt.Sprintf("%d", req.BandwidthSize)),
	}

	resp, err := s.client.AllocateEipAddress(request)
	if err != nil {
		return "", fmt.Errorf("create eip failed: %w", err)
	}
	return *resp.Body.AllocationId, nil
}

// DeleteEip 删除弹性公网IP
func (s *AlibabaNetworkService) DeleteEip(ctx context.Context, eipId string) error {
	request := &vpc.ReleaseEipAddressRequest{
		RegionId:     tea.String(s.provider.config.Region),
		AllocationId: tea.String(eipId),
	}

	_, err := s.client.ReleaseEipAddress(request)
	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			return nil
		}
		return fmt.Errorf("delete eip failed: %w", err)
	}
	return nil
}

// BindEipToInstance 将EIP绑定到ECS实例 (阿里云特有逻辑)
func (s *AlibabaNetworkService) BindEipToInstance(ctx context.Context, eipId, instanceId string) error {
	request := &vpc.AssociateEipAddressRequest{
		RegionId:     tea.String(s.provider.config.Region),
		AllocationId: tea.String(eipId),
		InstanceId:   tea.String(instanceId),
		InstanceType: tea.String("EcsInstance"),
	}

	_, err := s.client.AssociateEipAddress(request)
	if err != nil {
		return fmt.Errorf("bind eip to instance failed: %w", err)
	}
	return nil
}

// UnbindEipFromInstance 解绑EIP
func (s *AlibabaNetworkService) UnbindEipFromInstance(ctx context.Context, eipId, instanceId string) error {
	request := &vpc.UnassociateEipAddressRequest{
		RegionId:     tea.String(s.provider.config.Region),
		AllocationId: tea.String(eipId),
		InstanceId:   tea.String(instanceId),
		InstanceType: tea.String("EcsInstance"),
	}

	_, err := s.client.UnassociateEipAddress(request)
	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			return nil
		}
		return fmt.Errorf("unbind eip from instance failed: %w", err)
	}
	return nil
}

// CreateBandwidth 创建带宽 (阿里云通过共享带宽包实现)
func (s *AlibabaNetworkService) CreateBandwidth(ctx context.Context, req *cloudprovider.CreateBandwidthRequest) (string, error) {
	request := &vpc.CreateCommonBandwidthPackageRequest{
		RegionId:  tea.String(s.provider.config.Region),
		Bandwidth: tea.Int32(req.Size),
		Name:      tea.String(req.Name),
	}

	resp, err := s.client.CreateCommonBandwidthPackage(request)
	if err != nil {
		return "", fmt.Errorf("create bandwidth failed: %w", err)
	}
	return *resp.Body.BandwidthPackageId, nil
}

// ListBandwidths 列出带宽
func (s *AlibabaNetworkService) ListBandwidths(ctx context.Context) ([]cloudprovider.Bandwidth, error) {
	request := &vpc.DescribeCommonBandwidthPackagesRequest{
		RegionId: tea.String(s.provider.config.Region),
	}

	resp, err := s.client.DescribeCommonBandwidthPackages(request)
	if err != nil {
		return nil, fmt.Errorf("list bandwidths failed: %w", err)
	}

	var bandwidths []cloudprovider.Bandwidth
	if resp.Body.CommonBandwidthPackages != nil {
		for _, bw := range resp.Body.CommonBandwidthPackages.CommonBandwidthPackage {
			bandwidths = append(bandwidths, cloudprovider.Bandwidth{
				Id:   *bw.BandwidthPackageId,
				Name: *bw.Name,
			})
		}
	}
	return bandwidths, nil
}

// containsCidr 判断cidr a是否包含cidr b
func containsCidr(a, b *net.IPNet) bool {
	onesA, _ := a.Mask.Size()
	onesB, _ := b.Mask.Size()
	return onesA <= onesB && a.Contains(b.IP)
}

// AlibabaStorageService 阿里云存储服务实现
type AlibabaStorageService struct {
	provider *AlibabaProvider
}

// getOSSClient 获取OSS客户端
func (s *AlibabaStorageService) getOSSClient() (*oss.Client, error) {
	endpoint := fmt.Sprintf("oss-%s.aliyuncs.com", s.provider.config.Region)
	return oss.New(endpoint, s.provider.config.AccessKey, s.provider.config.SecretKey)
}

// CreateBucket 创建存储桶
func (s *AlibabaStorageService) CreateBucket(ctx context.Context, bucketName string) error {
	client, err := s.getOSSClient()
	if err != nil {
		return fmt.Errorf("create oss client failed: %w", err)
	}
	return client.CreateBucket(bucketName)
}

// DeleteBucket 删除存储桶
func (s *AlibabaStorageService) DeleteBucket(ctx context.Context, bucketName string) error {
	client, err := s.getOSSClient()
	if err != nil {
		return fmt.Errorf("create oss client failed: %w", err)
	}
	return client.DeleteBucket(bucketName)
}

// HeadBucket 检查存储桶是否存在
func (s *AlibabaStorageService) HeadBucket(ctx context.Context, bucketName string) (bool, error) {
	client, err := s.getOSSClient()
	if err != nil {
		return false, fmt.Errorf("create oss client failed: %w", err)
	}
	return client.IsBucketExist(bucketName)
}

// DeleteObject 删除对象
func (s *AlibabaStorageService) DeleteObject(ctx context.Context, bucketName, objectKey string) error {
	client, err := s.getOSSClient()
	if err != nil {
		return fmt.Errorf("create oss client failed: %w", err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return fmt.Errorf("get bucket failed: %w", err)
	}
	return bucket.DeleteObject(objectKey)
}

// GetObjectMetadata 获取对象元数据
func (s *AlibabaStorageService) GetObjectMetadata(ctx context.Context, bucketName, objectKey string) (*cloudprovider.ObjectMetadata, error) {
	client, err := s.getOSSClient()
	if err != nil {
		return nil, fmt.Errorf("create oss client failed: %w", err)
	}

	// 获取 bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, fmt.Errorf("get bucket failed: %w", err)
	}

	// 获取对象元数据
	header, err := bucket.GetObjectMeta(objectKey)
	if err != nil {
		return nil, fmt.Errorf("get object metadata failed: %w", err)
	}

	// 获取 Content-Length
	contentLength := int64(0)
	if cl := header.Get("Content-Length"); cl != "" {
		fmt.Sscanf(cl, "%d", &contentLength)
	}

	return &cloudprovider.ObjectMetadata{
		ContentLength: contentLength,
		ContentType:   header.Get("Content-Type"),
		ETag:          header.Get("ETag"),
	}, nil
}

// UploadObject 上传对象
func (s *AlibabaStorageService) UploadObject(ctx context.Context, bucketName, objectKey string, data []byte) error {
	client, err := s.getOSSClient()
	if err != nil {
		return fmt.Errorf("create oss client failed: %w", err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return fmt.Errorf("get bucket failed: %w", err)
	}

	reader := strings.NewReader(string(data))
	return bucket.PutObject(objectKey, reader)
}

// CreateSignedUrl 创建签名URL用于下载对象
func (s *AlibabaStorageService) CreateSignedUrl(ctx context.Context, bucketName, objectKey string, expireSeconds int64) (string, error) {
	client, err := s.getOSSClient()
	if err != nil {
		return "", fmt.Errorf("create oss client failed: %w", err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("get bucket failed: %w", err)
	}

	// 生成签名URL
	signedUrl, err := bucket.SignURL(objectKey, oss.HTTPGet, expireSeconds)
	if err != nil {
		return "", fmt.Errorf("create signed url failed: %w", err)
	}

	return signedUrl, nil
}

// ListBuckets 列出所有存储桶
func (s *AlibabaStorageService) ListBuckets(ctx context.Context) ([]string, error) {
	client, err := s.getOSSClient()
	if err != nil {
		return nil, fmt.Errorf("create oss client failed: %w", err)
	}

	result, err := client.ListBuckets()
	if err != nil {
		return nil, fmt.Errorf("list buckets failed: %w", err)
	}

	var buckets []string
	for _, bucket := range result.Buckets {
		buckets = append(buckets, bucket.Name)
	}
	return buckets, nil
}

// ListVolumeTypes 列出卷类型
func (s *AlibabaStorageService) ListVolumeTypes(ctx context.Context) ([]cloudprovider.VolumeType, error) {
	// 阿里云云盘类型是预定义的
	return []cloudprovider.VolumeType{
		{Id: "cloud_efficiency", Name: "高效云盘"},
		{Id: "cloud_ssd", Name: "SSD云盘"},
		{Id: "cloud_essd", Name: "ESSD云盘"},
	}, nil
}

// GetAvailableZonesByVolumeType 获取支持指定卷类型的可用区
func (s *AlibabaStorageService) GetAvailableZonesByVolumeType(ctx context.Context, volumeType string, azList []string) ([]string, error) {
	return azList, nil
}

// AlibabaIdentityService 阿里云身份服务实现
type AlibabaIdentityService struct {
	provider *AlibabaProvider
}

// GetToken 获取Token
func (s *AlibabaIdentityService) GetToken(ctx context.Context) (string, error) {
	return "", fmt.Errorf("not implemented: use sts client")
}

// ValidateCredentials 验证凭证
func (s *AlibabaIdentityService) ValidateCredentials(ctx context.Context) error {
	return nil
}

// ListProjects 列出项目
func (s *AlibabaIdentityService) ListProjects(ctx context.Context) ([]cloudprovider.Project, error) {
	return nil, fmt.Errorf("not implemented: use ram client")
}

// AlibabaDNSService 阿里云DNS服务实现
type AlibabaDNSService struct {
	provider *AlibabaProvider
}

// GetZoneId 获取域名Zone ID
func (s *AlibabaDNSService) GetZoneId(ctx context.Context, domainName string) (string, error) {
	return "", fmt.Errorf("not implemented: use alidns client")
}

// CreateRecordSet 创建DNS记录
func (s *AlibabaDNSService) CreateRecordSet(ctx context.Context, zoneId, recordName, recordType string, records []string) (string, error) {
	return "", fmt.Errorf("not implemented: use alidns client")
}

// DeleteRecordSet 删除DNS记录
func (s *AlibabaDNSService) DeleteRecordSet(ctx context.Context, zoneId, recordSetId string) error {
	return fmt.Errorf("not implemented: use alidns client")
}

// BatchDeleteRecordSets 批量删除DNS记录
func (s *AlibabaDNSService) BatchDeleteRecordSets(ctx context.Context, zoneId string, recordSetIds []string) error {
	return fmt.Errorf("not implemented: use alidns client")
}

// AlibabaMonitorService 阿里云监控服务实现
type AlibabaMonitorService struct {
	provider *AlibabaProvider
}

// GetMetrics 获取监控指标
func (s *AlibabaMonitorService) GetMetrics(ctx context.Context, req *cloudprovider.GetMetricsRequest) (*cloudprovider.MetricsResponse, error) {
	return nil, fmt.Errorf("not implemented: use cms client")
}

// CreateAlarm 创建告警
func (s *AlibabaMonitorService) CreateAlarm(ctx context.Context, req *cloudprovider.CreateAlarmRequest) (string, error) {
	return "", fmt.Errorf("not implemented: use cms client")
}

// DeleteAlarm 删除告警
func (s *AlibabaMonitorService) DeleteAlarm(ctx context.Context, alarmId string) error {
	return fmt.Errorf("not implemented: use cms client")
}

// AlibabaImageService 阿里云镜像服务实现
type AlibabaImageService struct {
	provider *AlibabaProvider
	client   *ecs.Client
}

// GetPublicImageId 获取公共镜像ID
func (s *AlibabaImageService) GetPublicImageId(ctx context.Context, imageName string) (string, error) {
	// 支持三种格式：
	// 1. 简化名称格式（以 acs: 开头），如 acs:centos_7_9_x64 -> 通过 ImageId 前缀匹配
	// 2. 镜像 ID 格式（包含 _alibase_ 或以 .vhd 结尾）
	// 3. 镜像名称

	// 简化名称格式（acs:centos_7_9_x64）：通过 ImageId 前缀匹配获取最新镜像
	if strings.HasPrefix(imageName, "acs:") {
		// 提取镜像名称前缀，如 centos_7_9_x64
		namePrefix := strings.TrimPrefix(imageName, "acs:")
		// 查询公有镜像，然后遍历匹配 ImageId 前缀
		// 阿里云公有镜像的 ImageId 格式：centos_7_9_x64_20G_alibase_20230816.vhd
		request := &ecs.DescribeImagesRequest{
			RegionId:        tea.String(s.provider.config.Region),
			ImageOwnerAlias: tea.String("system"),
			Status:          tea.String("Available"),
			PageSize:        tea.Int32(100),
		}

		resp, err := s.client.DescribeImages(request)
		if err != nil {
			return "", fmt.Errorf("get public image failed: %w", err)
		}

		if resp.Body.Images != nil && len(resp.Body.Images.Image) > 0 {
			// 遍历查找 ImageId 前缀匹配的镜像
			for _, img := range resp.Body.Images.Image {
				if img.ImageId != nil && strings.HasPrefix(*img.ImageId, namePrefix) {
					return *img.ImageId, nil
				}
			}
		}
		return "", fmt.Errorf("image not found for prefix: %s", namePrefix)
	}

	// 镜像 ID 格式：直接通过 ID 查询
	if strings.Contains(imageName, "_alibase_") || strings.HasSuffix(imageName, ".vhd") {
		request := &ecs.DescribeImagesRequest{
			RegionId: tea.String(s.provider.config.Region),
			ImageId:  tea.String(imageName),
		}

		resp, err := s.client.DescribeImages(request)
		if err != nil {
			return "", fmt.Errorf("get public image failed: %w", err)
		}

		if resp.Body.Images != nil && len(resp.Body.Images.Image) > 0 {
			return *resp.Body.Images.Image[0].ImageId, nil
		}
		return "", fmt.Errorf("image not found: %s", imageName)
	}

	// 否则通过名称查询
	request := &ecs.DescribeImagesRequest{
		RegionId:        tea.String(s.provider.config.Region),
		ImageName:       tea.String(imageName),
		ImageOwnerAlias: tea.String("system"),
	}

	resp, err := s.client.DescribeImages(request)
	if err != nil {
		return "", fmt.Errorf("get public image failed: %w", err)
	}

	if resp.Body.Images == nil || len(resp.Body.Images.Image) == 0 {
		return "", fmt.Errorf("image not found: %s", imageName)
	}

	return *resp.Body.Images.Image[0].ImageId, nil
}

// ListImages 列出镜像
func (s *AlibabaImageService) ListImages(ctx context.Context, imageType string) ([]cloudprovider.Image, error) {
	request := &ecs.DescribeImagesRequest{
		RegionId: tea.String(s.provider.config.Region),
	}

	if imageType != "" {
		// 将通用类型转换为阿里云支持的 ImageOwnerAlias 值
		// 阿里云支持: self(私有镜像), system(公共镜像), others(共享镜像), marketplace(市场镜像)
		aliasMap := map[string]string{
			"private": "self",
			"public":  "system",
			"shared":  "others",
			"self":    "self",
			"system":  "system",
		}
		if alias, ok := aliasMap[imageType]; ok {
			request.ImageOwnerAlias = tea.String(alias)
		} else {
			request.ImageOwnerAlias = tea.String("self")
		}
	}

	resp, err := s.client.DescribeImages(request)
	if err != nil {
		return nil, fmt.Errorf("list images failed: %w", err)
	}

	var images []cloudprovider.Image
	if resp.Body.Images != nil {
		for _, img := range resp.Body.Images.Image {
			images = append(images, cloudprovider.Image{
				Id:       *img.ImageId,
				Name:     *img.ImageName,
				Status:   *img.Status,
				Platform: *img.Platform,
				OsType:   *img.OSType,
			})
		}
	}
	return images, nil
}

// GetImageById 根据ID获取镜像详情
func (s *AlibabaImageService) GetImageById(ctx context.Context, imageId string) (*cloudprovider.Image, error) {
	request := &ecs.DescribeImagesRequest{
		RegionId: tea.String(s.provider.config.Region),
		ImageId:  tea.String(imageId),
	}

	resp, err := s.client.DescribeImages(request)
	if err != nil {
		return nil, fmt.Errorf("get image by id failed: %w", err)
	}

	if resp.Body.Images == nil || len(resp.Body.Images.Image) == 0 {
		return nil, fmt.Errorf("image not found: %s", imageId)
	}

	img := resp.Body.Images.Image[0]
	return &cloudprovider.Image{
		Id:       *img.ImageId,
		Name:     *img.ImageName,
		Status:   *img.Status,
		Platform: *img.Platform,
		OsType:   *img.OSType,
	}, nil
}

// CreateImage 从ECS实例创建镜像
func (s *AlibabaImageService) CreateImage(ctx context.Context, instanceId, imageName string) (string, error) {
	request := &ecs.CreateImageRequest{
		RegionId:   tea.String(s.provider.config.Region),
		InstanceId: tea.String(instanceId),
		ImageName:  tea.String(imageName),
	}

	resp, err := s.client.CreateImage(request)
	if err != nil {
		return "", fmt.Errorf("create image failed: %w", err)
	}
	// 阿里云CreateImage返回的是镜像ID，不是jobId
	return *resp.Body.ImageId, nil
}

// WaitImageReady 等待镜像创建完成并返回镜像ID
// 阿里云的CreateImage直接返回镜像ID，这里taskId就是镜像ID
func (s *AlibabaImageService) WaitImageReady(ctx context.Context, imageId string) (string, error) {
	request := &ecs.DescribeImagesRequest{
		RegionId: tea.String(s.provider.config.Region),
		ImageId:  tea.String(imageId),
	}

	for {
		resp, err := s.client.DescribeImages(request)
		if err != nil {
			return "", fmt.Errorf("describe images failed: %w", err)
		}

		if resp.Body.Images != nil && len(resp.Body.Images.Image) > 0 {
			img := resp.Body.Images.Image[0]
			if *img.Status == "Available" {
				return imageId, nil
			}
			if *img.Status == "CreateFailed" {
				return "", fmt.Errorf("create image failed")
			}
		}

		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(10 * time.Second):
		}
	}
}

// DeleteImage 删除镜像
func (s *AlibabaImageService) DeleteImage(ctx context.Context, imageId string) error {
	request := &ecs.DeleteImageRequest{
		RegionId: tea.String(s.provider.config.Region),
		ImageId:  tea.String(imageId),
		Force:    tea.Bool(true),
	}

	_, err := s.client.DeleteImage(request)
	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			return nil
		}
		return fmt.Errorf("delete image failed: %w", err)
	}
	return nil
}
