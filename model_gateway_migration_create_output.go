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

// GatewayMigrationCreateOutput struct for GatewayMigrationCreateOutput
type GatewayMigrationCreateOutput struct {
	MigrationName *string `json:"migration_name,omitempty"`
}

// NewGatewayMigrationCreateOutput instantiates a new GatewayMigrationCreateOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMigrationCreateOutput() *GatewayMigrationCreateOutput {
	this := GatewayMigrationCreateOutput{}
	return &this
}

// NewGatewayMigrationCreateOutputWithDefaults instantiates a new GatewayMigrationCreateOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMigrationCreateOutputWithDefaults() *GatewayMigrationCreateOutput {
	this := GatewayMigrationCreateOutput{}
	return &this
}

// GetMigrationName returns the MigrationName field value if set, zero value otherwise.
func (o *GatewayMigrationCreateOutput) GetMigrationName() string {
	if o == nil || o.MigrationName == nil {
		var ret string
		return ret
	}
	return *o.MigrationName
}

// GetMigrationNameOk returns a tuple with the MigrationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigrationCreateOutput) GetMigrationNameOk() (*string, bool) {
	if o == nil || o.MigrationName == nil {
		return nil, false
	}
	return o.MigrationName, true
}

// HasMigrationName returns a boolean if a field has been set.
func (o *GatewayMigrationCreateOutput) HasMigrationName() bool {
	if o != nil && o.MigrationName != nil {
		return true
	}

	return false
}

// SetMigrationName gets a reference to the given string and assigns it to the MigrationName field.
func (o *GatewayMigrationCreateOutput) SetMigrationName(v string) {
	o.MigrationName = &v
}

func (o GatewayMigrationCreateOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MigrationName != nil {
		toSerialize["migration_name"] = o.MigrationName
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayMigrationCreateOutput struct {
	value *GatewayMigrationCreateOutput
	isSet bool
}

func (v NullableGatewayMigrationCreateOutput) Get() *GatewayMigrationCreateOutput {
	return v.value
}

func (v *NullableGatewayMigrationCreateOutput) Set(val *GatewayMigrationCreateOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMigrationCreateOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMigrationCreateOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMigrationCreateOutput(val *GatewayMigrationCreateOutput) *NullableGatewayMigrationCreateOutput {
	return &NullableGatewayMigrationCreateOutput{value: val, isSet: true}
}

func (v NullableGatewayMigrationCreateOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMigrationCreateOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


