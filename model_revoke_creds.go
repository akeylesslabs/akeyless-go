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

// RevokeCreds revokeCreds will permanently revoke the credentials associated with the provided token or profile.
type RevokeCreds struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewRevokeCreds instantiates a new RevokeCreds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokeCreds() *RevokeCreds {
	this := RevokeCreds{}
	var json bool = false
	this.Json = &json
	return &this
}

// NewRevokeCredsWithDefaults instantiates a new RevokeCreds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeCredsWithDefaults() *RevokeCreds {
	this := RevokeCreds{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *RevokeCreds) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokeCreds) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *RevokeCreds) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *RevokeCreds) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RevokeCreds) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokeCreds) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RevokeCreds) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RevokeCreds) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *RevokeCreds) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokeCreds) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *RevokeCreds) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *RevokeCreds) SetUidToken(v string) {
	o.UidToken = &v
}

func (o RevokeCreds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableRevokeCreds struct {
	value *RevokeCreds
	isSet bool
}

func (v NullableRevokeCreds) Get() *RevokeCreds {
	return v.value
}

func (v *NullableRevokeCreds) Set(val *RevokeCreds) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokeCreds) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokeCreds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokeCreds(val *RevokeCreds) *NullableRevokeCreds {
	return &NullableRevokeCreds{value: val, isSet: true}
}

func (v NullableRevokeCreds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokeCreds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


