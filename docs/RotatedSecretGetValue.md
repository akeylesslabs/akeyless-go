# RotatedSecretGetValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Get rotated secret value of specific Host (relevant only for Linked Target) | [optional] 
**IgnoreCache** | Pointer to **string** | Retrieve the Secret value without checking the Gateway&#39;s cache [true/false]. This flag is only relevant when using the RestAPI | [optional] [default to "false"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Secret name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | Secret version | [optional] 

## Methods

### NewRotatedSecretGetValue

`func NewRotatedSecretGetValue(name string, ) *RotatedSecretGetValue`

NewRotatedSecretGetValue instantiates a new RotatedSecretGetValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretGetValueWithDefaults

`func NewRotatedSecretGetValueWithDefaults() *RotatedSecretGetValue`

NewRotatedSecretGetValueWithDefaults instantiates a new RotatedSecretGetValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RotatedSecretGetValue) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RotatedSecretGetValue) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RotatedSecretGetValue) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RotatedSecretGetValue) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIgnoreCache

`func (o *RotatedSecretGetValue) GetIgnoreCache() string`

GetIgnoreCache returns the IgnoreCache field if non-nil, zero value otherwise.

### GetIgnoreCacheOk

`func (o *RotatedSecretGetValue) GetIgnoreCacheOk() (*string, bool)`

GetIgnoreCacheOk returns a tuple with the IgnoreCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCache

`func (o *RotatedSecretGetValue) SetIgnoreCache(v string)`

SetIgnoreCache sets IgnoreCache field to given value.

### HasIgnoreCache

`func (o *RotatedSecretGetValue) HasIgnoreCache() bool`

HasIgnoreCache returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretGetValue) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretGetValue) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretGetValue) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretGetValue) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretGetValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretGetValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretGetValue) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *RotatedSecretGetValue) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretGetValue) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretGetValue) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretGetValue) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretGetValue) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretGetValue) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretGetValue) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretGetValue) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *RotatedSecretGetValue) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RotatedSecretGetValue) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RotatedSecretGetValue) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RotatedSecretGetValue) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


