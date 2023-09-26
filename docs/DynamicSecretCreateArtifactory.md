# DynamicSecretCreateArtifactory

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
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateArtifactory

`func NewDynamicSecretCreateArtifactory(artifactoryTokenAudience string, artifactoryTokenScope string, name string, ) *DynamicSecretCreateArtifactory`

NewDynamicSecretCreateArtifactory instantiates a new DynamicSecretCreateArtifactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateArtifactoryWithDefaults

`func NewDynamicSecretCreateArtifactoryWithDefaults() *DynamicSecretCreateArtifactory`

NewDynamicSecretCreateArtifactoryWithDefaults instantiates a new DynamicSecretCreateArtifactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactoryAdminName

`func (o *DynamicSecretCreateArtifactory) GetArtifactoryAdminName() string`

GetArtifactoryAdminName returns the ArtifactoryAdminName field if non-nil, zero value otherwise.

### GetArtifactoryAdminNameOk

`func (o *DynamicSecretCreateArtifactory) GetArtifactoryAdminNameOk() (*string, bool)`

GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminName

`func (o *DynamicSecretCreateArtifactory) SetArtifactoryAdminName(v string)`

SetArtifactoryAdminName sets ArtifactoryAdminName field to given value.

### HasArtifactoryAdminName

`func (o *DynamicSecretCreateArtifactory) HasArtifactoryAdminName() bool`

HasArtifactoryAdminName returns a boolean if a field has been set.

### GetArtifactoryAdminPwd

`func (o *DynamicSecretCreateArtifactory) GetArtifactoryAdminPwd() string`

GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field if non-nil, zero value otherwise.

### GetArtifactoryAdminPwdOk

`func (o *DynamicSecretCreateArtifactory) GetArtifactoryAdminPwdOk() (*string, bool)`

GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminPwd

`func (o *DynamicSecretCreateArtifactory) SetArtifactoryAdminPwd(v string)`

SetArtifactoryAdminPwd sets ArtifactoryAdminPwd field to given value.

### HasArtifactoryAdminPwd

`func (o *DynamicSecretCreateArtifactory) HasArtifactoryAdminPwd() bool`

HasArtifactoryAdminPwd returns a boolean if a field has been set.

### GetArtifactoryTokenAudience

`func (o *DynamicSecretCreateArtifactory) GetArtifactoryTokenAudience() string`

GetArtifactoryTokenAudience returns the ArtifactoryTokenAudience field if non-nil, zero value otherwise.

### GetArtifactoryTokenAudienceOk

`func (o *DynamicSecretCreateArtifactory) GetArtifactoryTokenAudienceOk() (*string, bool)`

GetArtifactoryTokenAudienceOk returns a tuple with the ArtifactoryTokenAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryTokenAudience

`func (o *DynamicSecretCreateArtifactory) SetArtifactoryTokenAudience(v string)`

SetArtifactoryTokenAudience sets ArtifactoryTokenAudience field to given value.


### GetArtifactoryTokenScope

`func (o *DynamicSecretCreateArtifactory) GetArtifactoryTokenScope() string`

GetArtifactoryTokenScope returns the ArtifactoryTokenScope field if non-nil, zero value otherwise.

### GetArtifactoryTokenScopeOk

`func (o *DynamicSecretCreateArtifactory) GetArtifactoryTokenScopeOk() (*string, bool)`

GetArtifactoryTokenScopeOk returns a tuple with the ArtifactoryTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryTokenScope

`func (o *DynamicSecretCreateArtifactory) SetArtifactoryTokenScope(v string)`

SetArtifactoryTokenScope sets ArtifactoryTokenScope field to given value.


### GetBaseUrl

`func (o *DynamicSecretCreateArtifactory) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DynamicSecretCreateArtifactory) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DynamicSecretCreateArtifactory) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *DynamicSecretCreateArtifactory) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateArtifactory) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateArtifactory) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateArtifactory) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateArtifactory) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateArtifactory) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateArtifactory) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateArtifactory) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateArtifactory) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateArtifactory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateArtifactory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateArtifactory) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateArtifactory) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateArtifactory) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateArtifactory) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateArtifactory) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateArtifactory) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateArtifactory) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateArtifactory) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateArtifactory) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateArtifactory) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateArtifactory) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateArtifactory) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateArtifactory) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateArtifactory) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateArtifactory) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateArtifactory) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateArtifactory) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateArtifactory) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateArtifactory) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateArtifactory) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateArtifactory) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateArtifactory) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateArtifactory) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateArtifactory) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateArtifactory) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


