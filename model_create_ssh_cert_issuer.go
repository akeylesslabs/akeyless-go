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

// CreateSSHCertIssuer struct for CreateSSHCertIssuer
type CreateSSHCertIssuer struct {
	// Users allowed to fetch the certificate, e.g root,ubuntu
	AllowedUsers string `json:"allowed-users"`
	// Signed certificates with extensions, e.g permit-port-forwarding=\\\"\\\"
	Extensions *map[string]string `json:"extensions,omitempty"`
	// A metadata about the issuer
	Metadata *string `json:"metadata,omitempty"`
	// SSH certificate issuer name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Signed certificates with principal, e.g example_role1,example_role2
	Principals *string `json:"principals,omitempty"`
	SecureAccessBastionApi *string `json:"secure-access-bastion-api,omitempty"`
	SecureAccessBastionSsh *string `json:"secure-access-bastion-ssh,omitempty"`
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	SecureAccessSshCredsUser *string `json:"secure-access-ssh-creds-user,omitempty"`
	// A key to sign the certificate with
	SignerKeyName string `json:"signer-key-name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The requested Time To Live for the certificate, use second units
	Ttl int64 `json:"ttl"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewCreateSSHCertIssuer instantiates a new CreateSSHCertIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSSHCertIssuer(allowedUsers string, name string, signerKeyName string, ttl int64, ) *CreateSSHCertIssuer {
	this := CreateSSHCertIssuer{}
	this.AllowedUsers = allowedUsers
	this.Name = name
	this.SignerKeyName = signerKeyName
	this.Ttl = ttl
	return &this
}

// NewCreateSSHCertIssuerWithDefaults instantiates a new CreateSSHCertIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSSHCertIssuerWithDefaults() *CreateSSHCertIssuer {
	this := CreateSSHCertIssuer{}
	return &this
}

// GetAllowedUsers returns the AllowedUsers field value
func (o *CreateSSHCertIssuer) GetAllowedUsers() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AllowedUsers
}

// GetAllowedUsersOk returns a tuple with the AllowedUsers field value
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetAllowedUsersOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AllowedUsers, true
}

// SetAllowedUsers sets field value
func (o *CreateSSHCertIssuer) SetAllowedUsers(v string) {
	o.AllowedUsers = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetExtensions() map[string]string {
	if o == nil || o.Extensions == nil {
		var ret map[string]string
		return ret
	}
	return *o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetExtensionsOk() (*map[string]string, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]string and assigns it to the Extensions field.
func (o *CreateSSHCertIssuer) SetExtensions(v map[string]string) {
	o.Extensions = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreateSSHCertIssuer) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *CreateSSHCertIssuer) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSSHCertIssuer) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateSSHCertIssuer) SetPassword(v string) {
	o.Password = &v
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetPrincipals() string {
	if o == nil || o.Principals == nil {
		var ret string
		return ret
	}
	return *o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetPrincipalsOk() (*string, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasPrincipals() bool {
	if o != nil && o.Principals != nil {
		return true
	}

	return false
}

// SetPrincipals gets a reference to the given string and assigns it to the Principals field.
func (o *CreateSSHCertIssuer) SetPrincipals(v string) {
	o.Principals = &v
}

// GetSecureAccessBastionApi returns the SecureAccessBastionApi field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetSecureAccessBastionApi() string {
	if o == nil || o.SecureAccessBastionApi == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionApi
}

// GetSecureAccessBastionApiOk returns a tuple with the SecureAccessBastionApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetSecureAccessBastionApiOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionApi == nil {
		return nil, false
	}
	return o.SecureAccessBastionApi, true
}

// HasSecureAccessBastionApi returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasSecureAccessBastionApi() bool {
	if o != nil && o.SecureAccessBastionApi != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionApi gets a reference to the given string and assigns it to the SecureAccessBastionApi field.
func (o *CreateSSHCertIssuer) SetSecureAccessBastionApi(v string) {
	o.SecureAccessBastionApi = &v
}

// GetSecureAccessBastionSsh returns the SecureAccessBastionSsh field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetSecureAccessBastionSsh() string {
	if o == nil || o.SecureAccessBastionSsh == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionSsh
}

// GetSecureAccessBastionSshOk returns a tuple with the SecureAccessBastionSsh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetSecureAccessBastionSshOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionSsh == nil {
		return nil, false
	}
	return o.SecureAccessBastionSsh, true
}

// HasSecureAccessBastionSsh returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasSecureAccessBastionSsh() bool {
	if o != nil && o.SecureAccessBastionSsh != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionSsh gets a reference to the given string and assigns it to the SecureAccessBastionSsh field.
func (o *CreateSSHCertIssuer) SetSecureAccessBastionSsh(v string) {
	o.SecureAccessBastionSsh = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *CreateSSHCertIssuer) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *CreateSSHCertIssuer) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessSshCredsUser returns the SecureAccessSshCredsUser field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetSecureAccessSshCredsUser() string {
	if o == nil || o.SecureAccessSshCredsUser == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessSshCredsUser
}

// GetSecureAccessSshCredsUserOk returns a tuple with the SecureAccessSshCredsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetSecureAccessSshCredsUserOk() (*string, bool) {
	if o == nil || o.SecureAccessSshCredsUser == nil {
		return nil, false
	}
	return o.SecureAccessSshCredsUser, true
}

// HasSecureAccessSshCredsUser returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasSecureAccessSshCredsUser() bool {
	if o != nil && o.SecureAccessSshCredsUser != nil {
		return true
	}

	return false
}

// SetSecureAccessSshCredsUser gets a reference to the given string and assigns it to the SecureAccessSshCredsUser field.
func (o *CreateSSHCertIssuer) SetSecureAccessSshCredsUser(v string) {
	o.SecureAccessSshCredsUser = &v
}

// GetSignerKeyName returns the SignerKeyName field value
func (o *CreateSSHCertIssuer) GetSignerKeyName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SignerKeyName
}

// GetSignerKeyNameOk returns a tuple with the SignerKeyName field value
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetSignerKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SignerKeyName, true
}

// SetSignerKeyName sets field value
func (o *CreateSSHCertIssuer) SetSignerKeyName(v string) {
	o.SignerKeyName = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateSSHCertIssuer) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value
func (o *CreateSSHCertIssuer) GetTtl() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetTtlOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *CreateSSHCertIssuer) SetTtl(v int64) {
	o.Ttl = v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateSSHCertIssuer) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateSSHCertIssuer) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHCertIssuer) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateSSHCertIssuer) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateSSHCertIssuer) SetUsername(v string) {
	o.Username = &v
}

func (o CreateSSHCertIssuer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowed-users"] = o.AllowedUsers
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSSHCertIssuer struct {
	value *CreateSSHCertIssuer
	isSet bool
}

func (v NullableCreateSSHCertIssuer) Get() *CreateSSHCertIssuer {
	return v.value
}

func (v *NullableCreateSSHCertIssuer) Set(val *CreateSSHCertIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSSHCertIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSSHCertIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSSHCertIssuer(val *CreateSSHCertIssuer) *NullableCreateSSHCertIssuer {
	return &NullableCreateSSHCertIssuer{value: val, isSet: true}
}

func (v NullableCreateSSHCertIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSSHCertIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


