# ManagedKeyDetailsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsProvidedByUser** | Pointer to **bool** |  | [optional] 
**IsUnexportable** | Pointer to **bool** |  | [optional] 
**KeyState** | Pointer to **string** | ItemState defines the different states an Item can be in | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**ManagedKeyId** | Pointer to **string** |  | [optional] 
**TargetAliasHelper** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]ManagedKeyTargetInfo**](ManagedKeyTargetInfo.md) |  | [optional] 

## Methods

### NewManagedKeyDetailsInfo

`func NewManagedKeyDetailsInfo() *ManagedKeyDetailsInfo`

NewManagedKeyDetailsInfo instantiates a new ManagedKeyDetailsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedKeyDetailsInfoWithDefaults

`func NewManagedKeyDetailsInfoWithDefaults() *ManagedKeyDetailsInfo`

NewManagedKeyDetailsInfoWithDefaults instantiates a new ManagedKeyDetailsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsProvidedByUser

`func (o *ManagedKeyDetailsInfo) GetIsProvidedByUser() bool`

GetIsProvidedByUser returns the IsProvidedByUser field if non-nil, zero value otherwise.

### GetIsProvidedByUserOk

`func (o *ManagedKeyDetailsInfo) GetIsProvidedByUserOk() (*bool, bool)`

GetIsProvidedByUserOk returns a tuple with the IsProvidedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvidedByUser

`func (o *ManagedKeyDetailsInfo) SetIsProvidedByUser(v bool)`

SetIsProvidedByUser sets IsProvidedByUser field to given value.

### HasIsProvidedByUser

`func (o *ManagedKeyDetailsInfo) HasIsProvidedByUser() bool`

HasIsProvidedByUser returns a boolean if a field has been set.

### GetIsUnexportable

`func (o *ManagedKeyDetailsInfo) GetIsUnexportable() bool`

GetIsUnexportable returns the IsUnexportable field if non-nil, zero value otherwise.

### GetIsUnexportableOk

`func (o *ManagedKeyDetailsInfo) GetIsUnexportableOk() (*bool, bool)`

GetIsUnexportableOk returns a tuple with the IsUnexportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnexportable

`func (o *ManagedKeyDetailsInfo) SetIsUnexportable(v bool)`

SetIsUnexportable sets IsUnexportable field to given value.

### HasIsUnexportable

`func (o *ManagedKeyDetailsInfo) HasIsUnexportable() bool`

HasIsUnexportable returns a boolean if a field has been set.

### GetKeyState

`func (o *ManagedKeyDetailsInfo) GetKeyState() string`

GetKeyState returns the KeyState field if non-nil, zero value otherwise.

### GetKeyStateOk

`func (o *ManagedKeyDetailsInfo) GetKeyStateOk() (*string, bool)`

GetKeyStateOk returns a tuple with the KeyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyState

`func (o *ManagedKeyDetailsInfo) SetKeyState(v string)`

SetKeyState sets KeyState field to given value.

### HasKeyState

`func (o *ManagedKeyDetailsInfo) HasKeyState() bool`

HasKeyState returns a boolean if a field has been set.

### GetKeyType

`func (o *ManagedKeyDetailsInfo) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ManagedKeyDetailsInfo) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ManagedKeyDetailsInfo) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ManagedKeyDetailsInfo) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetManagedKeyId

`func (o *ManagedKeyDetailsInfo) GetManagedKeyId() string`

GetManagedKeyId returns the ManagedKeyId field if non-nil, zero value otherwise.

### GetManagedKeyIdOk

`func (o *ManagedKeyDetailsInfo) GetManagedKeyIdOk() (*string, bool)`

GetManagedKeyIdOk returns a tuple with the ManagedKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyId

`func (o *ManagedKeyDetailsInfo) SetManagedKeyId(v string)`

SetManagedKeyId sets ManagedKeyId field to given value.

### HasManagedKeyId

`func (o *ManagedKeyDetailsInfo) HasManagedKeyId() bool`

HasManagedKeyId returns a boolean if a field has been set.

### GetTargetAliasHelper

`func (o *ManagedKeyDetailsInfo) GetTargetAliasHelper() string`

GetTargetAliasHelper returns the TargetAliasHelper field if non-nil, zero value otherwise.

### GetTargetAliasHelperOk

`func (o *ManagedKeyDetailsInfo) GetTargetAliasHelperOk() (*string, bool)`

GetTargetAliasHelperOk returns a tuple with the TargetAliasHelper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAliasHelper

`func (o *ManagedKeyDetailsInfo) SetTargetAliasHelper(v string)`

SetTargetAliasHelper sets TargetAliasHelper field to given value.

### HasTargetAliasHelper

`func (o *ManagedKeyDetailsInfo) HasTargetAliasHelper() bool`

HasTargetAliasHelper returns a boolean if a field has been set.

### GetTargets

`func (o *ManagedKeyDetailsInfo) GetTargets() []ManagedKeyTargetInfo`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ManagedKeyDetailsInfo) GetTargetsOk() (*[]ManagedKeyTargetInfo, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ManagedKeyDetailsInfo) SetTargets(v []ManagedKeyTargetInfo)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ManagedKeyDetailsInfo) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


