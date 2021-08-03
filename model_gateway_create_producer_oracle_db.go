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

// GatewayCreateProducerOracleDb gatewayCreateProducerOracleDb is a command that creates oracle db producer
type GatewayCreateProducerOracleDb struct {
	// (Optional) DB server certificates
	DbServerCertificates *string `json:"db-server-certificates,omitempty"`
	// (Optional) Server name for certificate verification
	DbServerName *string `json:"db-server-name,omitempty"`
	// Producer name
	Name string `json:"name"`
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
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewGatewayCreateProducerOracleDb instantiates a new GatewayCreateProducerOracleDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerOracleDb(name string, ) *GatewayCreateProducerOracleDb {
	this := GatewayCreateProducerOracleDb{}
	this.Name = name
	var oracleHost string = "127.0.0.1"
	this.OracleHost = &oracleHost
	var oraclePort string = "1521"
	this.OraclePort = &oraclePort
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerOracleDbWithDefaults instantiates a new GatewayCreateProducerOracleDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerOracleDbWithDefaults() *GatewayCreateProducerOracleDb {
	this := GatewayCreateProducerOracleDb{}
	var oracleHost string = "127.0.0.1"
	this.OracleHost = &oracleHost
	var oraclePort string = "1521"
	this.OraclePort = &oraclePort
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDbServerCertificates returns the DbServerCertificates field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetDbServerCertificates() string {
	if o == nil || o.DbServerCertificates == nil {
		var ret string
		return ret
	}
	return *o.DbServerCertificates
}

// GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetDbServerCertificatesOk() (*string, bool) {
	if o == nil || o.DbServerCertificates == nil {
		return nil, false
	}
	return o.DbServerCertificates, true
}

// HasDbServerCertificates returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasDbServerCertificates() bool {
	if o != nil && o.DbServerCertificates != nil {
		return true
	}

	return false
}

// SetDbServerCertificates gets a reference to the given string and assigns it to the DbServerCertificates field.
func (o *GatewayCreateProducerOracleDb) SetDbServerCertificates(v string) {
	o.DbServerCertificates = &v
}

// GetDbServerName returns the DbServerName field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetDbServerName() string {
	if o == nil || o.DbServerName == nil {
		var ret string
		return ret
	}
	return *o.DbServerName
}

// GetDbServerNameOk returns a tuple with the DbServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetDbServerNameOk() (*string, bool) {
	if o == nil || o.DbServerName == nil {
		return nil, false
	}
	return o.DbServerName, true
}

// HasDbServerName returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasDbServerName() bool {
	if o != nil && o.DbServerName != nil {
		return true
	}

	return false
}

// SetDbServerName gets a reference to the given string and assigns it to the DbServerName field.
func (o *GatewayCreateProducerOracleDb) SetDbServerName(v string) {
	o.DbServerName = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerOracleDb) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerOracleDb) SetName(v string) {
	o.Name = v
}

// GetOracleHost returns the OracleHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetOracleHost() string {
	if o == nil || o.OracleHost == nil {
		var ret string
		return ret
	}
	return *o.OracleHost
}

// GetOracleHostOk returns a tuple with the OracleHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetOracleHostOk() (*string, bool) {
	if o == nil || o.OracleHost == nil {
		return nil, false
	}
	return o.OracleHost, true
}

// HasOracleHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasOracleHost() bool {
	if o != nil && o.OracleHost != nil {
		return true
	}

	return false
}

// SetOracleHost gets a reference to the given string and assigns it to the OracleHost field.
func (o *GatewayCreateProducerOracleDb) SetOracleHost(v string) {
	o.OracleHost = &v
}

// GetOraclePassword returns the OraclePassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetOraclePassword() string {
	if o == nil || o.OraclePassword == nil {
		var ret string
		return ret
	}
	return *o.OraclePassword
}

// GetOraclePasswordOk returns a tuple with the OraclePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetOraclePasswordOk() (*string, bool) {
	if o == nil || o.OraclePassword == nil {
		return nil, false
	}
	return o.OraclePassword, true
}

// HasOraclePassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasOraclePassword() bool {
	if o != nil && o.OraclePassword != nil {
		return true
	}

	return false
}

// SetOraclePassword gets a reference to the given string and assigns it to the OraclePassword field.
func (o *GatewayCreateProducerOracleDb) SetOraclePassword(v string) {
	o.OraclePassword = &v
}

