# GetDynamicSecretValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | Optional input as &#x60;key-value&#x60; pairs | [optional] 
**Host** | Pointer to **string** | Host | [optional] 
**Name** | **string** | Dynamic secret name | 
**Target** | Pointer to **string** | Target Name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGetDynamicSecretValue

`func NewGetDynamicSecretValue(name string, ) *GetDynamicSecretValue`

NewGetDynamicSecretValue instantiates a new GetDynamicSecretValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDynamicSecretValueWithDefaults

`func NewGetDynamicSecretValueWithDefaults() *GetDynamicSecretValue`

NewGetDynamicSecretValueWithDefaults instantiates a new GetDynamicSecretValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *GetDynamicSecretValue) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *GetDynamicSecretValue) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *GetDynamicSecretValue) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *GetDynamicSecretValue) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetHost

`func (o *GetDynamicSecretValue) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetDynamicSecretValue) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetDynamicSecretValue) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetDynamicSecretValue) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *GetDynamicSecretValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDynamicSecretValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDynamicSecretValue) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *GetDynamicSecretValue) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetDynamicSecretValue) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetDynamicSecretValue) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GetDynamicSecretValue) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetToken

`func (o *GetDynamicSecretValue) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetDynamicSecretValue) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetDynamicSecretValue) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetDynamicSecretValue) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GetDynamicSecretValue) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GetDynamicSecretValue) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GetDynamicSecretValue) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GetDynamicSecretValue) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


