# ListGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Filter by auth method name or part of it | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**PaginationToken** | Pointer to **string** | Next page reference | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewListGroups

`func NewListGroups() *ListGroups`

NewListGroups instantiates a new ListGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGroupsWithDefaults

`func NewListGroupsWithDefaults() *ListGroups`

NewListGroupsWithDefaults instantiates a new ListGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ListGroups) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ListGroups) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ListGroups) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ListGroups) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetJson

`func (o *ListGroups) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ListGroups) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ListGroups) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ListGroups) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetPaginationToken

`func (o *ListGroups) GetPaginationToken() string`

GetPaginationToken returns the PaginationToken field if non-nil, zero value otherwise.

### GetPaginationTokenOk

`func (o *ListGroups) GetPaginationTokenOk() (*string, bool)`

GetPaginationTokenOk returns a tuple with the PaginationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationToken

`func (o *ListGroups) SetPaginationToken(v string)`

SetPaginationToken sets PaginationToken field to given value.

### HasPaginationToken

`func (o *ListGroups) HasPaginationToken() bool`

HasPaginationToken returns a boolean if a field has been set.

### GetToken

`func (o *ListGroups) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListGroups) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListGroups) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ListGroups) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ListGroups) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ListGroups) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ListGroups) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ListGroups) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


