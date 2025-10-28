# DynamicSecretUpdateGitlab

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
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | Pointer to **string** | Access Token TTL | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDynamicSecretUpdateGitlab

`func NewDynamicSecretUpdateGitlab(gitlabAccessType string, gitlabTokenScopes string, name string, ) *DynamicSecretUpdateGitlab`

NewDynamicSecretUpdateGitlab instantiates a new DynamicSecretUpdateGitlab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateGitlabWithDefaults

`func NewDynamicSecretUpdateGitlabWithDefaults() *DynamicSecretUpdateGitlab`

NewDynamicSecretUpdateGitlabWithDefaults instantiates a new DynamicSecretUpdateGitlab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretUpdateGitlab) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateGitlab) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateGitlab) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateGitlab) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateGitlab) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateGitlab) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateGitlab) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateGitlab) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGitlabAccessToken

`func (o *DynamicSecretUpdateGitlab) GetGitlabAccessToken() string`

GetGitlabAccessToken returns the GitlabAccessToken field if non-nil, zero value otherwise.

### GetGitlabAccessTokenOk

`func (o *DynamicSecretUpdateGitlab) GetGitlabAccessTokenOk() (*string, bool)`

GetGitlabAccessTokenOk returns a tuple with the GitlabAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabAccessToken

`func (o *DynamicSecretUpdateGitlab) SetGitlabAccessToken(v string)`

SetGitlabAccessToken sets GitlabAccessToken field to given value.

### HasGitlabAccessToken

`func (o *DynamicSecretUpdateGitlab) HasGitlabAccessToken() bool`

HasGitlabAccessToken returns a boolean if a field has been set.

### GetGitlabAccessType

`func (o *DynamicSecretUpdateGitlab) GetGitlabAccessType() string`

GetGitlabAccessType returns the GitlabAccessType field if non-nil, zero value otherwise.

### GetGitlabAccessTypeOk

`func (o *DynamicSecretUpdateGitlab) GetGitlabAccessTypeOk() (*string, bool)`

GetGitlabAccessTypeOk returns a tuple with the GitlabAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabAccessType

`func (o *DynamicSecretUpdateGitlab) SetGitlabAccessType(v string)`

SetGitlabAccessType sets GitlabAccessType field to given value.


### GetGitlabCertificate

`func (o *DynamicSecretUpdateGitlab) GetGitlabCertificate() string`

GetGitlabCertificate returns the GitlabCertificate field if non-nil, zero value otherwise.

### GetGitlabCertificateOk

`func (o *DynamicSecretUpdateGitlab) GetGitlabCertificateOk() (*string, bool)`

GetGitlabCertificateOk returns a tuple with the GitlabCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabCertificate

`func (o *DynamicSecretUpdateGitlab) SetGitlabCertificate(v string)`

SetGitlabCertificate sets GitlabCertificate field to given value.

### HasGitlabCertificate

`func (o *DynamicSecretUpdateGitlab) HasGitlabCertificate() bool`

HasGitlabCertificate returns a boolean if a field has been set.

### GetGitlabRole

`func (o *DynamicSecretUpdateGitlab) GetGitlabRole() string`

GetGitlabRole returns the GitlabRole field if non-nil, zero value otherwise.

### GetGitlabRoleOk

`func (o *DynamicSecretUpdateGitlab) GetGitlabRoleOk() (*string, bool)`

GetGitlabRoleOk returns a tuple with the GitlabRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabRole

`func (o *DynamicSecretUpdateGitlab) SetGitlabRole(v string)`

SetGitlabRole sets GitlabRole field to given value.

### HasGitlabRole

`func (o *DynamicSecretUpdateGitlab) HasGitlabRole() bool`

HasGitlabRole returns a boolean if a field has been set.

### GetGitlabTokenScopes

`func (o *DynamicSecretUpdateGitlab) GetGitlabTokenScopes() string`

GetGitlabTokenScopes returns the GitlabTokenScopes field if non-nil, zero value otherwise.

### GetGitlabTokenScopesOk

`func (o *DynamicSecretUpdateGitlab) GetGitlabTokenScopesOk() (*string, bool)`

GetGitlabTokenScopesOk returns a tuple with the GitlabTokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabTokenScopes

`func (o *DynamicSecretUpdateGitlab) SetGitlabTokenScopes(v string)`

SetGitlabTokenScopes sets GitlabTokenScopes field to given value.


### GetGitlabUrl

`func (o *DynamicSecretUpdateGitlab) GetGitlabUrl() string`

GetGitlabUrl returns the GitlabUrl field if non-nil, zero value otherwise.

### GetGitlabUrlOk

`func (o *DynamicSecretUpdateGitlab) GetGitlabUrlOk() (*string, bool)`

GetGitlabUrlOk returns a tuple with the GitlabUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabUrl

`func (o *DynamicSecretUpdateGitlab) SetGitlabUrl(v string)`

SetGitlabUrl sets GitlabUrl field to given value.

### HasGitlabUrl

`func (o *DynamicSecretUpdateGitlab) HasGitlabUrl() bool`

HasGitlabUrl returns a boolean if a field has been set.

### GetGroupName

`func (o *DynamicSecretUpdateGitlab) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DynamicSecretUpdateGitlab) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DynamicSecretUpdateGitlab) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DynamicSecretUpdateGitlab) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetInstallationOrganization

`func (o *DynamicSecretUpdateGitlab) GetInstallationOrganization() string`

GetInstallationOrganization returns the InstallationOrganization field if non-nil, zero value otherwise.

### GetInstallationOrganizationOk

`func (o *DynamicSecretUpdateGitlab) GetInstallationOrganizationOk() (*string, bool)`

GetInstallationOrganizationOk returns a tuple with the InstallationOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationOrganization

`func (o *DynamicSecretUpdateGitlab) SetInstallationOrganization(v string)`

SetInstallationOrganization sets InstallationOrganization field to given value.

### HasInstallationOrganization

`func (o *DynamicSecretUpdateGitlab) HasInstallationOrganization() bool`

HasInstallationOrganization returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *DynamicSecretUpdateGitlab) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *DynamicSecretUpdateGitlab) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *DynamicSecretUpdateGitlab) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *DynamicSecretUpdateGitlab) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateGitlab) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateGitlab) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateGitlab) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateGitlab) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateGitlab) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateGitlab) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateGitlab) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateGitlab) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateGitlab) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateGitlab) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateGitlab) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateGitlab) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateGitlab) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateGitlab) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateGitlab) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateGitlab) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateGitlab) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateGitlab) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateGitlab) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateGitlab) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateGitlab) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateGitlab) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateGitlab) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *DynamicSecretUpdateGitlab) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DynamicSecretUpdateGitlab) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DynamicSecretUpdateGitlab) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DynamicSecretUpdateGitlab) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateGitlab) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateGitlab) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateGitlab) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateGitlab) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


