/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 0.0.1
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// OAuth2AccessRules OAuth2AccessRules contains access rules specific to OAuth2 authentication method.
type OAuth2AccessRules struct {
	// The audience in the JWT.
	Audience *string `json:"audience,omitempty"`
	// The claims that login is restricted to.
	BoundClaims *[]OAuth2CustomClaim `json:"bound_claims,omitempty"`
	// The clients ids that login is restricted to.
	BoundClientsId *[]string `json:"bound_clients_id,omitempty"`
	// Issuer URL
	Issuer *string `json:"issuer,omitempty"`
	// The URL to the JSON Web Key Set (JWKS) that containing the public keys that should be used to verify any JSON Web Token (JWT) issued by the authorization server.
	JwksUri *string `json:"jwks_uri,omitempty"`
}

// NewOAuth2AccessRules instantiates a new OAuth2AccessRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2AccessRules() *OAuth2AccessRules {
	this := OAuth2AccessRules{}
	return &this
}

// NewOAuth2AccessRulesWithDefaults instantiates a new OAuth2AccessRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2AccessRulesWithDefaults() *OAuth2AccessRules {
	this := OAuth2AccessRules{}
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *OAuth2AccessRules) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AccessRules) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *OAuth2AccessRules) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *OAuth2AccessRules) SetAudience(v string) {
	o.Audience = &v
}

// GetBoundClaims returns the BoundClaims field value if set, zero value otherwise.
func (o *OAuth2AccessRules) GetBoundClaims() []OAuth2CustomClaim {
	if o == nil || o.BoundClaims == nil {
		var ret []OAuth2CustomClaim
		return ret
	}
	return *o.BoundClaims
}

// GetBoundClaimsOk returns a tuple with the BoundClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AccessRules) GetBoundClaimsOk() (*[]OAuth2CustomClaim, bool) {
	if o == nil || o.BoundClaims == nil {
		return nil, false
	}
	return o.BoundClaims, true
}

// HasBoundClaims returns a boolean if a field has been set.
func (o *OAuth2AccessRules) HasBoundClaims() bool {
	if o != nil && o.BoundClaims != nil {
		return true
	}

	return false
}

// SetBoundClaims gets a reference to the given []OAuth2CustomClaim and assigns it to the BoundClaims field.
func (o *OAuth2AccessRules) SetBoundClaims(v []OAuth2CustomClaim) {
	o.BoundClaims = &v
}

// GetBoundClientsId returns the BoundClientsId field value if set, zero value otherwise.
func (o *OAuth2AccessRules) GetBoundClientsId() []string {
	if o == nil || o.BoundClientsId == nil {
		var ret []string
		return ret
	}
	return *o.BoundClientsId
}

// GetBoundClientsIdOk returns a tuple with the BoundClientsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AccessRules) GetBoundClientsIdOk() (*[]string, bool) {
	if o == nil || o.BoundClientsId == nil {
		return nil, false
	}
	return o.BoundClientsId, true
}

// HasBoundClientsId returns a boolean if a field has been set.
func (o *OAuth2AccessRules) HasBoundClientsId() bool {
	if o != nil && o.BoundClientsId != nil {
		return true
	}

	return false
}

// SetBoundClientsId gets a reference to the given []string and assigns it to the BoundClientsId field.
func (o *OAuth2AccessRules) SetBoundClientsId(v []string) {
	o.BoundClientsId = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OAuth2AccessRules) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AccessRules) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OAuth2AccessRules) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OAuth2AccessRules) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *OAuth2AccessRules) GetJwksUri() string {
	if o == nil || o.JwksUri == nil {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AccessRules) GetJwksUriOk() (*string, bool) {
	if o == nil || o.JwksUri == nil {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *OAuth2AccessRules) HasJwksUri() bool {
	if o != nil && o.JwksUri != nil {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *OAuth2AccessRules) SetJwksUri(v string) {
	o.JwksUri = &v
}

func (o OAuth2AccessRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.BoundClaims != nil {
		toSerialize["bound_claims"] = o.BoundClaims
	}
	if o.BoundClientsId != nil {
		toSerialize["bound_clients_id"] = o.BoundClientsId
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.JwksUri != nil {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2AccessRules struct {
	value *OAuth2AccessRules
	isSet bool
}

func (v NullableOAuth2AccessRules) Get() *OAuth2AccessRules {
	return v.value
}

func (v *NullableOAuth2AccessRules) Set(val *OAuth2AccessRules) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2AccessRules) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2AccessRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2AccessRules(val *OAuth2AccessRules) *NullableOAuth2AccessRules {
	return &NullableOAuth2AccessRules{value: val, isSet: true}
}

func (v NullableOAuth2AccessRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2AccessRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


