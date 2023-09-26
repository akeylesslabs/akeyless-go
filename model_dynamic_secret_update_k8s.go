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

// DynamicSecretUpdateK8s dynamicSecretUpdateK8s is a command that updates k8s dynamic secret
type DynamicSecretUpdateK8s struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Comma-separated list of allowed K8S namespaces for the generated ServiceAccount (relevant only for k8s-service-account-type=dynamic)
	K8sAllowedNamespaces *string `json:"k8s-allowed-namespaces,omitempty"`
	// K8S cluster CA certificate
	K8sClusterCaCert *string `json:"k8s-cluster-ca-cert,omitempty"`
	// K8S cluster URL endpoint
	K8sClusterEndpoint *string `json:"k8s-cluster-endpoint,omitempty"`
	// K8S cluster Bearer token
	K8sClusterToken *string `json:"k8s-cluster-token,omitempty"`
	// K8S Namespace where the ServiceAccount exists.
	K8sNamespace *string `json:"k8s-namespace,omitempty"`
	// The pre-existing Role or ClusterRole name to bind the generated ServiceAccount to (relevant only for k8s-service-account-type=dynamic)
	K8sPredefinedRoleName *string `json:"k8s-predefined-role-name,omitempty"`
	// Specifies the type of the pre-existing K8S role [Role, ClusterRole] (relevant only for k8s-service-account-type=dynamic)
	K8sPredefinedRoleType *string `json:"k8s-predefined-role-type,omitempty"`
	// Path to yaml file that contains definitions of K8S role and role binding (relevant only for k8s-service-account-type=dynamic)
	K8sRolebindingYamlDef *string `json:"k8s-rolebinding-yaml-def,omitempty"`
	// K8S ServiceAccount to extract token from.
	K8sServiceAccount *string `json:"k8s-service-account,omitempty"`
	// K8S ServiceAccount type [fixed, dynamic].
	K8sServiceAccountType *string `json:"k8s-service-account-type,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret new name
	NewName *string `json:"new-name,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Enable Port forwarding while using CLI access
	SecureAccessAllowPortForwading *bool `json:"secure-access-allow-port-forwading,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// The K8s cluster endpoint URL
	SecureAccessClusterEndpoint *string `json:"secure-access-cluster-endpoint,omitempty"`
	// The K8s dashboard url
	SecureAccessDashboardUrl *string `json:"secure-access-dashboard-url,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// Secure browser via Akeyless Web Access Bastion
	SecureAccessWebBrowsing *bool `json:"secure-access-web-browsing,omitempty"`
	// Web-Proxy via Akeyless Web Access Bastion
	SecureAccessWebProxy *bool `json:"secure-access-web-proxy,omitempty"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Use the GW's service account
	UseGwServiceAccount *bool `json:"use-gw-service-account,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewDynamicSecretUpdateK8s instantiates a new DynamicSecretUpdateK8s object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdateK8s(name string, ) *DynamicSecretUpdateK8s {
	this := DynamicSecretUpdateK8s{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var secureAccessWebBrowsing bool = false
	this.SecureAccessWebBrowsing = &secureAccessWebBrowsing
	var secureAccessWebProxy bool = false
	this.SecureAccessWebProxy = &secureAccessWebProxy
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretUpdateK8sWithDefaults instantiates a new DynamicSecretUpdateK8s object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdateK8sWithDefaults() *DynamicSecretUpdateK8s {
	this := DynamicSecretUpdateK8s{}
	var json bool = false
	this.Json = &json
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var secureAccessWebBrowsing bool = false
	this.SecureAccessWebBrowsing = &secureAccessWebBrowsing
	var secureAccessWebProxy bool = false
	this.SecureAccessWebProxy = &secureAccessWebProxy
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretUpdateK8s) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretUpdateK8s) SetJson(v bool) {
	o.Json = &v
}

// GetK8sAllowedNamespaces returns the K8sAllowedNamespaces field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sAllowedNamespaces() string {
	if o == nil || o.K8sAllowedNamespaces == nil {
		var ret string
		return ret
	}
	return *o.K8sAllowedNamespaces
}

