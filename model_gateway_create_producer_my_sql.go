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

// checks if the GatewayCreateProducerMySQL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayCreateProducerMySQL{}

// GatewayCreateProducerMySQL gatewayCreateProducerMySQL is a command that creates mysql producer
type GatewayCreateProducerMySQL struct {
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
	// MySQL Revocation statements
	MysqlRevocationStatements *string `json:"mysql-revocation-statements,omitempty"`
	// MySQL Creation statements
	MysqlScreationStatements *string `json:"mysql-screation-statements,omitempty"`
	// MySQL Username
	MysqlUsername *string `json:"mysql-username,omitempty"`
	// Producer name
	Name string `json:"name"`
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
	// Enable/Disable SSL [true/false]
	Ssl *bool `json:"ssl,omitempty"`
	// SSL connection certificate
	SslCertificate *string `json:"ssl-certificate,omitempty"`
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

// NewGatewayCreateProducerMySQL instantiates a new GatewayCreateProducerMySQL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerMySQL(name string) *GatewayCreateProducerMySQL {
	this := GatewayCreateProducerMySQL{}
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

// NewGatewayCreateProducerMySQLWithDefaults instantiates a new GatewayCreateProducerMySQL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerMySQLWithDefaults() *GatewayCreateProducerMySQL {
	this := GatewayCreateProducerMySQL{}
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
func (o *GatewayCreateProducerMySQL) GetDbServerCertificates() string {
	if o == nil || IsNil(o.DbServerCertificates) {
		var ret string
		return ret
	}
	return *o.DbServerCertificates
}

// GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetDbServerCertificatesOk() (*string, bool) {
	if o == nil || IsNil(o.DbServerCertificates) {
		return nil, false
	}
	return o.DbServerCertificates, true
}

// HasDbServerCertificates returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasDbServerCertificates() bool {
	if o != nil && !IsNil(o.DbServerCertificates) {
		return true
	}

	return false
}

// SetDbServerCertificates gets a reference to the given string and assigns it to the DbServerCertificates field.
func (o *GatewayCreateProducerMySQL) SetDbServerCertificates(v string) {
	o.DbServerCertificates = &v
}

// GetDbServerName returns the DbServerName field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetDbServerName() string {
	if o == nil || IsNil(o.DbServerName) {
		var ret string
		return ret
	}
	return *o.DbServerName
}

// GetDbServerNameOk returns a tuple with the DbServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetDbServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.DbServerName) {
		return nil, false
	}
	return o.DbServerName, true
}

// HasDbServerName returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasDbServerName() bool {
	if o != nil && !IsNil(o.DbServerName) {
		return true
	}

	return false
}

// SetDbServerName gets a reference to the given string and assigns it to the DbServerName field.
func (o *GatewayCreateProducerMySQL) SetDbServerName(v string) {
	o.DbServerName = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetDeleteProtection() string {
	if o == nil || IsNil(o.DeleteProtection) {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || IsNil(o.DeleteProtection) {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasDeleteProtection() bool {
	if o != nil && !IsNil(o.DeleteProtection) {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerMySQL) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerMySQL) SetJson(v bool) {
	o.Json = &v
}

// GetMysqlDbname returns the MysqlDbname field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetMysqlDbname() string {
	if o == nil || IsNil(o.MysqlDbname) {
		var ret string
		return ret
	}
	return *o.MysqlDbname
}

// GetMysqlDbnameOk returns a tuple with the MysqlDbname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetMysqlDbnameOk() (*string, bool) {
	if o == nil || IsNil(o.MysqlDbname) {
		return nil, false
	}
	return o.MysqlDbname, true
}

// HasMysqlDbname returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasMysqlDbname() bool {
	if o != nil && !IsNil(o.MysqlDbname) {
		return true
	}

	return false
}

// SetMysqlDbname gets a reference to the given string and assigns it to the MysqlDbname field.
func (o *GatewayCreateProducerMySQL) SetMysqlDbname(v string) {
	o.MysqlDbname = &v
}

// GetMysqlHost returns the MysqlHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetMysqlHost() string {
	if o == nil || IsNil(o.MysqlHost) {
		var ret string
		return ret
	}
	return *o.MysqlHost
}

// GetMysqlHostOk returns a tuple with the MysqlHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetMysqlHostOk() (*string, bool) {
	if o == nil || IsNil(o.MysqlHost) {
		return nil, false
	}
	return o.MysqlHost, true
}

// HasMysqlHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasMysqlHost() bool {
	if o != nil && !IsNil(o.MysqlHost) {
		return true
	}

	return false
}

// SetMysqlHost gets a reference to the given string and assigns it to the MysqlHost field.
func (o *GatewayCreateProducerMySQL) SetMysqlHost(v string) {
	o.MysqlHost = &v
}

