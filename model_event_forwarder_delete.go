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

// EventForwarderDelete eventForwarderDelete is a command that delete event forwarder
type EventForwarderDelete struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// EventForwarder name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewEventForwarderDelete instantiates a new EventForwarderDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventForwarderDelete(name string, ) *EventForwarderDelete {
	this := EventForwarderDelete{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewEventForwarderDeleteWithDefaults instantiates a new EventForwarderDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventForwarderDeleteWithDefaults() *EventForwarderDelete {
	this := EventForwarderDelete{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *EventForwarderDelete) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderDelete) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *EventForwarderDelete) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *EventForwarderDelete) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *EventForwarderDelete) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventForwarderDelete) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventForwarderDelete) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EventForwarderDelete) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderDelete) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EventForwarderDelete) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EventForwarderDelete) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EventForwarderDelete) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventForwarderDelete) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EventForwarderDelete) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EventForwarderDelete) SetUidToken(v string) {
	o.UidToken = &v
}

func (o EventForwarderDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableEventForwarderDelete struct {
	value *EventForwarderDelete
	isSet bool
}

func (v NullableEventForwarderDelete) Get() *EventForwarderDelete {
	return v.value
}

func (v *NullableEventForwarderDelete) Set(val *EventForwarderDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableEventForwarderDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableEventForwarderDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventForwarderDelete(val *EventForwarderDelete) *NullableEventForwarderDelete {
	return &NullableEventForwarderDelete{value: val, isSet: true}
}

func (v NullableEventForwarderDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventForwarderDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


