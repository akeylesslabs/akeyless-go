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

// CreateAuthMethodSAML createAuthMethodSAML is a command that creates a new auth method that will be available to authenticate using SAML.
type CreateAuthMethodSAML struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// Allowed redirect URIs after the authentication
	AllowedRedirectUri *[]string `json:"allowed-redirect-uri,omitempty"`
	// A CIDR whitelist with the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// if true: enforce role-association must include sub claims
	ForceSubClaims *bool `json:"force-sub-claims,omitempty"`
	// IDP metadata url
	IdpMetadataUrl *string `json:"idp-metadata-url,omitempty"`
	// IDP metadata xml data
	IdpMetadataXmlData *string `json:"idp-metadata-xml-data,omitempty"`
	// Jwt TTL
	JwtTtl *int64 `json:"jwt-ttl,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// A unique identifier (ID) value should be configured for OAuth2, LDAP and SAML authentication method types and is usually a value such as the email, username, or upn for example. Whenever a user logs in with a token, these authentication types issue a \"sub claim\" that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization.
	UniqueIdentifier string `json:"unique-identifier"`
}

// NewCreateAuthMethodSAML instantiates a new CreateAuthMethodSAML object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodSAML(name string, uniqueIdentifier string, ) *CreateAuthMethodSAML {
	this := CreateAuthMethodSAML{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	this.Name = name
	this.UniqueIdentifier = uniqueIdentifier
	return &this
}

// NewCreateAuthMethodSAMLWithDefaults instantiates a new CreateAuthMethodSAML object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodSAMLWithDefaults() *CreateAuthMethodSAML {
	this := CreateAuthMethodSAML{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *CreateAuthMethodSAML) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *CreateAuthMethodSAML) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *CreateAuthMethodSAML) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetAllowedRedirectUri returns the AllowedRedirectUri field value if set, zero value otherwise.
func (o *CreateAuthMethodSAML) GetAllowedRedirectUri() []string {
	if o == nil || o.AllowedRedirectUri == nil {
		var ret []string
		return ret
	}
	return *o.AllowedRedirectUri
}

// GetAllowedRedirectUriOk returns a tuple with the AllowedRedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetAllowedRedirectUriOk() (*[]string, bool) {
	if o == nil || o.AllowedRedirectUri == nil {
		return nil, false
	}
	return o.AllowedRedirectUri, true
}

// HasAllowedRedirectUri returns a boolean if a field has been set.
func (o *CreateAuthMethodSAML) HasAllowedRedirectUri() bool {
	if o != nil && o.AllowedRedirectUri != nil {
		return true
	}

	return false
}

// SetAllowedRedirectUri gets a reference to the given []string and assigns it to the AllowedRedirectUri field.
func (o *CreateAuthMethodSAML) SetAllowedRedirectUri(v []string) {
	o.AllowedRedirectUri = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *CreateAuthMethodSAML) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *CreateAuthMethodSAML) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *CreateAuthMethodSAML) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *CreateAuthMethodSAML) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *CreateAuthMethodSAML) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *CreateAuthMethodSAML) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetIdpMetadataUrl returns the IdpMetadataUrl field value if set, zero value otherwise.
func (o *CreateAuthMethodSAML) GetIdpMetadataUrl() string {
	if o == nil || o.IdpMetadataUrl == nil {
		var ret string
		return ret
	}
	return *o.IdpMetadataUrl
}

// GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetIdpMetadataUrlOk() (*string, bool) {
	if o == nil || o.IdpMetadataUrl == nil {
		return nil, false
	}
	return o.IdpMetadataUrl, true
}

// HasIdpMetadataUrl returns a boolean if a field has been set.
func (o *CreateAuthMethodSAML) HasIdpMetadataUrl() bool {
	if o != nil && o.IdpMetadataUrl != nil {
		return true
	}

	return false
}

// SetIdpMetadataUrl gets a reference to the given string and assigns it to the IdpMetadataUrl field.
func (o *CreateAuthMethodSAML) SetIdpMetadataUrl(v string) {
	o.IdpMetadataUrl = &v
}

// GetIdpMetadataXmlData returns the IdpMetadataXmlData field value if set, zero value otherwise.
func (o *CreateAuthMethodSAML) GetIdpMetadataXmlData() string {
	if o == nil || o.IdpMetadataXmlData == nil {
		var ret string
		return ret
	}
	return *o.IdpMetadataXmlData
}

// GetIdpMetadataXmlDataOk returns a tuple with the IdpMetadataXmlData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetIdpMetadataXmlDataOk() (*string, bool) {
	if o == nil || o.IdpMetadataXmlData == nil {
		return nil, false
	}
	return o.IdpMetadataXmlData, true
}

// HasIdpMetadataXmlData returns a boolean if a field has been set.
func (o *CreateAuthMethodSAML) HasIdpMetadataXmlData() bool {
	if o != nil && o.IdpMetadataXmlData != nil {
		return true
	}

	return false
}

// SetIdpMetadataXmlData gets a reference to the given string and assigns it to the IdpMetadataXmlData field.
func (o *CreateAuthMethodSAML) SetIdpMetadataXmlData(v string) {
	o.IdpMetadataXmlData = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *CreateAuthMethodSAML) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *CreateAuthMethodSAML) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *CreateAuthMethodSAML) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *CreateAuthMethodSAML) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAuthMethodSAML) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateAuthMethodSAML) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateAuthMethodSAML) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateAuthMethodSAML) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateAuthMethodSAML) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateAuthMethodSAML) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateAuthMethodSAML) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUniqueIdentifier returns the UniqueIdentifier field value
func (o *CreateAuthMethodSAML) GetUniqueIdentifier() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniqueIdentifier
}

// GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodSAML) GetUniqueIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniqueIdentifier, true
}

// SetUniqueIdentifier sets field value
func (o *CreateAuthMethodSAML) SetUniqueIdentifier(v string) {
	o.UniqueIdentifier = v
}

func (o CreateAuthMethodSAML) MarshalJSON() ([]byte, error) {
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
	if o.ForceSubClaims != nil {
		toSerialize["force-sub-claims"] = o.ForceSubClaims
	}
	if o.IdpMetadataUrl != nil {
		toSerialize["idp-metadata-url"] = o.IdpMetadataUrl
	}
	if o.IdpMetadataXmlData != nil {
		toSerialize["idp-metadata-xml-data"] = o.IdpMetadataXmlData
	}
	if o.JwtTtl != nil {
		toSerialize["jwt-ttl"] = o.JwtTtl
	}
	if true {
		toSerialize["name"] = o.Name
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

type NullableCreateAuthMethodSAML struct {
	value *CreateAuthMethodSAML
	isSet bool
}

func (v NullableCreateAuthMethodSAML) Get() *CreateAuthMethodSAML {
	return v.value
}

func (v *NullableCreateAuthMethodSAML) Set(val *CreateAuthMethodSAML) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodSAML) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodSAML) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodSAML(val *CreateAuthMethodSAML) *NullableCreateAuthMethodSAML {
	return &NullableCreateAuthMethodSAML{value: val, isSet: true}
}

func (v NullableCreateAuthMethodSAML) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodSAML) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


