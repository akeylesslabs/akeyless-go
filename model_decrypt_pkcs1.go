/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 0.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// DecryptPKCS1 struct for DecryptPKCS1
type DecryptPKCS1 struct {
	// Ciphertext to be decrypted in base64 encoded format
	Ciphertext string `json:"ciphertext"`
	// The name of the RSA key to use in the decryption process
	KeyName string `json:"key-name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDecryptPKCS1 instantiates a new DecryptPKCS1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecryptPKCS1(ciphertext string, keyName string, ) *DecryptPKCS1 {
	this := DecryptPKCS1{}
	this.Ciphertext = ciphertext
	this.KeyName = keyName
	return &this
}

// NewDecryptPKCS1WithDefaults instantiates a new DecryptPKCS1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptPKCS1WithDefaults() *DecryptPKCS1 {
	this := DecryptPKCS1{}
	return &this
}

// GetCiphertext returns the Ciphertext field value
func (o *DecryptPKCS1) GetCiphertext() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Ciphertext
}

// GetCiphertextOk returns a tuple with the Ciphertext field value
// and a boolean to check if the value has been set.
func (o *DecryptPKCS1) GetCiphertextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ciphertext, true
}

// SetCiphertext sets field value
func (o *DecryptPKCS1) SetCiphertext(v string) {
	o.Ciphertext = v
}

// GetKeyName returns the KeyName field value
func (o *DecryptPKCS1) GetKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *DecryptPKCS1) GetKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *DecryptPKCS1) SetKeyName(v string) {
	o.KeyName = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DecryptPKCS1) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptPKCS1) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DecryptPKCS1) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DecryptPKCS1) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DecryptPKCS1) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptPKCS1) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DecryptPKCS1) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DecryptPKCS1) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DecryptPKCS1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ciphertext"] = o.Ciphertext
	}
	if true {
		toSerialize["key-name"] = o.KeyName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDecryptPKCS1 struct {
	value *DecryptPKCS1
	isSet bool
}

func (v NullableDecryptPKCS1) Get() *DecryptPKCS1 {
	return v.value
}

func (v *NullableDecryptPKCS1) Set(val *DecryptPKCS1) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptPKCS1) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptPKCS1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptPKCS1(val *DecryptPKCS1) *NullableDecryptPKCS1 {
	return &NullableDecryptPKCS1{value: val, isSet: true}
}

func (v NullableDecryptPKCS1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptPKCS1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


