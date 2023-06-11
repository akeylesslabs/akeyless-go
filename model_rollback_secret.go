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

// RollbackSecret struct for RollbackSecret
type RollbackSecret struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Secret name
	Name string `json:"name"`
	// Old secret version to rollback to
	OldVersion int32 `json:"old-version"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewRollbackSecret instantiates a new RollbackSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackSecret(name string, oldVersion int32) *RollbackSecret {
	this := RollbackSecret{}
	var json bool = false
	this.Json = &json
	this.Name = name
	this.OldVersion = oldVersion
	return &this
}

// NewRollbackSecretWithDefaults instantiates a new RollbackSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackSecretWithDefaults() *RollbackSecret {
	this := RollbackSecret{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *RollbackSecret) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackSecret) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *RollbackSecret) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *RollbackSecret) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *RollbackSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RollbackSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RollbackSecret) SetName(v string) {
	o.Name = v
}

// GetOldVersion returns the OldVersion field value
func (o *RollbackSecret) GetOldVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OldVersion
}

// GetOldVersionOk returns a tuple with the OldVersion field value
// and a boolean to check if the value has been set.
func (o *RollbackSecret) GetOldVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldVersion, true
}

// SetOldVersion sets field value
func (o *RollbackSecret) SetOldVersion(v int32) {
	o.OldVersion = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RollbackSecret) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackSecret) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RollbackSecret) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RollbackSecret) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *RollbackSecret) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackSecret) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *RollbackSecret) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *RollbackSecret) SetUidToken(v string) {
	o.UidToken = &v
}

func (o RollbackSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["old-version"] = o.OldVersion
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableRollbackSecret struct {
	value *RollbackSecret
	isSet bool
}

func (v NullableRollbackSecret) Get() *RollbackSecret {
	return v.value
}

func (v *NullableRollbackSecret) Set(val *RollbackSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackSecret(val *RollbackSecret) *NullableRollbackSecret {
	return &NullableRollbackSecret{value: val, isSet: true}
}

func (v NullableRollbackSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


