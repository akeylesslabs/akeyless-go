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

// VerifyPKICertWithClassicKey struct for VerifyPKICertWithClassicKey
type VerifyPKICertWithClassicKey struct {
	// The name of the key to use in the verify PKICert process
	DisplayId string `json:"display-id"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// PkiCert
	PkiCert string `json:"pki-cert"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// classic key version
	Version int32 `json:"version"`
}

// NewVerifyPKICertWithClassicKey instantiates a new VerifyPKICertWithClassicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyPKICertWithClassicKey(displayId string, pkiCert string, version int32, ) *VerifyPKICertWithClassicKey {
	this := VerifyPKICertWithClassicKey{}
	this.DisplayId = displayId
	var json bool = false
	this.Json = &json
	this.PkiCert = pkiCert
	this.Version = version
	return &this
}

// NewVerifyPKICertWithClassicKeyWithDefaults instantiates a new VerifyPKICertWithClassicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyPKICertWithClassicKeyWithDefaults() *VerifyPKICertWithClassicKey {
	this := VerifyPKICertWithClassicKey{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetDisplayId returns the DisplayId field value
func (o *VerifyPKICertWithClassicKey) GetDisplayId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value
// and a boolean to check if the value has been set.
func (o *VerifyPKICertWithClassicKey) GetDisplayIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayId, true
}

// SetDisplayId sets field value
func (o *VerifyPKICertWithClassicKey) SetDisplayId(v string) {
	o.DisplayId = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *VerifyPKICertWithClassicKey) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKICertWithClassicKey) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *VerifyPKICertWithClassicKey) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *VerifyPKICertWithClassicKey) SetJson(v bool) {
	o.Json = &v
}

// GetPkiCert returns the PkiCert field value
func (o *VerifyPKICertWithClassicKey) GetPkiCert() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PkiCert
}

// GetPkiCertOk returns a tuple with the PkiCert field value
// and a boolean to check if the value has been set.
func (o *VerifyPKICertWithClassicKey) GetPkiCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiCert, true
}

// SetPkiCert sets field value
func (o *VerifyPKICertWithClassicKey) SetPkiCert(v string) {
	o.PkiCert = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *VerifyPKICertWithClassicKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKICertWithClassicKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *VerifyPKICertWithClassicKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *VerifyPKICertWithClassicKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *VerifyPKICertWithClassicKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKICertWithClassicKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *VerifyPKICertWithClassicKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *VerifyPKICertWithClassicKey) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value
func (o *VerifyPKICertWithClassicKey) GetVersion() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VerifyPKICertWithClassicKey) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VerifyPKICertWithClassicKey) SetVersion(v int32) {
	o.Version = v
}

func (o VerifyPKICertWithClassicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["pki-cert"] = o.PkiCert
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyPKICertWithClassicKey struct {
	value *VerifyPKICertWithClassicKey
	isSet bool
}

func (v NullableVerifyPKICertWithClassicKey) Get() *VerifyPKICertWithClassicKey {
	return v.value
}

func (v *NullableVerifyPKICertWithClassicKey) Set(val *VerifyPKICertWithClassicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyPKICertWithClassicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyPKICertWithClassicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyPKICertWithClassicKey(val *VerifyPKICertWithClassicKey) *NullableVerifyPKICertWithClassicKey {
	return &NullableVerifyPKICertWithClassicKey{value: val, isSet: true}
}

func (v NullableVerifyPKICertWithClassicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyPKICertWithClassicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


