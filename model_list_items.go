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

// ListItems struct for ListItems
type ListItems struct {
	// for personal password manager
	Accessibility *string `json:"accessibility,omitempty"`
	// Filter by item name or part of it
	Filter *string `json:"filter,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	MinimalView *bool `json:"minimal-view,omitempty"`
	// Next page reference
	PaginationToken *string `json:"pagination-token,omitempty"`
	// Path to folder
	Path *string `json:"path,omitempty"`
	SubTypes *[]string `json:"sub_types,omitempty"`
	// Filter by item tag
	Tag *string `json:"tag,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The item types list of the requested items. In case it is empty, all types of items will be returned. options: [key, static-secret, dynamic-secret]
	Type *[]string `json:"type,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewListItems instantiates a new ListItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListItems() *ListItems {
	this := ListItems{}
	return &this
}

// NewListItemsWithDefaults instantiates a new ListItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListItemsWithDefaults() *ListItems {
	this := ListItems{}
	return &this
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *ListItems) GetAccessibility() string {
	if o == nil || o.Accessibility == nil {
		var ret string
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetAccessibilityOk() (*string, bool) {
	if o == nil || o.Accessibility == nil {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *ListItems) HasAccessibility() bool {
	if o != nil && o.Accessibility != nil {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given string and assigns it to the Accessibility field.
func (o *ListItems) SetAccessibility(v string) {
	o.Accessibility = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ListItems) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ListItems) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *ListItems) SetFilter(v string) {
	o.Filter = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *ListItems) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *ListItems) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *ListItems) SetJson(v bool) {
	o.Json = &v
}

// GetMinimalView returns the MinimalView field value if set, zero value otherwise.
func (o *ListItems) GetMinimalView() bool {
	if o == nil || o.MinimalView == nil {
		var ret bool
		return ret
	}
	return *o.MinimalView
}

// GetMinimalViewOk returns a tuple with the MinimalView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetMinimalViewOk() (*bool, bool) {
	if o == nil || o.MinimalView == nil {
		return nil, false
	}
	return o.MinimalView, true
}

// HasMinimalView returns a boolean if a field has been set.
func (o *ListItems) HasMinimalView() bool {
	if o != nil && o.MinimalView != nil {
		return true
	}

	return false
}

// SetMinimalView gets a reference to the given bool and assigns it to the MinimalView field.
func (o *ListItems) SetMinimalView(v bool) {
	o.MinimalView = &v
}

// GetPaginationToken returns the PaginationToken field value if set, zero value otherwise.
func (o *ListItems) GetPaginationToken() string {
	if o == nil || o.PaginationToken == nil {
		var ret string
		return ret
	}
	return *o.PaginationToken
}

// GetPaginationTokenOk returns a tuple with the PaginationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetPaginationTokenOk() (*string, bool) {
	if o == nil || o.PaginationToken == nil {
		return nil, false
	}
	return o.PaginationToken, true
}

// HasPaginationToken returns a boolean if a field has been set.
func (o *ListItems) HasPaginationToken() bool {
	if o != nil && o.PaginationToken != nil {
		return true
	}

	return false
}

// SetPaginationToken gets a reference to the given string and assigns it to the PaginationToken field.
func (o *ListItems) SetPaginationToken(v string) {
	o.PaginationToken = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ListItems) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ListItems) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ListItems) SetPath(v string) {
	o.Path = &v
}

// GetSubTypes returns the SubTypes field value if set, zero value otherwise.
func (o *ListItems) GetSubTypes() []string {
	if o == nil || o.SubTypes == nil {
		var ret []string
		return ret
	}
	return *o.SubTypes
}

// GetSubTypesOk returns a tuple with the SubTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetSubTypesOk() (*[]string, bool) {
	if o == nil || o.SubTypes == nil {
		return nil, false
	}
	return o.SubTypes, true
}

// HasSubTypes returns a boolean if a field has been set.
func (o *ListItems) HasSubTypes() bool {
	if o != nil && o.SubTypes != nil {
		return true
	}

	return false
}

// SetSubTypes gets a reference to the given []string and assigns it to the SubTypes field.
func (o *ListItems) SetSubTypes(v []string) {
	o.SubTypes = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ListItems) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ListItems) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ListItems) SetTag(v string) {
	o.Tag = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ListItems) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ListItems) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ListItems) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListItems) GetType() []string {
	if o == nil || o.Type == nil {
		var ret []string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetTypeOk() (*[]string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListItems) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given []string and assigns it to the Type field.
func (o *ListItems) SetType(v []string) {
	o.Type = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *ListItems) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItems) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *ListItems) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *ListItems) SetUidToken(v string) {
	o.UidToken = &v
}

func (o ListItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessibility != nil {
		toSerialize["accessibility"] = o.Accessibility
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.MinimalView != nil {
		toSerialize["minimal-view"] = o.MinimalView
	}
	if o.PaginationToken != nil {
		toSerialize["pagination-token"] = o.PaginationToken
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.SubTypes != nil {
		toSerialize["sub_types"] = o.SubTypes
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
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

type NullableListItems struct {
	value *ListItems
	isSet bool
}

func (v NullableListItems) Get() *ListItems {
	return v.value
}

func (v *NullableListItems) Set(val *ListItems) {
	v.value = val
	v.isSet = true
}

func (v NullableListItems) IsSet() bool {
	return v.isSet
}

func (v *NullableListItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListItems(val *ListItems) *NullableListItems {
	return &NullableListItems{value: val, isSet: true}
}

func (v NullableListItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


