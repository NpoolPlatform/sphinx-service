# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/all.proto](#npool/all.proto)
    - [AccountAddress](#sphinx.v1.AccountAddress)
    - [AccountBalance](#sphinx.v1.AccountBalance)
    - [ApplyTransactionRequest](#sphinx.v1.ApplyTransactionRequest)
    - [CoinInfoList](#sphinx.v1.CoinInfoList)
    - [CoinInfoRow](#sphinx.v1.CoinInfoRow)
    - [GetBalanceRequest](#sphinx.v1.GetBalanceRequest)
    - [GetCoinInfoRequest](#sphinx.v1.GetCoinInfoRequest)
    - [GetCoinInfosRequest](#sphinx.v1.GetCoinInfosRequest)
    - [RegisterAccountRequest](#sphinx.v1.RegisterAccountRequest)
    - [SuccessCode](#sphinx.v1.SuccessCode)
  
    - [Trading](#sphinx.v1.Trading)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/all.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/all.proto



<a name="sphinx.v1.AccountAddress"></a>

### AccountAddress
RegisterAccount 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  |  |






<a name="sphinx.v1.AccountBalance"></a>

### AccountBalance
GetBalance 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  |  |
| timestamp_utc | [uint64](#uint64) |  |  |
| amount_int | [int64](#int64) |  |  |
| amount_digits | [int32](#int32) |  |  |
| amount_string | [string](#string) |  |  |






<a name="sphinx.v1.ApplyTransactionRequest"></a>

### ApplyTransactionRequest
ApplyTransaction 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| transaction_id_insite | [string](#string) |  |  |
| address_from | [string](#string) |  |  |
| address_to | [string](#string) |  |  |
| amount_int | [int64](#int64) |  |  |
| amount_digits | [int32](#int32) |  |  |






<a name="sphinx.v1.CoinInfoList"></a>

### CoinInfoList
GetCoinInfos 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list | [CoinInfoRow](#sphinx.v1.CoinInfoRow) | repeated |  |






<a name="sphinx.v1.CoinInfoRow"></a>

### CoinInfoRow
GetCoinInfo 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| need_signinfo | [bool](#bool) |  |  |
| name | [string](#string) |  |  |
| unit | [string](#string) |  |  |






<a name="sphinx.v1.GetBalanceRequest"></a>

### GetBalanceRequest
GetBalance 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  |  |
| timestamp_utc | [uint64](#uint64) |  |  |






<a name="sphinx.v1.GetCoinInfoRequest"></a>

### GetCoinInfoRequest
GetCoinInfo 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |






<a name="sphinx.v1.GetCoinInfosRequest"></a>

### GetCoinInfosRequest
GetCoinInfos 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_ids | [int32](#int32) | repeated |  |






<a name="sphinx.v1.RegisterAccountRequest"></a>

### RegisterAccountRequest
RegisterAccount 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |






<a name="sphinx.v1.SuccessCode"></a>

### SuccessCode
通用返回码


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |





 

 

 


<a name="sphinx.v1.Trading"></a>

### Trading
交易服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCoinInfo | [GetCoinInfoRequest](#sphinx.v1.GetCoinInfoRequest) | [CoinInfoRow](#sphinx.v1.CoinInfoRow) | 查询单个币种 |
| GetCoinInfos | [GetCoinInfosRequest](#sphinx.v1.GetCoinInfosRequest) | [CoinInfoList](#sphinx.v1.CoinInfoList) | 查询全部币种 |
| RegisterAccount | [RegisterAccountRequest](#sphinx.v1.RegisterAccountRequest) | [AccountAddress](#sphinx.v1.AccountAddress) | 创建账户 |
| GetBalance | [GetBalanceRequest](#sphinx.v1.GetBalanceRequest) | [AccountBalance](#sphinx.v1.AccountBalance) | 余额查询 |
| ApplyTransaction | [ApplyTransactionRequest](#sphinx.v1.ApplyTransactionRequest) | [SuccessCode](#sphinx.v1.SuccessCode) | 转账 / 提现 |

 



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

