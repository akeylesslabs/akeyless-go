# ListItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**AdvancedFilter** | Pointer to **string** | Filter by item name/username/website or part of it | [optional] 
**AutoPagination** | Pointer to **string** | Retrieve all items using pagination, when disabled retrieving only first 1000 items | [optional] [default to "enabled"]
**Filter** | Pointer to **string** | Filter by item name or part of it | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MinimalView** | Pointer to **bool** | Show only basic information of the items | [optional] 
**ModifiedAfter** | Pointer to **int64** | List only secrets modified after specified date (in unix time) | [optional] 
**PaginationToken** | Pointer to **string** | Next page reference | [optional] 
**Path** | Pointer to **string** | Path to folder | [optional] 
**SraOnly** | Pointer to **bool** | Filter by items with SRA functionality enabled | [optional] [default to false]
**SubTypes** | Pointer to **[]string** |  | [optional] 
**Tag** | Pointer to **string** | Filter by item tag | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **[]string** | The item types list of the requested items. In case it is empty, all types of items will be returned. options: [key, static-secret, dynamic-secret, classic-key] | [optional] 
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

### GetAccessibility

`func (o *ListItems) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *ListItems) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *ListItems) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *ListItems) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetAdvancedFilter

`func (o *ListItems) GetAdvancedFilter() string`

GetAdvancedFilter returns the AdvancedFilter field if non-nil, zero value otherwise.

### GetAdvancedFilterOk

`func (o *ListItems) GetAdvancedFilterOk() (*string, bool)`

GetAdvancedFilterOk returns a tuple with the AdvancedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedFilter

`func (o *ListItems) SetAdvancedFilter(v string)`

SetAdvancedFilter sets AdvancedFilter field to given value.

### HasAdvancedFilter

`func (o *ListItems) HasAdvancedFilter() bool`

HasAdvancedFilter returns a boolean if a field has been set.

### GetAutoPagination

`func (o *ListItems) GetAutoPagination() string`

GetAutoPagination returns the AutoPagination field if non-nil, zero value otherwise.

### GetAutoPaginationOk

`func (o *ListItems) GetAutoPaginationOk() (*string, bool)`

GetAutoPaginationOk returns a tuple with the AutoPagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPagination

`func (o *ListItems) SetAutoPagination(v string)`

SetAutoPagination sets AutoPagination field to given value.

### HasAutoPagination

`func (o *ListItems) HasAutoPagination() bool`

HasAutoPagination returns a boolean if a field has been set.

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

### GetJson

`func (o *ListItems) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ListItems) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ListItems) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ListItems) HasJson() bool`

HasJson returns a boolean if a field has been set.

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

### GetModifiedAfter

`func (o *ListItems) GetModifiedAfter() int64`

GetModifiedAfter returns the ModifiedAfter field if non-nil, zero value otherwise.

### GetModifiedAfterOk

`func (o *ListItems) GetModifiedAfterOk() (*int64, bool)`

GetModifiedAfterOk returns a tuple with the ModifiedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAfter

`func (o *ListItems) SetModifiedAfter(v int64)`

SetModifiedAfter sets ModifiedAfter field to given value.

### HasModifiedAfter

`func (o *ListItems) HasModifiedAfter() bool`

HasModifiedAfter returns a boolean if a field has been set.

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

### GetSraOnly

`func (o *ListItems) GetSraOnly() bool`

GetSraOnly returns the SraOnly field if non-nil, zero value otherwise.

### GetSraOnlyOk

`func (o *ListItems) GetSraOnlyOk() (*bool, bool)`

GetSraOnlyOk returns a tuple with the SraOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSraOnly

`func (o *ListItems) SetSraOnly(v bool)`

SetSraOnly sets SraOnly field to given value.

### HasSraOnly

`func (o *ListItems) HasSraOnly() bool`

HasSraOnly returns a boolean if a field has been set.

### GetSubTypes

`func (o *ListItems) GetSubTypes() []string`

GetSubTypes returns the SubTypes field if non-nil, zero value otherwise.

### GetSubTypesOk

`func (o *ListItems) GetSubTypesOk() (*[]string, bool)`

GetSubTypesOk returns a tuple with the SubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTypes

`func (o *ListItems) SetSubTypes(v []string)`

SetSubTypes sets SubTypes field to given value.

### HasSubTypes

`func (o *ListItems) HasSubTypes() bool`

HasSubTypes returns a boolean if a field has been set.

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


