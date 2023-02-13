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

// GatewayCreateProducerPostgreSQL gatewayCreateProducerPostgreSQL is a command that creates postgresql producer
type GatewayCreateProducerPostgreSQL struct {
	// PostgreSQL Creation statements
	CreationStatements *string `json:"creation-statements,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
	// PostgreSQL DB Name
	PostgresqlDbName *string `json:"postgresql-db-name,omitempty"`
	// PostgreSQL Host
	PostgresqlHost *string `json:"postgresql-host,omitempty"`
	// PostgreSQL Password
	PostgresqlPassword *string `json:"postgresql-password,omitempty"`
	// PostgreSQL Port
	PostgresqlPort *string `json:"postgresql-port,omitempty"`
	// PostgreSQL Username
	PostgresqlUsername *string `json:"postgresql-username,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKey *string `json:"producer-encryption-key,omitempty"`
	// PostgreSQL Revocation statements
	RevocationStatement *string `json:"revocation-statement,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// The DB schema
	SecureAccessDbSchema *string `json:"secure-access-db-schema,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target DB servers for connections
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// Enable/Disable SSL [true/false]
	Ssl *bool `json:"ssl,omitempty"`
	// List of the tags attached to this secret
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

// NewGatewayCreateProducerPostgreSQL instantiates a new GatewayCreateProducerPostgreSQL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerPostgreSQL(name string, ) *GatewayCreateProducerPostgreSQL {
	this := GatewayCreateProducerPostgreSQL{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var postgresqlHost string = "127.0.0.1"
	this.PostgresqlHost = &postgresqlHost
	var postgresqlPort string = "5432"
	this.PostgresqlPort = &postgresqlPort
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var ssl bool = false
	this.Ssl = &ssl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerPostgreSQLWithDefaults instantiates a new GatewayCreateProducerPostgreSQL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerPostgreSQLWithDefaults() *GatewayCreateProducerPostgreSQL {
	this := GatewayCreateProducerPostgreSQL{}
	var json bool = false
	this.Json = &json
	var postgresqlHost string = "127.0.0.1"
	this.PostgresqlHost = &postgresqlHost
	var postgresqlPort string = "5432"
	this.PostgresqlPort = &postgresqlPort
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var ssl bool = false
	this.Ssl = &ssl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetCreationStatements returns the CreationStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetCreationStatements() string {
	if o == nil || o.CreationStatements == nil {
		var ret string
		return ret
	}
	return *o.CreationStatements
}

// GetCreationStatementsOk returns a tuple with the CreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetCreationStatementsOk() (*string, bool) {
	if o == nil || o.CreationStatements == nil {
		return nil, false
	}
	return o.CreationStatements, true
}

// HasCreationStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasCreationStatements() bool {
	if o != nil && o.CreationStatements != nil {
		return true
	}

	return false
}

// SetCreationStatements gets a reference to the given string and assigns it to the CreationStatements field.
func (o *GatewayCreateProducerPostgreSQL) SetCreationStatements(v string) {
	o.CreationStatements = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerPostgreSQL) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerPostgreSQL) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerPostgreSQL) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerPostgreSQL) SetName(v string) {
	o.Name = v
}

// GetPostgresqlDbName returns the PostgresqlDbName field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlDbName() string {
	if o == nil || o.PostgresqlDbName == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlDbName
}

// GetPostgresqlDbNameOk returns a tuple with the PostgresqlDbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlDbNameOk() (*string, bool) {
	if o == nil || o.PostgresqlDbName == nil {
		return nil, false
	}
	return o.PostgresqlDbName, true
}

// HasPostgresqlDbName returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasPostgresqlDbName() bool {
	if o != nil && o.PostgresqlDbName != nil {
		return true
	}

	return false
}

// SetPostgresqlDbName gets a reference to the given string and assigns it to the PostgresqlDbName field.
func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlDbName(v string) {
	o.PostgresqlDbName = &v
}

// GetPostgresqlHost returns the PostgresqlHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlHost() string {
	if o == nil || o.PostgresqlHost == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlHost
}

// GetPostgresqlHostOk returns a tuple with the PostgresqlHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlHostOk() (*string, bool) {
	if o == nil || o.PostgresqlHost == nil {
		return nil, false
	}
	return o.PostgresqlHost, true
}

// HasPostgresqlHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasPostgresqlHost() bool {
	if o != nil && o.PostgresqlHost != nil {
		return true
	}

	return false
}

// SetPostgresqlHost gets a reference to the given string and assigns it to the PostgresqlHost field.
func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlHost(v string) {
	o.PostgresqlHost = &v
}

// GetPostgresqlPassword returns the PostgresqlPassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlPassword() string {
	if o == nil || o.PostgresqlPassword == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlPassword
}

// GetPostgresqlPasswordOk returns a tuple with the PostgresqlPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlPasswordOk() (*string, bool) {
	if o == nil || o.PostgresqlPassword == nil {
		return nil, false
	}
	return o.PostgresqlPassword, true
}

// HasPostgresqlPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasPostgresqlPassword() bool {
	if o != nil && o.PostgresqlPassword != nil {
		return true
	}

	return false
}

// SetPostgresqlPassword gets a reference to the given string and assigns it to the PostgresqlPassword field.
func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlPassword(v string) {
	o.PostgresqlPassword = &v
}

// GetPostgresqlPort returns the PostgresqlPort field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlPort() string {
	if o == nil || o.PostgresqlPort == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlPort
}

// GetPostgresqlPortOk returns a tuple with the PostgresqlPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlPortOk() (*string, bool) {
	if o == nil || o.PostgresqlPort == nil {
		return nil, false
	}
	return o.PostgresqlPort, true
}

