// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package alibaba 阿里云ScalingService实现
package alibaba

import (
	"context"
	"fmt"
	"strings"
	"time"

	"scase.io/cloudprovider"

	ess "github.com/alibabacloud-go/ess-20220222/client"
	"github.com/alibabacloud-go/tea/tea"
)

const (
	waitScalingGroupStableTimes    = 60
	waitScalingGroupStableInterval = 10 * time.Second
)

// AlibabaScalingService 阿里云弹性伸缩服务实现
type AlibabaScalingService struct {
	provider *AlibabaProvider
	client   *ess.Client
}

// CreateScalingConfig 创建伸缩配置
func (s *AlibabaScalingService) CreateScalingConfig(ctx context.Context, req *cloudprovider.CreateScalingConfigRequest) (string, error) {
	request := &ess.CreateScalingConfigurationRequest{
		ScalingConfigurationName: tea.String(req.Name),
		ImageId:                  tea.String(req.ImageId),
		InstanceTypes:            tea.StringSlice(req.FlavorIds),
		SecurityGroupIds:         tea.StringSlice(req.SecurityGroups),
		KeyPairName:              tea.String(req.KeyName),
		UserData:                 tea.String(req.UserData),
	}

	resp, err := s.client.CreateScalingConfiguration(request)
	if err != nil {
		return "", fmt.Errorf("create scaling configuration failed: %w", err)
	}
	return *resp.Body.ScalingConfigurationId, nil
}

// DeleteScalingConfig 删除伸缩配置
func (s *AlibabaScalingService) DeleteScalingConfig(ctx context.Context, configId string) error {
	request := &ess.DeleteScalingConfigurationRequest{
		ScalingConfigurationId: tea.String(configId),
	}
	_, err := s.client.DeleteScalingConfiguration(request)
	if err != nil {
		// 忽略不存在的错误
		if strings.Contains(err.Error(), "NotFound") {
			return nil
		}
		return fmt.Errorf("delete scaling configuration failed: %w", err)
	}
	return nil
}

// CreateScalingGroup 创建伸缩组
func (s *AlibabaScalingService) CreateScalingGroup(ctx context.Context, req *cloudprovider.CreateScalingGroupRequest) (string, error) {
	request := &ess.CreateScalingGroupRequest{
		ScalingGroupName: tea.String(req.Name),
		VSwitchId:        tea.String(req.SubnetId),
		MinSize:          tea.Int32(req.MinInstanceNumber),
		MaxSize:          tea.Int32(req.MaxInstanceNumber),
		DesiredCapacity:  tea.Int32(req.DesireInstanceNumber),
	}

	resp, err := s.client.CreateScalingGroup(request)
	if err != nil {
		return "", fmt.Errorf("create scaling group failed: %w", err)
	}
	return *resp.Body.ScalingGroupId, nil
}

// DeleteScalingGroup 删除伸缩组
func (s *AlibabaScalingService) DeleteScalingGroup(ctx context.Context, groupId string, forceDelete bool) error {
	// 先禁用伸缩组
	if err := s.DisableScalingGroup(ctx, groupId); err != nil {
		// 忽略不存在的错误
		if !strings.Contains(err.Error(), "NotFound") {
			return err
		}
	}

	request := &ess.DeleteScalingGroupRequest{
		ScalingGroupId: tea.String(groupId),
		ForceDelete:    tea.Bool(forceDelete),
	}
	_, err := s.client.DeleteScalingGroup(request)
	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			return nil
		}
		return fmt.Errorf("delete scaling group failed: %w", err)
	}
	return nil
}

// EnableScalingGroup 启用伸缩组
func (s *AlibabaScalingService) EnableScalingGroup(ctx context.Context, groupId string) error {
	request := &ess.EnableScalingGroupRequest{
		ScalingGroupId: tea.String(groupId),
	}
	_, err := s.client.EnableScalingGroup(request)
	if err != nil {
		return fmt.Errorf("enable scaling group failed: %w", err)
	}
	return nil
}

// DisableScalingGroup 禁用伸缩组
func (s *AlibabaScalingService) DisableScalingGroup(ctx context.Context, groupId string) error {
	request := &ess.DisableScalingGroupRequest{
		ScalingGroupId: tea.String(groupId),
	}
	_, err := s.client.DisableScalingGroup(request)
	if err != nil {
		return fmt.Errorf("disable scaling group failed: %w", err)
	}
	return nil
}