// GetK8sAllowedNamespacesOk returns a tuple with the K8sAllowedNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sAllowedNamespacesOk() (*string, bool) {
	if o == nil || o.K8sAllowedNamespaces == nil {
		return nil, false
	}
	return o.K8sAllowedNamespaces, true
}

// HasK8sAllowedNamespaces returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sAllowedNamespaces() bool {
	if o != nil && o.K8sAllowedNamespaces != nil {
		return true
	}

	return false
}

// SetK8sAllowedNamespaces gets a reference to the given string and assigns it to the K8sAllowedNamespaces field.
func (o *DynamicSecretUpdateK8s) SetK8sAllowedNamespaces(v string) {
	o.K8sAllowedNamespaces = &v
}

// GetK8sClusterCaCert returns the K8sClusterCaCert field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sClusterCaCert() string {
	if o == nil || o.K8sClusterCaCert == nil {
		var ret string
		return ret
	}
	return *o.K8sClusterCaCert
}

// GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sClusterCaCertOk() (*string, bool) {
	if o == nil || o.K8sClusterCaCert == nil {
		return nil, false
	}
	return o.K8sClusterCaCert, true
}

// HasK8sClusterCaCert returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sClusterCaCert() bool {
	if o != nil && o.K8sClusterCaCert != nil {
		return true
	}

	return false
}

// SetK8sClusterCaCert gets a reference to the given string and assigns it to the K8sClusterCaCert field.
func (o *DynamicSecretUpdateK8s) SetK8sClusterCaCert(v string) {
	o.K8sClusterCaCert = &v
}

// GetK8sClusterEndpoint returns the K8sClusterEndpoint field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sClusterEndpoint() string {
	if o == nil || o.K8sClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.K8sClusterEndpoint
}

// GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sClusterEndpointOk() (*string, bool) {
	if o == nil || o.K8sClusterEndpoint == nil {
		return nil, false
	}
	return o.K8sClusterEndpoint, true
}

// HasK8sClusterEndpoint returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sClusterEndpoint() bool {
	if o != nil && o.K8sClusterEndpoint != nil {
		return true
	}

	return false
}

// SetK8sClusterEndpoint gets a reference to the given string and assigns it to the K8sClusterEndpoint field.
func (o *DynamicSecretUpdateK8s) SetK8sClusterEndpoint(v string) {
	o.K8sClusterEndpoint = &v
}

// GetK8sClusterToken returns the K8sClusterToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sClusterToken() string {
	if o == nil || o.K8sClusterToken == nil {
		var ret string
		return ret
	}
	return *o.K8sClusterToken
}

// GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sClusterTokenOk() (*string, bool) {
	if o == nil || o.K8sClusterToken == nil {
		return nil, false
	}
	return o.K8sClusterToken, true
}

// HasK8sClusterToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sClusterToken() bool {
	if o != nil && o.K8sClusterToken != nil {
		return true
	}

	return false
}

// SetK8sClusterToken gets a reference to the given string and assigns it to the K8sClusterToken field.
func (o *DynamicSecretUpdateK8s) SetK8sClusterToken(v string) {
	o.K8sClusterToken = &v
}

// GetK8sNamespace returns the K8sNamespace field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sNamespace() string {
	if o == nil || o.K8sNamespace == nil {
		var ret string
		return ret
	}
	return *o.K8sNamespace
}

// GetK8sNamespaceOk returns a tuple with the K8sNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sNamespaceOk() (*string, bool) {
	if o == nil || o.K8sNamespace == nil {
		return nil, false
	}
	return o.K8sNamespace, true
}

// HasK8sNamespace returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sNamespace() bool {
	if o != nil && o.K8sNamespace != nil {
		return true
	}

	return false
}

// SetK8sNamespace gets a reference to the given string and assigns it to the K8sNamespace field.
func (o *DynamicSecretUpdateK8s) SetK8sNamespace(v string) {
	o.K8sNamespace = &v
}

// GetK8sPredefinedRoleName returns the K8sPredefinedRoleName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sPredefinedRoleName() string {
	if o == nil || o.K8sPredefinedRoleName == nil {
		var ret string
		return ret
	}
	return *o.K8sPredefinedRoleName
}

