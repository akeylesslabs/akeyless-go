# GatewayCreateMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKey** | Pointer to **string** | AWS Secret Access Key (relevant only for AWS migration) | [optional] 
**AwsKeyId** | Pointer to **string** | AWS Access Key ID with sufficient permissions to get all secrets, e.g. &#39;arn:aws:secretsmanager:[Region]:[AccountId]:secret:[/path/to/secrets/_*]&#39; (relevant only for AWS migration) | [optional] 
**AwsRegion** | Pointer to **string** | AWS region of the required Secrets Manager (relevant only for AWS migration) | [optional] 
**AzureClientId** | Pointer to **string** | Azure Key Vault Access client ID, should be Azure AD App with a service principal (relevant only for Azure Key Vault migration) | [optional] 
**AzureKvName** | Pointer to **string** | Azure Key Vault Name (relevant only for Azure Key Vault migration) | [optional] 
**AzureSecret** | Pointer to **string** | Azure Key Vault secret (relevant only for Azure Key Vault migration) | [optional] 
**AzureTenantId** | Pointer to **string** | Azure Key Vault Access tenant ID (relevant only for Azure Key Vault migration) | [optional] 
**GcpKey** | Pointer to **string** | Base64-encoded GCP Service Account private key text with sufficient permissions to Secrets Manager, Minimum required permission is Secret Manager Secret Accessor, e.g. &#39;roles/secretmanager.secretAccessor&#39; (relevant only for GCP migration) | [optional] 
**HashiJson** | Pointer to **string** | Import secret key as json value or independent secrets (relevant only for HasiCorp Vault migration) | [optional] 
**HashiNs** | Pointer to **[]string** | HashiCorp Vault Namespaces is a comma-separated list of namespaces which need to be imported into Akeyless Vault. For every provided namespace, all its child namespaces are imported as well, e.g. nmsp/subnmsp1/subnmsp2,nmsp/anothernmsp. By default, import all namespaces (relevant only for HasiCorp Vault migration) | [optional] 
**HashiToken** | Pointer to **string** | HashiCorp Vault access token with sufficient permissions to preform list &amp; read operations on secrets objects (relevant only for HasiCorp Vault migration) | [optional] 
**HashiUrl** | Pointer to **string** | HashiCorp Vault API URL, e.g. https://vault-mgr01:8200 (relevant only for HasiCorp Vault migration) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**K8sCaCertificate** | Pointer to **[]int32** | For Certificate Authentication method K8s Cluster CA certificate (relevant only for K8s migration with Certificate Authentication method) | [optional] 
**K8sClientCertificate** | Pointer to **[]int32** | K8s Client certificate with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Certificate Authentication method) | [optional] 
**K8sClientKey** | Pointer to **[]int32** | K8s Client key (relevant only for K8s migration with Certificate Authentication method) | [optional] 
**K8sNamespace** | Pointer to **string** | K8s Namespace, Use this field to import secrets from a particular namespace only. By default, the secrets are imported from all namespaces (relevant only for K8s migration) | [optional] 
**K8sPassword** | Pointer to **string** | K8s Client password (relevant only for K8s migration with Password Authentication method) | [optional] 
**K8sSkipSystem** | Pointer to **bool** | K8s Skip Control Plane Secrets, This option allows to avoid importing secrets from system namespaces (relevant only for K8s migration) | [optional] 
**K8sToken** | Pointer to **string** | For Token Authentication method K8s Bearer Token with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Token Authentication method) | [optional] 
**K8sUrl** | Pointer to **string** | K8s API Server URL, e.g. https://k8s-api.mycompany.com:6443 (relevant only for K8s migration) | [optional] 
**K8sUsername** | Pointer to **string** | For Password Authentication method K8s Client username with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Password Authentication method) | [optional] 
**Name** | **string** | Migration name | 
**OpEmail** | Pointer to **string** | 1Password user email to connect to the API | [optional] 
**OpPassword** | Pointer to **string** | 1Password user password to connect to the API | [optional] 
**OpSecretKey** | Pointer to **string** | 1Password user secret key to connect to the API | [optional] 
**OpUrl** | Pointer to **string** | 1Password api container url | [optional] 
**OpVaults** | Pointer to **[]string** | 1Password list of vault to get the items from | [optional] 
**ProtectionKey** | Pointer to **string** | The name of the key that protects the classic key value (if empty, the account default key will be used) | [optional] 
**TargetLocation** | Pointer to **string** | Target location in Akeyless for imported secrets | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **string** | Migration type (hashi/aws/gcp/k8s/azure_kv/1password) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayCreateMigration

