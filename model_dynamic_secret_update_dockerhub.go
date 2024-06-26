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

// DynamicSecretUpdateDockerhub dynamicSecretUpdateDockerhub is a command that updates a DOCKERHUB dynamic secret
type DynamicSecretUpdateDockerhub struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// DockerhubPassword is either the user's password access token to manage the repository
	DockerhubPassword *string `json:"dockerhub-password,omitempty"`
	// Access token scopes list (comma-separated) to give the dynamic secret valid options are in \"repo:admin\", \"repo:write\", \"repo:read\", \"repo:public_read\"
	DockerhubTokenScopes *string `json:"dockerhub-token-scopes,omitempty"`
	// DockerhubUsername is the name of the user in dockerhub
	DockerhubUsername *string `json:"dockerhub-username,omitempty"`
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

// NewDynamicSecretUpdateDockerhub instantiates a new DynamicSecretUpdateDockerhub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdateDockerhub(name string, ) *DynamicSecretUpdateDockerhub {
	this := DynamicSecretUpdateDockerhub{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretUpdateDockerhubWithDefaults instantiates a new DynamicSecretUpdateDockerhub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdateDockerhubWithDefaults() *DynamicSecretUpdateDockerhub {
	this := DynamicSecretUpdateDockerhub{}
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretUpdateDockerhub) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretUpdateDockerhub) SetDescription(v string) {
	o.Description = &v
}

// GetDockerhubPassword returns the DockerhubPassword field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetDockerhubPassword() string {
	if o == nil || o.DockerhubPassword == nil {
		var ret string
		return ret
	}
	return *o.DockerhubPassword
}

// GetDockerhubPasswordOk returns a tuple with the DockerhubPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetDockerhubPasswordOk() (*string, bool) {
	if o == nil || o.DockerhubPassword == nil {
		return nil, false
	}
	return o.DockerhubPassword, true
}

// HasDockerhubPassword returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasDockerhubPassword() bool {
	if o != nil && o.DockerhubPassword != nil {
		return true
	}

	return false
}

// SetDockerhubPassword gets a reference to the given string and assigns it to the DockerhubPassword field.
func (o *DynamicSecretUpdateDockerhub) SetDockerhubPassword(v string) {
	o.DockerhubPassword = &v
}

// GetDockerhubTokenScopes returns the DockerhubTokenScopes field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetDockerhubTokenScopes() string {
	if o == nil || o.DockerhubTokenScopes == nil {
		var ret string
		return ret
	}
	return *o.DockerhubTokenScopes
}

// GetDockerhubTokenScopesOk returns a tuple with the DockerhubTokenScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetDockerhubTokenScopesOk() (*string, bool) {
	if o == nil || o.DockerhubTokenScopes == nil {
		return nil, false
	}
	return o.DockerhubTokenScopes, true
}

// HasDockerhubTokenScopes returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasDockerhubTokenScopes() bool {
	if o != nil && o.DockerhubTokenScopes != nil {
		return true
	}

	return false
}

// SetDockerhubTokenScopes gets a reference to the given string and assigns it to the DockerhubTokenScopes field.
func (o *DynamicSecretUpdateDockerhub) SetDockerhubTokenScopes(v string) {
	o.DockerhubTokenScopes = &v
}

// GetDockerhubUsername returns the DockerhubUsername field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetDockerhubUsername() string {
	if o == nil || o.DockerhubUsername == nil {
		var ret string
		return ret
	}
	return *o.DockerhubUsername
}

// GetDockerhubUsernameOk returns a tuple with the DockerhubUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetDockerhubUsernameOk() (*string, bool) {
	if o == nil || o.DockerhubUsername == nil {
		return nil, false
	}
	return o.DockerhubUsername, true
}

// HasDockerhubUsername returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasDockerhubUsername() bool {
	if o != nil && o.DockerhubUsername != nil {
		return true
	}

	return false
}

// SetDockerhubUsername gets a reference to the given string and assigns it to the DockerhubUsername field.
func (o *DynamicSecretUpdateDockerhub) SetDockerhubUsername(v string) {
	o.DockerhubUsername = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretUpdateDockerhub) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretUpdateDockerhub) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretUpdateDockerhub) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DynamicSecretUpdateDockerhub) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretUpdateDockerhub) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretUpdateDockerhub) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretUpdateDockerhub) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretUpdateDockerhub) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretUpdateDockerhub) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateDockerhub) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateDockerhub) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateDockerhub) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretUpdateDockerhub) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretUpdateDockerhub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DockerhubPassword != nil {
		toSerialize["dockerhub-password"] = o.DockerhubPassword
	}
	if o.DockerhubTokenScopes != nil {
		toSerialize["dockerhub-token-scopes"] = o.DockerhubTokenScopes
	}
	if o.DockerhubUsername != nil {
		toSerialize["dockerhub-username"] = o.DockerhubUsername
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

type NullableDynamicSecretUpdateDockerhub struct {
	value *DynamicSecretUpdateDockerhub
	isSet bool
}

func (v NullableDynamicSecretUpdateDockerhub) Get() *DynamicSecretUpdateDockerhub {
	return v.value
}

func (v *NullableDynamicSecretUpdateDockerhub) Set(val *DynamicSecretUpdateDockerhub) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdateDockerhub) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdateDockerhub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdateDockerhub(val *DynamicSecretUpdateDockerhub) *NullableDynamicSecretUpdateDockerhub {
	return &NullableDynamicSecretUpdateDockerhub{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdateDockerhub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdateDockerhub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


