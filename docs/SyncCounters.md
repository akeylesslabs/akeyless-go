# SyncCounters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int64** |  | [optional] 
**Failed** | Pointer to **int64** |  | [optional] 
**Skipped** | Pointer to **int64** |  | [optional] 

## Methods

### NewSyncCounters

`func NewSyncCounters() *SyncCounters`

NewSyncCounters instantiates a new SyncCounters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncCountersWithDefaults

`func NewSyncCountersWithDefaults() *SyncCounters`

NewSyncCountersWithDefaults instantiates a new SyncCounters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SyncCounters) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SyncCounters) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SyncCounters) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SyncCounters) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFailed

`func (o *SyncCounters) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *SyncCounters) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *SyncCounters) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *SyncCounters) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetSkipped

`func (o *SyncCounters) GetSkipped() int64`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *SyncCounters) GetSkippedOk() (*int64, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *SyncCounters) SetSkipped(v int64)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *SyncCounters) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