`func NewGatewayCreateMigration(name string, ) *GatewayCreateMigration`

NewGatewayCreateMigration instantiates a new GatewayCreateMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateMigrationWithDefaults

`func NewGatewayCreateMigrationWithDefaults() *GatewayCreateMigration`

NewGatewayCreateMigrationWithDefaults instantiates a new GatewayCreateMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKey

`func (o *GatewayCreateMigration) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *GatewayCreateMigration) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *GatewayCreateMigration) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.

### HasAwsKey

`func (o *GatewayCreateMigration) HasAwsKey() bool`

HasAwsKey returns a boolean if a field has been set.

### GetAwsKeyId

`func (o *GatewayCreateMigration) GetAwsKeyId() string`

GetAwsKeyId returns the AwsKeyId field if non-nil, zero value otherwise.

### GetAwsKeyIdOk

`func (o *GatewayCreateMigration) GetAwsKeyIdOk() (*string, bool)`

GetAwsKeyIdOk returns a tuple with the AwsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKeyId

`func (o *GatewayCreateMigration) SetAwsKeyId(v string)`

SetAwsKeyId sets AwsKeyId field to given value.

### HasAwsKeyId

`func (o *GatewayCreateMigration) HasAwsKeyId() bool`

HasAwsKeyId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *GatewayCreateMigration) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *GatewayCreateMigration) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *GatewayCreateMigration) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *GatewayCreateMigration) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAzureClientId

`func (o *GatewayCreateMigration) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *GatewayCreateMigration) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *GatewayCreateMigration) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *GatewayCreateMigration) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureKvName

`func (o *GatewayCreateMigration) GetAzureKvName() string`

GetAzureKvName returns the AzureKvName field if non-nil, zero value otherwise.

### GetAzureKvNameOk

`func (o *GatewayCreateMigration) GetAzureKvNameOk() (*string, bool)`

GetAzureKvNameOk returns a tuple with the AzureKvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKvName

`func (o *GatewayCreateMigration) SetAzureKvName(v string)`

SetAzureKvName sets AzureKvName field to given value.

### HasAzureKvName

`func (o *GatewayCreateMigration) HasAzureKvName() bool`

HasAzureKvName returns a boolean if a field has been set.

### GetAzureSecret

`func (o *GatewayCreateMigration) GetAzureSecret() string`

GetAzureSecret returns the AzureSecret field if non-nil, zero value otherwise.

### GetAzureSecretOk

`func (o *GatewayCreateMigration) GetAzureSecretOk() (*string, bool)`

GetAzureSecretOk returns a tuple with the AzureSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSecret

`func (o *GatewayCreateMigration) SetAzureSecret(v string)`

SetAzureSecret sets AzureSecret field to given value.

### HasAzureSecret

`func (o *GatewayCreateMigration) HasAzureSecret() bool`

HasAzureSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *GatewayCreateMigration) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *GatewayCreateMigration) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *GatewayCreateMigration) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *GatewayCreateMigration) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetGcpKey

