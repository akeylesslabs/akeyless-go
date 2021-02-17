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

// GetSubAdminsListReplyObj struct for GetSubAdminsListReplyObj
type GetSubAdminsListReplyObj struct {
	SubAdmins *[]string `json:"sub_admins,omitempty"`
}

// NewGetSubAdminsListReplyObj instantiates a new GetSubAdminsListReplyObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAdminsListReplyObj() *GetSubAdminsListReplyObj {
	this := GetSubAdminsListReplyObj{}
	return &this
}

// NewGetSubAdminsListReplyObjWithDefaults instantiates a new GetSubAdminsListReplyObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAdminsListReplyObjWithDefaults() *GetSubAdminsListReplyObj {
	this := GetSubAdminsListReplyObj{}
	return &this
}

// GetSubAdmins returns the SubAdmins field value if set, zero value otherwise.
func (o *GetSubAdminsListReplyObj) GetSubAdmins() []string {
	if o == nil || o.SubAdmins == nil {
		var ret []string
		return ret
	}
	return *o.SubAdmins
}

// GetSubAdminsOk returns a tuple with the SubAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAdminsListReplyObj) GetSubAdminsOk() (*[]string, bool) {
	if o == nil || o.SubAdmins == nil {
		return nil, false
	}
	return o.SubAdmins, true
}

// HasSubAdmins returns a boolean if a field has been set.
func (o *GetSubAdminsListReplyObj) HasSubAdmins() bool {
	if o != nil && o.SubAdmins != nil {
		return true
	}

	return false
}

// SetSubAdmins gets a reference to the given []string and assigns it to the SubAdmins field.
func (o *GetSubAdminsListReplyObj) SetSubAdmins(v []string) {
	o.SubAdmins = &v
}

func (o GetSubAdminsListReplyObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubAdmins != nil {
		toSerialize["sub_admins"] = o.SubAdmins
	}
	return json.Marshal(toSerialize)
}

type NullableGetSubAdminsListReplyObj struct {
	value *GetSubAdminsListReplyObj
	isSet bool
}

func (v NullableGetSubAdminsListReplyObj) Get() *GetSubAdminsListReplyObj {
	return v.value
}

func (v *NullableGetSubAdminsListReplyObj) Set(val *GetSubAdminsListReplyObj) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAdminsListReplyObj) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAdminsListReplyObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAdminsListReplyObj(val *GetSubAdminsListReplyObj) *NullableGetSubAdminsListReplyObj {
	return &NullableGetSubAdminsListReplyObj{value: val, isSet: true}
}

func (v NullableGetSubAdminsListReplyObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAdminsListReplyObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


