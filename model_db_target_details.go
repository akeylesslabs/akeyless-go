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

// DbTargetDetails DbTargetDetails
type DbTargetDetails struct {
	DbHostName *string `json:"db_host_name,omitempty"`
	DbName *string `json:"db_name,omitempty"`
	DbPort *string `json:"db_port,omitempty"`
	// (Optional) Private Key in PEM format
	DbPrivateKey *string `json:"db_private_key,omitempty"`
	DbPrivateKeyPassphrase *string `json:"db_private_key_passphrase,omitempty"`
	DbPwd *string `json:"db_pwd,omitempty"`
	// (Optional) DBServerCertificates defines the set of root certificate authorities that clients use when verifying server certificates. If DBServerCertificates is empty, TLS uses the host's root CA set.
	DbServerCertificates *string `json:"db_server_certificates,omitempty"`
	// (Optional) ServerName is used to verify the hostname on the returned certificates unless InsecureSkipVerify is given. It is also included in the client's handshake to support virtual hosting unless it is an IP address.
	DbServerName *string `json:"db_server_name,omitempty"`
	DbUserName *string `json:"db_user_name,omitempty"`
	SfAccount *string `json:"sf_account,omitempty"`
	// (Optional) SSLConnectionCertificate defines the certificate for SSL connection. Must be base64 certificate loaded by UI using file loader field
	SslConnectionCertificate *string `json:"ssl_connection_certificate,omitempty"`
	// (Optional) SSLConnectionMode defines if SSL mode will be used to connect to DB
	SslConnectionMode *bool `json:"ssl_connection_mode,omitempty"`
}

// NewDbTargetDetails instantiates a new DbTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbTargetDetails() *DbTargetDetails {
	this := DbTargetDetails{}
	return &this
}

// NewDbTargetDetailsWithDefaults instantiates a new DbTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbTargetDetailsWithDefaults() *DbTargetDetails {
	this := DbTargetDetails{}
	return &this
}

// GetDbHostName returns the DbHostName field value if set, zero value otherwise.
func (o *DbTargetDetails) GetDbHostName() string {
	if o == nil || o.DbHostName == nil {
		var ret string
		return ret
	}
	return *o.DbHostName
}

// GetDbHostNameOk returns a tuple with the DbHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetDbHostNameOk() (*string, bool) {
	if o == nil || o.DbHostName == nil {
		return nil, false
	}
	return o.DbHostName, true
}

// HasDbHostName returns a boolean if a field has been set.
func (o *DbTargetDetails) HasDbHostName() bool {
	if o != nil && o.DbHostName != nil {
		return true
	}

	return false
}

// SetDbHostName gets a reference to the given string and assigns it to the DbHostName field.
func (o *DbTargetDetails) SetDbHostName(v string) {
	o.DbHostName = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *DbTargetDetails) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *DbTargetDetails) HasDbName() bool {
	if o != nil && o.DbName != nil {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *DbTargetDetails) SetDbName(v string) {
	o.DbName = &v
}

// GetDbPort returns the DbPort field value if set, zero value otherwise.
func (o *DbTargetDetails) GetDbPort() string {
	if o == nil || o.DbPort == nil {
		var ret string
		return ret
	}
	return *o.DbPort
}

// GetDbPortOk returns a tuple with the DbPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetDbPortOk() (*string, bool) {
	if o == nil || o.DbPort == nil {
		return nil, false
	}
	return o.DbPort, true
}

// HasDbPort returns a boolean if a field has been set.
func (o *DbTargetDetails) HasDbPort() bool {
	if o != nil && o.DbPort != nil {
		return true
	}

	return false
}

// SetDbPort gets a reference to the given string and assigns it to the DbPort field.
func (o *DbTargetDetails) SetDbPort(v string) {
	o.DbPort = &v
}

// GetDbPrivateKey returns the DbPrivateKey field value if set, zero value otherwise.
func (o *DbTargetDetails) GetDbPrivateKey() string {
	if o == nil || o.DbPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.DbPrivateKey
}

// GetDbPrivateKeyOk returns a tuple with the DbPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetDbPrivateKeyOk() (*string, bool) {
	if o == nil || o.DbPrivateKey == nil {
		return nil, false
	}
	return o.DbPrivateKey, true
}

// HasDbPrivateKey returns a boolean if a field has been set.
func (o *DbTargetDetails) HasDbPrivateKey() bool {
	if o != nil && o.DbPrivateKey != nil {
		return true
	}

	return false
}

// SetDbPrivateKey gets a reference to the given string and assigns it to the DbPrivateKey field.
func (o *DbTargetDetails) SetDbPrivateKey(v string) {
	o.DbPrivateKey = &v
}

