# ListItemsInPathOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folders** | Pointer to **[]string** |  | [optional] 
**Items** | Pointer to [**[]Item**](Item.md) |  | [optional] 
**NextPage** | Pointer to **string** |  | [optional] 

## Methods

### NewListItemsInPathOutput

`func NewListItemsInPathOutput() *ListItemsInPathOutput`

NewListItemsInPathOutput instantiates a new ListItemsInPathOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemsInPathOutputWithDefaults

`func NewListItemsInPathOutputWithDefaults() *ListItemsInPathOutput`

NewListItemsInPathOutputWithDefaults instantiates a new ListItemsInPathOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolders

`func (o *ListItemsInPathOutput) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *ListItemsInPathOutput) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *ListItemsInPathOutput) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *ListItemsInPathOutput) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetItems

`func (o *ListItemsInPathOutput) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListItemsInPathOutput) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListItemsInPathOutput) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListItemsInPathOutput) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextPage

`func (o *ListItemsInPathOutput) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ListItemsInPathOutput) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ListItemsInPathOutput) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ListItemsInPathOutput) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


