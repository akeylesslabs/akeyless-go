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

// CreateGlobalSignAtlasTarget struct for CreateGlobalSignAtlasTarget
type CreateGlobalSignAtlasTarget struct {
	// API Key of the GlobalSign Atlas account
	ApiKey string `json:"api-key"`
	// API Secret of the GlobalSign Atlas account
	ApiSecret string `json:"api-secret"`
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Set the maximum number of versions, limited by the account settings defaults.
	MaxVersions *string `json:"max-versions,omitempty"`
	// Mutual TLS Certificate contents of the GlobalSign Atlas account encoded in base64, either mtls-cert-file-path or mtls-cert-data-base64 must be supplied
	MtlsCertDataBase64 *string `json:"mtls-cert-data-base64,omitempty"`
	// Mutual TLS Key contents of the GlobalSign Atlas account encoded in base64, either mtls-key-file-path or mtls-data-base64 must be supplied
	MtlsKeyDataBase64 *string `json:"mtls-key-data-base64,omitempty"`
	// Target name
	Name string `json:"name"`
	// Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h.
	Timeout *string `json:"timeout,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateGlobalSignAtlasTarget instantiates a new CreateGlobalSignAtlasTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGlobalSignAtlasTarget(apiKey string, apiSecret string, name string, ) *CreateGlobalSignAtlasTarget {
	this := CreateGlobalSignAtlasTarget{}
	this.ApiKey = apiKey
	this.ApiSecret = apiSecret
	var json bool = false
	this.Json = &json
	this.Name = name
	var timeout string = "5m"
	this.Timeout = &timeout
	return &this
}

// NewCreateGlobalSignAtlasTargetWithDefaults instantiates a new CreateGlobalSignAtlasTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGlobalSignAtlasTargetWithDefaults() *CreateGlobalSignAtlasTarget {
	this := CreateGlobalSignAtlasTarget{}
	var json bool = false
	this.Json = &json
	var timeout string = "5m"
	this.Timeout = &timeout
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *CreateGlobalSignAtlasTarget) GetApiKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetApiKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *CreateGlobalSignAtlasTarget) SetApiKey(v string) {
	o.ApiKey = v
}

// GetApiSecret returns the ApiSecret field value
func (o *CreateGlobalSignAtlasTarget) GetApiSecret() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApiSecret
}

// GetApiSecretOk returns a tuple with the ApiSecret field value
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetApiSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiSecret, true
}

// SetApiSecret sets field value
func (o *CreateGlobalSignAtlasTarget) SetApiSecret(v string) {
	o.ApiSecret = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateGlobalSignAtlasTarget) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateGlobalSignAtlasTarget) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateGlobalSignAtlasTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateGlobalSignAtlasTarget) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *CreateGlobalSignAtlasTarget) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetMtlsCertDataBase64 returns the MtlsCertDataBase64 field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetMtlsCertDataBase64() string {
	if o == nil || o.MtlsCertDataBase64 == nil {
		var ret string
		return ret
	}
	return *o.MtlsCertDataBase64
}

// GetMtlsCertDataBase64Ok returns a tuple with the MtlsCertDataBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetMtlsCertDataBase64Ok() (*string, bool) {
	if o == nil || o.MtlsCertDataBase64 == nil {
		return nil, false
	}
	return o.MtlsCertDataBase64, true
}

// HasMtlsCertDataBase64 returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasMtlsCertDataBase64() bool {
	if o != nil && o.MtlsCertDataBase64 != nil {
		return true
	}

	return false
}

// SetMtlsCertDataBase64 gets a reference to the given string and assigns it to the MtlsCertDataBase64 field.
func (o *CreateGlobalSignAtlasTarget) SetMtlsCertDataBase64(v string) {
	o.MtlsCertDataBase64 = &v
}

// GetMtlsKeyDataBase64 returns the MtlsKeyDataBase64 field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetMtlsKeyDataBase64() string {
	if o == nil || o.MtlsKeyDataBase64 == nil {
		var ret string
		return ret
	}
	return *o.MtlsKeyDataBase64
}

// GetMtlsKeyDataBase64Ok returns a tuple with the MtlsKeyDataBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetMtlsKeyDataBase64Ok() (*string, bool) {
	if o == nil || o.MtlsKeyDataBase64 == nil {
		return nil, false
	}
	return o.MtlsKeyDataBase64, true
}

// HasMtlsKeyDataBase64 returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasMtlsKeyDataBase64() bool {
	if o != nil && o.MtlsKeyDataBase64 != nil {
		return true
	}

	return false
}

// SetMtlsKeyDataBase64 gets a reference to the given string and assigns it to the MtlsKeyDataBase64 field.
func (o *CreateGlobalSignAtlasTarget) SetMtlsKeyDataBase64(v string) {
	o.MtlsKeyDataBase64 = &v
}

// GetName returns the Name field value
func (o *CreateGlobalSignAtlasTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGlobalSignAtlasTarget) SetName(v string) {
	o.Name = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetTimeout() string {
	if o == nil || o.Timeout == nil {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetTimeoutOk() (*string, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *CreateGlobalSignAtlasTarget) SetTimeout(v string) {
	o.Timeout = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateGlobalSignAtlasTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateGlobalSignAtlasTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalSignAtlasTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateGlobalSignAtlasTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateGlobalSignAtlasTarget) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateGlobalSignAtlasTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api-key"] = o.ApiKey
	}
	if true {
		toSerialize["api-secret"] = o.ApiSecret
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.MaxVersions != nil {
		toSerialize["max-versions"] = o.MaxVersions
	}
	if o.MtlsCertDataBase64 != nil {
		toSerialize["mtls-cert-data-base64"] = o.MtlsCertDataBase64
	}
	if o.MtlsKeyDataBase64 != nil {
		toSerialize["mtls-key-data-base64"] = o.MtlsKeyDataBase64
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGlobalSignAtlasTarget struct {
	value *CreateGlobalSignAtlasTarget
	isSet bool
}

func (v NullableCreateGlobalSignAtlasTarget) Get() *CreateGlobalSignAtlasTarget {
	return v.value
}

func (v *NullableCreateGlobalSignAtlasTarget) Set(val *CreateGlobalSignAtlasTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGlobalSignAtlasTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGlobalSignAtlasTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGlobalSignAtlasTarget(val *CreateGlobalSignAtlasTarget) *NullableCreateGlobalSignAtlasTarget {
	return &NullableCreateGlobalSignAtlasTarget{value: val, isSet: true}
}

func (v NullableCreateGlobalSignAtlasTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGlobalSignAtlasTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

