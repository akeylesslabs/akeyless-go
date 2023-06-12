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

// checks if the UpdateAWSTargetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAWSTargetDetails{}

// UpdateAWSTargetDetails struct for UpdateAWSTargetDetails
type UpdateAWSTargetDetails struct {
	// The aws secret access key
	AccessKey *string `json:"access-key,omitempty"`
	// The aws access key id
	AccessKeyId *string `json:"access-key-id,omitempty"`
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
	// The aws region
	Region *string `json:"region,omitempty"`
	// Required only for temporary security credentials retrieved via STS, otherwise it can be an empty string
	SessionToken *string `json:"session-token,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewUpdateAWSTargetDetails instantiates a new UpdateAWSTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAWSTargetDetails(name string) *UpdateAWSTargetDetails {
	this := UpdateAWSTargetDetails{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var region string = "us-east-2"
	this.Region = &region
	return &this
}

// NewUpdateAWSTargetDetailsWithDefaults instantiates a new UpdateAWSTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAWSTargetDetailsWithDefaults() *UpdateAWSTargetDetails {
	this := UpdateAWSTargetDetails{}
	var json bool = false
	this.Json = &json
	var region string = "us-east-2"
	this.Region = &region
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *UpdateAWSTargetDetails) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetAccessKeyId() string {
	if o == nil || IsNil(o.AccessKeyId) {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKeyId) {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasAccessKeyId() bool {
	if o != nil && !IsNil(o.AccessKeyId) {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *UpdateAWSTargetDetails) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateAWSTargetDetails) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetKeepPrevVersion() string {
	if o == nil || IsNil(o.KeepPrevVersion) {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KeepPrevVersion) {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasKeepPrevVersion() bool {
	if o != nil && !IsNil(o.KeepPrevVersion) {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateAWSTargetDetails) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetName returns the Name field value
func (o *UpdateAWSTargetDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateAWSTargetDetails) SetName(v string) {
	o.Name = v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetNewVersion() bool {
	if o == nil || IsNil(o.NewVersion) {
		var ret bool
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetNewVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.NewVersion) {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasNewVersion() bool {
	if o != nil && !IsNil(o.NewVersion) {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given bool and assigns it to the NewVersion field.
func (o *UpdateAWSTargetDetails) SetNewVersion(v bool) {
	o.NewVersion = &v
}

// GetProtectionKey returns the ProtectionKey field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetProtectionKey() string {
	if o == nil || IsNil(o.ProtectionKey) {
		var ret string
		return ret
	}
	return *o.ProtectionKey
}

// GetProtectionKeyOk returns a tuple with the ProtectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetProtectionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProtectionKey) {
		return nil, false
	}
	return o.ProtectionKey, true
}

// HasProtectionKey returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasProtectionKey() bool {
	if o != nil && !IsNil(o.ProtectionKey) {
		return true
	}

	return false
}

// SetProtectionKey gets a reference to the given string and assigns it to the ProtectionKey field.
func (o *UpdateAWSTargetDetails) SetProtectionKey(v string) {
	o.ProtectionKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *UpdateAWSTargetDetails) SetRegion(v string) {
	o.Region = &v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetSessionToken() string {
	if o == nil || IsNil(o.SessionToken) {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetSessionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SessionToken) {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasSessionToken() bool {
	if o != nil && !IsNil(o.SessionToken) {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *UpdateAWSTargetDetails) SetSessionToken(v string) {
	o.SessionToken = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateAWSTargetDetails) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateAWSTargetDetails) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAWSTargetDetails) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateAWSTargetDetails) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateAWSTargetDetails) SetUidToken(v string) {
	o.UidToken = &v
}

func (o UpdateAWSTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAWSTargetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKey) {
		toSerialize["access-key"] = o.AccessKey
	}
	if !IsNil(o.AccessKeyId) {
		toSerialize["access-key-id"] = o.AccessKeyId
	}
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
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.SessionToken) {
		toSerialize["session-token"] = o.SessionToken
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	return toSerialize, nil
}

type NullableUpdateAWSTargetDetails struct {
	value *UpdateAWSTargetDetails
	isSet bool
}

func (v NullableUpdateAWSTargetDetails) Get() *UpdateAWSTargetDetails {
	return v.value
}

func (v *NullableUpdateAWSTargetDetails) Set(val *UpdateAWSTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAWSTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAWSTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAWSTargetDetails(val *UpdateAWSTargetDetails) *NullableUpdateAWSTargetDetails {
	return &NullableUpdateAWSTargetDetails{value: val, isSet: true}
}

func (v NullableUpdateAWSTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAWSTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