// GetMysqlPassword returns the MysqlPassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetMysqlPassword() string {
	if o == nil || IsNil(o.MysqlPassword) {
		var ret string
		return ret
	}
	return *o.MysqlPassword
}

// GetMysqlPasswordOk returns a tuple with the MysqlPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetMysqlPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.MysqlPassword) {
		return nil, false
	}
	return o.MysqlPassword, true
}

// HasMysqlPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasMysqlPassword() bool {
	if o != nil && !IsNil(o.MysqlPassword) {
		return true
	}

	return false
}

// SetMysqlPassword gets a reference to the given string and assigns it to the MysqlPassword field.
func (o *GatewayCreateProducerMySQL) SetMysqlPassword(v string) {
	o.MysqlPassword = &v
}

// GetMysqlPort returns the MysqlPort field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetMysqlPort() string {
	if o == nil || IsNil(o.MysqlPort) {
		var ret string
		return ret
	}
	return *o.MysqlPort
}

// GetMysqlPortOk returns a tuple with the MysqlPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetMysqlPortOk() (*string, bool) {
	if o == nil || IsNil(o.MysqlPort) {
		return nil, false
	}
	return o.MysqlPort, true
}

// HasMysqlPort returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasMysqlPort() bool {
	if o != nil && !IsNil(o.MysqlPort) {
		return true
	}

	return false
}

// SetMysqlPort gets a reference to the given string and assigns it to the MysqlPort field.
func (o *GatewayCreateProducerMySQL) SetMysqlPort(v string) {
	o.MysqlPort = &v
}

// GetMysqlRevocationStatements returns the MysqlRevocationStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetMysqlRevocationStatements() string {
	if o == nil || IsNil(o.MysqlRevocationStatements) {
		var ret string
		return ret
	}
	return *o.MysqlRevocationStatements
}

// GetMysqlRevocationStatementsOk returns a tuple with the MysqlRevocationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetMysqlRevocationStatementsOk() (*string, bool) {
	if o == nil || IsNil(o.MysqlRevocationStatements) {
		return nil, false
	}
	return o.MysqlRevocationStatements, true
}

// HasMysqlRevocationStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasMysqlRevocationStatements() bool {
	if o != nil && !IsNil(o.MysqlRevocationStatements) {
		return true
	}

	return false
}

// SetMysqlRevocationStatements gets a reference to the given string and assigns it to the MysqlRevocationStatements field.
func (o *GatewayCreateProducerMySQL) SetMysqlRevocationStatements(v string) {
	o.MysqlRevocationStatements = &v
}

// GetMysqlScreationStatements returns the MysqlScreationStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetMysqlScreationStatements() string {
	if o == nil || IsNil(o.MysqlScreationStatements) {
		var ret string
		return ret
	}
	return *o.MysqlScreationStatements
}

// GetMysqlScreationStatementsOk returns a tuple with the MysqlScreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetMysqlScreationStatementsOk() (*string, bool) {
	if o == nil || IsNil(o.MysqlScreationStatements) {
		return nil, false
	}
	return o.MysqlScreationStatements, true
}

// HasMysqlScreationStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasMysqlScreationStatements() bool {
	if o != nil && !IsNil(o.MysqlScreationStatements) {
		return true
	}

	return false
}

// SetMysqlScreationStatements gets a reference to the given string and assigns it to the MysqlScreationStatements field.
func (o *GatewayCreateProducerMySQL) SetMysqlScreationStatements(v string) {
	o.MysqlScreationStatements = &v
}

// GetMysqlUsername returns the MysqlUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetMysqlUsername() string {
	if o == nil || IsNil(o.MysqlUsername) {
		var ret string
		return ret
	}
	return *o.MysqlUsername
}

// GetMysqlUsernameOk returns a tuple with the MysqlUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetMysqlUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.MysqlUsername) {
		return nil, false
	}
	return o.MysqlUsername, true
}

// HasMysqlUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasMysqlUsername() bool {
	if o != nil && !IsNil(o.MysqlUsername) {
		return true
	}

	return false
}

