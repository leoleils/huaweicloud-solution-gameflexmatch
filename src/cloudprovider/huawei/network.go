// Copyright (c) Huawei Technologies Co., Ltd. 2022-2022. All rights reserved.

// Package huawei 华为云NetworkService实现
package huawei

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"scase.io/cloudprovider"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	vpcmodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
)

const defaultCallIntervalTimes = 5 * time.Second

// HuaweiNetworkService 华为云网络服务实现
type HuaweiNetworkService struct {
	provider  *HuaweiProvider
	vpcClient *vpc.VpcClient
}

// CreateVpc 创建VPC
func (s *HuaweiNetworkService) CreateVpc(ctx context.Context, req *cloudprovider.CreateVpcRequest) (string, error) {
	request := &vpcmodel.CreateVpcRequest{
		Body: &vpcmodel.CreateVpcRequestBody{
			Vpc: &vpcmodel.CreateVpcOption{
				Cidr:                &req.Cidr,
				Name:                &req.Name,
				EnterpriseProjectId: &req.EnterpriseProjectId,
			},
		},
	}

	resp, err := s.vpcClient.CreateVpc(request)
	if err != nil {
		return "", fmt.Errorf("create vpc failed: %w", err)
	}
	return resp.Vpc.Id, nil
}

// GetVpc 获取VPC
func (s *HuaweiNetworkService) GetVpc(ctx context.Context, vpcId string) (*cloudprovider.Vpc, error) {
	request := &vpcmodel.ShowVpcRequest{VpcId: vpcId}
	resp, err := s.vpcClient.ShowVpc(request)
	if err != nil {
		return nil, fmt.Errorf("get vpc failed: %w", err)
	}

	return &cloudprovider.Vpc{
		Id:     resp.Vpc.Id,
		Name:   resp.Vpc.Name,
		Cidr:   resp.Vpc.Cidr,
		Status: resp.Vpc.Status.Value(),
	}, nil
}

// GetVpcByName 根据名称获取VPC
func (s *HuaweiNetworkService) GetVpcByName(ctx context.Context, name string) (*cloudprovider.Vpc, error) {
	request := &vpcmodel.ListVpcsRequest{}
	resp, err := s.vpcClient.ListVpcs(request)
	if err != nil {
		return nil, fmt.Errorf("list vpcs failed: %w", err)
	}

	if resp.Vpcs != nil {
		for _, v := range *resp.Vpcs {
			if v.Name == name {
				return &cloudprovider.Vpc{
					Id:     v.Id,
					Name:   v.Name,
					Cidr:   v.Cidr,
					Status: v.Status.Value(),
				}, nil
			}
		}
	}
	return nil, nil
}

// DeleteVpc 删除VPC
func (s *HuaweiNetworkService) DeleteVpc(ctx context.Context, vpcId string) error {
	request := &vpcmodel.DeleteVpcRequest{VpcId: vpcId}
	_, err := s.vpcClient.DeleteVpc(request)
	if err != nil {
		respErr, ok := err.(*sdkerr.ServiceResponseError)
		if ok && respErr.StatusCode == http.StatusNotFound {
			return nil
		}
		return fmt.Errorf("delete vpc failed: %w", err)
	}
	return nil
}

// WaitVpcReady 等待VPC就绪
func (s *HuaweiNetworkService) WaitVpcReady(ctx context.Context, vpcId string) error {
	for {
		request := &vpcmodel.ShowVpcRequest{VpcId: vpcId}
		resp, err := s.vpcClient.ShowVpc(request)
		if err != nil {
			return fmt.Errorf("show vpc failed: %w", err)
		}

		if resp.Vpc.Status == vpcmodel.GetVpcStatusEnum().OK {
			return nil
		}

		if resp.Vpc.Status == vpcmodel.GetVpcStatusEnum().ERROR {
			return fmt.Errorf("vpc status error")
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(defaultCallIntervalTimes):
		}
	}
}

// WaitVpcDeleted 等待VPC删除完成
func (s *HuaweiNetworkService) WaitVpcDeleted(ctx context.Context, vpcId string) error {
	for {
		request := &vpcmodel.ShowVpcRequest{VpcId: vpcId}
		_, err := s.vpcClient.ShowVpc(request)
		if err != nil {
			respErr, ok := err.(*sdkerr.ServiceResponseError)
			if ok && respErr.StatusCode == http.StatusNotFound {
				return nil
			}
			return fmt.Errorf("show vpc failed: %w", err)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(defaultCallIntervalTimes):
		}
	}
}

