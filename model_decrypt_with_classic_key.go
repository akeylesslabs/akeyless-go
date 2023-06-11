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

// DecryptWithClassicKey struct for DecryptWithClassicKey
type DecryptWithClassicKey struct {
	// Ciphertext to be decrypted in base64 encoded format
	Ciphertext string `json:"ciphertext"`
	// The name of the key to use in the encryption process
	DisplayId string `json:"display-id"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// classic key version
	Version int32 `json:"version"`
}

// NewDecryptWithClassicKey instantiates a new DecryptWithClassicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecryptWithClassicKey(ciphertext string, displayId string, version int32) *DecryptWithClassicKey {
	this := DecryptWithClassicKey{}
	this.Ciphertext = ciphertext
	this.DisplayId = displayId
	var json bool = false
	this.Json = &json
	this.Version = version
	return &this
}

// NewDecryptWithClassicKeyWithDefaults instantiates a new DecryptWithClassicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptWithClassicKeyWithDefaults() *DecryptWithClassicKey {
	this := DecryptWithClassicKey{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetCiphertext returns the Ciphertext field value
func (o *DecryptWithClassicKey) GetCiphertext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ciphertext
}

// GetCiphertextOk returns a tuple with the Ciphertext field value
// and a boolean to check if the value has been set.
func (o *DecryptWithClassicKey) GetCiphertextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ciphertext, true
}

// SetCiphertext sets field value
func (o *DecryptWithClassicKey) SetCiphertext(v string) {
	o.Ciphertext = v
}

// GetDisplayId returns the DisplayId field value
func (o *DecryptWithClassicKey) GetDisplayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value
// and a boolean to check if the value has been set.
func (o *DecryptWithClassicKey) GetDisplayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayId, true
}

// SetDisplayId sets field value
func (o *DecryptWithClassicKey) SetDisplayId(v string) {
	o.DisplayId = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DecryptWithClassicKey) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptWithClassicKey) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DecryptWithClassicKey) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DecryptWithClassicKey) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DecryptWithClassicKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptWithClassicKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DecryptWithClassicKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DecryptWithClassicKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DecryptWithClassicKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptWithClassicKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DecryptWithClassicKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DecryptWithClassicKey) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value
func (o *DecryptWithClassicKey) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DecryptWithClassicKey) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *DecryptWithClassicKey) SetVersion(v int32) {
	o.Version = v
}

func (o DecryptWithClassicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ciphertext"] = o.Ciphertext
	}
	if true {
		toSerialize["display-id"] = o.DisplayId
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
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDecryptWithClassicKey struct {
	value *DecryptWithClassicKey
	isSet bool
}

func (v NullableDecryptWithClassicKey) Get() *DecryptWithClassicKey {
	return v.value
}

func (v *NullableDecryptWithClassicKey) Set(val *DecryptWithClassicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptWithClassicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptWithClassicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptWithClassicKey(val *DecryptWithClassicKey) *NullableDecryptWithClassicKey {
	return &NullableDecryptWithClassicKey{value: val, isSet: true}
}

func (v NullableDecryptWithClassicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptWithClassicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


