# AssocTargetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**KeyOperations** | Pointer to **[]string** | A list of allowed operations for the key (required for azure targets) | [optional] 
**KeyringName** | Pointer to **string** | Keyring name of the GCP KMS (required for gcp targets) | [optional] 
**KmsAlgorithm** | Pointer to **string** | Algorithm of the key in GCP KMS (required for gcp targets) | [optional] 
**LocationId** | Pointer to **string** | Location id of the GCP KMS (required for gcp targets) | [optional] 
**Name** | **string** | The item to associate | 
**ProjectId** | Pointer to **string** | Project id of the GCP KMS (required for gcp targets) | [optional] 
**Purpose** | Pointer to **string** | Purpose of the key in GCP KMS (required for gcp targets) | [optional] 
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