// CreateSubnet 创建子网
func (s *HuaweiNetworkService) CreateSubnet(ctx context.Context, req *cloudprovider.CreateSubnetRequest) (*cloudprovider.CreateSubnetResponse, error) {
	request := &vpcmodel.CreateSubnetRequest{
		Body: &vpcmodel.CreateSubnetRequestBody{
			Subnet: &vpcmodel.CreateSubnetOption{
				Cidr:         req.Cidr,
				Name:         req.Name,
				VpcId:        req.VpcId,
				GatewayIp:    req.GatewayIp,
				PrimaryDns:   &req.PrimaryDns,
				SecondaryDns: &req.SecondaryDns,
				DnsList:      &req.DnsList,
			},
		},
	}

	resp, err := s.vpcClient.CreateSubnet(request)
	if err != nil {
		return nil, fmt.Errorf("create subnet failed: %w", err)
	}

	return &cloudprovider.CreateSubnetResponse{
		SubnetId:        resp.Subnet.Id,
		NeutronSubnetId: resp.Subnet.NeutronSubnetId,
	}, nil
}

// GetSubnet 获取子网
func (s *HuaweiNetworkService) GetSubnet(ctx context.Context, subnetId string) (*cloudprovider.Subnet, error) {
	request := &vpcmodel.ShowSubnetRequest{SubnetId: subnetId}
	resp, err := s.vpcClient.ShowSubnet(request)
	if err != nil {
		return nil, fmt.Errorf("get subnet failed: %w", err)
	}

	return &cloudprovider.Subnet{
		Id:               resp.Subnet.Id,
		Name:             resp.Subnet.Name,
		Cidr:             resp.Subnet.Cidr,
		VpcId:            resp.Subnet.VpcId,
		GatewayIp:        resp.Subnet.GatewayIp,
		AvailabilityZone: resp.Subnet.AvailabilityZone,
		Status:           resp.Subnet.Status.Value(),
		NeutronSubnetId:  resp.Subnet.NeutronSubnetId,
	}, nil
}

// ListSubnets 列出子网
func (s *HuaweiNetworkService) ListSubnets(ctx context.Context, vpcId string) ([]cloudprovider.Subnet, error) {
	request := &vpcmodel.ListSubnetsRequest{VpcId: &vpcId}
	resp, err := s.vpcClient.ListSubnets(request)
	if err != nil {
		return nil, fmt.Errorf("list subnets failed: %w", err)
	}

	var subnets []cloudprovider.Subnet
	if resp.Subnets != nil {
		for _, subnet := range *resp.Subnets {
			subnets = append(subnets, cloudprovider.Subnet{
				Id:               subnet.Id,
				Name:             subnet.Name,
				Cidr:             subnet.Cidr,
				VpcId:            subnet.VpcId,
				GatewayIp:        subnet.GatewayIp,
				AvailabilityZone: subnet.AvailabilityZone,
				Status:           subnet.Status.Value(),
				NeutronSubnetId:  subnet.NeutronSubnetId,
			})
		}
	}
	return subnets, nil
}

// DeleteSubnet 删除子网
func (s *HuaweiNetworkService) DeleteSubnet(ctx context.Context, vpcId, subnetId string) error {
	request := &vpcmodel.DeleteSubnetRequest{
		VpcId:    vpcId,
		SubnetId: subnetId,
	}
	_, err := s.vpcClient.DeleteSubnet(request)
	if err != nil {
		respErr, ok := err.(*sdkerr.ServiceResponseError)
		if ok && respErr.StatusCode == http.StatusNotFound {
			return nil
		}
		return fmt.Errorf("delete subnet failed: %w", err)
	}
	return nil
}

// WaitSubnetReady 等待子网就绪
func (s *HuaweiNetworkService) WaitSubnetReady(ctx context.Context, subnetId string) error {
	for {
		request := &vpcmodel.ShowSubnetRequest{SubnetId: subnetId}
		resp, err := s.vpcClient.ShowSubnet(request)
		if err != nil {
			return fmt.Errorf("show subnet failed: %w", err)
		}

		if resp.Subnet.Status == vpcmodel.GetSubnetStatusEnum().ACTIVE {
			return nil
		}

		if resp.Subnet.Status == vpcmodel.GetSubnetStatusEnum().ERROR {
			return fmt.Errorf("subnet status error")
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(defaultCallIntervalTimes):
		}
	}
}

// WaitSubnetDeleted 等待子网删除完成
func (s *HuaweiNetworkService) WaitSubnetDeleted(ctx context.Context, subnetId string) error {
	for {
		request := &vpcmodel.ShowSubnetRequest{SubnetId: subnetId}
		_, err := s.vpcClient.ShowSubnet(request)
		if err != nil {
			respErr, ok := err.(*sdkerr.ServiceResponseError)
			if ok && respErr.StatusCode == http.StatusNotFound {
				return nil
			}
			return fmt.Errorf("show subnet failed: %w", err)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(defaultCallIntervalTimes):
		}
	}
}

