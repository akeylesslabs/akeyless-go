# GatewayUpdateProducerArtifactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactoryAdminName** | Pointer to **string** | Artifactory Admin Name | [optional] 
**ArtifactoryAdminPwd** | Pointer to **string** | Artifactory Admin password | [optional] 
**ArtifactoryTokenAudience** | **string** | Token Audience | 
**ArtifactoryTokenScope** | **string** | Token Scope | 
**BaseUrl** | Pointer to **string** | Base URL | [optional] 
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayUpdateProducerArtifactory

`func NewGatewayUpdateProducerArtifactory(artifactoryTokenAudience string, artifactoryTokenScope string, name string, ) *GatewayUpdateProducerArtifactory`

NewGatewayUpdateProducerArtifactory instantiates a new GatewayUpdateProducerArtifactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerArtifactoryWithDefaults

`func NewGatewayUpdateProducerArtifactoryWithDefaults() *GatewayUpdateProducerArtifactory`

NewGatewayUpdateProducerArtifactoryWithDefaults instantiates a new GatewayUpdateProducerArtifactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactoryAdminName

`func (o *GatewayUpdateProducerArtifactory) GetArtifactoryAdminName() string`

GetArtifactoryAdminName returns the ArtifactoryAdminName field if non-nil, zero value otherwise.

### GetArtifactoryAdminNameOk

`func (o *GatewayUpdateProducerArtifactory) GetArtifactoryAdminNameOk() (*string, bool)`

GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminName

`func (o *GatewayUpdateProducerArtifactory) SetArtifactoryAdminName(v string)`

SetArtifactoryAdminName sets ArtifactoryAdminName field to given value.

### HasArtifactoryAdminName

`func (o *GatewayUpdateProducerArtifactory) HasArtifactoryAdminName() bool`

HasArtifactoryAdminName returns a boolean if a field has been set.

### GetArtifactoryAdminPwd

`func (o *GatewayUpdateProducerArtifactory) GetArtifactoryAdminPwd() string`

GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field if non-nil, zero value otherwise.

### GetArtifactoryAdminPwdOk

`func (o *GatewayUpdateProducerArtifactory) GetArtifactoryAdminPwdOk() (*string, bool)`

GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminPwd

`func (o *GatewayUpdateProducerArtifactory) SetArtifactoryAdminPwd(v string)`

SetArtifactoryAdminPwd sets ArtifactoryAdminPwd field to given value.

### HasArtifactoryAdminPwd

`func (o *GatewayUpdateProducerArtifactory) HasArtifactoryAdminPwd() bool`

HasArtifactoryAdminPwd returns a boolean if a field has been set.

### GetArtifactoryTokenAudience

`func (o *GatewayUpdateProducerArtifactory) GetArtifactoryTokenAudience() string`

GetArtifactoryTokenAudience returns the ArtifactoryTokenAudience field if non-nil, zero value otherwise.

### GetArtifactoryTokenAudienceOk

`func (o *GatewayUpdateProducerArtifactory) GetArtifactoryTokenAudienceOk() (*string, bool)`

GetArtifactoryTokenAudienceOk returns a tuple with the ArtifactoryTokenAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryTokenAudience

`func (o *GatewayUpdateProducerArtifactory) SetArtifactoryTokenAudience(v string)`

SetArtifactoryTokenAudience sets ArtifactoryTokenAudience field to given value.


### GetArtifactoryTokenScope

`func (o *GatewayUpdateProducerArtifactory) GetArtifactoryTokenScope() string`

GetArtifactoryTokenScope returns the ArtifactoryTokenScope field if non-nil, zero value otherwise.

### GetArtifactoryTokenScopeOk

`func (o *GatewayUpdateProducerArtifactory) GetArtifactoryTokenScopeOk() (*string, bool)`

GetArtifactoryTokenScopeOk returns a tuple with the ArtifactoryTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryTokenScope

`func (o *GatewayUpdateProducerArtifactory) SetArtifactoryTokenScope(v string)`

SetArtifactoryTokenScope sets ArtifactoryTokenScope field to given value.


### GetBaseUrl

`func (o *GatewayUpdateProducerArtifactory) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *GatewayUpdateProducerArtifactory) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *GatewayUpdateProducerArtifactory) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *GatewayUpdateProducerArtifactory) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerArtifactory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerArtifactory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerArtifactory) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerArtifactory) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerArtifactory) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerArtifactory) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerArtifactory) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *GatewayUpdateProducerArtifactory) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayUpdateProducerArtifactory) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayUpdateProducerArtifactory) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayUpdateProducerArtifactory) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerArtifactory) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerArtifactory) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerArtifactory) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerArtifactory) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerArtifactory) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerArtifactory) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerArtifactory) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerArtifactory) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerArtifactory) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerArtifactory) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerArtifactory) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerArtifactory) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerArtifactory) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerArtifactory) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerArtifactory) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerArtifactory) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerArtifactory) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerArtifactory) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerArtifactory) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerArtifactory) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerArtifactory) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerArtifactory) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerArtifactory) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerArtifactory) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayUpdateProducerArtifactory) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayUpdateProducerArtifactory) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayUpdateProducerArtifactory) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayUpdateProducerArtifactory) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


