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

// checks if the UpdateDockerhubTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDockerhubTarget{}

// UpdateDockerhubTarget struct for UpdateDockerhubTarget
type UpdateDockerhubTarget struct {
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	DockerhubPassword *string `json:"dockerhub-password,omitempty"`
	DockerhubUsername *string `json:"dockerhub-username,omitempty"`
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
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Deprecated
	UpdateVersion *bool `json:"update-version,omitempty"`
}

// NewUpdateDockerhubTarget instantiates a new UpdateDockerhubTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDockerhubTarget(name string) *UpdateDockerhubTarget {
	this := UpdateDockerhubTarget{}
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewUpdateDockerhubTargetWithDefaults instantiates a new UpdateDockerhubTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDockerhubTargetWithDefaults() *UpdateDockerhubTarget {
	this := UpdateDockerhubTarget{}
	var json bool = false
	this.Json = &json
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateDockerhubTarget) SetComment(v string) {
	o.Comment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateDockerhubTarget) SetDescription(v string) {
	o.Description = &v
}

// GetDockerhubPassword returns the DockerhubPassword field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetDockerhubPassword() string {
	if o == nil || IsNil(o.DockerhubPassword) {
		var ret string
		return ret
	}
	return *o.DockerhubPassword
}

// GetDockerhubPasswordOk returns a tuple with the DockerhubPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetDockerhubPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DockerhubPassword) {
		return nil, false
	}
	return o.DockerhubPassword, true
}

// HasDockerhubPassword returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasDockerhubPassword() bool {
	if o != nil && !IsNil(o.DockerhubPassword) {
		return true
	}

	return false
}

// SetDockerhubPassword gets a reference to the given string and assigns it to the DockerhubPassword field.
func (o *UpdateDockerhubTarget) SetDockerhubPassword(v string) {
	o.DockerhubPassword = &v
}

// GetDockerhubUsername returns the DockerhubUsername field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetDockerhubUsername() string {
	if o == nil || IsNil(o.DockerhubUsername) {
		var ret string
		return ret
	}
	return *o.DockerhubUsername
}

// GetDockerhubUsernameOk returns a tuple with the DockerhubUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetDockerhubUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.DockerhubUsername) {
		return nil, false
	}
	return o.DockerhubUsername, true
}

// HasDockerhubUsername returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasDockerhubUsername() bool {
	if o != nil && !IsNil(o.DockerhubUsername) {
		return true
	}

	return false
}

// SetDockerhubUsername gets a reference to the given string and assigns it to the DockerhubUsername field.
func (o *UpdateDockerhubTarget) SetDockerhubUsername(v string) {
	o.DockerhubUsername = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *UpdateDockerhubTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetKeepPrevVersion() string {
	if o == nil || IsNil(o.KeepPrevVersion) {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KeepPrevVersion) {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasKeepPrevVersion() bool {
	if o != nil && !IsNil(o.KeepPrevVersion) {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateDockerhubTarget) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateDockerhubTarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *UpdateDockerhubTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateDockerhubTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetNewName() string {
	if o == nil || IsNil(o.NewName) {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetNewNameOk() (*string, bool) {
	if o == nil || IsNil(o.NewName) {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasNewName() bool {
	if o != nil && !IsNil(o.NewName) {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateDockerhubTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateDockerhubTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateDockerhubTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateDockerhubTarget) GetUpdateVersion() bool {
	if o == nil || IsNil(o.UpdateVersion) {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDockerhubTarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateVersion) {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateDockerhubTarget) HasUpdateVersion() bool {
	if o != nil && !IsNil(o.UpdateVersion) {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateDockerhubTarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

func (o UpdateDockerhubTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDockerhubTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DockerhubPassword) {
		toSerialize["dockerhub-password"] = o.DockerhubPassword
	}
	if !IsNil(o.DockerhubUsername) {
		toSerialize["dockerhub-username"] = o.DockerhubUsername
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.KeepPrevVersion) {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NewName) {
		toSerialize["new-name"] = o.NewName
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	if !IsNil(o.UpdateVersion) {
		toSerialize["update-version"] = o.UpdateVersion
	}
	return toSerialize, nil
}

type NullableUpdateDockerhubTarget struct {
	value *UpdateDockerhubTarget
	isSet bool
}

func (v NullableUpdateDockerhubTarget) Get() *UpdateDockerhubTarget {
	return v.value
}

func (v *NullableUpdateDockerhubTarget) Set(val *UpdateDockerhubTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDockerhubTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDockerhubTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDockerhubTarget(val *UpdateDockerhubTarget) *NullableUpdateDockerhubTarget {
	return &NullableUpdateDockerhubTarget{value: val, isSet: true}
}

func (v NullableUpdateDockerhubTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDockerhubTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


