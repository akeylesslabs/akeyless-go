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

// DsUpdateLdap dsUpdateLdap is a command that updates ldap dynamic secret
type DsUpdateLdap struct {
	// Bind DN
	BindDn *string `json:"bind-dn,omitempty"`
	// Bind DN Password
	BindDnPassword *string `json:"bind-dn-password,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Externally provided username [true/false]
	ExternalUsername *string `json:"external-username,omitempty"`
	// Group DN which the temporary user should be added
	GroupDn *string `json:"group-dn,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// CA Certificate File Content
	LdapCaCert *string `json:"ldap-ca-cert,omitempty"`
	// LDAP Server URL
	LdapUrl *string `json:"ldap-url,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret new name
	NewName *string `json:"new-name,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Add tags attached to this object
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
}

// NewDsUpdateLdap instantiates a new DsUpdateLdap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDsUpdateLdap(name string, ) *DsUpdateLdap {
	this := DsUpdateLdap{}
	var externalUsername string = "false"
	this.ExternalUsername = &externalUsername
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDsUpdateLdapWithDefaults instantiates a new DsUpdateLdap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDsUpdateLdapWithDefaults() *DsUpdateLdap {
	this := DsUpdateLdap{}
	var externalUsername string = "false"
	this.ExternalUsername = &externalUsername
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetBindDn returns the BindDn field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetBindDn() string {
	if o == nil || o.BindDn == nil {
		var ret string
		return ret
	}
	return *o.BindDn
}

// GetBindDnOk returns a tuple with the BindDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetBindDnOk() (*string, bool) {
	if o == nil || o.BindDn == nil {
		return nil, false
	}
	return o.BindDn, true
}

// HasBindDn returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasBindDn() bool {
	if o != nil && o.BindDn != nil {
		return true
	}

	return false
}

// SetBindDn gets a reference to the given string and assigns it to the BindDn field.
func (o *DsUpdateLdap) SetBindDn(v string) {
	o.BindDn = &v
}

// GetBindDnPassword returns the BindDnPassword field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetBindDnPassword() string {
	if o == nil || o.BindDnPassword == nil {
		var ret string
		return ret
	}
	return *o.BindDnPassword
}

// GetBindDnPasswordOk returns a tuple with the BindDnPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetBindDnPasswordOk() (*string, bool) {
	if o == nil || o.BindDnPassword == nil {
		return nil, false
	}
	return o.BindDnPassword, true
}

// HasBindDnPassword returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasBindDnPassword() bool {
	if o != nil && o.BindDnPassword != nil {
		return true
	}

	return false
}

// SetBindDnPassword gets a reference to the given string and assigns it to the BindDnPassword field.
func (o *DsUpdateLdap) SetBindDnPassword(v string) {
	o.BindDnPassword = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DsUpdateLdap) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetExternalUsername returns the ExternalUsername field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetExternalUsername() string {
	if o == nil || o.ExternalUsername == nil {
		var ret string
		return ret
	}
	return *o.ExternalUsername
}

// GetExternalUsernameOk returns a tuple with the ExternalUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetExternalUsernameOk() (*string, bool) {
	if o == nil || o.ExternalUsername == nil {
		return nil, false
	}
	return o.ExternalUsername, true
}

// HasExternalUsername returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasExternalUsername() bool {
	if o != nil && o.ExternalUsername != nil {
		return true
	}

	return false
}

// SetExternalUsername gets a reference to the given string and assigns it to the ExternalUsername field.
func (o *DsUpdateLdap) SetExternalUsername(v string) {
	o.ExternalUsername = &v
}

// GetGroupDn returns the GroupDn field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetGroupDn() string {
	if o == nil || o.GroupDn == nil {
		var ret string
		return ret
	}
	return *o.GroupDn
}

// GetGroupDnOk returns a tuple with the GroupDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetGroupDnOk() (*string, bool) {
	if o == nil || o.GroupDn == nil {
		return nil, false
	}
	return o.GroupDn, true
}

// HasGroupDn returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasGroupDn() bool {
	if o != nil && o.GroupDn != nil {
		return true
	}

	return false
}

// SetGroupDn gets a reference to the given string and assigns it to the GroupDn field.
func (o *DsUpdateLdap) SetGroupDn(v string) {
	o.GroupDn = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DsUpdateLdap) SetJson(v bool) {
	o.Json = &v
}

// GetLdapCaCert returns the LdapCaCert field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetLdapCaCert() string {
	if o == nil || o.LdapCaCert == nil {
		var ret string
		return ret
	}
	return *o.LdapCaCert
}

// GetLdapCaCertOk returns a tuple with the LdapCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetLdapCaCertOk() (*string, bool) {
	if o == nil || o.LdapCaCert == nil {
		return nil, false
	}
	return o.LdapCaCert, true
}

