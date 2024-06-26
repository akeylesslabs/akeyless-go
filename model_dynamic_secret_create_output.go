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

// DynamicSecretCreateOutput struct for DynamicSecretCreateOutput
type DynamicSecretCreateOutput struct {
	DynamicSecretDetails *DSProducerDetails `json:"dynamic_secret_details,omitempty"`
}

// NewDynamicSecretCreateOutput instantiates a new DynamicSecretCreateOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretCreateOutput() *DynamicSecretCreateOutput {
	this := DynamicSecretCreateOutput{}
	return &this
}

// NewDynamicSecretCreateOutputWithDefaults instantiates a new DynamicSecretCreateOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretCreateOutputWithDefaults() *DynamicSecretCreateOutput {
	this := DynamicSecretCreateOutput{}
	return &this
}

// GetDynamicSecretDetails returns the DynamicSecretDetails field value if set, zero value otherwise.
func (o *DynamicSecretCreateOutput) GetDynamicSecretDetails() DSProducerDetails {
	if o == nil || o.DynamicSecretDetails == nil {
		var ret DSProducerDetails
		return ret
	}
	return *o.DynamicSecretDetails
}

// GetDynamicSecretDetailsOk returns a tuple with the DynamicSecretDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateOutput) GetDynamicSecretDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || o.DynamicSecretDetails == nil {
		return nil, false
	}
	return o.DynamicSecretDetails, true
}

// HasDynamicSecretDetails returns a boolean if a field has been set.
func (o *DynamicSecretCreateOutput) HasDynamicSecretDetails() bool {
	if o != nil && o.DynamicSecretDetails != nil {
		return true
	}

	return false
}

// SetDynamicSecretDetails gets a reference to the given DSProducerDetails and assigns it to the DynamicSecretDetails field.
func (o *DynamicSecretCreateOutput) SetDynamicSecretDetails(v DSProducerDetails) {
	o.DynamicSecretDetails = &v
}

func (o DynamicSecretCreateOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DynamicSecretDetails != nil {
		toSerialize["dynamic_secret_details"] = o.DynamicSecretDetails
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretCreateOutput struct {
	value *DynamicSecretCreateOutput
	isSet bool
}

func (v NullableDynamicSecretCreateOutput) Get() *DynamicSecretCreateOutput {
	return v.value
}

func (v *NullableDynamicSecretCreateOutput) Set(val *DynamicSecretCreateOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretCreateOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretCreateOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretCreateOutput(val *DynamicSecretCreateOutput) *NullableDynamicSecretCreateOutput {
	return &NullableDynamicSecretCreateOutput{value: val, isSet: true}
}

func (v NullableDynamicSecretCreateOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretCreateOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


