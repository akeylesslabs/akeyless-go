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

// UpdateGlobalSignTarget struct for UpdateGlobalSignTarget
type UpdateGlobalSignTarget struct {
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Email of the GlobalSign GCC account contact
	ContactEmail string `json:"contact-email"`
	// First name of the GlobalSign GCC account contact
	ContactFirstName string `json:"contact-first-name"`
	// Last name of the GlobalSign GCC account contact
	ContactLastName string `json:"contact-last-name"`
	// Telephone of the GlobalSign GCC account contact
	ContactPhone string `json:"contact-phone"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Whether to keep previous version [true/false]. If not set, use default according to account settings
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// New target name
	NewName *string `json:"new-name,omitempty"`
	// Password of the GlobalSign GCC account
	Password string `json:"password"`
	// Profile ID of the GlobalSign GCC account
	ProfileId string `json:"profile-id"`
	// Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h.
	Timeout *string `json:"timeout,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Deprecated
	UpdateVersion *bool `json:"update-version,omitempty"`
	// Username of the GlobalSign GCC account
	Username string `json:"username"`
}

// NewUpdateGlobalSignTarget instantiates a new UpdateGlobalSignTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGlobalSignTarget(contactEmail string, contactFirstName string, contactLastName string, contactPhone string, name string, password string, profileId string, username string) *UpdateGlobalSignTarget {
	this := UpdateGlobalSignTarget{}
	this.ContactEmail = contactEmail
	this.ContactFirstName = contactFirstName
	this.ContactLastName = contactLastName
	this.ContactPhone = contactPhone
	var json bool = false
	this.Json = &json
	this.Name = name
	this.Password = password
	this.ProfileId = profileId
	var timeout string = "5m"
	this.Timeout = &timeout
	this.Username = username
	return &this
}

// NewUpdateGlobalSignTargetWithDefaults instantiates a new UpdateGlobalSignTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGlobalSignTargetWithDefaults() *UpdateGlobalSignTarget {
	this := UpdateGlobalSignTarget{}
	var json bool = false
	this.Json = &json
	var timeout string = "5m"
	this.Timeout = &timeout
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateGlobalSignTarget) SetComment(v string) {
	o.Comment = &v
}

// GetContactEmail returns the ContactEmail field value
func (o *UpdateGlobalSignTarget) GetContactEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactEmail, true
}

// SetContactEmail sets field value
func (o *UpdateGlobalSignTarget) SetContactEmail(v string) {
	o.ContactEmail = v
}

// GetContactFirstName returns the ContactFirstName field value
func (o *UpdateGlobalSignTarget) GetContactFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactFirstName
}

// GetContactFirstNameOk returns a tuple with the ContactFirstName field value
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetContactFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactFirstName, true
}

// SetContactFirstName sets field value
func (o *UpdateGlobalSignTarget) SetContactFirstName(v string) {
	o.ContactFirstName = v
}

// GetContactLastName returns the ContactLastName field value
func (o *UpdateGlobalSignTarget) GetContactLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactLastName
}

// GetContactLastNameOk returns a tuple with the ContactLastName field value
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetContactLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactLastName, true
}

// SetContactLastName sets field value
func (o *UpdateGlobalSignTarget) SetContactLastName(v string) {
	o.ContactLastName = v
}

// GetContactPhone returns the ContactPhone field value
func (o *UpdateGlobalSignTarget) GetContactPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactPhone
}

// GetContactPhoneOk returns a tuple with the ContactPhone field value
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetContactPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactPhone, true
}

// SetContactPhone sets field value
func (o *UpdateGlobalSignTarget) SetContactPhone(v string) {
	o.ContactPhone = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateGlobalSignTarget) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateGlobalSignTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateGlobalSignTarget) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateGlobalSignTarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *UpdateGlobalSignTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateGlobalSignTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateGlobalSignTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetPassword returns the Password field value
func (o *UpdateGlobalSignTarget) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UpdateGlobalSignTarget) SetPassword(v string) {
	o.Password = v
}

// GetProfileId returns the ProfileId field value
func (o *UpdateGlobalSignTarget) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *UpdateGlobalSignTarget) SetProfileId(v string) {
	o.ProfileId = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetTimeout() string {
	if o == nil || o.Timeout == nil {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetTimeoutOk() (*string, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *UpdateGlobalSignTarget) SetTimeout(v string) {
	o.Timeout = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateGlobalSignTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateGlobalSignTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateGlobalSignTarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateGlobalSignTarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateGlobalSignTarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

// GetUsername returns the Username field value
func (o *UpdateGlobalSignTarget) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UpdateGlobalSignTarget) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UpdateGlobalSignTarget) SetUsername(v string) {
	o.Username = v
}

func (o UpdateGlobalSignTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["contact-email"] = o.ContactEmail
	}
	if true {
		toSerialize["contact-first-name"] = o.ContactFirstName
	}
	if true {
		toSerialize["contact-last-name"] = o.ContactLastName
	}
	if true {
		toSerialize["contact-phone"] = o.ContactPhone
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["profile-id"] = o.ProfileId
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UpdateVersion != nil {
		toSerialize["update-version"] = o.UpdateVersion
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateGlobalSignTarget struct {
	value *UpdateGlobalSignTarget
	isSet bool
}

func (v NullableUpdateGlobalSignTarget) Get() *UpdateGlobalSignTarget {
	return v.value
}

func (v *NullableUpdateGlobalSignTarget) Set(val *UpdateGlobalSignTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGlobalSignTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGlobalSignTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGlobalSignTarget(val *UpdateGlobalSignTarget) *NullableUpdateGlobalSignTarget {
	return &NullableUpdateGlobalSignTarget{value: val, isSet: true}
}

func (v NullableUpdateGlobalSignTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGlobalSignTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


