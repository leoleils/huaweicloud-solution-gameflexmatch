// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package huawei 华为云ScalingService实现
package huawei

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"scase.io/cloudprovider"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	as "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/as/v1"
	asmodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/as/v1/model"
)

const (
	waitGroupStableTimes           = 60
	eachWaitDurationForGroupStable = 10 * time.Second
	getAsScalingInstanceLimit      = 100
	batchRemoveAsInstancesLimit    = 50
	asBatchRemoveFiledErrorCode    = "AS.4030"
	asInstanceNotExistErrorCode    = "AS.4006"
)

// HuaweiScalingService 华为云弹性伸缩服务实现
type HuaweiScalingService struct {
	provider *HuaweiProvider
	client   *as.AsClient
}

// CreateScalingConfig 创建伸缩配置
func (s *HuaweiScalingService) CreateScalingConfig(ctx context.Context, req *cloudprovider.CreateScalingConfigRequest) (string, error) {
	flavorIds := strings.Join(req.FlavorIds, ",")

	var securityGroups []asmodel.SecurityGroups
	for _, sg := range req.SecurityGroups {
		securityGroups = append(securityGroups, asmodel.SecurityGroups{Id: sg})
	}

	name := req.Name
	imageId := req.ImageId
	keyName := req.KeyName
	userData := req.UserData

	request := &asmodel.CreateScalingConfigRequest{
		Body: &asmodel.CreateScalingConfigOption{
			ScalingConfigurationName: &name,
			InstanceConfig: &asmodel.InstanceConfig{
				FlavorRef:      &flavorIds,
				ImageRef:       &imageId,
				Disk:           s.buildDisks(req.Disks),
				PublicIp:       s.buildPublicIp(req.Eip),
				KeyName:        &keyName,
				UserData:       &userData,
				SecurityGroups: &securityGroups,
			},
		},
	}

	resp, err := s.client.CreateScalingConfig(request)
	if err != nil {
		return "", fmt.Errorf("create scaling config failed: %w", err)
	}
	return *resp.ScalingConfigurationId, nil
}

// DeleteScalingConfig 删除伸缩配置
func (s *HuaweiScalingService) DeleteScalingConfig(ctx context.Context, configId string) error {
	request := &asmodel.DeleteScalingConfigRequest{
		ScalingConfigurationId: configId,
	}
	_, err := s.client.DeleteScalingConfig(request)
	if err != nil {
		respErr, ok := err.(*sdkerr.ServiceResponseError)
		if ok && respErr.StatusCode == http.StatusNotFound {
			return nil
		}
		return fmt.Errorf("delete scaling config failed: %w", err)
	}
	return nil
}

// CreateScalingGroup 创建伸缩组
func (s *HuaweiScalingService) CreateScalingGroup(ctx context.Context, req *cloudprovider.CreateScalingGroupRequest) (string, error) {
	deleteEIP := req.DeletePublicIp
	deleteVol := req.DeleteVolume

	request := &asmodel.CreateScalingGroupRequest{
		Body: &asmodel.CreateScalingGroupOption{
			ScalingGroupName:       req.Name,
			ScalingConfigurationId: req.ConfigId,
			MinInstanceNumber:      &req.MinInstanceNumber,
			MaxInstanceNumber:      &req.MaxInstanceNumber,
			DesireInstanceNumber:   &req.DesireInstanceNumber,
			VpcId:                  req.VpcId,
			Networks: []asmodel.Networks{{
				Id: req.SubnetId,
			}},
			DeletePublicip:      &deleteEIP,
			DeleteVolume:        &deleteVol,
			EnterpriseProjectId: &req.EnterpriseProjectId,
			IamAgencyName:       &req.IamAgencyName,
		},
	}

	resp, err := s.client.CreateScalingGroup(request)
	if err != nil {
		return "", fmt.Errorf("create scaling group failed: %w", err)
	}
	return *resp.ScalingGroupId, nil
}

