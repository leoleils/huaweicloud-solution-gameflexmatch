// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package huawei 华为云ComputeService实现
package huawei

import (
	"context"
	"fmt"
	"net/http"

	"scase.io/cloudprovider"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecsmodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
)

// HuaweiComputeService 华为云计算服务实现
type HuaweiComputeService struct {
	provider *HuaweiProvider
	client   *ecs.EcsClient
}

// CreateInstance 创建虚拟机实例
func (s *HuaweiComputeService) CreateInstance(ctx context.Context, req *cloudprovider.CreateInstanceRequest) (*cloudprovider.CreateInstanceResponse, error) {
	// 构建标签
	var tags []ecsmodel.PrePaidServerTag
	for _, tag := range req.Tags {
		tags = append(tags, ecsmodel.PrePaidServerTag{
			Key:   tag.Key,
			Value: tag.Value,
		})
	}

	// 构建磁盘
	rootVolume, dataVolumes := s.buildVolumes(req.Disks)

	// 构建网卡
	nics := []ecsmodel.PrePaidServerNic{{SubnetId: req.SubnetId}}

	// 构建安全组
	var securityGroups []ecsmodel.PrePaidServerSecurityGroup
	for _, sgId := range req.SecurityGroupIds {
		securityGroups = append(securityGroups, ecsmodel.PrePaidServerSecurityGroup{Id: &sgId})
	}

	// 构建EIP
	var publicip *ecsmodel.PrePaidServerPublicip
	if req.Eip != nil {
		publicip = s.buildPublicIp(req.Eip)
	}

	// 构建元数据
	metadata := make(map[string]string)
	if req.AgencyName != "" {
		metadata["agency_name"] = req.AgencyName
	}

	ecsCount := int32(req.Count)
	isAutoRename := false
	batchCreateInMultiAz := true

	serverBody := &ecsmodel.PrePaidServer{
		ImageRef:             req.ImageId,
		FlavorRef:            req.FlavorId,
		AvailabilityZone:     &req.AvailabilityZone,
		Name:                 req.Name,
		Count:                &ecsCount,
		IsAutoRename:         &isAutoRename,
		KeyName:              &req.KeyName,
		Vpcid:                req.VpcId,
		Nics:                 nics,
		Publicip:             publicip,
		RootVolume:           rootVolume,
		DataVolumes:          &dataVolumes,
		ServerTags:           &tags,
		SecurityGroups:       &securityGroups,
		BatchCreateInMultiAz: &batchCreateInMultiAz,
		UserData:             &req.UserData,
		Metadata:             metadata,
	}

	if req.EnterpriseProjectId != "" {
		serverBody.Extendparam = &ecsmodel.PrePaidServerExtendParam{
			EnterpriseProjectId: &req.EnterpriseProjectId,
		}
	}

	request := &ecsmodel.CreateServersRequest{
		Body: &ecsmodel.CreateServersRequestBody{
			Server: serverBody,
		},
	}

	response, err := s.client.CreateServers(request)
	if err != nil {
		return nil, fmt.Errorf("create ecs instance failed: %w", err)
	}

	return &cloudprovider.CreateInstanceResponse{
		JobId:       *response.JobId,
		InstanceIds: *response.ServerIds,
	}, nil
}

// DeleteInstance 删除虚拟机实例
func (s *HuaweiComputeService) DeleteInstance(ctx context.Context, instanceId string) error {
	deleteEIP := true
	deleteVol := true
	request := &ecsmodel.DeleteServersRequest{
		Body: &ecsmodel.DeleteServersRequestBody{
			DeletePublicip: &deleteEIP,
			DeleteVolume:   &deleteVol,
			Servers:        []ecsmodel.ServerId{{Id: instanceId}},
		},
	}

	_, err := s.client.DeleteServers(request)
	if err != nil {
		respErr, ok := err.(*sdkerr.ServiceResponseError)
		if ok && respErr.StatusCode == http.StatusNotFound {
			return nil // 已删除
		}
		return fmt.Errorf("delete ecs instance failed: %w", err)
	}
	return nil
}

