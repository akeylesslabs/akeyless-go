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

// checks if the ListRolesOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRolesOutput{}

// ListRolesOutput struct for ListRolesOutput
type ListRolesOutput struct {
	NextPage *string `json:"next_page,omitempty"`
	Roles []Role `json:"roles,omitempty"`
}

// NewListRolesOutput instantiates a new ListRolesOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRolesOutput() *ListRolesOutput {
	this := ListRolesOutput{}
	return &this
}

// NewListRolesOutputWithDefaults instantiates a new ListRolesOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRolesOutputWithDefaults() *ListRolesOutput {
	this := ListRolesOutput{}
	return &this
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *ListRolesOutput) GetNextPage() string {
	if o == nil || IsNil(o.NextPage) {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesOutput) GetNextPageOk() (*string, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *ListRolesOutput) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *ListRolesOutput) SetNextPage(v string) {
	o.NextPage = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ListRolesOutput) GetRoles() []Role {
	if o == nil || IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesOutput) GetRolesOk() ([]Role, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ListRolesOutput) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *ListRolesOutput) SetRoles(v []Role) {
	o.Roles = v
}

func (o ListRolesOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRolesOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPage) {
		toSerialize["next_page"] = o.NextPage
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableListRolesOutput struct {
	value *ListRolesOutput
	isSet bool
}

func (v NullableListRolesOutput) Get() *ListRolesOutput {
	return v.value
}

func (v *NullableListRolesOutput) Set(val *ListRolesOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListRolesOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListRolesOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRolesOutput(val *ListRolesOutput) *NullableListRolesOutput {
	return &NullableListRolesOutput{value: val, isSet: true}
}

func (v NullableListRolesOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRolesOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


