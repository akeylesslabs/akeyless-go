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

// ListTargets struct for ListTargets
type ListTargets struct {
	// Filter by auth method name or part of it
	Filter *string `json:"filter,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Next page reference
	PaginationToken *string `json:"pagination-token,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The target types list . In case it is empty, all types of targets will be returned. options: [hanadb cassandra aws ssh gke eks mysql mongodb snowflake mssql redshift artifactory azure rabbitmq k8s venafi gcp oracle dockerhub ldap github chef web salesforce postgres]
	Type *[]string `json:"type,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewListTargets instantiates a new ListTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTargets() *ListTargets {
	this := ListTargets{}
	var json bool = false
	this.Json = &json
	return &this
}

// NewListTargetsWithDefaults instantiates a new ListTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTargetsWithDefaults() *ListTargets {
	this := ListTargets{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ListTargets) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTargets) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ListTargets) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *ListTargets) SetFilter(v string) {
	o.Filter = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ListTargets) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTargets) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ListTargets) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *ListTargets) SetJson(v bool) {
	o.Json = &v
}

// GetPaginationToken returns the PaginationToken field value if set, zero value otherwise.
func (o *ListTargets) GetPaginationToken() string {
	if o == nil || o.PaginationToken == nil {
		var ret string
		return ret
	}
	return *o.PaginationToken
}

// GetPaginationTokenOk returns a tuple with the PaginationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTargets) GetPaginationTokenOk() (*string, bool) {
	if o == nil || o.PaginationToken == nil {
		return nil, false
	}
	return o.PaginationToken, true
}

// HasPaginationToken returns a boolean if a field has been set.
func (o *ListTargets) HasPaginationToken() bool {
	if o != nil && o.PaginationToken != nil {
		return true
	}

	return false
}

// SetPaginationToken gets a reference to the given string and assigns it to the PaginationToken field.
func (o *ListTargets) SetPaginationToken(v string) {
	o.PaginationToken = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ListTargets) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTargets) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ListTargets) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ListTargets) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListTargets) GetType() []string {
	if o == nil || o.Type == nil {
		var ret []string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTargets) GetTypeOk() (*[]string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListTargets) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given []string and assigns it to the Type field.
func (o *ListTargets) SetType(v []string) {
	o.Type = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *ListTargets) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTargets) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *ListTargets) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *ListTargets) SetUidToken(v string) {
	o.UidToken = &v
}

func (o ListTargets) MarshalJSON() ([]byte, error) {
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableListTargets struct {
	value *ListTargets
	isSet bool
}

func (v NullableListTargets) Get() *ListTargets {
	return v.value
}

func (v *NullableListTargets) Set(val *ListTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableListTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableListTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTargets(val *ListTargets) *NullableListTargets {
	return &NullableListTargets{value: val, isSet: true}
}

func (v NullableListTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


