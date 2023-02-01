# GatewayUpdateProducerGithub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item | [optional] 
**GithubAppId** | Pointer to **int64** | Github app id | [optional] 
**GithubAppPrivateKey** | Pointer to **string** | App private key | [optional] 
**GithubBaseUrl** | Pointer to **string** | Base URL | [optional] 
**InstallationId** | Pointer to **int64** | Github app installation id | [optional] 
**InstallationRepository** | Pointer to **string** | Repository that the app installation has access to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenPermissions** | Pointer to **[]string** | Optional - installation token&#39;s allowed permissions | [optional] 
**TokenRepositories** | Pointer to **[]string** | Optional - installation token&#39;s allowed repositories | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateProducerGithub

`func NewGatewayUpdateProducerGithub(name string, ) *GatewayUpdateProducerGithub`

NewGatewayUpdateProducerGithub instantiates a new GatewayUpdateProducerGithub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerGithubWithDefaults

`func NewGatewayUpdateProducerGithubWithDefaults() *GatewayUpdateProducerGithub`

NewGatewayUpdateProducerGithubWithDefaults instantiates a new GatewayUpdateProducerGithub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *GatewayUpdateProducerGithub) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerGithub) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerGithub) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerGithub) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetGithubAppId

`func (o *GatewayUpdateProducerGithub) GetGithubAppId() int64`

GetGithubAppId returns the GithubAppId field if non-nil, zero value otherwise.

### GetGithubAppIdOk

`func (o *GatewayUpdateProducerGithub) GetGithubAppIdOk() (*int64, bool)`

GetGithubAppIdOk returns a tuple with the GithubAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppId

`func (o *GatewayUpdateProducerGithub) SetGithubAppId(v int64)`

SetGithubAppId sets GithubAppId field to given value.

### HasGithubAppId

`func (o *GatewayUpdateProducerGithub) HasGithubAppId() bool`

HasGithubAppId returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *GatewayUpdateProducerGithub) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *GatewayUpdateProducerGithub) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *GatewayUpdateProducerGithub) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *GatewayUpdateProducerGithub) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetGithubBaseUrl

`func (o *GatewayUpdateProducerGithub) GetGithubBaseUrl() string`

GetGithubBaseUrl returns the GithubBaseUrl field if non-nil, zero value otherwise.

### GetGithubBaseUrlOk

`func (o *GatewayUpdateProducerGithub) GetGithubBaseUrlOk() (*string, bool)`

GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubBaseUrl

`func (o *GatewayUpdateProducerGithub) SetGithubBaseUrl(v string)`

SetGithubBaseUrl sets GithubBaseUrl field to given value.

### HasGithubBaseUrl

`func (o *GatewayUpdateProducerGithub) HasGithubBaseUrl() bool`

HasGithubBaseUrl returns a boolean if a field has been set.

### GetInstallationId

`func (o *GatewayUpdateProducerGithub) GetInstallationId() int64`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *GatewayUpdateProducerGithub) GetInstallationIdOk() (*int64, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *GatewayUpdateProducerGithub) SetInstallationId(v int64)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *GatewayUpdateProducerGithub) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetInstallationRepository

`func (o *GatewayUpdateProducerGithub) GetInstallationRepository() string`

GetInstallationRepository returns the InstallationRepository field if non-nil, zero value otherwise.

### GetInstallationRepositoryOk

`func (o *GatewayUpdateProducerGithub) GetInstallationRepositoryOk() (*string, bool)`

GetInstallationRepositoryOk returns a tuple with the InstallationRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationRepository

`func (o *GatewayUpdateProducerGithub) SetInstallationRepository(v string)`

SetInstallationRepository sets InstallationRepository field to given value.

### HasInstallationRepository

`func (o *GatewayUpdateProducerGithub) HasInstallationRepository() bool`

HasInstallationRepository returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateProducerGithub) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateProducerGithub) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateProducerGithub) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateProducerGithub) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerGithub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerGithub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerGithub) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerGithub) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerGithub) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerGithub) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerGithub) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerGithub) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerGithub) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerGithub) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerGithub) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerGithub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerGithub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerGithub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerGithub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenPermissions

`func (o *GatewayUpdateProducerGithub) GetTokenPermissions() []string`

GetTokenPermissions returns the TokenPermissions field if non-nil, zero value otherwise.

### GetTokenPermissionsOk

`func (o *GatewayUpdateProducerGithub) GetTokenPermissionsOk() (*[]string, bool)`

GetTokenPermissionsOk returns a tuple with the TokenPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPermissions

`func (o *GatewayUpdateProducerGithub) SetTokenPermissions(v []string)`

SetTokenPermissions sets TokenPermissions field to given value.

### HasTokenPermissions

`func (o *GatewayUpdateProducerGithub) HasTokenPermissions() bool`

HasTokenPermissions returns a boolean if a field has been set.

### GetTokenRepositories

`func (o *GatewayUpdateProducerGithub) GetTokenRepositories() []string`

GetTokenRepositories returns the TokenRepositories field if non-nil, zero value otherwise.

### GetTokenRepositoriesOk

`func (o *GatewayUpdateProducerGithub) GetTokenRepositoriesOk() (*[]string, bool)`

GetTokenRepositoriesOk returns a tuple with the TokenRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRepositories

`func (o *GatewayUpdateProducerGithub) SetTokenRepositories(v []string)`

SetTokenRepositories sets TokenRepositories field to given value.

### HasTokenRepositories

`func (o *GatewayUpdateProducerGithub) HasTokenRepositories() bool`

HasTokenRepositories returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerGithub) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerGithub) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerGithub) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerGithub) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


