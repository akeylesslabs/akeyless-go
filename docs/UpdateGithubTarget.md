# UpdateGithubTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GithubAppId** | Pointer to **int64** | Github app id | [optional] 
**GithubAppPrivateKey** | Pointer to **string** | App private key | [optional] 
**GithubBaseUrl** | Pointer to **string** | Base URL | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**KeepPrevVersion** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 

## Methods

### NewUpdateGithubTarget

`func NewUpdateGithubTarget(name string, ) *UpdateGithubTarget`

NewUpdateGithubTarget instantiates a new UpdateGithubTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGithubTargetWithDefaults

`func NewUpdateGithubTargetWithDefaults() *UpdateGithubTarget`

NewUpdateGithubTargetWithDefaults instantiates a new UpdateGithubTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateGithubTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateGithubTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateGithubTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateGithubTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGithubTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGithubTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGithubTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGithubTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGithubAppId

`func (o *UpdateGithubTarget) GetGithubAppId() int64`

GetGithubAppId returns the GithubAppId field if non-nil, zero value otherwise.

### GetGithubAppIdOk

`func (o *UpdateGithubTarget) GetGithubAppIdOk() (*int64, bool)`

GetGithubAppIdOk returns a tuple with the GithubAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppId

`func (o *UpdateGithubTarget) SetGithubAppId(v int64)`

SetGithubAppId sets GithubAppId field to given value.

### HasGithubAppId

`func (o *UpdateGithubTarget) HasGithubAppId() bool`

HasGithubAppId returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *UpdateGithubTarget) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *UpdateGithubTarget) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *UpdateGithubTarget) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *UpdateGithubTarget) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetGithubBaseUrl

`func (o *UpdateGithubTarget) GetGithubBaseUrl() string`

GetGithubBaseUrl returns the GithubBaseUrl field if non-nil, zero value otherwise.

### GetGithubBaseUrlOk

`func (o *UpdateGithubTarget) GetGithubBaseUrlOk() (*string, bool)`

GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubBaseUrl

`func (o *UpdateGithubTarget) SetGithubBaseUrl(v string)`

SetGithubBaseUrl sets GithubBaseUrl field to given value.

### HasGithubBaseUrl

`func (o *UpdateGithubTarget) HasGithubBaseUrl() bool`

HasGithubBaseUrl returns a boolean if a field has been set.

### GetJson

`func (o *UpdateGithubTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateGithubTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateGithubTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateGithubTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateGithubTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateGithubTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateGithubTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateGithubTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateGithubTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateGithubTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateGithubTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateGithubTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateGithubTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGithubTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGithubTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateGithubTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateGithubTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateGithubTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateGithubTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateGithubTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateGithubTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateGithubTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateGithubTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateGithubTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateGithubTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateGithubTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateGithubTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateGithubTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateGithubTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateGithubTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateGithubTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


