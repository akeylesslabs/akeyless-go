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

// GatewayDeleteMigration gatewayDeleteMigration is a command that delete migration
type GatewayDeleteMigration struct {
	// Migration ID
	Id string `json:"id"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayDeleteMigration instantiates a new GatewayDeleteMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayDeleteMigration(id string, ) *GatewayDeleteMigration {
	this := GatewayDeleteMigration{}
	this.Id = id
	var json bool = false
	this.Json = &json
	return &this
}

// NewGatewayDeleteMigrationWithDefaults instantiates a new GatewayDeleteMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayDeleteMigrationWithDefaults() *GatewayDeleteMigration {
	this := GatewayDeleteMigration{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetId returns the Id field value
func (o *GatewayDeleteMigration) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GatewayDeleteMigration) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GatewayDeleteMigration) SetId(v string) {
	o.Id = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayDeleteMigration) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteMigration) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayDeleteMigration) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayDeleteMigration) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayDeleteMigration) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteMigration) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayDeleteMigration) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayDeleteMigration) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayDeleteMigration) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteMigration) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayDeleteMigration) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayDeleteMigration) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayDeleteMigration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayDeleteMigration struct {
	value *GatewayDeleteMigration
	isSet bool
}

func (v NullableGatewayDeleteMigration) Get() *GatewayDeleteMigration {
	return v.value
}

func (v *NullableGatewayDeleteMigration) Set(val *GatewayDeleteMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayDeleteMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayDeleteMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayDeleteMigration(val *GatewayDeleteMigration) *NullableGatewayDeleteMigration {
	return &NullableGatewayDeleteMigration{value: val, isSet: true}
}

func (v NullableGatewayDeleteMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayDeleteMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


