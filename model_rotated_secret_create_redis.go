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

// RotatedSecretCreateRedis struct for RotatedSecretCreateRedis
type RotatedSecretCreateRedis struct {
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
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Set the maximum number of versions, limited by the account settings defaults.
	MaxVersions *string `json:"max-versions,omitempty"`
	// Rotated secret name
	Name string `json:"name"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// rotated-username password (relevant only for rotator-type=password)
	RotatedPassword *string `json:"rotated-password,omitempty"`
	// username to be rotated, if selected use-self-creds at rotator-creds-type, this username will try to rotate it's own password, if use-target-creds is selected, target credentials will be use to rotate the rotated-password (relevant only for rotator-type=password)
	RotatedUsername *string `json:"rotated-username,omitempty"`
	// The Hour of the rotation in UTC
	RotationHour *int32 `json:"rotation-hour,omitempty"`
	// The number of days to wait between every automatic key rotation (1-365)
	RotationInterval *string `json:"rotation-interval,omitempty"`
	// The rotator type. options: [target/password]
	RotatorType string `json:"rotator-type"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName string `json:"target-name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewRotatedSecretCreateRedis instantiates a new RotatedSecretCreateRedis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotatedSecretCreateRedis(name string, rotatorType string, targetName string, ) *RotatedSecretCreateRedis {
	this := RotatedSecretCreateRedis{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var json bool = false
	this.Json = &json
	this.Name = name
	this.RotatorType = rotatorType
	this.TargetName = targetName
	return &this
}

// NewRotatedSecretCreateRedisWithDefaults instantiates a new RotatedSecretCreateRedis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotatedSecretCreateRedisWithDefaults() *RotatedSecretCreateRedis {
	this := RotatedSecretCreateRedis{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var json bool = false
	this.Json = &json
	return &this
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetAuthenticationCredentials() string {
	if o == nil || o.AuthenticationCredentials == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil || o.AuthenticationCredentials == nil {
		return nil, false
	}
	return o.AuthenticationCredentials, true
}

// HasAuthenticationCredentials returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasAuthenticationCredentials() bool {
	if o != nil && o.AuthenticationCredentials != nil {
		return true
	}

	return false
}

// SetAuthenticationCredentials gets a reference to the given string and assigns it to the AuthenticationCredentials field.
func (o *RotatedSecretCreateRedis) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = &v
}

// GetAutoRotate returns the AutoRotate field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetAutoRotate() string {
	if o == nil || o.AutoRotate == nil {
		var ret string
		return ret
	}
	return *o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetAutoRotateOk() (*string, bool) {
	if o == nil || o.AutoRotate == nil {
		return nil, false
	}
	return o.AutoRotate, true
}

// HasAutoRotate returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasAutoRotate() bool {
	if o != nil && o.AutoRotate != nil {
		return true
	}

	return false
}

// SetAutoRotate gets a reference to the given string and assigns it to the AutoRotate field.
func (o *RotatedSecretCreateRedis) SetAutoRotate(v string) {
	o.AutoRotate = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *RotatedSecretCreateRedis) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RotatedSecretCreateRedis) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *RotatedSecretCreateRedis) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RotatedSecretCreateRedis) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *RotatedSecretCreateRedis) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetName returns the Name field value
func (o *RotatedSecretCreateRedis) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RotatedSecretCreateRedis) SetName(v string) {
	o.Name = v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *RotatedSecretCreateRedis) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetRotatedPassword returns the RotatedPassword field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetRotatedPassword() string {
	if o == nil || o.RotatedPassword == nil {
		var ret string
		return ret
	}
	return *o.RotatedPassword
}

// GetRotatedPasswordOk returns a tuple with the RotatedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetRotatedPasswordOk() (*string, bool) {
	if o == nil || o.RotatedPassword == nil {
		return nil, false
	}
	return o.RotatedPassword, true
}

// HasRotatedPassword returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasRotatedPassword() bool {
	if o != nil && o.RotatedPassword != nil {
		return true
	}

	return false
}

// SetRotatedPassword gets a reference to the given string and assigns it to the RotatedPassword field.
func (o *RotatedSecretCreateRedis) SetRotatedPassword(v string) {
	o.RotatedPassword = &v
}

// GetRotatedUsername returns the RotatedUsername field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetRotatedUsername() string {
	if o == nil || o.RotatedUsername == nil {
		var ret string
		return ret
	}
	return *o.RotatedUsername
}

// GetRotatedUsernameOk returns a tuple with the RotatedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetRotatedUsernameOk() (*string, bool) {
	if o == nil || o.RotatedUsername == nil {
		return nil, false
	}
	return o.RotatedUsername, true
}

// HasRotatedUsername returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasRotatedUsername() bool {
	if o != nil && o.RotatedUsername != nil {
		return true
	}

	return false
}

// SetRotatedUsername gets a reference to the given string and assigns it to the RotatedUsername field.
func (o *RotatedSecretCreateRedis) SetRotatedUsername(v string) {
	o.RotatedUsername = &v
}

// GetRotationHour returns the RotationHour field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetRotationHour() int32 {
	if o == nil || o.RotationHour == nil {
		var ret int32
		return ret
	}
	return *o.RotationHour
}

// GetRotationHourOk returns a tuple with the RotationHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetRotationHourOk() (*int32, bool) {
	if o == nil || o.RotationHour == nil {
		return nil, false
	}
	return o.RotationHour, true
}

// HasRotationHour returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasRotationHour() bool {
	if o != nil && o.RotationHour != nil {
		return true
	}

	return false
}

// SetRotationHour gets a reference to the given int32 and assigns it to the RotationHour field.
func (o *RotatedSecretCreateRedis) SetRotationHour(v int32) {
	o.RotationHour = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetRotationInterval() string {
	if o == nil || o.RotationInterval == nil {
		var ret string
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetRotationIntervalOk() (*string, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given string and assigns it to the RotationInterval field.
func (o *RotatedSecretCreateRedis) SetRotationInterval(v string) {
	o.RotationInterval = &v
}

// GetRotatorType returns the RotatorType field value
func (o *RotatedSecretCreateRedis) GetRotatorType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RotatorType
}

// GetRotatorTypeOk returns a tuple with the RotatorType field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetRotatorTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RotatorType, true
}

// SetRotatorType sets field value
func (o *RotatedSecretCreateRedis) SetRotatorType(v string) {
	o.RotatorType = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RotatedSecretCreateRedis) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value
func (o *RotatedSecretCreateRedis) GetTargetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetTargetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetName, true
}

// SetTargetName sets field value
func (o *RotatedSecretCreateRedis) SetTargetName(v string) {
	o.TargetName = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RotatedSecretCreateRedis) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *RotatedSecretCreateRedis) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateRedis) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *RotatedSecretCreateRedis) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *RotatedSecretCreateRedis) SetUidToken(v string) {
	o.UidToken = &v
}

func (o RotatedSecretCreateRedis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.MaxVersions != nil {
		toSerialize["max-versions"] = o.MaxVersions
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.PasswordLength != nil {
		toSerialize["password-length"] = o.PasswordLength
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
	if true {
		toSerialize["rotator-type"] = o.RotatorType
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["target-name"] = o.TargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableRotatedSecretCreateRedis struct {
	value *RotatedSecretCreateRedis
	isSet bool
}

func (v NullableRotatedSecretCreateRedis) Get() *RotatedSecretCreateRedis {
	return v.value
}

func (v *NullableRotatedSecretCreateRedis) Set(val *RotatedSecretCreateRedis) {
	v.value = val
	v.isSet = true
}

func (v NullableRotatedSecretCreateRedis) IsSet() bool {
	return v.isSet
}

func (v *NullableRotatedSecretCreateRedis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotatedSecretCreateRedis(val *RotatedSecretCreateRedis) *NullableRotatedSecretCreateRedis {
	return &NullableRotatedSecretCreateRedis{value: val, isSet: true}
}

func (v NullableRotatedSecretCreateRedis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotatedSecretCreateRedis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


