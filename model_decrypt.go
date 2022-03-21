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

// Decrypt decrypt is a command that decrypts ciphertext into plaintext by using an AES key.
type Decrypt struct {
	// Ciphertext to be decrypted in base64 encoded format
	Ciphertext *string `json:"ciphertext,omitempty"`
	// The display id of the key to use in the decryption process
	DisplayId *string `json:"display-id,omitempty"`
	// The encryption context. If this was specified in the encrypt command, it must be specified here or the decryption operation will fail
	EncryptionContext *map[string]string `json:"encryption-context,omitempty"`
	// The name of the key to use in the decryption process
	KeyName *string `json:"key-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDecrypt instantiates a new Decrypt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecrypt() *Decrypt {
	this := Decrypt{}
	return &this
}

// NewDecryptWithDefaults instantiates a new Decrypt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptWithDefaults() *Decrypt {
	this := Decrypt{}
	return &this
}

// GetCiphertext returns the Ciphertext field value if set, zero value otherwise.
func (o *Decrypt) GetCiphertext() string {
	if o == nil || o.Ciphertext == nil {
		var ret string
		return ret
	}
	return *o.Ciphertext
}

// GetCiphertextOk returns a tuple with the Ciphertext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Decrypt) GetCiphertextOk() (*string, bool) {
	if o == nil || o.Ciphertext == nil {
		return nil, false
	}
	return o.Ciphertext, true
}

// HasCiphertext returns a boolean if a field has been set.
func (o *Decrypt) HasCiphertext() bool {
	if o != nil && o.Ciphertext != nil {
		return true
	}

	return false
}

// SetCiphertext gets a reference to the given string and assigns it to the Ciphertext field.
func (o *Decrypt) SetCiphertext(v string) {
	o.Ciphertext = &v
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *Decrypt) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Decrypt) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *Decrypt) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *Decrypt) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetEncryptionContext returns the EncryptionContext field value if set, zero value otherwise.
func (o *Decrypt) GetEncryptionContext() map[string]string {
	if o == nil || o.EncryptionContext == nil {
		var ret map[string]string
		return ret
	}
	return *o.EncryptionContext
}

// GetEncryptionContextOk returns a tuple with the EncryptionContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Decrypt) GetEncryptionContextOk() (*map[string]string, bool) {
	if o == nil || o.EncryptionContext == nil {
		return nil, false
	}
	return o.EncryptionContext, true
}

// HasEncryptionContext returns a boolean if a field has been set.
func (o *Decrypt) HasEncryptionContext() bool {
	if o != nil && o.EncryptionContext != nil {
		return true
	}

	return false
}

// SetEncryptionContext gets a reference to the given map[string]string and assigns it to the EncryptionContext field.
func (o *Decrypt) SetEncryptionContext(v map[string]string) {
	o.EncryptionContext = &v
}

// GetKeyName returns the KeyName field value if set, zero value otherwise.
func (o *Decrypt) GetKeyName() string {
	if o == nil || o.KeyName == nil {
		var ret string
		return ret
	}
	return *o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Decrypt) GetKeyNameOk() (*string, bool) {
	if o == nil || o.KeyName == nil {
		return nil, false
	}
	return o.KeyName, true
}

// HasKeyName returns a boolean if a field has been set.
func (o *Decrypt) HasKeyName() bool {
	if o != nil && o.KeyName != nil {
		return true
	}

	return false
}

// SetKeyName gets a reference to the given string and assigns it to the KeyName field.
func (o *Decrypt) SetKeyName(v string) {
	o.KeyName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Decrypt) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Decrypt) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Decrypt) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Decrypt) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *Decrypt) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Decrypt) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *Decrypt) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *Decrypt) SetUidToken(v string) {
	o.UidToken = &v
}

func (o Decrypt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ciphertext != nil {
		toSerialize["ciphertext"] = o.Ciphertext
	}
	if o.DisplayId != nil {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.EncryptionContext != nil {
		toSerialize["encryption-context"] = o.EncryptionContext
	}
	if o.KeyName != nil {
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

type NullableDecrypt struct {
	value *Decrypt
	isSet bool
}

func (v NullableDecrypt) Get() *Decrypt {
	return v.value
}

func (v *NullableDecrypt) Set(val *Decrypt) {
	v.value = val
	v.isSet = true
}

func (v NullableDecrypt) IsSet() bool {
	return v.isSet
}

func (v *NullableDecrypt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecrypt(val *Decrypt) *NullableDecrypt {
	return &NullableDecrypt{value: val, isSet: true}
}

func (v NullableDecrypt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecrypt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


