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

// DsUpdateRdp dsUpdateRdp is a command that updates rdp dynamic secret
type DsUpdateRdp struct {
	// AllowUserExtendSession
	AllowUserExtendSession *int64 `json:"allow-user-extend-session,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Allow access using externally (IdP) provided username [true/false]
	FixedUserOnly *string `json:"fixed-user-only,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret name
	NewName *string `json:"new-name,omitempty"`
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
	// RD Gateway server
	SecureAccessRdGatewayServer *string `json:"secure-access-rd-gateway-server,omitempty"`
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

// NewDsUpdateRdp instantiates a new DsUpdateRdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDsUpdateRdp(name string, ) *DsUpdateRdp {
	this := DsUpdateRdp{}
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

// NewDsUpdateRdpWithDefaults instantiates a new DsUpdateRdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDsUpdateRdpWithDefaults() *DsUpdateRdp {
	this := DsUpdateRdp{}
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
func (o *DsUpdateRdp) GetAllowUserExtendSession() int64 {
	if o == nil || o.AllowUserExtendSession == nil {
		var ret int64
		return ret
	}
	return *o.AllowUserExtendSession
}

// GetAllowUserExtendSessionOk returns a tuple with the AllowUserExtendSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetAllowUserExtendSessionOk() (*int64, bool) {
	if o == nil || o.AllowUserExtendSession == nil {
		return nil, false
	}
	return o.AllowUserExtendSession, true
}

// HasAllowUserExtendSession returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasAllowUserExtendSession() bool {
	if o != nil && o.AllowUserExtendSession != nil {
		return true
	}

	return false
}

// SetAllowUserExtendSession gets a reference to the given int64 and assigns it to the AllowUserExtendSession field.
func (o *DsUpdateRdp) SetAllowUserExtendSession(v int64) {
	o.AllowUserExtendSession = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DsUpdateRdp) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetFixedUserOnly returns the FixedUserOnly field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetFixedUserOnly() string {
	if o == nil || o.FixedUserOnly == nil {
		var ret string
		return ret
	}
	return *o.FixedUserOnly
}

// GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetFixedUserOnlyOk() (*string, bool) {
	if o == nil || o.FixedUserOnly == nil {
		return nil, false
	}
	return o.FixedUserOnly, true
}

// HasFixedUserOnly returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasFixedUserOnly() bool {
	if o != nil && o.FixedUserOnly != nil {
		return true
	}

	return false
}

// SetFixedUserOnly gets a reference to the given string and assigns it to the FixedUserOnly field.
func (o *DsUpdateRdp) SetFixedUserOnly(v string) {
	o.FixedUserOnly = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DsUpdateRdp) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DsUpdateRdp) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DsUpdateRdp) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DsUpdateRdp) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DsUpdateRdp) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRdpAdminName returns the RdpAdminName field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetRdpAdminName() string {
	if o == nil || o.RdpAdminName == nil {
		var ret string
		return ret
	}
	return *o.RdpAdminName
}

// GetRdpAdminNameOk returns a tuple with the RdpAdminName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetRdpAdminNameOk() (*string, bool) {
	if o == nil || o.RdpAdminName == nil {
		return nil, false
	}
	return o.RdpAdminName, true
}

// HasRdpAdminName returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasRdpAdminName() bool {
	if o != nil && o.RdpAdminName != nil {
		return true
	}

	return false
}

// SetRdpAdminName gets a reference to the given string and assigns it to the RdpAdminName field.
func (o *DsUpdateRdp) SetRdpAdminName(v string) {
	o.RdpAdminName = &v
}

// GetRdpAdminPwd returns the RdpAdminPwd field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetRdpAdminPwd() string {
	if o == nil || o.RdpAdminPwd == nil {
		var ret string
		return ret
	}
	return *o.RdpAdminPwd
}

// GetRdpAdminPwdOk returns a tuple with the RdpAdminPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetRdpAdminPwdOk() (*string, bool) {
	if o == nil || o.RdpAdminPwd == nil {
		return nil, false
	}
	return o.RdpAdminPwd, true
}

// HasRdpAdminPwd returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasRdpAdminPwd() bool {
	if o != nil && o.RdpAdminPwd != nil {
		return true
	}

	return false
}

