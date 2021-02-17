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

// GatewayCreateProducerArtifactory gatewayCreateProducerArtifactory is a command that creates artifactory producer
type GatewayCreateProducerArtifactory struct {
	// Artifactory Admin Name
	ArtifactoryAdminName string `json:"artifactory-admin-name"`
	// Artifactory Admin password
	ArtifactoryAdminPwd string `json:"artifactory-admin-pwd"`
	// Token Audience
	ArtifactoryTokenAudience string `json:"artifactory-token-audience"`
	// Token Scope
	ArtifactoryTokenScope string `json:"artifactory-token-scope"`
	// Base URL
	BaseUrl string `json:"base-url"`
	// Gateway url
	GatewayUrl *string `json:"gateway-url,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayCreateProducerArtifactory instantiates a new GatewayCreateProducerArtifactory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerArtifactory(artifactoryAdminName string, artifactoryAdminPwd string, artifactoryTokenAudience string, artifactoryTokenScope string, baseUrl string, name string, ) *GatewayCreateProducerArtifactory {
	this := GatewayCreateProducerArtifactory{}
	this.ArtifactoryAdminName = artifactoryAdminName
	this.ArtifactoryAdminPwd = artifactoryAdminPwd
	this.ArtifactoryTokenAudience = artifactoryTokenAudience
	this.ArtifactoryTokenScope = artifactoryTokenScope
	this.BaseUrl = baseUrl
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerArtifactoryWithDefaults instantiates a new GatewayCreateProducerArtifactory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerArtifactoryWithDefaults() *GatewayCreateProducerArtifactory {
	this := GatewayCreateProducerArtifactory{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetArtifactoryAdminName returns the ArtifactoryAdminName field value
func (o *GatewayCreateProducerArtifactory) GetArtifactoryAdminName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactoryAdminName
}

// GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetArtifactoryAdminNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactoryAdminName, true
}

// SetArtifactoryAdminName sets field value
func (o *GatewayCreateProducerArtifactory) SetArtifactoryAdminName(v string) {
	o.ArtifactoryAdminName = v
}

// GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field value
func (o *GatewayCreateProducerArtifactory) GetArtifactoryAdminPwd() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactoryAdminPwd
}

// GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetArtifactoryAdminPwdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactoryAdminPwd, true
}

// SetArtifactoryAdminPwd sets field value
func (o *GatewayCreateProducerArtifactory) SetArtifactoryAdminPwd(v string) {
	o.ArtifactoryAdminPwd = v
}

// GetArtifactoryTokenAudience returns the ArtifactoryTokenAudience field value
func (o *GatewayCreateProducerArtifactory) GetArtifactoryTokenAudience() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactoryTokenAudience
}

// GetArtifactoryTokenAudienceOk returns a tuple with the ArtifactoryTokenAudience field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetArtifactoryTokenAudienceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactoryTokenAudience, true
}

// SetArtifactoryTokenAudience sets field value
func (o *GatewayCreateProducerArtifactory) SetArtifactoryTokenAudience(v string) {
	o.ArtifactoryTokenAudience = v
}

// GetArtifactoryTokenScope returns the ArtifactoryTokenScope field value
func (o *GatewayCreateProducerArtifactory) GetArtifactoryTokenScope() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactoryTokenScope
}

// GetArtifactoryTokenScopeOk returns a tuple with the ArtifactoryTokenScope field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetArtifactoryTokenScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactoryTokenScope, true
}

// SetArtifactoryTokenScope sets field value
func (o *GatewayCreateProducerArtifactory) SetArtifactoryTokenScope(v string) {
	o.ArtifactoryTokenScope = v
}

// GetBaseUrl returns the BaseUrl field value
func (o *GatewayCreateProducerArtifactory) GetBaseUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetBaseUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *GatewayCreateProducerArtifactory) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetGatewayUrl returns the GatewayUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerArtifactory) GetGatewayUrl() string {
	if o == nil || o.GatewayUrl == nil {
		var ret string
		return ret
	}
	return *o.GatewayUrl
}

// GetGatewayUrlOk returns a tuple with the GatewayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetGatewayUrlOk() (*string, bool) {
	if o == nil || o.GatewayUrl == nil {
		return nil, false
	}
	return o.GatewayUrl, true
}

// HasGatewayUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerArtifactory) HasGatewayUrl() bool {
	if o != nil && o.GatewayUrl != nil {
		return true
	}

	return false
}

// SetGatewayUrl gets a reference to the given string and assigns it to the GatewayUrl field.
func (o *GatewayCreateProducerArtifactory) SetGatewayUrl(v string) {
	o.GatewayUrl = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerArtifactory) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerArtifactory) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerArtifactory) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerArtifactory) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerArtifactory) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerArtifactory) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerArtifactory) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerArtifactory) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerArtifactory) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerArtifactory) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerArtifactory) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerArtifactory) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerArtifactory) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerArtifactory) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerArtifactory) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerArtifactory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["artifactory-admin-name"] = o.ArtifactoryAdminName
	}
	if true {
		toSerialize["artifactory-admin-pwd"] = o.ArtifactoryAdminPwd
	}
	if true {
		toSerialize["artifactory-token-audience"] = o.ArtifactoryTokenAudience
	}
	if true {
		toSerialize["artifactory-token-scope"] = o.ArtifactoryTokenScope
	}
	if true {
		toSerialize["base-url"] = o.BaseUrl
	}
	if o.GatewayUrl != nil {
		toSerialize["gateway-url"] = o.GatewayUrl
	}
	if true {
		toSerialize["name"] = o.Name
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
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerArtifactory struct {
	value *GatewayCreateProducerArtifactory
	isSet bool
}

func (v NullableGatewayCreateProducerArtifactory) Get() *GatewayCreateProducerArtifactory {
	return v.value
}

func (v *NullableGatewayCreateProducerArtifactory) Set(val *GatewayCreateProducerArtifactory) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerArtifactory) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerArtifactory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerArtifactory(val *GatewayCreateProducerArtifactory) *NullableGatewayCreateProducerArtifactory {
	return &NullableGatewayCreateProducerArtifactory{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerArtifactory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerArtifactory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

