# DynamicSecretCreateGithub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GithubAppId** | Pointer to **int64** | Github app id | [optional] 
**GithubAppPrivateKey** | Pointer to **string** | App private key | [optional] 
**GithubBaseUrl** | Pointer to **string** | Base URL | [optional] [default to "https://api.github.com/"]
**InstallationId** | Pointer to **int64** | Github app installation id | [optional] 
**InstallationOrganization** | Pointer to **string** | Optional, instead of installation id, set a GitHub organization name | [optional] 
**InstallationRepository** | Pointer to **string** | Optional, instead of installation id, set a GitHub repository &#39;&lt;owner&gt;/&lt;repo-name&gt; | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenPermissions** | Pointer to **[]string** | Optional - installation token&#39;s allowed permissions | [optional] 
**TokenRepositories** | Pointer to **[]string** | Optional - installation token&#39;s allowed repositories | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDynamicSecretCreateGithub

`func NewDynamicSecretCreateGithub(name string, ) *DynamicSecretCreateGithub`

NewDynamicSecretCreateGithub instantiates a new DynamicSecretCreateGithub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateGithubWithDefaults

`func NewDynamicSecretCreateGithubWithDefaults() *DynamicSecretCreateGithub`

NewDynamicSecretCreateGithubWithDefaults instantiates a new DynamicSecretCreateGithub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretCreateGithub) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateGithub) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateGithub) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateGithub) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateGithub) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateGithub) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateGithub) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateGithub) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGithubAppId

`func (o *DynamicSecretCreateGithub) GetGithubAppId() int64`

GetGithubAppId returns the GithubAppId field if non-nil, zero value otherwise.

### GetGithubAppIdOk

`func (o *DynamicSecretCreateGithub) GetGithubAppIdOk() (*int64, bool)`

GetGithubAppIdOk returns a tuple with the GithubAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppId

`func (o *DynamicSecretCreateGithub) SetGithubAppId(v int64)`

SetGithubAppId sets GithubAppId field to given value.

### HasGithubAppId

`func (o *DynamicSecretCreateGithub) HasGithubAppId() bool`

HasGithubAppId returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *DynamicSecretCreateGithub) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *DynamicSecretCreateGithub) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *DynamicSecretCreateGithub) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *DynamicSecretCreateGithub) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetGithubBaseUrl

`func (o *DynamicSecretCreateGithub) GetGithubBaseUrl() string`

GetGithubBaseUrl returns the GithubBaseUrl field if non-nil, zero value otherwise.

### GetGithubBaseUrlOk

`func (o *DynamicSecretCreateGithub) GetGithubBaseUrlOk() (*string, bool)`

GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubBaseUrl

`func (o *DynamicSecretCreateGithub) SetGithubBaseUrl(v string)`

SetGithubBaseUrl sets GithubBaseUrl field to given value.

### HasGithubBaseUrl

`func (o *DynamicSecretCreateGithub) HasGithubBaseUrl() bool`

HasGithubBaseUrl returns a boolean if a field has been set.

### GetInstallationId

`func (o *DynamicSecretCreateGithub) GetInstallationId() int64`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *DynamicSecretCreateGithub) GetInstallationIdOk() (*int64, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *DynamicSecretCreateGithub) SetInstallationId(v int64)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *DynamicSecretCreateGithub) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetInstallationOrganization

`func (o *DynamicSecretCreateGithub) GetInstallationOrganization() string`

GetInstallationOrganization returns the InstallationOrganization field if non-nil, zero value otherwise.

### GetInstallationOrganizationOk

`func (o *DynamicSecretCreateGithub) GetInstallationOrganizationOk() (*string, bool)`

GetInstallationOrganizationOk returns a tuple with the InstallationOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationOrganization

`func (o *DynamicSecretCreateGithub) SetInstallationOrganization(v string)`

SetInstallationOrganization sets InstallationOrganization field to given value.

### HasInstallationOrganization

`func (o *DynamicSecretCreateGithub) HasInstallationOrganization() bool`

HasInstallationOrganization returns a boolean if a field has been set.

### GetInstallationRepository

`func (o *DynamicSecretCreateGithub) GetInstallationRepository() string`

GetInstallationRepository returns the InstallationRepository field if non-nil, zero value otherwise.

### GetInstallationRepositoryOk

`func (o *DynamicSecretCreateGithub) GetInstallationRepositoryOk() (*string, bool)`

GetInstallationRepositoryOk returns a tuple with the InstallationRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationRepository

`func (o *DynamicSecretCreateGithub) SetInstallationRepository(v string)`

SetInstallationRepository sets InstallationRepository field to given value.

### HasInstallationRepository

`func (o *DynamicSecretCreateGithub) HasInstallationRepository() bool`

HasInstallationRepository returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateGithub) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateGithub) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateGithub) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateGithub) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateGithub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateGithub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateGithub) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *DynamicSecretCreateGithub) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateGithub) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateGithub) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateGithub) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateGithub) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateGithub) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateGithub) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateGithub) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateGithub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateGithub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateGithub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateGithub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenPermissions

`func (o *DynamicSecretCreateGithub) GetTokenPermissions() []string`

GetTokenPermissions returns the TokenPermissions field if non-nil, zero value otherwise.

### GetTokenPermissionsOk

`func (o *DynamicSecretCreateGithub) GetTokenPermissionsOk() (*[]string, bool)`

GetTokenPermissionsOk returns a tuple with the TokenPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPermissions

`func (o *DynamicSecretCreateGithub) SetTokenPermissions(v []string)`

SetTokenPermissions sets TokenPermissions field to given value.

### HasTokenPermissions

`func (o *DynamicSecretCreateGithub) HasTokenPermissions() bool`

HasTokenPermissions returns a boolean if a field has been set.

### GetTokenRepositories

`func (o *DynamicSecretCreateGithub) GetTokenRepositories() []string`

GetTokenRepositories returns the TokenRepositories field if non-nil, zero value otherwise.

### GetTokenRepositoriesOk

`func (o *DynamicSecretCreateGithub) GetTokenRepositoriesOk() (*[]string, bool)`

GetTokenRepositoriesOk returns a tuple with the TokenRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRepositories

`func (o *DynamicSecretCreateGithub) SetTokenRepositories(v []string)`

SetTokenRepositories sets TokenRepositories field to given value.

### HasTokenRepositories

`func (o *DynamicSecretCreateGithub) HasTokenRepositories() bool`

HasTokenRepositories returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateGithub) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateGithub) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateGithub) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateGithub) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


