# RotatedSecretDeleteSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Rotated secret name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | **string** | Universal Secret Connector name | 

## Methods

### NewRotatedSecretDeleteSync

`func NewRotatedSecretDeleteSync(name string, uscName string, ) *RotatedSecretDeleteSync`

NewRotatedSecretDeleteSync instantiates a new RotatedSecretDeleteSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretDeleteSyncWithDefaults

`func NewRotatedSecretDeleteSyncWithDefaults() *RotatedSecretDeleteSync`

NewRotatedSecretDeleteSyncWithDefaults instantiates a new RotatedSecretDeleteSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *RotatedSecretDeleteSync) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretDeleteSync) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretDeleteSync) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretDeleteSync) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretDeleteSync) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretDeleteSync) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretDeleteSync) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *RotatedSecretDeleteSync) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretDeleteSync) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretDeleteSync) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretDeleteSync) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretDeleteSync) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretDeleteSync) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretDeleteSync) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretDeleteSync) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *RotatedSecretDeleteSync) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *RotatedSecretDeleteSync) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *RotatedSecretDeleteSync) SetUscName(v string)`

SetUscName sets UscName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


