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
    . "github.com/jdcloud-api/jdcloud-sdk-go/services/streambus/apis"
    "encoding/json"
    "errors"
)

type StreambusClient struct {
    JDCloudClient
}

func NewStreambusClient(credential *Credential) *StreambusClient {
    if credential == nil {
        return nil
    }

    config := NewConfig()
    config.SetEndpoint("streambus.jdcloud-api.com")

    return &StreambusClient{
        JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "streambus",
            Revision:    "0.2.4",
            Logger:      NewDefaultLogger(LOG_INFO),
        }}
}

func (c *StreambusClient) SetConfig(config *Config) {
    c.Config = *config
}

func (c *StreambusClient) SetLogger(logger Logger) {
    c.Logger = logger
}

/* 创建topic */
func (c *StreambusClient) AddTopic(request *AddTopicRequest) (*AddTopicResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &AddTopicResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 更新topic */
func (c *StreambusClient) UpdateTopic(request *UpdateTopicRequest) (*UpdateTopicResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &UpdateTopicResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询topic */
func (c *StreambusClient) GetTopicList(request *GetTopicListRequest) (*GetTopicListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &GetTopicListResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}
