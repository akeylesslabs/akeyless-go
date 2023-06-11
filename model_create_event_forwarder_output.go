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

// checks if the CreateEventForwarderOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEventForwarderOutput{}

// CreateEventForwarderOutput struct for CreateEventForwarderOutput
type CreateEventForwarderOutput struct {
	EventForwarderId *int64 `json:"event_forwarder_id,omitempty"`
}

// NewCreateEventForwarderOutput instantiates a new CreateEventForwarderOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEventForwarderOutput() *CreateEventForwarderOutput {
	this := CreateEventForwarderOutput{}
	return &this
}

// NewCreateEventForwarderOutputWithDefaults instantiates a new CreateEventForwarderOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEventForwarderOutputWithDefaults() *CreateEventForwarderOutput {
	this := CreateEventForwarderOutput{}
	return &this
}

// GetEventForwarderId returns the EventForwarderId field value if set, zero value otherwise.
func (o *CreateEventForwarderOutput) GetEventForwarderId() int64 {
	if o == nil || IsNil(o.EventForwarderId) {
		var ret int64
		return ret
	}
	return *o.EventForwarderId
}

// GetEventForwarderIdOk returns a tuple with the EventForwarderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventForwarderOutput) GetEventForwarderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EventForwarderId) {
		return nil, false
	}
	return o.EventForwarderId, true
}

// HasEventForwarderId returns a boolean if a field has been set.
func (o *CreateEventForwarderOutput) HasEventForwarderId() bool {
	if o != nil && !IsNil(o.EventForwarderId) {
		return true
	}

	return false
}

// SetEventForwarderId gets a reference to the given int64 and assigns it to the EventForwarderId field.
func (o *CreateEventForwarderOutput) SetEventForwarderId(v int64) {
	o.EventForwarderId = &v
}

func (o CreateEventForwarderOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEventForwarderOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventForwarderId) {
		toSerialize["event_forwarder_id"] = o.EventForwarderId
	}
	return toSerialize, nil
}

type NullableCreateEventForwarderOutput struct {
	value *CreateEventForwarderOutput
	isSet bool
}

func (v NullableCreateEventForwarderOutput) Get() *CreateEventForwarderOutput {
	return v.value
}

func (v *NullableCreateEventForwarderOutput) Set(val *CreateEventForwarderOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEventForwarderOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEventForwarderOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEventForwarderOutput(val *CreateEventForwarderOutput) *NullableCreateEventForwarderOutput {
	return &NullableCreateEventForwarderOutput{value: val, isSet: true}
}

func (v NullableCreateEventForwarderOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEventForwarderOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


