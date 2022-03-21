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

// UpdateDBTarget struct for UpdateDBTarget
type UpdateDBTarget struct {
	// Comment about the target
	Comment *string `json:"comment,omitempty"`
	DbName *string `json:"db-name,omitempty"`
	// (Optional) DB server certificates
	DbServerCertificates *string `json:"db-server-certificates,omitempty"`
	// (Optional) Server name for certificate verification
	DbServerName *string `json:"db-server-name,omitempty"`
	DbType string `json:"db-type"`
	Host *string `json:"host,omitempty"`
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
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
	// New target name
	NewName *string `json:"new-name,omitempty"`
	OracleServiceName *string `json:"oracle-service-name,omitempty"`
	Port *string `json:"port,omitempty"`
	Pwd *string `json:"pwd,omitempty"`
	SnowflakeAccount *string `json:"snowflake-account,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
	// Deprecated
	UpdateVersion *bool `json:"update-version,omitempty"`
	UserName *string `json:"user-name,omitempty"`
}

// NewUpdateDBTarget instantiates a new UpdateDBTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDBTarget(dbType string, name string, ) *UpdateDBTarget {
	this := UpdateDBTarget{}
	this.DbType = dbType
	this.Name = name
	return &this
}

// NewUpdateDBTargetWithDefaults instantiates a new UpdateDBTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDBTargetWithDefaults() *UpdateDBTarget {
	this := UpdateDBTarget{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateDBTarget) SetComment(v string) {
	o.Comment = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasDbName() bool {
	if o != nil && o.DbName != nil {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *UpdateDBTarget) SetDbName(v string) {
	o.DbName = &v
}

// GetDbServerCertificates returns the DbServerCertificates field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetDbServerCertificates() string {
	if o == nil || o.DbServerCertificates == nil {
		var ret string
		return ret
	}
	return *o.DbServerCertificates
}

// GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetDbServerCertificatesOk() (*string, bool) {
	if o == nil || o.DbServerCertificates == nil {
		return nil, false
	}
	return o.DbServerCertificates, true
}

// HasDbServerCertificates returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasDbServerCertificates() bool {
	if o != nil && o.DbServerCertificates != nil {
		return true
	}

	return false
}

// SetDbServerCertificates gets a reference to the given string and assigns it to the DbServerCertificates field.
func (o *UpdateDBTarget) SetDbServerCertificates(v string) {
	o.DbServerCertificates = &v
}

// GetDbServerName returns the DbServerName field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetDbServerName() string {
	if o == nil || o.DbServerName == nil {
		var ret string
		return ret
	}
	return *o.DbServerName
}

// GetDbServerNameOk returns a tuple with the DbServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetDbServerNameOk() (*string, bool) {
	if o == nil || o.DbServerName == nil {
		return nil, false
	}
	return o.DbServerName, true
}

// HasDbServerName returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasDbServerName() bool {
	if o != nil && o.DbServerName != nil {
		return true
	}

	return false
}

// SetDbServerName gets a reference to the given string and assigns it to the DbServerName field.
func (o *UpdateDBTarget) SetDbServerName(v string) {
	o.DbServerName = &v
}

// GetDbType returns the DbType field value
func (o *UpdateDBTarget) GetDbType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbType
}

// GetDbTypeOk returns a tuple with the DbType field value
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetDbTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbType, true
}

// SetDbType sets field value
func (o *UpdateDBTarget) SetDbType(v string) {
	o.DbType = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UpdateDBTarget) SetHost(v string) {
	o.Host = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *UpdateDBTarget) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateDBTarget) SetKey(v string) {
	o.Key = &v
}

// GetMongodbAtlas returns the MongodbAtlas field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetMongodbAtlas() bool {
	if o == nil || o.MongodbAtlas == nil {
		var ret bool
		return ret
	}
	return *o.MongodbAtlas
}

// GetMongodbAtlasOk returns a tuple with the MongodbAtlas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetMongodbAtlasOk() (*bool, bool) {
	if o == nil || o.MongodbAtlas == nil {
		return nil, false
	}
	return o.MongodbAtlas, true
}

// HasMongodbAtlas returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasMongodbAtlas() bool {
	if o != nil && o.MongodbAtlas != nil {
		return true
	}

	return false
}

// SetMongodbAtlas gets a reference to the given bool and assigns it to the MongodbAtlas field.
func (o *UpdateDBTarget) SetMongodbAtlas(v bool) {
	o.MongodbAtlas = &v
}

// GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetMongodbAtlasApiPrivateKey() string {
	if o == nil || o.MongodbAtlasApiPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.MongodbAtlasApiPrivateKey
}

// GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetMongodbAtlasApiPrivateKeyOk() (*string, bool) {
	if o == nil || o.MongodbAtlasApiPrivateKey == nil {
		return nil, false
	}
	return o.MongodbAtlasApiPrivateKey, true
}

// HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasMongodbAtlasApiPrivateKey() bool {
	if o != nil && o.MongodbAtlasApiPrivateKey != nil {
		return true
	}

	return false
}

// SetMongodbAtlasApiPrivateKey gets a reference to the given string and assigns it to the MongodbAtlasApiPrivateKey field.
func (o *UpdateDBTarget) SetMongodbAtlasApiPrivateKey(v string) {
	o.MongodbAtlasApiPrivateKey = &v
}

// GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetMongodbAtlasApiPublicKey() string {
	if o == nil || o.MongodbAtlasApiPublicKey == nil {
		var ret string
		return ret
	}
	return *o.MongodbAtlasApiPublicKey
}

// GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetMongodbAtlasApiPublicKeyOk() (*string, bool) {
	if o == nil || o.MongodbAtlasApiPublicKey == nil {
		return nil, false
	}
	return o.MongodbAtlasApiPublicKey, true
}

// HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasMongodbAtlasApiPublicKey() bool {
	if o != nil && o.MongodbAtlasApiPublicKey != nil {
		return true
	}

	return false
}

// SetMongodbAtlasApiPublicKey gets a reference to the given string and assigns it to the MongodbAtlasApiPublicKey field.
func (o *UpdateDBTarget) SetMongodbAtlasApiPublicKey(v string) {
	o.MongodbAtlasApiPublicKey = &v
}

// GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetMongodbAtlasProjectId() string {
	if o == nil || o.MongodbAtlasProjectId == nil {
		var ret string
		return ret
	}
	return *o.MongodbAtlasProjectId
}

// GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetMongodbAtlasProjectIdOk() (*string, bool) {
	if o == nil || o.MongodbAtlasProjectId == nil {
		return nil, false
	}
	return o.MongodbAtlasProjectId, true
}

// HasMongodbAtlasProjectId returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasMongodbAtlasProjectId() bool {
	if o != nil && o.MongodbAtlasProjectId != nil {
		return true
	}

	return false
}

// SetMongodbAtlasProjectId gets a reference to the given string and assigns it to the MongodbAtlasProjectId field.
func (o *UpdateDBTarget) SetMongodbAtlasProjectId(v string) {
	o.MongodbAtlasProjectId = &v
}

// GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetMongodbDefaultAuthDb() string {
	if o == nil || o.MongodbDefaultAuthDb == nil {
		var ret string
		return ret
	}
	return *o.MongodbDefaultAuthDb
}

// GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetMongodbDefaultAuthDbOk() (*string, bool) {
	if o == nil || o.MongodbDefaultAuthDb == nil {
		return nil, false
	}
	return o.MongodbDefaultAuthDb, true
}

// HasMongodbDefaultAuthDb returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasMongodbDefaultAuthDb() bool {
	if o != nil && o.MongodbDefaultAuthDb != nil {
		return true
	}

	return false
}

// SetMongodbDefaultAuthDb gets a reference to the given string and assigns it to the MongodbDefaultAuthDb field.
func (o *UpdateDBTarget) SetMongodbDefaultAuthDb(v string) {
	o.MongodbDefaultAuthDb = &v
}

// GetMongodbUriOptions returns the MongodbUriOptions field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetMongodbUriOptions() string {
	if o == nil || o.MongodbUriOptions == nil {
		var ret string
		return ret
	}
	return *o.MongodbUriOptions
}

// GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetMongodbUriOptionsOk() (*string, bool) {
	if o == nil || o.MongodbUriOptions == nil {
		return nil, false
	}
	return o.MongodbUriOptions, true
}

// HasMongodbUriOptions returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasMongodbUriOptions() bool {
	if o != nil && o.MongodbUriOptions != nil {
		return true
	}

	return false
}

// SetMongodbUriOptions gets a reference to the given string and assigns it to the MongodbUriOptions field.
func (o *UpdateDBTarget) SetMongodbUriOptions(v string) {
	o.MongodbUriOptions = &v
}

// GetName returns the Name field value
func (o *UpdateDBTarget) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateDBTarget) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateDBTarget) SetNewName(v string) {
	o.NewName = &v
}

// GetOracleServiceName returns the OracleServiceName field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetOracleServiceName() string {
	if o == nil || o.OracleServiceName == nil {
		var ret string
		return ret
	}
	return *o.OracleServiceName
}

// GetOracleServiceNameOk returns a tuple with the OracleServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetOracleServiceNameOk() (*string, bool) {
	if o == nil || o.OracleServiceName == nil {
		return nil, false
	}
	return o.OracleServiceName, true
}

// HasOracleServiceName returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasOracleServiceName() bool {
	if o != nil && o.OracleServiceName != nil {
		return true
	}

	return false
}

// SetOracleServiceName gets a reference to the given string and assigns it to the OracleServiceName field.
func (o *UpdateDBTarget) SetOracleServiceName(v string) {
	o.OracleServiceName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *UpdateDBTarget) SetPort(v string) {
	o.Port = &v
}

// GetPwd returns the Pwd field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetPwd() string {
	if o == nil || o.Pwd == nil {
		var ret string
		return ret
	}
	return *o.Pwd
}

// GetPwdOk returns a tuple with the Pwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetPwdOk() (*string, bool) {
	if o == nil || o.Pwd == nil {
		return nil, false
	}
	return o.Pwd, true
}

// HasPwd returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasPwd() bool {
	if o != nil && o.Pwd != nil {
		return true
	}

	return false
}

// SetPwd gets a reference to the given string and assigns it to the Pwd field.
func (o *UpdateDBTarget) SetPwd(v string) {
	o.Pwd = &v
}

// GetSnowflakeAccount returns the SnowflakeAccount field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetSnowflakeAccount() string {
	if o == nil || o.SnowflakeAccount == nil {
		var ret string
		return ret
	}
	return *o.SnowflakeAccount
}

// GetSnowflakeAccountOk returns a tuple with the SnowflakeAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetSnowflakeAccountOk() (*string, bool) {
	if o == nil || o.SnowflakeAccount == nil {
		return nil, false
	}
	return o.SnowflakeAccount, true
}

// HasSnowflakeAccount returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasSnowflakeAccount() bool {
	if o != nil && o.SnowflakeAccount != nil {
		return true
	}

	return false
}

// SetSnowflakeAccount gets a reference to the given string and assigns it to the SnowflakeAccount field.
func (o *UpdateDBTarget) SetSnowflakeAccount(v string) {
	o.SnowflakeAccount = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdateDBTarget) SetToken(v string) {
	o.Token = &v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *UpdateDBTarget) SetUidToken(v string) {
	o.UidToken = &v
}

// GetUpdateVersion returns the UpdateVersion field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetUpdateVersion() bool {
	if o == nil || o.UpdateVersion == nil {
		var ret bool
		return ret
	}
	return *o.UpdateVersion
}

// GetUpdateVersionOk returns a tuple with the UpdateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetUpdateVersionOk() (*bool, bool) {
	if o == nil || o.UpdateVersion == nil {
		return nil, false
	}
	return o.UpdateVersion, true
}

// HasUpdateVersion returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasUpdateVersion() bool {
	if o != nil && o.UpdateVersion != nil {
		return true
	}

	return false
}

// SetUpdateVersion gets a reference to the given bool and assigns it to the UpdateVersion field.
func (o *UpdateDBTarget) SetUpdateVersion(v bool) {
	o.UpdateVersion = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UpdateDBTarget) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDBTarget) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UpdateDBTarget) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UpdateDBTarget) SetUserName(v string) {
	o.UserName = &v
}

func (o UpdateDBTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
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
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
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
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
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
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	if o.UpdateVersion != nil {
		toSerialize["update-version"] = o.UpdateVersion
	}
	if o.UserName != nil {
		toSerialize["user-name"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDBTarget struct {
	value *UpdateDBTarget
	isSet bool
}

func (v NullableUpdateDBTarget) Get() *UpdateDBTarget {
	return v.value
}

func (v *NullableUpdateDBTarget) Set(val *UpdateDBTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDBTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDBTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDBTarget(val *UpdateDBTarget) *NullableUpdateDBTarget {
	return &NullableUpdateDBTarget{value: val, isSet: true}
}

func (v NullableUpdateDBTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDBTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


