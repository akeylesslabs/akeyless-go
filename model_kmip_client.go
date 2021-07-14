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

// KMIPClient struct for KMIPClient
type KMIPClient struct {
	Environment *string `json:"environment,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Rules *[]PathRule `json:"rules,omitempty"`
}

// NewKMIPClient instantiates a new KMIPClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKMIPClient() *KMIPClient {
	this := KMIPClient{}
	return &this
}

// NewKMIPClientWithDefaults instantiates a new KMIPClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKMIPClientWithDefaults() *KMIPClient {
	this := KMIPClient{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *KMIPClient) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KMIPClient) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *KMIPClient) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *KMIPClient) SetEnvironment(v string) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KMIPClient) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KMIPClient) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KMIPClient) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KMIPClient) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KMIPClient) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KMIPClient) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KMIPClient) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KMIPClient) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *KMIPClient) GetRules() []PathRule {
	if o == nil || o.Rules == nil {
		var ret []PathRule
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KMIPClient) GetRulesOk() (*[]PathRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *KMIPClient) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []PathRule and assigns it to the Rules field.
func (o *KMIPClient) SetRules(v []PathRule) {
	o.Rules = &v
}

func (o KMIPClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableKMIPClient struct {
	value *KMIPClient
	isSet bool
}

func (v NullableKMIPClient) Get() *KMIPClient {
	return v.value
}

func (v *NullableKMIPClient) Set(val *KMIPClient) {
	v.value = val
	v.isSet = true
}

func (v NullableKMIPClient) IsSet() bool {
	return v.isSet
}

func (v *NullableKMIPClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKMIPClient(val *KMIPClient) *NullableKMIPClient {
	return &NullableKMIPClient{value: val, isSet: true}
}

func (v NullableKMIPClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKMIPClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


