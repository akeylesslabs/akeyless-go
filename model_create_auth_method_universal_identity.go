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

// CreateAuthMethodUniversalIdentity createAuthMethodUniversalIdentity is a command that creates a new auth method that will be able to authenticate using Akeyless Universal Identity.
type CreateAuthMethodUniversalIdentity struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// A CIDR whitelist with the IPs that the access is restricted to
	BoundIps *[]string `json:"bound-ips,omitempty"`
	// Deny from root to create children
	DenyInheritance *bool `json:"deny-inheritance,omitempty"`
	// Deny from the token to rotate
	DenyRotate *bool `json:"deny-rotate,omitempty"`
	// if true: enforce role-association must include sub claims
	ForceSubClaims *bool `json:"force-sub-claims,omitempty"`
	// A CIDR whitelist with the GW IPs that the access is restricted to
	GwBoundIps *[]string `json:"gw-bound-ips,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Jwt TTL
	JwtTtl *int64 `json:"jwt-ttl,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Token ttl
	Ttl *int32 `json:"ttl,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewCreateAuthMethodUniversalIdentity instantiates a new CreateAuthMethodUniversalIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodUniversalIdentity(name string, ) *CreateAuthMethodUniversalIdentity {
	this := CreateAuthMethodUniversalIdentity{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var json bool = false
	this.Json = &json
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	this.Name = name
	var ttl int32 = 60
	this.Ttl = &ttl
	return &this
}

// NewCreateAuthMethodUniversalIdentityWithDefaults instantiates a new CreateAuthMethodUniversalIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodUniversalIdentityWithDefaults() *CreateAuthMethodUniversalIdentity {
	this := CreateAuthMethodUniversalIdentity{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var json bool = false
	this.Json = &json
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	var ttl int32 = 60
	this.Ttl = &ttl
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetAccessExpires() int64 {
	if o == nil || o.AccessExpires == nil {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || o.AccessExpires == nil {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasAccessExpires() bool {
	if o != nil && o.AccessExpires != nil {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *CreateAuthMethodUniversalIdentity) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetBoundIps() []string {
	if o == nil || o.BoundIps == nil {
		var ret []string
		return ret
	}
	return *o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetBoundIpsOk() (*[]string, bool) {
	if o == nil || o.BoundIps == nil {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasBoundIps() bool {
	if o != nil && o.BoundIps != nil {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *CreateAuthMethodUniversalIdentity) SetBoundIps(v []string) {
	o.BoundIps = &v
}

// GetDenyInheritance returns the DenyInheritance field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetDenyInheritance() bool {
	if o == nil || o.DenyInheritance == nil {
		var ret bool
		return ret
	}
	return *o.DenyInheritance
}

// GetDenyInheritanceOk returns a tuple with the DenyInheritance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetDenyInheritanceOk() (*bool, bool) {
	if o == nil || o.DenyInheritance == nil {
		return nil, false
	}
	return o.DenyInheritance, true
}

// HasDenyInheritance returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasDenyInheritance() bool {
	if o != nil && o.DenyInheritance != nil {
		return true
	}

	return false
}

// SetDenyInheritance gets a reference to the given bool and assigns it to the DenyInheritance field.
func (o *CreateAuthMethodUniversalIdentity) SetDenyInheritance(v bool) {
	o.DenyInheritance = &v
}

// GetDenyRotate returns the DenyRotate field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetDenyRotate() bool {
	if o == nil || o.DenyRotate == nil {
		var ret bool
		return ret
	}
	return *o.DenyRotate
}

// GetDenyRotateOk returns a tuple with the DenyRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetDenyRotateOk() (*bool, bool) {
	if o == nil || o.DenyRotate == nil {
		return nil, false
	}
	return o.DenyRotate, true
}

// HasDenyRotate returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasDenyRotate() bool {
	if o != nil && o.DenyRotate != nil {
		return true
	}

	return false
}

// SetDenyRotate gets a reference to the given bool and assigns it to the DenyRotate field.
func (o *CreateAuthMethodUniversalIdentity) SetDenyRotate(v bool) {
	o.DenyRotate = &v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetForceSubClaims() bool {
	if o == nil || o.ForceSubClaims == nil {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || o.ForceSubClaims == nil {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasForceSubClaims() bool {
	if o != nil && o.ForceSubClaims != nil {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *CreateAuthMethodUniversalIdentity) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetGwBoundIps returns the GwBoundIps field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetGwBoundIps() []string {
	if o == nil || o.GwBoundIps == nil {
		var ret []string
		return ret
	}
	return *o.GwBoundIps
}

// GetGwBoundIpsOk returns a tuple with the GwBoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetGwBoundIpsOk() (*[]string, bool) {
	if o == nil || o.GwBoundIps == nil {
		return nil, false
	}
	return o.GwBoundIps, true
}

// HasGwBoundIps returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasGwBoundIps() bool {
	if o != nil && o.GwBoundIps != nil {
		return true
	}

	return false
}

// SetGwBoundIps gets a reference to the given []string and assigns it to the GwBoundIps field.
func (o *CreateAuthMethodUniversalIdentity) SetGwBoundIps(v []string) {
	o.GwBoundIps = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateAuthMethodUniversalIdentity) SetJson(v bool) {
	o.Json = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetJwtTtl() int64 {
	if o == nil || o.JwtTtl == nil {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetJwtTtlOk() (*int64, bool) {
	if o == nil || o.JwtTtl == nil {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasJwtTtl() bool {
	if o != nil && o.JwtTtl != nil {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *CreateAuthMethodUniversalIdentity) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *CreateAuthMethodUniversalIdentity) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAuthMethodUniversalIdentity) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateAuthMethodUniversalIdentity) SetToken(v string) {
	o.Token = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *CreateAuthMethodUniversalIdentity) SetTtl(v int32) {
	o.Ttl = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateAuthMethodUniversalIdentity) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodUniversalIdentity) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateAuthMethodUniversalIdentity) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateAuthMethodUniversalIdentity) SetUidToken(v string) {
	o.UidToken = &v
}

func (o CreateAuthMethodUniversalIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessExpires != nil {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if o.BoundIps != nil {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if o.DenyInheritance != nil {
		toSerialize["deny-inheritance"] = o.DenyInheritance
	}
	if o.DenyRotate != nil {
		toSerialize["deny-rotate"] = o.DenyRotate
	}
	if o.ForceSubClaims != nil {
		toSerialize["force-sub-claims"] = o.ForceSubClaims
	}
	if o.GwBoundIps != nil {
		toSerialize["gw-bound-ips"] = o.GwBoundIps
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
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
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAuthMethodUniversalIdentity struct {
	value *CreateAuthMethodUniversalIdentity
	isSet bool
}

func (v NullableCreateAuthMethodUniversalIdentity) Get() *CreateAuthMethodUniversalIdentity {
	return v.value
}

func (v *NullableCreateAuthMethodUniversalIdentity) Set(val *CreateAuthMethodUniversalIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodUniversalIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodUniversalIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodUniversalIdentity(val *CreateAuthMethodUniversalIdentity) *NullableCreateAuthMethodUniversalIdentity {
	return &NullableCreateAuthMethodUniversalIdentity{value: val, isSet: true}
}

func (v NullableCreateAuthMethodUniversalIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodUniversalIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


