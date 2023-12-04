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

// CertificateChainInfo struct for CertificateChainInfo
type CertificateChainInfo struct {
	CertificateChain *[]CertificateInfo `json:"certificate_chain,omitempty"`
	CertificateFormat *string `json:"certificate_format,omitempty"`
	CertificateIssuerItemId *int64 `json:"certificate_issuer_item_id,omitempty"`
	CertificateIssuerName *string `json:"certificate_issuer_name,omitempty"`
	CertificatePem *string `json:"certificate_pem,omitempty"`
	CertificateStatus *string `json:"certificate_status,omitempty"`
	ExpirationEvents *[]CertificateExpirationEvent `json:"expiration_events,omitempty"`
}

// NewCertificateChainInfo instantiates a new CertificateChainInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateChainInfo() *CertificateChainInfo {
	this := CertificateChainInfo{}
	return &this
}

// NewCertificateChainInfoWithDefaults instantiates a new CertificateChainInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateChainInfoWithDefaults() *CertificateChainInfo {
	this := CertificateChainInfo{}
	return &this
}

// GetCertificateChain returns the CertificateChain field value if set, zero value otherwise.
func (o *CertificateChainInfo) GetCertificateChain() []CertificateInfo {
	if o == nil || o.CertificateChain == nil {
		var ret []CertificateInfo
		return ret
	}
	return *o.CertificateChain
}

// GetCertificateChainOk returns a tuple with the CertificateChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateChainInfo) GetCertificateChainOk() (*[]CertificateInfo, bool) {
	if o == nil || o.CertificateChain == nil {
		return nil, false
	}
	return o.CertificateChain, true
}

// HasCertificateChain returns a boolean if a field has been set.
func (o *CertificateChainInfo) HasCertificateChain() bool {
	if o != nil && o.CertificateChain != nil {
		return true
	}

	return false
}

// SetCertificateChain gets a reference to the given []CertificateInfo and assigns it to the CertificateChain field.
func (o *CertificateChainInfo) SetCertificateChain(v []CertificateInfo) {
	o.CertificateChain = &v
}

// GetCertificateFormat returns the CertificateFormat field value if set, zero value otherwise.
func (o *CertificateChainInfo) GetCertificateFormat() string {
	if o == nil || o.CertificateFormat == nil {
		var ret string
		return ret
	}
	return *o.CertificateFormat
}

// GetCertificateFormatOk returns a tuple with the CertificateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateChainInfo) GetCertificateFormatOk() (*string, bool) {
	if o == nil || o.CertificateFormat == nil {
		return nil, false
	}
	return o.CertificateFormat, true
}

// HasCertificateFormat returns a boolean if a field has been set.
func (o *CertificateChainInfo) HasCertificateFormat() bool {
	if o != nil && o.CertificateFormat != nil {
		return true
	}

	return false
}

// SetCertificateFormat gets a reference to the given string and assigns it to the CertificateFormat field.
func (o *CertificateChainInfo) SetCertificateFormat(v string) {
	o.CertificateFormat = &v
}

// GetCertificateIssuerItemId returns the CertificateIssuerItemId field value if set, zero value otherwise.
func (o *CertificateChainInfo) GetCertificateIssuerItemId() int64 {
	if o == nil || o.CertificateIssuerItemId == nil {
		var ret int64
		return ret
	}
	return *o.CertificateIssuerItemId
}

// GetCertificateIssuerItemIdOk returns a tuple with the CertificateIssuerItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateChainInfo) GetCertificateIssuerItemIdOk() (*int64, bool) {
	if o == nil || o.CertificateIssuerItemId == nil {
		return nil, false
	}
	return o.CertificateIssuerItemId, true
}

// HasCertificateIssuerItemId returns a boolean if a field has been set.
func (o *CertificateChainInfo) HasCertificateIssuerItemId() bool {
	if o != nil && o.CertificateIssuerItemId != nil {
		return true
	}

	return false
}

// SetCertificateIssuerItemId gets a reference to the given int64 and assigns it to the CertificateIssuerItemId field.
func (o *CertificateChainInfo) SetCertificateIssuerItemId(v int64) {
	o.CertificateIssuerItemId = &v
}

// GetCertificateIssuerName returns the CertificateIssuerName field value if set, zero value otherwise.
func (o *CertificateChainInfo) GetCertificateIssuerName() string {
	if o == nil || o.CertificateIssuerName == nil {
		var ret string
		return ret
	}
	return *o.CertificateIssuerName
}

// GetCertificateIssuerNameOk returns a tuple with the CertificateIssuerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateChainInfo) GetCertificateIssuerNameOk() (*string, bool) {
	if o == nil || o.CertificateIssuerName == nil {
		return nil, false
	}
	return o.CertificateIssuerName, true
}

