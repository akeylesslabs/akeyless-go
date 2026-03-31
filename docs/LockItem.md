# LockItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **string** | Comma-separated blocked actions (read, update) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**LockTtl** | Pointer to **int64** | Lock TTL in minutes | [optional] 
**Name** | **string** | Item name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewLockItem

`func NewLockItem(name string, ) *LockItem`

NewLockItem instantiates a new LockItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockItemWithDefaults

`func NewLockItemWithDefaults() *LockItem`

NewLockItemWithDefaults instantiates a new LockItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *LockItem) GetActions() string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *LockItem) GetActionsOk() (*string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *LockItem) SetActions(v string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *LockItem) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetJson

`func (o *LockItem) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *LockItem) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *LockItem) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *LockItem) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLockTtl

`func (o *LockItem) GetLockTtl() int64`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *LockItem) GetLockTtlOk() (*int64, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *LockItem) SetLockTtl(v int64)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *LockItem) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetName

`func (o *LockItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LockItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LockItem) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *LockItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LockItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LockItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LockItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *LockItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *LockItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *LockItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *LockItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


