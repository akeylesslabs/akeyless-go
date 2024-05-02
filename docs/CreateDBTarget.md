# CreateDBTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DBDefinedConnectionType** | Pointer to **string** |  | [optional] 
**AzureClientId** | Pointer to **string** | (Optional) Client id (relevant for \&quot;cloud-service-provider\&quot; only) | [optional] 
**AzureClientSecret** | Pointer to **string** | (Optional) Client secret (relevant for \&quot;cloud-service-provider\&quot; only) | [optional] 
**AzureTenantId** | Pointer to **string** | (Optional) Tenant id (relevant for \&quot;cloud-service-provider\&quot; only) | [optional] 
**CloudServiceProvider** | Pointer to **string** | (Optional) Cloud service provider (currently only supports Azure) | [optional] 
**ClusterMode** | Pointer to **bool** | Cluster Mode | [optional] 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**ConnectionType** | **string** | (Optional) Type of connection to mssql database [credentials/cloud-identity] | [default to "credentials"]
**DbName** | Pointer to **string** |  | [optional] 
**DbServerCertificates** | Pointer to **string** | (Optional) DB server certificates | [optional] 
**DbServerName** | Pointer to **string** | (Optional) Server name for certificate verification | [optional] 
**DbType** | **string** |  | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**MongodbAtlas** | Pointer to **bool** |  | [optional] 
**MongodbAtlasApiPrivateKey** | Pointer to **string** | MongoDB Atlas private key | [optional] 
**MongodbAtlasApiPublicKey** | Pointer to **string** | MongoDB Atlas public key | [optional] 
**MongodbAtlasProjectId** | Pointer to **string** | MongoDB Atlas project ID | [optional] 
**MongodbDefaultAuthDb** | Pointer to **string** | MongoDB server default authentication database | [optional] 
**MongodbUriOptions** | Pointer to **string** | MongoDB server URI options | [optional] 
**Name** | **string** | Target name | 
**OracleServiceName** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Pwd** | Pointer to **string** |  | [optional] 
**SnowflakeAccount** | Pointer to **string** |  | [optional] 
**SnowflakeApiPrivateKey** | Pointer to **string** | RSA Private key (base64 encoded) | [optional] 
**SnowflakeApiPrivateKeyPassword** | Pointer to **string** | The Private key passphrase | [optional] 
**Ssl** | Pointer to **bool** | Enable/Disable SSL [true/false] | [optional] [default to false]
**SslCertificate** | Pointer to **string** | SSL connection certificate | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDBTarget

`func NewCreateDBTarget(connectionType string, dbType string, name string, ) *CreateDBTarget`

NewCreateDBTarget instantiates a new CreateDBTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDBTargetWithDefaults

`func NewCreateDBTargetWithDefaults() *CreateDBTarget`

NewCreateDBTargetWithDefaults instantiates a new CreateDBTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDBDefinedConnectionType

`func (o *CreateDBTarget) GetDBDefinedConnectionType() string`

GetDBDefinedConnectionType returns the DBDefinedConnectionType field if non-nil, zero value otherwise.

### GetDBDefinedConnectionTypeOk

`func (o *CreateDBTarget) GetDBDefinedConnectionTypeOk() (*string, bool)`

GetDBDefinedConnectionTypeOk returns a tuple with the DBDefinedConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDBDefinedConnectionType

`func (o *CreateDBTarget) SetDBDefinedConnectionType(v string)`

SetDBDefinedConnectionType sets DBDefinedConnectionType field to given value.

### HasDBDefinedConnectionType

`func (o *CreateDBTarget) HasDBDefinedConnectionType() bool`

HasDBDefinedConnectionType returns a boolean if a field has been set.

### GetAzureClientId

`func (o *CreateDBTarget) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *CreateDBTarget) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *CreateDBTarget) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *CreateDBTarget) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureClientSecret

`func (o *CreateDBTarget) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *CreateDBTarget) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *CreateDBTarget) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### HasAzureClientSecret

`func (o *CreateDBTarget) HasAzureClientSecret() bool`

HasAzureClientSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *CreateDBTarget) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *CreateDBTarget) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *CreateDBTarget) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *CreateDBTarget) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetCloudServiceProvider

