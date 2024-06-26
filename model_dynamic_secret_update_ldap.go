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

// DynamicSecretUpdateLdap dynamicSecretUpdateLdap is a command that updates ldap dynamic secret
type DynamicSecretUpdateLdap struct {
	ProviderType *string `json:"ProviderType,omitempty"`
	// Bind DN
	BindDn *string `json:"bind-dn,omitempty"`
	// Bind DN Password
	BindDnPassword *string `json:"bind-dn-password,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Externally provided username [true/false]
	ExternalUsername *string `json:"external-username,omitempty"`
	// Group DN which the temporary user should be added
	GroupDn *string `json:"group-dn,omitempty"`
	// Host provider type [explicit/target], Default Host provider is explicit, Relevant only for Secure Remote Access of ssh cert issuer, ldap rotated secret and ldap dynamic secret
	HostProvider *string `json:"host-provider,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// CA Certificate File Content
	LdapCaCert *string `json:"ldap-ca-cert,omitempty"`
	// LDAP Server URL
	LdapUrl *string `json:"ldap-url,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret new name
	NewName *string `json:"new-name,omitempty"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers)
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	// RD Gateway server
	SecureAccessRdGatewayServer *string `json:"secure-access-rd-gateway-server,omitempty"`
	// Required when the Dynamic Secret is used for a domain user
	SecureAccessRdpDomain *string `json:"secure-access-rdp-domain,omitempty"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// A list of linked targets to be associated, Relevant only for Secure Remote Access for ssh cert issuer, ldap rotated secret and ldap dynamic secret, To specify multiple targets use argument multiple times
	Target *[]string `json:"target,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Token expiration
	TokenExpiration *string `json:"token-expiration,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User Attribute
	UserAttribute *string `json:"user-attribute,omitempty"`
	// User DN
	UserDn *string `json:"user-dn,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewDynamicSecretUpdateLdap instantiates a new DynamicSecretUpdateLdap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdateLdap(name string, ) *DynamicSecretUpdateLdap {
	this := DynamicSecretUpdateLdap{}
	var externalUsername string = "false"
	this.ExternalUsername = &externalUsername
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretUpdateLdapWithDefaults instantiates a new DynamicSecretUpdateLdap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdateLdapWithDefaults() *DynamicSecretUpdateLdap {
	this := DynamicSecretUpdateLdap{}
	var externalUsername string = "false"
	this.ExternalUsername = &externalUsername
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetProviderType returns the ProviderType field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetProviderType() string {
	if o == nil || o.ProviderType == nil {
		var ret string
		return ret
	}
	return *o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetProviderTypeOk() (*string, bool) {
	if o == nil || o.ProviderType == nil {
		return nil, false
	}
	return o.ProviderType, true
}

// HasProviderType returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasProviderType() bool {
	if o != nil && o.ProviderType != nil {
		return true
	}

	return false
}

// SetProviderType gets a reference to the given string and assigns it to the ProviderType field.
func (o *DynamicSecretUpdateLdap) SetProviderType(v string) {
	o.ProviderType = &v
}

// GetBindDn returns the BindDn field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetBindDn() string {
	if o == nil || o.BindDn == nil {
		var ret string
		return ret
	}
	return *o.BindDn
}

// GetBindDnOk returns a tuple with the BindDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetBindDnOk() (*string, bool) {
	if o == nil || o.BindDn == nil {
		return nil, false
	}
	return o.BindDn, true
}

// HasBindDn returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasBindDn() bool {
	if o != nil && o.BindDn != nil {
		return true
	}

	return false
}

// SetBindDn gets a reference to the given string and assigns it to the BindDn field.
func (o *DynamicSecretUpdateLdap) SetBindDn(v string) {
	o.BindDn = &v
}

// GetBindDnPassword returns the BindDnPassword field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetBindDnPassword() string {
	if o == nil || o.BindDnPassword == nil {
		var ret string
		return ret
	}
	return *o.BindDnPassword
}

// GetBindDnPasswordOk returns a tuple with the BindDnPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetBindDnPasswordOk() (*string, bool) {
	if o == nil || o.BindDnPassword == nil {
		return nil, false
	}
	return o.BindDnPassword, true
}

// HasBindDnPassword returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasBindDnPassword() bool {
	if o != nil && o.BindDnPassword != nil {
		return true
	}

	return false
}

// SetBindDnPassword gets a reference to the given string and assigns it to the BindDnPassword field.
func (o *DynamicSecretUpdateLdap) SetBindDnPassword(v string) {
	o.BindDnPassword = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretUpdateLdap) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretUpdateLdap) SetDescription(v string) {
	o.Description = &v
}

// GetExternalUsername returns the ExternalUsername field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetExternalUsername() string {
	if o == nil || o.ExternalUsername == nil {
		var ret string
		return ret
	}
	return *o.ExternalUsername
}

// GetExternalUsernameOk returns a tuple with the ExternalUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetExternalUsernameOk() (*string, bool) {
	if o == nil || o.ExternalUsername == nil {
		return nil, false
	}
	return o.ExternalUsername, true
}

