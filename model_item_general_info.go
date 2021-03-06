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

// ItemGeneralInfo struct for ItemGeneralInfo
type ItemGeneralInfo struct {
	CertIssueDetails *CertificateIssueInfo `json:"cert_issue_details,omitempty"`
	ClassicKeyDetails *ClassicKeyDetailsInfo `json:"classic_key_details,omitempty"`
	DynamicSecretProducerDetails *DynamicSecretProducerInfo `json:"dynamic_secret_producer_details,omitempty"`
	RotatedSecretDetails *RotatedSecretDetailsInfo `json:"rotated_secret_details,omitempty"`
}

// NewItemGeneralInfo instantiates a new ItemGeneralInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemGeneralInfo() *ItemGeneralInfo {
	this := ItemGeneralInfo{}
	return &this
}

// NewItemGeneralInfoWithDefaults instantiates a new ItemGeneralInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemGeneralInfoWithDefaults() *ItemGeneralInfo {
	this := ItemGeneralInfo{}
	return &this
}

// GetCertIssueDetails returns the CertIssueDetails field value if set, zero value otherwise.
func (o *ItemGeneralInfo) GetCertIssueDetails() CertificateIssueInfo {
	if o == nil || o.CertIssueDetails == nil {
		var ret CertificateIssueInfo
		return ret
	}
	return *o.CertIssueDetails
}

// GetCertIssueDetailsOk returns a tuple with the CertIssueDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGeneralInfo) GetCertIssueDetailsOk() (*CertificateIssueInfo, bool) {
	if o == nil || o.CertIssueDetails == nil {
		return nil, false
	}
	return o.CertIssueDetails, true
}

// HasCertIssueDetails returns a boolean if a field has been set.
func (o *ItemGeneralInfo) HasCertIssueDetails() bool {
	if o != nil && o.CertIssueDetails != nil {
		return true
	}

	return false
}

// SetCertIssueDetails gets a reference to the given CertificateIssueInfo and assigns it to the CertIssueDetails field.
func (o *ItemGeneralInfo) SetCertIssueDetails(v CertificateIssueInfo) {
	o.CertIssueDetails = &v
}

// GetClassicKeyDetails returns the ClassicKeyDetails field value if set, zero value otherwise.
func (o *ItemGeneralInfo) GetClassicKeyDetails() ClassicKeyDetailsInfo {
	if o == nil || o.ClassicKeyDetails == nil {
		var ret ClassicKeyDetailsInfo
		return ret
	}
	return *o.ClassicKeyDetails
}

// GetClassicKeyDetailsOk returns a tuple with the ClassicKeyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGeneralInfo) GetClassicKeyDetailsOk() (*ClassicKeyDetailsInfo, bool) {
	if o == nil || o.ClassicKeyDetails == nil {
		return nil, false
	}
	return o.ClassicKeyDetails, true
}

// HasClassicKeyDetails returns a boolean if a field has been set.
func (o *ItemGeneralInfo) HasClassicKeyDetails() bool {
	if o != nil && o.ClassicKeyDetails != nil {
		return true
	}

	return false
}

// SetClassicKeyDetails gets a reference to the given ClassicKeyDetailsInfo and assigns it to the ClassicKeyDetails field.
func (o *ItemGeneralInfo) SetClassicKeyDetails(v ClassicKeyDetailsInfo) {
	o.ClassicKeyDetails = &v
}

// GetDynamicSecretProducerDetails returns the DynamicSecretProducerDetails field value if set, zero value otherwise.
func (o *ItemGeneralInfo) GetDynamicSecretProducerDetails() DynamicSecretProducerInfo {
	if o == nil || o.DynamicSecretProducerDetails == nil {
		var ret DynamicSecretProducerInfo
		return ret
	}
	return *o.DynamicSecretProducerDetails
}

// GetDynamicSecretProducerDetailsOk returns a tuple with the DynamicSecretProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGeneralInfo) GetDynamicSecretProducerDetailsOk() (*DynamicSecretProducerInfo, bool) {
	if o == nil || o.DynamicSecretProducerDetails == nil {
		return nil, false
	}
	return o.DynamicSecretProducerDetails, true
}

// HasDynamicSecretProducerDetails returns a boolean if a field has been set.
func (o *ItemGeneralInfo) HasDynamicSecretProducerDetails() bool {
	if o != nil && o.DynamicSecretProducerDetails != nil {
		return true
	}

	return false
}

// SetDynamicSecretProducerDetails gets a reference to the given DynamicSecretProducerInfo and assigns it to the DynamicSecretProducerDetails field.
func (o *ItemGeneralInfo) SetDynamicSecretProducerDetails(v DynamicSecretProducerInfo) {
	o.DynamicSecretProducerDetails = &v
}

// GetRotatedSecretDetails returns the RotatedSecretDetails field value if set, zero value otherwise.
func (o *ItemGeneralInfo) GetRotatedSecretDetails() RotatedSecretDetailsInfo {
	if o == nil || o.RotatedSecretDetails == nil {
		var ret RotatedSecretDetailsInfo
		return ret
	}
	return *o.RotatedSecretDetails
}

// GetRotatedSecretDetailsOk returns a tuple with the RotatedSecretDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemGeneralInfo) GetRotatedSecretDetailsOk() (*RotatedSecretDetailsInfo, bool) {
	if o == nil || o.RotatedSecretDetails == nil {
		return nil, false
	}
	return o.RotatedSecretDetails, true
}

// HasRotatedSecretDetails returns a boolean if a field has been set.
func (o *ItemGeneralInfo) HasRotatedSecretDetails() bool {
	if o != nil && o.RotatedSecretDetails != nil {
		return true
	}

	return false
}

// SetRotatedSecretDetails gets a reference to the given RotatedSecretDetailsInfo and assigns it to the RotatedSecretDetails field.
func (o *ItemGeneralInfo) SetRotatedSecretDetails(v RotatedSecretDetailsInfo) {
	o.RotatedSecretDetails = &v
}

func (o ItemGeneralInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertIssueDetails != nil {
		toSerialize["cert_issue_details"] = o.CertIssueDetails
	}
	if o.ClassicKeyDetails != nil {
		toSerialize["classic_key_details"] = o.ClassicKeyDetails
	}
	if o.DynamicSecretProducerDetails != nil {
		toSerialize["dynamic_secret_producer_details"] = o.DynamicSecretProducerDetails
	}
	if o.RotatedSecretDetails != nil {
		toSerialize["rotated_secret_details"] = o.RotatedSecretDetails
	}
	return json.Marshal(toSerialize)
}

type NullableItemGeneralInfo struct {
	value *ItemGeneralInfo
	isSet bool
}

func (v NullableItemGeneralInfo) Get() *ItemGeneralInfo {
	return v.value
}

func (v *NullableItemGeneralInfo) Set(val *ItemGeneralInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableItemGeneralInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableItemGeneralInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemGeneralInfo(val *ItemGeneralInfo) *NullableItemGeneralInfo {
	return &NullableItemGeneralInfo{value: val, isSet: true}
}

func (v NullableItemGeneralInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemGeneralInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


