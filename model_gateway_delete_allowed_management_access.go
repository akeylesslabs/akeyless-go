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

// GatewayDeleteAllowedManagementAccess gatewayDeleteAllowedManagementAccess is a command that deletes sub-admins
type GatewayDeleteAllowedManagementAccess struct {
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// SubAdminID to be removed
	SubAdminId string `json:"sub-admin-id"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewGatewayDeleteAllowedManagementAccess instantiates a new GatewayDeleteAllowedManagementAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayDeleteAllowedManagementAccess(subAdminId string, ) *GatewayDeleteAllowedManagementAccess {
	this := GatewayDeleteAllowedManagementAccess{}
	this.SubAdminId = subAdminId
	return &this
}

// NewGatewayDeleteAllowedManagementAccessWithDefaults instantiates a new GatewayDeleteAllowedManagementAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayDeleteAllowedManagementAccessWithDefaults() *GatewayDeleteAllowedManagementAccess {
	this := GatewayDeleteAllowedManagementAccess{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayDeleteAllowedManagementAccess) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteAllowedManagementAccess) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayDeleteAllowedManagementAccess) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayDeleteAllowedManagementAccess) SetPassword(v string) {
	o.Password = &v
}

// GetSubAdminId returns the SubAdminId field value
func (o *GatewayDeleteAllowedManagementAccess) GetSubAdminId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SubAdminId
}

// GetSubAdminIdOk returns a tuple with the SubAdminId field value
// and a boolean to check if the value has been set.
func (o *GatewayDeleteAllowedManagementAccess) GetSubAdminIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubAdminId, true
}

// SetSubAdminId sets field value
func (o *GatewayDeleteAllowedManagementAccess) SetSubAdminId(v string) {
	o.SubAdminId = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayDeleteAllowedManagementAccess) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteAllowedManagementAccess) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayDeleteAllowedManagementAccess) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayDeleteAllowedManagementAccess) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayDeleteAllowedManagementAccess) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteAllowedManagementAccess) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayDeleteAllowedManagementAccess) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayDeleteAllowedManagementAccess) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayDeleteAllowedManagementAccess) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayDeleteAllowedManagementAccess) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayDeleteAllowedManagementAccess) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayDeleteAllowedManagementAccess) SetUsername(v string) {
	o.Username = &v
}

func (o GatewayDeleteAllowedManagementAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["sub-admin-id"] = o.SubAdminId
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
	return json.Marshal(toSerialize)
}

type NullableGatewayDeleteAllowedManagementAccess struct {
	value *GatewayDeleteAllowedManagementAccess
	isSet bool
}

func (v NullableGatewayDeleteAllowedManagementAccess) Get() *GatewayDeleteAllowedManagementAccess {
	return v.value
}

func (v *NullableGatewayDeleteAllowedManagementAccess) Set(val *GatewayDeleteAllowedManagementAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayDeleteAllowedManagementAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayDeleteAllowedManagementAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayDeleteAllowedManagementAccess(val *GatewayDeleteAllowedManagementAccess) *NullableGatewayDeleteAllowedManagementAccess {
	return &NullableGatewayDeleteAllowedManagementAccess{value: val, isSet: true}
}

func (v NullableGatewayDeleteAllowedManagementAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayDeleteAllowedManagementAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

