# AssocTargetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificatePath** | Pointer to **string** | A path on the target to store the certificate pem file (relevant only for certificate provisioning) | [optional] 
**ChainPath** | Pointer to **string** | A path on the target to store the full chain pem file (relevant only for certificate provisioning) | [optional] 
**DisablePreviousKeyVersion** | Pointer to **bool** | Automatically disable previous key version (required for azure targets) | [optional] [default to false]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyOperations** | Pointer to **[]string** | A list of allowed operations for the key (required for azure targets) | [optional] 
**KeyringName** | Pointer to **string** | Keyring name of the GCP KMS (required for gcp targets) | [optional] 
**KmsAlgorithm** | Pointer to **string** | Algorithm of the key in GCP KMS (required for gcp targets) | [optional] 
**LocationId** | Pointer to **string** | Location id of the GCP KMS (required for gcp targets) | [optional] 
**MultiRegion** | Pointer to **string** | Set to &#39;true&#39; to create a multi-region managed key. (Relevant only for Classic Key AWS targets) | [optional] [default to "false"]
**Name** | **string** | The item to associate | 
**PostProvisionCommand** | Pointer to **string** | A custom command to run on the remote target after successful provisioning (relevant only for certificate provisioning) | [optional] 
**PrivateKeyPath** | Pointer to **string** | A path on the target to store the private key (relevant only for certificate provisioning) | [optional] 
**ProjectId** | Pointer to **string** | Project id of the GCP KMS (required for gcp targets) | [optional] 
**Purpose** | Pointer to **string** | Purpose of the key in GCP KMS (required for gcp targets) | [optional] 
**Regions** | Pointer to **[]string** | The list of regions to create a copy of the key in (relevant for aws targets) | [optional] 
**SraAssociation** | Pointer to **bool** | Is the target to associate is for sra, relevant only for linked target association for ldap rotated secret | [optional] [default to false]
**TargetName** | **string** | The target to associate | 
**TenantSecretType** | Pointer to **string** | The tenant secret type [Data/SearchIndex/Analytics] (required for salesforce targets) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**VaultName** | Pointer to **string** | Name of the vault used (required for azure targets) | [optional] 

## Methods

### NewAssocTargetItem

`func NewAssocTargetItem(name string, targetName string, ) *AssocTargetItem`

NewAssocTargetItem instantiates a new AssocTargetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssocTargetItemWithDefaults

`func NewAssocTargetItemWithDefaults() *AssocTargetItem`

NewAssocTargetItemWithDefaults instantiates a new AssocTargetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificatePath

`func (o *AssocTargetItem) GetCertificatePath() string`

GetCertificatePath returns the CertificatePath field if non-nil, zero value otherwise.

### GetCertificatePathOk

`func (o *AssocTargetItem) GetCertificatePathOk() (*string, bool)`

GetCertificatePathOk returns a tuple with the CertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePath

`func (o *AssocTargetItem) SetCertificatePath(v string)`

SetCertificatePath sets CertificatePath field to given value.

### HasCertificatePath

`func (o *AssocTargetItem) HasCertificatePath() bool`

HasCertificatePath returns a boolean if a field has been set.

### GetChainPath

`func (o *AssocTargetItem) GetChainPath() string`

GetChainPath returns the ChainPath field if non-nil, zero value otherwise.

### GetChainPathOk

`func (o *AssocTargetItem) GetChainPathOk() (*string, bool)`

GetChainPathOk returns a tuple with the ChainPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainPath

`func (o *AssocTargetItem) SetChainPath(v string)`

SetChainPath sets ChainPath field to given value.

### HasChainPath

`func (o *AssocTargetItem) HasChainPath() bool`

HasChainPath returns a boolean if a field has been set.

### GetDisablePreviousKeyVersion

`func (o *AssocTargetItem) GetDisablePreviousKeyVersion() bool`

GetDisablePreviousKeyVersion returns the DisablePreviousKeyVersion field if non-nil, zero value otherwise.

### GetDisablePreviousKeyVersionOk

`func (o *AssocTargetItem) GetDisablePreviousKeyVersionOk() (*bool, bool)`

GetDisablePreviousKeyVersionOk returns a tuple with the DisablePreviousKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePreviousKeyVersion

`func (o *AssocTargetItem) SetDisablePreviousKeyVersion(v bool)`

SetDisablePreviousKeyVersion sets DisablePreviousKeyVersion field to given value.

### HasDisablePreviousKeyVersion

`func (o *AssocTargetItem) HasDisablePreviousKeyVersion() bool`

HasDisablePreviousKeyVersion returns a boolean if a field has been set.

### GetJson

