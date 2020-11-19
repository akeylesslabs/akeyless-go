/*
 * includes all operations supported by Akeyless.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// AuthOutput struct for AuthOutput
type AuthOutput struct {
	Token *string `json:"token,omitempty"`
}

// NewAuthOutput instantiates a new AuthOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthOutput() *AuthOutput {
	this := AuthOutput{}
	return &this
}

// NewAuthOutputWithDefaults instantiates a new AuthOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthOutputWithDefaults() *AuthOutput {
	this := AuthOutput{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthOutput) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOutput) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthOutput) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AuthOutput) SetToken(v string) {
	o.Token = &v
}

func (o AuthOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableAuthOutput struct {
	value *AuthOutput
	isSet bool
}

func (v NullableAuthOutput) Get() *AuthOutput {
	return v.value
}

func (v *NullableAuthOutput) Set(val *AuthOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthOutput(val *AuthOutput) *NullableAuthOutput {
	return &NullableAuthOutput{value: val, isSet: true}
}

func (v NullableAuthOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


