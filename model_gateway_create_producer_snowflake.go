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

// GatewayCreateProducerSnowflake gatewayCreateProducerSnowflakeCmd is a command that creates a Snowflake producer
type GatewayCreateProducerSnowflake struct {
	// Account name
	Account string `json:"account"`
	// Database name
	DbName string `json:"db-name"`
	// Producer name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// User role
	Role *string `json:"role,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
	// Warehouse name
	Warehouse *string `json:"warehouse,omitempty"`
}

// NewGatewayCreateProducerSnowflake instantiates a new GatewayCreateProducerSnowflake object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerSnowflake(account string, dbName string, name string, ) *GatewayCreateProducerSnowflake {
	this := GatewayCreateProducerSnowflake{}
	this.Account = account
	this.DbName = dbName
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
	var userTtl string = "24h"
	this.UserTtl = &userTtl
	return &this
}

// GetAccount returns the Account field value
func (o *GatewayCreateProducerSnowflake) GetAccount() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *GatewayCreateProducerSnowflake) SetAccount(v string) {
	o.Account = v
}

// GetDbName returns the DbName field value
func (o *GatewayCreateProducerSnowflake) GetDbName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetDbNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbName, true
}

// SetDbName sets field value
func (o *GatewayCreateProducerSnowflake) SetDbName(v string) {
	o.DbName = v
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

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayCreateProducerSnowflake) SetPassword(v string) {
	o.Password = &v
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

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayCreateProducerSnowflake) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerSnowflake) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerSnowflake) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayCreateProducerSnowflake) SetUsername(v string) {
	o.Username = &v
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
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["db-name"] = o.DbName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
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


