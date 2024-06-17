# CreateHashiVaultTarget

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

### NewCreateHashiVaultTarget

`func NewCreateHashiVaultTarget(name string, ) *CreateHashiVaultTarget`

NewCreateHashiVaultTarget instantiates a new CreateHashiVaultTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHashiVaultTargetWithDefaults

`func NewCreateHashiVaultTargetWithDefaults() *CreateHashiVaultTarget`

NewCreateHashiVaultTargetWithDefaults instantiates a new CreateHashiVaultTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateHashiVaultTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateHashiVaultTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateHashiVaultTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateHashiVaultTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHashiUrl

`func (o *CreateHashiVaultTarget) GetHashiUrl() string`

GetHashiUrl returns the HashiUrl field if non-nil, zero value otherwise.

### GetHashiUrlOk

`func (o *CreateHashiVaultTarget) GetHashiUrlOk() (*string, bool)`

GetHashiUrlOk returns a tuple with the HashiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiUrl

`func (o *CreateHashiVaultTarget) SetHashiUrl(v string)`

SetHashiUrl sets HashiUrl field to given value.

### HasHashiUrl

`func (o *CreateHashiVaultTarget) HasHashiUrl() bool`

HasHashiUrl returns a boolean if a field has been set.

### GetJson

`func (o *CreateHashiVaultTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateHashiVaultTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateHashiVaultTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateHashiVaultTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *CreateHashiVaultTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateHashiVaultTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateHashiVaultTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateHashiVaultTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *CreateHashiVaultTarget) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *CreateHashiVaultTarget) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *CreateHashiVaultTarget) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *CreateHashiVaultTarget) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *CreateHashiVaultTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHashiVaultTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHashiVaultTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *CreateHashiVaultTarget) GetNamespace() []string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateHashiVaultTarget) GetNamespaceOk() (*[]string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateHashiVaultTarget) SetNamespace(v []string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateHashiVaultTarget) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetToken

`func (o *CreateHashiVaultTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateHashiVaultTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateHashiVaultTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateHashiVaultTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateHashiVaultTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateHashiVaultTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateHashiVaultTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateHashiVaultTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVaultToken

`func (o *CreateHashiVaultTarget) GetVaultToken() string`

GetVaultToken returns the VaultToken field if non-nil, zero value otherwise.

### GetVaultTokenOk

`func (o *CreateHashiVaultTarget) GetVaultTokenOk() (*string, bool)`

GetVaultTokenOk returns a tuple with the VaultToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultToken

`func (o *CreateHashiVaultTarget) SetVaultToken(v string)`

SetVaultToken sets VaultToken field to given value.

### HasVaultToken

`func (o *CreateHashiVaultTarget) HasVaultToken() bool`

HasVaultToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


