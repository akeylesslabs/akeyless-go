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

// RotatedSecretCreateMysql struct for RotatedSecretCreateMysql
type RotatedSecretCreateMysql struct {
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
	// Rotated secret name
	Name string `json:"name"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// Rotate the value of the secret after SRA session ends [true/false]
	RotateAfterDisconnect *string `json:"rotate-after-disconnect,omitempty"`
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
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// The DB name (relevant only for DB Dynamic-Secret)
	SecureAccessDbName *string `json:"secure-access-db-name,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers)
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName string `json:"target-name"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewRotatedSecretCreateMysql instantiates a new RotatedSecretCreateMysql object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotatedSecretCreateMysql(name string, rotatorType string, targetName string, ) *RotatedSecretCreateMysql {
	this := RotatedSecretCreateMysql{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var json bool = false
	this.Json = &json
	this.Name = name
	var rotateAfterDisconnect string = "false"
	this.RotateAfterDisconnect = &rotateAfterDisconnect
	this.RotatorType = rotatorType
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	this.TargetName = targetName
	return &this
}

// NewRotatedSecretCreateMysqlWithDefaults instantiates a new RotatedSecretCreateMysql object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotatedSecretCreateMysqlWithDefaults() *RotatedSecretCreateMysql {
	this := RotatedSecretCreateMysql{}
	var authenticationCredentials string = "use-user-creds"
	this.AuthenticationCredentials = &authenticationCredentials
	var json bool = false
	this.Json = &json
	var rotateAfterDisconnect string = "false"
	this.RotateAfterDisconnect = &rotateAfterDisconnect
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	return &this
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetAuthenticationCredentials() string {
	if o == nil || o.AuthenticationCredentials == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil || o.AuthenticationCredentials == nil {
		return nil, false
	}
	return o.AuthenticationCredentials, true
}

// HasAuthenticationCredentials returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasAuthenticationCredentials() bool {
	if o != nil && o.AuthenticationCredentials != nil {
		return true
	}

	return false
}

// SetAuthenticationCredentials gets a reference to the given string and assigns it to the AuthenticationCredentials field.
func (o *RotatedSecretCreateMysql) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = &v
}

// GetAutoRotate returns the AutoRotate field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetAutoRotate() string {
	if o == nil || o.AutoRotate == nil {
		var ret string
		return ret
	}
	return *o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetAutoRotateOk() (*string, bool) {
	if o == nil || o.AutoRotate == nil {
		return nil, false
	}
	return o.AutoRotate, true
}

// HasAutoRotate returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasAutoRotate() bool {
	if o != nil && o.AutoRotate != nil {
		return true
	}

	return false
}

// SetAutoRotate gets a reference to the given string and assigns it to the AutoRotate field.
func (o *RotatedSecretCreateMysql) SetAutoRotate(v string) {
	o.AutoRotate = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *RotatedSecretCreateMysql) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RotatedSecretCreateMysql) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *RotatedSecretCreateMysql) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RotatedSecretCreateMysql) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *RotatedSecretCreateMysql) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RotatedSecretCreateMysql) SetName(v string) {
	o.Name = v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *RotatedSecretCreateMysql) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetRotateAfterDisconnect returns the RotateAfterDisconnect field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetRotateAfterDisconnect() string {
	if o == nil || o.RotateAfterDisconnect == nil {
		var ret string
		return ret
	}
	return *o.RotateAfterDisconnect
}

// GetRotateAfterDisconnectOk returns a tuple with the RotateAfterDisconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetRotateAfterDisconnectOk() (*string, bool) {
	if o == nil || o.RotateAfterDisconnect == nil {
		return nil, false
	}
	return o.RotateAfterDisconnect, true
}

// HasRotateAfterDisconnect returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasRotateAfterDisconnect() bool {
	if o != nil && o.RotateAfterDisconnect != nil {
		return true
	}

	return false
}

// SetRotateAfterDisconnect gets a reference to the given string and assigns it to the RotateAfterDisconnect field.
func (o *RotatedSecretCreateMysql) SetRotateAfterDisconnect(v string) {
	o.RotateAfterDisconnect = &v
}

// GetRotatedPassword returns the RotatedPassword field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetRotatedPassword() string {
	if o == nil || o.RotatedPassword == nil {
		var ret string
		return ret
	}
	return *o.RotatedPassword
}

// GetRotatedPasswordOk returns a tuple with the RotatedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetRotatedPasswordOk() (*string, bool) {
	if o == nil || o.RotatedPassword == nil {
		return nil, false
	}
	return o.RotatedPassword, true
}

// HasRotatedPassword returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasRotatedPassword() bool {
	if o != nil && o.RotatedPassword != nil {
		return true
	}

	return false
}

// SetRotatedPassword gets a reference to the given string and assigns it to the RotatedPassword field.
func (o *RotatedSecretCreateMysql) SetRotatedPassword(v string) {
	o.RotatedPassword = &v
}

// GetRotatedUsername returns the RotatedUsername field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetRotatedUsername() string {
	if o == nil || o.RotatedUsername == nil {
		var ret string
		return ret
	}
	return *o.RotatedUsername
}

// GetRotatedUsernameOk returns a tuple with the RotatedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetRotatedUsernameOk() (*string, bool) {
	if o == nil || o.RotatedUsername == nil {
		return nil, false
	}
	return o.RotatedUsername, true
}

// HasRotatedUsername returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasRotatedUsername() bool {
	if o != nil && o.RotatedUsername != nil {
		return true
	}

	return false
}

// SetRotatedUsername gets a reference to the given string and assigns it to the RotatedUsername field.
func (o *RotatedSecretCreateMysql) SetRotatedUsername(v string) {
	o.RotatedUsername = &v
}

// GetRotationHour returns the RotationHour field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetRotationHour() int32 {
	if o == nil || o.RotationHour == nil {
		var ret int32
		return ret
	}
	return *o.RotationHour
}

// GetRotationHourOk returns a tuple with the RotationHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetRotationHourOk() (*int32, bool) {
	if o == nil || o.RotationHour == nil {
		return nil, false
	}
	return o.RotationHour, true
}

// HasRotationHour returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasRotationHour() bool {
	if o != nil && o.RotationHour != nil {
		return true
	}

	return false
}

// SetRotationHour gets a reference to the given int32 and assigns it to the RotationHour field.
func (o *RotatedSecretCreateMysql) SetRotationHour(v int32) {
	o.RotationHour = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetRotationInterval() string {
	if o == nil || o.RotationInterval == nil {
		var ret string
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetRotationIntervalOk() (*string, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given string and assigns it to the RotationInterval field.
func (o *RotatedSecretCreateMysql) SetRotationInterval(v string) {
	o.RotationInterval = &v
}

// GetRotatorType returns the RotatorType field value
func (o *RotatedSecretCreateMysql) GetRotatorType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RotatorType
}

// GetRotatorTypeOk returns a tuple with the RotatorType field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetRotatorTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RotatorType, true
}

// SetRotatorType sets field value
func (o *RotatedSecretCreateMysql) SetRotatorType(v string) {
	o.RotatorType = v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *RotatedSecretCreateMysql) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessDbName returns the SecureAccessDbName field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetSecureAccessDbName() string {
	if o == nil || o.SecureAccessDbName == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDbName
}

// GetSecureAccessDbNameOk returns a tuple with the SecureAccessDbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetSecureAccessDbNameOk() (*string, bool) {
	if o == nil || o.SecureAccessDbName == nil {
		return nil, false
	}
	return o.SecureAccessDbName, true
}

// HasSecureAccessDbName returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasSecureAccessDbName() bool {
	if o != nil && o.SecureAccessDbName != nil {
		return true
	}

	return false
}

// SetSecureAccessDbName gets a reference to the given string and assigns it to the SecureAccessDbName field.
func (o *RotatedSecretCreateMysql) SetSecureAccessDbName(v string) {
	o.SecureAccessDbName = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *RotatedSecretCreateMysql) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *RotatedSecretCreateMysql) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *RotatedSecretCreateMysql) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RotatedSecretCreateMysql) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value
func (o *RotatedSecretCreateMysql) GetTargetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetTargetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetName, true
}

// SetTargetName sets field value
func (o *RotatedSecretCreateMysql) SetTargetName(v string) {
	o.TargetName = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RotatedSecretCreateMysql) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *RotatedSecretCreateMysql) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RotatedSecretCreateMysql) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *RotatedSecretCreateMysql) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *RotatedSecretCreateMysql) SetUidToken(v string) {
	o.UidToken = &v
}

func (o RotatedSecretCreateMysql) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["name"] = o.Name
	}
	if o.PasswordLength != nil {
		toSerialize["password-length"] = o.PasswordLength
	}
	if o.RotateAfterDisconnect != nil {
		toSerialize["rotate-after-disconnect"] = o.RotateAfterDisconnect
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
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessDbName != nil {
		toSerialize["secure-access-db-name"] = o.SecureAccessDbName
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessHost != nil {
		toSerialize["secure-access-host"] = o.SecureAccessHost
	}
	if o.SecureAccessWeb != nil {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
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

type NullableRotatedSecretCreateMysql struct {
	value *RotatedSecretCreateMysql
	isSet bool
}

func (v NullableRotatedSecretCreateMysql) Get() *RotatedSecretCreateMysql {
	return v.value
}

func (v *NullableRotatedSecretCreateMysql) Set(val *RotatedSecretCreateMysql) {
	v.value = val
	v.isSet = true
}

func (v NullableRotatedSecretCreateMysql) IsSet() bool {
	return v.isSet
}

func (v *NullableRotatedSecretCreateMysql) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotatedSecretCreateMysql(val *RotatedSecretCreateMysql) *NullableRotatedSecretCreateMysql {
	return &NullableRotatedSecretCreateMysql{value: val, isSet: true}
}

func (v NullableRotatedSecretCreateMysql) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotatedSecretCreateMysql) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


