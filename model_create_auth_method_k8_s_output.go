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

// checks if the CreateAuthMethodK8SOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAuthMethodK8SOutput{}

// CreateAuthMethodK8SOutput struct for CreateAuthMethodK8SOutput
type CreateAuthMethodK8SOutput struct {
	AccessId *string `json:"access_id,omitempty"`
	PrvKey *string `json:"prv_key,omitempty"`
}

// NewCreateAuthMethodK8SOutput instantiates a new CreateAuthMethodK8SOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodK8SOutput() *CreateAuthMethodK8SOutput {
	this := CreateAuthMethodK8SOutput{}
	return &this
}

// NewCreateAuthMethodK8SOutputWithDefaults instantiates a new CreateAuthMethodK8SOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodK8SOutputWithDefaults() *CreateAuthMethodK8SOutput {
	this := CreateAuthMethodK8SOutput{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *CreateAuthMethodK8SOutput) GetAccessId() string {
	if o == nil || IsNil(o.AccessId) {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodK8SOutput) GetAccessIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessId) {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *CreateAuthMethodK8SOutput) HasAccessId() bool {
	if o != nil && !IsNil(o.AccessId) {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *CreateAuthMethodK8SOutput) SetAccessId(v string) {
	o.AccessId = &v
}

// GetPrvKey returns the PrvKey field value if set, zero value otherwise.
func (o *CreateAuthMethodK8SOutput) GetPrvKey() string {
	if o == nil || IsNil(o.PrvKey) {
		var ret string
		return ret
	}
	return *o.PrvKey
}

// GetPrvKeyOk returns a tuple with the PrvKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodK8SOutput) GetPrvKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrvKey) {
		return nil, false
	}
	return o.PrvKey, true
}

// HasPrvKey returns a boolean if a field has been set.
func (o *CreateAuthMethodK8SOutput) HasPrvKey() bool {
	if o != nil && !IsNil(o.PrvKey) {
		return true
	}

	return false
}

// SetPrvKey gets a reference to the given string and assigns it to the PrvKey field.
func (o *CreateAuthMethodK8SOutput) SetPrvKey(v string) {
	o.PrvKey = &v
}

func (o CreateAuthMethodK8SOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAuthMethodK8SOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessId) {
		toSerialize["access_id"] = o.AccessId
	}
	if !IsNil(o.PrvKey) {
		toSerialize["prv_key"] = o.PrvKey
	}
	return toSerialize, nil
}

type NullableCreateAuthMethodK8SOutput struct {
	value *CreateAuthMethodK8SOutput
	isSet bool
}

func (v NullableCreateAuthMethodK8SOutput) Get() *CreateAuthMethodK8SOutput {
	return v.value
}

func (v *NullableCreateAuthMethodK8SOutput) Set(val *CreateAuthMethodK8SOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodK8SOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodK8SOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodK8SOutput(val *CreateAuthMethodK8SOutput) *NullableCreateAuthMethodK8SOutput {
	return &NullableCreateAuthMethodK8SOutput{value: val, isSet: true}
}

func (v NullableCreateAuthMethodK8SOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodK8SOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


