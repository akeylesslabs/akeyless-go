/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// checks if the GatewayUpdateProducerRabbitMQ type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayUpdateProducerRabbitMQ{}

// GatewayUpdateProducerRabbitMQ gatewayUpdateProducerRabbitMQ is a command that updates rabbitmq producer
type GatewayUpdateProducerRabbitMQ struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
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
	Tags []string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayUpdateProducerRabbitMQ instantiates a new GatewayUpdateProducerRabbitMQ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerRabbitMQ(name string) *GatewayUpdateProducerRabbitMQ {
	this := GatewayUpdateProducerRabbitMQ{}
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

// NewGatewayUpdateProducerRabbitMQWithDefaults instantiates a new GatewayUpdateProducerRabbitMQ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerRabbitMQWithDefaults() *GatewayUpdateProducerRabbitMQ {
	this := GatewayUpdateProducerRabbitMQ{}
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
func (o *GatewayUpdateProducerRabbitMQ) GetDeleteProtection() string {
	if o == nil || IsNil(o.DeleteProtection) {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || IsNil(o.DeleteProtection) {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasDeleteProtection() bool {
	if o != nil && !IsNil(o.DeleteProtection) {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerRabbitMQ) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerRabbitMQ) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerRabbitMQ) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerRabbitMQ) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetNewName() string {
	if o == nil || IsNil(o.NewName) {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetNewNameOk() (*string, bool) {
	if o == nil || IsNil(o.NewName) {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasNewName() bool {
	if o != nil && !IsNil(o.NewName) {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerRabbitMQ) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetProducerEncryptionKeyName() string {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasProducerEncryptionKeyName() bool {
	if o != nil && !IsNil(o.ProducerEncryptionKeyName) {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerRabbitMQ) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRabbitmqAdminPwd returns the RabbitmqAdminPwd field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqAdminPwd() string {
	if o == nil || IsNil(o.RabbitmqAdminPwd) {
		var ret string
		return ret
	}
	return *o.RabbitmqAdminPwd
}

// GetRabbitmqAdminPwdOk returns a tuple with the RabbitmqAdminPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqAdminPwdOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqAdminPwd) {
		return nil, false
	}
	return o.RabbitmqAdminPwd, true
}

// HasRabbitmqAdminPwd returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasRabbitmqAdminPwd() bool {
	if o != nil && !IsNil(o.RabbitmqAdminPwd) {
		return true
	}

	return false
}

// SetRabbitmqAdminPwd gets a reference to the given string and assigns it to the RabbitmqAdminPwd field.
func (o *GatewayUpdateProducerRabbitMQ) SetRabbitmqAdminPwd(v string) {
	o.RabbitmqAdminPwd = &v
}

// GetRabbitmqAdminUser returns the RabbitmqAdminUser field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqAdminUser() string {
	if o == nil || IsNil(o.RabbitmqAdminUser) {
		var ret string
		return ret
	}
	return *o.RabbitmqAdminUser
}

// GetRabbitmqAdminUserOk returns a tuple with the RabbitmqAdminUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqAdminUserOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqAdminUser) {
		return nil, false
	}
	return o.RabbitmqAdminUser, true
}

// HasRabbitmqAdminUser returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasRabbitmqAdminUser() bool {
	if o != nil && !IsNil(o.RabbitmqAdminUser) {
		return true
	}

	return false
}

// SetRabbitmqAdminUser gets a reference to the given string and assigns it to the RabbitmqAdminUser field.
func (o *GatewayUpdateProducerRabbitMQ) SetRabbitmqAdminUser(v string) {
	o.RabbitmqAdminUser = &v
}

// GetRabbitmqServerUri returns the RabbitmqServerUri field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqServerUri() string {
	if o == nil || IsNil(o.RabbitmqServerUri) {
		var ret string
		return ret
	}
	return *o.RabbitmqServerUri
}

// GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqServerUriOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqServerUri) {
		return nil, false
	}
	return o.RabbitmqServerUri, true
}

