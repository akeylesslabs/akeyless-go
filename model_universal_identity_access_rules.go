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

// checks if the UniversalIdentityAccessRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniversalIdentityAccessRules{}

// UniversalIdentityAccessRules struct for UniversalIdentityAccessRules
type UniversalIdentityAccessRules struct {
	DenyInheritance *bool `json:"deny_inheritance,omitempty"`
	DenyRotate *bool `json:"deny_rotate,omitempty"`
	Ttl *int32 `json:"ttl,omitempty"`
}

// NewUniversalIdentityAccessRules instantiates a new UniversalIdentityAccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniversalIdentityAccessRules() *UniversalIdentityAccessRules {
	this := UniversalIdentityAccessRules{}
	return &this
}

// NewUniversalIdentityAccessRulesWithDefaults instantiates a new UniversalIdentityAccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniversalIdentityAccessRulesWithDefaults() *UniversalIdentityAccessRules {
	this := UniversalIdentityAccessRules{}
	return &this
}

// GetDenyInheritance returns the DenyInheritance field value if set, zero value otherwise.
func (o *UniversalIdentityAccessRules) GetDenyInheritance() bool {
	if o == nil || IsNil(o.DenyInheritance) {
		var ret bool
		return ret
	}
	return *o.DenyInheritance
}

// GetDenyInheritanceOk returns a tuple with the DenyInheritance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalIdentityAccessRules) GetDenyInheritanceOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyInheritance) {
		return nil, false
	}
	return o.DenyInheritance, true
}

// HasDenyInheritance returns a boolean if a field has been set.
func (o *UniversalIdentityAccessRules) HasDenyInheritance() bool {
	if o != nil && !IsNil(o.DenyInheritance) {
		return true
	}

	return false
}

// SetDenyInheritance gets a reference to the given bool and assigns it to the DenyInheritance field.
func (o *UniversalIdentityAccessRules) SetDenyInheritance(v bool) {
	o.DenyInheritance = &v
}

// GetDenyRotate returns the DenyRotate field value if set, zero value otherwise.
func (o *UniversalIdentityAccessRules) GetDenyRotate() bool {
	if o == nil || IsNil(o.DenyRotate) {
		var ret bool
		return ret
	}
	return *o.DenyRotate
}

// GetDenyRotateOk returns a tuple with the DenyRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalIdentityAccessRules) GetDenyRotateOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyRotate) {
		return nil, false
	}
	return o.DenyRotate, true
}

// HasDenyRotate returns a boolean if a field has been set.
func (o *UniversalIdentityAccessRules) HasDenyRotate() bool {
	if o != nil && !IsNil(o.DenyRotate) {
		return true
	}

	return false
}

// SetDenyRotate gets a reference to the given bool and assigns it to the DenyRotate field.
func (o *UniversalIdentityAccessRules) SetDenyRotate(v bool) {
	o.DenyRotate = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *UniversalIdentityAccessRules) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalIdentityAccessRules) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *UniversalIdentityAccessRules) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *UniversalIdentityAccessRules) SetTtl(v int32) {
	o.Ttl = &v
}

func (o UniversalIdentityAccessRules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniversalIdentityAccessRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DenyInheritance) {
		toSerialize["deny_inheritance"] = o.DenyInheritance
	}
	if !IsNil(o.DenyRotate) {
		toSerialize["deny_rotate"] = o.DenyRotate
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	return toSerialize, nil
}

type NullableUniversalIdentityAccessRules struct {
	value *UniversalIdentityAccessRules
	isSet bool
}

func (v NullableUniversalIdentityAccessRules) Get() *UniversalIdentityAccessRules {
	return v.value
}

func (v *NullableUniversalIdentityAccessRules) Set(val *UniversalIdentityAccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableUniversalIdentityAccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableUniversalIdentityAccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniversalIdentityAccessRules(val *UniversalIdentityAccessRules) *NullableUniversalIdentityAccessRules {
	return &NullableUniversalIdentityAccessRules{value: val, isSet: true}
}

func (v NullableUniversalIdentityAccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniversalIdentityAccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