// GetK8sPredefinedRoleNameOk returns a tuple with the K8sPredefinedRoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sPredefinedRoleNameOk() (*string, bool) {
	if o == nil || o.K8sPredefinedRoleName == nil {
		return nil, false
	}
	return o.K8sPredefinedRoleName, true
}

// HasK8sPredefinedRoleName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sPredefinedRoleName() bool {
	if o != nil && o.K8sPredefinedRoleName != nil {
		return true
	}

	return false
}

// SetK8sPredefinedRoleName gets a reference to the given string and assigns it to the K8sPredefinedRoleName field.
func (o *DynamicSecretUpdateK8s) SetK8sPredefinedRoleName(v string) {
	o.K8sPredefinedRoleName = &v
}

// GetK8sPredefinedRoleType returns the K8sPredefinedRoleType field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sPredefinedRoleType() string {
	if o == nil || o.K8sPredefinedRoleType == nil {
		var ret string
		return ret
	}
	return *o.K8sPredefinedRoleType
}

// GetK8sPredefinedRoleTypeOk returns a tuple with the K8sPredefinedRoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sPredefinedRoleTypeOk() (*string, bool) {
	if o == nil || o.K8sPredefinedRoleType == nil {
		return nil, false
	}
	return o.K8sPredefinedRoleType, true
}

// HasK8sPredefinedRoleType returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sPredefinedRoleType() bool {
	if o != nil && o.K8sPredefinedRoleType != nil {
		return true
	}

	return false
}

// SetK8sPredefinedRoleType gets a reference to the given string and assigns it to the K8sPredefinedRoleType field.
func (o *DynamicSecretUpdateK8s) SetK8sPredefinedRoleType(v string) {
	o.K8sPredefinedRoleType = &v
}

// GetK8sRolebindingYamlDef returns the K8sRolebindingYamlDef field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sRolebindingYamlDef() string {
	if o == nil || o.K8sRolebindingYamlDef == nil {
		var ret string
		return ret
	}
	return *o.K8sRolebindingYamlDef
}

// GetK8sRolebindingYamlDefOk returns a tuple with the K8sRolebindingYamlDef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sRolebindingYamlDefOk() (*string, bool) {
	if o == nil || o.K8sRolebindingYamlDef == nil {
		return nil, false
	}
	return o.K8sRolebindingYamlDef, true
}

// HasK8sRolebindingYamlDef returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sRolebindingYamlDef() bool {
	if o != nil && o.K8sRolebindingYamlDef != nil {
		return true
	}

	return false
}

// SetK8sRolebindingYamlDef gets a reference to the given string and assigns it to the K8sRolebindingYamlDef field.
func (o *DynamicSecretUpdateK8s) SetK8sRolebindingYamlDef(v string) {
	o.K8sRolebindingYamlDef = &v
}

// GetK8sServiceAccount returns the K8sServiceAccount field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sServiceAccount() string {
	if o == nil || o.K8sServiceAccount == nil {
		var ret string
		return ret
	}
	return *o.K8sServiceAccount
}

// GetK8sServiceAccountOk returns a tuple with the K8sServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sServiceAccountOk() (*string, bool) {
	if o == nil || o.K8sServiceAccount == nil {
		return nil, false
	}
	return o.K8sServiceAccount, true
}

// HasK8sServiceAccount returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sServiceAccount() bool {
	if o != nil && o.K8sServiceAccount != nil {
		return true
	}

	return false
}

// SetK8sServiceAccount gets a reference to the given string and assigns it to the K8sServiceAccount field.
func (o *DynamicSecretUpdateK8s) SetK8sServiceAccount(v string) {
	o.K8sServiceAccount = &v
}

// GetK8sServiceAccountType returns the K8sServiceAccountType field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetK8sServiceAccountType() string {
	if o == nil || o.K8sServiceAccountType == nil {
		var ret string
		return ret
	}
	return *o.K8sServiceAccountType
}

// GetK8sServiceAccountTypeOk returns a tuple with the K8sServiceAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetK8sServiceAccountTypeOk() (*string, bool) {
	if o == nil || o.K8sServiceAccountType == nil {
		return nil, false
	}
	return o.K8sServiceAccountType, true
}

// HasK8sServiceAccountType returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasK8sServiceAccountType() bool {
	if o != nil && o.K8sServiceAccountType != nil {
		return true
	}

	return false
}

