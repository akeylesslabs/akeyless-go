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

// CreateGithubTarget struct for CreateGithubTarget
type CreateGithubTarget struct {
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	// Github app id
	GithubAppId *int64 `json:"github-app-id,omitempty"`
	// App private key
	GithubAppPrivateKey *string `json:"github-app-private-key,omitempty"`
	// Base URL
	GithubBaseUrl *string `json:"github-base-url,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateGithubTarget instantiates a new CreateGithubTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGithubTarget(name string, ) *CreateGithubTarget {
	this := CreateGithubTarget{}
	this.Name = name
	return &this
}

// NewCreateGithubTargetWithDefaults instantiates a new CreateGithubTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGithubTargetWithDefaults() *CreateGithubTarget {
	this := CreateGithubTarget{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateGithubTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGithubTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateGithubTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateGithubTarget) SetComment(v string) {
	o.Comment = &v
}

// GetGithubAppId returns the GithubAppId field value if set, zero value otherwise.
func (o *CreateGithubTarget) GetGithubAppId() int64 {
	if o == nil || o.GithubAppId == nil {
		var ret int64
		return ret
	}
	return *o.GithubAppId
}

// GetGithubAppIdOk returns a tuple with the GithubAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGithubTarget) GetGithubAppIdOk() (*int64, bool) {
	if o == nil || o.GithubAppId == nil {
		return nil, false
	}
	return o.GithubAppId, true
}

// HasGithubAppId returns a boolean if a field has been set.
func (o *CreateGithubTarget) HasGithubAppId() bool {
	if o != nil && o.GithubAppId != nil {
		return true
	}

	return false
}

// SetGithubAppId gets a reference to the given int64 and assigns it to the GithubAppId field.
func (o *CreateGithubTarget) SetGithubAppId(v int64) {
	o.GithubAppId = &v
}

// GetGithubAppPrivateKey returns the GithubAppPrivateKey field value if set, zero value otherwise.
func (o *CreateGithubTarget) GetGithubAppPrivateKey() string {
	if o == nil || o.GithubAppPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.GithubAppPrivateKey
}

// GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGithubTarget) GetGithubAppPrivateKeyOk() (*string, bool) {
	if o == nil || o.GithubAppPrivateKey == nil {
		return nil, false
	}
	return o.GithubAppPrivateKey, true
}

// HasGithubAppPrivateKey returns a boolean if a field has been set.
func (o *CreateGithubTarget) HasGithubAppPrivateKey() bool {
	if o != nil && o.GithubAppPrivateKey != nil {
		return true
	}

	return false
}

// SetGithubAppPrivateKey gets a reference to the given string and assigns it to the GithubAppPrivateKey field.
func (o *CreateGithubTarget) SetGithubAppPrivateKey(v string) {
	o.GithubAppPrivateKey = &v
}

// GetGithubBaseUrl returns the GithubBaseUrl field value if set, zero value otherwise.
func (o *CreateGithubTarget) GetGithubBaseUrl() string {
	if o == nil || o.GithubBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.GithubBaseUrl
}

// GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGithubTarget) GetGithubBaseUrlOk() (*string, bool) {
	if o == nil || o.GithubBaseUrl == nil {
		return nil, false
	}
	return o.GithubBaseUrl, true
}

// HasGithubBaseUrl returns a boolean if a field has been set.
func (o *CreateGithubTarget) HasGithubBaseUrl() bool {
	if o != nil && o.GithubBaseUrl != nil {
		return true
	}

	return false
}

// SetGithubBaseUrl gets a reference to the given string and assigns it to the GithubBaseUrl field.
func (o *CreateGithubTarget) SetGithubBaseUrl(v string) {
	o.GithubBaseUrl = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateGithubTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGithubTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateGithubTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateGithubTarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *CreateGithubTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGithubTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGithubTarget) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateGithubTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGithubTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateGithubTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateGithubTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateGithubTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGithubTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateGithubTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateGithubTarget) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateGithubTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.GithubAppId != nil {
		toSerialize["github-app-id"] = o.GithubAppId
	}
	if o.GithubAppPrivateKey != nil {
		toSerialize["github-app-private-key"] = o.GithubAppPrivateKey
	}
	if o.GithubBaseUrl != nil {
		toSerialize["github-base-url"] = o.GithubBaseUrl
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGithubTarget struct {
	value *CreateGithubTarget
	isSet bool
}

func (v NullableCreateGithubTarget) Get() *CreateGithubTarget {
	return v.value
}

func (v *NullableCreateGithubTarget) Set(val *CreateGithubTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGithubTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGithubTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGithubTarget(val *CreateGithubTarget) *NullableCreateGithubTarget {
	return &NullableCreateGithubTarget{value: val, isSet: true}
}

func (v NullableCreateGithubTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGithubTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

