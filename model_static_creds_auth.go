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

// StaticCredsAuth staticCredsAuth is a command that creates a temporary access profile using the provided static credentials.
type StaticCredsAuth struct {
	// Akeyless JWT token
	AccessId *string `json:"access-id,omitempty"`
	// Akeyless JWT token
	AdminEmail *string `json:"admin-email,omitempty"`
	// Akeyless JWT token
	Creds *string `json:"creds,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
}

// NewStaticCredsAuth instantiates a new StaticCredsAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticCredsAuth() *StaticCredsAuth {
	this := StaticCredsAuth{}
	var json bool = false
	this.Json = &json
	return &this
}

// NewStaticCredsAuthWithDefaults instantiates a new StaticCredsAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticCredsAuthWithDefaults() *StaticCredsAuth {
	this := StaticCredsAuth{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *StaticCredsAuth) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCredsAuth) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *StaticCredsAuth) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *StaticCredsAuth) SetAccessId(v string) {
	o.AccessId = &v
}

// GetAdminEmail returns the AdminEmail field value if set, zero value otherwise.
func (o *StaticCredsAuth) GetAdminEmail() string {
	if o == nil || o.AdminEmail == nil {
		var ret string
		return ret
	}
	return *o.AdminEmail
}

// GetAdminEmailOk returns a tuple with the AdminEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCredsAuth) GetAdminEmailOk() (*string, bool) {
	if o == nil || o.AdminEmail == nil {
		return nil, false
	}
	return o.AdminEmail, true
}

// HasAdminEmail returns a boolean if a field has been set.
func (o *StaticCredsAuth) HasAdminEmail() bool {
	if o != nil && o.AdminEmail != nil {
		return true
	}

	return false
}

// SetAdminEmail gets a reference to the given string and assigns it to the AdminEmail field.
func (o *StaticCredsAuth) SetAdminEmail(v string) {
	o.AdminEmail = &v
}

// GetCreds returns the Creds field value if set, zero value otherwise.
func (o *StaticCredsAuth) GetCreds() string {
	if o == nil || o.Creds == nil {
		var ret string
		return ret
	}
	return *o.Creds
}

// GetCredsOk returns a tuple with the Creds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCredsAuth) GetCredsOk() (*string, bool) {
	if o == nil || o.Creds == nil {
		return nil, false
	}
	return o.Creds, true
}

// HasCreds returns a boolean if a field has been set.
func (o *StaticCredsAuth) HasCreds() bool {
	if o != nil && o.Creds != nil {
		return true
	}

	return false
}

// SetCreds gets a reference to the given string and assigns it to the Creds field.
func (o *StaticCredsAuth) SetCreds(v string) {
	o.Creds = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *StaticCredsAuth) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCredsAuth) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *StaticCredsAuth) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *StaticCredsAuth) SetJson(v bool) {
	o.Json = &v
}

func (o StaticCredsAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access-id"] = o.AccessId
	}
	if o.AdminEmail != nil {
		toSerialize["admin-email"] = o.AdminEmail
	}
	if o.Creds != nil {
		toSerialize["creds"] = o.Creds
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	return json.Marshal(toSerialize)
}

type NullableStaticCredsAuth struct {
	value *StaticCredsAuth
	isSet bool
}

func (v NullableStaticCredsAuth) Get() *StaticCredsAuth {
	return v.value
}

func (v *NullableStaticCredsAuth) Set(val *StaticCredsAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticCredsAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticCredsAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticCredsAuth(val *StaticCredsAuth) *NullableStaticCredsAuth {
	return &NullableStaticCredsAuth{value: val, isSet: true}
}

func (v NullableStaticCredsAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticCredsAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


