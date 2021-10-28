# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/all.proto](#npool/all.proto)
    - [BroadcastScriptRequest](#sphinx.v1.BroadcastScriptRequest)
    - [GetListRequest](#sphinx.v1.GetListRequest)
    - [GetSignInfoRequest](#sphinx.v1.GetSignInfoRequest)
    - [NodeInfo](#sphinx.v1.NodeInfo)
    - [ProcessReviewRequest](#sphinx.v1.ProcessReviewRequest)
    - [RechargeUserRequest](#sphinx.v1.RechargeUserRequest)
    - [RegisterUserRequest](#sphinx.v1.RegisterUserRequest)
    - [ReviewList](#sphinx.v1.ReviewList)
    - [ReviewRow](#sphinx.v1.ReviewRow)
    - [ScriptInfo](#sphinx.v1.ScriptInfo)
    - [SetAdminPermissionRequest](#sphinx.v1.SetAdminPermissionRequest)
    - [SignInfo](#sphinx.v1.SignInfo)
    - [StatusScriptRequest](#sphinx.v1.StatusScriptRequest)
    - [SuccessCode](#sphinx.v1.SuccessCode)
    - [UserAddress](#sphinx.v1.UserAddress)
    - [WithdrawApplyRequest](#sphinx.v1.WithdrawApplyRequest)
    - [WithdrawConfirmedRequest](#sphinx.v1.WithdrawConfirmedRequest)
  
    - [Review](#sphinx.v1.Review)
    - [Trading](#sphinx.v1.Trading)
    - [WalletAgent](#sphinx.v1.WalletAgent)
    - [WalletNode](#sphinx.v1.WalletNode)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/all.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/all.proto



<a name="sphinx.v1.BroadcastScriptRequest"></a>

### BroadcastScriptRequest
BroadcastScript 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| review_id | [uint32](#uint32) |  |  |
| transaction_script | [string](#string) |  |  |






<a name="sphinx.v1.GetListRequest"></a>

### GetListRequest
GetList 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| return_all | [bool](#bool) |  |  |






<a name="sphinx.v1.GetSignInfoRequest"></a>

### GetSignInfoRequest
GetSignInfo 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address_send | [string](#string) |  |  |






<a name="sphinx.v1.NodeInfo"></a>

### NodeInfo
AcceptNode 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| ip_uint64 | [uint64](#uint64) |  |  |
| ip_is_public | [bool](#bool) |  |  |
| mac_address | [string](#string) |  |  |
| location | [string](#string) |  |  |






<a name="sphinx.v1.ProcessReviewRequest"></a>

### ProcessReviewRequest
ProcessReview 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin_id | [int32](#int32) |  |  |
| is_reject | [bool](#bool) |  |  |
| is_approve | [bool](#bool) |  |  |
| review_ids | [uint32](#uint32) | repeated |  |






<a name="sphinx.v1.RechargeUserRequest"></a>

### RechargeUserRequest
RechargeUser 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int32](#int32) |  |  |
| coin_id | [int32](#int32) |  |  |
| amount | [double](#double) |  |  |
| datetime | [uint64](#uint64) |  |  |
| address | [string](#string) |  |  |






<a name="sphinx.v1.RegisterUserRequest"></a>

### RegisterUserRequest
RegisterUser 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int32](#int32) |  |  |
| coin_id | [int32](#int32) |  |  |






<a name="sphinx.v1.ReviewList"></a>

### ReviewList
GetList 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list | [ReviewRow](#sphinx.v1.ReviewRow) | repeated |  |






<a name="sphinx.v1.ReviewRow"></a>

### ReviewRow
单行审核数据（暂定）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint32](#uint32) |  |  |
| user_id | [int32](#int32) |  |  |
| coin_id | [int32](#int32) |  |  |
| amount | [double](#double) |  |  |
| datetime | [int64](#int64) |  |  |
| sig_salt | [uint64](#uint64) |  |  |
| sign_user | [string](#string) |  |  |
| sign_admin | [string](#string) |  |  |
| address | [string](#string) |  |  |






<a name="sphinx.v1.ScriptInfo"></a>

### ScriptInfo
StatusScript 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data_json | [string](#string) |  |  |






<a name="sphinx.v1.SetAdminPermissionRequest"></a>

### SetAdminPermissionRequest
SetAdminPermission 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| admin_group_id | [int32](#int32) |  |  |
| is_revoke | [bool](#bool) |  |  |






<a name="sphinx.v1.SignInfo"></a>

### SignInfo
GetSignInfo 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data_json | [string](#string) |  |  |






<a name="sphinx.v1.StatusScriptRequest"></a>

### StatusScriptRequest
StatusScript 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| review_id | [uint32](#uint32) |  |  |
| transaction_id | [string](#string) |  |  |






<a name="sphinx.v1.SuccessCode"></a>

### SuccessCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="sphinx.v1.UserAddress"></a>

### UserAddress
RegisterUser 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int32](#int32) |  |  |
| coin_id | [int32](#int32) |  |  |
| datetime | [uint64](#uint64) |  |  |
| address | [string](#string) |  |  |






<a name="sphinx.v1.WithdrawApplyRequest"></a>

### WithdrawApplyRequest
WithdrawApply 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int32](#int32) |  |  |
| coin_id | [int32](#int32) |  |  |
| amount | [double](#double) |  |  |
| datetime | [uint64](#uint64) |  |  |
| sig_salt | [uint64](#uint64) |  |  |
| sign_user | [string](#string) |  |  |
| address | [string](#string) |  |  |






<a name="sphinx.v1.WithdrawConfirmedRequest"></a>

### WithdrawConfirmedRequest
WithdrawConfirmed 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| review_id | [uint32](#uint32) |  |  |





 

 

 


<a name="sphinx.v1.Review"></a>

### Review
审核服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetList | [GetListRequest](#sphinx.v1.GetListRequest) | [ReviewList](#sphinx.v1.ReviewList) | 获取待审核列表 |
| SetAdminPermission | [SetAdminPermissionRequest](#sphinx.v1.SetAdminPermissionRequest) | [SuccessCode](#sphinx.v1.SuccessCode) | 设置管理员的审核权限 |
| ProcessReview | [ProcessReviewRequest](#sphinx.v1.ProcessReviewRequest) | [ReviewList](#sphinx.v1.ReviewList) | 审核交易 |


<a name="sphinx.v1.Trading"></a>

### Trading
交易服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterUser | [RegisterUserRequest](#sphinx.v1.RegisterUserRequest) | [UserAddress](#sphinx.v1.UserAddress) | 创建账户 |
| WithdrawApply | [WithdrawApplyRequest](#sphinx.v1.WithdrawApplyRequest) | [SuccessCode](#sphinx.v1.SuccessCode) | 用户申请提现 |
| RechargeUser | [RechargeUserRequest](#sphinx.v1.RechargeUserRequest) | [SuccessCode](#sphinx.v1.SuccessCode) | （非主动） 充值购买 - 异步扫描收款地址通知回调 |
| WithdrawConfirmed | [WithdrawConfirmedRequest](#sphinx.v1.WithdrawConfirmedRequest) | [SuccessCode](#sphinx.v1.SuccessCode) | （非主动） 确认提现 - 从消息队列中获取确认提现的数据并操作 |


<a name="sphinx.v1.WalletAgent"></a>

### WalletAgent
钱包代理

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AcceptNode | [NodeInfo](#sphinx.v1.NodeInfo) | [SuccessCode](#sphinx.v1.SuccessCode) | 接收健康报告 |


<a name="sphinx.v1.WalletNode"></a>

### WalletNode
钱包节点

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSignInfo | [GetSignInfoRequest](#sphinx.v1.GetSignInfoRequest) | [SignInfo](#sphinx.v1.SignInfo) | 获取预签名信息 |
| BroadcastScript | [BroadcastScriptRequest](#sphinx.v1.BroadcastScriptRequest) | [SuccessCode](#sphinx.v1.SuccessCode) | 广播交易 |
| StatusScript | [StatusScriptRequest](#sphinx.v1.StatusScriptRequest) | [ScriptInfo](#sphinx.v1.ScriptInfo) | 查询交易状态 |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

