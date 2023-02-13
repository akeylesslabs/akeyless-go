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

// UpdateAuthMethodOAuth2 updateAuthMethodOAuth2 is a command that updates a new auth method that will be able to authenticate using Oauth2.
type UpdateAuthMethodOAuth2 struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// The audience in the JWT
	Audience *string `json:"audience,omitempty"`
	// The clients ids that the access is restricted to
	BoundClientIds *[]string `json:"bound-client-ids,omitempty"`
	// A CIDR whitelist with the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// if true: enforce role-association must include sub claims
	ForceSubClaims *bool `json:"force-sub-claims,omitempty"`
	// A CIDR whitelist with the GW IPs that the access is restricted to
	GwBoundIps *[]string `json:"gw-bound-ips,omitempty"`
	// Issuer URL
	Issuer *string `json:"issuer,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The URL to the JSON Web Key Set (JWKS) that containing the public keys that should be used to verify any JSON Web Token (JWT) issued by the authorization server.
	JwksUri string `json:"jwks-uri"`
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
	// A unique identifier (ID) value should be configured for OAuth2, LDAP and SAML authentication method types and is usually a value such as the email, username, or upn for example. Whenever a user logs in with a token, these authentication types issue a \"sub claim\" that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization.
	UniqueIdentifier string `json:"unique-identifier"`
}

// NewUpdateAuthMethodOAuth2 instantiates a new UpdateAuthMethodOAuth2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthMethodOAuth2(jwksUri string, name string, uniqueIdentifier string, ) *UpdateAuthMethodOAuth2 {
	this := UpdateAuthMethodOAuth2{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var json bool = false
	this.Json = &json
	this.JwksUri = jwksUri
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	this.Name = name
	this.UniqueIdentifier = uniqueIdentifier
	return &this
}

// NewUpdateAuthMethodOAuth2WithDefaults instantiates a new UpdateAuthMethodOAuth2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthMethodOAuth2WithDefaults() *UpdateAuthMethodOAuth2 {
	this := UpdateAuthMethodOAuth2{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var json bool = false
	this.Json = &json
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *UpdateAuthMethodOAuth2) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *UpdateAuthMethodOAuth2) SetAudience(v string) {
	o.Audience = &v
}

// GetBoundClientIds returns the BoundClientIds field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetBoundClientIds() []string {
	if o == nil || o.BoundClientIds == nil {
		var ret []string
		return ret
	}
	return *o.BoundClientIds
}

// GetBoundClientIdsOk returns a tuple with the BoundClientIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetBoundClientIdsOk() (*[]string, bool) {
	if o == nil || o.BoundClientIds == nil {
		return nil, false
	}
	return o.BoundClientIds, true
}

// HasBoundClientIds returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasBoundClientIds() bool {
	if o != nil && o.BoundClientIds != nil {
		return true
	}

	return false
}

// SetBoundClientIds gets a reference to the given []string and assigns it to the BoundClientIds field.
func (o *UpdateAuthMethodOAuth2) SetBoundClientIds(v []string) {
	o.BoundClientIds = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *UpdateAuthMethodOAuth2) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *UpdateAuthMethodOAuth2) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetGwBoundIps returns the GwBoundIps field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetGwBoundIps() []string {
	if o == nil || o.GwBoundIps == nil {
		var ret []string
		return ret
	}
	return *o.GwBoundIps
}

// GetGwBoundIpsOk returns a tuple with the GwBoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetGwBoundIpsOk() (*[]string, bool) {
	if o == nil || o.GwBoundIps == nil {
		return nil, false
	}
	return o.GwBoundIps, true
}

// HasGwBoundIps returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasGwBoundIps() bool {
	if o != nil && o.GwBoundIps != nil {
		return true
	}

	return false
}

// SetGwBoundIps gets a reference to the given []string and assigns it to the GwBoundIps field.
func (o *UpdateAuthMethodOAuth2) SetGwBoundIps(v []string) {
	o.GwBoundIps = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *UpdateAuthMethodOAuth2) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateAuthMethodOAuth2) SetJson(v bool) {
	o.Json = &v
}

// GetJwksUri returns the JwksUri field value
func (o *UpdateAuthMethodOAuth2) GetJwksUri() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetJwksUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.JwksUri, true
}

// SetJwksUri sets field value
func (o *UpdateAuthMethodOAuth2) SetJwksUri(v string) {
	o.JwksUri = v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *UpdateAuthMethodOAuth2) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *UpdateAuthMethodOAuth2) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateAuthMethodOAuth2) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateAuthMethodOAuth2) SetNewName(v string) {
	o.NewName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateAuthMethodOAuth2) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateAuthMethodOAuth2) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateAuthMethodOAuth2) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateAuthMethodOAuth2) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUniqueIdentifier returns the UniqueIdentifier field value
func (o *UpdateAuthMethodOAuth2) GetUniqueIdentifier() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UniqueIdentifier
}

// GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field value
// and a boolean to check if the value has been set.
func (o *UpdateAuthMethodOAuth2) GetUniqueIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UniqueIdentifier, true
}

// SetUniqueIdentifier sets field value
func (o *UpdateAuthMethodOAuth2) SetUniqueIdentifier(v string) {
	o.UniqueIdentifier = v
}

func (o UpdateAuthMethodOAuth2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.BoundClientIds != nil {
		toSerialize["bound-client-ids"] = o.BoundClientIds
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if o.ForceSubClaims != nil {
		toSerialize["force-sub-claims"] = o.ForceSubClaims
	}
	if o.GwBoundIps != nil {
		toSerialize["gw-bound-ips"] = o.GwBoundIps
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["jwks-uri"] = o.JwksUri
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

type NullableUpdateAuthMethodOAuth2 struct {
	value *UpdateAuthMethodOAuth2
	isSet bool
}

func (v NullableUpdateAuthMethodOAuth2) Get() *UpdateAuthMethodOAuth2 {
	return v.value
}

func (v *NullableUpdateAuthMethodOAuth2) Set(val *UpdateAuthMethodOAuth2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthMethodOAuth2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthMethodOAuth2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthMethodOAuth2(val *UpdateAuthMethodOAuth2) *NullableUpdateAuthMethodOAuth2 {
	return &NullableUpdateAuthMethodOAuth2{value: val, isSet: true}
}

func (v NullableUpdateAuthMethodOAuth2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthMethodOAuth2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


