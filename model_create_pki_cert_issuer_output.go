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

// CreatePKICertIssuerOutput struct for CreatePKICertIssuerOutput
type CreatePKICertIssuerOutput struct {
	Name *string `json:"name,omitempty"`
}

// NewCreatePKICertIssuerOutput instantiates a new CreatePKICertIssuerOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePKICertIssuerOutput() *CreatePKICertIssuerOutput {
	this := CreatePKICertIssuerOutput{}
	return &this
}

// NewCreatePKICertIssuerOutputWithDefaults instantiates a new CreatePKICertIssuerOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePKICertIssuerOutputWithDefaults() *CreatePKICertIssuerOutput {
	this := CreatePKICertIssuerOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreatePKICertIssuerOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePKICertIssuerOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreatePKICertIssuerOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreatePKICertIssuerOutput) SetName(v string) {
	o.Name = &v
}

func (o CreatePKICertIssuerOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePKICertIssuerOutput struct {
	value *CreatePKICertIssuerOutput
	isSet bool
}

func (v NullableCreatePKICertIssuerOutput) Get() *CreatePKICertIssuerOutput {
	return v.value
}

func (v *NullableCreatePKICertIssuerOutput) Set(val *CreatePKICertIssuerOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePKICertIssuerOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePKICertIssuerOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePKICertIssuerOutput(val *CreatePKICertIssuerOutput) *NullableCreatePKICertIssuerOutput {
	return &NullableCreatePKICertIssuerOutput{value: val, isSet: true}
}

func (v NullableCreatePKICertIssuerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePKICertIssuerOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

