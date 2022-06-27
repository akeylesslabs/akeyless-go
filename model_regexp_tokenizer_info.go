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

// RegexpTokenizerInfo RegexpTokenizerInfo represents a general Regexp tokenization template
type RegexpTokenizerInfo struct {
	// The Alphabet used for the tokenization
	Alphabet *string `json:"alphabet,omitempty"`
	// Transformation to perform on the decrypted data
	DecryptionTemplate *string `json:"decryption_template,omitempty"`
	// Transformation to perform on the encrypted data, if the required output template doesn't match the input string The output Should still be valid for the Pattern, otherwise the secret would be able to be decrypted.
	EncryptionTemplate *string `json:"encryption_template,omitempty"`
	// Regexp pattern to extract and deposit the text/encdata
	Pattern *string `json:"pattern,omitempty"`
}

// NewRegexpTokenizerInfo instantiates a new RegexpTokenizerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegexpTokenizerInfo() *RegexpTokenizerInfo {
	this := RegexpTokenizerInfo{}
	return &this
}

// NewRegexpTokenizerInfoWithDefaults instantiates a new RegexpTokenizerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegexpTokenizerInfoWithDefaults() *RegexpTokenizerInfo {
	this := RegexpTokenizerInfo{}
	return &this
}

// GetAlphabet returns the Alphabet field value if set, zero value otherwise.
func (o *RegexpTokenizerInfo) GetAlphabet() string {
	if o == nil || o.Alphabet == nil {
		var ret string
		return ret
	}
	return *o.Alphabet
}

// GetAlphabetOk returns a tuple with the Alphabet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegexpTokenizerInfo) GetAlphabetOk() (*string, bool) {
	if o == nil || o.Alphabet == nil {
		return nil, false
	}
	return o.Alphabet, true
}

// HasAlphabet returns a boolean if a field has been set.
func (o *RegexpTokenizerInfo) HasAlphabet() bool {
	if o != nil && o.Alphabet != nil {
		return true
	}

	return false
}

// SetAlphabet gets a reference to the given string and assigns it to the Alphabet field.
func (o *RegexpTokenizerInfo) SetAlphabet(v string) {
	o.Alphabet = &v
}

// GetDecryptionTemplate returns the DecryptionTemplate field value if set, zero value otherwise.
func (o *RegexpTokenizerInfo) GetDecryptionTemplate() string {
	if o == nil || o.DecryptionTemplate == nil {
		var ret string
		return ret
	}
	return *o.DecryptionTemplate
}

// GetDecryptionTemplateOk returns a tuple with the DecryptionTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegexpTokenizerInfo) GetDecryptionTemplateOk() (*string, bool) {
	if o == nil || o.DecryptionTemplate == nil {
		return nil, false
	}
	return o.DecryptionTemplate, true
}

// HasDecryptionTemplate returns a boolean if a field has been set.
func (o *RegexpTokenizerInfo) HasDecryptionTemplate() bool {
	if o != nil && o.DecryptionTemplate != nil {
		return true
	}

	return false
}

// SetDecryptionTemplate gets a reference to the given string and assigns it to the DecryptionTemplate field.
func (o *RegexpTokenizerInfo) SetDecryptionTemplate(v string) {
	o.DecryptionTemplate = &v
}

// GetEncryptionTemplate returns the EncryptionTemplate field value if set, zero value otherwise.
func (o *RegexpTokenizerInfo) GetEncryptionTemplate() string {
	if o == nil || o.EncryptionTemplate == nil {
		var ret string
		return ret
	}
	return *o.EncryptionTemplate
}

// GetEncryptionTemplateOk returns a tuple with the EncryptionTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegexpTokenizerInfo) GetEncryptionTemplateOk() (*string, bool) {
	if o == nil || o.EncryptionTemplate == nil {
		return nil, false
	}
	return o.EncryptionTemplate, true
}

// HasEncryptionTemplate returns a boolean if a field has been set.
func (o *RegexpTokenizerInfo) HasEncryptionTemplate() bool {
	if o != nil && o.EncryptionTemplate != nil {
		return true
	}

	return false
}

// SetEncryptionTemplate gets a reference to the given string and assigns it to the EncryptionTemplate field.
func (o *RegexpTokenizerInfo) SetEncryptionTemplate(v string) {
	o.EncryptionTemplate = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *RegexpTokenizerInfo) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegexpTokenizerInfo) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *RegexpTokenizerInfo) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *RegexpTokenizerInfo) SetPattern(v string) {
	o.Pattern = &v
}

func (o RegexpTokenizerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alphabet != nil {
		toSerialize["alphabet"] = o.Alphabet
	}
	if o.DecryptionTemplate != nil {
		toSerialize["decryption_template"] = o.DecryptionTemplate
	}
	if o.EncryptionTemplate != nil {
		toSerialize["encryption_template"] = o.EncryptionTemplate
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	return json.Marshal(toSerialize)
}

type NullableRegexpTokenizerInfo struct {
	value *RegexpTokenizerInfo
	isSet bool
}

func (v NullableRegexpTokenizerInfo) Get() *RegexpTokenizerInfo {
	return v.value
}

func (v *NullableRegexpTokenizerInfo) Set(val *RegexpTokenizerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRegexpTokenizerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRegexpTokenizerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegexpTokenizerInfo(val *RegexpTokenizerInfo) *NullableRegexpTokenizerInfo {
	return &NullableRegexpTokenizerInfo{value: val, isSet: true}
}

func (v NullableRegexpTokenizerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegexpTokenizerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

