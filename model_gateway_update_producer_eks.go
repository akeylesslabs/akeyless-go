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

// GatewayUpdateProducerEks gatewayUpdateProducerEks is a command that updates eks producer
type GatewayUpdateProducerEks struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Access Key ID
	EksAccessKeyId *string `json:"eks-access-key-id,omitempty"`
	// IAM assume role
	EksAssumeRole *string `json:"eks-assume-role,omitempty"`
	// EKS cluster CA certificate
	EksClusterCaCert *string `json:"eks-cluster-ca-cert,omitempty"`
	// EKS cluster URL endpoint
	EksClusterEndpoint *string `json:"eks-cluster-endpoint,omitempty"`
	// EKS cluster name
	EksClusterName *string `json:"eks-cluster-name,omitempty"`
	// Region
	EksRegion *string `json:"eks-region,omitempty"`
	// Secret Access Key
	EksSecretAccessKey *string `json:"eks-secret-access-key,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
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

// NewGatewayUpdateProducerEks instantiates a new GatewayUpdateProducerEks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerEks(name string, ) *GatewayUpdateProducerEks {
	this := GatewayUpdateProducerEks{}
	var eksRegion string = "us-east-2"
	this.EksRegion = &eksRegion
	var json bool = false
	this.Json = &json
	this.Name = name
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "15m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerEksWithDefaults instantiates a new GatewayUpdateProducerEks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerEksWithDefaults() *GatewayUpdateProducerEks {
	this := GatewayUpdateProducerEks{}
	var eksRegion string = "us-east-2"
	this.EksRegion = &eksRegion
	var json bool = false
	this.Json = &json
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "15m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerEks) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetEksAccessKeyId returns the EksAccessKeyId field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetEksAccessKeyId() string {
	if o == nil || o.EksAccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.EksAccessKeyId
}

// GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetEksAccessKeyIdOk() (*string, bool) {
	if o == nil || o.EksAccessKeyId == nil {
		return nil, false
	}
	return o.EksAccessKeyId, true
}

// HasEksAccessKeyId returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasEksAccessKeyId() bool {
	if o != nil && o.EksAccessKeyId != nil {
		return true
	}

	return false
}

// SetEksAccessKeyId gets a reference to the given string and assigns it to the EksAccessKeyId field.
func (o *GatewayUpdateProducerEks) SetEksAccessKeyId(v string) {
	o.EksAccessKeyId = &v
}

// GetEksAssumeRole returns the EksAssumeRole field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetEksAssumeRole() string {
	if o == nil || o.EksAssumeRole == nil {
		var ret string
		return ret
	}
	return *o.EksAssumeRole
}

// GetEksAssumeRoleOk returns a tuple with the EksAssumeRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetEksAssumeRoleOk() (*string, bool) {
	if o == nil || o.EksAssumeRole == nil {
		return nil, false
	}
	return o.EksAssumeRole, true
}

// HasEksAssumeRole returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasEksAssumeRole() bool {
	if o != nil && o.EksAssumeRole != nil {
		return true
	}

	return false
}

// SetEksAssumeRole gets a reference to the given string and assigns it to the EksAssumeRole field.
func (o *GatewayUpdateProducerEks) SetEksAssumeRole(v string) {
	o.EksAssumeRole = &v
}

// GetEksClusterCaCert returns the EksClusterCaCert field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetEksClusterCaCert() string {
	if o == nil || o.EksClusterCaCert == nil {
		var ret string
		return ret
	}
	return *o.EksClusterCaCert
}

// GetEksClusterCaCertOk returns a tuple with the EksClusterCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetEksClusterCaCertOk() (*string, bool) {
	if o == nil || o.EksClusterCaCert == nil {
		return nil, false
	}
	return o.EksClusterCaCert, true
}

// HasEksClusterCaCert returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasEksClusterCaCert() bool {
	if o != nil && o.EksClusterCaCert != nil {
		return true
	}

	return false
}

// SetEksClusterCaCert gets a reference to the given string and assigns it to the EksClusterCaCert field.
func (o *GatewayUpdateProducerEks) SetEksClusterCaCert(v string) {
	o.EksClusterCaCert = &v
}

// GetEksClusterEndpoint returns the EksClusterEndpoint field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetEksClusterEndpoint() string {
	if o == nil || o.EksClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.EksClusterEndpoint
}

// GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetEksClusterEndpointOk() (*string, bool) {
	if o == nil || o.EksClusterEndpoint == nil {
		return nil, false
	}
	return o.EksClusterEndpoint, true
}

// HasEksClusterEndpoint returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasEksClusterEndpoint() bool {
	if o != nil && o.EksClusterEndpoint != nil {
		return true
	}

	return false
}

// SetEksClusterEndpoint gets a reference to the given string and assigns it to the EksClusterEndpoint field.
func (o *GatewayUpdateProducerEks) SetEksClusterEndpoint(v string) {
	o.EksClusterEndpoint = &v
}

// GetEksClusterName returns the EksClusterName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetEksClusterName() string {
	if o == nil || o.EksClusterName == nil {
		var ret string
		return ret
	}
	return *o.EksClusterName
}

// GetEksClusterNameOk returns a tuple with the EksClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetEksClusterNameOk() (*string, bool) {
	if o == nil || o.EksClusterName == nil {
		return nil, false
	}
	return o.EksClusterName, true
}

// HasEksClusterName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasEksClusterName() bool {
	if o != nil && o.EksClusterName != nil {
		return true
	}

	return false
}

// SetEksClusterName gets a reference to the given string and assigns it to the EksClusterName field.
func (o *GatewayUpdateProducerEks) SetEksClusterName(v string) {
	o.EksClusterName = &v
}

// GetEksRegion returns the EksRegion field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetEksRegion() string {
	if o == nil || o.EksRegion == nil {
		var ret string
		return ret
	}
	return *o.EksRegion
}

// GetEksRegionOk returns a tuple with the EksRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetEksRegionOk() (*string, bool) {
	if o == nil || o.EksRegion == nil {
		return nil, false
	}
	return o.EksRegion, true
}

// HasEksRegion returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasEksRegion() bool {
	if o != nil && o.EksRegion != nil {
		return true
	}

	return false
}

// SetEksRegion gets a reference to the given string and assigns it to the EksRegion field.
func (o *GatewayUpdateProducerEks) SetEksRegion(v string) {
	o.EksRegion = &v
}

// GetEksSecretAccessKey returns the EksSecretAccessKey field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetEksSecretAccessKey() string {
	if o == nil || o.EksSecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.EksSecretAccessKey
}

// GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetEksSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.EksSecretAccessKey == nil {
		return nil, false
	}
	return o.EksSecretAccessKey, true
}

// HasEksSecretAccessKey returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasEksSecretAccessKey() bool {
	if o != nil && o.EksSecretAccessKey != nil {
		return true
	}

	return false
}

// SetEksSecretAccessKey gets a reference to the given string and assigns it to the EksSecretAccessKey field.
func (o *GatewayUpdateProducerEks) SetEksSecretAccessKey(v string) {
	o.EksSecretAccessKey = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerEks) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerEks) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerEks) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerEks) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerEks) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetSecureAccessAllowPortForwading() bool {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAllowPortForwading
}

// GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetSecureAccessAllowPortForwadingOk() (*bool, bool) {
	if o == nil || o.SecureAccessAllowPortForwading == nil {
		return nil, false
	}
	return o.SecureAccessAllowPortForwading, true
}

// HasSecureAccessAllowPortForwading returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasSecureAccessAllowPortForwading() bool {
	if o != nil && o.SecureAccessAllowPortForwading != nil {
		return true
	}

	return false
}

// SetSecureAccessAllowPortForwading gets a reference to the given bool and assigns it to the SecureAccessAllowPortForwading field.
func (o *GatewayUpdateProducerEks) SetSecureAccessAllowPortForwading(v bool) {
	o.SecureAccessAllowPortForwading = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayUpdateProducerEks) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetSecureAccessClusterEndpoint() string {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessClusterEndpoint
}

// GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetSecureAccessClusterEndpointOk() (*string, bool) {
	if o == nil || o.SecureAccessClusterEndpoint == nil {
		return nil, false
	}
	return o.SecureAccessClusterEndpoint, true
}

// HasSecureAccessClusterEndpoint returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasSecureAccessClusterEndpoint() bool {
	if o != nil && o.SecureAccessClusterEndpoint != nil {
		return true
	}

	return false
}

// SetSecureAccessClusterEndpoint gets a reference to the given string and assigns it to the SecureAccessClusterEndpoint field.
func (o *GatewayUpdateProducerEks) SetSecureAccessClusterEndpoint(v string) {
	o.SecureAccessClusterEndpoint = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayUpdateProducerEks) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayUpdateProducerEks) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerEks) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerEks) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerEks) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerEks) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerEks) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerEks) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerEks) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerEks) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerEks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.EksAccessKeyId != nil {
		toSerialize["eks-access-key-id"] = o.EksAccessKeyId
	}
	if o.EksAssumeRole != nil {
		toSerialize["eks-assume-role"] = o.EksAssumeRole
	}
	if o.EksClusterCaCert != nil {
		toSerialize["eks-cluster-ca-cert"] = o.EksClusterCaCert
	}
	if o.EksClusterEndpoint != nil {
		toSerialize["eks-cluster-endpoint"] = o.EksClusterEndpoint
	}
	if o.EksClusterName != nil {
		toSerialize["eks-cluster-name"] = o.EksClusterName
	}
	if o.EksRegion != nil {
		toSerialize["eks-region"] = o.EksRegion
	}
	if o.EksSecretAccessKey != nil {
		toSerialize["eks-secret-access-key"] = o.EksSecretAccessKey
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
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

type NullableGatewayUpdateProducerEks struct {
	value *GatewayUpdateProducerEks
	isSet bool
}

func (v NullableGatewayUpdateProducerEks) Get() *GatewayUpdateProducerEks {
	return v.value
}

func (v *NullableGatewayUpdateProducerEks) Set(val *GatewayUpdateProducerEks) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerEks) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerEks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerEks(val *GatewayUpdateProducerEks) *NullableGatewayUpdateProducerEks {
	return &NullableGatewayUpdateProducerEks{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerEks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerEks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


