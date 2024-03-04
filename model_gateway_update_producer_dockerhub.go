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

// GatewayUpdateProducerDockerhub gatewayUpdateProducerDockerhub is a command that updates a DOCKERHUB producer [Deprecated: Use dynamic-secret-update-dockerhub command]
type GatewayUpdateProducerDockerhub struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
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
	// Dynamic secret name
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

// NewGatewayUpdateProducerDockerhub instantiates a new GatewayUpdateProducerDockerhub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerDockerhub(name string, ) *GatewayUpdateProducerDockerhub {
	this := GatewayUpdateProducerDockerhub{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerDockerhubWithDefaults instantiates a new GatewayUpdateProducerDockerhub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerDockerhubWithDefaults() *GatewayUpdateProducerDockerhub {
	this := GatewayUpdateProducerDockerhub{}
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerDockerhub) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDockerhubPassword returns the DockerhubPassword field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetDockerhubPassword() string {
	if o == nil || o.DockerhubPassword == nil {
		var ret string
		return ret
	}
	return *o.DockerhubPassword
}

// GetDockerhubPasswordOk returns a tuple with the DockerhubPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetDockerhubPasswordOk() (*string, bool) {
	if o == nil || o.DockerhubPassword == nil {
		return nil, false
	}
	return o.DockerhubPassword, true
}

// HasDockerhubPassword returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasDockerhubPassword() bool {
	if o != nil && o.DockerhubPassword != nil {
		return true
	}

	return false
}

// SetDockerhubPassword gets a reference to the given string and assigns it to the DockerhubPassword field.
func (o *GatewayUpdateProducerDockerhub) SetDockerhubPassword(v string) {
	o.DockerhubPassword = &v
}

// GetDockerhubTokenScopes returns the DockerhubTokenScopes field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetDockerhubTokenScopes() string {
	if o == nil || o.DockerhubTokenScopes == nil {
		var ret string
		return ret
	}
	return *o.DockerhubTokenScopes
}

// GetDockerhubTokenScopesOk returns a tuple with the DockerhubTokenScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetDockerhubTokenScopesOk() (*string, bool) {
	if o == nil || o.DockerhubTokenScopes == nil {
		return nil, false
	}
	return o.DockerhubTokenScopes, true
}

// HasDockerhubTokenScopes returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasDockerhubTokenScopes() bool {
	if o != nil && o.DockerhubTokenScopes != nil {
		return true
	}

	return false
}

// SetDockerhubTokenScopes gets a reference to the given string and assigns it to the DockerhubTokenScopes field.
func (o *GatewayUpdateProducerDockerhub) SetDockerhubTokenScopes(v string) {
	o.DockerhubTokenScopes = &v
}

// GetDockerhubUsername returns the DockerhubUsername field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetDockerhubUsername() string {
	if o == nil || o.DockerhubUsername == nil {
		var ret string
		return ret
	}
	return *o.DockerhubUsername
}

// GetDockerhubUsernameOk returns a tuple with the DockerhubUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetDockerhubUsernameOk() (*string, bool) {
	if o == nil || o.DockerhubUsername == nil {
		return nil, false
	}
	return o.DockerhubUsername, true
}

// HasDockerhubUsername returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasDockerhubUsername() bool {
	if o != nil && o.DockerhubUsername != nil {
		return true
	}

	return false
}

// SetDockerhubUsername gets a reference to the given string and assigns it to the DockerhubUsername field.
func (o *GatewayUpdateProducerDockerhub) SetDockerhubUsername(v string) {
	o.DockerhubUsername = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerDockerhub) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerDockerhub) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerDockerhub) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerDockerhub) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerDockerhub) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerDockerhub) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerDockerhub) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerDockerhub) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerDockerhub) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerDockerhub) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerDockerhub) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerDockerhub) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerDockerhub) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerDockerhub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
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

type NullableGatewayUpdateProducerDockerhub struct {
	value *GatewayUpdateProducerDockerhub
	isSet bool
}

func (v NullableGatewayUpdateProducerDockerhub) Get() *GatewayUpdateProducerDockerhub {
	return v.value
}

func (v *NullableGatewayUpdateProducerDockerhub) Set(val *GatewayUpdateProducerDockerhub) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerDockerhub) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerDockerhub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerDockerhub(val *GatewayUpdateProducerDockerhub) *NullableGatewayUpdateProducerDockerhub {
	return &NullableGatewayUpdateProducerDockerhub{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerDockerhub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerDockerhub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


