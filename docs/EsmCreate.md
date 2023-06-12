# EsmCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryValue** | Pointer to **bool** | Use this option if the external secret value is a base64 encoded binary | [optional] 
**Description** | Pointer to **string** | Description of the external secret | [optional] 
**EsmName** | **string** | Name of the External Secrets Manager item | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**SecretName** | **string** | Name for the new external secret | 
**Tags** | Pointer to **map[string]string** | Tags for the external secret | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Value** | **string** | Value of the external secret item, either text or base64 encoded binary | 

## Methods

### NewEsmCreate

`func NewEsmCreate(esmName string, secretName string, value string, ) *EsmCreate`

NewEsmCreate instantiates a new EsmCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsmCreateWithDefaults

`func NewEsmCreateWithDefaults() *EsmCreate`

NewEsmCreateWithDefaults instantiates a new EsmCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryValue

`func (o *EsmCreate) GetBinaryValue() bool`

GetBinaryValue returns the BinaryValue field if non-nil, zero value otherwise.

### GetBinaryValueOk

`func (o *EsmCreate) GetBinaryValueOk() (*bool, bool)`

GetBinaryValueOk returns a tuple with the BinaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryValue

`func (o *EsmCreate) SetBinaryValue(v bool)`

SetBinaryValue sets BinaryValue field to given value.

### HasBinaryValue

`func (o *EsmCreate) HasBinaryValue() bool`

HasBinaryValue returns a boolean if a field has been set.

### GetDescription

`func (o *EsmCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsmCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsmCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsmCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEsmName

`func (o *EsmCreate) GetEsmName() string`

GetEsmName returns the EsmName field if non-nil, zero value otherwise.

### GetEsmNameOk

`func (o *EsmCreate) GetEsmNameOk() (*string, bool)`

GetEsmNameOk returns a tuple with the EsmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsmName

`func (o *EsmCreate) SetEsmName(v string)`

SetEsmName sets EsmName field to given value.


### GetJson

`func (o *EsmCreate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EsmCreate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EsmCreate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EsmCreate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSecretName

`func (o *EsmCreate) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *EsmCreate) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *EsmCreate) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetTags

`func (o *EsmCreate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EsmCreate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EsmCreate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EsmCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *EsmCreate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EsmCreate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EsmCreate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EsmCreate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EsmCreate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EsmCreate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EsmCreate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EsmCreate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetValue

`func (o *EsmCreate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EsmCreate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EsmCreate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


