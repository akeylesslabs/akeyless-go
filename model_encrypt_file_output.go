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

// checks if the EncryptFileOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptFileOutput{}

// EncryptFileOutput struct for EncryptFileOutput
type EncryptFileOutput struct {
	Result *string `json:"result,omitempty"`
}

// NewEncryptFileOutput instantiates a new EncryptFileOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptFileOutput() *EncryptFileOutput {
	this := EncryptFileOutput{}
	return &this
}

// NewEncryptFileOutputWithDefaults instantiates a new EncryptFileOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptFileOutputWithDefaults() *EncryptFileOutput {
	this := EncryptFileOutput{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *EncryptFileOutput) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptFileOutput) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *EncryptFileOutput) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *EncryptFileOutput) SetResult(v string) {
	o.Result = &v
}

func (o EncryptFileOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptFileOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableEncryptFileOutput struct {
	value *EncryptFileOutput
	isSet bool
}

func (v NullableEncryptFileOutput) Get() *EncryptFileOutput {
	return v.value
}

func (v *NullableEncryptFileOutput) Set(val *EncryptFileOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptFileOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptFileOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptFileOutput(val *EncryptFileOutput) *NullableEncryptFileOutput {
	return &NullableEncryptFileOutput{value: val, isSet: true}
}

func (v NullableEncryptFileOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptFileOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


