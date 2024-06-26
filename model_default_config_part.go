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

// DefaultConfigPart struct for DefaultConfigPart
type DefaultConfigPart struct {
	CertificateAccessId *string `json:"certificate_access_id,omitempty"`
	DefaultProtectionKeyId *int64 `json:"default_protection_key_id,omitempty"`
	DefaultSecretLocation *string `json:"default_secret_location,omitempty"`
	OidcAccessId *string `json:"oidc_access_id,omitempty"`
	SamlAccessId *string `json:"saml_access_id,omitempty"`
}

// NewDefaultConfigPart instantiates a new DefaultConfigPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultConfigPart() *DefaultConfigPart {
	this := DefaultConfigPart{}
	return &this
}

// NewDefaultConfigPartWithDefaults instantiates a new DefaultConfigPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultConfigPartWithDefaults() *DefaultConfigPart {
	this := DefaultConfigPart{}
	return &this
}

// GetCertificateAccessId returns the CertificateAccessId field value if set, zero value otherwise.
func (o *DefaultConfigPart) GetCertificateAccessId() string {
	if o == nil || o.CertificateAccessId == nil {
		var ret string
		return ret
	}
	return *o.CertificateAccessId
}

// GetCertificateAccessIdOk returns a tuple with the CertificateAccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultConfigPart) GetCertificateAccessIdOk() (*string, bool) {
	if o == nil || o.CertificateAccessId == nil {
		return nil, false
	}
	return o.CertificateAccessId, true
}

// HasCertificateAccessId returns a boolean if a field has been set.
func (o *DefaultConfigPart) HasCertificateAccessId() bool {
	if o != nil && o.CertificateAccessId != nil {
		return true
	}

	return false
}

// SetCertificateAccessId gets a reference to the given string and assigns it to the CertificateAccessId field.
func (o *DefaultConfigPart) SetCertificateAccessId(v string) {
	o.CertificateAccessId = &v
}

// GetDefaultProtectionKeyId returns the DefaultProtectionKeyId field value if set, zero value otherwise.
func (o *DefaultConfigPart) GetDefaultProtectionKeyId() int64 {
	if o == nil || o.DefaultProtectionKeyId == nil {
		var ret int64
		return ret
	}
	return *o.DefaultProtectionKeyId
}

// GetDefaultProtectionKeyIdOk returns a tuple with the DefaultProtectionKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultConfigPart) GetDefaultProtectionKeyIdOk() (*int64, bool) {
	if o == nil || o.DefaultProtectionKeyId == nil {
		return nil, false
	}
	return o.DefaultProtectionKeyId, true
}

// HasDefaultProtectionKeyId returns a boolean if a field has been set.
func (o *DefaultConfigPart) HasDefaultProtectionKeyId() bool {
	if o != nil && o.DefaultProtectionKeyId != nil {
		return true
	}

	return false
}

// SetDefaultProtectionKeyId gets a reference to the given int64 and assigns it to the DefaultProtectionKeyId field.
func (o *DefaultConfigPart) SetDefaultProtectionKeyId(v int64) {
	o.DefaultProtectionKeyId = &v
}

// GetDefaultSecretLocation returns the DefaultSecretLocation field value if set, zero value otherwise.
func (o *DefaultConfigPart) GetDefaultSecretLocation() string {
	if o == nil || o.DefaultSecretLocation == nil {
		var ret string
		return ret
	}
	return *o.DefaultSecretLocation
}

// GetDefaultSecretLocationOk returns a tuple with the DefaultSecretLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultConfigPart) GetDefaultSecretLocationOk() (*string, bool) {
	if o == nil || o.DefaultSecretLocation == nil {
		return nil, false
	}
	return o.DefaultSecretLocation, true
}

// HasDefaultSecretLocation returns a boolean if a field has been set.
func (o *DefaultConfigPart) HasDefaultSecretLocation() bool {
	if o != nil && o.DefaultSecretLocation != nil {
		return true
	}

	return false
}

// SetDefaultSecretLocation gets a reference to the given string and assigns it to the DefaultSecretLocation field.
func (o *DefaultConfigPart) SetDefaultSecretLocation(v string) {
	o.DefaultSecretLocation = &v
}

// GetOidcAccessId returns the OidcAccessId field value if set, zero value otherwise.
func (o *DefaultConfigPart) GetOidcAccessId() string {
	if o == nil || o.OidcAccessId == nil {
		var ret string
		return ret
	}
	return *o.OidcAccessId
}

// GetOidcAccessIdOk returns a tuple with the OidcAccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultConfigPart) GetOidcAccessIdOk() (*string, bool) {
	if o == nil || o.OidcAccessId == nil {
		return nil, false
	}
	return o.OidcAccessId, true
}

// HasOidcAccessId returns a boolean if a field has been set.
func (o *DefaultConfigPart) HasOidcAccessId() bool {
	if o != nil && o.OidcAccessId != nil {
		return true
	}

	return false
}

// SetOidcAccessId gets a reference to the given string and assigns it to the OidcAccessId field.
func (o *DefaultConfigPart) SetOidcAccessId(v string) {
	o.OidcAccessId = &v
}

// GetSamlAccessId returns the SamlAccessId field value if set, zero value otherwise.
func (o *DefaultConfigPart) GetSamlAccessId() string {
	if o == nil || o.SamlAccessId == nil {
		var ret string
		return ret
	}
	return *o.SamlAccessId
}

// GetSamlAccessIdOk returns a tuple with the SamlAccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultConfigPart) GetSamlAccessIdOk() (*string, bool) {
	if o == nil || o.SamlAccessId == nil {
		return nil, false
	}
	return o.SamlAccessId, true
}

// HasSamlAccessId returns a boolean if a field has been set.
func (o *DefaultConfigPart) HasSamlAccessId() bool {
	if o != nil && o.SamlAccessId != nil {
		return true
	}

	return false
}

// SetSamlAccessId gets a reference to the given string and assigns it to the SamlAccessId field.
func (o *DefaultConfigPart) SetSamlAccessId(v string) {
	o.SamlAccessId = &v
}

func (o DefaultConfigPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertificateAccessId != nil {
		toSerialize["certificate_access_id"] = o.CertificateAccessId
	}
	if o.DefaultProtectionKeyId != nil {
		toSerialize["default_protection_key_id"] = o.DefaultProtectionKeyId
	}
	if o.DefaultSecretLocation != nil {
		toSerialize["default_secret_location"] = o.DefaultSecretLocation
	}
	if o.OidcAccessId != nil {
		toSerialize["oidc_access_id"] = o.OidcAccessId
	}
	if o.SamlAccessId != nil {
		toSerialize["saml_access_id"] = o.SamlAccessId
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultConfigPart struct {
	value *DefaultConfigPart
	isSet bool
}

func (v NullableDefaultConfigPart) Get() *DefaultConfigPart {
	return v.value
}

func (v *NullableDefaultConfigPart) Set(val *DefaultConfigPart) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultConfigPart) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultConfigPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultConfigPart(val *DefaultConfigPart) *NullableDefaultConfigPart {
	return &NullableDefaultConfigPart{value: val, isSet: true}
}

func (v NullableDefaultConfigPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultConfigPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


