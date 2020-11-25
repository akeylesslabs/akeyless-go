/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// CreateAuthMethodLDAP createAuthMethodLDAP is a command that creates a new auth method that will be able to authenticate using LDAP.
type CreateAuthMethodLDAP struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// A CIDR whitelist of the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// A public key generated for LDAP authentication method on Akeyless [RSA2048]
	PublicKeyFilePath string `json:"public-key-file-path"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateAuthMethodLDAP instantiates a new CreateAuthMethodLDAP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodLDAP(name string, publicKeyFilePath string, ) *CreateAuthMethodLDAP {
	this := CreateAuthMethodLDAP{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	this.Name = name
	this.PublicKeyFilePath = publicKeyFilePath
	return &this
}

// NewCreateAuthMethodLDAPWithDefaults instantiates a new CreateAuthMethodLDAP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodLDAPWithDefaults() *CreateAuthMethodLDAP {
	this := CreateAuthMethodLDAP{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *CreateAuthMethodLDAP) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *CreateAuthMethodLDAP) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetName returns the Name field value
func (o *CreateAuthMethodLDAP) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAuthMethodLDAP) SetName(v string) {
	o.Name = v
}

// GetPublicKeyFilePath returns the PublicKeyFilePath field value
func (o *CreateAuthMethodLDAP) GetPublicKeyFilePath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PublicKeyFilePath
}

// GetPublicKeyFilePathOk returns a tuple with the PublicKeyFilePath field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetPublicKeyFilePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicKeyFilePath, true
}

// SetPublicKeyFilePath sets field value
func (o *CreateAuthMethodLDAP) SetPublicKeyFilePath(v string) {
	o.PublicKeyFilePath = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateAuthMethodLDAP) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateAuthMethodLDAP) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateAuthMethodLDAP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["public-key-file-path"] = o.PublicKeyFilePath
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAuthMethodLDAP struct {
	value *CreateAuthMethodLDAP
	isSet bool
}

func (v NullableCreateAuthMethodLDAP) Get() *CreateAuthMethodLDAP {
	return v.value
}

func (v *NullableCreateAuthMethodLDAP) Set(val *CreateAuthMethodLDAP) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodLDAP) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodLDAP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodLDAP(val *CreateAuthMethodLDAP) *NullableCreateAuthMethodLDAP {
	return &NullableCreateAuthMethodLDAP{value: val, isSet: true}
}

func (v NullableCreateAuthMethodLDAP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodLDAP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


