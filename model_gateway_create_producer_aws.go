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

// GatewayCreateProducerAws gatewayCreateProducerAws is a command that creates aws producer
type GatewayCreateProducerAws struct {
	// Access Key ID
	AccessKeyId string `json:"access-key-id"`
	AccessMode *string `json:"access-mode,omitempty"`
	// Secret Access Key
	AccessSecretKey string `json:"access-secret-key"`
	// Admin credentials rotation interval (days)
	AdminRotationIntervalDays *int64 `json:"admin-rotation-interval-days,omitempty"`
	// AWS Role ARNs to be used in the Assume Role operation (relevant only for assume_role mode)
	AwsRoleArns *string `json:"aws-role-arns,omitempty"`
	// AWS User console access
	AwsUserConsoleAccess *bool `json:"aws-user-console-access,omitempty"`
	// AWS User groups
	AwsUserGroups *string `json:"aws-user-groups,omitempty"`
	// AWS User policies
	AwsUserPolicies *string `json:"aws-user-policies,omitempty"`
	// AWS User programmatic access
	AwsUserProgrammaticAccess *bool `json:"aws-user-programmatic-access,omitempty"`
	// Automatic admin credentials rotation
	EnableAdminRotation *bool `json:"enable-admin-rotation,omitempty"`
	// Gateway url
	GatewayUrl *string `json:"gateway-url,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Region
	Region *string `json:"region,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewGatewayCreateProducerAws instantiates a new GatewayCreateProducerAws object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerAws(accessKeyId string, accessSecretKey string, name string, ) *GatewayCreateProducerAws {
	this := GatewayCreateProducerAws{}
	this.AccessKeyId = accessKeyId
	this.AccessSecretKey = accessSecretKey
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var awsUserConsoleAccess bool = false
	this.AwsUserConsoleAccess = &awsUserConsoleAccess
	var awsUserProgrammaticAccess bool = true
	this.AwsUserProgrammaticAccess = &awsUserProgrammaticAccess
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	this.Name = name
	var region string = "us-east-2"
	this.Region = &region
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerAwsWithDefaults instantiates a new GatewayCreateProducerAws object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerAwsWithDefaults() *GatewayCreateProducerAws {
	this := GatewayCreateProducerAws{}
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var awsUserConsoleAccess bool = false
	this.AwsUserConsoleAccess = &awsUserConsoleAccess
	var awsUserProgrammaticAccess bool = true
	this.AwsUserProgrammaticAccess = &awsUserProgrammaticAccess
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	var region string = "us-east-2"
	this.Region = &region
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *GatewayCreateProducerAws) GetAccessKeyId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetAccessKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value
func (o *GatewayCreateProducerAws) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetAccessMode returns the AccessMode field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetAccessMode() string {
	if o == nil || o.AccessMode == nil {
		var ret string
		return ret
	}
	return *o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetAccessModeOk() (*string, bool) {
	if o == nil || o.AccessMode == nil {
		return nil, false
	}
	return o.AccessMode, true
}

// HasAccessMode returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasAccessMode() bool {
	if o != nil && o.AccessMode != nil {
		return true
	}

	return false
}

// SetAccessMode gets a reference to the given string and assigns it to the AccessMode field.
func (o *GatewayCreateProducerAws) SetAccessMode(v string) {
	o.AccessMode = &v
}

// GetAccessSecretKey returns the AccessSecretKey field value
func (o *GatewayCreateProducerAws) GetAccessSecretKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccessSecretKey
}

// GetAccessSecretKeyOk returns a tuple with the AccessSecretKey field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetAccessSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessSecretKey, true
}

// SetAccessSecretKey sets field value
func (o *GatewayCreateProducerAws) SetAccessSecretKey(v string) {
	o.AccessSecretKey = v
}

// GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetAdminRotationIntervalDays() int64 {
	if o == nil || o.AdminRotationIntervalDays == nil {
		var ret int64
		return ret
	}
	return *o.AdminRotationIntervalDays
}

// GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetAdminRotationIntervalDaysOk() (*int64, bool) {
	if o == nil || o.AdminRotationIntervalDays == nil {
		return nil, false
	}
	return o.AdminRotationIntervalDays, true
}

