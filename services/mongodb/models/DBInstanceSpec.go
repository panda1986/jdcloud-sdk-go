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


type DBInstanceSpec struct {

    /* 实例名称，只支持数字、字母、英文下划线、中文，且不少于2字符不超过32字符。 (Optional) */
    InstanceName *string `json:"instanceName"`

    /* 数据库类型，MongoDB (Optional) */
    Engine *string `json:"engine"`

    /* 数据库版本，3.2 (Optional) */
    EngineVersion *string `json:"engineVersion"`

    /* 实例规格代码。mongo.s1.small：1核2G;mongo.s1.medium：2核4G;mongo.s1.large：4核8G;mongo.s1.xlarge：8核16G;mongo.s2.2xlarge：8核32G;mongo.s2.4xlarge：16核64G;  */
    InstanceClass string `json:"instanceClass"`

    /* 存储空间，单位GB，取值10-1000,10的倍数。  */
    InstanceStorageGB int `json:"instanceStorageGB"`

    /* 是否选择多可用区部署  */
    MultiAZ bool `json:"multiAZ"`

    /* 可用区ID，必填，第一个ID为primary所在可用区ID，第二个为secondary，第三个为hidden。multiAZ选择是，则primary与secondary的可用区ID需相同，且与hidden不同；multiAZ选择否，三个节点写相同的可用区ID。  */
    AzId []string `json:"azId"`

    /* VPCID  */
    VpcId string `json:"vpcId"`

    /* 子网ID  */
    SubnetId string `json:"subnetId"`

    /* 密码，必须包含且只支持字母及数字，不少于8字符不超过16字符。 (Optional) */
    Password *string `json:"password"`

    /* 按备份创建使用的具体备份ID (Optional) */
    BackupId *string `json:"backupId"`

    /* 基于一个实例的备份创建新实例，如填写则restoreTime也需要填写。 (Optional) */
    OriginDBInstanceId *string `json:"originDBInstanceId"`

    /* 用户指定备份保留周期内的任意时间点，如2011-06-11T16:00:00Z，非必填，与backupId互斥。 (Optional) */
    RestoreTime *string `json:"restoreTime"`
}
