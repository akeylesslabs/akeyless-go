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

// CreateGodaddyTargetOutput struct for CreateGodaddyTargetOutput
type CreateGodaddyTargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewCreateGodaddyTargetOutput instantiates a new CreateGodaddyTargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGodaddyTargetOutput() *CreateGodaddyTargetOutput {
	this := CreateGodaddyTargetOutput{}
	return &this
}

// NewCreateGodaddyTargetOutputWithDefaults instantiates a new CreateGodaddyTargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGodaddyTargetOutputWithDefaults() *CreateGodaddyTargetOutput {
	this := CreateGodaddyTargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CreateGodaddyTargetOutput) GetTargetId() int64 {
	if o == nil || o.TargetId == nil {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CreateGodaddyTargetOutput) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *CreateGodaddyTargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o CreateGodaddyTargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetId != nil {
		toSerialize["target_id"] = o.TargetId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGodaddyTargetOutput struct {
	value *CreateGodaddyTargetOutput
	isSet bool
}

func (v NullableCreateGodaddyTargetOutput) Get() *CreateGodaddyTargetOutput {
	return v.value
}

func (v *NullableCreateGodaddyTargetOutput) Set(val *CreateGodaddyTargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGodaddyTargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGodaddyTargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGodaddyTargetOutput(val *CreateGodaddyTargetOutput) *NullableCreateGodaddyTargetOutput {
	return &NullableCreateGodaddyTargetOutput{value: val, isSet: true}
}

func (v NullableCreateGodaddyTargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGodaddyTargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


