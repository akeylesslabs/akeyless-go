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

// RotatedSecretUpdateAws struct for RotatedSecretUpdateAws
type RotatedSecretUpdateAws struct {
	// List of the new tags that will be attached to this item
	AddTag *[]string `json:"add-tag,omitempty"`
	// API ID to rotate (relevant only for rotator-type=api-key)
	ApiId *string `json:"api-id,omitempty"`
	// API key to rotate (relevant only for rotator-type=api-key)
	ApiKey *string `json:"api-key,omitempty"`
	// The credentials to connect with use-user-creds/use-target-creds
	AuthenticationCredentials *string `json:"authentication-credentials,omitempty"`
	// Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation [true/false]
	AutoRotate *string `json:"auto-rotate,omitempty"`
	// Aws Region
	AwsRegion *string `json:"aws-region,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Create a new access key without deleting the old key from AWS for backup (relevant only for AWS) [true/false]
	GraceRotation *string `json:"grace-rotation,omitempty"`
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
	// Rotate the value of the secret after SRA session ends [true/false]
	RotateAfterDisconnect *string `json:"rotate-after-disconnect,omitempty"`
	// The Hour of the rotation in UTC
	RotationHour *int32 `json:"rotation-hour,omitempty"`
	// The number of days to wait between every automatic key rotation (1-365)
	RotationInterval *string `json:"rotation-interval,omitempty"`
	// The AWS account id
	SecureAccessAwsAccountId *string `json:"secure-access-aws-account-id,omitempty"`
	// The AWS native cli
	SecureAccessAwsNativeCli *bool `json:"secure-access-aws-native-cli,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewRotatedSecretUpdateAws instantiates a new RotatedSecretUpdateAws object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotatedSecretUpdateAws(name string, ) *RotatedSecretUpdateAws {
	this := RotatedSecretUpdateAws{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var awsRegion string = "us-east-2"
	this.AwsRegion = &awsRegion
	var description string = "default_metadata"
	this.Description = &description
	var json bool = false
	this.Json = &json
	this.Name = name
	var rotateAfterDisconnect string = "false"
	this.RotateAfterDisconnect = &rotateAfterDisconnect
	return &this
}

// NewRotatedSecretUpdateAwsWithDefaults instantiates a new RotatedSecretUpdateAws object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotatedSecretUpdateAwsWithDefaults() *RotatedSecretUpdateAws {
	this := RotatedSecretUpdateAws{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var awsRegion string = "us-east-2"
	this.AwsRegion = &awsRegion
	var description string = "default_metadata"
	this.Description = &description
	var json bool = false
	this.Json = &json
	var rotateAfterDisconnect string = "false"
	this.RotateAfterDisconnect = &rotateAfterDisconnect
	return &this
}

// GetAddTag returns the AddTag field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetAddTag() []string {
	if o == nil || o.AddTag == nil {
		var ret []string
		return ret
	}
	return *o.AddTag
}

// GetAddTagOk returns a tuple with the AddTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetAddTagOk() (*[]string, bool) {
	if o == nil || o.AddTag == nil {
		return nil, false
	}
	return o.AddTag, true
}

// HasAddTag returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasAddTag() bool {
	if o != nil && o.AddTag != nil {
		return true
	}

	return false
}

