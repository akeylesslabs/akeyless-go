# RotatedSecretDetailsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletePreviousVersionInDays** | Pointer to **int32** |  | [optional] 
**GwClusterId** | Pointer to **int64** |  | [optional] 
**LastRotationError** | Pointer to **string** |  | [optional] 
**NumberOfVersionsToSave** | Pointer to **int32** |  | [optional] 
**RotationHour** | Pointer to **int32** |  | [optional] 
**RotationIntervalMin** | Pointer to **bool** |  | [optional] 
**RotationStatement** | Pointer to **string** |  | [optional] 
**RotatorCredsType** | Pointer to **string** |  | [optional] 
**RotatorStatus** | Pointer to **string** | RotationStatus defines types of rotation Status | [optional] 
**RotatorType** | Pointer to **string** |  | [optional] 
**SamePassword** | Pointer to **bool** |  | [optional] 

## Methods

### NewRotatedSecretDetailsInfo

`func NewRotatedSecretDetailsInfo() *RotatedSecretDetailsInfo`

NewRotatedSecretDetailsInfo instantiates a new RotatedSecretDetailsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretDetailsInfoWithDefaults

`func NewRotatedSecretDetailsInfoWithDefaults() *RotatedSecretDetailsInfo`

NewRotatedSecretDetailsInfoWithDefaults instantiates a new RotatedSecretDetailsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletePreviousVersionInDays

`func (o *RotatedSecretDetailsInfo) GetDeletePreviousVersionInDays() int32`

GetDeletePreviousVersionInDays returns the DeletePreviousVersionInDays field if non-nil, zero value otherwise.

### GetDeletePreviousVersionInDaysOk

`func (o *RotatedSecretDetailsInfo) GetDeletePreviousVersionInDaysOk() (*int32, bool)`

GetDeletePreviousVersionInDaysOk returns a tuple with the DeletePreviousVersionInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletePreviousVersionInDays

`func (o *RotatedSecretDetailsInfo) SetDeletePreviousVersionInDays(v int32)`

SetDeletePreviousVersionInDays sets DeletePreviousVersionInDays field to given value.

### HasDeletePreviousVersionInDays

`func (o *RotatedSecretDetailsInfo) HasDeletePreviousVersionInDays() bool`

HasDeletePreviousVersionInDays returns a boolean if a field has been set.

### GetGwClusterId

`func (o *RotatedSecretDetailsInfo) GetGwClusterId() int64`

GetGwClusterId returns the GwClusterId field if non-nil, zero value otherwise.

### GetGwClusterIdOk

`func (o *RotatedSecretDetailsInfo) GetGwClusterIdOk() (*int64, bool)`

GetGwClusterIdOk returns a tuple with the GwClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwClusterId

`func (o *RotatedSecretDetailsInfo) SetGwClusterId(v int64)`

SetGwClusterId sets GwClusterId field to given value.

### HasGwClusterId

`func (o *RotatedSecretDetailsInfo) HasGwClusterId() bool`

HasGwClusterId returns a boolean if a field has been set.

### GetLastRotationError

`func (o *RotatedSecretDetailsInfo) GetLastRotationError() string`

GetLastRotationError returns the LastRotationError field if non-nil, zero value otherwise.

### GetLastRotationErrorOk

`func (o *RotatedSecretDetailsInfo) GetLastRotationErrorOk() (*string, bool)`

GetLastRotationErrorOk returns a tuple with the LastRotationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotationError

`func (o *RotatedSecretDetailsInfo) SetLastRotationError(v string)`

SetLastRotationError sets LastRotationError field to given value.

### HasLastRotationError

`func (o *RotatedSecretDetailsInfo) HasLastRotationError() bool`

HasLastRotationError returns a boolean if a field has been set.

### GetNumberOfVersionsToSave

`func (o *RotatedSecretDetailsInfo) GetNumberOfVersionsToSave() int32`

GetNumberOfVersionsToSave returns the NumberOfVersionsToSave field if non-nil, zero value otherwise.

### GetNumberOfVersionsToSaveOk

`func (o *RotatedSecretDetailsInfo) GetNumberOfVersionsToSaveOk() (*int32, bool)`

GetNumberOfVersionsToSaveOk returns a tuple with the NumberOfVersionsToSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVersionsToSave

`func (o *RotatedSecretDetailsInfo) SetNumberOfVersionsToSave(v int32)`

SetNumberOfVersionsToSave sets NumberOfVersionsToSave field to given value.

### HasNumberOfVersionsToSave

`func (o *RotatedSecretDetailsInfo) HasNumberOfVersionsToSave() bool`

