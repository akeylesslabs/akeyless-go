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

// checks if the UpdateAuthMethodK8SOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAuthMethodK8SOutput{}

// UpdateAuthMethodK8SOutput struct for UpdateAuthMethodK8SOutput
type UpdateAuthMethodK8SOutput struct {
	PrvKey *string `json:"prv_key,omitempty"`
}

// NewUpdateAuthMethodK8SOutput instantiates a new UpdateAuthMethodK8SOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthMethodK8SOutput() *UpdateAuthMethodK8SOutput {
	this := UpdateAuthMethodK8SOutput{}
	return &this
}

// NewUpdateAuthMethodK8SOutputWithDefaults instantiates a new UpdateAuthMethodK8SOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthMethodK8SOutputWithDefaults() *UpdateAuthMethodK8SOutput {
	this := UpdateAuthMethodK8SOutput{}
	return &this
}

// GetPrvKey returns the PrvKey field value if set, zero value otherwise.
func (o *UpdateAuthMethodK8SOutput) GetPrvKey() string {
	if o == nil || IsNil(o.PrvKey) {
		var ret string
		return ret
	}
	return *o.PrvKey
}

// GetPrvKeyOk returns a tuple with the PrvKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodK8SOutput) GetPrvKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrvKey) {
		return nil, false
	}
	return o.PrvKey, true
}

// HasPrvKey returns a boolean if a field has been set.
func (o *UpdateAuthMethodK8SOutput) HasPrvKey() bool {
	if o != nil && !IsNil(o.PrvKey) {
		return true
	}

	return false
}

// SetPrvKey gets a reference to the given string and assigns it to the PrvKey field.
func (o *UpdateAuthMethodK8SOutput) SetPrvKey(v string) {
	o.PrvKey = &v
}

func (o UpdateAuthMethodK8SOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAuthMethodK8SOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrvKey) {
		toSerialize["prv_key"] = o.PrvKey
	}
	return toSerialize, nil
}

type NullableUpdateAuthMethodK8SOutput struct {
	value *UpdateAuthMethodK8SOutput
	isSet bool
}

func (v NullableUpdateAuthMethodK8SOutput) Get() *UpdateAuthMethodK8SOutput {
	return v.value
}

func (v *NullableUpdateAuthMethodK8SOutput) Set(val *UpdateAuthMethodK8SOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthMethodK8SOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthMethodK8SOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthMethodK8SOutput(val *UpdateAuthMethodK8SOutput) *NullableUpdateAuthMethodK8SOutput {
	return &NullableUpdateAuthMethodK8SOutput{value: val, isSet: true}
}

func (v NullableUpdateAuthMethodK8SOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthMethodK8SOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


