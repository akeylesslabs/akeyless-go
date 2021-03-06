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

// GatewayAddSubAdminsOutput struct for GatewayAddSubAdminsOutput
type GatewayAddSubAdminsOutput struct {
	SubAdmins *[]string `json:"sub-admins,omitempty"`
}

// NewGatewayAddSubAdminsOutput instantiates a new GatewayAddSubAdminsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayAddSubAdminsOutput() *GatewayAddSubAdminsOutput {
	this := GatewayAddSubAdminsOutput{}
	return &this
}

// NewGatewayAddSubAdminsOutputWithDefaults instantiates a new GatewayAddSubAdminsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayAddSubAdminsOutputWithDefaults() *GatewayAddSubAdminsOutput {
	this := GatewayAddSubAdminsOutput{}
	return &this
}

// GetSubAdmins returns the SubAdmins field value if set, zero value otherwise.
func (o *GatewayAddSubAdminsOutput) GetSubAdmins() []string {
	if o == nil || o.SubAdmins == nil {
		var ret []string
		return ret
	}
	return *o.SubAdmins
}

// GetSubAdminsOk returns a tuple with the SubAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayAddSubAdminsOutput) GetSubAdminsOk() (*[]string, bool) {
	if o == nil || o.SubAdmins == nil {
		return nil, false
	}
	return o.SubAdmins, true
}

// HasSubAdmins returns a boolean if a field has been set.
func (o *GatewayAddSubAdminsOutput) HasSubAdmins() bool {
	if o != nil && o.SubAdmins != nil {
		return true
	}

	return false
}

// SetSubAdmins gets a reference to the given []string and assigns it to the SubAdmins field.
func (o *GatewayAddSubAdminsOutput) SetSubAdmins(v []string) {
	o.SubAdmins = &v
}

func (o GatewayAddSubAdminsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubAdmins != nil {
		toSerialize["sub-admins"] = o.SubAdmins
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayAddSubAdminsOutput struct {
	value *GatewayAddSubAdminsOutput
	isSet bool
}

func (v NullableGatewayAddSubAdminsOutput) Get() *GatewayAddSubAdminsOutput {
	return v.value
}

func (v *NullableGatewayAddSubAdminsOutput) Set(val *GatewayAddSubAdminsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayAddSubAdminsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayAddSubAdminsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayAddSubAdminsOutput(val *GatewayAddSubAdminsOutput) *NullableGatewayAddSubAdminsOutput {
	return &NullableGatewayAddSubAdminsOutput{value: val, isSet: true}
}

func (v NullableGatewayAddSubAdminsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayAddSubAdminsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


