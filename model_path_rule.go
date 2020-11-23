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

// PathRule struct for PathRule
type PathRule struct {
	// The approved/denied capabilities in the path
	Capabilities *[]string `json:"capabilities,omitempty"`
	// The path the rule refers to
	Path *string `json:"path,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewPathRule instantiates a new PathRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathRule() *PathRule {
	this := PathRule{}
	return &this
}

// NewPathRuleWithDefaults instantiates a new PathRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathRuleWithDefaults() *PathRule {
	this := PathRule{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *PathRule) GetCapabilities() []string {
	if o == nil || o.Capabilities == nil {
		var ret []string
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetCapabilitiesOk() (*[]string, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *PathRule) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *PathRule) SetCapabilities(v []string) {
	o.Capabilities = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PathRule) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PathRule) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PathRule) SetPath(v string) {
	o.Path = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PathRule) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathRule) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PathRule) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PathRule) SetType(v string) {
	o.Type = &v
}

func (o PathRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePathRule struct {
	value *PathRule
	isSet bool
}

func (v NullablePathRule) Get() *PathRule {
	return v.value
}

func (v *NullablePathRule) Set(val *PathRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePathRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePathRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathRule(val *PathRule) *NullablePathRule {
	return &NullablePathRule{value: val, isSet: true}
}

func (v NullablePathRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