`func (o *GatewayCreateMigration) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *GatewayCreateMigration) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *GatewayCreateMigration) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *GatewayCreateMigration) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetHashiJson

`func (o *GatewayCreateMigration) GetHashiJson() string`

GetHashiJson returns the HashiJson field if non-nil, zero value otherwise.

### GetHashiJsonOk

`func (o *GatewayCreateMigration) GetHashiJsonOk() (*string, bool)`

GetHashiJsonOk returns a tuple with the HashiJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiJson

`func (o *GatewayCreateMigration) SetHashiJson(v string)`

SetHashiJson sets HashiJson field to given value.

### HasHashiJson

`func (o *GatewayCreateMigration) HasHashiJson() bool`

HasHashiJson returns a boolean if a field has been set.

### GetHashiNs

`func (o *GatewayCreateMigration) GetHashiNs() []string`

GetHashiNs returns the HashiNs field if non-nil, zero value otherwise.

### GetHashiNsOk

`func (o *GatewayCreateMigration) GetHashiNsOk() (*[]string, bool)`

GetHashiNsOk returns a tuple with the HashiNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiNs

`func (o *GatewayCreateMigration) SetHashiNs(v []string)`

SetHashiNs sets HashiNs field to given value.

### HasHashiNs

`func (o *GatewayCreateMigration) HasHashiNs() bool`

HasHashiNs returns a boolean if a field has been set.

### GetHashiToken

`func (o *GatewayCreateMigration) GetHashiToken() string`

GetHashiToken returns the HashiToken field if non-nil, zero value otherwise.

### GetHashiTokenOk

`func (o *GatewayCreateMigration) GetHashiTokenOk() (*string, bool)`

GetHashiTokenOk returns a tuple with the HashiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiToken

`func (o *GatewayCreateMigration) SetHashiToken(v string)`

SetHashiToken sets HashiToken field to given value.

### HasHashiToken

`func (o *GatewayCreateMigration) HasHashiToken() bool`

HasHashiToken returns a boolean if a field has been set.

### GetHashiUrl

`func (o *GatewayCreateMigration) GetHashiUrl() string`

GetHashiUrl returns the HashiUrl field if non-nil, zero value otherwise.

### GetHashiUrlOk

`func (o *GatewayCreateMigration) GetHashiUrlOk() (*string, bool)`

GetHashiUrlOk returns a tuple with the HashiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiUrl

`func (o *GatewayCreateMigration) SetHashiUrl(v string)`

SetHashiUrl sets HashiUrl field to given value.

### HasHashiUrl

`func (o *GatewayCreateMigration) HasHashiUrl() bool`

HasHashiUrl returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateMigration) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateMigration) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateMigration) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateMigration) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sCaCertificate

`func (o *GatewayCreateMigration) GetK8sCaCertificate() []int32`

GetK8sCaCertificate returns the K8sCaCertificate field if non-nil, zero value otherwise.

### GetK8sCaCertificateOk

`func (o *GatewayCreateMigration) GetK8sCaCertificateOk() (*[]int32, bool)`

GetK8sCaCertificateOk returns a tuple with the K8sCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sCaCertificate

`func (o *GatewayCreateMigration) SetK8sCaCertificate(v []int32)`

SetK8sCaCertificate sets K8sCaCertificate field to given value.

### HasK8sCaCertificate

`func (o *GatewayCreateMigration) HasK8sCaCertificate() bool`

HasK8sCaCertificate returns a boolean if a field has been set.

### GetK8sClientCertificate

`func (o *GatewayCreateMigration) GetK8sClientCertificate() []int32`

GetK8sClientCertificate returns the K8sClientCertificate field if non-nil, zero value otherwise.

### GetK8sClientCertificateOk

`func (o *GatewayCreateMigration) GetK8sClientCertificateOk() (*[]int32, bool)`

GetK8sClientCertificateOk returns a tuple with the K8sClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientCertificate

`func (o *GatewayCreateMigration) SetK8sClientCertificate(v []int32)`

SetK8sClientCertificate sets K8sClientCertificate field to given value.

### HasK8sClientCertificate

`func (o *GatewayCreateMigration) HasK8sClientCertificate() bool`

HasK8sClientCertificate returns a boolean if a field has been set.

### GetK8sClientKey

`func (o *GatewayCreateMigration) GetK8sClientKey() []int32`

GetK8sClientKey returns the K8sClientKey field if non-nil, zero value otherwise.

### GetK8sClientKeyOk

`func (o *GatewayCreateMigration) GetK8sClientKeyOk() (*[]int32, bool)`

GetK8sClientKeyOk returns a tuple with the K8sClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientKey

`func (o *GatewayCreateMigration) SetK8sClientKey(v []int32)`

SetK8sClientKey sets K8sClientKey field to given value.

### HasK8sClientKey

`func (o *GatewayCreateMigration) HasK8sClientKey() bool`

HasK8sClientKey returns a boolean if a field has been set.

### GetK8sNamespace

`func (o *GatewayCreateMigration) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *GatewayCreateMigration) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *GatewayCreateMigration) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *GatewayCreateMigration) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### GetK8sPassword

