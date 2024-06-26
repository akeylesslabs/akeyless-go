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

// ChefTargetDetails ChefTargetDetails
type ChefTargetDetails struct {
	ChefServerHostName *string `json:"chef_server_host_name,omitempty"`
	ChefServerKey *string `json:"chef_server_key,omitempty"`
	ChefServerPort *string `json:"chef_server_port,omitempty"`
	ChefServerUrl *string `json:"chef_server_url,omitempty"`
	ChefServerUsername *string `json:"chef_server_username,omitempty"`
	ChefSkipSsl *bool `json:"chef_skip_ssl,omitempty"`
}

// NewChefTargetDetails instantiates a new ChefTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChefTargetDetails() *ChefTargetDetails {
	this := ChefTargetDetails{}
	return &this
}

// NewChefTargetDetailsWithDefaults instantiates a new ChefTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChefTargetDetailsWithDefaults() *ChefTargetDetails {
	this := ChefTargetDetails{}
	return &this
}

// GetChefServerHostName returns the ChefServerHostName field value if set, zero value otherwise.
func (o *ChefTargetDetails) GetChefServerHostName() string {
	if o == nil || o.ChefServerHostName == nil {
		var ret string
		return ret
	}
	return *o.ChefServerHostName
}

// GetChefServerHostNameOk returns a tuple with the ChefServerHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChefTargetDetails) GetChefServerHostNameOk() (*string, bool) {
	if o == nil || o.ChefServerHostName == nil {
		return nil, false
	}
	return o.ChefServerHostName, true
}

// HasChefServerHostName returns a boolean if a field has been set.
func (o *ChefTargetDetails) HasChefServerHostName() bool {
	if o != nil && o.ChefServerHostName != nil {
		return true
	}

	return false
}

// SetChefServerHostName gets a reference to the given string and assigns it to the ChefServerHostName field.
func (o *ChefTargetDetails) SetChefServerHostName(v string) {
	o.ChefServerHostName = &v
}

// GetChefServerKey returns the ChefServerKey field value if set, zero value otherwise.
func (o *ChefTargetDetails) GetChefServerKey() string {
	if o == nil || o.ChefServerKey == nil {
		var ret string
		return ret
	}
	return *o.ChefServerKey
}

// GetChefServerKeyOk returns a tuple with the ChefServerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChefTargetDetails) GetChefServerKeyOk() (*string, bool) {
	if o == nil || o.ChefServerKey == nil {
		return nil, false
	}
	return o.ChefServerKey, true
}

// HasChefServerKey returns a boolean if a field has been set.
func (o *ChefTargetDetails) HasChefServerKey() bool {
	if o != nil && o.ChefServerKey != nil {
		return true
	}

	return false
}

// SetChefServerKey gets a reference to the given string and assigns it to the ChefServerKey field.
func (o *ChefTargetDetails) SetChefServerKey(v string) {
	o.ChefServerKey = &v
}

// GetChefServerPort returns the ChefServerPort field value if set, zero value otherwise.
func (o *ChefTargetDetails) GetChefServerPort() string {
	if o == nil || o.ChefServerPort == nil {
		var ret string
		return ret
	}
	return *o.ChefServerPort
}

// GetChefServerPortOk returns a tuple with the ChefServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChefTargetDetails) GetChefServerPortOk() (*string, bool) {
	if o == nil || o.ChefServerPort == nil {
		return nil, false
	}
	return o.ChefServerPort, true
}

// HasChefServerPort returns a boolean if a field has been set.
func (o *ChefTargetDetails) HasChefServerPort() bool {
	if o != nil && o.ChefServerPort != nil {
		return true
	}

	return false
}

// SetChefServerPort gets a reference to the given string and assigns it to the ChefServerPort field.
func (o *ChefTargetDetails) SetChefServerPort(v string) {
	o.ChefServerPort = &v
}

// GetChefServerUrl returns the ChefServerUrl field value if set, zero value otherwise.
func (o *ChefTargetDetails) GetChefServerUrl() string {
	if o == nil || o.ChefServerUrl == nil {
		var ret string
		return ret
	}
	return *o.ChefServerUrl
}

