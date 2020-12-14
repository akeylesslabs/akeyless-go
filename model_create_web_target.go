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

// CreateWebTarget struct for CreateWebTarget
type CreateWebTarget struct {
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	// Target name
	Name string `json:"name"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	ProtectionKey *string `json:"protection_key,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The url
	Url *string `json:"url,omitempty"`
}

// NewCreateWebTarget instantiates a new CreateWebTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebTarget(name string, ) *CreateWebTarget {
	this := CreateWebTarget{}
	this.Name = name
	return &this
}

// NewCreateWebTargetWithDefaults instantiates a new CreateWebTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebTargetWithDefaults() *CreateWebTarget {
	this := CreateWebTarget{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateWebTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateWebTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateWebTarget) SetComment(v string) {
	o.Comment = &v
}

// GetName returns the Name field value
func (o *CreateWebTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWebTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWebTarget) SetName(v string) {
	o.Name = v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *CreateWebTarget) GetProtectionKey() string {
	if o == nil || o.ProtectionKey == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebTarget) GetProtectionKeyOk() (*string, bool) {
	if o == nil || o.ProtectionKey == nil {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *CreateWebTarget) HasProtectionKey() bool {
	if o != nil && o.ProtectionKey != nil {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *CreateWebTarget) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateWebTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateWebTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateWebTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateWebTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateWebTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateWebTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateWebTarget) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebTarget) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateWebTarget) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateWebTarget) SetUrl(v string) {
	o.Url = &v
}

func (o CreateWebTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProtectionKey != nil {
		toSerialize["protection_key"] = o.ProtectionKey
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWebTarget struct {
	value *CreateWebTarget
	isSet bool
}

func (v NullableCreateWebTarget) Get() *CreateWebTarget {
	return v.value
}

func (v *NullableCreateWebTarget) Set(val *CreateWebTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebTarget(val *CreateWebTarget) *NullableCreateWebTarget {
	return &NullableCreateWebTarget{value: val, isSet: true}
}

func (v NullableCreateWebTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


