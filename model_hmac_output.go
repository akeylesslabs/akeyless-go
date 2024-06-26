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

// HmacOutput struct for HmacOutput
type HmacOutput struct {
	Result *string `json:"result,omitempty"`
}

// NewHmacOutput instantiates a new HmacOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHmacOutput() *HmacOutput {
	this := HmacOutput{}
	return &this
}

// NewHmacOutputWithDefaults instantiates a new HmacOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHmacOutputWithDefaults() *HmacOutput {
	this := HmacOutput{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *HmacOutput) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HmacOutput) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *HmacOutput) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *HmacOutput) SetResult(v string) {
	o.Result = &v
}

func (o HmacOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableHmacOutput struct {
	value *HmacOutput
	isSet bool
}

func (v NullableHmacOutput) Get() *HmacOutput {
	return v.value
}

func (v *NullableHmacOutput) Set(val *HmacOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableHmacOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableHmacOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHmacOutput(val *HmacOutput) *NullableHmacOutput {
	return &NullableHmacOutput{value: val, isSet: true}
}

func (v NullableHmacOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHmacOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


