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

// GetPKICertificateOutput struct for GetPKICertificateOutput
type GetPKICertificateOutput struct {
	Path *string `json:"path,omitempty"`
}

// NewGetPKICertificateOutput instantiates a new GetPKICertificateOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPKICertificateOutput() *GetPKICertificateOutput {
	this := GetPKICertificateOutput{}
	return &this
}

// NewGetPKICertificateOutputWithDefaults instantiates a new GetPKICertificateOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPKICertificateOutputWithDefaults() *GetPKICertificateOutput {
	this := GetPKICertificateOutput{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GetPKICertificateOutput) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPKICertificateOutput) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GetPKICertificateOutput) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GetPKICertificateOutput) SetPath(v string) {
	o.Path = &v
}

func (o GetPKICertificateOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableGetPKICertificateOutput struct {
	value *GetPKICertificateOutput
	isSet bool
}

func (v NullableGetPKICertificateOutput) Get() *GetPKICertificateOutput {
	return v.value
}

func (v *NullableGetPKICertificateOutput) Set(val *GetPKICertificateOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPKICertificateOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPKICertificateOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPKICertificateOutput(val *GetPKICertificateOutput) *NullableGetPKICertificateOutput {
	return &NullableGetPKICertificateOutput{value: val, isSet: true}
}

func (v NullableGetPKICertificateOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPKICertificateOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