// HasExternalUsername returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasExternalUsername() bool {
	if o != nil && o.ExternalUsername != nil {
		return true
	}

	return false
}

// SetExternalUsername gets a reference to the given string and assigns it to the ExternalUsername field.
func (o *DynamicSecretUpdateLdap) SetExternalUsername(v string) {
	o.ExternalUsername = &v
}

// GetGroupDn returns the GroupDn field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetGroupDn() string {
	if o == nil || o.GroupDn == nil {
		var ret string
		return ret
	}
	return *o.GroupDn
}

// GetGroupDnOk returns a tuple with the GroupDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetGroupDnOk() (*string, bool) {
	if o == nil || o.GroupDn == nil {
		return nil, false
	}
	return o.GroupDn, true
}

// HasGroupDn returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasGroupDn() bool {
	if o != nil && o.GroupDn != nil {
		return true
	}

	return false
}

// SetGroupDn gets a reference to the given string and assigns it to the GroupDn field.
func (o *DynamicSecretUpdateLdap) SetGroupDn(v string) {
	o.GroupDn = &v
}

// GetHostProvider returns the HostProvider field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetHostProvider() string {
	if o == nil || o.HostProvider == nil {
		var ret string
		return ret
	}
	return *o.HostProvider
}

// GetHostProviderOk returns a tuple with the HostProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetHostProviderOk() (*string, bool) {
	if o == nil || o.HostProvider == nil {
		return nil, false
	}
	return o.HostProvider, true
}

// HasHostProvider returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasHostProvider() bool {
	if o != nil && o.HostProvider != nil {
		return true
	}

	return false
}

// SetHostProvider gets a reference to the given string and assigns it to the HostProvider field.
func (o *DynamicSecretUpdateLdap) SetHostProvider(v string) {
	o.HostProvider = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretUpdateLdap) SetJson(v bool) {
	o.Json = &v
}

// GetLdapCaCert returns the LdapCaCert field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetLdapCaCert() string {
	if o == nil || o.LdapCaCert == nil {
		var ret string
		return ret
	}
	return *o.LdapCaCert
}

// GetLdapCaCertOk returns a tuple with the LdapCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetLdapCaCertOk() (*string, bool) {
	if o == nil || o.LdapCaCert == nil {
		return nil, false
	}
	return o.LdapCaCert, true
}

// HasLdapCaCert returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasLdapCaCert() bool {
	if o != nil && o.LdapCaCert != nil {
		return true
	}

	return false
}

// SetLdapCaCert gets a reference to the given string and assigns it to the LdapCaCert field.
func (o *DynamicSecretUpdateLdap) SetLdapCaCert(v string) {
	o.LdapCaCert = &v
}

// GetLdapUrl returns the LdapUrl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetLdapUrl() string {
	if o == nil || o.LdapUrl == nil {
		var ret string
		return ret
	}
	return *o.LdapUrl
}

// GetLdapUrlOk returns a tuple with the LdapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetLdapUrlOk() (*string, bool) {
	if o == nil || o.LdapUrl == nil {
		return nil, false
	}
	return o.LdapUrl, true
}

// HasLdapUrl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasLdapUrl() bool {
	if o != nil && o.LdapUrl != nil {
		return true
	}

	return false
}

// SetLdapUrl gets a reference to the given string and assigns it to the LdapUrl field.
func (o *DynamicSecretUpdateLdap) SetLdapUrl(v string) {
	o.LdapUrl = &v
}

// GetName returns the Name field value
func (o *DynamicSecretUpdateLdap) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretUpdateLdap) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DynamicSecretUpdateLdap) SetNewName(v string) {
	o.NewName = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *DynamicSecretUpdateLdap) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretUpdateLdap) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DynamicSecretUpdateLdap) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *DynamicSecretUpdateLdap) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessRdGatewayServer returns the SecureAccessRdGatewayServer field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetSecureAccessRdGatewayServer() string {
	if o == nil || o.SecureAccessRdGatewayServer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessRdGatewayServer
}

// GetSecureAccessRdGatewayServerOk returns a tuple with the SecureAccessRdGatewayServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetSecureAccessRdGatewayServerOk() (*string, bool) {
	if o == nil || o.SecureAccessRdGatewayServer == nil {
		return nil, false
	}
	return o.SecureAccessRdGatewayServer, true
}

// HasSecureAccessRdGatewayServer returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasSecureAccessRdGatewayServer() bool {
	if o != nil && o.SecureAccessRdGatewayServer != nil {
		return true
	}

	return false
}

// SetSecureAccessRdGatewayServer gets a reference to the given string and assigns it to the SecureAccessRdGatewayServer field.
func (o *DynamicSecretUpdateLdap) SetSecureAccessRdGatewayServer(v string) {
	o.SecureAccessRdGatewayServer = &v
}

// GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetSecureAccessRdpDomain() string {
	if o == nil || o.SecureAccessRdpDomain == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessRdpDomain
}

// GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetSecureAccessRdpDomainOk() (*string, bool) {
	if o == nil || o.SecureAccessRdpDomain == nil {
		return nil, false
	}
	return o.SecureAccessRdpDomain, true
}

// HasSecureAccessRdpDomain returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasSecureAccessRdpDomain() bool {
	if o != nil && o.SecureAccessRdpDomain != nil {
		return true
	}

	return false
}

// SetSecureAccessRdpDomain gets a reference to the given string and assigns it to the SecureAccessRdpDomain field.
func (o *DynamicSecretUpdateLdap) SetSecureAccessRdpDomain(v string) {
	o.SecureAccessRdpDomain = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretUpdateLdap) SetTags(v []string) {
	o.Tags = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetTarget() []string {
	if o == nil || o.Target == nil {
		var ret []string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetTargetOk() (*[]string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given []string and assigns it to the Target field.
func (o *DynamicSecretUpdateLdap) SetTarget(v []string) {
	o.Target = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretUpdateLdap) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretUpdateLdap) SetToken(v string) {
	o.Token = &v
}

// GetTokenExpiration returns the TokenExpiration field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetTokenExpiration() string {
	if o == nil || o.TokenExpiration == nil {
		var ret string
		return ret
	}
	return *o.TokenExpiration
}

// GetTokenExpirationOk returns a tuple with the TokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetTokenExpirationOk() (*string, bool) {
	if o == nil || o.TokenExpiration == nil {
		return nil, false
	}
	return o.TokenExpiration, true
}

// HasTokenExpiration returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasTokenExpiration() bool {
	if o != nil && o.TokenExpiration != nil {
		return true
	}

	return false
}

// SetTokenExpiration gets a reference to the given string and assigns it to the TokenExpiration field.
func (o *DynamicSecretUpdateLdap) SetTokenExpiration(v string) {
	o.TokenExpiration = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretUpdateLdap) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserAttribute returns the UserAttribute field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetUserAttribute() string {
	if o == nil || o.UserAttribute == nil {
		var ret string
		return ret
	}
	return *o.UserAttribute
}

// GetUserAttributeOk returns a tuple with the UserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetUserAttributeOk() (*string, bool) {
	if o == nil || o.UserAttribute == nil {
		return nil, false
	}
	return o.UserAttribute, true
}

// HasUserAttribute returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasUserAttribute() bool {
	if o != nil && o.UserAttribute != nil {
		return true
	}

	return false
}

// SetUserAttribute gets a reference to the given string and assigns it to the UserAttribute field.
func (o *DynamicSecretUpdateLdap) SetUserAttribute(v string) {
	o.UserAttribute = &v
}

// GetUserDn returns the UserDn field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetUserDn() string {
	if o == nil || o.UserDn == nil {
		var ret string
		return ret
	}
	return *o.UserDn
}

// GetUserDnOk returns a tuple with the UserDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetUserDnOk() (*string, bool) {
	if o == nil || o.UserDn == nil {
		return nil, false
	}
	return o.UserDn, true
}

// HasUserDn returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasUserDn() bool {
	if o != nil && o.UserDn != nil {
		return true
	}

	return false
}

// SetUserDn gets a reference to the given string and assigns it to the UserDn field.
func (o *DynamicSecretUpdateLdap) SetUserDn(v string) {
	o.UserDn = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateLdap) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateLdap) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateLdap) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretUpdateLdap) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretUpdateLdap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProviderType != nil {
		toSerialize["ProviderType"] = o.ProviderType
	}
	if o.BindDn != nil {
		toSerialize["bind-dn"] = o.BindDn
	}
	if o.BindDnPassword != nil {
		toSerialize["bind-dn-password"] = o.BindDnPassword
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExternalUsername != nil {
		toSerialize["external-username"] = o.ExternalUsername
	}
	if o.GroupDn != nil {
		toSerialize["group-dn"] = o.GroupDn
	}
	if o.HostProvider != nil {
		toSerialize["host-provider"] = o.HostProvider
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.LdapCaCert != nil {
		toSerialize["ldap-ca-cert"] = o.LdapCaCert
	}
	if o.LdapUrl != nil {
		toSerialize["ldap-url"] = o.LdapUrl
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
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.TargetName != nil {
		toSerialize["target-name"] = o.TargetName
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.TokenExpiration != nil {
		toSerialize["token-expiration"] = o.TokenExpiration
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UserAttribute != nil {
		toSerialize["user-attribute"] = o.UserAttribute
	}
	if o.UserDn != nil {
		toSerialize["user-dn"] = o.UserDn
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretUpdateLdap struct {
	value *DynamicSecretUpdateLdap
	isSet bool
}

func (v NullableDynamicSecretUpdateLdap) Get() *DynamicSecretUpdateLdap {
	return v.value
}

func (v *NullableDynamicSecretUpdateLdap) Set(val *DynamicSecretUpdateLdap) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdateLdap) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdateLdap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdateLdap(val *DynamicSecretUpdateLdap) *NullableDynamicSecretUpdateLdap {
	return &NullableDynamicSecretUpdateLdap{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdateLdap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdateLdap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


