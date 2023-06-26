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

// EsmDelete esmDelete is a command that deletes a secret from an External Secrets Manager
type EsmDelete struct {
	// Name of the External Secrets Manager item
	EsmName string `json:"esm-name"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The external secret id (or name, for AWS, Azure or K8s targets) to delete
	SecretId string `json:"secret-id"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewEsmDelete instantiates a new EsmDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsmDelete(esmName string, secretId string, ) *EsmDelete {
	this := EsmDelete{}
	this.EsmName = esmName
	var json bool = false
	this.Json = &json
	this.SecretId = secretId
	return &this
}

// NewEsmDeleteWithDefaults instantiates a new EsmDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsmDeleteWithDefaults() *EsmDelete {
	this := EsmDelete{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetEsmName returns the EsmName field value
func (o *EsmDelete) GetEsmName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EsmName
}

// GetEsmNameOk returns a tuple with the EsmName field value
// and a boolean to check if the value has been set.
func (o *EsmDelete) GetEsmNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EsmName, true
}

// SetEsmName sets field value
func (o *EsmDelete) SetEsmName(v string) {
	o.EsmName = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *EsmDelete) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsmDelete) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *EsmDelete) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *EsmDelete) SetJson(v bool) {
	o.Json = &v
}

// GetSecretId returns the SecretId field value
func (o *EsmDelete) GetSecretId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value
// and a boolean to check if the value has been set.
func (o *EsmDelete) GetSecretIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecretId, true
}

// SetSecretId sets field value
func (o *EsmDelete) SetSecretId(v string) {
	o.SecretId = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EsmDelete) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsmDelete) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EsmDelete) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EsmDelete) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *EsmDelete) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsmDelete) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *EsmDelete) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *EsmDelete) SetUidToken(v string) {
	o.UidToken = &v
}

func (o EsmDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["esm-name"] = o.EsmName
	}
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
	return json.Marshal(toSerialize)
}

type NullableEsmDelete struct {
	value *EsmDelete
	isSet bool
}

func (v NullableEsmDelete) Get() *EsmDelete {
	return v.value
}

func (v *NullableEsmDelete) Set(val *EsmDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableEsmDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableEsmDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsmDelete(val *EsmDelete) *NullableEsmDelete {
	return &NullableEsmDelete{value: val, isSet: true}
}

func (v NullableEsmDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsmDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