// GetDbPrivateKeyPassphrase returns the DbPrivateKeyPassphrase field value if set, zero value otherwise.
func (o *DbTargetDetails) GetDbPrivateKeyPassphrase() string {
	if o == nil || o.DbPrivateKeyPassphrase == nil {
		var ret string
		return ret
	}
	return *o.DbPrivateKeyPassphrase
}

// GetDbPrivateKeyPassphraseOk returns a tuple with the DbPrivateKeyPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetDbPrivateKeyPassphraseOk() (*string, bool) {
	if o == nil || o.DbPrivateKeyPassphrase == nil {
		return nil, false
	}
	return o.DbPrivateKeyPassphrase, true
}

// HasDbPrivateKeyPassphrase returns a boolean if a field has been set.
func (o *DbTargetDetails) HasDbPrivateKeyPassphrase() bool {
	if o != nil && o.DbPrivateKeyPassphrase != nil {
		return true
	}

	return false
}

// SetDbPrivateKeyPassphrase gets a reference to the given string and assigns it to the DbPrivateKeyPassphrase field.
func (o *DbTargetDetails) SetDbPrivateKeyPassphrase(v string) {
	o.DbPrivateKeyPassphrase = &v
}

// GetDbPwd returns the DbPwd field value if set, zero value otherwise.
func (o *DbTargetDetails) GetDbPwd() string {
	if o == nil || o.DbPwd == nil {
		var ret string
		return ret
	}
	return *o.DbPwd
}

// GetDbPwdOk returns a tuple with the DbPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetDbPwdOk() (*string, bool) {
	if o == nil || o.DbPwd == nil {
		return nil, false
	}
	return o.DbPwd, true
}

// HasDbPwd returns a boolean if a field has been set.
func (o *DbTargetDetails) HasDbPwd() bool {
	if o != nil && o.DbPwd != nil {
		return true
	}

	return false
}

// SetDbPwd gets a reference to the given string and assigns it to the DbPwd field.
func (o *DbTargetDetails) SetDbPwd(v string) {
	o.DbPwd = &v
}

// GetDbServerCertificates returns the DbServerCertificates field value if set, zero value otherwise.
func (o *DbTargetDetails) GetDbServerCertificates() string {
	if o == nil || o.DbServerCertificates == nil {
		var ret string
		return ret
	}
	return *o.DbServerCertificates
}

// GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetDbServerCertificatesOk() (*string, bool) {
	if o == nil || o.DbServerCertificates == nil {
		return nil, false
	}
	return o.DbServerCertificates, true
}

// HasDbServerCertificates returns a boolean if a field has been set.
func (o *DbTargetDetails) HasDbServerCertificates() bool {
	if o != nil && o.DbServerCertificates != nil {
		return true
	}

	return false
}

// SetDbServerCertificates gets a reference to the given string and assigns it to the DbServerCertificates field.
func (o *DbTargetDetails) SetDbServerCertificates(v string) {
	o.DbServerCertificates = &v
}

// GetDbServerName returns the DbServerName field value if set, zero value otherwise.
func (o *DbTargetDetails) GetDbServerName() string {
	if o == nil || o.DbServerName == nil {
		var ret string
		return ret
	}
	return *o.DbServerName
}

// GetDbServerNameOk returns a tuple with the DbServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetDbServerNameOk() (*string, bool) {
	if o == nil || o.DbServerName == nil {
		return nil, false
	}
	return o.DbServerName, true
}

// HasDbServerName returns a boolean if a field has been set.
func (o *DbTargetDetails) HasDbServerName() bool {
	if o != nil && o.DbServerName != nil {
		return true
	}

	return false
}

// SetDbServerName gets a reference to the given string and assigns it to the DbServerName field.
func (o *DbTargetDetails) SetDbServerName(v string) {
	o.DbServerName = &v
}

// GetDbUserName returns the DbUserName field value if set, zero value otherwise.
func (o *DbTargetDetails) GetDbUserName() string {
	if o == nil || o.DbUserName == nil {
		var ret string
		return ret
	}
	return *o.DbUserName
}

// GetDbUserNameOk returns a tuple with the DbUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetDbUserNameOk() (*string, bool) {
	if o == nil || o.DbUserName == nil {
		return nil, false
	}
	return o.DbUserName, true
}

// HasDbUserName returns a boolean if a field has been set.
func (o *DbTargetDetails) HasDbUserName() bool {
	if o != nil && o.DbUserName != nil {
		return true
	}

	return false
}

// SetDbUserName gets a reference to the given string and assigns it to the DbUserName field.
func (o *DbTargetDetails) SetDbUserName(v string) {
	o.DbUserName = &v
}

