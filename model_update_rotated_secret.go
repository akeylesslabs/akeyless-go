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

// UpdateRotatedSecret updateRotatedSecret is a command that updates rotated secret. [Deprecated: Use gateway-update-item command]
type UpdateRotatedSecret struct {
	// List of the new tags that will be attached to this item
	AddTag *[]string `json:"add-tag,omitempty"`
	ApiId *string `json:"api-id,omitempty"`
	ApiKey *string `json:"api-key,omitempty"`
	// Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation
	AutoRotate *string `json:"auto-rotate,omitempty"`
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Secret name
	Name string `json:"name"`
	// New item metadata
	NewMetadata *string `json:"new-metadata,omitempty"`
	// New item name
	NewName *string `json:"new-name,omitempty"`
	// Whether to create a new version of not
	NewVersion *bool `json:"new-version,omitempty"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// List of the existent tags that will be removed from this item
	RmTag *[]string `json:"rm-tag,omitempty"`
	RotatedPassword *string `json:"rotated-password,omitempty"`
	RotatedUsername *string `json:"rotated-username,omitempty"`
	RotationHour *int32 `json:"rotation-hour,omitempty"`
	// The number of days to wait between every automatic key rotation (7-365)
	RotationInterval *string `json:"rotation-interval,omitempty"`
	RotatorCredsType *string `json:"rotator-creds-type,omitempty"`
	RotatorCustomCmd *string `json:"rotator-custom-cmd,omitempty"`
	// Deprecated: use RotatedPassword
	SshPassword *string `json:"ssh-password,omitempty"`
	// Deprecated: use RotatedUser
	SshUsername *string `json:"ssh-username,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewUpdateRotatedSecret instantiates a new UpdateRotatedSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRotatedSecret(name string, ) *UpdateRotatedSecret {
	this := UpdateRotatedSecret{}
	this.Name = name
	var newMetadata string = "default_metadata"
	this.NewMetadata = &newMetadata
	var newVersion bool = false
	this.NewVersion = &newVersion
	return &this
}

// NewUpdateRotatedSecretWithDefaults instantiates a new UpdateRotatedSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRotatedSecretWithDefaults() *UpdateRotatedSecret {
	this := UpdateRotatedSecret{}
	var newMetadata string = "default_metadata"
	this.NewMetadata = &newMetadata
	var newVersion bool = false
	this.NewVersion = &newVersion
	return &this
}

// GetAddTag returns the AddTag field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetAddTag() []string {
	if o == nil || o.AddTag == nil {
		var ret []string
		return ret
	}
	return *o.AddTag
}

// GetAddTagOk returns a tuple with the AddTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetAddTagOk() (*[]string, bool) {
	if o == nil || o.AddTag == nil {
		return nil, false
	}
	return o.AddTag, true
}

// HasAddTag returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasAddTag() bool {
	if o != nil && o.AddTag != nil {
		return true
	}

	return false
}