// HasRabbitmqServerUri returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasRabbitmqServerUri() bool {
	if o != nil && !IsNil(o.RabbitmqServerUri) {
		return true
	}

	return false
}

// SetRabbitmqServerUri gets a reference to the given string and assigns it to the RabbitmqServerUri field.
func (o *GatewayUpdateProducerRabbitMQ) SetRabbitmqServerUri(v string) {
	o.RabbitmqServerUri = &v
}

// GetRabbitmqUserConfPermission returns the RabbitmqUserConfPermission field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserConfPermission() string {
	if o == nil || IsNil(o.RabbitmqUserConfPermission) {
		var ret string
		return ret
	}
	return *o.RabbitmqUserConfPermission
}

// GetRabbitmqUserConfPermissionOk returns a tuple with the RabbitmqUserConfPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserConfPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqUserConfPermission) {
		return nil, false
	}
	return o.RabbitmqUserConfPermission, true
}

// HasRabbitmqUserConfPermission returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasRabbitmqUserConfPermission() bool {
	if o != nil && !IsNil(o.RabbitmqUserConfPermission) {
		return true
	}

	return false
}

// SetRabbitmqUserConfPermission gets a reference to the given string and assigns it to the RabbitmqUserConfPermission field.
func (o *GatewayUpdateProducerRabbitMQ) SetRabbitmqUserConfPermission(v string) {
	o.RabbitmqUserConfPermission = &v
}

// GetRabbitmqUserReadPermission returns the RabbitmqUserReadPermission field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserReadPermission() string {
	if o == nil || IsNil(o.RabbitmqUserReadPermission) {
		var ret string
		return ret
	}
	return *o.RabbitmqUserReadPermission
}

// GetRabbitmqUserReadPermissionOk returns a tuple with the RabbitmqUserReadPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserReadPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqUserReadPermission) {
		return nil, false
	}
	return o.RabbitmqUserReadPermission, true
}

// HasRabbitmqUserReadPermission returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasRabbitmqUserReadPermission() bool {
	if o != nil && !IsNil(o.RabbitmqUserReadPermission) {
		return true
	}

	return false
}

// SetRabbitmqUserReadPermission gets a reference to the given string and assigns it to the RabbitmqUserReadPermission field.
func (o *GatewayUpdateProducerRabbitMQ) SetRabbitmqUserReadPermission(v string) {
	o.RabbitmqUserReadPermission = &v
}

// GetRabbitmqUserTags returns the RabbitmqUserTags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserTags() string {
	if o == nil || IsNil(o.RabbitmqUserTags) {
		var ret string
		return ret
	}
	return *o.RabbitmqUserTags
}

// GetRabbitmqUserTagsOk returns a tuple with the RabbitmqUserTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserTagsOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqUserTags) {
		return nil, false
	}
	return o.RabbitmqUserTags, true
}

// HasRabbitmqUserTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasRabbitmqUserTags() bool {
	if o != nil && !IsNil(o.RabbitmqUserTags) {
		return true
	}

	return false
}

// SetRabbitmqUserTags gets a reference to the given string and assigns it to the RabbitmqUserTags field.
func (o *GatewayUpdateProducerRabbitMQ) SetRabbitmqUserTags(v string) {
	o.RabbitmqUserTags = &v
}

// GetRabbitmqUserVhost returns the RabbitmqUserVhost field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserVhost() string {
	if o == nil || IsNil(o.RabbitmqUserVhost) {
		var ret string
		return ret
	}
	return *o.RabbitmqUserVhost
}

// GetRabbitmqUserVhostOk returns a tuple with the RabbitmqUserVhost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserVhostOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqUserVhost) {
		return nil, false
	}
	return o.RabbitmqUserVhost, true
}

