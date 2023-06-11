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

// checks if the CreateGKETargetOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGKETargetOutput{}

// CreateGKETargetOutput struct for CreateGKETargetOutput
type CreateGKETargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewCreateGKETargetOutput instantiates a new CreateGKETargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGKETargetOutput() *CreateGKETargetOutput {
	this := CreateGKETargetOutput{}
	return &this
}

// NewCreateGKETargetOutputWithDefaults instantiates a new CreateGKETargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGKETargetOutputWithDefaults() *CreateGKETargetOutput {
	this := CreateGKETargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CreateGKETargetOutput) GetTargetId() int64 {
	if o == nil || IsNil(o.TargetId) {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGKETargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CreateGKETargetOutput) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *CreateGKETargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o CreateGKETargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGKETargetOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	return toSerialize, nil
}

type NullableCreateGKETargetOutput struct {
	value *CreateGKETargetOutput
	isSet bool
}

func (v NullableCreateGKETargetOutput) Get() *CreateGKETargetOutput {
	return v.value
}

func (v *NullableCreateGKETargetOutput) Set(val *CreateGKETargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGKETargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGKETargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGKETargetOutput(val *CreateGKETargetOutput) *NullableCreateGKETargetOutput {
	return &NullableCreateGKETargetOutput{value: val, isSet: true}
}

func (v NullableCreateGKETargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGKETargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


