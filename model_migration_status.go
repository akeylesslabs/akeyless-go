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

// MigrationStatus struct for MigrationStatus
type MigrationStatus struct {
	LastMessages *map[string]string `json:"last_messages,omitempty"`
	LastReports *map[string]string `json:"last_reports,omitempty"`
	LastStatuses *map[string]string `json:"last_statuses,omitempty"`
}

// NewMigrationStatus instantiates a new MigrationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationStatus() *MigrationStatus {
	this := MigrationStatus{}
	return &this
}

// NewMigrationStatusWithDefaults instantiates a new MigrationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationStatusWithDefaults() *MigrationStatus {
	this := MigrationStatus{}
	return &this
}

// GetLastMessages returns the LastMessages field value if set, zero value otherwise.
func (o *MigrationStatus) GetLastMessages() map[string]string {
	if o == nil || o.LastMessages == nil {
		var ret map[string]string
		return ret
	}
	return *o.LastMessages
}

// GetLastMessagesOk returns a tuple with the LastMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatus) GetLastMessagesOk() (*map[string]string, bool) {
	if o == nil || o.LastMessages == nil {
		return nil, false
	}
	return o.LastMessages, true
}

// HasLastMessages returns a boolean if a field has been set.
func (o *MigrationStatus) HasLastMessages() bool {
	if o != nil && o.LastMessages != nil {
		return true
	}

	return false
}

// SetLastMessages gets a reference to the given map[string]string and assigns it to the LastMessages field.
func (o *MigrationStatus) SetLastMessages(v map[string]string) {
	o.LastMessages = &v
}

// GetLastReports returns the LastReports field value if set, zero value otherwise.
func (o *MigrationStatus) GetLastReports() map[string]string {
	if o == nil || o.LastReports == nil {
		var ret map[string]string
		return ret
	}
	return *o.LastReports
}

// GetLastReportsOk returns a tuple with the LastReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatus) GetLastReportsOk() (*map[string]string, bool) {
	if o == nil || o.LastReports == nil {
		return nil, false
	}
	return o.LastReports, true
}

// HasLastReports returns a boolean if a field has been set.
func (o *MigrationStatus) HasLastReports() bool {
	if o != nil && o.LastReports != nil {
		return true
	}

	return false
}

// SetLastReports gets a reference to the given map[string]string and assigns it to the LastReports field.
func (o *MigrationStatus) SetLastReports(v map[string]string) {
	o.LastReports = &v
}

// GetLastStatuses returns the LastStatuses field value if set, zero value otherwise.
func (o *MigrationStatus) GetLastStatuses() map[string]string {
	if o == nil || o.LastStatuses == nil {
		var ret map[string]string
		return ret
	}
	return *o.LastStatuses
}

// GetLastStatusesOk returns a tuple with the LastStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationStatus) GetLastStatusesOk() (*map[string]string, bool) {
	if o == nil || o.LastStatuses == nil {
		return nil, false
	}
	return o.LastStatuses, true
}

// HasLastStatuses returns a boolean if a field has been set.
func (o *MigrationStatus) HasLastStatuses() bool {
	if o != nil && o.LastStatuses != nil {
		return true
	}

	return false
}

// SetLastStatuses gets a reference to the given map[string]string and assigns it to the LastStatuses field.
func (o *MigrationStatus) SetLastStatuses(v map[string]string) {
	o.LastStatuses = &v
}

func (o MigrationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastMessages != nil {
		toSerialize["last_messages"] = o.LastMessages
	}
	if o.LastReports != nil {
		toSerialize["last_reports"] = o.LastReports
	}
	if o.LastStatuses != nil {
		toSerialize["last_statuses"] = o.LastStatuses
	}
	return json.Marshal(toSerialize)
}

type NullableMigrationStatus struct {
	value *MigrationStatus
	isSet bool
}

func (v NullableMigrationStatus) Get() *MigrationStatus {
	return v.value
}

func (v *NullableMigrationStatus) Set(val *MigrationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationStatus(val *MigrationStatus) *NullableMigrationStatus {
	return &NullableMigrationStatus{value: val, isSet: true}
}

func (v NullableMigrationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


