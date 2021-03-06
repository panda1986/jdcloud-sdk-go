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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type CreateAlarmRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 幂等性校验参数，最长32位，值不变则返回值不会变  */
    ClientToken string `json:"clientToken"`

    /*   */
    CreateAlarmSpec *monitor.CreateAlarmSpec `json:"createAlarmSpec"`
}

/*
 * param regionId: 地域 Id (Required)
 * param clientToken: 幂等性校验参数，最长32位，值不变则返回值不会变 (Required)
 * param createAlarmSpec:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateAlarmRequest(
    regionId string,
    clientToken string,
    createAlarmSpec *monitor.CreateAlarmSpec,
) *CreateAlarmRequest {

	return &CreateAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/alarms",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClientToken: clientToken,
        CreateAlarmSpec: createAlarmSpec,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param clientToken: 幂等性校验参数，最长32位，值不变则返回值不会变 (Required)
 * param createAlarmSpec:  (Required)
 */
func NewCreateAlarmRequestWithAllParams(
    regionId string,
    clientToken string,
    createAlarmSpec *monitor.CreateAlarmSpec,
) *CreateAlarmRequest {

    return &CreateAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientToken: clientToken,
        CreateAlarmSpec: createAlarmSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateAlarmRequestWithoutParam() *CreateAlarmRequest {

    return &CreateAlarmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *CreateAlarmRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clientToken: 幂等性校验参数，最长32位，值不变则返回值不会变(Required) */
func (r *CreateAlarmRequest) SetClientToken(clientToken string) {
    r.ClientToken = clientToken
}

/* param createAlarmSpec: (Required) */
func (r *CreateAlarmRequest) SetCreateAlarmSpec(createAlarmSpec *monitor.CreateAlarmSpec) {
    r.CreateAlarmSpec = createAlarmSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAlarmRequest) GetRegionId() string {
    return r.RegionId
}

type CreateAlarmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateAlarmResult `json:"result"`
}

type CreateAlarmResult struct {
    AlarmIdList []string `json:"alarmIdList"`
}