`func (o *GatewayCreateMigration) GetK8sPassword() string`

GetK8sPassword returns the K8sPassword field if non-nil, zero value otherwise.

### GetK8sPasswordOk

`func (o *GatewayCreateMigration) GetK8sPasswordOk() (*string, bool)`

GetK8sPasswordOk returns a tuple with the K8sPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPassword

`func (o *GatewayCreateMigration) SetK8sPassword(v string)`

SetK8sPassword sets K8sPassword field to given value.

### HasK8sPassword

`func (o *GatewayCreateMigration) HasK8sPassword() bool`

HasK8sPassword returns a boolean if a field has been set.

### GetK8sSkipSystem

`func (o *GatewayCreateMigration) GetK8sSkipSystem() bool`

GetK8sSkipSystem returns the K8sSkipSystem field if non-nil, zero value otherwise.

### GetK8sSkipSystemOk

`func (o *GatewayCreateMigration) GetK8sSkipSystemOk() (*bool, bool)`

GetK8sSkipSystemOk returns a tuple with the K8sSkipSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sSkipSystem

`func (o *GatewayCreateMigration) SetK8sSkipSystem(v bool)`

SetK8sSkipSystem sets K8sSkipSystem field to given value.

### HasK8sSkipSystem

`func (o *GatewayCreateMigration) HasK8sSkipSystem() bool`

HasK8sSkipSystem returns a boolean if a field has been set.

### GetK8sToken

`func (o *GatewayCreateMigration) GetK8sToken() string`

GetK8sToken returns the K8sToken field if non-nil, zero value otherwise.

### GetK8sTokenOk

`func (o *GatewayCreateMigration) GetK8sTokenOk() (*string, bool)`

GetK8sTokenOk returns a tuple with the K8sToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sToken

`func (o *GatewayCreateMigration) SetK8sToken(v string)`

SetK8sToken sets K8sToken field to given value.

### HasK8sToken

`func (o *GatewayCreateMigration) HasK8sToken() bool`

HasK8sToken returns a boolean if a field has been set.

### GetK8sUrl

`func (o *GatewayCreateMigration) GetK8sUrl() string`

GetK8sUrl returns the K8sUrl field if non-nil, zero value otherwise.

### GetK8sUrlOk

`func (o *GatewayCreateMigration) GetK8sUrlOk() (*string, bool)`

GetK8sUrlOk returns a tuple with the K8sUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sUrl

`func (o *GatewayCreateMigration) SetK8sUrl(v string)`

SetK8sUrl sets K8sUrl field to given value.

### HasK8sUrl

`func (o *GatewayCreateMigration) HasK8sUrl() bool`

HasK8sUrl returns a boolean if a field has been set.

### GetK8sUsername

`func (o *GatewayCreateMigration) GetK8sUsername() string`

GetK8sUsername returns the K8sUsername field if non-nil, zero value otherwise.

### GetK8sUsernameOk

`func (o *GatewayCreateMigration) GetK8sUsernameOk() (*string, bool)`

GetK8sUsernameOk returns a tuple with the K8sUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sUsername

`func (o *GatewayCreateMigration) SetK8sUsername(v string)`

SetK8sUsername sets K8sUsername field to given value.

### HasK8sUsername

`func (o *GatewayCreateMigration) HasK8sUsername() bool`

