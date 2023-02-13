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

// GatewayUpdateProducerMySQL gatewayUpdateProducerMySQL is a command that updates mysql producer
type GatewayUpdateProducerMySQL struct {
	// (Optional) DB server certificates
	DbServerCertificates *string `json:"db-server-certificates,omitempty"`
	// (Optional) Server name for certificate verification
	DbServerName *string `json:"db-server-name,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// MySQL DB Name
	MysqlDbname *string `json:"mysql-dbname,omitempty"`
	// MySQL Host
	MysqlHost *string `json:"mysql-host,omitempty"`
	// MySQL Password
	MysqlPassword *string `json:"mysql-password,omitempty"`
	// MySQL Port
	MysqlPort *string `json:"mysql-port,omitempty"`
	// MySQL Creation statements
	MysqlScreationStatements *string `json:"mysql-screation-statements,omitempty"`
	// MySQL Username
	MysqlUsername *string `json:"mysql-username,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target DB servers for connections
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// Enable/Disable SSL [true/false]
	Ssl *bool `json:"ssl,omitempty"`
	// SSL connection certificate
	SslCertificate *string `json:"ssl-certificate,omitempty"`
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

// NewGatewayUpdateProducerMySQL instantiates a new GatewayUpdateProducerMySQL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerMySQL(name string, ) *GatewayUpdateProducerMySQL {
	this := GatewayUpdateProducerMySQL{}
	var json bool = false
	this.Json = &json
	var mysqlHost string = "127.0.0.1"
	this.MysqlHost = &mysqlHost
	var mysqlPort string = "3306"
	this.MysqlPort = &mysqlPort
	this.Name = name
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var ssl bool = false
	this.Ssl = &ssl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerMySQLWithDefaults instantiates a new GatewayUpdateProducerMySQL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerMySQLWithDefaults() *GatewayUpdateProducerMySQL {
	this := GatewayUpdateProducerMySQL{}
	var json bool = false
	this.Json = &json
	var mysqlHost string = "127.0.0.1"
	this.MysqlHost = &mysqlHost
	var mysqlPort string = "3306"
	this.MysqlPort = &mysqlPort
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var ssl bool = false
	this.Ssl = &ssl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDbServerCertificates returns the DbServerCertificates field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetDbServerCertificates() string {
	if o == nil || o.DbServerCertificates == nil {
		var ret string
		return ret
	}
	return *o.DbServerCertificates
}

// GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetDbServerCertificatesOk() (*string, bool) {
	if o == nil || o.DbServerCertificates == nil {
		return nil, false
	}
	return o.DbServerCertificates, true
}

// HasDbServerCertificates returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasDbServerCertificates() bool {
	if o != nil && o.DbServerCertificates != nil {
		return true
	}

	return false
}

// SetDbServerCertificates gets a reference to the given string and assigns it to the DbServerCertificates field.
func (o *GatewayUpdateProducerMySQL) SetDbServerCertificates(v string) {
	o.DbServerCertificates = &v
}

// GetDbServerName returns the DbServerName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetDbServerName() string {
	if o == nil || o.DbServerName == nil {
		var ret string
		return ret
	}
	return *o.DbServerName
}

// GetDbServerNameOk returns a tuple with the DbServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetDbServerNameOk() (*string, bool) {
	if o == nil || o.DbServerName == nil {
		return nil, false
	}
	return o.DbServerName, true
}

// HasDbServerName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasDbServerName() bool {
	if o != nil && o.DbServerName != nil {
		return true
	}

	return false
}

// SetDbServerName gets a reference to the given string and assigns it to the DbServerName field.
func (o *GatewayUpdateProducerMySQL) SetDbServerName(v string) {
	o.DbServerName = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerMySQL) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateProducerMySQL) SetJson(v bool) {
	o.Json = &v
}

// GetMysqlDbname returns the MysqlDbname field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetMysqlDbname() string {
	if o == nil || o.MysqlDbname == nil {
		var ret string
		return ret
	}
	return *o.MysqlDbname
}

// GetMysqlDbnameOk returns a tuple with the MysqlDbname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetMysqlDbnameOk() (*string, bool) {
	if o == nil || o.MysqlDbname == nil {
		return nil, false
	}
	return o.MysqlDbname, true
}

// HasMysqlDbname returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasMysqlDbname() bool {
	if o != nil && o.MysqlDbname != nil {
		return true
	}

	return false
}

// SetMysqlDbname gets a reference to the given string and assigns it to the MysqlDbname field.
func (o *GatewayUpdateProducerMySQL) SetMysqlDbname(v string) {
	o.MysqlDbname = &v
}

// GetMysqlHost returns the MysqlHost field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetMysqlHost() string {
	if o == nil || o.MysqlHost == nil {
		var ret string
		return ret
	}
	return *o.MysqlHost
}

// GetMysqlHostOk returns a tuple with the MysqlHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetMysqlHostOk() (*string, bool) {
	if o == nil || o.MysqlHost == nil {
		return nil, false
	}
	return o.MysqlHost, true
}

// HasMysqlHost returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasMysqlHost() bool {
	if o != nil && o.MysqlHost != nil {
		return true
	}

	return false
}

// SetMysqlHost gets a reference to the given string and assigns it to the MysqlHost field.
func (o *GatewayUpdateProducerMySQL) SetMysqlHost(v string) {
	o.MysqlHost = &v
}

// GetMysqlPassword returns the MysqlPassword field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetMysqlPassword() string {
	if o == nil || o.MysqlPassword == nil {
		var ret string
		return ret
	}
	return *o.MysqlPassword
}

