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

// GatewayCreateProducerGcp gatewayCreateProducerGcp is a command that creates a GCP producer
type GatewayCreateProducerGcp struct {
	GcpCredType *string `json:"gcp-cred-type,omitempty"`
	// Base64-encoded service account private key text
	GcpKey *string `json:"gcp-key,omitempty"`
	// Service account key algorithm, e.g. KEY_ALG_RSA_1024
	GcpKeyAlgo *string `json:"gcp-key-algo,omitempty"`
	// GCP service account email
	GcpSaEmail *string `json:"gcp-sa-email,omitempty"`
	// Access token scopes list, e.g. scope1,scope2
	GcpTokenScopes *string `json:"gcp-token-scopes,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
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

// NewGatewayCreateProducerGcp instantiates a new GatewayCreateProducerGcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerGcp(name string, ) *GatewayCreateProducerGcp {
	this := GatewayCreateProducerGcp{}
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerGcpWithDefaults instantiates a new GatewayCreateProducerGcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerGcpWithDefaults() *GatewayCreateProducerGcp {
	this := GatewayCreateProducerGcp{}
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetGcpCredType returns the GcpCredType field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetGcpCredType() string {
	if o == nil || o.GcpCredType == nil {
		var ret string
		return ret
	}
	return *o.GcpCredType
}

// GetGcpCredTypeOk returns a tuple with the GcpCredType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetGcpCredTypeOk() (*string, bool) {
	if o == nil || o.GcpCredType == nil {
		return nil, false
	}
	return o.GcpCredType, true
}

// HasGcpCredType returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasGcpCredType() bool {
	if o != nil && o.GcpCredType != nil {
		return true
	}

	return false
}

// SetGcpCredType gets a reference to the given string and assigns it to the GcpCredType field.
func (o *GatewayCreateProducerGcp) SetGcpCredType(v string) {
	o.GcpCredType = &v
}

// GetGcpKey returns the GcpKey field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetGcpKey() string {
	if o == nil || o.GcpKey == nil {
		var ret string
		return ret
	}
	return *o.GcpKey
}

// GetGcpKeyOk returns a tuple with the GcpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetGcpKeyOk() (*string, bool) {
	if o == nil || o.GcpKey == nil {
		return nil, false
	}
	return o.GcpKey, true
}

// HasGcpKey returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasGcpKey() bool {
	if o != nil && o.GcpKey != nil {
		return true
	}

	return false
}

// SetGcpKey gets a reference to the given string and assigns it to the GcpKey field.
func (o *GatewayCreateProducerGcp) SetGcpKey(v string) {
	o.GcpKey = &v
}

// GetGcpKeyAlgo returns the GcpKeyAlgo field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetGcpKeyAlgo() string {
	if o == nil || o.GcpKeyAlgo == nil {
		var ret string
		return ret
	}
	return *o.GcpKeyAlgo
}

// GetGcpKeyAlgoOk returns a tuple with the GcpKeyAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetGcpKeyAlgoOk() (*string, bool) {
	if o == nil || o.GcpKeyAlgo == nil {
		return nil, false
	}
	return o.GcpKeyAlgo, true
}

// HasGcpKeyAlgo returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasGcpKeyAlgo() bool {
	if o != nil && o.GcpKeyAlgo != nil {
		return true
	}

	return false
}

// SetGcpKeyAlgo gets a reference to the given string and assigns it to the GcpKeyAlgo field.
func (o *GatewayCreateProducerGcp) SetGcpKeyAlgo(v string) {
	o.GcpKeyAlgo = &v
}

// GetGcpSaEmail returns the GcpSaEmail field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetGcpSaEmail() string {
	if o == nil || o.GcpSaEmail == nil {
		var ret string
		return ret
	}
	return *o.GcpSaEmail
}

// GetGcpSaEmailOk returns a tuple with the GcpSaEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetGcpSaEmailOk() (*string, bool) {
	if o == nil || o.GcpSaEmail == nil {
		return nil, false
	}
	return o.GcpSaEmail, true
}

// HasGcpSaEmail returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasGcpSaEmail() bool {
	if o != nil && o.GcpSaEmail != nil {
		return true
	}

	return false
}

// SetGcpSaEmail gets a reference to the given string and assigns it to the GcpSaEmail field.
func (o *GatewayCreateProducerGcp) SetGcpSaEmail(v string) {
	o.GcpSaEmail = &v
}

// GetGcpTokenScopes returns the GcpTokenScopes field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetGcpTokenScopes() string {
	if o == nil || o.GcpTokenScopes == nil {
		var ret string
		return ret
	}
	return *o.GcpTokenScopes
}

// GetGcpTokenScopesOk returns a tuple with the GcpTokenScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetGcpTokenScopesOk() (*string, bool) {
	if o == nil || o.GcpTokenScopes == nil {
		return nil, false
	}
	return o.GcpTokenScopes, true
}

// HasGcpTokenScopes returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasGcpTokenScopes() bool {
	if o != nil && o.GcpTokenScopes != nil {
		return true
	}

	return false
}

// SetGcpTokenScopes gets a reference to the given string and assigns it to the GcpTokenScopes field.
func (o *GatewayCreateProducerGcp) SetGcpTokenScopes(v string) {
	o.GcpTokenScopes = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerGcp) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerGcp) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerGcp) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerGcp) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerGcp) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerGcp) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerGcp) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerGcp) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerGcp) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerGcp) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerGcp) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerGcp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GcpCredType != nil {
		toSerialize["gcp-cred-type"] = o.GcpCredType
	}
	if o.GcpKey != nil {
		toSerialize["gcp-key"] = o.GcpKey
	}
	if o.GcpKeyAlgo != nil {
		toSerialize["gcp-key-algo"] = o.GcpKeyAlgo
	}
	if o.GcpSaEmail != nil {
		toSerialize["gcp-sa-email"] = o.GcpSaEmail
	}
	if o.GcpTokenScopes != nil {
		toSerialize["gcp-token-scopes"] = o.GcpTokenScopes
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
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

type NullableGatewayCreateProducerGcp struct {
	value *GatewayCreateProducerGcp
	isSet bool
}

func (v NullableGatewayCreateProducerGcp) Get() *GatewayCreateProducerGcp {
	return v.value
}

func (v *NullableGatewayCreateProducerGcp) Set(val *GatewayCreateProducerGcp) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerGcp) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerGcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerGcp(val *GatewayCreateProducerGcp) *NullableGatewayCreateProducerGcp {
	return &NullableGatewayCreateProducerGcp{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerGcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerGcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