`func (o *AssocTargetItem) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AssocTargetItem) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AssocTargetItem) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AssocTargetItem) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyOperations

`func (o *AssocTargetItem) GetKeyOperations() []string`

GetKeyOperations returns the KeyOperations field if non-nil, zero value otherwise.

### GetKeyOperationsOk

`func (o *AssocTargetItem) GetKeyOperationsOk() (*[]string, bool)`

GetKeyOperationsOk returns a tuple with the KeyOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyOperations

`func (o *AssocTargetItem) SetKeyOperations(v []string)`

SetKeyOperations sets KeyOperations field to given value.

### HasKeyOperations

`func (o *AssocTargetItem) HasKeyOperations() bool`

HasKeyOperations returns a boolean if a field has been set.

### GetKeyringName

`func (o *AssocTargetItem) GetKeyringName() string`

GetKeyringName returns the KeyringName field if non-nil, zero value otherwise.

### GetKeyringNameOk

`func (o *AssocTargetItem) GetKeyringNameOk() (*string, bool)`

GetKeyringNameOk returns a tuple with the KeyringName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyringName

`func (o *AssocTargetItem) SetKeyringName(v string)`

SetKeyringName sets KeyringName field to given value.

### HasKeyringName

`func (o *AssocTargetItem) HasKeyringName() bool`

HasKeyringName returns a boolean if a field has been set.

### GetKmsAlgorithm

`func (o *AssocTargetItem) GetKmsAlgorithm() string`

GetKmsAlgorithm returns the KmsAlgorithm field if non-nil, zero value otherwise.

### GetKmsAlgorithmOk

`func (o *AssocTargetItem) GetKmsAlgorithmOk() (*string, bool)`

GetKmsAlgorithmOk returns a tuple with the KmsAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsAlgorithm

`func (o *AssocTargetItem) SetKmsAlgorithm(v string)`

SetKmsAlgorithm sets KmsAlgorithm field to given value.

### HasKmsAlgorithm

`func (o *AssocTargetItem) HasKmsAlgorithm() bool`

HasKmsAlgorithm returns a boolean if a field has been set.

### GetLocationId

`func (o *AssocTargetItem) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *AssocTargetItem) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *AssocTargetItem) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *AssocTargetItem) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetMultiRegion

`func (o *AssocTargetItem) GetMultiRegion() string`

GetMultiRegion returns the MultiRegion field if non-nil, zero value otherwise.

### GetMultiRegionOk

`func (o *AssocTargetItem) GetMultiRegionOk() (*string, bool)`

GetMultiRegionOk returns a tuple with the MultiRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiRegion

`func (o *AssocTargetItem) SetMultiRegion(v string)`

SetMultiRegion sets MultiRegion field to given value.

### HasMultiRegion

`func (o *AssocTargetItem) HasMultiRegion() bool`

HasMultiRegion returns a boolean if a field has been set.

### GetName

`func (o *AssocTargetItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssocTargetItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssocTargetItem) SetName(v string)`

SetName sets Name field to given value.


### GetPostProvisionCommand

`func (o *AssocTargetItem) GetPostProvisionCommand() string`

GetPostProvisionCommand returns the PostProvisionCommand field if non-nil, zero value otherwise.

### GetPostProvisionCommandOk

`func (o *AssocTargetItem) GetPostProvisionCommandOk() (*string, bool)`

GetPostProvisionCommandOk returns a tuple with the PostProvisionCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostProvisionCommand

`func (o *AssocTargetItem) SetPostProvisionCommand(v string)`

SetPostProvisionCommand sets PostProvisionCommand field to given value.

### HasPostProvisionCommand

`func (o *AssocTargetItem) HasPostProvisionCommand() bool`

HasPostProvisionCommand returns a boolean if a field has been set.

### GetPrivateKeyPath

`func (o *AssocTargetItem) GetPrivateKeyPath() string`

GetPrivateKeyPath returns the PrivateKeyPath field if non-nil, zero value otherwise.

### GetPrivateKeyPathOk

`func (o *AssocTargetItem) GetPrivateKeyPathOk() (*string, bool)`

GetPrivateKeyPathOk returns a tuple with the PrivateKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPath

`func (o *AssocTargetItem) SetPrivateKeyPath(v string)`

SetPrivateKeyPath sets PrivateKeyPath field to given value.

### HasPrivateKeyPath

`func (o *AssocTargetItem) HasPrivateKeyPath() bool`

HasPrivateKeyPath returns a boolean if a field has been set.

### GetProjectId

`func (o *AssocTargetItem) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AssocTargetItem) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AssocTargetItem) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AssocTargetItem) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPurpose

`func (o *AssocTargetItem) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *AssocTargetItem) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *AssocTargetItem) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *AssocTargetItem) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetRegions

`func (o *AssocTargetItem) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AssocTargetItem) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AssocTargetItem) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AssocTargetItem) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSraAssociation

`func (o *AssocTargetItem) GetSraAssociation() bool`

GetSraAssociation returns the SraAssociation field if non-nil, zero value otherwise.

### GetSraAssociationOk

`func (o *AssocTargetItem) GetSraAssociationOk() (*bool, bool)`

GetSraAssociationOk returns a tuple with the SraAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSraAssociation

`func (o *AssocTargetItem) SetSraAssociation(v bool)`

SetSraAssociation sets SraAssociation field to given value.

### HasSraAssociation

`func (o *AssocTargetItem) HasSraAssociation() bool`

HasSraAssociation returns a boolean if a field has been set.

### GetTargetName

`func (o *AssocTargetItem) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *AssocTargetItem) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *AssocTargetItem) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetTenantSecretType

`func (o *AssocTargetItem) GetTenantSecretType() string`

GetTenantSecretType returns the TenantSecretType field if non-nil, zero value otherwise.

### GetTenantSecretTypeOk

`func (o *AssocTargetItem) GetTenantSecretTypeOk() (*string, bool)`

GetTenantSecretTypeOk returns a tuple with the TenantSecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantSecretType

`func (o *AssocTargetItem) SetTenantSecretType(v string)`

SetTenantSecretType sets TenantSecretType field to given value.

### HasTenantSecretType

`func (o *AssocTargetItem) HasTenantSecretType() bool`

HasTenantSecretType returns a boolean if a field has been set.

### GetToken

`func (o *AssocTargetItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AssocTargetItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AssocTargetItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AssocTargetItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AssocTargetItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AssocTargetItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AssocTargetItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AssocTargetItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVaultName

`func (o *AssocTargetItem) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *AssocTargetItem) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *AssocTargetItem) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *AssocTargetItem) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