// SetK8sServiceAccountType gets a reference to the given string and assigns it to the K8sServiceAccountType field.
func (o *DynamicSecretUpdateK8s) SetK8sServiceAccountType(v string) {
	o.K8sServiceAccountType = &v
}

// GetName returns the Name field value
func (o *DynamicSecretUpdateK8s) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretUpdateK8s) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DynamicSecretUpdateK8s) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretUpdateK8s) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetSecureAccessAllowPortForwading() bool {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAllowPortForwading
}

// GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetSecureAccessAllowPortForwadingOk() (*bool, bool) {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		return nil, false
	}
	return o.SecureAccessAllowPortForwading, true
}

// HasSecureAccessAllowPortForwading returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasSecureAccessAllowPortForwading() bool {
	if o != nil && o.SecureAccessAllowPortForwading != nil {
		return true
	}

	return false
}

// SetSecureAccessAllowPortForwading gets a reference to the given bool and assigns it to the SecureAccessAllowPortForwading field.
func (o *DynamicSecretUpdateK8s) SetSecureAccessAllowPortForwading(v bool) {
	o.SecureAccessAllowPortForwading = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *DynamicSecretUpdateK8s) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetSecureAccessClusterEndpoint() string {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessClusterEndpoint
}

// GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetSecureAccessClusterEndpointOk() (*string, bool) {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		return nil, false
	}
	return o.SecureAccessClusterEndpoint, true
}

// HasSecureAccessClusterEndpoint returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasSecureAccessClusterEndpoint() bool {
	if o != nil && o.SecureAccessClusterEndpoint != nil {
		return true
	}

	return false
}

// SetSecureAccessClusterEndpoint gets a reference to the given string and assigns it to the SecureAccessClusterEndpoint field.
func (o *DynamicSecretUpdateK8s) SetSecureAccessClusterEndpoint(v string) {
	o.SecureAccessClusterEndpoint = &v
}

// GetSecureAccessDashboardUrl returns the SecureAccessDashboardUrl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetSecureAccessDashboardUrl() string {
	if o == nil || o.SecureAccessDashboardUrl == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDashboardUrl
}

// GetSecureAccessDashboardUrlOk returns a tuple with the SecureAccessDashboardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetSecureAccessDashboardUrlOk() (*string, bool) {
	if o == nil || o.SecureAccessDashboardUrl == nil {
		return nil, false
	}
	return o.SecureAccessDashboardUrl, true
}

// HasSecureAccessDashboardUrl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasSecureAccessDashboardUrl() bool {
	if o != nil && o.SecureAccessDashboardUrl != nil {
		return true
	}

	return false
}

// SetSecureAccessDashboardUrl gets a reference to the given string and assigns it to the SecureAccessDashboardUrl field.
func (o *DynamicSecretUpdateK8s) SetSecureAccessDashboardUrl(v string) {
	o.SecureAccessDashboardUrl = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DynamicSecretUpdateK8s) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *DynamicSecretUpdateK8s) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetSecureAccessWebBrowsing() bool {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebBrowsing
}

// GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetSecureAccessWebBrowsingOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		return nil, false
	}
	return o.SecureAccessWebBrowsing, true
}

// HasSecureAccessWebBrowsing returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasSecureAccessWebBrowsing() bool {
	if o != nil && o.SecureAccessWebBrowsing != nil {
		return true
	}

	return false
}

// SetSecureAccessWebBrowsing gets a reference to the given bool and assigns it to the SecureAccessWebBrowsing field.
func (o *DynamicSecretUpdateK8s) SetSecureAccessWebBrowsing(v bool) {
	o.SecureAccessWebBrowsing = &v
}

// GetSecureAccessWebProxy returns the SecureAccessWebProxy field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetSecureAccessWebProxy() bool {
	if o == nil || o.SecureAccessWebProxy == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebProxy
}

// GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetSecureAccessWebProxyOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebProxy == nil {
		return nil, false
	}
	return o.SecureAccessWebProxy, true
}

// HasSecureAccessWebProxy returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasSecureAccessWebProxy() bool {
	if o != nil && o.SecureAccessWebProxy != nil {
		return true
	}

	return false
}

