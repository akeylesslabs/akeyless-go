# GatewayRevokeTmpUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Deprecated: has no effect | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | **string** | Producer Name | 
**RevokeAll** | Pointer to **bool** | Revoke All Temp Creds | [optional] 
**SoftDelete** | Pointer to **bool** | Soft Delete | [optional] 
**TmpCredsId** | **string** | Tmp Creds ID | [default to "demo_default_tmp_creds_id_for_sdk_bc"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayRevokeTmpUsers

`func NewGatewayRevokeTmpUsers(name string, tmpCredsId string, ) *GatewayRevokeTmpUsers`

NewGatewayRevokeTmpUsers instantiates a new GatewayRevokeTmpUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRevokeTmpUsersWithDefaults

`func NewGatewayRevokeTmpUsersWithDefaults() *GatewayRevokeTmpUsers`

NewGatewayRevokeTmpUsersWithDefaults instantiates a new GatewayRevokeTmpUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *GatewayRevokeTmpUsers) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GatewayRevokeTmpUsers) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GatewayRevokeTmpUsers) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GatewayRevokeTmpUsers) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *GatewayRevokeTmpUsers) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayRevokeTmpUsers) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayRevokeTmpUsers) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayRevokeTmpUsers) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayRevokeTmpUsers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayRevokeTmpUsers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayRevokeTmpUsers) SetName(v string)`

SetName sets Name field to given value.


### GetRevokeAll

`func (o *GatewayRevokeTmpUsers) GetRevokeAll() bool`

GetRevokeAll returns the RevokeAll field if non-nil, zero value otherwise.

### GetRevokeAllOk

`func (o *GatewayRevokeTmpUsers) GetRevokeAllOk() (*bool, bool)`

GetRevokeAllOk returns a tuple with the RevokeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeAll

`func (o *GatewayRevokeTmpUsers) SetRevokeAll(v bool)`

SetRevokeAll sets RevokeAll field to given value.

### HasRevokeAll

`func (o *GatewayRevokeTmpUsers) HasRevokeAll() bool`

HasRevokeAll returns a boolean if a field has been set.

### GetSoftDelete

`func (o *GatewayRevokeTmpUsers) GetSoftDelete() bool`

GetSoftDelete returns the SoftDelete field if non-nil, zero value otherwise.

### GetSoftDeleteOk

`func (o *GatewayRevokeTmpUsers) GetSoftDeleteOk() (*bool, bool)`

GetSoftDeleteOk returns a tuple with the SoftDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDelete

`func (o *GatewayRevokeTmpUsers) SetSoftDelete(v bool)`

SetSoftDelete sets SoftDelete field to given value.

### HasSoftDelete

`func (o *GatewayRevokeTmpUsers) HasSoftDelete() bool`

HasSoftDelete returns a boolean if a field has been set.

### GetTmpCredsId

`func (o *GatewayRevokeTmpUsers) GetTmpCredsId() string`

GetTmpCredsId returns the TmpCredsId field if non-nil, zero value otherwise.

### GetTmpCredsIdOk

`func (o *GatewayRevokeTmpUsers) GetTmpCredsIdOk() (*string, bool)`

GetTmpCredsIdOk returns a tuple with the TmpCredsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpCredsId

`func (o *GatewayRevokeTmpUsers) SetTmpCredsId(v string)`

SetTmpCredsId sets TmpCredsId field to given value.


### GetToken

`func (o *GatewayRevokeTmpUsers) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayRevokeTmpUsers) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayRevokeTmpUsers) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayRevokeTmpUsers) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayRevokeTmpUsers) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayRevokeTmpUsers) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayRevokeTmpUsers) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayRevokeTmpUsers) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