// SetRdpAdminPwd gets a reference to the given string and assigns it to the RdpAdminPwd field.
func (o *DsUpdateRdp) SetRdpAdminPwd(v string) {
	o.RdpAdminPwd = &v
}

// GetRdpHostName returns the RdpHostName field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetRdpHostName() string {
	if o == nil || o.RdpHostName == nil {
		var ret string
		return ret
	}
	return *o.RdpHostName
}

// GetRdpHostNameOk returns a tuple with the RdpHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetRdpHostNameOk() (*string, bool) {
	if o == nil || o.RdpHostName == nil {
		return nil, false
	}
	return o.RdpHostName, true
}

// HasRdpHostName returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasRdpHostName() bool {
	if o != nil && o.RdpHostName != nil {
		return true
	}

	return false
}

// SetRdpHostName gets a reference to the given string and assigns it to the RdpHostName field.
func (o *DsUpdateRdp) SetRdpHostName(v string) {
	o.RdpHostName = &v
}

// GetRdpHostPort returns the RdpHostPort field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetRdpHostPort() string {
	if o == nil || o.RdpHostPort == nil {
		var ret string
		return ret
	}
	return *o.RdpHostPort
}

// GetRdpHostPortOk returns a tuple with the RdpHostPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetRdpHostPortOk() (*string, bool) {
	if o == nil || o.RdpHostPort == nil {
		return nil, false
	}
	return o.RdpHostPort, true
}

// HasRdpHostPort returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasRdpHostPort() bool {
	if o != nil && o.RdpHostPort != nil {
		return true
	}

	return false
}

// SetRdpHostPort gets a reference to the given string and assigns it to the RdpHostPort field.
func (o *DsUpdateRdp) SetRdpHostPort(v string) {
	o.RdpHostPort = &v
}

// GetRdpUserGroups returns the RdpUserGroups field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetRdpUserGroups() string {
	if o == nil || o.RdpUserGroups == nil {
		var ret string
		return ret
	}
	return *o.RdpUserGroups
}

// GetRdpUserGroupsOk returns a tuple with the RdpUserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetRdpUserGroupsOk() (*string, bool) {
	if o == nil || o.RdpUserGroups == nil {
		return nil, false
	}
	return o.RdpUserGroups, true
}

// HasRdpUserGroups returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasRdpUserGroups() bool {
	if o != nil && o.RdpUserGroups != nil {
		return true
	}

	return false
}

// SetRdpUserGroups gets a reference to the given string and assigns it to the RdpUserGroups field.
func (o *DsUpdateRdp) SetRdpUserGroups(v string) {
	o.RdpUserGroups = &v
}

// GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetSecureAccessAllowExternalUser() bool {
	if o == nil || o.SecureAccessAllowExternalUser == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessAllowExternalUser
}

// GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetSecureAccessAllowExternalUserOk() (*bool, bool) {
	if o == nil || o.SecureAccessAllowExternalUser == nil {
		return nil, false
	}
	return o.SecureAccessAllowExternalUser, true
}

// HasSecureAccessAllowExternalUser returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasSecureAccessAllowExternalUser() bool {
	if o != nil && o.SecureAccessAllowExternalUser != nil {
		return true
	}

	return false
}

// SetSecureAccessAllowExternalUser gets a reference to the given bool and assigns it to the SecureAccessAllowExternalUser field.
func (o *DsUpdateRdp) SetSecureAccessAllowExternalUser(v bool) {
	o.SecureAccessAllowExternalUser = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DsUpdateRdp) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *DsUpdateRdp) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessRdGatewayServer returns the SecureAccessRdGatewayServer field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetSecureAccessRdGatewayServer() string {
	if o == nil || o.SecureAccessRdGatewayServer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessRdGatewayServer
}

// GetSecureAccessRdGatewayServerOk returns a tuple with the SecureAccessRdGatewayServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetSecureAccessRdGatewayServerOk() (*string, bool) {
	if o == nil || o.SecureAccessRdGatewayServer == nil {
		return nil, false
	}
	return o.SecureAccessRdGatewayServer, true
}

// HasSecureAccessRdGatewayServer returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasSecureAccessRdGatewayServer() bool {
	if o != nil && o.SecureAccessRdGatewayServer != nil {
		return true
	}

	return false
}