// GetScalingGroup 获取伸缩组信息
func (s *AlibabaScalingService) GetScalingGroup(ctx context.Context, groupId string) (*cloudprovider.ScalingGroup, error) {
	request := &ess.DescribeScalingGroupsRequest{
		ScalingGroupIds: tea.StringSlice([]string{groupId}),
	}

	resp, err := s.client.DescribeScalingGroups(request)
	if err != nil {
		return nil, fmt.Errorf("get scaling group failed: %w", err)
	}

	if resp.Body.ScalingGroups == nil || len(resp.Body.ScalingGroups) == 0 {
		return nil, fmt.Errorf("scaling group not found: %s", groupId)
	}

	sg := resp.Body.ScalingGroups[0]
	return &cloudprovider.ScalingGroup{
		Id:                    *sg.ScalingGroupId,
		Name:                  *sg.ScalingGroupName,
		Status:                *sg.LifecycleState,
		CurrentInstanceNumber: int32(*sg.TotalCapacity),
		DesireInstanceNumber:  int32(*sg.DesiredCapacity),
		MinInstanceNumber:     *sg.MinSize,
		MaxInstanceNumber:     *sg.MaxSize,
		ConfigId:              *sg.ActiveScalingConfigurationId,
		VpcId:                 *sg.VpcId,
		SubnetId:              *sg.VSwitchId,
	}, nil
}

// UpdateScalingGroup 更新伸缩组
func (s *AlibabaScalingService) UpdateScalingGroup(ctx context.Context, groupId string, req *cloudprovider.UpdateScalingGroupRequest) error {
	request := &ess.ModifyScalingGroupRequest{
		ScalingGroupId: tea.String(groupId),
	}

	if req.DesireInstanceNumber != nil {
		request.DesiredCapacity = tea.Int32(*req.DesireInstanceNumber)
	}
	if req.MinInstanceNumber != nil {
		request.MinSize = req.MinInstanceNumber
	}
	if req.MaxInstanceNumber != nil {
		request.MaxSize = req.MaxInstanceNumber
	}

	_, err := s.client.ModifyScalingGroup(request)
	if err != nil {
		return fmt.Errorf("update scaling group failed: %w", err)
	}
	return nil
}

// WaitScalingGroupStable 等待伸缩组稳定
func (s *AlibabaScalingService) WaitScalingGroupStable(ctx context.Context, groupId string) error {
	for i := 0; i < waitScalingGroupStableTimes; i++ {
		sg, err := s.GetScalingGroup(ctx, groupId)
		if err != nil {
			return err
		}

		if sg.CurrentInstanceNumber == sg.DesireInstanceNumber {
			return nil
		}
		time.Sleep(waitScalingGroupStableInterval)
	}
	return nil
}

// ListScalingInstances 列出伸缩组实例
func (s *AlibabaScalingService) ListScalingInstances(ctx context.Context, groupId string) ([]string, error) {
	request := &ess.DescribeScalingInstancesRequest{
		ScalingGroupId: tea.String(groupId),
	}

	resp, err := s.client.DescribeScalingInstances(request)
	if err != nil {
		return nil, fmt.Errorf("list scaling instances failed: %w", err)
	}

	var instanceIds []string
	if resp.Body.ScalingInstances != nil {
		for _, inst := range resp.Body.ScalingInstances {
			if inst.InstanceId != nil {
				instanceIds = append(instanceIds, *inst.InstanceId)
			}
		}
	}
	return instanceIds, nil
}

// BatchRemoveScalingInstances 批量移除伸缩组实例
func (s *AlibabaScalingService) BatchRemoveScalingInstances(ctx context.Context, groupId string, instanceIds []string, deleteInstance bool) error {
	request := &ess.RemoveInstancesRequest{
		ScalingGroupId:          tea.String(groupId),
		InstanceIds:             tea.StringSlice(instanceIds),
		DecreaseDesiredCapacity: tea.Bool(true),
	}

	_, err := s.client.RemoveInstances(request)
	if err != nil {
		return fmt.Errorf("remove scaling instances failed: %w", err)
	}
	return nil
}

// UpdateScalingGroupTags 更新伸缩组标签
func (s *AlibabaScalingService) UpdateScalingGroupTags(ctx context.Context, groupId string, tags []cloudprovider.Tag) error {
	// 阿里云ESS标签API需要使用TagResources
	return s.CreateScalingGroupTags(ctx, groupId, tags)
}

// CreateScalingGroupTags 创建伸缩组标签
func (s *AlibabaScalingService) CreateScalingGroupTags(ctx context.Context, groupId string, tags []cloudprovider.Tag) error {
	// TODO: 实现阿里云ESS标签API
	return nil
}

// DeleteScalingGroupTags 删除伸缩组标签
func (s *AlibabaScalingService) DeleteScalingGroupTags(ctx context.Context, groupId string, tags []cloudprovider.Tag) error {
	// TODO: 实现阿里云ESS标签API
	return nil
}

// convertVolumeType 转换卷类型
func (s *AlibabaScalingService) convertVolumeType(volumeType string) string {
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
