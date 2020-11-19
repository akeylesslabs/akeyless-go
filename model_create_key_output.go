/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 0.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// CreateKeyOutput struct for CreateKeyOutput
type CreateKeyOutput struct {
	FragmentResults *[]int64 `json:"fragment_results,omitempty"`
}

// NewCreateKeyOutput instantiates a new CreateKeyOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeyOutput() *CreateKeyOutput {
	this := CreateKeyOutput{}
	return &this
}

// NewCreateKeyOutputWithDefaults instantiates a new CreateKeyOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeyOutputWithDefaults() *CreateKeyOutput {
	this := CreateKeyOutput{}
	return &this
}

// GetFragmentResults returns the FragmentResults field value if set, zero value otherwise.
func (o *CreateKeyOutput) GetFragmentResults() []int64 {
	if o == nil || o.FragmentResults == nil {
		var ret []int64
		return ret
	}
	return *o.FragmentResults
}

// GetFragmentResultsOk returns a tuple with the FragmentResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKeyOutput) GetFragmentResultsOk() (*[]int64, bool) {
	if o == nil || o.FragmentResults == nil {
		return nil, false
	}
	return o.FragmentResults, true
}

// HasFragmentResults returns a boolean if a field has been set.
func (o *CreateKeyOutput) HasFragmentResults() bool {
	if o != nil && o.FragmentResults != nil {
		return true
	}

	return false
}

// SetFragmentResults gets a reference to the given []int64 and assigns it to the FragmentResults field.
func (o *CreateKeyOutput) SetFragmentResults(v []int64) {
	o.FragmentResults = &v
}

func (o CreateKeyOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FragmentResults != nil {
		toSerialize["fragment_results"] = o.FragmentResults
	}
	return json.Marshal(toSerialize)
}

type NullableCreateKeyOutput struct {
	value *CreateKeyOutput
	isSet bool
}

func (v NullableCreateKeyOutput) Get() *CreateKeyOutput {
	return v.value
}

func (v *NullableCreateKeyOutput) Set(val *CreateKeyOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeyOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeyOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeyOutput(val *CreateKeyOutput) *NullableCreateKeyOutput {
	return &NullableCreateKeyOutput{value: val, isSet: true}
}

func (v NullableCreateKeyOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeyOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


