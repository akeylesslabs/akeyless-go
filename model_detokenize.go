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

// Detokenize detokenize is a command that decrypts text with a tokenizer
type Detokenize struct {
	// Data to be decrypted
	Ciphertext string `json:"ciphertext"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The name of the tokenizer to use in the decryption process
	TokenizerName string `json:"tokenizer-name"`
	// Base64 encoded tweak for vaultless encryption
	Tweak *string `json:"tweak,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDetokenize instantiates a new Detokenize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetokenize(ciphertext string, tokenizerName string, ) *Detokenize {
	this := Detokenize{}
	this.Ciphertext = ciphertext
	this.TokenizerName = tokenizerName
	return &this
}

// NewDetokenizeWithDefaults instantiates a new Detokenize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetokenizeWithDefaults() *Detokenize {
	this := Detokenize{}
	return &this
}

// GetCiphertext returns the Ciphertext field value
func (o *Detokenize) GetCiphertext() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Ciphertext
}

// GetCiphertextOk returns a tuple with the Ciphertext field value
// and a boolean to check if the value has been set.
func (o *Detokenize) GetCiphertextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ciphertext, true
}

// SetCiphertext sets field value
func (o *Detokenize) SetCiphertext(v string) {
	o.Ciphertext = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Detokenize) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detokenize) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Detokenize) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Detokenize) SetToken(v string) {
	o.Token = &v
}

// GetTokenizerName returns the TokenizerName field value
func (o *Detokenize) GetTokenizerName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TokenizerName
}

// GetTokenizerNameOk returns a tuple with the TokenizerName field value
// and a boolean to check if the value has been set.
func (o *Detokenize) GetTokenizerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenizerName, true
}

// SetTokenizerName sets field value
func (o *Detokenize) SetTokenizerName(v string) {
	o.TokenizerName = v
}

// GetTweak returns the Tweak field value if set, zero value otherwise.
func (o *Detokenize) GetTweak() string {
	if o == nil || o.Tweak == nil {
		var ret string
		return ret
	}
	return *o.Tweak
}

// GetTweakOk returns a tuple with the Tweak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detokenize) GetTweakOk() (*string, bool) {
	if o == nil || o.Tweak == nil {
		return nil, false
	}
	return o.Tweak, true
}

// HasTweak returns a boolean if a field has been set.
func (o *Detokenize) HasTweak() bool {
	if o != nil && o.Tweak != nil {
		return true
	}

	return false
}

// SetTweak gets a reference to the given string and assigns it to the Tweak field.
func (o *Detokenize) SetTweak(v string) {
	o.Tweak = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *Detokenize) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detokenize) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *Detokenize) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *Detokenize) SetUidToken(v string) {
	o.UidToken = &v
}

func (o Detokenize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ciphertext"] = o.Ciphertext
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["tokenizer-name"] = o.TokenizerName
	}
	if o.Tweak != nil {
		toSerialize["tweak"] = o.Tweak
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDetokenize struct {
	value *Detokenize
	isSet bool
}

func (v NullableDetokenize) Get() *Detokenize {
	return v.value
}

func (v *NullableDetokenize) Set(val *Detokenize) {
	v.value = val
	v.isSet = true
}

func (v NullableDetokenize) IsSet() bool {
	return v.isSet
}

func (v *NullableDetokenize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetokenize(val *Detokenize) *NullableDetokenize {
	return &NullableDetokenize{value: val, isSet: true}
}

func (v NullableDetokenize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetokenize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

