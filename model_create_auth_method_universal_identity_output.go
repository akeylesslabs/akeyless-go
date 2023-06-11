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

// checks if the CreateAuthMethodUniversalIdentityOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAuthMethodUniversalIdentityOutput{}

// CreateAuthMethodUniversalIdentityOutput struct for CreateAuthMethodUniversalIdentityOutput
type CreateAuthMethodUniversalIdentityOutput struct {
	AccessId *string `json:"access_id,omitempty"`
}

// NewCreateAuthMethodUniversalIdentityOutput instantiates a new CreateAuthMethodUniversalIdentityOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodUniversalIdentityOutput() *CreateAuthMethodUniversalIdentityOutput {
	this := CreateAuthMethodUniversalIdentityOutput{}
	return &this
}

// NewCreateAuthMethodUniversalIdentityOutputWithDefaults instantiates a new CreateAuthMethodUniversalIdentityOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodUniversalIdentityOutputWithDefaults() *CreateAuthMethodUniversalIdentityOutput {
	this := CreateAuthMethodUniversalIdentityOutput{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentityOutput) GetAccessId() string {
	if o == nil || IsNil(o.AccessId) {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentityOutput) GetAccessIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessId) {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentityOutput) HasAccessId() bool {
	if o != nil && !IsNil(o.AccessId) {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *CreateAuthMethodUniversalIdentityOutput) SetAccessId(v string) {
	o.AccessId = &v
}

func (o CreateAuthMethodUniversalIdentityOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAuthMethodUniversalIdentityOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessId) {
		toSerialize["access_id"] = o.AccessId
	}
	return toSerialize, nil
}

type NullableCreateAuthMethodUniversalIdentityOutput struct {
	value *CreateAuthMethodUniversalIdentityOutput
	isSet bool
}

func (v NullableCreateAuthMethodUniversalIdentityOutput) Get() *CreateAuthMethodUniversalIdentityOutput {
	return v.value
}

func (v *NullableCreateAuthMethodUniversalIdentityOutput) Set(val *CreateAuthMethodUniversalIdentityOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodUniversalIdentityOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodUniversalIdentityOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodUniversalIdentityOutput(val *CreateAuthMethodUniversalIdentityOutput) *NullableCreateAuthMethodUniversalIdentityOutput {
	return &NullableCreateAuthMethodUniversalIdentityOutput{value: val, isSet: true}
}

func (v NullableCreateAuthMethodUniversalIdentityOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodUniversalIdentityOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


