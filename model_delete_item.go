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

// DeleteItem struct for DeleteItem
type DeleteItem struct {
	// When delete-in-days=-1, must be set
	DeleteImmediately *bool `json:"delete-immediately,omitempty"`
	// The number of days to wait before deleting the item (relevant for keys only)
	DeleteInDays *int64 `json:"delete-in-days,omitempty"`
	// Item name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The specific version you want to delete - 0=last version, -1=entire item with all versions (default)
	Version *int32 `json:"version,omitempty"`
}

// NewDeleteItem instantiates a new DeleteItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteItem(name string, ) *DeleteItem {
	this := DeleteItem{}
	var deleteImmediately bool = false
	this.DeleteImmediately = &deleteImmediately
	var deleteInDays int64 = 7
	this.DeleteInDays = &deleteInDays
	this.Name = name
	var version int32 = -1
	this.Version = &version
	return &this
}

// NewDeleteItemWithDefaults instantiates a new DeleteItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteItemWithDefaults() *DeleteItem {
	this := DeleteItem{}
	var deleteImmediately bool = false
	this.DeleteImmediately = &deleteImmediately
	var deleteInDays int64 = 7
	this.DeleteInDays = &deleteInDays
	var version int32 = -1
	this.Version = &version
	return &this
}

// GetDeleteImmediately returns the DeleteImmediately field value if set, zero value otherwise.
func (o *DeleteItem) GetDeleteImmediately() bool {
	if o == nil || o.DeleteImmediately == nil {
		var ret bool
		return ret
	}
	return *o.DeleteImmediately
}

// GetDeleteImmediatelyOk returns a tuple with the DeleteImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteItem) GetDeleteImmediatelyOk() (*bool, bool) {
	if o == nil || o.DeleteImmediately == nil {
		return nil, false
	}
	return o.DeleteImmediately, true
}

// HasDeleteImmediately returns a boolean if a field has been set.
func (o *DeleteItem) HasDeleteImmediately() bool {
	if o != nil && o.DeleteImmediately != nil {
		return true
	}

	return false
}

// SetDeleteImmediately gets a reference to the given bool and assigns it to the DeleteImmediately field.
func (o *DeleteItem) SetDeleteImmediately(v bool) {
	o.DeleteImmediately = &v
}

// GetDeleteInDays returns the DeleteInDays field value if set, zero value otherwise.
func (o *DeleteItem) GetDeleteInDays() int64 {
	if o == nil || o.DeleteInDays == nil {
		var ret int64
		return ret
	}
	return *o.DeleteInDays
}

// GetDeleteInDaysOk returns a tuple with the DeleteInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteItem) GetDeleteInDaysOk() (*int64, bool) {
	if o == nil || o.DeleteInDays == nil {
		return nil, false
	}
	return o.DeleteInDays, true
}

// HasDeleteInDays returns a boolean if a field has been set.
func (o *DeleteItem) HasDeleteInDays() bool {
	if o != nil && o.DeleteInDays != nil {
		return true
	}

	return false
}

// SetDeleteInDays gets a reference to the given int64 and assigns it to the DeleteInDays field.
func (o *DeleteItem) SetDeleteInDays(v int64) {
	o.DeleteInDays = &v
}

// GetName returns the Name field value
func (o *DeleteItem) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeleteItem) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeleteItem) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DeleteItem) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteItem) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DeleteItem) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DeleteItem) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DeleteItem) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteItem) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DeleteItem) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DeleteItem) SetUidToken(v string) {
	o.UidToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeleteItem) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteItem) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeleteItem) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DeleteItem) SetVersion(v int32) {
	o.Version = &v
}

func (o DeleteItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteImmediately != nil {
		toSerialize["delete-immediately"] = o.DeleteImmediately
	}
	if o.DeleteInDays != nil {
		toSerialize["delete-in-days"] = o.DeleteInDays
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
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteItem struct {
	value *DeleteItem
	isSet bool
}

func (v NullableDeleteItem) Get() *DeleteItem {
	return v.value
}

func (v *NullableDeleteItem) Set(val *DeleteItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteItem(val *DeleteItem) *NullableDeleteItem {
	return &NullableDeleteItem{value: val, isSet: true}
}

func (v NullableDeleteItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


