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

// UidGenerateToken uidGenerateToken is a command that generates a new token using Akeyless Universal Identity.
type UidGenerateToken struct {
	// The universal identity auth method name
	AuthMethodName string `json:"auth-method-name"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUidGenerateToken instantiates a new UidGenerateToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUidGenerateToken(authMethodName string, ) *UidGenerateToken {
	this := UidGenerateToken{}
	this.AuthMethodName = authMethodName
	var json bool = false
	this.Json = &json
	return &this
}

// NewUidGenerateTokenWithDefaults instantiates a new UidGenerateToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUidGenerateTokenWithDefaults() *UidGenerateToken {
	this := UidGenerateToken{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAuthMethodName returns the AuthMethodName field value
func (o *UidGenerateToken) GetAuthMethodName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AuthMethodName
}

// GetAuthMethodNameOk returns a tuple with the AuthMethodName field value
// and a boolean to check if the value has been set.
func (o *UidGenerateToken) GetAuthMethodNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthMethodName, true
}

// SetAuthMethodName sets field value
func (o *UidGenerateToken) SetAuthMethodName(v string) {
	o.AuthMethodName = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UidGenerateToken) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UidGenerateToken) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UidGenerateToken) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UidGenerateToken) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UidGenerateToken) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UidGenerateToken) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UidGenerateToken) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UidGenerateToken) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UidGenerateToken) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UidGenerateToken) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UidGenerateToken) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UidGenerateToken) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UidGenerateToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auth-method-name"] = o.AuthMethodName
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableUidGenerateToken struct {
	value *UidGenerateToken
	isSet bool
}

func (v NullableUidGenerateToken) Get() *UidGenerateToken {
	return v.value
}

func (v *NullableUidGenerateToken) Set(val *UidGenerateToken) {
	v.value = val
	v.isSet = true
}

func (v NullableUidGenerateToken) IsSet() bool {
	return v.isSet
}

func (v *NullableUidGenerateToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUidGenerateToken(val *UidGenerateToken) *NullableUidGenerateToken {
	return &NullableUidGenerateToken{value: val, isSet: true}
}

func (v NullableUidGenerateToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUidGenerateToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


