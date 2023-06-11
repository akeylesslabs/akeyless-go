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

// checks if the SetItemState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetItemState{}

// SetItemState struct for SetItemState
type SetItemState struct {
	// Desired item state (Enabled, Disabled)
	DesiredState string `json:"desired-state"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Current item name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The specific version you want to update: 0=item level state (default)
	Version *int32 `json:"version,omitempty"`
}

// NewSetItemState instantiates a new SetItemState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetItemState(desiredState string, name string) *SetItemState {
	this := SetItemState{}
	this.DesiredState = desiredState
	var json bool = false
	this.Json = &json
	this.Name = name
	var version int32 = 0
	this.Version = &version
	return &this
}

// NewSetItemStateWithDefaults instantiates a new SetItemState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetItemStateWithDefaults() *SetItemState {
	this := SetItemState{}
	var json bool = false
	this.Json = &json
	var version int32 = 0
	this.Version = &version
	return &this
}

// GetDesiredState returns the DesiredState field value
func (o *SetItemState) GetDesiredState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value
// and a boolean to check if the value has been set.
func (o *SetItemState) GetDesiredStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredState, true
}

// SetDesiredState sets field value
func (o *SetItemState) SetDesiredState(v string) {
	o.DesiredState = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *SetItemState) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetItemState) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *SetItemState) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *SetItemState) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *SetItemState) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SetItemState) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SetItemState) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SetItemState) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetItemState) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SetItemState) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SetItemState) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *SetItemState) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetItemState) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *SetItemState) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *SetItemState) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SetItemState) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetItemState) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SetItemState) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SetItemState) SetVersion(v int32) {
	o.Version = &v
}

func (o SetItemState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetItemState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["desired-state"] = o.DesiredState
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableSetItemState struct {
	value *SetItemState
	isSet bool
}

func (v NullableSetItemState) Get() *SetItemState {
	return v.value
}

func (v *NullableSetItemState) Set(val *SetItemState) {
	v.value = val
	v.isSet = true
}

func (v NullableSetItemState) IsSet() bool {
	return v.isSet
}

func (v *NullableSetItemState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetItemState(val *SetItemState) *NullableSetItemState {
	return &NullableSetItemState{value: val, isSet: true}
}

func (v NullableSetItemState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetItemState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


