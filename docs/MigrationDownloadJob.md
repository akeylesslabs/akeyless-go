# MigrationDownloadJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to [**MigrationDownloadSummary**](MigrationDownloadSummary.md) |  | [optional] 
**TotalItems** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewMigrationDownloadJob

`func NewMigrationDownloadJob() *MigrationDownloadJob`

NewMigrationDownloadJob instantiates a new MigrationDownloadJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationDownloadJobWithDefaults

`func NewMigrationDownloadJobWithDefaults() *MigrationDownloadJob`

NewMigrationDownloadJobWithDefaults instantiates a new MigrationDownloadJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *MigrationDownloadJob) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MigrationDownloadJob) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MigrationDownloadJob) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MigrationDownloadJob) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStartedAt

`func (o *MigrationDownloadJob) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MigrationDownloadJob) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MigrationDownloadJob) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MigrationDownloadJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetSummary

`func (o *MigrationDownloadJob) GetSummary() MigrationDownloadSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *MigrationDownloadJob) GetSummaryOk() (*MigrationDownloadSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *MigrationDownloadJob) SetSummary(v MigrationDownloadSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *MigrationDownloadJob) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTotalItems

`func (o *MigrationDownloadJob) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *MigrationDownloadJob) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *MigrationDownloadJob) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *MigrationDownloadJob) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetType

`func (o *MigrationDownloadJob) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MigrationDownloadJob) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MigrationDownloadJob) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MigrationDownloadJob) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


