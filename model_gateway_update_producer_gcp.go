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

// GatewayUpdateProducerGcp gatewayUpdateProducerGcp is a command that updates a GCP producer [Deprecated: Use dynamic-secret-update-gcp command]
type GatewayUpdateProducerGcp struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	GcpCredType *string `json:"gcp-cred-type,omitempty"`
	// Base64-encoded service account private key text
	GcpKey *string `json:"gcp-key,omitempty"`
	// Service account key algorithm, e.g. KEY_ALG_RSA_1024
	GcpKeyAlgo *string `json:"gcp-key-algo,omitempty"`
	// The email of the fixed service acocunt to generate keys or tokens for. (revelant for service-account-type=fixed)
	GcpSaEmail *string `json:"gcp-sa-email,omitempty"`
	// Access token scopes list, e.g. scope1,scope2
	GcpTokenScopes *string `json:"gcp-token-scopes,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret name
	NewName *string `json:"new-name,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Role binding definitions in json format
	RoleBinding *string `json:"role-binding,omitempty"`
	// The type of the gcp dynamic secret. Options[fixed, dynamic]
	ServiceAccountType string `json:"service-account-type"`
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

// NewGatewayUpdateProducerGcp instantiates a new GatewayUpdateProducerGcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerGcp(name string, serviceAccountType string, ) *GatewayUpdateProducerGcp {
	this := GatewayUpdateProducerGcp{}
	var json bool = false
	this.Json = &json
	this.Name = name
	this.ServiceAccountType = serviceAccountType
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerGcpWithDefaults instantiates a new GatewayUpdateProducerGcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerGcpWithDefaults() *GatewayUpdateProducerGcp {
	this := GatewayUpdateProducerGcp{}
	var json bool = false
	this.Json = &json
	var serviceAccountType string = "fixed"
	this.ServiceAccountType = serviceAccountType
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerGcp) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetGcpCredType returns the GcpCredType field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetGcpCredType() string {
	if o == nil || o.GcpCredType == nil {
		var ret string
		return ret
	}
	return *o.GcpCredType
}

// GetGcpCredTypeOk returns a tuple with the GcpCredType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetGcpCredTypeOk() (*string, bool) {
	if o == nil || o.GcpCredType == nil {
		return nil, false
	}
	return o.GcpCredType, true
}

// HasGcpCredType returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasGcpCredType() bool {
	if o != nil && o.GcpCredType != nil {
		return true
	}

	return false
}

// SetGcpCredType gets a reference to the given string and assigns it to the GcpCredType field.
func (o *GatewayUpdateProducerGcp) SetGcpCredType(v string) {
	o.GcpCredType = &v
}

// GetGcpKey returns the GcpKey field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetGcpKey() string {
	if o == nil || o.GcpKey == nil {
		var ret string
		return ret
	}
	return *o.GcpKey
}

// GetGcpKeyOk returns a tuple with the GcpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetGcpKeyOk() (*string, bool) {
	if o == nil || o.GcpKey == nil {
		return nil, false
	}
	return o.GcpKey, true
}

// HasGcpKey returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasGcpKey() bool {
	if o != nil && o.GcpKey != nil {
		return true
	}

	return false
}

// SetGcpKey gets a reference to the given string and assigns it to the GcpKey field.
func (o *GatewayUpdateProducerGcp) SetGcpKey(v string) {
	o.GcpKey = &v
}

// GetGcpKeyAlgo returns the GcpKeyAlgo field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetGcpKeyAlgo() string {
	if o == nil || o.GcpKeyAlgo == nil {
		var ret string
		return ret
	}
	return *o.GcpKeyAlgo
}

// GetGcpKeyAlgoOk returns a tuple with the GcpKeyAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetGcpKeyAlgoOk() (*string, bool) {
	if o == nil || o.GcpKeyAlgo == nil {
		return nil, false
	}
	return o.GcpKeyAlgo, true
}

// HasGcpKeyAlgo returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasGcpKeyAlgo() bool {
	if o != nil && o.GcpKeyAlgo != nil {
		return true
	}

	return false
}

