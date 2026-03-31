# LockTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **string** | Comma-separated blocked actions (read, update) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**LockTtl** | Pointer to **int64** | Lock TTL in minutes | [optional] 
**Name** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewLockTarget

`func NewLockTarget(name string, ) *LockTarget`

NewLockTarget instantiates a new LockTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockTargetWithDefaults

`func NewLockTargetWithDefaults() *LockTarget`

NewLockTargetWithDefaults instantiates a new LockTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *LockTarget) GetActions() string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *LockTarget) GetActionsOk() (*string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *LockTarget) SetActions(v string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *LockTarget) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetJson

`func (o *LockTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *LockTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *LockTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *LockTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLockTtl

`func (o *LockTarget) GetLockTtl() int64`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *LockTarget) GetLockTtlOk() (*int64, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *LockTarget) SetLockTtl(v int64)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *LockTarget) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetName

`func (o *LockTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LockTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LockTarget) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *LockTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LockTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LockTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LockTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *LockTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *LockTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *LockTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *LockTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


