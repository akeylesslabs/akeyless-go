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

// checks if the GatewayUpdateProducerAzureOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayUpdateProducerAzureOutput{}

// GatewayUpdateProducerAzureOutput struct for GatewayUpdateProducerAzureOutput
type GatewayUpdateProducerAzureOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayUpdateProducerAzureOutput instantiates a new GatewayUpdateProducerAzureOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerAzureOutput() *GatewayUpdateProducerAzureOutput {
	this := GatewayUpdateProducerAzureOutput{}
	return &this
}

// NewGatewayUpdateProducerAzureOutputWithDefaults instantiates a new GatewayUpdateProducerAzureOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerAzureOutputWithDefaults() *GatewayUpdateProducerAzureOutput {
	this := GatewayUpdateProducerAzureOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayUpdateProducerAzureOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || IsNil(o.ProducerDetails) {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerAzureOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || IsNil(o.ProducerDetails) {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayUpdateProducerAzureOutput) HasProducerDetails() bool {
	if o != nil && !IsNil(o.ProducerDetails) {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayUpdateProducerAzureOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayUpdateProducerAzureOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayUpdateProducerAzureOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProducerDetails) {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return toSerialize, nil
}

type NullableGatewayUpdateProducerAzureOutput struct {
	value *GatewayUpdateProducerAzureOutput
	isSet bool
}

func (v NullableGatewayUpdateProducerAzureOutput) Get() *GatewayUpdateProducerAzureOutput {
	return v.value
}

func (v *NullableGatewayUpdateProducerAzureOutput) Set(val *GatewayUpdateProducerAzureOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerAzureOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerAzureOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerAzureOutput(val *GatewayUpdateProducerAzureOutput) *NullableGatewayUpdateProducerAzureOutput {
	return &NullableGatewayUpdateProducerAzureOutput{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerAzureOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerAzureOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


