// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package alibaba 阿里云ComputeService实现
package alibaba

import (
	"context"
	"fmt"
	"strings"

	"scase.io/cloudprovider"

	ecs "github.com/alibabacloud-go/ecs-20140526/v7/client"
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
	// 阿里云 DescribeInstanceTypes 需要指定 RegionId 来获取当前地域可用的实例规格
	request := &ecs.DescribeAvailableResourceRequest{
		RegionId:            tea.String(s.provider.config.Region),
		DestinationResource: tea.String("InstanceType"),
		InstanceChargeType:  tea.String("PostPaid"),
	}

	resp, err := s.client.DescribeAvailableResource(request)
	if err != nil {
		return nil, fmt.Errorf("list flavors failed: %w", err)
	}

	result := make(map[string][]string)
	if resp.Body.AvailableZones != nil {
		for _, zone := range resp.Body.AvailableZones.AvailableZone {
			if zone.AvailableResources == nil {
				continue
			}
			for _, resource := range zone.AvailableResources.AvailableResource {
				if resource.SupportedResources == nil {
					continue
				}
				for _, supported := range resource.SupportedResources.SupportedResource {
					if supported.Status != nil && *supported.Status == "Available" {
						instanceTypeId := tea.StringValue(supported.Value)
						if instanceTypeId != "" {
							// 按规格系列分组 (如 ecs.c6 -> c6 系列)
							series := s.extractInstanceSeries(instanceTypeId)
							result[series] = append(result[series], instanceTypeId)
						}
					}
				}
			}
		}
	}

	// 如果结果为空，尝试使用 DescribeInstanceTypes 获取所有规格
	if len(result) == 0 {
		request2 := &ecs.DescribeInstanceTypesRequest{}
		resp2, err := s.client.DescribeInstanceTypes(request2)
		if err != nil {
			return nil, fmt.Errorf("list instance types failed: %w", err)
		}
		if resp2.Body.InstanceTypes != nil {
			for _, flavor := range resp2.Body.InstanceTypes.InstanceType {
				instanceTypeId := tea.StringValue(flavor.InstanceTypeId)
				series := s.extractInstanceSeries(instanceTypeId)
				result[series] = append(result[series], instanceTypeId)
			}
		}
	}

	return result, nil
}

// extractInstanceSeries 提取实例规格系列名称
// 例如: ecs.c6.large -> ecs.c6, ecs.g7.xlarge -> ecs.g7
func (s *AlibabaComputeService) extractInstanceSeries(instanceTypeId string) string {
	parts := strings.Split(instanceTypeId, ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return instanceTypeId
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

// convertVolumeType 转换卷类型 - 将华为云磁盘类型转换为阿里云磁盘类型
func (s *AlibabaComputeService) convertVolumeType(volumeType string) string {
	switch volumeType {
	// 华为云 SSD -> 阿里云 cloud_ssd
	case cloudprovider.VolumeTypeSSD:
		return "cloud_ssd"
	// 华为云 ESSD -> 阿里云 cloud_essd
	case cloudprovider.VolumeTypeESSD:
		return "cloud_essd"
	// 华为云 GPSSD -> 阿里云 cloud_essd
	case cloudprovider.VolumeTypeGPSSD:
		return "cloud_essd"
	// 华为云 SATA/SAS -> 阿里云 cloud_essd
	case cloudprovider.VolumeTypeSATA, cloudprovider.VolumeTypeSAS:
		return "cloud_essd"
	// 阿里云原生类型直接返回
	case "cloud_efficiency", "cloud_ssd", "cloud_essd", "cloud_essd_entry", "cloud_auto":
		return volumeType
	default:
		// 默认使用ESSD云盘
		return "cloud_essd"
	}
}

// convertInstance 转换实例信息
func (s *AlibabaComputeService) convertInstance(inst *ecs.DescribeInstancesResponseBodyInstancesInstance) *cloudprovider.Instance {
	instance := &cloudprovider.Instance{
		Id:               *inst.InstanceId,
		Name:             *inst.InstanceName,
		Status:           s.convertInstanceStatus(*inst.Status),
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

// convertInstanceStatus 将阿里云ECS状态转换为标准状态
// 阿里云状态: Pending, Running, Starting, Stopping, Stopped
// 标准状态: RUNNING, STOPPED, SHUTOFF, DELETED
func (s *AlibabaComputeService) convertInstanceStatus(status string) string {
	switch strings.ToLower(status) {
	case "running":
		return cloudprovider.InstanceStatusRunning
	case "stopped":
		return cloudprovider.InstanceStatusStopped
	case "stopping":
		return "STOPPING"
	case "starting":
		return "STARTING"
	case "pending":
		return "PENDING"
	default:
		return strings.ToUpper(status)
	}
}
