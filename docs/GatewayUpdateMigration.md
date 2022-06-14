# GatewayUpdateMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKey** | Pointer to **string** | AWS Secret Access Key | [optional] 
**AwsKeyId** | Pointer to **string** | AWS Access Key ID | [optional] 
**AwsRegion** | Pointer to **string** | AWS region | [optional] 
**AzureClientId** | Pointer to **string** | Azure KV Access client ID | [optional] 
**AzureKvName** | Pointer to **string** | Azure Key Vault Name | [optional] 
**AzureSecret** | Pointer to **string** | Azure KV secret | [optional] 
**AzureTenantId** | Pointer to **string** | Azure KV Access tenant ID | [optional] 
**GcpKey** | Pointer to **string** | Base64-encoded service account private key text | [optional] 
**HashiJson** | Pointer to **string** | Import secret key as json value or independent secrets | [optional] 
**HashiNs** | Pointer to **[]string** | Hashi namespaces | [optional] 
**HashiToken** | Pointer to **string** | Hashi token | [optional] 
**HashiUrl** | Pointer to **string** | Hashi url | [optional] 
**Id** | Pointer to **string** | Migration ID | [optional] 
**K8sCaCertificate** | Pointer to **[]int32** | For Certificate Authentication method K8s Cluster CA certificate | [optional] 
**K8sClientCertificate** | Pointer to **[]int32** | K8s Client certificate | [optional] 
**K8sClientKey** | Pointer to **[]int32** | K8s Client key | [optional] 
**K8sNamespace** | Pointer to **string** | K8s Namespace | [optional] 
**K8sPassword** | Pointer to **string** | K8s client password | [optional] 
**K8sSkipSystem** | Pointer to **bool** | K8s Skip Control Plane Secrets | [optional] 
**K8sToken** | Pointer to **string** | For Token Authentication method K8s Bearer Token | [optional] 
**K8sUrl** | Pointer to **string** | K8s Endpoint URL | [optional] 
**K8sUsername** | Pointer to **string** | For Password Authentication method K8s client username | [optional] 
**Name** | **string** | Migration name | 
**ProtectionKey** | Pointer to **string** | The name of the key that protects the classic key value (if empty, the account default key will be used) | [optional] 
**TargetLocation** | Pointer to **string** | Target location in Akeyless for imported secrets | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **string** | Migration type, can be: hashi/aws/gcp/k8s/azure_kv | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateMigration

`func NewGatewayUpdateMigration(name string, ) *GatewayUpdateMigration`

NewGatewayUpdateMigration instantiates a new GatewayUpdateMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateMigrationWithDefaults

`func NewGatewayUpdateMigrationWithDefaults() *GatewayUpdateMigration`

NewGatewayUpdateMigrationWithDefaults instantiates a new GatewayUpdateMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKey

`func (o *GatewayUpdateMigration) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *GatewayUpdateMigration) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *GatewayUpdateMigration) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.

### HasAwsKey

`func (o *GatewayUpdateMigration) HasAwsKey() bool`

HasAwsKey returns a boolean if a field has been set.

### GetAwsKeyId

`func (o *GatewayUpdateMigration) GetAwsKeyId() string`

GetAwsKeyId returns the AwsKeyId field if non-nil, zero value otherwise.

### GetAwsKeyIdOk

`func (o *GatewayUpdateMigration) GetAwsKeyIdOk() (*string, bool)`

GetAwsKeyIdOk returns a tuple with the AwsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKeyId

`func (o *GatewayUpdateMigration) SetAwsKeyId(v string)`

SetAwsKeyId sets AwsKeyId field to given value.

### HasAwsKeyId

`func (o *GatewayUpdateMigration) HasAwsKeyId() bool`

HasAwsKeyId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *GatewayUpdateMigration) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *GatewayUpdateMigration) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *GatewayUpdateMigration) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *GatewayUpdateMigration) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAzureClientId

`func (o *GatewayUpdateMigration) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *GatewayUpdateMigration) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *GatewayUpdateMigration) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *GatewayUpdateMigration) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureKvName

`func (o *GatewayUpdateMigration) GetAzureKvName() string`

GetAzureKvName returns the AzureKvName field if non-nil, zero value otherwise.

### GetAzureKvNameOk

`func (o *GatewayUpdateMigration) GetAzureKvNameOk() (*string, bool)`

GetAzureKvNameOk returns a tuple with the AzureKvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKvName

`func (o *GatewayUpdateMigration) SetAzureKvName(v string)`

SetAzureKvName sets AzureKvName field to given value.

### HasAzureKvName

`func (o *GatewayUpdateMigration) HasAzureKvName() bool`

HasAzureKvName returns a boolean if a field has been set.

### GetAzureSecret

`func (o *GatewayUpdateMigration) GetAzureSecret() string`

GetAzureSecret returns the AzureSecret field if non-nil, zero value otherwise.

### GetAzureSecretOk

`func (o *GatewayUpdateMigration) GetAzureSecretOk() (*string, bool)`

