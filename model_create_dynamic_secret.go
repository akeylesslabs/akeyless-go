/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// CreateDynamicSecret struct for CreateDynamicSecret
type CreateDynamicSecret struct {
	// The name of a key that used to encrypt the dynamic secret values (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Metadata about the dynamic secret
	Metadata *string `json:"metadata,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// List of the tags attached to this secret
	Tags *[]string `json:"tags,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateDynamicSecret instantiates a new CreateDynamicSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDynamicSecret(name string, ) *CreateDynamicSecret {
	this := CreateDynamicSecret{}
	var metadata string = "None"
	this.Metadata = &metadata
	this.Name = name
	return &this
}

// NewCreateDynamicSecretWithDefaults instantiates a new CreateDynamicSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDynamicSecretWithDefaults() *CreateDynamicSecret {
	this := CreateDynamicSecret{}
	var metadata string = "None"
	this.Metadata = &metadata
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateDynamicSecret) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDynamicSecret) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateDynamicSecret) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateDynamicSecret) SetKey(v string) {
	o.Key = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateDynamicSecret) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDynamicSecret) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateDynamicSecret) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreateDynamicSecret) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *CreateDynamicSecret) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDynamicSecret) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDynamicSecret) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateDynamicSecret) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDynamicSecret) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateDynamicSecret) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateDynamicSecret) SetTags(v []string) {
	o.Tags = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateDynamicSecret) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDynamicSecret) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateDynamicSecret) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateDynamicSecret) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateDynamicSecret) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDynamicSecret) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateDynamicSecret) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateDynamicSecret) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateDynamicSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
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
	return json.Marshal(toSerialize)
}

type NullableCreateDynamicSecret struct {
	value *CreateDynamicSecret
	isSet bool
}

func (v NullableCreateDynamicSecret) Get() *CreateDynamicSecret {
	return v.value
}

func (v *NullableCreateDynamicSecret) Set(val *CreateDynamicSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDynamicSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDynamicSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDynamicSecret(val *CreateDynamicSecret) *NullableCreateDynamicSecret {
	return &NullableCreateDynamicSecret{value: val, isSet: true}
}

func (v NullableCreateDynamicSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDynamicSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


