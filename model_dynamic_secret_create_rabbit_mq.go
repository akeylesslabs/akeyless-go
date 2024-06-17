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

// DynamicSecretCreateRabbitMq dynamicSecretCreateRabbitMq is a command that creates rabbitmq dynamic secret
type DynamicSecretCreateRabbitMq struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// RabbitMQ Admin password
	RabbitmqAdminPwd *string `json:"rabbitmq-admin-pwd,omitempty"`
	// RabbitMQ Admin User
	RabbitmqAdminUser *string `json:"rabbitmq-admin-user,omitempty"`
	// Server URI
	RabbitmqServerUri *string `json:"rabbitmq-server-uri,omitempty"`
	// User configuration permission
	RabbitmqUserConfPermission *string `json:"rabbitmq-user-conf-permission,omitempty"`
	// User read permission
	RabbitmqUserReadPermission *string `json:"rabbitmq-user-read-permission,omitempty"`
	// User Tags
	RabbitmqUserTags *string `json:"rabbitmq-user-tags,omitempty"`
	// User Virtual Host
	RabbitmqUserVhost *string `json:"rabbitmq-user-vhost,omitempty"`
	// User write permission
	RabbitmqUserWritePermission *string `json:"rabbitmq-user-write-permission,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Destination URL to inject secrets
	SecureAccessUrl *string `json:"secure-access-url,omitempty"`
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

// NewDynamicSecretCreateRabbitMq instantiates a new DynamicSecretCreateRabbitMq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretCreateRabbitMq(name string, ) *DynamicSecretCreateRabbitMq {
	this := DynamicSecretCreateRabbitMq{}
	var json bool = false
	this.Json = &json
	this.Name = name
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

// NewDynamicSecretCreateRabbitMqWithDefaults instantiates a new DynamicSecretCreateRabbitMq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretCreateRabbitMqWithDefaults() *DynamicSecretCreateRabbitMq {
	this := DynamicSecretCreateRabbitMq{}
	var json bool = false
	this.Json = &json
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

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretCreateRabbitMq) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretCreateRabbitMq) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretCreateRabbitMq) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretCreateRabbitMq) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretCreateRabbitMq) SetName(v string) {
	o.Name = v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *DynamicSecretCreateRabbitMq) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretCreateRabbitMq) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRabbitmqAdminPwd returns the RabbitmqAdminPwd field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqAdminPwd() string {
	if o == nil || o.RabbitmqAdminPwd == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqAdminPwd
}

// GetRabbitmqAdminPwdOk returns a tuple with the RabbitmqAdminPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqAdminPwdOk() (*string, bool) {
	if o == nil || o.RabbitmqAdminPwd == nil {
		return nil, false
	}
	return o.RabbitmqAdminPwd, true
}

// HasRabbitmqAdminPwd returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasRabbitmqAdminPwd() bool {
	if o != nil && o.RabbitmqAdminPwd != nil {
		return true
	}

	return false
}

// SetRabbitmqAdminPwd gets a reference to the given string and assigns it to the RabbitmqAdminPwd field.
func (o *DynamicSecretCreateRabbitMq) SetRabbitmqAdminPwd(v string) {
	o.RabbitmqAdminPwd = &v
}

// GetRabbitmqAdminUser returns the RabbitmqAdminUser field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqAdminUser() string {
	if o == nil || o.RabbitmqAdminUser == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqAdminUser
}

// GetRabbitmqAdminUserOk returns a tuple with the RabbitmqAdminUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqAdminUserOk() (*string, bool) {
	if o == nil || o.RabbitmqAdminUser == nil {
		return nil, false
	}
	return o.RabbitmqAdminUser, true
}

// HasRabbitmqAdminUser returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasRabbitmqAdminUser() bool {
	if o != nil && o.RabbitmqAdminUser != nil {
		return true
	}

	return false
}

// SetRabbitmqAdminUser gets a reference to the given string and assigns it to the RabbitmqAdminUser field.
func (o *DynamicSecretCreateRabbitMq) SetRabbitmqAdminUser(v string) {
	o.RabbitmqAdminUser = &v
}

// GetRabbitmqServerUri returns the RabbitmqServerUri field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqServerUri() string {
	if o == nil || o.RabbitmqServerUri == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqServerUri
}

// GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqServerUriOk() (*string, bool) {
	if o == nil || o.RabbitmqServerUri == nil {
		return nil, false
	}
	return o.RabbitmqServerUri, true
}

// HasRabbitmqServerUri returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasRabbitmqServerUri() bool {
	if o != nil && o.RabbitmqServerUri != nil {
		return true
	}

	return false
}

// SetRabbitmqServerUri gets a reference to the given string and assigns it to the RabbitmqServerUri field.
func (o *DynamicSecretCreateRabbitMq) SetRabbitmqServerUri(v string) {
	o.RabbitmqServerUri = &v
}

// GetRabbitmqUserConfPermission returns the RabbitmqUserConfPermission field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserConfPermission() string {
	if o == nil || o.RabbitmqUserConfPermission == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqUserConfPermission
}

// GetRabbitmqUserConfPermissionOk returns a tuple with the RabbitmqUserConfPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserConfPermissionOk() (*string, bool) {
	if o == nil || o.RabbitmqUserConfPermission == nil {
		return nil, false
	}
	return o.RabbitmqUserConfPermission, true
}

// HasRabbitmqUserConfPermission returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserConfPermission() bool {
	if o != nil && o.RabbitmqUserConfPermission != nil {
		return true
	}

	return false
}

// SetRabbitmqUserConfPermission gets a reference to the given string and assigns it to the RabbitmqUserConfPermission field.
func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserConfPermission(v string) {
	o.RabbitmqUserConfPermission = &v
}

// GetRabbitmqUserReadPermission returns the RabbitmqUserReadPermission field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserReadPermission() string {
	if o == nil || o.RabbitmqUserReadPermission == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqUserReadPermission
}

// GetRabbitmqUserReadPermissionOk returns a tuple with the RabbitmqUserReadPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserReadPermissionOk() (*string, bool) {
	if o == nil || o.RabbitmqUserReadPermission == nil {
		return nil, false
	}
	return o.RabbitmqUserReadPermission, true
}

// HasRabbitmqUserReadPermission returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserReadPermission() bool {
	if o != nil && o.RabbitmqUserReadPermission != nil {
		return true
	}

	return false
}

// SetRabbitmqUserReadPermission gets a reference to the given string and assigns it to the RabbitmqUserReadPermission field.
func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserReadPermission(v string) {
	o.RabbitmqUserReadPermission = &v
}

// GetRabbitmqUserTags returns the RabbitmqUserTags field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserTags() string {
	if o == nil || o.RabbitmqUserTags == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqUserTags
}

// GetRabbitmqUserTagsOk returns a tuple with the RabbitmqUserTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserTagsOk() (*string, bool) {
	if o == nil || o.RabbitmqUserTags == nil {
		return nil, false
	}
	return o.RabbitmqUserTags, true
}

// HasRabbitmqUserTags returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserTags() bool {
	if o != nil && o.RabbitmqUserTags != nil {
		return true
	}

	return false
}

// SetRabbitmqUserTags gets a reference to the given string and assigns it to the RabbitmqUserTags field.
func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserTags(v string) {
	o.RabbitmqUserTags = &v
}

// GetRabbitmqUserVhost returns the RabbitmqUserVhost field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserVhost() string {
	if o == nil || o.RabbitmqUserVhost == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqUserVhost
}

// GetRabbitmqUserVhostOk returns a tuple with the RabbitmqUserVhost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserVhostOk() (*string, bool) {
	if o == nil || o.RabbitmqUserVhost == nil {
		return nil, false
	}
	return o.RabbitmqUserVhost, true
}

// HasRabbitmqUserVhost returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserVhost() bool {
	if o != nil && o.RabbitmqUserVhost != nil {
		return true
	}

	return false
}

// SetRabbitmqUserVhost gets a reference to the given string and assigns it to the RabbitmqUserVhost field.
func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserVhost(v string) {
	o.RabbitmqUserVhost = &v
}

// GetRabbitmqUserWritePermission returns the RabbitmqUserWritePermission field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserWritePermission() string {
	if o == nil || o.RabbitmqUserWritePermission == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqUserWritePermission
}

// GetRabbitmqUserWritePermissionOk returns a tuple with the RabbitmqUserWritePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetRabbitmqUserWritePermissionOk() (*string, bool) {
	if o == nil || o.RabbitmqUserWritePermission == nil {
		return nil, false
	}
	return o.RabbitmqUserWritePermission, true
}

// HasRabbitmqUserWritePermission returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasRabbitmqUserWritePermission() bool {
	if o != nil && o.RabbitmqUserWritePermission != nil {
		return true
	}

	return false
}

// SetRabbitmqUserWritePermission gets a reference to the given string and assigns it to the RabbitmqUserWritePermission field.
func (o *DynamicSecretCreateRabbitMq) SetRabbitmqUserWritePermission(v string) {
	o.RabbitmqUserWritePermission = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DynamicSecretCreateRabbitMq) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessUrl returns the SecureAccessUrl field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessUrl() string {
	if o == nil || o.SecureAccessUrl == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessUrl
}

// GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessUrlOk() (*string, bool) {
	if o == nil || o.SecureAccessUrl == nil {
		return nil, false
	}
	return o.SecureAccessUrl, true
}

// HasSecureAccessUrl returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasSecureAccessUrl() bool {
	if o != nil && o.SecureAccessUrl != nil {
		return true
	}

	return false
}

// SetSecureAccessUrl gets a reference to the given string and assigns it to the SecureAccessUrl field.
func (o *DynamicSecretCreateRabbitMq) SetSecureAccessUrl(v string) {
	o.SecureAccessUrl = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *DynamicSecretCreateRabbitMq) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebBrowsing() bool {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebBrowsing
}

// GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebBrowsingOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		return nil, false
	}
	return o.SecureAccessWebBrowsing, true
}

// HasSecureAccessWebBrowsing returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasSecureAccessWebBrowsing() bool {
	if o != nil && o.SecureAccessWebBrowsing != nil {
		return true
	}

	return false
}

// SetSecureAccessWebBrowsing gets a reference to the given bool and assigns it to the SecureAccessWebBrowsing field.
func (o *DynamicSecretCreateRabbitMq) SetSecureAccessWebBrowsing(v bool) {
	o.SecureAccessWebBrowsing = &v
}

// GetSecureAccessWebProxy returns the SecureAccessWebProxy field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebProxy() bool {
	if o == nil || o.SecureAccessWebProxy == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebProxy
}

// GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetSecureAccessWebProxyOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebProxy == nil {
		return nil, false
	}
	return o.SecureAccessWebProxy, true
}

// HasSecureAccessWebProxy returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasSecureAccessWebProxy() bool {
	if o != nil && o.SecureAccessWebProxy != nil {
		return true
	}

	return false
}

// SetSecureAccessWebProxy gets a reference to the given bool and assigns it to the SecureAccessWebProxy field.
func (o *DynamicSecretCreateRabbitMq) SetSecureAccessWebProxy(v bool) {
	o.SecureAccessWebProxy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretCreateRabbitMq) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretCreateRabbitMq) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretCreateRabbitMq) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretCreateRabbitMq) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretCreateRabbitMq) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRabbitMq) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretCreateRabbitMq) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretCreateRabbitMq) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretCreateRabbitMq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.PasswordLength != nil {
		toSerialize["password-length"] = o.PasswordLength
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.RabbitmqAdminPwd != nil {
		toSerialize["rabbitmq-admin-pwd"] = o.RabbitmqAdminPwd
	}
	if o.RabbitmqAdminUser != nil {
		toSerialize["rabbitmq-admin-user"] = o.RabbitmqAdminUser
	}
	if o.RabbitmqServerUri != nil {
		toSerialize["rabbitmq-server-uri"] = o.RabbitmqServerUri
	}
	if o.RabbitmqUserConfPermission != nil {
		toSerialize["rabbitmq-user-conf-permission"] = o.RabbitmqUserConfPermission
	}
	if o.RabbitmqUserReadPermission != nil {
		toSerialize["rabbitmq-user-read-permission"] = o.RabbitmqUserReadPermission
	}
	if o.RabbitmqUserTags != nil {
		toSerialize["rabbitmq-user-tags"] = o.RabbitmqUserTags
	}
	if o.RabbitmqUserVhost != nil {
		toSerialize["rabbitmq-user-vhost"] = o.RabbitmqUserVhost
	}
	if o.RabbitmqUserWritePermission != nil {
		toSerialize["rabbitmq-user-write-permission"] = o.RabbitmqUserWritePermission
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessUrl != nil {
		toSerialize["secure-access-url"] = o.SecureAccessUrl
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

type NullableDynamicSecretCreateRabbitMq struct {
	value *DynamicSecretCreateRabbitMq
	isSet bool
}

func (v NullableDynamicSecretCreateRabbitMq) Get() *DynamicSecretCreateRabbitMq {
	return v.value
}

func (v *NullableDynamicSecretCreateRabbitMq) Set(val *DynamicSecretCreateRabbitMq) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretCreateRabbitMq) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretCreateRabbitMq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretCreateRabbitMq(val *DynamicSecretCreateRabbitMq) *NullableDynamicSecretCreateRabbitMq {
	return &NullableDynamicSecretCreateRabbitMq{value: val, isSet: true}
}

func (v NullableDynamicSecretCreateRabbitMq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretCreateRabbitMq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

