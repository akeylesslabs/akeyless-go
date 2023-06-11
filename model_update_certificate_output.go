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

// checks if the UpdateCertificateOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCertificateOutput{}

// UpdateCertificateOutput struct for UpdateCertificateOutput
type UpdateCertificateOutput struct {
	Name *string `json:"name,omitempty"`
}

// NewUpdateCertificateOutput instantiates a new UpdateCertificateOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCertificateOutput() *UpdateCertificateOutput {
	this := UpdateCertificateOutput{}
	return &this
}

// NewUpdateCertificateOutputWithDefaults instantiates a new UpdateCertificateOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCertificateOutputWithDefaults() *UpdateCertificateOutput {
	this := UpdateCertificateOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateCertificateOutput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCertificateOutput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateCertificateOutput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateCertificateOutput) SetName(v string) {
	o.Name = &v
}

func (o UpdateCertificateOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCertificateOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateCertificateOutput struct {
	value *UpdateCertificateOutput
	isSet bool
}

func (v NullableUpdateCertificateOutput) Get() *UpdateCertificateOutput {
	return v.value
}

func (v *NullableUpdateCertificateOutput) Set(val *UpdateCertificateOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCertificateOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCertificateOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCertificateOutput(val *UpdateCertificateOutput) *NullableUpdateCertificateOutput {
	return &NullableUpdateCertificateOutput{value: val, isSet: true}
}

func (v NullableUpdateCertificateOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCertificateOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


