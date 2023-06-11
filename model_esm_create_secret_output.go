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

// checks if the EsmCreateSecretOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsmCreateSecretOutput{}

// EsmCreateSecretOutput struct for EsmCreateSecretOutput
type EsmCreateSecretOutput struct {
	SecretId *string `json:"secret_id,omitempty"`
	VersionId *string `json:"version_id,omitempty"`
}

// NewEsmCreateSecretOutput instantiates a new EsmCreateSecretOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsmCreateSecretOutput() *EsmCreateSecretOutput {
	this := EsmCreateSecretOutput{}
	return &this
}

// NewEsmCreateSecretOutputWithDefaults instantiates a new EsmCreateSecretOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsmCreateSecretOutputWithDefaults() *EsmCreateSecretOutput {
	this := EsmCreateSecretOutput{}
	return &this
}

// GetSecretId returns the SecretId field value if set, zero value otherwise.
func (o *EsmCreateSecretOutput) GetSecretId() string {
	if o == nil || IsNil(o.SecretId) {
		var ret string
		return ret
	}
	return *o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsmCreateSecretOutput) GetSecretIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecretId) {
		return nil, false
	}
	return o.SecretId, true
}

// HasSecretId returns a boolean if a field has been set.
func (o *EsmCreateSecretOutput) HasSecretId() bool {
	if o != nil && !IsNil(o.SecretId) {
		return true
	}

	return false
}

// SetSecretId gets a reference to the given string and assigns it to the SecretId field.
func (o *EsmCreateSecretOutput) SetSecretId(v string) {
	o.SecretId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *EsmCreateSecretOutput) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsmCreateSecretOutput) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *EsmCreateSecretOutput) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *EsmCreateSecretOutput) SetVersionId(v string) {
	o.VersionId = &v
}

func (o EsmCreateSecretOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsmCreateSecretOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecretId) {
		toSerialize["secret_id"] = o.SecretId
	}
	if !IsNil(o.VersionId) {
		toSerialize["version_id"] = o.VersionId
	}
	return toSerialize, nil
}

type NullableEsmCreateSecretOutput struct {
	value *EsmCreateSecretOutput
	isSet bool
}

func (v NullableEsmCreateSecretOutput) Get() *EsmCreateSecretOutput {
	return v.value
}

func (v *NullableEsmCreateSecretOutput) Set(val *EsmCreateSecretOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableEsmCreateSecretOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableEsmCreateSecretOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsmCreateSecretOutput(val *EsmCreateSecretOutput) *NullableEsmCreateSecretOutput {
	return &NullableEsmCreateSecretOutput{value: val, isSet: true}
}

func (v NullableEsmCreateSecretOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsmCreateSecretOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


