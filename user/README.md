# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/user/user.proto](#proto/user/user.proto)
    - [AddRequest](#.AddRequest)
    - [AddResponse](#.AddResponse)
    - [BaseUser](#.BaseUser)
    - [CreateTokenRequest](#.CreateTokenRequest)
    - [CreateTokenResponse](#.CreateTokenResponse)
    - [GetInfoByUniqueIdRequest](#.GetInfoByUniqueIdRequest)
    - [GetInfoByUniqueIdResponse](#.GetInfoByUniqueIdResponse)
    - [GetInfoByUserIdRequest](#.GetInfoByUserIdRequest)
    - [GetInfoByUserIdResponse](#.GetInfoByUserIdResponse)
    - [GetListByUserIdRequest](#.GetListByUserIdRequest)
    - [GetListByUserIdResponse](#.GetListByUserIdResponse)
    - [GetUserInfoByTokenRequest](#.GetUserInfoByTokenRequest)
    - [GetUserInfoByTokenResponse](#.GetUserInfoByTokenResponse)
    - [UserInfo](#.UserInfo)
  
    - [User](#.User)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/user/user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/user/user.proto



<a name=".AddRequest"></a>

### AddRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unique_id | [string](#string) |  | 唯一标识符 unique_id |
| nickname | [string](#string) |  | 昵称 |
| avatar | [string](#string) |  | 头像 |






<a name=".AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  | 用户ID |






<a name=".BaseUser"></a>

### BaseUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  | 用户ID |
| nickname | [string](#string) |  | 昵称 |
| avatar | [string](#string) |  | 头像 |






<a name=".CreateTokenRequest"></a>

### CreateTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  | 用户ID |






<a name=".CreateTokenResponse"></a>

### CreateTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | token令牌 |
| expired_at | [string](#string) |  | 过期时间 |






<a name=".GetInfoByUniqueIdRequest"></a>

### GetInfoByUniqueIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unique_id | [string](#string) |  | 唯一标识符 unique_id |






<a name=".GetInfoByUniqueIdResponse"></a>

### GetInfoByUniqueIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [UserInfo](#UserInfo) |  | 用户详情 |






<a name=".GetInfoByUserIdRequest"></a>

### GetInfoByUserIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  | 用户ID |






<a name=".GetInfoByUserIdResponse"></a>

### GetInfoByUserIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [UserInfo](#UserInfo) |  | 用户详情 |






<a name=".GetListByUserIdRequest"></a>

### GetListByUserIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_ids | [int64](#int64) | repeated | 用户ID数组 |






<a name=".GetListByUserIdResponse"></a>

### GetListByUserIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [BaseUser](#BaseUser) | repeated | 用户详情数组 |






<a name=".GetUserInfoByTokenRequest"></a>

### GetUserInfoByTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | token令牌 |






<a name=".GetUserInfoByTokenResponse"></a>

### GetUserInfoByTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [UserInfo](#UserInfo) |  | 用户详情 |






<a name=".UserInfo"></a>

### UserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  | 用户ID |
| unique_id | [string](#string) |  | 唯一标识符 unique_id |
| nickname | [string](#string) |  | 昵称 |
| avatar | [string](#string) |  | 头像 |





 

 

 


<a name=".User"></a>

### User


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetInfoByUniqueId | [.GetInfoByUniqueIdRequest](#GetInfoByUniqueIdRequest) | [.GetInfoByUniqueIdResponse](#GetInfoByUniqueIdResponse) | 根据 unique_id 获取用户详情 |
| GetInfoByUserId | [.GetInfoByUserIdRequest](#GetInfoByUserIdRequest) | [.GetInfoByUserIdResponse](#GetInfoByUserIdResponse) | 根据 user_id 获取用户详情 |
| GetListByUserId | [.GetListByUserIdRequest](#GetListByUserIdRequest) | [.GetListByUserIdResponse](#GetListByUserIdResponse) | 批量获取用户信息 |
| Add | [.AddRequest](#AddRequest) | [.AddResponse](#AddResponse) | 新增用户 |
| CreateToken | [.CreateTokenRequest](#CreateTokenRequest) | [.CreateTokenResponse](#CreateTokenResponse) | 获取用户token令牌 |
| GetUserInfoByToken | [.GetUserInfoByTokenRequest](#GetUserInfoByTokenRequest) | [.GetUserInfoByTokenResponse](#GetUserInfoByTokenResponse) | 根据 token 获取用户详情 |

 



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

