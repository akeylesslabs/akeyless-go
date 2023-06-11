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

// checks if the CreateAuthMethodOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAuthMethodOutput{}

// CreateAuthMethodOutput struct for CreateAuthMethodOutput
type CreateAuthMethodOutput struct {
	AccessId *string `json:"access_id,omitempty"`
	AccessKey *string `json:"access_key,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewCreateAuthMethodOutput instantiates a new CreateAuthMethodOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodOutput() *CreateAuthMethodOutput {
	this := CreateAuthMethodOutput{}
	return &this
}

// NewCreateAuthMethodOutputWithDefaults instantiates a new CreateAuthMethodOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodOutputWithDefaults() *CreateAuthMethodOutput {
	this := CreateAuthMethodOutput{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *CreateAuthMethodOutput) GetAccessId() string {
	if o == nil || IsNil(o.AccessId) {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodOutput) GetAccessIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessId) {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *CreateAuthMethodOutput) HasAccessId() bool {
	if o != nil && !IsNil(o.AccessId) {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *CreateAuthMethodOutput) SetAccessId(v string) {
	o.AccessId = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *CreateAuthMethodOutput) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodOutput) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *CreateAuthMethodOutput) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *CreateAuthMethodOutput) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateAuthMethodOutput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodOutput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateAuthMethodOutput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateAuthMethodOutput) SetName(v string) {
	o.Name = &v
}

func (o CreateAuthMethodOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAuthMethodOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessId) {
		toSerialize["access_id"] = o.AccessId
	}
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateAuthMethodOutput struct {
	value *CreateAuthMethodOutput
	isSet bool
}

func (v NullableCreateAuthMethodOutput) Get() *CreateAuthMethodOutput {
	return v.value
}

func (v *NullableCreateAuthMethodOutput) Set(val *CreateAuthMethodOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodOutput(val *CreateAuthMethodOutput) *NullableCreateAuthMethodOutput {
	return &NullableCreateAuthMethodOutput{value: val, isSet: true}
}

func (v NullableCreateAuthMethodOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


