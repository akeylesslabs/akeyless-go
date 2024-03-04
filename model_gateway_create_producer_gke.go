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

// GatewayCreateProducerGke gatewayCreateProducerGke is a command that creates gke producer [Deprecated: Use dynamic-secret-create-gke command]
type GatewayCreateProducerGke struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// GKE Service Account key file path
	GkeAccountKey *string `json:"gke-account-key,omitempty"`
	// GKE cluster CA certificate
	GkeClusterCert *string `json:"gke-cluster-cert,omitempty"`
	// GKE cluster URL endpoint
	GkeClusterEndpoint *string `json:"gke-cluster-endpoint,omitempty"`
	// GKE cluster name
	GkeClusterName *string `json:"gke-cluster-name,omitempty"`
	// GKE service account email
	GkeServiceAccountEmail *string `json:"gke-service-account-email,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Enable Port forwarding while using CLI access
	SecureAccessAllowPortForwading *bool `json:"secure-access-allow-port-forwading,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// The K8s cluster endpoint URL
	SecureAccessClusterEndpoint *string `json:"secure-access-cluster-endpoint,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// Add tags attached to this object
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

// NewGatewayCreateProducerGke instantiates a new GatewayCreateProducerGke object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerGke(name string, ) *GatewayCreateProducerGke {
	this := GatewayCreateProducerGke{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerGkeWithDefaults instantiates a new GatewayCreateProducerGke object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerGkeWithDefaults() *GatewayCreateProducerGke {
	this := GatewayCreateProducerGke{}
	var json bool = false
	this.Json = &json
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerGke) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetGkeAccountKey returns the GkeAccountKey field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetGkeAccountKey() string {
	if o == nil || o.GkeAccountKey == nil {
		var ret string
		return ret
	}
	return *o.GkeAccountKey
}

// GetGkeAccountKeyOk returns a tuple with the GkeAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetGkeAccountKeyOk() (*string, bool) {
	if o == nil || o.GkeAccountKey == nil {
		return nil, false
	}
	return o.GkeAccountKey, true
}

// HasGkeAccountKey returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasGkeAccountKey() bool {
	if o != nil && o.GkeAccountKey != nil {
		return true
	}

	return false
}

// SetGkeAccountKey gets a reference to the given string and assigns it to the GkeAccountKey field.
func (o *GatewayCreateProducerGke) SetGkeAccountKey(v string) {
	o.GkeAccountKey = &v
}

// GetGkeClusterCert returns the GkeClusterCert field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetGkeClusterCert() string {
	if o == nil || o.GkeClusterCert == nil {
		var ret string
		return ret
	}
	return *o.GkeClusterCert
}

// GetGkeClusterCertOk returns a tuple with the GkeClusterCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetGkeClusterCertOk() (*string, bool) {
	if o == nil || o.GkeClusterCert == nil {
		return nil, false
	}
	return o.GkeClusterCert, true
}

// HasGkeClusterCert returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasGkeClusterCert() bool {
	if o != nil && o.GkeClusterCert != nil {
		return true
	}

	return false
}

// SetGkeClusterCert gets a reference to the given string and assigns it to the GkeClusterCert field.
func (o *GatewayCreateProducerGke) SetGkeClusterCert(v string) {
	o.GkeClusterCert = &v
}

// GetGkeClusterEndpoint returns the GkeClusterEndpoint field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetGkeClusterEndpoint() string {
	if o == nil || o.GkeClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.GkeClusterEndpoint
}

// GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetGkeClusterEndpointOk() (*string, bool) {
	if o == nil || o.GkeClusterEndpoint == nil {
		return nil, false
	}
	return o.GkeClusterEndpoint, true
}

// HasGkeClusterEndpoint returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasGkeClusterEndpoint() bool {
	if o != nil && o.GkeClusterEndpoint != nil {
		return true
	}

	return false
}

// SetGkeClusterEndpoint gets a reference to the given string and assigns it to the GkeClusterEndpoint field.
func (o *GatewayCreateProducerGke) SetGkeClusterEndpoint(v string) {
	o.GkeClusterEndpoint = &v
}

// GetGkeClusterName returns the GkeClusterName field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetGkeClusterName() string {
	if o == nil || o.GkeClusterName == nil {
		var ret string
		return ret
	}
	return *o.GkeClusterName
}

// GetGkeClusterNameOk returns a tuple with the GkeClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetGkeClusterNameOk() (*string, bool) {
	if o == nil || o.GkeClusterName == nil {
		return nil, false
	}
	return o.GkeClusterName, true
}

// HasGkeClusterName returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasGkeClusterName() bool {
	if o != nil && o.GkeClusterName != nil {
		return true
	}

	return false
}

// SetGkeClusterName gets a reference to the given string and assigns it to the GkeClusterName field.
func (o *GatewayCreateProducerGke) SetGkeClusterName(v string) {
	o.GkeClusterName = &v
}

// GetGkeServiceAccountEmail returns the GkeServiceAccountEmail field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetGkeServiceAccountEmail() string {
	if o == nil || o.GkeServiceAccountEmail == nil {
		var ret string
		return ret
	}
	return *o.GkeServiceAccountEmail
}

// GetGkeServiceAccountEmailOk returns a tuple with the GkeServiceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetGkeServiceAccountEmailOk() (*string, bool) {
	if o == nil || o.GkeServiceAccountEmail == nil {
		return nil, false
	}
	return o.GkeServiceAccountEmail, true
}

// HasGkeServiceAccountEmail returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasGkeServiceAccountEmail() bool {
	if o != nil && o.GkeServiceAccountEmail != nil {
		return true
	}

	return false
}

// SetGkeServiceAccountEmail gets a reference to the given string and assigns it to the GkeServiceAccountEmail field.
func (o *GatewayCreateProducerGke) SetGkeServiceAccountEmail(v string) {
	o.GkeServiceAccountEmail = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerGke) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerGke) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerGke) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerGke) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetSecureAccessAllowPortForwading() bool {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAllowPortForwading
}

// GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetSecureAccessAllowPortForwadingOk() (*bool, bool) {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		return nil, false
	}
	return o.SecureAccessAllowPortForwading, true
}

// HasSecureAccessAllowPortForwading returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasSecureAccessAllowPortForwading() bool {
	if o != nil && o.SecureAccessAllowPortForwading != nil {
		return true
	}

	return false
}

// SetSecureAccessAllowPortForwading gets a reference to the given bool and assigns it to the SecureAccessAllowPortForwading field.
func (o *GatewayCreateProducerGke) SetSecureAccessAllowPortForwading(v bool) {
	o.SecureAccessAllowPortForwading = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayCreateProducerGke) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetSecureAccessClusterEndpoint() string {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessClusterEndpoint
}

// GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetSecureAccessClusterEndpointOk() (*string, bool) {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		return nil, false
	}
	return o.SecureAccessClusterEndpoint, true
}

// HasSecureAccessClusterEndpoint returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasSecureAccessClusterEndpoint() bool {
	if o != nil && o.SecureAccessClusterEndpoint != nil {
		return true
	}

	return false
}

// SetSecureAccessClusterEndpoint gets a reference to the given string and assigns it to the SecureAccessClusterEndpoint field.
func (o *GatewayCreateProducerGke) SetSecureAccessClusterEndpoint(v string) {
	o.SecureAccessClusterEndpoint = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayCreateProducerGke) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayCreateProducerGke) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerGke) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerGke) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerGke) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerGke) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerGke) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGke) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerGke) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerGke) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerGke) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.GkeAccountKey != nil {
		toSerialize["gke-account-key"] = o.GkeAccountKey
	}
	if o.GkeClusterCert != nil {
		toSerialize["gke-cluster-cert"] = o.GkeClusterCert
	}
	if o.GkeClusterEndpoint != nil {
		toSerialize["gke-cluster-endpoint"] = o.GkeClusterEndpoint
	}
	if o.GkeClusterName != nil {
		toSerialize["gke-cluster-name"] = o.GkeClusterName
	}
	if o.GkeServiceAccountEmail != nil {
		toSerialize["gke-service-account-email"] = o.GkeServiceAccountEmail
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
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
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessWeb != nil {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
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

type NullableGatewayCreateProducerGke struct {
	value *GatewayCreateProducerGke
	isSet bool
}

func (v NullableGatewayCreateProducerGke) Get() *GatewayCreateProducerGke {
	return v.value
}

func (v *NullableGatewayCreateProducerGke) Set(val *GatewayCreateProducerGke) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerGke) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerGke) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerGke(val *GatewayCreateProducerGke) *NullableGatewayCreateProducerGke {
	return &NullableGatewayCreateProducerGke{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerGke) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerGke) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


