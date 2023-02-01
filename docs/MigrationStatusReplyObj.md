# MigrationStatusReplyObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationTime** | Pointer to **string** |  | [optional] 
**LastStatusMessage** | Pointer to **string** |  | [optional] 
**MaxNameLength** | Pointer to **int64** |  | [optional] 
**MaxValueLength** | Pointer to **int64** |  | [optional] 
**MigrationId** | Pointer to **string** |  | [optional] 
**MigrationItems** | Pointer to [**MigrationItems**](MigrationItems.md) |  | [optional] 
**MigrationName** | Pointer to **string** |  | [optional] 
**MigrationState** | Pointer to **string** |  | [optional] 
**MigrationType** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 

## Methods

### NewMigrationStatusReplyObj

`func NewMigrationStatusReplyObj() *MigrationStatusReplyObj`

NewMigrationStatusReplyObj instantiates a new MigrationStatusReplyObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationStatusReplyObjWithDefaults

`func NewMigrationStatusReplyObjWithDefaults() *MigrationStatusReplyObj`

NewMigrationStatusReplyObjWithDefaults instantiates a new MigrationStatusReplyObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationTime

`func (o *MigrationStatusReplyObj) GetDurationTime() string`

GetDurationTime returns the DurationTime field if non-nil, zero value otherwise.

### GetDurationTimeOk

`func (o *MigrationStatusReplyObj) GetDurationTimeOk() (*string, bool)`

GetDurationTimeOk returns a tuple with the DurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationTime

`func (o *MigrationStatusReplyObj) SetDurationTime(v string)`

SetDurationTime sets DurationTime field to given value.

### HasDurationTime

`func (o *MigrationStatusReplyObj) HasDurationTime() bool`

HasDurationTime returns a boolean if a field has been set.

### GetLastStatusMessage

`func (o *MigrationStatusReplyObj) GetLastStatusMessage() string`

GetLastStatusMessage returns the LastStatusMessage field if non-nil, zero value otherwise.

### GetLastStatusMessageOk

`func (o *MigrationStatusReplyObj) GetLastStatusMessageOk() (*string, bool)`

GetLastStatusMessageOk returns a tuple with the LastStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusMessage

`func (o *MigrationStatusReplyObj) SetLastStatusMessage(v string)`

SetLastStatusMessage sets LastStatusMessage field to given value.

### HasLastStatusMessage

`func (o *MigrationStatusReplyObj) HasLastStatusMessage() bool`

HasLastStatusMessage returns a boolean if a field has been set.

### GetMaxNameLength

`func (o *MigrationStatusReplyObj) GetMaxNameLength() int64`

GetMaxNameLength returns the MaxNameLength field if non-nil, zero value otherwise.

### GetMaxNameLengthOk

`func (o *MigrationStatusReplyObj) GetMaxNameLengthOk() (*int64, bool)`

GetMaxNameLengthOk returns a tuple with the MaxNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNameLength

`func (o *MigrationStatusReplyObj) SetMaxNameLength(v int64)`

SetMaxNameLength sets MaxNameLength field to given value.

### HasMaxNameLength

`func (o *MigrationStatusReplyObj) HasMaxNameLength() bool`

HasMaxNameLength returns a boolean if a field has been set.

### GetMaxValueLength

`func (o *MigrationStatusReplyObj) GetMaxValueLength() int64`

GetMaxValueLength returns the MaxValueLength field if non-nil, zero value otherwise.

### GetMaxValueLengthOk

`func (o *MigrationStatusReplyObj) GetMaxValueLengthOk() (*int64, bool)`

GetMaxValueLengthOk returns a tuple with the MaxValueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValueLength

`func (o *MigrationStatusReplyObj) SetMaxValueLength(v int64)`

SetMaxValueLength sets MaxValueLength field to given value.

### HasMaxValueLength

`func (o *MigrationStatusReplyObj) HasMaxValueLength() bool`

HasMaxValueLength returns a boolean if a field has been set.

### GetMigrationId

`func (o *MigrationStatusReplyObj) GetMigrationId() string`

GetMigrationId returns the MigrationId field if non-nil, zero value otherwise.

### GetMigrationIdOk

`func (o *MigrationStatusReplyObj) GetMigrationIdOk() (*string, bool)`

GetMigrationIdOk returns a tuple with the MigrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationId

`func (o *MigrationStatusReplyObj) SetMigrationId(v string)`

SetMigrationId sets MigrationId field to given value.

### HasMigrationId

`func (o *MigrationStatusReplyObj) HasMigrationId() bool`

HasMigrationId returns a boolean if a field has been set.

### GetMigrationItems

`func (o *MigrationStatusReplyObj) GetMigrationItems() MigrationItems`

GetMigrationItems returns the MigrationItems field if non-nil, zero value otherwise.

### GetMigrationItemsOk

`func (o *MigrationStatusReplyObj) GetMigrationItemsOk() (*MigrationItems, bool)`

GetMigrationItemsOk returns a tuple with the MigrationItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationItems

`func (o *MigrationStatusReplyObj) SetMigrationItems(v MigrationItems)`

SetMigrationItems sets MigrationItems field to given value.

### HasMigrationItems

`func (o *MigrationStatusReplyObj) HasMigrationItems() bool`

HasMigrationItems returns a boolean if a field has been set.

### GetMigrationName

`func (o *MigrationStatusReplyObj) GetMigrationName() string`

GetMigrationName returns the MigrationName field if non-nil, zero value otherwise.

### GetMigrationNameOk

`func (o *MigrationStatusReplyObj) GetMigrationNameOk() (*string, bool)`

GetMigrationNameOk returns a tuple with the MigrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationName

`func (o *MigrationStatusReplyObj) SetMigrationName(v string)`

SetMigrationName sets MigrationName field to given value.

### HasMigrationName

`func (o *MigrationStatusReplyObj) HasMigrationName() bool`

HasMigrationName returns a boolean if a field has been set.

### GetMigrationState

`func (o *MigrationStatusReplyObj) GetMigrationState() string`

GetMigrationState returns the MigrationState field if non-nil, zero value otherwise.

### GetMigrationStateOk

`func (o *MigrationStatusReplyObj) GetMigrationStateOk() (*string, bool)`

GetMigrationStateOk returns a tuple with the MigrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationState

`func (o *MigrationStatusReplyObj) SetMigrationState(v string)`

SetMigrationState sets MigrationState field to given value.

### HasMigrationState

`func (o *MigrationStatusReplyObj) HasMigrationState() bool`

HasMigrationState returns a boolean if a field has been set.

### GetMigrationType

`func (o *MigrationStatusReplyObj) GetMigrationType() string`

GetMigrationType returns the MigrationType field if non-nil, zero value otherwise.

### GetMigrationTypeOk

`func (o *MigrationStatusReplyObj) GetMigrationTypeOk() (*string, bool)`

GetMigrationTypeOk returns a tuple with the MigrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationType

`func (o *MigrationStatusReplyObj) SetMigrationType(v string)`

SetMigrationType sets MigrationType field to given value.

### HasMigrationType

`func (o *MigrationStatusReplyObj) HasMigrationType() bool`

HasMigrationType returns a boolean if a field has been set.

### GetStartTime

`func (o *MigrationStatusReplyObj) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MigrationStatusReplyObj) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MigrationStatusReplyObj) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MigrationStatusReplyObj) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


