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

// DataProtectionSection We need the fields to be pointers as we use the same struct for partial updates as well
type DataProtectionSection struct {
	EnableClassicKeyProtection *bool `json:"enable_classic_key_protection,omitempty"`
}

// NewDataProtectionSection instantiates a new DataProtectionSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataProtectionSection() *DataProtectionSection {
	this := DataProtectionSection{}
	return &this
}

// NewDataProtectionSectionWithDefaults instantiates a new DataProtectionSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataProtectionSectionWithDefaults() *DataProtectionSection {
	this := DataProtectionSection{}
	return &this
}

// GetEnableClassicKeyProtection returns the EnableClassicKeyProtection field value if set, zero value otherwise.
func (o *DataProtectionSection) GetEnableClassicKeyProtection() bool {
	if o == nil || o.EnableClassicKeyProtection == nil {
		var ret bool
		return ret
	}
	return *o.EnableClassicKeyProtection
}

// GetEnableClassicKeyProtectionOk returns a tuple with the EnableClassicKeyProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProtectionSection) GetEnableClassicKeyProtectionOk() (*bool, bool) {
	if o == nil || o.EnableClassicKeyProtection == nil {
		return nil, false
	}
	return o.EnableClassicKeyProtection, true
}

// HasEnableClassicKeyProtection returns a boolean if a field has been set.
func (o *DataProtectionSection) HasEnableClassicKeyProtection() bool {
	if o != nil && o.EnableClassicKeyProtection != nil {
		return true
	}

	return false
}

// SetEnableClassicKeyProtection gets a reference to the given bool and assigns it to the EnableClassicKeyProtection field.
func (o *DataProtectionSection) SetEnableClassicKeyProtection(v bool) {
	o.EnableClassicKeyProtection = &v
}

func (o DataProtectionSection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableClassicKeyProtection != nil {
		toSerialize["enable_classic_key_protection"] = o.EnableClassicKeyProtection
	}
	return json.Marshal(toSerialize)
}

type NullableDataProtectionSection struct {
	value *DataProtectionSection
	isSet bool
}

func (v NullableDataProtectionSection) Get() *DataProtectionSection {
	return v.value
}

func (v *NullableDataProtectionSection) Set(val *DataProtectionSection) {
	v.value = val
	v.isSet = true
}

func (v NullableDataProtectionSection) IsSet() bool {
	return v.isSet
}

func (v *NullableDataProtectionSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataProtectionSection(val *DataProtectionSection) *NullableDataProtectionSection {
	return &NullableDataProtectionSection{value: val, isSet: true}
}

func (v NullableDataProtectionSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataProtectionSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

