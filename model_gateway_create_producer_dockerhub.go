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

// checks if the GatewayCreateProducerDockerhub type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayCreateProducerDockerhub{}

// GatewayCreateProducerDockerhub gatewayCreateProducerDockerhub is a command that creates a DOCKERHUB producer
type GatewayCreateProducerDockerhub struct {
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
	// Producer name
	Name string `json:"name"`
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

// NewGatewayCreateProducerDockerhub instantiates a new GatewayCreateProducerDockerhub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerDockerhub(name string) *GatewayCreateProducerDockerhub {
	this := GatewayCreateProducerDockerhub{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerDockerhubWithDefaults instantiates a new GatewayCreateProducerDockerhub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerDockerhubWithDefaults() *GatewayCreateProducerDockerhub {
	this := GatewayCreateProducerDockerhub{}
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetDeleteProtection() string {
	if o == nil || IsNil(o.DeleteProtection) {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || IsNil(o.DeleteProtection) {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasDeleteProtection() bool {
	if o != nil && !IsNil(o.DeleteProtection) {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerDockerhub) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDockerhubPassword returns the DockerhubPassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetDockerhubPassword() string {
	if o == nil || IsNil(o.DockerhubPassword) {
		var ret string
		return ret
	}
	return *o.DockerhubPassword
}

// GetDockerhubPasswordOk returns a tuple with the DockerhubPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetDockerhubPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DockerhubPassword) {
		return nil, false
	}
	return o.DockerhubPassword, true
}

// HasDockerhubPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasDockerhubPassword() bool {
	if o != nil && !IsNil(o.DockerhubPassword) {
		return true
	}

	return false
}

// SetDockerhubPassword gets a reference to the given string and assigns it to the DockerhubPassword field.
func (o *GatewayCreateProducerDockerhub) SetDockerhubPassword(v string) {
	o.DockerhubPassword = &v
}

// GetDockerhubTokenScopes returns the DockerhubTokenScopes field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetDockerhubTokenScopes() string {
	if o == nil || IsNil(o.DockerhubTokenScopes) {
		var ret string
		return ret
	}
	return *o.DockerhubTokenScopes
}

// GetDockerhubTokenScopesOk returns a tuple with the DockerhubTokenScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetDockerhubTokenScopesOk() (*string, bool) {
	if o == nil || IsNil(o.DockerhubTokenScopes) {
		return nil, false
	}
	return o.DockerhubTokenScopes, true
}

// HasDockerhubTokenScopes returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasDockerhubTokenScopes() bool {
	if o != nil && !IsNil(o.DockerhubTokenScopes) {
		return true
	}

	return false
}

// SetDockerhubTokenScopes gets a reference to the given string and assigns it to the DockerhubTokenScopes field.
func (o *GatewayCreateProducerDockerhub) SetDockerhubTokenScopes(v string) {
	o.DockerhubTokenScopes = &v
}

// GetDockerhubUsername returns the DockerhubUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetDockerhubUsername() string {
	if o == nil || IsNil(o.DockerhubUsername) {
		var ret string
		return ret
	}
	return *o.DockerhubUsername
}

// GetDockerhubUsernameOk returns a tuple with the DockerhubUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetDockerhubUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.DockerhubUsername) {
		return nil, false
	}
	return o.DockerhubUsername, true
}

// HasDockerhubUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasDockerhubUsername() bool {
	if o != nil && !IsNil(o.DockerhubUsername) {
		return true
	}

	return false
}

// SetDockerhubUsername gets a reference to the given string and assigns it to the DockerhubUsername field.
func (o *GatewayCreateProducerDockerhub) SetDockerhubUsername(v string) {
	o.DockerhubUsername = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerDockerhub) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerDockerhub) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerDockerhub) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetProducerEncryptionKeyName() string {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasProducerEncryptionKeyName() bool {
	if o != nil && !IsNil(o.ProducerEncryptionKeyName) {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerDockerhub) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerDockerhub) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerDockerhub) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerDockerhub) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerDockerhub) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerDockerhub) GetUserTtl() string {
	if o == nil || IsNil(o.UserTtl) {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerDockerhub) GetUserTtlOk() (*string, bool) {
	if o == nil || IsNil(o.UserTtl) {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerDockerhub) HasUserTtl() bool {
	if o != nil && !IsNil(o.UserTtl) {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerDockerhub) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerDockerhub) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayCreateProducerDockerhub) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteProtection) {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if !IsNil(o.DockerhubPassword) {
		toSerialize["dockerhub-password"] = o.DockerhubPassword
	}
	if !IsNil(o.DockerhubTokenScopes) {
		toSerialize["dockerhub-token-scopes"] = o.DockerhubTokenScopes
	}
	if !IsNil(o.DockerhubUsername) {
		toSerialize["dockerhub-username"] = o.DockerhubUsername
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["name"] = o.Name
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

type NullableGatewayCreateProducerDockerhub struct {
	value *GatewayCreateProducerDockerhub
	isSet bool
}

func (v NullableGatewayCreateProducerDockerhub) Get() *GatewayCreateProducerDockerhub {
	return v.value
}

func (v *NullableGatewayCreateProducerDockerhub) Set(val *GatewayCreateProducerDockerhub) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerDockerhub) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerDockerhub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerDockerhub(val *GatewayCreateProducerDockerhub) *NullableGatewayCreateProducerDockerhub {
	return &NullableGatewayCreateProducerDockerhub{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerDockerhub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerDockerhub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


