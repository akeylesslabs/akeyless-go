/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// checks if the GatewayUpdateProducerArtifactory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayUpdateProducerArtifactory{}

// GatewayUpdateProducerArtifactory gatewayUpdateProducerArtifactory is a command that updates artifactory producer
type GatewayUpdateProducerArtifactory struct {
	// Artifactory Admin Name
	ArtifactoryAdminName *string `json:"artifactory-admin-name,omitempty"`
	// Artifactory Admin password
	ArtifactoryAdminPwd *string `json:"artifactory-admin-pwd,omitempty"`
	// Token Audience
	ArtifactoryTokenAudience string `json:"artifactory-token-audience"`
	// Token Scope
	ArtifactoryTokenScope string `json:"artifactory-token-scope"`
	// Base URL
	BaseUrl *string `json:"base-url,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Add tags attached to this object
	Tags []string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayUpdateProducerArtifactory instantiates a new GatewayUpdateProducerArtifactory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerArtifactory(artifactoryTokenAudience string, artifactoryTokenScope string, name string) *GatewayUpdateProducerArtifactory {
	this := GatewayUpdateProducerArtifactory{}
	this.ArtifactoryTokenAudience = artifactoryTokenAudience
	this.ArtifactoryTokenScope = artifactoryTokenScope
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerArtifactoryWithDefaults instantiates a new GatewayUpdateProducerArtifactory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerArtifactoryWithDefaults() *GatewayUpdateProducerArtifactory {
	this := GatewayUpdateProducerArtifactory{}
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetArtifactoryAdminName returns the ArtifactoryAdminName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetArtifactoryAdminName() string {
	if o == nil || IsNil(o.ArtifactoryAdminName) {
		var ret string
		return ret
	}
	return *o.ArtifactoryAdminName
}

// GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetArtifactoryAdminNameOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactoryAdminName) {
		return nil, false
	}
	return o.ArtifactoryAdminName, true
}

// HasArtifactoryAdminName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasArtifactoryAdminName() bool {
	if o != nil && !IsNil(o.ArtifactoryAdminName) {
		return true
	}

	return false
}

// SetArtifactoryAdminName gets a reference to the given string and assigns it to the ArtifactoryAdminName field.
func (o *GatewayUpdateProducerArtifactory) SetArtifactoryAdminName(v string) {
	o.ArtifactoryAdminName = &v
}

// GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetArtifactoryAdminPwd() string {
	if o == nil || IsNil(o.ArtifactoryAdminPwd) {
		var ret string
		return ret
	}
	return *o.ArtifactoryAdminPwd
}

// GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetArtifactoryAdminPwdOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactoryAdminPwd) {
		return nil, false
	}
	return o.ArtifactoryAdminPwd, true
}

// HasArtifactoryAdminPwd returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasArtifactoryAdminPwd() bool {
	if o != nil && !IsNil(o.ArtifactoryAdminPwd) {
		return true
	}

	return false
}

// SetArtifactoryAdminPwd gets a reference to the given string and assigns it to the ArtifactoryAdminPwd field.
func (o *GatewayUpdateProducerArtifactory) SetArtifactoryAdminPwd(v string) {
	o.ArtifactoryAdminPwd = &v
}

// GetArtifactoryTokenAudience returns the ArtifactoryTokenAudience field value
func (o *GatewayUpdateProducerArtifactory) GetArtifactoryTokenAudience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactoryTokenAudience
}

// GetArtifactoryTokenAudienceOk returns a tuple with the ArtifactoryTokenAudience field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetArtifactoryTokenAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactoryTokenAudience, true
}

// SetArtifactoryTokenAudience sets field value
func (o *GatewayUpdateProducerArtifactory) SetArtifactoryTokenAudience(v string) {
	o.ArtifactoryTokenAudience = v
}

// GetArtifactoryTokenScope returns the ArtifactoryTokenScope field value
func (o *GatewayUpdateProducerArtifactory) GetArtifactoryTokenScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactoryTokenScope
}

// GetArtifactoryTokenScopeOk returns a tuple with the ArtifactoryTokenScope field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetArtifactoryTokenScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactoryTokenScope, true
}

// SetArtifactoryTokenScope sets field value
func (o *GatewayUpdateProducerArtifactory) SetArtifactoryTokenScope(v string) {
	o.ArtifactoryTokenScope = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *GatewayUpdateProducerArtifactory) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetDeleteProtection() string {
	if o == nil || IsNil(o.DeleteProtection) {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || IsNil(o.DeleteProtection) {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasDeleteProtection() bool {
	if o != nil && !IsNil(o.DeleteProtection) {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerArtifactory) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerArtifactory) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerArtifactory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerArtifactory) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetNewName() string {
	if o == nil || IsNil(o.NewName) {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetNewNameOk() (*string, bool) {
	if o == nil || IsNil(o.NewName) {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasNewName() bool {
	if o != nil && !IsNil(o.NewName) {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerArtifactory) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetProducerEncryptionKeyName() string {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasProducerEncryptionKeyName() bool {
	if o != nil && !IsNil(o.ProducerEncryptionKeyName) {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerArtifactory) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerArtifactory) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerArtifactory) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerArtifactory) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerArtifactory) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerArtifactory) GetUserTtl() string {
	if o == nil || IsNil(o.UserTtl) {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerArtifactory) GetUserTtlOk() (*string, bool) {
	if o == nil || IsNil(o.UserTtl) {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerArtifactory) HasUserTtl() bool {
	if o != nil && !IsNil(o.UserTtl) {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerArtifactory) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerArtifactory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayUpdateProducerArtifactory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArtifactoryAdminName) {
		toSerialize["artifactory-admin-name"] = o.ArtifactoryAdminName
	}
	if !IsNil(o.ArtifactoryAdminPwd) {
		toSerialize["artifactory-admin-pwd"] = o.ArtifactoryAdminPwd
	}
	toSerialize["artifactory-token-audience"] = o.ArtifactoryTokenAudience
	toSerialize["artifactory-token-scope"] = o.ArtifactoryTokenScope
	if !IsNil(o.BaseUrl) {
		toSerialize["base-url"] = o.BaseUrl
	}
	if !IsNil(o.DeleteProtection) {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NewName) {
		toSerialize["new-name"] = o.NewName
	}
	if !IsNil(o.ProducerEncryptionKeyName) {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TargetName) {
		toSerialize["target-name"] = o.TargetName
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	if !IsNil(o.UserTtl) {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return toSerialize, nil
}

type NullableGatewayUpdateProducerArtifactory struct {
	value *GatewayUpdateProducerArtifactory
	isSet bool
}

func (v NullableGatewayUpdateProducerArtifactory) Get() *GatewayUpdateProducerArtifactory {
	return v.value
}

func (v *NullableGatewayUpdateProducerArtifactory) Set(val *GatewayUpdateProducerArtifactory) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerArtifactory) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerArtifactory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerArtifactory(val *GatewayUpdateProducerArtifactory) *NullableGatewayUpdateProducerArtifactory {
	return &NullableGatewayUpdateProducerArtifactory{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerArtifactory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerArtifactory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


