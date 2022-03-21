# CreateWebTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | Pointer to **string** | The url | [optional] 

## Methods

### NewCreateWebTarget

`func NewCreateWebTarget(name string, ) *CreateWebTarget`

NewCreateWebTarget instantiates a new CreateWebTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebTargetWithDefaults

`func NewCreateWebTargetWithDefaults() *CreateWebTarget`

NewCreateWebTargetWithDefaults instantiates a new CreateWebTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateWebTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateWebTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateWebTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateWebTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetKey

`func (o *CreateWebTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateWebTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateWebTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateWebTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateWebTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebTarget) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *CreateWebTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateWebTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateWebTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateWebTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateWebTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateWebTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateWebTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateWebTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *CreateWebTarget) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebTarget) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebTarget) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateWebTarget) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


