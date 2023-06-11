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

// GatewayMigrationGetOutput struct for GatewayMigrationGetOutput
type GatewayMigrationGetOutput struct {
	Body *MigrationsConfigPart `json:"body,omitempty"`
}

// NewGatewayMigrationGetOutput instantiates a new GatewayMigrationGetOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMigrationGetOutput() *GatewayMigrationGetOutput {
	this := GatewayMigrationGetOutput{}
	return &this
}

// NewGatewayMigrationGetOutputWithDefaults instantiates a new GatewayMigrationGetOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMigrationGetOutputWithDefaults() *GatewayMigrationGetOutput {
	this := GatewayMigrationGetOutput{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *GatewayMigrationGetOutput) GetBody() MigrationsConfigPart {
	if o == nil || o.Body == nil {
		var ret MigrationsConfigPart
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMigrationGetOutput) GetBodyOk() (*MigrationsConfigPart, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *GatewayMigrationGetOutput) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given MigrationsConfigPart and assigns it to the Body field.
func (o *GatewayMigrationGetOutput) SetBody(v MigrationsConfigPart) {
	o.Body = &v
}

func (o GatewayMigrationGetOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayMigrationGetOutput struct {
	value *GatewayMigrationGetOutput
	isSet bool
}

func (v NullableGatewayMigrationGetOutput) Get() *GatewayMigrationGetOutput {
	return v.value
}

func (v *NullableGatewayMigrationGetOutput) Set(val *GatewayMigrationGetOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMigrationGetOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMigrationGetOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMigrationGetOutput(val *GatewayMigrationGetOutput) *NullableGatewayMigrationGetOutput {
	return &NullableGatewayMigrationGetOutput{value: val, isSet: true}
}

func (v NullableGatewayMigrationGetOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMigrationGetOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


