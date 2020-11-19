# UpdateRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Role name | 
**NewComment** | Pointer to **string** | New comment about the role | [optional] [default to "default_comment"]
**NewName** | Pointer to **string** | New Role name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateRole

`func NewUpdateRole(name string, ) *UpdateRole`

NewUpdateRole instantiates a new UpdateRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleWithDefaults

`func NewUpdateRoleWithDefaults() *UpdateRole`

NewUpdateRoleWithDefaults instantiates a new UpdateRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRole) SetName(v string)`

SetName sets Name field to given value.


### GetNewComment

`func (o *UpdateRole) GetNewComment() string`

GetNewComment returns the NewComment field if non-nil, zero value otherwise.

### GetNewCommentOk

`func (o *UpdateRole) GetNewCommentOk() (*string, bool)`

GetNewCommentOk returns a tuple with the NewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewComment

`func (o *UpdateRole) SetNewComment(v string)`

SetNewComment sets NewComment field to given value.

### HasNewComment

`func (o *UpdateRole) HasNewComment() bool`

HasNewComment returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateRole) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateRole) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateRole) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateRole) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateRole) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateRole) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateRole) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateRole) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateRole) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateRole) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateRole) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateRole) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

