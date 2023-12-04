# RotateSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RotateAllServicesBoolean** | Pointer to **bool** |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Secret name (Rotated Secret or Custom Dynamic Secret) | 
**RotateAllServices** | Pointer to **string** | Rotate all associated services | [optional] [default to "false"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRotateSecret

`func NewRotateSecret(name string, ) *RotateSecret`

NewRotateSecret instantiates a new RotateSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateSecretWithDefaults

`func NewRotateSecretWithDefaults() *RotateSecret`

NewRotateSecretWithDefaults instantiates a new RotateSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRotateAllServicesBoolean

`func (o *RotateSecret) GetRotateAllServicesBoolean() bool`

GetRotateAllServicesBoolean returns the RotateAllServicesBoolean field if non-nil, zero value otherwise.

### GetRotateAllServicesBooleanOk

`func (o *RotateSecret) GetRotateAllServicesBooleanOk() (*bool, bool)`

GetRotateAllServicesBooleanOk returns a tuple with the RotateAllServicesBoolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAllServicesBoolean

`func (o *RotateSecret) SetRotateAllServicesBoolean(v bool)`

SetRotateAllServicesBoolean sets RotateAllServicesBoolean field to given value.

### HasRotateAllServicesBoolean

`func (o *RotateSecret) HasRotateAllServicesBoolean() bool`

HasRotateAllServicesBoolean returns a boolean if a field has been set.

### GetJson

`func (o *RotateSecret) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotateSecret) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotateSecret) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotateSecret) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RotateSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotateSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotateSecret) SetName(v string)`

SetName sets Name field to given value.


### GetRotateAllServices

`func (o *RotateSecret) GetRotateAllServices() string`

GetRotateAllServices returns the RotateAllServices field if non-nil, zero value otherwise.

### GetRotateAllServicesOk

`func (o *RotateSecret) GetRotateAllServicesOk() (*string, bool)`

GetRotateAllServicesOk returns a tuple with the RotateAllServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAllServices

`func (o *RotateSecret) SetRotateAllServices(v string)`

SetRotateAllServices sets RotateAllServices field to given value.

### HasRotateAllServices

`func (o *RotateSecret) HasRotateAllServices() bool`

HasRotateAllServices returns a boolean if a field has been set.

### GetToken

`func (o *RotateSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotateSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotateSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotateSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotateSecret) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotateSecret) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotateSecret) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotateSecret) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


