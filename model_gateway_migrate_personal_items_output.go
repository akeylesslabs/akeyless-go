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

// checks if the GatewayMigratePersonalItemsOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayMigratePersonalItemsOutput{}

// GatewayMigratePersonalItemsOutput struct for GatewayMigratePersonalItemsOutput
type GatewayMigratePersonalItemsOutput struct {
	MigrationItems *MigrationItems `json:"migration_items,omitempty"`
}

// NewGatewayMigratePersonalItemsOutput instantiates a new GatewayMigratePersonalItemsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMigratePersonalItemsOutput() *GatewayMigratePersonalItemsOutput {
	this := GatewayMigratePersonalItemsOutput{}
	return &this
}

// NewGatewayMigratePersonalItemsOutputWithDefaults instantiates a new GatewayMigratePersonalItemsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMigratePersonalItemsOutputWithDefaults() *GatewayMigratePersonalItemsOutput {
	this := GatewayMigratePersonalItemsOutput{}
	return &this
}

// GetMigrationItems returns the MigrationItems field value if set, zero value otherwise.
func (o *GatewayMigratePersonalItemsOutput) GetMigrationItems() MigrationItems {
	if o == nil || IsNil(o.MigrationItems) {
		var ret MigrationItems
		return ret
	}
	return *o.MigrationItems
}

// GetMigrationItemsOk returns a tuple with the MigrationItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigratePersonalItemsOutput) GetMigrationItemsOk() (*MigrationItems, bool) {
	if o == nil || IsNil(o.MigrationItems) {
		return nil, false
	}
	return o.MigrationItems, true
}

// HasMigrationItems returns a boolean if a field has been set.
func (o *GatewayMigratePersonalItemsOutput) HasMigrationItems() bool {
	if o != nil && !IsNil(o.MigrationItems) {
		return true
	}

	return false
}

// SetMigrationItems gets a reference to the given MigrationItems and assigns it to the MigrationItems field.
func (o *GatewayMigratePersonalItemsOutput) SetMigrationItems(v MigrationItems) {
	o.MigrationItems = &v
}

func (o GatewayMigratePersonalItemsOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayMigratePersonalItemsOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MigrationItems) {
		toSerialize["migration_items"] = o.MigrationItems
	}
	return toSerialize, nil
}

type NullableGatewayMigratePersonalItemsOutput struct {
	value *GatewayMigratePersonalItemsOutput
	isSet bool
}

func (v NullableGatewayMigratePersonalItemsOutput) Get() *GatewayMigratePersonalItemsOutput {
	return v.value
}

func (v *NullableGatewayMigratePersonalItemsOutput) Set(val *GatewayMigratePersonalItemsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMigratePersonalItemsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMigratePersonalItemsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMigratePersonalItemsOutput(val *GatewayMigratePersonalItemsOutput) *NullableGatewayMigratePersonalItemsOutput {
	return &NullableGatewayMigratePersonalItemsOutput{value: val, isSet: true}
}

func (v NullableGatewayMigratePersonalItemsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMigratePersonalItemsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