// CheckSubnetConflict 检查子网冲突
func (s *HuaweiNetworkService) CheckSubnetConflict(ctx context.Context, vpcId, cidr string) (bool, error) {
	// 获取VPC信息
	vpcReq := &vpcmodel.ShowVpcRequest{VpcId: vpcId}
	vpcResp, err := s.vpcClient.ShowVpc(vpcReq)
	if err != nil {
		return false, fmt.Errorf("show vpc failed: %w", err)
	}

	// 获取子网列表
	subnetReq := &vpcmodel.ListSubnetsRequest{VpcId: &vpcId}
	subnetResp, err := s.vpcClient.ListSubnets(subnetReq)
	if err != nil {
		return false, fmt.Errorf("list subnets failed: %w", err)
	}

	_, newCidr, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, fmt.Errorf("parse cidr failed: %w", err)
	}

	_, vpcCidr, err := net.ParseCIDR(vpcResp.Vpc.Cidr)
	if err != nil {
		return false, fmt.Errorf("parse vpc cidr failed: %w", err)
	}

	// 检查是否在VPC范围内
	if !containsCidr(vpcCidr, newCidr) {
		return false, nil
	}

	// 检查与现有子网是否冲突
	if subnetResp.Subnets != nil {
		for _, subnet := range *subnetResp.Subnets {
			_, existingCidr, err := net.ParseCIDR(subnet.Cidr)
			if err != nil {
				return false, fmt.Errorf("parse existing subnet cidr failed: %w", err)
			}

			if containsCidr(newCidr, existingCidr) || containsCidr(existingCidr, newCidr) {
				return true, nil
			}
		}
	}
	return false, nil
}

// CreateSecurityGroup 创建安全组
func (s *HuaweiNetworkService) CreateSecurityGroup(ctx context.Context, req *cloudprovider.CreateSecurityGroupRequest) (string, error) {
	request := &vpcmodel.CreateSecurityGroupRequest{
		Body: &vpcmodel.CreateSecurityGroupRequestBody{
			SecurityGroup: &vpcmodel.CreateSecurityGroupOption{
				Name:                req.Name,
				EnterpriseProjectId: &req.EnterpriseProjectId,
			},
		},
	}

	resp, err := s.vpcClient.CreateSecurityGroup(request)
	if err != nil {
		return "", fmt.Errorf("create security group failed: %w", err)
	}
	return resp.SecurityGroup.Id, nil
}

// GetSecurityGroup 获取安全组
func (s *HuaweiNetworkService) GetSecurityGroup(ctx context.Context, name string) (string, error) {
	request := &vpcmodel.ListSecurityGroupsRequest{}
	resp, err := s.vpcClient.ListSecurityGroups(request)
	if err != nil {
		return "", fmt.Errorf("list security groups failed: %w", err)
	}

	if resp.SecurityGroups != nil {
		for _, sg := range *resp.SecurityGroups {
			if sg.Name == name {
				return sg.Id, nil
			}
		}
	}
	return "", nil
}

// ListSecurityGroups 列出安全组
func (s *HuaweiNetworkService) ListSecurityGroups(ctx context.Context) ([]cloudprovider.SecurityGroup, error) {
	request := &vpcmodel.NeutronListSecurityGroupsRequest{}
	resp, err := s.vpcClient.NeutronListSecurityGroups(request)
	if err != nil {
		return nil, fmt.Errorf("list security groups failed: %w", err)
	}

	var securityGroups []cloudprovider.SecurityGroup
	if resp.SecurityGroups != nil {
		for _, sg := range *resp.SecurityGroups {
			securityGroups = append(securityGroups, cloudprovider.SecurityGroup{
				Id:   sg.Id,
				Name: sg.Name,
			})
		}
	}
	return securityGroups, nil
}

// CreateEip 创建弹性公网IP
func (s *HuaweiNetworkService) CreateEip(ctx context.Context, req *cloudprovider.CreateEipRequest) (string, error) {
	// 此处需要使用EIP客户端，简化处理
	return "", fmt.Errorf("not implemented: use eip client")
}

// DeleteEip 删除弹性公网IP
func (s *HuaweiNetworkService) DeleteEip(ctx context.Context, eipId string) error {
	// 此处需要使用EIP客户端，简化处理
	return fmt.Errorf("not implemented: use eip client")
}

// CreateBandwidth 创建带宽
func (s *HuaweiNetworkService) CreateBandwidth(ctx context.Context, req *cloudprovider.CreateBandwidthRequest) (string, error) {
	// 此处需要使用EIP客户端，简化处理
	return "", fmt.Errorf("not implemented: use eip client")
}

// ListBandwidths 列出带宽
func (s *HuaweiNetworkService) ListBandwidths(ctx context.Context) ([]cloudprovider.Bandwidth, error) {
	// 此处需要使用EIP客户端，简化处理
	return nil, fmt.Errorf("not implemented: use eip client")
}

// containsCidr 判断cidr a是否包含cidr b
func containsCidr(a, b *net.IPNet) bool {
	onesA, _ := a.Mask.Size()
	onesB, _ := b.Mask.Size()
	return onesA <= onesB && a.Contains(b.IP)
}
