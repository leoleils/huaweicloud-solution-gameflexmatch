// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package alibaba 阿里云ComputeService实现
package alibaba

import (
	"context"
	"fmt"
	"strings"

	"scase.io/cloudprovider"

	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

// AlibabaComputeService 阿里云计算服务实现
type AlibabaComputeService struct {
	provider *AlibabaProvider
	client   *ecs.Client
}

// CreateInstance 创建虚拟机实例
func (s *AlibabaComputeService) CreateInstance(ctx context.Context, req *cloudprovider.CreateInstanceRequest) (*cloudprovider.CreateInstanceResponse, error) {
	// 构建系统盘
	var systemDisk *ecs.RunInstancesRequestSystemDisk
	var dataDisks []*ecs.RunInstancesRequestDataDisk
	for _, disk := range req.Disks {
		if disk.DiskType == cloudprovider.DiskTypeSys {
			systemDisk = &ecs.RunInstancesRequestSystemDisk{
				Size:     tea.String(fmt.Sprintf("%d", disk.Size)),
				Category: tea.String(s.convertVolumeType(disk.VolumeType)),
			}
		} else {
			dataDisks = append(dataDisks, &ecs.RunInstancesRequestDataDisk{
				Size:     tea.Int32(disk.Size),
				Category: tea.String(s.convertVolumeType(disk.VolumeType)),
			})
		}
	}

	// 构建标签
	var tags []*ecs.RunInstancesRequestTag
	for _, tag := range req.Tags {
		tags = append(tags, &ecs.RunInstancesRequestTag{
			Key:   tea.String(tag.Key),
			Value: tea.String(tag.Value),
		})
	}

	request := &ecs.RunInstancesRequest{
		RegionId:        tea.String(s.provider.config.Region),
		ImageId:         tea.String(req.ImageId),
		InstanceType:    tea.String(req.FlavorId),
		VSwitchId:       tea.String(req.SubnetId),
		SecurityGroupId: tea.String(strings.Join(req.SecurityGroupIds, ",")),
		ZoneId:          tea.String(req.AvailabilityZone),
		KeyPairName:     tea.String(req.KeyName),
		UserData:        tea.String(req.UserData),
		Amount:          tea.Int32(int32(req.Count)),
		InstanceName:    tea.String(req.Name),
		SystemDisk:      systemDisk,
		DataDisk:        dataDisks,
		Tag:             tags,
	}

	resp, err := s.client.RunInstances(request)
	if err != nil {
		return nil, fmt.Errorf("create ecs instance failed: %w", err)
	}

	return &cloudprovider.CreateInstanceResponse{
		InstanceIds: tea.StringSliceValue(resp.Body.InstanceIdSets.InstanceIdSet),
	}, nil
}

// DeleteInstance 删除虚拟机实例
func (s *AlibabaComputeService) DeleteInstance(ctx context.Context, instanceId string) error {
	request := &ecs.DeleteInstanceRequest{
		InstanceId: tea.String(instanceId),
		Force:      tea.Bool(true),
	}
	_, err := s.client.DeleteInstance(request)
	if err != nil {
		return fmt.Errorf("delete ecs instance failed: %w", err)
	}
	return nil
}

// BatchDeleteInstances 批量删除虚拟机实例
func (s *AlibabaComputeService) BatchDeleteInstances(ctx context.Context, instanceIds []string) (string, error) {
	if len(instanceIds) == 0 {
		return "", nil
	}

	request := &ecs.DeleteInstancesRequest{
		RegionId:   tea.String(s.provider.config.Region),
		InstanceId: tea.StringSlice(instanceIds),
		Force:      tea.Bool(true),
	}

	_, err := s.client.DeleteInstances(request)
	if err != nil {
		return "", fmt.Errorf("batch delete ecs instances failed: %w", err)
	}
	return "", nil
}

// GetInstance 获取虚拟机实例信息
func (s *AlibabaComputeService) GetInstance(ctx context.Context, instanceId string) (*cloudprovider.Instance, error) {
	request := &ecs.DescribeInstancesRequest{
		RegionId:    tea.String(s.provider.config.Region),
		InstanceIds: tea.String(fmt.Sprintf("[\"%s\"]", instanceId)),
	}

	resp, err := s.client.DescribeInstances(request)
	if err != nil {
		return nil, fmt.Errorf("get ecs instance failed: %w", err)
	}

	if resp.Body.Instances == nil || len(resp.Body.Instances.Instance) == 0 {
		return nil, fmt.Errorf("instance not found: %s", instanceId)
	}

	return s.convertInstance(resp.Body.Instances.Instance[0]), nil
}

// ListInstances 列出虚拟机实例
func (s *AlibabaComputeService) ListInstances(ctx context.Context, req *cloudprovider.ListInstancesRequest) (*cloudprovider.ListInstancesResponse, error) {
	request := &ecs.DescribeInstancesRequest{
		RegionId: tea.String(s.provider.config.Region),
	}

	if req.Name != "" {
		request.InstanceName = tea.String(req.Name)
	}
	if req.Status != "" {
		request.Status = tea.String(req.Status)
	}
	if req.VpcId != "" {
		request.VpcId = tea.String(req.VpcId)
	}
	if req.Limit > 0 {
		request.PageSize = tea.Int32(int32(req.Limit))
	}
	if req.Offset > 0 {
		pageNumber := (req.Offset / req.Limit) + 1
		request.PageNumber = tea.Int32(int32(pageNumber))
	}

	resp, err := s.client.DescribeInstances(request)
	if err != nil {
		return nil, fmt.Errorf("list ecs instances failed: %w", err)
	}

	var instances []cloudprovider.Instance
	if resp.Body.Instances != nil {
		for _, inst := range resp.Body.Instances.Instance {
			instances = append(instances, *s.convertInstance(inst))
		}
	}

	return &cloudprovider.ListInstancesResponse{
		Instances:  instances,
		TotalCount: int(*resp.Body.TotalCount),
	}, nil
}

// StartInstance 启动虚拟机实例
func (s *AlibabaComputeService) StartInstance(ctx context.Context, instanceId string) error {
	request := &ecs.StartInstanceRequest{
		InstanceId: tea.String(instanceId),
	}
	_, err := s.client.StartInstance(request)
	if err != nil {
		return fmt.Errorf("start ecs instance failed: %w", err)
	}
	return nil
}

// StopInstance 停止虚拟机实例
func (s *AlibabaComputeService) StopInstance(ctx context.Context, instanceId string) error {
	request := &ecs.StopInstanceRequest{
		InstanceId: tea.String(instanceId),
		ForceStop:  tea.Bool(true),
	}
	_, err := s.client.StopInstance(request)
	if err != nil {
		return fmt.Errorf("stop ecs instance failed: %w", err)
	}
	return nil
}

// BatchStopInstances 批量停止虚拟机实例
func (s *AlibabaComputeService) BatchStopInstances(ctx context.Context, instanceIds []string) error {
	if len(instanceIds) == 0 {
		return nil
	}

	request := &ecs.StopInstancesRequest{
		RegionId:   tea.String(s.provider.config.Region),
		InstanceId: tea.StringSlice(instanceIds),
		ForceStop:  tea.Bool(true),
	}

	_, err := s.client.StopInstances(request)
	if err != nil {
		return fmt.Errorf("batch stop ecs instances failed: %w", err)
	}
	return nil
}

// ListAvailabilityZones 列出可用区
func (s *AlibabaComputeService) ListAvailabilityZones(ctx context.Context) ([]string, error) {
	request := &ecs.DescribeZonesRequest{
		RegionId: tea.String(s.provider.config.Region),
	}

	resp, err := s.client.DescribeZones(request)
	if err != nil {
		return nil, fmt.Errorf("list availability zones failed: %w", err)
	}

	var zones []string
	if resp.Body.Zones != nil {
		for _, zone := range resp.Body.Zones.Zone {
			zones = append(zones, *zone.ZoneId)
		}
	}
	return zones, nil
}

// ListFlavors 列出规格
func (s *AlibabaComputeService) ListFlavors(ctx context.Context) (map[string][]string, error) {
	request := &ecs.DescribeInstanceTypesRequest{}

	resp, err := s.client.DescribeInstanceTypes(request)
	if err != nil {
		return nil, fmt.Errorf("list flavors failed: %w", err)
	}

	result := make(map[string][]string)
	if resp.Body.InstanceTypes != nil {
		for _, flavor := range resp.Body.InstanceTypes.InstanceType {
			result[*flavor.InstanceTypeId] = []string{*flavor.InstanceTypeId}
		}
	}
	return result, nil
}

// GetJob 获取异步任务状态 (阿里云通常是同步接口)
func (s *AlibabaComputeService) GetJob(ctx context.Context, jobId string) (*cloudprovider.JobInfo, error) {
	// 阿里云ECS大部分接口是同步的，不需要查询Job
	return &cloudprovider.JobInfo{
		Id:     jobId,
		Status: cloudprovider.JobStatusSuccess,
	}, nil
}

// UpdateInstanceTags 更新实例标签
func (s *AlibabaComputeService) UpdateInstanceTags(ctx context.Context, instanceId string, tags []cloudprovider.Tag) error {
	var ecsTags []*ecs.AddTagsRequestTag
	for _, tag := range tags {
		ecsTags = append(ecsTags, &ecs.AddTagsRequestTag{
			Key:   tea.String(tag.Key),
			Value: tea.String(tag.Value),
		})
	}

	request := &ecs.AddTagsRequest{
		RegionId:     tea.String(s.provider.config.Region),
		ResourceType: tea.String("instance"),
		ResourceId:   tea.String(instanceId),
		Tag:          ecsTags,
	}

	_, err := s.client.AddTags(request)
	if err != nil {
		return fmt.Errorf("update instance tags failed: %w", err)
	}
	return nil
}

// ParseInstanceIPs 解析实例IP地址
func (s *AlibabaComputeService) ParseInstanceIPs(ctx context.Context, instance *cloudprovider.Instance, ipType string) ([]string, error) {
	if ipType == cloudprovider.IPTypePublic {
		return instance.PublicIPs, nil
	}
	return instance.PrivateIPs, nil
}

// convertVolumeType 转换卷类型
func (s *AlibabaComputeService) convertVolumeType(volumeType string) string {
	switch volumeType {
	case cloudprovider.VolumeTypeSSD:
		return "cloud_ssd"
	case cloudprovider.VolumeTypeESSD:
		return "cloud_essd"
	case cloudprovider.VolumeTypeGPSSD:
		return "cloud_essd"
	default:
		return "cloud_efficiency"
	}
}

// convertInstance 转换实例信息
func (s *AlibabaComputeService) convertInstance(inst *ecs.DescribeInstancesResponseBodyInstancesInstance) *cloudprovider.Instance {
	instance := &cloudprovider.Instance{
		Id:               *inst.InstanceId,
		Name:             *inst.InstanceName,
		Status:           *inst.Status,
		FlavorId:         *inst.InstanceType,
		ImageId:          *inst.ImageId,
		AvailabilityZone: *inst.ZoneId,
		RawData:          inst,
	}

	// 解析私网IP
	if inst.VpcAttributes != nil && inst.VpcAttributes.PrivateIpAddress != nil {
		instance.PrivateIPs = tea.StringSliceValue(inst.VpcAttributes.PrivateIpAddress.IpAddress)
		instance.VpcId = tea.StringValue(inst.VpcAttributes.VpcId)
		instance.SubnetId = tea.StringValue(inst.VpcAttributes.VSwitchId)
	}

	// 解析公网IP
	if inst.PublicIpAddress != nil {
		instance.PublicIPs = tea.StringSliceValue(inst.PublicIpAddress.IpAddress)
	}

	// 解析安全组
	if inst.SecurityGroupIds != nil {
		instance.SecurityGroupIds = tea.StringSliceValue(inst.SecurityGroupIds.SecurityGroupId)
	}

	return instance
}