// DeleteScalingGroup 删除伸缩组
func (s *HuaweiScalingService) DeleteScalingGroup(ctx context.Context, groupId string, forceDelete bool) error {
	// 先暂停伸缩组
	pauseReq := &asmodel.PauseScalingGroupRequest{
		ScalingGroupId: groupId,
		Body: &asmodel.PauseScalingGroupOption{
			Action: asmodel.GetPauseScalingGroupOptionActionEnum().PAUSE,
		},
	}
	_, err := s.client.PauseScalingGroup(pauseReq)
	if err != nil {
		respErr, ok := err.(*sdkerr.ServiceResponseError)
		if ok && respErr.StatusCode == http.StatusNotFound {
			return nil
		}
		return fmt.Errorf("pause scaling group failed: %w", err)
	}

	// 删除伸缩组
	forceDel := asmodel.GetDeleteScalingGroupRequestForceDeleteEnum().NO
	if forceDelete {
		forceDel = asmodel.GetDeleteScalingGroupRequestForceDeleteEnum().YES
	}

	delReq := &asmodel.DeleteScalingGroupRequest{
		ScalingGroupId: groupId,
		ForceDelete:    &forceDel,
	}
	_, err = s.client.DeleteScalingGroup(delReq)
	if err != nil {
		return fmt.Errorf("delete scaling group failed: %w", err)
	}
	return nil
}

// EnableScalingGroup 启用伸缩组
func (s *HuaweiScalingService) EnableScalingGroup(ctx context.Context, groupId string) error {
	request := &asmodel.ResumeScalingGroupRequest{
		ScalingGroupId: groupId,
		Body: &asmodel.ResumeScalingGroupOption{
			Action: asmodel.GetResumeScalingGroupOptionActionEnum().RESUME,
		},
	}
	_, err := s.client.ResumeScalingGroup(request)
	if err != nil {
		return fmt.Errorf("enable scaling group failed: %w", err)
	}
	return nil
}

// DisableScalingGroup 禁用伸缩组
func (s *HuaweiScalingService) DisableScalingGroup(ctx context.Context, groupId string) error {
	request := &asmodel.PauseScalingGroupRequest{
		ScalingGroupId: groupId,
		Body: &asmodel.PauseScalingGroupOption{
			Action: asmodel.GetPauseScalingGroupOptionActionEnum().PAUSE,
		},
	}
	_, err := s.client.PauseScalingGroup(request)
	if err != nil {
		return fmt.Errorf("disable scaling group failed: %w", err)
	}
	return nil
}

// GetScalingGroup 获取伸缩组信息
func (s *HuaweiScalingService) GetScalingGroup(ctx context.Context, groupId string) (*cloudprovider.ScalingGroup, error) {
	request := &asmodel.ShowScalingGroupRequest{ScalingGroupId: groupId}
	resp, err := s.client.ShowScalingGroup(request)
	if err != nil {
		return nil, fmt.Errorf("get scaling group failed: %w", err)
	}

	return &cloudprovider.ScalingGroup{
		Id:                    *resp.ScalingGroup.ScalingGroupId,
		Name:                  *resp.ScalingGroup.ScalingGroupName,
		Status:                resp.ScalingGroup.ScalingGroupStatus.Value(),
		CurrentInstanceNumber: *resp.ScalingGroup.CurrentInstanceNumber,
		DesireInstanceNumber:  *resp.ScalingGroup.DesireInstanceNumber,
		MinInstanceNumber:     *resp.ScalingGroup.MinInstanceNumber,
		MaxInstanceNumber:     *resp.ScalingGroup.MaxInstanceNumber,
		ConfigId:              *resp.ScalingGroup.ScalingConfigurationId,
		VpcId:                 *resp.ScalingGroup.VpcId,
	}, nil
}

