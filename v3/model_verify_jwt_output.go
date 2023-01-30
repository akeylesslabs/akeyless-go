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

// VerifyJWTOutput struct for VerifyJWTOutput
type VerifyJWTOutput struct {
	Result *string `json:"result,omitempty"`
}

// NewVerifyJWTOutput instantiates a new VerifyJWTOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyJWTOutput() *VerifyJWTOutput {
	this := VerifyJWTOutput{}
	return &this
}

// NewVerifyJWTOutputWithDefaults instantiates a new VerifyJWTOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyJWTOutputWithDefaults() *VerifyJWTOutput {
	this := VerifyJWTOutput{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *VerifyJWTOutput) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyJWTOutput) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *VerifyJWTOutput) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *VerifyJWTOutput) SetResult(v string) {
	o.Result = &v
}

func (o VerifyJWTOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyJWTOutput struct {
	value *VerifyJWTOutput
	isSet bool
}

func (v NullableVerifyJWTOutput) Get() *VerifyJWTOutput {
	return v.value
}

func (v *NullableVerifyJWTOutput) Set(val *VerifyJWTOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyJWTOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyJWTOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyJWTOutput(val *VerifyJWTOutput) *NullableVerifyJWTOutput {
	return &NullableVerifyJWTOutput{value: val, isSet: true}
}

func (v NullableVerifyJWTOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyJWTOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

