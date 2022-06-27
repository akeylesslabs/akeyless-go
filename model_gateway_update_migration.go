/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GatewayUpdateMigration gatewayUpdateMigration is a command that update migration
type GatewayUpdateMigration struct {
	// AWS Secret Access Key (relevant only for AWS migration)
	AwsKey *string `json:"aws-key,omitempty"`
	// AWS Access Key ID with sufficient permissions to get all secrets, e.g. 'arn:aws:secretsmanager:[Region]:[AccountId]:secret:[/path/to/secrets/_*]' (relevant only for AWS migration)
	AwsKeyId *string `json:"aws-key-id,omitempty"`
	// AWS region of the required Secrets Manager (relevant only for AWS migration)
	AwsRegion *string `json:"aws-region,omitempty"`
	// Azure Key Vault Access client ID, should be Azure AD App with a service principal (relevant only for Azure Key Vault migration)
	AzureClientId *string `json:"azure-client-id,omitempty"`
	// Azure Key Vault Name (relevant only for Azure Key Vault migration)
	AzureKvName *string `json:"azure-kv-name,omitempty"`
	// Azure Key Vault secret (relevant only for Azure Key Vault migration)
	AzureSecret *string `json:"azure-secret,omitempty"`
	// Azure Key Vault Access tenant ID (relevant only for Azure Key Vault migration)
	AzureTenantId *string `json:"azure-tenant-id,omitempty"`
	// Base64-encoded GCP Service Account private key text with sufficient permissions to Secrets Manager, Minimum required permission is Secret Manager Secret Accessor, e.g. 'roles/secretmanager.secretAccessor' (relevant only for GCP migration)
	GcpKey *string `json:"gcp-key,omitempty"`
	// Import secret key as json value or independent secrets (relevant only for HasiCorp Vault migration)
	HashiJson *string `json:"hashi-json,omitempty"`
	// HashiCorp Vault Namespaces is a comma-separated list of namespaces which need to be imported into Akeyless Vault. For every provided namespace, all its child namespaces are imported as well, e.g. nmsp/subnmsp1/subnmsp2,nmsp/anothernmsp. By default, import all namespaces (relevant only for HasiCorp Vault migration)
	HashiNs *[]string `json:"hashi-ns,omitempty"`
	// HashiCorp Vault access token with sufficient permissions to preform list & read operations on secrets objects (relevant only for HasiCorp Vault migration)
	HashiToken *string `json:"hashi-token,omitempty"`
	// HashiCorp Vault API URL, e.g. https://vault-mgr01:8200 (relevant only for HasiCorp Vault migration)
	HashiUrl *string `json:"hashi-url,omitempty"`
	// Migration ID (Can be retrieved with gateway-list-migration command)
	Id *string `json:"id,omitempty"`
	// For Certificate Authentication method K8s Cluster CA certificate (relevant only for K8s migration with Certificate Authentication method)
	K8sCaCertificate *[]int32 `json:"k8s-ca-certificate,omitempty"`
	// K8s Client certificate with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Certificate Authentication method)
	K8sClientCertificate *[]int32 `json:"k8s-client-certificate,omitempty"`
	// K8s Client key (relevant only for K8s migration with Certificate Authentication method)
	K8sClientKey *[]int32 `json:"k8s-client-key,omitempty"`
	// K8s Namespace, Use this field to import secrets from a particular namespace only. By default, the secrets are imported from all namespaces (relevant only for K8s migration)
	K8sNamespace *string `json:"k8s-namespace,omitempty"`
	// K8s Client password (relevant only for K8s migration with Password Authentication method)
	K8sPassword *string `json:"k8s-password,omitempty"`
	// K8s Skip Control Plane Secrets, This option allows to avoid importing secrets from system namespaces (relevant only for K8s migration)
	K8sSkipSystem *bool `json:"k8s-skip-system,omitempty"`
	// For Token Authentication method K8s Bearer Token with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Token Authentication method)
	K8sToken *string `json:"k8s-token,omitempty"`
	// K8s API Server URL, e.g. https://k8s-api.mycompany.com:6443 (relevant only for K8s migration)
	K8sUrl *string `json:"k8s-url,omitempty"`
	// For Password Authentication method K8s Client username with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Password Authentication method)
	K8sUsername *string `json:"k8s-username,omitempty"`
	// Migration name
	Name *string `json:"name,omitempty"`
	// New migration name
	NewName *string `json:"new_name,omitempty"`
	// The name of the key that protects the classic key value (if empty, the account default key will be used)
	ProtectionKey *string `json:"protection-key,omitempty"`
	// Target location in Akeyless for imported secrets
	TargetLocation *string `json:"target-location,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayUpdateMigration instantiates a new GatewayUpdateMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateMigration() *GatewayUpdateMigration {
	this := GatewayUpdateMigration{}
	return &this
}

// NewGatewayUpdateMigrationWithDefaults instantiates a new GatewayUpdateMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateMigrationWithDefaults() *GatewayUpdateMigration {
	this := GatewayUpdateMigration{}
	return &this
}

// GetAwsKey returns the AwsKey field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetAwsKey() string {
	if o == nil || o.AwsKey == nil {
		var ret string
		return ret
	}
	return *o.AwsKey
}

// GetAwsKeyOk returns a tuple with the AwsKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetAwsKeyOk() (*string, bool) {
	if o == nil || o.AwsKey == nil {
		return nil, false
	}
	return o.AwsKey, true
}

// HasAwsKey returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasAwsKey() bool {
	if o != nil && o.AwsKey != nil {
		return true
	}

	return false
}

// SetAwsKey gets a reference to the given string and assigns it to the AwsKey field.
func (o *GatewayUpdateMigration) SetAwsKey(v string) {
	o.AwsKey = &v
}

// GetAwsKeyId returns the AwsKeyId field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetAwsKeyId() string {
	if o == nil || o.AwsKeyId == nil {
		var ret string
		return ret
	}
	return *o.AwsKeyId
}

// GetAwsKeyIdOk returns a tuple with the AwsKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetAwsKeyIdOk() (*string, bool) {
	if o == nil || o.AwsKeyId == nil {
		return nil, false
	}
	return o.AwsKeyId, true
}

// HasAwsKeyId returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasAwsKeyId() bool {
	if o != nil && o.AwsKeyId != nil {
		return true
	}

	return false
}

// SetAwsKeyId gets a reference to the given string and assigns it to the AwsKeyId field.
func (o *GatewayUpdateMigration) SetAwsKeyId(v string) {
	o.AwsKeyId = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetAwsRegion() string {
	if o == nil || o.AwsRegion == nil {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetAwsRegionOk() (*string, bool) {
	if o == nil || o.AwsRegion == nil {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasAwsRegion() bool {
	if o != nil && o.AwsRegion != nil {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *GatewayUpdateMigration) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetAzureClientId() string {
	if o == nil || o.AzureClientId == nil {
		var ret string
		return ret
	}
	return *o.AzureClientId
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetAzureClientIdOk() (*string, bool) {
	if o == nil || o.AzureClientId == nil {
		return nil, false
	}
	return o.AzureClientId, true
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasAzureClientId() bool {
	if o != nil && o.AzureClientId != nil {
		return true
	}

	return false
}

// SetAzureClientId gets a reference to the given string and assigns it to the AzureClientId field.
func (o *GatewayUpdateMigration) SetAzureClientId(v string) {
	o.AzureClientId = &v
}

// GetAzureKvName returns the AzureKvName field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetAzureKvName() string {
	if o == nil || o.AzureKvName == nil {
		var ret string
		return ret
	}
	return *o.AzureKvName
}

// GetAzureKvNameOk returns a tuple with the AzureKvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetAzureKvNameOk() (*string, bool) {
	if o == nil || o.AzureKvName == nil {
		return nil, false
	}
	return o.AzureKvName, true
}

// HasAzureKvName returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasAzureKvName() bool {
	if o != nil && o.AzureKvName != nil {
		return true
	}

	return false
}

// SetAzureKvName gets a reference to the given string and assigns it to the AzureKvName field.
func (o *GatewayUpdateMigration) SetAzureKvName(v string) {
	o.AzureKvName = &v
}

// GetAzureSecret returns the AzureSecret field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetAzureSecret() string {
	if o == nil || o.AzureSecret == nil {
		var ret string
		return ret
	}
	return *o.AzureSecret
}

// GetAzureSecretOk returns a tuple with the AzureSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetAzureSecretOk() (*string, bool) {
	if o == nil || o.AzureSecret == nil {
		return nil, false
	}
	return o.AzureSecret, true
}

// HasAzureSecret returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasAzureSecret() bool {
	if o != nil && o.AzureSecret != nil {
		return true
	}

	return false
}

// SetAzureSecret gets a reference to the given string and assigns it to the AzureSecret field.
func (o *GatewayUpdateMigration) SetAzureSecret(v string) {
	o.AzureSecret = &v
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || o.AzureTenantId == nil {
		return nil, false
	}
	return o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId != nil {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *GatewayUpdateMigration) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetGcpKey returns the GcpKey field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetGcpKey() string {
	if o == nil || o.GcpKey == nil {
		var ret string
		return ret
	}
	return *o.GcpKey
}

// GetGcpKeyOk returns a tuple with the GcpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetGcpKeyOk() (*string, bool) {
	if o == nil || o.GcpKey == nil {
		return nil, false
	}
	return o.GcpKey, true
}

// HasGcpKey returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasGcpKey() bool {
	if o != nil && o.GcpKey != nil {
		return true
	}

	return false
}

// SetGcpKey gets a reference to the given string and assigns it to the GcpKey field.
func (o *GatewayUpdateMigration) SetGcpKey(v string) {
	o.GcpKey = &v
}

// GetHashiJson returns the HashiJson field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetHashiJson() string {
	if o == nil || o.HashiJson == nil {
		var ret string
		return ret
	}
	return *o.HashiJson
}

// GetHashiJsonOk returns a tuple with the HashiJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetHashiJsonOk() (*string, bool) {
	if o == nil || o.HashiJson == nil {
		return nil, false
	}
	return o.HashiJson, true
}

// HasHashiJson returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasHashiJson() bool {
	if o != nil && o.HashiJson != nil {
		return true
	}

	return false
}

// SetHashiJson gets a reference to the given string and assigns it to the HashiJson field.
func (o *GatewayUpdateMigration) SetHashiJson(v string) {
	o.HashiJson = &v
}

// GetHashiNs returns the HashiNs field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetHashiNs() []string {
	if o == nil || o.HashiNs == nil {
		var ret []string
		return ret
	}
	return *o.HashiNs
}

// GetHashiNsOk returns a tuple with the HashiNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetHashiNsOk() (*[]string, bool) {
	if o == nil || o.HashiNs == nil {
		return nil, false
	}
	return o.HashiNs, true
}

// HasHashiNs returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasHashiNs() bool {
	if o != nil && o.HashiNs != nil {
		return true
	}

	return false
}

// SetHashiNs gets a reference to the given []string and assigns it to the HashiNs field.
func (o *GatewayUpdateMigration) SetHashiNs(v []string) {
	o.HashiNs = &v
}

// GetHashiToken returns the HashiToken field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetHashiToken() string {
	if o == nil || o.HashiToken == nil {
		var ret string
		return ret
	}
	return *o.HashiToken
}

// GetHashiTokenOk returns a tuple with the HashiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetHashiTokenOk() (*string, bool) {
	if o == nil || o.HashiToken == nil {
		return nil, false
	}
	return o.HashiToken, true
}

// HasHashiToken returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasHashiToken() bool {
	if o != nil && o.HashiToken != nil {
		return true
	}

	return false
}

// SetHashiToken gets a reference to the given string and assigns it to the HashiToken field.
func (o *GatewayUpdateMigration) SetHashiToken(v string) {
	o.HashiToken = &v
}

// GetHashiUrl returns the HashiUrl field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetHashiUrl() string {
	if o == nil || o.HashiUrl == nil {
		var ret string
		return ret
	}
	return *o.HashiUrl
}

// GetHashiUrlOk returns a tuple with the HashiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetHashiUrlOk() (*string, bool) {
	if o == nil || o.HashiUrl == nil {
		return nil, false
	}
	return o.HashiUrl, true
}

// HasHashiUrl returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasHashiUrl() bool {
	if o != nil && o.HashiUrl != nil {
		return true
	}

	return false
}

// SetHashiUrl gets a reference to the given string and assigns it to the HashiUrl field.
func (o *GatewayUpdateMigration) SetHashiUrl(v string) {
	o.HashiUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GatewayUpdateMigration) SetId(v string) {
	o.Id = &v
}

// GetK8sCaCertificate returns the K8sCaCertificate field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetK8sCaCertificate() []int32 {
	if o == nil || o.K8sCaCertificate == nil {
		var ret []int32
		return ret
	}
	return *o.K8sCaCertificate
}

// GetK8sCaCertificateOk returns a tuple with the K8sCaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetK8sCaCertificateOk() (*[]int32, bool) {
	if o == nil || o.K8sCaCertificate == nil {
		return nil, false
	}
	return o.K8sCaCertificate, true
}

// HasK8sCaCertificate returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasK8sCaCertificate() bool {
	if o != nil && o.K8sCaCertificate != nil {
		return true
	}

	return false
}

// SetK8sCaCertificate gets a reference to the given []int32 and assigns it to the K8sCaCertificate field.
func (o *GatewayUpdateMigration) SetK8sCaCertificate(v []int32) {
	o.K8sCaCertificate = &v
}

// GetK8sClientCertificate returns the K8sClientCertificate field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetK8sClientCertificate() []int32 {
	if o == nil || o.K8sClientCertificate == nil {
		var ret []int32
		return ret
	}
	return *o.K8sClientCertificate
}

// GetK8sClientCertificateOk returns a tuple with the K8sClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetK8sClientCertificateOk() (*[]int32, bool) {
	if o == nil || o.K8sClientCertificate == nil {
		return nil, false
	}
	return o.K8sClientCertificate, true
}

// HasK8sClientCertificate returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasK8sClientCertificate() bool {
	if o != nil && o.K8sClientCertificate != nil {
		return true
	}

	return false
}

// SetK8sClientCertificate gets a reference to the given []int32 and assigns it to the K8sClientCertificate field.
func (o *GatewayUpdateMigration) SetK8sClientCertificate(v []int32) {
	o.K8sClientCertificate = &v
}

// GetK8sClientKey returns the K8sClientKey field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetK8sClientKey() []int32 {
	if o == nil || o.K8sClientKey == nil {
		var ret []int32
		return ret
	}
	return *o.K8sClientKey
}

// GetK8sClientKeyOk returns a tuple with the K8sClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetK8sClientKeyOk() (*[]int32, bool) {
	if o == nil || o.K8sClientKey == nil {
		return nil, false
	}
	return o.K8sClientKey, true
}

// HasK8sClientKey returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasK8sClientKey() bool {
	if o != nil && o.K8sClientKey != nil {
		return true
	}

	return false
}

// SetK8sClientKey gets a reference to the given []int32 and assigns it to the K8sClientKey field.
func (o *GatewayUpdateMigration) SetK8sClientKey(v []int32) {
	o.K8sClientKey = &v
}

// GetK8sNamespace returns the K8sNamespace field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetK8sNamespace() string {
	if o == nil || o.K8sNamespace == nil {
		var ret string
		return ret
	}
	return *o.K8sNamespace
}

// GetK8sNamespaceOk returns a tuple with the K8sNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetK8sNamespaceOk() (*string, bool) {
	if o == nil || o.K8sNamespace == nil {
		return nil, false
	}
	return o.K8sNamespace, true
}

// HasK8sNamespace returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasK8sNamespace() bool {
	if o != nil && o.K8sNamespace != nil {
		return true
	}

	return false
}

// SetK8sNamespace gets a reference to the given string and assigns it to the K8sNamespace field.
func (o *GatewayUpdateMigration) SetK8sNamespace(v string) {
	o.K8sNamespace = &v
}

// GetK8sPassword returns the K8sPassword field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetK8sPassword() string {
	if o == nil || o.K8sPassword == nil {
		var ret string
		return ret
	}
	return *o.K8sPassword
}

// GetK8sPasswordOk returns a tuple with the K8sPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetK8sPasswordOk() (*string, bool) {
	if o == nil || o.K8sPassword == nil {
		return nil, false
	}
	return o.K8sPassword, true
}

// HasK8sPassword returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasK8sPassword() bool {
	if o != nil && o.K8sPassword != nil {
		return true
	}

	return false
}

// SetK8sPassword gets a reference to the given string and assigns it to the K8sPassword field.
func (o *GatewayUpdateMigration) SetK8sPassword(v string) {
	o.K8sPassword = &v
}

// GetK8sSkipSystem returns the K8sSkipSystem field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetK8sSkipSystem() bool {
	if o == nil || o.K8sSkipSystem == nil {
		var ret bool
		return ret
	}
	return *o.K8sSkipSystem
}

// GetK8sSkipSystemOk returns a tuple with the K8sSkipSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetK8sSkipSystemOk() (*bool, bool) {
	if o == nil || o.K8sSkipSystem == nil {
		return nil, false
	}
	return o.K8sSkipSystem, true
}

// HasK8sSkipSystem returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasK8sSkipSystem() bool {
	if o != nil && o.K8sSkipSystem != nil {
		return true
	}

	return false
}

// SetK8sSkipSystem gets a reference to the given bool and assigns it to the K8sSkipSystem field.
func (o *GatewayUpdateMigration) SetK8sSkipSystem(v bool) {
	o.K8sSkipSystem = &v
}

// GetK8sToken returns the K8sToken field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetK8sToken() string {
	if o == nil || o.K8sToken == nil {
		var ret string
		return ret
	}
	return *o.K8sToken
}

// GetK8sTokenOk returns a tuple with the K8sToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetK8sTokenOk() (*string, bool) {
	if o == nil || o.K8sToken == nil {
		return nil, false
	}
	return o.K8sToken, true
}

// HasK8sToken returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasK8sToken() bool {
	if o != nil && o.K8sToken != nil {
		return true
	}

	return false
}

// SetK8sToken gets a reference to the given string and assigns it to the K8sToken field.
func (o *GatewayUpdateMigration) SetK8sToken(v string) {
	o.K8sToken = &v
}

// GetK8sUrl returns the K8sUrl field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetK8sUrl() string {
	if o == nil || o.K8sUrl == nil {
		var ret string
		return ret
	}
	return *o.K8sUrl
}

// GetK8sUrlOk returns a tuple with the K8sUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetK8sUrlOk() (*string, bool) {
	if o == nil || o.K8sUrl == nil {
		return nil, false
	}
	return o.K8sUrl, true
}

// HasK8sUrl returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasK8sUrl() bool {
	if o != nil && o.K8sUrl != nil {
		return true
	}

	return false
}

// SetK8sUrl gets a reference to the given string and assigns it to the K8sUrl field.
func (o *GatewayUpdateMigration) SetK8sUrl(v string) {
	o.K8sUrl = &v
}

// GetK8sUsername returns the K8sUsername field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetK8sUsername() string {
	if o == nil || o.K8sUsername == nil {
		var ret string
		return ret
	}
	return *o.K8sUsername
}

// GetK8sUsernameOk returns a tuple with the K8sUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetK8sUsernameOk() (*string, bool) {
	if o == nil || o.K8sUsername == nil {
		return nil, false
	}
	return o.K8sUsername, true
}

// HasK8sUsername returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasK8sUsername() bool {
	if o != nil && o.K8sUsername != nil {
		return true
	}

	return false
}

// SetK8sUsername gets a reference to the given string and assigns it to the K8sUsername field.
func (o *GatewayUpdateMigration) SetK8sUsername(v string) {
	o.K8sUsername = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GatewayUpdateMigration) SetName(v string) {
	o.Name = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateMigration) SetNewName(v string) {
	o.NewName = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetProtectionKey() string {
	if o == nil || o.ProtectionKey == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetProtectionKeyOk() (*string, bool) {
	if o == nil || o.ProtectionKey == nil {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasProtectionKey() bool {
	if o != nil && o.ProtectionKey != nil {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *GatewayUpdateMigration) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetTargetLocation returns the TargetLocation field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetTargetLocation() string {
	if o == nil || o.TargetLocation == nil {
		var ret string
		return ret
	}
	return *o.TargetLocation
}

// GetTargetLocationOk returns a tuple with the TargetLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetTargetLocationOk() (*string, bool) {
	if o == nil || o.TargetLocation == nil {
		return nil, false
	}
	return o.TargetLocation, true
}

// HasTargetLocation returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasTargetLocation() bool {
	if o != nil && o.TargetLocation != nil {
		return true
	}

	return false
}

// SetTargetLocation gets a reference to the given string and assigns it to the TargetLocation field.
func (o *GatewayUpdateMigration) SetTargetLocation(v string) {
	o.TargetLocation = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateMigration) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateMigration) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateMigration) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateMigration) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateMigration) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayUpdateMigration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsKey != nil {
		toSerialize["aws-key"] = o.AwsKey
	}
	if o.AwsKeyId != nil {
		toSerialize["aws-key-id"] = o.AwsKeyId
	}
	if o.AwsRegion != nil {
		toSerialize["aws-region"] = o.AwsRegion
	}
	if o.AzureClientId != nil {
		toSerialize["azure-client-id"] = o.AzureClientId
	}
	if o.AzureKvName != nil {
		toSerialize["azure-kv-name"] = o.AzureKvName
	}
	if o.AzureSecret != nil {
		toSerialize["azure-secret"] = o.AzureSecret
	}
	if o.AzureTenantId != nil {
		toSerialize["azure-tenant-id"] = o.AzureTenantId
	}
	if o.GcpKey != nil {
		toSerialize["gcp-key"] = o.GcpKey
	}
	if o.HashiJson != nil {
		toSerialize["hashi-json"] = o.HashiJson
	}
	if o.HashiNs != nil {
		toSerialize["hashi-ns"] = o.HashiNs
	}
	if o.HashiToken != nil {
		toSerialize["hashi-token"] = o.HashiToken
	}
	if o.HashiUrl != nil {
		toSerialize["hashi-url"] = o.HashiUrl
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.K8sCaCertificate != nil {
		toSerialize["k8s-ca-certificate"] = o.K8sCaCertificate
	}
	if o.K8sClientCertificate != nil {
		toSerialize["k8s-client-certificate"] = o.K8sClientCertificate
	}
	if o.K8sClientKey != nil {
		toSerialize["k8s-client-key"] = o.K8sClientKey
	}
	if o.K8sNamespace != nil {
		toSerialize["k8s-namespace"] = o.K8sNamespace
	}
	if o.K8sPassword != nil {
		toSerialize["k8s-password"] = o.K8sPassword
	}
	if o.K8sSkipSystem != nil {
		toSerialize["k8s-skip-system"] = o.K8sSkipSystem
	}
	if o.K8sToken != nil {
		toSerialize["k8s-token"] = o.K8sToken
	}
	if o.K8sUrl != nil {
		toSerialize["k8s-url"] = o.K8sUrl
	}
	if o.K8sUsername != nil {
		toSerialize["k8s-username"] = o.K8sUsername
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new_name"] = o.NewName
	}
	if o.ProtectionKey != nil {
		toSerialize["protection-key"] = o.ProtectionKey
	}
	if o.TargetLocation != nil {
		toSerialize["target-location"] = o.TargetLocation
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateMigration struct {
	value *GatewayUpdateMigration
	isSet bool
}

func (v NullableGatewayUpdateMigration) Get() *GatewayUpdateMigration {
	return v.value
}

func (v *NullableGatewayUpdateMigration) Set(val *GatewayUpdateMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateMigration(val *GatewayUpdateMigration) *NullableGatewayUpdateMigration {
	return &NullableGatewayUpdateMigration{value: val, isSet: true}
}

func (v NullableGatewayUpdateMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