HasNumberOfVersionsToSave returns a boolean if a field has been set.

### GetRotationHour

`func (o *RotatedSecretDetailsInfo) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *RotatedSecretDetailsInfo) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *RotatedSecretDetailsInfo) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *RotatedSecretDetailsInfo) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationIntervalMin

`func (o *RotatedSecretDetailsInfo) GetRotationIntervalMin() bool`

GetRotationIntervalMin returns the RotationIntervalMin field if non-nil, zero value otherwise.

### GetRotationIntervalMinOk

`func (o *RotatedSecretDetailsInfo) GetRotationIntervalMinOk() (*bool, bool)`

GetRotationIntervalMinOk returns a tuple with the RotationIntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationIntervalMin

`func (o *RotatedSecretDetailsInfo) SetRotationIntervalMin(v bool)`

SetRotationIntervalMin sets RotationIntervalMin field to given value.

### HasRotationIntervalMin

`func (o *RotatedSecretDetailsInfo) HasRotationIntervalMin() bool`

HasRotationIntervalMin returns a boolean if a field has been set.

### GetRotationStatement

`func (o *RotatedSecretDetailsInfo) GetRotationStatement() string`

GetRotationStatement returns the RotationStatement field if non-nil, zero value otherwise.

### GetRotationStatementOk

`func (o *RotatedSecretDetailsInfo) GetRotationStatementOk() (*string, bool)`

GetRotationStatementOk returns a tuple with the RotationStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationStatement

`func (o *RotatedSecretDetailsInfo) SetRotationStatement(v string)`

SetRotationStatement sets RotationStatement field to given value.

### HasRotationStatement

`func (o *RotatedSecretDetailsInfo) HasRotationStatement() bool`

HasRotationStatement returns a boolean if a field has been set.

### GetRotatorCredsType

`func (o *RotatedSecretDetailsInfo) GetRotatorCredsType() string`

GetRotatorCredsType returns the RotatorCredsType field if non-nil, zero value otherwise.

### GetRotatorCredsTypeOk

`func (o *RotatedSecretDetailsInfo) GetRotatorCredsTypeOk() (*string, bool)`

GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorCredsType

`func (o *RotatedSecretDetailsInfo) SetRotatorCredsType(v string)`

SetRotatorCredsType sets RotatorCredsType field to given value.

### HasRotatorCredsType

`func (o *RotatedSecretDetailsInfo) HasRotatorCredsType() bool`

HasRotatorCredsType returns a boolean if a field has been set.

### GetRotatorStatus

`func (o *RotatedSecretDetailsInfo) GetRotatorStatus() string`

GetRotatorStatus returns the RotatorStatus field if non-nil, zero value otherwise.

### GetRotatorStatusOk

`func (o *RotatedSecretDetailsInfo) GetRotatorStatusOk() (*string, bool)`

GetRotatorStatusOk returns a tuple with the RotatorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorStatus

`func (o *RotatedSecretDetailsInfo) SetRotatorStatus(v string)`

SetRotatorStatus sets RotatorStatus field to given value.

### HasRotatorStatus

`func (o *RotatedSecretDetailsInfo) HasRotatorStatus() bool`

HasRotatorStatus returns a boolean if a field has been set.

### GetRotatorType

`func (o *RotatedSecretDetailsInfo) GetRotatorType() string`

GetRotatorType returns the RotatorType field if non-nil, zero value otherwise.

### GetRotatorTypeOk

`func (o *RotatedSecretDetailsInfo) GetRotatorTypeOk() (*string, bool)`

GetRotatorTypeOk returns a tuple with the RotatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorType

`func (o *RotatedSecretDetailsInfo) SetRotatorType(v string)`

SetRotatorType sets RotatorType field to given value.

### HasRotatorType

`func (o *RotatedSecretDetailsInfo) HasRotatorType() bool`

HasRotatorType returns a boolean if a field has been set.

### GetSamePassword

`func (o *RotatedSecretDetailsInfo) GetSamePassword() bool`

GetSamePassword returns the SamePassword field if non-nil, zero value otherwise.

### GetSamePasswordOk

`func (o *RotatedSecretDetailsInfo) GetSamePasswordOk() (*bool, bool)`

GetSamePasswordOk returns a tuple with the SamePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePassword

`func (o *RotatedSecretDetailsInfo) SetSamePassword(v bool)`

SetSamePassword sets SamePassword field to given value.

### HasSamePassword

`func (o *RotatedSecretDetailsInfo) HasSamePassword() bool`

HasSamePassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


