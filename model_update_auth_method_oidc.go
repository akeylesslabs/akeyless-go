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

// UpdateAuthMethodOIDC updateAuthMethodOIDC is a command that updates a new auth method that will be available to authenticate using OIDC.
type UpdateAuthMethodOIDC struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// Allowed redirect URIs after the authentication
	AllowedRedirectUri *[]string `json:"allowed-redirect-uri,omitempty"`
	// A CIDR whitelist with the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// Client ID
	ClientId *string `json:"client-id,omitempty"`
	// Client Secret
	ClientSecret *string `json:"client-secret,omitempty"`
	// if true: enforce role-association must include sub claims
	ForceSubClaims *bool `json:"force-sub-claims,omitempty"`
	// Issuer URL
	Issuer *string `json:"issuer,omitempty"`
	// Jwt TTL
	JwtTtl *int64 `json:"jwt-ttl,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// Auth Method new name
	NewName *string `json:"new-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// A unique identifier (ID) value should be configured for OIDC, OAuth2, LDAP and SAML authentication method types and is usually a value such as the email, username, or upn for example. Whenever a user logs in with a token, these authentication types issue a \"sub claim\" that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization.
	UniqueIdentifier string `json:"unique-identifier"`
}

// NewUpdateAuthMethodOIDC instantiates a new UpdateAuthMethodOIDC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthMethodOIDC(name string, uniqueIdentifier string, ) *UpdateAuthMethodOIDC {
	this := UpdateAuthMethodOIDC{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	this.Name = name
	this.UniqueIdentifier = uniqueIdentifier
	return &this
}

// NewUpdateAuthMethodOIDCWithDefaults instantiates a new UpdateAuthMethodOIDC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthMethodOIDCWithDefaults() *UpdateAuthMethodOIDC {
	this := UpdateAuthMethodOIDC{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *UpdateAuthMethodOIDC) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetAllowedRedirectUri returns the AllowedRedirectUri field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetAllowedRedirectUri() []string {
	if o == nil || o.AllowedRedirectUri == nil {
		var ret []string
		return ret
	}
	return *o.AllowedRedirectUri
}

// GetAllowedRedirectUriOk returns a tuple with the AllowedRedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetAllowedRedirectUriOk() (*[]string, bool) {
	if o == nil || o.AllowedRedirectUri == nil {
		return nil, false
	}
	return o.AllowedRedirectUri, true
}

// HasAllowedRedirectUri returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasAllowedRedirectUri() bool {
	if o != nil && o.AllowedRedirectUri != nil {
		return true
	}

	return false
}

// SetAllowedRedirectUri gets a reference to the given []string and assigns it to the AllowedRedirectUri field.
func (o *UpdateAuthMethodOIDC) SetAllowedRedirectUri(v []string) {
	o.AllowedRedirectUri = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *UpdateAuthMethodOIDC) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *UpdateAuthMethodOIDC) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *UpdateAuthMethodOIDC) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *UpdateAuthMethodOIDC) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *UpdateAuthMethodOIDC) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *UpdateAuthMethodOIDC) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *UpdateAuthMethodOIDC) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateAuthMethodOIDC) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateAuthMethodOIDC) SetNewName(v string) {
	o.NewName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateAuthMethodOIDC) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateAuthMethodOIDC) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodOIDC) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateAuthMethodOIDC) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUniqueIdentifier returns the UniqueIdentifier field value
func (o *UpdateAuthMethodOIDC) GetUniqueIdentifier() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniqueIdentifier
}

// GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOIDC) GetUniqueIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniqueIdentifier, true
}

// SetUniqueIdentifier sets field value
func (o *UpdateAuthMethodOIDC) SetUniqueIdentifier(v string) {
	o.UniqueIdentifier = v
}

func (o UpdateAuthMethodOIDC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if o.AllowedRedirectUri != nil {
		toSerialize["allowed-redirect-uri"] = o.AllowedRedirectUri
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if o.ClientId != nil {
		toSerialize["client-id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client-secret"] = o.ClientSecret
	}
	if o.ForceSubClaims != nil {
		toSerialize["force-sub-claims"] = o.ForceSubClaims
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.JwtTtl != nil {
		toSerialize["jwt-ttl"] = o.JwtTtl
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
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if true {
		toSerialize["unique-identifier"] = o.UniqueIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAuthMethodOIDC struct {
	value *UpdateAuthMethodOIDC
	isSet bool
}

func (v NullableUpdateAuthMethodOIDC) Get() *UpdateAuthMethodOIDC {
	return v.value
}

func (v *NullableUpdateAuthMethodOIDC) Set(val *UpdateAuthMethodOIDC) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthMethodOIDC) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthMethodOIDC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthMethodOIDC(val *UpdateAuthMethodOIDC) *NullableUpdateAuthMethodOIDC {
	return &NullableUpdateAuthMethodOIDC{value: val, isSet: true}
}

func (v NullableUpdateAuthMethodOIDC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthMethodOIDC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


