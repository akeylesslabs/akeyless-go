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

// checks if the CreateDockerhubTargetOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDockerhubTargetOutput{}

// CreateDockerhubTargetOutput struct for CreateDockerhubTargetOutput
type CreateDockerhubTargetOutput struct {
	TargetId *int64 `json:"target_id,omitempty"`
}

// NewCreateDockerhubTargetOutput instantiates a new CreateDockerhubTargetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDockerhubTargetOutput() *CreateDockerhubTargetOutput {
	this := CreateDockerhubTargetOutput{}
	return &this
}

// NewCreateDockerhubTargetOutputWithDefaults instantiates a new CreateDockerhubTargetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDockerhubTargetOutputWithDefaults() *CreateDockerhubTargetOutput {
	this := CreateDockerhubTargetOutput{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CreateDockerhubTargetOutput) GetTargetId() int64 {
	if o == nil || IsNil(o.TargetId) {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDockerhubTargetOutput) GetTargetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CreateDockerhubTargetOutput) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *CreateDockerhubTargetOutput) SetTargetId(v int64) {
	o.TargetId = &v
}

func (o CreateDockerhubTargetOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDockerhubTargetOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	return toSerialize, nil
}

type NullableCreateDockerhubTargetOutput struct {
	value *CreateDockerhubTargetOutput
	isSet bool
}

func (v NullableCreateDockerhubTargetOutput) Get() *CreateDockerhubTargetOutput {
	return v.value
}

func (v *NullableCreateDockerhubTargetOutput) Set(val *CreateDockerhubTargetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDockerhubTargetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDockerhubTargetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDockerhubTargetOutput(val *CreateDockerhubTargetOutput) *NullableCreateDockerhubTargetOutput {
	return &NullableCreateDockerhubTargetOutput{value: val, isSet: true}
}

func (v NullableCreateDockerhubTargetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDockerhubTargetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