// UpdateScalingGroup 更新伸缩组
func (s *HuaweiScalingService) UpdateScalingGroup(ctx context.Context, groupId string, req *cloudprovider.UpdateScalingGroupRequest) error {
	request := &asmodel.UpdateScalingGroupRequest{
		ScalingGroupId: groupId,
		Body:           &asmodel.UpdateScalingGroupOption{},
	}

	if req.DesireInstanceNumber != nil {
		request.Body.DesireInstanceNumber = req.DesireInstanceNumber
	}
	if req.MinInstanceNumber != nil {
		request.Body.MinInstanceNumber = req.MinInstanceNumber
	}
	if req.MaxInstanceNumber != nil {
		request.Body.MaxInstanceNumber = req.MaxInstanceNumber
	}

	_, err := s.client.UpdateScalingGroup(request)
	if err != nil {
		return fmt.Errorf("update scaling group failed: %w", err)
	}
	return nil
}

// WaitScalingGroupStable 等待伸缩组稳定
func (s *HuaweiScalingService) WaitScalingGroupStable(ctx context.Context, groupId string) error {
	for i := 0; i < waitGroupStableTimes; i++ {
		resp, err := s.client.ShowScalingGroup(&asmodel.ShowScalingGroupRequest{ScalingGroupId: groupId})
		if err != nil {
			return fmt.Errorf("show scaling group failed: %w", err)
		}

		curNum := *resp.ScalingGroup.CurrentInstanceNumber
		desireNum := *resp.ScalingGroup.DesireInstanceNumber
		if curNum == desireNum {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(eachWaitDurationForGroupStable):
		}
	}
	return fmt.Errorf("scaling group not stable after %d attempts", waitGroupStableTimes)
}

// ListScalingInstances 列出伸缩组实例
func (s *HuaweiScalingService) ListScalingInstances(ctx context.Context, groupId string) ([]string, error) {
	// 等待伸缩组稳定
	if err := s.WaitScalingGroupStable(ctx, groupId); err != nil {
		return nil, err
	}

	var allInstanceIds []string
	var startNumber int32 = 0
	limit := int32(getAsScalingInstanceLimit)

	for {
		resp, err := s.client.ListScalingInstances(&asmodel.ListScalingInstancesRequest{
			ScalingGroupId: groupId,
			StartNumber:    &startNumber,
			Limit:          &limit,
		})
		if err != nil {
			return nil, fmt.Errorf("list scaling instances failed: %w", err)
		}

		if resp.ScalingGroupInstances == nil || len(*resp.ScalingGroupInstances) == 0 {
			break
		}

		for _, instance := range *resp.ScalingGroupInstances {
			if instance.InstanceId != nil {
				allInstanceIds = append(allInstanceIds, *instance.InstanceId)
			}
		}

		if int32(len(*resp.ScalingGroupInstances)) < limit {
			break
		}
		startNumber += limit
	}

	return allInstanceIds, nil
}

// BatchRemoveScalingInstances 批量移除伸缩组实例
func (s *HuaweiScalingService) BatchRemoveScalingInstances(ctx context.Context, groupId string, instanceIds []string, deleteInstance bool) error {
	if len(instanceIds) == 0 {
		return nil
	}

	// 分批处理
	for i := 0; i < len(instanceIds); i += batchRemoveAsInstancesLimit {
		end := i + batchRemoveAsInstancesLimit
		if end > len(instanceIds) {
			end = len(instanceIds)
		}

		batch := instanceIds[i:end]
		if err := s.batchRemoveInstances(ctx, groupId, batch, deleteInstance); err != nil {
			return err
		}
	}
	return nil
}

