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

// SignPKCS1 signPKCS1 is a command that calculates the signature of hashed data using RSASSA-PKCS1-V1_5-SIGN from RSA PKCS#1 v1.5.
type SignPKCS1 struct {
	// The name of the RSA key to use in the signing process
	KeyName string `json:"key-name"`
	// The message to be signed
	Message string `json:"message"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewSignPKCS1 instantiates a new SignPKCS1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignPKCS1(keyName string, message string, ) *SignPKCS1 {
	this := SignPKCS1{}
	this.KeyName = keyName
	this.Message = message
	return &this
}

// NewSignPKCS1WithDefaults instantiates a new SignPKCS1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignPKCS1WithDefaults() *SignPKCS1 {
	this := SignPKCS1{}
	return &this
}

// GetKeyName returns the KeyName field value
func (o *SignPKCS1) GetKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *SignPKCS1) GetKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *SignPKCS1) SetKeyName(v string) {
	o.KeyName = v
}

// GetMessage returns the Message field value
func (o *SignPKCS1) GetMessage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SignPKCS1) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SignPKCS1) SetMessage(v string) {
	o.Message = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SignPKCS1) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKCS1) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SignPKCS1) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SignPKCS1) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *SignPKCS1) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignPKCS1) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *SignPKCS1) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *SignPKCS1) SetUidToken(v string) {
	o.UidToken = &v
}

func (o SignPKCS1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key-name"] = o.KeyName
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableSignPKCS1 struct {
	value *SignPKCS1
	isSet bool
}

func (v NullableSignPKCS1) Get() *SignPKCS1 {
	return v.value
}

func (v *NullableSignPKCS1) Set(val *SignPKCS1) {
	v.value = val
	v.isSet = true
}

func (v NullableSignPKCS1) IsSet() bool {
	return v.isSet
}

func (v *NullableSignPKCS1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignPKCS1(val *SignPKCS1) *NullableSignPKCS1 {
	return &NullableSignPKCS1{value: val, isSet: true}
}

func (v NullableSignPKCS1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignPKCS1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


