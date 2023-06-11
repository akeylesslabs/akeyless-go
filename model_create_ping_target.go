/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// CreatePingTarget struct for CreatePingTarget
type CreatePingTarget struct {
	// Ping Federate administrative port
	AdministrativePort *string `json:"administrative-port,omitempty"`
	// Ping Federate authorization port
	AuthorizationPort *string `json:"authorization-port,omitempty"`
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// Ping Federate privileged user password
	Password *string `json:"password,omitempty"`
	// Ping URL
	PingUrl *string `json:"ping-url,omitempty"`
	// Ping Federate privileged user
	PrivilegedUser *string `json:"privileged-user,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreatePingTarget instantiates a new CreatePingTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePingTarget(name string) *CreatePingTarget {
	this := CreatePingTarget{}
	var administrativePort string = "9999"
	this.AdministrativePort = &administrativePort
	var authorizationPort string = "9031"
	this.AuthorizationPort = &authorizationPort
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewCreatePingTargetWithDefaults instantiates a new CreatePingTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePingTargetWithDefaults() *CreatePingTarget {
	this := CreatePingTarget{}
	var administrativePort string = "9999"
	this.AdministrativePort = &administrativePort
	var authorizationPort string = "9031"
	this.AuthorizationPort = &authorizationPort
	var json bool = false
	this.Json = &json
	return &this
}

// GetAdministrativePort returns the AdministrativePort field value if set, zero value otherwise.
func (o *CreatePingTarget) GetAdministrativePort() string {
	if o == nil || o.AdministrativePort == nil {
		var ret string
		return ret
	}
	return *o.AdministrativePort
}

// GetAdministrativePortOk returns a tuple with the AdministrativePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetAdministrativePortOk() (*string, bool) {
	if o == nil || o.AdministrativePort == nil {
		return nil, false
	}
	return o.AdministrativePort, true
}

// HasAdministrativePort returns a boolean if a field has been set.
func (o *CreatePingTarget) HasAdministrativePort() bool {
	if o != nil && o.AdministrativePort != nil {
		return true
	}

	return false
}

// SetAdministrativePort gets a reference to the given string and assigns it to the AdministrativePort field.
func (o *CreatePingTarget) SetAdministrativePort(v string) {
	o.AdministrativePort = &v
}

// GetAuthorizationPort returns the AuthorizationPort field value if set, zero value otherwise.
func (o *CreatePingTarget) GetAuthorizationPort() string {
	if o == nil || o.AuthorizationPort == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationPort
}

// GetAuthorizationPortOk returns a tuple with the AuthorizationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetAuthorizationPortOk() (*string, bool) {
	if o == nil || o.AuthorizationPort == nil {
		return nil, false
	}
	return o.AuthorizationPort, true
}

// HasAuthorizationPort returns a boolean if a field has been set.
func (o *CreatePingTarget) HasAuthorizationPort() bool {
	if o != nil && o.AuthorizationPort != nil {
		return true
	}

	return false
}

// SetAuthorizationPort gets a reference to the given string and assigns it to the AuthorizationPort field.
func (o *CreatePingTarget) SetAuthorizationPort(v string) {
	o.AuthorizationPort = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreatePingTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreatePingTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreatePingTarget) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreatePingTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreatePingTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreatePingTarget) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreatePingTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreatePingTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreatePingTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreatePingTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreatePingTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreatePingTarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *CreatePingTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePingTarget) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreatePingTarget) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreatePingTarget) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreatePingTarget) SetPassword(v string) {
	o.Password = &v
}

// GetPingUrl returns the PingUrl field value if set, zero value otherwise.
func (o *CreatePingTarget) GetPingUrl() string {
	if o == nil || o.PingUrl == nil {
		var ret string
		return ret
	}
	return *o.PingUrl
}

// GetPingUrlOk returns a tuple with the PingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetPingUrlOk() (*string, bool) {
	if o == nil || o.PingUrl == nil {
		return nil, false
	}
	return o.PingUrl, true
}

// HasPingUrl returns a boolean if a field has been set.
func (o *CreatePingTarget) HasPingUrl() bool {
	if o != nil && o.PingUrl != nil {
		return true
	}

	return false
}

// SetPingUrl gets a reference to the given string and assigns it to the PingUrl field.
func (o *CreatePingTarget) SetPingUrl(v string) {
	o.PingUrl = &v
}

// GetPrivilegedUser returns the PrivilegedUser field value if set, zero value otherwise.
func (o *CreatePingTarget) GetPrivilegedUser() string {
	if o == nil || o.PrivilegedUser == nil {
		var ret string
		return ret
	}
	return *o.PrivilegedUser
}

// GetPrivilegedUserOk returns a tuple with the PrivilegedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetPrivilegedUserOk() (*string, bool) {
	if o == nil || o.PrivilegedUser == nil {
		return nil, false
	}
	return o.PrivilegedUser, true
}

// HasPrivilegedUser returns a boolean if a field has been set.
func (o *CreatePingTarget) HasPrivilegedUser() bool {
	if o != nil && o.PrivilegedUser != nil {
		return true
	}

	return false
}

// SetPrivilegedUser gets a reference to the given string and assigns it to the PrivilegedUser field.
func (o *CreatePingTarget) SetPrivilegedUser(v string) {
	o.PrivilegedUser = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreatePingTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreatePingTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreatePingTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreatePingTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePingTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreatePingTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreatePingTarget) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreatePingTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdministrativePort != nil {
		toSerialize["administrative-port"] = o.AdministrativePort
	}
	if o.AuthorizationPort != nil {
		toSerialize["authorization-port"] = o.AuthorizationPort
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
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PingUrl != nil {
		toSerialize["ping-url"] = o.PingUrl
	}
	if o.PrivilegedUser != nil {
		toSerialize["privileged-user"] = o.PrivilegedUser
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePingTarget struct {
	value *CreatePingTarget
	isSet bool
}

func (v NullableCreatePingTarget) Get() *CreatePingTarget {
	return v.value
}

func (v *NullableCreatePingTarget) Set(val *CreatePingTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePingTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePingTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePingTarget(val *CreatePingTarget) *NullableCreatePingTarget {
	return &NullableCreatePingTarget{value: val, isSet: true}
}

func (v NullableCreatePingTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePingTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


