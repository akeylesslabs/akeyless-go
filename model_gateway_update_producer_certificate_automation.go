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

// GatewayUpdateProducerCertificateAutomation gatewayUpdateProducerCertificateAutomation is a command that updates a Certificate Automation dynamic secret producer to dynamically update certificates generated by Venafi or have Akeyless generated certificates using PKI be monitored by Venafi
type GatewayUpdateProducerCertificateAutomation struct {
	// Admin credentials rotation interval (days)
	AdminRotationIntervalDays *int64 `json:"admin-rotation-interval-days,omitempty"`
	// Allow subdomains
	AllowSubdomains *bool `json:"allow-subdomains,omitempty"`
	// Allowed domains
	AllowedDomains *[]string `json:"allowed-domains,omitempty"`
	// Auto generated folder
	AutoGeneratedFolder *string `json:"auto-generated-folder,omitempty"`
	// Automatic admin credentials rotation
	EnableAdminRotation *bool `json:"enable-admin-rotation,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Root first in chain
	RootFirstInChain *bool `json:"root-first-in-chain,omitempty"`
	// Use Akeyless PKI issuer or Venafi issuer
	SignUsingAkeylessPki *bool `json:"sign-using-akeyless-pki,omitempty"`
	// Signer key name
	SignerKeyName *string `json:"signer-key-name,omitempty"`
	// Store private key
	StorePrivateKey *bool `json:"store-private-key,omitempty"`
	// List of the tags attached to this secret
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL in time.Duration format (2160h / 129600m / etc...). When using sign-using-akeyless-pki certificates created will have this validity period, otherwise the user-ttl is taken from the Validity Period field of the Zone's' Issuing Template. When using cert-manager it is advised to have a TTL of above 60 days (1440h). For more information - https://cert-manager.io/docs/usage/certificate/
	UserTtl *string `json:"user-ttl,omitempty"`
	// Venafi API key
	VenafiApiKey *string `json:"venafi-api-key,omitempty"`
	// Venafi Baseurl
	VenafiBaseurl *string `json:"venafi-baseurl,omitempty"`
	// Venafi Password
	VenafiPassword *string `json:"venafi-password,omitempty"`
	// Venafi using TPP
	VenafiUseTpp *bool `json:"venafi-use-tpp,omitempty"`
	// Venafi Username
	VenafiUsername *string `json:"venafi-username,omitempty"`
	// Venafi Zone
	VenafiZone *string `json:"venafi-zone,omitempty"`
}

// NewGatewayUpdateProducerCertificateAutomation instantiates a new GatewayUpdateProducerCertificateAutomation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerCertificateAutomation(name string, ) *GatewayUpdateProducerCertificateAutomation {
	this := GatewayUpdateProducerCertificateAutomation{}
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	this.Name = name
	return &this
}

// NewGatewayUpdateProducerCertificateAutomationWithDefaults instantiates a new GatewayUpdateProducerCertificateAutomation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerCertificateAutomationWithDefaults() *GatewayUpdateProducerCertificateAutomation {
	this := GatewayUpdateProducerCertificateAutomation{}
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	return &this
}

// GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetAdminRotationIntervalDays() int64 {
	if o == nil || o.AdminRotationIntervalDays == nil {
		var ret int64
		return ret
	}
	return *o.AdminRotationIntervalDays
}

// GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetAdminRotationIntervalDaysOk() (*int64, bool) {
	if o == nil || o.AdminRotationIntervalDays == nil {
		return nil, false
	}
	return o.AdminRotationIntervalDays, true
}

// HasAdminRotationIntervalDays returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasAdminRotationIntervalDays() bool {
	if o != nil && o.AdminRotationIntervalDays != nil {
		return true
	}

	return false
}

// SetAdminRotationIntervalDays gets a reference to the given int64 and assigns it to the AdminRotationIntervalDays field.
func (o *GatewayUpdateProducerCertificateAutomation) SetAdminRotationIntervalDays(v int64) {
	o.AdminRotationIntervalDays = &v
}

// GetAllowSubdomains returns the AllowSubdomains field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetAllowSubdomains() bool {
	if o == nil || o.AllowSubdomains == nil {
		var ret bool
		return ret
	}
	return *o.AllowSubdomains
}

// GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetAllowSubdomainsOk() (*bool, bool) {
	if o == nil || o.AllowSubdomains == nil {
		return nil, false
	}
	return o.AllowSubdomains, true
}

// HasAllowSubdomains returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasAllowSubdomains() bool {
	if o != nil && o.AllowSubdomains != nil {
		return true
	}

	return false
}

// SetAllowSubdomains gets a reference to the given bool and assigns it to the AllowSubdomains field.
func (o *GatewayUpdateProducerCertificateAutomation) SetAllowSubdomains(v bool) {
	o.AllowSubdomains = &v
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetAllowedDomains() []string {
	if o == nil || o.AllowedDomains == nil {
		var ret []string
		return ret
	}
	return *o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetAllowedDomainsOk() (*[]string, bool) {
	if o == nil || o.AllowedDomains == nil {
		return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasAllowedDomains() bool {
	if o != nil && o.AllowedDomains != nil {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given []string and assigns it to the AllowedDomains field.
func (o *GatewayUpdateProducerCertificateAutomation) SetAllowedDomains(v []string) {
	o.AllowedDomains = &v
}

// GetAutoGeneratedFolder returns the AutoGeneratedFolder field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetAutoGeneratedFolder() string {
	if o == nil || o.AutoGeneratedFolder == nil {
		var ret string
		return ret
	}
	return *o.AutoGeneratedFolder
}

// GetAutoGeneratedFolderOk returns a tuple with the AutoGeneratedFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetAutoGeneratedFolderOk() (*string, bool) {
	if o == nil || o.AutoGeneratedFolder == nil {
		return nil, false
	}
	return o.AutoGeneratedFolder, true
}

// HasAutoGeneratedFolder returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasAutoGeneratedFolder() bool {
	if o != nil && o.AutoGeneratedFolder != nil {
		return true
	}

	return false
}

// SetAutoGeneratedFolder gets a reference to the given string and assigns it to the AutoGeneratedFolder field.
func (o *GatewayUpdateProducerCertificateAutomation) SetAutoGeneratedFolder(v string) {
	o.AutoGeneratedFolder = &v
}

// GetEnableAdminRotation returns the EnableAdminRotation field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetEnableAdminRotation() bool {
	if o == nil || o.EnableAdminRotation == nil {
		var ret bool
		return ret
	}
	return *o.EnableAdminRotation
}

// GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetEnableAdminRotationOk() (*bool, bool) {
	if o == nil || o.EnableAdminRotation == nil {
		return nil, false
	}
	return o.EnableAdminRotation, true
}

// HasEnableAdminRotation returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasEnableAdminRotation() bool {
	if o != nil && o.EnableAdminRotation != nil {
		return true
	}

	return false
}

// SetEnableAdminRotation gets a reference to the given bool and assigns it to the EnableAdminRotation field.
func (o *GatewayUpdateProducerCertificateAutomation) SetEnableAdminRotation(v bool) {
	o.EnableAdminRotation = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerCertificateAutomation) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerCertificateAutomation) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerCertificateAutomation) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerCertificateAutomation) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRootFirstInChain returns the RootFirstInChain field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetRootFirstInChain() bool {
	if o == nil || o.RootFirstInChain == nil {
		var ret bool
		return ret
	}
	return *o.RootFirstInChain
}

// GetRootFirstInChainOk returns a tuple with the RootFirstInChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetRootFirstInChainOk() (*bool, bool) {
	if o == nil || o.RootFirstInChain == nil {
		return nil, false
	}
	return o.RootFirstInChain, true
}

// HasRootFirstInChain returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasRootFirstInChain() bool {
	if o != nil && o.RootFirstInChain != nil {
		return true
	}

	return false
}

// SetRootFirstInChain gets a reference to the given bool and assigns it to the RootFirstInChain field.
func (o *GatewayUpdateProducerCertificateAutomation) SetRootFirstInChain(v bool) {
	o.RootFirstInChain = &v
}

// GetSignUsingAkeylessPki returns the SignUsingAkeylessPki field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetSignUsingAkeylessPki() bool {
	if o == nil || o.SignUsingAkeylessPki == nil {
		var ret bool
		return ret
	}
	return *o.SignUsingAkeylessPki
}

// GetSignUsingAkeylessPkiOk returns a tuple with the SignUsingAkeylessPki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetSignUsingAkeylessPkiOk() (*bool, bool) {
	if o == nil || o.SignUsingAkeylessPki == nil {
		return nil, false
	}
	return o.SignUsingAkeylessPki, true
}

// HasSignUsingAkeylessPki returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasSignUsingAkeylessPki() bool {
	if o != nil && o.SignUsingAkeylessPki != nil {
		return true
	}

	return false
}

// SetSignUsingAkeylessPki gets a reference to the given bool and assigns it to the SignUsingAkeylessPki field.
func (o *GatewayUpdateProducerCertificateAutomation) SetSignUsingAkeylessPki(v bool) {
	o.SignUsingAkeylessPki = &v
}

// GetSignerKeyName returns the SignerKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetSignerKeyName() string {
	if o == nil || o.SignerKeyName == nil {
		var ret string
		return ret
	}
	return *o.SignerKeyName
}

// GetSignerKeyNameOk returns a tuple with the SignerKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetSignerKeyNameOk() (*string, bool) {
	if o == nil || o.SignerKeyName == nil {
		return nil, false
	}
	return o.SignerKeyName, true
}

// HasSignerKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasSignerKeyName() bool {
	if o != nil && o.SignerKeyName != nil {
		return true
	}

	return false
}

// SetSignerKeyName gets a reference to the given string and assigns it to the SignerKeyName field.
func (o *GatewayUpdateProducerCertificateAutomation) SetSignerKeyName(v string) {
	o.SignerKeyName = &v
}

// GetStorePrivateKey returns the StorePrivateKey field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetStorePrivateKey() bool {
	if o == nil || o.StorePrivateKey == nil {
		var ret bool
		return ret
	}
	return *o.StorePrivateKey
}

// GetStorePrivateKeyOk returns a tuple with the StorePrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetStorePrivateKeyOk() (*bool, bool) {
	if o == nil || o.StorePrivateKey == nil {
		return nil, false
	}
	return o.StorePrivateKey, true
}

// HasStorePrivateKey returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasStorePrivateKey() bool {
	if o != nil && o.StorePrivateKey != nil {
		return true
	}

	return false
}

// SetStorePrivateKey gets a reference to the given bool and assigns it to the StorePrivateKey field.
func (o *GatewayUpdateProducerCertificateAutomation) SetStorePrivateKey(v bool) {
	o.StorePrivateKey = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerCertificateAutomation) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerCertificateAutomation) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerCertificateAutomation) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerCertificateAutomation) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerCertificateAutomation) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetVenafiApiKey returns the VenafiApiKey field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiApiKey() string {
	if o == nil || o.VenafiApiKey == nil {
		var ret string
		return ret
	}
	return *o.VenafiApiKey
}

// GetVenafiApiKeyOk returns a tuple with the VenafiApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiApiKeyOk() (*string, bool) {
	if o == nil || o.VenafiApiKey == nil {
		return nil, false
	}
	return o.VenafiApiKey, true
}

// HasVenafiApiKey returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasVenafiApiKey() bool {
	if o != nil && o.VenafiApiKey != nil {
		return true
	}

	return false
}

// SetVenafiApiKey gets a reference to the given string and assigns it to the VenafiApiKey field.
func (o *GatewayUpdateProducerCertificateAutomation) SetVenafiApiKey(v string) {
	o.VenafiApiKey = &v
}

// GetVenafiBaseurl returns the VenafiBaseurl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiBaseurl() string {
	if o == nil || o.VenafiBaseurl == nil {
		var ret string
		return ret
	}
	return *o.VenafiBaseurl
}

// GetVenafiBaseurlOk returns a tuple with the VenafiBaseurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiBaseurlOk() (*string, bool) {
	if o == nil || o.VenafiBaseurl == nil {
		return nil, false
	}
	return o.VenafiBaseurl, true
}

// HasVenafiBaseurl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasVenafiBaseurl() bool {
	if o != nil && o.VenafiBaseurl != nil {
		return true
	}

	return false
}

// SetVenafiBaseurl gets a reference to the given string and assigns it to the VenafiBaseurl field.
func (o *GatewayUpdateProducerCertificateAutomation) SetVenafiBaseurl(v string) {
	o.VenafiBaseurl = &v
}

// GetVenafiPassword returns the VenafiPassword field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiPassword() string {
	if o == nil || o.VenafiPassword == nil {
		var ret string
		return ret
	}
	return *o.VenafiPassword
}

// GetVenafiPasswordOk returns a tuple with the VenafiPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiPasswordOk() (*string, bool) {
	if o == nil || o.VenafiPassword == nil {
		return nil, false
	}
	return o.VenafiPassword, true
}

// HasVenafiPassword returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasVenafiPassword() bool {
	if o != nil && o.VenafiPassword != nil {
		return true
	}

	return false
}

// SetVenafiPassword gets a reference to the given string and assigns it to the VenafiPassword field.
func (o *GatewayUpdateProducerCertificateAutomation) SetVenafiPassword(v string) {
	o.VenafiPassword = &v
}

// GetVenafiUseTpp returns the VenafiUseTpp field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiUseTpp() bool {
	if o == nil || o.VenafiUseTpp == nil {
		var ret bool
		return ret
	}
	return *o.VenafiUseTpp
}

// GetVenafiUseTppOk returns a tuple with the VenafiUseTpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiUseTppOk() (*bool, bool) {
	if o == nil || o.VenafiUseTpp == nil {
		return nil, false
	}
	return o.VenafiUseTpp, true
}

// HasVenafiUseTpp returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasVenafiUseTpp() bool {
	if o != nil && o.VenafiUseTpp != nil {
		return true
	}

	return false
}

// SetVenafiUseTpp gets a reference to the given bool and assigns it to the VenafiUseTpp field.
func (o *GatewayUpdateProducerCertificateAutomation) SetVenafiUseTpp(v bool) {
	o.VenafiUseTpp = &v
}

// GetVenafiUsername returns the VenafiUsername field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiUsername() string {
	if o == nil || o.VenafiUsername == nil {
		var ret string
		return ret
	}
	return *o.VenafiUsername
}

// GetVenafiUsernameOk returns a tuple with the VenafiUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiUsernameOk() (*string, bool) {
	if o == nil || o.VenafiUsername == nil {
		return nil, false
	}
	return o.VenafiUsername, true
}

// HasVenafiUsername returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasVenafiUsername() bool {
	if o != nil && o.VenafiUsername != nil {
		return true
	}

	return false
}

// SetVenafiUsername gets a reference to the given string and assigns it to the VenafiUsername field.
func (o *GatewayUpdateProducerCertificateAutomation) SetVenafiUsername(v string) {
	o.VenafiUsername = &v
}

// GetVenafiZone returns the VenafiZone field value if set, zero value otherwise.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiZone() string {
	if o == nil || o.VenafiZone == nil {
		var ret string
		return ret
	}
	return *o.VenafiZone
}

// GetVenafiZoneOk returns a tuple with the VenafiZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerCertificateAutomation) GetVenafiZoneOk() (*string, bool) {
	if o == nil || o.VenafiZone == nil {
		return nil, false
	}
	return o.VenafiZone, true
}

// HasVenafiZone returns a boolean if a field has been set.
func (o *GatewayUpdateProducerCertificateAutomation) HasVenafiZone() bool {
	if o != nil && o.VenafiZone != nil {
		return true
	}

	return false
}

// SetVenafiZone gets a reference to the given string and assigns it to the VenafiZone field.
func (o *GatewayUpdateProducerCertificateAutomation) SetVenafiZone(v string) {
	o.VenafiZone = &v
}

func (o GatewayUpdateProducerCertificateAutomation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminRotationIntervalDays != nil {
		toSerialize["admin-rotation-interval-days"] = o.AdminRotationIntervalDays
	}
	if o.AllowSubdomains != nil {
		toSerialize["allow-subdomains"] = o.AllowSubdomains
	}
	if o.AllowedDomains != nil {
		toSerialize["allowed-domains"] = o.AllowedDomains
	}
	if o.AutoGeneratedFolder != nil {
		toSerialize["auto-generated-folder"] = o.AutoGeneratedFolder
	}
	if o.EnableAdminRotation != nil {
		toSerialize["enable-admin-rotation"] = o.EnableAdminRotation
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
	if o.RootFirstInChain != nil {
		toSerialize["root-first-in-chain"] = o.RootFirstInChain
	}
	if o.SignUsingAkeylessPki != nil {
		toSerialize["sign-using-akeyless-pki"] = o.SignUsingAkeylessPki
	}
	if o.SignerKeyName != nil {
		toSerialize["signer-key-name"] = o.SignerKeyName
	}
	if o.StorePrivateKey != nil {
		toSerialize["store-private-key"] = o.StorePrivateKey
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
	if o.VenafiApiKey != nil {
		toSerialize["venafi-api-key"] = o.VenafiApiKey
	}
	if o.VenafiBaseurl != nil {
		toSerialize["venafi-baseurl"] = o.VenafiBaseurl
	}
	if o.VenafiPassword != nil {
		toSerialize["venafi-password"] = o.VenafiPassword
	}
	if o.VenafiUseTpp != nil {
		toSerialize["venafi-use-tpp"] = o.VenafiUseTpp
	}
	if o.VenafiUsername != nil {
		toSerialize["venafi-username"] = o.VenafiUsername
	}
	if o.VenafiZone != nil {
		toSerialize["venafi-zone"] = o.VenafiZone
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateProducerCertificateAutomation struct {
	value *GatewayUpdateProducerCertificateAutomation
	isSet bool
}

func (v NullableGatewayUpdateProducerCertificateAutomation) Get() *GatewayUpdateProducerCertificateAutomation {
	return v.value
}

func (v *NullableGatewayUpdateProducerCertificateAutomation) Set(val *GatewayUpdateProducerCertificateAutomation) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerCertificateAutomation) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerCertificateAutomation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerCertificateAutomation(val *GatewayUpdateProducerCertificateAutomation) *NullableGatewayUpdateProducerCertificateAutomation {
	return &NullableGatewayUpdateProducerCertificateAutomation{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerCertificateAutomation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerCertificateAutomation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


