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

// DynamicSecretUpdateArtifactory dynamicSecretUpdateArtifactory is a command that updates artifactory dynamic secret
type DynamicSecretUpdateArtifactory struct {
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
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret new name
	NewName *string `json:"new-name,omitempty"`
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

// NewDynamicSecretUpdateArtifactory instantiates a new DynamicSecretUpdateArtifactory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdateArtifactory(artifactoryTokenAudience string, artifactoryTokenScope string, name string, ) *DynamicSecretUpdateArtifactory {
	this := DynamicSecretUpdateArtifactory{}
	this.ArtifactoryTokenAudience = artifactoryTokenAudience
	this.ArtifactoryTokenScope = artifactoryTokenScope
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretUpdateArtifactoryWithDefaults instantiates a new DynamicSecretUpdateArtifactory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdateArtifactoryWithDefaults() *DynamicSecretUpdateArtifactory {
	this := DynamicSecretUpdateArtifactory{}
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetArtifactoryAdminName returns the ArtifactoryAdminName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetArtifactoryAdminName() string {
	if o == nil || o.ArtifactoryAdminName == nil {
		var ret string
		return ret
	}
	return *o.ArtifactoryAdminName
}

// GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetArtifactoryAdminNameOk() (*string, bool) {
	if o == nil || o.ArtifactoryAdminName == nil {
		return nil, false
	}
	return o.ArtifactoryAdminName, true
}

// HasArtifactoryAdminName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasArtifactoryAdminName() bool {
	if o != nil && o.ArtifactoryAdminName != nil {
		return true
	}

	return false
}

// SetArtifactoryAdminName gets a reference to the given string and assigns it to the ArtifactoryAdminName field.
func (o *DynamicSecretUpdateArtifactory) SetArtifactoryAdminName(v string) {
	o.ArtifactoryAdminName = &v
}

// GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetArtifactoryAdminPwd() string {
	if o == nil || o.ArtifactoryAdminPwd == nil {
		var ret string
		return ret
	}
	return *o.ArtifactoryAdminPwd
}

// GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetArtifactoryAdminPwdOk() (*string, bool) {
	if o == nil || o.ArtifactoryAdminPwd == nil {
		return nil, false
	}
	return o.ArtifactoryAdminPwd, true
}

// HasArtifactoryAdminPwd returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasArtifactoryAdminPwd() bool {
	if o != nil && o.ArtifactoryAdminPwd != nil {
		return true
	}

	return false
}

// SetArtifactoryAdminPwd gets a reference to the given string and assigns it to the ArtifactoryAdminPwd field.
func (o *DynamicSecretUpdateArtifactory) SetArtifactoryAdminPwd(v string) {
	o.ArtifactoryAdminPwd = &v
}

// GetArtifactoryTokenAudience returns the ArtifactoryTokenAudience field value
func (o *DynamicSecretUpdateArtifactory) GetArtifactoryTokenAudience() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactoryTokenAudience
}

// GetArtifactoryTokenAudienceOk returns a tuple with the ArtifactoryTokenAudience field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetArtifactoryTokenAudienceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactoryTokenAudience, true
}

// SetArtifactoryTokenAudience sets field value
func (o *DynamicSecretUpdateArtifactory) SetArtifactoryTokenAudience(v string) {
	o.ArtifactoryTokenAudience = v
}

// GetArtifactoryTokenScope returns the ArtifactoryTokenScope field value
func (o *DynamicSecretUpdateArtifactory) GetArtifactoryTokenScope() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactoryTokenScope
}

// GetArtifactoryTokenScopeOk returns a tuple with the ArtifactoryTokenScope field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetArtifactoryTokenScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactoryTokenScope, true
}

// SetArtifactoryTokenScope sets field value
func (o *DynamicSecretUpdateArtifactory) SetArtifactoryTokenScope(v string) {
	o.ArtifactoryTokenScope = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *DynamicSecretUpdateArtifactory) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretUpdateArtifactory) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretUpdateArtifactory) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretUpdateArtifactory) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretUpdateArtifactory) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretUpdateArtifactory) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DynamicSecretUpdateArtifactory) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretUpdateArtifactory) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretUpdateArtifactory) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretUpdateArtifactory) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretUpdateArtifactory) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretUpdateArtifactory) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateArtifactory) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateArtifactory) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateArtifactory) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretUpdateArtifactory) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretUpdateArtifactory) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
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

type NullableDynamicSecretUpdateArtifactory struct {
	value *DynamicSecretUpdateArtifactory
	isSet bool
}

func (v NullableDynamicSecretUpdateArtifactory) Get() *DynamicSecretUpdateArtifactory {
	return v.value
}

func (v *NullableDynamicSecretUpdateArtifactory) Set(val *DynamicSecretUpdateArtifactory) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdateArtifactory) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdateArtifactory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdateArtifactory(val *DynamicSecretUpdateArtifactory) *NullableDynamicSecretUpdateArtifactory {
	return &NullableDynamicSecretUpdateArtifactory{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdateArtifactory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdateArtifactory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


