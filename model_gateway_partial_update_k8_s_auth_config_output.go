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

// GatewayPartialUpdateK8SAuthConfigOutput struct for GatewayPartialUpdateK8SAuthConfigOutput
type GatewayPartialUpdateK8SAuthConfigOutput struct {
	ClusterId *string `json:"cluster_id,omitempty"`
	PartsChange *ConfigChange `json:"parts_change,omitempty"`
	TotalHash *string `json:"total_hash,omitempty"`
}

// NewGatewayPartialUpdateK8SAuthConfigOutput instantiates a new GatewayPartialUpdateK8SAuthConfigOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayPartialUpdateK8SAuthConfigOutput() *GatewayPartialUpdateK8SAuthConfigOutput {
	this := GatewayPartialUpdateK8SAuthConfigOutput{}
	return &this
}

// NewGatewayPartialUpdateK8SAuthConfigOutputWithDefaults instantiates a new GatewayPartialUpdateK8SAuthConfigOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayPartialUpdateK8SAuthConfigOutputWithDefaults() *GatewayPartialUpdateK8SAuthConfigOutput {
	this := GatewayPartialUpdateK8SAuthConfigOutput{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetPartsChange returns the PartsChange field value if set, zero value otherwise.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) GetPartsChange() ConfigChange {
	if o == nil || o.PartsChange == nil {
		var ret ConfigChange
		return ret
	}
	return *o.PartsChange
}

// GetPartsChangeOk returns a tuple with the PartsChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) GetPartsChangeOk() (*ConfigChange, bool) {
	if o == nil || o.PartsChange == nil {
		return nil, false
	}
	return o.PartsChange, true
}

// HasPartsChange returns a boolean if a field has been set.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) HasPartsChange() bool {
	if o != nil && o.PartsChange != nil {
		return true
	}

	return false
}

// SetPartsChange gets a reference to the given ConfigChange and assigns it to the PartsChange field.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) SetPartsChange(v ConfigChange) {
	o.PartsChange = &v
}

// GetTotalHash returns the TotalHash field value if set, zero value otherwise.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) GetTotalHash() string {
	if o == nil || o.TotalHash == nil {
		var ret string
		return ret
	}
	return *o.TotalHash
}

// GetTotalHashOk returns a tuple with the TotalHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) GetTotalHashOk() (*string, bool) {
	if o == nil || o.TotalHash == nil {
		return nil, false
	}
	return o.TotalHash, true
}

// HasTotalHash returns a boolean if a field has been set.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) HasTotalHash() bool {
	if o != nil && o.TotalHash != nil {
		return true
	}

	return false
}

// SetTotalHash gets a reference to the given string and assigns it to the TotalHash field.
func (o *GatewayPartialUpdateK8SAuthConfigOutput) SetTotalHash(v string) {
	o.TotalHash = &v
}

func (o GatewayPartialUpdateK8SAuthConfigOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.PartsChange != nil {
		toSerialize["parts_change"] = o.PartsChange
	}
	if o.TotalHash != nil {
		toSerialize["total_hash"] = o.TotalHash
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayPartialUpdateK8SAuthConfigOutput struct {
	value *GatewayPartialUpdateK8SAuthConfigOutput
	isSet bool
}

func (v NullableGatewayPartialUpdateK8SAuthConfigOutput) Get() *GatewayPartialUpdateK8SAuthConfigOutput {
	return v.value
}

func (v *NullableGatewayPartialUpdateK8SAuthConfigOutput) Set(val *GatewayPartialUpdateK8SAuthConfigOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPartialUpdateK8SAuthConfigOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPartialUpdateK8SAuthConfigOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPartialUpdateK8SAuthConfigOutput(val *GatewayPartialUpdateK8SAuthConfigOutput) *NullableGatewayPartialUpdateK8SAuthConfigOutput {
	return &NullableGatewayPartialUpdateK8SAuthConfigOutput{value: val, isSet: true}
}

func (v NullableGatewayPartialUpdateK8SAuthConfigOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPartialUpdateK8SAuthConfigOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


