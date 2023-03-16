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

// GatewayCreateProducerRdp gatewayCreateProducerRdp is a command that creates rdp producer
type GatewayCreateProducerRdp struct {
	// AllowUserExtendSession
	AllowUserExtendSession *int64 `json:"allow-user-extend-session,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Allow access using externally (IdP) provided username [true/false]
	FixedUserOnly *string `json:"fixed-user-only,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// RDP Admin Name
	RdpAdminName *string `json:"rdp-admin-name,omitempty"`
	// RDP Admin password
	RdpAdminPwd *string `json:"rdp-admin-pwd,omitempty"`
	// Hostname
	RdpHostName *string `json:"rdp-host-name,omitempty"`
	// Port
	RdpHostPort *string `json:"rdp-host-port,omitempty"`
	// Groups
	RdpUserGroups *string `json:"rdp-user-groups,omitempty"`
	// Allow providing external user for a domain users
	SecureAccessAllowExternalUser *bool `json:"secure-access-allow-external-user,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers)
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	// Required when the Dynamic Secret is used for a domain user
	SecureAccessRdpDomain *string `json:"secure-access-rdp-domain,omitempty"`
	// Override the RDP Domain username
	SecureAccessRdpUser *string `json:"secure-access-rdp-user,omitempty"`
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
	// WarnBeforeUserExpiration
	WarnUserBeforeExpiration *int64 `json:"warn-user-before-expiration,omitempty"`
}

// NewGatewayCreateProducerRdp instantiates a new GatewayCreateProducerRdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerRdp(name string, ) *GatewayCreateProducerRdp {
	this := GatewayCreateProducerRdp{}
	var fixedUserOnly string = "false"
	this.FixedUserOnly = &fixedUserOnly
	var json bool = false
	this.Json = &json
	this.Name = name
	var rdpHostPort string = "22"
	this.RdpHostPort = &rdpHostPort
	var secureAccessAllowExternalUser bool = false
	this.SecureAccessAllowExternalUser = &secureAccessAllowExternalUser
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerRdpWithDefaults instantiates a new GatewayCreateProducerRdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerRdpWithDefaults() *GatewayCreateProducerRdp {
	this := GatewayCreateProducerRdp{}
	var fixedUserOnly string = "false"
	this.FixedUserOnly = &fixedUserOnly
	var json bool = false
	this.Json = &json
	var rdpHostPort string = "22"
	this.RdpHostPort = &rdpHostPort
	var secureAccessAllowExternalUser bool = false
	this.SecureAccessAllowExternalUser = &secureAccessAllowExternalUser
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetAllowUserExtendSession returns the AllowUserExtendSession field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetAllowUserExtendSession() int64 {
	if o == nil || o.AllowUserExtendSession == nil {
		var ret int64
		return ret
	}
	return *o.AllowUserExtendSession
}

// GetAllowUserExtendSessionOk returns a tuple with the AllowUserExtendSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetAllowUserExtendSessionOk() (*int64, bool) {
	if o == nil || o.AllowUserExtendSession == nil {
		return nil, false
	}
	return o.AllowUserExtendSession, true
}

// HasAllowUserExtendSession returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasAllowUserExtendSession() bool {
	if o != nil && o.AllowUserExtendSession != nil {
		return true
	}

	return false
}

// SetAllowUserExtendSession gets a reference to the given int64 and assigns it to the AllowUserExtendSession field.
func (o *GatewayCreateProducerRdp) SetAllowUserExtendSession(v int64) {
	o.AllowUserExtendSession = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerRdp) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetFixedUserOnly returns the FixedUserOnly field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetFixedUserOnly() string {
	if o == nil || o.FixedUserOnly == nil {
		var ret string
		return ret
	}
	return *o.FixedUserOnly
}

// GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetFixedUserOnlyOk() (*string, bool) {
	if o == nil || o.FixedUserOnly == nil {
		return nil, false
	}
	return o.FixedUserOnly, true
}

// HasFixedUserOnly returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasFixedUserOnly() bool {
	if o != nil && o.FixedUserOnly != nil {
		return true
	}

	return false
}

// SetFixedUserOnly gets a reference to the given string and assigns it to the FixedUserOnly field.
func (o *GatewayCreateProducerRdp) SetFixedUserOnly(v string) {
	o.FixedUserOnly = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerRdp) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerRdp) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerRdp) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerRdp) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRdpAdminName returns the RdpAdminName field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetRdpAdminName() string {
	if o == nil || o.RdpAdminName == nil {
		var ret string
		return ret
	}
	return *o.RdpAdminName
}

// GetRdpAdminNameOk returns a tuple with the RdpAdminName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetRdpAdminNameOk() (*string, bool) {
	if o == nil || o.RdpAdminName == nil {
		return nil, false
	}
	return o.RdpAdminName, true
}

// HasRdpAdminName returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasRdpAdminName() bool {
	if o != nil && o.RdpAdminName != nil {
		return true
	}

	return false
}

// SetRdpAdminName gets a reference to the given string and assigns it to the RdpAdminName field.
func (o *GatewayCreateProducerRdp) SetRdpAdminName(v string) {
	o.RdpAdminName = &v
}

// GetRdpAdminPwd returns the RdpAdminPwd field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetRdpAdminPwd() string {
	if o == nil || o.RdpAdminPwd == nil {
		var ret string
		return ret
	}
	return *o.RdpAdminPwd
}

// GetRdpAdminPwdOk returns a tuple with the RdpAdminPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetRdpAdminPwdOk() (*string, bool) {
	if o == nil || o.RdpAdminPwd == nil {
		return nil, false
	}
	return o.RdpAdminPwd, true
}

// HasRdpAdminPwd returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasRdpAdminPwd() bool {
	if o != nil && o.RdpAdminPwd != nil {
		return true
	}

	return false
}

// SetRdpAdminPwd gets a reference to the given string and assigns it to the RdpAdminPwd field.
func (o *GatewayCreateProducerRdp) SetRdpAdminPwd(v string) {
	o.RdpAdminPwd = &v
}

// GetRdpHostName returns the RdpHostName field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetRdpHostName() string {
	if o == nil || o.RdpHostName == nil {
		var ret string
		return ret
	}
	return *o.RdpHostName
}

// GetRdpHostNameOk returns a tuple with the RdpHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetRdpHostNameOk() (*string, bool) {
	if o == nil || o.RdpHostName == nil {
		return nil, false
	}
	return o.RdpHostName, true
}

// HasRdpHostName returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasRdpHostName() bool {
	if o != nil && o.RdpHostName != nil {
		return true
	}

	return false
}

// SetRdpHostName gets a reference to the given string and assigns it to the RdpHostName field.
func (o *GatewayCreateProducerRdp) SetRdpHostName(v string) {
	o.RdpHostName = &v
}

// GetRdpHostPort returns the RdpHostPort field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetRdpHostPort() string {
	if o == nil || o.RdpHostPort == nil {
		var ret string
		return ret
	}
	return *o.RdpHostPort
}

// GetRdpHostPortOk returns a tuple with the RdpHostPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetRdpHostPortOk() (*string, bool) {
	if o == nil || o.RdpHostPort == nil {
		return nil, false
	}
	return o.RdpHostPort, true
}

// HasRdpHostPort returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasRdpHostPort() bool {
	if o != nil && o.RdpHostPort != nil {
		return true
	}

	return false
}

// SetRdpHostPort gets a reference to the given string and assigns it to the RdpHostPort field.
func (o *GatewayCreateProducerRdp) SetRdpHostPort(v string) {
	o.RdpHostPort = &v
}

// GetRdpUserGroups returns the RdpUserGroups field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetRdpUserGroups() string {
	if o == nil || o.RdpUserGroups == nil {
		var ret string
		return ret
	}
	return *o.RdpUserGroups
}

// GetRdpUserGroupsOk returns a tuple with the RdpUserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetRdpUserGroupsOk() (*string, bool) {
	if o == nil || o.RdpUserGroups == nil {
		return nil, false
	}
	return o.RdpUserGroups, true
}

// HasRdpUserGroups returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasRdpUserGroups() bool {
	if o != nil && o.RdpUserGroups != nil {
		return true
	}

	return false
}

// SetRdpUserGroups gets a reference to the given string and assigns it to the RdpUserGroups field.
func (o *GatewayCreateProducerRdp) SetRdpUserGroups(v string) {
	o.RdpUserGroups = &v
}

// GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetSecureAccessAllowExternalUser() bool {
	if o == nil || o.SecureAccessAllowExternalUser == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAllowExternalUser
}

// GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetSecureAccessAllowExternalUserOk() (*bool, bool) {
	if o == nil || o.SecureAccessAllowExternalUser == nil {
		return nil, false
	}
	return o.SecureAccessAllowExternalUser, true
}

// HasSecureAccessAllowExternalUser returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasSecureAccessAllowExternalUser() bool {
	if o != nil && o.SecureAccessAllowExternalUser != nil {
		return true
	}

	return false
}

// SetSecureAccessAllowExternalUser gets a reference to the given bool and assigns it to the SecureAccessAllowExternalUser field.
func (o *GatewayCreateProducerRdp) SetSecureAccessAllowExternalUser(v bool) {
	o.SecureAccessAllowExternalUser = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayCreateProducerRdp) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *GatewayCreateProducerRdp) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetSecureAccessRdpDomain() string {
	if o == nil || o.SecureAccessRdpDomain == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessRdpDomain
}

// GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetSecureAccessRdpDomainOk() (*string, bool) {
	if o == nil || o.SecureAccessRdpDomain == nil {
		return nil, false
	}
	return o.SecureAccessRdpDomain, true
}

// HasSecureAccessRdpDomain returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasSecureAccessRdpDomain() bool {
	if o != nil && o.SecureAccessRdpDomain != nil {
		return true
	}

	return false
}

// SetSecureAccessRdpDomain gets a reference to the given string and assigns it to the SecureAccessRdpDomain field.
func (o *GatewayCreateProducerRdp) SetSecureAccessRdpDomain(v string) {
	o.SecureAccessRdpDomain = &v
}

// GetSecureAccessRdpUser returns the SecureAccessRdpUser field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetSecureAccessRdpUser() string {
	if o == nil || o.SecureAccessRdpUser == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessRdpUser
}

// GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetSecureAccessRdpUserOk() (*string, bool) {
	if o == nil || o.SecureAccessRdpUser == nil {
		return nil, false
	}
	return o.SecureAccessRdpUser, true
}

// HasSecureAccessRdpUser returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasSecureAccessRdpUser() bool {
	if o != nil && o.SecureAccessRdpUser != nil {
		return true
	}

	return false
}

// SetSecureAccessRdpUser gets a reference to the given string and assigns it to the SecureAccessRdpUser field.
func (o *GatewayCreateProducerRdp) SetSecureAccessRdpUser(v string) {
	o.SecureAccessRdpUser = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerRdp) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerRdp) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerRdp) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerRdp) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerRdp) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetWarnUserBeforeExpiration returns the WarnUserBeforeExpiration field value if set, zero value otherwise.
func (o *GatewayCreateProducerRdp) GetWarnUserBeforeExpiration() int64 {
	if o == nil || o.WarnUserBeforeExpiration == nil {
		var ret int64
		return ret
	}
	return *o.WarnUserBeforeExpiration
}

// GetWarnUserBeforeExpirationOk returns a tuple with the WarnUserBeforeExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRdp) GetWarnUserBeforeExpirationOk() (*int64, bool) {
	if o == nil || o.WarnUserBeforeExpiration == nil {
		return nil, false
	}
	return o.WarnUserBeforeExpiration, true
}

// HasWarnUserBeforeExpiration returns a boolean if a field has been set.
func (o *GatewayCreateProducerRdp) HasWarnUserBeforeExpiration() bool {
	if o != nil && o.WarnUserBeforeExpiration != nil {
		return true
	}

	return false
}

// SetWarnUserBeforeExpiration gets a reference to the given int64 and assigns it to the WarnUserBeforeExpiration field.
func (o *GatewayCreateProducerRdp) SetWarnUserBeforeExpiration(v int64) {
	o.WarnUserBeforeExpiration = &v
}

func (o GatewayCreateProducerRdp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowUserExtendSession != nil {
		toSerialize["allow-user-extend-session"] = o.AllowUserExtendSession
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.FixedUserOnly != nil {
		toSerialize["fixed-user-only"] = o.FixedUserOnly
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.RdpAdminName != nil {
		toSerialize["rdp-admin-name"] = o.RdpAdminName
	}
	if o.RdpAdminPwd != nil {
		toSerialize["rdp-admin-pwd"] = o.RdpAdminPwd
	}
	if o.RdpHostName != nil {
		toSerialize["rdp-host-name"] = o.RdpHostName
	}
	if o.RdpHostPort != nil {
		toSerialize["rdp-host-port"] = o.RdpHostPort
	}
	if o.RdpUserGroups != nil {
		toSerialize["rdp-user-groups"] = o.RdpUserGroups
	}
	if o.SecureAccessAllowExternalUser != nil {
		toSerialize["secure-access-allow-external-user"] = o.SecureAccessAllowExternalUser
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessHost != nil {
		toSerialize["secure-access-host"] = o.SecureAccessHost
	}
	if o.SecureAccessRdpDomain != nil {
		toSerialize["secure-access-rdp-domain"] = o.SecureAccessRdpDomain
	}
	if o.SecureAccessRdpUser != nil {
		toSerialize["secure-access-rdp-user"] = o.SecureAccessRdpUser
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
	if o.WarnUserBeforeExpiration != nil {
		toSerialize["warn-user-before-expiration"] = o.WarnUserBeforeExpiration
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerRdp struct {
	value *GatewayCreateProducerRdp
	isSet bool
}

func (v NullableGatewayCreateProducerRdp) Get() *GatewayCreateProducerRdp {
	return v.value
}

func (v *NullableGatewayCreateProducerRdp) Set(val *GatewayCreateProducerRdp) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerRdp) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerRdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerRdp(val *GatewayCreateProducerRdp) *NullableGatewayCreateProducerRdp {
	return &NullableGatewayCreateProducerRdp{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerRdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerRdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


