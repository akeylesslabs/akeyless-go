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

// checks if the EmailEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailEntry{}

// EmailEntry struct for EmailEntry
type EmailEntry struct {
	ToEmail *string `json:"to_email,omitempty"`
	ToName *string `json:"to_name,omitempty"`
}

// NewEmailEntry instantiates a new EmailEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailEntry() *EmailEntry {
	this := EmailEntry{}
	return &this
}

// NewEmailEntryWithDefaults instantiates a new EmailEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailEntryWithDefaults() *EmailEntry {
	this := EmailEntry{}
	return &this
}

// GetToEmail returns the ToEmail field value if set, zero value otherwise.
func (o *EmailEntry) GetToEmail() string {
	if o == nil || IsNil(o.ToEmail) {
		var ret string
		return ret
	}
	return *o.ToEmail
}

// GetToEmailOk returns a tuple with the ToEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailEntry) GetToEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ToEmail) {
		return nil, false
	}
	return o.ToEmail, true
}

// HasToEmail returns a boolean if a field has been set.
func (o *EmailEntry) HasToEmail() bool {
	if o != nil && !IsNil(o.ToEmail) {
		return true
	}

	return false
}

// SetToEmail gets a reference to the given string and assigns it to the ToEmail field.
func (o *EmailEntry) SetToEmail(v string) {
	o.ToEmail = &v
}

// GetToName returns the ToName field value if set, zero value otherwise.
func (o *EmailEntry) GetToName() string {
	if o == nil || IsNil(o.ToName) {
		var ret string
		return ret
	}
	return *o.ToName
}

// GetToNameOk returns a tuple with the ToName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailEntry) GetToNameOk() (*string, bool) {
	if o == nil || IsNil(o.ToName) {
		return nil, false
	}
	return o.ToName, true
}

// HasToName returns a boolean if a field has been set.
func (o *EmailEntry) HasToName() bool {
	if o != nil && !IsNil(o.ToName) {
		return true
	}

	return false
}

// SetToName gets a reference to the given string and assigns it to the ToName field.
func (o *EmailEntry) SetToName(v string) {
	o.ToName = &v
}

func (o EmailEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ToEmail) {
		toSerialize["to_email"] = o.ToEmail
	}
	if !IsNil(o.ToName) {
		toSerialize["to_name"] = o.ToName
	}
	return toSerialize, nil
}

type NullableEmailEntry struct {
	value *EmailEntry
	isSet bool
}

func (v NullableEmailEntry) Get() *EmailEntry {
	return v.value
}

func (v *NullableEmailEntry) Set(val *EmailEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailEntry(val *EmailEntry) *NullableEmailEntry {
	return &NullableEmailEntry{value: val, isSet: true}
}

func (v NullableEmailEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


