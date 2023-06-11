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

// GetEventForwarderOutput struct for GetEventForwarderOutput
type GetEventForwarderOutput struct {
	EventForwarder *NotiForwarder `json:"event_forwarder,omitempty"`
}

// NewGetEventForwarderOutput instantiates a new GetEventForwarderOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventForwarderOutput() *GetEventForwarderOutput {
	this := GetEventForwarderOutput{}
	return &this
}

// NewGetEventForwarderOutputWithDefaults instantiates a new GetEventForwarderOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventForwarderOutputWithDefaults() *GetEventForwarderOutput {
	this := GetEventForwarderOutput{}
	return &this
}

// GetEventForwarder returns the EventForwarder field value if set, zero value otherwise.
func (o *GetEventForwarderOutput) GetEventForwarder() NotiForwarder {
	if o == nil || o.EventForwarder == nil {
		var ret NotiForwarder
		return ret
	}
	return *o.EventForwarder
}

// GetEventForwarderOk returns a tuple with the EventForwarder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventForwarderOutput) GetEventForwarderOk() (*NotiForwarder, bool) {
	if o == nil || o.EventForwarder == nil {
		return nil, false
	}
	return o.EventForwarder, true
}

// HasEventForwarder returns a boolean if a field has been set.
func (o *GetEventForwarderOutput) HasEventForwarder() bool {
	if o != nil && o.EventForwarder != nil {
		return true
	}

	return false
}

// SetEventForwarder gets a reference to the given NotiForwarder and assigns it to the EventForwarder field.
func (o *GetEventForwarderOutput) SetEventForwarder(v NotiForwarder) {
	o.EventForwarder = &v
}

func (o GetEventForwarderOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventForwarder != nil {
		toSerialize["event_forwarder"] = o.EventForwarder
	}
	return json.Marshal(toSerialize)
}

type NullableGetEventForwarderOutput struct {
	value *GetEventForwarderOutput
	isSet bool
}

func (v NullableGetEventForwarderOutput) Get() *GetEventForwarderOutput {
	return v.value
}

func (v *NullableGetEventForwarderOutput) Set(val *GetEventForwarderOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventForwarderOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventForwarderOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventForwarderOutput(val *GetEventForwarderOutput) *NullableGetEventForwarderOutput {
	return &NullableGetEventForwarderOutput{value: val, isSet: true}
}

func (v NullableGetEventForwarderOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventForwarderOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


