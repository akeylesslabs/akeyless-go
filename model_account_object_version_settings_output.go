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

// AccountObjectVersionSettingsOutput struct for AccountObjectVersionSettingsOutput
type AccountObjectVersionSettingsOutput struct {
	DefaultVersioning *bool `json:"default-versioning,omitempty"`
	ForceNewVersions *bool `json:"force_new_versions,omitempty"`
	Items *[]ObjectVersionSettingsOutput `json:"items,omitempty"`
}

// NewAccountObjectVersionSettingsOutput instantiates a new AccountObjectVersionSettingsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountObjectVersionSettingsOutput() *AccountObjectVersionSettingsOutput {
	this := AccountObjectVersionSettingsOutput{}
	return &this
}

// NewAccountObjectVersionSettingsOutputWithDefaults instantiates a new AccountObjectVersionSettingsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountObjectVersionSettingsOutputWithDefaults() *AccountObjectVersionSettingsOutput {
	this := AccountObjectVersionSettingsOutput{}
	return &this
}

// GetDefaultVersioning returns the DefaultVersioning field value if set, zero value otherwise.
func (o *AccountObjectVersionSettingsOutput) GetDefaultVersioning() bool {
	if o == nil || o.DefaultVersioning == nil {
		var ret bool
		return ret
	}
	return *o.DefaultVersioning
}

// GetDefaultVersioningOk returns a tuple with the DefaultVersioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountObjectVersionSettingsOutput) GetDefaultVersioningOk() (*bool, bool) {
	if o == nil || o.DefaultVersioning == nil {
		return nil, false
	}
	return o.DefaultVersioning, true
}

// HasDefaultVersioning returns a boolean if a field has been set.
func (o *AccountObjectVersionSettingsOutput) HasDefaultVersioning() bool {
	if o != nil && o.DefaultVersioning != nil {
		return true
	}

	return false
}

// SetDefaultVersioning gets a reference to the given bool and assigns it to the DefaultVersioning field.
func (o *AccountObjectVersionSettingsOutput) SetDefaultVersioning(v bool) {
	o.DefaultVersioning = &v
}

// GetForceNewVersions returns the ForceNewVersions field value if set, zero value otherwise.
func (o *AccountObjectVersionSettingsOutput) GetForceNewVersions() bool {
	if o == nil || o.ForceNewVersions == nil {
		var ret bool
		return ret
	}
	return *o.ForceNewVersions
}

// GetForceNewVersionsOk returns a tuple with the ForceNewVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountObjectVersionSettingsOutput) GetForceNewVersionsOk() (*bool, bool) {
	if o == nil || o.ForceNewVersions == nil {
		return nil, false
	}
	return o.ForceNewVersions, true
}

// HasForceNewVersions returns a boolean if a field has been set.
func (o *AccountObjectVersionSettingsOutput) HasForceNewVersions() bool {
	if o != nil && o.ForceNewVersions != nil {
		return true
	}

	return false
}

// SetForceNewVersions gets a reference to the given bool and assigns it to the ForceNewVersions field.
func (o *AccountObjectVersionSettingsOutput) SetForceNewVersions(v bool) {
	o.ForceNewVersions = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AccountObjectVersionSettingsOutput) GetItems() []ObjectVersionSettingsOutput {
	if o == nil || o.Items == nil {
		var ret []ObjectVersionSettingsOutput
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountObjectVersionSettingsOutput) GetItemsOk() (*[]ObjectVersionSettingsOutput, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AccountObjectVersionSettingsOutput) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ObjectVersionSettingsOutput and assigns it to the Items field.
func (o *AccountObjectVersionSettingsOutput) SetItems(v []ObjectVersionSettingsOutput) {
	o.Items = &v
}

func (o AccountObjectVersionSettingsOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultVersioning != nil {
		toSerialize["default-versioning"] = o.DefaultVersioning
	}
	if o.ForceNewVersions != nil {
		toSerialize["force_new_versions"] = o.ForceNewVersions
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableAccountObjectVersionSettingsOutput struct {
	value *AccountObjectVersionSettingsOutput
	isSet bool
}

func (v NullableAccountObjectVersionSettingsOutput) Get() *AccountObjectVersionSettingsOutput {
	return v.value
}

func (v *NullableAccountObjectVersionSettingsOutput) Set(val *AccountObjectVersionSettingsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountObjectVersionSettingsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountObjectVersionSettingsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountObjectVersionSettingsOutput(val *AccountObjectVersionSettingsOutput) *NullableAccountObjectVersionSettingsOutput {
	return &NullableAccountObjectVersionSettingsOutput{value: val, isSet: true}
}

func (v NullableAccountObjectVersionSettingsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountObjectVersionSettingsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


