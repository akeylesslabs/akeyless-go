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

// CfInfo struct for CfInfo
type CfInfo struct {
	CfName *string `json:"cf_name,omitempty"`
	Description *string `json:"description,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewCfInfo instantiates a new CfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCfInfo() *CfInfo {
	this := CfInfo{}
	return &this
}

// NewCfInfoWithDefaults instantiates a new CfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCfInfoWithDefaults() *CfInfo {
	this := CfInfo{}
	return &this
}

// GetCfName returns the CfName field value if set, zero value otherwise.
func (o *CfInfo) GetCfName() string {
	if o == nil || o.CfName == nil {
		var ret string
		return ret
	}
	return *o.CfName
}

// GetCfNameOk returns a tuple with the CfName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfInfo) GetCfNameOk() (*string, bool) {
	if o == nil || o.CfName == nil {
		return nil, false
	}
	return o.CfName, true
}

// HasCfName returns a boolean if a field has been set.
func (o *CfInfo) HasCfName() bool {
	if o != nil && o.CfName != nil {
		return true
	}

	return false
}

// SetCfName gets a reference to the given string and assigns it to the CfName field.
func (o *CfInfo) SetCfName(v string) {
	o.CfName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CfInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CfInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CfInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *CfInfo) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfInfo) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *CfInfo) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *CfInfo) SetHash(v string) {
	o.Hash = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CfInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CfInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CfInfo) SetId(v string) {
	o.Id = &v
}

func (o CfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CfName != nil {
		toSerialize["cf_name"] = o.CfName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCfInfo struct {
	value *CfInfo
	isSet bool
}

func (v NullableCfInfo) Get() *CfInfo {
	return v.value
}

func (v *NullableCfInfo) Set(val *CfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCfInfo(val *CfInfo) *NullableCfInfo {
	return &NullableCfInfo{value: val, isSet: true}
}

func (v NullableCfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


