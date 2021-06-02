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

// GatewayCreateProducerMSSQL gatewayCreateProducerMSSQL is a command that creates mssql producer
type GatewayCreateProducerMSSQL struct {
	// Gateway url
	GatewayUrl *string `json:"gateway-url,omitempty"`
	// MSSQL Creation statements
	MssqlCreateStatements *string `json:"mssql-create-statements,omitempty"`
	// MSSQL Name
	MssqlDbname string `json:"mssql-dbname"`
	// MSSQL Host
	MssqlHost *string `json:"mssql-host,omitempty"`
	// MSSQL Password
	MssqlPassword string `json:"mssql-password"`
	// MSSQL Port
	MssqlPort *string `json:"mssql-port,omitempty"`
	// MSSQL Revocation statements
	MssqlRevocationStatements *string `json:"mssql-revocation-statements,omitempty"`
	// MSSQL Username
	MssqlUsername string `json:"mssql-username"`
	// Producer name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewGatewayCreateProducerMSSQL instantiates a new GatewayCreateProducerMSSQL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerMSSQL(mssqlDbname string, mssqlPassword string, mssqlUsername string, name string, ) *GatewayCreateProducerMSSQL {
	this := GatewayCreateProducerMSSQL{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	this.MssqlDbname = mssqlDbname
	var mssqlHost string = "127.0.0.1"
	this.MssqlHost = &mssqlHost
	this.MssqlPassword = mssqlPassword
	var mssqlPort string = "1433"
	this.MssqlPort = &mssqlPort
	this.MssqlUsername = mssqlUsername
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerMSSQLWithDefaults instantiates a new GatewayCreateProducerMSSQL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerMSSQLWithDefaults() *GatewayCreateProducerMSSQL {
	this := GatewayCreateProducerMSSQL{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	var mssqlHost string = "127.0.0.1"
	this.MssqlHost = &mssqlHost
	var mssqlPort string = "1433"
	this.MssqlPort = &mssqlPort
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetGatewayUrl returns the GatewayUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetGatewayUrl() string {
	if o == nil || o.GatewayUrl == nil {
		var ret string
		return ret
	}
	return *o.GatewayUrl
}

// GetGatewayUrlOk returns a tuple with the GatewayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetGatewayUrlOk() (*string, bool) {
	if o == nil || o.GatewayUrl == nil {
		return nil, false
	}
	return o.GatewayUrl, true
}

// HasGatewayUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasGatewayUrl() bool {
	if o != nil && o.GatewayUrl != nil {
		return true
	}

	return false
}

// SetGatewayUrl gets a reference to the given string and assigns it to the GatewayUrl field.
func (o *GatewayCreateProducerMSSQL) SetGatewayUrl(v string) {
	o.GatewayUrl = &v
}

// GetMssqlCreateStatements returns the MssqlCreateStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetMssqlCreateStatements() string {
	if o == nil || o.MssqlCreateStatements == nil {
		var ret string
		return ret
	}
	return *o.MssqlCreateStatements
}

// GetMssqlCreateStatementsOk returns a tuple with the MssqlCreateStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetMssqlCreateStatementsOk() (*string, bool) {
	if o == nil || o.MssqlCreateStatements == nil {
		return nil, false
	}
	return o.MssqlCreateStatements, true
}

// HasMssqlCreateStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasMssqlCreateStatements() bool {
	if o != nil && o.MssqlCreateStatements != nil {
		return true
	}

	return false
}

// SetMssqlCreateStatements gets a reference to the given string and assigns it to the MssqlCreateStatements field.
func (o *GatewayCreateProducerMSSQL) SetMssqlCreateStatements(v string) {
	o.MssqlCreateStatements = &v
}

// GetMssqlDbname returns the MssqlDbname field value
func (o *GatewayCreateProducerMSSQL) GetMssqlDbname() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MssqlDbname
}

// GetMssqlDbnameOk returns a tuple with the MssqlDbname field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetMssqlDbnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MssqlDbname, true
}

// SetMssqlDbname sets field value
func (o *GatewayCreateProducerMSSQL) SetMssqlDbname(v string) {
	o.MssqlDbname = v
}

// GetMssqlHost returns the MssqlHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetMssqlHost() string {
	if o == nil || o.MssqlHost == nil {
		var ret string
		return ret
	}
	return *o.MssqlHost
}

// GetMssqlHostOk returns a tuple with the MssqlHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetMssqlHostOk() (*string, bool) {
	if o == nil || o.MssqlHost == nil {
		return nil, false
	}
	return o.MssqlHost, true
}

// HasMssqlHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasMssqlHost() bool {
	if o != nil && o.MssqlHost != nil {
		return true
	}

	return false
}

