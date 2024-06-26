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

// RotatedSecretCreateAws struct for RotatedSecretCreateAws
type RotatedSecretCreateAws struct {
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
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Set the maximum number of versions, limited by the account settings defaults.
	MaxVersions *string `json:"max-versions,omitempty"`
	// Rotated secret name
	Name string `json:"name"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// Rotate the value of the secret after SRA session ends [true/false]
	RotateAfterDisconnect *string `json:"rotate-after-disconnect,omitempty"`
	// The Hour of the rotation in UTC
	RotationHour *int32 `json:"rotation-hour,omitempty"`
	// The number of days to wait between every automatic key rotation (1-365)
	RotationInterval *string `json:"rotation-interval,omitempty"`
	// The rotator type. options: [target/api-key]
	RotatorType string `json:"rotator-type"`
	// The AWS account id
	SecureAccessAwsAccountId *string `json:"secure-access-aws-account-id,omitempty"`
	// The AWS native cli
	SecureAccessAwsNativeCli *bool `json:"secure-access-aws-native-cli,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName string `json:"target-name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewRotatedSecretCreateAws instantiates a new RotatedSecretCreateAws object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotatedSecretCreateAws(name string, rotatorType string, targetName string, ) *RotatedSecretCreateAws {
	this := RotatedSecretCreateAws{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var awsRegion string = "us-east-2"
	this.AwsRegion = &awsRegion
	var json bool = false
	this.Json = &json
	this.Name = name
	var rotateAfterDisconnect string = "false"
	this.RotateAfterDisconnect = &rotateAfterDisconnect
	this.RotatorType = rotatorType
	this.TargetName = targetName
	return &this
}

// NewRotatedSecretCreateAwsWithDefaults instantiates a new RotatedSecretCreateAws object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotatedSecretCreateAwsWithDefaults() *RotatedSecretCreateAws {
	this := RotatedSecretCreateAws{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var awsRegion string = "us-east-2"
	this.AwsRegion = &awsRegion
	var json bool = false
	this.Json = &json
	var rotateAfterDisconnect string = "false"
	this.RotateAfterDisconnect = &rotateAfterDisconnect
	return &this
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetApiId() string {
	if o == nil || o.ApiId == nil {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetApiIdOk() (*string, bool) {
	if o == nil || o.ApiId == nil {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasApiId() bool {
	if o != nil && o.ApiId != nil {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *RotatedSecretCreateAws) SetApiId(v string) {
	o.ApiId = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *RotatedSecretCreateAws) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetAuthenticationCredentials() string {
	if o == nil || o.AuthenticationCredentials == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil || o.AuthenticationCredentials == nil {
		return nil, false
	}
	return o.AuthenticationCredentials, true
}

// HasAuthenticationCredentials returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasAuthenticationCredentials() bool {
	if o != nil && o.AuthenticationCredentials != nil {
		return true
	}

	return false
}

// SetAuthenticationCredentials gets a reference to the given string and assigns it to the AuthenticationCredentials field.
func (o *RotatedSecretCreateAws) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = &v
}

// GetAutoRotate returns the AutoRotate field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetAutoRotate() string {
	if o == nil || o.AutoRotate == nil {
		var ret string
		return ret
	}
	return *o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetAutoRotateOk() (*string, bool) {
	if o == nil || o.AutoRotate == nil {
		return nil, false
	}
	return o.AutoRotate, true
}

// HasAutoRotate returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasAutoRotate() bool {
	if o != nil && o.AutoRotate != nil {
		return true
	}

	return false
}

// SetAutoRotate gets a reference to the given string and assigns it to the AutoRotate field.
func (o *RotatedSecretCreateAws) SetAutoRotate(v string) {
	o.AutoRotate = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetAwsRegion() string {
	if o == nil || o.AwsRegion == nil {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetAwsRegionOk() (*string, bool) {
	if o == nil || o.AwsRegion == nil {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasAwsRegion() bool {
	if o != nil && o.AwsRegion != nil {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *RotatedSecretCreateAws) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *RotatedSecretCreateAws) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RotatedSecretCreateAws) SetDescription(v string) {
	o.Description = &v
}

// GetGraceRotation returns the GraceRotation field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetGraceRotation() string {
	if o == nil || o.GraceRotation == nil {
		var ret string
		return ret
	}
	return *o.GraceRotation
}

// GetGraceRotationOk returns a tuple with the GraceRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetGraceRotationOk() (*string, bool) {
	if o == nil || o.GraceRotation == nil {
		return nil, false
	}
	return o.GraceRotation, true
}

// HasGraceRotation returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasGraceRotation() bool {
	if o != nil && o.GraceRotation != nil {
		return true
	}

	return false
}

// SetGraceRotation gets a reference to the given string and assigns it to the GraceRotation field.
func (o *RotatedSecretCreateAws) SetGraceRotation(v string) {
	o.GraceRotation = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *RotatedSecretCreateAws) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RotatedSecretCreateAws) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *RotatedSecretCreateAws) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetName returns the Name field value
func (o *RotatedSecretCreateAws) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RotatedSecretCreateAws) SetName(v string) {
	o.Name = v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *RotatedSecretCreateAws) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetRotateAfterDisconnect returns the RotateAfterDisconnect field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetRotateAfterDisconnect() string {
	if o == nil || o.RotateAfterDisconnect == nil {
		var ret string
		return ret
	}
	return *o.RotateAfterDisconnect
}

// GetRotateAfterDisconnectOk returns a tuple with the RotateAfterDisconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetRotateAfterDisconnectOk() (*string, bool) {
	if o == nil || o.RotateAfterDisconnect == nil {
		return nil, false
	}
	return o.RotateAfterDisconnect, true
}

// HasRotateAfterDisconnect returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasRotateAfterDisconnect() bool {
	if o != nil && o.RotateAfterDisconnect != nil {
		return true
	}

	return false
}

// SetRotateAfterDisconnect gets a reference to the given string and assigns it to the RotateAfterDisconnect field.
func (o *RotatedSecretCreateAws) SetRotateAfterDisconnect(v string) {
	o.RotateAfterDisconnect = &v
}

// GetRotationHour returns the RotationHour field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetRotationHour() int32 {
	if o == nil || o.RotationHour == nil {
		var ret int32
		return ret
	}
	return *o.RotationHour
}

