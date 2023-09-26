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

// DsUpdateHanaDb dsUpdateHanaDb is a command that updates hanadb dynamic secret
type DsUpdateHanaDb struct {
	// Protection from accidental deletion of this item [true/false]
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
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret new name
	NewName *string `json:"new-name,omitempty"`
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

// NewDsUpdateHanaDb instantiates a new DsUpdateHanaDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDsUpdateHanaDb(name string, ) *DsUpdateHanaDb {
	this := DsUpdateHanaDb{}
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

// NewDsUpdateHanaDbWithDefaults instantiates a new DsUpdateHanaDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDsUpdateHanaDbWithDefaults() *DsUpdateHanaDb {
	this := DsUpdateHanaDb{}
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
func (o *DsUpdateHanaDb) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DsUpdateHanaDb) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetHanaDbname returns the HanaDbname field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetHanaDbname() string {
	if o == nil || o.HanaDbname == nil {
		var ret string
		return ret
	}
	return *o.HanaDbname
}

// GetHanaDbnameOk returns a tuple with the HanaDbname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetHanaDbnameOk() (*string, bool) {
	if o == nil || o.HanaDbname == nil {
		return nil, false
	}
	return o.HanaDbname, true
}

// HasHanaDbname returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasHanaDbname() bool {
	if o != nil && o.HanaDbname != nil {
		return true
	}

	return false
}

// SetHanaDbname gets a reference to the given string and assigns it to the HanaDbname field.
func (o *DsUpdateHanaDb) SetHanaDbname(v string) {
	o.HanaDbname = &v
}

// GetHanadbCreateStatements returns the HanadbCreateStatements field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetHanadbCreateStatements() string {
	if o == nil || o.HanadbCreateStatements == nil {
		var ret string
		return ret
	}
	return *o.HanadbCreateStatements
}

// GetHanadbCreateStatementsOk returns a tuple with the HanadbCreateStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetHanadbCreateStatementsOk() (*string, bool) {
	if o == nil || o.HanadbCreateStatements == nil {
		return nil, false
	}
	return o.HanadbCreateStatements, true
}

// HasHanadbCreateStatements returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasHanadbCreateStatements() bool {
	if o != nil && o.HanadbCreateStatements != nil {
		return true
	}

	return false
}

// SetHanadbCreateStatements gets a reference to the given string and assigns it to the HanadbCreateStatements field.
func (o *DsUpdateHanaDb) SetHanadbCreateStatements(v string) {
	o.HanadbCreateStatements = &v
}

// GetHanadbHost returns the HanadbHost field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetHanadbHost() string {
	if o == nil || o.HanadbHost == nil {
		var ret string
		return ret
	}
	return *o.HanadbHost
}

// GetHanadbHostOk returns a tuple with the HanadbHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetHanadbHostOk() (*string, bool) {
	if o == nil || o.HanadbHost == nil {
		return nil, false
	}
	return o.HanadbHost, true
}

// HasHanadbHost returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasHanadbHost() bool {
	if o != nil && o.HanadbHost != nil {
		return true
	}

	return false
}

// SetHanadbHost gets a reference to the given string and assigns it to the HanadbHost field.
func (o *DsUpdateHanaDb) SetHanadbHost(v string) {
	o.HanadbHost = &v
}

// GetHanadbPassword returns the HanadbPassword field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetHanadbPassword() string {
	if o == nil || o.HanadbPassword == nil {
		var ret string
		return ret
	}
	return *o.HanadbPassword
}

// GetHanadbPasswordOk returns a tuple with the HanadbPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetHanadbPasswordOk() (*string, bool) {
	if o == nil || o.HanadbPassword == nil {
		return nil, false
	}
	return o.HanadbPassword, true
}

// HasHanadbPassword returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasHanadbPassword() bool {
	if o != nil && o.HanadbPassword != nil {
		return true
	}

	return false
}

// SetHanadbPassword gets a reference to the given string and assigns it to the HanadbPassword field.
func (o *DsUpdateHanaDb) SetHanadbPassword(v string) {
	o.HanadbPassword = &v
}

// GetHanadbPort returns the HanadbPort field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetHanadbPort() string {
	if o == nil || o.HanadbPort == nil {
		var ret string
		return ret
	}
	return *o.HanadbPort
}

// GetHanadbPortOk returns a tuple with the HanadbPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetHanadbPortOk() (*string, bool) {
	if o == nil || o.HanadbPort == nil {
		return nil, false
	}
	return o.HanadbPort, true
}