// HasAdminRotationIntervalDays returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasAdminRotationIntervalDays() bool {
	if o != nil && o.AdminRotationIntervalDays != nil {
		return true
	}

	return false
}

// SetAdminRotationIntervalDays gets a reference to the given int64 and assigns it to the AdminRotationIntervalDays field.
func (o *GatewayCreateProducerAws) SetAdminRotationIntervalDays(v int64) {
	o.AdminRotationIntervalDays = &v
}

// GetAwsRoleArns returns the AwsRoleArns field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetAwsRoleArns() string {
	if o == nil || o.AwsRoleArns == nil {
		var ret string
		return ret
	}
	return *o.AwsRoleArns
}

// GetAwsRoleArnsOk returns a tuple with the AwsRoleArns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetAwsRoleArnsOk() (*string, bool) {
	if o == nil || o.AwsRoleArns == nil {
		return nil, false
	}
	return o.AwsRoleArns, true
}

// HasAwsRoleArns returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasAwsRoleArns() bool {
	if o != nil && o.AwsRoleArns != nil {
		return true
	}

	return false
}

// SetAwsRoleArns gets a reference to the given string and assigns it to the AwsRoleArns field.
func (o *GatewayCreateProducerAws) SetAwsRoleArns(v string) {
	o.AwsRoleArns = &v
}

// GetAwsUserConsoleAccess returns the AwsUserConsoleAccess field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetAwsUserConsoleAccess() bool {
	if o == nil || o.AwsUserConsoleAccess == nil {
		var ret bool
		return ret
	}
	return *o.AwsUserConsoleAccess
}

// GetAwsUserConsoleAccessOk returns a tuple with the AwsUserConsoleAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetAwsUserConsoleAccessOk() (*bool, bool) {
	if o == nil || o.AwsUserConsoleAccess == nil {
		return nil, false
	}
	return o.AwsUserConsoleAccess, true
}

// HasAwsUserConsoleAccess returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasAwsUserConsoleAccess() bool {
	if o != nil && o.AwsUserConsoleAccess != nil {
		return true
	}

	return false
}

// SetAwsUserConsoleAccess gets a reference to the given bool and assigns it to the AwsUserConsoleAccess field.
func (o *GatewayCreateProducerAws) SetAwsUserConsoleAccess(v bool) {
	o.AwsUserConsoleAccess = &v
}

// GetAwsUserGroups returns the AwsUserGroups field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetAwsUserGroups() string {
	if o == nil || o.AwsUserGroups == nil {
		var ret string
		return ret
	}
	return *o.AwsUserGroups
}

// GetAwsUserGroupsOk returns a tuple with the AwsUserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetAwsUserGroupsOk() (*string, bool) {
	if o == nil || o.AwsUserGroups == nil {
		return nil, false
	}
	return o.AwsUserGroups, true
}

// HasAwsUserGroups returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasAwsUserGroups() bool {
	if o != nil && o.AwsUserGroups != nil {
		return true
	}

	return false
}

// SetAwsUserGroups gets a reference to the given string and assigns it to the AwsUserGroups field.
func (o *GatewayCreateProducerAws) SetAwsUserGroups(v string) {
	o.AwsUserGroups = &v
}

// GetAwsUserPolicies returns the AwsUserPolicies field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetAwsUserPolicies() string {
	if o == nil || o.AwsUserPolicies == nil {
		var ret string
		return ret
	}
	return *o.AwsUserPolicies
}

// GetAwsUserPoliciesOk returns a tuple with the AwsUserPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetAwsUserPoliciesOk() (*string, bool) {
	if o == nil || o.AwsUserPolicies == nil {
		return nil, false
	}
	return o.AwsUserPolicies, true
}

// HasAwsUserPolicies returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasAwsUserPolicies() bool {
	if o != nil && o.AwsUserPolicies != nil {
		return true
	}

	return false
}

// SetAwsUserPolicies gets a reference to the given string and assigns it to the AwsUserPolicies field.
func (o *GatewayCreateProducerAws) SetAwsUserPolicies(v string) {
	o.AwsUserPolicies = &v
}

