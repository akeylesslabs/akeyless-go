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

// checks if the GatewayGetK8SAuthConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayGetK8SAuthConfig{}

// GatewayGetK8SAuthConfig gatewayGetK8SAuth is a command that gets k8s auth config
type GatewayGetK8SAuthConfig struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// K8S Auth config name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayGetK8SAuthConfig instantiates a new GatewayGetK8SAuthConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayGetK8SAuthConfig(name string) *GatewayGetK8SAuthConfig {
	this := GatewayGetK8SAuthConfig{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewGatewayGetK8SAuthConfigWithDefaults instantiates a new GatewayGetK8SAuthConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayGetK8SAuthConfigWithDefaults() *GatewayGetK8SAuthConfig {
	this := GatewayGetK8SAuthConfig{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfig) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfig) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfig) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayGetK8SAuthConfig) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayGetK8SAuthConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayGetK8SAuthConfig) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfig) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfig) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfig) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayGetK8SAuthConfig) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayGetK8SAuthConfig) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayGetK8SAuthConfig) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayGetK8SAuthConfig) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayGetK8SAuthConfig) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayGetK8SAuthConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayGetK8SAuthConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableGatewayGetK8SAuthConfig struct {
	value *GatewayGetK8SAuthConfig
	isSet bool
}

func (v NullableGatewayGetK8SAuthConfig) Get() *GatewayGetK8SAuthConfig {
	return v.value
}

func (v *NullableGatewayGetK8SAuthConfig) Set(val *GatewayGetK8SAuthConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayGetK8SAuthConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayGetK8SAuthConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayGetK8SAuthConfig(val *GatewayGetK8SAuthConfig) *NullableGatewayGetK8SAuthConfig {
	return &NullableGatewayGetK8SAuthConfig{value: val, isSet: true}
}

func (v NullableGatewayGetK8SAuthConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayGetK8SAuthConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


