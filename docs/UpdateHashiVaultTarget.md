# UpdateHashiVaultTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**HashiUrl** | Pointer to **string** | HashiCorp Vault API URL, e.g. https://vault-mgr01:8200 | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Namespace** | Pointer to **[]string** | Comma-separated list of vault namespaces | [optional] 
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 
**VaultToken** | Pointer to **string** | Vault access token with sufficient permissions | [optional] 

## Methods

### NewUpdateHashiVaultTarget

`func NewUpdateHashiVaultTarget(name string, ) *UpdateHashiVaultTarget`

NewUpdateHashiVaultTarget instantiates a new UpdateHashiVaultTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHashiVaultTargetWithDefaults

`func NewUpdateHashiVaultTargetWithDefaults() *UpdateHashiVaultTarget`

NewUpdateHashiVaultTargetWithDefaults instantiates a new UpdateHashiVaultTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateHashiVaultTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateHashiVaultTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateHashiVaultTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateHashiVaultTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHashiUrl

`func (o *UpdateHashiVaultTarget) GetHashiUrl() string`

GetHashiUrl returns the HashiUrl field if non-nil, zero value otherwise.

### GetHashiUrlOk

`func (o *UpdateHashiVaultTarget) GetHashiUrlOk() (*string, bool)`

GetHashiUrlOk returns a tuple with the HashiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiUrl

`func (o *UpdateHashiVaultTarget) SetHashiUrl(v string)`

SetHashiUrl sets HashiUrl field to given value.

### HasHashiUrl

`func (o *UpdateHashiVaultTarget) HasHashiUrl() bool`

HasHashiUrl returns a boolean if a field has been set.

### GetJson

`func (o *UpdateHashiVaultTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateHashiVaultTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateHashiVaultTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateHashiVaultTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateHashiVaultTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateHashiVaultTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateHashiVaultTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateHashiVaultTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateHashiVaultTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateHashiVaultTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateHashiVaultTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateHashiVaultTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *UpdateHashiVaultTarget) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *UpdateHashiVaultTarget) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *UpdateHashiVaultTarget) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *UpdateHashiVaultTarget) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *UpdateHashiVaultTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateHashiVaultTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateHashiVaultTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *UpdateHashiVaultTarget) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpdateHashiVaultTarget) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpdateHashiVaultTarget) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UpdateHashiVaultTarget) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateHashiVaultTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateHashiVaultTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateHashiVaultTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateHashiVaultTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateHashiVaultTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateHashiVaultTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateHashiVaultTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateHashiVaultTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateHashiVaultTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateHashiVaultTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateHashiVaultTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateHashiVaultTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateHashiVaultTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateHashiVaultTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateHashiVaultTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateHashiVaultTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.

### GetVaultToken

`func (o *UpdateHashiVaultTarget) GetVaultToken() string`

GetVaultToken returns the VaultToken field if non-nil, zero value otherwise.

### GetVaultTokenOk

`func (o *UpdateHashiVaultTarget) GetVaultTokenOk() (*string, bool)`

GetVaultTokenOk returns a tuple with the VaultToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultToken

`func (o *UpdateHashiVaultTarget) SetVaultToken(v string)`

SetVaultToken sets VaultToken field to given value.

### HasVaultToken

`func (o *UpdateHashiVaultTarget) HasVaultToken() bool`

HasVaultToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


