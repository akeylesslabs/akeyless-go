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

// UpdateAuthMethodOutput struct for UpdateAuthMethodOutput
type UpdateAuthMethodOutput struct {
	Name *string `json:"name,omitempty"`
}

// NewUpdateAuthMethodOutput instantiates a new UpdateAuthMethodOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthMethodOutput() *UpdateAuthMethodOutput {
	this := UpdateAuthMethodOutput{}
	return &this
}

// NewUpdateAuthMethodOutputWithDefaults instantiates a new UpdateAuthMethodOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthMethodOutputWithDefaults() *UpdateAuthMethodOutput {
	this := UpdateAuthMethodOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAuthMethodOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAuthMethodOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAuthMethodOutput) SetName(v string) {
	o.Name = &v
}

func (o UpdateAuthMethodOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAuthMethodOutput struct {
	value *UpdateAuthMethodOutput
	isSet bool
}

func (v NullableUpdateAuthMethodOutput) Get() *UpdateAuthMethodOutput {
	return v.value
}

func (v *NullableUpdateAuthMethodOutput) Set(val *UpdateAuthMethodOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthMethodOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthMethodOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthMethodOutput(val *UpdateAuthMethodOutput) *NullableUpdateAuthMethodOutput {
	return &NullableUpdateAuthMethodOutput{value: val, isSet: true}
}

func (v NullableUpdateAuthMethodOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthMethodOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

