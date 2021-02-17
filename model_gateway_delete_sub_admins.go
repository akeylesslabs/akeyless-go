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

// GatewayDeleteSubAdmins gatewayDeleteSubAdmins is a command that deletes sub-admins
type GatewayDeleteSubAdmins struct {
	// Gateway url
	GatewayUrl *string `json:"gateway-url,omitempty"`
	// SubAdmins to be removed
	SubAdmin []string `json:"sub-admin"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayDeleteSubAdmins instantiates a new GatewayDeleteSubAdmins object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayDeleteSubAdmins(subAdmin []string, ) *GatewayDeleteSubAdmins {
	this := GatewayDeleteSubAdmins{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	this.SubAdmin = subAdmin
	return &this
}

// NewGatewayDeleteSubAdminsWithDefaults instantiates a new GatewayDeleteSubAdmins object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayDeleteSubAdminsWithDefaults() *GatewayDeleteSubAdmins {
	this := GatewayDeleteSubAdmins{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	return &this
}

// GetGatewayUrl returns the GatewayUrl field value if set, zero value otherwise.
func (o *GatewayDeleteSubAdmins) GetGatewayUrl() string {
	if o == nil || o.GatewayUrl == nil {
		var ret string
		return ret
	}
	return *o.GatewayUrl
}

// GetGatewayUrlOk returns a tuple with the GatewayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteSubAdmins) GetGatewayUrlOk() (*string, bool) {
	if o == nil || o.GatewayUrl == nil {
		return nil, false
	}
	return o.GatewayUrl, true
}

// HasGatewayUrl returns a boolean if a field has been set.
func (o *GatewayDeleteSubAdmins) HasGatewayUrl() bool {
	if o != nil && o.GatewayUrl != nil {
		return true
	}

	return false
}

// SetGatewayUrl gets a reference to the given string and assigns it to the GatewayUrl field.
func (o *GatewayDeleteSubAdmins) SetGatewayUrl(v string) {
	o.GatewayUrl = &v
}

// GetSubAdmin returns the SubAdmin field value
func (o *GatewayDeleteSubAdmins) GetSubAdmin() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.SubAdmin
}

// GetSubAdminOk returns a tuple with the SubAdmin field value
// and a boolean to check if the value has been set.
func (o *GatewayDeleteSubAdmins) GetSubAdminOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubAdmin, true
}

// SetSubAdmin sets field value
func (o *GatewayDeleteSubAdmins) SetSubAdmin(v []string) {
	o.SubAdmin = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayDeleteSubAdmins) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteSubAdmins) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayDeleteSubAdmins) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayDeleteSubAdmins) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayDeleteSubAdmins) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteSubAdmins) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayDeleteSubAdmins) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayDeleteSubAdmins) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayDeleteSubAdmins) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayUrl != nil {
		toSerialize["gateway-url"] = o.GatewayUrl
	}
	if true {
		toSerialize["sub-admin"] = o.SubAdmin
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayDeleteSubAdmins struct {
	value *GatewayDeleteSubAdmins
	isSet bool
}

func (v NullableGatewayDeleteSubAdmins) Get() *GatewayDeleteSubAdmins {
	return v.value
}

func (v *NullableGatewayDeleteSubAdmins) Set(val *GatewayDeleteSubAdmins) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayDeleteSubAdmins) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayDeleteSubAdmins) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayDeleteSubAdmins(val *GatewayDeleteSubAdmins) *NullableGatewayDeleteSubAdmins {
	return &NullableGatewayDeleteSubAdmins{value: val, isSet: true}
}

func (v NullableGatewayDeleteSubAdmins) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayDeleteSubAdmins) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

