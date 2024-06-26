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

// RotationSecretMaxInterval struct for RotationSecretMaxInterval
type RotationSecretMaxInterval struct {
	Enable *bool `json:"enable,omitempty"`
	MaxIntervalByDays *int32 `json:"max_interval_by_days,omitempty"`
}

// NewRotationSecretMaxInterval instantiates a new RotationSecretMaxInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotationSecretMaxInterval() *RotationSecretMaxInterval {
	this := RotationSecretMaxInterval{}
	return &this
}

// NewRotationSecretMaxIntervalWithDefaults instantiates a new RotationSecretMaxInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotationSecretMaxIntervalWithDefaults() *RotationSecretMaxInterval {
	this := RotationSecretMaxInterval{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *RotationSecretMaxInterval) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotationSecretMaxInterval) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *RotationSecretMaxInterval) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *RotationSecretMaxInterval) SetEnable(v bool) {
	o.Enable = &v
}

// GetMaxIntervalByDays returns the MaxIntervalByDays field value if set, zero value otherwise.
func (o *RotationSecretMaxInterval) GetMaxIntervalByDays() int32 {
	if o == nil || o.MaxIntervalByDays == nil {
		var ret int32
		return ret
	}
	return *o.MaxIntervalByDays
}

// GetMaxIntervalByDaysOk returns a tuple with the MaxIntervalByDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotationSecretMaxInterval) GetMaxIntervalByDaysOk() (*int32, bool) {
	if o == nil || o.MaxIntervalByDays == nil {
		return nil, false
	}
	return o.MaxIntervalByDays, true
}

// HasMaxIntervalByDays returns a boolean if a field has been set.
func (o *RotationSecretMaxInterval) HasMaxIntervalByDays() bool {
	if o != nil && o.MaxIntervalByDays != nil {
		return true
	}

	return false
}

// SetMaxIntervalByDays gets a reference to the given int32 and assigns it to the MaxIntervalByDays field.
func (o *RotationSecretMaxInterval) SetMaxIntervalByDays(v int32) {
	o.MaxIntervalByDays = &v
}

func (o RotationSecretMaxInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.MaxIntervalByDays != nil {
		toSerialize["max_interval_by_days"] = o.MaxIntervalByDays
	}
	return json.Marshal(toSerialize)
}

type NullableRotationSecretMaxInterval struct {
	value *RotationSecretMaxInterval
	isSet bool
}

func (v NullableRotationSecretMaxInterval) Get() *RotationSecretMaxInterval {
	return v.value
}

func (v *NullableRotationSecretMaxInterval) Set(val *RotationSecretMaxInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableRotationSecretMaxInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableRotationSecretMaxInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotationSecretMaxInterval(val *RotationSecretMaxInterval) *NullableRotationSecretMaxInterval {
	return &NullableRotationSecretMaxInterval{value: val, isSet: true}
}

func (v NullableRotationSecretMaxInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotationSecretMaxInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


