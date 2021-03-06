// Copyright 2018 JDCLOUD.COM
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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    baseanti "github.com/jdcloud-api/jdcloud-sdk-go/services/baseanti/models"
)

type DescribeIpResourceFlowRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 公网ip  */
    Ip string `json:"ip"`

    /* 查询的结束时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param regionId: Region ID (Required)
 * param ip: 公网ip (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIpResourceFlowRequest(
    regionId string,
    ip string,
) *DescribeIpResourceFlowRequest {

	return &DescribeIpResourceFlowRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ipResources/{ip}/monitorFlow",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Ip: ip,
	}
}

/*
 * param regionId: Region ID (Required)
 * param ip: 公网ip (Required)
 * param endTime: 查询的结束时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ (Optional)
 */
func NewDescribeIpResourceFlowRequestWithAllParams(
    regionId string,
    ip string,
    endTime *string,
) *DescribeIpResourceFlowRequest {

    return &DescribeIpResourceFlowRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources/{ip}/monitorFlow",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Ip: ip,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIpResourceFlowRequestWithoutParam() *DescribeIpResourceFlowRequest {

    return &DescribeIpResourceFlowRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources/{ip}/monitorFlow",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeIpResourceFlowRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param ip: 公网ip(Required) */
func (r *DescribeIpResourceFlowRequest) SetIp(ip string) {
    r.Ip = ip
}

/* param endTime: 查询的结束时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ(Optional) */
func (r *DescribeIpResourceFlowRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIpResourceFlowRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIpResourceFlowResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIpResourceFlowResult `json:"result"`
}

type DescribeIpResourceFlowResult struct {
    Data baseanti.IpResourceFlow `json:"data"`
}