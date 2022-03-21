# CreateArtifactoryTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactoryAdminName** | **string** | Artifactory Admin Name | 
**ArtifactoryAdminPwd** | **string** | Artifactory Admin password | 
**BaseUrl** | **string** | Base URL | 
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateArtifactoryTarget

`func NewCreateArtifactoryTarget(artifactoryAdminName string, artifactoryAdminPwd string, baseUrl string, name string, ) *CreateArtifactoryTarget`

NewCreateArtifactoryTarget instantiates a new CreateArtifactoryTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateArtifactoryTargetWithDefaults

`func NewCreateArtifactoryTargetWithDefaults() *CreateArtifactoryTarget`

NewCreateArtifactoryTargetWithDefaults instantiates a new CreateArtifactoryTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactoryAdminName

`func (o *CreateArtifactoryTarget) GetArtifactoryAdminName() string`

GetArtifactoryAdminName returns the ArtifactoryAdminName field if non-nil, zero value otherwise.

### GetArtifactoryAdminNameOk

`func (o *CreateArtifactoryTarget) GetArtifactoryAdminNameOk() (*string, bool)`

GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminName

`func (o *CreateArtifactoryTarget) SetArtifactoryAdminName(v string)`

SetArtifactoryAdminName sets ArtifactoryAdminName field to given value.


### GetArtifactoryAdminPwd

`func (o *CreateArtifactoryTarget) GetArtifactoryAdminPwd() string`

GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field if non-nil, zero value otherwise.

### GetArtifactoryAdminPwdOk

`func (o *CreateArtifactoryTarget) GetArtifactoryAdminPwdOk() (*string, bool)`

GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminPwd

`func (o *CreateArtifactoryTarget) SetArtifactoryAdminPwd(v string)`

SetArtifactoryAdminPwd sets ArtifactoryAdminPwd field to given value.


### GetBaseUrl

`func (o *CreateArtifactoryTarget) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *CreateArtifactoryTarget) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *CreateArtifactoryTarget) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetComment

`func (o *CreateArtifactoryTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateArtifactoryTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateArtifactoryTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateArtifactoryTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetKey

`func (o *CreateArtifactoryTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateArtifactoryTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateArtifactoryTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateArtifactoryTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateArtifactoryTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateArtifactoryTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateArtifactoryTarget) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *CreateArtifactoryTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateArtifactoryTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateArtifactoryTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateArtifactoryTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateArtifactoryTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateArtifactoryTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateArtifactoryTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateArtifactoryTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


