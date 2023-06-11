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

// checks if the UpdateLdapTargetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLdapTargetDetails{}

// UpdateLdapTargetDetails struct for UpdateLdapTargetDetails
type UpdateLdapTargetDetails struct {
	BindDn *string `json:"bind-dn,omitempty"`
	BindDnPassword *string `json:"bind-dn-password,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Whether to keep previous version [true/false]. If not set, use default according to account settings
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	Key *string `json:"key,omitempty"`
	LdapCaCert *string `json:"ldap-ca-cert,omitempty"`
	LdapUrl *string `json:"ldap-url,omitempty"`
	// Target name
	Name string `json:"name"`
	// Deprecated
	NewVersion *bool `json:"new-version,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	ProtectionKey *string `json:"protection_key,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	TokenExpiration *string `json:"token-expiration,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateLdapTargetDetails instantiates a new UpdateLdapTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLdapTargetDetails(name string) *UpdateLdapTargetDetails {
	this := UpdateLdapTargetDetails{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewUpdateLdapTargetDetailsWithDefaults instantiates a new UpdateLdapTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLdapTargetDetailsWithDefaults() *UpdateLdapTargetDetails {
	this := UpdateLdapTargetDetails{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetBindDn returns the BindDn field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetBindDn() string {
	if o == nil || IsNil(o.BindDn) {
		var ret string
		return ret
	}
	return *o.BindDn
}

// GetBindDnOk returns a tuple with the BindDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetBindDnOk() (*string, bool) {
	if o == nil || IsNil(o.BindDn) {
		return nil, false
	}
	return o.BindDn, true
}

// HasBindDn returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasBindDn() bool {
	if o != nil && !IsNil(o.BindDn) {
		return true
	}

	return false
}

// SetBindDn gets a reference to the given string and assigns it to the BindDn field.
func (o *UpdateLdapTargetDetails) SetBindDn(v string) {
	o.BindDn = &v
}

// GetBindDnPassword returns the BindDnPassword field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetBindDnPassword() string {
	if o == nil || IsNil(o.BindDnPassword) {
		var ret string
		return ret
	}
	return *o.BindDnPassword
}

// GetBindDnPasswordOk returns a tuple with the BindDnPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetBindDnPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.BindDnPassword) {
		return nil, false
	}
	return o.BindDnPassword, true
}

// HasBindDnPassword returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasBindDnPassword() bool {
	if o != nil && !IsNil(o.BindDnPassword) {
		return true
	}

	return false
}

// SetBindDnPassword gets a reference to the given string and assigns it to the BindDnPassword field.
func (o *UpdateLdapTargetDetails) SetBindDnPassword(v string) {
	o.BindDnPassword = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateLdapTargetDetails) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetKeepPrevVersion() string {
	if o == nil || IsNil(o.KeepPrevVersion) {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KeepPrevVersion) {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasKeepPrevVersion() bool {
	if o != nil && !IsNil(o.KeepPrevVersion) {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateLdapTargetDetails) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateLdapTargetDetails) SetKey(v string) {
	o.Key = &v
}

// GetLdapCaCert returns the LdapCaCert field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetLdapCaCert() string {
	if o == nil || IsNil(o.LdapCaCert) {
		var ret string
		return ret
	}
	return *o.LdapCaCert
}

// GetLdapCaCertOk returns a tuple with the LdapCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetLdapCaCertOk() (*string, bool) {
	if o == nil || IsNil(o.LdapCaCert) {
		return nil, false
	}
	return o.LdapCaCert, true
}

// HasLdapCaCert returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasLdapCaCert() bool {
	if o != nil && !IsNil(o.LdapCaCert) {
		return true
	}

	return false
}

// SetLdapCaCert gets a reference to the given string and assigns it to the LdapCaCert field.
func (o *UpdateLdapTargetDetails) SetLdapCaCert(v string) {
	o.LdapCaCert = &v
}

// GetLdapUrl returns the LdapUrl field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetLdapUrl() string {
	if o == nil || IsNil(o.LdapUrl) {
		var ret string
		return ret
	}
	return *o.LdapUrl
}

// GetLdapUrlOk returns a tuple with the LdapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetLdapUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LdapUrl) {
		return nil, false
	}
	return o.LdapUrl, true
}

// HasLdapUrl returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasLdapUrl() bool {
	if o != nil && !IsNil(o.LdapUrl) {
		return true
	}

	return false
}

// SetLdapUrl gets a reference to the given string and assigns it to the LdapUrl field.
func (o *UpdateLdapTargetDetails) SetLdapUrl(v string) {
	o.LdapUrl = &v
}

// GetName returns the Name field value
func (o *UpdateLdapTargetDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateLdapTargetDetails) SetName(v string) {
	o.Name = v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetNewVersion() bool {
	if o == nil || IsNil(o.NewVersion) {
		var ret bool
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetNewVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.NewVersion) {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasNewVersion() bool {
	if o != nil && !IsNil(o.NewVersion) {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given bool and assigns it to the NewVersion field.
func (o *UpdateLdapTargetDetails) SetNewVersion(v bool) {
	o.NewVersion = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetProtectionKey() string {
	if o == nil || IsNil(o.ProtectionKey) {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetProtectionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProtectionKey) {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasProtectionKey() bool {
	if o != nil && !IsNil(o.ProtectionKey) {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *UpdateLdapTargetDetails) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateLdapTargetDetails) SetToken(v string) {
	o.Token = &v
}

// GetTokenExpiration returns the TokenExpiration field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetTokenExpiration() string {
	if o == nil || IsNil(o.TokenExpiration) {
		var ret string
		return ret
	}
	return *o.TokenExpiration
}

// GetTokenExpirationOk returns a tuple with the TokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetTokenExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.TokenExpiration) {
		return nil, false
	}
	return o.TokenExpiration, true
}

// HasTokenExpiration returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasTokenExpiration() bool {
	if o != nil && !IsNil(o.TokenExpiration) {
		return true
	}

	return false
}

// SetTokenExpiration gets a reference to the given string and assigns it to the TokenExpiration field.
func (o *UpdateLdapTargetDetails) SetTokenExpiration(v string) {
	o.TokenExpiration = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateLdapTargetDetails) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLdapTargetDetails) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateLdapTargetDetails) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateLdapTargetDetails) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateLdapTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLdapTargetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BindDn) {
		toSerialize["bind-dn"] = o.BindDn
	}
	if !IsNil(o.BindDnPassword) {
		toSerialize["bind-dn-password"] = o.BindDnPassword
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.KeepPrevVersion) {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.LdapCaCert) {
		toSerialize["ldap-ca-cert"] = o.LdapCaCert
	}
	if !IsNil(o.LdapUrl) {
		toSerialize["ldap-url"] = o.LdapUrl
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NewVersion) {
		toSerialize["new-version"] = o.NewVersion
	}
	if !IsNil(o.ProtectionKey) {
		toSerialize["protection_key"] = o.ProtectionKey
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.TokenExpiration) {
		toSerialize["token-expiration"] = o.TokenExpiration
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableUpdateLdapTargetDetails struct {
	value *UpdateLdapTargetDetails
	isSet bool
}

func (v NullableUpdateLdapTargetDetails) Get() *UpdateLdapTargetDetails {
	return v.value
}

func (v *NullableUpdateLdapTargetDetails) Set(val *UpdateLdapTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLdapTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLdapTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLdapTargetDetails(val *UpdateLdapTargetDetails) *NullableUpdateLdapTargetDetails {
	return &NullableUpdateLdapTargetDetails{value: val, isSet: true}
}

func (v NullableUpdateLdapTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLdapTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


