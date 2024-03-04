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

// DynamicSecretUpdateSnowflake GatewayUpdateProducerSnowflakeCmd is a command that updates a Snowflake dynamic secret
type DynamicSecretUpdateSnowflake struct {
	// Account name
	Account *string `json:"account,omitempty"`
	// Database Password
	AccountPassword *string `json:"account-password,omitempty"`
	// Database Username
	AccountUsername *string `json:"account-username,omitempty"`
	// Database name
	DbName *string `json:"db-name,omitempty"`
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
	// RSA Private key (base64 encoded)
	PrivateKey *string `json:"private-key,omitempty"`
	// The Private key passphrase
	PrivateKeyPassphrase *string `json:"private-key-passphrase,omitempty"`
	// User role
	Role *string `json:"role,omitempty"`
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
	// Warehouse name
	Warehouse *string `json:"warehouse,omitempty"`
}

// NewDynamicSecretUpdateSnowflake instantiates a new DynamicSecretUpdateSnowflake object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretUpdateSnowflake(name string, ) *DynamicSecretUpdateSnowflake {
	this := DynamicSecretUpdateSnowflake{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "24h"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretUpdateSnowflakeWithDefaults instantiates a new DynamicSecretUpdateSnowflake object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretUpdateSnowflakeWithDefaults() *DynamicSecretUpdateSnowflake {
	this := DynamicSecretUpdateSnowflake{}
	var json bool = false
	this.Json = &json
	var userTtl string = "24h"
	this.UserTtl = &userTtl
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *DynamicSecretUpdateSnowflake) SetAccount(v string) {
	o.Account = &v
}

// GetAccountPassword returns the AccountPassword field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetAccountPassword() string {
	if o == nil || o.AccountPassword == nil {
		var ret string
		return ret
	}
	return *o.AccountPassword
}

// GetAccountPasswordOk returns a tuple with the AccountPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetAccountPasswordOk() (*string, bool) {
	if o == nil || o.AccountPassword == nil {
		return nil, false
	}
	return o.AccountPassword, true
}

// HasAccountPassword returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasAccountPassword() bool {
	if o != nil && o.AccountPassword != nil {
		return true
	}

	return false
}

// SetAccountPassword gets a reference to the given string and assigns it to the AccountPassword field.
func (o *DynamicSecretUpdateSnowflake) SetAccountPassword(v string) {
	o.AccountPassword = &v
}

// GetAccountUsername returns the AccountUsername field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetAccountUsername() string {
	if o == nil || o.AccountUsername == nil {
		var ret string
		return ret
	}
	return *o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetAccountUsernameOk() (*string, bool) {
	if o == nil || o.AccountUsername == nil {
		return nil, false
	}
	return o.AccountUsername, true
}

// HasAccountUsername returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasAccountUsername() bool {
	if o != nil && o.AccountUsername != nil {
		return true
	}

	return false
}

// SetAccountUsername gets a reference to the given string and assigns it to the AccountUsername field.
func (o *DynamicSecretUpdateSnowflake) SetAccountUsername(v string) {
	o.AccountUsername = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasDbName() bool {
	if o != nil && o.DbName != nil {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *DynamicSecretUpdateSnowflake) SetDbName(v string) {
	o.DbName = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretUpdateSnowflake) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretUpdateSnowflake) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretUpdateSnowflake) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretUpdateSnowflake) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretUpdateSnowflake) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *DynamicSecretUpdateSnowflake) SetNewName(v string) {
	o.NewName = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *DynamicSecretUpdateSnowflake) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *DynamicSecretUpdateSnowflake) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetPrivateKeyPassphrase() string {
	if o == nil || o.PrivateKeyPassphrase == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyPassphrase
}

// GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetPrivateKeyPassphraseOk() (*string, bool) {
	if o == nil || o.PrivateKeyPassphrase == nil {
		return nil, false
	}
	return o.PrivateKeyPassphrase, true
}

// HasPrivateKeyPassphrase returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasPrivateKeyPassphrase() bool {
	if o != nil && o.PrivateKeyPassphrase != nil {
		return true
	}

	return false
}

// SetPrivateKeyPassphrase gets a reference to the given string and assigns it to the PrivateKeyPassphrase field.
func (o *DynamicSecretUpdateSnowflake) SetPrivateKeyPassphrase(v string) {
	o.PrivateKeyPassphrase = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *DynamicSecretUpdateSnowflake) SetRole(v string) {
	o.Role = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretUpdateSnowflake) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretUpdateSnowflake) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretUpdateSnowflake) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretUpdateSnowflake) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretUpdateSnowflake) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *DynamicSecretUpdateSnowflake) GetWarehouse() string {
	if o == nil || o.Warehouse == nil {
		var ret string
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretUpdateSnowflake) GetWarehouseOk() (*string, bool) {
	if o == nil || o.Warehouse == nil {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *DynamicSecretUpdateSnowflake) HasWarehouse() bool {
	if o != nil && o.Warehouse != nil {
		return true
	}

	return false
}

// SetWarehouse gets a reference to the given string and assigns it to the Warehouse field.
func (o *DynamicSecretUpdateSnowflake) SetWarehouse(v string) {
	o.Warehouse = &v
}

func (o DynamicSecretUpdateSnowflake) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.AccountPassword != nil {
		toSerialize["account-password"] = o.AccountPassword
	}
	if o.AccountUsername != nil {
		toSerialize["account-username"] = o.AccountUsername
	}
	if o.DbName != nil {
		toSerialize["db-name"] = o.DbName
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
	if o.PrivateKey != nil {
		toSerialize["private-key"] = o.PrivateKey
	}
	if o.PrivateKeyPassphrase != nil {
		toSerialize["private-key-passphrase"] = o.PrivateKeyPassphrase
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
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
	if o.Warehouse != nil {
		toSerialize["warehouse"] = o.Warehouse
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretUpdateSnowflake struct {
	value *DynamicSecretUpdateSnowflake
	isSet bool
}

func (v NullableDynamicSecretUpdateSnowflake) Get() *DynamicSecretUpdateSnowflake {
	return v.value
}

func (v *NullableDynamicSecretUpdateSnowflake) Set(val *DynamicSecretUpdateSnowflake) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretUpdateSnowflake) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretUpdateSnowflake) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretUpdateSnowflake(val *DynamicSecretUpdateSnowflake) *NullableDynamicSecretUpdateSnowflake {
	return &NullableDynamicSecretUpdateSnowflake{value: val, isSet: true}
}

func (v NullableDynamicSecretUpdateSnowflake) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretUpdateSnowflake) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


