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

// CreateAuthMethodOIDCOutput struct for CreateAuthMethodOIDCOutput
type CreateAuthMethodOIDCOutput struct {
	AccessId *string `json:"access_id,omitempty"`
}

// NewCreateAuthMethodOIDCOutput instantiates a new CreateAuthMethodOIDCOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodOIDCOutput() *CreateAuthMethodOIDCOutput {
	this := CreateAuthMethodOIDCOutput{}
	return &this
}

// NewCreateAuthMethodOIDCOutputWithDefaults instantiates a new CreateAuthMethodOIDCOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodOIDCOutputWithDefaults() *CreateAuthMethodOIDCOutput {
	this := CreateAuthMethodOIDCOutput{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *CreateAuthMethodOIDCOutput) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodOIDCOutput) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *CreateAuthMethodOIDCOutput) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *CreateAuthMethodOIDCOutput) SetAccessId(v string) {
	o.AccessId = &v
}

func (o CreateAuthMethodOIDCOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access_id"] = o.AccessId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAuthMethodOIDCOutput struct {
	value *CreateAuthMethodOIDCOutput
	isSet bool
}

func (v NullableCreateAuthMethodOIDCOutput) Get() *CreateAuthMethodOIDCOutput {
	return v.value
}

func (v *NullableCreateAuthMethodOIDCOutput) Set(val *CreateAuthMethodOIDCOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodOIDCOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodOIDCOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodOIDCOutput(val *CreateAuthMethodOIDCOutput) *NullableCreateAuthMethodOIDCOutput {
	return &NullableCreateAuthMethodOIDCOutput{value: val, isSet: true}
}

func (v NullableCreateAuthMethodOIDCOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodOIDCOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


