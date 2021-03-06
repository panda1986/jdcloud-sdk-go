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

package models


type UserAccessKey struct {

    /* 用户通行id (Optional) */
    Id int `json:"id"`

    /* 用户名 (Optional) */
    Pin string `json:"pin"`

    /* 用户账号 (Optional) */
    Account string `json:"account"`

    /* 访问密钥，AccessKey用于程序方式调用云服务API (Optional) */
    AccessKey string `json:"accessKey"`

    /* AccessKeySecret是用来验证用户的密钥 (Optional) */
    AccessKeySecret string `json:"accessKeySecret"`

    /* 到期时间 (Optional) */
    Expire string `json:"expire"`

    /* 创建时间 (Optional) */
    Created string `json:"created"`

    /* 更新时间 (Optional) */
    Modified string `json:"modified"`

    /* 更新操作人 (Optional) */
    Modifier string `json:"modifier"`

    /* 状态 (Optional) */
    State int `json:"state"`

    /*  (Optional) */
    Yn int `json:"yn"`
}
