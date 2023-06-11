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

// checks if the CreateAuthMethodLDAP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAuthMethodLDAP{}

// CreateAuthMethodLDAP createAuthMethodLDAP is a command that creates a new auth method that will be able to authenticate using LDAP.
type CreateAuthMethodLDAP struct {
	// Access expiration date in Unix timestamp (select 0 for access without expiry date)
	AccessExpires *int64 `json:"access-expires,omitempty"`
	// A CIDR whitelist with the IPs that the access is restricted to
	BoundIps []string `json:"bound-ips,omitempty"`
	// if true: enforce role-association must include sub claims
	ForceSubClaims *bool `json:"force-sub-claims,omitempty"`
	// Automatically generate key-pair for LDAP configuration. If set to false, a public key needs to be provided [true/false]
	GenKey *string `json:"gen-key,omitempty"`
	// A CIDR whitelist with the GW IPs that the access is restricted to
	GwBoundIps []string `json:"gw-bound-ips,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Jwt TTL
	JwtTtl *int64 `json:"jwt-ttl,omitempty"`
	// Auth Method name
	Name string `json:"name"`
	// A public key generated for LDAP authentication method on Akeyless in base64 or PEM format [RSA2048]
	PublicKeyData *string `json:"public-key-data,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// A unique identifier (ID) value should be configured for OAuth2, LDAP and SAML authentication method types and is usually a value such as the email, username, or upn for example. Whenever a user logs in with a token, these authentication types issue a \"sub claim\" that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization.
	UniqueIdentifier *string `json:"unique-identifier,omitempty"`
}

// NewCreateAuthMethodLDAP instantiates a new CreateAuthMethodLDAP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthMethodLDAP(name string) *CreateAuthMethodLDAP {
	this := CreateAuthMethodLDAP{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var genKey string = "true"
	this.GenKey = &genKey
	var json bool = false
	this.Json = &json
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	this.Name = name
	var uniqueIdentifier string = "users"
	this.UniqueIdentifier = &uniqueIdentifier
	return &this
}

// NewCreateAuthMethodLDAPWithDefaults instantiates a new CreateAuthMethodLDAP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthMethodLDAPWithDefaults() *CreateAuthMethodLDAP {
	this := CreateAuthMethodLDAP{}
	var accessExpires int64 = 0
	this.AccessExpires = &accessExpires
	var genKey string = "true"
	this.GenKey = &genKey
	var json bool = false
	this.Json = &json
	var jwtTtl int64 = 0
	this.JwtTtl = &jwtTtl
	var uniqueIdentifier string = "users"
	this.UniqueIdentifier = &uniqueIdentifier
	return &this
}

// GetAccessExpires returns the AccessExpires field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetAccessExpires() int64 {
	if o == nil || IsNil(o.AccessExpires) {
		var ret int64
		return ret
	}
	return *o.AccessExpires
}

// GetAccessExpiresOk returns a tuple with the AccessExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetAccessExpiresOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessExpires) {
		return nil, false
	}
	return o.AccessExpires, true
}

// HasAccessExpires returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasAccessExpires() bool {
	if o != nil && !IsNil(o.AccessExpires) {
		return true
	}

	return false
}

// SetAccessExpires gets a reference to the given int64 and assigns it to the AccessExpires field.
func (o *CreateAuthMethodLDAP) SetAccessExpires(v int64) {
	o.AccessExpires = &v
}

// GetBoundIps returns the BoundIps field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetBoundIps() []string {
	if o == nil || IsNil(o.BoundIps) {
		var ret []string
		return ret
	}
	return o.BoundIps
}

// GetBoundIpsOk returns a tuple with the BoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetBoundIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.BoundIps) {
		return nil, false
	}
	return o.BoundIps, true
}

// HasBoundIps returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasBoundIps() bool {
	if o != nil && !IsNil(o.BoundIps) {
		return true
	}

	return false
}

// SetBoundIps gets a reference to the given []string and assigns it to the BoundIps field.
func (o *CreateAuthMethodLDAP) SetBoundIps(v []string) {
	o.BoundIps = v
}

// GetForceSubClaims returns the ForceSubClaims field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetForceSubClaims() bool {
	if o == nil || IsNil(o.ForceSubClaims) {
		var ret bool
		return ret
	}
	return *o.ForceSubClaims
}

// GetForceSubClaimsOk returns a tuple with the ForceSubClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetForceSubClaimsOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceSubClaims) {
		return nil, false
	}
	return o.ForceSubClaims, true
}

// HasForceSubClaims returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasForceSubClaims() bool {
	if o != nil && !IsNil(o.ForceSubClaims) {
		return true
	}

	return false
}

// SetForceSubClaims gets a reference to the given bool and assigns it to the ForceSubClaims field.
func (o *CreateAuthMethodLDAP) SetForceSubClaims(v bool) {
	o.ForceSubClaims = &v
}

// GetGenKey returns the GenKey field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetGenKey() string {
	if o == nil || IsNil(o.GenKey) {
		var ret string
		return ret
	}
	return *o.GenKey
}

// GetGenKeyOk returns a tuple with the GenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetGenKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GenKey) {
		return nil, false
	}
	return o.GenKey, true
}

// HasGenKey returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasGenKey() bool {
	if o != nil && !IsNil(o.GenKey) {
		return true
	}

	return false
}