// BatchDeleteInstances 批量删除虚拟机实例
func (s *HuaweiComputeService) BatchDeleteInstances(ctx context.Context, instanceIds []string) (string, error) {
	if len(instanceIds) == 0 {
		return "", nil
	}

	deleteEIP := true
	deleteVol := true
	var serverIds []ecsmodel.ServerId
	for _, id := range instanceIds {
		serverIds = append(serverIds, ecsmodel.ServerId{Id: id})
	}

	request := &ecsmodel.DeleteServersRequest{
		Body: &ecsmodel.DeleteServersRequestBody{
			DeletePublicip: &deleteEIP,
			DeleteVolume:   &deleteVol,
			Servers:        serverIds,
		},
	}

	response, err := s.client.DeleteServers(request)
	if err != nil {
		respErr, ok := err.(*sdkerr.ServiceResponseError)
		if ok && respErr.StatusCode == http.StatusNotFound {
			return "", nil
		}
		return "", fmt.Errorf("batch delete ecs instances failed: %w", err)
	}
	return *response.JobId, nil
}

// GetInstance 获取虚拟机实例信息
func (s *HuaweiComputeService) GetInstance(ctx context.Context, instanceId string) (*cloudprovider.Instance, error) {
	request := &ecsmodel.ShowServerRequest{ServerId: instanceId}
	response, err := s.client.ShowServer(request)
	if err != nil {
		return nil, fmt.Errorf("get ecs instance failed: %w", err)
	}
	return s.convertInstance(response.Server), nil
}

// ListInstances 列出虚拟机实例
func (s *HuaweiComputeService) ListInstances(ctx context.Context, req *cloudprovider.ListInstancesRequest) (*cloudprovider.ListInstancesResponse, error) {
	request := &ecsmodel.ListServersDetailsRequest{}
	if req.Name != "" {
		request.Name = &req.Name
	}
	if req.Status != "" {
		request.Status = &req.Status
	}
	if req.Limit > 0 {
		limit := int32(req.Limit)
		request.Limit = &limit
	}
	if req.Offset > 0 {
		offset := int32(req.Offset)
		request.Offset = &offset
	}

	response, err := s.client.ListServersDetails(request)
	if err != nil {
		return nil, fmt.Errorf("list ecs instances failed: %w", err)
	}

	var instances []cloudprovider.Instance
	if response.Servers != nil {
		for _, server := range *response.Servers {
			instances = append(instances, *s.convertServerDetail(&server))
		}
	}

	return &cloudprovider.ListInstancesResponse{
		Instances:  instances,
		TotalCount: int(*response.Count),
	}, nil
}

// StartInstance 启动虚拟机实例
func (s *HuaweiComputeService) StartInstance(ctx context.Context, instanceId string) error {
	request := &ecsmodel.BatchStartServersRequest{
		Body: &ecsmodel.BatchStartServersRequestBody{
			OsStart: &ecsmodel.BatchStartServersOption{
				Servers: []ecsmodel.ServerId{{Id: instanceId}},
			},
		},
	}
	_, err := s.client.BatchStartServers(request)
	if err != nil {
		return fmt.Errorf("start ecs instance failed: %w", err)
	}
	return nil
}

// StopInstance 停止虚拟机实例
func (s *HuaweiComputeService) StopInstance(ctx context.Context, instanceId string) error {
	request := &ecsmodel.BatchStopServersRequest{
		Body: &ecsmodel.BatchStopServersRequestBody{
			OsStop: &ecsmodel.BatchStopServersOption{
				Servers: []ecsmodel.ServerId{{Id: instanceId}},
			},
		},
	}
	_, err := s.client.BatchStopServers(request)
	if err != nil {
		return fmt.Errorf("stop ecs instance failed: %w", err)
	}
	return nil
}

// BatchStopInstances 批量停止虚拟机实例
func (s *HuaweiComputeService) BatchStopInstances(ctx context.Context, instanceIds []string) error {
	if len(instanceIds) == 0 {
		return nil
	}

	var serverIds []ecsmodel.ServerId
	for _, id := range instanceIds {
		serverIds = append(serverIds, ecsmodel.ServerId{Id: id})
	}

	request := &ecsmodel.BatchStopServersRequest{
		Body: &ecsmodel.BatchStopServersRequestBody{
			OsStop: &ecsmodel.BatchStopServersOption{
				Servers: serverIds,
			},
		},
	}
	_, err := s.client.BatchStopServers(request)
	if err != nil {
		return fmt.Errorf("batch stop ecs instances failed: %w", err)
	}
	return nil
}