// SetAddTag gets a reference to the given []string and assigns it to the AddTag field.
func (o *RotatedSecretUpdateAws) SetAddTag(v []string) {
	o.AddTag = &v
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetApiId() string {
	if o == nil || o.ApiId == nil {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetApiIdOk() (*string, bool) {
	if o == nil || o.ApiId == nil {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasApiId() bool {
	if o != nil && o.ApiId != nil {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *RotatedSecretUpdateAws) SetApiId(v string) {
	o.ApiId = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *RotatedSecretUpdateAws) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetAuthenticationCredentials() string {
	if o == nil || o.AuthenticationCredentials == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil || o.AuthenticationCredentials == nil {
		return nil, false
	}
	return o.AuthenticationCredentials, true
}

// HasAuthenticationCredentials returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasAuthenticationCredentials() bool {
	if o != nil && o.AuthenticationCredentials != nil {
		return true
	}

	return false
}

// SetAuthenticationCredentials gets a reference to the given string and assigns it to the AuthenticationCredentials field.
func (o *RotatedSecretUpdateAws) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = &v
}

// GetAutoRotate returns the AutoRotate field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetAutoRotate() string {
	if o == nil || o.AutoRotate == nil {
		var ret string
		return ret
	}
	return *o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetAutoRotateOk() (*string, bool) {
	if o == nil || o.AutoRotate == nil {
		return nil, false
	}
	return o.AutoRotate, true
}

// HasAutoRotate returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasAutoRotate() bool {
	if o != nil && o.AutoRotate != nil {
		return true
	}

	return false
}

// SetAutoRotate gets a reference to the given string and assigns it to the AutoRotate field.
func (o *RotatedSecretUpdateAws) SetAutoRotate(v string) {
	o.AutoRotate = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetAwsRegion() string {
	if o == nil || o.AwsRegion == nil {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetAwsRegionOk() (*string, bool) {
	if o == nil || o.AwsRegion == nil {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasAwsRegion() bool {
	if o != nil && o.AwsRegion != nil {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *RotatedSecretUpdateAws) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *RotatedSecretUpdateAws) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RotatedSecretUpdateAws) SetDescription(v string) {
	o.Description = &v
}

// GetGraceRotation returns the GraceRotation field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetGraceRotation() string {
	if o == nil || o.GraceRotation == nil {
		var ret string
		return ret
	}
	return *o.GraceRotation
}

// GetGraceRotationOk returns a tuple with the GraceRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetGraceRotationOk() (*string, bool) {
	if o == nil || o.GraceRotation == nil {
		return nil, false
	}
	return o.GraceRotation, true
}

// HasGraceRotation returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasGraceRotation() bool {
	if o != nil && o.GraceRotation != nil {
		return true
	}

	return false
}

// SetGraceRotation gets a reference to the given string and assigns it to the GraceRotation field.
func (o *RotatedSecretUpdateAws) SetGraceRotation(v string) {
	o.GraceRotation = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *RotatedSecretUpdateAws) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *RotatedSecretUpdateAws) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RotatedSecretUpdateAws) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *RotatedSecretUpdateAws) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetName returns the Name field value
func (o *RotatedSecretUpdateAws) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RotatedSecretUpdateAws) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *RotatedSecretUpdateAws) SetNewName(v string) {
	o.NewName = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *RotatedSecretUpdateAws) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetRmTag returns the RmTag field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetRmTag() []string {
	if o == nil || o.RmTag == nil {
		var ret []string
		return ret
	}
	return *o.RmTag
}

// GetRmTagOk returns a tuple with the RmTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetRmTagOk() (*[]string, bool) {
	if o == nil || o.RmTag == nil {
		return nil, false
	}
	return o.RmTag, true
}

// HasRmTag returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasRmTag() bool {
	if o != nil && o.RmTag != nil {
		return true
	}

	return false
}

// SetRmTag gets a reference to the given []string and assigns it to the RmTag field.
func (o *RotatedSecretUpdateAws) SetRmTag(v []string) {
	o.RmTag = &v
}

// GetRotateAfterDisconnect returns the RotateAfterDisconnect field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetRotateAfterDisconnect() string {
	if o == nil || o.RotateAfterDisconnect == nil {
		var ret string
		return ret
	}
	return *o.RotateAfterDisconnect
}

// GetRotateAfterDisconnectOk returns a tuple with the RotateAfterDisconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetRotateAfterDisconnectOk() (*string, bool) {
	if o == nil || o.RotateAfterDisconnect == nil {
		return nil, false
	}
	return o.RotateAfterDisconnect, true
}

// HasRotateAfterDisconnect returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasRotateAfterDisconnect() bool {
	if o != nil && o.RotateAfterDisconnect != nil {
		return true
	}

	return false
}

// SetRotateAfterDisconnect gets a reference to the given string and assigns it to the RotateAfterDisconnect field.
func (o *RotatedSecretUpdateAws) SetRotateAfterDisconnect(v string) {
	o.RotateAfterDisconnect = &v
}

// GetRotationHour returns the RotationHour field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetRotationHour() int32 {
	if o == nil || o.RotationHour == nil {
		var ret int32
		return ret
	}
	return *o.RotationHour
}

// GetRotationHourOk returns a tuple with the RotationHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetRotationHourOk() (*int32, bool) {
	if o == nil || o.RotationHour == nil {
		return nil, false
	}
	return o.RotationHour, true
}

