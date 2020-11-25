/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GetSSHCertificateOutput struct for GetSSHCertificateOutput
type GetSSHCertificateOutput struct {
	Data *string `json:"data,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewGetSSHCertificateOutput instantiates a new GetSSHCertificateOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSSHCertificateOutput() *GetSSHCertificateOutput {
	this := GetSSHCertificateOutput{}
	return &this
}

// NewGetSSHCertificateOutputWithDefaults instantiates a new GetSSHCertificateOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSSHCertificateOutputWithDefaults() *GetSSHCertificateOutput {
	this := GetSSHCertificateOutput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSSHCertificateOutput) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSSHCertificateOutput) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSSHCertificateOutput) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *GetSSHCertificateOutput) SetData(v string) {
	o.Data = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GetSSHCertificateOutput) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSSHCertificateOutput) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GetSSHCertificateOutput) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GetSSHCertificateOutput) SetPath(v string) {
	o.Path = &v
}

func (o GetSSHCertificateOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableGetSSHCertificateOutput struct {
	value *GetSSHCertificateOutput
	isSet bool
}

func (v NullableGetSSHCertificateOutput) Get() *GetSSHCertificateOutput {
	return v.value
}

func (v *NullableGetSSHCertificateOutput) Set(val *GetSSHCertificateOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSSHCertificateOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSSHCertificateOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSSHCertificateOutput(val *GetSSHCertificateOutput) *NullableGetSSHCertificateOutput {
	return &NullableGetSSHCertificateOutput{value: val, isSet: true}
}

func (v NullableGetSSHCertificateOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSSHCertificateOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


