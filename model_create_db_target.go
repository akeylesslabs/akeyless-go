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

// CreateDBTarget struct for CreateDBTarget
type CreateDBTarget struct {
	DBDefinedConnectionType *string `json:"DBDefinedConnectionType,omitempty"`
	// (Optional) Client id (relevant for \"cloud-service-provider\" only)
	AzureClientId *string `json:"azure-client-id,omitempty"`
	// (Optional) Client secret (relevant for \"cloud-service-provider\" only)
	AzureClientSecret *string `json:"azure-client-secret,omitempty"`
	// (Optional) Tenant id (relevant for \"cloud-service-provider\" only)
	AzureTenantId *string `json:"azure-tenant-id,omitempty"`
	// (Optional) Cloud service provider (currently only supports Azure)
	CloudServiceProvider *string `json:"cloud-service-provider,omitempty"`
	// Cluster Mode
	ClusterMode *bool `json:"cluster-mode,omitempty"`
	// Deprecated - use description
	Comment *string `json:"comment,omitempty"`
	// (Optional) Type of connection to mssql database [credentials/cloud-identity]
	ConnectionType string `json:"connection-type"`
	DbName *string `json:"db-name,omitempty"`
	// (Optional) DB server certificates
	DbServerCertificates *string `json:"db-server-certificates,omitempty"`
	// (Optional) Server name for certificate verification
	DbServerName *string `json:"db-server-name,omitempty"`
	DbType string `json:"db-type"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	Host *string `json:"host,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Set the maximum number of versions, limited by the account settings defaults.
	MaxVersions *string `json:"max-versions,omitempty"`
	MongodbAtlas *bool `json:"mongodb-atlas,omitempty"`
	// MongoDB Atlas private key
	MongodbAtlasApiPrivateKey *string `json:"mongodb-atlas-api-private-key,omitempty"`
	// MongoDB Atlas public key
	MongodbAtlasApiPublicKey *string `json:"mongodb-atlas-api-public-key,omitempty"`
	// MongoDB Atlas project ID
	MongodbAtlasProjectId *string `json:"mongodb-atlas-project-id,omitempty"`
	// MongoDB server default authentication database
	MongodbDefaultAuthDb *string `json:"mongodb-default-auth-db,omitempty"`
	// MongoDB server URI options
	MongodbUriOptions *string `json:"mongodb-uri-options,omitempty"`
	// Target name
	Name string `json:"name"`
	OracleServiceName *string `json:"oracle-service-name,omitempty"`
	Port *string `json:"port,omitempty"`
	Pwd *string `json:"pwd,omitempty"`
	SnowflakeAccount *string `json:"snowflake-account,omitempty"`
	// RSA Private key (base64 encoded)
	SnowflakeApiPrivateKey *string `json:"snowflake-api-private-key,omitempty"`
	// The Private key passphrase
	SnowflakeApiPrivateKeyPassword *string `json:"snowflake-api-private-key-password,omitempty"`
	// Enable/Disable SSL [true/false]
	Ssl *bool `json:"ssl,omitempty"`
	// SSL connection certificate
	SslCertificate *string `json:"ssl-certificate,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	UserName *string `json:"user-name,omitempty"`
}

// NewCreateDBTarget instantiates a new CreateDBTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDBTarget(connectionType string, dbType string, name string, ) *CreateDBTarget {
	this := CreateDBTarget{}
	this.ConnectionType = connectionType
	this.DbType = dbType
	var json bool = false
	this.Json = &json
	this.Name = name
	var ssl bool = false
	this.Ssl = &ssl
	return &this
}

// NewCreateDBTargetWithDefaults instantiates a new CreateDBTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDBTargetWithDefaults() *CreateDBTarget {
	this := CreateDBTarget{}
	var connectionType string = "credentials"
	this.ConnectionType = connectionType
	var json bool = false
	this.Json = &json
	var ssl bool = false
	this.Ssl = &ssl
	return &this
}

// GetDBDefinedConnectionType returns the DBDefinedConnectionType field value if set, zero value otherwise.
func (o *CreateDBTarget) GetDBDefinedConnectionType() string {
	if o == nil || o.DBDefinedConnectionType == nil {
		var ret string
		return ret
	}
	return *o.DBDefinedConnectionType
}

// GetDBDefinedConnectionTypeOk returns a tuple with the DBDefinedConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetDBDefinedConnectionTypeOk() (*string, bool) {
	if o == nil || o.DBDefinedConnectionType == nil {
		return nil, false
	}
	return o.DBDefinedConnectionType, true
}

// HasDBDefinedConnectionType returns a boolean if a field has been set.
func (o *CreateDBTarget) HasDBDefinedConnectionType() bool {
	if o != nil && o.DBDefinedConnectionType != nil {
		return true
	}

	return false
}

// SetDBDefinedConnectionType gets a reference to the given string and assigns it to the DBDefinedConnectionType field.
func (o *CreateDBTarget) SetDBDefinedConnectionType(v string) {
	o.DBDefinedConnectionType = &v
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise.
func (o *CreateDBTarget) GetAzureClientId() string {
	if o == nil || o.AzureClientId == nil {
		var ret string
		return ret
	}
	return *o.AzureClientId
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetAzureClientIdOk() (*string, bool) {
	if o == nil || o.AzureClientId == nil {
		return nil, false
	}
	return o.AzureClientId, true
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *CreateDBTarget) HasAzureClientId() bool {
	if o != nil && o.AzureClientId != nil {
		return true
	}

	return false
}

// SetAzureClientId gets a reference to the given string and assigns it to the AzureClientId field.
func (o *CreateDBTarget) SetAzureClientId(v string) {
	o.AzureClientId = &v
}

// GetAzureClientSecret returns the AzureClientSecret field value if set, zero value otherwise.
func (o *CreateDBTarget) GetAzureClientSecret() string {
	if o == nil || o.AzureClientSecret == nil {
		var ret string
		return ret
	}
	return *o.AzureClientSecret
}

// GetAzureClientSecretOk returns a tuple with the AzureClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetAzureClientSecretOk() (*string, bool) {
	if o == nil || o.AzureClientSecret == nil {
		return nil, false
	}
	return o.AzureClientSecret, true
}

// HasAzureClientSecret returns a boolean if a field has been set.
func (o *CreateDBTarget) HasAzureClientSecret() bool {
	if o != nil && o.AzureClientSecret != nil {
		return true
	}

	return false
}

// SetAzureClientSecret gets a reference to the given string and assigns it to the AzureClientSecret field.
func (o *CreateDBTarget) SetAzureClientSecret(v string) {
	o.AzureClientSecret = &v
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *CreateDBTarget) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || o.AzureTenantId == nil {
		return nil, false
	}
	return o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *CreateDBTarget) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId != nil {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *CreateDBTarget) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetCloudServiceProvider returns the CloudServiceProvider field value if set, zero value otherwise.
func (o *CreateDBTarget) GetCloudServiceProvider() string {
	if o == nil || o.CloudServiceProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudServiceProvider
}

// GetCloudServiceProviderOk returns a tuple with the CloudServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetCloudServiceProviderOk() (*string, bool) {
	if o == nil || o.CloudServiceProvider == nil {
		return nil, false
	}
	return o.CloudServiceProvider, true
}

// HasCloudServiceProvider returns a boolean if a field has been set.
func (o *CreateDBTarget) HasCloudServiceProvider() bool {
	if o != nil && o.CloudServiceProvider != nil {
		return true
	}

	return false
}

// SetCloudServiceProvider gets a reference to the given string and assigns it to the CloudServiceProvider field.
func (o *CreateDBTarget) SetCloudServiceProvider(v string) {
	o.CloudServiceProvider = &v
}

// GetClusterMode returns the ClusterMode field value if set, zero value otherwise.
func (o *CreateDBTarget) GetClusterMode() bool {
	if o == nil || o.ClusterMode == nil {
		var ret bool
		return ret
	}
	return *o.ClusterMode
}

// GetClusterModeOk returns a tuple with the ClusterMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetClusterModeOk() (*bool, bool) {
	if o == nil || o.ClusterMode == nil {
		return nil, false
	}
	return o.ClusterMode, true
}

// HasClusterMode returns a boolean if a field has been set.
func (o *CreateDBTarget) HasClusterMode() bool {
	if o != nil && o.ClusterMode != nil {
		return true
	}

	return false
}

// SetClusterMode gets a reference to the given bool and assigns it to the ClusterMode field.
func (o *CreateDBTarget) SetClusterMode(v bool) {
	o.ClusterMode = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateDBTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateDBTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateDBTarget) SetComment(v string) {
	o.Comment = &v
}

// GetConnectionType returns the ConnectionType field value
func (o *CreateDBTarget) GetConnectionType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetConnectionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *CreateDBTarget) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *CreateDBTarget) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *CreateDBTarget) HasDbName() bool {
	if o != nil && o.DbName != nil {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *CreateDBTarget) SetDbName(v string) {
	o.DbName = &v
}

// GetDbServerCertificates returns the DbServerCertificates field value if set, zero value otherwise.
func (o *CreateDBTarget) GetDbServerCertificates() string {
	if o == nil || o.DbServerCertificates == nil {
		var ret string
		return ret
	}
	return *o.DbServerCertificates
}

// GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetDbServerCertificatesOk() (*string, bool) {
	if o == nil || o.DbServerCertificates == nil {
		return nil, false
	}
	return o.DbServerCertificates, true
}

// HasDbServerCertificates returns a boolean if a field has been set.
func (o *CreateDBTarget) HasDbServerCertificates() bool {
	if o != nil && o.DbServerCertificates != nil {
		return true
	}

	return false
}

// SetDbServerCertificates gets a reference to the given string and assigns it to the DbServerCertificates field.
func (o *CreateDBTarget) SetDbServerCertificates(v string) {
	o.DbServerCertificates = &v
}

// GetDbServerName returns the DbServerName field value if set, zero value otherwise.
func (o *CreateDBTarget) GetDbServerName() string {
	if o == nil || o.DbServerName == nil {
		var ret string
		return ret
	}
	return *o.DbServerName
}

// GetDbServerNameOk returns a tuple with the DbServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetDbServerNameOk() (*string, bool) {
	if o == nil || o.DbServerName == nil {
		return nil, false
	}
	return o.DbServerName, true
}

// HasDbServerName returns a boolean if a field has been set.
func (o *CreateDBTarget) HasDbServerName() bool {
	if o != nil && o.DbServerName != nil {
		return true
	}

	return false
}

// SetDbServerName gets a reference to the given string and assigns it to the DbServerName field.
func (o *CreateDBTarget) SetDbServerName(v string) {
	o.DbServerName = &v
}

// GetDbType returns the DbType field value
func (o *CreateDBTarget) GetDbType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbType
}

// GetDbTypeOk returns a tuple with the DbType field value
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetDbTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbType, true
}

// SetDbType sets field value
func (o *CreateDBTarget) SetDbType(v string) {
	o.DbType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDBTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDBTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDBTarget) SetDescription(v string) {
	o.Description = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CreateDBTarget) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CreateDBTarget) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CreateDBTarget) SetHost(v string) {
	o.Host = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *CreateDBTarget) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *CreateDBTarget) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *CreateDBTarget) SetJson(v bool) {
	o.Json = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateDBTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateDBTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateDBTarget) SetKey(v string) {
	o.Key = &v
}

// GetMaxVersions returns the MaxVersions field value if set, zero value otherwise.
func (o *CreateDBTarget) GetMaxVersions() string {
	if o == nil || o.MaxVersions == nil {
		var ret string
		return ret
	}
	return *o.MaxVersions
}

// GetMaxVersionsOk returns a tuple with the MaxVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetMaxVersionsOk() (*string, bool) {
	if o == nil || o.MaxVersions == nil {
		return nil, false
	}
	return o.MaxVersions, true
}

// HasMaxVersions returns a boolean if a field has been set.
func (o *CreateDBTarget) HasMaxVersions() bool {
	if o != nil && o.MaxVersions != nil {
		return true
	}

	return false
}

// SetMaxVersions gets a reference to the given string and assigns it to the MaxVersions field.
func (o *CreateDBTarget) SetMaxVersions(v string) {
	o.MaxVersions = &v
}

// GetMongodbAtlas returns the MongodbAtlas field value if set, zero value otherwise.
func (o *CreateDBTarget) GetMongodbAtlas() bool {
	if o == nil || o.MongodbAtlas == nil {
		var ret bool
		return ret
	}
	return *o.MongodbAtlas
}

// GetMongodbAtlasOk returns a tuple with the MongodbAtlas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetMongodbAtlasOk() (*bool, bool) {
	if o == nil || o.MongodbAtlas == nil {
		return nil, false
	}
	return o.MongodbAtlas, true
}

// HasMongodbAtlas returns a boolean if a field has been set.
func (o *CreateDBTarget) HasMongodbAtlas() bool {
	if o != nil && o.MongodbAtlas != nil {
		return true
	}

	return false
}

// SetMongodbAtlas gets a reference to the given bool and assigns it to the MongodbAtlas field.
func (o *CreateDBTarget) SetMongodbAtlas(v bool) {
	o.MongodbAtlas = &v
}

// GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field value if set, zero value otherwise.
func (o *CreateDBTarget) GetMongodbAtlasApiPrivateKey() string {
	if o == nil || o.MongodbAtlasApiPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.MongodbAtlasApiPrivateKey
}

// GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetMongodbAtlasApiPrivateKeyOk() (*string, bool) {
	if o == nil || o.MongodbAtlasApiPrivateKey == nil {
		return nil, false
	}
	return o.MongodbAtlasApiPrivateKey, true
}

// HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.
func (o *CreateDBTarget) HasMongodbAtlasApiPrivateKey() bool {
	if o != nil && o.MongodbAtlasApiPrivateKey != nil {
		return true
	}

	return false
}

// SetMongodbAtlasApiPrivateKey gets a reference to the given string and assigns it to the MongodbAtlasApiPrivateKey field.
func (o *CreateDBTarget) SetMongodbAtlasApiPrivateKey(v string) {
	o.MongodbAtlasApiPrivateKey = &v
}

// GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field value if set, zero value otherwise.
func (o *CreateDBTarget) GetMongodbAtlasApiPublicKey() string {
	if o == nil || o.MongodbAtlasApiPublicKey == nil {
		var ret string
		return ret
	}
	return *o.MongodbAtlasApiPublicKey
}

// GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetMongodbAtlasApiPublicKeyOk() (*string, bool) {
	if o == nil || o.MongodbAtlasApiPublicKey == nil {
		return nil, false
	}
	return o.MongodbAtlasApiPublicKey, true
}

// HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.
func (o *CreateDBTarget) HasMongodbAtlasApiPublicKey() bool {
	if o != nil && o.MongodbAtlasApiPublicKey != nil {
		return true
	}

	return false
}

// SetMongodbAtlasApiPublicKey gets a reference to the given string and assigns it to the MongodbAtlasApiPublicKey field.
func (o *CreateDBTarget) SetMongodbAtlasApiPublicKey(v string) {
	o.MongodbAtlasApiPublicKey = &v
}

// GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field value if set, zero value otherwise.
func (o *CreateDBTarget) GetMongodbAtlasProjectId() string {
	if o == nil || o.MongodbAtlasProjectId == nil {
		var ret string
		return ret
	}
	return *o.MongodbAtlasProjectId
}

// GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetMongodbAtlasProjectIdOk() (*string, bool) {
	if o == nil || o.MongodbAtlasProjectId == nil {
		return nil, false
	}
	return o.MongodbAtlasProjectId, true
}

// HasMongodbAtlasProjectId returns a boolean if a field has been set.
func (o *CreateDBTarget) HasMongodbAtlasProjectId() bool {
	if o != nil && o.MongodbAtlasProjectId != nil {
		return true
	}

	return false
}

// SetMongodbAtlasProjectId gets a reference to the given string and assigns it to the MongodbAtlasProjectId field.
func (o *CreateDBTarget) SetMongodbAtlasProjectId(v string) {
	o.MongodbAtlasProjectId = &v
}

// GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field value if set, zero value otherwise.
func (o *CreateDBTarget) GetMongodbDefaultAuthDb() string {
	if o == nil || o.MongodbDefaultAuthDb == nil {
		var ret string
		return ret
	}
	return *o.MongodbDefaultAuthDb
}

// GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetMongodbDefaultAuthDbOk() (*string, bool) {
	if o == nil || o.MongodbDefaultAuthDb == nil {
		return nil, false
	}
	return o.MongodbDefaultAuthDb, true
}

// HasMongodbDefaultAuthDb returns a boolean if a field has been set.
func (o *CreateDBTarget) HasMongodbDefaultAuthDb() bool {
	if o != nil && o.MongodbDefaultAuthDb != nil {
		return true
	}

	return false
}

// SetMongodbDefaultAuthDb gets a reference to the given string and assigns it to the MongodbDefaultAuthDb field.
func (o *CreateDBTarget) SetMongodbDefaultAuthDb(v string) {
	o.MongodbDefaultAuthDb = &v
}

// GetMongodbUriOptions returns the MongodbUriOptions field value if set, zero value otherwise.
func (o *CreateDBTarget) GetMongodbUriOptions() string {
	if o == nil || o.MongodbUriOptions == nil {
		var ret string
		return ret
	}
	return *o.MongodbUriOptions
}

// GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetMongodbUriOptionsOk() (*string, bool) {
	if o == nil || o.MongodbUriOptions == nil {
		return nil, false
	}
	return o.MongodbUriOptions, true
}

// HasMongodbUriOptions returns a boolean if a field has been set.
func (o *CreateDBTarget) HasMongodbUriOptions() bool {
	if o != nil && o.MongodbUriOptions != nil {
		return true
	}

	return false
}

// SetMongodbUriOptions gets a reference to the given string and assigns it to the MongodbUriOptions field.
func (o *CreateDBTarget) SetMongodbUriOptions(v string) {
	o.MongodbUriOptions = &v
}

// GetName returns the Name field value
func (o *CreateDBTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDBTarget) SetName(v string) {
	o.Name = v
}

// GetOracleServiceName returns the OracleServiceName field value if set, zero value otherwise.
func (o *CreateDBTarget) GetOracleServiceName() string {
	if o == nil || o.OracleServiceName == nil {
		var ret string
		return ret
	}
	return *o.OracleServiceName
}

// GetOracleServiceNameOk returns a tuple with the OracleServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetOracleServiceNameOk() (*string, bool) {
	if o == nil || o.OracleServiceName == nil {
		return nil, false
	}
	return o.OracleServiceName, true
}

// HasOracleServiceName returns a boolean if a field has been set.
func (o *CreateDBTarget) HasOracleServiceName() bool {
	if o != nil && o.OracleServiceName != nil {
		return true
	}

	return false
}

// SetOracleServiceName gets a reference to the given string and assigns it to the OracleServiceName field.
func (o *CreateDBTarget) SetOracleServiceName(v string) {
	o.OracleServiceName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CreateDBTarget) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CreateDBTarget) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *CreateDBTarget) SetPort(v string) {
	o.Port = &v
}

// GetPwd returns the Pwd field value if set, zero value otherwise.
func (o *CreateDBTarget) GetPwd() string {
	if o == nil || o.Pwd == nil {
		var ret string
		return ret
	}
	return *o.Pwd
}

// GetPwdOk returns a tuple with the Pwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetPwdOk() (*string, bool) {
	if o == nil || o.Pwd == nil {
		return nil, false
	}
	return o.Pwd, true
}

// HasPwd returns a boolean if a field has been set.
func (o *CreateDBTarget) HasPwd() bool {
	if o != nil && o.Pwd != nil {
		return true
	}

	return false
}

// SetPwd gets a reference to the given string and assigns it to the Pwd field.
func (o *CreateDBTarget) SetPwd(v string) {
	o.Pwd = &v
}

// GetSnowflakeAccount returns the SnowflakeAccount field value if set, zero value otherwise.
func (o *CreateDBTarget) GetSnowflakeAccount() string {
	if o == nil || o.SnowflakeAccount == nil {
		var ret string
		return ret
	}
	return *o.SnowflakeAccount
}

// GetSnowflakeAccountOk returns a tuple with the SnowflakeAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetSnowflakeAccountOk() (*string, bool) {
	if o == nil || o.SnowflakeAccount == nil {
		return nil, false
	}
	return o.SnowflakeAccount, true
}

// HasSnowflakeAccount returns a boolean if a field has been set.
func (o *CreateDBTarget) HasSnowflakeAccount() bool {
	if o != nil && o.SnowflakeAccount != nil {
		return true
	}

	return false
}

// SetSnowflakeAccount gets a reference to the given string and assigns it to the SnowflakeAccount field.
func (o *CreateDBTarget) SetSnowflakeAccount(v string) {
	o.SnowflakeAccount = &v
}

// GetSnowflakeApiPrivateKey returns the SnowflakeApiPrivateKey field value if set, zero value otherwise.
func (o *CreateDBTarget) GetSnowflakeApiPrivateKey() string {
	if o == nil || o.SnowflakeApiPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.SnowflakeApiPrivateKey
}

// GetSnowflakeApiPrivateKeyOk returns a tuple with the SnowflakeApiPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetSnowflakeApiPrivateKeyOk() (*string, bool) {
	if o == nil || o.SnowflakeApiPrivateKey == nil {
		return nil, false
	}
	return o.SnowflakeApiPrivateKey, true
}

// HasSnowflakeApiPrivateKey returns a boolean if a field has been set.
func (o *CreateDBTarget) HasSnowflakeApiPrivateKey() bool {
	if o != nil && o.SnowflakeApiPrivateKey != nil {
		return true
	}

	return false
}

// SetSnowflakeApiPrivateKey gets a reference to the given string and assigns it to the SnowflakeApiPrivateKey field.
func (o *CreateDBTarget) SetSnowflakeApiPrivateKey(v string) {
	o.SnowflakeApiPrivateKey = &v
}

// GetSnowflakeApiPrivateKeyPassword returns the SnowflakeApiPrivateKeyPassword field value if set, zero value otherwise.
func (o *CreateDBTarget) GetSnowflakeApiPrivateKeyPassword() string {
	if o == nil || o.SnowflakeApiPrivateKeyPassword == nil {
		var ret string
		return ret
	}
	return *o.SnowflakeApiPrivateKeyPassword
}

// GetSnowflakeApiPrivateKeyPasswordOk returns a tuple with the SnowflakeApiPrivateKeyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetSnowflakeApiPrivateKeyPasswordOk() (*string, bool) {
	if o == nil || o.SnowflakeApiPrivateKeyPassword == nil {
		return nil, false
	}
	return o.SnowflakeApiPrivateKeyPassword, true
}

// HasSnowflakeApiPrivateKeyPassword returns a boolean if a field has been set.
func (o *CreateDBTarget) HasSnowflakeApiPrivateKeyPassword() bool {
	if o != nil && o.SnowflakeApiPrivateKeyPassword != nil {
		return true
	}

	return false
}

// SetSnowflakeApiPrivateKeyPassword gets a reference to the given string and assigns it to the SnowflakeApiPrivateKeyPassword field.
func (o *CreateDBTarget) SetSnowflakeApiPrivateKeyPassword(v string) {
	o.SnowflakeApiPrivateKeyPassword = &v
}

// GetSsl returns the Ssl field value if set, zero value otherwise.
func (o *CreateDBTarget) GetSsl() bool {
	if o == nil || o.Ssl == nil {
		var ret bool
		return ret
	}
	return *o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetSslOk() (*bool, bool) {
	if o == nil || o.Ssl == nil {
		return nil, false
	}
	return o.Ssl, true
}

// HasSsl returns a boolean if a field has been set.
func (o *CreateDBTarget) HasSsl() bool {
	if o != nil && o.Ssl != nil {
		return true
	}

	return false
}

// SetSsl gets a reference to the given bool and assigns it to the Ssl field.
func (o *CreateDBTarget) SetSsl(v bool) {
	o.Ssl = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *CreateDBTarget) GetSslCertificate() string {
	if o == nil || o.SslCertificate == nil {
		var ret string
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetSslCertificateOk() (*string, bool) {
	if o == nil || o.SslCertificate == nil {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *CreateDBTarget) HasSslCertificate() bool {
	if o != nil && o.SslCertificate != nil {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given string and assigns it to the SslCertificate field.
func (o *CreateDBTarget) SetSslCertificate(v string) {
	o.SslCertificate = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateDBTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateDBTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateDBTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *CreateDBTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *CreateDBTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *CreateDBTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *CreateDBTarget) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDBTarget) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *CreateDBTarget) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *CreateDBTarget) SetUserName(v string) {
	o.UserName = &v
}

func (o CreateDBTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DBDefinedConnectionType != nil {
		toSerialize["DBDefinedConnectionType"] = o.DBDefinedConnectionType
	}
	if o.AzureClientId != nil {
		toSerialize["azure-client-id"] = o.AzureClientId
	}
	if o.AzureClientSecret != nil {
		toSerialize["azure-client-secret"] = o.AzureClientSecret
	}
	if o.AzureTenantId != nil {
		toSerialize["azure-tenant-id"] = o.AzureTenantId
	}
	if o.CloudServiceProvider != nil {
		toSerialize["cloud-service-provider"] = o.CloudServiceProvider
	}
	if o.ClusterMode != nil {
		toSerialize["cluster-mode"] = o.ClusterMode
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["connection-type"] = o.ConnectionType
	}
	if o.DbName != nil {
		toSerialize["db-name"] = o.DbName
	}
	if o.DbServerCertificates != nil {
		toSerialize["db-server-certificates"] = o.DbServerCertificates
	}
	if o.DbServerName != nil {
		toSerialize["db-server-name"] = o.DbServerName
	}
	if true {
		toSerialize["db-type"] = o.DbType
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
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.MaxVersions != nil {
		toSerialize["max-versions"] = o.MaxVersions
	}
	if o.MongodbAtlas != nil {
		toSerialize["mongodb-atlas"] = o.MongodbAtlas
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
	if o.MongodbDefaultAuthDb != nil {
		toSerialize["mongodb-default-auth-db"] = o.MongodbDefaultAuthDb
	}
	if o.MongodbUriOptions != nil {
		toSerialize["mongodb-uri-options"] = o.MongodbUriOptions
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OracleServiceName != nil {
		toSerialize["oracle-service-name"] = o.OracleServiceName
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Pwd != nil {
		toSerialize["pwd"] = o.Pwd
	}
	if o.SnowflakeAccount != nil {
		toSerialize["snowflake-account"] = o.SnowflakeAccount
	}
	if o.SnowflakeApiPrivateKey != nil {
		toSerialize["snowflake-api-private-key"] = o.SnowflakeApiPrivateKey
	}
	if o.SnowflakeApiPrivateKeyPassword != nil {
		toSerialize["snowflake-api-private-key-password"] = o.SnowflakeApiPrivateKeyPassword
	}
	if o.Ssl != nil {
		toSerialize["ssl"] = o.Ssl
	}
	if o.SslCertificate != nil {
		toSerialize["ssl-certificate"] = o.SslCertificate
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UserName != nil {
		toSerialize["user-name"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDBTarget struct {
	value *CreateDBTarget
	isSet bool
}

func (v NullableCreateDBTarget) Get() *CreateDBTarget {
	return v.value
}

func (v *NullableCreateDBTarget) Set(val *CreateDBTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDBTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDBTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDBTarget(val *CreateDBTarget) *NullableCreateDBTarget {
	return &NullableCreateDBTarget{value: val, isSet: true}
}

func (v NullableCreateDBTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDBTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


