# EsmUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryValue** | Pointer to **bool** | Use this option if the external secret value is a base64 encoded binary | [optional] 
**Description** | Pointer to **string** | Description of the external secret | [optional] 
**EsmName** | **string** | Name of the External Secrets Manager item | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**SecretId** | **string** | The external secret id (or name, for AWS, Azure or K8s targets) to update | 
**Tags** | Pointer to **map[string]string** | Tags for the external secret | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Value** | **string** | Value of the external secret item, either text or base64 encoded binary | 

## Methods

### NewEsmUpdate

`func NewEsmUpdate(esmName string, secretId string, value string, ) *EsmUpdate`

NewEsmUpdate instantiates a new EsmUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsmUpdateWithDefaults

`func NewEsmUpdateWithDefaults() *EsmUpdate`

NewEsmUpdateWithDefaults instantiates a new EsmUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryValue

`func (o *EsmUpdate) GetBinaryValue() bool`

GetBinaryValue returns the BinaryValue field if non-nil, zero value otherwise.

### GetBinaryValueOk

`func (o *EsmUpdate) GetBinaryValueOk() (*bool, bool)`

GetBinaryValueOk returns a tuple with the BinaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryValue

`func (o *EsmUpdate) SetBinaryValue(v bool)`

SetBinaryValue sets BinaryValue field to given value.

### HasBinaryValue

`func (o *EsmUpdate) HasBinaryValue() bool`

HasBinaryValue returns a boolean if a field has been set.

### GetDescription

`func (o *EsmUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsmUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsmUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsmUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEsmName

`func (o *EsmUpdate) GetEsmName() string`

GetEsmName returns the EsmName field if non-nil, zero value otherwise.

### GetEsmNameOk

`func (o *EsmUpdate) GetEsmNameOk() (*string, bool)`

GetEsmNameOk returns a tuple with the EsmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsmName

`func (o *EsmUpdate) SetEsmName(v string)`

SetEsmName sets EsmName field to given value.


### GetJson

`func (o *EsmUpdate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EsmUpdate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EsmUpdate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EsmUpdate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSecretId

`func (o *EsmUpdate) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *EsmUpdate) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *EsmUpdate) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetTags

`func (o *EsmUpdate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EsmUpdate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EsmUpdate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EsmUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *EsmUpdate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EsmUpdate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EsmUpdate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EsmUpdate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EsmUpdate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EsmUpdate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EsmUpdate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EsmUpdate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetValue

`func (o *EsmUpdate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EsmUpdate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EsmUpdate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


