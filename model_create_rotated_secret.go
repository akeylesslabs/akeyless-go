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

// CreateRotatedSecret struct for CreateRotatedSecret
type CreateRotatedSecret struct {
	ApiId *string `json:"api-id,omitempty"`
	ApiKey *string `json:"api-key,omitempty"`
	// Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation
	AutoRotate *string `json:"auto-rotate,omitempty"`
	CustomPayload *string `json:"custom-payload,omitempty"`
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Metadata about the secret
	Metadata *string `json:"metadata,omitempty"`
	// Secret name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	RotatedPassword *string `json:"rotated-password,omitempty"`
	RotatedUsername *string `json:"rotated-username,omitempty"`
	RotationHour *int32 `json:"rotation-hour,omitempty"`
	// The number of days to wait between every automatic key rotation (1-365)
	RotationInterval *string `json:"rotation-interval,omitempty"`
	RotatorCredsType *string `json:"rotator-creds-type,omitempty"`
	RotatorCustomCmd *string `json:"rotator-custom-cmd,omitempty"`
	// Rotator Type
	RotatorType string `json:"rotator-type"`
	// Deprecated: use RotatedPassword
	SshPassword *string `json:"ssh-password,omitempty"`
	// Deprecated: use RotatedUser
	SshUsername *string `json:"ssh-username,omitempty"`
	// List of the tags attached to this secret
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName string `json:"target-name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User Attribute
	UserAttribute *string `json:"user-attribute,omitempty"`
	// User DN
	UserDn *string `json:"user-dn,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewCreateRotatedSecret instantiates a new CreateRotatedSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRotatedSecret(name string, rotatorType string, targetName string, ) *CreateRotatedSecret {
	this := CreateRotatedSecret{}
	this.Name = name
	this.RotatorType = rotatorType
	this.TargetName = targetName
	return &this
}

// NewCreateRotatedSecretWithDefaults instantiates a new CreateRotatedSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRotatedSecretWithDefaults() *CreateRotatedSecret {
	this := CreateRotatedSecret{}
	return &this
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetApiId() string {
	if o == nil || o.ApiId == nil {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetApiIdOk() (*string, bool) {
	if o == nil || o.ApiId == nil {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasApiId() bool {
	if o != nil && o.ApiId != nil {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *CreateRotatedSecret) SetApiId(v string) {
	o.ApiId = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *CreateRotatedSecret) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAutoRotate returns the AutoRotate field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetAutoRotate() string {
	if o == nil || o.AutoRotate == nil {
		var ret string
		return ret
	}
	return *o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetAutoRotateOk() (*string, bool) {
	if o == nil || o.AutoRotate == nil {
		return nil, false
	}
	return o.AutoRotate, true
}

// HasAutoRotate returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasAutoRotate() bool {
	if o != nil && o.AutoRotate != nil {
		return true
	}

	return false
}

// SetAutoRotate gets a reference to the given string and assigns it to the AutoRotate field.
func (o *CreateRotatedSecret) SetAutoRotate(v string) {
	o.AutoRotate = &v
}

// GetCustomPayload returns the CustomPayload field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetCustomPayload() string {
	if o == nil || o.CustomPayload == nil {
		var ret string
		return ret
	}
	return *o.CustomPayload
}

// GetCustomPayloadOk returns a tuple with the CustomPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetCustomPayloadOk() (*string, bool) {
	if o == nil || o.CustomPayload == nil {
		return nil, false
	}
	return o.CustomPayload, true
}

// HasCustomPayload returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasCustomPayload() bool {
	if o != nil && o.CustomPayload != nil {
		return true
	}

	return false
}

// SetCustomPayload gets a reference to the given string and assigns it to the CustomPayload field.
func (o *CreateRotatedSecret) SetCustomPayload(v string) {
	o.CustomPayload = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateRotatedSecret) SetKey(v string) {
	o.Key = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *CreateRotatedSecret) SetMetadata(v string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *CreateRotatedSecret) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRotatedSecret) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateRotatedSecret) SetPassword(v string) {
	o.Password = &v
}

// GetRotatedPassword returns the RotatedPassword field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetRotatedPassword() string {
	if o == nil || o.RotatedPassword == nil {
		var ret string
		return ret
	}
	return *o.RotatedPassword
}

// GetRotatedPasswordOk returns a tuple with the RotatedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetRotatedPasswordOk() (*string, bool) {
	if o == nil || o.RotatedPassword == nil {
		return nil, false
	}
	return o.RotatedPassword, true
}

