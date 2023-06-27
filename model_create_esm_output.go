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

// CreateESMOutput struct for CreateESMOutput
type CreateESMOutput struct {
	ExternalSecretManagerId *int64 `json:"external_secret_manager_id,omitempty"`
	ExternalSecretManagerName *string `json:"external_secret_manager_name,omitempty"`
}

// NewCreateESMOutput instantiates a new CreateESMOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateESMOutput() *CreateESMOutput {
	this := CreateESMOutput{}
	return &this
}

// NewCreateESMOutputWithDefaults instantiates a new CreateESMOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateESMOutputWithDefaults() *CreateESMOutput {
	this := CreateESMOutput{}
	return &this
}

// GetExternalSecretManagerId returns the ExternalSecretManagerId field value if set, zero value otherwise.
func (o *CreateESMOutput) GetExternalSecretManagerId() int64 {
	if o == nil || o.ExternalSecretManagerId == nil {
		var ret int64
		return ret
	}
	return *o.ExternalSecretManagerId
}

// GetExternalSecretManagerIdOk returns a tuple with the ExternalSecretManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESMOutput) GetExternalSecretManagerIdOk() (*int64, bool) {
	if o == nil || o.ExternalSecretManagerId == nil {
		return nil, false
	}
	return o.ExternalSecretManagerId, true
}

// HasExternalSecretManagerId returns a boolean if a field has been set.
func (o *CreateESMOutput) HasExternalSecretManagerId() bool {
	if o != nil && o.ExternalSecretManagerId != nil {
		return true
	}

	return false
}

// SetExternalSecretManagerId gets a reference to the given int64 and assigns it to the ExternalSecretManagerId field.
func (o *CreateESMOutput) SetExternalSecretManagerId(v int64) {
	o.ExternalSecretManagerId = &v
}

// GetExternalSecretManagerName returns the ExternalSecretManagerName field value if set, zero value otherwise.
func (o *CreateESMOutput) GetExternalSecretManagerName() string {
	if o == nil || o.ExternalSecretManagerName == nil {
		var ret string
		return ret
	}
	return *o.ExternalSecretManagerName
}

// GetExternalSecretManagerNameOk returns a tuple with the ExternalSecretManagerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateESMOutput) GetExternalSecretManagerNameOk() (*string, bool) {
	if o == nil || o.ExternalSecretManagerName == nil {
		return nil, false
	}
	return o.ExternalSecretManagerName, true
}

// HasExternalSecretManagerName returns a boolean if a field has been set.
func (o *CreateESMOutput) HasExternalSecretManagerName() bool {
	if o != nil && o.ExternalSecretManagerName != nil {
		return true
	}

	return false
}

// SetExternalSecretManagerName gets a reference to the given string and assigns it to the ExternalSecretManagerName field.
func (o *CreateESMOutput) SetExternalSecretManagerName(v string) {
	o.ExternalSecretManagerName = &v
}

func (o CreateESMOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalSecretManagerId != nil {
		toSerialize["external_secret_manager_id"] = o.ExternalSecretManagerId
	}
	if o.ExternalSecretManagerName != nil {
		toSerialize["external_secret_manager_name"] = o.ExternalSecretManagerName
	}
	return json.Marshal(toSerialize)
}

type NullableCreateESMOutput struct {
	value *CreateESMOutput
	isSet bool
}

func (v NullableCreateESMOutput) Get() *CreateESMOutput {
	return v.value
}

func (v *NullableCreateESMOutput) Set(val *CreateESMOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateESMOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateESMOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateESMOutput(val *CreateESMOutput) *NullableCreateESMOutput {
	return &NullableCreateESMOutput{value: val, isSet: true}
}

func (v NullableCreateESMOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateESMOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


