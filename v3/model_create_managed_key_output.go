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

// CreateManagedKeyOutput struct for CreateManagedKeyOutput
type CreateManagedKeyOutput struct {
	ManagedKeyId *int64 `json:"managed_key_id,omitempty"`
	ManagedKeyName *string `json:"managed_key_name,omitempty"`
	ManagedKeyType *string `json:"managed_key_type,omitempty"`
}

// NewCreateManagedKeyOutput instantiates a new CreateManagedKeyOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateManagedKeyOutput() *CreateManagedKeyOutput {
	this := CreateManagedKeyOutput{}
	return &this
}

// NewCreateManagedKeyOutputWithDefaults instantiates a new CreateManagedKeyOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateManagedKeyOutputWithDefaults() *CreateManagedKeyOutput {
	this := CreateManagedKeyOutput{}
	return &this
}

// GetManagedKeyId returns the ManagedKeyId field value if set, zero value otherwise.
func (o *CreateManagedKeyOutput) GetManagedKeyId() int64 {
	if o == nil || o.ManagedKeyId == nil {
		var ret int64
		return ret
	}
	return *o.ManagedKeyId
}

// GetManagedKeyIdOk returns a tuple with the ManagedKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateManagedKeyOutput) GetManagedKeyIdOk() (*int64, bool) {
	if o == nil || o.ManagedKeyId == nil {
		return nil, false
	}
	return o.ManagedKeyId, true
}

// HasManagedKeyId returns a boolean if a field has been set.
func (o *CreateManagedKeyOutput) HasManagedKeyId() bool {
	if o != nil && o.ManagedKeyId != nil {
		return true
	}

	return false
}

// SetManagedKeyId gets a reference to the given int64 and assigns it to the ManagedKeyId field.
func (o *CreateManagedKeyOutput) SetManagedKeyId(v int64) {
	o.ManagedKeyId = &v
}

// GetManagedKeyName returns the ManagedKeyName field value if set, zero value otherwise.
func (o *CreateManagedKeyOutput) GetManagedKeyName() string {
	if o == nil || o.ManagedKeyName == nil {
		var ret string
		return ret
	}
	return *o.ManagedKeyName
}

// GetManagedKeyNameOk returns a tuple with the ManagedKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateManagedKeyOutput) GetManagedKeyNameOk() (*string, bool) {
	if o == nil || o.ManagedKeyName == nil {
		return nil, false
	}
	return o.ManagedKeyName, true
}

// HasManagedKeyName returns a boolean if a field has been set.
func (o *CreateManagedKeyOutput) HasManagedKeyName() bool {
	if o != nil && o.ManagedKeyName != nil {
		return true
	}

	return false
}

// SetManagedKeyName gets a reference to the given string and assigns it to the ManagedKeyName field.
func (o *CreateManagedKeyOutput) SetManagedKeyName(v string) {
	o.ManagedKeyName = &v
}

// GetManagedKeyType returns the ManagedKeyType field value if set, zero value otherwise.
func (o *CreateManagedKeyOutput) GetManagedKeyType() string {
	if o == nil || o.ManagedKeyType == nil {
		var ret string
		return ret
	}
	return *o.ManagedKeyType
}

// GetManagedKeyTypeOk returns a tuple with the ManagedKeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateManagedKeyOutput) GetManagedKeyTypeOk() (*string, bool) {
	if o == nil || o.ManagedKeyType == nil {
		return nil, false
	}
	return o.ManagedKeyType, true
}

// HasManagedKeyType returns a boolean if a field has been set.
func (o *CreateManagedKeyOutput) HasManagedKeyType() bool {
	if o != nil && o.ManagedKeyType != nil {
		return true
	}

	return false
}

// SetManagedKeyType gets a reference to the given string and assigns it to the ManagedKeyType field.
func (o *CreateManagedKeyOutput) SetManagedKeyType(v string) {
	o.ManagedKeyType = &v
}

func (o CreateManagedKeyOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ManagedKeyId != nil {
		toSerialize["managed_key_id"] = o.ManagedKeyId
	}
	if o.ManagedKeyName != nil {
		toSerialize["managed_key_name"] = o.ManagedKeyName
	}
	if o.ManagedKeyType != nil {
		toSerialize["managed_key_type"] = o.ManagedKeyType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateManagedKeyOutput struct {
	value *CreateManagedKeyOutput
	isSet bool
}

func (v NullableCreateManagedKeyOutput) Get() *CreateManagedKeyOutput {
	return v.value
}

func (v *NullableCreateManagedKeyOutput) Set(val *CreateManagedKeyOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateManagedKeyOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateManagedKeyOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateManagedKeyOutput(val *CreateManagedKeyOutput) *NullableCreateManagedKeyOutput {
	return &NullableCreateManagedKeyOutput{value: val, isSet: true}
}

func (v NullableCreateManagedKeyOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateManagedKeyOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