// SetGcpKeyAlgo gets a reference to the given string and assigns it to the GcpKeyAlgo field.
func (o *GatewayUpdateProducerGcp) SetGcpKeyAlgo(v string) {
	o.GcpKeyAlgo = &v
}

// GetGcpSaEmail returns the GcpSaEmail field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetGcpSaEmail() string {
	if o == nil || o.GcpSaEmail == nil {
		var ret string
		return ret
	}
	return *o.GcpSaEmail
}

// GetGcpSaEmailOk returns a tuple with the GcpSaEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetGcpSaEmailOk() (*string, bool) {
	if o == nil || o.GcpSaEmail == nil {
		return nil, false
	}
	return o.GcpSaEmail, true
}

// HasGcpSaEmail returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasGcpSaEmail() bool {
	if o != nil && o.GcpSaEmail != nil {
		return true
	}

	return false
}

// SetGcpSaEmail gets a reference to the given string and assigns it to the GcpSaEmail field.
func (o *GatewayUpdateProducerGcp) SetGcpSaEmail(v string) {
	o.GcpSaEmail = &v
}

// GetGcpTokenScopes returns the GcpTokenScopes field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetGcpTokenScopes() string {
	if o == nil || o.GcpTokenScopes == nil {
		var ret string
		return ret
	}
	return *o.GcpTokenScopes
}

// GetGcpTokenScopesOk returns a tuple with the GcpTokenScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetGcpTokenScopesOk() (*string, bool) {
	if o == nil || o.GcpTokenScopes == nil {
		return nil, false
	}
	return o.GcpTokenScopes, true
}

// HasGcpTokenScopes returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasGcpTokenScopes() bool {
	if o != nil && o.GcpTokenScopes != nil {
		return true
	}

	return false
}

// SetGcpTokenScopes gets a reference to the given string and assigns it to the GcpTokenScopes field.
func (o *GatewayUpdateProducerGcp) SetGcpTokenScopes(v string) {
	o.GcpTokenScopes = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerGcp) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerGcp) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerGcp) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerGcp) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerGcp) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRoleBinding returns the RoleBinding field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetRoleBinding() string {
	if o == nil || o.RoleBinding == nil {
		var ret string
		return ret
	}
	return *o.RoleBinding
}

// GetRoleBindingOk returns a tuple with the RoleBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetRoleBindingOk() (*string, bool) {
	if o == nil || o.RoleBinding == nil {
		return nil, false
	}
	return o.RoleBinding, true
}

// HasRoleBinding returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasRoleBinding() bool {
	if o != nil && o.RoleBinding != nil {
		return true
	}

	return false
}

// SetRoleBinding gets a reference to the given string and assigns it to the RoleBinding field.
func (o *GatewayUpdateProducerGcp) SetRoleBinding(v string) {
	o.RoleBinding = &v
}

// GetServiceAccountType returns the ServiceAccountType field value
func (o *GatewayUpdateProducerGcp) GetServiceAccountType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServiceAccountType
}

// GetServiceAccountTypeOk returns a tuple with the ServiceAccountType field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetServiceAccountTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceAccountType, true
}

// SetServiceAccountType sets field value
func (o *GatewayUpdateProducerGcp) SetServiceAccountType(v string) {
	o.ServiceAccountType = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerGcp) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerGcp) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerGcp) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerGcp) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerGcp) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerGcp) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerGcp) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerGcp) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerGcp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
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
	if o.RoleBinding != nil {
		toSerialize["role-binding"] = o.RoleBinding
	}
	if true {
		toSerialize["service-account-type"] = o.ServiceAccountType
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

type NullableGatewayUpdateProducerGcp struct {
	value *GatewayUpdateProducerGcp
	isSet bool
}

func (v NullableGatewayUpdateProducerGcp) Get() *GatewayUpdateProducerGcp {
	return v.value
}

func (v *NullableGatewayUpdateProducerGcp) Set(val *GatewayUpdateProducerGcp) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerGcp) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerGcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerGcp(val *GatewayUpdateProducerGcp) *NullableGatewayUpdateProducerGcp {
	return &NullableGatewayUpdateProducerGcp{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerGcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerGcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


