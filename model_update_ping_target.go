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

// UpdatePingTarget struct for UpdatePingTarget
type UpdatePingTarget struct {
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
	// Whether to keep previous version [true/false]. If not set, use default according to account settings
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Set the maximum number of versions, limited by the account settings defaults.
	MaxVersions *string `json:"max-versions,omitempty"`
	// Target name
	Name string `json:"name"`
	// New target name
	NewName *string `json:"new-name,omitempty"`
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
	// Deprecated
	UpdateVersion *bool `json:"update-version,omitempty"`
}

// NewUpdatePingTarget instantiates a new UpdatePingTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePingTarget(name string, ) *UpdatePingTarget {
	this := UpdatePingTarget{}
	var administrativePort string = "9999"
	this.AdministrativePort = &administrativePort
	var authorizationPort string = "9031"
	this.AuthorizationPort = &authorizationPort
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewUpdatePingTargetWithDefaults instantiates a new UpdatePingTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePingTargetWithDefaults() *UpdatePingTarget {
	this := UpdatePingTarget{}
	var administrativePort string = "9999"
	this.AdministrativePort = &administrativePort
	var authorizationPort string = "9031"
	this.AuthorizationPort = &authorizationPort
	var json bool = false
	this.Json = &json
	return &this
}

// GetAdministrativePort returns the AdministrativePort field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetAdministrativePort() string {
	if o == nil || o.AdministrativePort == nil {
		var ret string
		return ret
	}
	return *o.AdministrativePort
}

// GetAdministrativePortOk returns a tuple with the AdministrativePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetAdministrativePortOk() (*string, bool) {
	if o == nil || o.AdministrativePort == nil {
		return nil, false
	}
	return o.AdministrativePort, true
}

// HasAdministrativePort returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasAdministrativePort() bool {
	if o != nil && o.AdministrativePort != nil {
		return true
	}

	return false
}

// SetAdministrativePort gets a reference to the given string and assigns it to the AdministrativePort field.
func (o *UpdatePingTarget) SetAdministrativePort(v string) {
	o.AdministrativePort = &v
}

// GetAuthorizationPort returns the AuthorizationPort field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetAuthorizationPort() string {
	if o == nil || o.AuthorizationPort == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationPort
}

// GetAuthorizationPortOk returns a tuple with the AuthorizationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetAuthorizationPortOk() (*string, bool) {
	if o == nil || o.AuthorizationPort == nil {
		return nil, false
	}
	return o.AuthorizationPort, true
}

// HasAuthorizationPort returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasAuthorizationPort() bool {
	if o != nil && o.AuthorizationPort != nil {
		return true
	}

	return false
}

// SetAuthorizationPort gets a reference to the given string and assigns it to the AuthorizationPort field.
func (o *UpdatePingTarget) SetAuthorizationPort(v string) {
	o.AuthorizationPort = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdatePingTarget) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdatePingTarget) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdatePingTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdatePingTarget) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdatePingTarget) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *UpdatePingTarget) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetName returns the Name field value
func (o *UpdatePingTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdatePingTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdatePingTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdatePingTarget) SetPassword(v string) {
	o.Password = &v
}

// GetPingUrl returns the PingUrl field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetPingUrl() string {
	if o == nil || o.PingUrl == nil {
		var ret string
		return ret
	}
	return *o.PingUrl
}

// GetPingUrlOk returns a tuple with the PingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetPingUrlOk() (*string, bool) {
	if o == nil || o.PingUrl == nil {
		return nil, false
	}
	return o.PingUrl, true
}

// HasPingUrl returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasPingUrl() bool {
	if o != nil && o.PingUrl != nil {
		return true
	}

	return false
}

// SetPingUrl gets a reference to the given string and assigns it to the PingUrl field.
func (o *UpdatePingTarget) SetPingUrl(v string) {
	o.PingUrl = &v
}

// GetPrivilegedUser returns the PrivilegedUser field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetPrivilegedUser() string {
	if o == nil || o.PrivilegedUser == nil {
		var ret string
		return ret
	}
	return *o.PrivilegedUser
}

// GetPrivilegedUserOk returns a tuple with the PrivilegedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetPrivilegedUserOk() (*string, bool) {
	if o == nil || o.PrivilegedUser == nil {
		return nil, false
	}
	return o.PrivilegedUser, true
}

// HasPrivilegedUser returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasPrivilegedUser() bool {
	if o != nil && o.PrivilegedUser != nil {
		return true
	}

	return false
}

// SetPrivilegedUser gets a reference to the given string and assigns it to the PrivilegedUser field.
func (o *UpdatePingTarget) SetPrivilegedUser(v string) {
	o.PrivilegedUser = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdatePingTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdatePingTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdatePingTarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePingTarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdatePingTarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdatePingTarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

func (o UpdatePingTarget) MarshalJSON() ([]byte, error) {
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
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.MaxVersions != nil {
		toSerialize["max-versions"] = o.MaxVersions
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
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
	if o.UpdateVersion != nil {
		toSerialize["update-version"] = o.UpdateVersion
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePingTarget struct {
	value *UpdatePingTarget
	isSet bool
}

func (v NullableUpdatePingTarget) Get() *UpdatePingTarget {
	return v.value
}

func (v *NullableUpdatePingTarget) Set(val *UpdatePingTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePingTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePingTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePingTarget(val *UpdatePingTarget) *NullableUpdatePingTarget {
	return &NullableUpdatePingTarget{value: val, isSet: true}
}

func (v NullableUpdatePingTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePingTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


