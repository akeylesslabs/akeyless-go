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

// checks if the LDAPAccessRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LDAPAccessRules{}

// LDAPAccessRules struct for LDAPAccessRules
type LDAPAccessRules struct {
	Alg *string `json:"alg,omitempty"`
	// Generate public/private key (the private key is required for the LDAP Auth Config in the Akeyless Gateway)
	GenKeyPair *string `json:"gen_key_pair,omitempty"`
	// The public key value of LDAP.
	Key *string `json:"key,omitempty"`
	// A unique identifier to distinguish different users
	UniqueIdentifier *string `json:"unique_identifier,omitempty"`
}

// NewLDAPAccessRules instantiates a new LDAPAccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPAccessRules() *LDAPAccessRules {
	this := LDAPAccessRules{}
	return &this
}

// NewLDAPAccessRulesWithDefaults instantiates a new LDAPAccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPAccessRulesWithDefaults() *LDAPAccessRules {
	this := LDAPAccessRules{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *LDAPAccessRules) GetAlg() string {
	if o == nil || IsNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPAccessRules) GetAlgOk() (*string, bool) {
	if o == nil || IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *LDAPAccessRules) HasAlg() bool {
	if o != nil && !IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *LDAPAccessRules) SetAlg(v string) {
	o.Alg = &v
}

// GetGenKeyPair returns the GenKeyPair field value if set, zero value otherwise.
func (o *LDAPAccessRules) GetGenKeyPair() string {
	if o == nil || IsNil(o.GenKeyPair) {
		var ret string
		return ret
	}
	return *o.GenKeyPair
}

// GetGenKeyPairOk returns a tuple with the GenKeyPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPAccessRules) GetGenKeyPairOk() (*string, bool) {
	if o == nil || IsNil(o.GenKeyPair) {
		return nil, false
	}
	return o.GenKeyPair, true
}

// HasGenKeyPair returns a boolean if a field has been set.
func (o *LDAPAccessRules) HasGenKeyPair() bool {
	if o != nil && !IsNil(o.GenKeyPair) {
		return true
	}

	return false
}

// SetGenKeyPair gets a reference to the given string and assigns it to the GenKeyPair field.
func (o *LDAPAccessRules) SetGenKeyPair(v string) {
	o.GenKeyPair = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LDAPAccessRules) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPAccessRules) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LDAPAccessRules) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LDAPAccessRules) SetKey(v string) {
	o.Key = &v
}

// GetUniqueIdentifier returns the UniqueIdentifier field value if set, zero value otherwise.
func (o *LDAPAccessRules) GetUniqueIdentifier() string {
	if o == nil || IsNil(o.UniqueIdentifier) {
		var ret string
		return ret
	}
	return *o.UniqueIdentifier
}

// GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPAccessRules) GetUniqueIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueIdentifier) {
		return nil, false
	}
	return o.UniqueIdentifier, true
}

// HasUniqueIdentifier returns a boolean if a field has been set.
func (o *LDAPAccessRules) HasUniqueIdentifier() bool {
	if o != nil && !IsNil(o.UniqueIdentifier) {
		return true
	}

	return false
}

// SetUniqueIdentifier gets a reference to the given string and assigns it to the UniqueIdentifier field.
func (o *LDAPAccessRules) SetUniqueIdentifier(v string) {
	o.UniqueIdentifier = &v
}

func (o LDAPAccessRules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LDAPAccessRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alg) {
		toSerialize["alg"] = o.Alg
	}
	if !IsNil(o.GenKeyPair) {
		toSerialize["gen_key_pair"] = o.GenKeyPair
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.UniqueIdentifier) {
		toSerialize["unique_identifier"] = o.UniqueIdentifier
	}
	return toSerialize, nil
}

type NullableLDAPAccessRules struct {
	value *LDAPAccessRules
	isSet bool
}

func (v NullableLDAPAccessRules) Get() *LDAPAccessRules {
	return v.value
}

func (v *NullableLDAPAccessRules) Set(val *LDAPAccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPAccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPAccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPAccessRules(val *LDAPAccessRules) *NullableLDAPAccessRules {
	return &NullableLDAPAccessRules{value: val, isSet: true}
}

func (v NullableLDAPAccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPAccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


