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

// checks if the ExportClassicKeyOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportClassicKeyOutput{}

// ExportClassicKeyOutput struct for ExportClassicKeyOutput
type ExportClassicKeyOutput struct {
	CertificatePem *string `json:"certificatePem,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewExportClassicKeyOutput instantiates a new ExportClassicKeyOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportClassicKeyOutput() *ExportClassicKeyOutput {
	this := ExportClassicKeyOutput{}
	return &this
}

// NewExportClassicKeyOutputWithDefaults instantiates a new ExportClassicKeyOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportClassicKeyOutputWithDefaults() *ExportClassicKeyOutput {
	this := ExportClassicKeyOutput{}
	return &this
}

// GetCertificatePem returns the CertificatePem field value if set, zero value otherwise.
func (o *ExportClassicKeyOutput) GetCertificatePem() string {
	if o == nil || IsNil(o.CertificatePem) {
		var ret string
		return ret
	}
	return *o.CertificatePem
}

// GetCertificatePemOk returns a tuple with the CertificatePem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportClassicKeyOutput) GetCertificatePemOk() (*string, bool) {
	if o == nil || IsNil(o.CertificatePem) {
		return nil, false
	}
	return o.CertificatePem, true
}

// HasCertificatePem returns a boolean if a field has been set.
func (o *ExportClassicKeyOutput) HasCertificatePem() bool {
	if o != nil && !IsNil(o.CertificatePem) {
		return true
	}

	return false
}

// SetCertificatePem gets a reference to the given string and assigns it to the CertificatePem field.
func (o *ExportClassicKeyOutput) SetCertificatePem(v string) {
	o.CertificatePem = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ExportClassicKeyOutput) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportClassicKeyOutput) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ExportClassicKeyOutput) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ExportClassicKeyOutput) SetKey(v string) {
	o.Key = &v
}

func (o ExportClassicKeyOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportClassicKeyOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificatePem) {
		toSerialize["certificatePem"] = o.CertificatePem
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableExportClassicKeyOutput struct {
	value *ExportClassicKeyOutput
	isSet bool
}

func (v NullableExportClassicKeyOutput) Get() *ExportClassicKeyOutput {
	return v.value
}

func (v *NullableExportClassicKeyOutput) Set(val *ExportClassicKeyOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableExportClassicKeyOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableExportClassicKeyOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportClassicKeyOutput(val *ExportClassicKeyOutput) *NullableExportClassicKeyOutput {
	return &NullableExportClassicKeyOutput{value: val, isSet: true}
}

func (v NullableExportClassicKeyOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportClassicKeyOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


