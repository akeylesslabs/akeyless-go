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

// checks if the KmipServerSetup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KmipServerSetup{}

// KmipServerSetup struct for KmipServerSetup
type KmipServerSetup struct {
	// Server certificate TTL in days
	CertificateTtl *int64 `json:"certificate-ttl,omitempty"`
	// Hostname
	Hostname string `json:"hostname"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Root path of KMIP Resources
	Root string `json:"root"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewKmipServerSetup instantiates a new KmipServerSetup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipServerSetup(hostname string, root string) *KmipServerSetup {
	this := KmipServerSetup{}
	var certificateTtl int64 = 90
	this.CertificateTtl = &certificateTtl
	this.Hostname = hostname
	var json bool = false
	this.Json = &json
	this.Root = root
	return &this
}

// NewKmipServerSetupWithDefaults instantiates a new KmipServerSetup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipServerSetupWithDefaults() *KmipServerSetup {
	this := KmipServerSetup{}
	var certificateTtl int64 = 90
	this.CertificateTtl = &certificateTtl
	var json bool = false
	this.Json = &json
	return &this
}

// GetCertificateTtl returns the CertificateTtl field value if set, zero value otherwise.
func (o *KmipServerSetup) GetCertificateTtl() int64 {
	if o == nil || IsNil(o.CertificateTtl) {
		var ret int64
		return ret
	}
	return *o.CertificateTtl
}

// GetCertificateTtlOk returns a tuple with the CertificateTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipServerSetup) GetCertificateTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.CertificateTtl) {
		return nil, false
	}
	return o.CertificateTtl, true
}

// HasCertificateTtl returns a boolean if a field has been set.
func (o *KmipServerSetup) HasCertificateTtl() bool {
	if o != nil && !IsNil(o.CertificateTtl) {
		return true
	}

	return false
}

// SetCertificateTtl gets a reference to the given int64 and assigns it to the CertificateTtl field.
func (o *KmipServerSetup) SetCertificateTtl(v int64) {
	o.CertificateTtl = &v
}

// GetHostname returns the Hostname field value
func (o *KmipServerSetup) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *KmipServerSetup) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *KmipServerSetup) SetHostname(v string) {
	o.Hostname = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *KmipServerSetup) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipServerSetup) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *KmipServerSetup) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *KmipServerSetup) SetJson(v bool) {
	o.Json = &v
}

// GetRoot returns the Root field value
func (o *KmipServerSetup) GetRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Root
}

// GetRootOk returns a tuple with the Root field value
// and a boolean to check if the value has been set.
func (o *KmipServerSetup) GetRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Root, true
}

// SetRoot sets field value
func (o *KmipServerSetup) SetRoot(v string) {
	o.Root = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KmipServerSetup) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipServerSetup) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KmipServerSetup) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KmipServerSetup) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *KmipServerSetup) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipServerSetup) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *KmipServerSetup) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *KmipServerSetup) SetUidToken(v string) {
	o.UidToken = &v
}

func (o KmipServerSetup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KmipServerSetup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificateTtl) {
		toSerialize["certificate-ttl"] = o.CertificateTtl
	}
	toSerialize["hostname"] = o.Hostname
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["root"] = o.Root
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableKmipServerSetup struct {
	value *KmipServerSetup
	isSet bool
}

func (v NullableKmipServerSetup) Get() *KmipServerSetup {
	return v.value
}

func (v *NullableKmipServerSetup) Set(val *KmipServerSetup) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipServerSetup) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipServerSetup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipServerSetup(val *KmipServerSetup) *NullableKmipServerSetup {
	return &NullableKmipServerSetup{value: val, isSet: true}
}

func (v NullableKmipServerSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipServerSetup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