// SetGenKey gets a reference to the given string and assigns it to the GenKey field.
func (o *CreateAuthMethodLDAP) SetGenKey(v string) {
	o.GenKey = &v
}

// GetGwBoundIps returns the GwBoundIps field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetGwBoundIps() []string {
	if o == nil || IsNil(o.GwBoundIps) {
		var ret []string
		return ret
	}
	return o.GwBoundIps
}

// GetGwBoundIpsOk returns a tuple with the GwBoundIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetGwBoundIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.GwBoundIps) {
		return nil, false
	}
	return o.GwBoundIps, true
}

// HasGwBoundIps returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasGwBoundIps() bool {
	if o != nil && !IsNil(o.GwBoundIps) {
		return true
	}

	return false
}

// SetGwBoundIps gets a reference to the given []string and assigns it to the GwBoundIps field.
func (o *CreateAuthMethodLDAP) SetGwBoundIps(v []string) {
	o.GwBoundIps = v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateAuthMethodLDAP) SetJson(v bool) {
	o.Json = &v
}

// GetJwtTtl returns the JwtTtl field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetJwtTtl() int64 {
	if o == nil || IsNil(o.JwtTtl) {
		var ret int64
		return ret
	}
	return *o.JwtTtl
}

// GetJwtTtlOk returns a tuple with the JwtTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetJwtTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.JwtTtl) {
		return nil, false
	}
	return o.JwtTtl, true
}

// HasJwtTtl returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasJwtTtl() bool {
	if o != nil && !IsNil(o.JwtTtl) {
		return true
	}

	return false
}

// SetJwtTtl gets a reference to the given int64 and assigns it to the JwtTtl field.
func (o *CreateAuthMethodLDAP) SetJwtTtl(v int64) {
	o.JwtTtl = &v
}

// GetName returns the Name field value
func (o *CreateAuthMethodLDAP) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAuthMethodLDAP) SetName(v string) {
	o.Name = v
}

// GetPublicKeyData returns the PublicKeyData field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetPublicKeyData() string {
	if o == nil || IsNil(o.PublicKeyData) {
		var ret string
		return ret
	}
	return *o.PublicKeyData
}

// GetPublicKeyDataOk returns a tuple with the PublicKeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetPublicKeyDataOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKeyData) {
		return nil, false
	}
	return o.PublicKeyData, true
}

// HasPublicKeyData returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasPublicKeyData() bool {
	if o != nil && !IsNil(o.PublicKeyData) {
		return true
	}

	return false
}

// SetPublicKeyData gets a reference to the given string and assigns it to the PublicKeyData field.
func (o *CreateAuthMethodLDAP) SetPublicKeyData(v string) {
	o.PublicKeyData = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateAuthMethodLDAP) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateAuthMethodLDAP) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUniqueIdentifier returns the UniqueIdentifier field value if set, zero value otherwise.
func (o *CreateAuthMethodLDAP) GetUniqueIdentifier() string {
	if o == nil || IsNil(o.UniqueIdentifier) {
		var ret string
		return ret
	}
	return *o.UniqueIdentifier
}

// GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthMethodLDAP) GetUniqueIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueIdentifier) {
		return nil, false
	}
	return o.UniqueIdentifier, true
}

// HasUniqueIdentifier returns a boolean if a field has been set.
func (o *CreateAuthMethodLDAP) HasUniqueIdentifier() bool {
	if o != nil && !IsNil(o.UniqueIdentifier) {
		return true
	}

	return false
}

// SetUniqueIdentifier gets a reference to the given string and assigns it to the UniqueIdentifier field.
func (o *CreateAuthMethodLDAP) SetUniqueIdentifier(v string) {
	o.UniqueIdentifier = &v
}

func (o CreateAuthMethodLDAP) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAuthMethodLDAP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessExpires) {
		toSerialize["access-expires"] = o.AccessExpires
	}
	if !IsNil(o.BoundIps) {
		toSerialize["bound-ips"] = o.BoundIps
	}
	if !IsNil(o.ForceSubClaims) {
		toSerialize["force-sub-claims"] = o.ForceSubClaims
	}
	if !IsNil(o.GenKey) {
		toSerialize["gen-key"] = o.GenKey
	}
	if !IsNil(o.GwBoundIps) {
		toSerialize["gw-bound-ips"] = o.GwBoundIps
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.JwtTtl) {
		toSerialize["jwt-ttl"] = o.JwtTtl
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PublicKeyData) {
		toSerialize["public-key-data"] = o.PublicKeyData
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	if !IsNil(o.UniqueIdentifier) {
		toSerialize["unique-identifier"] = o.UniqueIdentifier
	}
	return toSerialize, nil
}

type NullableCreateAuthMethodLDAP struct {
	value *CreateAuthMethodLDAP
	isSet bool
}

func (v NullableCreateAuthMethodLDAP) Get() *CreateAuthMethodLDAP {
	return v.value
}

func (v *NullableCreateAuthMethodLDAP) Set(val *CreateAuthMethodLDAP) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthMethodLDAP) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthMethodLDAP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthMethodLDAP(val *CreateAuthMethodLDAP) *NullableCreateAuthMethodLDAP {
	return &NullableCreateAuthMethodLDAP{value: val, isSet: true}
}

func (v NullableCreateAuthMethodLDAP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthMethodLDAP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