// HasRabbitmqUserVhost returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasRabbitmqUserVhost() bool {
	if o != nil && !IsNil(o.RabbitmqUserVhost) {
		return true
	}

	return false
}

// SetRabbitmqUserVhost gets a reference to the given string and assigns it to the RabbitmqUserVhost field.
func (o *GatewayUpdateProducerRabbitMQ) SetRabbitmqUserVhost(v string) {
	o.RabbitmqUserVhost = &v
}

// GetRabbitmqUserWritePermission returns the RabbitmqUserWritePermission field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserWritePermission() string {
	if o == nil || IsNil(o.RabbitmqUserWritePermission) {
		var ret string
		return ret
	}
	return *o.RabbitmqUserWritePermission
}

// GetRabbitmqUserWritePermissionOk returns a tuple with the RabbitmqUserWritePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetRabbitmqUserWritePermissionOk() (*string, bool) {
	if o == nil || IsNil(o.RabbitmqUserWritePermission) {
		return nil, false
	}
	return o.RabbitmqUserWritePermission, true
}

// HasRabbitmqUserWritePermission returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasRabbitmqUserWritePermission() bool {
	if o != nil && !IsNil(o.RabbitmqUserWritePermission) {
		return true
	}

	return false
}

// SetRabbitmqUserWritePermission gets a reference to the given string and assigns it to the RabbitmqUserWritePermission field.
func (o *GatewayUpdateProducerRabbitMQ) SetRabbitmqUserWritePermission(v string) {
	o.RabbitmqUserWritePermission = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessEnable() string {
	if o == nil || IsNil(o.SecureAccessEnable) {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || IsNil(o.SecureAccessEnable) {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasSecureAccessEnable() bool {
	if o != nil && !IsNil(o.SecureAccessEnable) {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayUpdateProducerRabbitMQ) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessUrl returns the SecureAccessUrl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessUrl() string {
	if o == nil || IsNil(o.SecureAccessUrl) {
		var ret string
		return ret
	}
	return *o.SecureAccessUrl
}

// GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SecureAccessUrl) {
		return nil, false
	}
	return o.SecureAccessUrl, true
}

// HasSecureAccessUrl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasSecureAccessUrl() bool {
	if o != nil && !IsNil(o.SecureAccessUrl) {
		return true
	}

	return false
}

// SetSecureAccessUrl gets a reference to the given string and assigns it to the SecureAccessUrl field.
func (o *GatewayUpdateProducerRabbitMQ) SetSecureAccessUrl(v string) {
	o.SecureAccessUrl = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessWeb() bool {
	if o == nil || IsNil(o.SecureAccessWeb) {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureAccessWeb) {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasSecureAccessWeb() bool {
	if o != nil && !IsNil(o.SecureAccessWeb) {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayUpdateProducerRabbitMQ) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessWebBrowsing() bool {
	if o == nil || IsNil(o.SecureAccessWebBrowsing) {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebBrowsing
}

// GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessWebBrowsingOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureAccessWebBrowsing) {
		return nil, false
	}
	return o.SecureAccessWebBrowsing, true
}

// HasSecureAccessWebBrowsing returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasSecureAccessWebBrowsing() bool {
	if o != nil && !IsNil(o.SecureAccessWebBrowsing) {
		return true
	}

	return false
}

// SetSecureAccessWebBrowsing gets a reference to the given bool and assigns it to the SecureAccessWebBrowsing field.
func (o *GatewayUpdateProducerRabbitMQ) SetSecureAccessWebBrowsing(v bool) {
	o.SecureAccessWebBrowsing = &v
}

// GetSecureAccessWebProxy returns the SecureAccessWebProxy field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessWebProxy() bool {
	if o == nil || IsNil(o.SecureAccessWebProxy) {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebProxy
}

// GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetSecureAccessWebProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureAccessWebProxy) {
		return nil, false
	}
	return o.SecureAccessWebProxy, true
}

// HasSecureAccessWebProxy returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasSecureAccessWebProxy() bool {
	if o != nil && !IsNil(o.SecureAccessWebProxy) {
		return true
	}

	return false
}

// SetSecureAccessWebProxy gets a reference to the given bool and assigns it to the SecureAccessWebProxy field.
func (o *GatewayUpdateProducerRabbitMQ) SetSecureAccessWebProxy(v bool) {
	o.SecureAccessWebProxy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerRabbitMQ) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerRabbitMQ) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerRabbitMQ) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerRabbitMQ) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerRabbitMQ) GetUserTtl() string {
	if o == nil || IsNil(o.UserTtl) {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerRabbitMQ) GetUserTtlOk() (*string, bool) {
	if o == nil || IsNil(o.UserTtl) {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerRabbitMQ) HasUserTtl() bool {
	if o != nil && !IsNil(o.UserTtl) {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerRabbitMQ) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerRabbitMQ) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayUpdateProducerRabbitMQ) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteProtection) {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NewName) {
		toSerialize["new-name"] = o.NewName
	}
	if !IsNil(o.ProducerEncryptionKeyName) {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if !IsNil(o.RabbitmqAdminPwd) {
		toSerialize["rabbitmq-admin-pwd"] = o.RabbitmqAdminPwd
	}
	if !IsNil(o.RabbitmqAdminUser) {
		toSerialize["rabbitmq-admin-user"] = o.RabbitmqAdminUser
	}
	if !IsNil(o.RabbitmqServerUri) {
		toSerialize["rabbitmq-server-uri"] = o.RabbitmqServerUri
	}
	if !IsNil(o.RabbitmqUserConfPermission) {
		toSerialize["rabbitmq-user-conf-permission"] = o.RabbitmqUserConfPermission
	}
	if !IsNil(o.RabbitmqUserReadPermission) {
		toSerialize["rabbitmq-user-read-permission"] = o.RabbitmqUserReadPermission
	}
	if !IsNil(o.RabbitmqUserTags) {
		toSerialize["rabbitmq-user-tags"] = o.RabbitmqUserTags
	}
	if !IsNil(o.RabbitmqUserVhost) {
		toSerialize["rabbitmq-user-vhost"] = o.RabbitmqUserVhost
	}
	if !IsNil(o.RabbitmqUserWritePermission) {
		toSerialize["rabbitmq-user-write-permission"] = o.RabbitmqUserWritePermission
	}
	if !IsNil(o.SecureAccessEnable) {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if !IsNil(o.SecureAccessUrl) {
		toSerialize["secure-access-url"] = o.SecureAccessUrl
	}
	if !IsNil(o.SecureAccessWeb) {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
	}
	if !IsNil(o.SecureAccessWebBrowsing) {
		toSerialize["secure-access-web-browsing"] = o.SecureAccessWebBrowsing
	}
	if !IsNil(o.SecureAccessWebProxy) {
		toSerialize["secure-access-web-proxy"] = o.SecureAccessWebProxy
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TargetName) {
		toSerialize["target-name"] = o.TargetName
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UidToken) {
		toSerialize["uid-token"] = o.UidToken
	}
	if !IsNil(o.UserTtl) {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return toSerialize, nil
}

type NullableGatewayUpdateProducerRabbitMQ struct {
	value *GatewayUpdateProducerRabbitMQ
	isSet bool
}

func (v NullableGatewayUpdateProducerRabbitMQ) Get() *GatewayUpdateProducerRabbitMQ {
	return v.value
}

func (v *NullableGatewayUpdateProducerRabbitMQ) Set(val *GatewayUpdateProducerRabbitMQ) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerRabbitMQ) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerRabbitMQ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerRabbitMQ(val *GatewayUpdateProducerRabbitMQ) *NullableGatewayUpdateProducerRabbitMQ {
	return &NullableGatewayUpdateProducerRabbitMQ{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerRabbitMQ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerRabbitMQ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


