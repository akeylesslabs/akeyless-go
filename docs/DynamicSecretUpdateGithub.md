# DynamicSecretUpdateGithub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**GithubAppId** | Pointer to **int64** | Github app id | [optional] 
**GithubAppPrivateKey** | Pointer to **string** | App private key | [optional] 
**GithubBaseUrl** | Pointer to **string** | Base URL | [optional] [default to "https://api.github.com/"]
**InstallationId** | Pointer to **int64** | Github app installation id | [optional] 
**InstallationRepository** | Pointer to **string** | Repository that the app installation has access to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenPermissions** | Pointer to **[]string** | Optional - installation token&#39;s allowed permissions | [optional] 
**TokenRepositories** | Pointer to **[]string** | Optional - installation token&#39;s allowed repositories | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDynamicSecretUpdateGithub

`func NewDynamicSecretUpdateGithub(name string, ) *DynamicSecretUpdateGithub`

NewDynamicSecretUpdateGithub instantiates a new DynamicSecretUpdateGithub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateGithubWithDefaults

`func NewDynamicSecretUpdateGithubWithDefaults() *DynamicSecretUpdateGithub`

NewDynamicSecretUpdateGithubWithDefaults instantiates a new DynamicSecretUpdateGithub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretUpdateGithub) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateGithub) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateGithub) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateGithub) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetGithubAppId

`func (o *DynamicSecretUpdateGithub) GetGithubAppId() int64`

GetGithubAppId returns the GithubAppId field if non-nil, zero value otherwise.

### GetGithubAppIdOk

`func (o *DynamicSecretUpdateGithub) GetGithubAppIdOk() (*int64, bool)`

GetGithubAppIdOk returns a tuple with the GithubAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppId

`func (o *DynamicSecretUpdateGithub) SetGithubAppId(v int64)`

SetGithubAppId sets GithubAppId field to given value.

### HasGithubAppId

`func (o *DynamicSecretUpdateGithub) HasGithubAppId() bool`

HasGithubAppId returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *DynamicSecretUpdateGithub) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *DynamicSecretUpdateGithub) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *DynamicSecretUpdateGithub) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *DynamicSecretUpdateGithub) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetGithubBaseUrl

`func (o *DynamicSecretUpdateGithub) GetGithubBaseUrl() string`

GetGithubBaseUrl returns the GithubBaseUrl field if non-nil, zero value otherwise.

### GetGithubBaseUrlOk

`func (o *DynamicSecretUpdateGithub) GetGithubBaseUrlOk() (*string, bool)`

GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubBaseUrl

`func (o *DynamicSecretUpdateGithub) SetGithubBaseUrl(v string)`

SetGithubBaseUrl sets GithubBaseUrl field to given value.

### HasGithubBaseUrl

`func (o *DynamicSecretUpdateGithub) HasGithubBaseUrl() bool`

HasGithubBaseUrl returns a boolean if a field has been set.

### GetInstallationId

`func (o *DynamicSecretUpdateGithub) GetInstallationId() int64`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *DynamicSecretUpdateGithub) GetInstallationIdOk() (*int64, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *DynamicSecretUpdateGithub) SetInstallationId(v int64)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *DynamicSecretUpdateGithub) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetInstallationRepository

`func (o *DynamicSecretUpdateGithub) GetInstallationRepository() string`

GetInstallationRepository returns the InstallationRepository field if non-nil, zero value otherwise.

### GetInstallationRepositoryOk

`func (o *DynamicSecretUpdateGithub) GetInstallationRepositoryOk() (*string, bool)`

GetInstallationRepositoryOk returns a tuple with the InstallationRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationRepository

`func (o *DynamicSecretUpdateGithub) SetInstallationRepository(v string)`

SetInstallationRepository sets InstallationRepository field to given value.

### HasInstallationRepository

`func (o *DynamicSecretUpdateGithub) HasInstallationRepository() bool`

HasInstallationRepository returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateGithub) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateGithub) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateGithub) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateGithub) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateGithub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateGithub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateGithub) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateGithub) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateGithub) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateGithub) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateGithub) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateGithub) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateGithub) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateGithub) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateGithub) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateGithub) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateGithub) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateGithub) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateGithub) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateGithub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateGithub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateGithub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateGithub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenPermissions

`func (o *DynamicSecretUpdateGithub) GetTokenPermissions() []string`

GetTokenPermissions returns the TokenPermissions field if non-nil, zero value otherwise.

### GetTokenPermissionsOk

`func (o *DynamicSecretUpdateGithub) GetTokenPermissionsOk() (*[]string, bool)`

GetTokenPermissionsOk returns a tuple with the TokenPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPermissions

`func (o *DynamicSecretUpdateGithub) SetTokenPermissions(v []string)`

SetTokenPermissions sets TokenPermissions field to given value.

### HasTokenPermissions

`func (o *DynamicSecretUpdateGithub) HasTokenPermissions() bool`

HasTokenPermissions returns a boolean if a field has been set.

### GetTokenRepositories

`func (o *DynamicSecretUpdateGithub) GetTokenRepositories() []string`

GetTokenRepositories returns the TokenRepositories field if non-nil, zero value otherwise.

### GetTokenRepositoriesOk

`func (o *DynamicSecretUpdateGithub) GetTokenRepositoriesOk() (*[]string, bool)`

GetTokenRepositoriesOk returns a tuple with the TokenRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRepositories

`func (o *DynamicSecretUpdateGithub) SetTokenRepositories(v []string)`

SetTokenRepositories sets TokenRepositories field to given value.

### HasTokenRepositories

`func (o *DynamicSecretUpdateGithub) HasTokenRepositories() bool`

HasTokenRepositories returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateGithub) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateGithub) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateGithub) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateGithub) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


