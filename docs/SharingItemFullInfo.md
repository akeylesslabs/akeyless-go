# SharingItemFullInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigners** | Pointer to [**[]RuleAssigner**](RuleAssigner.md) |  | [optional] 
**Capabilities** | Pointer to **[]string** | The approved/denied capabilities in the path | [optional] 
**Cb** | Pointer to **int32** |  | [optional] 
**IsLimitAccess** | Pointer to **bool** | flag that indicate that this rule is allowed to be access RemainingAccess of times. | [optional] 
**ItemId** | Pointer to **int64** | The item id this rule directly refers to (when applicable) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumberOfAccessUsed** | Pointer to **int64** |  | [optional] 
**NumberOfAllowedAccess** | Pointer to **int64** |  | [optional] 
**Path** | Pointer to **string** | The path the rule refers to | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**Ttl** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSharingItemFullInfo

`func NewSharingItemFullInfo() *SharingItemFullInfo`

NewSharingItemFullInfo instantiates a new SharingItemFullInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingItemFullInfoWithDefaults

`func NewSharingItemFullInfoWithDefaults() *SharingItemFullInfo`

NewSharingItemFullInfoWithDefaults instantiates a new SharingItemFullInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigners

`func (o *SharingItemFullInfo) GetAssigners() []RuleAssigner`

GetAssigners returns the Assigners field if non-nil, zero value otherwise.

### GetAssignersOk

`func (o *SharingItemFullInfo) GetAssignersOk() (*[]RuleAssigner, bool)`

GetAssignersOk returns a tuple with the Assigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigners

`func (o *SharingItemFullInfo) SetAssigners(v []RuleAssigner)`

SetAssigners sets Assigners field to given value.

### HasAssigners

`func (o *SharingItemFullInfo) HasAssigners() bool`

HasAssigners returns a boolean if a field has been set.

### GetCapabilities

`func (o *SharingItemFullInfo) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SharingItemFullInfo) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SharingItemFullInfo) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *SharingItemFullInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCb

`func (o *SharingItemFullInfo) GetCb() int32`

GetCb returns the Cb field if non-nil, zero value otherwise.

### GetCbOk

`func (o *SharingItemFullInfo) GetCbOk() (*int32, bool)`

GetCbOk returns a tuple with the Cb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCb

`func (o *SharingItemFullInfo) SetCb(v int32)`

SetCb sets Cb field to given value.

### HasCb

`func (o *SharingItemFullInfo) HasCb() bool`

HasCb returns a boolean if a field has been set.

### GetIsLimitAccess

`func (o *SharingItemFullInfo) GetIsLimitAccess() bool`

GetIsLimitAccess returns the IsLimitAccess field if non-nil, zero value otherwise.

### GetIsLimitAccessOk

`func (o *SharingItemFullInfo) GetIsLimitAccessOk() (*bool, bool)`

GetIsLimitAccessOk returns a tuple with the IsLimitAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimitAccess

`func (o *SharingItemFullInfo) SetIsLimitAccess(v bool)`

SetIsLimitAccess sets IsLimitAccess field to given value.

### HasIsLimitAccess

`func (o *SharingItemFullInfo) HasIsLimitAccess() bool`

HasIsLimitAccess returns a boolean if a field has been set.

### GetItemId

`func (o *SharingItemFullInfo) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SharingItemFullInfo) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SharingItemFullInfo) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SharingItemFullInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetName

`func (o *SharingItemFullInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharingItemFullInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharingItemFullInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharingItemFullInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfAccessUsed

`func (o *SharingItemFullInfo) GetNumberOfAccessUsed() int64`

GetNumberOfAccessUsed returns the NumberOfAccessUsed field if non-nil, zero value otherwise.

### GetNumberOfAccessUsedOk

`func (o *SharingItemFullInfo) GetNumberOfAccessUsedOk() (*int64, bool)`

GetNumberOfAccessUsedOk returns a tuple with the NumberOfAccessUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAccessUsed

`func (o *SharingItemFullInfo) SetNumberOfAccessUsed(v int64)`

SetNumberOfAccessUsed sets NumberOfAccessUsed field to given value.

### HasNumberOfAccessUsed

`func (o *SharingItemFullInfo) HasNumberOfAccessUsed() bool`

HasNumberOfAccessUsed returns a boolean if a field has been set.

### GetNumberOfAllowedAccess

`func (o *SharingItemFullInfo) GetNumberOfAllowedAccess() int64`

GetNumberOfAllowedAccess returns the NumberOfAllowedAccess field if non-nil, zero value otherwise.

### GetNumberOfAllowedAccessOk

`func (o *SharingItemFullInfo) GetNumberOfAllowedAccessOk() (*int64, bool)`

GetNumberOfAllowedAccessOk returns a tuple with the NumberOfAllowedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAllowedAccess

`func (o *SharingItemFullInfo) SetNumberOfAllowedAccess(v int64)`

SetNumberOfAllowedAccess sets NumberOfAllowedAccess field to given value.

### HasNumberOfAllowedAccess

`func (o *SharingItemFullInfo) HasNumberOfAllowedAccess() bool`

HasNumberOfAllowedAccess returns a boolean if a field has been set.

### GetPath

`func (o *SharingItemFullInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SharingItemFullInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SharingItemFullInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SharingItemFullInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStartTime

`func (o *SharingItemFullInfo) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SharingItemFullInfo) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SharingItemFullInfo) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SharingItemFullInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTtl

`func (o *SharingItemFullInfo) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SharingItemFullInfo) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SharingItemFullInfo) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SharingItemFullInfo) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *SharingItemFullInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SharingItemFullInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SharingItemFullInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SharingItemFullInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


