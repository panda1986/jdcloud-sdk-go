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
    jmr "github.com/jdcloud-api/jdcloud-sdk-go/services/jmr/models"
)

type IsValidJobNameRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*   */
    JmrJobViewModel *jmr.JmrJobViewModel `json:"jmrJobViewModel"`
}

/*
 * param regionId: 地域ID (Required)
 * param jmrJobViewModel:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewIsValidJobNameRequest(
    regionId string,
    jmrJobViewModel *jmr.JmrJobViewModel,
) *IsValidJobNameRequest {

	return &IsValidJobNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/jobName:Validate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        JmrJobViewModel: jmrJobViewModel,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param jmrJobViewModel:  (Required)
 */
func NewIsValidJobNameRequestWithAllParams(
    regionId string,
    jmrJobViewModel *jmr.JmrJobViewModel,
) *IsValidJobNameRequest {

    return &IsValidJobNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jobName:Validate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        JmrJobViewModel: jmrJobViewModel,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewIsValidJobNameRequestWithoutParam() *IsValidJobNameRequest {

    return &IsValidJobNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jobName:Validate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *IsValidJobNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param jmrJobViewModel: (Required) */
func (r *IsValidJobNameRequest) SetJmrJobViewModel(jmrJobViewModel *jmr.JmrJobViewModel) {
    r.JmrJobViewModel = jmrJobViewModel
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r IsValidJobNameRequest) GetRegionId() string {
    return r.RegionId
}

type IsValidJobNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result IsValidJobNameResult `json:"result"`
}

type IsValidJobNameResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
}