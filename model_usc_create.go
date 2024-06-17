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

// UscCreate uscCreate is a command that creates a new secret in a Universal Secrets Connector
type UscCreate struct {
	// Use this option if the universal secrets value is a base64 encoded binary
	BinaryValue *bool `json:"binary-value,omitempty"`
	// Description of the universal secrets
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The namespace (relevant for Hashi vault target)
	Namespace *string `json:"namespace,omitempty"`
	// Name for the new universal secrets
	SecretName string `json:"secret-name"`
	// Tags for the universal secrets
	Tags *map[string]string `json:"tags,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Name of the Universal Secrets Connector item
	UscName string `json:"usc-name"`
	// Value of the universal secrets item, either text or base64 encoded binary
	Value string `json:"value"`
}

// NewUscCreate instantiates a new UscCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUscCreate(secretName string, uscName string, value string, ) *UscCreate {
	this := UscCreate{}
	var json bool = false
	this.Json = &json
	this.SecretName = secretName
	this.UscName = uscName
	this.Value = value
	return &this
}

// NewUscCreateWithDefaults instantiates a new UscCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUscCreateWithDefaults() *UscCreate {
	this := UscCreate{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetBinaryValue returns the BinaryValue field value if set, zero value otherwise.
func (o *UscCreate) GetBinaryValue() bool {
	if o == nil || o.BinaryValue == nil {
		var ret bool
		return ret
	}
	return *o.BinaryValue
}

// GetBinaryValueOk returns a tuple with the BinaryValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscCreate) GetBinaryValueOk() (*bool, bool) {
	if o == nil || o.BinaryValue == nil {
		return nil, false
	}
	return o.BinaryValue, true
}

// HasBinaryValue returns a boolean if a field has been set.
func (o *UscCreate) HasBinaryValue() bool {
	if o != nil && o.BinaryValue != nil {
		return true
	}

	return false
}

// SetBinaryValue gets a reference to the given bool and assigns it to the BinaryValue field.
func (o *UscCreate) SetBinaryValue(v bool) {
	o.BinaryValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UscCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UscCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UscCreate) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UscCreate) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscCreate) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UscCreate) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UscCreate) SetJson(v bool) {
	o.Json = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *UscCreate) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscCreate) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *UscCreate) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *UscCreate) SetNamespace(v string) {
	o.Namespace = &v
}

// GetSecretName returns the SecretName field value
func (o *UscCreate) GetSecretName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *UscCreate) GetSecretNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecretName, true
}

// SetSecretName sets field value
func (o *UscCreate) SetSecretName(v string) {
	o.SecretName = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UscCreate) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscCreate) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UscCreate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *UscCreate) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UscCreate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscCreate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UscCreate) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UscCreate) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UscCreate) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UscCreate) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UscCreate) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UscCreate) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUscName returns the UscName field value
func (o *UscCreate) GetUscName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UscName
}

// GetUscNameOk returns a tuple with the UscName field value
// and a boolean to check if the value has been set.
func (o *UscCreate) GetUscNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UscName, true
}

// SetUscName sets field value
func (o *UscCreate) SetUscName(v string) {
	o.UscName = v
}

// GetValue returns the Value field value
func (o *UscCreate) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UscCreate) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UscCreate) SetValue(v string) {
	o.Value = v
}

func (o UscCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BinaryValue != nil {
		toSerialize["binary-value"] = o.BinaryValue
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if true {
		toSerialize["secret-name"] = o.SecretName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if true {
		toSerialize["usc-name"] = o.UscName
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUscCreate struct {
	value *UscCreate
	isSet bool
}

func (v NullableUscCreate) Get() *UscCreate {
	return v.value
}

func (v *NullableUscCreate) Set(val *UscCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableUscCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableUscCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUscCreate(val *UscCreate) *NullableUscCreate {
	return &NullableUscCreate{value: val, isSet: true}
}

func (v NullableUscCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUscCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