// ListAvailabilityZones 列出可用区
func (s *HuaweiComputeService) ListAvailabilityZones(ctx context.Context) ([]string, error) {
	response, err := s.client.NovaListAvailabilityZones(&ecsmodel.NovaListAvailabilityZonesRequest{})
	if err != nil {
		return nil, fmt.Errorf("list availability zones failed: %w", err)
	}

	var azList []string
	if response.AvailabilityZoneInfo != nil {
		for _, az := range *response.AvailabilityZoneInfo {
			azList = append(azList, az.ZoneName)
		}
	}
	return azList, nil
}

// ListFlavors 列出规格
func (s *HuaweiComputeService) ListFlavors(ctx context.Context) (map[string][]string, error) {
	request := &ecsmodel.ListFlavorsRequest{}
	response, err := s.client.ListFlavors(request)
	if err != nil {
		return nil, fmt.Errorf("list flavors failed: %w", err)
	}

	result := make(map[string][]string)
	if response.Flavors != nil {
		for _, flavor := range *response.Flavors {
			result[flavor.Id] = []string{flavor.Name}
		}
	}
	return result, nil
}

// GetJob 获取异步任务状态
func (s *HuaweiComputeService) GetJob(ctx context.Context, jobId string) (*cloudprovider.JobInfo, error) {
	request := &ecsmodel.ShowJobRequest{JobId: jobId}
	response, err := s.client.ShowJob(request)
	if err != nil {
		return nil, fmt.Errorf("get job failed: %w", err)
	}

	jobInfo := &cloudprovider.JobInfo{
		Id:     *response.JobId,
		Status: response.Status.Value(),
	}

	if response.Entities != nil && response.Entities.SubJobs != nil {
		for _, subJob := range *response.Entities.SubJobs {
			subJobInfo := cloudprovider.SubJobInfo{
				Id:     *subJob.JobId,
				Status: subJob.Status.Value(),
			}
			if subJob.Entities != nil && subJob.Entities.ServerId != nil {
				subJobInfo.ServerId = *subJob.Entities.ServerId
			}
			jobInfo.SubJobs = append(jobInfo.SubJobs, subJobInfo)
		}
	}
	return jobInfo, nil
}

// UpdateInstanceTags 更新实例标签
func (s *HuaweiComputeService) UpdateInstanceTags(ctx context.Context, instanceId string, tags []cloudprovider.Tag) error {
	var ecsTags []ecsmodel.BatchAddServerTag
	for _, tag := range tags {
		ecsTags = append(ecsTags, ecsmodel.BatchAddServerTag{
			Key:   tag.Key,
			Value: tag.Value,
		})
	}

	request := &ecsmodel.BatchCreateServerTagsRequest{
		ServerId: instanceId,
		Body: &ecsmodel.BatchCreateServerTagsRequestBody{
			Tags:   ecsTags,
			Action: ecsmodel.GetBatchCreateServerTagsRequestBodyActionEnum().CREATE,
		},
	}

	_, err := s.client.BatchCreateServerTags(request)
	if err != nil {
		return fmt.Errorf("update instance tags failed: %w", err)
	}
	return nil
}

// ParseInstanceIPs 解析实例IP地址
func (s *HuaweiComputeService) ParseInstanceIPs(ctx context.Context, instance *cloudprovider.Instance, ipType string) ([]string, error) {
	if ipType == cloudprovider.IPTypePublic {
		return instance.PublicIPs, nil
	}
	return instance.PrivateIPs, nil
}

// buildVolumes 构建磁盘配置
func (s *HuaweiComputeService) buildVolumes(disks []cloudprovider.Disk) (*ecsmodel.PrePaidServerRootVolume, []ecsmodel.PrePaidServerDataVolume) {
	var rootVolume *ecsmodel.PrePaidServerRootVolume
	var dataVolumes []ecsmodel.PrePaidServerDataVolume

	for _, disk := range disks {
		if disk.DiskType == cloudprovider.DiskTypeSys {
			rootVolume = &ecsmodel.PrePaidServerRootVolume{
				Size:       &disk.Size,
				Volumetype: s.getRootVolumeType(disk.VolumeType),
			}
		} else {
			dataVolumes = append(dataVolumes, ecsmodel.PrePaidServerDataVolume{
				Volumetype: s.getDataVolumeType(disk.VolumeType),
				Size:       disk.Size,
			})
		}
	}
	return rootVolume, dataVolumes
}