// GetOraclePort returns the OraclePort field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetOraclePort() string {
	if o == nil || o.OraclePort == nil {
		var ret string
		return ret
	}
	return *o.OraclePort
}

// GetOraclePortOk returns a tuple with the OraclePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetOraclePortOk() (*string, bool) {
	if o == nil || o.OraclePort == nil {
		return nil, false
	}
	return o.OraclePort, true
}

// HasOraclePort returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasOraclePort() bool {
	if o != nil && o.OraclePort != nil {
		return true
	}

	return false
}

// SetOraclePort gets a reference to the given string and assigns it to the OraclePort field.
func (o *GatewayCreateProducerOracleDb) SetOraclePort(v string) {
	o.OraclePort = &v
}

// GetOracleScreationStatements returns the OracleScreationStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetOracleScreationStatements() string {
	if o == nil || o.OracleScreationStatements == nil {
		var ret string
		return ret
	}
	return *o.OracleScreationStatements
}

// GetOracleScreationStatementsOk returns a tuple with the OracleScreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetOracleScreationStatementsOk() (*string, bool) {
	if o == nil || o.OracleScreationStatements == nil {
		return nil, false
	}
	return o.OracleScreationStatements, true
}

// HasOracleScreationStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasOracleScreationStatements() bool {
	if o != nil && o.OracleScreationStatements != nil {
		return true
	}

	return false
}

// SetOracleScreationStatements gets a reference to the given string and assigns it to the OracleScreationStatements field.
func (o *GatewayCreateProducerOracleDb) SetOracleScreationStatements(v string) {
	o.OracleScreationStatements = &v
}

// GetOracleServiceName returns the OracleServiceName field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetOracleServiceName() string {
	if o == nil || o.OracleServiceName == nil {
		var ret string
		return ret
	}
	return *o.OracleServiceName
}

// GetOracleServiceNameOk returns a tuple with the OracleServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetOracleServiceNameOk() (*string, bool) {
	if o == nil || o.OracleServiceName == nil {
		return nil, false
	}
	return o.OracleServiceName, true
}

// HasOracleServiceName returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasOracleServiceName() bool {
	if o != nil && o.OracleServiceName != nil {
		return true
	}

	return false
}

// SetOracleServiceName gets a reference to the given string and assigns it to the OracleServiceName field.
func (o *GatewayCreateProducerOracleDb) SetOracleServiceName(v string) {
	o.OracleServiceName = &v
}

// GetOracleUsername returns the OracleUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetOracleUsername() string {
	if o == nil || o.OracleUsername == nil {
		var ret string
		return ret
	}
	return *o.OracleUsername
}

// GetOracleUsernameOk returns a tuple with the OracleUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetOracleUsernameOk() (*string, bool) {
	if o == nil || o.OracleUsername == nil {
		return nil, false
	}
	return o.OracleUsername, true
}

// HasOracleUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasOracleUsername() bool {
	if o != nil && o.OracleUsername != nil {
		return true
	}

	return false
}

// SetOracleUsername gets a reference to the given string and assigns it to the OracleUsername field.
func (o *GatewayCreateProducerOracleDb) SetOracleUsername(v string) {
	o.OracleUsername = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayCreateProducerOracleDb) SetPassword(v string) {
	o.Password = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerOracleDb) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerOracleDb) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerOracleDb) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerOracleDb) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerOracleDb) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayCreateProducerOracleDb) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerOracleDb) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerOracleDb) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayCreateProducerOracleDb) SetUsername(v string) {
	o.Username = &v
}

func (o GatewayCreateProducerOracleDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DbServerCertificates != nil {
		toSerialize["db-server-certificates"] = o.DbServerCertificates
	}
	if o.DbServerName != nil {
		toSerialize["db-server-name"] = o.DbServerName
	}
	if true {
		toSerialize["name"] = o.Name
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
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerOracleDb struct {
	value *GatewayCreateProducerOracleDb
	isSet bool
}

func (v NullableGatewayCreateProducerOracleDb) Get() *GatewayCreateProducerOracleDb {
	return v.value
}

func (v *NullableGatewayCreateProducerOracleDb) Set(val *GatewayCreateProducerOracleDb) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerOracleDb) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerOracleDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerOracleDb(val *GatewayCreateProducerOracleDb) *NullableGatewayCreateProducerOracleDb {
	return &NullableGatewayCreateProducerOracleDb{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerOracleDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerOracleDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