// GetMysqlPasswordOk returns a tuple with the MysqlPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetMysqlPasswordOk() (*string, bool) {
	if o == nil || o.MysqlPassword == nil {
		return nil, false
	}
	return o.MysqlPassword, true
}

// HasMysqlPassword returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasMysqlPassword() bool {
	if o != nil && o.MysqlPassword != nil {
		return true
	}

	return false
}

// SetMysqlPassword gets a reference to the given string and assigns it to the MysqlPassword field.
func (o *GatewayUpdateProducerMySQL) SetMysqlPassword(v string) {
	o.MysqlPassword = &v
}

// GetMysqlPort returns the MysqlPort field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetMysqlPort() string {
	if o == nil || o.MysqlPort == nil {
		var ret string
		return ret
	}
	return *o.MysqlPort
}

// GetMysqlPortOk returns a tuple with the MysqlPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetMysqlPortOk() (*string, bool) {
	if o == nil || o.MysqlPort == nil {
		return nil, false
	}
	return o.MysqlPort, true
}

// HasMysqlPort returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasMysqlPort() bool {
	if o != nil && o.MysqlPort != nil {
		return true
	}

	return false
}

// SetMysqlPort gets a reference to the given string and assigns it to the MysqlPort field.
func (o *GatewayUpdateProducerMySQL) SetMysqlPort(v string) {
	o.MysqlPort = &v
}

// GetMysqlScreationStatements returns the MysqlScreationStatements field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetMysqlScreationStatements() string {
	if o == nil || o.MysqlScreationStatements == nil {
		var ret string
		return ret
	}
	return *o.MysqlScreationStatements
}

// GetMysqlScreationStatementsOk returns a tuple with the MysqlScreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetMysqlScreationStatementsOk() (*string, bool) {
	if o == nil || o.MysqlScreationStatements == nil {
		return nil, false
	}
	return o.MysqlScreationStatements, true
}

// HasMysqlScreationStatements returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasMysqlScreationStatements() bool {
	if o != nil && o.MysqlScreationStatements != nil {
		return true
	}

	return false
}

// SetMysqlScreationStatements gets a reference to the given string and assigns it to the MysqlScreationStatements field.
func (o *GatewayUpdateProducerMySQL) SetMysqlScreationStatements(v string) {
	o.MysqlScreationStatements = &v
}

// GetMysqlUsername returns the MysqlUsername field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetMysqlUsername() string {
	if o == nil || o.MysqlUsername == nil {
		var ret string
		return ret
	}
	return *o.MysqlUsername
}

// GetMysqlUsernameOk returns a tuple with the MysqlUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetMysqlUsernameOk() (*string, bool) {
	if o == nil || o.MysqlUsername == nil {
		return nil, false
	}
	return o.MysqlUsername, true
}

// HasMysqlUsername returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasMysqlUsername() bool {
	if o != nil && o.MysqlUsername != nil {
		return true
	}

	return false
}

// SetMysqlUsername gets a reference to the given string and assigns it to the MysqlUsername field.
func (o *GatewayUpdateProducerMySQL) SetMysqlUsername(v string) {
	o.MysqlUsername = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerMySQL) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerMySQL) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerMySQL) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayUpdateProducerMySQL) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayUpdateProducerMySQL) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayUpdateProducerMySQL) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *GatewayUpdateProducerMySQL) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayUpdateProducerMySQL) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetSsl() bool {
	if o == nil || o.Ssl == nil {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetSslOk() (*bool, bool) {
	if o == nil || o.Ssl == nil {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasSsl() bool {
	if o != nil && o.Ssl != nil {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *GatewayUpdateProducerMySQL) SetSsl(v bool) {
	o.Ssl = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetSslCertificate() string {
	if o == nil || o.SslCertificate == nil {
		var ret string
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetSslCertificateOk() (*string, bool) {
	if o == nil || o.SslCertificate == nil {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasSslCertificate() bool {
	if o != nil && o.SslCertificate != nil {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given string and assigns it to the SslCertificate field.
func (o *GatewayUpdateProducerMySQL) SetSslCertificate(v string) {
	o.SslCertificate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerMySQL) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerMySQL) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerMySQL) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerMySQL) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerMySQL) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerMySQL) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerMySQL) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerMySQL) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayUpdateProducerMySQL) MarshalJSON() ([]byte, error) {
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
	if o.MysqlDbname != nil {
		toSerialize["mysql-dbname"] = o.MysqlDbname
	}
	if o.MysqlHost != nil {
		toSerialize["mysql-host"] = o.MysqlHost
	}
	if o.MysqlPassword != nil {
		toSerialize["mysql-password"] = o.MysqlPassword
	}
	if o.MysqlPort != nil {
		toSerialize["mysql-port"] = o.MysqlPort
	}
	if o.MysqlScreationStatements != nil {
		toSerialize["mysql-screation-statements"] = o.MysqlScreationStatements
	}
	if o.MysqlUsername != nil {
		toSerialize["mysql-username"] = o.MysqlUsername
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
	if o.Ssl != nil {
		toSerialize["ssl"] = o.Ssl
	}
	if o.SslCertificate != nil {
		toSerialize["ssl-certificate"] = o.SslCertificate
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

type NullableGatewayUpdateProducerMySQL struct {
	value *GatewayUpdateProducerMySQL
	isSet bool
}

func (v NullableGatewayUpdateProducerMySQL) Get() *GatewayUpdateProducerMySQL {
	return v.value
}

func (v *NullableGatewayUpdateProducerMySQL) Set(val *GatewayUpdateProducerMySQL) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerMySQL) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerMySQL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerMySQL(val *GatewayUpdateProducerMySQL) *NullableGatewayUpdateProducerMySQL {
	return &NullableGatewayUpdateProducerMySQL{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerMySQL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerMySQL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


