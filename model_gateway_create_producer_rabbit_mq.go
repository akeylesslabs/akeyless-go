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

// GatewayCreateProducerRabbitMQ gatewayCreateProducerRabbitMQ is a command that creates rabbitmq producer
type GatewayCreateProducerRabbitMQ struct {
	// Gateway url
	GatewayUrl *string `json:"gateway-url,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// RabbitMQ Admin password
	RabbitmqAdminPwd string `json:"rabbitmq-admin-pwd"`
	// RabbitMQ Admin User
	RabbitmqAdminUser string `json:"rabbitmq-admin-user"`
	// Server URI
	RabbitmqServerUri string `json:"rabbitmq-server-uri"`
	// User configuration permission
	RabbitmqUserConfPermission string `json:"rabbitmq-user-conf-permission"`
	// User read permission
	RabbitmqUserReadPermission string `json:"rabbitmq-user-read-permission"`
	// User Tags
	RabbitmqUserTags *string `json:"rabbitmq-user-tags,omitempty"`
	// User Virtual Host
	RabbitmqUserVhost *string `json:"rabbitmq-user-vhost,omitempty"`
	// User write permission
	RabbitmqUserWritePermission string `json:"rabbitmq-user-write-permission"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayCreateProducerRabbitMQ instantiates a new GatewayCreateProducerRabbitMQ object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerRabbitMQ(name string, rabbitmqAdminPwd string, rabbitmqAdminUser string, rabbitmqServerUri string, rabbitmqUserConfPermission string, rabbitmqUserReadPermission string, rabbitmqUserWritePermission string, ) *GatewayCreateProducerRabbitMQ {
	this := GatewayCreateProducerRabbitMQ{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	this.Name = name
	this.RabbitmqAdminPwd = rabbitmqAdminPwd
	this.RabbitmqAdminUser = rabbitmqAdminUser
	this.RabbitmqServerUri = rabbitmqServerUri
	this.RabbitmqUserConfPermission = rabbitmqUserConfPermission
	this.RabbitmqUserReadPermission = rabbitmqUserReadPermission
	this.RabbitmqUserWritePermission = rabbitmqUserWritePermission
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerRabbitMQWithDefaults instantiates a new GatewayCreateProducerRabbitMQ object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerRabbitMQWithDefaults() *GatewayCreateProducerRabbitMQ {
	this := GatewayCreateProducerRabbitMQ{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetGatewayUrl returns the GatewayUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerRabbitMQ) GetGatewayUrl() string {
	if o == nil || o.GatewayUrl == nil {
		var ret string
		return ret
	}
	return *o.GatewayUrl
}

// GetGatewayUrlOk returns a tuple with the GatewayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetGatewayUrlOk() (*string, bool) {
	if o == nil || o.GatewayUrl == nil {
		return nil, false
	}
	return o.GatewayUrl, true
}

// HasGatewayUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerRabbitMQ) HasGatewayUrl() bool {
	if o != nil && o.GatewayUrl != nil {
		return true
	}

	return false
}

// SetGatewayUrl gets a reference to the given string and assigns it to the GatewayUrl field.
func (o *GatewayCreateProducerRabbitMQ) SetGatewayUrl(v string) {
	o.GatewayUrl = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerRabbitMQ) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerRabbitMQ) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerRabbitMQ) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerRabbitMQ) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerRabbitMQ) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetRabbitmqAdminPwd returns the RabbitmqAdminPwd field value
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqAdminPwd() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RabbitmqAdminPwd
}

// GetRabbitmqAdminPwdOk returns a tuple with the RabbitmqAdminPwd field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqAdminPwdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RabbitmqAdminPwd, true
}

// SetRabbitmqAdminPwd sets field value
func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqAdminPwd(v string) {
	o.RabbitmqAdminPwd = v
}

// GetRabbitmqAdminUser returns the RabbitmqAdminUser field value
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqAdminUser() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RabbitmqAdminUser
}

// GetRabbitmqAdminUserOk returns a tuple with the RabbitmqAdminUser field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqAdminUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RabbitmqAdminUser, true
}

// SetRabbitmqAdminUser sets field value
func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqAdminUser(v string) {
	o.RabbitmqAdminUser = v
}

// GetRabbitmqServerUri returns the RabbitmqServerUri field value
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqServerUri() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RabbitmqServerUri
}

// GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqServerUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RabbitmqServerUri, true
}

// SetRabbitmqServerUri sets field value
func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqServerUri(v string) {
	o.RabbitmqServerUri = v
}

// GetRabbitmqUserConfPermission returns the RabbitmqUserConfPermission field value
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserConfPermission() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RabbitmqUserConfPermission
}

// GetRabbitmqUserConfPermissionOk returns a tuple with the RabbitmqUserConfPermission field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserConfPermissionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RabbitmqUserConfPermission, true
}

