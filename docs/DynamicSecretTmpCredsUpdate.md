# DynamicSecretTmpCredsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Host | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewTtlMin** | **int64** | New TTL in Minutes | 
**TmpCredsId** | **string** | Tmp Creds ID | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDynamicSecretTmpCredsUpdate

`func NewDynamicSecretTmpCredsUpdate(host string, name string, newTtlMin int64, tmpCredsId string, ) *DynamicSecretTmpCredsUpdate`

NewDynamicSecretTmpCredsUpdate instantiates a new DynamicSecretTmpCredsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretTmpCredsUpdateWithDefaults

`func NewDynamicSecretTmpCredsUpdateWithDefaults() *DynamicSecretTmpCredsUpdate`

NewDynamicSecretTmpCredsUpdateWithDefaults instantiates a new DynamicSecretTmpCredsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *DynamicSecretTmpCredsUpdate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DynamicSecretTmpCredsUpdate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DynamicSecretTmpCredsUpdate) SetHost(v string)`

SetHost sets Host field to given value.


### GetJson

`func (o *DynamicSecretTmpCredsUpdate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretTmpCredsUpdate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretTmpCredsUpdate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretTmpCredsUpdate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretTmpCredsUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretTmpCredsUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretTmpCredsUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetNewTtlMin

`func (o *DynamicSecretTmpCredsUpdate) GetNewTtlMin() int64`

GetNewTtlMin returns the NewTtlMin field if non-nil, zero value otherwise.

### GetNewTtlMinOk

`func (o *DynamicSecretTmpCredsUpdate) GetNewTtlMinOk() (*int64, bool)`

GetNewTtlMinOk returns a tuple with the NewTtlMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTtlMin

`func (o *DynamicSecretTmpCredsUpdate) SetNewTtlMin(v int64)`

SetNewTtlMin sets NewTtlMin field to given value.


### GetTmpCredsId

`func (o *DynamicSecretTmpCredsUpdate) GetTmpCredsId() string`

GetTmpCredsId returns the TmpCredsId field if non-nil, zero value otherwise.

### GetTmpCredsIdOk

`func (o *DynamicSecretTmpCredsUpdate) GetTmpCredsIdOk() (*string, bool)`

GetTmpCredsIdOk returns a tuple with the TmpCredsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpCredsId

`func (o *DynamicSecretTmpCredsUpdate) SetTmpCredsId(v string)`

SetTmpCredsId sets TmpCredsId field to given value.


### GetToken

`func (o *DynamicSecretTmpCredsUpdate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretTmpCredsUpdate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretTmpCredsUpdate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretTmpCredsUpdate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretTmpCredsUpdate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretTmpCredsUpdate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretTmpCredsUpdate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretTmpCredsUpdate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


