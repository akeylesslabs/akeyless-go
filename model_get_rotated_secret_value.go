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

// GetRotatedSecretValue getRotatedSecretValue is a command that creates a new Akeyless profile. [Deprecated: Use rotated-secret get-value command]
type GetRotatedSecretValue struct {
	// Get rotated secret value of specific Host (relevant only for Linked Target)
	Host *string `json:"host,omitempty"`
	// Retrieve the Secret value without checking the Gateway's cache [true/false]. This flag is only relevant when using the RestAPI
	IgnoreCache *string `json:"ignore-cache,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Secret name
	Names string `json:"names"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Secret version
	Version *int32 `json:"version,omitempty"`
}

// NewGetRotatedSecretValue instantiates a new GetRotatedSecretValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRotatedSecretValue(names string, ) *GetRotatedSecretValue {
	this := GetRotatedSecretValue{}
	var ignoreCache string = "false"
	this.IgnoreCache = &ignoreCache
	var json bool = false
	this.Json = &json
	this.Names = names
	return &this
}

// NewGetRotatedSecretValueWithDefaults instantiates a new GetRotatedSecretValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRotatedSecretValueWithDefaults() *GetRotatedSecretValue {
	this := GetRotatedSecretValue{}
	var ignoreCache string = "false"
	this.IgnoreCache = &ignoreCache
	var json bool = false
	this.Json = &json
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GetRotatedSecretValue) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRotatedSecretValue) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GetRotatedSecretValue) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *GetRotatedSecretValue) SetHost(v string) {
	o.Host = &v
}

// GetIgnoreCache returns the IgnoreCache field value if set, zero value otherwise.
func (o *GetRotatedSecretValue) GetIgnoreCache() string {
	if o == nil || o.IgnoreCache == nil {
		var ret string
		return ret
	}
	return *o.IgnoreCache
}

// GetIgnoreCacheOk returns a tuple with the IgnoreCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRotatedSecretValue) GetIgnoreCacheOk() (*string, bool) {
	if o == nil || o.IgnoreCache == nil {
		return nil, false
	}
	return o.IgnoreCache, true
}

// HasIgnoreCache returns a boolean if a field has been set.
func (o *GetRotatedSecretValue) HasIgnoreCache() bool {
	if o != nil && o.IgnoreCache != nil {
		return true
	}

	return false
}

// SetIgnoreCache gets a reference to the given string and assigns it to the IgnoreCache field.
func (o *GetRotatedSecretValue) SetIgnoreCache(v string) {
	o.IgnoreCache = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GetRotatedSecretValue) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRotatedSecretValue) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GetRotatedSecretValue) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GetRotatedSecretValue) SetJson(v bool) {
	o.Json = &v
}

// GetNames returns the Names field value
func (o *GetRotatedSecretValue) GetNames() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *GetRotatedSecretValue) GetNamesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value
func (o *GetRotatedSecretValue) SetNames(v string) {
	o.Names = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetRotatedSecretValue) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRotatedSecretValue) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetRotatedSecretValue) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetRotatedSecretValue) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GetRotatedSecretValue) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRotatedSecretValue) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GetRotatedSecretValue) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GetRotatedSecretValue) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetRotatedSecretValue) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRotatedSecretValue) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetRotatedSecretValue) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *GetRotatedSecretValue) SetVersion(v int32) {
	o.Version = &v
}

func (o GetRotatedSecretValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.IgnoreCache != nil {
		toSerialize["ignore-cache"] = o.IgnoreCache
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["names"] = o.Names
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableGetRotatedSecretValue struct {
	value *GetRotatedSecretValue
	isSet bool
}

func (v NullableGetRotatedSecretValue) Get() *GetRotatedSecretValue {
	return v.value
}

func (v *NullableGetRotatedSecretValue) Set(val *GetRotatedSecretValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRotatedSecretValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRotatedSecretValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRotatedSecretValue(val *GetRotatedSecretValue) *NullableGetRotatedSecretValue {
	return &NullableGetRotatedSecretValue{value: val, isSet: true}
}

func (v NullableGetRotatedSecretValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRotatedSecretValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