// SetAddTag gets a reference to the given []string and assigns it to the AddTag field.
func (o *UpdateRotatedSecret) SetAddTag(v []string) {
	o.AddTag = &v
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetApiId() string {
	if o == nil || o.ApiId == nil {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetApiIdOk() (*string, bool) {
	if o == nil || o.ApiId == nil {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasApiId() bool {
	if o != nil && o.ApiId != nil {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *UpdateRotatedSecret) SetApiId(v string) {
	o.ApiId = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *UpdateRotatedSecret) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAutoRotate returns the AutoRotate field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetAutoRotate() string {
	if o == nil || o.AutoRotate == nil {
		var ret string
		return ret
	}
	return *o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetAutoRotateOk() (*string, bool) {
	if o == nil || o.AutoRotate == nil {
		return nil, false
	}
	return o.AutoRotate, true
}

// HasAutoRotate returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasAutoRotate() bool {
	if o != nil && o.AutoRotate != nil {
		return true
	}

	return false
}

// SetAutoRotate gets a reference to the given string and assigns it to the AutoRotate field.
func (o *UpdateRotatedSecret) SetAutoRotate(v string) {
	o.AutoRotate = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateRotatedSecret) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *UpdateRotatedSecret) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateRotatedSecret) SetName(v string) {
	o.Name = v
}

// GetNewMetadata returns the NewMetadata field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetNewMetadata() string {
	if o == nil || o.NewMetadata == nil {
		var ret string
		return ret
	}
	return *o.NewMetadata
}

// GetNewMetadataOk returns a tuple with the NewMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetNewMetadataOk() (*string, bool) {
	if o == nil || o.NewMetadata == nil {
		return nil, false
	}
	return o.NewMetadata, true
}

// HasNewMetadata returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasNewMetadata() bool {
	if o != nil && o.NewMetadata != nil {
		return true
	}

	return false
}

// SetNewMetadata gets a reference to the given string and assigns it to the NewMetadata field.
func (o *UpdateRotatedSecret) SetNewMetadata(v string) {
	o.NewMetadata = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateRotatedSecret) SetNewName(v string) {
	o.NewName = &v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetNewVersion() bool {
	if o == nil || o.NewVersion == nil {
		var ret bool
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetNewVersionOk() (*bool, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given bool and assigns it to the NewVersion field.
func (o *UpdateRotatedSecret) SetNewVersion(v bool) {
	o.NewVersion = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateRotatedSecret) SetPassword(v string) {
	o.Password = &v
}

// GetRmTag returns the RmTag field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetRmTag() []string {
	if o == nil || o.RmTag == nil {
		var ret []string
		return ret
	}
	return *o.RmTag
}

// GetRmTagOk returns a tuple with the RmTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetRmTagOk() (*[]string, bool) {
	if o == nil || o.RmTag == nil {
		return nil, false
	}
	return o.RmTag, true
}

// HasRmTag returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasRmTag() bool {
	if o != nil && o.RmTag != nil {
		return true
	}

	return false
}

// SetRmTag gets a reference to the given []string and assigns it to the RmTag field.
func (o *UpdateRotatedSecret) SetRmTag(v []string) {
	o.RmTag = &v
}

// GetRotatedPassword returns the RotatedPassword field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetRotatedPassword() string {
	if o == nil || o.RotatedPassword == nil {
		var ret string
		return ret
	}
	return *o.RotatedPassword
}

// GetRotatedPasswordOk returns a tuple with the RotatedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetRotatedPasswordOk() (*string, bool) {
	if o == nil || o.RotatedPassword == nil {
		return nil, false
	}
	return o.RotatedPassword, true
}

// HasRotatedPassword returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasRotatedPassword() bool {
	if o != nil && o.RotatedPassword != nil {
		return true
	}

	return false
}

// SetRotatedPassword gets a reference to the given string and assigns it to the RotatedPassword field.
func (o *UpdateRotatedSecret) SetRotatedPassword(v string) {
	o.RotatedPassword = &v
}

// GetRotatedUsername returns the RotatedUsername field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetRotatedUsername() string {
	if o == nil || o.RotatedUsername == nil {
		var ret string
		return ret
	}
	return *o.RotatedUsername
}

// GetRotatedUsernameOk returns a tuple with the RotatedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetRotatedUsernameOk() (*string, bool) {
	if o == nil || o.RotatedUsername == nil {
		return nil, false
	}
	return o.RotatedUsername, true
}

// HasRotatedUsername returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasRotatedUsername() bool {
	if o != nil && o.RotatedUsername != nil {
		return true
	}

	return false
}

// SetRotatedUsername gets a reference to the given string and assigns it to the RotatedUsername field.
func (o *UpdateRotatedSecret) SetRotatedUsername(v string) {
	o.RotatedUsername = &v
}

// GetRotationHour returns the RotationHour field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetRotationHour() int32 {
	if o == nil || o.RotationHour == nil {
		var ret int32
		return ret
	}
	return *o.RotationHour
}

// GetRotationHourOk returns a tuple with the RotationHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetRotationHourOk() (*int32, bool) {
	if o == nil || o.RotationHour == nil {
		return nil, false
	}
	return o.RotationHour, true
}

// HasRotationHour returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasRotationHour() bool {
	if o != nil && o.RotationHour != nil {
		return true
	}

	return false
}

