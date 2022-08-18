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

// CreateAuthMethodLDAPOutput struct for CreateAuthMethodLDAPOutput
type CreateAuthMethodLDAPOutput struct {
	AccessId *string `json:"access_id,omitempty"`
	PrvKey *string `json:"prv_key,omitempty"`
}

// NewCreateAuthMethodLDAPOutput instantiates a new CreateAuthMethodLDAPOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodLDAPOutput() *CreateAuthMethodLDAPOutput {
	this := CreateAuthMethodLDAPOutput{}
	return &this
}

// NewCreateAuthMethodLDAPOutputWithDefaults instantiates a new CreateAuthMethodLDAPOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodLDAPOutputWithDefaults() *CreateAuthMethodLDAPOutput {
	this := CreateAuthMethodLDAPOutput{}
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAPOutput) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAPOutput) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAPOutput) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *CreateAuthMethodLDAPOutput) SetAccessId(v string) {
	o.AccessId = &v
}

// GetPrvKey returns the PrvKey field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAPOutput) GetPrvKey() string {
	if o == nil || o.PrvKey == nil {
		var ret string
		return ret
	}
	return *o.PrvKey
}

// GetPrvKeyOk returns a tuple with the PrvKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAPOutput) GetPrvKeyOk() (*string, bool) {
	if o == nil || o.PrvKey == nil {
		return nil, false
	}
	return o.PrvKey, true
}

// HasPrvKey returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAPOutput) HasPrvKey() bool {
	if o != nil && o.PrvKey != nil {
		return true
	}

	return false
}

// SetPrvKey gets a reference to the given string and assigns it to the PrvKey field.
func (o *CreateAuthMethodLDAPOutput) SetPrvKey(v string) {
	o.PrvKey = &v
}

func (o CreateAuthMethodLDAPOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access_id"] = o.AccessId
	}
	if o.PrvKey != nil {
		toSerialize["prv_key"] = o.PrvKey
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAuthMethodLDAPOutput struct {
	value *CreateAuthMethodLDAPOutput
	isSet bool
}

func (v NullableCreateAuthMethodLDAPOutput) Get() *CreateAuthMethodLDAPOutput {
	return v.value
}

func (v *NullableCreateAuthMethodLDAPOutput) Set(val *CreateAuthMethodLDAPOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodLDAPOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodLDAPOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodLDAPOutput(val *CreateAuthMethodLDAPOutput) *NullableCreateAuthMethodLDAPOutput {
	return &NullableCreateAuthMethodLDAPOutput{value: val, isSet: true}
}

func (v NullableCreateAuthMethodLDAPOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodLDAPOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


