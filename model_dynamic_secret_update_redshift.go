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

// DynamicSecretUpdateRedshift dynamicSecretUpdateRedshift is a command that updates redshift dynamic secret
type DynamicSecretUpdateRedshift struct {
	// Redshift Creation statements
	CreationStatements *string `json:"creation-statements,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic secret name
	NewName *string `json:"new-name,omitempty"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
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
	// Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts)
	SecureAccessHost *[]string `json:"secure-access-host,omitempty"`
	// Enable/Disable SSL [true/false]
	Ssl *bool `json:"ssl,omitempty"`
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

// NewDynamicSecretUpdateRedshift instantiates a new DynamicSecretUpdateRedshift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdateRedshift(name string, ) *DynamicSecretUpdateRedshift {
	this := DynamicSecretUpdateRedshift{}
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

// NewDynamicSecretUpdateRedshiftWithDefaults instantiates a new DynamicSecretUpdateRedshift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdateRedshiftWithDefaults() *DynamicSecretUpdateRedshift {
	this := DynamicSecretUpdateRedshift{}
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
func (o *DynamicSecretUpdateRedshift) GetCreationStatements() string {
	if o == nil || o.CreationStatements == nil {
		var ret string
		return ret
	}
	return *o.CreationStatements
}

// GetCreationStatementsOk returns a tuple with the CreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetCreationStatementsOk() (*string, bool) {
	if o == nil || o.CreationStatements == nil {
		return nil, false
	}
	return o.CreationStatements, true
}

// HasCreationStatements returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasCreationStatements() bool {
	if o != nil && o.CreationStatements != nil {
		return true
	}

	return false
}

// SetCreationStatements gets a reference to the given string and assigns it to the CreationStatements field.
func (o *DynamicSecretUpdateRedshift) SetCreationStatements(v string) {
	o.CreationStatements = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretUpdateRedshift) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretUpdateRedshift) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretUpdateRedshift) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretUpdateRedshift) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretUpdateRedshift) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DynamicSecretUpdateRedshift) SetNewName(v string) {
	o.NewName = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *DynamicSecretUpdateRedshift) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetProducerEncryptionKey returns the ProducerEncryptionKey field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetProducerEncryptionKey() string {
	if o == nil || o.ProducerEncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKey
}

// GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetProducerEncryptionKeyOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKey == nil {
		return nil, false
	}
	return o.ProducerEncryptionKey, true
}

// HasProducerEncryptionKey returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasProducerEncryptionKey() bool {
	if o != nil && o.ProducerEncryptionKey != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKey gets a reference to the given string and assigns it to the ProducerEncryptionKey field.
func (o *DynamicSecretUpdateRedshift) SetProducerEncryptionKey(v string) {
	o.ProducerEncryptionKey = &v
}

// GetRedshiftDbName returns the RedshiftDbName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetRedshiftDbName() string {
	if o == nil || o.RedshiftDbName == nil {
		var ret string
		return ret
	}
	return *o.RedshiftDbName
}

// GetRedshiftDbNameOk returns a tuple with the RedshiftDbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetRedshiftDbNameOk() (*string, bool) {
	if o == nil || o.RedshiftDbName == nil {
		return nil, false
	}
	return o.RedshiftDbName, true
}

// HasRedshiftDbName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasRedshiftDbName() bool {
	if o != nil && o.RedshiftDbName != nil {
		return true
	}

	return false
}

// SetRedshiftDbName gets a reference to the given string and assigns it to the RedshiftDbName field.
func (o *DynamicSecretUpdateRedshift) SetRedshiftDbName(v string) {
	o.RedshiftDbName = &v
}

// GetRedshiftHost returns the RedshiftHost field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetRedshiftHost() string {
	if o == nil || o.RedshiftHost == nil {
		var ret string
		return ret
	}
	return *o.RedshiftHost
}

// GetRedshiftHostOk returns a tuple with the RedshiftHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetRedshiftHostOk() (*string, bool) {
	if o == nil || o.RedshiftHost == nil {
		return nil, false
	}
	return o.RedshiftHost, true
}

// HasRedshiftHost returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasRedshiftHost() bool {
	if o != nil && o.RedshiftHost != nil {
		return true
	}

	return false
}

