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

// UpdateGodaddyTarget struct for UpdateGodaddyTarget
type UpdateGodaddyTarget struct {
	// Key of the api credentials to the Godaddy account
	ApiKey string `json:"api-key"`
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// ImapFQDN of the IMAP service, FQDN or IPv4 address. Must be FQDN if the IMAP is using TLS
	ImapFqdn string `json:"imap-fqdn"`
	// ImapPassword to access the IMAP service
	ImapPassword string `json:"imap-password"`
	// ImapPort of the IMAP service
	ImapPort *string `json:"imap-port,omitempty"`
	// ImapUsername to access the IMAP service
	ImapUsername string `json:"imap-username"`
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
	// Secret of the api credentials to the Godaddy account
	Secret string `json:"secret"`
	// Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h.
	Timeout *string `json:"timeout,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Deprecated
	UpdateVersion *bool `json:"update-version,omitempty"`
}

// NewUpdateGodaddyTarget instantiates a new UpdateGodaddyTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGodaddyTarget(apiKey string, imapFqdn string, imapPassword string, imapUsername string, name string, secret string, ) *UpdateGodaddyTarget {
	this := UpdateGodaddyTarget{}
	this.ApiKey = apiKey
	this.ImapFqdn = imapFqdn
	this.ImapPassword = imapPassword
	var imapPort string = "993"
	this.ImapPort = &imapPort
	this.ImapUsername = imapUsername
	var json bool = false
	this.Json = &json
	this.Name = name
	this.Secret = secret
	var timeout string = "5m"
	this.Timeout = &timeout
	return &this
}

// NewUpdateGodaddyTargetWithDefaults instantiates a new UpdateGodaddyTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGodaddyTargetWithDefaults() *UpdateGodaddyTarget {
	this := UpdateGodaddyTarget{}
	var imapPort string = "993"
	this.ImapPort = &imapPort
	var json bool = false
	this.Json = &json
	var timeout string = "5m"
	this.Timeout = &timeout
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *UpdateGodaddyTarget) GetApiKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetApiKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *UpdateGodaddyTarget) SetApiKey(v string) {
	o.ApiKey = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateGodaddyTarget) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateGodaddyTarget) SetDescription(v string) {
	o.Description = &v
}

// GetImapFqdn returns the ImapFqdn field value
func (o *UpdateGodaddyTarget) GetImapFqdn() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ImapFqdn
}

// GetImapFqdnOk returns a tuple with the ImapFqdn field value
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetImapFqdnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImapFqdn, true
}

// SetImapFqdn sets field value
func (o *UpdateGodaddyTarget) SetImapFqdn(v string) {
	o.ImapFqdn = v
}

// GetImapPassword returns the ImapPassword field value
func (o *UpdateGodaddyTarget) GetImapPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ImapPassword
}

// GetImapPasswordOk returns a tuple with the ImapPassword field value
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetImapPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImapPassword, true
}

// SetImapPassword sets field value
func (o *UpdateGodaddyTarget) SetImapPassword(v string) {
	o.ImapPassword = v
}

// GetImapPort returns the ImapPort field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetImapPort() string {
	if o == nil || o.ImapPort == nil {
		var ret string
		return ret
	}
	return *o.ImapPort
}

// GetImapPortOk returns a tuple with the ImapPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetImapPortOk() (*string, bool) {
	if o == nil || o.ImapPort == nil {
		return nil, false
	}
	return o.ImapPort, true
}

// HasImapPort returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasImapPort() bool {
	if o != nil && o.ImapPort != nil {
		return true
	}

	return false
}

// SetImapPort gets a reference to the given string and assigns it to the ImapPort field.
func (o *UpdateGodaddyTarget) SetImapPort(v string) {
	o.ImapPort = &v
}

// GetImapUsername returns the ImapUsername field value
func (o *UpdateGodaddyTarget) GetImapUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ImapUsername
}

// GetImapUsernameOk returns a tuple with the ImapUsername field value
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetImapUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImapUsername, true
}

// SetImapUsername sets field value
func (o *UpdateGodaddyTarget) SetImapUsername(v string) {
	o.ImapUsername = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateGodaddyTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateGodaddyTarget) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateGodaddyTarget) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *UpdateGodaddyTarget) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetName returns the Name field value
func (o *UpdateGodaddyTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateGodaddyTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateGodaddyTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetSecret returns the Secret field value
func (o *UpdateGodaddyTarget) GetSecret() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *UpdateGodaddyTarget) SetSecret(v string) {
	o.Secret = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetTimeout() string {
	if o == nil || o.Timeout == nil {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetTimeoutOk() (*string, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *UpdateGodaddyTarget) SetTimeout(v string) {
	o.Timeout = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateGodaddyTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateGodaddyTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateGodaddyTarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGodaddyTarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateGodaddyTarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateGodaddyTarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

func (o UpdateGodaddyTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api-key"] = o.ApiKey
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["imap-fqdn"] = o.ImapFqdn
	}
	if true {
		toSerialize["imap-password"] = o.ImapPassword
	}
	if o.ImapPort != nil {
		toSerialize["imap-port"] = o.ImapPort
	}
	if true {
		toSerialize["imap-username"] = o.ImapUsername
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
	if true {
		toSerialize["secret"] = o.Secret
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
	if o.UpdateVersion != nil {
		toSerialize["update-version"] = o.UpdateVersion
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateGodaddyTarget struct {
	value *UpdateGodaddyTarget
	isSet bool
}

func (v NullableUpdateGodaddyTarget) Get() *UpdateGodaddyTarget {
	return v.value
}

func (v *NullableUpdateGodaddyTarget) Set(val *UpdateGodaddyTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGodaddyTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGodaddyTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGodaddyTarget(val *UpdateGodaddyTarget) *NullableUpdateGodaddyTarget {
	return &NullableUpdateGodaddyTarget{value: val, isSet: true}
}

func (v NullableUpdateGodaddyTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGodaddyTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


