/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GatewayCreateProducerMongo gatewayCreateProducerMongo is a command that creates either mongodb  producer or mongodb atlas producer
type GatewayCreateProducerMongo struct {
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// MongoDB Atlas private key
	MongodbAtlasApiPrivateKey *string `json:"mongodb-atlas-api-private-key,omitempty"`
	// MongoDB Atlas public key
	MongodbAtlasApiPublicKey *string `json:"mongodb-atlas-api-public-key,omitempty"`
	// MongoDB Atlas project ID
	MongodbAtlasProjectId *string `json:"mongodb-atlas-project-id,omitempty"`
	// MongoDB custom data
	MongodbCustomData *string `json:"mongodb-custom-data,omitempty"`
	// MongoDB server default authentication database
	MongodbDefaultAuthDb *string `json:"mongodb-default-auth-db,omitempty"`
	// MongoDB server host and port
	MongodbHostPort *string `json:"mongodb-host-port,omitempty"`
	// MongoDB Name
	MongodbName *string `json:"mongodb-name,omitempty"`
	// MongoDB server password. You will prompted to provide a password if it will not appear in CLI parameters
	MongodbPassword *string `json:"mongodb-password,omitempty"`
	// MongoDB Roles
	MongodbRoles *string `json:"mongodb-roles,omitempty"`
	// MongoDB server URI
	MongodbServerUri *string `json:"mongodb-server-uri,omitempty"`
	// MongoDB server URI options
	MongodbUriOptions *string `json:"mongodb-uri-options,omitempty"`
	// MongoDB server username
	MongodbUsername *string `json:"mongodb-username,omitempty"`
	// Producer name
	Name string `json:"name"`
	// Encrypt producer with following key
	ProducerEncryptionKeyName *string `json:"producer-encryption-key-name,omitempty"`
	// Path to the SSH Certificate Issuer for your Akeyless Bastion
	SecureAccessBastionIssuer *string `json:"secure-access-bastion-issuer,omitempty"`
	// Enable/Disable secure remote access [true/false]
	SecureAccessEnable *string `json:"secure-access-enable,omitempty"`
	// Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts)
	SecureAccessHost []string `json:"secure-access-host,omitempty"`
	// Enable Web Secure Remote Access
	SecureAccessWeb *bool `json:"secure-access-web,omitempty"`
	// Add tags attached to this object
	Tags []string `json:"tags,omitempty"`
	// Target name
	TargetName *string `json:"target-name,omitempty"`
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
func NewGatewayCreateProducerMongo(name string) *GatewayCreateProducerMongo {
	this := GatewayCreateProducerMongo{}
	var json bool = false
	this.Json = &json
	var mongodbRoles string = "[]"
	this.MongodbRoles = &mongodbRoles
	this.Name = name
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// NewGatewayCreateProducerMongoWithDefaults instantiates a new GatewayCreateProducerMongo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateProducerMongoWithDefaults() *GatewayCreateProducerMongo {
	this := GatewayCreateProducerMongo{}
	var json bool = false
	this.Json = &json
	var mongodbRoles string = "[]"
	this.MongodbRoles = &mongodbRoles
	var secureAccessWeb bool = false
	this.SecureAccessWeb = &secureAccessWeb
	var userTtl string = "60m"
	this.UserTtl = &userTtl
	return &this
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayCreateProducerMongo) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayCreateProducerMongo) SetJson(v bool) {
	o.Json = &v
}

// GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbAtlasApiPrivateKey() string {
	if o == nil || o.MongodbAtlasApiPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.MongodbAtlasApiPrivateKey
}

// GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbAtlasApiPrivateKeyOk() (*string, bool) {
	if o == nil || o.MongodbAtlasApiPrivateKey == nil {
		return nil, false
	}
	return o.MongodbAtlasApiPrivateKey, true
}

// HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbAtlasApiPrivateKey() bool {
	if o != nil && o.MongodbAtlasApiPrivateKey != nil {
		return true
	}

	return false
}

// SetMongodbAtlasApiPrivateKey gets a reference to the given string and assigns it to the MongodbAtlasApiPrivateKey field.
func (o *GatewayCreateProducerMongo) SetMongodbAtlasApiPrivateKey(v string) {
	o.MongodbAtlasApiPrivateKey = &v
}

// GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbAtlasApiPublicKey() string {
	if o == nil || o.MongodbAtlasApiPublicKey == nil {
		var ret string
		return ret
	}
	return *o.MongodbAtlasApiPublicKey
}

// GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbAtlasApiPublicKeyOk() (*string, bool) {
	if o == nil || o.MongodbAtlasApiPublicKey == nil {
		return nil, false
	}
	return o.MongodbAtlasApiPublicKey, true
}

// HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbAtlasApiPublicKey() bool {
	if o != nil && o.MongodbAtlasApiPublicKey != nil {
		return true
	}

	return false
}

// SetMongodbAtlasApiPublicKey gets a reference to the given string and assigns it to the MongodbAtlasApiPublicKey field.
func (o *GatewayCreateProducerMongo) SetMongodbAtlasApiPublicKey(v string) {
	o.MongodbAtlasApiPublicKey = &v
}

// GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbAtlasProjectId() string {
	if o == nil || o.MongodbAtlasProjectId == nil {
		var ret string
		return ret
	}
	return *o.MongodbAtlasProjectId
}

// GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbAtlasProjectIdOk() (*string, bool) {
	if o == nil || o.MongodbAtlasProjectId == nil {
		return nil, false
	}
	return o.MongodbAtlasProjectId, true
}

// HasMongodbAtlasProjectId returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbAtlasProjectId() bool {
	if o != nil && o.MongodbAtlasProjectId != nil {
		return true
	}

	return false
}

// SetMongodbAtlasProjectId gets a reference to the given string and assigns it to the MongodbAtlasProjectId field.
func (o *GatewayCreateProducerMongo) SetMongodbAtlasProjectId(v string) {
	o.MongodbAtlasProjectId = &v
}

// GetMongodbCustomData returns the MongodbCustomData field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbCustomData() string {
	if o == nil || o.MongodbCustomData == nil {
		var ret string
		return ret
	}
	return *o.MongodbCustomData
}

// GetMongodbCustomDataOk returns a tuple with the MongodbCustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbCustomDataOk() (*string, bool) {
	if o == nil || o.MongodbCustomData == nil {
		return nil, false
	}
	return o.MongodbCustomData, true
}

// HasMongodbCustomData returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbCustomData() bool {
	if o != nil && o.MongodbCustomData != nil {
		return true
	}

	return false
}

// SetMongodbCustomData gets a reference to the given string and assigns it to the MongodbCustomData field.
func (o *GatewayCreateProducerMongo) SetMongodbCustomData(v string) {
	o.MongodbCustomData = &v
}

// GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbDefaultAuthDb() string {
	if o == nil || o.MongodbDefaultAuthDb == nil {
		var ret string
		return ret
	}
	return *o.MongodbDefaultAuthDb
}

// GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbDefaultAuthDbOk() (*string, bool) {
	if o == nil || o.MongodbDefaultAuthDb == nil {
		return nil, false
	}
	return o.MongodbDefaultAuthDb, true
}

// HasMongodbDefaultAuthDb returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbDefaultAuthDb() bool {
	if o != nil && o.MongodbDefaultAuthDb != nil {
		return true
	}

	return false
}

// SetMongodbDefaultAuthDb gets a reference to the given string and assigns it to the MongodbDefaultAuthDb field.
func (o *GatewayCreateProducerMongo) SetMongodbDefaultAuthDb(v string) {
	o.MongodbDefaultAuthDb = &v
}

// GetMongodbHostPort returns the MongodbHostPort field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbHostPort() string {
	if o == nil || o.MongodbHostPort == nil {
		var ret string
		return ret
	}
	return *o.MongodbHostPort
}

// GetMongodbHostPortOk returns a tuple with the MongodbHostPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbHostPortOk() (*string, bool) {
	if o == nil || o.MongodbHostPort == nil {
		return nil, false
	}
	return o.MongodbHostPort, true
}

// HasMongodbHostPort returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbHostPort() bool {
	if o != nil && o.MongodbHostPort != nil {
		return true
	}

	return false
}

// SetMongodbHostPort gets a reference to the given string and assigns it to the MongodbHostPort field.
func (o *GatewayCreateProducerMongo) SetMongodbHostPort(v string) {
	o.MongodbHostPort = &v
}

// GetMongodbName returns the MongodbName field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbName() string {
	if o == nil || o.MongodbName == nil {
		var ret string
		return ret
	}
	return *o.MongodbName
}

// GetMongodbNameOk returns a tuple with the MongodbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbNameOk() (*string, bool) {
	if o == nil || o.MongodbName == nil {
		return nil, false
	}
	return o.MongodbName, true
}

// HasMongodbName returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbName() bool {
	if o != nil && o.MongodbName != nil {
		return true
	}

	return false
}

// SetMongodbName gets a reference to the given string and assigns it to the MongodbName field.
func (o *GatewayCreateProducerMongo) SetMongodbName(v string) {
	o.MongodbName = &v
}

