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

// CreateGodaddyTarget struct for CreateGodaddyTarget
type CreateGodaddyTarget struct {
	// Key of the api credentials to the Godaddy account
	ApiKey string `json:"api-key"`
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
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Set the maximum number of versions, limited by the account settings defaults.
	MaxVersions *string `json:"max-versions,omitempty"`
	// Target name
	Name string `json:"name"`
	// Secret of the api credentials to the Godaddy account
	Secret string `json:"secret"`
	// Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h.
	Timeout *string `json:"timeout,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateGodaddyTarget instantiates a new CreateGodaddyTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGodaddyTarget(apiKey string, imapFqdn string, imapPassword string, imapUsername string, name string, secret string, ) *CreateGodaddyTarget {
	this := CreateGodaddyTarget{}
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

// NewCreateGodaddyTargetWithDefaults instantiates a new CreateGodaddyTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGodaddyTargetWithDefaults() *CreateGodaddyTarget {
	this := CreateGodaddyTarget{}
	var imapPort string = "993"
	this.ImapPort = &imapPort
	var json bool = false
	this.Json = &json
	var timeout string = "5m"
	this.Timeout = &timeout
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *CreateGodaddyTarget) GetApiKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetApiKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *CreateGodaddyTarget) SetApiKey(v string) {
	o.ApiKey = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateGodaddyTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGodaddyTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateGodaddyTarget) SetDescription(v string) {
	o.Description = &v
}

// GetImapFqdn returns the ImapFqdn field value
func (o *CreateGodaddyTarget) GetImapFqdn() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ImapFqdn
}

// GetImapFqdnOk returns a tuple with the ImapFqdn field value
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetImapFqdnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImapFqdn, true
}

// SetImapFqdn sets field value
func (o *CreateGodaddyTarget) SetImapFqdn(v string) {
	o.ImapFqdn = v
}

// GetImapPassword returns the ImapPassword field value
func (o *CreateGodaddyTarget) GetImapPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ImapPassword
}

// GetImapPasswordOk returns a tuple with the ImapPassword field value
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetImapPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImapPassword, true
}

// SetImapPassword sets field value
func (o *CreateGodaddyTarget) SetImapPassword(v string) {
	o.ImapPassword = v
}

// GetImapPort returns the ImapPort field value if set, zero value otherwise.
func (o *CreateGodaddyTarget) GetImapPort() string {
	if o == nil || o.ImapPort == nil {
		var ret string
		return ret
	}
	return *o.ImapPort
}

// GetImapPortOk returns a tuple with the ImapPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetImapPortOk() (*string, bool) {
	if o == nil || o.ImapPort == nil {
		return nil, false
	}
	return o.ImapPort, true
}

// HasImapPort returns a boolean if a field has been set.
func (o *CreateGodaddyTarget) HasImapPort() bool {
	if o != nil && o.ImapPort != nil {
		return true
	}

	return false
}

// SetImapPort gets a reference to the given string and assigns it to the ImapPort field.
func (o *CreateGodaddyTarget) SetImapPort(v string) {
	o.ImapPort = &v
}

// GetImapUsername returns the ImapUsername field value
func (o *CreateGodaddyTarget) GetImapUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ImapUsername
}

// GetImapUsernameOk returns a tuple with the ImapUsername field value
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetImapUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImapUsername, true
}

// SetImapUsername sets field value
func (o *CreateGodaddyTarget) SetImapUsername(v string) {
	o.ImapUsername = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateGodaddyTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateGodaddyTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateGodaddyTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateGodaddyTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateGodaddyTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateGodaddyTarget) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *CreateGodaddyTarget) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *CreateGodaddyTarget) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *CreateGodaddyTarget) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetName returns the Name field value
func (o *CreateGodaddyTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGodaddyTarget) SetName(v string) {
	o.Name = v
}

// GetSecret returns the Secret field value
func (o *CreateGodaddyTarget) GetSecret() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateGodaddyTarget) SetSecret(v string) {
	o.Secret = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *CreateGodaddyTarget) GetTimeout() string {
	if o == nil || o.Timeout == nil {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetTimeoutOk() (*string, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *CreateGodaddyTarget) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *CreateGodaddyTarget) SetTimeout(v string) {
	o.Timeout = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateGodaddyTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateGodaddyTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateGodaddyTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateGodaddyTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGodaddyTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateGodaddyTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateGodaddyTarget) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateGodaddyTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api-key"] = o.ApiKey
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
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.MaxVersions != nil {
		toSerialize["max-versions"] = o.MaxVersions
	}
	if true {
		toSerialize["name"] = o.Name
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
	return json.Marshal(toSerialize)
}

type NullableCreateGodaddyTarget struct {
	value *CreateGodaddyTarget
	isSet bool
}

func (v NullableCreateGodaddyTarget) Get() *CreateGodaddyTarget {
	return v.value
}

func (v *NullableCreateGodaddyTarget) Set(val *CreateGodaddyTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGodaddyTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGodaddyTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGodaddyTarget(val *CreateGodaddyTarget) *NullableCreateGodaddyTarget {
	return &NullableCreateGodaddyTarget{value: val, isSet: true}
}

func (v NullableCreateGodaddyTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGodaddyTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