// HasRotatedPassword returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasRotatedPassword() bool {
	if o != nil && o.RotatedPassword != nil {
		return true
	}

	return false
}

// SetRotatedPassword gets a reference to the given string and assigns it to the RotatedPassword field.
func (o *CreateRotatedSecret) SetRotatedPassword(v string) {
	o.RotatedPassword = &v
}

// GetRotatedUsername returns the RotatedUsername field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetRotatedUsername() string {
	if o == nil || o.RotatedUsername == nil {
		var ret string
		return ret
	}
	return *o.RotatedUsername
}

// GetRotatedUsernameOk returns a tuple with the RotatedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetRotatedUsernameOk() (*string, bool) {
	if o == nil || o.RotatedUsername == nil {
		return nil, false
	}
	return o.RotatedUsername, true
}

// HasRotatedUsername returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasRotatedUsername() bool {
	if o != nil && o.RotatedUsername != nil {
		return true
	}

	return false
}

// SetRotatedUsername gets a reference to the given string and assigns it to the RotatedUsername field.
func (o *CreateRotatedSecret) SetRotatedUsername(v string) {
	o.RotatedUsername = &v
}

// GetRotationHour returns the RotationHour field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetRotationHour() int32 {
	if o == nil || o.RotationHour == nil {
		var ret int32
		return ret
	}
	return *o.RotationHour
}

// GetRotationHourOk returns a tuple with the RotationHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetRotationHourOk() (*int32, bool) {
	if o == nil || o.RotationHour == nil {
		return nil, false
	}
	return o.RotationHour, true
}

// HasRotationHour returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasRotationHour() bool {
	if o != nil && o.RotationHour != nil {
		return true
	}

	return false
}

// SetRotationHour gets a reference to the given int32 and assigns it to the RotationHour field.
func (o *CreateRotatedSecret) SetRotationHour(v int32) {
	o.RotationHour = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetRotationInterval() string {
	if o == nil || o.RotationInterval == nil {
		var ret string
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetRotationIntervalOk() (*string, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given string and assigns it to the RotationInterval field.
func (o *CreateRotatedSecret) SetRotationInterval(v string) {
	o.RotationInterval = &v
}

// GetRotatorCredsType returns the RotatorCredsType field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetRotatorCredsType() string {
	if o == nil || o.RotatorCredsType == nil {
		var ret string
		return ret
	}
	return *o.RotatorCredsType
}

// GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetRotatorCredsTypeOk() (*string, bool) {
	if o == nil || o.RotatorCredsType == nil {
		return nil, false
	}
	return o.RotatorCredsType, true
}

// HasRotatorCredsType returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasRotatorCredsType() bool {
	if o != nil && o.RotatorCredsType != nil {
		return true
	}

	return false
}

// SetRotatorCredsType gets a reference to the given string and assigns it to the RotatorCredsType field.
func (o *CreateRotatedSecret) SetRotatorCredsType(v string) {
	o.RotatorCredsType = &v
}

// GetRotatorCustomCmd returns the RotatorCustomCmd field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetRotatorCustomCmd() string {
	if o == nil || o.RotatorCustomCmd == nil {
		var ret string
		return ret
	}
	return *o.RotatorCustomCmd
}

// GetRotatorCustomCmdOk returns a tuple with the RotatorCustomCmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetRotatorCustomCmdOk() (*string, bool) {
	if o == nil || o.RotatorCustomCmd == nil {
		return nil, false
	}
	return o.RotatorCustomCmd, true
}

// HasRotatorCustomCmd returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasRotatorCustomCmd() bool {
	if o != nil && o.RotatorCustomCmd != nil {
		return true
	}

	return false
}

// SetRotatorCustomCmd gets a reference to the given string and assigns it to the RotatorCustomCmd field.
func (o *CreateRotatedSecret) SetRotatorCustomCmd(v string) {
	o.RotatorCustomCmd = &v
}

// GetRotatorType returns the RotatorType field value
func (o *CreateRotatedSecret) GetRotatorType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RotatorType
}

// GetRotatorTypeOk returns a tuple with the RotatorType field value
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetRotatorTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RotatorType, true
}

