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

// DescribeAssoc struct for DescribeAssoc
type DescribeAssoc struct {
	// The association id
	AssocId string `json:"assoc-id"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewDescribeAssoc instantiates a new DescribeAssoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAssoc(assocId string, ) *DescribeAssoc {
	this := DescribeAssoc{}
	this.AssocId = assocId
	var json bool = false
	this.Json = &json
	return &this
}

// NewDescribeAssocWithDefaults instantiates a new DescribeAssoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAssocWithDefaults() *DescribeAssoc {
	this := DescribeAssoc{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAssocId returns the AssocId field value
func (o *DescribeAssoc) GetAssocId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AssocId
}

// GetAssocIdOk returns a tuple with the AssocId field value
// and a boolean to check if the value has been set.
func (o *DescribeAssoc) GetAssocIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssocId, true
}

// SetAssocId sets field value
func (o *DescribeAssoc) SetAssocId(v string) {
	o.AssocId = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DescribeAssoc) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAssoc) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DescribeAssoc) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DescribeAssoc) SetJson(v bool) {
	o.Json = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DescribeAssoc) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAssoc) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DescribeAssoc) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DescribeAssoc) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DescribeAssoc) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAssoc) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DescribeAssoc) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DescribeAssoc) SetUidToken(v string) {
	o.UidToken = &v
}

func (o DescribeAssoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assoc-id"] = o.AssocId
	}
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

type NullableDescribeAssoc struct {
	value *DescribeAssoc
	isSet bool
}

func (v NullableDescribeAssoc) Get() *DescribeAssoc {
	return v.value
}

func (v *NullableDescribeAssoc) Set(val *DescribeAssoc) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAssoc) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAssoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAssoc(val *DescribeAssoc) *NullableDescribeAssoc {
	return &NullableDescribeAssoc{value: val, isSet: true}
}

func (v NullableDescribeAssoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAssoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


