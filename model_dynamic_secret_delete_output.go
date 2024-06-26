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

// DynamicSecretDeleteOutput struct for DynamicSecretDeleteOutput
type DynamicSecretDeleteOutput struct {
	DynamicSecretName *string `json:"dynamic_secret_name,omitempty"`
}

// NewDynamicSecretDeleteOutput instantiates a new DynamicSecretDeleteOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretDeleteOutput() *DynamicSecretDeleteOutput {
	this := DynamicSecretDeleteOutput{}
	return &this
}

// NewDynamicSecretDeleteOutputWithDefaults instantiates a new DynamicSecretDeleteOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretDeleteOutputWithDefaults() *DynamicSecretDeleteOutput {
	this := DynamicSecretDeleteOutput{}
	return &this
}

// GetDynamicSecretName returns the DynamicSecretName field value if set, zero value otherwise.
func (o *DynamicSecretDeleteOutput) GetDynamicSecretName() string {
	if o == nil || o.DynamicSecretName == nil {
		var ret string
		return ret
	}
	return *o.DynamicSecretName
}

// GetDynamicSecretNameOk returns a tuple with the DynamicSecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretDeleteOutput) GetDynamicSecretNameOk() (*string, bool) {
	if o == nil || o.DynamicSecretName == nil {
		return nil, false
	}
	return o.DynamicSecretName, true
}

// HasDynamicSecretName returns a boolean if a field has been set.
func (o *DynamicSecretDeleteOutput) HasDynamicSecretName() bool {
	if o != nil && o.DynamicSecretName != nil {
		return true
	}

	return false
}

// SetDynamicSecretName gets a reference to the given string and assigns it to the DynamicSecretName field.
func (o *DynamicSecretDeleteOutput) SetDynamicSecretName(v string) {
	o.DynamicSecretName = &v
}

func (o DynamicSecretDeleteOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DynamicSecretName != nil {
		toSerialize["dynamic_secret_name"] = o.DynamicSecretName
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretDeleteOutput struct {
	value *DynamicSecretDeleteOutput
	isSet bool
}

func (v NullableDynamicSecretDeleteOutput) Get() *DynamicSecretDeleteOutput {
	return v.value
}

func (v *NullableDynamicSecretDeleteOutput) Set(val *DynamicSecretDeleteOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretDeleteOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretDeleteOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretDeleteOutput(val *DynamicSecretDeleteOutput) *NullableDynamicSecretDeleteOutput {
	return &NullableDynamicSecretDeleteOutput{value: val, isSet: true}
}

func (v NullableDynamicSecretDeleteOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretDeleteOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