func (s *HuaweiScalingService) batchRemoveInstances(ctx context.Context, groupId string, instanceIds []string, deleteInstance bool) error {
	// 等待伸缩组稳定
	if err := s.WaitScalingGroupStable(ctx, groupId); err != nil {
		return err
	}

	instanceDelete := asmodel.GetBatchRemoveInstancesOptionInstanceDeleteEnum().NO
	if deleteInstance {
		instanceDelete = asmodel.GetBatchRemoveInstancesOptionInstanceDeleteEnum().YES
	}

	_, err := s.client.BatchRemoveScalingInstances(&asmodel.BatchRemoveScalingInstancesRequest{
		ScalingGroupId: groupId,
		Body: &asmodel.BatchRemoveInstancesOption{
			InstancesId:    instanceIds,
			InstanceDelete: &instanceDelete,
			Action:         asmodel.GetBatchRemoveInstancesOptionActionEnum().REMOVE,
		},
	})
	if err != nil {
		respErr, ok := err.(*sdkerr.ServiceResponseError)
		if ok && (respErr.StatusCode == http.StatusNotFound ||
			(respErr.ErrorCode == asBatchRemoveFiledErrorCode &&
				strings.Contains(respErr.ErrorMessage, asInstanceNotExistErrorCode))) {
			return nil
		}
		return fmt.Errorf("batch remove scaling instances failed: %w", err)
	}
	return nil
}

// UpdateScalingGroupTags 更新伸缩组标签
func (s *HuaweiScalingService) UpdateScalingGroupTags(ctx context.Context, groupId string, tags []cloudprovider.Tag) error {
	// 先获取现有标签
	listReq := &asmodel.ListScalingTagInfosByResourceIdRequest{
		ResourceType: asmodel.GetListScalingTagInfosByResourceIdRequestResourceTypeEnum().SCALING_GROUP_TAG,
		ResourceId:   groupId,
	}
	listResp, err := s.client.ListScalingTagInfosByResourceId(listReq)
	if err != nil {
		return fmt.Errorf("list scaling group tags failed: %w", err)
	}

	// 计算需要创建和删除的标签
	existingTags := make(map[string]string)
	if listResp.Tags != nil {
		for _, tag := range *listResp.Tags {
			existingTags[tag.Key] = *tag.Value
		}
	}

	newTags := make(map[string]string)
	for _, tag := range tags {
		newTags[tag.Key] = tag.Value
	}

	// 删除不再需要的标签
	var deleteTags []asmodel.TagsSingleValue
	for key, value := range existingTags {
		if _, ok := newTags[key]; !ok {
			deleteTags = append(deleteTags, asmodel.TagsSingleValue{Key: key, Value: &value})
		}
	}
	if len(deleteTags) > 0 {
		if err := s.DeleteScalingGroupTags(ctx, groupId, s.convertToCloudproviderTags(deleteTags)); err != nil {
			return err
		}
	}

	// 创建/更新标签
	if len(tags) > 0 {
		if err := s.CreateScalingGroupTags(ctx, groupId, tags); err != nil {
			return err
		}
	}

	return nil
}

// CreateScalingGroupTags 创建伸缩组标签
func (s *HuaweiScalingService) CreateScalingGroupTags(ctx context.Context, groupId string, tags []cloudprovider.Tag) error {
	var asTags []asmodel.TagsSingleValue
	for _, tag := range tags {
		value := tag.Value
		asTags = append(asTags, asmodel.TagsSingleValue{Key: tag.Key, Value: &value})
	}

	request := &asmodel.CreateScalingTagInfoRequest{
		ResourceType: asmodel.GetCreateScalingTagInfoRequestResourceTypeEnum().SCALING_GROUP_TAG,
		ResourceId:   groupId,
		Body: &asmodel.CreateTagsOption{
			Tags:   asTags,
			Action: asmodel.GetCreateTagsOptionActionEnum().CREATE,
		},
	}
	_, err := s.client.CreateScalingTagInfo(request)
	if err != nil {
		return fmt.Errorf("create scaling group tags failed: %w", err)
	}
	return nil
}

