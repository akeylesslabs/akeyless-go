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

// GatewayUpdateProducerMSSQL gatewayUpdateProducerMSSQL is a command that updates mssql producer [Deprecated: Use dynamic-secret-update-mssql command]
type GatewayUpdateProducerMSSQL struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// MSSQL Creation statements
	MssqlCreateStatements *string `json:"mssql-create-statements,omitempty"`
	// MSSQL Name
	MssqlDbname *string `json:"mssql-dbname,omitempty"`
	// MSSQL Host
	MssqlHost *string `json:"mssql-host,omitempty"`
	// MSSQL Password
	MssqlPassword *string `json:"mssql-password,omitempty"`
	// MSSQL Port
	MssqlPort *string `json:"mssql-port,omitempty"`
	// MSSQL Revocation statements
	MssqlRevocationStatements *string `json:"mssql-revocation-statements,omitempty"`
	// MSSQL Username
	MssqlUsername *string `json:"mssql-username,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret name
	NewName *string `json:"new-name,omitempty"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// The DB schema
	SecureAccessDbSchema *string `json:"secure-access-db-schema,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts)
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
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
}

// NewGatewayUpdateProducerMSSQL instantiates a new GatewayUpdateProducerMSSQL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerMSSQL(name string, ) *GatewayUpdateProducerMSSQL {
	this := GatewayUpdateProducerMSSQL{}
	var json bool = false
	this.Json = &json
	var mssqlHost string = "127.0.0.1"
	this.MssqlHost = &mssqlHost
	var mssqlPort string = "1433"
	this.MssqlPort = &mssqlPort
	this.Name = name
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerMSSQLWithDefaults instantiates a new GatewayUpdateProducerMSSQL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerMSSQLWithDefaults() *GatewayUpdateProducerMSSQL {
	this := GatewayUpdateProducerMSSQL{}
	var json bool = false
	this.Json = &json
	var mssqlHost string = "127.0.0.1"
	this.MssqlHost = &mssqlHost
	var mssqlPort string = "1433"
	this.MssqlPort = &mssqlPort
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerMSSQL) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerMSSQL) SetJson(v bool) {
	o.Json = &v
}

// GetMssqlCreateStatements returns the MssqlCreateStatements field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetMssqlCreateStatements() string {
	if o == nil || o.MssqlCreateStatements == nil {
		var ret string
		return ret
	}
	return *o.MssqlCreateStatements
}

// GetMssqlCreateStatementsOk returns a tuple with the MssqlCreateStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetMssqlCreateStatementsOk() (*string, bool) {
	if o == nil || o.MssqlCreateStatements == nil {
		return nil, false
	}
	return o.MssqlCreateStatements, true
}

// HasMssqlCreateStatements returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasMssqlCreateStatements() bool {
	if o != nil && o.MssqlCreateStatements != nil {
		return true
	}

	return false
}

// SetMssqlCreateStatements gets a reference to the given string and assigns it to the MssqlCreateStatements field.
func (o *GatewayUpdateProducerMSSQL) SetMssqlCreateStatements(v string) {
	o.MssqlCreateStatements = &v
}

// GetMssqlDbname returns the MssqlDbname field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetMssqlDbname() string {
	if o == nil || o.MssqlDbname == nil {
		var ret string
		return ret
	}
	return *o.MssqlDbname
}

// GetMssqlDbnameOk returns a tuple with the MssqlDbname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetMssqlDbnameOk() (*string, bool) {
	if o == nil || o.MssqlDbname == nil {
		return nil, false
	}
	return o.MssqlDbname, true
}

// HasMssqlDbname returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasMssqlDbname() bool {
	if o != nil && o.MssqlDbname != nil {
		return true
	}

	return false
}

// SetMssqlDbname gets a reference to the given string and assigns it to the MssqlDbname field.
func (o *GatewayUpdateProducerMSSQL) SetMssqlDbname(v string) {
	o.MssqlDbname = &v
}

// GetMssqlHost returns the MssqlHost field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetMssqlHost() string {
	if o == nil || o.MssqlHost == nil {
		var ret string
		return ret
	}
	return *o.MssqlHost
}

// GetMssqlHostOk returns a tuple with the MssqlHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetMssqlHostOk() (*string, bool) {
	if o == nil || o.MssqlHost == nil {
		return nil, false
	}
	return o.MssqlHost, true
}

// HasMssqlHost returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasMssqlHost() bool {
	if o != nil && o.MssqlHost != nil {
		return true
	}

	return false
}

// SetMssqlHost gets a reference to the given string and assigns it to the MssqlHost field.
func (o *GatewayUpdateProducerMSSQL) SetMssqlHost(v string) {
	o.MssqlHost = &v
}

// GetMssqlPassword returns the MssqlPassword field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetMssqlPassword() string {
	if o == nil || o.MssqlPassword == nil {
		var ret string
		return ret
	}
	return *o.MssqlPassword
}

// GetMssqlPasswordOk returns a tuple with the MssqlPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetMssqlPasswordOk() (*string, bool) {
	if o == nil || o.MssqlPassword == nil {
		return nil, false
	}
	return o.MssqlPassword, true
}

// HasMssqlPassword returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasMssqlPassword() bool {
	if o != nil && o.MssqlPassword != nil {
		return true
	}

	return false
}

// SetMssqlPassword gets a reference to the given string and assigns it to the MssqlPassword field.
func (o *GatewayUpdateProducerMSSQL) SetMssqlPassword(v string) {
	o.MssqlPassword = &v
}

// GetMssqlPort returns the MssqlPort field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetMssqlPort() string {
	if o == nil || o.MssqlPort == nil {
		var ret string
		return ret
	}
	return *o.MssqlPort
}

// GetMssqlPortOk returns a tuple with the MssqlPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetMssqlPortOk() (*string, bool) {
	if o == nil || o.MssqlPort == nil {
		return nil, false
	}
	return o.MssqlPort, true
}

// HasMssqlPort returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasMssqlPort() bool {
	if o != nil && o.MssqlPort != nil {
		return true
	}

	return false
}

// SetMssqlPort gets a reference to the given string and assigns it to the MssqlPort field.
func (o *GatewayUpdateProducerMSSQL) SetMssqlPort(v string) {
	o.MssqlPort = &v
}

// GetMssqlRevocationStatements returns the MssqlRevocationStatements field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetMssqlRevocationStatements() string {
	if o == nil || o.MssqlRevocationStatements == nil {
		var ret string
		return ret
	}
	return *o.MssqlRevocationStatements
}

// GetMssqlRevocationStatementsOk returns a tuple with the MssqlRevocationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetMssqlRevocationStatementsOk() (*string, bool) {
	if o == nil || o.MssqlRevocationStatements == nil {
		return nil, false
	}
	return o.MssqlRevocationStatements, true
}

// HasMssqlRevocationStatements returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasMssqlRevocationStatements() bool {
	if o != nil && o.MssqlRevocationStatements != nil {
		return true
	}

	return false
}

// SetMssqlRevocationStatements gets a reference to the given string and assigns it to the MssqlRevocationStatements field.
func (o *GatewayUpdateProducerMSSQL) SetMssqlRevocationStatements(v string) {
	o.MssqlRevocationStatements = &v
}

// GetMssqlUsername returns the MssqlUsername field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetMssqlUsername() string {
	if o == nil || o.MssqlUsername == nil {
		var ret string
		return ret
	}
	return *o.MssqlUsername
}

// GetMssqlUsernameOk returns a tuple with the MssqlUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetMssqlUsernameOk() (*string, bool) {
	if o == nil || o.MssqlUsername == nil {
		return nil, false
	}
	return o.MssqlUsername, true
}

// HasMssqlUsername returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasMssqlUsername() bool {
	if o != nil && o.MssqlUsername != nil {
		return true
	}

	return false
}

// SetMssqlUsername gets a reference to the given string and assigns it to the MssqlUsername field.
func (o *GatewayUpdateProducerMSSQL) SetMssqlUsername(v string) {
	o.MssqlUsername = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerMSSQL) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerMSSQL) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerMSSQL) SetNewName(v string) {
	o.NewName = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *GatewayUpdateProducerMSSQL) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerMSSQL) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayUpdateProducerMSSQL) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessDbSchema returns the SecureAccessDbSchema field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessDbSchema() string {
	if o == nil || o.SecureAccessDbSchema == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDbSchema
}

// GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessDbSchemaOk() (*string, bool) {
	if o == nil || o.SecureAccessDbSchema == nil {
		return nil, false
	}
	return o.SecureAccessDbSchema, true
}

// HasSecureAccessDbSchema returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasSecureAccessDbSchema() bool {
	if o != nil && o.SecureAccessDbSchema != nil {
		return true
	}

	return false
}

// SetSecureAccessDbSchema gets a reference to the given string and assigns it to the SecureAccessDbSchema field.
func (o *GatewayUpdateProducerMSSQL) SetSecureAccessDbSchema(v string) {
	o.SecureAccessDbSchema = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayUpdateProducerMSSQL) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *GatewayUpdateProducerMSSQL) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayUpdateProducerMSSQL) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerMSSQL) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerMSSQL) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerMSSQL) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerMSSQL) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMSSQL) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMSSQL) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMSSQL) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerMSSQL) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerMSSQL) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.MssqlCreateStatements != nil {
		toSerialize["mssql-create-statements"] = o.MssqlCreateStatements
	}
	if o.MssqlDbname != nil {
		toSerialize["mssql-dbname"] = o.MssqlDbname
	}
	if o.MssqlHost != nil {
		toSerialize["mssql-host"] = o.MssqlHost
	}
	if o.MssqlPassword != nil {
		toSerialize["mssql-password"] = o.MssqlPassword
	}
	if o.MssqlPort != nil {
		toSerialize["mssql-port"] = o.MssqlPort
	}
	if o.MssqlRevocationStatements != nil {
		toSerialize["mssql-revocation-statements"] = o.MssqlRevocationStatements
	}
	if o.MssqlUsername != nil {
		toSerialize["mssql-username"] = o.MssqlUsername
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
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessDbSchema != nil {
		toSerialize["secure-access-db-schema"] = o.SecureAccessDbSchema
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

type NullableGatewayUpdateProducerMSSQL struct {
	value *GatewayUpdateProducerMSSQL
	isSet bool
}

func (v NullableGatewayUpdateProducerMSSQL) Get() *GatewayUpdateProducerMSSQL {
	return v.value
}

func (v *NullableGatewayUpdateProducerMSSQL) Set(val *GatewayUpdateProducerMSSQL) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerMSSQL) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerMSSQL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerMSSQL(val *GatewayUpdateProducerMSSQL) *NullableGatewayUpdateProducerMSSQL {
	return &NullableGatewayUpdateProducerMSSQL{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerMSSQL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerMSSQL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


