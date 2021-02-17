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

// GatewayListProducers gatewayListProducers is a command that returns producer
type GatewayListProducers struct {
	// Gateway url
	GatewayUrl *string `json:"gateway-url,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayListProducers instantiates a new GatewayListProducers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayListProducers() *GatewayListProducers {
	this := GatewayListProducers{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	return &this
}

// NewGatewayListProducersWithDefaults instantiates a new GatewayListProducers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayListProducersWithDefaults() *GatewayListProducers {
	this := GatewayListProducers{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	return &this
}

// GetGatewayUrl returns the GatewayUrl field value if set, zero value otherwise.
func (o *GatewayListProducers) GetGatewayUrl() string {
	if o == nil || o.GatewayUrl == nil {
		var ret string
		return ret
	}
	return *o.GatewayUrl
}

// GetGatewayUrlOk returns a tuple with the GatewayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListProducers) GetGatewayUrlOk() (*string, bool) {
	if o == nil || o.GatewayUrl == nil {
		return nil, false
	}
	return o.GatewayUrl, true
}

// HasGatewayUrl returns a boolean if a field has been set.
func (o *GatewayListProducers) HasGatewayUrl() bool {
	if o != nil && o.GatewayUrl != nil {
		return true
	}

	return false
}

// SetGatewayUrl gets a reference to the given string and assigns it to the GatewayUrl field.
func (o *GatewayListProducers) SetGatewayUrl(v string) {
	o.GatewayUrl = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayListProducers) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListProducers) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayListProducers) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayListProducers) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayListProducers) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayListProducers) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayListProducers) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayListProducers) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayListProducers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayUrl != nil {
		toSerialize["gateway-url"] = o.GatewayUrl
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayListProducers struct {
	value *GatewayListProducers
	isSet bool
}

func (v NullableGatewayListProducers) Get() *GatewayListProducers {
	return v.value
}

func (v *NullableGatewayListProducers) Set(val *GatewayListProducers) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayListProducers) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayListProducers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayListProducers(val *GatewayListProducers) *NullableGatewayListProducers {
	return &NullableGatewayListProducers{value: val, isSet: true}
}

func (v NullableGatewayListProducers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayListProducers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


