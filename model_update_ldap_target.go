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

// UpdateLdapTarget struct for UpdateLdapTarget
type UpdateLdapTarget struct {
	BindDn *string `json:"bind-dn,omitempty"`
	BindDnPassword *string `json:"bind-dn-password,omitempty"`
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	LdapCaCert *string `json:"ldap-ca-cert,omitempty"`
	LdapUrl *string `json:"ldap-url,omitempty"`
	// Target name
	Name string `json:"name"`
	// New target name
	NewName *string `json:"new-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	TokenExpiration *string `json:"token-expiration,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Deprecated
	UpdateVersion *bool `json:"update-version,omitempty"`
}

// NewUpdateLdapTarget instantiates a new UpdateLdapTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLdapTarget(name string, ) *UpdateLdapTarget {
	this := UpdateLdapTarget{}
	this.Name = name
	return &this
}

// NewUpdateLdapTargetWithDefaults instantiates a new UpdateLdapTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLdapTargetWithDefaults() *UpdateLdapTarget {
	this := UpdateLdapTarget{}
	return &this
}

// GetBindDn returns the BindDn field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetBindDn() string {
	if o == nil || o.BindDn == nil {
		var ret string
		return ret
	}
	return *o.BindDn
}

// GetBindDnOk returns a tuple with the BindDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetBindDnOk() (*string, bool) {
	if o == nil || o.BindDn == nil {
		return nil, false
	}
	return o.BindDn, true
}

// HasBindDn returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasBindDn() bool {
	if o != nil && o.BindDn != nil {
		return true
	}

	return false
}

// SetBindDn gets a reference to the given string and assigns it to the BindDn field.
func (o *UpdateLdapTarget) SetBindDn(v string) {
	o.BindDn = &v
}

// GetBindDnPassword returns the BindDnPassword field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetBindDnPassword() string {
	if o == nil || o.BindDnPassword == nil {
		var ret string
		return ret
	}
	return *o.BindDnPassword
}

// GetBindDnPasswordOk returns a tuple with the BindDnPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetBindDnPasswordOk() (*string, bool) {
	if o == nil || o.BindDnPassword == nil {
		return nil, false
	}
	return o.BindDnPassword, true
}

// HasBindDnPassword returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasBindDnPassword() bool {
	if o != nil && o.BindDnPassword != nil {
		return true
	}

	return false
}

// SetBindDnPassword gets a reference to the given string and assigns it to the BindDnPassword field.
func (o *UpdateLdapTarget) SetBindDnPassword(v string) {
	o.BindDnPassword = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateLdapTarget) SetComment(v string) {
	o.Comment = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateLdapTarget) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateLdapTarget) SetKey(v string) {
	o.Key = &v
}

// GetLdapCaCert returns the LdapCaCert field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetLdapCaCert() string {
	if o == nil || o.LdapCaCert == nil {
		var ret string
		return ret
	}
	return *o.LdapCaCert
}

// GetLdapCaCertOk returns a tuple with the LdapCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetLdapCaCertOk() (*string, bool) {
	if o == nil || o.LdapCaCert == nil {
		return nil, false
	}
	return o.LdapCaCert, true
}

// HasLdapCaCert returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasLdapCaCert() bool {
	if o != nil && o.LdapCaCert != nil {
		return true
	}

	return false
}

// SetLdapCaCert gets a reference to the given string and assigns it to the LdapCaCert field.
func (o *UpdateLdapTarget) SetLdapCaCert(v string) {
	o.LdapCaCert = &v
}

// GetLdapUrl returns the LdapUrl field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetLdapUrl() string {
	if o == nil || o.LdapUrl == nil {
		var ret string
		return ret
	}
	return *o.LdapUrl
}

// GetLdapUrlOk returns a tuple with the LdapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetLdapUrlOk() (*string, bool) {
	if o == nil || o.LdapUrl == nil {
		return nil, false
	}
	return o.LdapUrl, true
}

// HasLdapUrl returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasLdapUrl() bool {
	if o != nil && o.LdapUrl != nil {
		return true
	}

	return false
}

// SetLdapUrl gets a reference to the given string and assigns it to the LdapUrl field.
func (o *UpdateLdapTarget) SetLdapUrl(v string) {
	o.LdapUrl = &v
}

// GetName returns the Name field value
func (o *UpdateLdapTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateLdapTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateLdapTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateLdapTarget) SetToken(v string) {
	o.Token = &v
}

// GetTokenExpiration returns the TokenExpiration field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetTokenExpiration() string {
	if o == nil || o.TokenExpiration == nil {
		var ret string
		return ret
	}
	return *o.TokenExpiration
}

// GetTokenExpirationOk returns a tuple with the TokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetTokenExpirationOk() (*string, bool) {
	if o == nil || o.TokenExpiration == nil {
		return nil, false
	}
	return o.TokenExpiration, true
}

// HasTokenExpiration returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasTokenExpiration() bool {
	if o != nil && o.TokenExpiration != nil {
		return true
	}

	return false
}

// SetTokenExpiration gets a reference to the given string and assigns it to the TokenExpiration field.
func (o *UpdateLdapTarget) SetTokenExpiration(v string) {
	o.TokenExpiration = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateLdapTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateLdapTarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateLdapTarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateLdapTarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

func (o UpdateLdapTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BindDn != nil {
		toSerialize["bind-dn"] = o.BindDn
	}
	if o.BindDnPassword != nil {
		toSerialize["bind-dn-password"] = o.BindDnPassword
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
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
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.TokenExpiration != nil {
		toSerialize["token-expiration"] = o.TokenExpiration
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UpdateVersion != nil {
		toSerialize["update-version"] = o.UpdateVersion
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateLdapTarget struct {
	value *UpdateLdapTarget
	isSet bool
}

func (v NullableUpdateLdapTarget) Get() *UpdateLdapTarget {
	return v.value
}

func (v *NullableUpdateLdapTarget) Set(val *UpdateLdapTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLdapTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLdapTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLdapTarget(val *UpdateLdapTarget) *NullableUpdateLdapTarget {
	return &NullableUpdateLdapTarget{value: val, isSet: true}
}

func (v NullableUpdateLdapTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLdapTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

