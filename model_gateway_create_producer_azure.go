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

// GatewayCreateProducerAzure gatewayCreateProducerAzure is a command that creates azure producer
type GatewayCreateProducerAzure struct {
	// Azure App Object Id
	AppObjId *string `json:"app-obj-id,omitempty"`
	// Azure Client ID
	AzureClientId *string `json:"azure-client-id,omitempty"`
	// Azure Client Secret
	AzureClientSecret *string `json:"azure-client-secret,omitempty"`
	// Azure Tenant ID
	AzureTenantId *string `json:"azure-tenant-id,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// FixedUserClaimKeyname
	FixedUserClaimKeyname *string `json:"fixed-user-claim-keyname,omitempty"`
	// Fixed user
	FixedUserOnly *bool `json:"fixed-user-only,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
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
	// User Group Object Id
	UserGroupObjId *string `json:"user-group-obj-id,omitempty"`
	// Azure User portal access
	UserPortalAccess *bool `json:"user-portal-access,omitempty"`
	// User Principal Name
	UserPrincipalName *string `json:"user-principal-name,omitempty"`
	// Azure User programmatic access
	UserProgrammaticAccess *bool `json:"user-programmatic-access,omitempty"`
	// User Role Template Id
	UserRoleTemplateId *string `json:"user-role-template-id,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayCreateProducerAzure instantiates a new GatewayCreateProducerAzure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerAzure(name string) *GatewayCreateProducerAzure {
	this := GatewayCreateProducerAzure{}
	var fixedUserClaimKeyname string = "false"
	this.FixedUserClaimKeyname = &fixedUserClaimKeyname
	var fixedUserOnly bool = false
	this.FixedUserOnly = &fixedUserOnly
	var json bool = false
	this.Json = &json
	this.Name = name
	var secureAccessWeb bool = true
	this.SecureAccessWeb = &secureAccessWeb
	var secureAccessWebBrowsing bool = false
	this.SecureAccessWebBrowsing = &secureAccessWebBrowsing
	var secureAccessWebProxy bool = false
	this.SecureAccessWebProxy = &secureAccessWebProxy
	var userPortalAccess bool = false
	this.UserPortalAccess = &userPortalAccess
	var userProgrammaticAccess bool = false
	this.UserProgrammaticAccess = &userProgrammaticAccess
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerAzureWithDefaults instantiates a new GatewayCreateProducerAzure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerAzureWithDefaults() *GatewayCreateProducerAzure {
	this := GatewayCreateProducerAzure{}
	var fixedUserClaimKeyname string = "false"
	this.FixedUserClaimKeyname = &fixedUserClaimKeyname
	var fixedUserOnly bool = false
	this.FixedUserOnly = &fixedUserOnly
	var json bool = false
	this.Json = &json
	var secureAccessWeb bool = true
	this.SecureAccessWeb = &secureAccessWeb
	var secureAccessWebBrowsing bool = false
	this.SecureAccessWebBrowsing = &secureAccessWebBrowsing
	var secureAccessWebProxy bool = false
	this.SecureAccessWebProxy = &secureAccessWebProxy
	var userPortalAccess bool = false
	this.UserPortalAccess = &userPortalAccess
	var userProgrammaticAccess bool = false
	this.UserProgrammaticAccess = &userProgrammaticAccess
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetAppObjId returns the AppObjId field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetAppObjId() string {
	if o == nil || o.AppObjId == nil {
		var ret string
		return ret
	}
	return *o.AppObjId
}

// GetAppObjIdOk returns a tuple with the AppObjId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetAppObjIdOk() (*string, bool) {
	if o == nil || o.AppObjId == nil {
		return nil, false
	}
	return o.AppObjId, true
}

// HasAppObjId returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasAppObjId() bool {
	if o != nil && o.AppObjId != nil {
		return true
	}

	return false
}

// SetAppObjId gets a reference to the given string and assigns it to the AppObjId field.
func (o *GatewayCreateProducerAzure) SetAppObjId(v string) {
	o.AppObjId = &v
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetAzureClientId() string {
	if o == nil || o.AzureClientId == nil {
		var ret string
		return ret
	}
	return *o.AzureClientId
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetAzureClientIdOk() (*string, bool) {
	if o == nil || o.AzureClientId == nil {
		return nil, false
	}
	return o.AzureClientId, true
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasAzureClientId() bool {
	if o != nil && o.AzureClientId != nil {
		return true
	}

	return false
}

// SetAzureClientId gets a reference to the given string and assigns it to the AzureClientId field.
func (o *GatewayCreateProducerAzure) SetAzureClientId(v string) {
	o.AzureClientId = &v
}

// GetAzureClientSecret returns the AzureClientSecret field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetAzureClientSecret() string {
	if o == nil || o.AzureClientSecret == nil {
		var ret string
		return ret
	}
	return *o.AzureClientSecret
}

// GetAzureClientSecretOk returns a tuple with the AzureClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetAzureClientSecretOk() (*string, bool) {
	if o == nil || o.AzureClientSecret == nil {
		return nil, false
	}
	return o.AzureClientSecret, true
}

// HasAzureClientSecret returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasAzureClientSecret() bool {
	if o != nil && o.AzureClientSecret != nil {
		return true
	}

	return false
}

// SetAzureClientSecret gets a reference to the given string and assigns it to the AzureClientSecret field.
func (o *GatewayCreateProducerAzure) SetAzureClientSecret(v string) {
	o.AzureClientSecret = &v
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || o.AzureTenantId == nil {
		return nil, false
	}
	return o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId != nil {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *GatewayCreateProducerAzure) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerAzure) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetFixedUserClaimKeyname returns the FixedUserClaimKeyname field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetFixedUserClaimKeyname() string {
	if o == nil || o.FixedUserClaimKeyname == nil {
		var ret string
		return ret
	}
	return *o.FixedUserClaimKeyname
}

// GetFixedUserClaimKeynameOk returns a tuple with the FixedUserClaimKeyname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetFixedUserClaimKeynameOk() (*string, bool) {
	if o == nil || o.FixedUserClaimKeyname == nil {
		return nil, false
	}
	return o.FixedUserClaimKeyname, true
}

// HasFixedUserClaimKeyname returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasFixedUserClaimKeyname() bool {
	if o != nil && o.FixedUserClaimKeyname != nil {
		return true
	}

	return false
}

// SetFixedUserClaimKeyname gets a reference to the given string and assigns it to the FixedUserClaimKeyname field.
func (o *GatewayCreateProducerAzure) SetFixedUserClaimKeyname(v string) {
	o.FixedUserClaimKeyname = &v
}

// GetFixedUserOnly returns the FixedUserOnly field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetFixedUserOnly() bool {
	if o == nil || o.FixedUserOnly == nil {
		var ret bool
		return ret
	}
	return *o.FixedUserOnly
}

// GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetFixedUserOnlyOk() (*bool, bool) {
	if o == nil || o.FixedUserOnly == nil {
		return nil, false
	}
	return o.FixedUserOnly, true
}

// HasFixedUserOnly returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasFixedUserOnly() bool {
	if o != nil && o.FixedUserOnly != nil {
		return true
	}

	return false
}

// SetFixedUserOnly gets a reference to the given bool and assigns it to the FixedUserOnly field.
func (o *GatewayCreateProducerAzure) SetFixedUserOnly(v bool) {
	o.FixedUserOnly = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerAzure) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerAzure) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerAzure) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerAzure) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayCreateProducerAzure) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayCreateProducerAzure) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetSecureAccessWebBrowsing() bool {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebBrowsing
}

// GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetSecureAccessWebBrowsingOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebBrowsing == nil {
		return nil, false
	}
	return o.SecureAccessWebBrowsing, true
}

// HasSecureAccessWebBrowsing returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasSecureAccessWebBrowsing() bool {
	if o != nil && o.SecureAccessWebBrowsing != nil {
		return true
	}

	return false
}

// SetSecureAccessWebBrowsing gets a reference to the given bool and assigns it to the SecureAccessWebBrowsing field.
func (o *GatewayCreateProducerAzure) SetSecureAccessWebBrowsing(v bool) {
	o.SecureAccessWebBrowsing = &v
}

// GetSecureAccessWebProxy returns the SecureAccessWebProxy field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetSecureAccessWebProxy() bool {
	if o == nil || o.SecureAccessWebProxy == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWebProxy
}

// GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetSecureAccessWebProxyOk() (*bool, bool) {
	if o == nil || o.SecureAccessWebProxy == nil {
		return nil, false
	}
	return o.SecureAccessWebProxy, true
}

// HasSecureAccessWebProxy returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasSecureAccessWebProxy() bool {
	if o != nil && o.SecureAccessWebProxy != nil {
		return true
	}

	return false
}

// SetSecureAccessWebProxy gets a reference to the given bool and assigns it to the SecureAccessWebProxy field.
func (o *GatewayCreateProducerAzure) SetSecureAccessWebProxy(v bool) {
	o.SecureAccessWebProxy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerAzure) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerAzure) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerAzure) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerAzure) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserGroupObjId returns the UserGroupObjId field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserGroupObjId() string {
	if o == nil || o.UserGroupObjId == nil {
		var ret string
		return ret
	}
	return *o.UserGroupObjId
}

// GetUserGroupObjIdOk returns a tuple with the UserGroupObjId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserGroupObjIdOk() (*string, bool) {
	if o == nil || o.UserGroupObjId == nil {
		return nil, false
	}
	return o.UserGroupObjId, true
}

// HasUserGroupObjId returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserGroupObjId() bool {
	if o != nil && o.UserGroupObjId != nil {
		return true
	}

	return false
}

// SetUserGroupObjId gets a reference to the given string and assigns it to the UserGroupObjId field.
func (o *GatewayCreateProducerAzure) SetUserGroupObjId(v string) {
	o.UserGroupObjId = &v
}

// GetUserPortalAccess returns the UserPortalAccess field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserPortalAccess() bool {
	if o == nil || o.UserPortalAccess == nil {
		var ret bool
		return ret
	}
	return *o.UserPortalAccess
}

// GetUserPortalAccessOk returns a tuple with the UserPortalAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserPortalAccessOk() (*bool, bool) {
	if o == nil || o.UserPortalAccess == nil {
		return nil, false
	}
	return o.UserPortalAccess, true
}

// HasUserPortalAccess returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserPortalAccess() bool {
	if o != nil && o.UserPortalAccess != nil {
		return true
	}

	return false
}

// SetUserPortalAccess gets a reference to the given bool and assigns it to the UserPortalAccess field.
func (o *GatewayCreateProducerAzure) SetUserPortalAccess(v bool) {
	o.UserPortalAccess = &v
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil || o.UserPrincipalName == nil {
		return nil, false
	}
	return o.UserPrincipalName, true
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName != nil {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.
func (o *GatewayCreateProducerAzure) SetUserPrincipalName(v string) {
	o.UserPrincipalName = &v
}

// GetUserProgrammaticAccess returns the UserProgrammaticAccess field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserProgrammaticAccess() bool {
	if o == nil || o.UserProgrammaticAccess == nil {
		var ret bool
		return ret
	}
	return *o.UserProgrammaticAccess
}

// GetUserProgrammaticAccessOk returns a tuple with the UserProgrammaticAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserProgrammaticAccessOk() (*bool, bool) {
	if o == nil || o.UserProgrammaticAccess == nil {
		return nil, false
	}
	return o.UserProgrammaticAccess, true
}

// HasUserProgrammaticAccess returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserProgrammaticAccess() bool {
	if o != nil && o.UserProgrammaticAccess != nil {
		return true
	}

	return false
}

// SetUserProgrammaticAccess gets a reference to the given bool and assigns it to the UserProgrammaticAccess field.
func (o *GatewayCreateProducerAzure) SetUserProgrammaticAccess(v bool) {
	o.UserProgrammaticAccess = &v
}

// GetUserRoleTemplateId returns the UserRoleTemplateId field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserRoleTemplateId() string {
	if o == nil || o.UserRoleTemplateId == nil {
		var ret string
		return ret
	}
	return *o.UserRoleTemplateId
}

// GetUserRoleTemplateIdOk returns a tuple with the UserRoleTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserRoleTemplateIdOk() (*string, bool) {
	if o == nil || o.UserRoleTemplateId == nil {
		return nil, false
	}
	return o.UserRoleTemplateId, true
}

// HasUserRoleTemplateId returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserRoleTemplateId() bool {
	if o != nil && o.UserRoleTemplateId != nil {
		return true
	}

	return false
}

// SetUserRoleTemplateId gets a reference to the given string and assigns it to the UserRoleTemplateId field.
func (o *GatewayCreateProducerAzure) SetUserRoleTemplateId(v string) {
	o.UserRoleTemplateId = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerAzure) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerAzure) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerAzure) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerAzure) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerAzure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppObjId != nil {
		toSerialize["app-obj-id"] = o.AppObjId
	}
	if o.AzureClientId != nil {
		toSerialize["azure-client-id"] = o.AzureClientId
	}
	if o.AzureClientSecret != nil {
		toSerialize["azure-client-secret"] = o.AzureClientSecret
	}
	if o.AzureTenantId != nil {
		toSerialize["azure-tenant-id"] = o.AzureTenantId
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.FixedUserClaimKeyname != nil {
		toSerialize["fixed-user-claim-keyname"] = o.FixedUserClaimKeyname
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
	if o.UserGroupObjId != nil {
		toSerialize["user-group-obj-id"] = o.UserGroupObjId
	}
	if o.UserPortalAccess != nil {
		toSerialize["user-portal-access"] = o.UserPortalAccess
	}
	if o.UserPrincipalName != nil {
		toSerialize["user-principal-name"] = o.UserPrincipalName
	}
	if o.UserProgrammaticAccess != nil {
		toSerialize["user-programmatic-access"] = o.UserProgrammaticAccess
	}
	if o.UserRoleTemplateId != nil {
		toSerialize["user-role-template-id"] = o.UserRoleTemplateId
	}
	if o.UserTtl != nil {
		toSerialize["user-ttl"] = o.UserTtl
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerAzure struct {
	value *GatewayCreateProducerAzure
	isSet bool
}

func (v NullableGatewayCreateProducerAzure) Get() *GatewayCreateProducerAzure {
	return v.value
}

func (v *NullableGatewayCreateProducerAzure) Set(val *GatewayCreateProducerAzure) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerAzure) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerAzure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerAzure(val *GatewayCreateProducerAzure) *NullableGatewayCreateProducerAzure {
	return &NullableGatewayCreateProducerAzure{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerAzure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerAzure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


