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

// DynamicSecretProducerInfo DynamicSecretProducerInfo The dynamic secret producer info This parameter relevant and required only in case of create update dynamic secret.
type DynamicSecretProducerInfo struct {
	GwClusterId *int64 `json:"gw_cluster_id,omitempty"`
	ProducerMetadata *string `json:"producer_metadata,omitempty"`
	// RotationStatus defines types of rotation Status
	ProducerStatus *string `json:"producer_status,omitempty"`
	ProducerType *string `json:"producer_type,omitempty"`
}

// NewDynamicSecretProducerInfo instantiates a new DynamicSecretProducerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretProducerInfo() *DynamicSecretProducerInfo {
	this := DynamicSecretProducerInfo{}
	return &this
}

// NewDynamicSecretProducerInfoWithDefaults instantiates a new DynamicSecretProducerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretProducerInfoWithDefaults() *DynamicSecretProducerInfo {
	this := DynamicSecretProducerInfo{}
	return &this
}

// GetGwClusterId returns the GwClusterId field value if set, zero value otherwise.
func (o *DynamicSecretProducerInfo) GetGwClusterId() int64 {
	if o == nil || o.GwClusterId == nil {
		var ret int64
		return ret
	}
	return *o.GwClusterId
}

// GetGwClusterIdOk returns a tuple with the GwClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretProducerInfo) GetGwClusterIdOk() (*int64, bool) {
	if o == nil || o.GwClusterId == nil {
		return nil, false
	}
	return o.GwClusterId, true
}

// HasGwClusterId returns a boolean if a field has been set.
func (o *DynamicSecretProducerInfo) HasGwClusterId() bool {
	if o != nil && o.GwClusterId != nil {
		return true
	}

	return false
}

// SetGwClusterId gets a reference to the given int64 and assigns it to the GwClusterId field.
func (o *DynamicSecretProducerInfo) SetGwClusterId(v int64) {
	o.GwClusterId = &v
}

// GetProducerMetadata returns the ProducerMetadata field value if set, zero value otherwise.
func (o *DynamicSecretProducerInfo) GetProducerMetadata() string {
	if o == nil || o.ProducerMetadata == nil {
		var ret string
		return ret
	}
	return *o.ProducerMetadata
}

// GetProducerMetadataOk returns a tuple with the ProducerMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretProducerInfo) GetProducerMetadataOk() (*string, bool) {
	if o == nil || o.ProducerMetadata == nil {
		return nil, false
	}
	return o.ProducerMetadata, true
}

// HasProducerMetadata returns a boolean if a field has been set.
func (o *DynamicSecretProducerInfo) HasProducerMetadata() bool {
	if o != nil && o.ProducerMetadata != nil {
		return true
	}

	return false
}

// SetProducerMetadata gets a reference to the given string and assigns it to the ProducerMetadata field.
func (o *DynamicSecretProducerInfo) SetProducerMetadata(v string) {
	o.ProducerMetadata = &v
}

// GetProducerStatus returns the ProducerStatus field value if set, zero value otherwise.
func (o *DynamicSecretProducerInfo) GetProducerStatus() string {
	if o == nil || o.ProducerStatus == nil {
		var ret string
		return ret
	}
	return *o.ProducerStatus
}

// GetProducerStatusOk returns a tuple with the ProducerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretProducerInfo) GetProducerStatusOk() (*string, bool) {
	if o == nil || o.ProducerStatus == nil {
		return nil, false
	}
	return o.ProducerStatus, true
}

// HasProducerStatus returns a boolean if a field has been set.
func (o *DynamicSecretProducerInfo) HasProducerStatus() bool {
	if o != nil && o.ProducerStatus != nil {
		return true
	}

	return false
}

// SetProducerStatus gets a reference to the given string and assigns it to the ProducerStatus field.
func (o *DynamicSecretProducerInfo) SetProducerStatus(v string) {
	o.ProducerStatus = &v
}

// GetProducerType returns the ProducerType field value if set, zero value otherwise.
func (o *DynamicSecretProducerInfo) GetProducerType() string {
	if o == nil || o.ProducerType == nil {
		var ret string
		return ret
	}
	return *o.ProducerType
}

// GetProducerTypeOk returns a tuple with the ProducerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretProducerInfo) GetProducerTypeOk() (*string, bool) {
	if o == nil || o.ProducerType == nil {
		return nil, false
	}
	return o.ProducerType, true
}

// HasProducerType returns a boolean if a field has been set.
func (o *DynamicSecretProducerInfo) HasProducerType() bool {
	if o != nil && o.ProducerType != nil {
		return true
	}

	return false
}

// SetProducerType gets a reference to the given string and assigns it to the ProducerType field.
func (o *DynamicSecretProducerInfo) SetProducerType(v string) {
	o.ProducerType = &v
}

func (o DynamicSecretProducerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GwClusterId != nil {
		toSerialize["gw_cluster_id"] = o.GwClusterId
	}
	if o.ProducerMetadata != nil {
		toSerialize["producer_metadata"] = o.ProducerMetadata
	}
	if o.ProducerStatus != nil {
		toSerialize["producer_status"] = o.ProducerStatus
	}
	if o.ProducerType != nil {
		toSerialize["producer_type"] = o.ProducerType
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretProducerInfo struct {
	value *DynamicSecretProducerInfo
	isSet bool
}

func (v NullableDynamicSecretProducerInfo) Get() *DynamicSecretProducerInfo {
	return v.value
}

func (v *NullableDynamicSecretProducerInfo) Set(val *DynamicSecretProducerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretProducerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretProducerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretProducerInfo(val *DynamicSecretProducerInfo) *NullableDynamicSecretProducerInfo {
	return &NullableDynamicSecretProducerInfo{value: val, isSet: true}
}

func (v NullableDynamicSecretProducerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretProducerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


