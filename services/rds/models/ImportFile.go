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


type ImportFile struct {

    /* 文件名称 (Optional) */
    Name string `json:"name"`

    /* 如果该文件是共享文件，则有全局ID，如不是共享文件，则为空。该全局ID在文件删除时，需要用到 (Optional) */
    SharedFileGid string `json:"sharedFileGid"`

    /* 文件大小，单位Byte (Optional) */
    SizeByte int `json:"sizeByte"`

    /* 文件上传完成时间，格式为：YYYY-MM-DD HH:mm:ss (Optional) */
    UploadTime string `json:"uploadTime"`

    /* 是否所属当前实例.<br> 1：当前实例；<br>0：不是当前实例，为共享文件 (Optional) */
    IsLocal string `json:"isLocal"`
}
