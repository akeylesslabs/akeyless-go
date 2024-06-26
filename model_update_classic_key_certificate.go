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

// UpdateClassicKeyCertificate UpdateClassicKeyCertificate is a command that updates the certificate for a classic key
type UpdateClassicKeyCertificate struct {
	// PEM Certificate in a Base64 format. Used for updating RSA keys' certificates.
	CertFileData *string `json:"cert-file-data,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// ClassicKey name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateClassicKeyCertificate instantiates a new UpdateClassicKeyCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateClassicKeyCertificate(name string, ) *UpdateClassicKeyCertificate {
	this := UpdateClassicKeyCertificate{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewUpdateClassicKeyCertificateWithDefaults instantiates a new UpdateClassicKeyCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateClassicKeyCertificateWithDefaults() *UpdateClassicKeyCertificate {
	this := UpdateClassicKeyCertificate{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetCertFileData returns the CertFileData field value if set, zero value otherwise.
func (o *UpdateClassicKeyCertificate) GetCertFileData() string {
	if o == nil || o.CertFileData == nil {
		var ret string
		return ret
	}
	return *o.CertFileData
}

// GetCertFileDataOk returns a tuple with the CertFileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClassicKeyCertificate) GetCertFileDataOk() (*string, bool) {
	if o == nil || o.CertFileData == nil {
		return nil, false
	}
	return o.CertFileData, true
}

// HasCertFileData returns a boolean if a field has been set.
func (o *UpdateClassicKeyCertificate) HasCertFileData() bool {
	if o != nil && o.CertFileData != nil {
		return true
	}

	return false
}

// SetCertFileData gets a reference to the given string and assigns it to the CertFileData field.
func (o *UpdateClassicKeyCertificate) SetCertFileData(v string) {
	o.CertFileData = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateClassicKeyCertificate) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClassicKeyCertificate) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateClassicKeyCertificate) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateClassicKeyCertificate) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *UpdateClassicKeyCertificate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateClassicKeyCertificate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateClassicKeyCertificate) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateClassicKeyCertificate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClassicKeyCertificate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateClassicKeyCertificate) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateClassicKeyCertificate) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateClassicKeyCertificate) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClassicKeyCertificate) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateClassicKeyCertificate) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateClassicKeyCertificate) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateClassicKeyCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertFileData != nil {
		toSerialize["cert-file-data"] = o.CertFileData
	}
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

type NullableUpdateClassicKeyCertificate struct {
	value *UpdateClassicKeyCertificate
	isSet bool
}

func (v NullableUpdateClassicKeyCertificate) Get() *UpdateClassicKeyCertificate {
	return v.value
}

func (v *NullableUpdateClassicKeyCertificate) Set(val *UpdateClassicKeyCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateClassicKeyCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateClassicKeyCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateClassicKeyCertificate(val *UpdateClassicKeyCertificate) *NullableUpdateClassicKeyCertificate {
	return &NullableUpdateClassicKeyCertificate{value: val, isSet: true}
}

func (v NullableUpdateClassicKeyCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateClassicKeyCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


