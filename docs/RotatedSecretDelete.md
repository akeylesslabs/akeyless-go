# RotatedSecretDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Rotated secret name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | The specific version you want to delete, -1&#x3D;entire item with all versions (default) | [optional] [default to -1]

## Methods

### NewRotatedSecretDelete

`func NewRotatedSecretDelete(name string, ) *RotatedSecretDelete`

NewRotatedSecretDelete instantiates a new RotatedSecretDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretDeleteWithDefaults

`func NewRotatedSecretDeleteWithDefaults() *RotatedSecretDelete`

NewRotatedSecretDeleteWithDefaults instantiates a new RotatedSecretDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *RotatedSecretDelete) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretDelete) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretDelete) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretDelete) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretDelete) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretDelete) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretDelete) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *RotatedSecretDelete) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretDelete) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretDelete) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretDelete) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretDelete) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretDelete) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretDelete) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretDelete) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *RotatedSecretDelete) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RotatedSecretDelete) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RotatedSecretDelete) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RotatedSecretDelete) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