// GetRotationHourOk returns a tuple with the RotationHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetRotationHourOk() (*int32, bool) {
	if o == nil || o.RotationHour == nil {
		return nil, false
	}
	return o.RotationHour, true
}

// HasRotationHour returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasRotationHour() bool {
	if o != nil && o.RotationHour != nil {
		return true
	}

	return false
}

// SetRotationHour gets a reference to the given int32 and assigns it to the RotationHour field.
func (o *RotatedSecretCreateAws) SetRotationHour(v int32) {
	o.RotationHour = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetRotationInterval() string {
	if o == nil || o.RotationInterval == nil {
		var ret string
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetRotationIntervalOk() (*string, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given string and assigns it to the RotationInterval field.
func (o *RotatedSecretCreateAws) SetRotationInterval(v string) {
	o.RotationInterval = &v
}

// GetRotatorType returns the RotatorType field value
func (o *RotatedSecretCreateAws) GetRotatorType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RotatorType
}

// GetRotatorTypeOk returns a tuple with the RotatorType field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetRotatorTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RotatorType, true
}

// SetRotatorType sets field value
func (o *RotatedSecretCreateAws) SetRotatorType(v string) {
	o.RotatorType = v
}

// GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetSecureAccessAwsAccountId() string {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessAwsAccountId
}

// GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetSecureAccessAwsAccountIdOk() (*string, bool) {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		return nil, false
	}
	return o.SecureAccessAwsAccountId, true
}

// HasSecureAccessAwsAccountId returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasSecureAccessAwsAccountId() bool {
	if o != nil && o.SecureAccessAwsAccountId != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsAccountId gets a reference to the given string and assigns it to the SecureAccessAwsAccountId field.
func (o *RotatedSecretCreateAws) SetSecureAccessAwsAccountId(v string) {
	o.SecureAccessAwsAccountId = &v
}

// GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetSecureAccessAwsNativeCli() bool {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAwsNativeCli
}

// GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetSecureAccessAwsNativeCliOk() (*bool, bool) {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		return nil, false
	}
	return o.SecureAccessAwsNativeCli, true
}

// HasSecureAccessAwsNativeCli returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasSecureAccessAwsNativeCli() bool {
	if o != nil && o.SecureAccessAwsNativeCli != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsNativeCli gets a reference to the given bool and assigns it to the SecureAccessAwsNativeCli field.
func (o *RotatedSecretCreateAws) SetSecureAccessAwsNativeCli(v bool) {
	o.SecureAccessAwsNativeCli = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *RotatedSecretCreateAws) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *RotatedSecretCreateAws) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RotatedSecretCreateAws) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value
func (o *RotatedSecretCreateAws) GetTargetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetTargetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetName, true
}

// SetTargetName sets field value
func (o *RotatedSecretCreateAws) SetTargetName(v string) {
	o.TargetName = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RotatedSecretCreateAws) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *RotatedSecretCreateAws) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateAws) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *RotatedSecretCreateAws) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *RotatedSecretCreateAws) SetUidToken(v string) {
	o.UidToken = &v
}

func (o RotatedSecretCreateAws) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.RotateAfterDisconnect != nil {
		toSerialize["rotate-after-disconnect"] = o.RotateAfterDisconnect
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

type NullableRotatedSecretCreateAws struct {
	value *RotatedSecretCreateAws
	isSet bool
}

func (v NullableRotatedSecretCreateAws) Get() *RotatedSecretCreateAws {
	return v.value
}

func (v *NullableRotatedSecretCreateAws) Set(val *RotatedSecretCreateAws) {
	v.value = val
	v.isSet = true
}

func (v NullableRotatedSecretCreateAws) IsSet() bool {
	return v.isSet
}

func (v *NullableRotatedSecretCreateAws) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotatedSecretCreateAws(val *RotatedSecretCreateAws) *NullableRotatedSecretCreateAws {
	return &NullableRotatedSecretCreateAws{value: val, isSet: true}
}

func (v NullableRotatedSecretCreateAws) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotatedSecretCreateAws) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


