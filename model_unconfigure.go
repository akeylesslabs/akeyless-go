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

// Unconfigure struct for Unconfigure
type Unconfigure struct {
	// The profile name to be removed
	Profile *string `json:"profile,omitempty"`
}

// NewUnconfigure instantiates a new Unconfigure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnconfigure() *Unconfigure {
	this := Unconfigure{}
	var profile string = "default"
	this.Profile = &profile
	return &this
}

// NewUnconfigureWithDefaults instantiates a new Unconfigure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnconfigureWithDefaults() *Unconfigure {
	this := Unconfigure{}
	var profile string = "default"
	this.Profile = &profile
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *Unconfigure) GetProfile() string {
	if o == nil || o.Profile == nil {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unconfigure) GetProfileOk() (*string, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *Unconfigure) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *Unconfigure) SetProfile(v string) {
	o.Profile = &v
}

func (o Unconfigure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	return json.Marshal(toSerialize)
}

type NullableUnconfigure struct {
	value *Unconfigure
	isSet bool
}

func (v NullableUnconfigure) Get() *Unconfigure {
	return v.value
}

func (v *NullableUnconfigure) Set(val *Unconfigure) {
	v.value = val
	v.isSet = true
}

func (v NullableUnconfigure) IsSet() bool {
	return v.isSet
}

func (v *NullableUnconfigure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnconfigure(val *Unconfigure) *NullableUnconfigure {
	return &NullableUnconfigure{value: val, isSet: true}
}

func (v NullableUnconfigure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnconfigure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


