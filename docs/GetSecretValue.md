# GetSecretValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | **[]string** | Secret name | 
**PrettyPrint** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | Secret version | [optional] 

## Methods

### NewGetSecretValue

`func NewGetSecretValue(names []string, ) *GetSecretValue`

NewGetSecretValue instantiates a new GetSecretValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecretValueWithDefaults

`func NewGetSecretValueWithDefaults() *GetSecretValue`

NewGetSecretValueWithDefaults instantiates a new GetSecretValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *GetSecretValue) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GetSecretValue) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GetSecretValue) SetNames(v []string)`

SetNames sets Names field to given value.


### GetPrettyPrint

`func (o *GetSecretValue) GetPrettyPrint() bool`

GetPrettyPrint returns the PrettyPrint field if non-nil, zero value otherwise.

### GetPrettyPrintOk

`func (o *GetSecretValue) GetPrettyPrintOk() (*bool, bool)`

GetPrettyPrintOk returns a tuple with the PrettyPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrettyPrint

`func (o *GetSecretValue) SetPrettyPrint(v bool)`

SetPrettyPrint sets PrettyPrint field to given value.

### HasPrettyPrint

`func (o *GetSecretValue) HasPrettyPrint() bool`

HasPrettyPrint returns a boolean if a field has been set.

### GetToken

`func (o *GetSecretValue) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetSecretValue) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetSecretValue) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetSecretValue) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GetSecretValue) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GetSecretValue) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GetSecretValue) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GetSecretValue) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *GetSecretValue) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetSecretValue) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetSecretValue) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetSecretValue) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


