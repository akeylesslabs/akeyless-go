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

// DynamicSecretCreateHanaDb dynamicSecretCreateHanaDb is a command that creates hanadb dynamic secret
type DynamicSecretCreateHanaDb struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
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
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
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

// NewDynamicSecretCreateHanaDb instantiates a new DynamicSecretCreateHanaDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretCreateHanaDb(name string, ) *DynamicSecretCreateHanaDb {
	this := DynamicSecretCreateHanaDb{}
	var hanadbHost string = "127.0.0.1"
	this.HanadbHost = &hanadbHost
	var hanadbPort string = "443"
	this.HanadbPort = &hanadbPort
	var json bool = false
	this.Json = &json
	this.Name = name
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretCreateHanaDbWithDefaults instantiates a new DynamicSecretCreateHanaDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretCreateHanaDbWithDefaults() *DynamicSecretCreateHanaDb {
	this := DynamicSecretCreateHanaDb{}
	var hanadbHost string = "127.0.0.1"
	this.HanadbHost = &hanadbHost
	var hanadbPort string = "443"
	this.HanadbPort = &hanadbPort
	var json bool = false
	this.Json = &json
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretCreateHanaDb) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretCreateHanaDb) SetDescription(v string) {
	o.Description = &v
}

// GetHanaDbname returns the HanaDbname field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetHanaDbname() string {
	if o == nil || o.HanaDbname == nil {
		var ret string
		return ret
	}
	return *o.HanaDbname
}

// GetHanaDbnameOk returns a tuple with the HanaDbname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetHanaDbnameOk() (*string, bool) {
	if o == nil || o.HanaDbname == nil {
		return nil, false
	}
	return o.HanaDbname, true
}

// HasHanaDbname returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasHanaDbname() bool {
	if o != nil && o.HanaDbname != nil {
		return true
	}

	return false
}

// SetHanaDbname gets a reference to the given string and assigns it to the HanaDbname field.
func (o *DynamicSecretCreateHanaDb) SetHanaDbname(v string) {
	o.HanaDbname = &v
}

// GetHanadbCreateStatements returns the HanadbCreateStatements field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetHanadbCreateStatements() string {
	if o == nil || o.HanadbCreateStatements == nil {
		var ret string
		return ret
	}
	return *o.HanadbCreateStatements
}

// GetHanadbCreateStatementsOk returns a tuple with the HanadbCreateStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetHanadbCreateStatementsOk() (*string, bool) {
	if o == nil || o.HanadbCreateStatements == nil {
		return nil, false
	}
	return o.HanadbCreateStatements, true
}

// HasHanadbCreateStatements returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasHanadbCreateStatements() bool {
	if o != nil && o.HanadbCreateStatements != nil {
		return true
	}

	return false
}

// SetHanadbCreateStatements gets a reference to the given string and assigns it to the HanadbCreateStatements field.
func (o *DynamicSecretCreateHanaDb) SetHanadbCreateStatements(v string) {
	o.HanadbCreateStatements = &v
}

// GetHanadbHost returns the HanadbHost field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetHanadbHost() string {
	if o == nil || o.HanadbHost == nil {
		var ret string
		return ret
	}
	return *o.HanadbHost
}

// GetHanadbHostOk returns a tuple with the HanadbHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetHanadbHostOk() (*string, bool) {
	if o == nil || o.HanadbHost == nil {
		return nil, false
	}
	return o.HanadbHost, true
}

// HasHanadbHost returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasHanadbHost() bool {
	if o != nil && o.HanadbHost != nil {
		return true
	}

	return false
}

// SetHanadbHost gets a reference to the given string and assigns it to the HanadbHost field.
func (o *DynamicSecretCreateHanaDb) SetHanadbHost(v string) {
	o.HanadbHost = &v
}

// GetHanadbPassword returns the HanadbPassword field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetHanadbPassword() string {
	if o == nil || o.HanadbPassword == nil {
		var ret string
		return ret
	}
	return *o.HanadbPassword
}

// GetHanadbPasswordOk returns a tuple with the HanadbPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetHanadbPasswordOk() (*string, bool) {
	if o == nil || o.HanadbPassword == nil {
		return nil, false
	}
	return o.HanadbPassword, true
}

// HasHanadbPassword returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasHanadbPassword() bool {
	if o != nil && o.HanadbPassword != nil {
		return true
	}

	return false
}

// SetHanadbPassword gets a reference to the given string and assigns it to the HanadbPassword field.
func (o *DynamicSecretCreateHanaDb) SetHanadbPassword(v string) {
	o.HanadbPassword = &v
}

