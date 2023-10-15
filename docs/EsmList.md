# EsmList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsmName** | **string** | Name of the External Secrets Manager item | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewEsmList

`func NewEsmList(esmName string, ) *EsmList`

NewEsmList instantiates a new EsmList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsmListWithDefaults

`func NewEsmListWithDefaults() *EsmList`

NewEsmListWithDefaults instantiates a new EsmList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsmName

`func (o *EsmList) GetEsmName() string`

GetEsmName returns the EsmName field if non-nil, zero value otherwise.

### GetEsmNameOk

`func (o *EsmList) GetEsmNameOk() (*string, bool)`

GetEsmNameOk returns a tuple with the EsmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsmName

`func (o *EsmList) SetEsmName(v string)`

SetEsmName sets EsmName field to given value.


### GetJson

`func (o *EsmList) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EsmList) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EsmList) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EsmList) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *EsmList) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EsmList) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EsmList) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EsmList) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EsmList) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EsmList) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EsmList) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EsmList) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


