# DsCreateGithub

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
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenPermissions** | Pointer to **[]string** | Optional - installation token&#39;s allowed permissions | [optional] 
**TokenRepositories** | Pointer to **[]string** | Optional - installation token&#39;s allowed repositories | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDsCreateGithub

`func NewDsCreateGithub(name string, ) *DsCreateGithub`

NewDsCreateGithub instantiates a new DsCreateGithub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsCreateGithubWithDefaults

`func NewDsCreateGithubWithDefaults() *DsCreateGithub`

NewDsCreateGithubWithDefaults instantiates a new DsCreateGithub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DsCreateGithub) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DsCreateGithub) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DsCreateGithub) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DsCreateGithub) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetGithubAppId

`func (o *DsCreateGithub) GetGithubAppId() int64`

GetGithubAppId returns the GithubAppId field if non-nil, zero value otherwise.

### GetGithubAppIdOk

`func (o *DsCreateGithub) GetGithubAppIdOk() (*int64, bool)`

GetGithubAppIdOk returns a tuple with the GithubAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppId

`func (o *DsCreateGithub) SetGithubAppId(v int64)`

SetGithubAppId sets GithubAppId field to given value.

### HasGithubAppId

`func (o *DsCreateGithub) HasGithubAppId() bool`

HasGithubAppId returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *DsCreateGithub) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *DsCreateGithub) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *DsCreateGithub) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *DsCreateGithub) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetGithubBaseUrl

`func (o *DsCreateGithub) GetGithubBaseUrl() string`

GetGithubBaseUrl returns the GithubBaseUrl field if non-nil, zero value otherwise.

### GetGithubBaseUrlOk

`func (o *DsCreateGithub) GetGithubBaseUrlOk() (*string, bool)`

GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubBaseUrl

`func (o *DsCreateGithub) SetGithubBaseUrl(v string)`

SetGithubBaseUrl sets GithubBaseUrl field to given value.

### HasGithubBaseUrl

`func (o *DsCreateGithub) HasGithubBaseUrl() bool`

HasGithubBaseUrl returns a boolean if a field has been set.

### GetInstallationId

`func (o *DsCreateGithub) GetInstallationId() int64`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *DsCreateGithub) GetInstallationIdOk() (*int64, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *DsCreateGithub) SetInstallationId(v int64)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *DsCreateGithub) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetInstallationRepository

`func (o *DsCreateGithub) GetInstallationRepository() string`

GetInstallationRepository returns the InstallationRepository field if non-nil, zero value otherwise.

### GetInstallationRepositoryOk

`func (o *DsCreateGithub) GetInstallationRepositoryOk() (*string, bool)`

GetInstallationRepositoryOk returns a tuple with the InstallationRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationRepository

`func (o *DsCreateGithub) SetInstallationRepository(v string)`

SetInstallationRepository sets InstallationRepository field to given value.

### HasInstallationRepository

`func (o *DsCreateGithub) HasInstallationRepository() bool`

HasInstallationRepository returns a boolean if a field has been set.

### GetJson

`func (o *DsCreateGithub) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DsCreateGithub) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DsCreateGithub) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DsCreateGithub) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DsCreateGithub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DsCreateGithub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DsCreateGithub) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *DsCreateGithub) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DsCreateGithub) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DsCreateGithub) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DsCreateGithub) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DsCreateGithub) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DsCreateGithub) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DsCreateGithub) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DsCreateGithub) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DsCreateGithub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DsCreateGithub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DsCreateGithub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DsCreateGithub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenPermissions

`func (o *DsCreateGithub) GetTokenPermissions() []string`

GetTokenPermissions returns the TokenPermissions field if non-nil, zero value otherwise.

### GetTokenPermissionsOk

`func (o *DsCreateGithub) GetTokenPermissionsOk() (*[]string, bool)`

GetTokenPermissionsOk returns a tuple with the TokenPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPermissions

`func (o *DsCreateGithub) SetTokenPermissions(v []string)`

SetTokenPermissions sets TokenPermissions field to given value.

### HasTokenPermissions

`func (o *DsCreateGithub) HasTokenPermissions() bool`

HasTokenPermissions returns a boolean if a field has been set.

### GetTokenRepositories

`func (o *DsCreateGithub) GetTokenRepositories() []string`

GetTokenRepositories returns the TokenRepositories field if non-nil, zero value otherwise.

### GetTokenRepositoriesOk

`func (o *DsCreateGithub) GetTokenRepositoriesOk() (*[]string, bool)`

GetTokenRepositoriesOk returns a tuple with the TokenRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRepositories

`func (o *DsCreateGithub) SetTokenRepositories(v []string)`

SetTokenRepositories sets TokenRepositories field to given value.

### HasTokenRepositories

`func (o *DsCreateGithub) HasTokenRepositories() bool`

HasTokenRepositories returns a boolean if a field has been set.

### GetUidToken

`func (o *DsCreateGithub) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DsCreateGithub) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DsCreateGithub) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DsCreateGithub) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


