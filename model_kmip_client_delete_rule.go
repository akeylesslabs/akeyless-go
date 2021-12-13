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

// KmipClientDeleteRule struct for KmipClientDeleteRule
type KmipClientDeleteRule struct {
	ClientId *string `json:"client-id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Access path
	Path string `json:"path"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewKmipClientDeleteRule instantiates a new KmipClientDeleteRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipClientDeleteRule(path string, ) *KmipClientDeleteRule {
	this := KmipClientDeleteRule{}
	this.Path = path
	return &this
}

// NewKmipClientDeleteRuleWithDefaults instantiates a new KmipClientDeleteRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipClientDeleteRuleWithDefaults() *KmipClientDeleteRule {
	this := KmipClientDeleteRule{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *KmipClientDeleteRule) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipClientDeleteRule) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *KmipClientDeleteRule) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *KmipClientDeleteRule) SetClientId(v string) {
	o.ClientId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KmipClientDeleteRule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipClientDeleteRule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KmipClientDeleteRule) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KmipClientDeleteRule) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *KmipClientDeleteRule) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipClientDeleteRule) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *KmipClientDeleteRule) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *KmipClientDeleteRule) SetPassword(v string) {
	o.Password = &v
}

// GetPath returns the Path field value
func (o *KmipClientDeleteRule) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *KmipClientDeleteRule) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *KmipClientDeleteRule) SetPath(v string) {
	o.Path = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *KmipClientDeleteRule) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipClientDeleteRule) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *KmipClientDeleteRule) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *KmipClientDeleteRule) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *KmipClientDeleteRule) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipClientDeleteRule) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *KmipClientDeleteRule) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *KmipClientDeleteRule) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *KmipClientDeleteRule) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipClientDeleteRule) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *KmipClientDeleteRule) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *KmipClientDeleteRule) SetUsername(v string) {
	o.Username = &v
}

func (o KmipClientDeleteRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client-id"] = o.ClientId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableKmipClientDeleteRule struct {
	value *KmipClientDeleteRule
	isSet bool
}

func (v NullableKmipClientDeleteRule) Get() *KmipClientDeleteRule {
	return v.value
}

func (v *NullableKmipClientDeleteRule) Set(val *KmipClientDeleteRule) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipClientDeleteRule) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipClientDeleteRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipClientDeleteRule(val *KmipClientDeleteRule) *NullableKmipClientDeleteRule {
	return &NullableKmipClientDeleteRule{value: val, isSet: true}
}

func (v NullableKmipClientDeleteRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipClientDeleteRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

