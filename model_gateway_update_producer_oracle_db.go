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

// GatewayUpdateProducerOracleDb gatewayUpdateProducerOracleDb is a command that updates oracle db producer
type GatewayUpdateProducerOracleDb struct {
	// (Optional) DB server certificates
	DbServerCertificates *string `json:"db-server-certificates,omitempty"`
	// (Optional) Server name for certificate verification
	DbServerName *string `json:"db-server-name,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// Oracle Host
	OracleHost *string `json:"oracle-host,omitempty"`
	// Oracle Password
	OraclePassword *string `json:"oracle-password,omitempty"`
	// Oracle Port
	OraclePort *string `json:"oracle-port,omitempty"`
	// Oracle Creation statements
	OracleScreationStatements *string `json:"oracle-screation-statements,omitempty"`
	// Oracle DB Name
	OracleServiceName *string `json:"oracle-service-name,omitempty"`
	// Oracle Username
	OracleUsername *string `json:"oracle-username,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts)
	SecureAccessHost []string `json:"secure-access-host,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
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

// NewGatewayUpdateProducerOracleDb instantiates a new GatewayUpdateProducerOracleDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerOracleDb(name string) *GatewayUpdateProducerOracleDb {
	this := GatewayUpdateProducerOracleDb{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var oracleHost string = "127.0.0.1"
	this.OracleHost = &oracleHost
	var oraclePort string = "1521"
	this.OraclePort = &oraclePort
	var secureAccessEnable string = "false"
	this.SecureAccessEnable = &secureAccessEnable
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerOracleDbWithDefaults instantiates a new GatewayUpdateProducerOracleDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerOracleDbWithDefaults() *GatewayUpdateProducerOracleDb {
	this := GatewayUpdateProducerOracleDb{}
	var json bool = false
	this.Json = &json
	var oracleHost string = "127.0.0.1"
	this.OracleHost = &oracleHost
	var oraclePort string = "1521"
	this.OraclePort = &oraclePort
	var secureAccessEnable string = "false"
	this.SecureAccessEnable = &secureAccessEnable
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDbServerCertificates returns the DbServerCertificates field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetDbServerCertificates() string {
	if o == nil || o.DbServerCertificates == nil {
		var ret string
		return ret
	}
	return *o.DbServerCertificates
}

// GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetDbServerCertificatesOk() (*string, bool) {
	if o == nil || o.DbServerCertificates == nil {
		return nil, false
	}
	return o.DbServerCertificates, true
}

// HasDbServerCertificates returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasDbServerCertificates() bool {
	if o != nil && o.DbServerCertificates != nil {
		return true
	}

	return false
}

// SetDbServerCertificates gets a reference to the given string and assigns it to the DbServerCertificates field.
func (o *GatewayUpdateProducerOracleDb) SetDbServerCertificates(v string) {
	o.DbServerCertificates = &v
}

// GetDbServerName returns the DbServerName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetDbServerName() string {
	if o == nil || o.DbServerName == nil {
		var ret string
		return ret
	}
	return *o.DbServerName
}

// GetDbServerNameOk returns a tuple with the DbServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetDbServerNameOk() (*string, bool) {
	if o == nil || o.DbServerName == nil {
		return nil, false
	}
	return o.DbServerName, true
}

// HasDbServerName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasDbServerName() bool {
	if o != nil && o.DbServerName != nil {
		return true
	}

	return false
}

// SetDbServerName gets a reference to the given string and assigns it to the DbServerName field.
func (o *GatewayUpdateProducerOracleDb) SetDbServerName(v string) {
	o.DbServerName = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerOracleDb) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerOracleDb) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerOracleDb) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerOracleDb) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerOracleDb) SetNewName(v string) {
	o.NewName = &v
}

// GetOracleHost returns the OracleHost field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetOracleHost() string {
	if o == nil || o.OracleHost == nil {
		var ret string
		return ret
	}
	return *o.OracleHost
}

// GetOracleHostOk returns a tuple with the OracleHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetOracleHostOk() (*string, bool) {
	if o == nil || o.OracleHost == nil {
		return nil, false
	}
	return o.OracleHost, true
}

// HasOracleHost returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasOracleHost() bool {
	if o != nil && o.OracleHost != nil {
		return true
	}

	return false
}

// SetOracleHost gets a reference to the given string and assigns it to the OracleHost field.
func (o *GatewayUpdateProducerOracleDb) SetOracleHost(v string) {
	o.OracleHost = &v
}

// GetOraclePassword returns the OraclePassword field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetOraclePassword() string {
	if o == nil || o.OraclePassword == nil {
		var ret string
		return ret
	}
	return *o.OraclePassword
}

// GetOraclePasswordOk returns a tuple with the OraclePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetOraclePasswordOk() (*string, bool) {
	if o == nil || o.OraclePassword == nil {
		return nil, false
	}
	return o.OraclePassword, true
}

// HasOraclePassword returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasOraclePassword() bool {
	if o != nil && o.OraclePassword != nil {
		return true
	}

	return false
}

// SetOraclePassword gets a reference to the given string and assigns it to the OraclePassword field.
func (o *GatewayUpdateProducerOracleDb) SetOraclePassword(v string) {
	o.OraclePassword = &v
}

// GetOraclePort returns the OraclePort field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetOraclePort() string {
	if o == nil || o.OraclePort == nil {
		var ret string
		return ret
	}
	return *o.OraclePort
}

// GetOraclePortOk returns a tuple with the OraclePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetOraclePortOk() (*string, bool) {
	if o == nil || o.OraclePort == nil {
		return nil, false
	}
	return o.OraclePort, true
}

// HasOraclePort returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasOraclePort() bool {
	if o != nil && o.OraclePort != nil {
		return true
	}

	return false
}

// SetOraclePort gets a reference to the given string and assigns it to the OraclePort field.
func (o *GatewayUpdateProducerOracleDb) SetOraclePort(v string) {
	o.OraclePort = &v
}

// GetOracleScreationStatements returns the OracleScreationStatements field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetOracleScreationStatements() string {
	if o == nil || o.OracleScreationStatements == nil {
		var ret string
		return ret
	}
	return *o.OracleScreationStatements
}

// GetOracleScreationStatementsOk returns a tuple with the OracleScreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetOracleScreationStatementsOk() (*string, bool) {
	if o == nil || o.OracleScreationStatements == nil {
		return nil, false
	}
	return o.OracleScreationStatements, true
}

// HasOracleScreationStatements returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasOracleScreationStatements() bool {
	if o != nil && o.OracleScreationStatements != nil {
		return true
	}

	return false
}

// SetOracleScreationStatements gets a reference to the given string and assigns it to the OracleScreationStatements field.
func (o *GatewayUpdateProducerOracleDb) SetOracleScreationStatements(v string) {
	o.OracleScreationStatements = &v
}

// GetOracleServiceName returns the OracleServiceName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetOracleServiceName() string {
	if o == nil || o.OracleServiceName == nil {
		var ret string
		return ret
	}
	return *o.OracleServiceName
}

// GetOracleServiceNameOk returns a tuple with the OracleServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetOracleServiceNameOk() (*string, bool) {
	if o == nil || o.OracleServiceName == nil {
		return nil, false
	}
	return o.OracleServiceName, true
}

// HasOracleServiceName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasOracleServiceName() bool {
	if o != nil && o.OracleServiceName != nil {
		return true
	}

	return false
}

// SetOracleServiceName gets a reference to the given string and assigns it to the OracleServiceName field.
func (o *GatewayUpdateProducerOracleDb) SetOracleServiceName(v string) {
	o.OracleServiceName = &v
}

// GetOracleUsername returns the OracleUsername field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetOracleUsername() string {
	if o == nil || o.OracleUsername == nil {
		var ret string
		return ret
	}
	return *o.OracleUsername
}

// GetOracleUsernameOk returns a tuple with the OracleUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetOracleUsernameOk() (*string, bool) {
	if o == nil || o.OracleUsername == nil {
		return nil, false
	}
	return o.OracleUsername, true
}

// HasOracleUsername returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasOracleUsername() bool {
	if o != nil && o.OracleUsername != nil {
		return true
	}

	return false
}

// SetOracleUsername gets a reference to the given string and assigns it to the OracleUsername field.
func (o *GatewayUpdateProducerOracleDb) SetOracleUsername(v string) {
	o.OracleUsername = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerOracleDb) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayUpdateProducerOracleDb) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayUpdateProducerOracleDb) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetSecureAccessHostOk() ([]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *GatewayUpdateProducerOracleDb) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayUpdateProducerOracleDb) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerOracleDb) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerOracleDb) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerOracleDb) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerOracleDb) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerOracleDb) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerOracleDb) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerOracleDb) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerOracleDb) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerOracleDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DbServerCertificates != nil {
		toSerialize["db-server-certificates"] = o.DbServerCertificates
	}
	if o.DbServerName != nil {
		toSerialize["db-server-name"] = o.DbServerName
	}
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
	if o.OracleHost != nil {
		toSerialize["oracle-host"] = o.OracleHost
	}
	if o.OraclePassword != nil {
		toSerialize["oracle-password"] = o.OraclePassword
	}
	if o.OraclePort != nil {
		toSerialize["oracle-port"] = o.OraclePort
	}
	if o.OracleScreationStatements != nil {
		toSerialize["oracle-screation-statements"] = o.OracleScreationStatements
	}
	if o.OracleServiceName != nil {
		toSerialize["oracle-service-name"] = o.OracleServiceName
	}
	if o.OracleUsername != nil {
		toSerialize["oracle-username"] = o.OracleUsername
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessHost != nil {
		toSerialize["secure-access-host"] = o.SecureAccessHost
	}
	if o.SecureAccessWeb != nil {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
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

type NullableGatewayUpdateProducerOracleDb struct {
	value *GatewayUpdateProducerOracleDb
	isSet bool
}

func (v NullableGatewayUpdateProducerOracleDb) Get() *GatewayUpdateProducerOracleDb {
	return v.value
}

func (v *NullableGatewayUpdateProducerOracleDb) Set(val *GatewayUpdateProducerOracleDb) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerOracleDb) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerOracleDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerOracleDb(val *GatewayUpdateProducerOracleDb) *NullableGatewayUpdateProducerOracleDb {
	return &NullableGatewayUpdateProducerOracleDb{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerOracleDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerOracleDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


