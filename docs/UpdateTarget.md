# UpdateTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | **string** | Target name | 
**NewComment** | Pointer to **string** | New comment about the target | [optional] [default to "default_comment"]
**NewName** | Pointer to **string** | New Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateTarget

`func NewUpdateTarget(name string, ) *UpdateTarget`

NewUpdateTarget instantiates a new UpdateTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTargetWithDefaults

`func NewUpdateTargetWithDefaults() *UpdateTarget`

NewUpdateTargetWithDefaults instantiates a new UpdateTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *UpdateTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *UpdateTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewComment

`func (o *UpdateTarget) GetNewComment() string`

GetNewComment returns the NewComment field if non-nil, zero value otherwise.

### GetNewCommentOk

`func (o *UpdateTarget) GetNewCommentOk() (*string, bool)`

GetNewCommentOk returns a tuple with the NewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewComment

`func (o *UpdateTarget) SetNewComment(v string)`

SetNewComment sets NewComment field to given value.

### HasNewComment

`func (o *UpdateTarget) HasNewComment() bool`

HasNewComment returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