// SetMssqlHost gets a reference to the given string and assigns it to the MssqlHost field.
func (o *GatewayCreateProducerMSSQL) SetMssqlHost(v string) {
	o.MssqlHost = &v
}

// GetMssqlPassword returns the MssqlPassword field value
func (o *GatewayCreateProducerMSSQL) GetMssqlPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MssqlPassword
}

// GetMssqlPasswordOk returns a tuple with the MssqlPassword field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetMssqlPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MssqlPassword, true
}

// SetMssqlPassword sets field value
func (o *GatewayCreateProducerMSSQL) SetMssqlPassword(v string) {
	o.MssqlPassword = v
}

// GetMssqlPort returns the MssqlPort field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetMssqlPort() string {
	if o == nil || o.MssqlPort == nil {
		var ret string
		return ret
	}
	return *o.MssqlPort
}

// GetMssqlPortOk returns a tuple with the MssqlPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetMssqlPortOk() (*string, bool) {
	if o == nil || o.MssqlPort == nil {
		return nil, false
	}
	return o.MssqlPort, true
}

// HasMssqlPort returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasMssqlPort() bool {
	if o != nil && o.MssqlPort != nil {
		return true
	}

	return false
}

// SetMssqlPort gets a reference to the given string and assigns it to the MssqlPort field.
func (o *GatewayCreateProducerMSSQL) SetMssqlPort(v string) {
	o.MssqlPort = &v
}

// GetMssqlRevocationStatements returns the MssqlRevocationStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetMssqlRevocationStatements() string {
	if o == nil || o.MssqlRevocationStatements == nil {
		var ret string
		return ret
	}
	return *o.MssqlRevocationStatements
}

// GetMssqlRevocationStatementsOk returns a tuple with the MssqlRevocationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetMssqlRevocationStatementsOk() (*string, bool) {
	if o == nil || o.MssqlRevocationStatements == nil {
		return nil, false
	}
	return o.MssqlRevocationStatements, true
}

// HasMssqlRevocationStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasMssqlRevocationStatements() bool {
	if o != nil && o.MssqlRevocationStatements != nil {
		return true
	}

	return false
}

// SetMssqlRevocationStatements gets a reference to the given string and assigns it to the MssqlRevocationStatements field.
func (o *GatewayCreateProducerMSSQL) SetMssqlRevocationStatements(v string) {
	o.MssqlRevocationStatements = &v
}

// GetMssqlUsername returns the MssqlUsername field value
func (o *GatewayCreateProducerMSSQL) GetMssqlUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MssqlUsername
}

// GetMssqlUsernameOk returns a tuple with the MssqlUsername field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetMssqlUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MssqlUsername, true
}

// SetMssqlUsername sets field value
func (o *GatewayCreateProducerMSSQL) SetMssqlUsername(v string) {
	o.MssqlUsername = v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerMSSQL) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerMSSQL) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayCreateProducerMSSQL) SetPassword(v string) {
	o.Password = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerMSSQL) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerMSSQL) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerMSSQL) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerMSSQL) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayCreateProducerMSSQL) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMSSQL) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerMSSQL) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayCreateProducerMSSQL) SetUsername(v string) {
	o.Username = &v
}

func (o GatewayCreateProducerMSSQL) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayUrl != nil {
		toSerialize["gateway-url"] = o.GatewayUrl
	}
	if o.MssqlCreateStatements != nil {
		toSerialize["mssql-create-statements"] = o.MssqlCreateStatements
	}
	if true {
		toSerialize["mssql-dbname"] = o.MssqlDbname
	}
	if o.MssqlHost != nil {
		toSerialize["mssql-host"] = o.MssqlHost
	}
	if true {
		toSerialize["mssql-password"] = o.MssqlPassword
	}
	if o.MssqlPort != nil {
		toSerialize["mssql-port"] = o.MssqlPort
	}
	if o.MssqlRevocationStatements != nil {
		toSerialize["mssql-revocation-statements"] = o.MssqlRevocationStatements
	}
	if true {
		toSerialize["mssql-username"] = o.MssqlUsername
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
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

type NullableGatewayCreateProducerMSSQL struct {
	value *GatewayCreateProducerMSSQL
	isSet bool
}

func (v NullableGatewayCreateProducerMSSQL) Get() *GatewayCreateProducerMSSQL {
	return v.value
}

func (v *NullableGatewayCreateProducerMSSQL) Set(val *GatewayCreateProducerMSSQL) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerMSSQL) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerMSSQL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerMSSQL(val *GatewayCreateProducerMSSQL) *NullableGatewayCreateProducerMSSQL {
	return &NullableGatewayCreateProducerMSSQL{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerMSSQL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerMSSQL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


