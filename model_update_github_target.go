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

// UpdateGithubTarget struct for UpdateGithubTarget
type UpdateGithubTarget struct {
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	// Github app id
	GithubAppId *int64 `json:"github-app-id,omitempty"`
	// App private key
	GithubAppPrivateKey *string `json:"github-app-private-key,omitempty"`
	// Base URL
	GithubBaseUrl *string `json:"github-base-url,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Target name
	Name string `json:"name"`
	// New target name
	NewName *string `json:"new-name,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Create new version for the target
	UpdateVersion *bool `json:"update-version,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewUpdateGithubTarget instantiates a new UpdateGithubTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGithubTarget(name string, ) *UpdateGithubTarget {
	this := UpdateGithubTarget{}
	this.Name = name
	var updateVersion bool = false
	this.UpdateVersion = &updateVersion
	return &this
}

// NewUpdateGithubTargetWithDefaults instantiates a new UpdateGithubTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGithubTargetWithDefaults() *UpdateGithubTarget {
	this := UpdateGithubTarget{}
	var updateVersion bool = false
	this.UpdateVersion = &updateVersion
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateGithubTarget) SetComment(v string) {
	o.Comment = &v
}

// GetGithubAppId returns the GithubAppId field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetGithubAppId() int64 {
	if o == nil || o.GithubAppId == nil {
		var ret int64
		return ret
	}
	return *o.GithubAppId
}

// GetGithubAppIdOk returns a tuple with the GithubAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetGithubAppIdOk() (*int64, bool) {
	if o == nil || o.GithubAppId == nil {
		return nil, false
	}
	return o.GithubAppId, true
}

// HasGithubAppId returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasGithubAppId() bool {
	if o != nil && o.GithubAppId != nil {
		return true
	}

	return false
}

// SetGithubAppId gets a reference to the given int64 and assigns it to the GithubAppId field.
func (o *UpdateGithubTarget) SetGithubAppId(v int64) {
	o.GithubAppId = &v
}

// GetGithubAppPrivateKey returns the GithubAppPrivateKey field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetGithubAppPrivateKey() string {
	if o == nil || o.GithubAppPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.GithubAppPrivateKey
}

// GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetGithubAppPrivateKeyOk() (*string, bool) {
	if o == nil || o.GithubAppPrivateKey == nil {
		return nil, false
	}
	return o.GithubAppPrivateKey, true
}

// HasGithubAppPrivateKey returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasGithubAppPrivateKey() bool {
	if o != nil && o.GithubAppPrivateKey != nil {
		return true
	}

	return false
}

// SetGithubAppPrivateKey gets a reference to the given string and assigns it to the GithubAppPrivateKey field.
func (o *UpdateGithubTarget) SetGithubAppPrivateKey(v string) {
	o.GithubAppPrivateKey = &v
}

// GetGithubBaseUrl returns the GithubBaseUrl field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetGithubBaseUrl() string {
	if o == nil || o.GithubBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.GithubBaseUrl
}

// GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetGithubBaseUrlOk() (*string, bool) {
	if o == nil || o.GithubBaseUrl == nil {
		return nil, false
	}
	return o.GithubBaseUrl, true
}

// HasGithubBaseUrl returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasGithubBaseUrl() bool {
	if o != nil && o.GithubBaseUrl != nil {
		return true
	}

	return false
}

// SetGithubBaseUrl gets a reference to the given string and assigns it to the GithubBaseUrl field.
func (o *UpdateGithubTarget) SetGithubBaseUrl(v string) {
	o.GithubBaseUrl = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateGithubTarget) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *UpdateGithubTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateGithubTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateGithubTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateGithubTarget) SetPassword(v string) {
	o.Password = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateGithubTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateGithubTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateGithubTarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateGithubTarget) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubTarget) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateGithubTarget) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateGithubTarget) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateGithubTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.GithubAppId != nil {
		toSerialize["github-app-id"] = o.GithubAppId
	}
	if o.GithubAppPrivateKey != nil {
		toSerialize["github-app-private-key"] = o.GithubAppPrivateKey
	}
	if o.GithubBaseUrl != nil {
		toSerialize["github-base-url"] = o.GithubBaseUrl
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
	if o.Password != nil {
		toSerialize["password"] = o.Password
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateGithubTarget struct {
	value *UpdateGithubTarget
	isSet bool
}

func (v NullableUpdateGithubTarget) Get() *UpdateGithubTarget {
	return v.value
}

func (v *NullableUpdateGithubTarget) Set(val *UpdateGithubTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGithubTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGithubTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGithubTarget(val *UpdateGithubTarget) *NullableUpdateGithubTarget {
	return &NullableUpdateGithubTarget{value: val, isSet: true}
}

func (v NullableUpdateGithubTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGithubTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


