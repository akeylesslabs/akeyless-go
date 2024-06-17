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

// PingTargetDetails PingTargetDetails
type PingTargetDetails struct {
	AdministrativePort *string `json:"administrative_port,omitempty"`
	AuthorizationPort *string `json:"authorization_port,omitempty"`
	PingUrl *string `json:"ping_url,omitempty"`
	PrivilegedUser *string `json:"privileged_user,omitempty"`
	UserPassword *string `json:"user_password,omitempty"`
}

// NewPingTargetDetails instantiates a new PingTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingTargetDetails() *PingTargetDetails {
	this := PingTargetDetails{}
	return &this
}

// NewPingTargetDetailsWithDefaults instantiates a new PingTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingTargetDetailsWithDefaults() *PingTargetDetails {
	this := PingTargetDetails{}
	return &this
}

// GetAdministrativePort returns the AdministrativePort field value if set, zero value otherwise.
func (o *PingTargetDetails) GetAdministrativePort() string {
	if o == nil || o.AdministrativePort == nil {
		var ret string
		return ret
	}
	return *o.AdministrativePort
}

// GetAdministrativePortOk returns a tuple with the AdministrativePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingTargetDetails) GetAdministrativePortOk() (*string, bool) {
	if o == nil || o.AdministrativePort == nil {
		return nil, false
	}
	return o.AdministrativePort, true
}

// HasAdministrativePort returns a boolean if a field has been set.
func (o *PingTargetDetails) HasAdministrativePort() bool {
	if o != nil && o.AdministrativePort != nil {
		return true
	}

	return false
}

// SetAdministrativePort gets a reference to the given string and assigns it to the AdministrativePort field.
func (o *PingTargetDetails) SetAdministrativePort(v string) {
	o.AdministrativePort = &v
}

// GetAuthorizationPort returns the AuthorizationPort field value if set, zero value otherwise.
func (o *PingTargetDetails) GetAuthorizationPort() string {
	if o == nil || o.AuthorizationPort == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationPort
}

// GetAuthorizationPortOk returns a tuple with the AuthorizationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingTargetDetails) GetAuthorizationPortOk() (*string, bool) {
	if o == nil || o.AuthorizationPort == nil {
		return nil, false
	}
	return o.AuthorizationPort, true
}

// HasAuthorizationPort returns a boolean if a field has been set.
func (o *PingTargetDetails) HasAuthorizationPort() bool {
	if o != nil && o.AuthorizationPort != nil {
		return true
	}

	return false
}

// SetAuthorizationPort gets a reference to the given string and assigns it to the AuthorizationPort field.
func (o *PingTargetDetails) SetAuthorizationPort(v string) {
	o.AuthorizationPort = &v
}

// GetPingUrl returns the PingUrl field value if set, zero value otherwise.
func (o *PingTargetDetails) GetPingUrl() string {
	if o == nil || o.PingUrl == nil {
		var ret string
		return ret
	}
	return *o.PingUrl
}

// GetPingUrlOk returns a tuple with the PingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingTargetDetails) GetPingUrlOk() (*string, bool) {
	if o == nil || o.PingUrl == nil {
		return nil, false
	}
	return o.PingUrl, true
}

// HasPingUrl returns a boolean if a field has been set.
func (o *PingTargetDetails) HasPingUrl() bool {
	if o != nil && o.PingUrl != nil {
		return true
	}

	return false
}

// SetPingUrl gets a reference to the given string and assigns it to the PingUrl field.
func (o *PingTargetDetails) SetPingUrl(v string) {
	o.PingUrl = &v
}

// GetPrivilegedUser returns the PrivilegedUser field value if set, zero value otherwise.
func (o *PingTargetDetails) GetPrivilegedUser() string {
	if o == nil || o.PrivilegedUser == nil {
		var ret string
		return ret
	}
	return *o.PrivilegedUser
}

// GetPrivilegedUserOk returns a tuple with the PrivilegedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingTargetDetails) GetPrivilegedUserOk() (*string, bool) {
	if o == nil || o.PrivilegedUser == nil {
		return nil, false
	}
	return o.PrivilegedUser, true
}

// HasPrivilegedUser returns a boolean if a field has been set.
func (o *PingTargetDetails) HasPrivilegedUser() bool {
	if o != nil && o.PrivilegedUser != nil {
		return true
	}

	return false
}

// SetPrivilegedUser gets a reference to the given string and assigns it to the PrivilegedUser field.
func (o *PingTargetDetails) SetPrivilegedUser(v string) {
	o.PrivilegedUser = &v
}

// GetUserPassword returns the UserPassword field value if set, zero value otherwise.
func (o *PingTargetDetails) GetUserPassword() string {
	if o == nil || o.UserPassword == nil {
		var ret string
		return ret
	}
	return *o.UserPassword
}

// GetUserPasswordOk returns a tuple with the UserPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingTargetDetails) GetUserPasswordOk() (*string, bool) {
	if o == nil || o.UserPassword == nil {
		return nil, false
	}
	return o.UserPassword, true
}

// HasUserPassword returns a boolean if a field has been set.
func (o *PingTargetDetails) HasUserPassword() bool {
	if o != nil && o.UserPassword != nil {
		return true
	}

	return false
}

// SetUserPassword gets a reference to the given string and assigns it to the UserPassword field.
func (o *PingTargetDetails) SetUserPassword(v string) {
	o.UserPassword = &v
}

func (o PingTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdministrativePort != nil {
		toSerialize["administrative_port"] = o.AdministrativePort
	}
	if o.AuthorizationPort != nil {
		toSerialize["authorization_port"] = o.AuthorizationPort
	}
	if o.PingUrl != nil {
		toSerialize["ping_url"] = o.PingUrl
	}
	if o.PrivilegedUser != nil {
		toSerialize["privileged_user"] = o.PrivilegedUser
	}
	if o.UserPassword != nil {
		toSerialize["user_password"] = o.UserPassword
	}
	return json.Marshal(toSerialize)
}

type NullablePingTargetDetails struct {
	value *PingTargetDetails
	isSet bool
}

func (v NullablePingTargetDetails) Get() *PingTargetDetails {
	return v.value
}

func (v *NullablePingTargetDetails) Set(val *PingTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePingTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePingTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingTargetDetails(val *PingTargetDetails) *NullablePingTargetDetails {
	return &NullablePingTargetDetails{value: val, isSet: true}
}

func (v NullablePingTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