// HasLdapCaCert returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasLdapCaCert() bool {
	if o != nil && o.LdapCaCert != nil {
		return true
	}

	return false
}

// SetLdapCaCert gets a reference to the given string and assigns it to the LdapCaCert field.
func (o *DsUpdateLdap) SetLdapCaCert(v string) {
	o.LdapCaCert = &v
}

// GetLdapUrl returns the LdapUrl field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetLdapUrl() string {
	if o == nil || o.LdapUrl == nil {
		var ret string
		return ret
	}
	return *o.LdapUrl
}

// GetLdapUrlOk returns a tuple with the LdapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetLdapUrlOk() (*string, bool) {
	if o == nil || o.LdapUrl == nil {
		return nil, false
	}
	return o.LdapUrl, true
}

// HasLdapUrl returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasLdapUrl() bool {
	if o != nil && o.LdapUrl != nil {
		return true
	}

	return false
}

// SetLdapUrl gets a reference to the given string and assigns it to the LdapUrl field.
func (o *DsUpdateLdap) SetLdapUrl(v string) {
	o.LdapUrl = &v
}

// GetName returns the Name field value
func (o *DsUpdateLdap) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DsUpdateLdap) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DsUpdateLdap) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DsUpdateLdap) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DsUpdateLdap) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DsUpdateLdap) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DsUpdateLdap) SetToken(v string) {
	o.Token = &v
}

// GetTokenExpiration returns the TokenExpiration field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetTokenExpiration() string {
	if o == nil || o.TokenExpiration == nil {
		var ret string
		return ret
	}
	return *o.TokenExpiration
}

// GetTokenExpirationOk returns a tuple with the TokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetTokenExpirationOk() (*string, bool) {
	if o == nil || o.TokenExpiration == nil {
		return nil, false
	}
	return o.TokenExpiration, true
}

// HasTokenExpiration returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasTokenExpiration() bool {
	if o != nil && o.TokenExpiration != nil {
		return true
	}

	return false
}

// SetTokenExpiration gets a reference to the given string and assigns it to the TokenExpiration field.
func (o *DsUpdateLdap) SetTokenExpiration(v string) {
	o.TokenExpiration = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DsUpdateLdap) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserAttribute returns the UserAttribute field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetUserAttribute() string {
	if o == nil || o.UserAttribute == nil {
		var ret string
		return ret
	}
	return *o.UserAttribute
}

// GetUserAttributeOk returns a tuple with the UserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetUserAttributeOk() (*string, bool) {
	if o == nil || o.UserAttribute == nil {
		return nil, false
	}
	return o.UserAttribute, true
}

// HasUserAttribute returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasUserAttribute() bool {
	if o != nil && o.UserAttribute != nil {
		return true
	}

	return false
}

// SetUserAttribute gets a reference to the given string and assigns it to the UserAttribute field.
func (o *DsUpdateLdap) SetUserAttribute(v string) {
	o.UserAttribute = &v
}

// GetUserDn returns the UserDn field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetUserDn() string {
	if o == nil || o.UserDn == nil {
		var ret string
		return ret
	}
	return *o.UserDn
}

// GetUserDnOk returns a tuple with the UserDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetUserDnOk() (*string, bool) {
	if o == nil || o.UserDn == nil {
		return nil, false
	}
	return o.UserDn, true
}

// HasUserDn returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasUserDn() bool {
	if o != nil && o.UserDn != nil {
		return true
	}

	return false
}

// SetUserDn gets a reference to the given string and assigns it to the UserDn field.
func (o *DsUpdateLdap) SetUserDn(v string) {
	o.UserDn = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DsUpdateLdap) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateLdap) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DsUpdateLdap) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DsUpdateLdap) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DsUpdateLdap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BindDn != nil {
		toSerialize["bind-dn"] = o.BindDn
	}
	if o.BindDnPassword != nil {
		toSerialize["bind-dn-password"] = o.BindDnPassword
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.ExternalUsername != nil {
		toSerialize["external-username"] = o.ExternalUsername
	}
	if o.GroupDn != nil {
		toSerialize["group-dn"] = o.GroupDn
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
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
	return json.Marshal(toSerialize)
}

type NullableDsUpdateLdap struct {
	value *DsUpdateLdap
	isSet bool
}

func (v NullableDsUpdateLdap) Get() *DsUpdateLdap {
	return v.value
}

func (v *NullableDsUpdateLdap) Set(val *DsUpdateLdap) {
	v.value = val
	v.isSet = true
}

func (v NullableDsUpdateLdap) IsSet() bool {
	return v.isSet
}

func (v *NullableDsUpdateLdap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDsUpdateLdap(val *DsUpdateLdap) *NullableDsUpdateLdap {
	return &NullableDsUpdateLdap{value: val, isSet: true}
}

func (v NullableDsUpdateLdap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDsUpdateLdap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