// GetSfAccount returns the SfAccount field value if set, zero value otherwise.
func (o *DbTargetDetails) GetSfAccount() string {
	if o == nil || o.SfAccount == nil {
		var ret string
		return ret
	}
	return *o.SfAccount
}

// GetSfAccountOk returns a tuple with the SfAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetSfAccountOk() (*string, bool) {
	if o == nil || o.SfAccount == nil {
		return nil, false
	}
	return o.SfAccount, true
}

// HasSfAccount returns a boolean if a field has been set.
func (o *DbTargetDetails) HasSfAccount() bool {
	if o != nil && o.SfAccount != nil {
		return true
	}

	return false
}

// SetSfAccount gets a reference to the given string and assigns it to the SfAccount field.
func (o *DbTargetDetails) SetSfAccount(v string) {
	o.SfAccount = &v
}

// GetSslConnectionCertificate returns the SslConnectionCertificate field value if set, zero value otherwise.
func (o *DbTargetDetails) GetSslConnectionCertificate() string {
	if o == nil || o.SslConnectionCertificate == nil {
		var ret string
		return ret
	}
	return *o.SslConnectionCertificate
}

// GetSslConnectionCertificateOk returns a tuple with the SslConnectionCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetSslConnectionCertificateOk() (*string, bool) {
	if o == nil || o.SslConnectionCertificate == nil {
		return nil, false
	}
	return o.SslConnectionCertificate, true
}

// HasSslConnectionCertificate returns a boolean if a field has been set.
func (o *DbTargetDetails) HasSslConnectionCertificate() bool {
	if o != nil && o.SslConnectionCertificate != nil {
		return true
	}

	return false
}

// SetSslConnectionCertificate gets a reference to the given string and assigns it to the SslConnectionCertificate field.
func (o *DbTargetDetails) SetSslConnectionCertificate(v string) {
	o.SslConnectionCertificate = &v
}

// GetSslConnectionMode returns the SslConnectionMode field value if set, zero value otherwise.
func (o *DbTargetDetails) GetSslConnectionMode() bool {
	if o == nil || o.SslConnectionMode == nil {
		var ret bool
		return ret
	}
	return *o.SslConnectionMode
}

// GetSslConnectionModeOk returns a tuple with the SslConnectionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbTargetDetails) GetSslConnectionModeOk() (*bool, bool) {
	if o == nil || o.SslConnectionMode == nil {
		return nil, false
	}
	return o.SslConnectionMode, true
}

// HasSslConnectionMode returns a boolean if a field has been set.
func (o *DbTargetDetails) HasSslConnectionMode() bool {
	if o != nil && o.SslConnectionMode != nil {
		return true
	}

	return false
}

// SetSslConnectionMode gets a reference to the given bool and assigns it to the SslConnectionMode field.
func (o *DbTargetDetails) SetSslConnectionMode(v bool) {
	o.SslConnectionMode = &v
}

func (o DbTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DbHostName != nil {
		toSerialize["db_host_name"] = o.DbHostName
	}
	if o.DbName != nil {
		toSerialize["db_name"] = o.DbName
	}
	if o.DbPort != nil {
		toSerialize["db_port"] = o.DbPort
	}
	if o.DbPrivateKey != nil {
		toSerialize["db_private_key"] = o.DbPrivateKey
	}
	if o.DbPrivateKeyPassphrase != nil {
		toSerialize["db_private_key_passphrase"] = o.DbPrivateKeyPassphrase
	}
	if o.DbPwd != nil {
		toSerialize["db_pwd"] = o.DbPwd
	}
	if o.DbServerCertificates != nil {
		toSerialize["db_server_certificates"] = o.DbServerCertificates
	}
	if o.DbServerName != nil {
		toSerialize["db_server_name"] = o.DbServerName
	}
	if o.DbUserName != nil {
		toSerialize["db_user_name"] = o.DbUserName
	}
	if o.SfAccount != nil {
		toSerialize["sf_account"] = o.SfAccount
	}
	if o.SslConnectionCertificate != nil {
		toSerialize["ssl_connection_certificate"] = o.SslConnectionCertificate
	}
	if o.SslConnectionMode != nil {
		toSerialize["ssl_connection_mode"] = o.SslConnectionMode
	}
	return json.Marshal(toSerialize)
}

type NullableDbTargetDetails struct {
	value *DbTargetDetails
	isSet bool
}

func (v NullableDbTargetDetails) Get() *DbTargetDetails {
	return v.value
}

func (v *NullableDbTargetDetails) Set(val *DbTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDbTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDbTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbTargetDetails(val *DbTargetDetails) *NullableDbTargetDetails {
	return &NullableDbTargetDetails{value: val, isSet: true}
}

func (v NullableDbTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


