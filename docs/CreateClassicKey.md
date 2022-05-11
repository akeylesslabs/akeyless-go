# CreateClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | Classic Key type; options: [AES128GCM, AES256GCM, AES128SIV, AES256SIV, RSA1024, RSA2048, RSA3072, RSA4096, EC256, EC384] | 
**CertFileData** | Pointer to **string** | Certificate in a PEM format. | [optional] 
**KeyData** | Pointer to **string** | Base64-encoded classic key value | [optional] 
**KeyOperations** | Pointer to **[]string** | A list of allowed operations for the key (required for azure targets) | [optional] 
**KeyringName** | Pointer to **string** | Keyring name of the GCP KMS (required for gcp targets) | [optional] 
**KmsAlgorithm** | Pointer to **string** | Algorithm of the key in GCP KMS (required for gcp targets) | [optional] 
**LocationId** | Pointer to **string** | Location id of the GCP KMS (required for gcp targets) | [optional] 
**Metadata** | Pointer to **string** | Metadata about the classic key | [optional] 
**Name** | **string** | ClassicKey name | 
**ProjectId** | Pointer to **string** | Project id of the GCP KMS (required for gcp targets) | [optional] 
**ProtectionKeyName** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Purpose** | Pointer to **string** | Purpose of the key in GCP KMS (required for gcp targets) | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this classic key | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**TenantSecretType** | Pointer to **string** | The tenant secret type [Data/SearchIndex/Analytics] (required for salesforce targets) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**VaultName** | Pointer to **string** | Name of the vault used (required for azure targets) | [optional] 

## Methods

### NewCreateClassicKey

`func NewCreateClassicKey(alg string, name string, ) *CreateClassicKey`

NewCreateClassicKey instantiates a new CreateClassicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClassicKeyWithDefaults

`func NewCreateClassicKeyWithDefaults() *CreateClassicKey`

NewCreateClassicKeyWithDefaults instantiates a new CreateClassicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *CreateClassicKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *CreateClassicKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *CreateClassicKey) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetCertFileData

`func (o *CreateClassicKey) GetCertFileData() string`

GetCertFileData returns the CertFileData field if non-nil, zero value otherwise.

### GetCertFileDataOk

`func (o *CreateClassicKey) GetCertFileDataOk() (*string, bool)`

GetCertFileDataOk returns a tuple with the CertFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFileData

`func (o *CreateClassicKey) SetCertFileData(v string)`

SetCertFileData sets CertFileData field to given value.

### HasCertFileData

`func (o *CreateClassicKey) HasCertFileData() bool`

HasCertFileData returns a boolean if a field has been set.

### GetKeyData

`func (o *CreateClassicKey) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *CreateClassicKey) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *CreateClassicKey) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *CreateClassicKey) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetKeyOperations

`func (o *CreateClassicKey) GetKeyOperations() []string`

GetKeyOperations returns the KeyOperations field if non-nil, zero value otherwise.

### GetKeyOperationsOk

`func (o *CreateClassicKey) GetKeyOperationsOk() (*[]string, bool)`

GetKeyOperationsOk returns a tuple with the KeyOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyOperations

`func (o *CreateClassicKey) SetKeyOperations(v []string)`

SetKeyOperations sets KeyOperations field to given value.

### HasKeyOperations

`func (o *CreateClassicKey) HasKeyOperations() bool`

HasKeyOperations returns a boolean if a field has been set.

### GetKeyringName

`func (o *CreateClassicKey) GetKeyringName() string`

GetKeyringName returns the KeyringName field if non-nil, zero value otherwise.

### GetKeyringNameOk

`func (o *CreateClassicKey) GetKeyringNameOk() (*string, bool)`

GetKeyringNameOk returns a tuple with the KeyringName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyringName

`func (o *CreateClassicKey) SetKeyringName(v string)`

SetKeyringName sets KeyringName field to given value.

### HasKeyringName

`func (o *CreateClassicKey) HasKeyringName() bool`

HasKeyringName returns a boolean if a field has been set.

### GetKmsAlgorithm

`func (o *CreateClassicKey) GetKmsAlgorithm() string`

GetKmsAlgorithm returns the KmsAlgorithm field if non-nil, zero value otherwise.

### GetKmsAlgorithmOk

`func (o *CreateClassicKey) GetKmsAlgorithmOk() (*string, bool)`

GetKmsAlgorithmOk returns a tuple with the KmsAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsAlgorithm

`func (o *CreateClassicKey) SetKmsAlgorithm(v string)`

SetKmsAlgorithm sets KmsAlgorithm field to given value.

### HasKmsAlgorithm

`func (o *CreateClassicKey) HasKmsAlgorithm() bool`

HasKmsAlgorithm returns a boolean if a field has been set.

### GetLocationId

`func (o *CreateClassicKey) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CreateClassicKey) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CreateClassicKey) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *CreateClassicKey) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateClassicKey) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateClassicKey) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateClassicKey) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateClassicKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateClassicKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClassicKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClassicKey) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *CreateClassicKey) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateClassicKey) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateClassicKey) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateClassicKey) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProtectionKeyName

`func (o *CreateClassicKey) GetProtectionKeyName() string`

GetProtectionKeyName returns the ProtectionKeyName field if non-nil, zero value otherwise.

### GetProtectionKeyNameOk

`func (o *CreateClassicKey) GetProtectionKeyNameOk() (*string, bool)`

GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyName

`func (o *CreateClassicKey) SetProtectionKeyName(v string)`

SetProtectionKeyName sets ProtectionKeyName field to given value.

### HasProtectionKeyName

`func (o *CreateClassicKey) HasProtectionKeyName() bool`

HasProtectionKeyName returns a boolean if a field has been set.

### GetPurpose

`func (o *CreateClassicKey) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreateClassicKey) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreateClassicKey) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *CreateClassicKey) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetTags

`func (o *CreateClassicKey) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateClassicKey) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateClassicKey) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateClassicKey) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *CreateClassicKey) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *CreateClassicKey) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *CreateClassicKey) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *CreateClassicKey) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTenantSecretType

`func (o *CreateClassicKey) GetTenantSecretType() string`

GetTenantSecretType returns the TenantSecretType field if non-nil, zero value otherwise.

### GetTenantSecretTypeOk

`func (o *CreateClassicKey) GetTenantSecretTypeOk() (*string, bool)`

GetTenantSecretTypeOk returns a tuple with the TenantSecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantSecretType

`func (o *CreateClassicKey) SetTenantSecretType(v string)`

SetTenantSecretType sets TenantSecretType field to given value.

### HasTenantSecretType

`func (o *CreateClassicKey) HasTenantSecretType() bool`

HasTenantSecretType returns a boolean if a field has been set.

### GetToken

`func (o *CreateClassicKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateClassicKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateClassicKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateClassicKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateClassicKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateClassicKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateClassicKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateClassicKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVaultName

`func (o *CreateClassicKey) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *CreateClassicKey) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *CreateClassicKey) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *CreateClassicKey) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


