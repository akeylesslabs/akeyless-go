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

// checks if the DeleteAuthMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteAuthMethod{}

// DeleteAuthMethod struct for DeleteAuthMethod
type DeleteAuthMethod struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDeleteAuthMethod instantiates a new DeleteAuthMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAuthMethod(name string) *DeleteAuthMethod {
	this := DeleteAuthMethod{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewDeleteAuthMethodWithDefaults instantiates a new DeleteAuthMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAuthMethodWithDefaults() *DeleteAuthMethod {
	this := DeleteAuthMethod{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DeleteAuthMethod) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAuthMethod) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DeleteAuthMethod) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DeleteAuthMethod) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DeleteAuthMethod) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeleteAuthMethod) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeleteAuthMethod) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DeleteAuthMethod) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAuthMethod) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DeleteAuthMethod) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DeleteAuthMethod) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DeleteAuthMethod) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAuthMethod) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DeleteAuthMethod) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DeleteAuthMethod) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DeleteAuthMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteAuthMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

type NullableDeleteAuthMethod struct {
	value *DeleteAuthMethod
	isSet bool
}

func (v NullableDeleteAuthMethod) Get() *DeleteAuthMethod {
	return v.value
}

func (v *NullableDeleteAuthMethod) Set(val *DeleteAuthMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAuthMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAuthMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAuthMethod(val *DeleteAuthMethod) *NullableDeleteAuthMethod {
	return &NullableDeleteAuthMethod{value: val, isSet: true}
}

func (v NullableDeleteAuthMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAuthMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


