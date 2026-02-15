# RotatedSecretDetailsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletePreviousVersionInDays** | Pointer to **int32** |  | [optional] 
**EnableCustomPasswordPolicy** | Pointer to **bool** |  | [optional] 
**GraceRotation** | Pointer to **bool** |  | [optional] 
**GraceRotationHour** | Pointer to **int32** |  | [optional] 
**GraceRotationInterval** | Pointer to **int32** |  | [optional] 
**GraceRotationTiming** | Pointer to **string** |  | [optional] 
**GwClusterId** | Pointer to **int64** |  | [optional] 
**IisAppsDetails** | Pointer to [**[]WindowsService**](WindowsService.md) |  | [optional] 
**LastRotationError** | Pointer to **string** |  | [optional] 
**ManagedByAkeyless** | Pointer to **bool** |  | [optional] 
**MaxVersions** | Pointer to **int64** |  | [optional] 
**NextAutoRotateType** | Pointer to **string** |  | [optional] 
**NumberOfVersionsToSave** | Pointer to **int32** |  | [optional] 
**PublicKeyRemotePath** | Pointer to **string** |  | [optional] 
**RotationHour** | Pointer to **int32** |  | [optional] 
**RotationIntervalMin** | Pointer to **bool** |  | [optional] 
**RotationStatement** | Pointer to **string** |  | [optional] 
**RotatorCredsType** | Pointer to **string** |  | [optional] 
**RotatorStatus** | Pointer to **string** | RotationStatus defines types of rotation Status | [optional] 
**RotatorType** | Pointer to **string** |  | [optional] 
**SamePassword** | Pointer to **bool** |  | [optional] 
**ServicesDetails** | Pointer to [**[]WindowsService**](WindowsService.md) |  | [optional] 
**TimeoutSeconds** | Pointer to **int64** |  | [optional] 

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

### GetEnableCustomPasswordPolicy

`func (o *RotatedSecretDetailsInfo) GetEnableCustomPasswordPolicy() bool`

GetEnableCustomPasswordPolicy returns the EnableCustomPasswordPolicy field if non-nil, zero value otherwise.

### GetEnableCustomPasswordPolicyOk

`func (o *RotatedSecretDetailsInfo) GetEnableCustomPasswordPolicyOk() (*bool, bool)`

GetEnableCustomPasswordPolicyOk returns a tuple with the EnableCustomPasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomPasswordPolicy

`func (o *RotatedSecretDetailsInfo) SetEnableCustomPasswordPolicy(v bool)`

SetEnableCustomPasswordPolicy sets EnableCustomPasswordPolicy field to given value.

### HasEnableCustomPasswordPolicy

`func (o *RotatedSecretDetailsInfo) HasEnableCustomPasswordPolicy() bool`

HasEnableCustomPasswordPolicy returns a boolean if a field has been set.

### GetGraceRotation

`func (o *RotatedSecretDetailsInfo) GetGraceRotation() bool`

GetGraceRotation returns the GraceRotation field if non-nil, zero value otherwise.

### GetGraceRotationOk

`func (o *RotatedSecretDetailsInfo) GetGraceRotationOk() (*bool, bool)`

GetGraceRotationOk returns a tuple with the GraceRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotation

`func (o *RotatedSecretDetailsInfo) SetGraceRotation(v bool)`

SetGraceRotation sets GraceRotation field to given value.

### HasGraceRotation

`func (o *RotatedSecretDetailsInfo) HasGraceRotation() bool`

HasGraceRotation returns a boolean if a field has been set.

### GetGraceRotationHour

`func (o *RotatedSecretDetailsInfo) GetGraceRotationHour() int32`

GetGraceRotationHour returns the GraceRotationHour field if non-nil, zero value otherwise.

### GetGraceRotationHourOk

`func (o *RotatedSecretDetailsInfo) GetGraceRotationHourOk() (*int32, bool)`

GetGraceRotationHourOk returns a tuple with the GraceRotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotationHour

`func (o *RotatedSecretDetailsInfo) SetGraceRotationHour(v int32)`

SetGraceRotationHour sets GraceRotationHour field to given value.

### HasGraceRotationHour

`func (o *RotatedSecretDetailsInfo) HasGraceRotationHour() bool`

HasGraceRotationHour returns a boolean if a field has been set.

### GetGraceRotationInterval

`func (o *RotatedSecretDetailsInfo) GetGraceRotationInterval() int32`

GetGraceRotationInterval returns the GraceRotationInterval field if non-nil, zero value otherwise.

### GetGraceRotationIntervalOk

`func (o *RotatedSecretDetailsInfo) GetGraceRotationIntervalOk() (*int32, bool)`

GetGraceRotationIntervalOk returns a tuple with the GraceRotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotationInterval

`func (o *RotatedSecretDetailsInfo) SetGraceRotationInterval(v int32)`

SetGraceRotationInterval sets GraceRotationInterval field to given value.

### HasGraceRotationInterval

`func (o *RotatedSecretDetailsInfo) HasGraceRotationInterval() bool`

HasGraceRotationInterval returns a boolean if a field has been set.

### GetGraceRotationTiming

`func (o *RotatedSecretDetailsInfo) GetGraceRotationTiming() string`

