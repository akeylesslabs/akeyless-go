# FolderSyncOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedItems** | Pointer to **map[string]string** |  | [optional] 
**FolderName** | Pointer to **string** |  | [optional] 
**SkippedItems** | Pointer to **[]string** |  | [optional] 
**SyncedItems** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFolderSyncOutput

`func NewFolderSyncOutput() *FolderSyncOutput`

NewFolderSyncOutput instantiates a new FolderSyncOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderSyncOutputWithDefaults

`func NewFolderSyncOutputWithDefaults() *FolderSyncOutput`

NewFolderSyncOutputWithDefaults instantiates a new FolderSyncOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedItems

`func (o *FolderSyncOutput) GetFailedItems() map[string]string`

GetFailedItems returns the FailedItems field if non-nil, zero value otherwise.

### GetFailedItemsOk

`func (o *FolderSyncOutput) GetFailedItemsOk() (*map[string]string, bool)`

GetFailedItemsOk returns a tuple with the FailedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedItems

`func (o *FolderSyncOutput) SetFailedItems(v map[string]string)`

SetFailedItems sets FailedItems field to given value.

### HasFailedItems

`func (o *FolderSyncOutput) HasFailedItems() bool`

HasFailedItems returns a boolean if a field has been set.

### GetFolderName

`func (o *FolderSyncOutput) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *FolderSyncOutput) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *FolderSyncOutput) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *FolderSyncOutput) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### GetSkippedItems

`func (o *FolderSyncOutput) GetSkippedItems() []string`

GetSkippedItems returns the SkippedItems field if non-nil, zero value otherwise.

### GetSkippedItemsOk

`func (o *FolderSyncOutput) GetSkippedItemsOk() (*[]string, bool)`

GetSkippedItemsOk returns a tuple with the SkippedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedItems

`func (o *FolderSyncOutput) SetSkippedItems(v []string)`

SetSkippedItems sets SkippedItems field to given value.

### HasSkippedItems

`func (o *FolderSyncOutput) HasSkippedItems() bool`

HasSkippedItems returns a boolean if a field has been set.

### GetSyncedItems

`func (o *FolderSyncOutput) GetSyncedItems() []string`

GetSyncedItems returns the SyncedItems field if non-nil, zero value otherwise.

### GetSyncedItemsOk

`func (o *FolderSyncOutput) GetSyncedItemsOk() (*[]string, bool)`

GetSyncedItemsOk returns a tuple with the SyncedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedItems

`func (o *FolderSyncOutput) SetSyncedItems(v []string)`

SetSyncedItems sets SyncedItems field to given value.

### HasSyncedItems

`func (o *FolderSyncOutput) HasSyncedItems() bool`

HasSyncedItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