`func (o *CreateDBTarget) GetCloudServiceProvider() string`

GetCloudServiceProvider returns the CloudServiceProvider field if non-nil, zero value otherwise.

### GetCloudServiceProviderOk

`func (o *CreateDBTarget) GetCloudServiceProviderOk() (*string, bool)`

GetCloudServiceProviderOk returns a tuple with the CloudServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudServiceProvider

`func (o *CreateDBTarget) SetCloudServiceProvider(v string)`

SetCloudServiceProvider sets CloudServiceProvider field to given value.

### HasCloudServiceProvider

`func (o *CreateDBTarget) HasCloudServiceProvider() bool`

HasCloudServiceProvider returns a boolean if a field has been set.

### GetClusterMode

`func (o *CreateDBTarget) GetClusterMode() bool`

GetClusterMode returns the ClusterMode field if non-nil, zero value otherwise.

### GetClusterModeOk

`func (o *CreateDBTarget) GetClusterModeOk() (*bool, bool)`

GetClusterModeOk returns a tuple with the ClusterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMode

`func (o *CreateDBTarget) SetClusterMode(v bool)`

SetClusterMode sets ClusterMode field to given value.

### HasClusterMode

`func (o *CreateDBTarget) HasClusterMode() bool`

HasClusterMode returns a boolean if a field has been set.

### GetComment

`func (o *CreateDBTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateDBTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateDBTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateDBTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConnectionType

`func (o *CreateDBTarget) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateDBTarget) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateDBTarget) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetDbName

`func (o *CreateDBTarget) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *CreateDBTarget) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *CreateDBTarget) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *CreateDBTarget) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbServerCertificates

`func (o *CreateDBTarget) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *CreateDBTarget) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *CreateDBTarget) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *CreateDBTarget) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *CreateDBTarget) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *CreateDBTarget) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *CreateDBTarget) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *CreateDBTarget) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDbType

`func (o *CreateDBTarget) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *CreateDBTarget) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *CreateDBTarget) SetDbType(v string)`

SetDbType sets DbType field to given value.


### GetDescription

`func (o *CreateDBTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDBTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDBTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDBTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *CreateDBTarget) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateDBTarget) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateDBTarget) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateDBTarget) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *CreateDBTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateDBTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateDBTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateDBTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *CreateDBTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateDBTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateDBTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateDBTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *CreateDBTarget) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *CreateDBTarget) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *CreateDBTarget) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *CreateDBTarget) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetMongodbAtlas

`func (o *CreateDBTarget) GetMongodbAtlas() bool`

GetMongodbAtlas returns the MongodbAtlas field if non-nil, zero value otherwise.

### GetMongodbAtlasOk

`func (o *CreateDBTarget) GetMongodbAtlasOk() (*bool, bool)`

GetMongodbAtlasOk returns a tuple with the MongodbAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlas

`func (o *CreateDBTarget) SetMongodbAtlas(v bool)`

SetMongodbAtlas sets MongodbAtlas field to given value.

### HasMongodbAtlas

`func (o *CreateDBTarget) HasMongodbAtlas() bool`

HasMongodbAtlas returns a boolean if a field has been set.

### GetMongodbAtlasApiPrivateKey

`func (o *CreateDBTarget) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *CreateDBTarget) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *CreateDBTarget) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *CreateDBTarget) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *CreateDBTarget) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *CreateDBTarget) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *CreateDBTarget) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *CreateDBTarget) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *CreateDBTarget) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *CreateDBTarget) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *CreateDBTarget) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *CreateDBTarget) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *CreateDBTarget) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *CreateDBTarget) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *CreateDBTarget) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *CreateDBTarget) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *CreateDBTarget) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *CreateDBTarget) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *CreateDBTarget) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *CreateDBTarget) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetName

`func (o *CreateDBTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDBTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDBTarget) SetName(v string)`

SetName sets Name field to given value.


### GetOracleServiceName

`func (o *CreateDBTarget) GetOracleServiceName() string`

GetOracleServiceName returns the OracleServiceName field if non-nil, zero value otherwise.

