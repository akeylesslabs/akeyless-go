/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// UpdateGroupOutput struct for UpdateGroupOutput
type UpdateGroupOutput struct {
	GroupName *string `json:"group_name,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewUpdateGroupOutput instantiates a new UpdateGroupOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGroupOutput() *UpdateGroupOutput {
	this := UpdateGroupOutput{}
	return &this
}

// NewUpdateGroupOutputWithDefaults instantiates a new UpdateGroupOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGroupOutputWithDefaults() *UpdateGroupOutput {
	this := UpdateGroupOutput{}
	return &this
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *UpdateGroupOutput) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupOutput) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *UpdateGroupOutput) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *UpdateGroupOutput) SetGroupName(v string) {
	o.GroupName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateGroupOutput) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupOutput) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateGroupOutput) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateGroupOutput) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateGroupOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGroupOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateGroupOutput) SetName(v string) {
	o.Name = &v
}

func (o UpdateGroupOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupName != nil {
		toSerialize["group_name"] = o.GroupName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateGroupOutput struct {
	value *UpdateGroupOutput
	isSet bool
}

func (v NullableUpdateGroupOutput) Get() *UpdateGroupOutput {
	return v.value
}

func (v *NullableUpdateGroupOutput) Set(val *UpdateGroupOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupOutput(val *UpdateGroupOutput) *NullableUpdateGroupOutput {
	return &NullableUpdateGroupOutput{value: val, isSet: true}
}

func (v NullableUpdateGroupOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