// getRootVolumeType 获取系统盘类型
func (s *HuaweiComputeService) getRootVolumeType(volumeType string) ecsmodel.PrePaidServerRootVolumeVolumetype {
	volumeTypeEnum := ecsmodel.GetPrePaidServerRootVolumeVolumetypeEnum()
	switch volumeType {
	case cloudprovider.VolumeTypeSATA:
		return volumeTypeEnum.SATA
	case cloudprovider.VolumeTypeSAS:
		return volumeTypeEnum.SAS
	case cloudprovider.VolumeTypeSSD:
		return volumeTypeEnum.SSD
	case cloudprovider.VolumeTypeGPSSD:
		return volumeTypeEnum.GPSSD
	case cloudprovider.VolumeTypeESSD:
		return volumeTypeEnum.ESSD
	default:
		return volumeTypeEnum.GPSSD
	}
}

// getDataVolumeType 获取数据盘类型
func (s *HuaweiComputeService) getDataVolumeType(volumeType string) ecsmodel.PrePaidServerDataVolumeVolumetype {
	volumeTypeEnum := ecsmodel.GetPrePaidServerDataVolumeVolumetypeEnum()
	switch volumeType {
	case cloudprovider.VolumeTypeSATA:
		return volumeTypeEnum.SATA
	case cloudprovider.VolumeTypeSAS:
		return volumeTypeEnum.SAS
	case cloudprovider.VolumeTypeSSD:
		return volumeTypeEnum.SSD
	case cloudprovider.VolumeTypeGPSSD:
		return volumeTypeEnum.GPSSD
	case cloudprovider.VolumeTypeESSD:
		return volumeTypeEnum.ESSD
	default:
		return volumeTypeEnum.GPSSD
	}
}

// buildPublicIp 构建公网IP配置
func (s *HuaweiComputeService) buildPublicIp(eip *cloudprovider.EipConfig) *ecsmodel.PrePaidServerPublicip {
	shareType := ecsmodel.GetPrePaidServerEipBandwidthSharetypeEnum().PER
	if eip.ShareType == cloudprovider.ShareTypeWhole {
		shareType = ecsmodel.GetPrePaidServerEipBandwidthSharetypeEnum().WHOLE
	}

	bandwidth := ecsmodel.PrePaidServerEipBandwidth{
		Size:      &eip.BandwidthSize,
		Sharetype: shareType,
		Id:        &eip.BandwidthId,
	}

	return &ecsmodel.PrePaidServerPublicip{
		Eip: &ecsmodel.PrePaidServerEip{
			Iptype:    eip.IpType,
			Bandwidth: &bandwidth,
		},
	}
}

// convertInstance 转换实例信息
func (s *HuaweiComputeService) convertInstance(server *ecsmodel.ServerDetail) *cloudprovider.Instance {
	instance := &cloudprovider.Instance{
		Id:       server.Id,
		Name:     server.Name,
		Status:   server.Status,
		FlavorId: server.Flavor.Id,
		ImageId:  server.Image.Id,
		RawData:  server,
	}

	// 解析IP地址
	for _, addrs := range server.Addresses {
		for _, addr := range addrs {
			if *addr.OSEXTIPStype == ecsmodel.GetServerAddressOSEXTIPStypeEnum().FIXED {
				instance.PrivateIPs = append(instance.PrivateIPs, addr.Addr)
			} else if *addr.OSEXTIPStype == ecsmodel.GetServerAddressOSEXTIPStypeEnum().FLOATING {
				instance.PublicIPs = append(instance.PublicIPs, addr.Addr)
			}
		}
	}

	return instance
}

// convertServerDetail 转换服务器详情
func (s *HuaweiComputeService) convertServerDetail(server *ecsmodel.ServerDetail) *cloudprovider.Instance {
	return s.convertInstance(server)
}