// GetMongodbPassword returns the MongodbPassword field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbPassword() string {
	if o == nil || o.MongodbPassword == nil {
		var ret string
		return ret
	}
	return *o.MongodbPassword
}

// GetMongodbPasswordOk returns a tuple with the MongodbPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbPasswordOk() (*string, bool) {
	if o == nil || o.MongodbPassword == nil {
		return nil, false
	}
	return o.MongodbPassword, true
}

// HasMongodbPassword returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbPassword() bool {
	if o != nil && o.MongodbPassword != nil {
		return true
	}

	return false
}

// SetMongodbPassword gets a reference to the given string and assigns it to the MongodbPassword field.
func (o *GatewayCreateProducerMongo) SetMongodbPassword(v string) {
	o.MongodbPassword = &v
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

// GetMongodbServerUri returns the MongodbServerUri field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbServerUri() string {
	if o == nil || o.MongodbServerUri == nil {
		var ret string
		return ret
	}
	return *o.MongodbServerUri
}

// GetMongodbServerUriOk returns a tuple with the MongodbServerUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbServerUriOk() (*string, bool) {
	if o == nil || o.MongodbServerUri == nil {
		return nil, false
	}
	return o.MongodbServerUri, true
}

// HasMongodbServerUri returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbServerUri() bool {
	if o != nil && o.MongodbServerUri != nil {
		return true
	}

	return false
}

// SetMongodbServerUri gets a reference to the given string and assigns it to the MongodbServerUri field.
func (o *GatewayCreateProducerMongo) SetMongodbServerUri(v string) {
	o.MongodbServerUri = &v
}

// GetMongodbUriOptions returns the MongodbUriOptions field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbUriOptions() string {
	if o == nil || o.MongodbUriOptions == nil {
		var ret string
		return ret
	}
	return *o.MongodbUriOptions
}

// GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbUriOptionsOk() (*string, bool) {
	if o == nil || o.MongodbUriOptions == nil {
		return nil, false
	}
	return o.MongodbUriOptions, true
}

// HasMongodbUriOptions returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbUriOptions() bool {
	if o != nil && o.MongodbUriOptions != nil {
		return true
	}

	return false
}

// SetMongodbUriOptions gets a reference to the given string and assigns it to the MongodbUriOptions field.
func (o *GatewayCreateProducerMongo) SetMongodbUriOptions(v string) {
	o.MongodbUriOptions = &v
}

// GetMongodbUsername returns the MongodbUsername field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetMongodbUsername() string {
	if o == nil || o.MongodbUsername == nil {
		var ret string
		return ret
	}
	return *o.MongodbUsername
}

// GetMongodbUsernameOk returns a tuple with the MongodbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetMongodbUsernameOk() (*string, bool) {
	if o == nil || o.MongodbUsername == nil {
		return nil, false
	}
	return o.MongodbUsername, true
}

// HasMongodbUsername returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasMongodbUsername() bool {
	if o != nil && o.MongodbUsername != nil {
		return true
	}

	return false
}

// SetMongodbUsername gets a reference to the given string and assigns it to the MongodbUsername field.
func (o *GatewayCreateProducerMongo) SetMongodbUsername(v string) {
	o.MongodbUsername = &v
}

// GetName returns the Name field value
func (o *GatewayCreateProducerMongo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetNameOk() (*string, bool) {
	if o == nil {
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

// GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetSecureAccessBastionIssuer() string {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessBastionIssuer
}

// GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetSecureAccessBastionIssuerOk() (*string, bool) {
	if o == nil || o.SecureAccessBastionIssuer == nil {
		return nil, false
	}
	return o.SecureAccessBastionIssuer, true
}

// HasSecureAccessBastionIssuer returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasSecureAccessBastionIssuer() bool {
	if o != nil && o.SecureAccessBastionIssuer != nil {
		return true
	}

	return false
}

// SetSecureAccessBastionIssuer gets a reference to the given string and assigns it to the SecureAccessBastionIssuer field.
func (o *GatewayCreateProducerMongo) SetSecureAccessBastionIssuer(v string) {
	o.SecureAccessBastionIssuer = &v
}

// GetSecureAccessEnable returns the SecureAccessEnable field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetSecureAccessEnable() string {
	if o == nil || o.SecureAccessEnable == nil {
		var ret string
		return ret
	}
	return *o.SecureAccessEnable
}

// GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetSecureAccessEnableOk() (*string, bool) {
	if o == nil || o.SecureAccessEnable == nil {
		return nil, false
	}
	return o.SecureAccessEnable, true
}

