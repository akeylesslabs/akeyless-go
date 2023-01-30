# ListAuthMethodsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethods** | Pointer to [**[]AuthMethod**](AuthMethod.md) |  | [optional] 
**NextPage** | Pointer to **string** |  | [optional] 

## Methods

### NewListAuthMethodsOutput

`func NewListAuthMethodsOutput() *ListAuthMethodsOutput`

NewListAuthMethodsOutput instantiates a new ListAuthMethodsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthMethodsOutputWithDefaults

`func NewListAuthMethodsOutputWithDefaults() *ListAuthMethodsOutput`

NewListAuthMethodsOutputWithDefaults instantiates a new ListAuthMethodsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethods

`func (o *ListAuthMethodsOutput) GetAuthMethods() []AuthMethod`

GetAuthMethods returns the AuthMethods field if non-nil, zero value otherwise.

### GetAuthMethodsOk

`func (o *ListAuthMethodsOutput) GetAuthMethodsOk() (*[]AuthMethod, bool)`

GetAuthMethodsOk returns a tuple with the AuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethods

`func (o *ListAuthMethodsOutput) SetAuthMethods(v []AuthMethod)`

SetAuthMethods sets AuthMethods field to given value.

### HasAuthMethods

`func (o *ListAuthMethodsOutput) HasAuthMethods() bool`

HasAuthMethods returns a boolean if a field has been set.

### GetNextPage

`func (o *ListAuthMethodsOutput) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ListAuthMethodsOutput) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ListAuthMethodsOutput) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ListAuthMethodsOutput) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