// HasHanadbPort returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasHanadbPort() bool {
	if o != nil && o.HanadbPort != nil {
		return true
	}

	return false
}

// SetHanadbPort gets a reference to the given string and assigns it to the HanadbPort field.
func (o *DsUpdateHanaDb) SetHanadbPort(v string) {
	o.HanadbPort = &v
}

// GetHanadbRevocationStatements returns the HanadbRevocationStatements field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetHanadbRevocationStatements() string {
	if o == nil || o.HanadbRevocationStatements == nil {
		var ret string
		return ret
	}
	return *o.HanadbRevocationStatements
}

// GetHanadbRevocationStatementsOk returns a tuple with the HanadbRevocationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetHanadbRevocationStatementsOk() (*string, bool) {
	if o == nil || o.HanadbRevocationStatements == nil {
		return nil, false
	}
	return o.HanadbRevocationStatements, true
}

// HasHanadbRevocationStatements returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasHanadbRevocationStatements() bool {
	if o != nil && o.HanadbRevocationStatements != nil {
		return true
	}

	return false
}

// SetHanadbRevocationStatements gets a reference to the given string and assigns it to the HanadbRevocationStatements field.
func (o *DsUpdateHanaDb) SetHanadbRevocationStatements(v string) {
	o.HanadbRevocationStatements = &v
}

// GetHanadbUsername returns the HanadbUsername field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetHanadbUsername() string {
	if o == nil || o.HanadbUsername == nil {
		var ret string
		return ret
	}
	return *o.HanadbUsername
}

// GetHanadbUsernameOk returns a tuple with the HanadbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetHanadbUsernameOk() (*string, bool) {
	if o == nil || o.HanadbUsername == nil {
		return nil, false
	}
	return o.HanadbUsername, true
}

// HasHanadbUsername returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasHanadbUsername() bool {
	if o != nil && o.HanadbUsername != nil {
		return true
	}

	return false
}

// SetHanadbUsername gets a reference to the given string and assigns it to the HanadbUsername field.
func (o *DsUpdateHanaDb) SetHanadbUsername(v string) {
	o.HanadbUsername = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DsUpdateHanaDb) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DsUpdateHanaDb) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DsUpdateHanaDb) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DsUpdateHanaDb) SetNewName(v string) {
	o.NewName = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DsUpdateHanaDb) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *DsUpdateHanaDb) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessDbSchema returns the SecureAccessDbSchema field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetSecureAccessDbSchema() string {
	if o == nil || o.SecureAccessDbSchema == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessDbSchema
}

// GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetSecureAccessDbSchemaOk() (*string, bool) {
	if o == nil || o.SecureAccessDbSchema == nil {
		return nil, false
	}
	return o.SecureAccessDbSchema, true
}

// HasSecureAccessDbSchema returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasSecureAccessDbSchema() bool {
	if o != nil && o.SecureAccessDbSchema != nil {
		return true
	}

	return false
}

// SetSecureAccessDbSchema gets a reference to the given string and assigns it to the SecureAccessDbSchema field.
func (o *DsUpdateHanaDb) SetSecureAccessDbSchema(v string) {
	o.SecureAccessDbSchema = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DsUpdateHanaDb) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *DsUpdateHanaDb) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *DsUpdateHanaDb) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DsUpdateHanaDb) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DsUpdateHanaDb) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DsUpdateHanaDb) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DsUpdateHanaDb) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DsUpdateHanaDb) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsUpdateHanaDb) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DsUpdateHanaDb) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DsUpdateHanaDb) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DsUpdateHanaDb) MarshalJSON() ([]byte, error) {
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
	if o.Json != nil {
		toSerialize["json"] = o.Json
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

type NullableDsUpdateHanaDb struct {
	value *DsUpdateHanaDb
	isSet bool
}

func (v NullableDsUpdateHanaDb) Get() *DsUpdateHanaDb {
	return v.value
}

func (v *NullableDsUpdateHanaDb) Set(val *DsUpdateHanaDb) {
	v.value = val
	v.isSet = true
}

func (v NullableDsUpdateHanaDb) IsSet() bool {
	return v.isSet
}

func (v *NullableDsUpdateHanaDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDsUpdateHanaDb(val *DsUpdateHanaDb) *NullableDsUpdateHanaDb {
	return &NullableDsUpdateHanaDb{value: val, isSet: true}
}

func (v NullableDsUpdateHanaDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDsUpdateHanaDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


