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

// RotatedSecretUpdateRedis struct for RotatedSecretUpdateRedis
type RotatedSecretUpdateRedis struct {
	// List of the new tags that will be attached to this item
	AddTag *[]string `json:"add-tag,omitempty"`
	// The credentials to connect with use-user-creds/use-target-creds
	AuthenticationCredentials *string `json:"authentication-credentials,omitempty"`
	// Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation [true/false]
	AutoRotate *string `json:"auto-rotate,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Whether to keep previous version [true/false]. If not set, use default according to account settings
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Set the maximum number of versions, limited by the account settings defaults.
	MaxVersions *string `json:"max-versions,omitempty"`
	// Rotated secret name
	Name string `json:"name"`
	// New item name
	NewName *string `json:"new-name,omitempty"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// List of the existent tags that will be removed from this item
	RmTag *[]string `json:"rm-tag,omitempty"`
	// rotated-username password (relevant only for rotator-type=password)
	RotatedPassword *string `json:"rotated-password,omitempty"`
	// username to be rotated, if selected use-self-creds at rotator-creds-type, this username will try to rotate it's own password, if use-target-creds is selected, target credentials will be use to rotate the rotated-password (relevant only for rotator-type=password)
	RotatedUsername *string `json:"rotated-username,omitempty"`
	// The Hour of the rotation in UTC
	RotationHour *int32 `json:"rotation-hour,omitempty"`
	// The number of days to wait between every automatic key rotation (1-365)
	RotationInterval *string `json:"rotation-interval,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewRotatedSecretUpdateRedis instantiates a new RotatedSecretUpdateRedis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotatedSecretUpdateRedis(name string, ) *RotatedSecretUpdateRedis {
	this := RotatedSecretUpdateRedis{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var description string = "default_metadata"
	this.Description = &description
	var json bool = false
	this.Json = &json
	this.Name = name
	return &this
}

// NewRotatedSecretUpdateRedisWithDefaults instantiates a new RotatedSecretUpdateRedis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotatedSecretUpdateRedisWithDefaults() *RotatedSecretUpdateRedis {
	this := RotatedSecretUpdateRedis{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var description string = "default_metadata"
	this.Description = &description
	var json bool = false
	this.Json = &json
	return &this
}

// GetAddTag returns the AddTag field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetAddTag() []string {
	if o == nil || o.AddTag == nil {
		var ret []string
		return ret
	}
	return *o.AddTag
}

// GetAddTagOk returns a tuple with the AddTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetAddTagOk() (*[]string, bool) {
	if o == nil || o.AddTag == nil {
		return nil, false
	}
	return o.AddTag, true
}

// HasAddTag returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasAddTag() bool {
	if o != nil && o.AddTag != nil {
		return true
	}

	return false
}

// SetAddTag gets a reference to the given []string and assigns it to the AddTag field.
func (o *RotatedSecretUpdateRedis) SetAddTag(v []string) {
	o.AddTag = &v
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetAuthenticationCredentials() string {
	if o == nil || o.AuthenticationCredentials == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil || o.AuthenticationCredentials == nil {
		return nil, false
	}
	return o.AuthenticationCredentials, true
}

// HasAuthenticationCredentials returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasAuthenticationCredentials() bool {
	if o != nil && o.AuthenticationCredentials != nil {
		return true
	}

	return false
}

// SetAuthenticationCredentials gets a reference to the given string and assigns it to the AuthenticationCredentials field.
func (o *RotatedSecretUpdateRedis) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = &v
}

// GetAutoRotate returns the AutoRotate field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetAutoRotate() string {
	if o == nil || o.AutoRotate == nil {
		var ret string
		return ret
	}
	return *o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetAutoRotateOk() (*string, bool) {
	if o == nil || o.AutoRotate == nil {
		return nil, false
	}
	return o.AutoRotate, true
}

// HasAutoRotate returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasAutoRotate() bool {
	if o != nil && o.AutoRotate != nil {
		return true
	}

	return false
}

// SetAutoRotate gets a reference to the given string and assigns it to the AutoRotate field.
func (o *RotatedSecretUpdateRedis) SetAutoRotate(v string) {
	o.AutoRotate = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *RotatedSecretUpdateRedis) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RotatedSecretUpdateRedis) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *RotatedSecretUpdateRedis) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *RotatedSecretUpdateRedis) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RotatedSecretUpdateRedis) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *RotatedSecretUpdateRedis) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetName returns the Name field value
func (o *RotatedSecretUpdateRedis) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RotatedSecretUpdateRedis) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *RotatedSecretUpdateRedis) SetNewName(v string) {
	o.NewName = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *RotatedSecretUpdateRedis) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetRmTag returns the RmTag field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetRmTag() []string {
	if o == nil || o.RmTag == nil {
		var ret []string
		return ret
	}
	return *o.RmTag
}

// GetRmTagOk returns a tuple with the RmTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetRmTagOk() (*[]string, bool) {
	if o == nil || o.RmTag == nil {
		return nil, false
	}
	return o.RmTag, true
}

// HasRmTag returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasRmTag() bool {
	if o != nil && o.RmTag != nil {
		return true
	}

	return false
}

// SetRmTag gets a reference to the given []string and assigns it to the RmTag field.
func (o *RotatedSecretUpdateRedis) SetRmTag(v []string) {
	o.RmTag = &v
}

// GetRotatedPassword returns the RotatedPassword field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetRotatedPassword() string {
	if o == nil || o.RotatedPassword == nil {
		var ret string
		return ret
	}
	return *o.RotatedPassword
}

// GetRotatedPasswordOk returns a tuple with the RotatedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetRotatedPasswordOk() (*string, bool) {
	if o == nil || o.RotatedPassword == nil {
		return nil, false
	}
	return o.RotatedPassword, true
}

// HasRotatedPassword returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasRotatedPassword() bool {
	if o != nil && o.RotatedPassword != nil {
		return true
	}

	return false
}

// SetRotatedPassword gets a reference to the given string and assigns it to the RotatedPassword field.
func (o *RotatedSecretUpdateRedis) SetRotatedPassword(v string) {
	o.RotatedPassword = &v
}

// GetRotatedUsername returns the RotatedUsername field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetRotatedUsername() string {
	if o == nil || o.RotatedUsername == nil {
		var ret string
		return ret
	}
	return *o.RotatedUsername
}

// GetRotatedUsernameOk returns a tuple with the RotatedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetRotatedUsernameOk() (*string, bool) {
	if o == nil || o.RotatedUsername == nil {
		return nil, false
	}
	return o.RotatedUsername, true
}

// HasRotatedUsername returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasRotatedUsername() bool {
	if o != nil && o.RotatedUsername != nil {
		return true
	}

	return false
}

// SetRotatedUsername gets a reference to the given string and assigns it to the RotatedUsername field.
func (o *RotatedSecretUpdateRedis) SetRotatedUsername(v string) {
	o.RotatedUsername = &v
}

// GetRotationHour returns the RotationHour field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetRotationHour() int32 {
	if o == nil || o.RotationHour == nil {
		var ret int32
		return ret
	}
	return *o.RotationHour
}

// GetRotationHourOk returns a tuple with the RotationHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetRotationHourOk() (*int32, bool) {
	if o == nil || o.RotationHour == nil {
		return nil, false
	}
	return o.RotationHour, true
}

// HasRotationHour returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasRotationHour() bool {
	if o != nil && o.RotationHour != nil {
		return true
	}

	return false
}

// SetRotationHour gets a reference to the given int32 and assigns it to the RotationHour field.
func (o *RotatedSecretUpdateRedis) SetRotationHour(v int32) {
	o.RotationHour = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetRotationInterval() string {
	if o == nil || o.RotationInterval == nil {
		var ret string
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetRotationIntervalOk() (*string, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given string and assigns it to the RotationInterval field.
func (o *RotatedSecretUpdateRedis) SetRotationInterval(v string) {
	o.RotationInterval = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RotatedSecretUpdateRedis) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *RotatedSecretUpdateRedis) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateRedis) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *RotatedSecretUpdateRedis) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *RotatedSecretUpdateRedis) SetUidToken(v string) {
	o.UidToken = &v
}

func (o RotatedSecretUpdateRedis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddTag != nil {
		toSerialize["add-tag"] = o.AddTag
	}
	if o.AuthenticationCredentials != nil {
		toSerialize["authentication-credentials"] = o.AuthenticationCredentials
	}
	if o.AutoRotate != nil {
		toSerialize["auto-rotate"] = o.AutoRotate
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
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
	if o.MaxVersions != nil {
		toSerialize["max-versions"] = o.MaxVersions
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.PasswordLength != nil {
		toSerialize["password-length"] = o.PasswordLength
	}
	if o.RmTag != nil {
		toSerialize["rm-tag"] = o.RmTag
	}
	if o.RotatedPassword != nil {
		toSerialize["rotated-password"] = o.RotatedPassword
	}
	if o.RotatedUsername != nil {
		toSerialize["rotated-username"] = o.RotatedUsername
	}
	if o.RotationHour != nil {
		toSerialize["rotation-hour"] = o.RotationHour
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
	return json.Marshal(toSerialize)
}

type NullableRotatedSecretUpdateRedis struct {
	value *RotatedSecretUpdateRedis
	isSet bool
}

func (v NullableRotatedSecretUpdateRedis) Get() *RotatedSecretUpdateRedis {
	return v.value
}

func (v *NullableRotatedSecretUpdateRedis) Set(val *RotatedSecretUpdateRedis) {
	v.value = val
	v.isSet = true
}

func (v NullableRotatedSecretUpdateRedis) IsSet() bool {
	return v.isSet
}

func (v *NullableRotatedSecretUpdateRedis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotatedSecretUpdateRedis(val *RotatedSecretUpdateRedis) *NullableRotatedSecretUpdateRedis {
	return &NullableRotatedSecretUpdateRedis{value: val, isSet: true}
}

func (v NullableRotatedSecretUpdateRedis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotatedSecretUpdateRedis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


