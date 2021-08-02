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

// KMIPConfigPart struct for KMIPConfigPart
type KMIPConfigPart struct {
	Clients *map[string]KMIPClient `json:"clients,omitempty"`
	ServerEnc *[]int32 `json:"server_enc,omitempty"`
}

// NewKMIPConfigPart instantiates a new KMIPConfigPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKMIPConfigPart() *KMIPConfigPart {
	this := KMIPConfigPart{}
	return &this
}

// NewKMIPConfigPartWithDefaults instantiates a new KMIPConfigPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKMIPConfigPartWithDefaults() *KMIPConfigPart {
	this := KMIPConfigPart{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *KMIPConfigPart) GetClients() map[string]KMIPClient {
	if o == nil || o.Clients == nil {
		var ret map[string]KMIPClient
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KMIPConfigPart) GetClientsOk() (*map[string]KMIPClient, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *KMIPConfigPart) HasClients() bool {
	if o != nil && o.Clients != nil {
		return true
	}

	return false
}

// SetClients gets a reference to the given map[string]KMIPClient and assigns it to the Clients field.
func (o *KMIPConfigPart) SetClients(v map[string]KMIPClient) {
	o.Clients = &v
}

// GetServerEnc returns the ServerEnc field value if set, zero value otherwise.
func (o *KMIPConfigPart) GetServerEnc() []int32 {
	if o == nil || o.ServerEnc == nil {
		var ret []int32
		return ret
	}
	return *o.ServerEnc
}

// GetServerEncOk returns a tuple with the ServerEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KMIPConfigPart) GetServerEncOk() (*[]int32, bool) {
	if o == nil || o.ServerEnc == nil {
		return nil, false
	}
	return o.ServerEnc, true
}

// HasServerEnc returns a boolean if a field has been set.
func (o *KMIPConfigPart) HasServerEnc() bool {
	if o != nil && o.ServerEnc != nil {
		return true
	}

	return false
}

// SetServerEnc gets a reference to the given []int32 and assigns it to the ServerEnc field.
func (o *KMIPConfigPart) SetServerEnc(v []int32) {
	o.ServerEnc = &v
}

func (o KMIPConfigPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}
	if o.ServerEnc != nil {
		toSerialize["server_enc"] = o.ServerEnc
	}
	return json.Marshal(toSerialize)
}

type NullableKMIPConfigPart struct {
	value *KMIPConfigPart
	isSet bool
}

func (v NullableKMIPConfigPart) Get() *KMIPConfigPart {
	return v.value
}

func (v *NullableKMIPConfigPart) Set(val *KMIPConfigPart) {
	v.value = val
	v.isSet = true
}

func (v NullableKMIPConfigPart) IsSet() bool {
	return v.isSet
}

func (v *NullableKMIPConfigPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKMIPConfigPart(val *KMIPConfigPart) *NullableKMIPConfigPart {
	return &NullableKMIPConfigPart{value: val, isSet: true}
}

func (v NullableKMIPConfigPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKMIPConfigPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


