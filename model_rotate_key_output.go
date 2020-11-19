/*
 * includes all operations supported by Akeyless.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
	"time"
)

// RotateKeyOutput RotateKeyOutput defines output of RotateKey operation
type RotateKeyOutput struct {
	NewItemVersion *int32 `json:"new_item_version,omitempty"`
	NextRotationDate *time.Time `json:"next_rotation_date,omitempty"`
}

// NewRotateKeyOutput instantiates a new RotateKeyOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotateKeyOutput() *RotateKeyOutput {
	this := RotateKeyOutput{}
	return &this
}

// NewRotateKeyOutputWithDefaults instantiates a new RotateKeyOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotateKeyOutputWithDefaults() *RotateKeyOutput {
	this := RotateKeyOutput{}
	return &this
}

// GetNewItemVersion returns the NewItemVersion field value if set, zero value otherwise.
func (o *RotateKeyOutput) GetNewItemVersion() int32 {
	if o == nil || o.NewItemVersion == nil {
		var ret int32
		return ret
	}
	return *o.NewItemVersion
}

// GetNewItemVersionOk returns a tuple with the NewItemVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotateKeyOutput) GetNewItemVersionOk() (*int32, bool) {
	if o == nil || o.NewItemVersion == nil {
		return nil, false
	}
	return o.NewItemVersion, true
}

// HasNewItemVersion returns a boolean if a field has been set.
func (o *RotateKeyOutput) HasNewItemVersion() bool {
	if o != nil && o.NewItemVersion != nil {
		return true
	}

	return false
}

// SetNewItemVersion gets a reference to the given int32 and assigns it to the NewItemVersion field.
func (o *RotateKeyOutput) SetNewItemVersion(v int32) {
	o.NewItemVersion = &v
}

// GetNextRotationDate returns the NextRotationDate field value if set, zero value otherwise.
func (o *RotateKeyOutput) GetNextRotationDate() time.Time {
	if o == nil || o.NextRotationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.NextRotationDate
}

// GetNextRotationDateOk returns a tuple with the NextRotationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotateKeyOutput) GetNextRotationDateOk() (*time.Time, bool) {
	if o == nil || o.NextRotationDate == nil {
		return nil, false
	}
	return o.NextRotationDate, true
}

// HasNextRotationDate returns a boolean if a field has been set.
func (o *RotateKeyOutput) HasNextRotationDate() bool {
	if o != nil && o.NextRotationDate != nil {
		return true
	}

	return false
}

// SetNextRotationDate gets a reference to the given time.Time and assigns it to the NextRotationDate field.
func (o *RotateKeyOutput) SetNextRotationDate(v time.Time) {
	o.NextRotationDate = &v
}

func (o RotateKeyOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewItemVersion != nil {
		toSerialize["new_item_version"] = o.NewItemVersion
	}
	if o.NextRotationDate != nil {
		toSerialize["next_rotation_date"] = o.NextRotationDate
	}
	return json.Marshal(toSerialize)
}

type NullableRotateKeyOutput struct {
	value *RotateKeyOutput
	isSet bool
}

func (v NullableRotateKeyOutput) Get() *RotateKeyOutput {
	return v.value
}

func (v *NullableRotateKeyOutput) Set(val *RotateKeyOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableRotateKeyOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRotateKeyOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotateKeyOutput(val *RotateKeyOutput) *NullableRotateKeyOutput {
	return &NullableRotateKeyOutput{value: val, isSet: true}
}

func (v NullableRotateKeyOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotateKeyOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


