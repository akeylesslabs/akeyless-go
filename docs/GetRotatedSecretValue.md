# GetRotatedSecretValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | **string** | Secret name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | Secret version | [optional] 

## Methods

### NewGetRotatedSecretValue

`func NewGetRotatedSecretValue(names string, ) *GetRotatedSecretValue`

NewGetRotatedSecretValue instantiates a new GetRotatedSecretValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRotatedSecretValueWithDefaults

`func NewGetRotatedSecretValueWithDefaults() *GetRotatedSecretValue`

NewGetRotatedSecretValueWithDefaults instantiates a new GetRotatedSecretValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *GetRotatedSecretValue) GetNames() string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GetRotatedSecretValue) GetNamesOk() (*string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GetRotatedSecretValue) SetNames(v string)`

SetNames sets Names field to given value.


### GetToken

`func (o *GetRotatedSecretValue) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetRotatedSecretValue) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetRotatedSecretValue) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetRotatedSecretValue) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GetRotatedSecretValue) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GetRotatedSecretValue) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GetRotatedSecretValue) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GetRotatedSecretValue) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *GetRotatedSecretValue) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRotatedSecretValue) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRotatedSecretValue) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetRotatedSecretValue) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


