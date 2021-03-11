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

// Auth struct for Auth
type Auth struct {
	// Access ID
	AccessId *string `json:"access-id,omitempty"`
	// Access key (relevant only for access-type=access_key)
	AccessKey *string `json:"access-key,omitempty"`
	// Access Type (access_key/password/saml/ldap/azure_ad/aws_iam/universal_identity/jwt/gcp)
	AccessType *string `json:"access-type,omitempty"`
	// Email (relevant only for access-type=password)
	AdminEmail *string `json:"admin-email,omitempty"`
	// Password (relevant only for access-type=password)
	AdminPassword *string `json:"admin-password,omitempty"`
	// The cloud identity (relevant only for access-type=azure_ad,aws_iam,gcp)
	CloudId *string `json:"cloud-id,omitempty"`
	// GCP JWT audience
	GcpAudience *string `json:"gcp-audience,omitempty"`
	// The Json Web Token (relevant only for access-type=jwt/oidc)
	Jwt *string `json:"jwt,omitempty"`
	// LDAP password (relevant only for access-type=ldap)
	LdapPassword *string `json:"ldap_password,omitempty"`
	// LDAP username (relevant only for access-type=ldap)
	LdapUsername *string `json:"ldap_username,omitempty"`
	// The universal_identity token (relevant only for access-type=universal_identity)
	UidToken *string `json:"uid_token,omitempty"`
}

// NewAuth instantiates a new Auth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuth() *Auth {
	this := Auth{}
	var accessType string = "access_key"
	this.AccessType = &accessType
	return &this
}

// NewAuthWithDefaults instantiates a new Auth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthWithDefaults() *Auth {
	this := Auth{}
	var accessType string = "access_key"
	this.AccessType = &accessType
	return &this
}

