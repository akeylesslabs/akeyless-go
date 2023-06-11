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

// checks if the GatewayCreateProducerPing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayCreateProducerPing{}

// GatewayCreateProducerPing gatewayCreateProducerPing is a command that creates ping producer
type GatewayCreateProducerPing struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
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
	PingGrantTypes []string `json:"ping-grant-types,omitempty"`
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
	PingRedirectUris []string `json:"ping-redirect-uris,omitempty"`
	// Limit the OAuth client to specific scopes list
	PingRestrictedScopes []string `json:"ping-restricted-scopes,omitempty"`
	// The signing algorithm that the client must use to sign its request objects [RS256,RS384,RS512,ES256,ES384,ES512,PS256,PS384,PS512] If no explicit value is given, the client can use any of the supported signing algorithms (relevant for PRIVATE_KEY_JWT authentication method)
	PingSigningAlgo *string `json:"ping-signing-algo,omitempty"`
	// Ping URL
	PingUrl *string `json:"ping-url,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Add tags attached to this object
	Tags []string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// The time from dynamic secret creation to expiration.
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayCreateProducerPing instantiates a new GatewayCreateProducerPing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerPing(name string) *GatewayCreateProducerPing {
	this := GatewayCreateProducerPing{}
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

// NewGatewayCreateProducerPingWithDefaults instantiates a new GatewayCreateProducerPing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerPingWithDefaults() *GatewayCreateProducerPing {
	this := GatewayCreateProducerPing{}
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
func (o *GatewayCreateProducerPing) GetDeleteProtection() string {
	if o == nil || IsNil(o.DeleteProtection) {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || IsNil(o.DeleteProtection) {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasDeleteProtection() bool {
	if o != nil && !IsNil(o.DeleteProtection) {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerPing) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerPing) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerPing) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerPing) SetName(v string) {
	o.Name = v
}

// GetPingAdministrativePort returns the PingAdministrativePort field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingAdministrativePort() string {
	if o == nil || IsNil(o.PingAdministrativePort) {
		var ret string
		return ret
	}
	return *o.PingAdministrativePort
}

// GetPingAdministrativePortOk returns a tuple with the PingAdministrativePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingAdministrativePortOk() (*string, bool) {
	if o == nil || IsNil(o.PingAdministrativePort) {
		return nil, false
	}
	return o.PingAdministrativePort, true
}

// HasPingAdministrativePort returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingAdministrativePort() bool {
	if o != nil && !IsNil(o.PingAdministrativePort) {
		return true
	}

	return false
}

// SetPingAdministrativePort gets a reference to the given string and assigns it to the PingAdministrativePort field.
func (o *GatewayCreateProducerPing) SetPingAdministrativePort(v string) {
	o.PingAdministrativePort = &v
}

// GetPingAtmId returns the PingAtmId field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingAtmId() string {
	if o == nil || IsNil(o.PingAtmId) {
		var ret string
		return ret
	}
	return *o.PingAtmId
}

// GetPingAtmIdOk returns a tuple with the PingAtmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingAtmIdOk() (*string, bool) {
	if o == nil || IsNil(o.PingAtmId) {
		return nil, false
	}
	return o.PingAtmId, true
}

// HasPingAtmId returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingAtmId() bool {
	if o != nil && !IsNil(o.PingAtmId) {
		return true
	}

	return false
}

// SetPingAtmId gets a reference to the given string and assigns it to the PingAtmId field.
func (o *GatewayCreateProducerPing) SetPingAtmId(v string) {
	o.PingAtmId = &v
}

// GetPingAuthorizationPort returns the PingAuthorizationPort field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingAuthorizationPort() string {
	if o == nil || IsNil(o.PingAuthorizationPort) {
		var ret string
		return ret
	}
	return *o.PingAuthorizationPort
}

// GetPingAuthorizationPortOk returns a tuple with the PingAuthorizationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingAuthorizationPortOk() (*string, bool) {
	if o == nil || IsNil(o.PingAuthorizationPort) {
		return nil, false
	}
	return o.PingAuthorizationPort, true
}

// HasPingAuthorizationPort returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingAuthorizationPort() bool {
	if o != nil && !IsNil(o.PingAuthorizationPort) {
		return true
	}

	return false
}

// SetPingAuthorizationPort gets a reference to the given string and assigns it to the PingAuthorizationPort field.
func (o *GatewayCreateProducerPing) SetPingAuthorizationPort(v string) {
	o.PingAuthorizationPort = &v
}

// GetPingCertSubjectDn returns the PingCertSubjectDn field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingCertSubjectDn() string {
	if o == nil || IsNil(o.PingCertSubjectDn) {
		var ret string
		return ret
	}
	return *o.PingCertSubjectDn
}

// GetPingCertSubjectDnOk returns a tuple with the PingCertSubjectDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingCertSubjectDnOk() (*string, bool) {
	if o == nil || IsNil(o.PingCertSubjectDn) {
		return nil, false
	}
	return o.PingCertSubjectDn, true
}

// HasPingCertSubjectDn returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingCertSubjectDn() bool {
	if o != nil && !IsNil(o.PingCertSubjectDn) {
		return true
	}

	return false
}

// SetPingCertSubjectDn gets a reference to the given string and assigns it to the PingCertSubjectDn field.
func (o *GatewayCreateProducerPing) SetPingCertSubjectDn(v string) {
	o.PingCertSubjectDn = &v
}

// GetPingClientAuthenticationType returns the PingClientAuthenticationType field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingClientAuthenticationType() string {
	if o == nil || IsNil(o.PingClientAuthenticationType) {
		var ret string
		return ret
	}
	return *o.PingClientAuthenticationType
}

// GetPingClientAuthenticationTypeOk returns a tuple with the PingClientAuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingClientAuthenticationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PingClientAuthenticationType) {
		return nil, false
	}
	return o.PingClientAuthenticationType, true
}

// HasPingClientAuthenticationType returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingClientAuthenticationType() bool {
	if o != nil && !IsNil(o.PingClientAuthenticationType) {
		return true
	}

	return false
}

// SetPingClientAuthenticationType gets a reference to the given string and assigns it to the PingClientAuthenticationType field.
func (o *GatewayCreateProducerPing) SetPingClientAuthenticationType(v string) {
	o.PingClientAuthenticationType = &v
}

// GetPingEnforceReplayPrevention returns the PingEnforceReplayPrevention field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingEnforceReplayPrevention() string {
	if o == nil || IsNil(o.PingEnforceReplayPrevention) {
		var ret string
		return ret
	}
	return *o.PingEnforceReplayPrevention
}

// GetPingEnforceReplayPreventionOk returns a tuple with the PingEnforceReplayPrevention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingEnforceReplayPreventionOk() (*string, bool) {
	if o == nil || IsNil(o.PingEnforceReplayPrevention) {
		return nil, false
	}
	return o.PingEnforceReplayPrevention, true
}

// HasPingEnforceReplayPrevention returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingEnforceReplayPrevention() bool {
	if o != nil && !IsNil(o.PingEnforceReplayPrevention) {
		return true
	}

	return false
}

// SetPingEnforceReplayPrevention gets a reference to the given string and assigns it to the PingEnforceReplayPrevention field.
func (o *GatewayCreateProducerPing) SetPingEnforceReplayPrevention(v string) {
	o.PingEnforceReplayPrevention = &v
}

// GetPingGrantTypes returns the PingGrantTypes field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingGrantTypes() []string {
	if o == nil || IsNil(o.PingGrantTypes) {
		var ret []string
		return ret
	}
	return o.PingGrantTypes
}

// GetPingGrantTypesOk returns a tuple with the PingGrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingGrantTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.PingGrantTypes) {
		return nil, false
	}
	return o.PingGrantTypes, true
}

// HasPingGrantTypes returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingGrantTypes() bool {
	if o != nil && !IsNil(o.PingGrantTypes) {
		return true
	}

	return false
}

// SetPingGrantTypes gets a reference to the given []string and assigns it to the PingGrantTypes field.
func (o *GatewayCreateProducerPing) SetPingGrantTypes(v []string) {
	o.PingGrantTypes = v
}

// GetPingIssuerDn returns the PingIssuerDn field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingIssuerDn() string {
	if o == nil || IsNil(o.PingIssuerDn) {
		var ret string
		return ret
	}
	return *o.PingIssuerDn
}

// GetPingIssuerDnOk returns a tuple with the PingIssuerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingIssuerDnOk() (*string, bool) {
	if o == nil || IsNil(o.PingIssuerDn) {
		return nil, false
	}
	return o.PingIssuerDn, true
}

// HasPingIssuerDn returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingIssuerDn() bool {
	if o != nil && !IsNil(o.PingIssuerDn) {
		return true
	}

	return false
}

// SetPingIssuerDn gets a reference to the given string and assigns it to the PingIssuerDn field.
func (o *GatewayCreateProducerPing) SetPingIssuerDn(v string) {
	o.PingIssuerDn = &v
}

// GetPingJwks returns the PingJwks field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingJwks() string {
	if o == nil || IsNil(o.PingJwks) {
		var ret string
		return ret
	}
	return *o.PingJwks
}

// GetPingJwksOk returns a tuple with the PingJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingJwksOk() (*string, bool) {
	if o == nil || IsNil(o.PingJwks) {
		return nil, false
	}
	return o.PingJwks, true
}

// HasPingJwks returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingJwks() bool {
	if o != nil && !IsNil(o.PingJwks) {
		return true
	}

	return false
}

// SetPingJwks gets a reference to the given string and assigns it to the PingJwks field.
func (o *GatewayCreateProducerPing) SetPingJwks(v string) {
	o.PingJwks = &v
}

// GetPingJwksUrl returns the PingJwksUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingJwksUrl() string {
	if o == nil || IsNil(o.PingJwksUrl) {
		var ret string
		return ret
	}
	return *o.PingJwksUrl
}

// GetPingJwksUrlOk returns a tuple with the PingJwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingJwksUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PingJwksUrl) {
		return nil, false
	}
	return o.PingJwksUrl, true
}

// HasPingJwksUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingJwksUrl() bool {
	if o != nil && !IsNil(o.PingJwksUrl) {
		return true
	}

	return false
}

// SetPingJwksUrl gets a reference to the given string and assigns it to the PingJwksUrl field.
func (o *GatewayCreateProducerPing) SetPingJwksUrl(v string) {
	o.PingJwksUrl = &v
}

// GetPingPassword returns the PingPassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingPassword() string {
	if o == nil || IsNil(o.PingPassword) {
		var ret string
		return ret
	}
	return *o.PingPassword
}

// GetPingPasswordOk returns a tuple with the PingPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.PingPassword) {
		return nil, false
	}
	return o.PingPassword, true
}

// HasPingPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingPassword() bool {
	if o != nil && !IsNil(o.PingPassword) {
		return true
	}

	return false
}

// SetPingPassword gets a reference to the given string and assigns it to the PingPassword field.
func (o *GatewayCreateProducerPing) SetPingPassword(v string) {
	o.PingPassword = &v
}

// GetPingPrivilegedUser returns the PingPrivilegedUser field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingPrivilegedUser() string {
	if o == nil || IsNil(o.PingPrivilegedUser) {
		var ret string
		return ret
	}
	return *o.PingPrivilegedUser
}

// GetPingPrivilegedUserOk returns a tuple with the PingPrivilegedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingPrivilegedUserOk() (*string, bool) {
	if o == nil || IsNil(o.PingPrivilegedUser) {
		return nil, false
	}
	return o.PingPrivilegedUser, true
}

// HasPingPrivilegedUser returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingPrivilegedUser() bool {
	if o != nil && !IsNil(o.PingPrivilegedUser) {
		return true
	}

	return false
}

// SetPingPrivilegedUser gets a reference to the given string and assigns it to the PingPrivilegedUser field.
func (o *GatewayCreateProducerPing) SetPingPrivilegedUser(v string) {
	o.PingPrivilegedUser = &v
}

// GetPingRedirectUris returns the PingRedirectUris field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingRedirectUris() []string {
	if o == nil || IsNil(o.PingRedirectUris) {
		var ret []string
		return ret
	}
	return o.PingRedirectUris
}

// GetPingRedirectUrisOk returns a tuple with the PingRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.PingRedirectUris) {
		return nil, false
	}
	return o.PingRedirectUris, true
}

// HasPingRedirectUris returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingRedirectUris() bool {
	if o != nil && !IsNil(o.PingRedirectUris) {
		return true
	}

	return false
}

// SetPingRedirectUris gets a reference to the given []string and assigns it to the PingRedirectUris field.
func (o *GatewayCreateProducerPing) SetPingRedirectUris(v []string) {
	o.PingRedirectUris = v
}

// GetPingRestrictedScopes returns the PingRestrictedScopes field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingRestrictedScopes() []string {
	if o == nil || IsNil(o.PingRestrictedScopes) {
		var ret []string
		return ret
	}
	return o.PingRestrictedScopes
}

// GetPingRestrictedScopesOk returns a tuple with the PingRestrictedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingRestrictedScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.PingRestrictedScopes) {
		return nil, false
	}
	return o.PingRestrictedScopes, true
}

// HasPingRestrictedScopes returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingRestrictedScopes() bool {
	if o != nil && !IsNil(o.PingRestrictedScopes) {
		return true
	}

	return false
}

// SetPingRestrictedScopes gets a reference to the given []string and assigns it to the PingRestrictedScopes field.
func (o *GatewayCreateProducerPing) SetPingRestrictedScopes(v []string) {
	o.PingRestrictedScopes = v
}

// GetPingSigningAlgo returns the PingSigningAlgo field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingSigningAlgo() string {
	if o == nil || IsNil(o.PingSigningAlgo) {
		var ret string
		return ret
	}
	return *o.PingSigningAlgo
}

// GetPingSigningAlgoOk returns a tuple with the PingSigningAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingSigningAlgoOk() (*string, bool) {
	if o == nil || IsNil(o.PingSigningAlgo) {
		return nil, false
	}
	return o.PingSigningAlgo, true
}

// HasPingSigningAlgo returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingSigningAlgo() bool {
	if o != nil && !IsNil(o.PingSigningAlgo) {
		return true
	}

	return false
}

// SetPingSigningAlgo gets a reference to the given string and assigns it to the PingSigningAlgo field.
func (o *GatewayCreateProducerPing) SetPingSigningAlgo(v string) {
	o.PingSigningAlgo = &v
}

// GetPingUrl returns the PingUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetPingUrl() string {
	if o == nil || IsNil(o.PingUrl) {
		var ret string
		return ret
	}
	return *o.PingUrl
}

// GetPingUrlOk returns a tuple with the PingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetPingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PingUrl) {
		return nil, false
	}
	return o.PingUrl, true
}

// HasPingUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasPingUrl() bool {
	if o != nil && !IsNil(o.PingUrl) {
		return true
	}

	return false
}

// SetPingUrl gets a reference to the given string and assigns it to the PingUrl field.
func (o *GatewayCreateProducerPing) SetPingUrl(v string) {
	o.PingUrl = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetProducerEncryptionKeyName() string {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasProducerEncryptionKeyName() bool {
	if o != nil && !IsNil(o.ProducerEncryptionKeyName) {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerPing) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerPing) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerPing) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerPing) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerPing) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerPing) GetUserTtl() string {
	if o == nil || IsNil(o.UserTtl) {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPing) GetUserTtlOk() (*string, bool) {
	if o == nil || IsNil(o.UserTtl) {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerPing) HasUserTtl() bool {
	if o != nil && !IsNil(o.UserTtl) {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerPing) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerPing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayCreateProducerPing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteProtection) {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PingAdministrativePort) {
		toSerialize["ping-administrative-port"] = o.PingAdministrativePort
	}
	if !IsNil(o.PingAtmId) {
		toSerialize["ping-atm-id"] = o.PingAtmId
	}
	if !IsNil(o.PingAuthorizationPort) {
		toSerialize["ping-authorization-port"] = o.PingAuthorizationPort
	}
	if !IsNil(o.PingCertSubjectDn) {
		toSerialize["ping-cert-subject-dn"] = o.PingCertSubjectDn
	}
	if !IsNil(o.PingClientAuthenticationType) {
		toSerialize["ping-client-authentication-type"] = o.PingClientAuthenticationType
	}
	if !IsNil(o.PingEnforceReplayPrevention) {
		toSerialize["ping-enforce-replay-prevention"] = o.PingEnforceReplayPrevention
	}
	if !IsNil(o.PingGrantTypes) {
		toSerialize["ping-grant-types"] = o.PingGrantTypes
	}
	if !IsNil(o.PingIssuerDn) {
		toSerialize["ping-issuer-dn"] = o.PingIssuerDn
	}
	if !IsNil(o.PingJwks) {
		toSerialize["ping-jwks"] = o.PingJwks
	}
	if !IsNil(o.PingJwksUrl) {
		toSerialize["ping-jwks-url"] = o.PingJwksUrl
	}
	if !IsNil(o.PingPassword) {
		toSerialize["ping-password"] = o.PingPassword
	}
	if !IsNil(o.PingPrivilegedUser) {
		toSerialize["ping-privileged-user"] = o.PingPrivilegedUser
	}
	if !IsNil(o.PingRedirectUris) {
		toSerialize["ping-redirect-uris"] = o.PingRedirectUris
	}
	if !IsNil(o.PingRestrictedScopes) {
		toSerialize["ping-restricted-scopes"] = o.PingRestrictedScopes
	}
	if !IsNil(o.PingSigningAlgo) {
		toSerialize["ping-signing-algo"] = o.PingSigningAlgo
	}
	if !IsNil(o.PingUrl) {
		toSerialize["ping-url"] = o.PingUrl
	}
	if !IsNil(o.ProducerEncryptionKeyName) {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
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

type NullableGatewayCreateProducerPing struct {
	value *GatewayCreateProducerPing
	isSet bool
}

func (v NullableGatewayCreateProducerPing) Get() *GatewayCreateProducerPing {
	return v.value
}

func (v *NullableGatewayCreateProducerPing) Set(val *GatewayCreateProducerPing) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerPing) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerPing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerPing(val *GatewayCreateProducerPing) *NullableGatewayCreateProducerPing {
	return &NullableGatewayCreateProducerPing{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerPing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerPing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


