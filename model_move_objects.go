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

// MoveObjects struct for MoveObjects
type MoveObjects struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The objects type to move (item/auth_method/role)
	ObjectsType *string `json:"objects-type,omitempty"`
	// Source path to move the objects from
	Source string `json:"source"`
	// Target path to move the objects to
	Target string `json:"target"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewMoveObjects instantiates a new MoveObjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveObjects(source string, target string) *MoveObjects {
	this := MoveObjects{}
	var json bool = false
	this.Json = &json
	var objectsType string = "item"
	this.ObjectsType = &objectsType
	this.Source = source
	this.Target = target
	return &this
}

// NewMoveObjectsWithDefaults instantiates a new MoveObjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveObjectsWithDefaults() *MoveObjects {
	this := MoveObjects{}
	var json bool = false
	this.Json = &json
	var objectsType string = "item"
	this.ObjectsType = &objectsType
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *MoveObjects) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveObjects) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *MoveObjects) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *MoveObjects) SetJson(v bool) {
	o.Json = &v
}

// GetObjectsType returns the ObjectsType field value if set, zero value otherwise.
func (o *MoveObjects) GetObjectsType() string {
	if o == nil || o.ObjectsType == nil {
		var ret string
		return ret
	}
	return *o.ObjectsType
}

// GetObjectsTypeOk returns a tuple with the ObjectsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveObjects) GetObjectsTypeOk() (*string, bool) {
	if o == nil || o.ObjectsType == nil {
		return nil, false
	}
	return o.ObjectsType, true
}

// HasObjectsType returns a boolean if a field has been set.
func (o *MoveObjects) HasObjectsType() bool {
	if o != nil && o.ObjectsType != nil {
		return true
	}

	return false
}

// SetObjectsType gets a reference to the given string and assigns it to the ObjectsType field.
func (o *MoveObjects) SetObjectsType(v string) {
	o.ObjectsType = &v
}

// GetSource returns the Source field value
func (o *MoveObjects) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *MoveObjects) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *MoveObjects) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value
func (o *MoveObjects) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *MoveObjects) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *MoveObjects) SetTarget(v string) {
	o.Target = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *MoveObjects) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveObjects) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *MoveObjects) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *MoveObjects) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *MoveObjects) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveObjects) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *MoveObjects) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *MoveObjects) SetUidToken(v string) {
	o.UidToken = &v
}

func (o MoveObjects) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.ObjectsType != nil {
		toSerialize["objects-type"] = o.ObjectsType
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableMoveObjects struct {
	value *MoveObjects
	isSet bool
}

func (v NullableMoveObjects) Get() *MoveObjects {
	return v.value
}

func (v *NullableMoveObjects) Set(val *MoveObjects) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveObjects) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveObjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveObjects(val *MoveObjects) *NullableMoveObjects {
	return &NullableMoveObjects{value: val, isSet: true}
}

func (v NullableMoveObjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveObjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