// HasSecureAccessEnable returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasSecureAccessEnable() bool {
	if o != nil && o.SecureAccessEnable != nil {
		return true
	}

	return false
}

// SetSecureAccessEnable gets a reference to the given string and assigns it to the SecureAccessEnable field.
func (o *GatewayCreateProducerMongo) SetSecureAccessEnable(v string) {
	o.SecureAccessEnable = &v
}

// GetSecureAccessHost returns the SecureAccessHost field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetSecureAccessHost() []string {
	if o == nil || o.SecureAccessHost == nil {
		var ret []string
		return ret
	}
	return o.SecureAccessHost
}

// GetSecureAccessHostOk returns a tuple with the SecureAccessHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetSecureAccessHostOk() ([]string, bool) {
	if o == nil || o.SecureAccessHost == nil {
		return nil, false
	}
	return o.SecureAccessHost, true
}

// HasSecureAccessHost returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasSecureAccessHost() bool {
	if o != nil && o.SecureAccessHost != nil {
		return true
	}

	return false
}

// SetSecureAccessHost gets a reference to the given []string and assigns it to the SecureAccessHost field.
func (o *GatewayCreateProducerMongo) SetSecureAccessHost(v []string) {
	o.SecureAccessHost = v
}

// GetSecureAccessWeb returns the SecureAccessWeb field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetSecureAccessWeb() bool {
	if o == nil || o.SecureAccessWeb == nil {
		var ret bool
		return ret
	}
	return *o.SecureAccessWeb
}

// GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetSecureAccessWebOk() (*bool, bool) {
	if o == nil || o.SecureAccessWeb == nil {
		return nil, false
	}
	return o.SecureAccessWeb, true
}

// HasSecureAccessWeb returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasSecureAccessWeb() bool {
	if o != nil && o.SecureAccessWeb != nil {
		return true
	}

	return false
}

// SetSecureAccessWeb gets a reference to the given bool and assigns it to the SecureAccessWeb field.
func (o *GatewayCreateProducerMongo) SetSecureAccessWeb(v bool) {
	o.SecureAccessWeb = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GatewayCreateProducerMongo) SetTags(v []string) {
	o.Tags = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *GatewayCreateProducerMongo) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreateProducerMongo) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *GatewayCreateProducerMongo) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *GatewayCreateProducerMongo) SetTargetName(v string) {
	o.TargetName = &v
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
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.MongodbAtlasApiPrivateKey != nil {
		toSerialize["mongodb-atlas-api-private-key"] = o.MongodbAtlasApiPrivateKey
	}
	if o.MongodbAtlasApiPublicKey != nil {
		toSerialize["mongodb-atlas-api-public-key"] = o.MongodbAtlasApiPublicKey
	}
	if o.MongodbAtlasProjectId != nil {
		toSerialize["mongodb-atlas-project-id"] = o.MongodbAtlasProjectId
	}
	if o.MongodbCustomData != nil {
		toSerialize["mongodb-custom-data"] = o.MongodbCustomData
	}
	if o.MongodbDefaultAuthDb != nil {
		toSerialize["mongodb-default-auth-db"] = o.MongodbDefaultAuthDb
	}
	if o.MongodbHostPort != nil {
		toSerialize["mongodb-host-port"] = o.MongodbHostPort
	}
	if o.MongodbName != nil {
		toSerialize["mongodb-name"] = o.MongodbName
	}
	if o.MongodbPassword != nil {
		toSerialize["mongodb-password"] = o.MongodbPassword
	}
	if o.MongodbRoles != nil {
		toSerialize["mongodb-roles"] = o.MongodbRoles
	}
	if o.MongodbServerUri != nil {
		toSerialize["mongodb-server-uri"] = o.MongodbServerUri
	}
	if o.MongodbUriOptions != nil {
		toSerialize["mongodb-uri-options"] = o.MongodbUriOptions
	}
	if o.MongodbUsername != nil {
		toSerialize["mongodb-username"] = o.MongodbUsername
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProducerEncryptionKeyName != nil {
		toSerialize["producer-encryption-key-name"] = o.ProducerEncryptionKeyName
	}
	if o.SecureAccessBastionIssuer != nil {
		toSerialize["secure-access-bastion-issuer"] = o.SecureAccessBastionIssuer
	}
	if o.SecureAccessEnable != nil {
		toSerialize["secure-access-enable"] = o.SecureAccessEnable
	}
	if o.SecureAccessHost != nil {
		toSerialize["secure-access-host"] = o.SecureAccessHost
	}
	if o.SecureAccessWeb != nil {
		toSerialize["secure-access-web"] = o.SecureAccessWeb
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


