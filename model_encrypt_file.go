/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// EncryptFile struct for EncryptFile
type EncryptFile struct {
	// name-value pair that specifies the encryption context to be used for authenticated encryption. If used here, the same value must be supplied to the decrypt command or decryption will fail
	EncryptionContext *map[string]string `json:"encryption-context,omitempty"`
	// Path to the file to be encrypted. If not provided, the content will be taken from stdin
	In string `json:"in"`
	// The name of the key to use in the encryption process
	KeyName string `json:"key-name"`
	// Path to the output file. If not provided, the output will be sent to stdout
	Out *string `json:"out,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewEncryptFile instantiates a new EncryptFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptFile(in string, keyName string, ) *EncryptFile {
	this := EncryptFile{}
	this.In = in
	this.KeyName = keyName
	return &this
}

// NewEncryptFileWithDefaults instantiates a new EncryptFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptFileWithDefaults() *EncryptFile {
	this := EncryptFile{}
	return &this
}

// GetEncryptionContext returns the EncryptionContext field value if set, zero value otherwise.
func (o *EncryptFile) GetEncryptionContext() map[string]string {
	if o == nil || o.EncryptionContext == nil {
		var ret map[string]string
		return ret
	}
	return *o.EncryptionContext
}

// GetEncryptionContextOk returns a tuple with the EncryptionContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetEncryptionContextOk() (*map[string]string, bool) {
	if o == nil || o.EncryptionContext == nil {
		return nil, false
	}
	return o.EncryptionContext, true
}

// HasEncryptionContext returns a boolean if a field has been set.
func (o *EncryptFile) HasEncryptionContext() bool {
	if o != nil && o.EncryptionContext != nil {
		return true
	}

	return false
}

// SetEncryptionContext gets a reference to the given map[string]string and assigns it to the EncryptionContext field.
func (o *EncryptFile) SetEncryptionContext(v map[string]string) {
	o.EncryptionContext = &v
}

// GetIn returns the In field value
func (o *EncryptFile) GetIn() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.In
}

// GetInOk returns a tuple with the In field value
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetInOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.In, true
}

// SetIn sets field value
func (o *EncryptFile) SetIn(v string) {
	o.In = v
}

// GetKeyName returns the KeyName field value
func (o *EncryptFile) GetKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *EncryptFile) SetKeyName(v string) {
	o.KeyName = v
}

// GetOut returns the Out field value if set, zero value otherwise.
func (o *EncryptFile) GetOut() string {
	if o == nil || o.Out == nil {
		var ret string
		return ret
	}
	return *o.Out
}

// GetOutOk returns a tuple with the Out field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetOutOk() (*string, bool) {
	if o == nil || o.Out == nil {
		return nil, false
	}
	return o.Out, true
}

// HasOut returns a boolean if a field has been set.
func (o *EncryptFile) HasOut() bool {
	if o != nil && o.Out != nil {
		return true
	}

	return false
}

// SetOut gets a reference to the given string and assigns it to the Out field.
func (o *EncryptFile) SetOut(v string) {
	o.Out = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EncryptFile) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EncryptFile) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EncryptFile) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EncryptFile) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFile) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EncryptFile) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EncryptFile) SetUidToken(v string) {
	o.UidToken = &v
}

func (o EncryptFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptionContext != nil {
		toSerialize["encryption-context"] = o.EncryptionContext
	}
	if true {
		toSerialize["in"] = o.In
	}
	if true {
		toSerialize["key-name"] = o.KeyName
	}
	if o.Out != nil {
		toSerialize["out"] = o.Out
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableEncryptFile struct {
	value *EncryptFile
	isSet bool
}

func (v NullableEncryptFile) Get() *EncryptFile {
	return v.value
}

func (v *NullableEncryptFile) Set(val *EncryptFile) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptFile) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptFile(val *EncryptFile) *NullableEncryptFile {
	return &NullableEncryptFile{value: val, isSet: true}
}

func (v NullableEncryptFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


