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

// DeleteRoleRuleOutput struct for DeleteRoleRuleOutput
type DeleteRoleRuleOutput struct {
	Result *string `json:"result,omitempty"`
}

// NewDeleteRoleRuleOutput instantiates a new DeleteRoleRuleOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRoleRuleOutput() *DeleteRoleRuleOutput {
	this := DeleteRoleRuleOutput{}
	return &this
}

// NewDeleteRoleRuleOutputWithDefaults instantiates a new DeleteRoleRuleOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRoleRuleOutputWithDefaults() *DeleteRoleRuleOutput {
	this := DeleteRoleRuleOutput{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DeleteRoleRuleOutput) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRoleRuleOutput) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DeleteRoleRuleOutput) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DeleteRoleRuleOutput) SetResult(v string) {
	o.Result = &v
}

func (o DeleteRoleRuleOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteRoleRuleOutput struct {
	value *DeleteRoleRuleOutput
	isSet bool
}

func (v NullableDeleteRoleRuleOutput) Get() *DeleteRoleRuleOutput {
	return v.value
}

func (v *NullableDeleteRoleRuleOutput) Set(val *DeleteRoleRuleOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRoleRuleOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRoleRuleOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRoleRuleOutput(val *DeleteRoleRuleOutput) *NullableDeleteRoleRuleOutput {
	return &NullableDeleteRoleRuleOutput{value: val, isSet: true}
}

func (v NullableDeleteRoleRuleOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRoleRuleOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


