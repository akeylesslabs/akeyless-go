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

// CreateDFCKey struct for CreateDFCKey
type CreateDFCKey struct {
	// DFCKey type; options: [AES128GCM, AES256GCM, AES128SIV, AES256SIV, AES128CBC, AES256CBC, RSA1024, RSA2048, RSA3072, RSA4096]
	Alg string `json:"alg"`
	// The customer fragment ID that will be used to create the DFC key (if empty, the key will be created independently of a customer fragment)
	CustomerFrgId *string `json:"customer-frg-id,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Deprecated - use description
	Metadata *string `json:"metadata,omitempty"`
	// DFCKey name
	Name string `json:"name"`
	// The number of fragments that the item will be split into (not includes customer fragment)
	SplitLevel *int64 `json:"split-level,omitempty"`
	// List of the tags attached to this DFC key
	Tag *[]string `json:"tag,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateDFCKey instantiates a new CreateDFCKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDFCKey(alg string, name string, ) *CreateDFCKey {
	this := CreateDFCKey{}
	this.Alg = alg
	var json bool = false
	this.Json = &json
	this.Name = name
	var splitLevel int64 = 3
	this.SplitLevel = &splitLevel
	return &this
}

// NewCreateDFCKeyWithDefaults instantiates a new CreateDFCKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDFCKeyWithDefaults() *CreateDFCKey {
	this := CreateDFCKey{}
	var json bool = false
	this.Json = &json
	var splitLevel int64 = 3
	this.SplitLevel = &splitLevel
	return &this
}

// GetAlg returns the Alg field value
func (o *CreateDFCKey) GetAlg() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetAlgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *CreateDFCKey) SetAlg(v string) {
	o.Alg = v
}

// GetCustomerFrgId returns the CustomerFrgId field value if set, zero value otherwise.
func (o *CreateDFCKey) GetCustomerFrgId() string {
	if o == nil || o.CustomerFrgId == nil {
		var ret string
		return ret
	}
	return *o.CustomerFrgId
}

// GetCustomerFrgIdOk returns a tuple with the CustomerFrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetCustomerFrgIdOk() (*string, bool) {
	if o == nil || o.CustomerFrgId == nil {
		return nil, false
	}
	return o.CustomerFrgId, true
}

// HasCustomerFrgId returns a boolean if a field has been set.
func (o *CreateDFCKey) HasCustomerFrgId() bool {
	if o != nil && o.CustomerFrgId != nil {
		return true
	}

	return false
}

// SetCustomerFrgId gets a reference to the given string and assigns it to the CustomerFrgId field.
func (o *CreateDFCKey) SetCustomerFrgId(v string) {
	o.CustomerFrgId = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *CreateDFCKey) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *CreateDFCKey) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *CreateDFCKey) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDFCKey) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDFCKey) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDFCKey) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateDFCKey) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateDFCKey) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateDFCKey) SetJson(v bool) {
	o.Json = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateDFCKey) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateDFCKey) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreateDFCKey) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *CreateDFCKey) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDFCKey) SetName(v string) {
	o.Name = v
}

// GetSplitLevel returns the SplitLevel field value if set, zero value otherwise.
func (o *CreateDFCKey) GetSplitLevel() int64 {
	if o == nil || o.SplitLevel == nil {
		var ret int64
		return ret
	}
	return *o.SplitLevel
}

// GetSplitLevelOk returns a tuple with the SplitLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetSplitLevelOk() (*int64, bool) {
	if o == nil || o.SplitLevel == nil {
		return nil, false
	}
	return o.SplitLevel, true
}

// HasSplitLevel returns a boolean if a field has been set.
func (o *CreateDFCKey) HasSplitLevel() bool {
	if o != nil && o.SplitLevel != nil {
		return true
	}

	return false
}

// SetSplitLevel gets a reference to the given int64 and assigns it to the SplitLevel field.
func (o *CreateDFCKey) SetSplitLevel(v int64) {
	o.SplitLevel = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateDFCKey) GetTag() []string {
	if o == nil || o.Tag == nil {
		var ret []string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetTagOk() (*[]string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateDFCKey) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *CreateDFCKey) SetTag(v []string) {
	o.Tag = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateDFCKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateDFCKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateDFCKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateDFCKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDFCKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateDFCKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateDFCKey) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateDFCKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if o.CustomerFrgId != nil {
		toSerialize["customer-frg-id"] = o.CustomerFrgId
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SplitLevel != nil {
		toSerialize["split-level"] = o.SplitLevel
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDFCKey struct {
	value *CreateDFCKey
	isSet bool
}

func (v NullableCreateDFCKey) Get() *CreateDFCKey {
	return v.value
}

func (v *NullableCreateDFCKey) Set(val *CreateDFCKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDFCKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDFCKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDFCKey(val *CreateDFCKey) *NullableCreateDFCKey {
	return &NullableCreateDFCKey{value: val, isSet: true}
}

func (v NullableCreateDFCKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDFCKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


