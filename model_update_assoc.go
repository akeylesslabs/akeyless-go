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

// UpdateAssoc updateAssoc is a command that updates the sub-claims of an association between role and auth method.
type UpdateAssoc struct {
	// The association id to be updated
	AssocId string `json:"assoc-id"`
	// Treat sub claims as case-sensitive [true/false]
	CaseSensitive *string `json:"case-sensitive,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// key/val of sub claims, e.g group=admins,developers
	SubClaims *map[string]string `json:"sub-claims,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateAssoc instantiates a new UpdateAssoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAssoc(assocId string, ) *UpdateAssoc {
	this := UpdateAssoc{}
	this.AssocId = assocId
	var caseSensitive string = "true"
	this.CaseSensitive = &caseSensitive
	var json bool = false
	this.Json = &json
	return &this
}

// NewUpdateAssocWithDefaults instantiates a new UpdateAssoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAssocWithDefaults() *UpdateAssoc {
	this := UpdateAssoc{}
	var caseSensitive string = "true"
	this.CaseSensitive = &caseSensitive
	var json bool = false
	this.Json = &json
	return &this
}

// GetAssocId returns the AssocId field value
func (o *UpdateAssoc) GetAssocId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AssocId
}

// GetAssocIdOk returns a tuple with the AssocId field value
// and a boolean to check if the value has been set.
func (o *UpdateAssoc) GetAssocIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssocId, true
}

// SetAssocId sets field value
func (o *UpdateAssoc) SetAssocId(v string) {
	o.AssocId = v
}

// GetCaseSensitive returns the CaseSensitive field value if set, zero value otherwise.
func (o *UpdateAssoc) GetCaseSensitive() string {
	if o == nil || o.CaseSensitive == nil {
		var ret string
		return ret
	}
	return *o.CaseSensitive
}

// GetCaseSensitiveOk returns a tuple with the CaseSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAssoc) GetCaseSensitiveOk() (*string, bool) {
	if o == nil || o.CaseSensitive == nil {
		return nil, false
	}
	return o.CaseSensitive, true
}

// HasCaseSensitive returns a boolean if a field has been set.
func (o *UpdateAssoc) HasCaseSensitive() bool {
	if o != nil && o.CaseSensitive != nil {
		return true
	}

	return false
}

// SetCaseSensitive gets a reference to the given string and assigns it to the CaseSensitive field.
func (o *UpdateAssoc) SetCaseSensitive(v string) {
	o.CaseSensitive = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateAssoc) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAssoc) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateAssoc) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateAssoc) SetJson(v bool) {
	o.Json = &v
}

// GetSubClaims returns the SubClaims field value if set, zero value otherwise.
func (o *UpdateAssoc) GetSubClaims() map[string]string {
	if o == nil || o.SubClaims == nil {
		var ret map[string]string
		return ret
	}
	return *o.SubClaims
}

// GetSubClaimsOk returns a tuple with the SubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAssoc) GetSubClaimsOk() (*map[string]string, bool) {
	if o == nil || o.SubClaims == nil {
		return nil, false
	}
	return o.SubClaims, true
}

// HasSubClaims returns a boolean if a field has been set.
func (o *UpdateAssoc) HasSubClaims() bool {
	if o != nil && o.SubClaims != nil {
		return true
	}

	return false
}

// SetSubClaims gets a reference to the given map[string]string and assigns it to the SubClaims field.
func (o *UpdateAssoc) SetSubClaims(v map[string]string) {
	o.SubClaims = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateAssoc) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAssoc) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateAssoc) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateAssoc) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateAssoc) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAssoc) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateAssoc) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateAssoc) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateAssoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assoc-id"] = o.AssocId
	}
	if o.CaseSensitive != nil {
		toSerialize["case-sensitive"] = o.CaseSensitive
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.SubClaims != nil {
		toSerialize["sub-claims"] = o.SubClaims
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAssoc struct {
	value *UpdateAssoc
	isSet bool
}

func (v NullableUpdateAssoc) Get() *UpdateAssoc {
	return v.value
}

func (v *NullableUpdateAssoc) Set(val *UpdateAssoc) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAssoc) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAssoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAssoc(val *UpdateAssoc) *NullableUpdateAssoc {
	return &NullableUpdateAssoc{value: val, isSet: true}
}

func (v NullableUpdateAssoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAssoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


