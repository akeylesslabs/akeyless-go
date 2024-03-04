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

// GatewayGetTmpUsers gatewayGetTmpUsers is a command that returns gateway configuration [Deprecated: Use dynamic-secret-tmp-creds-get command]
type GatewayGetTmpUsers struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayGetTmpUsers instantiates a new GatewayGetTmpUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayGetTmpUsers(name string, ) *GatewayGetTmpUsers {
	this := GatewayGetTmpUsers{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewGatewayGetTmpUsersWithDefaults instantiates a new GatewayGetTmpUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayGetTmpUsersWithDefaults() *GatewayGetTmpUsers {
	this := GatewayGetTmpUsers{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayGetTmpUsers) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetTmpUsers) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayGetTmpUsers) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayGetTmpUsers) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayGetTmpUsers) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayGetTmpUsers) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayGetTmpUsers) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayGetTmpUsers) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetTmpUsers) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayGetTmpUsers) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayGetTmpUsers) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayGetTmpUsers) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetTmpUsers) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayGetTmpUsers) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayGetTmpUsers) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayGetTmpUsers) MarshalJSON() ([]byte, error) {
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

type NullableGatewayGetTmpUsers struct {
	value *GatewayGetTmpUsers
	isSet bool
}

func (v NullableGatewayGetTmpUsers) Get() *GatewayGetTmpUsers {
	return v.value
}

func (v *NullableGatewayGetTmpUsers) Set(val *GatewayGetTmpUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayGetTmpUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayGetTmpUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayGetTmpUsers(val *GatewayGetTmpUsers) *NullableGatewayGetTmpUsers {
	return &NullableGatewayGetTmpUsers{value: val, isSet: true}
}

func (v NullableGatewayGetTmpUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayGetTmpUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


