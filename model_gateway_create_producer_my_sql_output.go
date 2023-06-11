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

// checks if the GatewayCreateProducerMySQLOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayCreateProducerMySQLOutput{}

// GatewayCreateProducerMySQLOutput struct for GatewayCreateProducerMySQLOutput
type GatewayCreateProducerMySQLOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayCreateProducerMySQLOutput instantiates a new GatewayCreateProducerMySQLOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerMySQLOutput() *GatewayCreateProducerMySQLOutput {
	this := GatewayCreateProducerMySQLOutput{}
	return &this
}

// NewGatewayCreateProducerMySQLOutputWithDefaults instantiates a new GatewayCreateProducerMySQLOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerMySQLOutputWithDefaults() *GatewayCreateProducerMySQLOutput {
	this := GatewayCreateProducerMySQLOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQLOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || IsNil(o.ProducerDetails) {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQLOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || IsNil(o.ProducerDetails) {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQLOutput) HasProducerDetails() bool {
	if o != nil && !IsNil(o.ProducerDetails) {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayCreateProducerMySQLOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayCreateProducerMySQLOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayCreateProducerMySQLOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProducerDetails) {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return toSerialize, nil
}

type NullableGatewayCreateProducerMySQLOutput struct {
	value *GatewayCreateProducerMySQLOutput
	isSet bool
}

func (v NullableGatewayCreateProducerMySQLOutput) Get() *GatewayCreateProducerMySQLOutput {
	return v.value
}

func (v *NullableGatewayCreateProducerMySQLOutput) Set(val *GatewayCreateProducerMySQLOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerMySQLOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerMySQLOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerMySQLOutput(val *GatewayCreateProducerMySQLOutput) *NullableGatewayCreateProducerMySQLOutput {
	return &NullableGatewayCreateProducerMySQLOutput{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerMySQLOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerMySQLOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