// GetAccessId returns the AccessId field value if set, zero value otherwise.
func (o *Auth) GetAccessId() string {
	if o == nil || o.AccessId == nil {
		var ret string
		return ret
	}
	return *o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetAccessIdOk() (*string, bool) {
	if o == nil || o.AccessId == nil {
		return nil, false
	}
	return o.AccessId, true
}

// HasAccessId returns a boolean if a field has been set.
func (o *Auth) HasAccessId() bool {
	if o != nil && o.AccessId != nil {
		return true
	}

	return false
}

// SetAccessId gets a reference to the given string and assigns it to the AccessId field.
func (o *Auth) SetAccessId(v string) {
	o.AccessId = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *Auth) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *Auth) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *Auth) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *Auth) GetAccessType() string {
	if o == nil || o.AccessType == nil {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetAccessTypeOk() (*string, bool) {
	if o == nil || o.AccessType == nil {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *Auth) HasAccessType() bool {
	if o != nil && o.AccessType != nil {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *Auth) SetAccessType(v string) {
	o.AccessType = &v
}

// GetAdminEmail returns the AdminEmail field value if set, zero value otherwise.
func (o *Auth) GetAdminEmail() string {
	if o == nil || o.AdminEmail == nil {
		var ret string
		return ret
	}
	return *o.AdminEmail
}

// GetAdminEmailOk returns a tuple with the AdminEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetAdminEmailOk() (*string, bool) {
	if o == nil || o.AdminEmail == nil {
		return nil, false
	}
	return o.AdminEmail, true
}

// HasAdminEmail returns a boolean if a field has been set.
func (o *Auth) HasAdminEmail() bool {
	if o != nil && o.AdminEmail != nil {
		return true
	}

	return false
}

// SetAdminEmail gets a reference to the given string and assigns it to the AdminEmail field.
func (o *Auth) SetAdminEmail(v string) {
	o.AdminEmail = &v
}

// GetAdminPassword returns the AdminPassword field value if set, zero value otherwise.
func (o *Auth) GetAdminPassword() string {
	if o == nil || o.AdminPassword == nil {
		var ret string
		return ret
	}
	return *o.AdminPassword
}

// GetAdminPasswordOk returns a tuple with the AdminPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetAdminPasswordOk() (*string, bool) {
	if o == nil || o.AdminPassword == nil {
		return nil, false
	}
	return o.AdminPassword, true
}

// HasAdminPassword returns a boolean if a field has been set.
func (o *Auth) HasAdminPassword() bool {
	if o != nil && o.AdminPassword != nil {
		return true
	}

	return false
}

// SetAdminPassword gets a reference to the given string and assigns it to the AdminPassword field.
func (o *Auth) SetAdminPassword(v string) {
	o.AdminPassword = &v
}

// GetCloudId returns the CloudId field value if set, zero value otherwise.
func (o *Auth) GetCloudId() string {
	if o == nil || o.CloudId == nil {
		var ret string
		return ret
	}
	return *o.CloudId
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetCloudIdOk() (*string, bool) {
	if o == nil || o.CloudId == nil {
		return nil, false
	}
	return o.CloudId, true
}

// HasCloudId returns a boolean if a field has been set.
func (o *Auth) HasCloudId() bool {
	if o != nil && o.CloudId != nil {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given string and assigns it to the CloudId field.
func (o *Auth) SetCloudId(v string) {
	o.CloudId = &v
}

// GetGcpAudience returns the GcpAudience field value if set, zero value otherwise.
func (o *Auth) GetGcpAudience() string {
	if o == nil || o.GcpAudience == nil {
		var ret string
		return ret
	}
	return *o.GcpAudience
}

// GetGcpAudienceOk returns a tuple with the GcpAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetGcpAudienceOk() (*string, bool) {
	if o == nil || o.GcpAudience == nil {
		return nil, false
	}
	return o.GcpAudience, true
}

// HasGcpAudience returns a boolean if a field has been set.
func (o *Auth) HasGcpAudience() bool {
	if o != nil && o.GcpAudience != nil {
		return true
	}

	return false
}

// SetGcpAudience gets a reference to the given string and assigns it to the GcpAudience field.
func (o *Auth) SetGcpAudience(v string) {
	o.GcpAudience = &v
}

// GetJwt returns the Jwt field value if set, zero value otherwise.
func (o *Auth) GetJwt() string {
	if o == nil || o.Jwt == nil {
		var ret string
		return ret
	}
	return *o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetJwtOk() (*string, bool) {
	if o == nil || o.Jwt == nil {
		return nil, false
	}
	return o.Jwt, true
}

// HasJwt returns a boolean if a field has been set.
func (o *Auth) HasJwt() bool {
	if o != nil && o.Jwt != nil {
		return true
	}

	return false
}

// SetJwt gets a reference to the given string and assigns it to the Jwt field.
func (o *Auth) SetJwt(v string) {
	o.Jwt = &v
}

// GetLdapPassword returns the LdapPassword field value if set, zero value otherwise.
func (o *Auth) GetLdapPassword() string {
	if o == nil || o.LdapPassword == nil {
		var ret string
		return ret
	}
	return *o.LdapPassword
}

// GetLdapPasswordOk returns a tuple with the LdapPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetLdapPasswordOk() (*string, bool) {
	if o == nil || o.LdapPassword == nil {
		return nil, false
	}
	return o.LdapPassword, true
}

// HasLdapPassword returns a boolean if a field has been set.
func (o *Auth) HasLdapPassword() bool {
	if o != nil && o.LdapPassword != nil {
		return true
	}

	return false
}

// SetLdapPassword gets a reference to the given string and assigns it to the LdapPassword field.
func (o *Auth) SetLdapPassword(v string) {
	o.LdapPassword = &v
}

// GetLdapUsername returns the LdapUsername field value if set, zero value otherwise.
func (o *Auth) GetLdapUsername() string {
	if o == nil || o.LdapUsername == nil {
		var ret string
		return ret
	}
	return *o.LdapUsername
}

// GetLdapUsernameOk returns a tuple with the LdapUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetLdapUsernameOk() (*string, bool) {
	if o == nil || o.LdapUsername == nil {
		return nil, false
	}
	return o.LdapUsername, true
}

// HasLdapUsername returns a boolean if a field has been set.
func (o *Auth) HasLdapUsername() bool {
	if o != nil && o.LdapUsername != nil {
		return true
	}

	return false
}

// SetLdapUsername gets a reference to the given string and assigns it to the LdapUsername field.
func (o *Auth) SetLdapUsername(v string) {
	o.LdapUsername = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *Auth) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *Auth) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *Auth) SetUidToken(v string) {
	o.UidToken = &v
}

func (o Auth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessId != nil {
		toSerialize["access-id"] = o.AccessId
	}
	if o.AccessKey != nil {
		toSerialize["access-key"] = o.AccessKey
	}
	if o.AccessType != nil {
		toSerialize["access-type"] = o.AccessType
	}
	if o.AdminEmail != nil {
		toSerialize["admin-email"] = o.AdminEmail
	}
	if o.AdminPassword != nil {
		toSerialize["admin-password"] = o.AdminPassword
	}
	if o.CloudId != nil {
		toSerialize["cloud-id"] = o.CloudId
	}
	if o.GcpAudience != nil {
		toSerialize["gcp-audience"] = o.GcpAudience
	}
	if o.Jwt != nil {
		toSerialize["jwt"] = o.Jwt
	}
	if o.LdapPassword != nil {
		toSerialize["ldap_password"] = o.LdapPassword
	}
	if o.LdapUsername != nil {
		toSerialize["ldap_username"] = o.LdapUsername
	}
	if o.UidToken != nil {
		toSerialize["uid_token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableAuth struct {
	value *Auth
	isSet bool
}

func (v NullableAuth) Get() *Auth {
	return v.value
}

func (v *NullableAuth) Set(val *Auth) {
	v.value = val
	v.isSet = true
}

func (v NullableAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuth(val *Auth) *NullableAuth {
	return &NullableAuth{value: val, isSet: true}
}

func (v NullableAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


