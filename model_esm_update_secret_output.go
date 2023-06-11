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

// checks if the EsmUpdateSecretOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsmUpdateSecretOutput{}

// EsmUpdateSecretOutput struct for EsmUpdateSecretOutput
type EsmUpdateSecretOutput struct {
	SecretId *string `json:"secret_id,omitempty"`
	VersionId *string `json:"version_id,omitempty"`
}

// NewEsmUpdateSecretOutput instantiates a new EsmUpdateSecretOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsmUpdateSecretOutput() *EsmUpdateSecretOutput {
	this := EsmUpdateSecretOutput{}
	return &this
}

// NewEsmUpdateSecretOutputWithDefaults instantiates a new EsmUpdateSecretOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsmUpdateSecretOutputWithDefaults() *EsmUpdateSecretOutput {
	this := EsmUpdateSecretOutput{}
	return &this
}

// GetSecretId returns the SecretId field value if set, zero value otherwise.
func (o *EsmUpdateSecretOutput) GetSecretId() string {
	if o == nil || IsNil(o.SecretId) {
		var ret string
		return ret
	}
	return *o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsmUpdateSecretOutput) GetSecretIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecretId) {
		return nil, false
	}
	return o.SecretId, true
}

// HasSecretId returns a boolean if a field has been set.
func (o *EsmUpdateSecretOutput) HasSecretId() bool {
	if o != nil && !IsNil(o.SecretId) {
		return true
	}

	return false
}

// SetSecretId gets a reference to the given string and assigns it to the SecretId field.
func (o *EsmUpdateSecretOutput) SetSecretId(v string) {
	o.SecretId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *EsmUpdateSecretOutput) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsmUpdateSecretOutput) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *EsmUpdateSecretOutput) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *EsmUpdateSecretOutput) SetVersionId(v string) {
	o.VersionId = &v
}

func (o EsmUpdateSecretOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsmUpdateSecretOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecretId) {
		toSerialize["secret_id"] = o.SecretId
	}
	if !IsNil(o.VersionId) {
		toSerialize["version_id"] = o.VersionId
	}
	return toSerialize, nil
}

type NullableEsmUpdateSecretOutput struct {
	value *EsmUpdateSecretOutput
	isSet bool
}

func (v NullableEsmUpdateSecretOutput) Get() *EsmUpdateSecretOutput {
	return v.value
}

func (v *NullableEsmUpdateSecretOutput) Set(val *EsmUpdateSecretOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableEsmUpdateSecretOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableEsmUpdateSecretOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsmUpdateSecretOutput(val *EsmUpdateSecretOutput) *NullableEsmUpdateSecretOutput {
	return &NullableEsmUpdateSecretOutput{value: val, isSet: true}
}

func (v NullableEsmUpdateSecretOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsmUpdateSecretOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


