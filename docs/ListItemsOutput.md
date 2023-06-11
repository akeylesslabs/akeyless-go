# ListItemsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Item**](Item.md) |  | [optional] 
**NextPage** | Pointer to **string** |  | [optional] 

## Methods

### NewListItemsOutput

`func NewListItemsOutput() *ListItemsOutput`

NewListItemsOutput instantiates a new ListItemsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemsOutputWithDefaults

`func NewListItemsOutputWithDefaults() *ListItemsOutput`

NewListItemsOutputWithDefaults instantiates a new ListItemsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListItemsOutput) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListItemsOutput) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListItemsOutput) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListItemsOutput) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextPage

`func (o *ListItemsOutput) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ListItemsOutput) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ListItemsOutput) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ListItemsOutput) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


