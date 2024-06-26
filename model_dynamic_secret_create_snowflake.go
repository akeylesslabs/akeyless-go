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

// DynamicSecretCreateSnowflake GatewayCreateProducerSnowflakeCmd is a command that creates a Snowflake dynamic secret
type DynamicSecretCreateSnowflake struct {
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

// NewDynamicSecretCreateSnowflake instantiates a new DynamicSecretCreateSnowflake object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretCreateSnowflake(name string, ) *DynamicSecretCreateSnowflake {
	this := DynamicSecretCreateSnowflake{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "24h"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretCreateSnowflakeWithDefaults instantiates a new DynamicSecretCreateSnowflake object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretCreateSnowflakeWithDefaults() *DynamicSecretCreateSnowflake {
	this := DynamicSecretCreateSnowflake{}
	var json bool = false
	this.Json = &json
	var userTtl string = "24h"
	this.UserTtl = &userTtl
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *DynamicSecretCreateSnowflake) SetAccount(v string) {
	o.Account = &v
}

// GetAccountPassword returns the AccountPassword field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetAccountPassword() string {
	if o == nil || o.AccountPassword == nil {
		var ret string
		return ret
	}
	return *o.AccountPassword
}

// GetAccountPasswordOk returns a tuple with the AccountPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetAccountPasswordOk() (*string, bool) {
	if o == nil || o.AccountPassword == nil {
		return nil, false
	}
	return o.AccountPassword, true
}

// HasAccountPassword returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasAccountPassword() bool {
	if o != nil && o.AccountPassword != nil {
		return true
	}

	return false
}

// SetAccountPassword gets a reference to the given string and assigns it to the AccountPassword field.
func (o *DynamicSecretCreateSnowflake) SetAccountPassword(v string) {
	o.AccountPassword = &v
}

// GetAccountUsername returns the AccountUsername field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetAccountUsername() string {
	if o == nil || o.AccountUsername == nil {
		var ret string
		return ret
	}
	return *o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetAccountUsernameOk() (*string, bool) {
	if o == nil || o.AccountUsername == nil {
		return nil, false
	}
	return o.AccountUsername, true
}

// HasAccountUsername returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasAccountUsername() bool {
	if o != nil && o.AccountUsername != nil {
		return true
	}

	return false
}

// SetAccountUsername gets a reference to the given string and assigns it to the AccountUsername field.
func (o *DynamicSecretCreateSnowflake) SetAccountUsername(v string) {
	o.AccountUsername = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasDbName() bool {
	if o != nil && o.DbName != nil {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *DynamicSecretCreateSnowflake) SetDbName(v string) {
	o.DbName = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretCreateSnowflake) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretCreateSnowflake) SetDescription(v string) {
	o.Description = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretCreateSnowflake) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretCreateSnowflake) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretCreateSnowflake) SetName(v string) {
	o.Name = v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *DynamicSecretCreateSnowflake) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *DynamicSecretCreateSnowflake) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetPrivateKeyPassphrase() string {
	if o == nil || o.PrivateKeyPassphrase == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyPassphrase
}

// GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetPrivateKeyPassphraseOk() (*string, bool) {
	if o == nil || o.PrivateKeyPassphrase == nil {
		return nil, false
	}
	return o.PrivateKeyPassphrase, true
}

// HasPrivateKeyPassphrase returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasPrivateKeyPassphrase() bool {
	if o != nil && o.PrivateKeyPassphrase != nil {
		return true
	}

	return false
}

// SetPrivateKeyPassphrase gets a reference to the given string and assigns it to the PrivateKeyPassphrase field.
func (o *DynamicSecretCreateSnowflake) SetPrivateKeyPassphrase(v string) {
	o.PrivateKeyPassphrase = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *DynamicSecretCreateSnowflake) SetRole(v string) {
	o.Role = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretCreateSnowflake) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretCreateSnowflake) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretCreateSnowflake) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretCreateSnowflake) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretCreateSnowflake) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *DynamicSecretCreateSnowflake) GetWarehouse() string {
	if o == nil || o.Warehouse == nil {
		var ret string
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateSnowflake) GetWarehouseOk() (*string, bool) {
	if o == nil || o.Warehouse == nil {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *DynamicSecretCreateSnowflake) HasWarehouse() bool {
	if o != nil && o.Warehouse != nil {
		return true
	}

	return false
}

// SetWarehouse gets a reference to the given string and assigns it to the Warehouse field.
func (o *DynamicSecretCreateSnowflake) SetWarehouse(v string) {
	o.Warehouse = &v
}

func (o DynamicSecretCreateSnowflake) MarshalJSON() ([]byte, error) {
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

type NullableDynamicSecretCreateSnowflake struct {
	value *DynamicSecretCreateSnowflake
	isSet bool
}

func (v NullableDynamicSecretCreateSnowflake) Get() *DynamicSecretCreateSnowflake {
	return v.value
}

func (v *NullableDynamicSecretCreateSnowflake) Set(val *DynamicSecretCreateSnowflake) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretCreateSnowflake) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretCreateSnowflake) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretCreateSnowflake(val *DynamicSecretCreateSnowflake) *NullableDynamicSecretCreateSnowflake {
	return &NullableDynamicSecretCreateSnowflake{value: val, isSet: true}
}

func (v NullableDynamicSecretCreateSnowflake) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretCreateSnowflake) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


