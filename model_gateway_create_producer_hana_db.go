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

// GatewayCreateProducerHanaDb gatewayCreateProducerHanaDb is a command that creates hanadb producer
type GatewayCreateProducerHanaDb struct {
	// Protection from accidental deletion of this item
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// HanaDb Name
	HanaDbname *string `json:"hana-dbname,omitempty"`
	// HanaDb Creation statements
	HanadbCreateStatements *string `json:"hanadb-create-statements,omitempty"`
	// HanaDb Host
	HanadbHost *string `json:"hanadb-host,omitempty"`
	// HanaDb Password
	HanadbPassword *string `json:"hanadb-password,omitempty"`
	// HanaDb Port
	HanadbPort *string `json:"hanadb-port,omitempty"`
	// HanaDb Revocation statements
	HanadbRevocationStatements *string `json:"hanadb-revocation-statements,omitempty"`
	// HanaDb Username
	HanadbUsername *string `json:"hanadb-username,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	SecureAccessDbSchema *string `json:"secure-access-db-schema,omitempty"`
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
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

// NewGatewayCreateProducerHanaDb instantiates a new GatewayCreateProducerHanaDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerHanaDb(name string, ) *GatewayCreateProducerHanaDb {
	this := GatewayCreateProducerHanaDb{}
	var hanadbHost string = "127.0.0.1"
	this.HanadbHost = &hanadbHost
	var hanadbPort string = "443"
	this.HanadbPort = &hanadbPort
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerHanaDbWithDefaults instantiates a new GatewayCreateProducerHanaDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerHanaDbWithDefaults() *GatewayCreateProducerHanaDb {
	this := GatewayCreateProducerHanaDb{}
	var hanadbHost string = "127.0.0.1"
	this.HanadbHost = &hanadbHost
	var hanadbPort string = "443"
	this.HanadbPort = &hanadbPort
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerHanaDb) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetHanaDbname returns the HanaDbname field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetHanaDbname() string {
	if o == nil || o.HanaDbname == nil {
		var ret string
		return ret
	}
	return *o.HanaDbname
}

// GetHanaDbnameOk returns a tuple with the HanaDbname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetHanaDbnameOk() (*string, bool) {
	if o == nil || o.HanaDbname == nil {
		return nil, false
	}
	return o.HanaDbname, true
}

// HasHanaDbname returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasHanaDbname() bool {
	if o != nil && o.HanaDbname != nil {
		return true
	}

	return false
}

// SetHanaDbname gets a reference to the given string and assigns it to the HanaDbname field.
func (o *GatewayCreateProducerHanaDb) SetHanaDbname(v string) {
	o.HanaDbname = &v
}

// GetHanadbCreateStatements returns the HanadbCreateStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetHanadbCreateStatements() string {
	if o == nil || o.HanadbCreateStatements == nil {
		var ret string
		return ret
	}
	return *o.HanadbCreateStatements
}

// GetHanadbCreateStatementsOk returns a tuple with the HanadbCreateStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetHanadbCreateStatementsOk() (*string, bool) {
	if o == nil || o.HanadbCreateStatements == nil {
		return nil, false
	}
	return o.HanadbCreateStatements, true
}

// HasHanadbCreateStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasHanadbCreateStatements() bool {
	if o != nil && o.HanadbCreateStatements != nil {
		return true
	}

	return false
}

// SetHanadbCreateStatements gets a reference to the given string and assigns it to the HanadbCreateStatements field.
func (o *GatewayCreateProducerHanaDb) SetHanadbCreateStatements(v string) {
	o.HanadbCreateStatements = &v
}

// GetHanadbHost returns the HanadbHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetHanadbHost() string {
	if o == nil || o.HanadbHost == nil {
		var ret string
		return ret
	}
	return *o.HanadbHost
}

// GetHanadbHostOk returns a tuple with the HanadbHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetHanadbHostOk() (*string, bool) {
	if o == nil || o.HanadbHost == nil {
		return nil, false
	}
	return o.HanadbHost, true
}

// HasHanadbHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasHanadbHost() bool {
	if o != nil && o.HanadbHost != nil {
		return true
	}

	return false
}

// SetHanadbHost gets a reference to the given string and assigns it to the HanadbHost field.
func (o *GatewayCreateProducerHanaDb) SetHanadbHost(v string) {
	o.HanadbHost = &v
}

// GetHanadbPassword returns the HanadbPassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetHanadbPassword() string {
	if o == nil || o.HanadbPassword == nil {
		var ret string
		return ret
	}
	return *o.HanadbPassword
}

// GetHanadbPasswordOk returns a tuple with the HanadbPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetHanadbPasswordOk() (*string, bool) {
	if o == nil || o.HanadbPassword == nil {
		return nil, false
	}
	return o.HanadbPassword, true
}

// HasHanadbPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasHanadbPassword() bool {
	if o != nil && o.HanadbPassword != nil {
		return true
	}

	return false
}

// SetHanadbPassword gets a reference to the given string and assigns it to the HanadbPassword field.
func (o *GatewayCreateProducerHanaDb) SetHanadbPassword(v string) {
	o.HanadbPassword = &v
}

// GetHanadbPort returns the HanadbPort field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetHanadbPort() string {
	if o == nil || o.HanadbPort == nil {
		var ret string
		return ret
	}
	return *o.HanadbPort
}

// GetHanadbPortOk returns a tuple with the HanadbPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetHanadbPortOk() (*string, bool) {
	if o == nil || o.HanadbPort == nil {
		return nil, false
	}
	return o.HanadbPort, true
}

// HasHanadbPort returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasHanadbPort() bool {
	if o != nil && o.HanadbPort != nil {
		return true
	}

	return false
}

// SetHanadbPort gets a reference to the given string and assigns it to the HanadbPort field.
func (o *GatewayCreateProducerHanaDb) SetHanadbPort(v string) {
	o.HanadbPort = &v
}

// GetHanadbRevocationStatements returns the HanadbRevocationStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetHanadbRevocationStatements() string {
	if o == nil || o.HanadbRevocationStatements == nil {
		var ret string
		return ret
	}
	return *o.HanadbRevocationStatements
}

// GetHanadbRevocationStatementsOk returns a tuple with the HanadbRevocationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetHanadbRevocationStatementsOk() (*string, bool) {
	if o == nil || o.HanadbRevocationStatements == nil {
		return nil, false
	}
	return o.HanadbRevocationStatements, true
}

// HasHanadbRevocationStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasHanadbRevocationStatements() bool {
	if o != nil && o.HanadbRevocationStatements != nil {
		return true
	}

	return false
}

// SetHanadbRevocationStatements gets a reference to the given string and assigns it to the HanadbRevocationStatements field.
func (o *GatewayCreateProducerHanaDb) SetHanadbRevocationStatements(v string) {
	o.HanadbRevocationStatements = &v
}

// GetHanadbUsername returns the HanadbUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetHanadbUsername() string {
	if o == nil || o.HanadbUsername == nil {
		var ret string
		return ret
	}
	return *o.HanadbUsername
}

// GetHanadbUsernameOk returns a tuple with the HanadbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetHanadbUsernameOk() (*string, bool) {
	if o == nil || o.HanadbUsername == nil {
		return nil, false
	}
	return o.HanadbUsername, true
}

// HasHanadbUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasHanadbUsername() bool {
	if o != nil && o.HanadbUsername != nil {
		return true
	}

	return false
}

// SetHanadbUsername gets a reference to the given string and assigns it to the HanadbUsername field.
func (o *GatewayCreateProducerHanaDb) SetHanadbUsername(v string) {
	o.HanadbUsername = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerHanaDb) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerHanaDb) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerHanaDb) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayCreateProducerHanaDb) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessDbSchema returns the SecureAccessDbSchema field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessDbSchema() string {
	if o == nil || o.SecureAccessDbSchema == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDbSchema
}

// GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessDbSchemaOk() (*string, bool) {
	if o == nil || o.SecureAccessDbSchema == nil {
		return nil, false
	}
	return o.SecureAccessDbSchema, true
}

// HasSecureAccessDbSchema returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasSecureAccessDbSchema() bool {
	if o != nil && o.SecureAccessDbSchema != nil {
		return true
	}

	return false
}

// SetSecureAccessDbSchema gets a reference to the given string and assigns it to the SecureAccessDbSchema field.
func (o *GatewayCreateProducerHanaDb) SetSecureAccessDbSchema(v string) {
	o.SecureAccessDbSchema = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayCreateProducerHanaDb) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *GatewayCreateProducerHanaDb) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayCreateProducerHanaDb) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerHanaDb) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerHanaDb) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerHanaDb) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerHanaDb) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerHanaDb) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerHanaDb) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerHanaDb) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerHanaDb) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerHanaDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.HanaDbname != nil {
		toSerialize["hana-dbname"] = o.HanaDbname
	}
	if o.HanadbCreateStatements != nil {
		toSerialize["hanadb-create-statements"] = o.HanadbCreateStatements
	}
	if o.HanadbHost != nil {
		toSerialize["hanadb-host"] = o.HanadbHost
	}
	if o.HanadbPassword != nil {
		toSerialize["hanadb-password"] = o.HanadbPassword
	}
	if o.HanadbPort != nil {
		toSerialize["hanadb-port"] = o.HanadbPort
	}
	if o.HanadbRevocationStatements != nil {
		toSerialize["hanadb-revocation-statements"] = o.HanadbRevocationStatements
	}
	if o.HanadbUsername != nil {
		toSerialize["hanadb-username"] = o.HanadbUsername
	}
	if true {
		toSerialize["name"] = o.Name
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

type NullableGatewayCreateProducerHanaDb struct {
	value *GatewayCreateProducerHanaDb
	isSet bool
}

func (v NullableGatewayCreateProducerHanaDb) Get() *GatewayCreateProducerHanaDb {
	return v.value
}

func (v *NullableGatewayCreateProducerHanaDb) Set(val *GatewayCreateProducerHanaDb) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerHanaDb) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerHanaDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerHanaDb(val *GatewayCreateProducerHanaDb) *NullableGatewayCreateProducerHanaDb {
	return &NullableGatewayCreateProducerHanaDb{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerHanaDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerHanaDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