// GetHanadbPort returns the HanadbPort field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetHanadbPort() string {
	if o == nil || o.HanadbPort == nil {
		var ret string
		return ret
	}
	return *o.HanadbPort
}

// GetHanadbPortOk returns a tuple with the HanadbPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetHanadbPortOk() (*string, bool) {
	if o == nil || o.HanadbPort == nil {
		return nil, false
	}
	return o.HanadbPort, true
}

// HasHanadbPort returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasHanadbPort() bool {
	if o != nil && o.HanadbPort != nil {
		return true
	}

	return false
}

// SetHanadbPort gets a reference to the given string and assigns it to the HanadbPort field.
func (o *DynamicSecretCreateHanaDb) SetHanadbPort(v string) {
	o.HanadbPort = &v
}

// GetHanadbRevocationStatements returns the HanadbRevocationStatements field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetHanadbRevocationStatements() string {
	if o == nil || o.HanadbRevocationStatements == nil {
		var ret string
		return ret
	}
	return *o.HanadbRevocationStatements
}

// GetHanadbRevocationStatementsOk returns a tuple with the HanadbRevocationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetHanadbRevocationStatementsOk() (*string, bool) {
	if o == nil || o.HanadbRevocationStatements == nil {
		return nil, false
	}
	return o.HanadbRevocationStatements, true
}

// HasHanadbRevocationStatements returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasHanadbRevocationStatements() bool {
	if o != nil && o.HanadbRevocationStatements != nil {
		return true
	}

	return false
}

// SetHanadbRevocationStatements gets a reference to the given string and assigns it to the HanadbRevocationStatements field.
func (o *DynamicSecretCreateHanaDb) SetHanadbRevocationStatements(v string) {
	o.HanadbRevocationStatements = &v
}

// GetHanadbUsername returns the HanadbUsername field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetHanadbUsername() string {
	if o == nil || o.HanadbUsername == nil {
		var ret string
		return ret
	}
	return *o.HanadbUsername
}

// GetHanadbUsernameOk returns a tuple with the HanadbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetHanadbUsernameOk() (*string, bool) {
	if o == nil || o.HanadbUsername == nil {
		return nil, false
	}
	return o.HanadbUsername, true
}

// HasHanadbUsername returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasHanadbUsername() bool {
	if o != nil && o.HanadbUsername != nil {
		return true
	}

	return false
}

// SetHanadbUsername gets a reference to the given string and assigns it to the HanadbUsername field.
func (o *DynamicSecretCreateHanaDb) SetHanadbUsername(v string) {
	o.HanadbUsername = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretCreateHanaDb) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretCreateHanaDb) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretCreateHanaDb) SetName(v string) {
	o.Name = v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *DynamicSecretCreateHanaDb) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretCreateHanaDb) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *DynamicSecretCreateHanaDb) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessDbSchema returns the SecureAccessDbSchema field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessDbSchema() string {
	if o == nil || o.SecureAccessDbSchema == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDbSchema
}

// GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessDbSchemaOk() (*string, bool) {
	if o == nil || o.SecureAccessDbSchema == nil {
		return nil, false
	}
	return o.SecureAccessDbSchema, true
}

// HasSecureAccessDbSchema returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasSecureAccessDbSchema() bool {
	if o != nil && o.SecureAccessDbSchema != nil {
		return true
	}

	return false
}

// SetSecureAccessDbSchema gets a reference to the given string and assigns it to the SecureAccessDbSchema field.
func (o *DynamicSecretCreateHanaDb) SetSecureAccessDbSchema(v string) {
	o.SecureAccessDbSchema = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DynamicSecretCreateHanaDb) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *DynamicSecretCreateHanaDb) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *DynamicSecretCreateHanaDb) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretCreateHanaDb) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretCreateHanaDb) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretCreateHanaDb) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretCreateHanaDb) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretCreateHanaDb) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateHanaDb) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretCreateHanaDb) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretCreateHanaDb) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretCreateHanaDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
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

type NullableDynamicSecretCreateHanaDb struct {
	value *DynamicSecretCreateHanaDb
	isSet bool
}

func (v NullableDynamicSecretCreateHanaDb) Get() *DynamicSecretCreateHanaDb {
	return v.value
}

func (v *NullableDynamicSecretCreateHanaDb) Set(val *DynamicSecretCreateHanaDb) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretCreateHanaDb) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretCreateHanaDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretCreateHanaDb(val *DynamicSecretCreateHanaDb) *NullableDynamicSecretCreateHanaDb {
	return &NullableDynamicSecretCreateHanaDb{value: val, isSet: true}
}

func (v NullableDynamicSecretCreateHanaDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretCreateHanaDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

