# TargetCreateHashiVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**HashiUrl** | Pointer to **string** | HashiCorp Vault API URL, e.g. https://vault-mgr01:8200 | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Namespace** | Pointer to **[]string** | Comma-separated list of vault namespaces | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**VaultToken** | Pointer to **string** | Vault access token with sufficient permissions | [optional] 

## Methods

### NewTargetCreateHashiVault

`func NewTargetCreateHashiVault(name string, ) *TargetCreateHashiVault`

NewTargetCreateHashiVault instantiates a new TargetCreateHashiVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateHashiVaultWithDefaults

`func NewTargetCreateHashiVaultWithDefaults() *TargetCreateHashiVault`

NewTargetCreateHashiVaultWithDefaults instantiates a new TargetCreateHashiVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TargetCreateHashiVault) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateHashiVault) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateHashiVault) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateHashiVault) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHashiUrl

`func (o *TargetCreateHashiVault) GetHashiUrl() string`

GetHashiUrl returns the HashiUrl field if non-nil, zero value otherwise.

### GetHashiUrlOk

`func (o *TargetCreateHashiVault) GetHashiUrlOk() (*string, bool)`

GetHashiUrlOk returns a tuple with the HashiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiUrl

`func (o *TargetCreateHashiVault) SetHashiUrl(v string)`

SetHashiUrl sets HashiUrl field to given value.

### HasHashiUrl

`func (o *TargetCreateHashiVault) HasHashiUrl() bool`

HasHashiUrl returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateHashiVault) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateHashiVault) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateHashiVault) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateHashiVault) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateHashiVault) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateHashiVault) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateHashiVault) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateHashiVault) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateHashiVault) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateHashiVault) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateHashiVault) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateHashiVault) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateHashiVault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateHashiVault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateHashiVault) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *TargetCreateHashiVault) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TargetCreateHashiVault) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TargetCreateHashiVault) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *TargetCreateHashiVault) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateHashiVault) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateHashiVault) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateHashiVault) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateHashiVault) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateHashiVault) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateHashiVault) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateHashiVault) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateHashiVault) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVaultToken

`func (o *TargetCreateHashiVault) GetVaultToken() string`

GetVaultToken returns the VaultToken field if non-nil, zero value otherwise.

### GetVaultTokenOk

`func (o *TargetCreateHashiVault) GetVaultTokenOk() (*string, bool)`

GetVaultTokenOk returns a tuple with the VaultToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultToken

`func (o *TargetCreateHashiVault) SetVaultToken(v string)`

SetVaultToken sets VaultToken field to given value.

### HasVaultToken

`func (o *TargetCreateHashiVault) HasVaultToken() bool`

HasVaultToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


