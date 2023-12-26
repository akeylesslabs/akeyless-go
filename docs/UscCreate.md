# UscCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryValue** | Pointer to **bool** | Use this option if the universal secrets value is a base64 encoded binary | [optional] 
**Description** | Pointer to **string** | Description of the universal secrets | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**SecretName** | **string** | Name for the new universal secrets | 
**Tags** | Pointer to **map[string]string** | Tags for the universal secrets | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | **string** | Name of the Universal Secrets Connector item | 
**Value** | **string** | Value of the universal secrets item, either text or base64 encoded binary | 

## Methods

### NewUscCreate

`func NewUscCreate(secretName string, uscName string, value string, ) *UscCreate`

NewUscCreate instantiates a new UscCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUscCreateWithDefaults

`func NewUscCreateWithDefaults() *UscCreate`

NewUscCreateWithDefaults instantiates a new UscCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryValue

`func (o *UscCreate) GetBinaryValue() bool`

GetBinaryValue returns the BinaryValue field if non-nil, zero value otherwise.

### GetBinaryValueOk

`func (o *UscCreate) GetBinaryValueOk() (*bool, bool)`

GetBinaryValueOk returns a tuple with the BinaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryValue

`func (o *UscCreate) SetBinaryValue(v bool)`

SetBinaryValue sets BinaryValue field to given value.

### HasBinaryValue

`func (o *UscCreate) HasBinaryValue() bool`

HasBinaryValue returns a boolean if a field has been set.

### GetDescription

`func (o *UscCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UscCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UscCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UscCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *UscCreate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UscCreate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UscCreate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UscCreate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSecretName

`func (o *UscCreate) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *UscCreate) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *UscCreate) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetTags

`func (o *UscCreate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UscCreate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UscCreate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UscCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *UscCreate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UscCreate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UscCreate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UscCreate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UscCreate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UscCreate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UscCreate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UscCreate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *UscCreate) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *UscCreate) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *UscCreate) SetUscName(v string)`

SetUscName sets UscName field to given value.


### GetValue

`func (o *UscCreate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UscCreate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UscCreate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


