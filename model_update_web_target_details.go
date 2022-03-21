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

// UpdateWebTargetDetails struct for UpdateWebTargetDetails
type UpdateWebTargetDetails struct {
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// Target name
	Name string `json:"name"`
	// Deprecated
	NewVersion *bool `json:"new-version,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	ProtectionKey *string `json:"protection_key,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewUpdateWebTargetDetails instantiates a new UpdateWebTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebTargetDetails(name string, ) *UpdateWebTargetDetails {
	this := UpdateWebTargetDetails{}
	this.Name = name
	return &this
}

// NewUpdateWebTargetDetailsWithDefaults instantiates a new UpdateWebTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebTargetDetailsWithDefaults() *UpdateWebTargetDetails {
	this := UpdateWebTargetDetails{}
	return &this
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateWebTargetDetails) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebTargetDetails) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateWebTargetDetails) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateWebTargetDetails) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetName returns the Name field value
func (o *UpdateWebTargetDetails) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateWebTargetDetails) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateWebTargetDetails) SetName(v string) {
	o.Name = v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *UpdateWebTargetDetails) GetNewVersion() bool {
	if o == nil || o.NewVersion == nil {
		var ret bool
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebTargetDetails) GetNewVersionOk() (*bool, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *UpdateWebTargetDetails) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given bool and assigns it to the NewVersion field.
func (o *UpdateWebTargetDetails) SetNewVersion(v bool) {
	o.NewVersion = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *UpdateWebTargetDetails) GetProtectionKey() string {
	if o == nil || o.ProtectionKey == nil {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebTargetDetails) GetProtectionKeyOk() (*string, bool) {
	if o == nil || o.ProtectionKey == nil {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *UpdateWebTargetDetails) HasProtectionKey() bool {
	if o != nil && o.ProtectionKey != nil {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *UpdateWebTargetDetails) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateWebTargetDetails) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebTargetDetails) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateWebTargetDetails) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateWebTargetDetails) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateWebTargetDetails) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebTargetDetails) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateWebTargetDetails) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateWebTargetDetails) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateWebTargetDetails) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebTargetDetails) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateWebTargetDetails) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateWebTargetDetails) SetUrl(v string) {
	o.Url = &v
}

func (o UpdateWebTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewVersion != nil {
		toSerialize["new-version"] = o.NewVersion
	}
	if o.ProtectionKey != nil {
		toSerialize["protection_key"] = o.ProtectionKey
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWebTargetDetails struct {
	value *UpdateWebTargetDetails
	isSet bool
}

func (v NullableUpdateWebTargetDetails) Get() *UpdateWebTargetDetails {
	return v.value
}

func (v *NullableUpdateWebTargetDetails) Set(val *UpdateWebTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebTargetDetails(val *UpdateWebTargetDetails) *NullableUpdateWebTargetDetails {
	return &NullableUpdateWebTargetDetails{value: val, isSet: true}
}

func (v NullableUpdateWebTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


