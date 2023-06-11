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

// SystemAccessCredsSettings SystemAccessCredsSettings describes system access credential settings for account by minutes
type SystemAccessCredsSettings struct {
	JwtTtlDefault *int64 `json:"jwt_ttl_default,omitempty"`
	JwtTtlMaximum *int64 `json:"jwt_ttl_maximum,omitempty"`
	JwtTtlMinimum *int64 `json:"jwt_ttl_minimum,omitempty"`
}

// NewSystemAccessCredsSettings instantiates a new SystemAccessCredsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemAccessCredsSettings() *SystemAccessCredsSettings {
	this := SystemAccessCredsSettings{}
	return &this
}

// NewSystemAccessCredsSettingsWithDefaults instantiates a new SystemAccessCredsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemAccessCredsSettingsWithDefaults() *SystemAccessCredsSettings {
	this := SystemAccessCredsSettings{}
	return &this
}

// GetJwtTtlDefault returns the JwtTtlDefault field value if set, zero value otherwise.
func (o *SystemAccessCredsSettings) GetJwtTtlDefault() int64 {
	if o == nil || o.JwtTtlDefault == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtlDefault
}

// GetJwtTtlDefaultOk returns a tuple with the JwtTtlDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAccessCredsSettings) GetJwtTtlDefaultOk() (*int64, bool) {
	if o == nil || o.JwtTtlDefault == nil {
		return nil, false
	}
	return o.JwtTtlDefault, true
}

// HasJwtTtlDefault returns a boolean if a field has been set.
func (o *SystemAccessCredsSettings) HasJwtTtlDefault() bool {
	if o != nil && o.JwtTtlDefault != nil {
		return true
	}

	return false
}

// SetJwtTtlDefault gets a reference to the given int64 and assigns it to the JwtTtlDefault field.
func (o *SystemAccessCredsSettings) SetJwtTtlDefault(v int64) {
	o.JwtTtlDefault = &v
}

// GetJwtTtlMaximum returns the JwtTtlMaximum field value if set, zero value otherwise.
func (o *SystemAccessCredsSettings) GetJwtTtlMaximum() int64 {
	if o == nil || o.JwtTtlMaximum == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtlMaximum
}

// GetJwtTtlMaximumOk returns a tuple with the JwtTtlMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAccessCredsSettings) GetJwtTtlMaximumOk() (*int64, bool) {
	if o == nil || o.JwtTtlMaximum == nil {
		return nil, false
	}
	return o.JwtTtlMaximum, true
}

// HasJwtTtlMaximum returns a boolean if a field has been set.
func (o *SystemAccessCredsSettings) HasJwtTtlMaximum() bool {
	if o != nil && o.JwtTtlMaximum != nil {
		return true
	}

	return false
}

// SetJwtTtlMaximum gets a reference to the given int64 and assigns it to the JwtTtlMaximum field.
func (o *SystemAccessCredsSettings) SetJwtTtlMaximum(v int64) {
	o.JwtTtlMaximum = &v
}

// GetJwtTtlMinimum returns the JwtTtlMinimum field value if set, zero value otherwise.
func (o *SystemAccessCredsSettings) GetJwtTtlMinimum() int64 {
	if o == nil || o.JwtTtlMinimum == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtlMinimum
}

// GetJwtTtlMinimumOk returns a tuple with the JwtTtlMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAccessCredsSettings) GetJwtTtlMinimumOk() (*int64, bool) {
	if o == nil || o.JwtTtlMinimum == nil {
		return nil, false
	}
	return o.JwtTtlMinimum, true
}

// HasJwtTtlMinimum returns a boolean if a field has been set.
func (o *SystemAccessCredsSettings) HasJwtTtlMinimum() bool {
	if o != nil && o.JwtTtlMinimum != nil {
		return true
	}

	return false
}

// SetJwtTtlMinimum gets a reference to the given int64 and assigns it to the JwtTtlMinimum field.
func (o *SystemAccessCredsSettings) SetJwtTtlMinimum(v int64) {
	o.JwtTtlMinimum = &v
}

func (o SystemAccessCredsSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JwtTtlDefault != nil {
		toSerialize["jwt_ttl_default"] = o.JwtTtlDefault
	}
	if o.JwtTtlMaximum != nil {
		toSerialize["jwt_ttl_maximum"] = o.JwtTtlMaximum
	}
	if o.JwtTtlMinimum != nil {
		toSerialize["jwt_ttl_minimum"] = o.JwtTtlMinimum
	}
	return json.Marshal(toSerialize)
}

type NullableSystemAccessCredsSettings struct {
	value *SystemAccessCredsSettings
	isSet bool
}

func (v NullableSystemAccessCredsSettings) Get() *SystemAccessCredsSettings {
	return v.value
}

func (v *NullableSystemAccessCredsSettings) Set(val *SystemAccessCredsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemAccessCredsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemAccessCredsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemAccessCredsSettings(val *SystemAccessCredsSettings) *NullableSystemAccessCredsSettings {
	return &NullableSystemAccessCredsSettings{value: val, isSet: true}
}

func (v NullableSystemAccessCredsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemAccessCredsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


