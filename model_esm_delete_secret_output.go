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

// EsmDeleteSecretOutput struct for EsmDeleteSecretOutput
type EsmDeleteSecretOutput struct {
	Name *string `json:"name,omitempty"`
}

// NewEsmDeleteSecretOutput instantiates a new EsmDeleteSecretOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsmDeleteSecretOutput() *EsmDeleteSecretOutput {
	this := EsmDeleteSecretOutput{}
	return &this
}

// NewEsmDeleteSecretOutputWithDefaults instantiates a new EsmDeleteSecretOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsmDeleteSecretOutputWithDefaults() *EsmDeleteSecretOutput {
	this := EsmDeleteSecretOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EsmDeleteSecretOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsmDeleteSecretOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EsmDeleteSecretOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EsmDeleteSecretOutput) SetName(v string) {
	o.Name = &v
}

func (o EsmDeleteSecretOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableEsmDeleteSecretOutput struct {
	value *EsmDeleteSecretOutput
	isSet bool
}

func (v NullableEsmDeleteSecretOutput) Get() *EsmDeleteSecretOutput {
	return v.value
}

func (v *NullableEsmDeleteSecretOutput) Set(val *EsmDeleteSecretOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableEsmDeleteSecretOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableEsmDeleteSecretOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsmDeleteSecretOutput(val *EsmDeleteSecretOutput) *NullableEsmDeleteSecretOutput {
	return &NullableEsmDeleteSecretOutput{value: val, isSet: true}
}

func (v NullableEsmDeleteSecretOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsmDeleteSecretOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

