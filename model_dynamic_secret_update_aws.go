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

// DynamicSecretUpdateAws dynamicSecretUpdateAws is a command that Updates aws dynamic secret
type DynamicSecretUpdateAws struct {
	AccessMode *string `json:"access-mode,omitempty"`
	// Admin credentials rotation interval (days)
	AdminRotationIntervalDays *int64 `json:"admin-rotation-interval-days,omitempty"`
	// Access Key ID
	AwsAccessKeyId *string `json:"aws-access-key-id,omitempty"`
	// Secret Access Key
	AwsAccessSecretKey *string `json:"aws-access-secret-key,omitempty"`
	// AWS Role ARNs to be used in the Assume Role operation (relevant only for assume_role mode)
	AwsRoleArns *string `json:"aws-role-arns,omitempty"`
	// AWS User console access
	AwsUserConsoleAccess *bool `json:"aws-user-console-access,omitempty"`
	// AWS User groups
	AwsUserGroups *string `json:"aws-user-groups,omitempty"`
	// AWS User policies
	AwsUserPolicies *string `json:"aws-user-policies,omitempty"`
	// Enable AWS User programmatic access
	AwsUserProgrammaticAccess *bool `json:"aws-user-programmatic-access,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Automatic admin credentials rotation
	EnableAdminRotation *bool `json:"enable-admin-rotation,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret new name
	NewName *string `json:"new-name,omitempty"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Region
	Region *string `json:"region,omitempty"`
	// The AWS account id
	SecureAccessAwsAccountId *string `json:"secure-access-aws-account-id,omitempty"`
	// The AWS native cli
	SecureAccessAwsNativeCli *bool `json:"secure-access-aws-native-cli,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// Secure browser via Akeyless Web Access Bastion
	SecureAccessWebBrowsing *bool `json:"secure-access-web-browsing,omitempty"`
	// Web-Proxy via Akeyless Web Access Bastion
	SecureAccessWebProxy *bool `json:"secure-access-web-proxy,omitempty"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewDynamicSecretUpdateAws instantiates a new DynamicSecretUpdateAws object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdateAws(name string, ) *DynamicSecretUpdateAws {
	this := DynamicSecretUpdateAws{}
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var awsUserConsoleAccess bool = false
	this.AwsUserConsoleAccess = &awsUserConsoleAccess
	var awsUserProgrammaticAccess bool = true
	this.AwsUserProgrammaticAccess = &awsUserProgrammaticAccess
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	var json bool = false
	this.Json = &json
	this.Name = name
	var region string = "us-east-2"
	this.Region = &region
	var secureAccessWeb bool = true
	this.SecureAccessWeb = &secureAccessWeb
	var secureAccessWebBrowsing bool = false
	this.SecureAccessWebBrowsing = &secureAccessWebBrowsing
	var secureAccessWebProxy bool = false
	this.SecureAccessWebProxy = &secureAccessWebProxy
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretUpdateAwsWithDefaults instantiates a new DynamicSecretUpdateAws object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdateAwsWithDefaults() *DynamicSecretUpdateAws {
	this := DynamicSecretUpdateAws{}
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var awsUserConsoleAccess bool = false
	this.AwsUserConsoleAccess = &awsUserConsoleAccess
	var awsUserProgrammaticAccess bool = true
	this.AwsUserProgrammaticAccess = &awsUserProgrammaticAccess
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	var json bool = false
	this.Json = &json
	var region string = "us-east-2"
	this.Region = &region
	var secureAccessWeb bool = true
	this.SecureAccessWeb = &secureAccessWeb
	var secureAccessWebBrowsing bool = false
	this.SecureAccessWebBrowsing = &secureAccessWebBrowsing
	var secureAccessWebProxy bool = false
	this.SecureAccessWebProxy = &secureAccessWebProxy
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetAccessMode returns the AccessMode field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetAccessMode() string {
	if o == nil || o.AccessMode == nil {
		var ret string
		return ret
	}
	return *o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetAccessModeOk() (*string, bool) {
	if o == nil || o.AccessMode == nil {
		return nil, false
	}
	return o.AccessMode, true
}

// HasAccessMode returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasAccessMode() bool {
	if o != nil && o.AccessMode != nil {
		return true
	}

	return false
}

// SetAccessMode gets a reference to the given string and assigns it to the AccessMode field.
func (o *DynamicSecretUpdateAws) SetAccessMode(v string) {
	o.AccessMode = &v
}

// GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetAdminRotationIntervalDays() int64 {
	if o == nil || o.AdminRotationIntervalDays == nil {
		var ret int64
		return ret
	}
	return *o.AdminRotationIntervalDays
}

// GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetAdminRotationIntervalDaysOk() (*int64, bool) {
	if o == nil || o.AdminRotationIntervalDays == nil {
		return nil, false
	}
	return o.AdminRotationIntervalDays, true
}

// HasAdminRotationIntervalDays returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasAdminRotationIntervalDays() bool {
	if o != nil && o.AdminRotationIntervalDays != nil {
		return true
	}

	return false
}

// SetAdminRotationIntervalDays gets a reference to the given int64 and assigns it to the AdminRotationIntervalDays field.
func (o *DynamicSecretUpdateAws) SetAdminRotationIntervalDays(v int64) {
	o.AdminRotationIntervalDays = &v
}

// GetAwsAccessKeyId returns the AwsAccessKeyId field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetAwsAccessKeyId() string {
	if o == nil || o.AwsAccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyId
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AwsAccessKeyId == nil {
		return nil, false
	}
	return o.AwsAccessKeyId, true
}

// HasAwsAccessKeyId returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasAwsAccessKeyId() bool {
	if o != nil && o.AwsAccessKeyId != nil {
		return true
	}

	return false
}

// SetAwsAccessKeyId gets a reference to the given string and assigns it to the AwsAccessKeyId field.
func (o *DynamicSecretUpdateAws) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId = &v
}

// GetAwsAccessSecretKey returns the AwsAccessSecretKey field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetAwsAccessSecretKey() string {
	if o == nil || o.AwsAccessSecretKey == nil {
		var ret string
		return ret
	}
	return *o.AwsAccessSecretKey
}

// GetAwsAccessSecretKeyOk returns a tuple with the AwsAccessSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetAwsAccessSecretKeyOk() (*string, bool) {
	if o == nil || o.AwsAccessSecretKey == nil {
		return nil, false
	}
	return o.AwsAccessSecretKey, true
}

// HasAwsAccessSecretKey returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasAwsAccessSecretKey() bool {
	if o != nil && o.AwsAccessSecretKey != nil {
		return true
	}

	return false
}

// SetAwsAccessSecretKey gets a reference to the given string and assigns it to the AwsAccessSecretKey field.
func (o *DynamicSecretUpdateAws) SetAwsAccessSecretKey(v string) {
	o.AwsAccessSecretKey = &v
}

// GetAwsRoleArns returns the AwsRoleArns field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetAwsRoleArns() string {
	if o == nil || o.AwsRoleArns == nil {
		var ret string
		return ret
	}
	return *o.AwsRoleArns
}

// GetAwsRoleArnsOk returns a tuple with the AwsRoleArns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetAwsRoleArnsOk() (*string, bool) {
	if o == nil || o.AwsRoleArns == nil {
		return nil, false
	}
	return o.AwsRoleArns, true
}

// HasAwsRoleArns returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasAwsRoleArns() bool {
	if o != nil && o.AwsRoleArns != nil {
		return true
	}

	return false
}

// SetAwsRoleArns gets a reference to the given string and assigns it to the AwsRoleArns field.
func (o *DynamicSecretUpdateAws) SetAwsRoleArns(v string) {
	o.AwsRoleArns = &v
}

// GetAwsUserConsoleAccess returns the AwsUserConsoleAccess field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetAwsUserConsoleAccess() bool {
	if o == nil || o.AwsUserConsoleAccess == nil {
		var ret bool
		return ret
	}
	return *o.AwsUserConsoleAccess
}

// GetAwsUserConsoleAccessOk returns a tuple with the AwsUserConsoleAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetAwsUserConsoleAccessOk() (*bool, bool) {
	if o == nil || o.AwsUserConsoleAccess == nil {
		return nil, false
	}
	return o.AwsUserConsoleAccess, true
}

// HasAwsUserConsoleAccess returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasAwsUserConsoleAccess() bool {
	if o != nil && o.AwsUserConsoleAccess != nil {
		return true
	}

	return false
}

// SetAwsUserConsoleAccess gets a reference to the given bool and assigns it to the AwsUserConsoleAccess field.
func (o *DynamicSecretUpdateAws) SetAwsUserConsoleAccess(v bool) {
	o.AwsUserConsoleAccess = &v
}

// GetAwsUserGroups returns the AwsUserGroups field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetAwsUserGroups() string {
	if o == nil || o.AwsUserGroups == nil {
		var ret string
		return ret
	}
	return *o.AwsUserGroups
}

// GetAwsUserGroupsOk returns a tuple with the AwsUserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetAwsUserGroupsOk() (*string, bool) {
	if o == nil || o.AwsUserGroups == nil {
		return nil, false
	}
	return o.AwsUserGroups, true
}

// HasAwsUserGroups returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasAwsUserGroups() bool {
	if o != nil && o.AwsUserGroups != nil {
		return true
	}

	return false
}

// SetAwsUserGroups gets a reference to the given string and assigns it to the AwsUserGroups field.
func (o *DynamicSecretUpdateAws) SetAwsUserGroups(v string) {
	o.AwsUserGroups = &v
}

// GetAwsUserPolicies returns the AwsUserPolicies field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetAwsUserPolicies() string {
	if o == nil || o.AwsUserPolicies == nil {
		var ret string
		return ret
	}
	return *o.AwsUserPolicies
}

// GetAwsUserPoliciesOk returns a tuple with the AwsUserPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetAwsUserPoliciesOk() (*string, bool) {
	if o == nil || o.AwsUserPolicies == nil {
		return nil, false
	}
	return o.AwsUserPolicies, true
}

// HasAwsUserPolicies returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasAwsUserPolicies() bool {
	if o != nil && o.AwsUserPolicies != nil {
		return true
	}

	return false
}

// SetAwsUserPolicies gets a reference to the given string and assigns it to the AwsUserPolicies field.
func (o *DynamicSecretUpdateAws) SetAwsUserPolicies(v string) {
	o.AwsUserPolicies = &v
}

// GetAwsUserProgrammaticAccess returns the AwsUserProgrammaticAccess field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetAwsUserProgrammaticAccess() bool {
	if o == nil || o.AwsUserProgrammaticAccess == nil {
		var ret bool
		return ret
	}
	return *o.AwsUserProgrammaticAccess
}

// GetAwsUserProgrammaticAccessOk returns a tuple with the AwsUserProgrammaticAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetAwsUserProgrammaticAccessOk() (*bool, bool) {
	if o == nil || o.AwsUserProgrammaticAccess == nil {
		return nil, false
	}
	return o.AwsUserProgrammaticAccess, true
}

// HasAwsUserProgrammaticAccess returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasAwsUserProgrammaticAccess() bool {
	if o != nil && o.AwsUserProgrammaticAccess != nil {
		return true
	}

	return false
}

// SetAwsUserProgrammaticAccess gets a reference to the given bool and assigns it to the AwsUserProgrammaticAccess field.
func (o *DynamicSecretUpdateAws) SetAwsUserProgrammaticAccess(v bool) {
	o.AwsUserProgrammaticAccess = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretUpdateAws) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretUpdateAws) SetDescription(v string) {
	o.Description = &v
}

// GetEnableAdminRotation returns the EnableAdminRotation field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetEnableAdminRotation() bool {
	if o == nil || o.EnableAdminRotation == nil {
		var ret bool
		return ret
	}
	return *o.EnableAdminRotation
}

// GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetEnableAdminRotationOk() (*bool, bool) {
	if o == nil || o.EnableAdminRotation == nil {
		return nil, false
	}
	return o.EnableAdminRotation, true
}

// HasEnableAdminRotation returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasEnableAdminRotation() bool {
	if o != nil && o.EnableAdminRotation != nil {
		return true
	}

	return false
}

// SetEnableAdminRotation gets a reference to the given bool and assigns it to the EnableAdminRotation field.
func (o *DynamicSecretUpdateAws) SetEnableAdminRotation(v bool) {
	o.EnableAdminRotation = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretUpdateAws) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretUpdateAws) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretUpdateAws) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DynamicSecretUpdateAws) SetNewName(v string) {
	o.NewName = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *DynamicSecretUpdateAws) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretUpdateAws) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *DynamicSecretUpdateAws) SetRegion(v string) {
	o.Region = &v
}

// GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetSecureAccessAwsAccountId() string {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessAwsAccountId
}

// GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetSecureAccessAwsAccountIdOk() (*string, bool) {
	if o == nil || o.SecureAccessAwsAccountId == nil {
		return nil, false
	}
	return o.SecureAccessAwsAccountId, true
}

// HasSecureAccessAwsAccountId returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasSecureAccessAwsAccountId() bool {
	if o != nil && o.SecureAccessAwsAccountId != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsAccountId gets a reference to the given string and assigns it to the SecureAccessAwsAccountId field.
func (o *DynamicSecretUpdateAws) SetSecureAccessAwsAccountId(v string) {
	o.SecureAccessAwsAccountId = &v
}

// GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetSecureAccessAwsNativeCli() bool {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAwsNativeCli
}

// GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetSecureAccessAwsNativeCliOk() (*bool, bool) {
	if o == nil || o.SecureAccessAwsNativeCli == nil {
		return nil, false
	}
	return o.SecureAccessAwsNativeCli, true
}

// HasSecureAccessAwsNativeCli returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasSecureAccessAwsNativeCli() bool {
	if o != nil && o.SecureAccessAwsNativeCli != nil {
		return true
	}

	return false
}

// SetSecureAccessAwsNativeCli gets a reference to the given bool and assigns it to the SecureAccessAwsNativeCli field.
func (o *DynamicSecretUpdateAws) SetSecureAccessAwsNativeCli(v bool) {
	o.SecureAccessAwsNativeCli = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *DynamicSecretUpdateAws) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DynamicSecretUpdateAws) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *DynamicSecretUpdateAws) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetSecureAccessWebBrowsing() bool {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebBrowsing
}

// GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetSecureAccessWebBrowsingOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		return nil, false
	}
	return o.SecureAccessWebBrowsing, true
}

// HasSecureAccessWebBrowsing returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasSecureAccessWebBrowsing() bool {
	if o != nil && o.SecureAccessWebBrowsing != nil {
		return true
	}

	return false
}

// SetSecureAccessWebBrowsing gets a reference to the given bool and assigns it to the SecureAccessWebBrowsing field.
func (o *DynamicSecretUpdateAws) SetSecureAccessWebBrowsing(v bool) {
	o.SecureAccessWebBrowsing = &v
}

// GetSecureAccessWebProxy returns the SecureAccessWebProxy field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetSecureAccessWebProxy() bool {
	if o == nil || o.SecureAccessWebProxy == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebProxy
}

// GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetSecureAccessWebProxyOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebProxy == nil {
		return nil, false
	}
	return o.SecureAccessWebProxy, true
}

// HasSecureAccessWebProxy returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasSecureAccessWebProxy() bool {
	if o != nil && o.SecureAccessWebProxy != nil {
		return true
	}

	return false
}

// SetSecureAccessWebProxy gets a reference to the given bool and assigns it to the SecureAccessWebProxy field.
func (o *DynamicSecretUpdateAws) SetSecureAccessWebProxy(v bool) {
	o.SecureAccessWebProxy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretUpdateAws) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretUpdateAws) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretUpdateAws) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretUpdateAws) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateAws) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateAws) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateAws) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretUpdateAws) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretUpdateAws) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessMode != nil {
		toSerialize["access-mode"] = o.AccessMode
	}
	if o.AdminRotationIntervalDays != nil {
		toSerialize["admin-rotation-interval-days"] = o.AdminRotationIntervalDays
	}
	if o.AwsAccessKeyId != nil {
		toSerialize["aws-access-key-id"] = o.AwsAccessKeyId
	}
	if o.AwsAccessSecretKey != nil {
		toSerialize["aws-access-secret-key"] = o.AwsAccessSecretKey
	}
	if o.AwsRoleArns != nil {
		toSerialize["aws-role-arns"] = o.AwsRoleArns
	}
	if o.AwsUserConsoleAccess != nil {
		toSerialize["aws-user-console-access"] = o.AwsUserConsoleAccess
	}
	if o.AwsUserGroups != nil {
		toSerialize["aws-user-groups"] = o.AwsUserGroups
	}
	if o.AwsUserPolicies != nil {
		toSerialize["aws-user-policies"] = o.AwsUserPolicies
	}
	if o.AwsUserProgrammaticAccess != nil {
		toSerialize["aws-user-programmatic-access"] = o.AwsUserProgrammaticAccess
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EnableAdminRotation != nil {
		toSerialize["enable-admin-rotation"] = o.EnableAdminRotation
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
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
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
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
	if o.SecureAccessWeb != nil {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
	}
	if o.SecureAccessWebBrowsing != nil {
		toSerialize["secure-access-web-browsing"] = o.SecureAccessWebBrowsing
	}
	if o.SecureAccessWebProxy != nil {
		toSerialize["secure-access-web-proxy"] = o.SecureAccessWebProxy
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetName != nil {
		toSerialize["target-name"] = o.TargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretUpdateAws struct {
	value *DynamicSecretUpdateAws
	isSet bool
}

func (v NullableDynamicSecretUpdateAws) Get() *DynamicSecretUpdateAws {
	return v.value
}

func (v *NullableDynamicSecretUpdateAws) Set(val *DynamicSecretUpdateAws) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdateAws) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdateAws) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdateAws(val *DynamicSecretUpdateAws) *NullableDynamicSecretUpdateAws {
	return &NullableDynamicSecretUpdateAws{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdateAws) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdateAws) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


