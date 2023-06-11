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

// checks if the Hmac type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Hmac{}

// Hmac struct for Hmac
type Hmac struct {
	// The display id of the key to use in the encryption process
	DisplayId *string `json:"display-id,omitempty"`
	// Hash function [sha-256,sha-512]
	HashFunction *string `json:"hash-function,omitempty"`
	// Select default assumed format for any plaintext input. Currently supported options: [base64]
	InputFormat *string `json:"input-format,omitempty"`
	// The item id of the key to use in the encryption process
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of the key to use in the encryption process
	KeyName string `json:"key-name"`
	// Data to perform hmac on
	Plaintext *string `json:"plaintext,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewHmac instantiates a new Hmac object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHmac(keyName string) *Hmac {
	this := Hmac{}
	var hashFunction string = "sha-256"
	this.HashFunction = &hashFunction
	var json bool = false
	this.Json = &json
	this.KeyName = keyName
	return &this
}

// NewHmacWithDefaults instantiates a new Hmac object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHmacWithDefaults() *Hmac {
	this := Hmac{}
	var hashFunction string = "sha-256"
	this.HashFunction = &hashFunction
	var json bool = false
	this.Json = &json
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *Hmac) GetDisplayId() string {
	if o == nil || IsNil(o.DisplayId) {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hmac) GetDisplayIdOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayId) {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *Hmac) HasDisplayId() bool {
	if o != nil && !IsNil(o.DisplayId) {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *Hmac) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetHashFunction returns the HashFunction field value if set, zero value otherwise.
func (o *Hmac) GetHashFunction() string {
	if o == nil || IsNil(o.HashFunction) {
		var ret string
		return ret
	}
	return *o.HashFunction
}

// GetHashFunctionOk returns a tuple with the HashFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hmac) GetHashFunctionOk() (*string, bool) {
	if o == nil || IsNil(o.HashFunction) {
		return nil, false
	}
	return o.HashFunction, true
}

// HasHashFunction returns a boolean if a field has been set.
func (o *Hmac) HasHashFunction() bool {
	if o != nil && !IsNil(o.HashFunction) {
		return true
	}

	return false
}

// SetHashFunction gets a reference to the given string and assigns it to the HashFunction field.
func (o *Hmac) SetHashFunction(v string) {
	o.HashFunction = &v
}

// GetInputFormat returns the InputFormat field value if set, zero value otherwise.
func (o *Hmac) GetInputFormat() string {
	if o == nil || IsNil(o.InputFormat) {
		var ret string
		return ret
	}
	return *o.InputFormat
}

// GetInputFormatOk returns a tuple with the InputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hmac) GetInputFormatOk() (*string, bool) {
	if o == nil || IsNil(o.InputFormat) {
		return nil, false
	}
	return o.InputFormat, true
}

// HasInputFormat returns a boolean if a field has been set.
func (o *Hmac) HasInputFormat() bool {
	if o != nil && !IsNil(o.InputFormat) {
		return true
	}

	return false
}

// SetInputFormat gets a reference to the given string and assigns it to the InputFormat field.
func (o *Hmac) SetInputFormat(v string) {
	o.InputFormat = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *Hmac) GetItemId() int64 {
	if o == nil || IsNil(o.ItemId) {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hmac) GetItemIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *Hmac) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *Hmac) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *Hmac) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hmac) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *Hmac) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *Hmac) SetJson(v bool) {
	o.Json = &v
}

// GetKeyName returns the KeyName field value
func (o *Hmac) GetKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *Hmac) GetKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *Hmac) SetKeyName(v string) {
	o.KeyName = v
}

// GetPlaintext returns the Plaintext field value if set, zero value otherwise.
func (o *Hmac) GetPlaintext() string {
	if o == nil || IsNil(o.Plaintext) {
		var ret string
		return ret
	}
	return *o.Plaintext
}

// GetPlaintextOk returns a tuple with the Plaintext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hmac) GetPlaintextOk() (*string, bool) {
	if o == nil || IsNil(o.Plaintext) {
		return nil, false
	}
	return o.Plaintext, true
}

// HasPlaintext returns a boolean if a field has been set.
func (o *Hmac) HasPlaintext() bool {
	if o != nil && !IsNil(o.Plaintext) {
		return true
	}

	return false
}

// SetPlaintext gets a reference to the given string and assigns it to the Plaintext field.
func (o *Hmac) SetPlaintext(v string) {
	o.Plaintext = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Hmac) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hmac) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Hmac) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Hmac) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *Hmac) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hmac) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *Hmac) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *Hmac) SetUidToken(v string) {
	o.UidToken = &v
}

func (o Hmac) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hmac) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayId) {
		toSerialize["display-id"] = o.DisplayId
	}
	if !IsNil(o.HashFunction) {
		toSerialize["hash-function"] = o.HashFunction
	}
	if !IsNil(o.InputFormat) {
		toSerialize["input-format"] = o.InputFormat
	}
	if !IsNil(o.ItemId) {
		toSerialize["item-id"] = o.ItemId
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["key-name"] = o.KeyName
	if !IsNil(o.Plaintext) {
		toSerialize["plaintext"] = o.Plaintext
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableHmac struct {
	value *Hmac
	isSet bool
}

func (v NullableHmac) Get() *Hmac {
	return v.value
}

func (v *NullableHmac) Set(val *Hmac) {
	v.value = val
	v.isSet = true
}

func (v NullableHmac) IsSet() bool {
	return v.isSet
}

func (v *NullableHmac) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHmac(val *Hmac) *NullableHmac {
	return &NullableHmac{value: val, isSet: true}
}

func (v NullableHmac) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHmac) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


