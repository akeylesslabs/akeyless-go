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

// DynamicSecretCreateArtifactory dynamicSecretCreateArtifactory is a command that creates artifactory dynamic secret
type DynamicSecretCreateArtifactory struct {
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
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
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

// NewDynamicSecretCreateArtifactory instantiates a new DynamicSecretCreateArtifactory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretCreateArtifactory(artifactoryTokenAudience string, artifactoryTokenScope string, name string, ) *DynamicSecretCreateArtifactory {
	this := DynamicSecretCreateArtifactory{}
	this.ArtifactoryTokenAudience = artifactoryTokenAudience
	this.ArtifactoryTokenScope = artifactoryTokenScope
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretCreateArtifactoryWithDefaults instantiates a new DynamicSecretCreateArtifactory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretCreateArtifactoryWithDefaults() *DynamicSecretCreateArtifactory {
	this := DynamicSecretCreateArtifactory{}
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetArtifactoryAdminName returns the ArtifactoryAdminName field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetArtifactoryAdminName() string {
	if o == nil || o.ArtifactoryAdminName == nil {
		var ret string
		return ret
	}
	return *o.ArtifactoryAdminName
}

// GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetArtifactoryAdminNameOk() (*string, bool) {
	if o == nil || o.ArtifactoryAdminName == nil {
		return nil, false
	}
	return o.ArtifactoryAdminName, true
}

// HasArtifactoryAdminName returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasArtifactoryAdminName() bool {
	if o != nil && o.ArtifactoryAdminName != nil {
		return true
	}

	return false
}

// SetArtifactoryAdminName gets a reference to the given string and assigns it to the ArtifactoryAdminName field.
func (o *DynamicSecretCreateArtifactory) SetArtifactoryAdminName(v string) {
	o.ArtifactoryAdminName = &v
}

// GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetArtifactoryAdminPwd() string {
	if o == nil || o.ArtifactoryAdminPwd == nil {
		var ret string
		return ret
	}
	return *o.ArtifactoryAdminPwd
}

// GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetArtifactoryAdminPwdOk() (*string, bool) {
	if o == nil || o.ArtifactoryAdminPwd == nil {
		return nil, false
	}
	return o.ArtifactoryAdminPwd, true
}

// HasArtifactoryAdminPwd returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasArtifactoryAdminPwd() bool {
	if o != nil && o.ArtifactoryAdminPwd != nil {
		return true
	}

	return false
}

// SetArtifactoryAdminPwd gets a reference to the given string and assigns it to the ArtifactoryAdminPwd field.
func (o *DynamicSecretCreateArtifactory) SetArtifactoryAdminPwd(v string) {
	o.ArtifactoryAdminPwd = &v
}

// GetArtifactoryTokenAudience returns the ArtifactoryTokenAudience field value
func (o *DynamicSecretCreateArtifactory) GetArtifactoryTokenAudience() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactoryTokenAudience
}

// GetArtifactoryTokenAudienceOk returns a tuple with the ArtifactoryTokenAudience field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetArtifactoryTokenAudienceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactoryTokenAudience, true
}

// SetArtifactoryTokenAudience sets field value
func (o *DynamicSecretCreateArtifactory) SetArtifactoryTokenAudience(v string) {
	o.ArtifactoryTokenAudience = v
}

// GetArtifactoryTokenScope returns the ArtifactoryTokenScope field value
func (o *DynamicSecretCreateArtifactory) GetArtifactoryTokenScope() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactoryTokenScope
}

// GetArtifactoryTokenScopeOk returns a tuple with the ArtifactoryTokenScope field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetArtifactoryTokenScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactoryTokenScope, true
}

// SetArtifactoryTokenScope sets field value
func (o *DynamicSecretCreateArtifactory) SetArtifactoryTokenScope(v string) {
	o.ArtifactoryTokenScope = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *DynamicSecretCreateArtifactory) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretCreateArtifactory) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretCreateArtifactory) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretCreateArtifactory) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretCreateArtifactory) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretCreateArtifactory) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretCreateArtifactory) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretCreateArtifactory) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretCreateArtifactory) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretCreateArtifactory) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretCreateArtifactory) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateArtifactory) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretCreateArtifactory) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretCreateArtifactory) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretCreateArtifactory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactoryAdminName != nil {
		toSerialize["artifactory-admin-name"] = o.ArtifactoryAdminName
	}
	if o.ArtifactoryAdminPwd != nil {
		toSerialize["artifactory-admin-pwd"] = o.ArtifactoryAdminPwd
	}
	if true {
		toSerialize["artifactory-token-audience"] = o.ArtifactoryTokenAudience
	}
	if true {
		toSerialize["artifactory-token-scope"] = o.ArtifactoryTokenScope
	}
	if o.BaseUrl != nil {
		toSerialize["base-url"] = o.BaseUrl
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
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

type NullableDynamicSecretCreateArtifactory struct {
	value *DynamicSecretCreateArtifactory
	isSet bool
}

func (v NullableDynamicSecretCreateArtifactory) Get() *DynamicSecretCreateArtifactory {
	return v.value
}

func (v *NullableDynamicSecretCreateArtifactory) Set(val *DynamicSecretCreateArtifactory) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretCreateArtifactory) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretCreateArtifactory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretCreateArtifactory(val *DynamicSecretCreateArtifactory) *NullableDynamicSecretCreateArtifactory {
	return &NullableDynamicSecretCreateArtifactory{value: val, isSet: true}
}

func (v NullableDynamicSecretCreateArtifactory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretCreateArtifactory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


