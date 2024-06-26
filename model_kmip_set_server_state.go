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

// KmipSetServerState kmipSetServerState is a command that sets the environment state to active/inactive.
type KmipSetServerState struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	State string `json:"state"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewKmipSetServerState instantiates a new KmipSetServerState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipSetServerState(state string, ) *KmipSetServerState {
	this := KmipSetServerState{}
	var json bool = false
	this.Json = &json
	this.State = state
	return &this
}

// NewKmipSetServerStateWithDefaults instantiates a new KmipSetServerState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipSetServerStateWithDefaults() *KmipSetServerState {
	this := KmipSetServerState{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *KmipSetServerState) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipSetServerState) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *KmipSetServerState) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *KmipSetServerState) SetJson(v bool) {
	o.Json = &v
}

// GetState returns the State field value
func (o *KmipSetServerState) GetState() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *KmipSetServerState) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *KmipSetServerState) SetState(v string) {
	o.State = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KmipSetServerState) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipSetServerState) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KmipSetServerState) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KmipSetServerState) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *KmipSetServerState) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipSetServerState) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *KmipSetServerState) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *KmipSetServerState) SetUidToken(v string) {
	o.UidToken = &v
}

func (o KmipSetServerState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableKmipSetServerState struct {
	value *KmipSetServerState
	isSet bool
}

func (v NullableKmipSetServerState) Get() *KmipSetServerState {
	return v.value
}

func (v *NullableKmipSetServerState) Set(val *KmipSetServerState) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipSetServerState) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipSetServerState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipSetServerState(val *KmipSetServerState) *NullableKmipSetServerState {
	return &NullableKmipSetServerState{value: val, isSet: true}
}

func (v NullableKmipSetServerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipSetServerState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


