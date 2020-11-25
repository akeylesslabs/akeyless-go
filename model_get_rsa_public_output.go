/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GetRSAPublicOutput struct for GetRSAPublicOutput
type GetRSAPublicOutput struct {
	Raw *string `json:"raw,omitempty"`
	Ssh *[]int32 `json:"ssh,omitempty"`
}

// NewGetRSAPublicOutput instantiates a new GetRSAPublicOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRSAPublicOutput() *GetRSAPublicOutput {
	this := GetRSAPublicOutput{}
	return &this
}

// NewGetRSAPublicOutputWithDefaults instantiates a new GetRSAPublicOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRSAPublicOutputWithDefaults() *GetRSAPublicOutput {
	this := GetRSAPublicOutput{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *GetRSAPublicOutput) GetRaw() string {
	if o == nil || o.Raw == nil {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRSAPublicOutput) GetRawOk() (*string, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *GetRSAPublicOutput) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *GetRSAPublicOutput) SetRaw(v string) {
	o.Raw = &v
}

// GetSsh returns the Ssh field value if set, zero value otherwise.
func (o *GetRSAPublicOutput) GetSsh() []int32 {
	if o == nil || o.Ssh == nil {
		var ret []int32
		return ret
	}
	return *o.Ssh
}

// GetSshOk returns a tuple with the Ssh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRSAPublicOutput) GetSshOk() (*[]int32, bool) {
	if o == nil || o.Ssh == nil {
		return nil, false
	}
	return o.Ssh, true
}

// HasSsh returns a boolean if a field has been set.
func (o *GetRSAPublicOutput) HasSsh() bool {
	if o != nil && o.Ssh != nil {
		return true
	}

	return false
}

// SetSsh gets a reference to the given []int32 and assigns it to the Ssh field.
func (o *GetRSAPublicOutput) SetSsh(v []int32) {
	o.Ssh = &v
}

func (o GetRSAPublicOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Ssh != nil {
		toSerialize["ssh"] = o.Ssh
	}
	return json.Marshal(toSerialize)
}

type NullableGetRSAPublicOutput struct {
	value *GetRSAPublicOutput
	isSet bool
}

func (v NullableGetRSAPublicOutput) Get() *GetRSAPublicOutput {
	return v.value
}

func (v *NullableGetRSAPublicOutput) Set(val *GetRSAPublicOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRSAPublicOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRSAPublicOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRSAPublicOutput(val *GetRSAPublicOutput) *NullableGetRSAPublicOutput {
	return &NullableGetRSAPublicOutput{value: val, isSet: true}
}

func (v NullableGetRSAPublicOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRSAPublicOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