// SetRabbitmqUserConfPermission sets field value
func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserConfPermission(v string) {
	o.RabbitmqUserConfPermission = v
}

// GetRabbitmqUserReadPermission returns the RabbitmqUserReadPermission field value
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserReadPermission() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RabbitmqUserReadPermission
}

// GetRabbitmqUserReadPermissionOk returns a tuple with the RabbitmqUserReadPermission field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserReadPermissionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RabbitmqUserReadPermission, true
}

// SetRabbitmqUserReadPermission sets field value
func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserReadPermission(v string) {
	o.RabbitmqUserReadPermission = v
}

// GetRabbitmqUserTags returns the RabbitmqUserTags field value if set, zero value otherwise.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserTags() string {
	if o == nil || o.RabbitmqUserTags == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqUserTags
}

// GetRabbitmqUserTagsOk returns a tuple with the RabbitmqUserTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserTagsOk() (*string, bool) {
	if o == nil || o.RabbitmqUserTags == nil {
		return nil, false
	}
	return o.RabbitmqUserTags, true
}

// HasRabbitmqUserTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqUserTags() bool {
	if o != nil && o.RabbitmqUserTags != nil {
		return true
	}

	return false
}

// SetRabbitmqUserTags gets a reference to the given string and assigns it to the RabbitmqUserTags field.
func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserTags(v string) {
	o.RabbitmqUserTags = &v
}

// GetRabbitmqUserVhost returns the RabbitmqUserVhost field value if set, zero value otherwise.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserVhost() string {
	if o == nil || o.RabbitmqUserVhost == nil {
		var ret string
		return ret
	}
	return *o.RabbitmqUserVhost
}

// GetRabbitmqUserVhostOk returns a tuple with the RabbitmqUserVhost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserVhostOk() (*string, bool) {
	if o == nil || o.RabbitmqUserVhost == nil {
		return nil, false
	}
	return o.RabbitmqUserVhost, true
}

// HasRabbitmqUserVhost returns a boolean if a field has been set.
func (o *GatewayCreateProducerRabbitMQ) HasRabbitmqUserVhost() bool {
	if o != nil && o.RabbitmqUserVhost != nil {
		return true
	}

	return false
}

// SetRabbitmqUserVhost gets a reference to the given string and assigns it to the RabbitmqUserVhost field.
func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserVhost(v string) {
	o.RabbitmqUserVhost = &v
}

// GetRabbitmqUserWritePermission returns the RabbitmqUserWritePermission field value
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserWritePermission() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RabbitmqUserWritePermission
}

// GetRabbitmqUserWritePermissionOk returns a tuple with the RabbitmqUserWritePermission field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetRabbitmqUserWritePermissionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RabbitmqUserWritePermission, true
}

// SetRabbitmqUserWritePermission sets field value
func (o *GatewayCreateProducerRabbitMQ) SetRabbitmqUserWritePermission(v string) {
	o.RabbitmqUserWritePermission = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerRabbitMQ) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerRabbitMQ) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerRabbitMQ) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerRabbitMQ) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerRabbitMQ) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerRabbitMQ) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerRabbitMQ) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerRabbitMQ) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerRabbitMQ) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerRabbitMQ) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerRabbitMQ) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayUrl != nil {
		toSerialize["gateway-url"] = o.GatewayUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if true {
		toSerialize["rabbitmq-admin-pwd"] = o.RabbitmqAdminPwd
	}
	if true {
		toSerialize["rabbitmq-admin-user"] = o.RabbitmqAdminUser
	}
	if true {
		toSerialize["rabbitmq-server-uri"] = o.RabbitmqServerUri
	}
	if true {
		toSerialize["rabbitmq-user-conf-permission"] = o.RabbitmqUserConfPermission
	}
	if true {
		toSerialize["rabbitmq-user-read-permission"] = o.RabbitmqUserReadPermission
	}
	if o.RabbitmqUserTags != nil {
		toSerialize["rabbitmq-user-tags"] = o.RabbitmqUserTags
	}
	if o.RabbitmqUserVhost != nil {
		toSerialize["rabbitmq-user-vhost"] = o.RabbitmqUserVhost
	}
	if true {
		toSerialize["rabbitmq-user-write-permission"] = o.RabbitmqUserWritePermission
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

type NullableGatewayCreateProducerRabbitMQ struct {
	value *GatewayCreateProducerRabbitMQ
	isSet bool
}

func (v NullableGatewayCreateProducerRabbitMQ) Get() *GatewayCreateProducerRabbitMQ {
	return v.value
}

func (v *NullableGatewayCreateProducerRabbitMQ) Set(val *GatewayCreateProducerRabbitMQ) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerRabbitMQ) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerRabbitMQ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerRabbitMQ(val *GatewayCreateProducerRabbitMQ) *NullableGatewayCreateProducerRabbitMQ {
	return &NullableGatewayCreateProducerRabbitMQ{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerRabbitMQ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerRabbitMQ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


