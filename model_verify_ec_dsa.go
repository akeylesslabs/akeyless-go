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

// VerifyEcDsa verifyEcDsa is a command that verifies an ECDSA signature using a sha hash algorithm matching the key size
type VerifyEcDsa struct {
	// The display id of the EC key to use for the verification process
	DisplayId *string `json:"display-id,omitempty"`
	// The item id of the EC key to use for the verification process
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of the EC key to use for the verification process
	KeyName *string `json:"key-name,omitempty"`
	// The message to be verified in a base64 format
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

// NewVerifyEcDsa instantiates a new VerifyEcDsa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyEcDsa(message string, signature string, ) *VerifyEcDsa {
	this := VerifyEcDsa{}
	var json bool = false
	this.Json = &json
	this.Message = message
	this.Signature = signature
	return &this
}

// NewVerifyEcDsaWithDefaults instantiates a new VerifyEcDsa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyEcDsaWithDefaults() *VerifyEcDsa {
	this := VerifyEcDsa{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *VerifyEcDsa) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *VerifyEcDsa) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *VerifyEcDsa) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *VerifyEcDsa) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *VerifyEcDsa) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *VerifyEcDsa) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *VerifyEcDsa) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *VerifyEcDsa) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *VerifyEcDsa) SetJson(v bool) {
	o.Json = &v
}

// GetKeyName returns the KeyName field value if set, zero value otherwise.
func (o *VerifyEcDsa) GetKeyName() string {
	if o == nil || o.KeyName == nil {
		var ret string
		return ret
	}
	return *o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetKeyNameOk() (*string, bool) {
	if o == nil || o.KeyName == nil {
		return nil, false
	}
	return o.KeyName, true
}

// HasKeyName returns a boolean if a field has been set.
func (o *VerifyEcDsa) HasKeyName() bool {
	if o != nil && o.KeyName != nil {
		return true
	}

	return false
}

// SetKeyName gets a reference to the given string and assigns it to the KeyName field.
func (o *VerifyEcDsa) SetKeyName(v string) {
	o.KeyName = &v
}

// GetMessage returns the Message field value
func (o *VerifyEcDsa) GetMessage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *VerifyEcDsa) SetMessage(v string) {
	o.Message = v
}

// GetPrehashed returns the Prehashed field value if set, zero value otherwise.
func (o *VerifyEcDsa) GetPrehashed() bool {
	if o == nil || o.Prehashed == nil {
		var ret bool
		return ret
	}
	return *o.Prehashed
}

// GetPrehashedOk returns a tuple with the Prehashed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetPrehashedOk() (*bool, bool) {
	if o == nil || o.Prehashed == nil {
		return nil, false
	}
	return o.Prehashed, true
}

// HasPrehashed returns a boolean if a field has been set.
func (o *VerifyEcDsa) HasPrehashed() bool {
	if o != nil && o.Prehashed != nil {
		return true
	}

	return false
}

// SetPrehashed gets a reference to the given bool and assigns it to the Prehashed field.
func (o *VerifyEcDsa) SetPrehashed(v bool) {
	o.Prehashed = &v
}

// GetSignature returns the Signature field value
func (o *VerifyEcDsa) GetSignature() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetSignatureOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *VerifyEcDsa) SetSignature(v string) {
	o.Signature = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *VerifyEcDsa) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *VerifyEcDsa) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *VerifyEcDsa) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *VerifyEcDsa) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *VerifyEcDsa) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *VerifyEcDsa) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VerifyEcDsa) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyEcDsa) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VerifyEcDsa) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *VerifyEcDsa) SetVersion(v int32) {
	o.Version = &v
}

func (o VerifyEcDsa) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayId != nil {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.ItemId != nil {
		toSerialize["item-id"] = o.ItemId
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.KeyName != nil {
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

type NullableVerifyEcDsa struct {
	value *VerifyEcDsa
	isSet bool
}

func (v NullableVerifyEcDsa) Get() *VerifyEcDsa {
	return v.value
}

func (v *NullableVerifyEcDsa) Set(val *VerifyEcDsa) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyEcDsa) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyEcDsa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyEcDsa(val *VerifyEcDsa) *NullableVerifyEcDsa {
	return &NullableVerifyEcDsa{value: val, isSet: true}
}

func (v NullableVerifyEcDsa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyEcDsa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


