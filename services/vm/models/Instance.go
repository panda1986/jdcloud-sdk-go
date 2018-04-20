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

package models

import . "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type Instance struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例类型 (Optional) */
    InstanceType string `json:"instanceType"`

    /* 主网卡所属VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 主网卡所属子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 主网卡主IP地址 (Optional) */
    PrivateIpAddress string `json:"privateIpAddress"`

    /* 主网卡主IP绑定弹性IP的ID (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 主网卡主IP绑定弹性IP的地址 (Optional) */
    ElasticIpAddress string `json:"elasticIpAddress"`

    /* 实例状态 (Optional) */
    Status string `json:"status"`

    /* 实例描述 (Optional) */
    Description string `json:"description"`

    /* 镜像ID (Optional) */
    ImageId string `json:"imageId"`

    /* 系统盘信息 (Optional) */
    SystemDisk InstanceDiskAttachment `json:"systemDisk"`

    /* 数据盘信息 (Optional) */
    DataDisks []InstanceDiskAttachment `json:"dataDisks"`

    /* 主网卡信息 (Optional) */
    PrimaryNetworkInterface InstanceNetworkInterfaceAttachment `json:"primaryNetworkInterface"`

    /* 创建时间 (Optional) */
    LaunchTime string `json:"launchTime"`

    /* 可用区 (Optional) */
    Az string `json:"az"`

    /* 密钥对名称 (Optional) */
    KeyNames []string `json:"keyNames"`

    /* 计费信息 (Optional) */
    Charge Charge `json:"charge"`
}