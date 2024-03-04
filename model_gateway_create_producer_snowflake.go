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

// GatewayCreateProducerSnowflake GatewayCreateProducerSnowflakeCmd is a command that creates a Snowflake producer [Deprecated: Use dynamic-secret-create-snowflake command]
type GatewayCreateProducerSnowflake struct {
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

// NewGatewayCreateProducerSnowflake instantiates a new GatewayCreateProducerSnowflake object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerSnowflake(name string, ) *GatewayCreateProducerSnowflake {
	this := GatewayCreateProducerSnowflake{}
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "24h"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerSnowflakeWithDefaults instantiates a new GatewayCreateProducerSnowflake object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerSnowflakeWithDefaults() *GatewayCreateProducerSnowflake {
	this := GatewayCreateProducerSnowflake{}
	var json bool = false
	this.Json = &json
	var userTtl string = "24h"
	this.UserTtl = &userTtl
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *GatewayCreateProducerSnowflake) SetAccount(v string) {
	o.Account = &v
}

// GetAccountPassword returns the AccountPassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetAccountPassword() string {
	if o == nil || o.AccountPassword == nil {
		var ret string
		return ret
	}
	return *o.AccountPassword
}

// GetAccountPasswordOk returns a tuple with the AccountPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetAccountPasswordOk() (*string, bool) {
	if o == nil || o.AccountPassword == nil {
		return nil, false
	}
	return o.AccountPassword, true
}

// HasAccountPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasAccountPassword() bool {
	if o != nil && o.AccountPassword != nil {
		return true
	}

	return false
}

// SetAccountPassword gets a reference to the given string and assigns it to the AccountPassword field.
func (o *GatewayCreateProducerSnowflake) SetAccountPassword(v string) {
	o.AccountPassword = &v
}

// GetAccountUsername returns the AccountUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetAccountUsername() string {
	if o == nil || o.AccountUsername == nil {
		var ret string
		return ret
	}
	return *o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetAccountUsernameOk() (*string, bool) {
	if o == nil || o.AccountUsername == nil {
		return nil, false
	}
	return o.AccountUsername, true
}

// HasAccountUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasAccountUsername() bool {
	if o != nil && o.AccountUsername != nil {
		return true
	}

	return false
}

// SetAccountUsername gets a reference to the given string and assigns it to the AccountUsername field.
func (o *GatewayCreateProducerSnowflake) SetAccountUsername(v string) {
	o.AccountUsername = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasDbName() bool {
	if o != nil && o.DbName != nil {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *GatewayCreateProducerSnowflake) SetDbName(v string) {
	o.DbName = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerSnowflake) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerSnowflake) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerSnowflake) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerSnowflake) SetName(v string) {
	o.Name = v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *GatewayCreateProducerSnowflake) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *GatewayCreateProducerSnowflake) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetPrivateKeyPassphrase() string {
	if o == nil || o.PrivateKeyPassphrase == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyPassphrase
}

// GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetPrivateKeyPassphraseOk() (*string, bool) {
	if o == nil || o.PrivateKeyPassphrase == nil {
		return nil, false
	}
	return o.PrivateKeyPassphrase, true
}

// HasPrivateKeyPassphrase returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasPrivateKeyPassphrase() bool {
	if o != nil && o.PrivateKeyPassphrase != nil {
		return true
	}

	return false
}

// SetPrivateKeyPassphrase gets a reference to the given string and assigns it to the PrivateKeyPassphrase field.
func (o *GatewayCreateProducerSnowflake) SetPrivateKeyPassphrase(v string) {
	o.PrivateKeyPassphrase = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *GatewayCreateProducerSnowflake) SetRole(v string) {
	o.Role = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerSnowflake) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerSnowflake) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerSnowflake) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerSnowflake) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerSnowflake) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetWarehouse() string {
	if o == nil || o.Warehouse == nil {
		var ret string
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetWarehouseOk() (*string, bool) {
	if o == nil || o.Warehouse == nil {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasWarehouse() bool {
	if o != nil && o.Warehouse != nil {
		return true
	}

	return false
}

// SetWarehouse gets a reference to the given string and assigns it to the Warehouse field.
func (o *GatewayCreateProducerSnowflake) SetWarehouse(v string) {
	o.Warehouse = &v
}

func (o GatewayCreateProducerSnowflake) MarshalJSON() ([]byte, error) {
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

type NullableGatewayCreateProducerSnowflake struct {
	value *GatewayCreateProducerSnowflake
	isSet bool
}

func (v NullableGatewayCreateProducerSnowflake) Get() *GatewayCreateProducerSnowflake {
	return v.value
}

func (v *NullableGatewayCreateProducerSnowflake) Set(val *GatewayCreateProducerSnowflake) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerSnowflake) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerSnowflake) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerSnowflake(val *GatewayCreateProducerSnowflake) *NullableGatewayCreateProducerSnowflake {
	return &NullableGatewayCreateProducerSnowflake{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerSnowflake) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerSnowflake) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