// HasPostgresqlPort returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasPostgresqlPort() bool {
	if o != nil && o.PostgresqlPort != nil {
		return true
	}

	return false
}

// SetPostgresqlPort gets a reference to the given string and assigns it to the PostgresqlPort field.
func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlPort(v string) {
	o.PostgresqlPort = &v
}

// GetPostgresqlUsername returns the PostgresqlUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlUsername() string {
	if o == nil || o.PostgresqlUsername == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlUsername
}

// GetPostgresqlUsernameOk returns a tuple with the PostgresqlUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlUsernameOk() (*string, bool) {
	if o == nil || o.PostgresqlUsername == nil {
		return nil, false
	}
	return o.PostgresqlUsername, true
}

// HasPostgresqlUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasPostgresqlUsername() bool {
	if o != nil && o.PostgresqlUsername != nil {
		return true
	}

	return false
}

// SetPostgresqlUsername gets a reference to the given string and assigns it to the PostgresqlUsername field.
func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlUsername(v string) {
	o.PostgresqlUsername = &v
}

// GetProducerEncryptionKey returns the ProducerEncryptionKey field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetProducerEncryptionKey() string {
	if o == nil || o.ProducerEncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKey
}

// GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetProducerEncryptionKeyOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKey == nil {
		return nil, false
	}
	return o.ProducerEncryptionKey, true
}

// HasProducerEncryptionKey returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasProducerEncryptionKey() bool {
	if o != nil && o.ProducerEncryptionKey != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKey gets a reference to the given string and assigns it to the ProducerEncryptionKey field.
func (o *GatewayCreateProducerPostgreSQL) SetProducerEncryptionKey(v string) {
	o.ProducerEncryptionKey = &v
}

// GetRevocationStatement returns the RevocationStatement field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetRevocationStatement() string {
	if o == nil || o.RevocationStatement == nil {
		var ret string
		return ret
	}
	return *o.RevocationStatement
}

// GetRevocationStatementOk returns a tuple with the RevocationStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetRevocationStatementOk() (*string, bool) {
	if o == nil || o.RevocationStatement == nil {
		return nil, false
	}
	return o.RevocationStatement, true
}

// HasRevocationStatement returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasRevocationStatement() bool {
	if o != nil && o.RevocationStatement != nil {
		return true
	}

	return false
}

// SetRevocationStatement gets a reference to the given string and assigns it to the RevocationStatement field.
func (o *GatewayCreateProducerPostgreSQL) SetRevocationStatement(v string) {
	o.RevocationStatement = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayCreateProducerPostgreSQL) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessDbSchema returns the SecureAccessDbSchema field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessDbSchema() string {
	if o == nil || o.SecureAccessDbSchema == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDbSchema
}

// GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessDbSchemaOk() (*string, bool) {
	if o == nil || o.SecureAccessDbSchema == nil {
		return nil, false
	}
	return o.SecureAccessDbSchema, true
}

// HasSecureAccessDbSchema returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasSecureAccessDbSchema() bool {
	if o != nil && o.SecureAccessDbSchema != nil {
		return true
	}

	return false
}

// SetSecureAccessDbSchema gets a reference to the given string and assigns it to the SecureAccessDbSchema field.
func (o *GatewayCreateProducerPostgreSQL) SetSecureAccessDbSchema(v string) {
	o.SecureAccessDbSchema = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayCreateProducerPostgreSQL) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *GatewayCreateProducerPostgreSQL) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayCreateProducerPostgreSQL) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetSsl() bool {
	if o == nil || o.Ssl == nil {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetSslOk() (*bool, bool) {
	if o == nil || o.Ssl == nil {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasSsl() bool {
	if o != nil && o.Ssl != nil {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *GatewayCreateProducerPostgreSQL) SetSsl(v bool) {
	o.Ssl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerPostgreSQL) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerPostgreSQL) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerPostgreSQL) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerPostgreSQL) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerPostgreSQL) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerPostgreSQL) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerPostgreSQL) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerPostgreSQL) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerPostgreSQL) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationStatements != nil {
		toSerialize["creation-statements"] = o.CreationStatements
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
	if o.PostgresqlDbName != nil {
		toSerialize["postgresql-db-name"] = o.PostgresqlDbName
	}
	if o.PostgresqlHost != nil {
		toSerialize["postgresql-host"] = o.PostgresqlHost
	}
	if o.PostgresqlPassword != nil {
		toSerialize["postgresql-password"] = o.PostgresqlPassword
	}
	if o.PostgresqlPort != nil {
		toSerialize["postgresql-port"] = o.PostgresqlPort
	}
	if o.PostgresqlUsername != nil {
		toSerialize["postgresql-username"] = o.PostgresqlUsername
	}
	if o.ProducerEncryptionKey != nil {
		toSerialize["producer-encryption-key"] = o.ProducerEncryptionKey
	}
	if o.RevocationStatement != nil {
		toSerialize["revocation-statement"] = o.RevocationStatement
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
	if o.Ssl != nil {
		toSerialize["ssl"] = o.Ssl
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

type NullableGatewayCreateProducerPostgreSQL struct {
	value *GatewayCreateProducerPostgreSQL
	isSet bool
}

func (v NullableGatewayCreateProducerPostgreSQL) Get() *GatewayCreateProducerPostgreSQL {
	return v.value
}

func (v *NullableGatewayCreateProducerPostgreSQL) Set(val *GatewayCreateProducerPostgreSQL) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerPostgreSQL) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerPostgreSQL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerPostgreSQL(val *GatewayCreateProducerPostgreSQL) *NullableGatewayCreateProducerPostgreSQL {
	return &NullableGatewayCreateProducerPostgreSQL{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerPostgreSQL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerPostgreSQL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


