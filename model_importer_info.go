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

// ImporterInfo struct for ImporterInfo
type ImporterInfo struct {
	ExternalItemId *string `json:"external_item_id,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewImporterInfo instantiates a new ImporterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImporterInfo() *ImporterInfo {
	this := ImporterInfo{}
	return &this
}

// NewImporterInfoWithDefaults instantiates a new ImporterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImporterInfoWithDefaults() *ImporterInfo {
	this := ImporterInfo{}
	return &this
}

// GetExternalItemId returns the ExternalItemId field value if set, zero value otherwise.
func (o *ImporterInfo) GetExternalItemId() string {
	if o == nil || o.ExternalItemId == nil {
		var ret string
		return ret
	}
	return *o.ExternalItemId
}

// GetExternalItemIdOk returns a tuple with the ExternalItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImporterInfo) GetExternalItemIdOk() (*string, bool) {
	if o == nil || o.ExternalItemId == nil {
		return nil, false
	}
	return o.ExternalItemId, true
}

// HasExternalItemId returns a boolean if a field has been set.
func (o *ImporterInfo) HasExternalItemId() bool {
	if o != nil && o.ExternalItemId != nil {
		return true
	}

	return false
}

// SetExternalItemId gets a reference to the given string and assigns it to the ExternalItemId field.
func (o *ImporterInfo) SetExternalItemId(v string) {
	o.ExternalItemId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ImporterInfo) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImporterInfo) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ImporterInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ImporterInfo) SetVersion(v int64) {
	o.Version = &v
}

func (o ImporterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalItemId != nil {
		toSerialize["external_item_id"] = o.ExternalItemId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableImporterInfo struct {
	value *ImporterInfo
	isSet bool
}

func (v NullableImporterInfo) Get() *ImporterInfo {
	return v.value
}

func (v *NullableImporterInfo) Set(val *ImporterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableImporterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableImporterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImporterInfo(val *ImporterInfo) *NullableImporterInfo {
	return &NullableImporterInfo{value: val, isSet: true}
}

func (v NullableImporterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImporterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


