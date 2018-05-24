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

package apis

import (
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    "reflect"
    streambus "github.com/jdcloud-api/jdcloud-sdk-go/services/streambus/models"
)

type GetTopicListRequest struct {

    JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    Keyword *string `json:"keyword"`
}

/*
 * param regionId: Region ID 
 * param keyword:  (Optional)
 */
func NewGetTopicListRequest(
    regionId string,
) *GetTopicListRequest {

	return &GetTopicListRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/topicList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *GetTopicListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *GetTopicListRequest) SetKeyword(keyword string) {
    r.Keyword = &keyword
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetTopicListRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type GetTopicListResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result GetTopicListResult `json:"result"`
}

type GetTopicListResult struct {
    Topic []streambus.TopicListInfo `json:"topic"`
}