// SetSecureAccessWebProxy gets a reference to the given bool and assigns it to the SecureAccessWebProxy field.
func (o *DynamicSecretUpdateK8s) SetSecureAccessWebProxy(v bool) {
	o.SecureAccessWebProxy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretUpdateK8s) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretUpdateK8s) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretUpdateK8s) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretUpdateK8s) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUseGwServiceAccount returns the UseGwServiceAccount field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetUseGwServiceAccount() bool {
	if o == nil || o.UseGwServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.UseGwServiceAccount
}

// GetUseGwServiceAccountOk returns a tuple with the UseGwServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetUseGwServiceAccountOk() (*bool, bool) {
	if o == nil || o.UseGwServiceAccount == nil {
		return nil, false
	}
	return o.UseGwServiceAccount, true
}

// HasUseGwServiceAccount returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasUseGwServiceAccount() bool {
	if o != nil && o.UseGwServiceAccount != nil {
		return true
	}

	return false
}

// SetUseGwServiceAccount gets a reference to the given bool and assigns it to the UseGwServiceAccount field.
func (o *DynamicSecretUpdateK8s) SetUseGwServiceAccount(v bool) {
	o.UseGwServiceAccount = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateK8s) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateK8s) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateK8s) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretUpdateK8s) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretUpdateK8s) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.K8sAllowedNamespaces != nil {
		toSerialize["k8s-allowed-namespaces"] = o.K8sAllowedNamespaces
	}
	if o.K8sClusterCaCert != nil {
		toSerialize["k8s-cluster-ca-cert"] = o.K8sClusterCaCert
	}
	if o.K8sClusterEndpoint != nil {
		toSerialize["k8s-cluster-endpoint"] = o.K8sClusterEndpoint
	}
	if o.K8sClusterToken != nil {
		toSerialize["k8s-cluster-token"] = o.K8sClusterToken
	}
	if o.K8sNamespace != nil {
		toSerialize["k8s-namespace"] = o.K8sNamespace
	}
	if o.K8sPredefinedRoleName != nil {
		toSerialize["k8s-predefined-role-name"] = o.K8sPredefinedRoleName
	}
	if o.K8sPredefinedRoleType != nil {
		toSerialize["k8s-predefined-role-type"] = o.K8sPredefinedRoleType
	}
	if o.K8sRolebindingYamlDef != nil {
		toSerialize["k8s-rolebinding-yaml-def"] = o.K8sRolebindingYamlDef
	}
	if o.K8sServiceAccount != nil {
		toSerialize["k8s-service-account"] = o.K8sServiceAccount
	}
	if o.K8sServiceAccountType != nil {
		toSerialize["k8s-service-account-type"] = o.K8sServiceAccountType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.SecureAccessAllowPortForwading != nil {
		toSerialize["secure-access-allow-port-forwading"] = o.SecureAccessAllowPortForwading
	}
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessClusterEndpoint != nil {
		toSerialize["secure-access-cluster-endpoint"] = o.SecureAccessClusterEndpoint
	}
	if o.SecureAccessDashboardUrl != nil {
		toSerialize["secure-access-dashboard-url"] = o.SecureAccessDashboardUrl
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessWeb != nil {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
	}
	if o.SecureAccessWebBrowsing != nil {
		toSerialize["secure-access-web-browsing"] = o.SecureAccessWebBrowsing
	}
	if o.SecureAccessWebProxy != nil {
		toSerialize["secure-access-web-proxy"] = o.SecureAccessWebProxy
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetName != nil {
		toSerialize["target-name"] = o.TargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UseGwServiceAccount != nil {
		toSerialize["use-gw-service-account"] = o.UseGwServiceAccount
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretUpdateK8s struct {
	value *DynamicSecretUpdateK8s
	isSet bool
}

func (v NullableDynamicSecretUpdateK8s) Get() *DynamicSecretUpdateK8s {
	return v.value
}

func (v *NullableDynamicSecretUpdateK8s) Set(val *DynamicSecretUpdateK8s) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdateK8s) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdateK8s) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdateK8s(val *DynamicSecretUpdateK8s) *NullableDynamicSecretUpdateK8s {
	return &NullableDynamicSecretUpdateK8s{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdateK8s) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdateK8s) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


