# ClassicKeyDetailsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassicKeyId** | Pointer to **string** |  | [optional] 
**IsProvidedByUser** | Pointer to **bool** |  | [optional] 
**IsUnexportable** | Pointer to **bool** |  | [optional] 
**KeyState** | Pointer to **string** | ItemState defines the different states an Item can be in | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**TargetAliasHelper** | Pointer to **string** |  | [optional] 
**TargetTypes** | Pointer to **[]string** |  | [optional] 
**Targets** | Pointer to [**[]ClassicKeyTargetInfo**](ClassicKeyTargetInfo.md) |  | [optional] 

## Methods

### NewClassicKeyDetailsInfo

`func NewClassicKeyDetailsInfo() *ClassicKeyDetailsInfo`

NewClassicKeyDetailsInfo instantiates a new ClassicKeyDetailsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassicKeyDetailsInfoWithDefaults

`func NewClassicKeyDetailsInfoWithDefaults() *ClassicKeyDetailsInfo`

NewClassicKeyDetailsInfoWithDefaults instantiates a new ClassicKeyDetailsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassicKeyId

`func (o *ClassicKeyDetailsInfo) GetClassicKeyId() string`

GetClassicKeyId returns the ClassicKeyId field if non-nil, zero value otherwise.

### GetClassicKeyIdOk

`func (o *ClassicKeyDetailsInfo) GetClassicKeyIdOk() (*string, bool)`

GetClassicKeyIdOk returns a tuple with the ClassicKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicKeyId

`func (o *ClassicKeyDetailsInfo) SetClassicKeyId(v string)`

SetClassicKeyId sets ClassicKeyId field to given value.

### HasClassicKeyId

`func (o *ClassicKeyDetailsInfo) HasClassicKeyId() bool`

HasClassicKeyId returns a boolean if a field has been set.

### GetIsProvidedByUser

`func (o *ClassicKeyDetailsInfo) GetIsProvidedByUser() bool`

GetIsProvidedByUser returns the IsProvidedByUser field if non-nil, zero value otherwise.

### GetIsProvidedByUserOk

`func (o *ClassicKeyDetailsInfo) GetIsProvidedByUserOk() (*bool, bool)`

GetIsProvidedByUserOk returns a tuple with the IsProvidedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvidedByUser

`func (o *ClassicKeyDetailsInfo) SetIsProvidedByUser(v bool)`

SetIsProvidedByUser sets IsProvidedByUser field to given value.

### HasIsProvidedByUser

`func (o *ClassicKeyDetailsInfo) HasIsProvidedByUser() bool`

HasIsProvidedByUser returns a boolean if a field has been set.

### GetIsUnexportable

`func (o *ClassicKeyDetailsInfo) GetIsUnexportable() bool`

GetIsUnexportable returns the IsUnexportable field if non-nil, zero value otherwise.

### GetIsUnexportableOk

`func (o *ClassicKeyDetailsInfo) GetIsUnexportableOk() (*bool, bool)`

GetIsUnexportableOk returns a tuple with the IsUnexportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnexportable

`func (o *ClassicKeyDetailsInfo) SetIsUnexportable(v bool)`

SetIsUnexportable sets IsUnexportable field to given value.

### HasIsUnexportable

`func (o *ClassicKeyDetailsInfo) HasIsUnexportable() bool`

HasIsUnexportable returns a boolean if a field has been set.

### GetKeyState

`func (o *ClassicKeyDetailsInfo) GetKeyState() string`

GetKeyState returns the KeyState field if non-nil, zero value otherwise.

### GetKeyStateOk

`func (o *ClassicKeyDetailsInfo) GetKeyStateOk() (*string, bool)`

GetKeyStateOk returns a tuple with the KeyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyState

`func (o *ClassicKeyDetailsInfo) SetKeyState(v string)`

SetKeyState sets KeyState field to given value.

### HasKeyState

`func (o *ClassicKeyDetailsInfo) HasKeyState() bool`

HasKeyState returns a boolean if a field has been set.

### GetKeyType

`func (o *ClassicKeyDetailsInfo) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ClassicKeyDetailsInfo) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ClassicKeyDetailsInfo) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ClassicKeyDetailsInfo) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetLastError

`func (o *ClassicKeyDetailsInfo) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ClassicKeyDetailsInfo) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ClassicKeyDetailsInfo) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ClassicKeyDetailsInfo) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetTargetAliasHelper

`func (o *ClassicKeyDetailsInfo) GetTargetAliasHelper() string`

GetTargetAliasHelper returns the TargetAliasHelper field if non-nil, zero value otherwise.

### GetTargetAliasHelperOk

`func (o *ClassicKeyDetailsInfo) GetTargetAliasHelperOk() (*string, bool)`

GetTargetAliasHelperOk returns a tuple with the TargetAliasHelper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAliasHelper

`func (o *ClassicKeyDetailsInfo) SetTargetAliasHelper(v string)`

SetTargetAliasHelper sets TargetAliasHelper field to given value.

### HasTargetAliasHelper

`func (o *ClassicKeyDetailsInfo) HasTargetAliasHelper() bool`

HasTargetAliasHelper returns a boolean if a field has been set.

### GetTargetTypes

`func (o *ClassicKeyDetailsInfo) GetTargetTypes() []string`

GetTargetTypes returns the TargetTypes field if non-nil, zero value otherwise.

### GetTargetTypesOk

`func (o *ClassicKeyDetailsInfo) GetTargetTypesOk() (*[]string, bool)`

GetTargetTypesOk returns a tuple with the TargetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTypes

`func (o *ClassicKeyDetailsInfo) SetTargetTypes(v []string)`

SetTargetTypes sets TargetTypes field to given value.

### HasTargetTypes

`func (o *ClassicKeyDetailsInfo) HasTargetTypes() bool`

HasTargetTypes returns a boolean if a field has been set.

### GetTargets

`func (o *ClassicKeyDetailsInfo) GetTargets() []ClassicKeyTargetInfo`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ClassicKeyDetailsInfo) GetTargetsOk() (*[]ClassicKeyTargetInfo, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ClassicKeyDetailsInfo) SetTargets(v []ClassicKeyTargetInfo)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ClassicKeyDetailsInfo) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


