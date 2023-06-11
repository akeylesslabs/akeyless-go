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

// checks if the OnePasswordMigration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnePasswordMigration{}

// OnePasswordMigration struct for OnePasswordMigration
type OnePasswordMigration struct {
	General *MigrationGeneral `json:"general,omitempty"`
	Payload *OnePasswordPayload `json:"payload,omitempty"`
}

// NewOnePasswordMigration instantiates a new OnePasswordMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnePasswordMigration() *OnePasswordMigration {
	this := OnePasswordMigration{}
	return &this
}

// NewOnePasswordMigrationWithDefaults instantiates a new OnePasswordMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnePasswordMigrationWithDefaults() *OnePasswordMigration {
	this := OnePasswordMigration{}
	return &this
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *OnePasswordMigration) GetGeneral() MigrationGeneral {
	if o == nil || IsNil(o.General) {
		var ret MigrationGeneral
		return ret
	}
	return *o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnePasswordMigration) GetGeneralOk() (*MigrationGeneral, bool) {
	if o == nil || IsNil(o.General) {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *OnePasswordMigration) HasGeneral() bool {
	if o != nil && !IsNil(o.General) {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given MigrationGeneral and assigns it to the General field.
func (o *OnePasswordMigration) SetGeneral(v MigrationGeneral) {
	o.General = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *OnePasswordMigration) GetPayload() OnePasswordPayload {
	if o == nil || IsNil(o.Payload) {
		var ret OnePasswordPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnePasswordMigration) GetPayloadOk() (*OnePasswordPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *OnePasswordMigration) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given OnePasswordPayload and assigns it to the Payload field.
func (o *OnePasswordMigration) SetPayload(v OnePasswordPayload) {
	o.Payload = &v
}

func (o OnePasswordMigration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnePasswordMigration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.General) {
		toSerialize["general"] = o.General
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableOnePasswordMigration struct {
	value *OnePasswordMigration
	isSet bool
}

func (v NullableOnePasswordMigration) Get() *OnePasswordMigration {
	return v.value
}

func (v *NullableOnePasswordMigration) Set(val *OnePasswordMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableOnePasswordMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableOnePasswordMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnePasswordMigration(val *OnePasswordMigration) *NullableOnePasswordMigration {
	return &NullableOnePasswordMigration{value: val, isSet: true}
}

func (v NullableOnePasswordMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnePasswordMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


