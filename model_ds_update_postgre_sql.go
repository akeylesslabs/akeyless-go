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

// DsUpdatePostgreSql dsUpdatePostgreSql is a command that updates postgresql dynamic secret
type DsUpdatePostgreSql struct {
	// PostgreSQL Creation statements
	CreationStatements *string `json:"creation-statements,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret name
	NewName *string `json:"new-name,omitempty"`
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
	// Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts)
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// Enable/Disable SSL [true/false]
	Ssl *bool `json:"ssl,omitempty"`
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

// NewDsUpdatePostgreSql instantiates a new DsUpdatePostgreSql object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDsUpdatePostgreSql(name string, ) *DsUpdatePostgreSql {
	this := DsUpdatePostgreSql{}
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

// NewDsUpdatePostgreSqlWithDefaults instantiates a new DsUpdatePostgreSql object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDsUpdatePostgreSqlWithDefaults() *DsUpdatePostgreSql {
	this := DsUpdatePostgreSql{}
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
func (o *DsUpdatePostgreSql) GetCreationStatements() string {
	if o == nil || o.CreationStatements == nil {
		var ret string
		return ret
	}
	return *o.CreationStatements
}

// GetCreationStatementsOk returns a tuple with the CreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetCreationStatementsOk() (*string, bool) {
	if o == nil || o.CreationStatements == nil {
		return nil, false
	}
	return o.CreationStatements, true
}

// HasCreationStatements returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasCreationStatements() bool {
	if o != nil && o.CreationStatements != nil {
		return true
	}

	return false
}

// SetCreationStatements gets a reference to the given string and assigns it to the CreationStatements field.
func (o *DsUpdatePostgreSql) SetCreationStatements(v string) {
	o.CreationStatements = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DsUpdatePostgreSql) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DsUpdatePostgreSql) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DsUpdatePostgreSql) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DsUpdatePostgreSql) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DsUpdatePostgreSql) SetNewName(v string) {
	o.NewName = &v
}

// GetPostgresqlDbName returns the PostgresqlDbName field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetPostgresqlDbName() string {
	if o == nil || o.PostgresqlDbName == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlDbName
}

// GetPostgresqlDbNameOk returns a tuple with the PostgresqlDbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetPostgresqlDbNameOk() (*string, bool) {
	if o == nil || o.PostgresqlDbName == nil {
		return nil, false
	}
	return o.PostgresqlDbName, true
}

// HasPostgresqlDbName returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasPostgresqlDbName() bool {
	if o != nil && o.PostgresqlDbName != nil {
		return true
	}

	return false
}

// SetPostgresqlDbName gets a reference to the given string and assigns it to the PostgresqlDbName field.
func (o *DsUpdatePostgreSql) SetPostgresqlDbName(v string) {
	o.PostgresqlDbName = &v
}

// GetPostgresqlHost returns the PostgresqlHost field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetPostgresqlHost() string {
	if o == nil || o.PostgresqlHost == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlHost
}

// GetPostgresqlHostOk returns a tuple with the PostgresqlHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetPostgresqlHostOk() (*string, bool) {
	if o == nil || o.PostgresqlHost == nil {
		return nil, false
	}
	return o.PostgresqlHost, true
}

// HasPostgresqlHost returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasPostgresqlHost() bool {
	if o != nil && o.PostgresqlHost != nil {
		return true
	}

	return false
}

// SetPostgresqlHost gets a reference to the given string and assigns it to the PostgresqlHost field.
func (o *DsUpdatePostgreSql) SetPostgresqlHost(v string) {
	o.PostgresqlHost = &v
}

// GetPostgresqlPassword returns the PostgresqlPassword field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetPostgresqlPassword() string {
	if o == nil || o.PostgresqlPassword == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlPassword
}

// GetPostgresqlPasswordOk returns a tuple with the PostgresqlPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetPostgresqlPasswordOk() (*string, bool) {
	if o == nil || o.PostgresqlPassword == nil {
		return nil, false
	}
	return o.PostgresqlPassword, true
}

// HasPostgresqlPassword returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasPostgresqlPassword() bool {
	if o != nil && o.PostgresqlPassword != nil {
		return true
	}

	return false
}

// SetPostgresqlPassword gets a reference to the given string and assigns it to the PostgresqlPassword field.
func (o *DsUpdatePostgreSql) SetPostgresqlPassword(v string) {
	o.PostgresqlPassword = &v
}

// GetPostgresqlPort returns the PostgresqlPort field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetPostgresqlPort() string {
	if o == nil || o.PostgresqlPort == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlPort
}

// GetPostgresqlPortOk returns a tuple with the PostgresqlPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetPostgresqlPortOk() (*string, bool) {
	if o == nil || o.PostgresqlPort == nil {
		return nil, false
	}
	return o.PostgresqlPort, true
}

// HasPostgresqlPort returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasPostgresqlPort() bool {
	if o != nil && o.PostgresqlPort != nil {
		return true
	}

	return false
}

// SetPostgresqlPort gets a reference to the given string and assigns it to the PostgresqlPort field.
func (o *DsUpdatePostgreSql) SetPostgresqlPort(v string) {
	o.PostgresqlPort = &v
}

// GetPostgresqlUsername returns the PostgresqlUsername field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetPostgresqlUsername() string {
	if o == nil || o.PostgresqlUsername == nil {
		var ret string
		return ret
	}
	return *o.PostgresqlUsername
}

// GetPostgresqlUsernameOk returns a tuple with the PostgresqlUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetPostgresqlUsernameOk() (*string, bool) {
	if o == nil || o.PostgresqlUsername == nil {
		return nil, false
	}
	return o.PostgresqlUsername, true
}

// HasPostgresqlUsername returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasPostgresqlUsername() bool {
	if o != nil && o.PostgresqlUsername != nil {
		return true
	}

	return false
}

// SetPostgresqlUsername gets a reference to the given string and assigns it to the PostgresqlUsername field.
func (o *DsUpdatePostgreSql) SetPostgresqlUsername(v string) {
	o.PostgresqlUsername = &v
}

// GetProducerEncryptionKey returns the ProducerEncryptionKey field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetProducerEncryptionKey() string {
	if o == nil || o.ProducerEncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKey
}

// GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetProducerEncryptionKeyOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKey == nil {
		return nil, false
	}
	return o.ProducerEncryptionKey, true
}

// HasProducerEncryptionKey returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasProducerEncryptionKey() bool {
	if o != nil && o.ProducerEncryptionKey != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKey gets a reference to the given string and assigns it to the ProducerEncryptionKey field.
func (o *DsUpdatePostgreSql) SetProducerEncryptionKey(v string) {
	o.ProducerEncryptionKey = &v
}

// GetRevocationStatement returns the RevocationStatement field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetRevocationStatement() string {
	if o == nil || o.RevocationStatement == nil {
		var ret string
		return ret
	}
	return *o.RevocationStatement
}

// GetRevocationStatementOk returns a tuple with the RevocationStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetRevocationStatementOk() (*string, bool) {
	if o == nil || o.RevocationStatement == nil {
		return nil, false
	}
	return o.RevocationStatement, true
}

// HasRevocationStatement returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasRevocationStatement() bool {
	if o != nil && o.RevocationStatement != nil {
		return true
	}

	return false
}

// SetRevocationStatement gets a reference to the given string and assigns it to the RevocationStatement field.
func (o *DsUpdatePostgreSql) SetRevocationStatement(v string) {
	o.RevocationStatement = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *DsUpdatePostgreSql) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessDbSchema returns the SecureAccessDbSchema field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetSecureAccessDbSchema() string {
	if o == nil || o.SecureAccessDbSchema == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDbSchema
}

// GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetSecureAccessDbSchemaOk() (*string, bool) {
	if o == nil || o.SecureAccessDbSchema == nil {
		return nil, false
	}
	return o.SecureAccessDbSchema, true
}

// HasSecureAccessDbSchema returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasSecureAccessDbSchema() bool {
	if o != nil && o.SecureAccessDbSchema != nil {
		return true
	}

	return false
}

// SetSecureAccessDbSchema gets a reference to the given string and assigns it to the SecureAccessDbSchema field.
func (o *DsUpdatePostgreSql) SetSecureAccessDbSchema(v string) {
	o.SecureAccessDbSchema = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DsUpdatePostgreSql) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *DsUpdatePostgreSql) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *DsUpdatePostgreSql) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetSsl() bool {
	if o == nil || o.Ssl == nil {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetSslOk() (*bool, bool) {
	if o == nil || o.Ssl == nil {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasSsl() bool {
	if o != nil && o.Ssl != nil {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *DsUpdatePostgreSql) SetSsl(v bool) {
	o.Ssl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DsUpdatePostgreSql) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DsUpdatePostgreSql) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DsUpdatePostgreSql) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DsUpdatePostgreSql) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DsUpdatePostgreSql) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdatePostgreSql) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DsUpdatePostgreSql) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DsUpdatePostgreSql) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DsUpdatePostgreSql) MarshalJSON() ([]byte, error) {
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
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
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

type NullableDsUpdatePostgreSql struct {
	value *DsUpdatePostgreSql
	isSet bool
}

func (v NullableDsUpdatePostgreSql) Get() *DsUpdatePostgreSql {
	return v.value
}

func (v *NullableDsUpdatePostgreSql) Set(val *DsUpdatePostgreSql) {
	v.value = val
	v.isSet = true
}

func (v NullableDsUpdatePostgreSql) IsSet() bool {
	return v.isSet
}

func (v *NullableDsUpdatePostgreSql) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDsUpdatePostgreSql(val *DsUpdatePostgreSql) *NullableDsUpdatePostgreSql {
	return &NullableDsUpdatePostgreSql{value: val, isSet: true}
}

func (v NullableDsUpdatePostgreSql) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDsUpdatePostgreSql) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


