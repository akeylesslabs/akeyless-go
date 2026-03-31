# LockingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** |  | [optional] 
**ExpireAt** | Pointer to **int64** |  | [optional] 
**LockedBy** | Pointer to **string** |  | [optional] 
**UniqueIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewLockingInfo

`func NewLockingInfo() *LockingInfo`

NewLockingInfo instantiates a new LockingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockingInfoWithDefaults

`func NewLockingInfoWithDefaults() *LockingInfo`

NewLockingInfoWithDefaults instantiates a new LockingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *LockingInfo) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *LockingInfo) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *LockingInfo) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *LockingInfo) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetExpireAt

`func (o *LockingInfo) GetExpireAt() int64`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *LockingInfo) GetExpireAtOk() (*int64, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *LockingInfo) SetExpireAt(v int64)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *LockingInfo) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetLockedBy

`func (o *LockingInfo) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *LockingInfo) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *LockingInfo) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *LockingInfo) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *LockingInfo) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *LockingInfo) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *LockingInfo) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *LockingInfo) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


