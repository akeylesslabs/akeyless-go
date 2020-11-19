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
)

// GenCustomerFragment struct for GenCustomerFragment
type GenCustomerFragment struct {
	// The Customer Fragment Description
	Description *string `json:"description,omitempty"`
}

// NewGenCustomerFragment instantiates a new GenCustomerFragment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenCustomerFragment() *GenCustomerFragment {
	this := GenCustomerFragment{}
	return &this
}

// NewGenCustomerFragmentWithDefaults instantiates a new GenCustomerFragment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenCustomerFragmentWithDefaults() *GenCustomerFragment {
	this := GenCustomerFragment{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GenCustomerFragment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenCustomerFragment) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GenCustomerFragment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GenCustomerFragment) SetDescription(v string) {
	o.Description = &v
}

func (o GenCustomerFragment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableGenCustomerFragment struct {
	value *GenCustomerFragment
	isSet bool
}

func (v NullableGenCustomerFragment) Get() *GenCustomerFragment {
	return v.value
}

func (v *NullableGenCustomerFragment) Set(val *GenCustomerFragment) {
	v.value = val
	v.isSet = true
}

func (v NullableGenCustomerFragment) IsSet() bool {
	return v.isSet
}

func (v *NullableGenCustomerFragment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenCustomerFragment(val *GenCustomerFragment) *NullableGenCustomerFragment {
	return &NullableGenCustomerFragment{value: val, isSet: true}
}

func (v NullableGenCustomerFragment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenCustomerFragment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


