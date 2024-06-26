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

// CustomerFragmentsConfigJson struct for CustomerFragmentsConfigJson
type CustomerFragmentsConfigJson struct {
	CustomerFragments *[]CustomerFragmentConfig `json:"customer_fragments,omitempty"`
}

// NewCustomerFragmentsConfigJson instantiates a new CustomerFragmentsConfigJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFragmentsConfigJson() *CustomerFragmentsConfigJson {
	this := CustomerFragmentsConfigJson{}
	return &this
}

// NewCustomerFragmentsConfigJsonWithDefaults instantiates a new CustomerFragmentsConfigJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFragmentsConfigJsonWithDefaults() *CustomerFragmentsConfigJson {
	this := CustomerFragmentsConfigJson{}
	return &this
}

// GetCustomerFragments returns the CustomerFragments field value if set, zero value otherwise.
func (o *CustomerFragmentsConfigJson) GetCustomerFragments() []CustomerFragmentConfig {
	if o == nil || o.CustomerFragments == nil {
		var ret []CustomerFragmentConfig
		return ret
	}
	return *o.CustomerFragments
}

// GetCustomerFragmentsOk returns a tuple with the CustomerFragments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFragmentsConfigJson) GetCustomerFragmentsOk() (*[]CustomerFragmentConfig, bool) {
	if o == nil || o.CustomerFragments == nil {
		return nil, false
	}
	return o.CustomerFragments, true
}

// HasCustomerFragments returns a boolean if a field has been set.
func (o *CustomerFragmentsConfigJson) HasCustomerFragments() bool {
	if o != nil && o.CustomerFragments != nil {
		return true
	}

	return false
}

// SetCustomerFragments gets a reference to the given []CustomerFragmentConfig and assigns it to the CustomerFragments field.
func (o *CustomerFragmentsConfigJson) SetCustomerFragments(v []CustomerFragmentConfig) {
	o.CustomerFragments = &v
}

func (o CustomerFragmentsConfigJson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerFragments != nil {
		toSerialize["customer_fragments"] = o.CustomerFragments
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerFragmentsConfigJson struct {
	value *CustomerFragmentsConfigJson
	isSet bool
}

func (v NullableCustomerFragmentsConfigJson) Get() *CustomerFragmentsConfigJson {
	return v.value
}

func (v *NullableCustomerFragmentsConfigJson) Set(val *CustomerFragmentsConfigJson) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFragmentsConfigJson) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFragmentsConfigJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFragmentsConfigJson(val *CustomerFragmentsConfigJson) *NullableCustomerFragmentsConfigJson {
	return &NullableCustomerFragmentsConfigJson{value: val, isSet: true}
}

func (v NullableCustomerFragmentsConfigJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFragmentsConfigJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


