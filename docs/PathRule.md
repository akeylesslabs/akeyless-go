# PathRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigners** | Pointer to [**[]RuleAssigner**](RuleAssigner.md) |  | [optional] 
**Capabilities** | Pointer to **[]string** | The approved/denied capabilities in the path | [optional] 
**IsLimitAccess** | Pointer to **bool** | flag that indicate that this rule is allowed to be access RemainingAccess of times. | [optional] 
**NumberOfAccessUsed** | Pointer to **int64** |  | [optional] 
**NumberOfAllowedAccess** | Pointer to **int64** |  | [optional] 
**Path** | Pointer to **string** | The path the rule refers to | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**Ttl** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPathRule

`func NewPathRule() *PathRule`

NewPathRule instantiates a new PathRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathRuleWithDefaults

`func NewPathRuleWithDefaults() *PathRule`

NewPathRuleWithDefaults instantiates a new PathRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigners

`func (o *PathRule) GetAssigners() []RuleAssigner`

GetAssigners returns the Assigners field if non-nil, zero value otherwise.

### GetAssignersOk

`func (o *PathRule) GetAssignersOk() (*[]RuleAssigner, bool)`

GetAssignersOk returns a tuple with the Assigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigners

`func (o *PathRule) SetAssigners(v []RuleAssigner)`

SetAssigners sets Assigners field to given value.

### HasAssigners

`func (o *PathRule) HasAssigners() bool`

HasAssigners returns a boolean if a field has been set.

### GetCapabilities

`func (o *PathRule) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PathRule) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PathRule) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *PathRule) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetIsLimitAccess

`func (o *PathRule) GetIsLimitAccess() bool`

GetIsLimitAccess returns the IsLimitAccess field if non-nil, zero value otherwise.

### GetIsLimitAccessOk

`func (o *PathRule) GetIsLimitAccessOk() (*bool, bool)`

GetIsLimitAccessOk returns a tuple with the IsLimitAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimitAccess

`func (o *PathRule) SetIsLimitAccess(v bool)`

SetIsLimitAccess sets IsLimitAccess field to given value.

### HasIsLimitAccess

`func (o *PathRule) HasIsLimitAccess() bool`

HasIsLimitAccess returns a boolean if a field has been set.

### GetNumberOfAccessUsed

`func (o *PathRule) GetNumberOfAccessUsed() int64`

GetNumberOfAccessUsed returns the NumberOfAccessUsed field if non-nil, zero value otherwise.

### GetNumberOfAccessUsedOk

`func (o *PathRule) GetNumberOfAccessUsedOk() (*int64, bool)`

GetNumberOfAccessUsedOk returns a tuple with the NumberOfAccessUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAccessUsed

`func (o *PathRule) SetNumberOfAccessUsed(v int64)`

SetNumberOfAccessUsed sets NumberOfAccessUsed field to given value.

### HasNumberOfAccessUsed

`func (o *PathRule) HasNumberOfAccessUsed() bool`

HasNumberOfAccessUsed returns a boolean if a field has been set.

### GetNumberOfAllowedAccess

`func (o *PathRule) GetNumberOfAllowedAccess() int64`

GetNumberOfAllowedAccess returns the NumberOfAllowedAccess field if non-nil, zero value otherwise.

### GetNumberOfAllowedAccessOk

`func (o *PathRule) GetNumberOfAllowedAccessOk() (*int64, bool)`

GetNumberOfAllowedAccessOk returns a tuple with the NumberOfAllowedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAllowedAccess

`func (o *PathRule) SetNumberOfAllowedAccess(v int64)`

SetNumberOfAllowedAccess sets NumberOfAllowedAccess field to given value.

### HasNumberOfAllowedAccess

`func (o *PathRule) HasNumberOfAllowedAccess() bool`

HasNumberOfAllowedAccess returns a boolean if a field has been set.

### GetPath

`func (o *PathRule) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PathRule) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PathRule) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PathRule) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStartTime

`func (o *PathRule) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PathRule) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PathRule) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *PathRule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTtl

`func (o *PathRule) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PathRule) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PathRule) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PathRule) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *PathRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PathRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PathRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PathRule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


