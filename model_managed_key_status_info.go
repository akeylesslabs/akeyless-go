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

// ManagedKeyStatusInfo struct for ManagedKeyStatusInfo
type ManagedKeyStatusInfo struct {
	KeyId *int64 `json:"key_id,omitempty"`
	LastError *string `json:"last_error,omitempty"`
	// ManagedKeyTargetStatus defines status of managed key target
	LastStatus *string `json:"last_status,omitempty"`
}

// NewManagedKeyStatusInfo instantiates a new ManagedKeyStatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedKeyStatusInfo() *ManagedKeyStatusInfo {
	this := ManagedKeyStatusInfo{}
	return &this
}

// NewManagedKeyStatusInfoWithDefaults instantiates a new ManagedKeyStatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedKeyStatusInfoWithDefaults() *ManagedKeyStatusInfo {
	this := ManagedKeyStatusInfo{}
	return &this
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *ManagedKeyStatusInfo) GetKeyId() int64 {
	if o == nil || o.KeyId == nil {
		var ret int64
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedKeyStatusInfo) GetKeyIdOk() (*int64, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *ManagedKeyStatusInfo) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given int64 and assigns it to the KeyId field.
func (o *ManagedKeyStatusInfo) SetKeyId(v int64) {
	o.KeyId = &v
}

// GetLastError returns the LastError field value if set, zero value otherwise.
func (o *ManagedKeyStatusInfo) GetLastError() string {
	if o == nil || o.LastError == nil {
		var ret string
		return ret
	}
	return *o.LastError
}

// GetLastErrorOk returns a tuple with the LastError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedKeyStatusInfo) GetLastErrorOk() (*string, bool) {
	if o == nil || o.LastError == nil {
		return nil, false
	}
	return o.LastError, true
}

// HasLastError returns a boolean if a field has been set.
func (o *ManagedKeyStatusInfo) HasLastError() bool {
	if o != nil && o.LastError != nil {
		return true
	}

	return false
}

// SetLastError gets a reference to the given string and assigns it to the LastError field.
func (o *ManagedKeyStatusInfo) SetLastError(v string) {
	o.LastError = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *ManagedKeyStatusInfo) GetLastStatus() string {
	if o == nil || o.LastStatus == nil {
		var ret string
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedKeyStatusInfo) GetLastStatusOk() (*string, bool) {
	if o == nil || o.LastStatus == nil {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *ManagedKeyStatusInfo) HasLastStatus() bool {
	if o != nil && o.LastStatus != nil {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given string and assigns it to the LastStatus field.
func (o *ManagedKeyStatusInfo) SetLastStatus(v string) {
	o.LastStatus = &v
}

func (o ManagedKeyStatusInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyId != nil {
		toSerialize["key_id"] = o.KeyId
	}
	if o.LastError != nil {
		toSerialize["last_error"] = o.LastError
	}
	if o.LastStatus != nil {
		toSerialize["last_status"] = o.LastStatus
	}
	return json.Marshal(toSerialize)
}

type NullableManagedKeyStatusInfo struct {
	value *ManagedKeyStatusInfo
	isSet bool
}

func (v NullableManagedKeyStatusInfo) Get() *ManagedKeyStatusInfo {
	return v.value
}

func (v *NullableManagedKeyStatusInfo) Set(val *ManagedKeyStatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedKeyStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedKeyStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedKeyStatusInfo(val *ManagedKeyStatusInfo) *NullableManagedKeyStatusInfo {
	return &NullableManagedKeyStatusInfo{value: val, isSet: true}
}

func (v NullableManagedKeyStatusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedKeyStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