// HasRotationHour returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasRotationHour() bool {
	if o != nil && o.RotationHour != nil {
		return true
	}

	return false
}

// SetRotationHour gets a reference to the given int32 and assigns it to the RotationHour field.
func (o *RotatedSecretUpdateAws) SetRotationHour(v int32) {
	o.RotationHour = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetRotationInterval() string {
	if o == nil || o.RotationInterval == nil {
		var ret string
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetRotationIntervalOk() (*string, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given string and assigns it to the RotationInterval field.
func (o *RotatedSecretUpdateAws) SetRotationInterval(v string) {
	o.RotationInterval = &v
}

// GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetSecureAccessAwsAccountId() string {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessAwsAccountId
}

// GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetSecureAccessAwsAccountIdOk() (*string, bool) {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		return nil, false
	}
	return o.SecureAccessAwsAccountId, true
}

// HasSecureAccessAwsAccountId returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasSecureAccessAwsAccountId() bool {
	if o != nil && o.SecureAccessAwsAccountId != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsAccountId gets a reference to the given string and assigns it to the SecureAccessAwsAccountId field.
func (o *RotatedSecretUpdateAws) SetSecureAccessAwsAccountId(v string) {
	o.SecureAccessAwsAccountId = &v
}

// GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetSecureAccessAwsNativeCli() bool {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAwsNativeCli
}

// GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetSecureAccessAwsNativeCliOk() (*bool, bool) {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		return nil, false
	}
	return o.SecureAccessAwsNativeCli, true
}

// HasSecureAccessAwsNativeCli returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasSecureAccessAwsNativeCli() bool {
	if o != nil && o.SecureAccessAwsNativeCli != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsNativeCli gets a reference to the given bool and assigns it to the SecureAccessAwsNativeCli field.
func (o *RotatedSecretUpdateAws) SetSecureAccessAwsNativeCli(v bool) {
	o.SecureAccessAwsNativeCli = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *RotatedSecretUpdateAws) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *RotatedSecretUpdateAws) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RotatedSecretUpdateAws) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *RotatedSecretUpdateAws) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretUpdateAws) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *RotatedSecretUpdateAws) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *RotatedSecretUpdateAws) SetUidToken(v string) {
	o.UidToken = &v
}

func (o RotatedSecretUpdateAws) MarshalJSON() ([]byte, error) {
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
	if o.AuthenticationCredentials != nil {
		toSerialize["authentication-credentials"] = o.AuthenticationCredentials
	}
	if o.AutoRotate != nil {
		toSerialize["auto-rotate"] = o.AutoRotate
	}
	if o.AwsRegion != nil {
		toSerialize["aws-region"] = o.AwsRegion
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GraceRotation != nil {
		toSerialize["grace-rotation"] = o.GraceRotation
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
	if o.RotateAfterDisconnect != nil {
		toSerialize["rotate-after-disconnect"] = o.RotateAfterDisconnect
	}
	if o.RotationHour != nil {
		toSerialize["rotation-hour"] = o.RotationHour
	}
	if o.RotationInterval != nil {
		toSerialize["rotation-interval"] = o.RotationInterval
	}
	if o.SecureAccessAwsAccountId != nil {
		toSerialize["secure-access-aws-account-id"] = o.SecureAccessAwsAccountId
	}
	if o.SecureAccessAwsNativeCli != nil {
		toSerialize["secure-access-aws-native-cli"] = o.SecureAccessAwsNativeCli
	}
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableRotatedSecretUpdateAws struct {
	value *RotatedSecretUpdateAws
	isSet bool
}

func (v NullableRotatedSecretUpdateAws) Get() *RotatedSecretUpdateAws {
	return v.value
}

func (v *NullableRotatedSecretUpdateAws) Set(val *RotatedSecretUpdateAws) {
	v.value = val
	v.isSet = true
}

func (v NullableRotatedSecretUpdateAws) IsSet() bool {
	return v.isSet
}

func (v *NullableRotatedSecretUpdateAws) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotatedSecretUpdateAws(val *RotatedSecretUpdateAws) *NullableRotatedSecretUpdateAws {
	return &NullableRotatedSecretUpdateAws{value: val, isSet: true}
}

func (v NullableRotatedSecretUpdateAws) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotatedSecretUpdateAws) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


