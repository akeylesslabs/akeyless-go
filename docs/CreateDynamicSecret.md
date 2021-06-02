# CreateDynamicSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The name of a key that used to encrypt the dynamic secret values (if empty, the account default protectionKey key will be used) | [optional] 
**Metadata** | Pointer to **string** | Metadata about the dynamic secret | [optional] [default to "None"]
**Name** | **string** | Dynamic secret name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreateDynamicSecret

`func NewCreateDynamicSecret(name string, ) *CreateDynamicSecret`

NewCreateDynamicSecret instantiates a new CreateDynamicSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDynamicSecretWithDefaults

`func NewCreateDynamicSecretWithDefaults() *CreateDynamicSecret`

NewCreateDynamicSecretWithDefaults instantiates a new CreateDynamicSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateDynamicSecret) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateDynamicSecret) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateDynamicSecret) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateDynamicSecret) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateDynamicSecret) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateDynamicSecret) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateDynamicSecret) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateDynamicSecret) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateDynamicSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDynamicSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDynamicSecret) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateDynamicSecret) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDynamicSecret) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDynamicSecret) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateDynamicSecret) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTags

`func (o *CreateDynamicSecret) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDynamicSecret) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDynamicSecret) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDynamicSecret) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *CreateDynamicSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateDynamicSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateDynamicSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateDynamicSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateDynamicSecret) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateDynamicSecret) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateDynamicSecret) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateDynamicSecret) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateDynamicSecret) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateDynamicSecret) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateDynamicSecret) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateDynamicSecret) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