// SetRotatorType sets field value
func (o *CreateRotatedSecret) SetRotatorType(v string) {
	o.RotatorType = v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetSshPassword() string {
	if o == nil || o.SshPassword == nil {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetSshPasswordOk() (*string, bool) {
	if o == nil || o.SshPassword == nil {
		return nil, false
	}
	return o.SshPassword, true
}

// HasSshPassword returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasSshPassword() bool {
	if o != nil && o.SshPassword != nil {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *CreateRotatedSecret) SetSshPassword(v string) {
	o.SshPassword = &v
}

// GetSshUsername returns the SshUsername field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetSshUsername() string {
	if o == nil || o.SshUsername == nil {
		var ret string
		return ret
	}
	return *o.SshUsername
}

// GetSshUsernameOk returns a tuple with the SshUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetSshUsernameOk() (*string, bool) {
	if o == nil || o.SshUsername == nil {
		return nil, false
	}
	return o.SshUsername, true
}

// HasSshUsername returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasSshUsername() bool {
	if o != nil && o.SshUsername != nil {
		return true
	}

	return false
}

// SetSshUsername gets a reference to the given string and assigns it to the SshUsername field.
func (o *CreateRotatedSecret) SetSshUsername(v string) {
	o.SshUsername = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateRotatedSecret) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value
func (o *CreateRotatedSecret) GetTargetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetTargetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetName, true
}

// SetTargetName sets field value
func (o *CreateRotatedSecret) SetTargetName(v string) {
	o.TargetName = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateRotatedSecret) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateRotatedSecret) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserAttribute returns the UserAttribute field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetUserAttribute() string {
	if o == nil || o.UserAttribute == nil {
		var ret string
		return ret
	}
	return *o.UserAttribute
}

// GetUserAttributeOk returns a tuple with the UserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetUserAttributeOk() (*string, bool) {
	if o == nil || o.UserAttribute == nil {
		return nil, false
	}
	return o.UserAttribute, true
}

// HasUserAttribute returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasUserAttribute() bool {
	if o != nil && o.UserAttribute != nil {
		return true
	}

	return false
}

// SetUserAttribute gets a reference to the given string and assigns it to the UserAttribute field.
func (o *CreateRotatedSecret) SetUserAttribute(v string) {
	o.UserAttribute = &v
}

// GetUserDn returns the UserDn field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetUserDn() string {
	if o == nil || o.UserDn == nil {
		var ret string
		return ret
	}
	return *o.UserDn
}

// GetUserDnOk returns a tuple with the UserDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetUserDnOk() (*string, bool) {
	if o == nil || o.UserDn == nil {
		return nil, false
	}
	return o.UserDn, true
}

// HasUserDn returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasUserDn() bool {
	if o != nil && o.UserDn != nil {
		return true
	}

	return false
}

// SetUserDn gets a reference to the given string and assigns it to the UserDn field.
func (o *CreateRotatedSecret) SetUserDn(v string) {
	o.UserDn = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateRotatedSecret) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRotatedSecret) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateRotatedSecret) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateRotatedSecret) SetUsername(v string) {
	o.Username = &v
}

func (o CreateRotatedSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiId != nil {
		toSerialize["api-id"] = o.ApiId
	}
	if o.ApiKey != nil {
		toSerialize["api-key"] = o.ApiKey
	}
	if o.AutoRotate != nil {
		toSerialize["auto-rotate"] = o.AutoRotate
	}
	if o.CustomPayload != nil {
		toSerialize["custom-payload"] = o.CustomPayload
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
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
	if o.RotatorCredsType != nil {
		toSerialize["rotator-creds-type"] = o.RotatorCredsType
	}
	if o.RotatorCustomCmd != nil {
		toSerialize["rotator-custom-cmd"] = o.RotatorCustomCmd
	}
	if true {
		toSerialize["rotator-type"] = o.RotatorType
	}
	if o.SshPassword != nil {
		toSerialize["ssh-password"] = o.SshPassword
	}
	if o.SshUsername != nil {
		toSerialize["ssh-username"] = o.SshUsername
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
	if o.UserAttribute != nil {
		toSerialize["user-attribute"] = o.UserAttribute
	}
	if o.UserDn != nil {
		toSerialize["user-dn"] = o.UserDn
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRotatedSecret struct {
	value *CreateRotatedSecret
	isSet bool
}

func (v NullableCreateRotatedSecret) Get() *CreateRotatedSecret {
	return v.value
}

func (v *NullableCreateRotatedSecret) Set(val *CreateRotatedSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRotatedSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRotatedSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRotatedSecret(val *CreateRotatedSecret) *NullableCreateRotatedSecret {
	return &NullableCreateRotatedSecret{value: val, isSet: true}
}

func (v NullableCreateRotatedSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRotatedSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


