# PolicyOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AllowedAlgorithms** | Pointer to **[]string** |  | [optional] 
**AllowedKeyNames** | Pointer to **[]string** |  | [optional] 
**AllowedKeyTypes** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxRotationIntervalDays** | Pointer to **int32** |  | [optional] 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyOutput

`func NewPolicyOutput() *PolicyOutput`

NewPolicyOutput instantiates a new PolicyOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyOutputWithDefaults

`func NewPolicyOutputWithDefaults() *PolicyOutput`

NewPolicyOutputWithDefaults instantiates a new PolicyOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PolicyOutput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PolicyOutput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PolicyOutput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PolicyOutput) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAllowedAlgorithms

`func (o *PolicyOutput) GetAllowedAlgorithms() []string`

GetAllowedAlgorithms returns the AllowedAlgorithms field if non-nil, zero value otherwise.

### GetAllowedAlgorithmsOk

`func (o *PolicyOutput) GetAllowedAlgorithmsOk() (*[]string, bool)`

GetAllowedAlgorithmsOk returns a tuple with the AllowedAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAlgorithms

`func (o *PolicyOutput) SetAllowedAlgorithms(v []string)`

SetAllowedAlgorithms sets AllowedAlgorithms field to given value.

### HasAllowedAlgorithms

`func (o *PolicyOutput) HasAllowedAlgorithms() bool`

HasAllowedAlgorithms returns a boolean if a field has been set.

### GetAllowedKeyNames

`func (o *PolicyOutput) GetAllowedKeyNames() []string`

GetAllowedKeyNames returns the AllowedKeyNames field if non-nil, zero value otherwise.

### GetAllowedKeyNamesOk

`func (o *PolicyOutput) GetAllowedKeyNamesOk() (*[]string, bool)`

GetAllowedKeyNamesOk returns a tuple with the AllowedKeyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKeyNames

`func (o *PolicyOutput) SetAllowedKeyNames(v []string)`

SetAllowedKeyNames sets AllowedKeyNames field to given value.

### HasAllowedKeyNames

`func (o *PolicyOutput) HasAllowedKeyNames() bool`

HasAllowedKeyNames returns a boolean if a field has been set.

### GetAllowedKeyTypes

`func (o *PolicyOutput) GetAllowedKeyTypes() []string`

GetAllowedKeyTypes returns the AllowedKeyTypes field if non-nil, zero value otherwise.

### GetAllowedKeyTypesOk

`func (o *PolicyOutput) GetAllowedKeyTypesOk() (*[]string, bool)`

GetAllowedKeyTypesOk returns a tuple with the AllowedKeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKeyTypes

`func (o *PolicyOutput) SetAllowedKeyTypes(v []string)`

SetAllowedKeyTypes sets AllowedKeyTypes field to given value.

### HasAllowedKeyTypes

`func (o *PolicyOutput) HasAllowedKeyTypes() bool`

HasAllowedKeyTypes returns a boolean if a field has been set.

### GetId

`func (o *PolicyOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxRotationIntervalDays

`func (o *PolicyOutput) GetMaxRotationIntervalDays() int32`

GetMaxRotationIntervalDays returns the MaxRotationIntervalDays field if non-nil, zero value otherwise.

### GetMaxRotationIntervalDaysOk

`func (o *PolicyOutput) GetMaxRotationIntervalDaysOk() (*int32, bool)`

GetMaxRotationIntervalDaysOk returns a tuple with the MaxRotationIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRotationIntervalDays

`func (o *PolicyOutput) SetMaxRotationIntervalDays(v int32)`

SetMaxRotationIntervalDays sets MaxRotationIntervalDays field to given value.

### HasMaxRotationIntervalDays

`func (o *PolicyOutput) HasMaxRotationIntervalDays() bool`

HasMaxRotationIntervalDays returns a boolean if a field has been set.

### GetObjectTypes

`func (o *PolicyOutput) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PolicyOutput) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PolicyOutput) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PolicyOutput) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetPath

`func (o *PolicyOutput) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PolicyOutput) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PolicyOutput) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PolicyOutput) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *PolicyOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyOutput) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


