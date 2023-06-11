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

// EncryptGPG struct for EncryptGPG
type EncryptGPG struct {
	// The display id of the key to use in the encryption process
	DisplayId *string `json:"display-id,omitempty"`
	// If specified, the plaintext input is assumed to be formatted accordingly. Current supported options: [base64]
	InputFormat *string `json:"input-format,omitempty"`
	// The item id of the key to use in the encryption process
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of the key to use in the encryption process
	KeyName string `json:"key-name"`
	// Data to be encrypted
	Plaintext string `json:"plaintext"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewEncryptGPG instantiates a new EncryptGPG object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptGPG(keyName string, plaintext string) *EncryptGPG {
	this := EncryptGPG{}
	var json bool = false
	this.Json = &json
	this.KeyName = keyName
	this.Plaintext = plaintext
	return &this
}

// NewEncryptGPGWithDefaults instantiates a new EncryptGPG object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptGPGWithDefaults() *EncryptGPG {
	this := EncryptGPG{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *EncryptGPG) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptGPG) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *EncryptGPG) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *EncryptGPG) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetInputFormat returns the InputFormat field value if set, zero value otherwise.
func (o *EncryptGPG) GetInputFormat() string {
	if o == nil || o.InputFormat == nil {
		var ret string
		return ret
	}
	return *o.InputFormat
}

// GetInputFormatOk returns a tuple with the InputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptGPG) GetInputFormatOk() (*string, bool) {
	if o == nil || o.InputFormat == nil {
		return nil, false
	}
	return o.InputFormat, true
}

// HasInputFormat returns a boolean if a field has been set.
func (o *EncryptGPG) HasInputFormat() bool {
	if o != nil && o.InputFormat != nil {
		return true
	}

	return false
}

// SetInputFormat gets a reference to the given string and assigns it to the InputFormat field.
func (o *EncryptGPG) SetInputFormat(v string) {
	o.InputFormat = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *EncryptGPG) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptGPG) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *EncryptGPG) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *EncryptGPG) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *EncryptGPG) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptGPG) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *EncryptGPG) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *EncryptGPG) SetJson(v bool) {
	o.Json = &v
}

// GetKeyName returns the KeyName field value
func (o *EncryptGPG) GetKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *EncryptGPG) GetKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *EncryptGPG) SetKeyName(v string) {
	o.KeyName = v
}

// GetPlaintext returns the Plaintext field value
func (o *EncryptGPG) GetPlaintext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plaintext
}

// GetPlaintextOk returns a tuple with the Plaintext field value
// and a boolean to check if the value has been set.
func (o *EncryptGPG) GetPlaintextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plaintext, true
}

// SetPlaintext sets field value
func (o *EncryptGPG) SetPlaintext(v string) {
	o.Plaintext = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EncryptGPG) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptGPG) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EncryptGPG) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EncryptGPG) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EncryptGPG) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptGPG) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EncryptGPG) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EncryptGPG) SetUidToken(v string) {
	o.UidToken = &v
}

func (o EncryptGPG) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayId != nil {
		toSerialize["display-id"] = o.DisplayId
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
	if true {
		toSerialize["plaintext"] = o.Plaintext
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableEncryptGPG struct {
	value *EncryptGPG
	isSet bool
}

func (v NullableEncryptGPG) Get() *EncryptGPG {
	return v.value
}

func (v *NullableEncryptGPG) Set(val *EncryptGPG) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptGPG) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptGPG) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptGPG(val *EncryptGPG) *NullableEncryptGPG {
	return &NullableEncryptGPG{value: val, isSet: true}
}

func (v NullableEncryptGPG) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptGPG) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


