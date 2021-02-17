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

// HashiPayload struct for HashiPayload
type HashiPayload struct {
	Namespaces *[]string `json:"namespaces,omitempty"`
	Token *string `json:"token,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewHashiPayload instantiates a new HashiPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashiPayload() *HashiPayload {
	this := HashiPayload{}
	return &this
}

// NewHashiPayloadWithDefaults instantiates a new HashiPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashiPayloadWithDefaults() *HashiPayload {
	this := HashiPayload{}
	return &this
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *HashiPayload) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return *o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashiPayload) GetNamespacesOk() (*[]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *HashiPayload) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *HashiPayload) SetNamespaces(v []string) {
	o.Namespaces = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *HashiPayload) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashiPayload) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *HashiPayload) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *HashiPayload) SetToken(v string) {
	o.Token = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HashiPayload) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashiPayload) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HashiPayload) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HashiPayload) SetUrl(v string) {
	o.Url = &v
}

func (o HashiPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableHashiPayload struct {
	value *HashiPayload
	isSet bool
}

func (v NullableHashiPayload) Get() *HashiPayload {
	return v.value
}

func (v *NullableHashiPayload) Set(val *HashiPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableHashiPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableHashiPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashiPayload(val *HashiPayload) *NullableHashiPayload {
	return &NullableHashiPayload{value: val, isSet: true}
}

func (v NullableHashiPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashiPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


