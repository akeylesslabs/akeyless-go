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

// UpdateRoleOutput struct for UpdateRoleOutput
type UpdateRoleOutput struct {
	Updated *bool `json:"updated,omitempty"`
}

// NewUpdateRoleOutput instantiates a new UpdateRoleOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleOutput() *UpdateRoleOutput {
	this := UpdateRoleOutput{}
	return &this
}

// NewUpdateRoleOutputWithDefaults instantiates a new UpdateRoleOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleOutputWithDefaults() *UpdateRoleOutput {
	this := UpdateRoleOutput{}
	return &this
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *UpdateRoleOutput) GetUpdated() bool {
	if o == nil || o.Updated == nil {
		var ret bool
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleOutput) GetUpdatedOk() (*bool, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *UpdateRoleOutput) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given bool and assigns it to the Updated field.
func (o *UpdateRoleOutput) SetUpdated(v bool) {
	o.Updated = &v
}

func (o UpdateRoleOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRoleOutput struct {
	value *UpdateRoleOutput
	isSet bool
}

func (v NullableUpdateRoleOutput) Get() *UpdateRoleOutput {
	return v.value
}

func (v *NullableUpdateRoleOutput) Set(val *UpdateRoleOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleOutput(val *UpdateRoleOutput) *NullableUpdateRoleOutput {
	return &NullableUpdateRoleOutput{value: val, isSet: true}
}

func (v NullableUpdateRoleOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


