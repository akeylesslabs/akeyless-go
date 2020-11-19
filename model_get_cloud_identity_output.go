/*
 * includes all operations supported by Akeyless.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GetCloudIdentityOutput struct for GetCloudIdentityOutput
type GetCloudIdentityOutput struct {
	Token *string `json:"token,omitempty"`
}

// NewGetCloudIdentityOutput instantiates a new GetCloudIdentityOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCloudIdentityOutput() *GetCloudIdentityOutput {
	this := GetCloudIdentityOutput{}
	return &this
}

// NewGetCloudIdentityOutputWithDefaults instantiates a new GetCloudIdentityOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCloudIdentityOutputWithDefaults() *GetCloudIdentityOutput {
	this := GetCloudIdentityOutput{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetCloudIdentityOutput) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudIdentityOutput) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetCloudIdentityOutput) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetCloudIdentityOutput) SetToken(v string) {
	o.Token = &v
}

func (o GetCloudIdentityOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableGetCloudIdentityOutput struct {
	value *GetCloudIdentityOutput
	isSet bool
}

func (v NullableGetCloudIdentityOutput) Get() *GetCloudIdentityOutput {
	return v.value
}

func (v *NullableGetCloudIdentityOutput) Set(val *GetCloudIdentityOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCloudIdentityOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCloudIdentityOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCloudIdentityOutput(val *GetCloudIdentityOutput) *NullableGetCloudIdentityOutput {
	return &NullableGetCloudIdentityOutput{value: val, isSet: true}
}

func (v NullableGetCloudIdentityOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCloudIdentityOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


