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

// EncryptWithClassicKey struct for EncryptWithClassicKey
type EncryptWithClassicKey struct {
	// The name of the key to use in the encryption process
	DisplayId string `json:"display-id"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Data to be encrypted
	Plaintext string `json:"plaintext"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// classic key version
	Version int32 `json:"version"`
}

// NewEncryptWithClassicKey instantiates a new EncryptWithClassicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptWithClassicKey(displayId string, plaintext string, version int32, ) *EncryptWithClassicKey {
	this := EncryptWithClassicKey{}
	this.DisplayId = displayId
	var json bool = false
	this.Json = &json
	this.Plaintext = plaintext
	this.Version = version
	return &this
}

// NewEncryptWithClassicKeyWithDefaults instantiates a new EncryptWithClassicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptWithClassicKeyWithDefaults() *EncryptWithClassicKey {
	this := EncryptWithClassicKey{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetDisplayId returns the DisplayId field value
func (o *EncryptWithClassicKey) GetDisplayId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value
// and a boolean to check if the value has been set.
func (o *EncryptWithClassicKey) GetDisplayIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayId, true
}

// SetDisplayId sets field value
func (o *EncryptWithClassicKey) SetDisplayId(v string) {
	o.DisplayId = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *EncryptWithClassicKey) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptWithClassicKey) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *EncryptWithClassicKey) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *EncryptWithClassicKey) SetJson(v bool) {
	o.Json = &v
}

// GetPlaintext returns the Plaintext field value
func (o *EncryptWithClassicKey) GetPlaintext() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Plaintext
}

// GetPlaintextOk returns a tuple with the Plaintext field value
// and a boolean to check if the value has been set.
func (o *EncryptWithClassicKey) GetPlaintextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Plaintext, true
}

// SetPlaintext sets field value
func (o *EncryptWithClassicKey) SetPlaintext(v string) {
	o.Plaintext = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EncryptWithClassicKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptWithClassicKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EncryptWithClassicKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EncryptWithClassicKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EncryptWithClassicKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptWithClassicKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EncryptWithClassicKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EncryptWithClassicKey) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value
func (o *EncryptWithClassicKey) GetVersion() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EncryptWithClassicKey) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EncryptWithClassicKey) SetVersion(v int32) {
	o.Version = v
}

func (o EncryptWithClassicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["plaintext"] = o.Plaintext
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

type NullableEncryptWithClassicKey struct {
	value *EncryptWithClassicKey
	isSet bool
}

func (v NullableEncryptWithClassicKey) Get() *EncryptWithClassicKey {
	return v.value
}

func (v *NullableEncryptWithClassicKey) Set(val *EncryptWithClassicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptWithClassicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptWithClassicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptWithClassicKey(val *EncryptWithClassicKey) *NullableEncryptWithClassicKey {
	return &NullableEncryptWithClassicKey{value: val, isSet: true}
}

func (v NullableEncryptWithClassicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptWithClassicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


