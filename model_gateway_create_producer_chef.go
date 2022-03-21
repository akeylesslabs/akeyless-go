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

// GatewayCreateProducerChef gatewayCreateProducerChef is a command that creates chef producer
type GatewayCreateProducerChef struct {
	// Organizations
	ChefOrgs *string `json:"chef-orgs,omitempty"`
	// Server key
	ChefServerKey *string `json:"chef-server-key,omitempty"`
	// Server URL
	ChefServerUrl *string `json:"chef-server-url,omitempty"`
	// Server username
	ChefServerUsername *string `json:"chef-server-username,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Skip SSL
	SkipSsl *bool `json:"skip-ssl,omitempty"`
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

// NewGatewayCreateProducerChef instantiates a new GatewayCreateProducerChef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerChef(name string, ) *GatewayCreateProducerChef {
	this := GatewayCreateProducerChef{}
	this.Name = name
	var skipSsl bool = true
	this.SkipSsl = &skipSsl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerChefWithDefaults instantiates a new GatewayCreateProducerChef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerChefWithDefaults() *GatewayCreateProducerChef {
	this := GatewayCreateProducerChef{}
	var skipSsl bool = true
	this.SkipSsl = &skipSsl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetChefOrgs returns the ChefOrgs field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetChefOrgs() string {
	if o == nil || o.ChefOrgs == nil {
		var ret string
		return ret
	}
	return *o.ChefOrgs
}

// GetChefOrgsOk returns a tuple with the ChefOrgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetChefOrgsOk() (*string, bool) {
	if o == nil || o.ChefOrgs == nil {
		return nil, false
	}
	return o.ChefOrgs, true
}

// HasChefOrgs returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasChefOrgs() bool {
	if o != nil && o.ChefOrgs != nil {
		return true
	}

	return false
}

// SetChefOrgs gets a reference to the given string and assigns it to the ChefOrgs field.
func (o *GatewayCreateProducerChef) SetChefOrgs(v string) {
	o.ChefOrgs = &v
}

// GetChefServerKey returns the ChefServerKey field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetChefServerKey() string {
	if o == nil || o.ChefServerKey == nil {
		var ret string
		return ret
	}
	return *o.ChefServerKey
}

// GetChefServerKeyOk returns a tuple with the ChefServerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetChefServerKeyOk() (*string, bool) {
	if o == nil || o.ChefServerKey == nil {
		return nil, false
	}
	return o.ChefServerKey, true
}

// HasChefServerKey returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasChefServerKey() bool {
	if o != nil && o.ChefServerKey != nil {
		return true
	}

	return false
}

// SetChefServerKey gets a reference to the given string and assigns it to the ChefServerKey field.
func (o *GatewayCreateProducerChef) SetChefServerKey(v string) {
	o.ChefServerKey = &v
}

// GetChefServerUrl returns the ChefServerUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetChefServerUrl() string {
	if o == nil || o.ChefServerUrl == nil {
		var ret string
		return ret
	}
	return *o.ChefServerUrl
}

// GetChefServerUrlOk returns a tuple with the ChefServerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetChefServerUrlOk() (*string, bool) {
	if o == nil || o.ChefServerUrl == nil {
		return nil, false
	}
	return o.ChefServerUrl, true
}

// HasChefServerUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasChefServerUrl() bool {
	if o != nil && o.ChefServerUrl != nil {
		return true
	}

	return false
}

// SetChefServerUrl gets a reference to the given string and assigns it to the ChefServerUrl field.
func (o *GatewayCreateProducerChef) SetChefServerUrl(v string) {
	o.ChefServerUrl = &v
}

// GetChefServerUsername returns the ChefServerUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetChefServerUsername() string {
	if o == nil || o.ChefServerUsername == nil {
		var ret string
		return ret
	}
	return *o.ChefServerUsername
}

// GetChefServerUsernameOk returns a tuple with the ChefServerUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetChefServerUsernameOk() (*string, bool) {
	if o == nil || o.ChefServerUsername == nil {
		return nil, false
	}
	return o.ChefServerUsername, true
}

// HasChefServerUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasChefServerUsername() bool {
	if o != nil && o.ChefServerUsername != nil {
		return true
	}

	return false
}

// SetChefServerUsername gets a reference to the given string and assigns it to the ChefServerUsername field.
func (o *GatewayCreateProducerChef) SetChefServerUsername(v string) {
	o.ChefServerUsername = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerChef) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerChef) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerChef) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSkipSsl returns the SkipSsl field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetSkipSsl() bool {
	if o == nil || o.SkipSsl == nil {
		var ret bool
		return ret
	}
	return *o.SkipSsl
}

// GetSkipSslOk returns a tuple with the SkipSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetSkipSslOk() (*bool, bool) {
	if o == nil || o.SkipSsl == nil {
		return nil, false
	}
	return o.SkipSsl, true
}

// HasSkipSsl returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasSkipSsl() bool {
	if o != nil && o.SkipSsl != nil {
		return true
	}

	return false
}

// SetSkipSsl gets a reference to the given bool and assigns it to the SkipSsl field.
func (o *GatewayCreateProducerChef) SetSkipSsl(v bool) {
	o.SkipSsl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerChef) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerChef) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerChef) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerChef) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerChef) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerChef) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerChef) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerChef) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerChef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChefOrgs != nil {
		toSerialize["chef-orgs"] = o.ChefOrgs
	}
	if o.ChefServerKey != nil {
		toSerialize["chef-server-key"] = o.ChefServerKey
	}
	if o.ChefServerUrl != nil {
		toSerialize["chef-server-url"] = o.ChefServerUrl
	}
	if o.ChefServerUsername != nil {
		toSerialize["chef-server-username"] = o.ChefServerUsername
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.SkipSsl != nil {
		toSerialize["skip-ssl"] = o.SkipSsl
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

type NullableGatewayCreateProducerChef struct {
	value *GatewayCreateProducerChef
	isSet bool
}

func (v NullableGatewayCreateProducerChef) Get() *GatewayCreateProducerChef {
	return v.value
}

func (v *NullableGatewayCreateProducerChef) Set(val *GatewayCreateProducerChef) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerChef) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerChef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerChef(val *GatewayCreateProducerChef) *NullableGatewayCreateProducerChef {
	return &NullableGatewayCreateProducerChef{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerChef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerChef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


