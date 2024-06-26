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

// RotateOidcClientOutput struct for RotateOidcClientOutput
type RotateOidcClientOutput struct {
	NewSecret *string `json:"new_secret,omitempty"`
}

// NewRotateOidcClientOutput instantiates a new RotateOidcClientOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotateOidcClientOutput() *RotateOidcClientOutput {
	this := RotateOidcClientOutput{}
	return &this
}

// NewRotateOidcClientOutputWithDefaults instantiates a new RotateOidcClientOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotateOidcClientOutputWithDefaults() *RotateOidcClientOutput {
	this := RotateOidcClientOutput{}
	return &this
}

// GetNewSecret returns the NewSecret field value if set, zero value otherwise.
func (o *RotateOidcClientOutput) GetNewSecret() string {
	if o == nil || o.NewSecret == nil {
		var ret string
		return ret
	}
	return *o.NewSecret
}

// GetNewSecretOk returns a tuple with the NewSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotateOidcClientOutput) GetNewSecretOk() (*string, bool) {
	if o == nil || o.NewSecret == nil {
		return nil, false
	}
	return o.NewSecret, true
}

// HasNewSecret returns a boolean if a field has been set.
func (o *RotateOidcClientOutput) HasNewSecret() bool {
	if o != nil && o.NewSecret != nil {
		return true
	}

	return false
}

// SetNewSecret gets a reference to the given string and assigns it to the NewSecret field.
func (o *RotateOidcClientOutput) SetNewSecret(v string) {
	o.NewSecret = &v
}

func (o RotateOidcClientOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewSecret != nil {
		toSerialize["new_secret"] = o.NewSecret
	}
	return json.Marshal(toSerialize)
}

type NullableRotateOidcClientOutput struct {
	value *RotateOidcClientOutput
	isSet bool
}

func (v NullableRotateOidcClientOutput) Get() *RotateOidcClientOutput {
	return v.value
}

func (v *NullableRotateOidcClientOutput) Set(val *RotateOidcClientOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableRotateOidcClientOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRotateOidcClientOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotateOidcClientOutput(val *RotateOidcClientOutput) *NullableRotateOidcClientOutput {
	return &NullableRotateOidcClientOutput{value: val, isSet: true}
}

func (v NullableRotateOidcClientOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotateOidcClientOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


