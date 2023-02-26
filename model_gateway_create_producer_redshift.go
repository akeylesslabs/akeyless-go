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

// GatewayCreateProducerRedshift gatewayCreateProducerRedshift is a command that creates redshift producer
type GatewayCreateProducerRedshift struct {
	// Redshift Creation statements
	CreationStatements *string `json:"creation-statements,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKey *string `json:"producer-encryption-key,omitempty"`
	// Redshift DB Name
	RedshiftDbName *string `json:"redshift-db-name,omitempty"`
	// Redshift Host
	RedshiftHost *string `json:"redshift-host,omitempty"`
	// Redshift Password
	RedshiftPassword *string `json:"redshift-password,omitempty"`
	// Redshift Port
	RedshiftPort *string `json:"redshift-port,omitempty"`
	// Redshift Username
	RedshiftUsername *string `json:"redshift-username,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target DB servers for connections
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
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

// NewGatewayCreateProducerRedshift instantiates a new GatewayCreateProducerRedshift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerRedshift(name string, ) *GatewayCreateProducerRedshift {
	this := GatewayCreateProducerRedshift{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var redshiftHost string = "127.0.0.1"
	this.RedshiftHost = &redshiftHost
	var redshiftPort string = "5439"
	this.RedshiftPort = &redshiftPort
	var ssl bool = false
	this.Ssl = &ssl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerRedshiftWithDefaults instantiates a new GatewayCreateProducerRedshift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerRedshiftWithDefaults() *GatewayCreateProducerRedshift {
	this := GatewayCreateProducerRedshift{}
	var json bool = false
	this.Json = &json
	var redshiftHost string = "127.0.0.1"
	this.RedshiftHost = &redshiftHost
	var redshiftPort string = "5439"
	this.RedshiftPort = &redshiftPort
	var ssl bool = false
	this.Ssl = &ssl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetCreationStatements returns the CreationStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetCreationStatements() string {
	if o == nil || o.CreationStatements == nil {
		var ret string
		return ret
	}
	return *o.CreationStatements
}

// GetCreationStatementsOk returns a tuple with the CreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetCreationStatementsOk() (*string, bool) {
	if o == nil || o.CreationStatements == nil {
		return nil, false
	}
	return o.CreationStatements, true
}

// HasCreationStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasCreationStatements() bool {
	if o != nil && o.CreationStatements != nil {
		return true
	}

	return false
}

// SetCreationStatements gets a reference to the given string and assigns it to the CreationStatements field.
func (o *GatewayCreateProducerRedshift) SetCreationStatements(v string) {
	o.CreationStatements = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerRedshift) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerRedshift) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerRedshift) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerRedshift) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKey returns the ProducerEncryptionKey field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetProducerEncryptionKey() string {
	if o == nil || o.ProducerEncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKey
}

// GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetProducerEncryptionKeyOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKey == nil {
		return nil, false
	}
	return o.ProducerEncryptionKey, true
}

// HasProducerEncryptionKey returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasProducerEncryptionKey() bool {
	if o != nil && o.ProducerEncryptionKey != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKey gets a reference to the given string and assigns it to the ProducerEncryptionKey field.
func (o *GatewayCreateProducerRedshift) SetProducerEncryptionKey(v string) {
	o.ProducerEncryptionKey = &v
}

// GetRedshiftDbName returns the RedshiftDbName field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetRedshiftDbName() string {
	if o == nil || o.RedshiftDbName == nil {
		var ret string
		return ret
	}
	return *o.RedshiftDbName
}

// GetRedshiftDbNameOk returns a tuple with the RedshiftDbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetRedshiftDbNameOk() (*string, bool) {
	if o == nil || o.RedshiftDbName == nil {
		return nil, false
	}
	return o.RedshiftDbName, true
}

// HasRedshiftDbName returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasRedshiftDbName() bool {
	if o != nil && o.RedshiftDbName != nil {
		return true
	}

	return false
}

// SetRedshiftDbName gets a reference to the given string and assigns it to the RedshiftDbName field.
func (o *GatewayCreateProducerRedshift) SetRedshiftDbName(v string) {
	o.RedshiftDbName = &v
}

// GetRedshiftHost returns the RedshiftHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetRedshiftHost() string {
	if o == nil || o.RedshiftHost == nil {
		var ret string
		return ret
	}
	return *o.RedshiftHost
}

// GetRedshiftHostOk returns a tuple with the RedshiftHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetRedshiftHostOk() (*string, bool) {
	if o == nil || o.RedshiftHost == nil {
		return nil, false
	}
	return o.RedshiftHost, true
}

// HasRedshiftHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasRedshiftHost() bool {
	if o != nil && o.RedshiftHost != nil {
		return true
	}

	return false
}

