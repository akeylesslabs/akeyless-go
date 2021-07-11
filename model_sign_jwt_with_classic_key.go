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

// SignJWTWithClassicKey struct for SignJWTWithClassicKey
type SignJWTWithClassicKey struct {
	// The name of the key to use in the sign JWT process
	DisplayId string `json:"display-id"`
	// JWTClaims
	JwtClaims string `json:"jwt-claims"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// SigningMethod
	SigningMethod string `json:"signing-method"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
	// classic key version
	Version int32 `json:"version"`
}

// NewSignJWTWithClassicKey instantiates a new SignJWTWithClassicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignJWTWithClassicKey(displayId string, jwtClaims string, signingMethod string, version int32, ) *SignJWTWithClassicKey {
	this := SignJWTWithClassicKey{}
	this.DisplayId = displayId
	this.JwtClaims = jwtClaims
	this.SigningMethod = signingMethod
	this.Version = version
	return &this
}

// NewSignJWTWithClassicKeyWithDefaults instantiates a new SignJWTWithClassicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignJWTWithClassicKeyWithDefaults() *SignJWTWithClassicKey {
	this := SignJWTWithClassicKey{}
	return &this
}

// GetDisplayId returns the DisplayId field value
func (o *SignJWTWithClassicKey) GetDisplayId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value
// and a boolean to check if the value has been set.
func (o *SignJWTWithClassicKey) GetDisplayIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayId, true
}

// SetDisplayId sets field value
func (o *SignJWTWithClassicKey) SetDisplayId(v string) {
	o.DisplayId = v
}

// GetJwtClaims returns the JwtClaims field value
func (o *SignJWTWithClassicKey) GetJwtClaims() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.JwtClaims
}

// GetJwtClaimsOk returns a tuple with the JwtClaims field value
// and a boolean to check if the value has been set.
func (o *SignJWTWithClassicKey) GetJwtClaimsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.JwtClaims, true
}

// SetJwtClaims sets field value
func (o *SignJWTWithClassicKey) SetJwtClaims(v string) {
	o.JwtClaims = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SignJWTWithClassicKey) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignJWTWithClassicKey) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SignJWTWithClassicKey) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SignJWTWithClassicKey) SetPassword(v string) {
	o.Password = &v
}

// GetSigningMethod returns the SigningMethod field value
func (o *SignJWTWithClassicKey) GetSigningMethod() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SigningMethod
}

// GetSigningMethodOk returns a tuple with the SigningMethod field value
// and a boolean to check if the value has been set.
func (o *SignJWTWithClassicKey) GetSigningMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SigningMethod, true
}

// SetSigningMethod sets field value
func (o *SignJWTWithClassicKey) SetSigningMethod(v string) {
	o.SigningMethod = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SignJWTWithClassicKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignJWTWithClassicKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SignJWTWithClassicKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SignJWTWithClassicKey) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *SignJWTWithClassicKey) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignJWTWithClassicKey) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *SignJWTWithClassicKey) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *SignJWTWithClassicKey) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SignJWTWithClassicKey) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignJWTWithClassicKey) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SignJWTWithClassicKey) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SignJWTWithClassicKey) SetUsername(v string) {
	o.Username = &v
}

// GetVersion returns the Version field value
func (o *SignJWTWithClassicKey) GetVersion() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SignJWTWithClassicKey) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SignJWTWithClassicKey) SetVersion(v int32) {
	o.Version = v
}

func (o SignJWTWithClassicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["display-id"] = o.DisplayId
	}
	if true {
		toSerialize["jwt-claims"] = o.JwtClaims
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["signing-method"] = o.SigningMethod
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSignJWTWithClassicKey struct {
	value *SignJWTWithClassicKey
	isSet bool
}

func (v NullableSignJWTWithClassicKey) Get() *SignJWTWithClassicKey {
	return v.value
}

func (v *NullableSignJWTWithClassicKey) Set(val *SignJWTWithClassicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSignJWTWithClassicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSignJWTWithClassicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignJWTWithClassicKey(val *SignJWTWithClassicKey) *NullableSignJWTWithClassicKey {
	return &NullableSignJWTWithClassicKey{value: val, isSet: true}
}

func (v NullableSignJWTWithClassicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignJWTWithClassicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