// GetAwsUserProgrammaticAccess returns the AwsUserProgrammaticAccess field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetAwsUserProgrammaticAccess() bool {
	if o == nil || o.AwsUserProgrammaticAccess == nil {
		var ret bool
		return ret
	}
	return *o.AwsUserProgrammaticAccess
}

// GetAwsUserProgrammaticAccessOk returns a tuple with the AwsUserProgrammaticAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetAwsUserProgrammaticAccessOk() (*bool, bool) {
	if o == nil || o.AwsUserProgrammaticAccess == nil {
		return nil, false
	}
	return o.AwsUserProgrammaticAccess, true
}

// HasAwsUserProgrammaticAccess returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasAwsUserProgrammaticAccess() bool {
	if o != nil && o.AwsUserProgrammaticAccess != nil {
		return true
	}

	return false
}

// SetAwsUserProgrammaticAccess gets a reference to the given bool and assigns it to the AwsUserProgrammaticAccess field.
func (o *GatewayCreateProducerAws) SetAwsUserProgrammaticAccess(v bool) {
	o.AwsUserProgrammaticAccess = &v
}

// GetEnableAdminRotation returns the EnableAdminRotation field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetEnableAdminRotation() bool {
	if o == nil || o.EnableAdminRotation == nil {
		var ret bool
		return ret
	}
	return *o.EnableAdminRotation
}

// GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetEnableAdminRotationOk() (*bool, bool) {
	if o == nil || o.EnableAdminRotation == nil {
		return nil, false
	}
	return o.EnableAdminRotation, true
}

// HasEnableAdminRotation returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasEnableAdminRotation() bool {
	if o != nil && o.EnableAdminRotation != nil {
		return true
	}

	return false
}

// SetEnableAdminRotation gets a reference to the given bool and assigns it to the EnableAdminRotation field.
func (o *GatewayCreateProducerAws) SetEnableAdminRotation(v bool) {
	o.EnableAdminRotation = &v
}

// GetGatewayUrl returns the GatewayUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetGatewayUrl() string {
	if o == nil || o.GatewayUrl == nil {
		var ret string
		return ret
	}
	return *o.GatewayUrl
}

// GetGatewayUrlOk returns a tuple with the GatewayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetGatewayUrlOk() (*string, bool) {
	if o == nil || o.GatewayUrl == nil {
		return nil, false
	}
	return o.GatewayUrl, true
}

// HasGatewayUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasGatewayUrl() bool {
	if o != nil && o.GatewayUrl != nil {
		return true
	}

	return false
}

// SetGatewayUrl gets a reference to the given string and assigns it to the GatewayUrl field.
func (o *GatewayCreateProducerAws) SetGatewayUrl(v string) {
	o.GatewayUrl = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerAws) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerAws) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayCreateProducerAws) SetPassword(v string) {
	o.Password = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerAws) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GatewayCreateProducerAws) SetRegion(v string) {
	o.Region = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerAws) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerAws) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerAws) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayCreateProducerAws) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAws) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerAws) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayCreateProducerAws) SetUsername(v string) {
	o.Username = &v
}

func (o GatewayCreateProducerAws) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access-key-id"] = o.AccessKeyId
	}
	if o.AccessMode != nil {
		toSerialize["access-mode"] = o.AccessMode
	}
	if true {
		toSerialize["access-secret-key"] = o.AccessSecretKey
	}
	if o.AdminRotationIntervalDays != nil {
		toSerialize["admin-rotation-interval-days"] = o.AdminRotationIntervalDays
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
	if o.EnableAdminRotation != nil {
		toSerialize["enable-admin-rotation"] = o.EnableAdminRotation
	}
	if o.GatewayUrl != nil {
		toSerialize["gateway-url"] = o.GatewayUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerAws struct {
	value *GatewayCreateProducerAws
	isSet bool
}

func (v NullableGatewayCreateProducerAws) Get() *GatewayCreateProducerAws {
	return v.value
}

func (v *NullableGatewayCreateProducerAws) Set(val *GatewayCreateProducerAws) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerAws) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerAws) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerAws(val *GatewayCreateProducerAws) *NullableGatewayCreateProducerAws {
	return &NullableGatewayCreateProducerAws{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerAws) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerAws) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


