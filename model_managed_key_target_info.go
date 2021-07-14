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

// ManagedKeyTargetInfo struct for ManagedKeyTargetInfo
type ManagedKeyTargetInfo struct {
	ExternalKmsId *ExternalKMSKeyId `json:"external_kms_id,omitempty"`
	KeyPurpose *[]string `json:"key_purpose,omitempty"`
	KeyStatuses *[]ManagedKeyStatusInfo `json:"key_statuses,omitempty"`
	TargetAssocId *string `json:"target_assoc_id,omitempty"`
}

// NewManagedKeyTargetInfo instantiates a new ManagedKeyTargetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedKeyTargetInfo() *ManagedKeyTargetInfo {
	this := ManagedKeyTargetInfo{}
	return &this
}

// NewManagedKeyTargetInfoWithDefaults instantiates a new ManagedKeyTargetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedKeyTargetInfoWithDefaults() *ManagedKeyTargetInfo {
	this := ManagedKeyTargetInfo{}
	return &this
}

// GetExternalKmsId returns the ExternalKmsId field value if set, zero value otherwise.
func (o *ManagedKeyTargetInfo) GetExternalKmsId() ExternalKMSKeyId {
	if o == nil || o.ExternalKmsId == nil {
		var ret ExternalKMSKeyId
		return ret
	}
	return *o.ExternalKmsId
}

// GetExternalKmsIdOk returns a tuple with the ExternalKmsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedKeyTargetInfo) GetExternalKmsIdOk() (*ExternalKMSKeyId, bool) {
	if o == nil || o.ExternalKmsId == nil {
		return nil, false
	}
	return o.ExternalKmsId, true
}

// HasExternalKmsId returns a boolean if a field has been set.
func (o *ManagedKeyTargetInfo) HasExternalKmsId() bool {
	if o != nil && o.ExternalKmsId != nil {
		return true
	}

	return false
}

// SetExternalKmsId gets a reference to the given ExternalKMSKeyId and assigns it to the ExternalKmsId field.
func (o *ManagedKeyTargetInfo) SetExternalKmsId(v ExternalKMSKeyId) {
	o.ExternalKmsId = &v
}

// GetKeyPurpose returns the KeyPurpose field value if set, zero value otherwise.
func (o *ManagedKeyTargetInfo) GetKeyPurpose() []string {
	if o == nil || o.KeyPurpose == nil {
		var ret []string
		return ret
	}
	return *o.KeyPurpose
}

// GetKeyPurposeOk returns a tuple with the KeyPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedKeyTargetInfo) GetKeyPurposeOk() (*[]string, bool) {
	if o == nil || o.KeyPurpose == nil {
		return nil, false
	}
	return o.KeyPurpose, true
}

// HasKeyPurpose returns a boolean if a field has been set.
func (o *ManagedKeyTargetInfo) HasKeyPurpose() bool {
	if o != nil && o.KeyPurpose != nil {
		return true
	}

	return false
}

// SetKeyPurpose gets a reference to the given []string and assigns it to the KeyPurpose field.
func (o *ManagedKeyTargetInfo) SetKeyPurpose(v []string) {
	o.KeyPurpose = &v
}

// GetKeyStatuses returns the KeyStatuses field value if set, zero value otherwise.
func (o *ManagedKeyTargetInfo) GetKeyStatuses() []ManagedKeyStatusInfo {
	if o == nil || o.KeyStatuses == nil {
		var ret []ManagedKeyStatusInfo
		return ret
	}
	return *o.KeyStatuses
}

// GetKeyStatusesOk returns a tuple with the KeyStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedKeyTargetInfo) GetKeyStatusesOk() (*[]ManagedKeyStatusInfo, bool) {
	if o == nil || o.KeyStatuses == nil {
		return nil, false
	}
	return o.KeyStatuses, true
}

// HasKeyStatuses returns a boolean if a field has been set.
func (o *ManagedKeyTargetInfo) HasKeyStatuses() bool {
	if o != nil && o.KeyStatuses != nil {
		return true
	}

	return false
}

// SetKeyStatuses gets a reference to the given []ManagedKeyStatusInfo and assigns it to the KeyStatuses field.
func (o *ManagedKeyTargetInfo) SetKeyStatuses(v []ManagedKeyStatusInfo) {
	o.KeyStatuses = &v
}

// GetTargetAssocId returns the TargetAssocId field value if set, zero value otherwise.
func (o *ManagedKeyTargetInfo) GetTargetAssocId() string {
	if o == nil || o.TargetAssocId == nil {
		var ret string
		return ret
	}
	return *o.TargetAssocId
}

// GetTargetAssocIdOk returns a tuple with the TargetAssocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedKeyTargetInfo) GetTargetAssocIdOk() (*string, bool) {
	if o == nil || o.TargetAssocId == nil {
		return nil, false
	}
	return o.TargetAssocId, true
}

// HasTargetAssocId returns a boolean if a field has been set.
func (o *ManagedKeyTargetInfo) HasTargetAssocId() bool {
	if o != nil && o.TargetAssocId != nil {
		return true
	}

	return false
}

// SetTargetAssocId gets a reference to the given string and assigns it to the TargetAssocId field.
func (o *ManagedKeyTargetInfo) SetTargetAssocId(v string) {
	o.TargetAssocId = &v
}

func (o ManagedKeyTargetInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalKmsId != nil {
		toSerialize["external_kms_id"] = o.ExternalKmsId
	}
	if o.KeyPurpose != nil {
		toSerialize["key_purpose"] = o.KeyPurpose
	}
	if o.KeyStatuses != nil {
		toSerialize["key_statuses"] = o.KeyStatuses
	}
	if o.TargetAssocId != nil {
		toSerialize["target_assoc_id"] = o.TargetAssocId
	}
	return json.Marshal(toSerialize)
}

type NullableManagedKeyTargetInfo struct {
	value *ManagedKeyTargetInfo
	isSet bool
}

func (v NullableManagedKeyTargetInfo) Get() *ManagedKeyTargetInfo {
	return v.value
}

func (v *NullableManagedKeyTargetInfo) Set(val *ManagedKeyTargetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedKeyTargetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedKeyTargetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedKeyTargetInfo(val *ManagedKeyTargetInfo) *NullableManagedKeyTargetInfo {
	return &NullableManagedKeyTargetInfo{value: val, isSet: true}
}

func (v NullableManagedKeyTargetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedKeyTargetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

