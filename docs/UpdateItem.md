# UpdateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**Name** | **string** | Current item name | 
**NewMetadata** | Pointer to **string** | New item metadata | [optional] [default to "default_metadata"]
**NewName** | Pointer to **string** | New item name | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewUpdateItem

`func NewUpdateItem(name string, ) *UpdateItem`

NewUpdateItem instantiates a new UpdateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateItemWithDefaults

`func NewUpdateItemWithDefaults() *UpdateItem`

NewUpdateItemWithDefaults instantiates a new UpdateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *UpdateItem) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *UpdateItem) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *UpdateItem) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *UpdateItem) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetName

`func (o *UpdateItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateItem) SetName(v string)`

SetName sets Name field to given value.


### GetNewMetadata

`func (o *UpdateItem) GetNewMetadata() string`

GetNewMetadata returns the NewMetadata field if non-nil, zero value otherwise.

### GetNewMetadataOk

`func (o *UpdateItem) GetNewMetadataOk() (*string, bool)`

GetNewMetadataOk returns a tuple with the NewMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMetadata

`func (o *UpdateItem) SetNewMetadata(v string)`

SetNewMetadata sets NewMetadata field to given value.

### HasNewMetadata

`func (o *UpdateItem) HasNewMetadata() bool`

HasNewMetadata returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateItem) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateItem) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateItem) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateItem) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateItem) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateItem) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateItem) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateItem) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRmTag

`func (o *UpdateItem) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *UpdateItem) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *UpdateItem) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *UpdateItem) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetToken

`func (o *UpdateItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateItem) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateItem) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateItem) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateItem) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


