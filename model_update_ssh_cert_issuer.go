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

// UpdateSSHCertIssuer struct for UpdateSSHCertIssuer
type UpdateSSHCertIssuer struct {
	// List of the new tags that will be attached to this item
	AddTag *[]string `json:"add-tag,omitempty"`
	// Users allowed to fetch the certificate, e.g root,ubuntu
	AllowedUsers string `json:"allowed-users"`
	// Signed certificates with extensions, e.g permit-port-forwarding=\\\"\\\"
	Extensions *map[string]string `json:"extensions,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// A metadata about the issuer
	Metadata *string `json:"metadata,omitempty"`
	// SSH certificate issuer name
	Name string `json:"name"`
	// New item name
	NewName *string `json:"new-name,omitempty"`
	// Signed certificates with principal, e.g example_role1,example_role2
	Principals *string `json:"principals,omitempty"`
	// List of the existent tags that will be removed from this item
	RmTag *[]string `json:"rm-tag,omitempty"`
	SecureAccessBastionApi *string `json:"secure-access-bastion-api,omitempty"`
	SecureAccessBastionSsh *string `json:"secure-access-bastion-ssh,omitempty"`
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	SecureAccessSshCredsUser *string `json:"secure-access-ssh-creds-user,omitempty"`
	SecureAccessUseInternalBastion *bool `json:"secure-access-use-internal-bastion,omitempty"`
	// A key to sign the certificate with
	SignerKeyName string `json:"signer-key-name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// he requested Time To Live for the certificate, in seconds
	Ttl int64 `json:"ttl"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateSSHCertIssuer instantiates a new UpdateSSHCertIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSSHCertIssuer(allowedUsers string, name string, signerKeyName string, ttl int64, ) *UpdateSSHCertIssuer {
	this := UpdateSSHCertIssuer{}
	this.AllowedUsers = allowedUsers
	this.Name = name
	this.SignerKeyName = signerKeyName
	this.Ttl = ttl
	return &this
}

// NewUpdateSSHCertIssuerWithDefaults instantiates a new UpdateSSHCertIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSSHCertIssuerWithDefaults() *UpdateSSHCertIssuer {
	this := UpdateSSHCertIssuer{}
	return &this
}

// GetAddTag returns the AddTag field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetAddTag() []string {
	if o == nil || o.AddTag == nil {
		var ret []string
		return ret
	}
	return *o.AddTag
}

// GetAddTagOk returns a tuple with the AddTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetAddTagOk() (*[]string, bool) {
	if o == nil || o.AddTag == nil {
		return nil, false
	}
	return o.AddTag, true
}

// HasAddTag returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasAddTag() bool {
	if o != nil && o.AddTag != nil {
		return true
	}

	return false
}

// SetAddTag gets a reference to the given []string and assigns it to the AddTag field.
func (o *UpdateSSHCertIssuer) SetAddTag(v []string) {
	o.AddTag = &v
}

// GetAllowedUsers returns the AllowedUsers field value
func (o *UpdateSSHCertIssuer) GetAllowedUsers() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AllowedUsers
}

// GetAllowedUsersOk returns a tuple with the AllowedUsers field value
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetAllowedUsersOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AllowedUsers, true
}

// SetAllowedUsers sets field value
func (o *UpdateSSHCertIssuer) SetAllowedUsers(v string) {
	o.AllowedUsers = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetExtensions() map[string]string {
	if o == nil || o.Extensions == nil {
		var ret map[string]string
		return ret
	}
	return *o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetExtensionsOk() (*map[string]string, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]string and assigns it to the Extensions field.
func (o *UpdateSSHCertIssuer) SetExtensions(v map[string]string) {
	o.Extensions = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateSSHCertIssuer) SetJson(v bool) {
	o.Json = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *UpdateSSHCertIssuer) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *UpdateSSHCertIssuer) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateSSHCertIssuer) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateSSHCertIssuer) SetNewName(v string) {
	o.NewName = &v
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetPrincipals() string {
	if o == nil || o.Principals == nil {
		var ret string
		return ret
	}
	return *o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetPrincipalsOk() (*string, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasPrincipals() bool {
	if o != nil && o.Principals != nil {
		return true
	}

	return false
}

// SetPrincipals gets a reference to the given string and assigns it to the Principals field.
func (o *UpdateSSHCertIssuer) SetPrincipals(v string) {
	o.Principals = &v
}

// GetRmTag returns the RmTag field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetRmTag() []string {
	if o == nil || o.RmTag == nil {
		var ret []string
		return ret
	}
	return *o.RmTag
}

// GetRmTagOk returns a tuple with the RmTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetRmTagOk() (*[]string, bool) {
	if o == nil || o.RmTag == nil {
		return nil, false
	}
	return o.RmTag, true
}

// HasRmTag returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasRmTag() bool {
	if o != nil && o.RmTag != nil {
		return true
	}

	return false
}

// SetRmTag gets a reference to the given []string and assigns it to the RmTag field.
func (o *UpdateSSHCertIssuer) SetRmTag(v []string) {
	o.RmTag = &v
}

// GetSecureAccessBastionApi returns the SecureAccessBastionApi field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetSecureAccessBastionApi() string {
	if o == nil || o.SecureAccessBastionApi == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionApi
}

// GetSecureAccessBastionApiOk returns a tuple with the SecureAccessBastionApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetSecureAccessBastionApiOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionApi == nil {
		return nil, false
	}
	return o.SecureAccessBastionApi, true
}

// HasSecureAccessBastionApi returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasSecureAccessBastionApi() bool {
	if o != nil && o.SecureAccessBastionApi != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionApi gets a reference to the given string and assigns it to the SecureAccessBastionApi field.
func (o *UpdateSSHCertIssuer) SetSecureAccessBastionApi(v string) {
	o.SecureAccessBastionApi = &v
}

// GetSecureAccessBastionSsh returns the SecureAccessBastionSsh field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetSecureAccessBastionSsh() string {
	if o == nil || o.SecureAccessBastionSsh == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionSsh
}

// GetSecureAccessBastionSshOk returns a tuple with the SecureAccessBastionSsh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetSecureAccessBastionSshOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionSsh == nil {
		return nil, false
	}
	return o.SecureAccessBastionSsh, true
}

// HasSecureAccessBastionSsh returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasSecureAccessBastionSsh() bool {
	if o != nil && o.SecureAccessBastionSsh != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionSsh gets a reference to the given string and assigns it to the SecureAccessBastionSsh field.
func (o *UpdateSSHCertIssuer) SetSecureAccessBastionSsh(v string) {
	o.SecureAccessBastionSsh = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *UpdateSSHCertIssuer) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *UpdateSSHCertIssuer) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessSshCredsUser returns the SecureAccessSshCredsUser field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetSecureAccessSshCredsUser() string {
	if o == nil || o.SecureAccessSshCredsUser == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessSshCredsUser
}

// GetSecureAccessSshCredsUserOk returns a tuple with the SecureAccessSshCredsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetSecureAccessSshCredsUserOk() (*string, bool) {
	if o == nil || o.SecureAccessSshCredsUser == nil {
		return nil, false
	}
	return o.SecureAccessSshCredsUser, true
}

// HasSecureAccessSshCredsUser returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasSecureAccessSshCredsUser() bool {
	if o != nil && o.SecureAccessSshCredsUser != nil {
		return true
	}

	return false
}

// SetSecureAccessSshCredsUser gets a reference to the given string and assigns it to the SecureAccessSshCredsUser field.
func (o *UpdateSSHCertIssuer) SetSecureAccessSshCredsUser(v string) {
	o.SecureAccessSshCredsUser = &v
}

// GetSecureAccessUseInternalBastion returns the SecureAccessUseInternalBastion field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetSecureAccessUseInternalBastion() bool {
	if o == nil || o.SecureAccessUseInternalBastion == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessUseInternalBastion
}

// GetSecureAccessUseInternalBastionOk returns a tuple with the SecureAccessUseInternalBastion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetSecureAccessUseInternalBastionOk() (*bool, bool) {
	if o == nil || o.SecureAccessUseInternalBastion == nil {
		return nil, false
	}
	return o.SecureAccessUseInternalBastion, true
}

// HasSecureAccessUseInternalBastion returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasSecureAccessUseInternalBastion() bool {
	if o != nil && o.SecureAccessUseInternalBastion != nil {
		return true
	}

	return false
}

// SetSecureAccessUseInternalBastion gets a reference to the given bool and assigns it to the SecureAccessUseInternalBastion field.
func (o *UpdateSSHCertIssuer) SetSecureAccessUseInternalBastion(v bool) {
	o.SecureAccessUseInternalBastion = &v
}

// GetSignerKeyName returns the SignerKeyName field value
func (o *UpdateSSHCertIssuer) GetSignerKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SignerKeyName
}

// GetSignerKeyNameOk returns a tuple with the SignerKeyName field value
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetSignerKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SignerKeyName, true
}

// SetSignerKeyName sets field value
func (o *UpdateSSHCertIssuer) SetSignerKeyName(v string) {
	o.SignerKeyName = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateSSHCertIssuer) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value
func (o *UpdateSSHCertIssuer) GetTtl() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetTtlOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *UpdateSSHCertIssuer) SetTtl(v int64) {
	o.Ttl = v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateSSHCertIssuer) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSSHCertIssuer) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateSSHCertIssuer) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateSSHCertIssuer) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateSSHCertIssuer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddTag != nil {
		toSerialize["add-tag"] = o.AddTag
	}
	if true {
		toSerialize["allowed-users"] = o.AllowedUsers
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
	}
	if o.RmTag != nil {
		toSerialize["rm-tag"] = o.RmTag
	}
	if o.SecureAccessBastionApi != nil {
		toSerialize["secure-access-bastion-api"] = o.SecureAccessBastionApi
	}
	if o.SecureAccessBastionSsh != nil {
		toSerialize["secure-access-bastion-ssh"] = o.SecureAccessBastionSsh
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessHost != nil {
		toSerialize["secure-access-host"] = o.SecureAccessHost
	}
	if o.SecureAccessSshCredsUser != nil {
		toSerialize["secure-access-ssh-creds-user"] = o.SecureAccessSshCredsUser
	}
	if o.SecureAccessUseInternalBastion != nil {
		toSerialize["secure-access-use-internal-bastion"] = o.SecureAccessUseInternalBastion
	}
	if true {
		toSerialize["signer-key-name"] = o.SignerKeyName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["ttl"] = o.Ttl
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSSHCertIssuer struct {
	value *UpdateSSHCertIssuer
	isSet bool
}

func (v NullableUpdateSSHCertIssuer) Get() *UpdateSSHCertIssuer {
	return v.value
}

func (v *NullableUpdateSSHCertIssuer) Set(val *UpdateSSHCertIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSSHCertIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSSHCertIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSSHCertIssuer(val *UpdateSSHCertIssuer) *NullableUpdateSSHCertIssuer {
	return &NullableUpdateSSHCertIssuer{value: val, isSet: true}
}

func (v NullableUpdateSSHCertIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSSHCertIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