// GetChefServerUrlOk returns a tuple with the ChefServerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChefTargetDetails) GetChefServerUrlOk() (*string, bool) {
	if o == nil || o.ChefServerUrl == nil {
		return nil, false
	}
	return o.ChefServerUrl, true
}

// HasChefServerUrl returns a boolean if a field has been set.
func (o *ChefTargetDetails) HasChefServerUrl() bool {
	if o != nil && o.ChefServerUrl != nil {
		return true
	}

	return false
}

// SetChefServerUrl gets a reference to the given string and assigns it to the ChefServerUrl field.
func (o *ChefTargetDetails) SetChefServerUrl(v string) {
	o.ChefServerUrl = &v
}

// GetChefServerUsername returns the ChefServerUsername field value if set, zero value otherwise.
func (o *ChefTargetDetails) GetChefServerUsername() string {
	if o == nil || o.ChefServerUsername == nil {
		var ret string
		return ret
	}
	return *o.ChefServerUsername
}

// GetChefServerUsernameOk returns a tuple with the ChefServerUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChefTargetDetails) GetChefServerUsernameOk() (*string, bool) {
	if o == nil || o.ChefServerUsername == nil {
		return nil, false
	}
	return o.ChefServerUsername, true
}

// HasChefServerUsername returns a boolean if a field has been set.
func (o *ChefTargetDetails) HasChefServerUsername() bool {
	if o != nil && o.ChefServerUsername != nil {
		return true
	}

	return false
}

// SetChefServerUsername gets a reference to the given string and assigns it to the ChefServerUsername field.
func (o *ChefTargetDetails) SetChefServerUsername(v string) {
	o.ChefServerUsername = &v
}

// GetChefSkipSsl returns the ChefSkipSsl field value if set, zero value otherwise.
func (o *ChefTargetDetails) GetChefSkipSsl() bool {
	if o == nil || o.ChefSkipSsl == nil {
		var ret bool
		return ret
	}
	return *o.ChefSkipSsl
}

// GetChefSkipSslOk returns a tuple with the ChefSkipSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChefTargetDetails) GetChefSkipSslOk() (*bool, bool) {
	if o == nil || o.ChefSkipSsl == nil {
		return nil, false
	}
	return o.ChefSkipSsl, true
}

// HasChefSkipSsl returns a boolean if a field has been set.
func (o *ChefTargetDetails) HasChefSkipSsl() bool {
	if o != nil && o.ChefSkipSsl != nil {
		return true
	}

	return false
}

// SetChefSkipSsl gets a reference to the given bool and assigns it to the ChefSkipSsl field.
func (o *ChefTargetDetails) SetChefSkipSsl(v bool) {
	o.ChefSkipSsl = &v
}

func (o ChefTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChefServerHostName != nil {
		toSerialize["chef_server_host_name"] = o.ChefServerHostName
	}
	if o.ChefServerKey != nil {
		toSerialize["chef_server_key"] = o.ChefServerKey
	}
	if o.ChefServerPort != nil {
		toSerialize["chef_server_port"] = o.ChefServerPort
	}
	if o.ChefServerUrl != nil {
		toSerialize["chef_server_url"] = o.ChefServerUrl
	}
	if o.ChefServerUsername != nil {
		toSerialize["chef_server_username"] = o.ChefServerUsername
	}
	if o.ChefSkipSsl != nil {
		toSerialize["chef_skip_ssl"] = o.ChefSkipSsl
	}
	return json.Marshal(toSerialize)
}

type NullableChefTargetDetails struct {
	value *ChefTargetDetails
	isSet bool
}

func (v NullableChefTargetDetails) Get() *ChefTargetDetails {
	return v.value
}

func (v *NullableChefTargetDetails) Set(val *ChefTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableChefTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableChefTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChefTargetDetails(val *ChefTargetDetails) *NullableChefTargetDetails {
	return &NullableChefTargetDetails{value: val, isSet: true}
}

func (v NullableChefTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChefTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