// SetRedshiftHost gets a reference to the given string and assigns it to the RedshiftHost field.
func (o *DynamicSecretUpdateRedshift) SetRedshiftHost(v string) {
	o.RedshiftHost = &v
}

// GetRedshiftPassword returns the RedshiftPassword field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetRedshiftPassword() string {
	if o == nil || o.RedshiftPassword == nil {
		var ret string
		return ret
	}
	return *o.RedshiftPassword
}

// GetRedshiftPasswordOk returns a tuple with the RedshiftPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetRedshiftPasswordOk() (*string, bool) {
	if o == nil || o.RedshiftPassword == nil {
		return nil, false
	}
	return o.RedshiftPassword, true
}

// HasRedshiftPassword returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasRedshiftPassword() bool {
	if o != nil && o.RedshiftPassword != nil {
		return true
	}

	return false
}

// SetRedshiftPassword gets a reference to the given string and assigns it to the RedshiftPassword field.
func (o *DynamicSecretUpdateRedshift) SetRedshiftPassword(v string) {
	o.RedshiftPassword = &v
}

// GetRedshiftPort returns the RedshiftPort field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetRedshiftPort() string {
	if o == nil || o.RedshiftPort == nil {
		var ret string
		return ret
	}
	return *o.RedshiftPort
}

// GetRedshiftPortOk returns a tuple with the RedshiftPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetRedshiftPortOk() (*string, bool) {
	if o == nil || o.RedshiftPort == nil {
		return nil, false
	}
	return o.RedshiftPort, true
}

// HasRedshiftPort returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasRedshiftPort() bool {
	if o != nil && o.RedshiftPort != nil {
		return true
	}

	return false
}

// SetRedshiftPort gets a reference to the given string and assigns it to the RedshiftPort field.
func (o *DynamicSecretUpdateRedshift) SetRedshiftPort(v string) {
	o.RedshiftPort = &v
}

// GetRedshiftUsername returns the RedshiftUsername field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetRedshiftUsername() string {
	if o == nil || o.RedshiftUsername == nil {
		var ret string
		return ret
	}
	return *o.RedshiftUsername
}

// GetRedshiftUsernameOk returns a tuple with the RedshiftUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetRedshiftUsernameOk() (*string, bool) {
	if o == nil || o.RedshiftUsername == nil {
		return nil, false
	}
	return o.RedshiftUsername, true
}

// HasRedshiftUsername returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasRedshiftUsername() bool {
	if o != nil && o.RedshiftUsername != nil {
		return true
	}

	return false
}

// SetRedshiftUsername gets a reference to the given string and assigns it to the RedshiftUsername field.
func (o *DynamicSecretUpdateRedshift) SetRedshiftUsername(v string) {
	o.RedshiftUsername = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *DynamicSecretUpdateRedshift) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return *o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetSecureAccessHostOk() (*[]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *DynamicSecretUpdateRedshift) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetSsl() bool {
	if o == nil || o.Ssl == nil {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetSslOk() (*bool, bool) {
	if o == nil || o.Ssl == nil {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasSsl() bool {
	if o != nil && o.Ssl != nil {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *DynamicSecretUpdateRedshift) SetSsl(v bool) {
	o.Ssl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretUpdateRedshift) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretUpdateRedshift) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretUpdateRedshift) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretUpdateRedshift) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateRedshift) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateRedshift) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateRedshift) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretUpdateRedshift) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DynamicSecretUpdateRedshift) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationStatements != nil {
		toSerialize["creation-statements"] = o.CreationStatements
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
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.PasswordLength != nil {
		toSerialize["password-length"] = o.PasswordLength
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

type NullableDynamicSecretUpdateRedshift struct {
	value *DynamicSecretUpdateRedshift
	isSet bool
}

func (v NullableDynamicSecretUpdateRedshift) Get() *DynamicSecretUpdateRedshift {
	return v.value
}

func (v *NullableDynamicSecretUpdateRedshift) Set(val *DynamicSecretUpdateRedshift) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdateRedshift) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdateRedshift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdateRedshift(val *DynamicSecretUpdateRedshift) *NullableDynamicSecretUpdateRedshift {
	return &NullableDynamicSecretUpdateRedshift{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdateRedshift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdateRedshift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


