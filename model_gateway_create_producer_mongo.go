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

// GatewayCreateProducerMongo gatewayCreateProducerMongo is a command that creates mongo producer
type GatewayCreateProducerMongo struct {
	// Gateway url
	GatewayUrl *string `json:"gateway-url,omitempty"`
	// MongoDB Name
	MongodbName string `json:"mongodb-name"`
	// MongoDB Roles
	MongodbRoles *string `json:"mongodb-roles,omitempty"`
	// Server URI
	MongodbServerUri string `json:"mongodb-server-uri"`
	// Producer name
	Name string `json:"name"`
	// Dynamic producer encryption key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// User TTL
	UserTtl *string `json:"user-ttl,omitempty"`
}

// NewGatewayCreateProducerMongo instantiates a new GatewayCreateProducerMongo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreateProducerMongo(mongodbName string, mongodbServerUri string, name string, ) *GatewayCreateProducerMongo {
	this := GatewayCreateProducerMongo{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	this.MongodbName = mongodbName
	var mongodbRoles string = "[]"
	this.MongodbRoles = &mongodbRoles
	this.MongodbServerUri = mongodbServerUri
	this.Name = name
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerMongoWithDefaults instantiates a new GatewayCreateProducerMongo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerMongoWithDefaults() *GatewayCreateProducerMongo {
	this := GatewayCreateProducerMongo{}
	var gatewayUrl string = "http://localhost:8000"
	this.GatewayUrl = &gatewayUrl
	var mongodbRoles string = "[]"
	this.MongodbRoles = &mongodbRoles
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetGatewayUrl returns the GatewayUrl field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetGatewayUrl() string {
	if o == nil || o.GatewayUrl == nil {
		var ret string
		return ret
	}
	return *o.GatewayUrl
}

// GetGatewayUrlOk returns a tuple with the GatewayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetGatewayUrlOk() (*string, bool) {
	if o == nil || o.GatewayUrl == nil {
		return nil, false
	}
	return o.GatewayUrl, true
}

// HasGatewayUrl returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasGatewayUrl() bool {
	if o != nil && o.GatewayUrl != nil {
		return true
	}

	return false
}

// SetGatewayUrl gets a reference to the given string and assigns it to the GatewayUrl field.
func (o *GatewayCreateProducerMongo) SetGatewayUrl(v string) {
	o.GatewayUrl = &v
}

// GetMongodbName returns the MongodbName field value
func (o *GatewayCreateProducerMongo) GetMongodbName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MongodbName
}

// GetMongodbNameOk returns a tuple with the MongodbName field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MongodbName, true
}

// SetMongodbName sets field value
func (o *GatewayCreateProducerMongo) SetMongodbName(v string) {
	o.MongodbName = v
}

// GetMongodbRoles returns the MongodbRoles field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbRoles() string {
	if o == nil || o.MongodbRoles == nil {
		var ret string
		return ret
	}
	return *o.MongodbRoles
}

// GetMongodbRolesOk returns a tuple with the MongodbRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbRolesOk() (*string, bool) {
	if o == nil || o.MongodbRoles == nil {
		return nil, false
	}
	return o.MongodbRoles, true
}

// HasMongodbRoles returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbRoles() bool {
	if o != nil && o.MongodbRoles != nil {
		return true
	}

	return false
}

// SetMongodbRoles gets a reference to the given string and assigns it to the MongodbRoles field.
func (o *GatewayCreateProducerMongo) SetMongodbRoles(v string) {
	o.MongodbRoles = &v
}

// GetMongodbServerUri returns the MongodbServerUri field value
func (o *GatewayCreateProducerMongo) GetMongodbServerUri() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MongodbServerUri
}

// GetMongodbServerUriOk returns a tuple with the MongodbServerUri field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbServerUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MongodbServerUri, true
}

// SetMongodbServerUri sets field value
func (o *GatewayCreateProducerMongo) SetMongodbServerUri(v string) {
	o.MongodbServerUri = v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerMongo) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayCreateProducerMongo) SetName(v string) {
	o.Name = v
}

// GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetProducerEncryptionKeyName() string {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		var ret string
		return ret
	}
	return *o.ProducerEncryptionKeyName
}

// GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetProducerEncryptionKeyNameOk() (*string, bool) {
	if o == nil || o.ProducerEncryptionKeyName == nil {
		return nil, false
	}
	return o.ProducerEncryptionKeyName, true
}

// HasProducerEncryptionKeyName returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasProducerEncryptionKeyName() bool {
	if o != nil && o.ProducerEncryptionKeyName != nil {
		return true
	}

	return false
}

// SetProducerEncryptionKeyName gets a reference to the given string and assigns it to the ProducerEncryptionKeyName field.
func (o *GatewayCreateProducerMongo) SetProducerEncryptionKeyName(v string) {
	o.ProducerEncryptionKeyName = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayCreateProducerMongo) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayCreateProducerMongo) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserTtl returns the UserTtl field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetUserTtl() string {
	if o == nil || o.UserTtl == nil {
		var ret string
		return ret
	}
	return *o.UserTtl
}

// GetUserTtlOk returns a tuple with the UserTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetUserTtlOk() (*string, bool) {
	if o == nil || o.UserTtl == nil {
		return nil, false
	}
	return o.UserTtl, true
}

// HasUserTtl returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasUserTtl() bool {
	if o != nil && o.UserTtl != nil {
		return true
	}

	return false
}

// SetUserTtl gets a reference to the given string and assigns it to the UserTtl field.
func (o *GatewayCreateProducerMongo) SetUserTtl(v string) {
	o.UserTtl = &v
}

func (o GatewayCreateProducerMongo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayUrl != nil {
		toSerialize["gateway-url"] = o.GatewayUrl
	}
	if true {
		toSerialize["mongodb-name"] = o.MongodbName
	}
	if o.MongodbRoles != nil {
		toSerialize["mongodb-roles"] = o.MongodbRoles
	}
	if true {
		toSerialize["mongodb-server-uri"] = o.MongodbServerUri
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
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

type NullableGatewayCreateProducerMongo struct {
	value *GatewayCreateProducerMongo
	isSet bool
}

func (v NullableGatewayCreateProducerMongo) Get() *GatewayCreateProducerMongo {
	return v.value
}

func (v *NullableGatewayCreateProducerMongo) Set(val *GatewayCreateProducerMongo) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreateProducerMongo) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreateProducerMongo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreateProducerMongo(val *GatewayCreateProducerMongo) *NullableGatewayCreateProducerMongo {
	return &NullableGatewayCreateProducerMongo{value: val, isSet: true}
}

func (v NullableGatewayCreateProducerMongo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreateProducerMongo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