// SetSecureAccessRdGatewayServer gets a reference to the given string and assigns it to the SecureAccessRdGatewayServer field.
func (o *DsUpdateRdp) SetSecureAccessRdGatewayServer(v string) {
	o.SecureAccessRdGatewayServer = &v
}

// GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetSecureAccessRdpDomain() string {
	if o == nil || o.SecureAccessRdpDomain == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessRdpDomain
}

// GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetSecureAccessRdpDomainOk() (*string, bool) {
	if o == nil || o.SecureAccessRdpDomain == nil {
		return nil, false
	}
	return o.SecureAccessRdpDomain, true
}

// HasSecureAccessRdpDomain returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasSecureAccessRdpDomain() bool {
	if o != nil && o.SecureAccessRdpDomain != nil {
		return true
	}

	return false
}

// SetSecureAccessRdpDomain gets a reference to the given string and assigns it to the SecureAccessRdpDomain field.
func (o *DsUpdateRdp) SetSecureAccessRdpDomain(v string) {
	o.SecureAccessRdpDomain = &v
}

// GetSecureAccessRdpUser returns the SecureAccessRdpUser field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetSecureAccessRdpUser() string {
	if o == nil || o.SecureAccessRdpUser == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessRdpUser
}

// GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetSecureAccessRdpUserOk() (*string, bool) {
	if o == nil || o.SecureAccessRdpUser == nil {
		return nil, false
	}
	return o.SecureAccessRdpUser, true
}

// HasSecureAccessRdpUser returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasSecureAccessRdpUser() bool {
	if o != nil && o.SecureAccessRdpUser != nil {
		return true
	}

	return false
}

// SetSecureAccessRdpUser gets a reference to the given string and assigns it to the SecureAccessRdpUser field.
func (o *DsUpdateRdp) SetSecureAccessRdpUser(v string) {
	o.SecureAccessRdpUser = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DsUpdateRdp) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DsUpdateRdp) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DsUpdateRdp) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DsUpdateRdp) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DsUpdateRdp) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetWarnUserBeforeExpiration returns the WarnUserBeforeExpiration field value if set, zero value otherwise.
func (o *DsUpdateRdp) GetWarnUserBeforeExpiration() int64 {
	if o == nil || o.WarnUserBeforeExpiration == nil {
		var ret int64
		return ret
	}
	return *o.WarnUserBeforeExpiration
}

// GetWarnUserBeforeExpirationOk returns a tuple with the WarnUserBeforeExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateRdp) GetWarnUserBeforeExpirationOk() (*int64, bool) {
	if o == nil || o.WarnUserBeforeExpiration == nil {
		return nil, false
	}
	return o.WarnUserBeforeExpiration, true
}

// HasWarnUserBeforeExpiration returns a boolean if a field has been set.
func (o *DsUpdateRdp) HasWarnUserBeforeExpiration() bool {
	if o != nil && o.WarnUserBeforeExpiration != nil {
		return true
	}

	return false
}

// SetWarnUserBeforeExpiration gets a reference to the given int64 and assigns it to the WarnUserBeforeExpiration field.
func (o *DsUpdateRdp) SetWarnUserBeforeExpiration(v int64) {
	o.WarnUserBeforeExpiration = &v
}

func (o DsUpdateRdp) MarshalJSON() ([]byte, error) {
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
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
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
	if o.SecureAccessRdGatewayServer != nil {
		toSerialize["secure-access-rd-gateway-server"] = o.SecureAccessRdGatewayServer
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

type NullableDsUpdateRdp struct {
	value *DsUpdateRdp
	isSet bool
}

func (v NullableDsUpdateRdp) Get() *DsUpdateRdp {
	return v.value
}

func (v *NullableDsUpdateRdp) Set(val *DsUpdateRdp) {
	v.value = val
	v.isSet = true
}

func (v NullableDsUpdateRdp) IsSet() bool {
	return v.isSet
}

func (v *NullableDsUpdateRdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDsUpdateRdp(val *DsUpdateRdp) *NullableDsUpdateRdp {
	return &NullableDsUpdateRdp{value: val, isSet: true}
}

func (v NullableDsUpdateRdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDsUpdateRdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


