/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// Rules struct for Rules
type Rules struct {
	// Is admin
	Admin *bool `json:"admin,omitempty"`
	// The path the rules refers to
	PathRules *[]PathRule `json:"path_rules,omitempty"`
}

// NewRules instantiates a new Rules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRules() *Rules {
	this := Rules{}
	return &this
}

// NewRulesWithDefaults instantiates a new Rules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesWithDefaults() *Rules {
	this := Rules{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *Rules) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rules) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *Rules) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *Rules) SetAdmin(v bool) {
	o.Admin = &v
}

// GetPathRules returns the PathRules field value if set, zero value otherwise.
func (o *Rules) GetPathRules() []PathRule {
	if o == nil || o.PathRules == nil {
		var ret []PathRule
		return ret
	}
	return *o.PathRules
}

// GetPathRulesOk returns a tuple with the PathRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rules) GetPathRulesOk() (*[]PathRule, bool) {
	if o == nil || o.PathRules == nil {
		return nil, false
	}
	return o.PathRules, true
}

// HasPathRules returns a boolean if a field has been set.
func (o *Rules) HasPathRules() bool {
	if o != nil && o.PathRules != nil {
		return true
	}

	return false
}

// SetPathRules gets a reference to the given []PathRule and assigns it to the PathRules field.
func (o *Rules) SetPathRules(v []PathRule) {
	o.PathRules = &v
}

func (o Rules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.PathRules != nil {
		toSerialize["path_rules"] = o.PathRules
	}
	return json.Marshal(toSerialize)
}

type NullableRules struct {
	value *Rules
	isSet bool
}

func (v NullableRules) Get() *Rules {
	return v.value
}

func (v *NullableRules) Set(val *Rules) {
	v.value = val
	v.isSet = true
}

func (v NullableRules) IsSet() bool {
	return v.isSet
}

func (v *NullableRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRules(val *Rules) *NullableRules {
	return &NullableRules{value: val, isSet: true}
}

func (v NullableRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


