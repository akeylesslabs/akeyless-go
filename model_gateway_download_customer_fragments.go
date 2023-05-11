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

// GatewayDownloadCustomerFragments gatewayDownloadCustomerFragments is a command that downloads customer fragments
type GatewayDownloadCustomerFragments struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayDownloadCustomerFragments instantiates a new GatewayDownloadCustomerFragments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayDownloadCustomerFragments() *GatewayDownloadCustomerFragments {
	this := GatewayDownloadCustomerFragments{}
	var json bool = false
	this.Json = &json
	return &this
}

// NewGatewayDownloadCustomerFragmentsWithDefaults instantiates a new GatewayDownloadCustomerFragments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayDownloadCustomerFragmentsWithDefaults() *GatewayDownloadCustomerFragments {
	this := GatewayDownloadCustomerFragments{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayDownloadCustomerFragments) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDownloadCustomerFragments) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayDownloadCustomerFragments) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayDownloadCustomerFragments) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayDownloadCustomerFragments) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDownloadCustomerFragments) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayDownloadCustomerFragments) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayDownloadCustomerFragments) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayDownloadCustomerFragments) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDownloadCustomerFragments) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayDownloadCustomerFragments) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayDownloadCustomerFragments) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayDownloadCustomerFragments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableGatewayDownloadCustomerFragments struct {
	value *GatewayDownloadCustomerFragments
	isSet bool
}

func (v NullableGatewayDownloadCustomerFragments) Get() *GatewayDownloadCustomerFragments {
	return v.value
}

func (v *NullableGatewayDownloadCustomerFragments) Set(val *GatewayDownloadCustomerFragments) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayDownloadCustomerFragments) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayDownloadCustomerFragments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayDownloadCustomerFragments(val *GatewayDownloadCustomerFragments) *NullableGatewayDownloadCustomerFragments {
	return &NullableGatewayDownloadCustomerFragments{value: val, isSet: true}
}

func (v NullableGatewayDownloadCustomerFragments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayDownloadCustomerFragments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


