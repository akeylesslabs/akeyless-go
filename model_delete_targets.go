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

// checks if the DeleteTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTargets{}

// DeleteTargets struct for DeleteTargets
type DeleteTargets struct {
	// Enforce deletion
	ForceDeletion *bool `json:"force-deletion,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Path to delete the targets from
	Path string `json:"path"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDeleteTargets instantiates a new DeleteTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTargets(path string) *DeleteTargets {
	this := DeleteTargets{}
	var forceDeletion bool = false
	this.ForceDeletion = &forceDeletion
	var json bool = false
	this.Json = &json
	this.Path = path
	return &this
}

// NewDeleteTargetsWithDefaults instantiates a new DeleteTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTargetsWithDefaults() *DeleteTargets {
	this := DeleteTargets{}
	var forceDeletion bool = false
	this.ForceDeletion = &forceDeletion
	var json bool = false
	this.Json = &json
	return &this
}

// GetForceDeletion returns the ForceDeletion field value if set, zero value otherwise.
func (o *DeleteTargets) GetForceDeletion() bool {
	if o == nil || IsNil(o.ForceDeletion) {
		var ret bool
		return ret
	}
	return *o.ForceDeletion
}

// GetForceDeletionOk returns a tuple with the ForceDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTargets) GetForceDeletionOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceDeletion) {
		return nil, false
	}
	return o.ForceDeletion, true
}

// HasForceDeletion returns a boolean if a field has been set.
func (o *DeleteTargets) HasForceDeletion() bool {
	if o != nil && !IsNil(o.ForceDeletion) {
		return true
	}

	return false
}

// SetForceDeletion gets a reference to the given bool and assigns it to the ForceDeletion field.
func (o *DeleteTargets) SetForceDeletion(v bool) {
	o.ForceDeletion = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DeleteTargets) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTargets) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DeleteTargets) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DeleteTargets) SetJson(v bool) {
	o.Json = &v
}

// GetPath returns the Path field value
func (o *DeleteTargets) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DeleteTargets) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DeleteTargets) SetPath(v string) {
	o.Path = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DeleteTargets) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTargets) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DeleteTargets) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DeleteTargets) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DeleteTargets) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTargets) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DeleteTargets) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DeleteTargets) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DeleteTargets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTargets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForceDeletion) {
		toSerialize["force-deletion"] = o.ForceDeletion
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["path"] = o.Path
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableDeleteTargets struct {
	value *DeleteTargets
	isSet bool
}

func (v NullableDeleteTargets) Get() *DeleteTargets {
	return v.value
}

func (v *NullableDeleteTargets) Set(val *DeleteTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTargets(val *DeleteTargets) *NullableDeleteTargets {
	return &NullableDeleteTargets{value: val, isSet: true}
}

func (v NullableDeleteTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