GetAzureSecretOk returns a tuple with the AzureSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSecret

`func (o *GatewayUpdateMigration) SetAzureSecret(v string)`

SetAzureSecret sets AzureSecret field to given value.

### HasAzureSecret

`func (o *GatewayUpdateMigration) HasAzureSecret() bool`

HasAzureSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *GatewayUpdateMigration) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *GatewayUpdateMigration) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *GatewayUpdateMigration) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *GatewayUpdateMigration) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetGcpKey

`func (o *GatewayUpdateMigration) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *GatewayUpdateMigration) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *GatewayUpdateMigration) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *GatewayUpdateMigration) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetHashiJson

`func (o *GatewayUpdateMigration) GetHashiJson() string`

GetHashiJson returns the HashiJson field if non-nil, zero value otherwise.

### GetHashiJsonOk

`func (o *GatewayUpdateMigration) GetHashiJsonOk() (*string, bool)`

GetHashiJsonOk returns a tuple with the HashiJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiJson

`func (o *GatewayUpdateMigration) SetHashiJson(v string)`

SetHashiJson sets HashiJson field to given value.

### HasHashiJson

`func (o *GatewayUpdateMigration) HasHashiJson() bool`

HasHashiJson returns a boolean if a field has been set.

### GetHashiNs

`func (o *GatewayUpdateMigration) GetHashiNs() []string`

GetHashiNs returns the HashiNs field if non-nil, zero value otherwise.

### GetHashiNsOk

`func (o *GatewayUpdateMigration) GetHashiNsOk() (*[]string, bool)`

GetHashiNsOk returns a tuple with the HashiNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiNs

`func (o *GatewayUpdateMigration) SetHashiNs(v []string)`

SetHashiNs sets HashiNs field to given value.

### HasHashiNs

`func (o *GatewayUpdateMigration) HasHashiNs() bool`

HasHashiNs returns a boolean if a field has been set.

### GetHashiToken

`func (o *GatewayUpdateMigration) GetHashiToken() string`

GetHashiToken returns the HashiToken field if non-nil, zero value otherwise.

### GetHashiTokenOk

`func (o *GatewayUpdateMigration) GetHashiTokenOk() (*string, bool)`

GetHashiTokenOk returns a tuple with the HashiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiToken

`func (o *GatewayUpdateMigration) SetHashiToken(v string)`

SetHashiToken sets HashiToken field to given value.

### HasHashiToken

`func (o *GatewayUpdateMigration) HasHashiToken() bool`

HasHashiToken returns a boolean if a field has been set.

### GetHashiUrl

`func (o *GatewayUpdateMigration) GetHashiUrl() string`

GetHashiUrl returns the HashiUrl field if non-nil, zero value otherwise.

### GetHashiUrlOk

`func (o *GatewayUpdateMigration) GetHashiUrlOk() (*string, bool)`

GetHashiUrlOk returns a tuple with the HashiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiUrl

`func (o *GatewayUpdateMigration) SetHashiUrl(v string)`

SetHashiUrl sets HashiUrl field to given value.

### HasHashiUrl

`func (o *GatewayUpdateMigration) HasHashiUrl() bool`

HasHashiUrl returns a boolean if a field has been set.

### GetId

`func (o *GatewayUpdateMigration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayUpdateMigration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayUpdateMigration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayUpdateMigration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetK8sCaCertificate

`func (o *GatewayUpdateMigration) GetK8sCaCertificate() []int32`

GetK8sCaCertificate returns the K8sCaCertificate field if non-nil, zero value otherwise.

### GetK8sCaCertificateOk

`func (o *GatewayUpdateMigration) GetK8sCaCertificateOk() (*[]int32, bool)`

GetK8sCaCertificateOk returns a tuple with the K8sCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sCaCertificate

`func (o *GatewayUpdateMigration) SetK8sCaCertificate(v []int32)`

SetK8sCaCertificate sets K8sCaCertificate field to given value.

### HasK8sCaCertificate

`func (o *GatewayUpdateMigration) HasK8sCaCertificate() bool`

HasK8sCaCertificate returns a boolean if a field has been set.

### GetK8sClientCertificate

`func (o *GatewayUpdateMigration) GetK8sClientCertificate() []int32`

GetK8sClientCertificate returns the K8sClientCertificate field if non-nil, zero value otherwise.

### GetK8sClientCertificateOk

`func (o *GatewayUpdateMigration) GetK8sClientCertificateOk() (*[]int32, bool)`

GetK8sClientCertificateOk returns a tuple with the K8sClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientCertificate

`func (o *GatewayUpdateMigration) SetK8sClientCertificate(v []int32)`

SetK8sClientCertificate sets K8sClientCertificate field to given value.

### HasK8sClientCertificate

`func (o *GatewayUpdateMigration) HasK8sClientCertificate() bool`

HasK8sClientCertificate returns a boolean if a field has been set.

### GetK8sClientKey

`func (o *GatewayUpdateMigration) GetK8sClientKey() []int32`

GetK8sClientKey returns the K8sClientKey field if non-nil, zero value otherwise.

### GetK8sClientKeyOk

`func (o *GatewayUpdateMigration) GetK8sClientKeyOk() (*[]int32, bool)`

GetK8sClientKeyOk returns a tuple with the K8sClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientKey

`func (o *GatewayUpdateMigration) SetK8sClientKey(v []int32)`

SetK8sClientKey sets K8sClientKey field to given value.

### HasK8sClientKey

`func (o *GatewayUpdateMigration) HasK8sClientKey() bool`

HasK8sClientKey returns a boolean if a field has been set.

### GetK8sNamespace

`func (o *GatewayUpdateMigration) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *GatewayUpdateMigration) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *GatewayUpdateMigration) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *GatewayUpdateMigration) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### GetK8sPassword

`func (o *GatewayUpdateMigration) GetK8sPassword() string`

GetK8sPassword returns the K8sPassword field if non-nil, zero value otherwise.

### GetK8sPasswordOk

`func (o *GatewayUpdateMigration) GetK8sPasswordOk() (*string, bool)`

GetK8sPasswordOk returns a tuple with the K8sPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPassword

`func (o *GatewayUpdateMigration) SetK8sPassword(v string)`

SetK8sPassword sets K8sPassword field to given value.

### HasK8sPassword

`func (o *GatewayUpdateMigration) HasK8sPassword() bool`

HasK8sPassword returns a boolean if a field has been set.

### GetK8sSkipSystem

`func (o *GatewayUpdateMigration) GetK8sSkipSystem() bool`

GetK8sSkipSystem returns the K8sSkipSystem field if non-nil, zero value otherwise.

### GetK8sSkipSystemOk

`func (o *GatewayUpdateMigration) GetK8sSkipSystemOk() (*bool, bool)`

GetK8sSkipSystemOk returns a tuple with the K8sSkipSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sSkipSystem

`func (o *GatewayUpdateMigration) SetK8sSkipSystem(v bool)`

SetK8sSkipSystem sets K8sSkipSystem field to given value.

### HasK8sSkipSystem

`func (o *GatewayUpdateMigration) HasK8sSkipSystem() bool`

HasK8sSkipSystem returns a boolean if a field has been set.

### GetK8sToken

`func (o *GatewayUpdateMigration) GetK8sToken() string`

GetK8sToken returns the K8sToken field if non-nil, zero value otherwise.

### GetK8sTokenOk

`func (o *GatewayUpdateMigration) GetK8sTokenOk() (*string, bool)`

GetK8sTokenOk returns a tuple with the K8sToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sToken

`func (o *GatewayUpdateMigration) SetK8sToken(v string)`

SetK8sToken sets K8sToken field to given value.

### HasK8sToken

`func (o *GatewayUpdateMigration) HasK8sToken() bool`

HasK8sToken returns a boolean if a field has been set.

### GetK8sUrl

`func (o *GatewayUpdateMigration) GetK8sUrl() string`

GetK8sUrl returns the K8sUrl field if non-nil, zero value otherwise.

### GetK8sUrlOk

`func (o *GatewayUpdateMigration) GetK8sUrlOk() (*string, bool)`

GetK8sUrlOk returns a tuple with the K8sUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sUrl

`func (o *GatewayUpdateMigration) SetK8sUrl(v string)`

SetK8sUrl sets K8sUrl field to given value.

### HasK8sUrl

`func (o *GatewayUpdateMigration) HasK8sUrl() bool`

HasK8sUrl returns a boolean if a field has been set.

### GetK8sUsername

`func (o *GatewayUpdateMigration) GetK8sUsername() string`

GetK8sUsername returns the K8sUsername field if non-nil, zero value otherwise.

### GetK8sUsernameOk

`func (o *GatewayUpdateMigration) GetK8sUsernameOk() (*string, bool)`

GetK8sUsernameOk returns a tuple with the K8sUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sUsername

`func (o *GatewayUpdateMigration) SetK8sUsername(v string)`

SetK8sUsername sets K8sUsername field to given value.

### HasK8sUsername

`func (o *GatewayUpdateMigration) HasK8sUsername() bool`

HasK8sUsername returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateMigration) SetName(v string)`

SetName sets Name field to given value.


### GetProtectionKey

`func (o *GatewayUpdateMigration) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *GatewayUpdateMigration) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *GatewayUpdateMigration) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *GatewayUpdateMigration) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTargetLocation

`func (o *GatewayUpdateMigration) GetTargetLocation() string`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *GatewayUpdateMigration) GetTargetLocationOk() (*string, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *GatewayUpdateMigration) SetTargetLocation(v string)`

SetTargetLocation sets TargetLocation field to given value.

### HasTargetLocation

`func (o *GatewayUpdateMigration) HasTargetLocation() bool`

HasTargetLocation returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateMigration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateMigration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateMigration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateMigration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *GatewayUpdateMigration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayUpdateMigration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayUpdateMigration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayUpdateMigration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateMigration) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateMigration) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateMigration) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateMigration) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


