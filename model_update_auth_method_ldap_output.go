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

// checks if the UpdateAuthMethodLDAPOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAuthMethodLDAPOutput{}

// UpdateAuthMethodLDAPOutput struct for UpdateAuthMethodLDAPOutput
type UpdateAuthMethodLDAPOutput struct {
	PrvKey *string `json:"prv_key,omitempty"`
}

// NewUpdateAuthMethodLDAPOutput instantiates a new UpdateAuthMethodLDAPOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthMethodLDAPOutput() *UpdateAuthMethodLDAPOutput {
	this := UpdateAuthMethodLDAPOutput{}
	return &this
}

// NewUpdateAuthMethodLDAPOutputWithDefaults instantiates a new UpdateAuthMethodLDAPOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthMethodLDAPOutputWithDefaults() *UpdateAuthMethodLDAPOutput {
	this := UpdateAuthMethodLDAPOutput{}
	return &this
}

// GetPrvKey returns the PrvKey field value if set, zero value otherwise.
func (o *UpdateAuthMethodLDAPOutput) GetPrvKey() string {
	if o == nil || IsNil(o.PrvKey) {
		var ret string
		return ret
	}
	return *o.PrvKey
}

// GetPrvKeyOk returns a tuple with the PrvKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodLDAPOutput) GetPrvKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrvKey) {
		return nil, false
	}
	return o.PrvKey, true
}

// HasPrvKey returns a boolean if a field has been set.
func (o *UpdateAuthMethodLDAPOutput) HasPrvKey() bool {
	if o != nil && !IsNil(o.PrvKey) {
		return true
	}

	return false
}

// SetPrvKey gets a reference to the given string and assigns it to the PrvKey field.
func (o *UpdateAuthMethodLDAPOutput) SetPrvKey(v string) {
	o.PrvKey = &v
}

func (o UpdateAuthMethodLDAPOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAuthMethodLDAPOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrvKey) {
		toSerialize["prv_key"] = o.PrvKey
	}
	return toSerialize, nil
}

type NullableUpdateAuthMethodLDAPOutput struct {
	value *UpdateAuthMethodLDAPOutput
	isSet bool
}

func (v NullableUpdateAuthMethodLDAPOutput) Get() *UpdateAuthMethodLDAPOutput {
	return v.value
}

func (v *NullableUpdateAuthMethodLDAPOutput) Set(val *UpdateAuthMethodLDAPOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthMethodLDAPOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthMethodLDAPOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthMethodLDAPOutput(val *UpdateAuthMethodLDAPOutput) *NullableUpdateAuthMethodLDAPOutput {
	return &NullableUpdateAuthMethodLDAPOutput{value: val, isSet: true}
}

func (v NullableUpdateAuthMethodLDAPOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthMethodLDAPOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


