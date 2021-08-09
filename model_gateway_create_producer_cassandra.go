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

// GatewayCreateProducerCassandra gatewayCreateProducerCassandra is a command that creates a Cassandra producer
type GatewayCreateProducerCassandra struct {
	// Cassandra creation statements
	CassandraCreationStatements *string `json:"cassandra-creation-statements,omitempty"`
	// Cassandra hosts IP or addresses, comma separated
	CassandraHosts *string `json:"cassandra-hosts,omitempty"`
	// Cassandra superuser password
	CassandraPassword *string `json:"cassandra-password,omitempty"`
	// Cassandra port
	CassandraPort *string `json:"cassandra-port,omitempty"`
	// Cassandra superuser username
	CassandraUsername *string `json:"cassandra-username,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Required only when the authentication process requires a username and password
	Password *string `json:"password,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
	// Required only when the authentication process requires a username and password
	Username *string `json:"username,omitempty"`
}

// NewGatewayCreateProducerCassandra instantiates a new GatewayCreateProducerCassandra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerCassandra(name string, ) *GatewayCreateProducerCassandra {
	this := GatewayCreateProducerCassandra{}
	var cassandraPort string = "9042"
	this.CassandraPort = &cassandraPort
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerCassandraWithDefaults instantiates a new GatewayCreateProducerCassandra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerCassandraWithDefaults() *GatewayCreateProducerCassandra {
	this := GatewayCreateProducerCassandra{}
	var cassandraPort string = "9042"
	this.CassandraPort = &cassandraPort
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetCassandraCreationStatements returns the CassandraCreationStatements field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetCassandraCreationStatements() string {
	if o == nil || o.CassandraCreationStatements == nil {
		var ret string
		return ret
	}
	return *o.CassandraCreationStatements
}

// GetCassandraCreationStatementsOk returns a tuple with the CassandraCreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetCassandraCreationStatementsOk() (*string, bool) {
	if o == nil || o.CassandraCreationStatements == nil {
		return nil, false
	}
	return o.CassandraCreationStatements, true
}

// HasCassandraCreationStatements returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasCassandraCreationStatements() bool {
	if o != nil && o.CassandraCreationStatements != nil {
		return true
	}

	return false
}

// SetCassandraCreationStatements gets a reference to the given string and assigns it to the CassandraCreationStatements field.
func (o *GatewayCreateProducerCassandra) SetCassandraCreationStatements(v string) {
	o.CassandraCreationStatements = &v
}

// GetCassandraHosts returns the CassandraHosts field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetCassandraHosts() string {
	if o == nil || o.CassandraHosts == nil {
		var ret string
		return ret
	}
	return *o.CassandraHosts
}

// GetCassandraHostsOk returns a tuple with the CassandraHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetCassandraHostsOk() (*string, bool) {
	if o == nil || o.CassandraHosts == nil {
		return nil, false
	}
	return o.CassandraHosts, true
}

// HasCassandraHosts returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasCassandraHosts() bool {
	if o != nil && o.CassandraHosts != nil {
		return true
	}

	return false
}

// SetCassandraHosts gets a reference to the given string and assigns it to the CassandraHosts field.
func (o *GatewayCreateProducerCassandra) SetCassandraHosts(v string) {
	o.CassandraHosts = &v
}

// GetCassandraPassword returns the CassandraPassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetCassandraPassword() string {
	if o == nil || o.CassandraPassword == nil {
		var ret string
		return ret
	}
	return *o.CassandraPassword
}

// GetCassandraPasswordOk returns a tuple with the CassandraPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetCassandraPasswordOk() (*string, bool) {
	if o == nil || o.CassandraPassword == nil {
		return nil, false
	}
	return o.CassandraPassword, true
}

// HasCassandraPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasCassandraPassword() bool {
	if o != nil && o.CassandraPassword != nil {
		return true
	}

	return false
}

// SetCassandraPassword gets a reference to the given string and assigns it to the CassandraPassword field.
func (o *GatewayCreateProducerCassandra) SetCassandraPassword(v string) {
	o.CassandraPassword = &v
}

// GetCassandraPort returns the CassandraPort field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetCassandraPort() string {
	if o == nil || o.CassandraPort == nil {
		var ret string
		return ret
	}
	return *o.CassandraPort
}

// GetCassandraPortOk returns a tuple with the CassandraPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetCassandraPortOk() (*string, bool) {
	if o == nil || o.CassandraPort == nil {
		return nil, false
	}
	return o.CassandraPort, true
}

// HasCassandraPort returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasCassandraPort() bool {
	if o != nil && o.CassandraPort != nil {
		return true
	}

	return false
}

// SetCassandraPort gets a reference to the given string and assigns it to the CassandraPort field.
func (o *GatewayCreateProducerCassandra) SetCassandraPort(v string) {
	o.CassandraPort = &v
}

// GetCassandraUsername returns the CassandraUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetCassandraUsername() string {
	if o == nil || o.CassandraUsername == nil {
		var ret string
		return ret
	}
	return *o.CassandraUsername
}

// GetCassandraUsernameOk returns a tuple with the CassandraUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetCassandraUsernameOk() (*string, bool) {
	if o == nil || o.CassandraUsername == nil {
		return nil, false
	}
	return o.CassandraUsername, true
}

// HasCassandraUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasCassandraUsername() bool {
	if o != nil && o.CassandraUsername != nil {
		return true
	}

	return false
}

// SetCassandraUsername gets a reference to the given string and assigns it to the CassandraUsername field.
func (o *GatewayCreateProducerCassandra) SetCassandraUsername(v string) {
	o.CassandraUsername = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerCassandra) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerCassandra) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GatewayCreateProducerCassandra) SetPassword(v string) {
	o.Password = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerCassandra) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerCassandra) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerCassandra) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerCassandra) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerCassandra) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GatewayCreateProducerCassandra) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerCassandra) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerCassandra) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GatewayCreateProducerCassandra) SetUsername(v string) {
	o.Username = &v
}

func (o GatewayCreateProducerCassandra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CassandraCreationStatements != nil {
		toSerialize["cassandra-creation-statements"] = o.CassandraCreationStatements
	}
	if o.CassandraHosts != nil {
		toSerialize["cassandra-hosts"] = o.CassandraHosts
	}
	if o.CassandraPassword != nil {
		toSerialize["cassandra-password"] = o.CassandraPassword
	}
	if o.CassandraPort != nil {
		toSerialize["cassandra-port"] = o.CassandraPort
	}
	if o.CassandraUsername != nil {
		toSerialize["cassandra-username"] = o.CassandraUsername
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayCreateProducerCassandra struct {
	value *GatewayCreateProducerCassandra
	isSet bool
}

func (v NullableGatewayCreateProducerCassandra) Get() *GatewayCreateProducerCassandra {
	return v.value
}

func (v *NullableGatewayCreateProducerCassandra) Set(val *GatewayCreateProducerCassandra) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerCassandra) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerCassandra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerCassandra(val *GatewayCreateProducerCassandra) *NullableGatewayCreateProducerCassandra {
	return &NullableGatewayCreateProducerCassandra{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerCassandra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerCassandra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


