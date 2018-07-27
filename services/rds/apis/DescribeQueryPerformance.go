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
    rds "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/models"
)

type DescribeQueryPerformanceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 查询类型，不同的查询类型按照相应的字段从高到低返回结果。支持如下类型：ExecutionCount：执行次数LastRows：上次返回行数ElapsedTime：平均执行时间CPUTime：平均CPU时间LogicalReads：平均逻辑读LogicalWrites：平均逻辑写PhysicalReads：平均物理读  */
    QueryType string `json:"queryType"`

    /* 只返回查询条件大于等于threshold的记录，默认为0 (Optional) */
    Threshold *int `json:"threshold"`

    /* 显示数据的页码，取值范围：[1,1000)，页码超过总页数时，显示最后一页，用于查询列表的接口 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为50，取值范围：[1,100]，只能为10的倍数 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param queryType: 查询类型，不同的查询类型按照相应的字段从高到低返回结果。支持如下类型：ExecutionCount：执行次数LastRows：上次返回行数ElapsedTime：平均执行时间CPUTime：平均CPU时间LogicalReads：平均逻辑读LogicalWrites：平均逻辑写PhysicalReads：平均物理读 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeQueryPerformanceRequest(
    regionId string,
    instanceId string,
    queryType string,
) *DescribeQueryPerformanceRequest {

	return &DescribeQueryPerformanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/performance:describeQueryPerformance",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        QueryType: queryType,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param queryType: 查询类型，不同的查询类型按照相应的字段从高到低返回结果。支持如下类型：ExecutionCount：执行次数LastRows：上次返回行数ElapsedTime：平均执行时间CPUTime：平均CPU时间LogicalReads：平均逻辑读LogicalWrites：平均逻辑写PhysicalReads：平均物理读 (Required)
 * param threshold: 只返回查询条件大于等于threshold的记录，默认为0 (Optional)
 * param pageNumber: 显示数据的页码，取值范围：[1,1000)，页码超过总页数时，显示最后一页，用于查询列表的接口 (Optional)
 * param pageSize: 每页显示的数据条数，默认为50，取值范围：[1,100]，只能为10的倍数 (Optional)
 */
func NewDescribeQueryPerformanceRequestWithAllParams(
    regionId string,
    instanceId string,
    queryType string,
    threshold *int,
    pageNumber *int,
    pageSize *int,
) *DescribeQueryPerformanceRequest {

    return &DescribeQueryPerformanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/performance:describeQueryPerformance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        QueryType: queryType,
        Threshold: threshold,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeQueryPerformanceRequestWithoutParam() *DescribeQueryPerformanceRequest {

    return &DescribeQueryPerformanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/performance:describeQueryPerformance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeQueryPerformanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: Instance ID(Required) */
func (r *DescribeQueryPerformanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param queryType: 查询类型，不同的查询类型按照相应的字段从高到低返回结果。支持如下类型：ExecutionCount：执行次数LastRows：上次返回行数ElapsedTime：平均执行时间CPUTime：平均CPU时间LogicalReads：平均逻辑读LogicalWrites：平均逻辑写PhysicalReads：平均物理读(Required) */
func (r *DescribeQueryPerformanceRequest) SetQueryType(queryType string) {
    r.QueryType = queryType
}

/* param threshold: 只返回查询条件大于等于threshold的记录，默认为0(Optional) */
func (r *DescribeQueryPerformanceRequest) SetThreshold(threshold int) {
    r.Threshold = &threshold
}

/* param pageNumber: 显示数据的页码，取值范围：[1,1000)，页码超过总页数时，显示最后一页，用于查询列表的接口(Optional) */
func (r *DescribeQueryPerformanceRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，默认为50，取值范围：[1,100]，只能为10的倍数(Optional) */
func (r *DescribeQueryPerformanceRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeQueryPerformanceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeQueryPerformanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeQueryPerformanceResult `json:"result"`
}

type DescribeQueryPerformanceResult struct {
    QueryPerformanceResult []rds.QueryPerformanceResult `json:"queryPerformanceResult"`
    TotalCount int `json:"totalCount"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
}