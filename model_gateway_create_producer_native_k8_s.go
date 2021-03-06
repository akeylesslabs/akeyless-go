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
	K8sClusterCaCert string `json:"k8s-cluster-ca-cert"`
	// K8S cluster URL endpoint
	K8sClusterEndpoint string `json:"k8s-cluster-endpoint"`
	// K8S cluster Bearer token
	K8sClusterToken string `json:"k8s-cluster-token"`
	// K8S namespace
	K8sNamespace *string `json:"k8s-namespace,omitempty"`
	// K8S service account
	K8sServiceAccount string `json:"k8s-service-account"`
	// Producer name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewGatewayCreateProducerNativeK8S instantiates a new GatewayCreateProducerNativeK8S object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerNativeK8S(k8sClusterCaCert string, k8sClusterEndpoint string, k8sClusterToken string, k8sServiceAccount string, name string, ) *GatewayCreateProducerNativeK8S {
	this := GatewayCreateProducerNativeK8S{}
	this.K8sClusterCaCert = k8sClusterCaCert
	this.K8sClusterEndpoint = k8sClusterEndpoint
	this.K8sClusterToken = k8sClusterToken
	this.K8sServiceAccount = k8sServiceAccount
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

// GetK8sClusterCaCert returns the K8sClusterCaCert field value
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterCaCert() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sClusterCaCert
}

// GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sClusterCaCert, true
}

// SetK8sClusterCaCert sets field value
func (o *GatewayCreateProducerNativeK8S) SetK8sClusterCaCert(v string) {
	o.K8sClusterCaCert = v
}

// GetK8sClusterEndpoint returns the K8sClusterEndpoint field value
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterEndpoint() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sClusterEndpoint
}

// GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sClusterEndpoint, true
}

// SetK8sClusterEndpoint sets field value
func (o *GatewayCreateProducerNativeK8S) SetK8sClusterEndpoint(v string) {
	o.K8sClusterEndpoint = v
}

// GetK8sClusterToken returns the K8sClusterToken field value
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterToken() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sClusterToken
}

// GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetK8sClusterTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sClusterToken, true
}

// SetK8sClusterToken sets field value
func (o *GatewayCreateProducerNativeK8S) SetK8sClusterToken(v string) {
	o.K8sClusterToken = v
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

// GetK8sServiceAccount returns the K8sServiceAccount field value
func (o *GatewayCreateProducerNativeK8S) GetK8sServiceAccount() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.K8sServiceAccount
}

// GetK8sServiceAccountOk returns a tuple with the K8sServiceAccount field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetK8sServiceAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.K8sServiceAccount, true
}

// SetK8sServiceAccount sets field value
func (o *GatewayCreateProducerNativeK8S) SetK8sServiceAccount(v string) {
	o.K8sServiceAccount = v
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

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayCreateProducerNativeK8S) SetPassword(v string) {
	o.Password = &v
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

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayCreateProducerNativeK8S) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerNativeK8S) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerNativeK8S) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayCreateProducerNativeK8S) SetUsername(v string) {
	o.Username = &v
}

func (o GatewayCreateProducerNativeK8S) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["k8s-cluster-ca-cert"] = o.K8sClusterCaCert
	}
	if true {
		toSerialize["k8s-cluster-endpoint"] = o.K8sClusterEndpoint
	}
	if true {
		toSerialize["k8s-cluster-token"] = o.K8sClusterToken
	}
	if o.K8sNamespace != nil {
		toSerialize["k8s-namespace"] = o.K8sNamespace
	}
	if true {
		toSerialize["k8s-service-account"] = o.K8sServiceAccount
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
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


