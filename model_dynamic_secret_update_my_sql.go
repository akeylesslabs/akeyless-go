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

// DynamicSecretUpdateMySql dynamicSecretUpdateMySql is a command that updates mysql dynamic secret
type DynamicSecretUpdateMySql struct {
	// (Optional) DB server certificates
	DbServerCertificates *string `json:"db-server-certificates,omitempty"`
	// (Optional) Server name for certificate verification
	DbServerName *string `json:"db-server-name,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
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
	// MySQL Revocation statements
	MysqlRevocationStatements *string `json:"mysql-revocation-statements,omitempty"`
	// MySQL Creation statements
	MysqlScreationStatements *string `json:"mysql-screation-statements,omitempty"`
	// MySQL Username
	MysqlUsername *string `json:"mysql-username,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret new name
	NewName *string `json:"new-name,omitempty"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts)
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// Enable/Disable SSL [true/false]
	Ssl *bool `json:"ssl,omitempty"`
	// SSL connection certificate
	SslCertificate *string `json:"ssl-certificate,omitempty"`
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

// NewDynamicSecretUpdateMySql instantiates a new DynamicSecretUpdateMySql object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdateMySql(name string, ) *DynamicSecretUpdateMySql {
	this := DynamicSecretUpdateMySql{}
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

// NewDynamicSecretUpdateMySqlWithDefaults instantiates a new DynamicSecretUpdateMySql object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdateMySqlWithDefaults() *DynamicSecretUpdateMySql {
	this := DynamicSecretUpdateMySql{}
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
func (o *DynamicSecretUpdateMySql) GetDbServerCertificates() string {
	if o == nil || o.DbServerCertificates == nil {
		var ret string
		return ret
	}
	return *o.DbServerCertificates
}

// GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetDbServerCertificatesOk() (*string, bool) {
	if o == nil || o.DbServerCertificates == nil {
		return nil, false
	}
	return o.DbServerCertificates, true
}

// HasDbServerCertificates returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasDbServerCertificates() bool {
	if o != nil && o.DbServerCertificates != nil {
		return true
	}

	return false
}

// SetDbServerCertificates gets a reference to the given string and assigns it to the DbServerCertificates field.
func (o *DynamicSecretUpdateMySql) SetDbServerCertificates(v string) {
	o.DbServerCertificates = &v
}

// GetDbServerName returns the DbServerName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetDbServerName() string {
	if o == nil || o.DbServerName == nil {
		var ret string
		return ret
	}
	return *o.DbServerName
}

// GetDbServerNameOk returns a tuple with the DbServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetDbServerNameOk() (*string, bool) {
	if o == nil || o.DbServerName == nil {
		return nil, false
	}
	return o.DbServerName, true
}

// HasDbServerName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasDbServerName() bool {
	if o != nil && o.DbServerName != nil {
		return true
	}

	return false
}

// SetDbServerName gets a reference to the given string and assigns it to the DbServerName field.
func (o *DynamicSecretUpdateMySql) SetDbServerName(v string) {
	o.DbServerName = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretUpdateMySql) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretUpdateMySql) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretUpdateMySql) SetJson(v bool) {
	o.Json = &v
}

// GetMysqlDbname returns the MysqlDbname field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetMysqlDbname() string {
	if o == nil || o.MysqlDbname == nil {
		var ret string
		return ret
	}
	return *o.MysqlDbname
}

// GetMysqlDbnameOk returns a tuple with the MysqlDbname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetMysqlDbnameOk() (*string, bool) {
	if o == nil || o.MysqlDbname == nil {
		return nil, false
	}
	return o.MysqlDbname, true
}

// HasMysqlDbname returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasMysqlDbname() bool {
	if o != nil && o.MysqlDbname != nil {
		return true
	}

	return false
}

// SetMysqlDbname gets a reference to the given string and assigns it to the MysqlDbname field.
func (o *DynamicSecretUpdateMySql) SetMysqlDbname(v string) {
	o.MysqlDbname = &v
}

// GetMysqlHost returns the MysqlHost field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetMysqlHost() string {
	if o == nil || o.MysqlHost == nil {
		var ret string
		return ret
	}
	return *o.MysqlHost
}

// GetMysqlHostOk returns a tuple with the MysqlHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetMysqlHostOk() (*string, bool) {
	if o == nil || o.MysqlHost == nil {
		return nil, false
	}
	return o.MysqlHost, true
}

// HasMysqlHost returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasMysqlHost() bool {
	if o != nil && o.MysqlHost != nil {
		return true
	}

	return false
}

// SetMysqlHost gets a reference to the given string and assigns it to the MysqlHost field.
func (o *DynamicSecretUpdateMySql) SetMysqlHost(v string) {
	o.MysqlHost = &v
}

// GetMysqlPassword returns the MysqlPassword field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetMysqlPassword() string {
	if o == nil || o.MysqlPassword == nil {
		var ret string
		return ret
	}
	return *o.MysqlPassword
}

// GetMysqlPasswordOk returns a tuple with the MysqlPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetMysqlPasswordOk() (*string, bool) {
	if o == nil || o.MysqlPassword == nil {
		return nil, false
	}
	return o.MysqlPassword, true
}

// HasMysqlPassword returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasMysqlPassword() bool {
	if o != nil && o.MysqlPassword != nil {
		return true
	}

	return false
}

// SetMysqlPassword gets a reference to the given string and assigns it to the MysqlPassword field.
func (o *DynamicSecretUpdateMySql) SetMysqlPassword(v string) {
	o.MysqlPassword = &v
}

// GetMysqlPort returns the MysqlPort field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetMysqlPort() string {
	if o == nil || o.MysqlPort == nil {
		var ret string
		return ret
	}
	return *o.MysqlPort
}

// GetMysqlPortOk returns a tuple with the MysqlPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetMysqlPortOk() (*string, bool) {
	if o == nil || o.MysqlPort == nil {
		return nil, false
	}
	return o.MysqlPort, true
}

// HasMysqlPort returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasMysqlPort() bool {
	if o != nil && o.MysqlPort != nil {
		return true
	}

	return false
}

// SetMysqlPort gets a reference to the given string and assigns it to the MysqlPort field.
func (o *DynamicSecretUpdateMySql) SetMysqlPort(v string) {
	o.MysqlPort = &v
}

// GetMysqlRevocationStatements returns the MysqlRevocationStatements field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetMysqlRevocationStatements() string {
	if o == nil || o.MysqlRevocationStatements == nil {
		var ret string
		return ret
	}
	return *o.MysqlRevocationStatements
}

// GetMysqlRevocationStatementsOk returns a tuple with the MysqlRevocationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetMysqlRevocationStatementsOk() (*string, bool) {
	if o == nil || o.MysqlRevocationStatements == nil {
		return nil, false
	}
	return o.MysqlRevocationStatements, true
}

// HasMysqlRevocationStatements returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasMysqlRevocationStatements() bool {
	if o != nil && o.MysqlRevocationStatements != nil {
		return true
	}

	return false
}

// SetMysqlRevocationStatements gets a reference to the given string and assigns it to the MysqlRevocationStatements field.
func (o *DynamicSecretUpdateMySql) SetMysqlRevocationStatements(v string) {
	o.MysqlRevocationStatements = &v
}

// GetMysqlScreationStatements returns the MysqlScreationStatements field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetMysqlScreationStatements() string {
	if o == nil || o.MysqlScreationStatements == nil {
		var ret string
		return ret
	}
	return *o.MysqlScreationStatements
}

// GetMysqlScreationStatementsOk returns a tuple with the MysqlScreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetMysqlScreationStatementsOk() (*string, bool) {
	if o == nil || o.MysqlScreationStatements == nil {
		return nil, false
	}
	return o.MysqlScreationStatements, true
}

// HasMysqlScreationStatements returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasMysqlScreationStatements() bool {
	if o != nil && o.MysqlScreationStatements != nil {
		return true
	}

	return false
}

// SetMysqlScreationStatements gets a reference to the given string and assigns it to the MysqlScreationStatements field.
func (o *DynamicSecretUpdateMySql) SetMysqlScreationStatements(v string) {
	o.MysqlScreationStatements = &v
}

// GetMysqlUsername returns the MysqlUsername field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetMysqlUsername() string {
	if o == nil || o.MysqlUsername == nil {
		var ret string
		return ret
	}
	return *o.MysqlUsername
}

// GetMysqlUsernameOk returns a tuple with the MysqlUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetMysqlUsernameOk() (*string, bool) {
	if o == nil || o.MysqlUsername == nil {
		return nil, false
	}
	return o.MysqlUsername, true
}

// HasMysqlUsername returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasMysqlUsername() bool {
	if o != nil && o.MysqlUsername != nil {
		return true
	}

	return false
}

// SetMysqlUsername gets a reference to the given string and assigns it to the MysqlUsername field.
func (o *DynamicSecretUpdateMySql) SetMysqlUsername(v string) {
	o.MysqlUsername = &v
}

// GetName returns the Name field value
func (o *DynamicSecretUpdateMySql) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretUpdateMySql) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DynamicSecretUpdateMySql) SetNewName(v string) {
	o.NewName = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *DynamicSecretUpdateMySql) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretUpdateMySql) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *DynamicSecretUpdateMySql) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DynamicSecretUpdateMySql) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *DynamicSecretUpdateMySql) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *DynamicSecretUpdateMySql) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetSsl() bool {
	if o == nil || o.Ssl == nil {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetSslOk() (*bool, bool) {
	if o == nil || o.Ssl == nil {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasSsl() bool {
	if o != nil && o.Ssl != nil {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *DynamicSecretUpdateMySql) SetSsl(v bool) {
	o.Ssl = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetSslCertificate() string {
	if o == nil || o.SslCertificate == nil {
		var ret string
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetSslCertificateOk() (*string, bool) {
	if o == nil || o.SslCertificate == nil {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasSslCertificate() bool {
	if o != nil && o.SslCertificate != nil {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given string and assigns it to the SslCertificate field.
func (o *DynamicSecretUpdateMySql) SetSslCertificate(v string) {
	o.SslCertificate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretUpdateMySql) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretUpdateMySql) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretUpdateMySql) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretUpdateMySql) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateMySql) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateMySql) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateMySql) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretUpdateMySql) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretUpdateMySql) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
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
	if o.MysqlRevocationStatements != nil {
		toSerialize["mysql-revocation-statements"] = o.MysqlRevocationStatements
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
	if o.PasswordLength != nil {
		toSerialize["password-length"] = o.PasswordLength
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

type NullableDynamicSecretUpdateMySql struct {
	value *DynamicSecretUpdateMySql
	isSet bool
}

func (v NullableDynamicSecretUpdateMySql) Get() *DynamicSecretUpdateMySql {
	return v.value
}

func (v *NullableDynamicSecretUpdateMySql) Set(val *DynamicSecretUpdateMySql) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdateMySql) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdateMySql) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdateMySql(val *DynamicSecretUpdateMySql) *NullableDynamicSecretUpdateMySql {
	return &NullableDynamicSecretUpdateMySql{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdateMySql) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdateMySql) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


