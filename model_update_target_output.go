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

// checks if the UpdateTargetOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTargetOutput{}

// UpdateTargetOutput struct for UpdateTargetOutput
type UpdateTargetOutput struct {
	Updated *bool `json:"updated,omitempty"`
}

// NewUpdateTargetOutput instantiates a new UpdateTargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTargetOutput() *UpdateTargetOutput {
	this := UpdateTargetOutput{}
	return &this
}

// NewUpdateTargetOutputWithDefaults instantiates a new UpdateTargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTargetOutputWithDefaults() *UpdateTargetOutput {
	this := UpdateTargetOutput{}
	return &this
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *UpdateTargetOutput) GetUpdated() bool {
	if o == nil || IsNil(o.Updated) {
		var ret bool
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTargetOutput) GetUpdatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *UpdateTargetOutput) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given bool and assigns it to the Updated field.
func (o *UpdateTargetOutput) SetUpdated(v bool) {
	o.Updated = &v
}

func (o UpdateTargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTargetOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	return toSerialize, nil
}

type NullableUpdateTargetOutput struct {
	value *UpdateTargetOutput
	isSet bool
}

func (v NullableUpdateTargetOutput) Get() *UpdateTargetOutput {
	return v.value
}

func (v *NullableUpdateTargetOutput) Set(val *UpdateTargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTargetOutput(val *UpdateTargetOutput) *NullableUpdateTargetOutput {
	return &NullableUpdateTargetOutput{value: val, isSet: true}
}

func (v NullableUpdateTargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


