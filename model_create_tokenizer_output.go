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

// checks if the CreateTokenizerOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTokenizerOutput{}

// CreateTokenizerOutput struct for CreateTokenizerOutput
type CreateTokenizerOutput struct {
	Name *string `json:"name,omitempty"`
}

// NewCreateTokenizerOutput instantiates a new CreateTokenizerOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokenizerOutput() *CreateTokenizerOutput {
	this := CreateTokenizerOutput{}
	return &this
}

// NewCreateTokenizerOutputWithDefaults instantiates a new CreateTokenizerOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokenizerOutputWithDefaults() *CreateTokenizerOutput {
	this := CreateTokenizerOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateTokenizerOutput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenizerOutput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateTokenizerOutput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateTokenizerOutput) SetName(v string) {
	o.Name = &v
}

func (o CreateTokenizerOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTokenizerOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateTokenizerOutput struct {
	value *CreateTokenizerOutput
	isSet bool
}

func (v NullableCreateTokenizerOutput) Get() *CreateTokenizerOutput {
	return v.value
}

func (v *NullableCreateTokenizerOutput) Set(val *CreateTokenizerOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokenizerOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokenizerOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokenizerOutput(val *CreateTokenizerOutput) *NullableCreateTokenizerOutput {
	return &NullableCreateTokenizerOutput{value: val, isSet: true}
}

func (v NullableCreateTokenizerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokenizerOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