// SetRedshiftHost gets a reference to the given string and assigns it to the RedshiftHost field.
func (o *GatewayCreateProducerRedshift) SetRedshiftHost(v string) {
	o.RedshiftHost = &v
}

// GetRedshiftPassword returns the RedshiftPassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetRedshiftPassword() string {
	if o == nil || o.RedshiftPassword == nil {
		var ret string
		return ret
	}
	return *o.RedshiftPassword
}

// GetRedshiftPasswordOk returns a tuple with the RedshiftPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetRedshiftPasswordOk() (*string, bool) {
	if o == nil || o.RedshiftPassword == nil {
		return nil, false
	}
	return o.RedshiftPassword, true
}

// HasRedshiftPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasRedshiftPassword() bool {
	if o != nil && o.RedshiftPassword != nil {
		return true
	}

	return false
}

// SetRedshiftPassword gets a reference to the given string and assigns it to the RedshiftPassword field.
func (o *GatewayCreateProducerRedshift) SetRedshiftPassword(v string) {
	o.RedshiftPassword = &v
}

// GetRedshiftPort returns the RedshiftPort field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetRedshiftPort() string {
	if o == nil || o.RedshiftPort == nil {
		var ret string
		return ret
	}
	return *o.RedshiftPort
}

// GetRedshiftPortOk returns a tuple with the RedshiftPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetRedshiftPortOk() (*string, bool) {
	if o == nil || o.RedshiftPort == nil {
		return nil, false
	}
	return o.RedshiftPort, true
}

// HasRedshiftPort returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasRedshiftPort() bool {
	if o != nil && o.RedshiftPort != nil {
		return true
	}

	return false
}

// SetRedshiftPort gets a reference to the given string and assigns it to the RedshiftPort field.
func (o *GatewayCreateProducerRedshift) SetRedshiftPort(v string) {
	o.RedshiftPort = &v
}

// GetRedshiftUsername returns the RedshiftUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetRedshiftUsername() string {
	if o == nil || o.RedshiftUsername == nil {
		var ret string
		return ret
	}
	return *o.RedshiftUsername
}

// GetRedshiftUsernameOk returns a tuple with the RedshiftUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetRedshiftUsernameOk() (*string, bool) {
	if o == nil || o.RedshiftUsername == nil {
		return nil, false
	}
	return o.RedshiftUsername, true
}

// HasRedshiftUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasRedshiftUsername() bool {
	if o != nil && o.RedshiftUsername != nil {
		return true
	}

	return false
}

// SetRedshiftUsername gets a reference to the given string and assigns it to the RedshiftUsername field.
func (o *GatewayCreateProducerRedshift) SetRedshiftUsername(v string) {
	o.RedshiftUsername = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayCreateProducerRedshift) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *GatewayCreateProducerRedshift) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetSsl() bool {
	if o == nil || o.Ssl == nil {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetSslOk() (*bool, bool) {
	if o == nil || o.Ssl == nil {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasSsl() bool {
	if o != nil && o.Ssl != nil {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *GatewayCreateProducerRedshift) SetSsl(v bool) {
	o.Ssl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerRedshift) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerRedshift) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerRedshift) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerRedshift) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerRedshift) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRedshift) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerRedshift) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerRedshift) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerRedshift) MarshalJSON() ([]byte, error) {
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
	if o.ProducerEncryptionKey != nil {
		toSerialize["producer-encryption-key"] = o.ProducerEncryptionKey
	}
	if o.RedshiftDbName != nil {
		toSerialize["redshift-db-name"] = o.RedshiftDbName
	}
	if o.RedshiftHost != nil {
		toSerialize["redshift-host"] = o.RedshiftHost
	}
	if o.RedshiftPassword != nil {
		toSerialize["redshift-password"] = o.RedshiftPassword
	}
	if o.RedshiftPort != nil {
		toSerialize["redshift-port"] = o.RedshiftPort
	}
	if o.RedshiftUsername != nil {
		toSerialize["redshift-username"] = o.RedshiftUsername
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessHost != nil {
		toSerialize["secure-access-host"] = o.SecureAccessHost
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

type NullableGatewayCreateProducerRedshift struct {
	value *GatewayCreateProducerRedshift
	isSet bool
}

func (v NullableGatewayCreateProducerRedshift) Get() *GatewayCreateProducerRedshift {
	return v.value
}

func (v *NullableGatewayCreateProducerRedshift) Set(val *GatewayCreateProducerRedshift) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerRedshift) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerRedshift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerRedshift(val *GatewayCreateProducerRedshift) *NullableGatewayCreateProducerRedshift {
	return &NullableGatewayCreateProducerRedshift{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerRedshift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerRedshift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


