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

// DynamicSecretUpdateOutput struct for DynamicSecretUpdateOutput
type DynamicSecretUpdateOutput struct {
	DynamicSecretDetails *DSProducerDetails `json:"dynamic_secret_details,omitempty"`
}

// NewDynamicSecretUpdateOutput instantiates a new DynamicSecretUpdateOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdateOutput() *DynamicSecretUpdateOutput {
	this := DynamicSecretUpdateOutput{}
	return &this
}

// NewDynamicSecretUpdateOutputWithDefaults instantiates a new DynamicSecretUpdateOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdateOutputWithDefaults() *DynamicSecretUpdateOutput {
	this := DynamicSecretUpdateOutput{}
	return &this
}

// GetDynamicSecretDetails returns the DynamicSecretDetails field value if set, zero value otherwise.
func (o *DynamicSecretUpdateOutput) GetDynamicSecretDetails() DSProducerDetails {
	if o == nil || o.DynamicSecretDetails == nil {
		var ret DSProducerDetails
		return ret
	}
	return *o.DynamicSecretDetails
}

// GetDynamicSecretDetailsOk returns a tuple with the DynamicSecretDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateOutput) GetDynamicSecretDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || o.DynamicSecretDetails == nil {
		return nil, false
	}
	return o.DynamicSecretDetails, true
}

// HasDynamicSecretDetails returns a boolean if a field has been set.
func (o *DynamicSecretUpdateOutput) HasDynamicSecretDetails() bool {
	if o != nil && o.DynamicSecretDetails != nil {
		return true
	}

	return false
}

// SetDynamicSecretDetails gets a reference to the given DSProducerDetails and assigns it to the DynamicSecretDetails field.
func (o *DynamicSecretUpdateOutput) SetDynamicSecretDetails(v DSProducerDetails) {
	o.DynamicSecretDetails = &v
}

func (o DynamicSecretUpdateOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DynamicSecretDetails != nil {
		toSerialize["dynamic_secret_details"] = o.DynamicSecretDetails
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretUpdateOutput struct {
	value *DynamicSecretUpdateOutput
	isSet bool
}

func (v NullableDynamicSecretUpdateOutput) Get() *DynamicSecretUpdateOutput {
	return v.value
}

func (v *NullableDynamicSecretUpdateOutput) Set(val *DynamicSecretUpdateOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdateOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdateOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdateOutput(val *DynamicSecretUpdateOutput) *NullableDynamicSecretUpdateOutput {
	return &NullableDynamicSecretUpdateOutput{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdateOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdateOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

