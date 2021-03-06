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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type CreateForwardRuleRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`

    /* 非网站类规则参数  */
    ForwardRuleSpec *ipanti.ForwardRuleSpec `json:"forwardRuleSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例id (Required)
 * param forwardRuleSpec: 非网站类规则参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateForwardRuleRequest(
    regionId string,
    instanceId string,
    forwardRuleSpec *ipanti.ForwardRuleSpec,
) *CreateForwardRuleRequest {

	return &CreateForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleSpec: forwardRuleSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例id (Required)
 * param forwardRuleSpec: 非网站类规则参数 (Required)
 */
func NewCreateForwardRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    forwardRuleSpec *ipanti.ForwardRuleSpec,
) *CreateForwardRuleRequest {

    return &CreateForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleSpec: forwardRuleSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateForwardRuleRequestWithoutParam() *CreateForwardRuleRequest {

    return &CreateForwardRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateForwardRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例id(Required) */
func (r *CreateForwardRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param forwardRuleSpec: 非网站类规则参数(Required) */
func (r *CreateForwardRuleRequest) SetForwardRuleSpec(forwardRuleSpec *ipanti.ForwardRuleSpec) {
    r.ForwardRuleSpec = forwardRuleSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateForwardRuleRequest) GetRegionId() string {
    return r.RegionId
}

type CreateForwardRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateForwardRuleResult `json:"result"`
}

type CreateForwardRuleResult struct {
}