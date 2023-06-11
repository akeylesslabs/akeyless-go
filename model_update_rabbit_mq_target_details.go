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

// checks if the UpdateRabbitMQTargetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRabbitMQTargetDetails{}

// UpdateRabbitMQTargetDetails struct for UpdateRabbitMQTargetDetails
type UpdateRabbitMQTargetDetails struct {
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Whether to keep previous version [true/false]. If not set, use default according to account settings
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// Target name
	Name string `json:"name"`
	// Deprecated
	NewVersion *bool `json:"new-version,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	ProtectionKey *string `json:"protection_key,omitempty"`
	RabbitmqServerPassword *string `json:"rabbitmq_server_password,omitempty"`
	RabbitmqServerUri *string `json:"rabbitmq_server_uri,omitempty"`
	RabbitmqServerUser *string `json:"rabbitmq_server_user,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateRabbitMQTargetDetails instantiates a new UpdateRabbitMQTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRabbitMQTargetDetails(name string) *UpdateRabbitMQTargetDetails {
	this := UpdateRabbitMQTargetDetails{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewUpdateRabbitMQTargetDetailsWithDefaults instantiates a new UpdateRabbitMQTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRabbitMQTargetDetailsWithDefaults() *UpdateRabbitMQTargetDetails {
	this := UpdateRabbitMQTargetDetails{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetDetails) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetDetails) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateRabbitMQTargetDetails) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetDetails) GetKeepPrevVersion() string {
	if o == nil || IsNil(o.KeepPrevVersion) {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KeepPrevVersion) {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetDetails) HasKeepPrevVersion() bool {
	if o != nil && !IsNil(o.KeepPrevVersion) {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateRabbitMQTargetDetails) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetName returns the Name field value
func (o *UpdateRabbitMQTargetDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateRabbitMQTargetDetails) SetName(v string) {
	o.Name = v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetDetails) GetNewVersion() bool {
	if o == nil || IsNil(o.NewVersion) {
		var ret bool
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetNewVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.NewVersion) {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetDetails) HasNewVersion() bool {
	if o != nil && !IsNil(o.NewVersion) {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given bool and assigns it to the NewVersion field.
func (o *UpdateRabbitMQTargetDetails) SetNewVersion(v bool) {
	o.NewVersion = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetDetails) GetProtectionKey() string {
	if o == nil || IsNil(o.ProtectionKey) {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetProtectionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProtectionKey) {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetDetails) HasProtectionKey() bool {
	if o != nil && !IsNil(o.ProtectionKey) {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *UpdateRabbitMQTargetDetails) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetRabbitmqServerPassword returns the RabbitmqServerPassword field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerPassword() string {
	if o == nil || IsNil(o.RabbitmqServerPassword) {
		var ret string
		return ret
	}
	return *o.RabbitmqServerPassword
}

// GetRabbitmqServerPasswordOk returns a tuple with the RabbitmqServerPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqServerPassword) {
		return nil, false
	}
	return o.RabbitmqServerPassword, true
}

// HasRabbitmqServerPassword returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetDetails) HasRabbitmqServerPassword() bool {
	if o != nil && !IsNil(o.RabbitmqServerPassword) {
		return true
	}

	return false
}

// SetRabbitmqServerPassword gets a reference to the given string and assigns it to the RabbitmqServerPassword field.
func (o *UpdateRabbitMQTargetDetails) SetRabbitmqServerPassword(v string) {
	o.RabbitmqServerPassword = &v
}

// GetRabbitmqServerUri returns the RabbitmqServerUri field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerUri() string {
	if o == nil || IsNil(o.RabbitmqServerUri) {
		var ret string
		return ret
	}
	return *o.RabbitmqServerUri
}

// GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerUriOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqServerUri) {
		return nil, false
	}
	return o.RabbitmqServerUri, true
}

// HasRabbitmqServerUri returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetDetails) HasRabbitmqServerUri() bool {
	if o != nil && !IsNil(o.RabbitmqServerUri) {
		return true
	}

	return false
}

// SetRabbitmqServerUri gets a reference to the given string and assigns it to the RabbitmqServerUri field.
func (o *UpdateRabbitMQTargetDetails) SetRabbitmqServerUri(v string) {
	o.RabbitmqServerUri = &v
}

// GetRabbitmqServerUser returns the RabbitmqServerUser field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerUser() string {
	if o == nil || IsNil(o.RabbitmqServerUser) {
		var ret string
		return ret
	}
	return *o.RabbitmqServerUser
}

// GetRabbitmqServerUserOk returns a tuple with the RabbitmqServerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetRabbitmqServerUserOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqServerUser) {
		return nil, false
	}
	return o.RabbitmqServerUser, true
}

// HasRabbitmqServerUser returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetDetails) HasRabbitmqServerUser() bool {
	if o != nil && !IsNil(o.RabbitmqServerUser) {
		return true
	}

	return false
}

// SetRabbitmqServerUser gets a reference to the given string and assigns it to the RabbitmqServerUser field.
func (o *UpdateRabbitMQTargetDetails) SetRabbitmqServerUser(v string) {
	o.RabbitmqServerUser = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetDetails) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetDetails) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateRabbitMQTargetDetails) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateRabbitMQTargetDetails) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRabbitMQTargetDetails) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateRabbitMQTargetDetails) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateRabbitMQTargetDetails) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateRabbitMQTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRabbitMQTargetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.KeepPrevVersion) {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NewVersion) {
		toSerialize["new-version"] = o.NewVersion
	}
	if !IsNil(o.ProtectionKey) {
		toSerialize["protection_key"] = o.ProtectionKey
	}
	if !IsNil(o.RabbitmqServerPassword) {
		toSerialize["rabbitmq_server_password"] = o.RabbitmqServerPassword
	}
	if !IsNil(o.RabbitmqServerUri) {
		toSerialize["rabbitmq_server_uri"] = o.RabbitmqServerUri
	}
	if !IsNil(o.RabbitmqServerUser) {
		toSerialize["rabbitmq_server_user"] = o.RabbitmqServerUser
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableUpdateRabbitMQTargetDetails struct {
	value *UpdateRabbitMQTargetDetails
	isSet bool
}

func (v NullableUpdateRabbitMQTargetDetails) Get() *UpdateRabbitMQTargetDetails {
	return v.value
}

func (v *NullableUpdateRabbitMQTargetDetails) Set(val *UpdateRabbitMQTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRabbitMQTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRabbitMQTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRabbitMQTargetDetails(val *UpdateRabbitMQTargetDetails) *NullableUpdateRabbitMQTargetDetails {
	return &NullableUpdateRabbitMQTargetDetails{value: val, isSet: true}
}

func (v NullableUpdateRabbitMQTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRabbitMQTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


