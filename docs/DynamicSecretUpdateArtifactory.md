# DynamicSecretUpdateArtifactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactoryAdminName** | Pointer to **string** | Artifactory Admin Name | [optional] 
**ArtifactoryAdminPwd** | Pointer to **string** | Artifactory Admin password | [optional] 
**ArtifactoryTokenAudience** | **string** | Token Audience | 
**ArtifactoryTokenScope** | **string** | Token Scope | 
**BaseUrl** | Pointer to **string** | Base URL | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretUpdateArtifactory

`func NewDynamicSecretUpdateArtifactory(artifactoryTokenAudience string, artifactoryTokenScope string, name string, ) *DynamicSecretUpdateArtifactory`

NewDynamicSecretUpdateArtifactory instantiates a new DynamicSecretUpdateArtifactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateArtifactoryWithDefaults

`func NewDynamicSecretUpdateArtifactoryWithDefaults() *DynamicSecretUpdateArtifactory`

NewDynamicSecretUpdateArtifactoryWithDefaults instantiates a new DynamicSecretUpdateArtifactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactoryAdminName

`func (o *DynamicSecretUpdateArtifactory) GetArtifactoryAdminName() string`

GetArtifactoryAdminName returns the ArtifactoryAdminName field if non-nil, zero value otherwise.

### GetArtifactoryAdminNameOk

`func (o *DynamicSecretUpdateArtifactory) GetArtifactoryAdminNameOk() (*string, bool)`

GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminName

`func (o *DynamicSecretUpdateArtifactory) SetArtifactoryAdminName(v string)`

SetArtifactoryAdminName sets ArtifactoryAdminName field to given value.

### HasArtifactoryAdminName

`func (o *DynamicSecretUpdateArtifactory) HasArtifactoryAdminName() bool`

HasArtifactoryAdminName returns a boolean if a field has been set.

### GetArtifactoryAdminPwd

`func (o *DynamicSecretUpdateArtifactory) GetArtifactoryAdminPwd() string`

GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field if non-nil, zero value otherwise.

### GetArtifactoryAdminPwdOk

`func (o *DynamicSecretUpdateArtifactory) GetArtifactoryAdminPwdOk() (*string, bool)`

GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminPwd

`func (o *DynamicSecretUpdateArtifactory) SetArtifactoryAdminPwd(v string)`

SetArtifactoryAdminPwd sets ArtifactoryAdminPwd field to given value.

### HasArtifactoryAdminPwd

`func (o *DynamicSecretUpdateArtifactory) HasArtifactoryAdminPwd() bool`

HasArtifactoryAdminPwd returns a boolean if a field has been set.

### GetArtifactoryTokenAudience

`func (o *DynamicSecretUpdateArtifactory) GetArtifactoryTokenAudience() string`

GetArtifactoryTokenAudience returns the ArtifactoryTokenAudience field if non-nil, zero value otherwise.

### GetArtifactoryTokenAudienceOk

`func (o *DynamicSecretUpdateArtifactory) GetArtifactoryTokenAudienceOk() (*string, bool)`

GetArtifactoryTokenAudienceOk returns a tuple with the ArtifactoryTokenAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryTokenAudience

`func (o *DynamicSecretUpdateArtifactory) SetArtifactoryTokenAudience(v string)`

SetArtifactoryTokenAudience sets ArtifactoryTokenAudience field to given value.


### GetArtifactoryTokenScope

`func (o *DynamicSecretUpdateArtifactory) GetArtifactoryTokenScope() string`

GetArtifactoryTokenScope returns the ArtifactoryTokenScope field if non-nil, zero value otherwise.

### GetArtifactoryTokenScopeOk

`func (o *DynamicSecretUpdateArtifactory) GetArtifactoryTokenScopeOk() (*string, bool)`

GetArtifactoryTokenScopeOk returns a tuple with the ArtifactoryTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryTokenScope

`func (o *DynamicSecretUpdateArtifactory) SetArtifactoryTokenScope(v string)`

SetArtifactoryTokenScope sets ArtifactoryTokenScope field to given value.


### GetBaseUrl

`func (o *DynamicSecretUpdateArtifactory) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DynamicSecretUpdateArtifactory) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DynamicSecretUpdateArtifactory) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *DynamicSecretUpdateArtifactory) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretUpdateArtifactory) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateArtifactory) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateArtifactory) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateArtifactory) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateArtifactory) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateArtifactory) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateArtifactory) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateArtifactory) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateArtifactory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateArtifactory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateArtifactory) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateArtifactory) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateArtifactory) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateArtifactory) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateArtifactory) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateArtifactory) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateArtifactory) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateArtifactory) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateArtifactory) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateArtifactory) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateArtifactory) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateArtifactory) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateArtifactory) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateArtifactory) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateArtifactory) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateArtifactory) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateArtifactory) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateArtifactory) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateArtifactory) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateArtifactory) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateArtifactory) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateArtifactory) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateArtifactory) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateArtifactory) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateArtifactory) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateArtifactory) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateArtifactory) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateArtifactory) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateArtifactory) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


