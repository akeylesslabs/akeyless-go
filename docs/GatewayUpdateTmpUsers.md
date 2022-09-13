# GatewayUpdateTmpUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | **string** | Producer Name | 
**NewTtlMin** | **int64** | New TTL in Minutes | 
**TmpCredsId** | **string** | Tmp Creds ID | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateTmpUsers

`func NewGatewayUpdateTmpUsers(name string, newTtlMin int64, tmpCredsId string, ) *GatewayUpdateTmpUsers`

NewGatewayUpdateTmpUsers instantiates a new GatewayUpdateTmpUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateTmpUsersWithDefaults

`func NewGatewayUpdateTmpUsersWithDefaults() *GatewayUpdateTmpUsers`

NewGatewayUpdateTmpUsersWithDefaults instantiates a new GatewayUpdateTmpUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *GatewayUpdateTmpUsers) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateTmpUsers) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateTmpUsers) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateTmpUsers) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateTmpUsers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateTmpUsers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateTmpUsers) SetName(v string)`

SetName sets Name field to given value.


### GetNewTtlMin

`func (o *GatewayUpdateTmpUsers) GetNewTtlMin() int64`

GetNewTtlMin returns the NewTtlMin field if non-nil, zero value otherwise.

### GetNewTtlMinOk

`func (o *GatewayUpdateTmpUsers) GetNewTtlMinOk() (*int64, bool)`

GetNewTtlMinOk returns a tuple with the NewTtlMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTtlMin

`func (o *GatewayUpdateTmpUsers) SetNewTtlMin(v int64)`

SetNewTtlMin sets NewTtlMin field to given value.


### GetTmpCredsId

`func (o *GatewayUpdateTmpUsers) GetTmpCredsId() string`

GetTmpCredsId returns the TmpCredsId field if non-nil, zero value otherwise.

### GetTmpCredsIdOk

`func (o *GatewayUpdateTmpUsers) GetTmpCredsIdOk() (*string, bool)`

GetTmpCredsIdOk returns a tuple with the TmpCredsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpCredsId

`func (o *GatewayUpdateTmpUsers) SetTmpCredsId(v string)`

SetTmpCredsId sets TmpCredsId field to given value.


### GetToken

`func (o *GatewayUpdateTmpUsers) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateTmpUsers) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateTmpUsers) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateTmpUsers) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateTmpUsers) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateTmpUsers) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateTmpUsers) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateTmpUsers) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


