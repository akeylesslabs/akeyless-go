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

// checks if the ConfigChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigChange{}

// ConfigChange struct for ConfigChange
type ConfigChange struct {
	ConfigHash *ConfigHash `json:"config_hash,omitempty"`
	LastChange *LastConfigChange `json:"last_change,omitempty"`
	LastStatus *LastStatusInfo `json:"last_status,omitempty"`
	RequiredActivity *RequiredActivity `json:"required_activity,omitempty"`
	UpdateStamp *int64 `json:"update_stamp,omitempty"`
}

// NewConfigChange instantiates a new ConfigChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigChange() *ConfigChange {
	this := ConfigChange{}
	return &this
}

// NewConfigChangeWithDefaults instantiates a new ConfigChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigChangeWithDefaults() *ConfigChange {
	this := ConfigChange{}
	return &this
}

// GetConfigHash returns the ConfigHash field value if set, zero value otherwise.
func (o *ConfigChange) GetConfigHash() ConfigHash {
	if o == nil || IsNil(o.ConfigHash) {
		var ret ConfigHash
		return ret
	}
	return *o.ConfigHash
}

// GetConfigHashOk returns a tuple with the ConfigHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigChange) GetConfigHashOk() (*ConfigHash, bool) {
	if o == nil || IsNil(o.ConfigHash) {
		return nil, false
	}
	return o.ConfigHash, true
}

// HasConfigHash returns a boolean if a field has been set.
func (o *ConfigChange) HasConfigHash() bool {
	if o != nil && !IsNil(o.ConfigHash) {
		return true
	}

	return false
}

// SetConfigHash gets a reference to the given ConfigHash and assigns it to the ConfigHash field.
func (o *ConfigChange) SetConfigHash(v ConfigHash) {
	o.ConfigHash = &v
}

// GetLastChange returns the LastChange field value if set, zero value otherwise.
func (o *ConfigChange) GetLastChange() LastConfigChange {
	if o == nil || IsNil(o.LastChange) {
		var ret LastConfigChange
		return ret
	}
	return *o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigChange) GetLastChangeOk() (*LastConfigChange, bool) {
	if o == nil || IsNil(o.LastChange) {
		return nil, false
	}
	return o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *ConfigChange) HasLastChange() bool {
	if o != nil && !IsNil(o.LastChange) {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given LastConfigChange and assigns it to the LastChange field.
func (o *ConfigChange) SetLastChange(v LastConfigChange) {
	o.LastChange = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *ConfigChange) GetLastStatus() LastStatusInfo {
	if o == nil || IsNil(o.LastStatus) {
		var ret LastStatusInfo
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigChange) GetLastStatusOk() (*LastStatusInfo, bool) {
	if o == nil || IsNil(o.LastStatus) {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *ConfigChange) HasLastStatus() bool {
	if o != nil && !IsNil(o.LastStatus) {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given LastStatusInfo and assigns it to the LastStatus field.
func (o *ConfigChange) SetLastStatus(v LastStatusInfo) {
	o.LastStatus = &v
}

// GetRequiredActivity returns the RequiredActivity field value if set, zero value otherwise.
func (o *ConfigChange) GetRequiredActivity() RequiredActivity {
	if o == nil || IsNil(o.RequiredActivity) {
		var ret RequiredActivity
		return ret
	}
	return *o.RequiredActivity
}

// GetRequiredActivityOk returns a tuple with the RequiredActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigChange) GetRequiredActivityOk() (*RequiredActivity, bool) {
	if o == nil || IsNil(o.RequiredActivity) {
		return nil, false
	}
	return o.RequiredActivity, true
}

// HasRequiredActivity returns a boolean if a field has been set.
func (o *ConfigChange) HasRequiredActivity() bool {
	if o != nil && !IsNil(o.RequiredActivity) {
		return true
	}

	return false
}

// SetRequiredActivity gets a reference to the given RequiredActivity and assigns it to the RequiredActivity field.
func (o *ConfigChange) SetRequiredActivity(v RequiredActivity) {
	o.RequiredActivity = &v
}

// GetUpdateStamp returns the UpdateStamp field value if set, zero value otherwise.
func (o *ConfigChange) GetUpdateStamp() int64 {
	if o == nil || IsNil(o.UpdateStamp) {
		var ret int64
		return ret
	}
	return *o.UpdateStamp
}

// GetUpdateStampOk returns a tuple with the UpdateStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigChange) GetUpdateStampOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateStamp) {
		return nil, false
	}
	return o.UpdateStamp, true
}

// HasUpdateStamp returns a boolean if a field has been set.
func (o *ConfigChange) HasUpdateStamp() bool {
	if o != nil && !IsNil(o.UpdateStamp) {
		return true
	}

	return false
}

// SetUpdateStamp gets a reference to the given int64 and assigns it to the UpdateStamp field.
func (o *ConfigChange) SetUpdateStamp(v int64) {
	o.UpdateStamp = &v
}

func (o ConfigChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigHash) {
		toSerialize["config_hash"] = o.ConfigHash
	}
	if !IsNil(o.LastChange) {
		toSerialize["last_change"] = o.LastChange
	}
	if !IsNil(o.LastStatus) {
		toSerialize["last_status"] = o.LastStatus
	}
	if !IsNil(o.RequiredActivity) {
		toSerialize["required_activity"] = o.RequiredActivity
	}
	if !IsNil(o.UpdateStamp) {
		toSerialize["update_stamp"] = o.UpdateStamp
	}
	return toSerialize, nil
}

type NullableConfigChange struct {
	value *ConfigChange
	isSet bool
}

func (v NullableConfigChange) Get() *ConfigChange {
	return v.value
}

func (v *NullableConfigChange) Set(val *ConfigChange) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigChange) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigChange(val *ConfigChange) *NullableConfigChange {
	return &NullableConfigChange{value: val, isSet: true}
}

func (v NullableConfigChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