// DeleteScalingGroupTags 删除伸缩组标签
func (s *HuaweiScalingService) DeleteScalingGroupTags(ctx context.Context, groupId string, tags []cloudprovider.Tag) error {
	var asTags []asmodel.TagsSingleValue
	for _, tag := range tags {
		value := tag.Value
		asTags = append(asTags, asmodel.TagsSingleValue{Key: tag.Key, Value: &value})
	}

	request := &asmodel.DeleteScalingTagInfoRequest{
		ResourceType: asmodel.GetDeleteScalingTagInfoRequestResourceTypeEnum().SCALING_GROUP_TAG,
		ResourceId:   groupId,
		Body: &asmodel.DeleteTagsOption{
			Tags:   asTags,
			Action: asmodel.GetDeleteTagsOptionActionEnum().DELETE,
		},
	}
	_, err := s.client.DeleteScalingTagInfo(request)
	if err != nil {
		return fmt.Errorf("delete scaling group tags failed: %w", err)
	}
	return nil
}

// buildDisks 构建磁盘配置
func (s *HuaweiScalingService) buildDisks(disks []cloudprovider.Disk) *[]asmodel.DiskInfo {
	var diskInfos []asmodel.DiskInfo
	for _, d := range disks {
		diskInfos = append(diskInfos, asmodel.DiskInfo{
			Size:       d.Size,
			VolumeType: s.getVolumeType(d.VolumeType),
			DiskType:   s.getDiskType(d.DiskType),
		})
	}
	return &diskInfos
}

// buildPublicIp 构建公网IP配置
func (s *HuaweiScalingService) buildPublicIp(eip *cloudprovider.EipConfig) *asmodel.PublicIp {
	if eip == nil {
		return nil
	}

	shareType := asmodel.GetBandwidthInfoShareTypeEnum().PER
	if eip.ShareType == cloudprovider.ShareTypeWhole {
		shareType = asmodel.GetBandwidthInfoShareTypeEnum().WHOLE
	}

	chargeMode := asmodel.GetBandwidthInfoChargingModeEnum().TRAFFIC
	if eip.ChargeMode == cloudprovider.ChargeModeBandwidth {
		chargeMode = asmodel.GetBandwidthInfoChargingModeEnum().BANDWIDTH
	}

	return &asmodel.PublicIp{
		Eip: &asmodel.EipInfo{
			IpType: asmodel.GetEipInfoIpTypeEnum().E_5_BGP,
			Bandwidth: &asmodel.BandwidthInfo{
				Size:         &eip.BandwidthSize,
				ShareType:    shareType,
				ChargingMode: &chargeMode,
				Id:           &eip.BandwidthId,
			},
		},
	}
}

// getVolumeType 获取卷类型
func (s *HuaweiScalingService) getVolumeType(volumeType string) asmodel.DiskInfoVolumeType {
	volumeTypeEnum := asmodel.GetDiskInfoVolumeTypeEnum()
	switch volumeType {
	case cloudprovider.VolumeTypeSATA:
		return volumeTypeEnum.SATA
	case cloudprovider.VolumeTypeSAS:
		return volumeTypeEnum.SAS
	case cloudprovider.VolumeTypeSSD:
		return volumeTypeEnum.SSD
	case cloudprovider.VolumeTypeGPSSD:
		return volumeTypeEnum.GPSSD
	default:
		return volumeTypeEnum.GPSSD
	}
}

// getDiskType 获取磁盘类型
func (s *HuaweiScalingService) getDiskType(diskType string) asmodel.DiskInfoDiskType {
	diskTypeEnum := asmodel.GetDiskInfoDiskTypeEnum()
	if diskType == cloudprovider.DiskTypeData {
		return diskTypeEnum.DATA
	}
	return diskTypeEnum.SYS
}

// convertToCloudproviderTags 转换标签
func (s *HuaweiScalingService) convertToCloudproviderTags(tags []asmodel.TagsSingleValue) []cloudprovider.Tag {
	var result []cloudprovider.Tag
	for _, tag := range tags {
		value := ""
		if tag.Value != nil {
			value = *tag.Value
		}
		result = append(result, cloudprovider.Tag{Key: tag.Key, Value: value})
	}
	return result
}
