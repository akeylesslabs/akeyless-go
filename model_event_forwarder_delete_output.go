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

// EventForwarderDeleteOutput struct for EventForwarderDeleteOutput
type EventForwarderDeleteOutput struct {
	EventForwarderName *string `json:"event_forwarder_name,omitempty"`
}

// NewEventForwarderDeleteOutput instantiates a new EventForwarderDeleteOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventForwarderDeleteOutput() *EventForwarderDeleteOutput {
	this := EventForwarderDeleteOutput{}
	return &this
}

// NewEventForwarderDeleteOutputWithDefaults instantiates a new EventForwarderDeleteOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventForwarderDeleteOutputWithDefaults() *EventForwarderDeleteOutput {
	this := EventForwarderDeleteOutput{}
	return &this
}

// GetEventForwarderName returns the EventForwarderName field value if set, zero value otherwise.
func (o *EventForwarderDeleteOutput) GetEventForwarderName() string {
	if o == nil || o.EventForwarderName == nil {
		var ret string
		return ret
	}
	return *o.EventForwarderName
}

// GetEventForwarderNameOk returns a tuple with the EventForwarderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderDeleteOutput) GetEventForwarderNameOk() (*string, bool) {
	if o == nil || o.EventForwarderName == nil {
		return nil, false
	}
	return o.EventForwarderName, true
}

// HasEventForwarderName returns a boolean if a field has been set.
func (o *EventForwarderDeleteOutput) HasEventForwarderName() bool {
	if o != nil && o.EventForwarderName != nil {
		return true
	}

	return false
}

// SetEventForwarderName gets a reference to the given string and assigns it to the EventForwarderName field.
func (o *EventForwarderDeleteOutput) SetEventForwarderName(v string) {
	o.EventForwarderName = &v
}

func (o EventForwarderDeleteOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventForwarderName != nil {
		toSerialize["event_forwarder_name"] = o.EventForwarderName
	}
	return json.Marshal(toSerialize)
}

type NullableEventForwarderDeleteOutput struct {
	value *EventForwarderDeleteOutput
	isSet bool
}

func (v NullableEventForwarderDeleteOutput) Get() *EventForwarderDeleteOutput {
	return v.value
}

func (v *NullableEventForwarderDeleteOutput) Set(val *EventForwarderDeleteOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableEventForwarderDeleteOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableEventForwarderDeleteOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventForwarderDeleteOutput(val *EventForwarderDeleteOutput) *NullableEventForwarderDeleteOutput {
	return &NullableEventForwarderDeleteOutput{value: val, isSet: true}
}

func (v NullableEventForwarderDeleteOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventForwarderDeleteOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

