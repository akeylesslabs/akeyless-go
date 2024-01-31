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

// VerifyPKCS1 struct for VerifyPKCS1
type VerifyPKCS1 struct {
	// The display id of the key to use in the verification process
	DisplayId *string `json:"display-id,omitempty"`
	// HashFunction defines the hash function (e.g. sha-256)
	HashFunction *string `json:"hash-function,omitempty"`
	// Select default assumed format for the plaintext message. Currently supported options: [base64]
	InputFormat *string `json:"input-format,omitempty"`
	// The item id of the key to use in the verification process
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of the RSA key to use in the verification process
	KeyName string `json:"key-name"`
	// The message to be verified
	Message string `json:"message"`
	// Markes that the message is already hashed
	Prehashed *bool `json:"prehashed,omitempty"`
	// The message's signature
	Signature string `json:"signature"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The version of the key to use for verification
	Version *int32 `json:"version,omitempty"`
}

// NewVerifyPKCS1 instantiates a new VerifyPKCS1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyPKCS1(keyName string, message string, signature string, ) *VerifyPKCS1 {
	this := VerifyPKCS1{}
	var json bool = false
	this.Json = &json
	this.KeyName = keyName
	this.Message = message
	this.Signature = signature
	return &this
}

// NewVerifyPKCS1WithDefaults instantiates a new VerifyPKCS1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyPKCS1WithDefaults() *VerifyPKCS1 {
	this := VerifyPKCS1{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *VerifyPKCS1) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *VerifyPKCS1) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *VerifyPKCS1) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetHashFunction returns the HashFunction field value if set, zero value otherwise.
func (o *VerifyPKCS1) GetHashFunction() string {
	if o == nil || o.HashFunction == nil {
		var ret string
		return ret
	}
	return *o.HashFunction
}

// GetHashFunctionOk returns a tuple with the HashFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetHashFunctionOk() (*string, bool) {
	if o == nil || o.HashFunction == nil {
		return nil, false
	}
	return o.HashFunction, true
}

// HasHashFunction returns a boolean if a field has been set.
func (o *VerifyPKCS1) HasHashFunction() bool {
	if o != nil && o.HashFunction != nil {
		return true
	}

	return false
}

// SetHashFunction gets a reference to the given string and assigns it to the HashFunction field.
func (o *VerifyPKCS1) SetHashFunction(v string) {
	o.HashFunction = &v
}

// GetInputFormat returns the InputFormat field value if set, zero value otherwise.
func (o *VerifyPKCS1) GetInputFormat() string {
	if o == nil || o.InputFormat == nil {
		var ret string
		return ret
	}
	return *o.InputFormat
}

// GetInputFormatOk returns a tuple with the InputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetInputFormatOk() (*string, bool) {
	if o == nil || o.InputFormat == nil {
		return nil, false
	}
	return o.InputFormat, true
}

// HasInputFormat returns a boolean if a field has been set.
func (o *VerifyPKCS1) HasInputFormat() bool {
	if o != nil && o.InputFormat != nil {
		return true
	}

	return false
}

// SetInputFormat gets a reference to the given string and assigns it to the InputFormat field.
func (o *VerifyPKCS1) SetInputFormat(v string) {
	o.InputFormat = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *VerifyPKCS1) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *VerifyPKCS1) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *VerifyPKCS1) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *VerifyPKCS1) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *VerifyPKCS1) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *VerifyPKCS1) SetJson(v bool) {
	o.Json = &v
}

// GetKeyName returns the KeyName field value
func (o *VerifyPKCS1) GetKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *VerifyPKCS1) SetKeyName(v string) {
	o.KeyName = v
}

// GetMessage returns the Message field value
func (o *VerifyPKCS1) GetMessage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *VerifyPKCS1) SetMessage(v string) {
	o.Message = v
}

// GetPrehashed returns the Prehashed field value if set, zero value otherwise.
func (o *VerifyPKCS1) GetPrehashed() bool {
	if o == nil || o.Prehashed == nil {
		var ret bool
		return ret
	}
	return *o.Prehashed
}

// GetPrehashedOk returns a tuple with the Prehashed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetPrehashedOk() (*bool, bool) {
	if o == nil || o.Prehashed == nil {
		return nil, false
	}
	return o.Prehashed, true
}

// HasPrehashed returns a boolean if a field has been set.
func (o *VerifyPKCS1) HasPrehashed() bool {
	if o != nil && o.Prehashed != nil {
		return true
	}

	return false
}

// SetPrehashed gets a reference to the given bool and assigns it to the Prehashed field.
func (o *VerifyPKCS1) SetPrehashed(v bool) {
	o.Prehashed = &v
}

// GetSignature returns the Signature field value
func (o *VerifyPKCS1) GetSignature() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetSignatureOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *VerifyPKCS1) SetSignature(v string) {
	o.Signature = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *VerifyPKCS1) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *VerifyPKCS1) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *VerifyPKCS1) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *VerifyPKCS1) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *VerifyPKCS1) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *VerifyPKCS1) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VerifyPKCS1) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPKCS1) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VerifyPKCS1) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *VerifyPKCS1) SetVersion(v int32) {
	o.Version = &v
}

func (o VerifyPKCS1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayId != nil {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.HashFunction != nil {
		toSerialize["hash-function"] = o.HashFunction
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
		toSerialize["message"] = o.Message
	}
	if o.Prehashed != nil {
		toSerialize["prehashed"] = o.Prehashed
	}
	if true {
		toSerialize["signature"] = o.Signature
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

type NullableVerifyPKCS1 struct {
	value *VerifyPKCS1
	isSet bool
}

func (v NullableVerifyPKCS1) Get() *VerifyPKCS1 {
	return v.value
}

func (v *NullableVerifyPKCS1) Set(val *VerifyPKCS1) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyPKCS1) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyPKCS1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyPKCS1(val *VerifyPKCS1) *NullableVerifyPKCS1 {
	return &NullableVerifyPKCS1{value: val, isSet: true}
}

func (v NullableVerifyPKCS1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyPKCS1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


