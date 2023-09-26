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

// DsCreateCertificateAutomation dsCreateCertificateAutomation is a command that creates a Certificate Automation dynamic secret to dynamically create certificates generated by Venafi or have Akeyless generated certificates using PKI be monitored by Venafi
type DsCreateCertificateAutomation struct {
	// Admin credentials rotation interval (days)
	AdminRotationIntervalDays *int64 `json:"admin-rotation-interval-days,omitempty"`
	// Allow subdomains
	AllowSubdomains *bool `json:"allow-subdomains,omitempty"`
	// Allowed domains
	AllowedDomains *[]string `json:"allowed-domains,omitempty"`
	// Auto generated folder
	AutoGeneratedFolder *string `json:"auto-generated-folder,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Automatic admin credentials rotation
	EnableAdminRotation *bool `json:"enable-admin-rotation,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
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
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL in time.Duration format (2160h / 129600m / etc...). When using sign-using-akeyless-pki certificates created will have this validity period, otherwise the user-ttl is taken from the Validity Period field of the Zone's' Issuing Template. When using cert-manager it is advised to have a TTL of above 60 days (1440h). For more information - https://cert-manager.io/docs/usage/certificate/
	UserTtl *string `json:"user-ttl,omitempty"`
	// Venafi Access Token to use to access the TPP environment (Relevant when using TPP)
	VenafiAccessToken *string `json:"venafi-access-token,omitempty"`
	// Venafi API key
	VenafiApiKey *string `json:"venafi-api-key,omitempty"`
	// Venafi Baseurl
	VenafiBaseurl *string `json:"venafi-baseurl,omitempty"`
	// Venafi Client ID that was used when the access token was generated
	VenafiClientId *string `json:"venafi-client-id,omitempty"`
	// Venafi Refresh Token to use when the Access Token is expired (Relevant when using TPP)
	VenafiRefreshToken *string `json:"venafi-refresh-token,omitempty"`
	// Venafi using TPP
	VenafiUseTpp *bool `json:"venafi-use-tpp,omitempty"`
	// Venafi Zone
	VenafiZone *string `json:"venafi-zone,omitempty"`
}

// NewDsCreateCertificateAutomation instantiates a new DsCreateCertificateAutomation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDsCreateCertificateAutomation(name string, ) *DsCreateCertificateAutomation {
	this := DsCreateCertificateAutomation{}
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "2160h"
	this.UserTtl = &userTtl
	var venafiClientId string = "akeyless"
	this.VenafiClientId = &venafiClientId
	return &this
}

// NewDsCreateCertificateAutomationWithDefaults instantiates a new DsCreateCertificateAutomation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDsCreateCertificateAutomationWithDefaults() *DsCreateCertificateAutomation {
	this := DsCreateCertificateAutomation{}
	var adminRotationIntervalDays int64 = 0
	this.AdminRotationIntervalDays = &adminRotationIntervalDays
	var enableAdminRotation bool = false
	this.EnableAdminRotation = &enableAdminRotation
	var json bool = false
	this.Json = &json
	var userTtl string = "2160h"
	this.UserTtl = &userTtl
	var venafiClientId string = "akeyless"
	this.VenafiClientId = &venafiClientId
	return &this
}

// GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetAdminRotationIntervalDays() int64 {
	if o == nil || o.AdminRotationIntervalDays == nil {
		var ret int64
		return ret
	}
	return *o.AdminRotationIntervalDays
}

// GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetAdminRotationIntervalDaysOk() (*int64, bool) {
	if o == nil || o.AdminRotationIntervalDays == nil {
		return nil, false
	}
	return o.AdminRotationIntervalDays, true
}

// HasAdminRotationIntervalDays returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasAdminRotationIntervalDays() bool {
	if o != nil && o.AdminRotationIntervalDays != nil {
		return true
	}

	return false
}

// SetAdminRotationIntervalDays gets a reference to the given int64 and assigns it to the AdminRotationIntervalDays field.
func (o *DsCreateCertificateAutomation) SetAdminRotationIntervalDays(v int64) {
	o.AdminRotationIntervalDays = &v
}

// GetAllowSubdomains returns the AllowSubdomains field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetAllowSubdomains() bool {
	if o == nil || o.AllowSubdomains == nil {
		var ret bool
		return ret
	}
	return *o.AllowSubdomains
}

// GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetAllowSubdomainsOk() (*bool, bool) {
	if o == nil || o.AllowSubdomains == nil {
		return nil, false
	}
	return o.AllowSubdomains, true
}

// HasAllowSubdomains returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasAllowSubdomains() bool {
	if o != nil && o.AllowSubdomains != nil {
		return true
	}

	return false
}

