# AdminsConfigPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminsMigrationStatus** | Pointer to **int64** |  | [optional] 
**AllowedAccess** | Pointer to [**map[string]AllowedAccessOld**](AllowedAccessOld.md) |  | [optional] 

## Methods

### NewAdminsConfigPart

`func NewAdminsConfigPart() *AdminsConfigPart`

NewAdminsConfigPart instantiates a new AdminsConfigPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminsConfigPartWithDefaults

`func NewAdminsConfigPartWithDefaults() *AdminsConfigPart`

NewAdminsConfigPartWithDefaults instantiates a new AdminsConfigPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminsMigrationStatus

`func (o *AdminsConfigPart) GetAdminsMigrationStatus() int64`

GetAdminsMigrationStatus returns the AdminsMigrationStatus field if non-nil, zero value otherwise.

### GetAdminsMigrationStatusOk

`func (o *AdminsConfigPart) GetAdminsMigrationStatusOk() (*int64, bool)`

GetAdminsMigrationStatusOk returns a tuple with the AdminsMigrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminsMigrationStatus

`func (o *AdminsConfigPart) SetAdminsMigrationStatus(v int64)`

SetAdminsMigrationStatus sets AdminsMigrationStatus field to given value.

### HasAdminsMigrationStatus

`func (o *AdminsConfigPart) HasAdminsMigrationStatus() bool`

HasAdminsMigrationStatus returns a boolean if a field has been set.

### GetAllowedAccess

`func (o *AdminsConfigPart) GetAllowedAccess() map[string]AllowedAccessOld`

GetAllowedAccess returns the AllowedAccess field if non-nil, zero value otherwise.

### GetAllowedAccessOk

`func (o *AdminsConfigPart) GetAllowedAccessOk() (*map[string]AllowedAccessOld, bool)`

GetAllowedAccessOk returns a tuple with the AllowedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAccess

`func (o *AdminsConfigPart) SetAllowedAccess(v map[string]AllowedAccessOld)`

SetAllowedAccess sets AllowedAccess field to given value.

### HasAllowedAccess

`func (o *AdminsConfigPart) HasAllowedAccess() bool`

HasAllowedAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


