# TargetCreateDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureClientId** | Pointer to **string** | (Optional) Client id (relevant for \&quot;cloud-service-provider\&quot; only) | [optional] 
**AzureClientSecret** | Pointer to **string** | (Optional) Client secret (relevant for \&quot;cloud-service-provider\&quot; only) | [optional] 
**AzureTenantId** | Pointer to **string** | (Optional) Tenant id (relevant for \&quot;cloud-service-provider\&quot; only) | [optional] 
**CloudServiceProvider** | Pointer to **string** | (Optional) Cloud service provider (currently only supports Azure) | [optional] 
**ClusterMode** | Pointer to **bool** | Cluster Mode | [optional] 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**ConnectionType** | **string** | Type of connection to mssql database [credentials/cloud-identity/wallet/parent-target] | [default to "credentials"]
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
**OracleServiceName** | Pointer to **string** | Oracle db service name | [optional] 
**OracleWalletLoginType** | Pointer to **string** | Oracle Wallet login type (password/mtls) | [optional] 
**OracleWalletP12FileData** | Pointer to **string** | Oracle wallet p12 file data in base64 | [optional] 
**OracleWalletSsoFileData** | Pointer to **string** | Oracle wallet sso file data in base64 | [optional] 
**ParentTargetName** | Pointer to **string** | Name of the parent target, relevant only when connection-type is parent-target | [optional] 
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

### NewTargetCreateDB

`func NewTargetCreateDB(connectionType string, dbType string, name string, ) *TargetCreateDB`

NewTargetCreateDB instantiates a new TargetCreateDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateDBWithDefaults

`func NewTargetCreateDBWithDefaults() *TargetCreateDB`

NewTargetCreateDBWithDefaults instantiates a new TargetCreateDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureClientId

`func (o *TargetCreateDB) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *TargetCreateDB) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *TargetCreateDB) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *TargetCreateDB) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureClientSecret

`func (o *TargetCreateDB) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *TargetCreateDB) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *TargetCreateDB) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### HasAzureClientSecret

`func (o *TargetCreateDB) HasAzureClientSecret() bool`

HasAzureClientSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *TargetCreateDB) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *TargetCreateDB) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *TargetCreateDB) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *TargetCreateDB) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetCloudServiceProvider

`func (o *TargetCreateDB) GetCloudServiceProvider() string`

GetCloudServiceProvider returns the CloudServiceProvider field if non-nil, zero value otherwise.

### GetCloudServiceProviderOk

`func (o *TargetCreateDB) GetCloudServiceProviderOk() (*string, bool)`

GetCloudServiceProviderOk returns a tuple with the CloudServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudServiceProvider

`func (o *TargetCreateDB) SetCloudServiceProvider(v string)`

SetCloudServiceProvider sets CloudServiceProvider field to given value.

### HasCloudServiceProvider

`func (o *TargetCreateDB) HasCloudServiceProvider() bool`

HasCloudServiceProvider returns a boolean if a field has been set.

### GetClusterMode

`func (o *TargetCreateDB) GetClusterMode() bool`

GetClusterMode returns the ClusterMode field if non-nil, zero value otherwise.

### GetClusterModeOk

`func (o *TargetCreateDB) GetClusterModeOk() (*bool, bool)`

GetClusterModeOk returns a tuple with the ClusterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMode

`func (o *TargetCreateDB) SetClusterMode(v bool)`

SetClusterMode sets ClusterMode field to given value.

### HasClusterMode

`func (o *TargetCreateDB) HasClusterMode() bool`

HasClusterMode returns a boolean if a field has been set.

### GetComment

`func (o *TargetCreateDB) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TargetCreateDB) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TargetCreateDB) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TargetCreateDB) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConnectionType

`func (o *TargetCreateDB) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *TargetCreateDB) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *TargetCreateDB) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetDbName

`func (o *TargetCreateDB) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TargetCreateDB) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TargetCreateDB) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *TargetCreateDB) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbServerCertificates

`func (o *TargetCreateDB) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *TargetCreateDB) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *TargetCreateDB) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *TargetCreateDB) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *TargetCreateDB) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *TargetCreateDB) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *TargetCreateDB) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *TargetCreateDB) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDbType

`func (o *TargetCreateDB) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *TargetCreateDB) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *TargetCreateDB) SetDbType(v string)`

SetDbType sets DbType field to given value.


### GetDescription