### GetOracleServiceNameOk

`func (o *CreateDBTarget) GetOracleServiceNameOk() (*string, bool)`

GetOracleServiceNameOk returns a tuple with the OracleServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleServiceName

`func (o *CreateDBTarget) SetOracleServiceName(v string)`

SetOracleServiceName sets OracleServiceName field to given value.

### HasOracleServiceName

`func (o *CreateDBTarget) HasOracleServiceName() bool`

HasOracleServiceName returns a boolean if a field has been set.

### GetPort

`func (o *CreateDBTarget) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateDBTarget) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateDBTarget) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateDBTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPwd

`func (o *CreateDBTarget) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *CreateDBTarget) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *CreateDBTarget) SetPwd(v string)`

SetPwd sets Pwd field to given value.

### HasPwd

`func (o *CreateDBTarget) HasPwd() bool`

HasPwd returns a boolean if a field has been set.

### GetSnowflakeAccount

`func (o *CreateDBTarget) GetSnowflakeAccount() string`

GetSnowflakeAccount returns the SnowflakeAccount field if non-nil, zero value otherwise.

### GetSnowflakeAccountOk

`func (o *CreateDBTarget) GetSnowflakeAccountOk() (*string, bool)`

GetSnowflakeAccountOk returns a tuple with the SnowflakeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeAccount

`func (o *CreateDBTarget) SetSnowflakeAccount(v string)`

SetSnowflakeAccount sets SnowflakeAccount field to given value.

### HasSnowflakeAccount

`func (o *CreateDBTarget) HasSnowflakeAccount() bool`

HasSnowflakeAccount returns a boolean if a field has been set.

### GetSnowflakeApiPrivateKey

`func (o *CreateDBTarget) GetSnowflakeApiPrivateKey() string`

GetSnowflakeApiPrivateKey returns the SnowflakeApiPrivateKey field if non-nil, zero value otherwise.

### GetSnowflakeApiPrivateKeyOk

`func (o *CreateDBTarget) GetSnowflakeApiPrivateKeyOk() (*string, bool)`

GetSnowflakeApiPrivateKeyOk returns a tuple with the SnowflakeApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeApiPrivateKey

`func (o *CreateDBTarget) SetSnowflakeApiPrivateKey(v string)`

SetSnowflakeApiPrivateKey sets SnowflakeApiPrivateKey field to given value.

### HasSnowflakeApiPrivateKey

`func (o *CreateDBTarget) HasSnowflakeApiPrivateKey() bool`

HasSnowflakeApiPrivateKey returns a boolean if a field has been set.

### GetSnowflakeApiPrivateKeyPassword

`func (o *CreateDBTarget) GetSnowflakeApiPrivateKeyPassword() string`

GetSnowflakeApiPrivateKeyPassword returns the SnowflakeApiPrivateKeyPassword field if non-nil, zero value otherwise.

### GetSnowflakeApiPrivateKeyPasswordOk

`func (o *CreateDBTarget) GetSnowflakeApiPrivateKeyPasswordOk() (*string, bool)`

GetSnowflakeApiPrivateKeyPasswordOk returns a tuple with the SnowflakeApiPrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeApiPrivateKeyPassword

`func (o *CreateDBTarget) SetSnowflakeApiPrivateKeyPassword(v string)`

SetSnowflakeApiPrivateKeyPassword sets SnowflakeApiPrivateKeyPassword field to given value.

### HasSnowflakeApiPrivateKeyPassword

`func (o *CreateDBTarget) HasSnowflakeApiPrivateKeyPassword() bool`

HasSnowflakeApiPrivateKeyPassword returns a boolean if a field has been set.

### GetSsl

`func (o *CreateDBTarget) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *CreateDBTarget) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *CreateDBTarget) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *CreateDBTarget) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *CreateDBTarget) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *CreateDBTarget) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *CreateDBTarget) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *CreateDBTarget) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetToken

`func (o *CreateDBTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateDBTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateDBTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateDBTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateDBTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateDBTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateDBTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateDBTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserName

`func (o *CreateDBTarget) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreateDBTarget) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreateDBTarget) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *CreateDBTarget) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


