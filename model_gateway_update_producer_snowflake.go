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

// GatewayUpdateProducerSnowflake gatewayUpdateProducerSnowflakeCmd is a command that updates a Snowflake producer
type GatewayUpdateProducerSnowflake struct {
	// Account name
	Account *string `json:"account,omitempty"`
	// Database Password
	AccountPassword *string `json:"account-password,omitempty"`
	// Database Username
	AccountUsername *string `json:"account-username,omitempty"`
	// Database name
	DbName *string `json:"db-name,omitempty"`
	// Protection from accidental deletion of this item
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Producer name
	NewName *string `json:"new-name,omitempty"`
	// User role
	Role *string `json:"role,omitempty"`
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
	// Warehouse name
	Warehouse *string `json:"warehouse,omitempty"`
}

// NewGatewayUpdateProducerSnowflake instantiates a new GatewayUpdateProducerSnowflake object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateProducerSnowflake(name string, ) *GatewayUpdateProducerSnowflake {
	this := GatewayUpdateProducerSnowflake{}
	this.Name = name
	var userTtl string = "24h"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayUpdateProducerSnowflakeWithDefaults instantiates a new GatewayUpdateProducerSnowflake object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateProducerSnowflakeWithDefaults() *GatewayUpdateProducerSnowflake {
	this := GatewayUpdateProducerSnowflake{}
	var userTtl string = "24h"
	this.UserTtl = &userTtl
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *GatewayUpdateProducerSnowflake) SetAccount(v string) {
	o.Account = &v
}

// GetAccountPassword returns the AccountPassword field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetAccountPassword() string {
	if o == nil || o.AccountPassword == nil {
		var ret string
		return ret
	}
	return *o.AccountPassword
}

// GetAccountPasswordOk returns a tuple with the AccountPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetAccountPasswordOk() (*string, bool) {
	if o == nil || o.AccountPassword == nil {
		return nil, false
	}
	return o.AccountPassword, true
}

// HasAccountPassword returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasAccountPassword() bool {
	if o != nil && o.AccountPassword != nil {
		return true
	}

	return false
}

// SetAccountPassword gets a reference to the given string and assigns it to the AccountPassword field.
func (o *GatewayUpdateProducerSnowflake) SetAccountPassword(v string) {
	o.AccountPassword = &v
}

// GetAccountUsername returns the AccountUsername field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetAccountUsername() string {
	if o == nil || o.AccountUsername == nil {
		var ret string
		return ret
	}
	return *o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetAccountUsernameOk() (*string, bool) {
	if o == nil || o.AccountUsername == nil {
		return nil, false
	}
	return o.AccountUsername, true
}

// HasAccountUsername returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasAccountUsername() bool {
	if o != nil && o.AccountUsername != nil {
		return true
	}

	return false
}

// SetAccountUsername gets a reference to the given string and assigns it to the AccountUsername field.
func (o *GatewayUpdateProducerSnowflake) SetAccountUsername(v string) {
	o.AccountUsername = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasDbName() bool {
	if o != nil && o.DbName != nil {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *GatewayUpdateProducerSnowflake) SetDbName(v string) {
	o.DbName = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateProducerSnowflake) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateProducerSnowflake) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateProducerSnowflake) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateProducerSnowflake) SetNewName(v string) {
	o.NewName = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *GatewayUpdateProducerSnowflake) SetRole(v string) {
	o.Role = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayUpdateProducerSnowflake) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayUpdateProducerSnowflake) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateProducerSnowflake) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateProducerSnowflake) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayUpdateProducerSnowflake) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *GatewayUpdateProducerSnowflake) GetWarehouse() string {
	if o == nil || o.Warehouse == nil {
		var ret string
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateProducerSnowflake) GetWarehouseOk() (*string, bool) {
	if o == nil || o.Warehouse == nil {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *GatewayUpdateProducerSnowflake) HasWarehouse() bool {
	if o != nil && o.Warehouse != nil {
		return true
	}

	return false
}

// SetWarehouse gets a reference to the given string and assigns it to the Warehouse field.
func (o *GatewayUpdateProducerSnowflake) SetWarehouse(v string) {
	o.Warehouse = &v
}

func (o GatewayUpdateProducerSnowflake) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
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

type NullableGatewayUpdateProducerSnowflake struct {
	value *GatewayUpdateProducerSnowflake
	isSet bool
}

func (v NullableGatewayUpdateProducerSnowflake) Get() *GatewayUpdateProducerSnowflake {
	return v.value
}

func (v *NullableGatewayUpdateProducerSnowflake) Set(val *GatewayUpdateProducerSnowflake) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateProducerSnowflake) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateProducerSnowflake) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateProducerSnowflake(val *GatewayUpdateProducerSnowflake) *NullableGatewayUpdateProducerSnowflake {
	return &NullableGatewayUpdateProducerSnowflake{value: val, isSet: true}
}

func (v NullableGatewayUpdateProducerSnowflake) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateProducerSnowflake) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


