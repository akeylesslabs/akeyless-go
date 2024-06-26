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

// Encrypt struct for Encrypt
type Encrypt struct {
	// The display id of the key to use in the encryption process
	DisplayId *string `json:"display-id,omitempty"`
	// name-value pair that specifies the encryption context to be used for authenticated encryption. If used here, the same value must be supplied to the decrypt command or decryption will fail
	EncryptionContext *map[string]string `json:"encryption-context,omitempty"`
	// Select default assumed format for any plaintext input. Currently supported options: [base64]
	InputFormat *string `json:"input-format,omitempty"`
	// The item id of the key to use in the encryption process
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of the key to use in the encryption process
	KeyName string `json:"key-name"`
	// Data to be encrypted
	Plaintext *string `json:"plaintext,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// key version (relevant only for classic key)
	Version *int32 `json:"version,omitempty"`
}

// NewEncrypt instantiates a new Encrypt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncrypt(keyName string, ) *Encrypt {
	this := Encrypt{}
	var json bool = false
	this.Json = &json
	this.KeyName = keyName
	return &this
}

// NewEncryptWithDefaults instantiates a new Encrypt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptWithDefaults() *Encrypt {
	this := Encrypt{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *Encrypt) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *Encrypt) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *Encrypt) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetEncryptionContext returns the EncryptionContext field value if set, zero value otherwise.
func (o *Encrypt) GetEncryptionContext() map[string]string {
	if o == nil || o.EncryptionContext == nil {
		var ret map[string]string
		return ret
	}
	return *o.EncryptionContext
}

// GetEncryptionContextOk returns a tuple with the EncryptionContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt) GetEncryptionContextOk() (*map[string]string, bool) {
	if o == nil || o.EncryptionContext == nil {
		return nil, false
	}
	return o.EncryptionContext, true
}

// HasEncryptionContext returns a boolean if a field has been set.
func (o *Encrypt) HasEncryptionContext() bool {
	if o != nil && o.EncryptionContext != nil {
		return true
	}

	return false
}

// SetEncryptionContext gets a reference to the given map[string]string and assigns it to the EncryptionContext field.
func (o *Encrypt) SetEncryptionContext(v map[string]string) {
	o.EncryptionContext = &v
}

// GetInputFormat returns the InputFormat field value if set, zero value otherwise.
func (o *Encrypt) GetInputFormat() string {
	if o == nil || o.InputFormat == nil {
		var ret string
		return ret
	}
	return *o.InputFormat
}

// GetInputFormatOk returns a tuple with the InputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt) GetInputFormatOk() (*string, bool) {
	if o == nil || o.InputFormat == nil {
		return nil, false
	}
	return o.InputFormat, true
}

// HasInputFormat returns a boolean if a field has been set.
func (o *Encrypt) HasInputFormat() bool {
	if o != nil && o.InputFormat != nil {
		return true
	}

	return false
}

// SetInputFormat gets a reference to the given string and assigns it to the InputFormat field.
func (o *Encrypt) SetInputFormat(v string) {
	o.InputFormat = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *Encrypt) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *Encrypt) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *Encrypt) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *Encrypt) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *Encrypt) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *Encrypt) SetJson(v bool) {
	o.Json = &v
}

// GetKeyName returns the KeyName field value
func (o *Encrypt) GetKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *Encrypt) GetKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *Encrypt) SetKeyName(v string) {
	o.KeyName = v
}

// GetPlaintext returns the Plaintext field value if set, zero value otherwise.
func (o *Encrypt) GetPlaintext() string {
	if o == nil || o.Plaintext == nil {
		var ret string
		return ret
	}
	return *o.Plaintext
}

// GetPlaintextOk returns a tuple with the Plaintext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt) GetPlaintextOk() (*string, bool) {
	if o == nil || o.Plaintext == nil {
		return nil, false
	}
	return o.Plaintext, true
}

// HasPlaintext returns a boolean if a field has been set.
func (o *Encrypt) HasPlaintext() bool {
	if o != nil && o.Plaintext != nil {
		return true
	}

	return false
}

// SetPlaintext gets a reference to the given string and assigns it to the Plaintext field.
func (o *Encrypt) SetPlaintext(v string) {
	o.Plaintext = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Encrypt) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Encrypt) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Encrypt) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *Encrypt) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *Encrypt) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *Encrypt) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Encrypt) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Encrypt) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Encrypt) SetVersion(v int32) {
	o.Version = &v
}

func (o Encrypt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayId != nil {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.EncryptionContext != nil {
		toSerialize["encryption-context"] = o.EncryptionContext
	}
	if o.InputFormat != nil {
		toSerialize["input-format"] = o.InputFormat
	}
	if o.ItemId != nil {
		toSerialize["item-id"] = o.ItemId
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["key-name"] = o.KeyName
	}
	if o.Plaintext != nil {
		toSerialize["plaintext"] = o.Plaintext
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableEncrypt struct {
	value *Encrypt
	isSet bool
}

func (v NullableEncrypt) Get() *Encrypt {
	return v.value
}

func (v *NullableEncrypt) Set(val *Encrypt) {
	v.value = val
	v.isSet = true
}

func (v NullableEncrypt) IsSet() bool {
	return v.isSet
}

func (v *NullableEncrypt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncrypt(val *Encrypt) *NullableEncrypt {
	return &NullableEncrypt{value: val, isSet: true}
}

func (v NullableEncrypt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncrypt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