// SetMysqlUsername gets a reference to the given string and assigns it to the MysqlUsername field.
func (o *GatewayCreateProducerMySQL) SetMysqlUsername(v string) {
	o.MysqlUsername = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerMySQL) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerMySQL) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetProducerEncryptionKeyName() string {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProducerEncryptionKeyName) {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasProducerEncryptionKeyName() bool {
	if o != nil && !IsNil(o.ProducerEncryptionKeyName) {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerMySQL) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetSecureAccessBastionIssuer() string {
	if o == nil || IsNil(o.SecureAccessBastionIssuer) {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.SecureAccessBastionIssuer) {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasSecureAccessBastionIssuer() bool {
	if o != nil && !IsNil(o.SecureAccessBastionIssuer) {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayCreateProducerMySQL) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetSecureAccessEnable() string {
	if o == nil || IsNil(o.SecureAccessEnable) {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || IsNil(o.SecureAccessEnable) {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasSecureAccessEnable() bool {
	if o != nil && !IsNil(o.SecureAccessEnable) {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayCreateProducerMySQL) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetSecureAccessHost() []string {
	if o == nil || IsNil(o.SecureAccessHost) {
		var ret []string
		return ret
	}
	return o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetSecureAccessHostOk() ([]string, bool) {
	if o == nil || IsNil(o.SecureAccessHost) {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasSecureAccessHost() bool {
	if o != nil && !IsNil(o.SecureAccessHost) {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *GatewayCreateProducerMySQL) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetSecureAccessWeb() bool {
	if o == nil || IsNil(o.SecureAccessWeb) {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureAccessWeb) {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasSecureAccessWeb() bool {
	if o != nil && !IsNil(o.SecureAccessWeb) {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayCreateProducerMySQL) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetSsl() bool {
	if o == nil || IsNil(o.Ssl) {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetSslOk() (*bool, bool) {
	if o == nil || IsNil(o.Ssl) {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasSsl() bool {
	if o != nil && !IsNil(o.Ssl) {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *GatewayCreateProducerMySQL) SetSsl(v bool) {
	o.Ssl = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetSslCertificate() string {
	if o == nil || IsNil(o.SslCertificate) {
		var ret string
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetSslCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.SslCertificate) {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasSslCertificate() bool {
	if o != nil && !IsNil(o.SslCertificate) {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given string and assigns it to the SslCertificate field.
func (o *GatewayCreateProducerMySQL) SetSslCertificate(v string) {
	o.SslCertificate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerMySQL) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerMySQL) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerMySQL) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetUidToken() string {
	if o == nil || IsNil(o.UidToken) {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetUidTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UidToken) {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasUidToken() bool {
	if o != nil && !IsNil(o.UidToken) {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerMySQL) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerMySQL) GetUserTtl() string {
	if o == nil || IsNil(o.UserTtl) {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMySQL) GetUserTtlOk() (*string, bool) {
	if o == nil || IsNil(o.UserTtl) {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerMySQL) HasUserTtl() bool {
	if o != nil && !IsNil(o.UserTtl) {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerMySQL) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerMySQL) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayCreateProducerMySQL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DbServerCertificates) {
		toSerialize["db-server-certificates"] = o.DbServerCertificates
	}
	if !IsNil(o.DbServerName) {
		toSerialize["db-server-name"] = o.DbServerName
	}
	if !IsNil(o.DeleteProtection) {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.MysqlDbname) {
		toSerialize["mysql-dbname"] = o.MysqlDbname
	}
	if !IsNil(o.MysqlHost) {
		toSerialize["mysql-host"] = o.MysqlHost
	}
	if !IsNil(o.MysqlPassword) {
		toSerialize["mysql-password"] = o.MysqlPassword
	}
	if !IsNil(o.MysqlPort) {
		toSerialize["mysql-port"] = o.MysqlPort
	}
	if !IsNil(o.MysqlRevocationStatements) {
		toSerialize["mysql-revocation-statements"] = o.MysqlRevocationStatements
	}
	if !IsNil(o.MysqlScreationStatements) {
		toSerialize["mysql-screation-statements"] = o.MysqlScreationStatements
	}
	if !IsNil(o.MysqlUsername) {
		toSerialize["mysql-username"] = o.MysqlUsername
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ProducerEncryptionKeyName) {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if !IsNil(o.SecureAccessBastionIssuer) {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if !IsNil(o.SecureAccessEnable) {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if !IsNil(o.SecureAccessHost) {
		toSerialize["secure-access-host"] = o.SecureAccessHost
	}
	if !IsNil(o.SecureAccessWeb) {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
	}
	if !IsNil(o.Ssl) {
		toSerialize["ssl"] = o.Ssl
	}
	if !IsNil(o.SslCertificate) {
		toSerialize["ssl-certificate"] = o.SslCertificate
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

type NullableGatewayCreateProducerMySQL struct {
	value *GatewayCreateProducerMySQL
	isSet bool
}

func (v NullableGatewayCreateProducerMySQL) Get() *GatewayCreateProducerMySQL {
	return v.value
}

func (v *NullableGatewayCreateProducerMySQL) Set(val *GatewayCreateProducerMySQL) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerMySQL) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerMySQL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerMySQL(val *GatewayCreateProducerMySQL) *NullableGatewayCreateProducerMySQL {
	return &NullableGatewayCreateProducerMySQL{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerMySQL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerMySQL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


