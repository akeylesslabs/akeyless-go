/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// checks if the AllowedAccessDeleteArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllowedAccessDeleteArgs{}

// AllowedAccessDeleteArgs struct for AllowedAccessDeleteArgs
type AllowedAccessDeleteArgs struct {
	// Allowed access name to delete
	Name string `json:"name"`
}

// NewAllowedAccessDeleteArgs instantiates a new AllowedAccessDeleteArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedAccessDeleteArgs(name string) *AllowedAccessDeleteArgs {
	this := AllowedAccessDeleteArgs{}
	this.Name = name
	return &this
}

// NewAllowedAccessDeleteArgsWithDefaults instantiates a new AllowedAccessDeleteArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedAccessDeleteArgsWithDefaults() *AllowedAccessDeleteArgs {
	this := AllowedAccessDeleteArgs{}
	return &this
}

// GetName returns the Name field value
func (o *AllowedAccessDeleteArgs) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AllowedAccessDeleteArgs) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AllowedAccessDeleteArgs) SetName(v string) {
	o.Name = v
}

func (o AllowedAccessDeleteArgs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllowedAccessDeleteArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAllowedAccessDeleteArgs struct {
	value *AllowedAccessDeleteArgs
	isSet bool
}

func (v NullableAllowedAccessDeleteArgs) Get() *AllowedAccessDeleteArgs {
	return v.value
}

func (v *NullableAllowedAccessDeleteArgs) Set(val *AllowedAccessDeleteArgs) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedAccessDeleteArgs) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedAccessDeleteArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedAccessDeleteArgs(val *AllowedAccessDeleteArgs) *NullableAllowedAccessDeleteArgs {
	return &NullableAllowedAccessDeleteArgs{value: val, isSet: true}
}

func (v NullableAllowedAccessDeleteArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedAccessDeleteArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


