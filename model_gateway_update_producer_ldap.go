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

// GatewayUpdateProducerLdap gatewayUpdateProducerLdap is a command that updates ldap producer
type GatewayUpdateProducerLdap struct {
	// Bind DN
	BindDn *string `json:"bind-dn,omitempty"`
	// Bind DN Password
	BindDnPassword *string `json:"bind-dn-password,omitempty"`
	// Fixed user
	ExternalUsername *string `json:"external-username,omitempty"`
	// CA Certificate File Content
	LdapCaCert *string `json:"ldap-ca-cert,omitempty"`
	// LDAP Server URL
	LdapUrl *string `json:"ldap-url,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// List of the tags attached to this secret
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Token expiration
	TokenExpiration *string `json:"token-expiration,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User Attribute
	UserAttribute *string `json:"user-attribute,omitempty"`
	// User DN
	UserDn *string `json:"user-dn,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewGatewayUpdateProducerLdap instantiates a new GatewayUpdateProducerLdap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerLdap(name string, ) *GatewayUpdateProducerLdap {
	this := GatewayUpdateProducerLdap{}
	var externalUsername string = "false"
	this.ExternalUsername = &externalUsername
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerLdapWithDefaults instantiates a new GatewayUpdateProducerLdap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerLdapWithDefaults() *GatewayUpdateProducerLdap {
	this := GatewayUpdateProducerLdap{}
	var externalUsername string = "false"
	this.ExternalUsername = &externalUsername
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetBindDn returns the BindDn field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetBindDn() string {
	if o == nil || o.BindDn == nil {
		var ret string
		return ret
	}
	return *o.BindDn
}

// GetBindDnOk returns a tuple with the BindDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetBindDnOk() (*string, bool) {
	if o == nil || o.BindDn == nil {
		return nil, false
	}
	return o.BindDn, true
}

// HasBindDn returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasBindDn() bool {
	if o != nil && o.BindDn != nil {
		return true
	}

	return false
}

// SetBindDn gets a reference to the given string and assigns it to the BindDn field.
func (o *GatewayUpdateProducerLdap) SetBindDn(v string) {
	o.BindDn = &v
}

// GetBindDnPassword returns the BindDnPassword field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetBindDnPassword() string {
	if o == nil || o.BindDnPassword == nil {
		var ret string
		return ret
	}
	return *o.BindDnPassword
}

// GetBindDnPasswordOk returns a tuple with the BindDnPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetBindDnPasswordOk() (*string, bool) {
	if o == nil || o.BindDnPassword == nil {
		return nil, false
	}
	return o.BindDnPassword, true
}

// HasBindDnPassword returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasBindDnPassword() bool {
	if o != nil && o.BindDnPassword != nil {
		return true
	}

	return false
}

// SetBindDnPassword gets a reference to the given string and assigns it to the BindDnPassword field.
func (o *GatewayUpdateProducerLdap) SetBindDnPassword(v string) {
	o.BindDnPassword = &v
}

// GetExternalUsername returns the ExternalUsername field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetExternalUsername() string {
	if o == nil || o.ExternalUsername == nil {
		var ret string
		return ret
	}
	return *o.ExternalUsername
}

// GetExternalUsernameOk returns a tuple with the ExternalUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetExternalUsernameOk() (*string, bool) {
	if o == nil || o.ExternalUsername == nil {
		return nil, false
	}
	return o.ExternalUsername, true
}

// HasExternalUsername returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasExternalUsername() bool {
	if o != nil && o.ExternalUsername != nil {
		return true
	}

	return false
}

// SetExternalUsername gets a reference to the given string and assigns it to the ExternalUsername field.
func (o *GatewayUpdateProducerLdap) SetExternalUsername(v string) {
	o.ExternalUsername = &v
}

// GetLdapCaCert returns the LdapCaCert field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetLdapCaCert() string {
	if o == nil || o.LdapCaCert == nil {
		var ret string
		return ret
	}
	return *o.LdapCaCert
}

// GetLdapCaCertOk returns a tuple with the LdapCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetLdapCaCertOk() (*string, bool) {
	if o == nil || o.LdapCaCert == nil {
		return nil, false
	}
	return o.LdapCaCert, true
}

// HasLdapCaCert returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasLdapCaCert() bool {
	if o != nil && o.LdapCaCert != nil {
		return true
	}

	return false
}

// SetLdapCaCert gets a reference to the given string and assigns it to the LdapCaCert field.
func (o *GatewayUpdateProducerLdap) SetLdapCaCert(v string) {
	o.LdapCaCert = &v
}

// GetLdapUrl returns the LdapUrl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetLdapUrl() string {
	if o == nil || o.LdapUrl == nil {
		var ret string
		return ret
	}
	return *o.LdapUrl
}

// GetLdapUrlOk returns a tuple with the LdapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetLdapUrlOk() (*string, bool) {
	if o == nil || o.LdapUrl == nil {
		return nil, false
	}
	return o.LdapUrl, true
}

// HasLdapUrl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasLdapUrl() bool {
	if o != nil && o.LdapUrl != nil {
		return true
	}

	return false
}

// SetLdapUrl gets a reference to the given string and assigns it to the LdapUrl field.
func (o *GatewayUpdateProducerLdap) SetLdapUrl(v string) {
	o.LdapUrl = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerLdap) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerLdap) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerLdap) SetNewName(v string) {
	o.NewName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayUpdateProducerLdap) SetPassword(v string) {
	o.Password = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerLdap) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerLdap) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerLdap) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerLdap) SetToken(v string) {
	o.Token = &v
}

// GetTokenExpiration returns the TokenExpiration field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetTokenExpiration() string {
	if o == nil || o.TokenExpiration == nil {
		var ret string
		return ret
	}
	return *o.TokenExpiration
}

// GetTokenExpirationOk returns a tuple with the TokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetTokenExpirationOk() (*string, bool) {
	if o == nil || o.TokenExpiration == nil {
		return nil, false
	}
	return o.TokenExpiration, true
}

// HasTokenExpiration returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasTokenExpiration() bool {
	if o != nil && o.TokenExpiration != nil {
		return true
	}

	return false
}

// SetTokenExpiration gets a reference to the given string and assigns it to the TokenExpiration field.
func (o *GatewayUpdateProducerLdap) SetTokenExpiration(v string) {
	o.TokenExpiration = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerLdap) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserAttribute returns the UserAttribute field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetUserAttribute() string {
	if o == nil || o.UserAttribute == nil {
		var ret string
		return ret
	}
	return *o.UserAttribute
}

// GetUserAttributeOk returns a tuple with the UserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetUserAttributeOk() (*string, bool) {
	if o == nil || o.UserAttribute == nil {
		return nil, false
	}
	return o.UserAttribute, true
}

// HasUserAttribute returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasUserAttribute() bool {
	if o != nil && o.UserAttribute != nil {
		return true
	}

	return false
}

// SetUserAttribute gets a reference to the given string and assigns it to the UserAttribute field.
func (o *GatewayUpdateProducerLdap) SetUserAttribute(v string) {
	o.UserAttribute = &v
}

// GetUserDn returns the UserDn field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetUserDn() string {
	if o == nil || o.UserDn == nil {
		var ret string
		return ret
	}
	return *o.UserDn
}

// GetUserDnOk returns a tuple with the UserDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetUserDnOk() (*string, bool) {
	if o == nil || o.UserDn == nil {
		return nil, false
	}
	return o.UserDn, true
}

// HasUserDn returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasUserDn() bool {
	if o != nil && o.UserDn != nil {
		return true
	}

	return false
}

// SetUserDn gets a reference to the given string and assigns it to the UserDn field.
func (o *GatewayUpdateProducerLdap) SetUserDn(v string) {
	o.UserDn = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerLdap) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayUpdateProducerLdap) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerLdap) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayUpdateProducerLdap) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayUpdateProducerLdap) SetUsername(v string) {
	o.Username = &v
}

func (o GatewayUpdateProducerLdap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BindDn != nil {
		toSerialize["bind-dn"] = o.BindDn
	}
	if o.BindDnPassword != nil {
		toSerialize["bind-dn-password"] = o.BindDnPassword
	}
	if o.ExternalUsername != nil {
		toSerialize["external-username"] = o.ExternalUsername
	}
	if o.LdapCaCert != nil {
		toSerialize["ldap-ca-cert"] = o.LdapCaCert
	}
	if o.LdapUrl != nil {
		toSerialize["ldap-url"] = o.LdapUrl
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
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetName != nil {
		toSerialize["target-name"] = o.TargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.TokenExpiration != nil {
		toSerialize["token-expiration"] = o.TokenExpiration
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UserAttribute != nil {
		toSerialize["user-attribute"] = o.UserAttribute
	}
	if o.UserDn != nil {
		toSerialize["user-dn"] = o.UserDn
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateProducerLdap struct {
	value *GatewayUpdateProducerLdap
	isSet bool
}

func (v NullableGatewayUpdateProducerLdap) Get() *GatewayUpdateProducerLdap {
	return v.value
}

func (v *NullableGatewayUpdateProducerLdap) Set(val *GatewayUpdateProducerLdap) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerLdap) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerLdap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerLdap(val *GatewayUpdateProducerLdap) *NullableGatewayUpdateProducerLdap {
	return &NullableGatewayUpdateProducerLdap{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerLdap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerLdap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

