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

// checks if the GCPSecretsMigration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GCPSecretsMigration{}

// GCPSecretsMigration struct for GCPSecretsMigration
type GCPSecretsMigration struct {
	General *MigrationGeneral `json:"general,omitempty"`
	Payload *GCPPayload `json:"payload,omitempty"`
}

// NewGCPSecretsMigration instantiates a new GCPSecretsMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPSecretsMigration() *GCPSecretsMigration {
	this := GCPSecretsMigration{}
	return &this
}

// NewGCPSecretsMigrationWithDefaults instantiates a new GCPSecretsMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPSecretsMigrationWithDefaults() *GCPSecretsMigration {
	this := GCPSecretsMigration{}
	return &this
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *GCPSecretsMigration) GetGeneral() MigrationGeneral {
	if o == nil || IsNil(o.General) {
		var ret MigrationGeneral
		return ret
	}
	return *o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSecretsMigration) GetGeneralOk() (*MigrationGeneral, bool) {
	if o == nil || IsNil(o.General) {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *GCPSecretsMigration) HasGeneral() bool {
	if o != nil && !IsNil(o.General) {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given MigrationGeneral and assigns it to the General field.
func (o *GCPSecretsMigration) SetGeneral(v MigrationGeneral) {
	o.General = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GCPSecretsMigration) GetPayload() GCPPayload {
	if o == nil || IsNil(o.Payload) {
		var ret GCPPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSecretsMigration) GetPayloadOk() (*GCPPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GCPSecretsMigration) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given GCPPayload and assigns it to the Payload field.
func (o *GCPSecretsMigration) SetPayload(v GCPPayload) {
	o.Payload = &v
}

func (o GCPSecretsMigration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GCPSecretsMigration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.General) {
		toSerialize["general"] = o.General
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableGCPSecretsMigration struct {
	value *GCPSecretsMigration
	isSet bool
}

func (v NullableGCPSecretsMigration) Get() *GCPSecretsMigration {
	return v.value
}

func (v *NullableGCPSecretsMigration) Set(val *GCPSecretsMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPSecretsMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPSecretsMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPSecretsMigration(val *GCPSecretsMigration) *NullableGCPSecretsMigration {
	return &NullableGCPSecretsMigration{value: val, isSet: true}
}

func (v NullableGCPSecretsMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPSecretsMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


