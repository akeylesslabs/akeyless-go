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

// DynamicSecretTmpCredsDelete dynamicSecretTmpCredsDelete is a command that deletes dynamic secret temp creds
type DynamicSecretTmpCredsDelete struct {
	// Host
	Host *string `json:"host,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Revoke All Temp Creds
	RevokeAll *bool `json:"revoke-all,omitempty"`
	// Soft Delete
	SoftDelete *bool `json:"soft-delete,omitempty"`
	// Tmp Creds ID
	TmpCredsId string `json:"tmp-creds-id"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDynamicSecretTmpCredsDelete instantiates a new DynamicSecretTmpCredsDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretTmpCredsDelete(name string, tmpCredsId string, ) *DynamicSecretTmpCredsDelete {
	this := DynamicSecretTmpCredsDelete{}
	var json bool = false
	this.Json = &json
	this.Name = name
	this.TmpCredsId = tmpCredsId
	return &this
}

// NewDynamicSecretTmpCredsDeleteWithDefaults instantiates a new DynamicSecretTmpCredsDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretTmpCredsDeleteWithDefaults() *DynamicSecretTmpCredsDelete {
	this := DynamicSecretTmpCredsDelete{}
	var json bool = false
	this.Json = &json
	var tmpCredsId string = "demo_default_tmp_creds_id_for_sdk_bc"
	this.TmpCredsId = tmpCredsId
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *DynamicSecretTmpCredsDelete) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretTmpCredsDelete) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *DynamicSecretTmpCredsDelete) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *DynamicSecretTmpCredsDelete) SetHost(v string) {
	o.Host = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretTmpCredsDelete) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretTmpCredsDelete) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretTmpCredsDelete) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretTmpCredsDelete) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretTmpCredsDelete) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretTmpCredsDelete) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretTmpCredsDelete) SetName(v string) {
	o.Name = v
}

// GetRevokeAll returns the RevokeAll field value if set, zero value otherwise.
func (o *DynamicSecretTmpCredsDelete) GetRevokeAll() bool {
	if o == nil || o.RevokeAll == nil {
		var ret bool
		return ret
	}
	return *o.RevokeAll
}

// GetRevokeAllOk returns a tuple with the RevokeAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretTmpCredsDelete) GetRevokeAllOk() (*bool, bool) {
	if o == nil || o.RevokeAll == nil {
		return nil, false
	}
	return o.RevokeAll, true
}

// HasRevokeAll returns a boolean if a field has been set.
func (o *DynamicSecretTmpCredsDelete) HasRevokeAll() bool {
	if o != nil && o.RevokeAll != nil {
		return true
	}

	return false
}

// SetRevokeAll gets a reference to the given bool and assigns it to the RevokeAll field.
func (o *DynamicSecretTmpCredsDelete) SetRevokeAll(v bool) {
	o.RevokeAll = &v
}

// GetSoftDelete returns the SoftDelete field value if set, zero value otherwise.
func (o *DynamicSecretTmpCredsDelete) GetSoftDelete() bool {
	if o == nil || o.SoftDelete == nil {
		var ret bool
		return ret
	}
	return *o.SoftDelete
}

// GetSoftDeleteOk returns a tuple with the SoftDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretTmpCredsDelete) GetSoftDeleteOk() (*bool, bool) {
	if o == nil || o.SoftDelete == nil {
		return nil, false
	}
	return o.SoftDelete, true
}

// HasSoftDelete returns a boolean if a field has been set.
func (o *DynamicSecretTmpCredsDelete) HasSoftDelete() bool {
	if o != nil && o.SoftDelete != nil {
		return true
	}

	return false
}

// SetSoftDelete gets a reference to the given bool and assigns it to the SoftDelete field.
func (o *DynamicSecretTmpCredsDelete) SetSoftDelete(v bool) {
	o.SoftDelete = &v
}

// GetTmpCredsId returns the TmpCredsId field value
func (o *DynamicSecretTmpCredsDelete) GetTmpCredsId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TmpCredsId
}

// GetTmpCredsIdOk returns a tuple with the TmpCredsId field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretTmpCredsDelete) GetTmpCredsIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TmpCredsId, true
}

// SetTmpCredsId sets field value
func (o *DynamicSecretTmpCredsDelete) SetTmpCredsId(v string) {
	o.TmpCredsId = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretTmpCredsDelete) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretTmpCredsDelete) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretTmpCredsDelete) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretTmpCredsDelete) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretTmpCredsDelete) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretTmpCredsDelete) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretTmpCredsDelete) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretTmpCredsDelete) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DynamicSecretTmpCredsDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.RevokeAll != nil {
		toSerialize["revoke-all"] = o.RevokeAll
	}
	if o.SoftDelete != nil {
		toSerialize["soft-delete"] = o.SoftDelete
	}
	if true {
		toSerialize["tmp-creds-id"] = o.TmpCredsId
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretTmpCredsDelete struct {
	value *DynamicSecretTmpCredsDelete
	isSet bool
}

func (v NullableDynamicSecretTmpCredsDelete) Get() *DynamicSecretTmpCredsDelete {
	return v.value
}

func (v *NullableDynamicSecretTmpCredsDelete) Set(val *DynamicSecretTmpCredsDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretTmpCredsDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretTmpCredsDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretTmpCredsDelete(val *DynamicSecretTmpCredsDelete) *NullableDynamicSecretTmpCredsDelete {
	return &NullableDynamicSecretTmpCredsDelete{value: val, isSet: true}
}

func (v NullableDynamicSecretTmpCredsDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretTmpCredsDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


