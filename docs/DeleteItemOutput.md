# DeleteItemOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletionDate** | Pointer to **time.Time** |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**ItemName** | Pointer to **string** |  | [optional] 
**VersionDeleted** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeleteItemOutput

`func NewDeleteItemOutput() *DeleteItemOutput`

NewDeleteItemOutput instantiates a new DeleteItemOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteItemOutputWithDefaults

`func NewDeleteItemOutputWithDefaults() *DeleteItemOutput`

NewDeleteItemOutputWithDefaults instantiates a new DeleteItemOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletionDate

`func (o *DeleteItemOutput) GetDeletionDate() time.Time`

GetDeletionDate returns the DeletionDate field if non-nil, zero value otherwise.

### GetDeletionDateOk

`func (o *DeleteItemOutput) GetDeletionDateOk() (*time.Time, bool)`

GetDeletionDateOk returns a tuple with the DeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDate

`func (o *DeleteItemOutput) SetDeletionDate(v time.Time)`

SetDeletionDate sets DeletionDate field to given value.

### HasDeletionDate

`func (o *DeleteItemOutput) HasDeletionDate() bool`

HasDeletionDate returns a boolean if a field has been set.

### GetItemId

`func (o *DeleteItemOutput) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DeleteItemOutput) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DeleteItemOutput) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DeleteItemOutput) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemName

`func (o *DeleteItemOutput) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *DeleteItemOutput) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *DeleteItemOutput) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *DeleteItemOutput) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetVersionDeleted

`func (o *DeleteItemOutput) GetVersionDeleted() int32`

GetVersionDeleted returns the VersionDeleted field if non-nil, zero value otherwise.

### GetVersionDeletedOk

`func (o *DeleteItemOutput) GetVersionDeletedOk() (*int32, bool)`

GetVersionDeletedOk returns a tuple with the VersionDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionDeleted

`func (o *DeleteItemOutput) SetVersionDeleted(v int32)`

SetVersionDeleted sets VersionDeleted field to given value.

### HasVersionDeleted

`func (o *DeleteItemOutput) HasVersionDeleted() bool`

HasVersionDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