// SetAllowSubdomains gets a reference to the given bool and assigns it to the AllowSubdomains field.
func (o *DsCreateCertificateAutomation) SetAllowSubdomains(v bool) {
	o.AllowSubdomains = &v
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetAllowedDomains() []string {
	if o == nil || o.AllowedDomains == nil {
		var ret []string
		return ret
	}
	return *o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetAllowedDomainsOk() (*[]string, bool) {
	if o == nil || o.AllowedDomains == nil {
		return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasAllowedDomains() bool {
	if o != nil && o.AllowedDomains != nil {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given []string and assigns it to the AllowedDomains field.
func (o *DsCreateCertificateAutomation) SetAllowedDomains(v []string) {
	o.AllowedDomains = &v
}

// GetAutoGeneratedFolder returns the AutoGeneratedFolder field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetAutoGeneratedFolder() string {
	if o == nil || o.AutoGeneratedFolder == nil {
		var ret string
		return ret
	}
	return *o.AutoGeneratedFolder
}

// GetAutoGeneratedFolderOk returns a tuple with the AutoGeneratedFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetAutoGeneratedFolderOk() (*string, bool) {
	if o == nil || o.AutoGeneratedFolder == nil {
		return nil, false
	}
	return o.AutoGeneratedFolder, true
}

// HasAutoGeneratedFolder returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasAutoGeneratedFolder() bool {
	if o != nil && o.AutoGeneratedFolder != nil {
		return true
	}

	return false
}

// SetAutoGeneratedFolder gets a reference to the given string and assigns it to the AutoGeneratedFolder field.
func (o *DsCreateCertificateAutomation) SetAutoGeneratedFolder(v string) {
	o.AutoGeneratedFolder = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DsCreateCertificateAutomation) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetEnableAdminRotation returns the EnableAdminRotation field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetEnableAdminRotation() bool {
	if o == nil || o.EnableAdminRotation == nil {
		var ret bool
		return ret
	}
	return *o.EnableAdminRotation
}

// GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetEnableAdminRotationOk() (*bool, bool) {
	if o == nil || o.EnableAdminRotation == nil {
		return nil, false
	}
	return o.EnableAdminRotation, true
}

// HasEnableAdminRotation returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasEnableAdminRotation() bool {
	if o != nil && o.EnableAdminRotation != nil {
		return true
	}

	return false
}

// SetEnableAdminRotation gets a reference to the given bool and assigns it to the EnableAdminRotation field.
func (o *DsCreateCertificateAutomation) SetEnableAdminRotation(v bool) {
	o.EnableAdminRotation = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DsCreateCertificateAutomation) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DsCreateCertificateAutomation) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DsCreateCertificateAutomation) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DsCreateCertificateAutomation) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRootFirstInChain returns the RootFirstInChain field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetRootFirstInChain() bool {
	if o == nil || o.RootFirstInChain == nil {
		var ret bool
		return ret
	}
	return *o.RootFirstInChain
}

// GetRootFirstInChainOk returns a tuple with the RootFirstInChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetRootFirstInChainOk() (*bool, bool) {
	if o == nil || o.RootFirstInChain == nil {
		return nil, false
	}
	return o.RootFirstInChain, true
}

// HasRootFirstInChain returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasRootFirstInChain() bool {
	if o != nil && o.RootFirstInChain != nil {
		return true
	}

	return false
}

// SetRootFirstInChain gets a reference to the given bool and assigns it to the RootFirstInChain field.
func (o *DsCreateCertificateAutomation) SetRootFirstInChain(v bool) {
	o.RootFirstInChain = &v
}

// GetSignUsingAkeylessPki returns the SignUsingAkeylessPki field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetSignUsingAkeylessPki() bool {
	if o == nil || o.SignUsingAkeylessPki == nil {
		var ret bool
		return ret
	}
	return *o.SignUsingAkeylessPki
}

// GetSignUsingAkeylessPkiOk returns a tuple with the SignUsingAkeylessPki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetSignUsingAkeylessPkiOk() (*bool, bool) {
	if o == nil || o.SignUsingAkeylessPki == nil {
		return nil, false
	}
	return o.SignUsingAkeylessPki, true
}

// HasSignUsingAkeylessPki returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasSignUsingAkeylessPki() bool {
	if o != nil && o.SignUsingAkeylessPki != nil {
		return true
	}

	return false
}

// SetSignUsingAkeylessPki gets a reference to the given bool and assigns it to the SignUsingAkeylessPki field.
func (o *DsCreateCertificateAutomation) SetSignUsingAkeylessPki(v bool) {
	o.SignUsingAkeylessPki = &v
}

// GetSignerKeyName returns the SignerKeyName field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetSignerKeyName() string {
	if o == nil || o.SignerKeyName == nil {
		var ret string
		return ret
	}
	return *o.SignerKeyName
}

