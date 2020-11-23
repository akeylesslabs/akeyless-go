/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// CustomerFragmentsJson struct for CustomerFragmentsJson
type CustomerFragmentsJson struct {
	CustomerFragments *[]CustomerFragment `json:"customer_fragments,omitempty"`
}

// NewCustomerFragmentsJson instantiates a new CustomerFragmentsJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFragmentsJson() *CustomerFragmentsJson {
	this := CustomerFragmentsJson{}
	return &this
}

// NewCustomerFragmentsJsonWithDefaults instantiates a new CustomerFragmentsJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFragmentsJsonWithDefaults() *CustomerFragmentsJson {
	this := CustomerFragmentsJson{}
	return &this
}

// GetCustomerFragments returns the CustomerFragments field value if set, zero value otherwise.
func (o *CustomerFragmentsJson) GetCustomerFragments() []CustomerFragment {
	if o == nil || o.CustomerFragments == nil {
		var ret []CustomerFragment
		return ret
	}
	return *o.CustomerFragments
}

// GetCustomerFragmentsOk returns a tuple with the CustomerFragments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFragmentsJson) GetCustomerFragmentsOk() (*[]CustomerFragment, bool) {
	if o == nil || o.CustomerFragments == nil {
		return nil, false
	}
	return o.CustomerFragments, true
}

// HasCustomerFragments returns a boolean if a field has been set.
func (o *CustomerFragmentsJson) HasCustomerFragments() bool {
	if o != nil && o.CustomerFragments != nil {
		return true
	}

	return false
}

// SetCustomerFragments gets a reference to the given []CustomerFragment and assigns it to the CustomerFragments field.
func (o *CustomerFragmentsJson) SetCustomerFragments(v []CustomerFragment) {
	o.CustomerFragments = &v
}

func (o CustomerFragmentsJson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerFragments != nil {
		toSerialize["customer_fragments"] = o.CustomerFragments
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerFragmentsJson struct {
	value *CustomerFragmentsJson
	isSet bool
}

func (v NullableCustomerFragmentsJson) Get() *CustomerFragmentsJson {
	return v.value
}

func (v *NullableCustomerFragmentsJson) Set(val *CustomerFragmentsJson) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFragmentsJson) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFragmentsJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFragmentsJson(val *CustomerFragmentsJson) *NullableCustomerFragmentsJson {
	return &NullableCustomerFragmentsJson{value: val, isSet: true}
}

func (v NullableCustomerFragmentsJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFragmentsJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


