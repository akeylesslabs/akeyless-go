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

// StaticSecretDetailsInfo struct for StaticSecretDetailsInfo
type StaticSecretDetailsInfo struct {
	Username *string `json:"username,omitempty"`
	// deprecated
	Website *string `json:"website,omitempty"`
	Websites []string `json:"websites,omitempty"`
}

// NewStaticSecretDetailsInfo instantiates a new StaticSecretDetailsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticSecretDetailsInfo() *StaticSecretDetailsInfo {
	this := StaticSecretDetailsInfo{}
	return &this
}

// NewStaticSecretDetailsInfoWithDefaults instantiates a new StaticSecretDetailsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticSecretDetailsInfoWithDefaults() *StaticSecretDetailsInfo {
	this := StaticSecretDetailsInfo{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *StaticSecretDetailsInfo) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSecretDetailsInfo) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *StaticSecretDetailsInfo) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *StaticSecretDetailsInfo) SetUsername(v string) {
	o.Username = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *StaticSecretDetailsInfo) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSecretDetailsInfo) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *StaticSecretDetailsInfo) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *StaticSecretDetailsInfo) SetWebsite(v string) {
	o.Website = &v
}

// GetWebsites returns the Websites field value if set, zero value otherwise.
func (o *StaticSecretDetailsInfo) GetWebsites() []string {
	if o == nil || o.Websites == nil {
		var ret []string
		return ret
	}
	return o.Websites
}

// GetWebsitesOk returns a tuple with the Websites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSecretDetailsInfo) GetWebsitesOk() ([]string, bool) {
	if o == nil || o.Websites == nil {
		return nil, false
	}
	return o.Websites, true
}

// HasWebsites returns a boolean if a field has been set.
func (o *StaticSecretDetailsInfo) HasWebsites() bool {
	if o != nil && o.Websites != nil {
		return true
	}

	return false
}

// SetWebsites gets a reference to the given []string and assigns it to the Websites field.
func (o *StaticSecretDetailsInfo) SetWebsites(v []string) {
	o.Websites = v
}

func (o StaticSecretDetailsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.Websites != nil {
		toSerialize["websites"] = o.Websites
	}
	return json.Marshal(toSerialize)
}

type NullableStaticSecretDetailsInfo struct {
	value *StaticSecretDetailsInfo
	isSet bool
}

func (v NullableStaticSecretDetailsInfo) Get() *StaticSecretDetailsInfo {
	return v.value
}

func (v *NullableStaticSecretDetailsInfo) Set(val *StaticSecretDetailsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticSecretDetailsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticSecretDetailsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticSecretDetailsInfo(val *StaticSecretDetailsInfo) *NullableStaticSecretDetailsInfo {
	return &NullableStaticSecretDetailsInfo{value: val, isSet: true}
}

func (v NullableStaticSecretDetailsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticSecretDetailsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


