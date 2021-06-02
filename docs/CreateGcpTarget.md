# CreateGcpTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**GcpKey** | Pointer to **string** | Base64-encoded service account private key text | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreateGcpTarget

`func NewCreateGcpTarget(name string, ) *CreateGcpTarget`

NewCreateGcpTarget instantiates a new CreateGcpTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGcpTargetWithDefaults

`func NewCreateGcpTargetWithDefaults() *CreateGcpTarget`

NewCreateGcpTargetWithDefaults instantiates a new CreateGcpTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateGcpTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateGcpTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateGcpTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateGcpTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGcpKey

`func (o *CreateGcpTarget) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *CreateGcpTarget) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *CreateGcpTarget) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *CreateGcpTarget) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetKey

`func (o *CreateGcpTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateGcpTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateGcpTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateGcpTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateGcpTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGcpTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGcpTarget) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateGcpTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateGcpTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateGcpTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateGcpTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *CreateGcpTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateGcpTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateGcpTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateGcpTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateGcpTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateGcpTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateGcpTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateGcpTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateGcpTarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateGcpTarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateGcpTarget) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateGcpTarget) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


