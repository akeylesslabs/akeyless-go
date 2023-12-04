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
	"time"
)

// APIKeyAccessRules struct for APIKeyAccessRules
type APIKeyAccessRules struct {
	Alg *string `json:"alg,omitempty"`
	// The public key value of the API-key.
	Key *string `json:"key,omitempty"`
	ModificationDate *time.Time `json:"modification_date,omitempty"`
}

// NewAPIKeyAccessRules instantiates a new APIKeyAccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIKeyAccessRules() *APIKeyAccessRules {
	this := APIKeyAccessRules{}
	return &this
}

// NewAPIKeyAccessRulesWithDefaults instantiates a new APIKeyAccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIKeyAccessRulesWithDefaults() *APIKeyAccessRules {
	this := APIKeyAccessRules{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *APIKeyAccessRules) GetAlg() string {
	if o == nil || o.Alg == nil {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeyAccessRules) GetAlgOk() (*string, bool) {
	if o == nil || o.Alg == nil {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *APIKeyAccessRules) HasAlg() bool {
	if o != nil && o.Alg != nil {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *APIKeyAccessRules) SetAlg(v string) {
	o.Alg = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *APIKeyAccessRules) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeyAccessRules) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *APIKeyAccessRules) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *APIKeyAccessRules) SetKey(v string) {
	o.Key = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *APIKeyAccessRules) GetModificationDate() time.Time {
	if o == nil || o.ModificationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeyAccessRules) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || o.ModificationDate == nil {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *APIKeyAccessRules) HasModificationDate() bool {
	if o != nil && o.ModificationDate != nil {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *APIKeyAccessRules) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

func (o APIKeyAccessRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alg != nil {
		toSerialize["alg"] = o.Alg
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.ModificationDate != nil {
		toSerialize["modification_date"] = o.ModificationDate
	}
	return json.Marshal(toSerialize)
}

type NullableAPIKeyAccessRules struct {
	value *APIKeyAccessRules
	isSet bool
}

func (v NullableAPIKeyAccessRules) Get() *APIKeyAccessRules {
	return v.value
}

func (v *NullableAPIKeyAccessRules) Set(val *APIKeyAccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIKeyAccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIKeyAccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIKeyAccessRules(val *APIKeyAccessRules) *NullableAPIKeyAccessRules {
	return &NullableAPIKeyAccessRules{value: val, isSet: true}
}

func (v NullableAPIKeyAccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIKeyAccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


