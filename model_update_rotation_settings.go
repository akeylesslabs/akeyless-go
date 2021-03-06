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

// UpdateRotationSettings updateRotationSettings is a command that updates rotations settings of an existing key
type UpdateRotationSettings struct {
	// Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation
	AutoRotate bool `json:"auto-rotate"`
	// Key name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// The number of days to wait between every automatic key rotation (7-365)
	RotationInterval *int64 `json:"rotation-interval,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewUpdateRotationSettings instantiates a new UpdateRotationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRotationSettings(autoRotate bool, name string, ) *UpdateRotationSettings {
	this := UpdateRotationSettings{}
	this.AutoRotate = autoRotate
	this.Name = name
	return &this
}

// NewUpdateRotationSettingsWithDefaults instantiates a new UpdateRotationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRotationSettingsWithDefaults() *UpdateRotationSettings {
	this := UpdateRotationSettings{}
	return &this
}

// GetAutoRotate returns the AutoRotate field value
func (o *UpdateRotationSettings) GetAutoRotate() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value
// and a boolean to check if the value has been set.
func (o *UpdateRotationSettings) GetAutoRotateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AutoRotate, true
}

// SetAutoRotate sets field value
func (o *UpdateRotationSettings) SetAutoRotate(v bool) {
	o.AutoRotate = v
}

// GetName returns the Name field value
func (o *UpdateRotationSettings) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRotationSettings) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateRotationSettings) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateRotationSettings) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotationSettings) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateRotationSettings) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateRotationSettings) SetPassword(v string) {
	o.Password = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *UpdateRotationSettings) GetRotationInterval() int64 {
	if o == nil || o.RotationInterval == nil {
		var ret int64
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotationSettings) GetRotationIntervalOk() (*int64, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *UpdateRotationSettings) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given int64 and assigns it to the RotationInterval field.
func (o *UpdateRotationSettings) SetRotationInterval(v int64) {
	o.RotationInterval = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateRotationSettings) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotationSettings) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateRotationSettings) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateRotationSettings) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateRotationSettings) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotationSettings) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateRotationSettings) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateRotationSettings) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateRotationSettings) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotationSettings) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateRotationSettings) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateRotationSettings) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateRotationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auto-rotate"] = o.AutoRotate
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.RotationInterval != nil {
		toSerialize["rotation-interval"] = o.RotationInterval
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

type NullableUpdateRotationSettings struct {
	value *UpdateRotationSettings
	isSet bool
}

func (v NullableUpdateRotationSettings) Get() *UpdateRotationSettings {
	return v.value
}

func (v *NullableUpdateRotationSettings) Set(val *UpdateRotationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRotationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRotationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRotationSettings(val *UpdateRotationSettings) *NullableUpdateRotationSettings {
	return &NullableUpdateRotationSettings{value: val, isSet: true}
}

func (v NullableUpdateRotationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRotationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


