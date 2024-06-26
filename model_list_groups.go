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

// ListGroups struct for ListGroups
type ListGroups struct {
	// Filter by auth method name or part of it
	Filter *string `json:"filter,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Next page reference
	PaginationToken *string `json:"pagination-token,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewListGroups instantiates a new ListGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGroups() *ListGroups {
	this := ListGroups{}
	var json bool = false
	this.Json = &json
	return &this
}

// NewListGroupsWithDefaults instantiates a new ListGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGroupsWithDefaults() *ListGroups {
	this := ListGroups{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ListGroups) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGroups) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ListGroups) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *ListGroups) SetFilter(v string) {
	o.Filter = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ListGroups) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGroups) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ListGroups) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *ListGroups) SetJson(v bool) {
	o.Json = &v
}

// GetPaginationToken returns the PaginationToken field value if set, zero value otherwise.
func (o *ListGroups) GetPaginationToken() string {
	if o == nil || o.PaginationToken == nil {
		var ret string
		return ret
	}
	return *o.PaginationToken
}

// GetPaginationTokenOk returns a tuple with the PaginationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGroups) GetPaginationTokenOk() (*string, bool) {
	if o == nil || o.PaginationToken == nil {
		return nil, false
	}
	return o.PaginationToken, true
}

// HasPaginationToken returns a boolean if a field has been set.
func (o *ListGroups) HasPaginationToken() bool {
	if o != nil && o.PaginationToken != nil {
		return true
	}

	return false
}

// SetPaginationToken gets a reference to the given string and assigns it to the PaginationToken field.
func (o *ListGroups) SetPaginationToken(v string) {
	o.PaginationToken = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ListGroups) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGroups) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ListGroups) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ListGroups) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *ListGroups) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGroups) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *ListGroups) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *ListGroups) SetUidToken(v string) {
	o.UidToken = &v
}

func (o ListGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.PaginationToken != nil {
		toSerialize["pagination-token"] = o.PaginationToken
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableListGroups struct {
	value *ListGroups
	isSet bool
}

func (v NullableListGroups) Get() *ListGroups {
	return v.value
}

func (v *NullableListGroups) Set(val *ListGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableListGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableListGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGroups(val *ListGroups) *NullableListGroups {
	return &NullableListGroups{value: val, isSet: true}
}

func (v NullableListGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