`func (o *TargetCreateDB) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateDB) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateDB) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateDB) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *TargetCreateDB) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetCreateDB) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetCreateDB) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TargetCreateDB) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateDB) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateDB) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateDB) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateDB) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateDB) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateDB) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateDB) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateDB) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateDB) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateDB) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateDB) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateDB) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetMongodbAtlas

`func (o *TargetCreateDB) GetMongodbAtlas() bool`

GetMongodbAtlas returns the MongodbAtlas field if non-nil, zero value otherwise.

### GetMongodbAtlasOk

`func (o *TargetCreateDB) GetMongodbAtlasOk() (*bool, bool)`

GetMongodbAtlasOk returns a tuple with the MongodbAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlas

`func (o *TargetCreateDB) SetMongodbAtlas(v bool)`

SetMongodbAtlas sets MongodbAtlas field to given value.

### HasMongodbAtlas

`func (o *TargetCreateDB) HasMongodbAtlas() bool`

HasMongodbAtlas returns a boolean if a field has been set.

### GetMongodbAtlasApiPrivateKey

`func (o *TargetCreateDB) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *TargetCreateDB) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *TargetCreateDB) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *TargetCreateDB) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *TargetCreateDB) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *TargetCreateDB) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *TargetCreateDB) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *TargetCreateDB) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *TargetCreateDB) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *TargetCreateDB) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *TargetCreateDB) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *TargetCreateDB) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *TargetCreateDB) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *TargetCreateDB) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *TargetCreateDB) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *TargetCreateDB) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *TargetCreateDB) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *TargetCreateDB) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *TargetCreateDB) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *TargetCreateDB) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateDB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateDB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateDB) SetName(v string)`

SetName sets Name field to given value.


### GetOracleServiceName

`func (o *TargetCreateDB) GetOracleServiceName() string`

GetOracleServiceName returns the OracleServiceName field if non-nil, zero value otherwise.

### GetOracleServiceNameOk

`func (o *TargetCreateDB) GetOracleServiceNameOk() (*string, bool)`

GetOracleServiceNameOk returns a tuple with the OracleServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleServiceName

`func (o *TargetCreateDB) SetOracleServiceName(v string)`

SetOracleServiceName sets OracleServiceName field to given value.

### HasOracleServiceName

`func (o *TargetCreateDB) HasOracleServiceName() bool`

HasOracleServiceName returns a boolean if a field has been set.

### GetOracleWalletLoginType

`func (o *TargetCreateDB) GetOracleWalletLoginType() string`

GetOracleWalletLoginType returns the OracleWalletLoginType field if non-nil, zero value otherwise.

### GetOracleWalletLoginTypeOk

`func (o *TargetCreateDB) GetOracleWalletLoginTypeOk() (*string, bool)`

GetOracleWalletLoginTypeOk returns a tuple with the OracleWalletLoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleWalletLoginType

`func (o *TargetCreateDB) SetOracleWalletLoginType(v string)`

SetOracleWalletLoginType sets OracleWalletLoginType field to given value.

### HasOracleWalletLoginType

`func (o *TargetCreateDB) HasOracleWalletLoginType() bool`

HasOracleWalletLoginType returns a boolean if a field has been set.

### GetOracleWalletP12FileData

`func (o *TargetCreateDB) GetOracleWalletP12FileData() string`

GetOracleWalletP12FileData returns the OracleWalletP12FileData field if non-nil, zero value otherwise.

### GetOracleWalletP12FileDataOk

`func (o *TargetCreateDB) GetOracleWalletP12FileDataOk() (*string, bool)`

GetOracleWalletP12FileDataOk returns a tuple with the OracleWalletP12FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleWalletP12FileData

`func (o *TargetCreateDB) SetOracleWalletP12FileData(v string)`

SetOracleWalletP12FileData sets OracleWalletP12FileData field to given value.

### HasOracleWalletP12FileData

`func (o *TargetCreateDB) HasOracleWalletP12FileData() bool`

HasOracleWalletP12FileData returns a boolean if a field has been set.

### GetOracleWalletSsoFileData

`func (o *TargetCreateDB) GetOracleWalletSsoFileData() string`

GetOracleWalletSsoFileData returns the OracleWalletSsoFileData field if non-nil, zero value otherwise.

### GetOracleWalletSsoFileDataOk

`func (o *TargetCreateDB) GetOracleWalletSsoFileDataOk() (*string, bool)`

GetOracleWalletSsoFileDataOk returns a tuple with the OracleWalletSsoFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleWalletSsoFileData

`func (o *TargetCreateDB) SetOracleWalletSsoFileData(v string)`

SetOracleWalletSsoFileData sets OracleWalletSsoFileData field to given value.

### HasOracleWalletSsoFileData

`func (o *TargetCreateDB) HasOracleWalletSsoFileData() bool`

HasOracleWalletSsoFileData returns a boolean if a field has been set.

### GetParentTargetName

`func (o *TargetCreateDB) GetParentTargetName() string`

GetParentTargetName returns the ParentTargetName field if non-nil, zero value otherwise.

### GetParentTargetNameOk

`func (o *TargetCreateDB) GetParentTargetNameOk() (*string, bool)`

GetParentTargetNameOk returns a tuple with the ParentTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTargetName

`func (o *TargetCreateDB) SetParentTargetName(v string)`

SetParentTargetName sets ParentTargetName field to given value.

### HasParentTargetName

`func (o *TargetCreateDB) HasParentTargetName() bool`

HasParentTargetName returns a boolean if a field has been set.

### GetPort

`func (o *TargetCreateDB) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetCreateDB) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetCreateDB) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetCreateDB) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPwd

`func (o *TargetCreateDB) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *TargetCreateDB) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *TargetCreateDB) SetPwd(v string)`

SetPwd sets Pwd field to given value.

### HasPwd

`func (o *TargetCreateDB) HasPwd() bool`

HasPwd returns a boolean if a field has been set.

### GetSnowflakeAccount

`func (o *TargetCreateDB) GetSnowflakeAccount() string`

GetSnowflakeAccount returns the SnowflakeAccount field if non-nil, zero value otherwise.

### GetSnowflakeAccountOk

`func (o *TargetCreateDB) GetSnowflakeAccountOk() (*string, bool)`

GetSnowflakeAccountOk returns a tuple with the SnowflakeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeAccount

`func (o *TargetCreateDB) SetSnowflakeAccount(v string)`

SetSnowflakeAccount sets SnowflakeAccount field to given value.

### HasSnowflakeAccount

`func (o *TargetCreateDB) HasSnowflakeAccount() bool`

HasSnowflakeAccount returns a boolean if a field has been set.

### GetSnowflakeApiPrivateKey

`func (o *TargetCreateDB) GetSnowflakeApiPrivateKey() string`

GetSnowflakeApiPrivateKey returns the SnowflakeApiPrivateKey field if non-nil, zero value otherwise.

### GetSnowflakeApiPrivateKeyOk

`func (o *TargetCreateDB) GetSnowflakeApiPrivateKeyOk() (*string, bool)`

GetSnowflakeApiPrivateKeyOk returns a tuple with the SnowflakeApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeApiPrivateKey

`func (o *TargetCreateDB) SetSnowflakeApiPrivateKey(v string)`

SetSnowflakeApiPrivateKey sets SnowflakeApiPrivateKey field to given value.

### HasSnowflakeApiPrivateKey

`func (o *TargetCreateDB) HasSnowflakeApiPrivateKey() bool`

HasSnowflakeApiPrivateKey returns a boolean if a field has been set.

### GetSnowflakeApiPrivateKeyPassword

`func (o *TargetCreateDB) GetSnowflakeApiPrivateKeyPassword() string`

GetSnowflakeApiPrivateKeyPassword returns the SnowflakeApiPrivateKeyPassword field if non-nil, zero value otherwise.

### GetSnowflakeApiPrivateKeyPasswordOk

`func (o *TargetCreateDB) GetSnowflakeApiPrivateKeyPasswordOk() (*string, bool)`

GetSnowflakeApiPrivateKeyPasswordOk returns a tuple with the SnowflakeApiPrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeApiPrivateKeyPassword

`func (o *TargetCreateDB) SetSnowflakeApiPrivateKeyPassword(v string)`

SetSnowflakeApiPrivateKeyPassword sets SnowflakeApiPrivateKeyPassword field to given value.

### HasSnowflakeApiPrivateKeyPassword

`func (o *TargetCreateDB) HasSnowflakeApiPrivateKeyPassword() bool`

HasSnowflakeApiPrivateKeyPassword returns a boolean if a field has been set.

### GetSsl

`func (o *TargetCreateDB) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *TargetCreateDB) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *TargetCreateDB) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *TargetCreateDB) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *TargetCreateDB) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *TargetCreateDB) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *TargetCreateDB) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *TargetCreateDB) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateDB) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateDB) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateDB) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateDB) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateDB) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateDB) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateDB) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateDB) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserName

`func (o *TargetCreateDB) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *TargetCreateDB) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *TargetCreateDB) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *TargetCreateDB) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


