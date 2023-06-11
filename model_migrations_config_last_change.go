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

// MigrationsConfigLastChange struct for MigrationsConfigLastChange
type MigrationsConfigLastChange struct {
	ChangedMigrations []string `json:"changed_migrations,omitempty"`
	CreatedMigrations []string `json:"created_migrations,omitempty"`
	DeletedMigrations []string `json:"deleted_migrations,omitempty"`
}

// NewMigrationsConfigLastChange instantiates a new MigrationsConfigLastChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationsConfigLastChange() *MigrationsConfigLastChange {
	this := MigrationsConfigLastChange{}
	return &this
}

// NewMigrationsConfigLastChangeWithDefaults instantiates a new MigrationsConfigLastChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationsConfigLastChangeWithDefaults() *MigrationsConfigLastChange {
	this := MigrationsConfigLastChange{}
	return &this
}

// GetChangedMigrations returns the ChangedMigrations field value if set, zero value otherwise.
func (o *MigrationsConfigLastChange) GetChangedMigrations() []string {
	if o == nil || o.ChangedMigrations == nil {
		var ret []string
		return ret
	}
	return o.ChangedMigrations
}

// GetChangedMigrationsOk returns a tuple with the ChangedMigrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationsConfigLastChange) GetChangedMigrationsOk() ([]string, bool) {
	if o == nil || o.ChangedMigrations == nil {
		return nil, false
	}
	return o.ChangedMigrations, true
}

// HasChangedMigrations returns a boolean if a field has been set.
func (o *MigrationsConfigLastChange) HasChangedMigrations() bool {
	if o != nil && o.ChangedMigrations != nil {
		return true
	}

	return false
}

// SetChangedMigrations gets a reference to the given []string and assigns it to the ChangedMigrations field.
func (o *MigrationsConfigLastChange) SetChangedMigrations(v []string) {
	o.ChangedMigrations = v
}

// GetCreatedMigrations returns the CreatedMigrations field value if set, zero value otherwise.
func (o *MigrationsConfigLastChange) GetCreatedMigrations() []string {
	if o == nil || o.CreatedMigrations == nil {
		var ret []string
		return ret
	}
	return o.CreatedMigrations
}

// GetCreatedMigrationsOk returns a tuple with the CreatedMigrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationsConfigLastChange) GetCreatedMigrationsOk() ([]string, bool) {
	if o == nil || o.CreatedMigrations == nil {
		return nil, false
	}
	return o.CreatedMigrations, true
}

// HasCreatedMigrations returns a boolean if a field has been set.
func (o *MigrationsConfigLastChange) HasCreatedMigrations() bool {
	if o != nil && o.CreatedMigrations != nil {
		return true
	}

	return false
}

// SetCreatedMigrations gets a reference to the given []string and assigns it to the CreatedMigrations field.
func (o *MigrationsConfigLastChange) SetCreatedMigrations(v []string) {
	o.CreatedMigrations = v
}

// GetDeletedMigrations returns the DeletedMigrations field value if set, zero value otherwise.
func (o *MigrationsConfigLastChange) GetDeletedMigrations() []string {
	if o == nil || o.DeletedMigrations == nil {
		var ret []string
		return ret
	}
	return o.DeletedMigrations
}

// GetDeletedMigrationsOk returns a tuple with the DeletedMigrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationsConfigLastChange) GetDeletedMigrationsOk() ([]string, bool) {
	if o == nil || o.DeletedMigrations == nil {
		return nil, false
	}
	return o.DeletedMigrations, true
}

// HasDeletedMigrations returns a boolean if a field has been set.
func (o *MigrationsConfigLastChange) HasDeletedMigrations() bool {
	if o != nil && o.DeletedMigrations != nil {
		return true
	}

	return false
}

// SetDeletedMigrations gets a reference to the given []string and assigns it to the DeletedMigrations field.
func (o *MigrationsConfigLastChange) SetDeletedMigrations(v []string) {
	o.DeletedMigrations = v
}

func (o MigrationsConfigLastChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangedMigrations != nil {
		toSerialize["changed_migrations"] = o.ChangedMigrations
	}
	if o.CreatedMigrations != nil {
		toSerialize["created_migrations"] = o.CreatedMigrations
	}
	if o.DeletedMigrations != nil {
		toSerialize["deleted_migrations"] = o.DeletedMigrations
	}
	return json.Marshal(toSerialize)
}

type NullableMigrationsConfigLastChange struct {
	value *MigrationsConfigLastChange
	isSet bool
}

func (v NullableMigrationsConfigLastChange) Get() *MigrationsConfigLastChange {
	return v.value
}

func (v *NullableMigrationsConfigLastChange) Set(val *MigrationsConfigLastChange) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationsConfigLastChange) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationsConfigLastChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationsConfigLastChange(val *MigrationsConfigLastChange) *NullableMigrationsConfigLastChange {
	return &NullableMigrationsConfigLastChange{value: val, isSet: true}
}

func (v NullableMigrationsConfigLastChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationsConfigLastChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