GetGraceRotationTiming returns the GraceRotationTiming field if non-nil, zero value otherwise.

### GetGraceRotationTimingOk

`func (o *RotatedSecretDetailsInfo) GetGraceRotationTimingOk() (*string, bool)`

GetGraceRotationTimingOk returns a tuple with the GraceRotationTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceRotationTiming

`func (o *RotatedSecretDetailsInfo) SetGraceRotationTiming(v string)`

SetGraceRotationTiming sets GraceRotationTiming field to given value.

### HasGraceRotationTiming

`func (o *RotatedSecretDetailsInfo) HasGraceRotationTiming() bool`

HasGraceRotationTiming returns a boolean if a field has been set.

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

### GetIisAppsDetails

`func (o *RotatedSecretDetailsInfo) GetIisAppsDetails() []WindowsService`

GetIisAppsDetails returns the IisAppsDetails field if non-nil, zero value otherwise.

### GetIisAppsDetailsOk

`func (o *RotatedSecretDetailsInfo) GetIisAppsDetailsOk() (*[]WindowsService, bool)`

GetIisAppsDetailsOk returns a tuple with the IisAppsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIisAppsDetails

`func (o *RotatedSecretDetailsInfo) SetIisAppsDetails(v []WindowsService)`

SetIisAppsDetails sets IisAppsDetails field to given value.

### HasIisAppsDetails

`func (o *RotatedSecretDetailsInfo) HasIisAppsDetails() bool`

HasIisAppsDetails returns a boolean if a field has been set.

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

### GetManagedByAkeyless

`func (o *RotatedSecretDetailsInfo) GetManagedByAkeyless() bool`

GetManagedByAkeyless returns the ManagedByAkeyless field if non-nil, zero value otherwise.

### GetManagedByAkeylessOk

`func (o *RotatedSecretDetailsInfo) GetManagedByAkeylessOk() (*bool, bool)`

GetManagedByAkeylessOk returns a tuple with the ManagedByAkeyless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedByAkeyless

`func (o *RotatedSecretDetailsInfo) SetManagedByAkeyless(v bool)`

SetManagedByAkeyless sets ManagedByAkeyless field to given value.

### HasManagedByAkeyless

`func (o *RotatedSecretDetailsInfo) HasManagedByAkeyless() bool`

HasManagedByAkeyless returns a boolean if a field has been set.

### GetMaxVersions

`func (o *RotatedSecretDetailsInfo) GetMaxVersions() int64`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *RotatedSecretDetailsInfo) GetMaxVersionsOk() (*int64, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *RotatedSecretDetailsInfo) SetMaxVersions(v int64)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *RotatedSecretDetailsInfo) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetNextAutoRotateType

`func (o *RotatedSecretDetailsInfo) GetNextAutoRotateType() string`

GetNextAutoRotateType returns the NextAutoRotateType field if non-nil, zero value otherwise.

### GetNextAutoRotateTypeOk

`func (o *RotatedSecretDetailsInfo) GetNextAutoRotateTypeOk() (*string, bool)`

GetNextAutoRotateTypeOk returns a tuple with the NextAutoRotateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAutoRotateType

`func (o *RotatedSecretDetailsInfo) SetNextAutoRotateType(v string)`

SetNextAutoRotateType sets NextAutoRotateType field to given value.

### HasNextAutoRotateType

`func (o *RotatedSecretDetailsInfo) HasNextAutoRotateType() bool`

HasNextAutoRotateType returns a boolean if a field has been set.

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

### GetPublicKeyRemotePath

`func (o *RotatedSecretDetailsInfo) GetPublicKeyRemotePath() string`

GetPublicKeyRemotePath returns the PublicKeyRemotePath field if non-nil, zero value otherwise.

### GetPublicKeyRemotePathOk

`func (o *RotatedSecretDetailsInfo) GetPublicKeyRemotePathOk() (*string, bool)`

GetPublicKeyRemotePathOk returns a tuple with the PublicKeyRemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyRemotePath

`func (o *RotatedSecretDetailsInfo) SetPublicKeyRemotePath(v string)`

SetPublicKeyRemotePath sets PublicKeyRemotePath field to given value.

### HasPublicKeyRemotePath

`func (o *RotatedSecretDetailsInfo) HasPublicKeyRemotePath() bool`

HasPublicKeyRemotePath returns a boolean if a field has been set.

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

### GetServicesDetails

`func (o *RotatedSecretDetailsInfo) GetServicesDetails() []WindowsService`

GetServicesDetails returns the ServicesDetails field if non-nil, zero value otherwise.

### GetServicesDetailsOk

`func (o *RotatedSecretDetailsInfo) GetServicesDetailsOk() (*[]WindowsService, bool)`

GetServicesDetailsOk returns a tuple with the ServicesDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesDetails

`func (o *RotatedSecretDetailsInfo) SetServicesDetails(v []WindowsService)`

SetServicesDetails sets ServicesDetails field to given value.

### HasServicesDetails

`func (o *RotatedSecretDetailsInfo) HasServicesDetails() bool`

HasServicesDetails returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *RotatedSecretDetailsInfo) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *RotatedSecretDetailsInfo) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *RotatedSecretDetailsInfo) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *RotatedSecretDetailsInfo) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


