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

// RotateKey of it.
type RotateKey struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Key name
	Name string `json:"name"`
	// The new pem encoded certificate for the classic key. relevant only for keys provided by user ('bring-your-own-key')
	NewCertPemData *string `json:"new-cert-pem-data,omitempty"`
	// The new base64 encoded value for the classic key. relevant only for keys provided by user ('bring-your-own-key')
	NewKeyData *string `json:"new-key-data,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewRotateKey instantiates a new RotateKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotateKey(name string, ) *RotateKey {
	this := RotateKey{}
	this.Name = name
	return &this
}

// NewRotateKeyWithDefaults instantiates a new RotateKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotateKeyWithDefaults() *RotateKey {
	this := RotateKey{}
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *RotateKey) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotateKey) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *RotateKey) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *RotateKey) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *RotateKey) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RotateKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RotateKey) SetName(v string) {
	o.Name = v
}

// GetNewCertPemData returns the NewCertPemData field value if set, zero value otherwise.
func (o *RotateKey) GetNewCertPemData() string {
	if o == nil || o.NewCertPemData == nil {
		var ret string
		return ret
	}
	return *o.NewCertPemData
}

// GetNewCertPemDataOk returns a tuple with the NewCertPemData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotateKey) GetNewCertPemDataOk() (*string, bool) {
	if o == nil || o.NewCertPemData == nil {
		return nil, false
	}
	return o.NewCertPemData, true
}

// HasNewCertPemData returns a boolean if a field has been set.
func (o *RotateKey) HasNewCertPemData() bool {
	if o != nil && o.NewCertPemData != nil {
		return true
	}

	return false
}

// SetNewCertPemData gets a reference to the given string and assigns it to the NewCertPemData field.
func (o *RotateKey) SetNewCertPemData(v string) {
	o.NewCertPemData = &v
}

// GetNewKeyData returns the NewKeyData field value if set, zero value otherwise.
func (o *RotateKey) GetNewKeyData() string {
	if o == nil || o.NewKeyData == nil {
		var ret string
		return ret
	}
	return *o.NewKeyData
}

// GetNewKeyDataOk returns a tuple with the NewKeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotateKey) GetNewKeyDataOk() (*string, bool) {
	if o == nil || o.NewKeyData == nil {
		return nil, false
	}
	return o.NewKeyData, true
}

// HasNewKeyData returns a boolean if a field has been set.
func (o *RotateKey) HasNewKeyData() bool {
	if o != nil && o.NewKeyData != nil {
		return true
	}

	return false
}

// SetNewKeyData gets a reference to the given string and assigns it to the NewKeyData field.
func (o *RotateKey) SetNewKeyData(v string) {
	o.NewKeyData = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RotateKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotateKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RotateKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RotateKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *RotateKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotateKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *RotateKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *RotateKey) SetUidToken(v string) {
	o.UidToken = &v
}

func (o RotateKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewCertPemData != nil {
		toSerialize["new-cert-pem-data"] = o.NewCertPemData
	}
	if o.NewKeyData != nil {
		toSerialize["new-key-data"] = o.NewKeyData
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableRotateKey struct {
	value *RotateKey
	isSet bool
}

func (v NullableRotateKey) Get() *RotateKey {
	return v.value
}

func (v *NullableRotateKey) Set(val *RotateKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRotateKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRotateKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotateKey(val *RotateKey) *NullableRotateKey {
	return &NullableRotateKey{value: val, isSet: true}
}

func (v NullableRotateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

