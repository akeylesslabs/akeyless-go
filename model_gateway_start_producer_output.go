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

// checks if the GatewayStartProducerOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayStartProducerOutput{}

// GatewayStartProducerOutput struct for GatewayStartProducerOutput
type GatewayStartProducerOutput struct {
	ProducerName *string `json:"producer_name,omitempty"`
}

// NewGatewayStartProducerOutput instantiates a new GatewayStartProducerOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayStartProducerOutput() *GatewayStartProducerOutput {
	this := GatewayStartProducerOutput{}
	return &this
}

// NewGatewayStartProducerOutputWithDefaults instantiates a new GatewayStartProducerOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayStartProducerOutputWithDefaults() *GatewayStartProducerOutput {
	this := GatewayStartProducerOutput{}
	return &this
}

// GetProducerName returns the ProducerName field value if set, zero value otherwise.
func (o *GatewayStartProducerOutput) GetProducerName() string {
	if o == nil || IsNil(o.ProducerName) {
		var ret string
		return ret
	}
	return *o.ProducerName
}

// GetProducerNameOk returns a tuple with the ProducerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayStartProducerOutput) GetProducerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProducerName) {
		return nil, false
	}
	return o.ProducerName, true
}

// HasProducerName returns a boolean if a field has been set.
func (o *GatewayStartProducerOutput) HasProducerName() bool {
	if o != nil && !IsNil(o.ProducerName) {
		return true
	}

	return false
}

// SetProducerName gets a reference to the given string and assigns it to the ProducerName field.
func (o *GatewayStartProducerOutput) SetProducerName(v string) {
	o.ProducerName = &v
}

func (o GatewayStartProducerOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayStartProducerOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProducerName) {
		toSerialize["producer_name"] = o.ProducerName
	}
	return toSerialize, nil
}

type NullableGatewayStartProducerOutput struct {
	value *GatewayStartProducerOutput
	isSet bool
}

func (v NullableGatewayStartProducerOutput) Get() *GatewayStartProducerOutput {
	return v.value
}

func (v *NullableGatewayStartProducerOutput) Set(val *GatewayStartProducerOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayStartProducerOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayStartProducerOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayStartProducerOutput(val *GatewayStartProducerOutput) *NullableGatewayStartProducerOutput {
	return &NullableGatewayStartProducerOutput{value: val, isSet: true}
}

func (v NullableGatewayStartProducerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayStartProducerOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


