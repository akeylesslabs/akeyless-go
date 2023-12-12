# DeleteItemsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedItems** | Pointer to **[]string** |  | [optional] 
**FailedDeletedItems** | Pointer to **map[string]string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteItemsOutput

`func NewDeleteItemsOutput() *DeleteItemsOutput`

NewDeleteItemsOutput instantiates a new DeleteItemsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteItemsOutputWithDefaults

`func NewDeleteItemsOutputWithDefaults() *DeleteItemsOutput`

NewDeleteItemsOutputWithDefaults instantiates a new DeleteItemsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedItems

`func (o *DeleteItemsOutput) GetDeletedItems() []string`

GetDeletedItems returns the DeletedItems field if non-nil, zero value otherwise.

### GetDeletedItemsOk

`func (o *DeleteItemsOutput) GetDeletedItemsOk() (*[]string, bool)`

GetDeletedItemsOk returns a tuple with the DeletedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedItems

`func (o *DeleteItemsOutput) SetDeletedItems(v []string)`

SetDeletedItems sets DeletedItems field to given value.

### HasDeletedItems

`func (o *DeleteItemsOutput) HasDeletedItems() bool`

HasDeletedItems returns a boolean if a field has been set.

### GetFailedDeletedItems

`func (o *DeleteItemsOutput) GetFailedDeletedItems() map[string]string`

GetFailedDeletedItems returns the FailedDeletedItems field if non-nil, zero value otherwise.

### GetFailedDeletedItemsOk

`func (o *DeleteItemsOutput) GetFailedDeletedItemsOk() (*map[string]string, bool)`

GetFailedDeletedItemsOk returns a tuple with the FailedDeletedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDeletedItems

`func (o *DeleteItemsOutput) SetFailedDeletedItems(v map[string]string)`

SetFailedDeletedItems sets FailedDeletedItems field to given value.

### HasFailedDeletedItems

`func (o *DeleteItemsOutput) HasFailedDeletedItems() bool`

HasFailedDeletedItems returns a boolean if a field has been set.

### GetPath

`func (o *DeleteItemsOutput) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeleteItemsOutput) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeleteItemsOutput) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DeleteItemsOutput) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


