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

// UscDelete uscDelete is a command that deletes a secret from a Universal Secrets Connector
type UscDelete struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The universal secrets id (or name, for AWS, Azure or K8s targets) to delete
	SecretId string `json:"secret-id"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Name of the Universal Secrets Connector item
	UscName string `json:"usc-name"`
}

// NewUscDelete instantiates a new UscDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUscDelete(secretId string, uscName string, ) *UscDelete {
	this := UscDelete{}
	var json bool = false
	this.Json = &json
	this.SecretId = secretId
	this.UscName = uscName
	return &this
}

// NewUscDeleteWithDefaults instantiates a new UscDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUscDeleteWithDefaults() *UscDelete {
	this := UscDelete{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UscDelete) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscDelete) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UscDelete) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UscDelete) SetJson(v bool) {
	o.Json = &v
}

// GetSecretId returns the SecretId field value
func (o *UscDelete) GetSecretId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value
// and a boolean to check if the value has been set.
func (o *UscDelete) GetSecretIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecretId, true
}

// SetSecretId sets field value
func (o *UscDelete) SetSecretId(v string) {
	o.SecretId = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UscDelete) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscDelete) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UscDelete) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UscDelete) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UscDelete) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscDelete) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UscDelete) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UscDelete) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUscName returns the UscName field value
func (o *UscDelete) GetUscName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UscName
}

// GetUscNameOk returns a tuple with the UscName field value
// and a boolean to check if the value has been set.
func (o *UscDelete) GetUscNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UscName, true
}

// SetUscName sets field value
func (o *UscDelete) SetUscName(v string) {
	o.UscName = v
}

func (o UscDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["secret-id"] = o.SecretId
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if true {
		toSerialize["usc-name"] = o.UscName
	}
	return json.Marshal(toSerialize)
}

type NullableUscDelete struct {
	value *UscDelete
	isSet bool
}

func (v NullableUscDelete) Get() *UscDelete {
	return v.value
}

func (v *NullableUscDelete) Set(val *UscDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableUscDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableUscDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUscDelete(val *UscDelete) *NullableUscDelete {
	return &NullableUscDelete{value: val, isSet: true}
}

func (v NullableUscDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUscDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


