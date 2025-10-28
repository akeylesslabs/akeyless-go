# DynamicSecretCreateGitlab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GitlabAccessToken** | Pointer to **string** | Gitlab access token | [optional] 
**GitlabAccessType** | **string** | Gitlab access token type [project,group] | 
**GitlabCertificate** | Pointer to **string** | Gitlab tls certificate (base64 encoded) | [optional] 
**GitlabRole** | Pointer to **string** | Gitlab role | [optional] 
**GitlabTokenScopes** | **string** | Comma-separated list of access token scopes to grant | 
**GitlabUrl** | Pointer to **string** | Gitlab base url | [optional] [default to "https://gitlab.com/"]
**GroupName** | Pointer to **string** | Gitlab group name, required for access-type&#x3D;group | [optional] 
**InstallationOrganization** | Pointer to **string** | Gitlab project name, required for access-type&#x3D;project | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | Pointer to **string** | Access Token TTL | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDynamicSecretCreateGitlab

`func NewDynamicSecretCreateGitlab(gitlabAccessType string, gitlabTokenScopes string, name string, ) *DynamicSecretCreateGitlab`

NewDynamicSecretCreateGitlab instantiates a new DynamicSecretCreateGitlab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateGitlabWithDefaults

`func NewDynamicSecretCreateGitlabWithDefaults() *DynamicSecretCreateGitlab`

NewDynamicSecretCreateGitlabWithDefaults instantiates a new DynamicSecretCreateGitlab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretCreateGitlab) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateGitlab) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateGitlab) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateGitlab) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateGitlab) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateGitlab) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateGitlab) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateGitlab) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGitlabAccessToken

`func (o *DynamicSecretCreateGitlab) GetGitlabAccessToken() string`

GetGitlabAccessToken returns the GitlabAccessToken field if non-nil, zero value otherwise.

### GetGitlabAccessTokenOk

`func (o *DynamicSecretCreateGitlab) GetGitlabAccessTokenOk() (*string, bool)`

GetGitlabAccessTokenOk returns a tuple with the GitlabAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabAccessToken

`func (o *DynamicSecretCreateGitlab) SetGitlabAccessToken(v string)`

SetGitlabAccessToken sets GitlabAccessToken field to given value.

### HasGitlabAccessToken

`func (o *DynamicSecretCreateGitlab) HasGitlabAccessToken() bool`

HasGitlabAccessToken returns a boolean if a field has been set.

### GetGitlabAccessType

`func (o *DynamicSecretCreateGitlab) GetGitlabAccessType() string`

GetGitlabAccessType returns the GitlabAccessType field if non-nil, zero value otherwise.

### GetGitlabAccessTypeOk

`func (o *DynamicSecretCreateGitlab) GetGitlabAccessTypeOk() (*string, bool)`

GetGitlabAccessTypeOk returns a tuple with the GitlabAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabAccessType

`func (o *DynamicSecretCreateGitlab) SetGitlabAccessType(v string)`

SetGitlabAccessType sets GitlabAccessType field to given value.


### GetGitlabCertificate

`func (o *DynamicSecretCreateGitlab) GetGitlabCertificate() string`

GetGitlabCertificate returns the GitlabCertificate field if non-nil, zero value otherwise.

### GetGitlabCertificateOk

`func (o *DynamicSecretCreateGitlab) GetGitlabCertificateOk() (*string, bool)`

GetGitlabCertificateOk returns a tuple with the GitlabCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabCertificate

`func (o *DynamicSecretCreateGitlab) SetGitlabCertificate(v string)`

SetGitlabCertificate sets GitlabCertificate field to given value.

### HasGitlabCertificate

`func (o *DynamicSecretCreateGitlab) HasGitlabCertificate() bool`

HasGitlabCertificate returns a boolean if a field has been set.

### GetGitlabRole

`func (o *DynamicSecretCreateGitlab) GetGitlabRole() string`

GetGitlabRole returns the GitlabRole field if non-nil, zero value otherwise.

### GetGitlabRoleOk

`func (o *DynamicSecretCreateGitlab) GetGitlabRoleOk() (*string, bool)`

GetGitlabRoleOk returns a tuple with the GitlabRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabRole

`func (o *DynamicSecretCreateGitlab) SetGitlabRole(v string)`

SetGitlabRole sets GitlabRole field to given value.

### HasGitlabRole

`func (o *DynamicSecretCreateGitlab) HasGitlabRole() bool`

HasGitlabRole returns a boolean if a field has been set.

### GetGitlabTokenScopes

`func (o *DynamicSecretCreateGitlab) GetGitlabTokenScopes() string`

GetGitlabTokenScopes returns the GitlabTokenScopes field if non-nil, zero value otherwise.

### GetGitlabTokenScopesOk

`func (o *DynamicSecretCreateGitlab) GetGitlabTokenScopesOk() (*string, bool)`

GetGitlabTokenScopesOk returns a tuple with the GitlabTokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabTokenScopes

`func (o *DynamicSecretCreateGitlab) SetGitlabTokenScopes(v string)`

SetGitlabTokenScopes sets GitlabTokenScopes field to given value.


### GetGitlabUrl

`func (o *DynamicSecretCreateGitlab) GetGitlabUrl() string`

GetGitlabUrl returns the GitlabUrl field if non-nil, zero value otherwise.

### GetGitlabUrlOk

`func (o *DynamicSecretCreateGitlab) GetGitlabUrlOk() (*string, bool)`

GetGitlabUrlOk returns a tuple with the GitlabUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabUrl

`func (o *DynamicSecretCreateGitlab) SetGitlabUrl(v string)`

SetGitlabUrl sets GitlabUrl field to given value.

### HasGitlabUrl

`func (o *DynamicSecretCreateGitlab) HasGitlabUrl() bool`

HasGitlabUrl returns a boolean if a field has been set.

### GetGroupName

`func (o *DynamicSecretCreateGitlab) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DynamicSecretCreateGitlab) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DynamicSecretCreateGitlab) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DynamicSecretCreateGitlab) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetInstallationOrganization

`func (o *DynamicSecretCreateGitlab) GetInstallationOrganization() string`

GetInstallationOrganization returns the InstallationOrganization field if non-nil, zero value otherwise.

### GetInstallationOrganizationOk

`func (o *DynamicSecretCreateGitlab) GetInstallationOrganizationOk() (*string, bool)`

GetInstallationOrganizationOk returns a tuple with the InstallationOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationOrganization

`func (o *DynamicSecretCreateGitlab) SetInstallationOrganization(v string)`

SetInstallationOrganization sets InstallationOrganization field to given value.

### HasInstallationOrganization

`func (o *DynamicSecretCreateGitlab) HasInstallationOrganization() bool`

HasInstallationOrganization returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *DynamicSecretCreateGitlab) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *DynamicSecretCreateGitlab) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *DynamicSecretCreateGitlab) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *DynamicSecretCreateGitlab) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateGitlab) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateGitlab) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateGitlab) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateGitlab) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateGitlab) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateGitlab) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateGitlab) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *DynamicSecretCreateGitlab) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateGitlab) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateGitlab) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateGitlab) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateGitlab) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateGitlab) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateGitlab) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateGitlab) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateGitlab) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateGitlab) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateGitlab) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateGitlab) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *DynamicSecretCreateGitlab) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DynamicSecretCreateGitlab) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DynamicSecretCreateGitlab) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DynamicSecretCreateGitlab) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateGitlab) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateGitlab) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateGitlab) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateGitlab) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