// GetSignerKeyNameOk returns a tuple with the SignerKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetSignerKeyNameOk() (*string, bool) {
	if o == nil || o.SignerKeyName == nil {
		return nil, false
	}
	return o.SignerKeyName, true
}

// HasSignerKeyName returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasSignerKeyName() bool {
	if o != nil && o.SignerKeyName != nil {
		return true
	}

	return false
}

// SetSignerKeyName gets a reference to the given string and assigns it to the SignerKeyName field.
func (o *DsCreateCertificateAutomation) SetSignerKeyName(v string) {
	o.SignerKeyName = &v
}

// GetStorePrivateKey returns the StorePrivateKey field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetStorePrivateKey() bool {
	if o == nil || o.StorePrivateKey == nil {
		var ret bool
		return ret
	}
	return *o.StorePrivateKey
}

// GetStorePrivateKeyOk returns a tuple with the StorePrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetStorePrivateKeyOk() (*bool, bool) {
	if o == nil || o.StorePrivateKey == nil {
		return nil, false
	}
	return o.StorePrivateKey, true
}

// HasStorePrivateKey returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasStorePrivateKey() bool {
	if o != nil && o.StorePrivateKey != nil {
		return true
	}

	return false
}

// SetStorePrivateKey gets a reference to the given bool and assigns it to the StorePrivateKey field.
func (o *DsCreateCertificateAutomation) SetStorePrivateKey(v bool) {
	o.StorePrivateKey = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DsCreateCertificateAutomation) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DsCreateCertificateAutomation) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DsCreateCertificateAutomation) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DsCreateCertificateAutomation) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DsCreateCertificateAutomation) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetVenafiAccessToken returns the VenafiAccessToken field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetVenafiAccessToken() string {
	if o == nil || o.VenafiAccessToken == nil {
		var ret string
		return ret
	}
	return *o.VenafiAccessToken
}

// GetVenafiAccessTokenOk returns a tuple with the VenafiAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetVenafiAccessTokenOk() (*string, bool) {
	if o == nil || o.VenafiAccessToken == nil {
		return nil, false
	}
	return o.VenafiAccessToken, true
}

// HasVenafiAccessToken returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasVenafiAccessToken() bool {
	if o != nil && o.VenafiAccessToken != nil {
		return true
	}

	return false
}

// SetVenafiAccessToken gets a reference to the given string and assigns it to the VenafiAccessToken field.
func (o *DsCreateCertificateAutomation) SetVenafiAccessToken(v string) {
	o.VenafiAccessToken = &v
}

// GetVenafiApiKey returns the VenafiApiKey field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetVenafiApiKey() string {
	if o == nil || o.VenafiApiKey == nil {
		var ret string
		return ret
	}
	return *o.VenafiApiKey
}

// GetVenafiApiKeyOk returns a tuple with the VenafiApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetVenafiApiKeyOk() (*string, bool) {
	if o == nil || o.VenafiApiKey == nil {
		return nil, false
	}
	return o.VenafiApiKey, true
}

// HasVenafiApiKey returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasVenafiApiKey() bool {
	if o != nil && o.VenafiApiKey != nil {
		return true
	}

	return false
}

// SetVenafiApiKey gets a reference to the given string and assigns it to the VenafiApiKey field.
func (o *DsCreateCertificateAutomation) SetVenafiApiKey(v string) {
	o.VenafiApiKey = &v
}

// GetVenafiBaseurl returns the VenafiBaseurl field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetVenafiBaseurl() string {
	if o == nil || o.VenafiBaseurl == nil {
		var ret string
		return ret
	}
	return *o.VenafiBaseurl
}

// GetVenafiBaseurlOk returns a tuple with the VenafiBaseurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetVenafiBaseurlOk() (*string, bool) {
	if o == nil || o.VenafiBaseurl == nil {
		return nil, false
	}
	return o.VenafiBaseurl, true
}

// HasVenafiBaseurl returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasVenafiBaseurl() bool {
	if o != nil && o.VenafiBaseurl != nil {
		return true
	}

	return false
}

// SetVenafiBaseurl gets a reference to the given string and assigns it to the VenafiBaseurl field.
func (o *DsCreateCertificateAutomation) SetVenafiBaseurl(v string) {
	o.VenafiBaseurl = &v
}

// GetVenafiClientId returns the VenafiClientId field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetVenafiClientId() string {
	if o == nil || o.VenafiClientId == nil {
		var ret string
		return ret
	}
	return *o.VenafiClientId
}

// GetVenafiClientIdOk returns a tuple with the VenafiClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetVenafiClientIdOk() (*string, bool) {
	if o == nil || o.VenafiClientId == nil {
		return nil, false
	}
	return o.VenafiClientId, true
}

// HasVenafiClientId returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasVenafiClientId() bool {
	if o != nil && o.VenafiClientId != nil {
		return true
	}

	return false
}

