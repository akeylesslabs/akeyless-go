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

// DynamicSecretUpdatePing dynamicSecretUpdatePing is a command that updates Ping dynamic secret
type DynamicSecretUpdatePing struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret new name
	NewName *string `json:"new-name,omitempty"`
	// Ping Federate administrative port
	PingAdministrativePort *string `json:"ping-administrative-port,omitempty"`
	// Set a specific Access Token Management (ATM) instance for the created OAuth Client by providing the ATM Id. If no explicit value is given, the default pingfederate server ATM will be set.
	PingAtmId *string `json:"ping-atm-id,omitempty"`
	// Ping Federate authorization port
	PingAuthorizationPort *string `json:"ping-authorization-port,omitempty"`
	// The subject DN of the client certificate. If no explicit value is given, the producer will create CA certificate and matched client certificate and return it as value. Used in conjunction with ping-issuer-dn (relevant for CLIENT_TLS_CERTIFICATE authentication method)
	PingCertSubjectDn *string `json:"ping-cert-subject-dn,omitempty"`
	// OAuth Client Authentication Type [CLIENT_SECRET, PRIVATE_KEY_JWT, CLIENT_TLS_CERTIFICATE]
	PingClientAuthenticationType *string `json:"ping-client-authentication-type,omitempty"`
	// Determines whether PingFederate requires a unique signed JWT from the client for each action (relevant for PRIVATE_KEY_JWT authentication method) [true/false]
	PingEnforceReplayPrevention *string `json:"ping-enforce-replay-prevention,omitempty"`
	// List of OAuth client grant types [IMPLICIT, AUTHORIZATION_CODE, CLIENT_CREDENTIALS, TOKEN_EXCHANGE, REFRESH_TOKEN, ASSERTION_GRANTS, PASSWORD, RESOURCE_OWNER_CREDENTIALS]. If no explicit value is given, AUTHORIZATION_CODE will be selected as default.
	PingGrantTypes *[]string `json:"ping-grant-types,omitempty"`
	// Issuer DN of trusted CA certificate that imported into Ping Federate server. You may select \\\"Trust Any\\\" to trust all the existing issuers in Ping Federate server. Used in conjunction with ping-cert-subject-dn (relevant for CLIENT_TLS_CERTIFICATE authentication method)
	PingIssuerDn *string `json:"ping-issuer-dn,omitempty"`
	// Base64-encoded JSON Web Key Set (JWKS). If no explicit value is given, the producer will create JWKs and matched signed JWT (Sign Algo: RS256) and return it as value (relevant for PRIVATE_KEY_JWT authentication method)
	PingJwks *string `json:"ping-jwks,omitempty"`
	// The URL of the JSON Web Key Set (JWKS). If no explicit value is given, the producer will create JWKs and matched signed JWT and return it as value (relevant for PRIVATE_KEY_JWT authentication method)
	PingJwksUrl *string `json:"ping-jwks-url,omitempty"`
	// Ping Federate privileged user password
	PingPassword *string `json:"ping-password,omitempty"`
	// Ping Federate privileged user
	PingPrivilegedUser *string `json:"ping-privileged-user,omitempty"`
	// List of URIs to which the OAuth authorization server may redirect the resource owner's user agent after authorization is obtained. At least one redirection URI is required for the AUTHORIZATION_CODE and IMPLICIT grant types.
	PingRedirectUris *[]string `json:"ping-redirect-uris,omitempty"`
	// Limit the OAuth client to specific scopes list
	PingRestrictedScopes *[]string `json:"ping-restricted-scopes,omitempty"`
	// The signing algorithm that the client must use to sign its request objects [RS256,RS384,RS512,ES256,ES384,ES512,PS256,PS384,PS512] If no explicit value is given, the client can use any of the supported signing algorithms (relevant for PRIVATE_KEY_JWT authentication method)
	PingSigningAlgo *string `json:"ping-signing-algo,omitempty"`
	// Ping URL
	PingUrl *string `json:"ping-url,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Add tags attached to this object
	Tags *[]string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The time from dynamic secret creation to expiration.
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewDynamicSecretUpdatePing instantiates a new DynamicSecretUpdatePing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdatePing(name string, ) *DynamicSecretUpdatePing {
	this := DynamicSecretUpdatePing{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var pingAdministrativePort string = "9999"
	this.PingAdministrativePort = &pingAdministrativePort
	var pingAuthorizationPort string = "9031"
	this.PingAuthorizationPort = &pingAuthorizationPort
	var pingClientAuthenticationType string = "CLIENT_SECRET"
	this.PingClientAuthenticationType = &pingClientAuthenticationType
	var pingEnforceReplayPrevention string = "false"
	this.PingEnforceReplayPrevention = &pingEnforceReplayPrevention
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretUpdatePingWithDefaults instantiates a new DynamicSecretUpdatePing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdatePingWithDefaults() *DynamicSecretUpdatePing {
	this := DynamicSecretUpdatePing{}
	var json bool = false
	this.Json = &json
	var pingAdministrativePort string = "9999"
	this.PingAdministrativePort = &pingAdministrativePort
	var pingAuthorizationPort string = "9031"
	this.PingAuthorizationPort = &pingAuthorizationPort
	var pingClientAuthenticationType string = "CLIENT_SECRET"
	this.PingClientAuthenticationType = &pingClientAuthenticationType
	var pingEnforceReplayPrevention string = "false"
	this.PingEnforceReplayPrevention = &pingEnforceReplayPrevention
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretUpdatePing) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretUpdatePing) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretUpdatePing) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretUpdatePing) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DynamicSecretUpdatePing) SetNewName(v string) {
	o.NewName = &v
}

// GetPingAdministrativePort returns the PingAdministrativePort field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingAdministrativePort() string {
	if o == nil || o.PingAdministrativePort == nil {
		var ret string
		return ret
	}
	return *o.PingAdministrativePort
}

// GetPingAdministrativePortOk returns a tuple with the PingAdministrativePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingAdministrativePortOk() (*string, bool) {
	if o == nil || o.PingAdministrativePort == nil {
		return nil, false
	}
	return o.PingAdministrativePort, true
}

// HasPingAdministrativePort returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingAdministrativePort() bool {
	if o != nil && o.PingAdministrativePort != nil {
		return true
	}

	return false
}

// SetPingAdministrativePort gets a reference to the given string and assigns it to the PingAdministrativePort field.
func (o *DynamicSecretUpdatePing) SetPingAdministrativePort(v string) {
	o.PingAdministrativePort = &v
}

// GetPingAtmId returns the PingAtmId field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingAtmId() string {
	if o == nil || o.PingAtmId == nil {
		var ret string
		return ret
	}
	return *o.PingAtmId
}

// GetPingAtmIdOk returns a tuple with the PingAtmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingAtmIdOk() (*string, bool) {
	if o == nil || o.PingAtmId == nil {
		return nil, false
	}
	return o.PingAtmId, true
}

// HasPingAtmId returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingAtmId() bool {
	if o != nil && o.PingAtmId != nil {
		return true
	}

	return false
}

// SetPingAtmId gets a reference to the given string and assigns it to the PingAtmId field.
func (o *DynamicSecretUpdatePing) SetPingAtmId(v string) {
	o.PingAtmId = &v
}

// GetPingAuthorizationPort returns the PingAuthorizationPort field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingAuthorizationPort() string {
	if o == nil || o.PingAuthorizationPort == nil {
		var ret string
		return ret
	}
	return *o.PingAuthorizationPort
}

// GetPingAuthorizationPortOk returns a tuple with the PingAuthorizationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingAuthorizationPortOk() (*string, bool) {
	if o == nil || o.PingAuthorizationPort == nil {
		return nil, false
	}
	return o.PingAuthorizationPort, true
}

// HasPingAuthorizationPort returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingAuthorizationPort() bool {
	if o != nil && o.PingAuthorizationPort != nil {
		return true
	}

	return false
}

// SetPingAuthorizationPort gets a reference to the given string and assigns it to the PingAuthorizationPort field.
func (o *DynamicSecretUpdatePing) SetPingAuthorizationPort(v string) {
	o.PingAuthorizationPort = &v
}

// GetPingCertSubjectDn returns the PingCertSubjectDn field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingCertSubjectDn() string {
	if o == nil || o.PingCertSubjectDn == nil {
		var ret string
		return ret
	}
	return *o.PingCertSubjectDn
}

// GetPingCertSubjectDnOk returns a tuple with the PingCertSubjectDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingCertSubjectDnOk() (*string, bool) {
	if o == nil || o.PingCertSubjectDn == nil {
		return nil, false
	}
	return o.PingCertSubjectDn, true
}

// HasPingCertSubjectDn returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingCertSubjectDn() bool {
	if o != nil && o.PingCertSubjectDn != nil {
		return true
	}

	return false
}

// SetPingCertSubjectDn gets a reference to the given string and assigns it to the PingCertSubjectDn field.
func (o *DynamicSecretUpdatePing) SetPingCertSubjectDn(v string) {
	o.PingCertSubjectDn = &v
}

// GetPingClientAuthenticationType returns the PingClientAuthenticationType field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingClientAuthenticationType() string {
	if o == nil || o.PingClientAuthenticationType == nil {
		var ret string
		return ret
	}
	return *o.PingClientAuthenticationType
}

// GetPingClientAuthenticationTypeOk returns a tuple with the PingClientAuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingClientAuthenticationTypeOk() (*string, bool) {
	if o == nil || o.PingClientAuthenticationType == nil {
		return nil, false
	}
	return o.PingClientAuthenticationType, true
}

// HasPingClientAuthenticationType returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingClientAuthenticationType() bool {
	if o != nil && o.PingClientAuthenticationType != nil {
		return true
	}

	return false
}

// SetPingClientAuthenticationType gets a reference to the given string and assigns it to the PingClientAuthenticationType field.
func (o *DynamicSecretUpdatePing) SetPingClientAuthenticationType(v string) {
	o.PingClientAuthenticationType = &v
}

// GetPingEnforceReplayPrevention returns the PingEnforceReplayPrevention field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingEnforceReplayPrevention() string {
	if o == nil || o.PingEnforceReplayPrevention == nil {
		var ret string
		return ret
	}
	return *o.PingEnforceReplayPrevention
}

// GetPingEnforceReplayPreventionOk returns a tuple with the PingEnforceReplayPrevention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingEnforceReplayPreventionOk() (*string, bool) {
	if o == nil || o.PingEnforceReplayPrevention == nil {
		return nil, false
	}
	return o.PingEnforceReplayPrevention, true
}

// HasPingEnforceReplayPrevention returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingEnforceReplayPrevention() bool {
	if o != nil && o.PingEnforceReplayPrevention != nil {
		return true
	}

	return false
}

// SetPingEnforceReplayPrevention gets a reference to the given string and assigns it to the PingEnforceReplayPrevention field.
func (o *DynamicSecretUpdatePing) SetPingEnforceReplayPrevention(v string) {
	o.PingEnforceReplayPrevention = &v
}

// GetPingGrantTypes returns the PingGrantTypes field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingGrantTypes() []string {
	if o == nil || o.PingGrantTypes == nil {
		var ret []string
		return ret
	}
	return *o.PingGrantTypes
}

// GetPingGrantTypesOk returns a tuple with the PingGrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingGrantTypesOk() (*[]string, bool) {
	if o == nil || o.PingGrantTypes == nil {
		return nil, false
	}
	return o.PingGrantTypes, true
}

// HasPingGrantTypes returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingGrantTypes() bool {
	if o != nil && o.PingGrantTypes != nil {
		return true
	}

	return false
}

// SetPingGrantTypes gets a reference to the given []string and assigns it to the PingGrantTypes field.
func (o *DynamicSecretUpdatePing) SetPingGrantTypes(v []string) {
	o.PingGrantTypes = &v
}

// GetPingIssuerDn returns the PingIssuerDn field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingIssuerDn() string {
	if o == nil || o.PingIssuerDn == nil {
		var ret string
		return ret
	}
	return *o.PingIssuerDn
}

// GetPingIssuerDnOk returns a tuple with the PingIssuerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingIssuerDnOk() (*string, bool) {
	if o == nil || o.PingIssuerDn == nil {
		return nil, false
	}
	return o.PingIssuerDn, true
}

// HasPingIssuerDn returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingIssuerDn() bool {
	if o != nil && o.PingIssuerDn != nil {
		return true
	}

	return false
}

// SetPingIssuerDn gets a reference to the given string and assigns it to the PingIssuerDn field.
func (o *DynamicSecretUpdatePing) SetPingIssuerDn(v string) {
	o.PingIssuerDn = &v
}

// GetPingJwks returns the PingJwks field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingJwks() string {
	if o == nil || o.PingJwks == nil {
		var ret string
		return ret
	}
	return *o.PingJwks
}

// GetPingJwksOk returns a tuple with the PingJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingJwksOk() (*string, bool) {
	if o == nil || o.PingJwks == nil {
		return nil, false
	}
	return o.PingJwks, true
}

// HasPingJwks returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingJwks() bool {
	if o != nil && o.PingJwks != nil {
		return true
	}

	return false
}

// SetPingJwks gets a reference to the given string and assigns it to the PingJwks field.
func (o *DynamicSecretUpdatePing) SetPingJwks(v string) {
	o.PingJwks = &v
}

// GetPingJwksUrl returns the PingJwksUrl field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingJwksUrl() string {
	if o == nil || o.PingJwksUrl == nil {
		var ret string
		return ret
	}
	return *o.PingJwksUrl
}

// GetPingJwksUrlOk returns a tuple with the PingJwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingJwksUrlOk() (*string, bool) {
	if o == nil || o.PingJwksUrl == nil {
		return nil, false
	}
	return o.PingJwksUrl, true
}

// HasPingJwksUrl returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingJwksUrl() bool {
	if o != nil && o.PingJwksUrl != nil {
		return true
	}

	return false
}

// SetPingJwksUrl gets a reference to the given string and assigns it to the PingJwksUrl field.
func (o *DynamicSecretUpdatePing) SetPingJwksUrl(v string) {
	o.PingJwksUrl = &v
}

// GetPingPassword returns the PingPassword field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingPassword() string {
	if o == nil || o.PingPassword == nil {
		var ret string
		return ret
	}
	return *o.PingPassword
}

// GetPingPasswordOk returns a tuple with the PingPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingPasswordOk() (*string, bool) {
	if o == nil || o.PingPassword == nil {
		return nil, false
	}
	return o.PingPassword, true
}

// HasPingPassword returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingPassword() bool {
	if o != nil && o.PingPassword != nil {
		return true
	}

	return false
}

// SetPingPassword gets a reference to the given string and assigns it to the PingPassword field.
func (o *DynamicSecretUpdatePing) SetPingPassword(v string) {
	o.PingPassword = &v
}

// GetPingPrivilegedUser returns the PingPrivilegedUser field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingPrivilegedUser() string {
	if o == nil || o.PingPrivilegedUser == nil {
		var ret string
		return ret
	}
	return *o.PingPrivilegedUser
}

// GetPingPrivilegedUserOk returns a tuple with the PingPrivilegedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingPrivilegedUserOk() (*string, bool) {
	if o == nil || o.PingPrivilegedUser == nil {
		return nil, false
	}
	return o.PingPrivilegedUser, true
}

// HasPingPrivilegedUser returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingPrivilegedUser() bool {
	if o != nil && o.PingPrivilegedUser != nil {
		return true
	}

	return false
}

// SetPingPrivilegedUser gets a reference to the given string and assigns it to the PingPrivilegedUser field.
func (o *DynamicSecretUpdatePing) SetPingPrivilegedUser(v string) {
	o.PingPrivilegedUser = &v
}

// GetPingRedirectUris returns the PingRedirectUris field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingRedirectUris() []string {
	if o == nil || o.PingRedirectUris == nil {
		var ret []string
		return ret
	}
	return *o.PingRedirectUris
}

// GetPingRedirectUrisOk returns a tuple with the PingRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingRedirectUrisOk() (*[]string, bool) {
	if o == nil || o.PingRedirectUris == nil {
		return nil, false
	}
	return o.PingRedirectUris, true
}

// HasPingRedirectUris returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingRedirectUris() bool {
	if o != nil && o.PingRedirectUris != nil {
		return true
	}

	return false
}

// SetPingRedirectUris gets a reference to the given []string and assigns it to the PingRedirectUris field.
func (o *DynamicSecretUpdatePing) SetPingRedirectUris(v []string) {
	o.PingRedirectUris = &v
}

// GetPingRestrictedScopes returns the PingRestrictedScopes field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingRestrictedScopes() []string {
	if o == nil || o.PingRestrictedScopes == nil {
		var ret []string
		return ret
	}
	return *o.PingRestrictedScopes
}

// GetPingRestrictedScopesOk returns a tuple with the PingRestrictedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingRestrictedScopesOk() (*[]string, bool) {
	if o == nil || o.PingRestrictedScopes == nil {
		return nil, false
	}
	return o.PingRestrictedScopes, true
}

// HasPingRestrictedScopes returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingRestrictedScopes() bool {
	if o != nil && o.PingRestrictedScopes != nil {
		return true
	}

	return false
}

// SetPingRestrictedScopes gets a reference to the given []string and assigns it to the PingRestrictedScopes field.
func (o *DynamicSecretUpdatePing) SetPingRestrictedScopes(v []string) {
	o.PingRestrictedScopes = &v
}

// GetPingSigningAlgo returns the PingSigningAlgo field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingSigningAlgo() string {
	if o == nil || o.PingSigningAlgo == nil {
		var ret string
		return ret
	}
	return *o.PingSigningAlgo
}

// GetPingSigningAlgoOk returns a tuple with the PingSigningAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingSigningAlgoOk() (*string, bool) {
	if o == nil || o.PingSigningAlgo == nil {
		return nil, false
	}
	return o.PingSigningAlgo, true
}

// HasPingSigningAlgo returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingSigningAlgo() bool {
	if o != nil && o.PingSigningAlgo != nil {
		return true
	}

	return false
}

// SetPingSigningAlgo gets a reference to the given string and assigns it to the PingSigningAlgo field.
func (o *DynamicSecretUpdatePing) SetPingSigningAlgo(v string) {
	o.PingSigningAlgo = &v
}

// GetPingUrl returns the PingUrl field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetPingUrl() string {
	if o == nil || o.PingUrl == nil {
		var ret string
		return ret
	}
	return *o.PingUrl
}

// GetPingUrlOk returns a tuple with the PingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetPingUrlOk() (*string, bool) {
	if o == nil || o.PingUrl == nil {
		return nil, false
	}
	return o.PingUrl, true
}

// HasPingUrl returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasPingUrl() bool {
	if o != nil && o.PingUrl != nil {
		return true
	}

	return false
}

// SetPingUrl gets a reference to the given string and assigns it to the PingUrl field.
func (o *DynamicSecretUpdatePing) SetPingUrl(v string) {
	o.PingUrl = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretUpdatePing) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretUpdatePing) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretUpdatePing) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretUpdatePing) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretUpdatePing) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretUpdatePing) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdatePing) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretUpdatePing) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretUpdatePing) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretUpdatePing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
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
	if o.PingAdministrativePort != nil {
		toSerialize["ping-administrative-port"] = o.PingAdministrativePort
	}
	if o.PingAtmId != nil {
		toSerialize["ping-atm-id"] = o.PingAtmId
	}
	if o.PingAuthorizationPort != nil {
		toSerialize["ping-authorization-port"] = o.PingAuthorizationPort
	}
	if o.PingCertSubjectDn != nil {
		toSerialize["ping-cert-subject-dn"] = o.PingCertSubjectDn
	}
	if o.PingClientAuthenticationType != nil {
		toSerialize["ping-client-authentication-type"] = o.PingClientAuthenticationType
	}
	if o.PingEnforceReplayPrevention != nil {
		toSerialize["ping-enforce-replay-prevention"] = o.PingEnforceReplayPrevention
	}
	if o.PingGrantTypes != nil {
		toSerialize["ping-grant-types"] = o.PingGrantTypes
	}
	if o.PingIssuerDn != nil {
		toSerialize["ping-issuer-dn"] = o.PingIssuerDn
	}
	if o.PingJwks != nil {
		toSerialize["ping-jwks"] = o.PingJwks
	}
	if o.PingJwksUrl != nil {
		toSerialize["ping-jwks-url"] = o.PingJwksUrl
	}
	if o.PingPassword != nil {
		toSerialize["ping-password"] = o.PingPassword
	}
	if o.PingPrivilegedUser != nil {
		toSerialize["ping-privileged-user"] = o.PingPrivilegedUser
	}
	if o.PingRedirectUris != nil {
		toSerialize["ping-redirect-uris"] = o.PingRedirectUris
	}
	if o.PingRestrictedScopes != nil {
		toSerialize["ping-restricted-scopes"] = o.PingRestrictedScopes
	}
	if o.PingSigningAlgo != nil {
		toSerialize["ping-signing-algo"] = o.PingSigningAlgo
	}
	if o.PingUrl != nil {
		toSerialize["ping-url"] = o.PingUrl
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
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

type NullableDynamicSecretUpdatePing struct {
	value *DynamicSecretUpdatePing
	isSet bool
}

func (v NullableDynamicSecretUpdatePing) Get() *DynamicSecretUpdatePing {
	return v.value
}

func (v *NullableDynamicSecretUpdatePing) Set(val *DynamicSecretUpdatePing) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdatePing) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdatePing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdatePing(val *DynamicSecretUpdatePing) *NullableDynamicSecretUpdatePing {
	return &NullableDynamicSecretUpdatePing{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdatePing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdatePing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


