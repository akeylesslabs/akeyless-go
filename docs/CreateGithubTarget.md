# CreateGithubTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**GithubAppId** | Pointer to **int64** | Github app id | [optional] 
**GithubAppPrivateKey** | Pointer to **string** | App private key | [optional] 
**GithubBaseUrl** | Pointer to **string** | Base URL | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateGithubTarget

`func NewCreateGithubTarget(name string, ) *CreateGithubTarget`

NewCreateGithubTarget instantiates a new CreateGithubTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGithubTargetWithDefaults

`func NewCreateGithubTargetWithDefaults() *CreateGithubTarget`

NewCreateGithubTargetWithDefaults instantiates a new CreateGithubTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateGithubTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateGithubTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateGithubTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateGithubTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGithubAppId

`func (o *CreateGithubTarget) GetGithubAppId() int64`

GetGithubAppId returns the GithubAppId field if non-nil, zero value otherwise.

### GetGithubAppIdOk

`func (o *CreateGithubTarget) GetGithubAppIdOk() (*int64, bool)`

GetGithubAppIdOk returns a tuple with the GithubAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppId

`func (o *CreateGithubTarget) SetGithubAppId(v int64)`

SetGithubAppId sets GithubAppId field to given value.

### HasGithubAppId

`func (o *CreateGithubTarget) HasGithubAppId() bool`

HasGithubAppId returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *CreateGithubTarget) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *CreateGithubTarget) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *CreateGithubTarget) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *CreateGithubTarget) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetGithubBaseUrl

`func (o *CreateGithubTarget) GetGithubBaseUrl() string`

GetGithubBaseUrl returns the GithubBaseUrl field if non-nil, zero value otherwise.

### GetGithubBaseUrlOk

`func (o *CreateGithubTarget) GetGithubBaseUrlOk() (*string, bool)`

GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubBaseUrl

`func (o *CreateGithubTarget) SetGithubBaseUrl(v string)`

SetGithubBaseUrl sets GithubBaseUrl field to given value.

### HasGithubBaseUrl

`func (o *CreateGithubTarget) HasGithubBaseUrl() bool`

HasGithubBaseUrl returns a boolean if a field has been set.

### GetKey

`func (o *CreateGithubTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateGithubTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateGithubTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateGithubTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateGithubTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGithubTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGithubTarget) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *CreateGithubTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateGithubTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateGithubTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateGithubTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateGithubTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateGithubTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateGithubTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateGithubTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