// SetVenafiClientId gets a reference to the given string and assigns it to the VenafiClientId field.
func (o *DsCreateCertificateAutomation) SetVenafiClientId(v string) {
	o.VenafiClientId = &v
}

// GetVenafiRefreshToken returns the VenafiRefreshToken field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetVenafiRefreshToken() string {
	if o == nil || o.VenafiRefreshToken == nil {
		var ret string
		return ret
	}
	return *o.VenafiRefreshToken
}

// GetVenafiRefreshTokenOk returns a tuple with the VenafiRefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetVenafiRefreshTokenOk() (*string, bool) {
	if o == nil || o.VenafiRefreshToken == nil {
		return nil, false
	}
	return o.VenafiRefreshToken, true
}

// HasVenafiRefreshToken returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasVenafiRefreshToken() bool {
	if o != nil && o.VenafiRefreshToken != nil {
		return true
	}

	return false
}

// SetVenafiRefreshToken gets a reference to the given string and assigns it to the VenafiRefreshToken field.
func (o *DsCreateCertificateAutomation) SetVenafiRefreshToken(v string) {
	o.VenafiRefreshToken = &v
}

// GetVenafiUseTpp returns the VenafiUseTpp field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetVenafiUseTpp() bool {
	if o == nil || o.VenafiUseTpp == nil {
		var ret bool
		return ret
	}
	return *o.VenafiUseTpp
}

// GetVenafiUseTppOk returns a tuple with the VenafiUseTpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetVenafiUseTppOk() (*bool, bool) {
	if o == nil || o.VenafiUseTpp == nil {
		return nil, false
	}
	return o.VenafiUseTpp, true
}

// HasVenafiUseTpp returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasVenafiUseTpp() bool {
	if o != nil && o.VenafiUseTpp != nil {
		return true
	}

	return false
}

// SetVenafiUseTpp gets a reference to the given bool and assigns it to the VenafiUseTpp field.
func (o *DsCreateCertificateAutomation) SetVenafiUseTpp(v bool) {
	o.VenafiUseTpp = &v
}

// GetVenafiZone returns the VenafiZone field value if set, zero value otherwise.
func (o *DsCreateCertificateAutomation) GetVenafiZone() string {
	if o == nil || o.VenafiZone == nil {
		var ret string
		return ret
	}
	return *o.VenafiZone
}

// GetVenafiZoneOk returns a tuple with the VenafiZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCertificateAutomation) GetVenafiZoneOk() (*string, bool) {
	if o == nil || o.VenafiZone == nil {
		return nil, false
	}
	return o.VenafiZone, true
}

// HasVenafiZone returns a boolean if a field has been set.
func (o *DsCreateCertificateAutomation) HasVenafiZone() bool {
	if o != nil && o.VenafiZone != nil {
		return true
	}

	return false
}

// SetVenafiZone gets a reference to the given string and assigns it to the VenafiZone field.
func (o *DsCreateCertificateAutomation) SetVenafiZone(v string) {
	o.VenafiZone = &v
}

func (o DsCreateCertificateAutomation) MarshalJSON() ([]byte, error) {
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
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
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
	if o.VenafiAccessToken != nil {
		toSerialize["venafi-access-token"] = o.VenafiAccessToken
	}
	if o.VenafiApiKey != nil {
		toSerialize["venafi-api-key"] = o.VenafiApiKey
	}
	if o.VenafiBaseurl != nil {
		toSerialize["venafi-baseurl"] = o.VenafiBaseurl
	}
	if o.VenafiClientId != nil {
		toSerialize["venafi-client-id"] = o.VenafiClientId
	}
	if o.VenafiRefreshToken != nil {
		toSerialize["venafi-refresh-token"] = o.VenafiRefreshToken
	}
	if o.VenafiUseTpp != nil {
		toSerialize["venafi-use-tpp"] = o.VenafiUseTpp
	}
	if o.VenafiZone != nil {
		toSerialize["venafi-zone"] = o.VenafiZone
	}
	return json.Marshal(toSerialize)
}

type NullableDsCreateCertificateAutomation struct {
	value *DsCreateCertificateAutomation
	isSet bool
}

func (v NullableDsCreateCertificateAutomation) Get() *DsCreateCertificateAutomation {
	return v.value
}

func (v *NullableDsCreateCertificateAutomation) Set(val *DsCreateCertificateAutomation) {
	v.value = val
	v.isSet = true
}

func (v NullableDsCreateCertificateAutomation) IsSet() bool {
	return v.isSet
}

func (v *NullableDsCreateCertificateAutomation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDsCreateCertificateAutomation(val *DsCreateCertificateAutomation) *NullableDsCreateCertificateAutomation {
	return &NullableDsCreateCertificateAutomation{value: val, isSet: true}
}

func (v NullableDsCreateCertificateAutomation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDsCreateCertificateAutomation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


