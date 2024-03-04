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

// DynamicSecretCreateRedis dynamicSecretCreateRedis is a command that creates Redis dynamic secret
type DynamicSecretCreateRedis struct {
	// A JSON array list of redis ACL rules to attach to the created user. For available rules see the ACL CAT command https://redis.io/commands/acl-cat By default the user will have permissions to read all keys '[\"~*\", \"+@read\"]'
	AclRules *string `json:"acl-rules,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Redis Host
	Host *string `json:"host,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Dynamic secret name
	Name string `json:"name"`
	// Redis Password
	Password *string `json:"password,omitempty"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// Redis Port
	Port *string `json:"port,omitempty"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Enable/Disable SSL [true/false]
	Ssl *bool `json:"ssl,omitempty"`
	// SSL CA certificate in base64 encoding generated from a trusted Certificate Authority (CA)
	SslCertificate *string `json:"ssl-certificate,omitempty"`
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
	// Redis Username
	Username *string `json:"username,omitempty"`
}

// NewDynamicSecretCreateRedis instantiates a new DynamicSecretCreateRedis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSecretCreateRedis(name string, ) *DynamicSecretCreateRedis {
	this := DynamicSecretCreateRedis{}
	var host string = "127.0.0.1"
	this.Host = &host
	var json bool = false
	this.Json = &json
	this.Name = name
	var port string = "6379"
	this.Port = &port
	var ssl bool = false
	this.Ssl = &ssl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewDynamicSecretCreateRedisWithDefaults instantiates a new DynamicSecretCreateRedis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSecretCreateRedisWithDefaults() *DynamicSecretCreateRedis {
	this := DynamicSecretCreateRedis{}
	var host string = "127.0.0.1"
	this.Host = &host
	var json bool = false
	this.Json = &json
	var port string = "6379"
	this.Port = &port
	var ssl bool = false
	this.Ssl = &ssl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetAclRules returns the AclRules field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetAclRules() string {
	if o == nil || o.AclRules == nil {
		var ret string
		return ret
	}
	return *o.AclRules
}

// GetAclRulesOk returns a tuple with the AclRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetAclRulesOk() (*string, bool) {
	if o == nil || o.AclRules == nil {
		return nil, false
	}
	return o.AclRules, true
}

// HasAclRules returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasAclRules() bool {
	if o != nil && o.AclRules != nil {
		return true
	}

	return false
}

// SetAclRules gets a reference to the given string and assigns it to the AclRules field.
func (o *DynamicSecretCreateRedis) SetAclRules(v string) {
	o.AclRules = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *DynamicSecretCreateRedis) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicSecretCreateRedis) SetDescription(v string) {
	o.Description = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *DynamicSecretCreateRedis) SetHost(v string) {
	o.Host = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *DynamicSecretCreateRedis) SetJson(v bool) {
	o.Json = &v
}

// GetName returns the Name field value
func (o *DynamicSecretCreateRedis) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicSecretCreateRedis) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DynamicSecretCreateRedis) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *DynamicSecretCreateRedis) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *DynamicSecretCreateRedis) SetPort(v string) {
	o.Port = &v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *DynamicSecretCreateRedis) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetSsl() bool {
	if o == nil || o.Ssl == nil {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetSslOk() (*bool, bool) {
	if o == nil || o.Ssl == nil {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasSsl() bool {
	if o != nil && o.Ssl != nil {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *DynamicSecretCreateRedis) SetSsl(v bool) {
	o.Ssl = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetSslCertificate() string {
	if o == nil || o.SslCertificate == nil {
		var ret string
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetSslCertificateOk() (*string, bool) {
	if o == nil || o.SslCertificate == nil {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasSslCertificate() bool {
	if o != nil && o.SslCertificate != nil {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given string and assigns it to the SslCertificate field.
func (o *DynamicSecretCreateRedis) SetSslCertificate(v string) {
	o.SslCertificate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DynamicSecretCreateRedis) SetTags(v []string) {
	o.Tags = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *DynamicSecretCreateRedis) SetTargetName(v string) {
	o.TargetName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DynamicSecretCreateRedis) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *DynamicSecretCreateRedis) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *DynamicSecretCreateRedis) SetUserTtl(v string) {
	o.UserTtl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DynamicSecretCreateRedis) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSecretCreateRedis) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DynamicSecretCreateRedis) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DynamicSecretCreateRedis) SetUsername(v string) {
	o.Username = &v
}

func (o DynamicSecretCreateRedis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AclRules != nil {
		toSerialize["acl-rules"] = o.AclRules
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PasswordLength != nil {
		toSerialize["password-length"] = o.PasswordLength
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.Ssl != nil {
		toSerialize["ssl"] = o.Ssl
	}
	if o.SslCertificate != nil {
		toSerialize["ssl-certificate"] = o.SslCertificate
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicSecretCreateRedis struct {
	value *DynamicSecretCreateRedis
	isSet bool
}

func (v NullableDynamicSecretCreateRedis) Get() *DynamicSecretCreateRedis {
	return v.value
}

func (v *NullableDynamicSecretCreateRedis) Set(val *DynamicSecretCreateRedis) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSecretCreateRedis) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSecretCreateRedis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSecretCreateRedis(val *DynamicSecretCreateRedis) *NullableDynamicSecretCreateRedis {
	return &NullableDynamicSecretCreateRedis{value: val, isSet: true}
}

func (v NullableDynamicSecretCreateRedis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSecretCreateRedis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


