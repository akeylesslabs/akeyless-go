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

// DescribePermissions describePermissions is a command that shows which permissions your have to a particular path.
type DescribePermissions struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Path to an object
	Path string `json:"path"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Type of object (item, am, role, target)
	Type string `json:"type"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDescribePermissions instantiates a new DescribePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribePermissions(path string, type_ string, ) *DescribePermissions {
	this := DescribePermissions{}
	var json bool = false
	this.Json = &json
	this.Path = path
	this.Type = type_
	return &this
}

// NewDescribePermissionsWithDefaults instantiates a new DescribePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribePermissionsWithDefaults() *DescribePermissions {
	this := DescribePermissions{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DescribePermissions) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePermissions) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DescribePermissions) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DescribePermissions) SetJson(v bool) {
	o.Json = &v
}

// GetPath returns the Path field value
func (o *DescribePermissions) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DescribePermissions) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DescribePermissions) SetPath(v string) {
	o.Path = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DescribePermissions) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePermissions) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DescribePermissions) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DescribePermissions) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value
func (o *DescribePermissions) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DescribePermissions) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DescribePermissions) SetType(v string) {
	o.Type = v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DescribePermissions) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribePermissions) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DescribePermissions) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DescribePermissions) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DescribePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDescribePermissions struct {
	value *DescribePermissions
	isSet bool
}

func (v NullableDescribePermissions) Get() *DescribePermissions {
	return v.value
}

func (v *NullableDescribePermissions) Set(val *DescribePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribePermissions(val *DescribePermissions) *NullableDescribePermissions {
	return &NullableDescribePermissions{value: val, isSet: true}
}

func (v NullableDescribePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


