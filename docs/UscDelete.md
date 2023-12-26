# UscDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**SecretId** | **string** | The universal secrets id (or name, for AWS, Azure or K8s targets) to delete | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | **string** | Name of the Universal Secrets Connector item | 

## Methods

### NewUscDelete

`func NewUscDelete(secretId string, uscName string, ) *UscDelete`

NewUscDelete instantiates a new UscDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUscDeleteWithDefaults

`func NewUscDeleteWithDefaults() *UscDelete`

NewUscDeleteWithDefaults instantiates a new UscDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *UscDelete) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UscDelete) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UscDelete) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UscDelete) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSecretId

`func (o *UscDelete) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *UscDelete) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *UscDelete) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetToken

`func (o *UscDelete) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UscDelete) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UscDelete) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UscDelete) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UscDelete) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UscDelete) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UscDelete) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UscDelete) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *UscDelete) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *UscDelete) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *UscDelete) SetUscName(v string)`

SetUscName sets UscName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


