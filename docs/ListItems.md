# ListItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Filter by item name or part of it | [optional] 
**MinimalView** | Pointer to **bool** |  | [optional] 
**PaginationToken** | Pointer to **string** | Next page reference | [optional] 
**Path** | Pointer to **string** | Path to folder | [optional] 
**Tag** | Pointer to **string** | Filter by item tag | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **[]string** | The item types list of the requested items. In case it is empty, all types of items will be returned. options: [key, static-secret, dynamic-secret] | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewListItems

`func NewListItems() *ListItems`

NewListItems instantiates a new ListItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemsWithDefaults

`func NewListItemsWithDefaults() *ListItems`

NewListItemsWithDefaults instantiates a new ListItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ListItems) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ListItems) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ListItems) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ListItems) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMinimalView

`func (o *ListItems) GetMinimalView() bool`

GetMinimalView returns the MinimalView field if non-nil, zero value otherwise.

### GetMinimalViewOk

`func (o *ListItems) GetMinimalViewOk() (*bool, bool)`

GetMinimalViewOk returns a tuple with the MinimalView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalView

`func (o *ListItems) SetMinimalView(v bool)`

SetMinimalView sets MinimalView field to given value.

### HasMinimalView

`func (o *ListItems) HasMinimalView() bool`

HasMinimalView returns a boolean if a field has been set.

### GetPaginationToken

`func (o *ListItems) GetPaginationToken() string`

GetPaginationToken returns the PaginationToken field if non-nil, zero value otherwise.

### GetPaginationTokenOk

`func (o *ListItems) GetPaginationTokenOk() (*string, bool)`

GetPaginationTokenOk returns a tuple with the PaginationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationToken

`func (o *ListItems) SetPaginationToken(v string)`

SetPaginationToken sets PaginationToken field to given value.

### HasPaginationToken

`func (o *ListItems) HasPaginationToken() bool`

HasPaginationToken returns a boolean if a field has been set.

### GetPath

`func (o *ListItems) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ListItems) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ListItems) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ListItems) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTag

`func (o *ListItems) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ListItems) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ListItems) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ListItems) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetToken

`func (o *ListItems) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListItems) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListItems) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ListItems) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *ListItems) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListItems) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListItems) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *ListItems) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *ListItems) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ListItems) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ListItems) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ListItems) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


