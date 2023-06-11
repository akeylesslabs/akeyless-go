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

// checks if the GatewayUpdateProducerMySQLOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayUpdateProducerMySQLOutput{}

// GatewayUpdateProducerMySQLOutput struct for GatewayUpdateProducerMySQLOutput
type GatewayUpdateProducerMySQLOutput struct {
	ProducerDetails *DSProducerDetails `json:"producer_details,omitempty"`
}

// NewGatewayUpdateProducerMySQLOutput instantiates a new GatewayUpdateProducerMySQLOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerMySQLOutput() *GatewayUpdateProducerMySQLOutput {
	this := GatewayUpdateProducerMySQLOutput{}
	return &this
}

// NewGatewayUpdateProducerMySQLOutputWithDefaults instantiates a new GatewayUpdateProducerMySQLOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerMySQLOutputWithDefaults() *GatewayUpdateProducerMySQLOutput {
	this := GatewayUpdateProducerMySQLOutput{}
	return &this
}

// GetProducerDetails returns the ProducerDetails field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQLOutput) GetProducerDetails() DSProducerDetails {
	if o == nil || IsNil(o.ProducerDetails) {
		var ret DSProducerDetails
		return ret
	}
	return *o.ProducerDetails
}

// GetProducerDetailsOk returns a tuple with the ProducerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQLOutput) GetProducerDetailsOk() (*DSProducerDetails, bool) {
	if o == nil || IsNil(o.ProducerDetails) {
		return nil, false
	}
	return o.ProducerDetails, true
}

// HasProducerDetails returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQLOutput) HasProducerDetails() bool {
	if o != nil && !IsNil(o.ProducerDetails) {
		return true
	}

	return false
}

// SetProducerDetails gets a reference to the given DSProducerDetails and assigns it to the ProducerDetails field.
func (o *GatewayUpdateProducerMySQLOutput) SetProducerDetails(v DSProducerDetails) {
	o.ProducerDetails = &v
}

func (o GatewayUpdateProducerMySQLOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayUpdateProducerMySQLOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProducerDetails) {
		toSerialize["producer_details"] = o.ProducerDetails
	}
	return toSerialize, nil
}

type NullableGatewayUpdateProducerMySQLOutput struct {
	value *GatewayUpdateProducerMySQLOutput
	isSet bool
}

func (v NullableGatewayUpdateProducerMySQLOutput) Get() *GatewayUpdateProducerMySQLOutput {
	return v.value
}

func (v *NullableGatewayUpdateProducerMySQLOutput) Set(val *GatewayUpdateProducerMySQLOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerMySQLOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerMySQLOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerMySQLOutput(val *GatewayUpdateProducerMySQLOutput) *NullableGatewayUpdateProducerMySQLOutput {
	return &NullableGatewayUpdateProducerMySQLOutput{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerMySQLOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerMySQLOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