HasK8sUsername returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateMigration) SetName(v string)`

SetName sets Name field to given value.


### GetOpEmail

`func (o *GatewayCreateMigration) GetOpEmail() string`

GetOpEmail returns the OpEmail field if non-nil, zero value otherwise.

### GetOpEmailOk

`func (o *GatewayCreateMigration) GetOpEmailOk() (*string, bool)`

GetOpEmailOk returns a tuple with the OpEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpEmail

`func (o *GatewayCreateMigration) SetOpEmail(v string)`

SetOpEmail sets OpEmail field to given value.

### HasOpEmail

`func (o *GatewayCreateMigration) HasOpEmail() bool`

HasOpEmail returns a boolean if a field has been set.

### GetOpPassword

`func (o *GatewayCreateMigration) GetOpPassword() string`

GetOpPassword returns the OpPassword field if non-nil, zero value otherwise.

### GetOpPasswordOk

`func (o *GatewayCreateMigration) GetOpPasswordOk() (*string, bool)`

GetOpPasswordOk returns a tuple with the OpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpPassword

`func (o *GatewayCreateMigration) SetOpPassword(v string)`

SetOpPassword sets OpPassword field to given value.

### HasOpPassword

`func (o *GatewayCreateMigration) HasOpPassword() bool`

HasOpPassword returns a boolean if a field has been set.

### GetOpSecretKey

`func (o *GatewayCreateMigration) GetOpSecretKey() string`

GetOpSecretKey returns the OpSecretKey field if non-nil, zero value otherwise.

### GetOpSecretKeyOk

`func (o *GatewayCreateMigration) GetOpSecretKeyOk() (*string, bool)`

GetOpSecretKeyOk returns a tuple with the OpSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpSecretKey

`func (o *GatewayCreateMigration) SetOpSecretKey(v string)`

SetOpSecretKey sets OpSecretKey field to given value.

### HasOpSecretKey

`func (o *GatewayCreateMigration) HasOpSecretKey() bool`

HasOpSecretKey returns a boolean if a field has been set.

### GetOpUrl

`func (o *GatewayCreateMigration) GetOpUrl() string`

GetOpUrl returns the OpUrl field if non-nil, zero value otherwise.

### GetOpUrlOk

`func (o *GatewayCreateMigration) GetOpUrlOk() (*string, bool)`

GetOpUrlOk returns a tuple with the OpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpUrl

`func (o *GatewayCreateMigration) SetOpUrl(v string)`

SetOpUrl sets OpUrl field to given value.

### HasOpUrl

`func (o *GatewayCreateMigration) HasOpUrl() bool`

HasOpUrl returns a boolean if a field has been set.

### GetOpVaults

`func (o *GatewayCreateMigration) GetOpVaults() []string`

GetOpVaults returns the OpVaults field if non-nil, zero value otherwise.

### GetOpVaultsOk

`func (o *GatewayCreateMigration) GetOpVaultsOk() (*[]string, bool)`

GetOpVaultsOk returns a tuple with the OpVaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpVaults

`func (o *GatewayCreateMigration) SetOpVaults(v []string)`

SetOpVaults sets OpVaults field to given value.

### HasOpVaults

`func (o *GatewayCreateMigration) HasOpVaults() bool`

HasOpVaults returns a boolean if a field has been set.

### GetProtectionKey

`func (o *GatewayCreateMigration) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *GatewayCreateMigration) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *GatewayCreateMigration) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *GatewayCreateMigration) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTargetLocation

`func (o *GatewayCreateMigration) GetTargetLocation() string`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *GatewayCreateMigration) GetTargetLocationOk() (*string, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *GatewayCreateMigration) SetTargetLocation(v string)`

SetTargetLocation sets TargetLocation field to given value.

### HasTargetLocation

`func (o *GatewayCreateMigration) HasTargetLocation() bool`

HasTargetLocation returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateMigration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateMigration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateMigration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateMigration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *GatewayCreateMigration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayCreateMigration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayCreateMigration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayCreateMigration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateMigration) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateMigration) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateMigration) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateMigration) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


