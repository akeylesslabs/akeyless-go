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

// ProvisionCertificate provisionCertificate is a command that provisions a certificate content to a target
type ProvisionCertificate struct {
	// Certificate display ID
	DisplayId *string `json:"display-id,omitempty"`
	// Certificate item ID
	ItemId *int64 `json:"item-id,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Certificate name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewProvisionCertificate instantiates a new ProvisionCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionCertificate(name string, ) *ProvisionCertificate {
	this := ProvisionCertificate{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewProvisionCertificateWithDefaults instantiates a new ProvisionCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionCertificateWithDefaults() *ProvisionCertificate {
	this := ProvisionCertificate{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *ProvisionCertificate) GetDisplayId() string {
	if o == nil || o.DisplayId == nil {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionCertificate) GetDisplayIdOk() (*string, bool) {
	if o == nil || o.DisplayId == nil {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *ProvisionCertificate) HasDisplayId() bool {
	if o != nil && o.DisplayId != nil {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *ProvisionCertificate) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ProvisionCertificate) GetItemId() int64 {
	if o == nil || o.ItemId == nil {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionCertificate) GetItemIdOk() (*int64, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ProvisionCertificate) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *ProvisionCertificate) SetItemId(v int64) {
	o.ItemId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ProvisionCertificate) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionCertificate) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ProvisionCertificate) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *ProvisionCertificate) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *ProvisionCertificate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProvisionCertificate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProvisionCertificate) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProvisionCertificate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionCertificate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProvisionCertificate) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProvisionCertificate) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *ProvisionCertificate) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionCertificate) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *ProvisionCertificate) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *ProvisionCertificate) SetUidToken(v string) {
	o.UidToken = &v
}

func (o ProvisionCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayId != nil {
		toSerialize["display-id"] = o.DisplayId
	}
	if o.ItemId != nil {
		toSerialize["item-id"] = o.ItemId
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

type NullableProvisionCertificate struct {
	value *ProvisionCertificate
	isSet bool
}

func (v NullableProvisionCertificate) Get() *ProvisionCertificate {
	return v.value
}

func (v *NullableProvisionCertificate) Set(val *ProvisionCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionCertificate(val *ProvisionCertificate) *NullableProvisionCertificate {
	return &NullableProvisionCertificate{value: val, isSet: true}
}

func (v NullableProvisionCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


