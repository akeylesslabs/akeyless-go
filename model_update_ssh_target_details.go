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

// UpdateSSHTargetDetails struct for UpdateSSHTargetDetails
type UpdateSSHTargetDetails struct {
	Host *string `json:"host,omitempty"`
	// Target name
	Name string `json:"name"`
	// Whether to create a new version of not
	NewVersion *bool `json:"new-version,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	Port *string `json:"port,omitempty"`
	PrivateKey *string `json:"private-key,omitempty"`
	PrivateKeyPassword *string `json:"private-key-password,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	ProtectionKey *string `json:"protection_key,omitempty"`
	SshPassword *string `json:"ssh-password,omitempty"`
	SshUsername *string `json:"ssh-username,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewUpdateSSHTargetDetails instantiates a new UpdateSSHTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSSHTargetDetails(name string, ) *UpdateSSHTargetDetails {
	this := UpdateSSHTargetDetails{}
	this.Name = name
	var newVersion bool = false
	this.NewVersion = &newVersion
	return &this
}

// NewUpdateSSHTargetDetailsWithDefaults instantiates a new UpdateSSHTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSSHTargetDetailsWithDefaults() *UpdateSSHTargetDetails {
	this := UpdateSSHTargetDetails{}
	var newVersion bool = false
	this.NewVersion = &newVersion
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UpdateSSHTargetDetails) SetHost(v string) {
	o.Host = &v
}

// GetName returns the Name field value
func (o *UpdateSSHTargetDetails) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateSSHTargetDetails) SetName(v string) {
	o.Name = v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetNewVersion() bool {
	if o == nil || o.NewVersion == nil {
		var ret bool
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetNewVersionOk() (*bool, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given bool and assigns it to the NewVersion field.
func (o *UpdateSSHTargetDetails) SetNewVersion(v bool) {
	o.NewVersion = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateSSHTargetDetails) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *UpdateSSHTargetDetails) SetPort(v string) {
	o.Port = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *UpdateSSHTargetDetails) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPrivateKeyPassword returns the PrivateKeyPassword field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetPrivateKeyPassword() string {
	if o == nil || o.PrivateKeyPassword == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyPassword
}

// GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetPrivateKeyPasswordOk() (*string, bool) {
	if o == nil || o.PrivateKeyPassword == nil {
		return nil, false
	}
	return o.PrivateKeyPassword, true
}

// HasPrivateKeyPassword returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasPrivateKeyPassword() bool {
	if o != nil && o.PrivateKeyPassword != nil {
		return true
	}

	return false
}

// SetPrivateKeyPassword gets a reference to the given string and assigns it to the PrivateKeyPassword field.
func (o *UpdateSSHTargetDetails) SetPrivateKeyPassword(v string) {
	o.PrivateKeyPassword = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetProtectionKey() string {
	if o == nil || o.ProtectionKey == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetProtectionKeyOk() (*string, bool) {
	if o == nil || o.ProtectionKey == nil {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasProtectionKey() bool {
	if o != nil && o.ProtectionKey != nil {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *UpdateSSHTargetDetails) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetSshPassword() string {
	if o == nil || o.SshPassword == nil {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetSshPasswordOk() (*string, bool) {
	if o == nil || o.SshPassword == nil {
		return nil, false
	}
	return o.SshPassword, true
}

// HasSshPassword returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasSshPassword() bool {
	if o != nil && o.SshPassword != nil {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *UpdateSSHTargetDetails) SetSshPassword(v string) {
	o.SshPassword = &v
}

// GetSshUsername returns the SshUsername field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetSshUsername() string {
	if o == nil || o.SshUsername == nil {
		var ret string
		return ret
	}
	return *o.SshUsername
}

// GetSshUsernameOk returns a tuple with the SshUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetSshUsernameOk() (*string, bool) {
	if o == nil || o.SshUsername == nil {
		return nil, false
	}
	return o.SshUsername, true
}

// HasSshUsername returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasSshUsername() bool {
	if o != nil && o.SshUsername != nil {
		return true
	}

	return false
}

// SetSshUsername gets a reference to the given string and assigns it to the SshUsername field.
func (o *UpdateSSHTargetDetails) SetSshUsername(v string) {
	o.SshUsername = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateSSHTargetDetails) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateSSHTargetDetails) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateSSHTargetDetails) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHTargetDetails) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateSSHTargetDetails) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateSSHTargetDetails) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateSSHTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewVersion != nil {
		toSerialize["new-version"] = o.NewVersion
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.PrivateKey != nil {
		toSerialize["private-key"] = o.PrivateKey
	}
	if o.PrivateKeyPassword != nil {
		toSerialize["private-key-password"] = o.PrivateKeyPassword
	}
	if o.ProtectionKey != nil {
		toSerialize["protection_key"] = o.ProtectionKey
	}
	if o.SshPassword != nil {
		toSerialize["ssh-password"] = o.SshPassword
	}
	if o.SshUsername != nil {
		toSerialize["ssh-username"] = o.SshUsername
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

type NullableUpdateSSHTargetDetails struct {
	value *UpdateSSHTargetDetails
	isSet bool
}

func (v NullableUpdateSSHTargetDetails) Get() *UpdateSSHTargetDetails {
	return v.value
}

func (v *NullableUpdateSSHTargetDetails) Set(val *UpdateSSHTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSSHTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSSHTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSSHTargetDetails(val *UpdateSSHTargetDetails) *NullableUpdateSSHTargetDetails {
	return &NullableUpdateSSHTargetDetails{value: val, isSet: true}
}

func (v NullableUpdateSSHTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSSHTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


