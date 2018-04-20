// Copyright 2018-2025 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package client

import (
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    . "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
    "encoding/json"
    "errors"
)

type VpcClient struct {
    JDCloudClient
}

func NewVpcClient(credential *Credential) *VpcClient {
    if credential == nil {
        return nil
    }

    config := NewConfig()
    config.SetEndpoint("vpc.jdcloud-api.com")

    return &VpcClient{
        JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "vpc",
            Revision:    "0.2.0",
            Logger:      NewDefaultLogger(LOG_INFO),
        }}
}

func (c *VpcClient) SetConfig(config *Config) {
    c.Config = *config
}

func (c *VpcClient) SetLogger(logger Logger) {
    c.Logger = logger
}

/* 查询弹性ip列表 */
func (c *VpcClient) DescribeElasticIps(request *DescribeElasticIpsRequest) (*DescribeElasticIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeElasticIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 给网卡分配secondaryIp接口 */
func (c *VpcClient) AssignSecondaryIps(request *AssignSecondaryIpsRequest) (*AssignSecondaryIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &AssignSecondaryIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询VpcPeering资源详情 */
func (c *VpcClient) DescribeVpcPeering(request *DescribeVpcPeeringRequest) (*DescribeVpcPeeringResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeVpcPeeringResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 给网卡解绑弹性Ip接口 */
func (c *VpcClient) DisassociateElasticIp(request *DisassociateElasticIpRequest) (*DisassociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DisassociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 删除弹性Ip */
func (c *VpcClient) DeleteElasticIp(request *DeleteElasticIpRequest) (*DeleteElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DeleteElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 创建一个或者多个弹性Ip */
func (c *VpcClient) CreateElasticIps(request *CreateElasticIpsRequest) (*CreateElasticIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &CreateElasticIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 给网卡删除secondaryIp接口 */
func (c *VpcClient) UnassignSecondaryIps(request *UnassignSecondaryIpsRequest) (*UnassignSecondaryIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &UnassignSecondaryIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询Vpc信息详情 */
func (c *VpcClient) DescribeVpc(request *DescribeVpcRequest) (*DescribeVpcResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeVpcResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询VpcPeering资源列表 */
func (c *VpcClient) DescribeVpcPeerings(request *DescribeVpcPeeringsRequest) (*DescribeVpcPeeringsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeVpcPeeringsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询安全组列表 */
func (c *VpcClient) DescribeNetworkSecurityGroups(request *DescribeNetworkSecurityGroupsRequest) (*DescribeNetworkSecurityGroupsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeNetworkSecurityGroupsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询子网列表 */
func (c *VpcClient) DescribeSubnets(request *DescribeSubnetsRequest) (*DescribeSubnetsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeSubnetsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询子网信息详情 */
func (c *VpcClient) DescribeSubnet(request *DescribeSubnetRequest) (*DescribeSubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeSubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 创建VpcPeering接口 */
func (c *VpcClient) CreateVpcPeering(request *CreateVpcPeeringRequest) (*CreateVpcPeeringResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &CreateVpcPeeringResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询安全组信息详情 */
func (c *VpcClient) DescribeNetworkSecurityGroup(request *DescribeNetworkSecurityGroupRequest) (*DescribeNetworkSecurityGroupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeNetworkSecurityGroupResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 修改VpcPeering接口 */
func (c *VpcClient) ModifyVpcPeering(request *ModifyVpcPeeringRequest) (*ModifyVpcPeeringResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ModifyVpcPeeringResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 给网卡绑定弹性Ip接口 */
func (c *VpcClient) AssociateElasticIp(request *AssociateElasticIpRequest) (*AssociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &AssociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 删除VpcPeering接口 */
func (c *VpcClient) DeleteVpcPeering(request *DeleteVpcPeeringRequest) (*DeleteVpcPeeringResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DeleteVpcPeeringResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* ElasticIp资源信息详情 */
func (c *VpcClient) DescribeElasticIp(request *DescribeElasticIpRequest) (*DescribeElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询私有网络列表 */
func (c *VpcClient) DescribeVpcs(request *DescribeVpcsRequest) (*DescribeVpcsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeVpcsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}
