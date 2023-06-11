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

// checks if the GatewayMessageQueueInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayMessageQueueInfo{}

// GatewayMessageQueueInfo struct for GatewayMessageQueueInfo
type GatewayMessageQueueInfo struct {
	BroadcastQueueNameA *string `json:"broadcast_queue_name_a,omitempty"`
	MqType *string `json:"mq_type,omitempty"`
	QueueName *string `json:"queue_name,omitempty"`
	QueueUrl *string `json:"queue_url,omitempty"`
}

// NewGatewayMessageQueueInfo instantiates a new GatewayMessageQueueInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMessageQueueInfo() *GatewayMessageQueueInfo {
	this := GatewayMessageQueueInfo{}
	return &this
}

// NewGatewayMessageQueueInfoWithDefaults instantiates a new GatewayMessageQueueInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMessageQueueInfoWithDefaults() *GatewayMessageQueueInfo {
	this := GatewayMessageQueueInfo{}
	return &this
}

// GetBroadcastQueueNameA returns the BroadcastQueueNameA field value if set, zero value otherwise.
func (o *GatewayMessageQueueInfo) GetBroadcastQueueNameA() string {
	if o == nil || IsNil(o.BroadcastQueueNameA) {
		var ret string
		return ret
	}
	return *o.BroadcastQueueNameA
}

// GetBroadcastQueueNameAOk returns a tuple with the BroadcastQueueNameA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMessageQueueInfo) GetBroadcastQueueNameAOk() (*string, bool) {
	if o == nil || IsNil(o.BroadcastQueueNameA) {
		return nil, false
	}
	return o.BroadcastQueueNameA, true
}

// HasBroadcastQueueNameA returns a boolean if a field has been set.
func (o *GatewayMessageQueueInfo) HasBroadcastQueueNameA() bool {
	if o != nil && !IsNil(o.BroadcastQueueNameA) {
		return true
	}

	return false
}

// SetBroadcastQueueNameA gets a reference to the given string and assigns it to the BroadcastQueueNameA field.
func (o *GatewayMessageQueueInfo) SetBroadcastQueueNameA(v string) {
	o.BroadcastQueueNameA = &v
}

// GetMqType returns the MqType field value if set, zero value otherwise.
func (o *GatewayMessageQueueInfo) GetMqType() string {
	if o == nil || IsNil(o.MqType) {
		var ret string
		return ret
	}
	return *o.MqType
}

// GetMqTypeOk returns a tuple with the MqType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMessageQueueInfo) GetMqTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MqType) {
		return nil, false
	}
	return o.MqType, true
}

// HasMqType returns a boolean if a field has been set.
func (o *GatewayMessageQueueInfo) HasMqType() bool {
	if o != nil && !IsNil(o.MqType) {
		return true
	}

	return false
}

// SetMqType gets a reference to the given string and assigns it to the MqType field.
func (o *GatewayMessageQueueInfo) SetMqType(v string) {
	o.MqType = &v
}

// GetQueueName returns the QueueName field value if set, zero value otherwise.
func (o *GatewayMessageQueueInfo) GetQueueName() string {
	if o == nil || IsNil(o.QueueName) {
		var ret string
		return ret
	}
	return *o.QueueName
}

// GetQueueNameOk returns a tuple with the QueueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMessageQueueInfo) GetQueueNameOk() (*string, bool) {
	if o == nil || IsNil(o.QueueName) {
		return nil, false
	}
	return o.QueueName, true
}

// HasQueueName returns a boolean if a field has been set.
func (o *GatewayMessageQueueInfo) HasQueueName() bool {
	if o != nil && !IsNil(o.QueueName) {
		return true
	}

	return false
}

// SetQueueName gets a reference to the given string and assigns it to the QueueName field.
func (o *GatewayMessageQueueInfo) SetQueueName(v string) {
	o.QueueName = &v
}

// GetQueueUrl returns the QueueUrl field value if set, zero value otherwise.
func (o *GatewayMessageQueueInfo) GetQueueUrl() string {
	if o == nil || IsNil(o.QueueUrl) {
		var ret string
		return ret
	}
	return *o.QueueUrl
}

// GetQueueUrlOk returns a tuple with the QueueUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMessageQueueInfo) GetQueueUrlOk() (*string, bool) {
	if o == nil || IsNil(o.QueueUrl) {
		return nil, false
	}
	return o.QueueUrl, true
}

// HasQueueUrl returns a boolean if a field has been set.
func (o *GatewayMessageQueueInfo) HasQueueUrl() bool {
	if o != nil && !IsNil(o.QueueUrl) {
		return true
	}

	return false
}

// SetQueueUrl gets a reference to the given string and assigns it to the QueueUrl field.
func (o *GatewayMessageQueueInfo) SetQueueUrl(v string) {
	o.QueueUrl = &v
}

func (o GatewayMessageQueueInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayMessageQueueInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BroadcastQueueNameA) {
		toSerialize["broadcast_queue_name_a"] = o.BroadcastQueueNameA
	}
	if !IsNil(o.MqType) {
		toSerialize["mq_type"] = o.MqType
	}
	if !IsNil(o.QueueName) {
		toSerialize["queue_name"] = o.QueueName
	}
	if !IsNil(o.QueueUrl) {
		toSerialize["queue_url"] = o.QueueUrl
	}
	return toSerialize, nil
}

type NullableGatewayMessageQueueInfo struct {
	value *GatewayMessageQueueInfo
	isSet bool
}

func (v NullableGatewayMessageQueueInfo) Get() *GatewayMessageQueueInfo {
	return v.value
}

func (v *NullableGatewayMessageQueueInfo) Set(val *GatewayMessageQueueInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMessageQueueInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMessageQueueInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMessageQueueInfo(val *GatewayMessageQueueInfo) *NullableGatewayMessageQueueInfo {
	return &NullableGatewayMessageQueueInfo{value: val, isSet: true}
}

func (v NullableGatewayMessageQueueInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMessageQueueInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