// HasCertificateIssuerName returns a boolean if a field has been set.
func (o *CertificateChainInfo) HasCertificateIssuerName() bool {
	if o != nil && o.CertificateIssuerName != nil {
		return true
	}

	return false
}

// SetCertificateIssuerName gets a reference to the given string and assigns it to the CertificateIssuerName field.
func (o *CertificateChainInfo) SetCertificateIssuerName(v string) {
	o.CertificateIssuerName = &v
}

// GetCertificatePem returns the CertificatePem field value if set, zero value otherwise.
func (o *CertificateChainInfo) GetCertificatePem() string {
	if o == nil || o.CertificatePem == nil {
		var ret string
		return ret
	}
	return *o.CertificatePem
}

// GetCertificatePemOk returns a tuple with the CertificatePem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateChainInfo) GetCertificatePemOk() (*string, bool) {
	if o == nil || o.CertificatePem == nil {
		return nil, false
	}
	return o.CertificatePem, true
}

// HasCertificatePem returns a boolean if a field has been set.
func (o *CertificateChainInfo) HasCertificatePem() bool {
	if o != nil && o.CertificatePem != nil {
		return true
	}

	return false
}

// SetCertificatePem gets a reference to the given string and assigns it to the CertificatePem field.
func (o *CertificateChainInfo) SetCertificatePem(v string) {
	o.CertificatePem = &v
}

// GetCertificateStatus returns the CertificateStatus field value if set, zero value otherwise.
func (o *CertificateChainInfo) GetCertificateStatus() string {
	if o == nil || o.CertificateStatus == nil {
		var ret string
		return ret
	}
	return *o.CertificateStatus
}

// GetCertificateStatusOk returns a tuple with the CertificateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateChainInfo) GetCertificateStatusOk() (*string, bool) {
	if o == nil || o.CertificateStatus == nil {
		return nil, false
	}
	return o.CertificateStatus, true
}

// HasCertificateStatus returns a boolean if a field has been set.
func (o *CertificateChainInfo) HasCertificateStatus() bool {
	if o != nil && o.CertificateStatus != nil {
		return true
	}

	return false
}

// SetCertificateStatus gets a reference to the given string and assigns it to the CertificateStatus field.
func (o *CertificateChainInfo) SetCertificateStatus(v string) {
	o.CertificateStatus = &v
}

// GetExpirationEvents returns the ExpirationEvents field value if set, zero value otherwise.
func (o *CertificateChainInfo) GetExpirationEvents() []CertificateExpirationEvent {
	if o == nil || o.ExpirationEvents == nil {
		var ret []CertificateExpirationEvent
		return ret
	}
	return *o.ExpirationEvents
}

// GetExpirationEventsOk returns a tuple with the ExpirationEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateChainInfo) GetExpirationEventsOk() (*[]CertificateExpirationEvent, bool) {
	if o == nil || o.ExpirationEvents == nil {
		return nil, false
	}
	return o.ExpirationEvents, true
}

// HasExpirationEvents returns a boolean if a field has been set.
func (o *CertificateChainInfo) HasExpirationEvents() bool {
	if o != nil && o.ExpirationEvents != nil {
		return true
	}

	return false
}

// SetExpirationEvents gets a reference to the given []CertificateExpirationEvent and assigns it to the ExpirationEvents field.
func (o *CertificateChainInfo) SetExpirationEvents(v []CertificateExpirationEvent) {
	o.ExpirationEvents = &v
}

func (o CertificateChainInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertificateChain != nil {
		toSerialize["certificate_chain"] = o.CertificateChain
	}
	if o.CertificateFormat != nil {
		toSerialize["certificate_format"] = o.CertificateFormat
	}
	if o.CertificateIssuerItemId != nil {
		toSerialize["certificate_issuer_item_id"] = o.CertificateIssuerItemId
	}
	if o.CertificateIssuerName != nil {
		toSerialize["certificate_issuer_name"] = o.CertificateIssuerName
	}
	if o.CertificatePem != nil {
		toSerialize["certificate_pem"] = o.CertificatePem
	}
	if o.CertificateStatus != nil {
		toSerialize["certificate_status"] = o.CertificateStatus
	}
	if o.ExpirationEvents != nil {
		toSerialize["expiration_events"] = o.ExpirationEvents
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateChainInfo struct {
	value *CertificateChainInfo
	isSet bool
}

func (v NullableCertificateChainInfo) Get() *CertificateChainInfo {
	return v.value
}

func (v *NullableCertificateChainInfo) Set(val *CertificateChainInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateChainInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateChainInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateChainInfo(val *CertificateChainInfo) *NullableCertificateChainInfo {
	return &NullableCertificateChainInfo{value: val, isSet: true}
}

func (v NullableCertificateChainInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateChainInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


