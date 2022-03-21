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

// GatewayCreateProducerNativeK8S gatewayCreateProducerNativeK8S is a command that creates k8s producer
type GatewayCreateProducerNativeK8S struct {
	// K8S cluster CA certificate
	K8sClusterCaCert *string `json:"k8s-cluster-ca-cert,omitempty"`
	// K8S cluster URL endpoint
	K8sClusterEndpoint *string `json:"k8s-cluster-endpoint,omitempty"`
	// K8S cluster Bearer token
	K8sClusterToken *string `json:"k8s-cluster-token,omitempty"`
	// K8S namespace
	K8sNamespace *string `json:"k8s-namespace,omitempty"`
	// K8S service account
	K8sServiceAccount *string `json:"k8s-service-account,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	SecureAccessAllowPortForwading *bool `json:"secure-access-allow-port-forwading,omitempty"`
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	SecureAccessClusterEndpoint *string `json:"secure-access-cluster-endpoint,omitempty"`
	SecureAccessDashboardUrl *string `json:"secure-access-dashboard-url,omitempty"`
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	SecureAccessWebBrowsing *bool `json:"secure-access-web-browsing,omitempty"`
	SecureAccessWebProxy *bool `json:"secure-access-web-proxy,omitempty"`
	// List of the tags attached to this secret
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayCreateProducerNativeK8S instantiates a new GatewayCreateProducerNativeK8S object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerNativeK8S(name string, ) *GatewayCreateProducerNativeK8S {
	this := GatewayCreateProducerNativeK8S{}
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerNativeK8SWithDefaults instantiates a new GatewayCreateProducerNativeK8S object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerNativeK8SWithDefaults() *GatewayCreateProducerNativeK8S {
	this := GatewayCreateProducerNativeK8S{}
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetK8sClusterCaCert returns the K8sClusterCaCert field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterCaCert() string {
	if o == nil || o.K8sClusterCaCert == nil {
		var ret string
		return ret
	}
	return *o.K8sClusterCaCert
}

// GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterCaCertOk() (*string, bool) {
	if o == nil || o.K8sClusterCaCert == nil {
		return nil, false
	}
	return o.K8sClusterCaCert, true
}

// HasK8sClusterCaCert returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasK8sClusterCaCert() bool {
	if o != nil && o.K8sClusterCaCert != nil {
		return true
	}

	return false
}

// SetK8sClusterCaCert gets a reference to the given string and assigns it to the K8sClusterCaCert field.
func (o *GatewayCreateProducerNativeK8S) SetK8sClusterCaCert(v string) {
	o.K8sClusterCaCert = &v
}

// GetK8sClusterEndpoint returns the K8sClusterEndpoint field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterEndpoint() string {
	if o == nil || o.K8sClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.K8sClusterEndpoint
}

// GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterEndpointOk() (*string, bool) {
	if o == nil || o.K8sClusterEndpoint == nil {
		return nil, false
	}
	return o.K8sClusterEndpoint, true
}

// HasK8sClusterEndpoint returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasK8sClusterEndpoint() bool {
	if o != nil && o.K8sClusterEndpoint != nil {
		return true
	}

	return false
}

// SetK8sClusterEndpoint gets a reference to the given string and assigns it to the K8sClusterEndpoint field.
func (o *GatewayCreateProducerNativeK8S) SetK8sClusterEndpoint(v string) {
	o.K8sClusterEndpoint = &v
}

// GetK8sClusterToken returns the K8sClusterToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterToken() string {
	if o == nil || o.K8sClusterToken == nil {
		var ret string
		return ret
	}
	return *o.K8sClusterToken
}

// GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterTokenOk() (*string, bool) {
	if o == nil || o.K8sClusterToken == nil {
		return nil, false
	}
	return o.K8sClusterToken, true
}

// HasK8sClusterToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasK8sClusterToken() bool {
	if o != nil && o.K8sClusterToken != nil {
		return true
	}

	return false
}

// SetK8sClusterToken gets a reference to the given string and assigns it to the K8sClusterToken field.
func (o *GatewayCreateProducerNativeK8S) SetK8sClusterToken(v string) {
	o.K8sClusterToken = &v
}

// GetK8sNamespace returns the K8sNamespace field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetK8sNamespace() string {
	if o == nil || o.K8sNamespace == nil {
		var ret string
		return ret
	}
	return *o.K8sNamespace
}

// GetK8sNamespaceOk returns a tuple with the K8sNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetK8sNamespaceOk() (*string, bool) {
	if o == nil || o.K8sNamespace == nil {
		return nil, false
	}
	return o.K8sNamespace, true
}

// HasK8sNamespace returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasK8sNamespace() bool {
	if o != nil && o.K8sNamespace != nil {
		return true
	}

	return false
}

// SetK8sNamespace gets a reference to the given string and assigns it to the K8sNamespace field.
func (o *GatewayCreateProducerNativeK8S) SetK8sNamespace(v string) {
	o.K8sNamespace = &v
}

// GetK8sServiceAccount returns the K8sServiceAccount field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetK8sServiceAccount() string {
	if o == nil || o.K8sServiceAccount == nil {
		var ret string
		return ret
	}
	return *o.K8sServiceAccount
}

// GetK8sServiceAccountOk returns a tuple with the K8sServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetK8sServiceAccountOk() (*string, bool) {
	if o == nil || o.K8sServiceAccount == nil {
		return nil, false
	}
	return o.K8sServiceAccount, true
}

// HasK8sServiceAccount returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasK8sServiceAccount() bool {
	if o != nil && o.K8sServiceAccount != nil {
		return true
	}

	return false
}

// SetK8sServiceAccount gets a reference to the given string and assigns it to the K8sServiceAccount field.
func (o *GatewayCreateProducerNativeK8S) SetK8sServiceAccount(v string) {
	o.K8sServiceAccount = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerNativeK8S) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerNativeK8S) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerNativeK8S) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessAllowPortForwading() bool {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAllowPortForwading
}

// GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessAllowPortForwadingOk() (*bool, bool) {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		return nil, false
	}
	return o.SecureAccessAllowPortForwading, true
}

// HasSecureAccessAllowPortForwading returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasSecureAccessAllowPortForwading() bool {
	if o != nil && o.SecureAccessAllowPortForwading != nil {
		return true
	}

	return false
}

// SetSecureAccessAllowPortForwading gets a reference to the given bool and assigns it to the SecureAccessAllowPortForwading field.
func (o *GatewayCreateProducerNativeK8S) SetSecureAccessAllowPortForwading(v bool) {
	o.SecureAccessAllowPortForwading = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayCreateProducerNativeK8S) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessClusterEndpoint() string {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessClusterEndpoint
}

// GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessClusterEndpointOk() (*string, bool) {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		return nil, false
	}
	return o.SecureAccessClusterEndpoint, true
}

// HasSecureAccessClusterEndpoint returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasSecureAccessClusterEndpoint() bool {
	if o != nil && o.SecureAccessClusterEndpoint != nil {
		return true
	}

	return false
}

// SetSecureAccessClusterEndpoint gets a reference to the given string and assigns it to the SecureAccessClusterEndpoint field.
func (o *GatewayCreateProducerNativeK8S) SetSecureAccessClusterEndpoint(v string) {
	o.SecureAccessClusterEndpoint = &v
}

// GetSecureAccessDashboardUrl returns the SecureAccessDashboardUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessDashboardUrl() string {
	if o == nil || o.SecureAccessDashboardUrl == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDashboardUrl
}

// GetSecureAccessDashboardUrlOk returns a tuple with the SecureAccessDashboardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessDashboardUrlOk() (*string, bool) {
	if o == nil || o.SecureAccessDashboardUrl == nil {
		return nil, false
	}
	return o.SecureAccessDashboardUrl, true
}

// HasSecureAccessDashboardUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasSecureAccessDashboardUrl() bool {
	if o != nil && o.SecureAccessDashboardUrl != nil {
		return true
	}

	return false
}

// SetSecureAccessDashboardUrl gets a reference to the given string and assigns it to the SecureAccessDashboardUrl field.
func (o *GatewayCreateProducerNativeK8S) SetSecureAccessDashboardUrl(v string) {
	o.SecureAccessDashboardUrl = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayCreateProducerNativeK8S) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayCreateProducerNativeK8S) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessWebBrowsing() bool {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebBrowsing
}

// GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessWebBrowsingOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		return nil, false
	}
	return o.SecureAccessWebBrowsing, true
}

// HasSecureAccessWebBrowsing returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasSecureAccessWebBrowsing() bool {
	if o != nil && o.SecureAccessWebBrowsing != nil {
		return true
	}

	return false
}

// SetSecureAccessWebBrowsing gets a reference to the given bool and assigns it to the SecureAccessWebBrowsing field.
func (o *GatewayCreateProducerNativeK8S) SetSecureAccessWebBrowsing(v bool) {
	o.SecureAccessWebBrowsing = &v
}

// GetSecureAccessWebProxy returns the SecureAccessWebProxy field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessWebProxy() bool {
	if o == nil || o.SecureAccessWebProxy == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebProxy
}

// GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetSecureAccessWebProxyOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebProxy == nil {
		return nil, false
	}
	return o.SecureAccessWebProxy, true
}

// HasSecureAccessWebProxy returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasSecureAccessWebProxy() bool {
	if o != nil && o.SecureAccessWebProxy != nil {
		return true
	}

	return false
}

// SetSecureAccessWebProxy gets a reference to the given bool and assigns it to the SecureAccessWebProxy field.
func (o *GatewayCreateProducerNativeK8S) SetSecureAccessWebProxy(v bool) {
	o.SecureAccessWebProxy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerNativeK8S) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerNativeK8S) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerNativeK8S) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerNativeK8S) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerNativeK8S) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerNativeK8S) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.K8sServiceAccount != nil {
		toSerialize["k8s-service-account"] = o.K8sServiceAccount
	}
	if true {
		toSerialize["name"] = o.Name
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
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerNativeK8S struct {
	value *GatewayCreateProducerNativeK8S
	isSet bool
}

func (v NullableGatewayCreateProducerNativeK8S) Get() *GatewayCreateProducerNativeK8S {
	return v.value
}

func (v *NullableGatewayCreateProducerNativeK8S) Set(val *GatewayCreateProducerNativeK8S) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerNativeK8S) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerNativeK8S) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerNativeK8S(val *GatewayCreateProducerNativeK8S) *NullableGatewayCreateProducerNativeK8S {
	return &NullableGatewayCreateProducerNativeK8S{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerNativeK8S) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerNativeK8S) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


