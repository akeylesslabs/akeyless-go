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

// ValidateTokenOutput struct for ValidateTokenOutput
type ValidateTokenOutput struct {
	Expiration *string `json:"expiration,omitempty"`
	IsValid *bool `json:"is_valid,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewValidateTokenOutput instantiates a new ValidateTokenOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateTokenOutput() *ValidateTokenOutput {
	this := ValidateTokenOutput{}
	return &this
}

// NewValidateTokenOutputWithDefaults instantiates a new ValidateTokenOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateTokenOutputWithDefaults() *ValidateTokenOutput {
	this := ValidateTokenOutput{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ValidateTokenOutput) GetExpiration() string {
	if o == nil || o.Expiration == nil {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateTokenOutput) GetExpirationOk() (*string, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ValidateTokenOutput) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *ValidateTokenOutput) SetExpiration(v string) {
	o.Expiration = &v
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *ValidateTokenOutput) GetIsValid() bool {
	if o == nil || o.IsValid == nil {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateTokenOutput) GetIsValidOk() (*bool, bool) {
	if o == nil || o.IsValid == nil {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *ValidateTokenOutput) HasIsValid() bool {
	if o != nil && o.IsValid != nil {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *ValidateTokenOutput) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ValidateTokenOutput) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateTokenOutput) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ValidateTokenOutput) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ValidateTokenOutput) SetReason(v string) {
	o.Reason = &v
}

func (o ValidateTokenOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if o.IsValid != nil {
		toSerialize["is_valid"] = o.IsValid
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableValidateTokenOutput struct {
	value *ValidateTokenOutput
	isSet bool
}

func (v NullableValidateTokenOutput) Get() *ValidateTokenOutput {
	return v.value
}

func (v *NullableValidateTokenOutput) Set(val *ValidateTokenOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateTokenOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateTokenOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateTokenOutput(val *ValidateTokenOutput) *NullableValidateTokenOutput {
	return &NullableValidateTokenOutput{value: val, isSet: true}
}

func (v NullableValidateTokenOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateTokenOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


