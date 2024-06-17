# UscUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryValue** | Pointer to **bool** | Use this option if the universal secrets value is a base64 encoded binary | [optional] 
**Description** | Pointer to **string** | Description of the universal secrets | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Namespace** | Pointer to **string** | The namespace (relevant for Hashi vault target) | [optional] 
**SecretId** | **string** | The universal secrets id (or name, for AWS, Azure, K8s or Hashi vault targets) to update | 
**Tags** | Pointer to **map[string]string** | Tags for the universal secrets | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | **string** | Name of the Universal Secrets Connector item | 
**Value** | **string** | Value of the universal secrets item, either text or base64 encoded binary | 

## Methods

### NewUscUpdate

`func NewUscUpdate(secretId string, uscName string, value string, ) *UscUpdate`

NewUscUpdate instantiates a new UscUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUscUpdateWithDefaults

`func NewUscUpdateWithDefaults() *UscUpdate`

NewUscUpdateWithDefaults instantiates a new UscUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryValue

`func (o *UscUpdate) GetBinaryValue() bool`

GetBinaryValue returns the BinaryValue field if non-nil, zero value otherwise.

### GetBinaryValueOk

`func (o *UscUpdate) GetBinaryValueOk() (*bool, bool)`

GetBinaryValueOk returns a tuple with the BinaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryValue

`func (o *UscUpdate) SetBinaryValue(v bool)`

SetBinaryValue sets BinaryValue field to given value.

### HasBinaryValue

`func (o *UscUpdate) HasBinaryValue() bool`

HasBinaryValue returns a boolean if a field has been set.

### GetDescription

`func (o *UscUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UscUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UscUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UscUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *UscUpdate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UscUpdate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UscUpdate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UscUpdate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetNamespace

`func (o *UscUpdate) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UscUpdate) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UscUpdate) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UscUpdate) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSecretId

`func (o *UscUpdate) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *UscUpdate) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *UscUpdate) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetTags

`func (o *UscUpdate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UscUpdate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UscUpdate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UscUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *UscUpdate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UscUpdate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UscUpdate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UscUpdate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UscUpdate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UscUpdate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UscUpdate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UscUpdate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *UscUpdate) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *UscUpdate) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *UscUpdate) SetUscName(v string)`

SetUscName sets UscName field to given value.


### GetValue

`func (o *UscUpdate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UscUpdate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UscUpdate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


