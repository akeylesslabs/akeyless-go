/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// ClientData struct for ClientData
type ClientData struct {
	ClientCertificateData *string `json:"clientCertificateData,omitempty"`
	ClientKeyData *string `json:"clientKeyData,omitempty"`
}

// NewClientData instantiates a new ClientData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientData() *ClientData {
	this := ClientData{}
	return &this
}

// NewClientDataWithDefaults instantiates a new ClientData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientDataWithDefaults() *ClientData {
	this := ClientData{}
	return &this
}

// GetClientCertificateData returns the ClientCertificateData field value if set, zero value otherwise.
func (o *ClientData) GetClientCertificateData() string {
	if o == nil || o.ClientCertificateData == nil {
		var ret string
		return ret
	}
	return *o.ClientCertificateData
}

// GetClientCertificateDataOk returns a tuple with the ClientCertificateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientData) GetClientCertificateDataOk() (*string, bool) {
	if o == nil || o.ClientCertificateData == nil {
		return nil, false
	}
	return o.ClientCertificateData, true
}

// HasClientCertificateData returns a boolean if a field has been set.
func (o *ClientData) HasClientCertificateData() bool {
	if o != nil && o.ClientCertificateData != nil {
		return true
	}

	return false
}

// SetClientCertificateData gets a reference to the given string and assigns it to the ClientCertificateData field.
func (o *ClientData) SetClientCertificateData(v string) {
	o.ClientCertificateData = &v
}

// GetClientKeyData returns the ClientKeyData field value if set, zero value otherwise.
func (o *ClientData) GetClientKeyData() string {
	if o == nil || o.ClientKeyData == nil {
		var ret string
		return ret
	}
	return *o.ClientKeyData
}

// GetClientKeyDataOk returns a tuple with the ClientKeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientData) GetClientKeyDataOk() (*string, bool) {
	if o == nil || o.ClientKeyData == nil {
		return nil, false
	}
	return o.ClientKeyData, true
}

// HasClientKeyData returns a boolean if a field has been set.
func (o *ClientData) HasClientKeyData() bool {
	if o != nil && o.ClientKeyData != nil {
		return true
	}

	return false
}

// SetClientKeyData gets a reference to the given string and assigns it to the ClientKeyData field.
func (o *ClientData) SetClientKeyData(v string) {
	o.ClientKeyData = &v
}

func (o ClientData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientCertificateData != nil {
		toSerialize["clientCertificateData"] = o.ClientCertificateData
	}
	if o.ClientKeyData != nil {
		toSerialize["clientKeyData"] = o.ClientKeyData
	}
	return json.Marshal(toSerialize)
}

type NullableClientData struct {
	value *ClientData
	isSet bool
}

func (v NullableClientData) Get() *ClientData {
	return v.value
}

func (v *NullableClientData) Set(val *ClientData) {
	v.value = val
	v.isSet = true
}

func (v NullableClientData) IsSet() bool {
	return v.isSet
}

func (v *NullableClientData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientData(val *ClientData) *NullableClientData {
	return &NullableClientData{value: val, isSet: true}
}

func (v NullableClientData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


