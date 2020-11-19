/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 0.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// UpdateItemOutput struct for UpdateItemOutput
type UpdateItemOutput struct {
	Updated *bool `json:"updated,omitempty"`
}

// NewUpdateItemOutput instantiates a new UpdateItemOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateItemOutput() *UpdateItemOutput {
	this := UpdateItemOutput{}
	return &this
}

// NewUpdateItemOutputWithDefaults instantiates a new UpdateItemOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateItemOutputWithDefaults() *UpdateItemOutput {
	this := UpdateItemOutput{}
	return &this
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *UpdateItemOutput) GetUpdated() bool {
	if o == nil || o.Updated == nil {
		var ret bool
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateItemOutput) GetUpdatedOk() (*bool, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *UpdateItemOutput) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given bool and assigns it to the Updated field.
func (o *UpdateItemOutput) SetUpdated(v bool) {
	o.Updated = &v
}

func (o UpdateItemOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateItemOutput struct {
	value *UpdateItemOutput
	isSet bool
}

func (v NullableUpdateItemOutput) Get() *UpdateItemOutput {
	return v.value
}

func (v *NullableUpdateItemOutput) Set(val *UpdateItemOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateItemOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateItemOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateItemOutput(val *UpdateItemOutput) *NullableUpdateItemOutput {
	return &NullableUpdateItemOutput{value: val, isSet: true}
}

func (v NullableUpdateItemOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateItemOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