// SetRotationHour gets a reference to the given int32 and assigns it to the RotationHour field.
func (o *UpdateRotatedSecret) SetRotationHour(v int32) {
	o.RotationHour = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetRotationInterval() string {
	if o == nil || o.RotationInterval == nil {
		var ret string
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetRotationIntervalOk() (*string, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given string and assigns it to the RotationInterval field.
func (o *UpdateRotatedSecret) SetRotationInterval(v string) {
	o.RotationInterval = &v
}

// GetRotatorCredsType returns the RotatorCredsType field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetRotatorCredsType() string {
	if o == nil || o.RotatorCredsType == nil {
		var ret string
		return ret
	}
	return *o.RotatorCredsType
}

// GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetRotatorCredsTypeOk() (*string, bool) {
	if o == nil || o.RotatorCredsType == nil {
		return nil, false
	}
	return o.RotatorCredsType, true
}

// HasRotatorCredsType returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasRotatorCredsType() bool {
	if o != nil && o.RotatorCredsType != nil {
		return true
	}

	return false
}

// SetRotatorCredsType gets a reference to the given string and assigns it to the RotatorCredsType field.
func (o *UpdateRotatedSecret) SetRotatorCredsType(v string) {
	o.RotatorCredsType = &v
}

// GetRotatorCustomCmd returns the RotatorCustomCmd field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetRotatorCustomCmd() string {
	if o == nil || o.RotatorCustomCmd == nil {
		var ret string
		return ret
	}
	return *o.RotatorCustomCmd
}

// GetRotatorCustomCmdOk returns a tuple with the RotatorCustomCmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetRotatorCustomCmdOk() (*string, bool) {
	if o == nil || o.RotatorCustomCmd == nil {
		return nil, false
	}
	return o.RotatorCustomCmd, true
}

// HasRotatorCustomCmd returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasRotatorCustomCmd() bool {
	if o != nil && o.RotatorCustomCmd != nil {
		return true
	}

	return false
}

// SetRotatorCustomCmd gets a reference to the given string and assigns it to the RotatorCustomCmd field.
func (o *UpdateRotatedSecret) SetRotatorCustomCmd(v string) {
	o.RotatorCustomCmd = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetSshPassword() string {
	if o == nil || o.SshPassword == nil {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetSshPasswordOk() (*string, bool) {
	if o == nil || o.SshPassword == nil {
		return nil, false
	}
	return o.SshPassword, true
}

// HasSshPassword returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasSshPassword() bool {
	if o != nil && o.SshPassword != nil {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *UpdateRotatedSecret) SetSshPassword(v string) {
	o.SshPassword = &v
}

// GetSshUsername returns the SshUsername field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetSshUsername() string {
	if o == nil || o.SshUsername == nil {
		var ret string
		return ret
	}
	return *o.SshUsername
}

// GetSshUsernameOk returns a tuple with the SshUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetSshUsernameOk() (*string, bool) {
	if o == nil || o.SshUsername == nil {
		return nil, false
	}
	return o.SshUsername, true
}

// HasSshUsername returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasSshUsername() bool {
	if o != nil && o.SshUsername != nil {
		return true
	}

	return false
}

// SetSshUsername gets a reference to the given string and assigns it to the SshUsername field.
func (o *UpdateRotatedSecret) SetSshUsername(v string) {
	o.SshUsername = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateRotatedSecret) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateRotatedSecret) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateRotatedSecret) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRotatedSecret) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateRotatedSecret) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateRotatedSecret) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateRotatedSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddTag != nil {
		toSerialize["add-tag"] = o.AddTag
	}
	if o.ApiId != nil {
		toSerialize["api-id"] = o.ApiId
	}
	if o.ApiKey != nil {
		toSerialize["api-key"] = o.ApiKey
	}
	if o.AutoRotate != nil {
		toSerialize["auto-rotate"] = o.AutoRotate
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewMetadata != nil {
		toSerialize["new-metadata"] = o.NewMetadata
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.NewVersion != nil {
		toSerialize["new-version"] = o.NewVersion
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
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
	if o.RotatorCredsType != nil {
		toSerialize["rotator-creds-type"] = o.RotatorCredsType
	}
	if o.RotatorCustomCmd != nil {
		toSerialize["rotator-custom-cmd"] = o.RotatorCustomCmd
	}
	if o.SshPassword != nil {
		toSerialize["ssh-password"] = o.SshPassword
	}
	if o.SshUsername != nil {
		toSerialize["ssh-username"] = o.SshUsername
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

type NullableUpdateRotatedSecret struct {
	value *UpdateRotatedSecret
	isSet bool
}

func (v NullableUpdateRotatedSecret) Get() *UpdateRotatedSecret {
	return v.value
}

func (v *NullableUpdateRotatedSecret) Set(val *UpdateRotatedSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRotatedSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRotatedSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRotatedSecret(val *UpdateRotatedSecret) *NullableUpdateRotatedSecret {
	return &NullableUpdateRotatedSecret{value: val, isSet: true}
}

func (v NullableUpdateRotatedSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRotatedSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


