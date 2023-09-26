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

// DsCreateCassandra dsCreateCassandra is a command that creates a Cassandra dynamic secret
type DsCreateCassandra struct {
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
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
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

// NewDsCreateCassandra instantiates a new DsCreateCassandra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDsCreateCassandra(name string, ) *DsCreateCassandra {
	this := DsCreateCassandra{}
	var cassandraPort string = "9042"
	this.CassandraPort = &cassandraPort
	var json bool = false
	this.Json = &json
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDsCreateCassandraWithDefaults instantiates a new DsCreateCassandra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDsCreateCassandraWithDefaults() *DsCreateCassandra {
	this := DsCreateCassandra{}
	var cassandraPort string = "9042"
	this.CassandraPort = &cassandraPort
	var json bool = false
	this.Json = &json
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetCassandraCreationStatements returns the CassandraCreationStatements field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetCassandraCreationStatements() string {
	if o == nil || o.CassandraCreationStatements == nil {
		var ret string
		return ret
	}
	return *o.CassandraCreationStatements
}

// GetCassandraCreationStatementsOk returns a tuple with the CassandraCreationStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetCassandraCreationStatementsOk() (*string, bool) {
	if o == nil || o.CassandraCreationStatements == nil {
		return nil, false
	}
	return o.CassandraCreationStatements, true
}

// HasCassandraCreationStatements returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasCassandraCreationStatements() bool {
	if o != nil && o.CassandraCreationStatements != nil {
		return true
	}

	return false
}

// SetCassandraCreationStatements gets a reference to the given string and assigns it to the CassandraCreationStatements field.
func (o *DsCreateCassandra) SetCassandraCreationStatements(v string) {
	o.CassandraCreationStatements = &v
}

// GetCassandraHosts returns the CassandraHosts field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetCassandraHosts() string {
	if o == nil || o.CassandraHosts == nil {
		var ret string
		return ret
	}
	return *o.CassandraHosts
}

// GetCassandraHostsOk returns a tuple with the CassandraHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetCassandraHostsOk() (*string, bool) {
	if o == nil || o.CassandraHosts == nil {
		return nil, false
	}
	return o.CassandraHosts, true
}

// HasCassandraHosts returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasCassandraHosts() bool {
	if o != nil && o.CassandraHosts != nil {
		return true
	}

	return false
}

// SetCassandraHosts gets a reference to the given string and assigns it to the CassandraHosts field.
func (o *DsCreateCassandra) SetCassandraHosts(v string) {
	o.CassandraHosts = &v
}

// GetCassandraPassword returns the CassandraPassword field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetCassandraPassword() string {
	if o == nil || o.CassandraPassword == nil {
		var ret string
		return ret
	}
	return *o.CassandraPassword
}

// GetCassandraPasswordOk returns a tuple with the CassandraPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetCassandraPasswordOk() (*string, bool) {
	if o == nil || o.CassandraPassword == nil {
		return nil, false
	}
	return o.CassandraPassword, true
}

// HasCassandraPassword returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasCassandraPassword() bool {
	if o != nil && o.CassandraPassword != nil {
		return true
	}

	return false
}

// SetCassandraPassword gets a reference to the given string and assigns it to the CassandraPassword field.
func (o *DsCreateCassandra) SetCassandraPassword(v string) {
	o.CassandraPassword = &v
}

// GetCassandraPort returns the CassandraPort field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetCassandraPort() string {
	if o == nil || o.CassandraPort == nil {
		var ret string
		return ret
	}
	return *o.CassandraPort
}

// GetCassandraPortOk returns a tuple with the CassandraPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetCassandraPortOk() (*string, bool) {
	if o == nil || o.CassandraPort == nil {
		return nil, false
	}
	return o.CassandraPort, true
}

// HasCassandraPort returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasCassandraPort() bool {
	if o != nil && o.CassandraPort != nil {
		return true
	}

	return false
}

// SetCassandraPort gets a reference to the given string and assigns it to the CassandraPort field.
func (o *DsCreateCassandra) SetCassandraPort(v string) {
	o.CassandraPort = &v
}

// GetCassandraUsername returns the CassandraUsername field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetCassandraUsername() string {
	if o == nil || o.CassandraUsername == nil {
		var ret string
		return ret
	}
	return *o.CassandraUsername
}

// GetCassandraUsernameOk returns a tuple with the CassandraUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetCassandraUsernameOk() (*string, bool) {
	if o == nil || o.CassandraUsername == nil {
		return nil, false
	}
	return o.CassandraUsername, true
}

// HasCassandraUsername returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasCassandraUsername() bool {
	if o != nil && o.CassandraUsername != nil {
		return true
	}

	return false
}

// SetCassandraUsername gets a reference to the given string and assigns it to the CassandraUsername field.
func (o *DsCreateCassandra) SetCassandraUsername(v string) {
	o.CassandraUsername = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DsCreateCassandra) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DsCreateCassandra) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DsCreateCassandra) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DsCreateCassandra) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DsCreateCassandra) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DsCreateCassandra) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DsCreateCassandra) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DsCreateCassandra) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DsCreateCassandra) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DsCreateCassandra) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DsCreateCassandra) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DsCreateCassandra) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DsCreateCassandra) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o DsCreateCassandra) MarshalJSON() ([]byte, error) {
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
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
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

type NullableDsCreateCassandra struct {
	value *DsCreateCassandra
	isSet bool
}

func (v NullableDsCreateCassandra) Get() *DsCreateCassandra {
	return v.value
}

func (v *NullableDsCreateCassandra) Set(val *DsCreateCassandra) {
	v.value = val
	v.isSet = true
}

func (v NullableDsCreateCassandra) IsSet() bool {
	return v.isSet
}

func (v *NullableDsCreateCassandra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDsCreateCassandra(val *DsCreateCassandra) *NullableDsCreateCassandra {
	return &NullableDsCreateCassandra{value: val, isSet: true}
}

func (v NullableDsCreateCassandra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDsCreateCassandra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


