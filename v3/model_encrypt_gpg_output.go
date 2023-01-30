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

// EncryptGPGOutput struct for EncryptGPGOutput
type EncryptGPGOutput struct {
	Result *string `json:"result,omitempty"`
}

// NewEncryptGPGOutput instantiates a new EncryptGPGOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptGPGOutput() *EncryptGPGOutput {
	this := EncryptGPGOutput{}
	return &this
}

// NewEncryptGPGOutputWithDefaults instantiates a new EncryptGPGOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptGPGOutputWithDefaults() *EncryptGPGOutput {
	this := EncryptGPGOutput{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *EncryptGPGOutput) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptGPGOutput) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *EncryptGPGOutput) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *EncryptGPGOutput) SetResult(v string) {
	o.Result = &v
}

func (o EncryptGPGOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableEncryptGPGOutput struct {
	value *EncryptGPGOutput
	isSet bool
}

func (v NullableEncryptGPGOutput) Get() *EncryptGPGOutput {
	return v.value
}

func (v *NullableEncryptGPGOutput) Set(val *EncryptGPGOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptGPGOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptGPGOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptGPGOutput(val *EncryptGPGOutput) *NullableEncryptGPGOutput {
	return &NullableEncryptGPGOutput{value: val, isSet: true}
}

func (v NullableEncryptGPGOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptGPGOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

