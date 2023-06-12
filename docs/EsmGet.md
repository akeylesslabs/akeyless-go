# EsmGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsmName** | **string** | Name of the External Secrets Manager item | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**SecretId** | **string** | The secret id (or name, for AWS, Azure or K8s targets) to get from the External Secrets Manager | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewEsmGet

`func NewEsmGet(esmName string, secretId string, ) *EsmGet`

NewEsmGet instantiates a new EsmGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsmGetWithDefaults

`func NewEsmGetWithDefaults() *EsmGet`

NewEsmGetWithDefaults instantiates a new EsmGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsmName

`func (o *EsmGet) GetEsmName() string`

GetEsmName returns the EsmName field if non-nil, zero value otherwise.

### GetEsmNameOk

`func (o *EsmGet) GetEsmNameOk() (*string, bool)`

GetEsmNameOk returns a tuple with the EsmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsmName

`func (o *EsmGet) SetEsmName(v string)`

SetEsmName sets EsmName field to given value.


### GetJson

`func (o *EsmGet) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *EsmGet) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *EsmGet) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *EsmGet) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSecretId

`func (o *EsmGet) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *EsmGet) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *EsmGet) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetToken

`func (o *EsmGet) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EsmGet) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EsmGet) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EsmGet) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *EsmGet) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *EsmGet) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *EsmGet) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *EsmGet) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


