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

// UpdateWebTargetOutput struct for UpdateWebTargetOutput
type UpdateWebTargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewUpdateWebTargetOutput instantiates a new UpdateWebTargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebTargetOutput() *UpdateWebTargetOutput {
	this := UpdateWebTargetOutput{}
	return &this
}

// NewUpdateWebTargetOutputWithDefaults instantiates a new UpdateWebTargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebTargetOutputWithDefaults() *UpdateWebTargetOutput {
	this := UpdateWebTargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *UpdateWebTargetOutput) GetTargetId() int64 {
	if o == nil || o.TargetId == nil {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebTargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *UpdateWebTargetOutput) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *UpdateWebTargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o UpdateWebTargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetId != nil {
		toSerialize["target_id"] = o.TargetId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWebTargetOutput struct {
	value *UpdateWebTargetOutput
	isSet bool
}

func (v NullableUpdateWebTargetOutput) Get() *UpdateWebTargetOutput {
	return v.value
}

func (v *NullableUpdateWebTargetOutput) Set(val *UpdateWebTargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebTargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebTargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebTargetOutput(val *UpdateWebTargetOutput) *NullableUpdateWebTargetOutput {
	return &NullableUpdateWebTargetOutput{value: val, isSet: true}
}

func (v NullableUpdateWebTargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebTargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


