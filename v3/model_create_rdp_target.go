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

// CreateRdpTarget struct for CreateRdpTarget
type CreateRdpTarget struct {
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	Host *string `json:"host,omitempty"`
	// Target name
	Name string `json:"name"`
	Pass *string `json:"pass,omitempty"`
	Port *string `json:"port,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	ProtectionKey *string `json:"protection_key,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	User *string `json:"user,omitempty"`
}

// NewCreateRdpTarget instantiates a new CreateRdpTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRdpTarget(name string, ) *CreateRdpTarget {
	this := CreateRdpTarget{}
	this.Name = name
	return &this
}

// NewCreateRdpTargetWithDefaults instantiates a new CreateRdpTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRdpTargetWithDefaults() *CreateRdpTarget {
	this := CreateRdpTarget{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateRdpTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRdpTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateRdpTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateRdpTarget) SetComment(v string) {
	o.Comment = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CreateRdpTarget) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRdpTarget) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CreateRdpTarget) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CreateRdpTarget) SetHost(v string) {
	o.Host = &v
}

// GetName returns the Name field value
func (o *CreateRdpTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRdpTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRdpTarget) SetName(v string) {
	o.Name = v
}

// GetPass returns the Pass field value if set, zero value otherwise.
func (o *CreateRdpTarget) GetPass() string {
	if o == nil || o.Pass == nil {
		var ret string
		return ret
	}
	return *o.Pass
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRdpTarget) GetPassOk() (*string, bool) {
	if o == nil || o.Pass == nil {
		return nil, false
	}
	return o.Pass, true
}

// HasPass returns a boolean if a field has been set.
func (o *CreateRdpTarget) HasPass() bool {
	if o != nil && o.Pass != nil {
		return true
	}

	return false
}

// SetPass gets a reference to the given string and assigns it to the Pass field.
func (o *CreateRdpTarget) SetPass(v string) {
	o.Pass = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CreateRdpTarget) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRdpTarget) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CreateRdpTarget) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *CreateRdpTarget) SetPort(v string) {
	o.Port = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *CreateRdpTarget) GetProtectionKey() string {
	if o == nil || o.ProtectionKey == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRdpTarget) GetProtectionKeyOk() (*string, bool) {
	if o == nil || o.ProtectionKey == nil {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *CreateRdpTarget) HasProtectionKey() bool {
	if o != nil && o.ProtectionKey != nil {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *CreateRdpTarget) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateRdpTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRdpTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateRdpTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateRdpTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateRdpTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRdpTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateRdpTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateRdpTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateRdpTarget) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRdpTarget) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateRdpTarget) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CreateRdpTarget) SetUser(v string) {
	o.User = &v
}

func (o CreateRdpTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Pass != nil {
		toSerialize["pass"] = o.Pass
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
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
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRdpTarget struct {
	value *CreateRdpTarget
	isSet bool
}

func (v NullableCreateRdpTarget) Get() *CreateRdpTarget {
	return v.value
}

func (v *NullableCreateRdpTarget) Set(val *CreateRdpTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRdpTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRdpTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRdpTarget(val *CreateRdpTarget) *NullableCreateRdpTarget {
	return &NullableCreateRdpTarget{value: val, isSet: true}
}

func (v NullableCreateRdpTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRdpTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


