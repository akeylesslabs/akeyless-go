# ListAuthMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaginationToken** | Pointer to **string** | Next page reference | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewListAuthMethods

`func NewListAuthMethods() *ListAuthMethods`

NewListAuthMethods instantiates a new ListAuthMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthMethodsWithDefaults

`func NewListAuthMethodsWithDefaults() *ListAuthMethods`

NewListAuthMethodsWithDefaults instantiates a new ListAuthMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaginationToken

`func (o *ListAuthMethods) GetPaginationToken() string`

GetPaginationToken returns the PaginationToken field if non-nil, zero value otherwise.

### GetPaginationTokenOk

`func (o *ListAuthMethods) GetPaginationTokenOk() (*string, bool)`

GetPaginationTokenOk returns a tuple with the PaginationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationToken

`func (o *ListAuthMethods) SetPaginationToken(v string)`

SetPaginationToken sets PaginationToken field to given value.

### HasPaginationToken

`func (o *ListAuthMethods) HasPaginationToken() bool`

HasPaginationToken returns a boolean if a field has been set.

### GetToken

`func (o *ListAuthMethods) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAuthMethods) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAuthMethods) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ListAuthMethods) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ListAuthMethods) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ListAuthMethods) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ListAuthMethods) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ListAuthMethods) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


