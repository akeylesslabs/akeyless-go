# TargetUpdateDB

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
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**MongodbAtlas** | Pointer to **bool** |  | [optional] 
**MongodbAtlasApiPrivateKey** | Pointer to **string** | MongoDB Atlas private key | [optional] 
**MongodbAtlasApiPublicKey** | Pointer to **string** | MongoDB Atlas public key | [optional] 
**MongodbAtlasProjectId** | Pointer to **string** | MongoDB Atlas project ID | [optional] 
**MongodbDefaultAuthDb** | Pointer to **string** | MongoDB server default authentication database | [optional] 
**MongodbUriOptions** | Pointer to **string** | MongoDB server URI options | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
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

### NewTargetUpdateDB

`func NewTargetUpdateDB(connectionType string, dbType string, name string, ) *TargetUpdateDB`

NewTargetUpdateDB instantiates a new TargetUpdateDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateDBWithDefaults

`func NewTargetUpdateDBWithDefaults() *TargetUpdateDB`

NewTargetUpdateDBWithDefaults instantiates a new TargetUpdateDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureClientId

`func (o *TargetUpdateDB) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *TargetUpdateDB) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *TargetUpdateDB) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *TargetUpdateDB) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureClientSecret

`func (o *TargetUpdateDB) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *TargetUpdateDB) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *TargetUpdateDB) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### HasAzureClientSecret

`func (o *TargetUpdateDB) HasAzureClientSecret() bool`

HasAzureClientSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *TargetUpdateDB) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *TargetUpdateDB) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *TargetUpdateDB) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *TargetUpdateDB) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetCloudServiceProvider

`func (o *TargetUpdateDB) GetCloudServiceProvider() string`

GetCloudServiceProvider returns the CloudServiceProvider field if non-nil, zero value otherwise.

### GetCloudServiceProviderOk

`func (o *TargetUpdateDB) GetCloudServiceProviderOk() (*string, bool)`

GetCloudServiceProviderOk returns a tuple with the CloudServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudServiceProvider

`func (o *TargetUpdateDB) SetCloudServiceProvider(v string)`

SetCloudServiceProvider sets CloudServiceProvider field to given value.

### HasCloudServiceProvider

`func (o *TargetUpdateDB) HasCloudServiceProvider() bool`

HasCloudServiceProvider returns a boolean if a field has been set.

### GetClusterMode

`func (o *TargetUpdateDB) GetClusterMode() bool`

GetClusterMode returns the ClusterMode field if non-nil, zero value otherwise.

### GetClusterModeOk

`func (o *TargetUpdateDB) GetClusterModeOk() (*bool, bool)`

GetClusterModeOk returns a tuple with the ClusterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMode

`func (o *TargetUpdateDB) SetClusterMode(v bool)`

SetClusterMode sets ClusterMode field to given value.

### HasClusterMode

`func (o *TargetUpdateDB) HasClusterMode() bool`

HasClusterMode returns a boolean if a field has been set.

### GetComment

`func (o *TargetUpdateDB) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TargetUpdateDB) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TargetUpdateDB) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TargetUpdateDB) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConnectionType

`func (o *TargetUpdateDB) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *TargetUpdateDB) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *TargetUpdateDB) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetDbName

`func (o *TargetUpdateDB) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TargetUpdateDB) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TargetUpdateDB) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *TargetUpdateDB) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbServerCertificates

`func (o *TargetUpdateDB) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *TargetUpdateDB) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *TargetUpdateDB) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *TargetUpdateDB) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *TargetUpdateDB) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *TargetUpdateDB) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *TargetUpdateDB) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *TargetUpdateDB) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDbType

`func (o *TargetUpdateDB) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *TargetUpdateDB) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *TargetUpdateDB) SetDbType(v string)`

SetDbType sets DbType field to given value.


### GetDescription

`func (o *TargetUpdateDB) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateDB) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateDB) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateDB) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *TargetUpdateDB) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetUpdateDB) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetUpdateDB) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TargetUpdateDB) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateDB) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateDB) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateDB) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateDB) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateDB) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateDB) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateDB) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateDB) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateDB) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateDB) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateDB) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateDB) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateDB) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateDB) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateDB) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateDB) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetMongodbAtlas

`func (o *TargetUpdateDB) GetMongodbAtlas() bool`

GetMongodbAtlas returns the MongodbAtlas field if non-nil, zero value otherwise.

### GetMongodbAtlasOk

`func (o *TargetUpdateDB) GetMongodbAtlasOk() (*bool, bool)`

GetMongodbAtlasOk returns a tuple with the MongodbAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlas

`func (o *TargetUpdateDB) SetMongodbAtlas(v bool)`

SetMongodbAtlas sets MongodbAtlas field to given value.

### HasMongodbAtlas

`func (o *TargetUpdateDB) HasMongodbAtlas() bool`

HasMongodbAtlas returns a boolean if a field has been set.

### GetMongodbAtlasApiPrivateKey

`func (o *TargetUpdateDB) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *TargetUpdateDB) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *TargetUpdateDB) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *TargetUpdateDB) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *TargetUpdateDB) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *TargetUpdateDB) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *TargetUpdateDB) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *TargetUpdateDB) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *TargetUpdateDB) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *TargetUpdateDB) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *TargetUpdateDB) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *TargetUpdateDB) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *TargetUpdateDB) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *TargetUpdateDB) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *TargetUpdateDB) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *TargetUpdateDB) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *TargetUpdateDB) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *TargetUpdateDB) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *TargetUpdateDB) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *TargetUpdateDB) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateDB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateDB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateDB) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateDB) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateDB) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateDB) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateDB) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetOracleServiceName

`func (o *TargetUpdateDB) GetOracleServiceName() string`

GetOracleServiceName returns the OracleServiceName field if non-nil, zero value otherwise.

### GetOracleServiceNameOk

`func (o *TargetUpdateDB) GetOracleServiceNameOk() (*string, bool)`

GetOracleServiceNameOk returns a tuple with the OracleServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleServiceName

`func (o *TargetUpdateDB) SetOracleServiceName(v string)`

SetOracleServiceName sets OracleServiceName field to given value.

### HasOracleServiceName

`func (o *TargetUpdateDB) HasOracleServiceName() bool`

HasOracleServiceName returns a boolean if a field has been set.

### GetOracleWalletLoginType

`func (o *TargetUpdateDB) GetOracleWalletLoginType() string`

GetOracleWalletLoginType returns the OracleWalletLoginType field if non-nil, zero value otherwise.

### GetOracleWalletLoginTypeOk

`func (o *TargetUpdateDB) GetOracleWalletLoginTypeOk() (*string, bool)`

GetOracleWalletLoginTypeOk returns a tuple with the OracleWalletLoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleWalletLoginType

`func (o *TargetUpdateDB) SetOracleWalletLoginType(v string)`

SetOracleWalletLoginType sets OracleWalletLoginType field to given value.

### HasOracleWalletLoginType

`func (o *TargetUpdateDB) HasOracleWalletLoginType() bool`

HasOracleWalletLoginType returns a boolean if a field has been set.

### GetOracleWalletP12FileData

`func (o *TargetUpdateDB) GetOracleWalletP12FileData() string`

GetOracleWalletP12FileData returns the OracleWalletP12FileData field if non-nil, zero value otherwise.

### GetOracleWalletP12FileDataOk

`func (o *TargetUpdateDB) GetOracleWalletP12FileDataOk() (*string, bool)`

GetOracleWalletP12FileDataOk returns a tuple with the OracleWalletP12FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleWalletP12FileData

`func (o *TargetUpdateDB) SetOracleWalletP12FileData(v string)`

SetOracleWalletP12FileData sets OracleWalletP12FileData field to given value.

### HasOracleWalletP12FileData

`func (o *TargetUpdateDB) HasOracleWalletP12FileData() bool`

HasOracleWalletP12FileData returns a boolean if a field has been set.

### GetOracleWalletSsoFileData

`func (o *TargetUpdateDB) GetOracleWalletSsoFileData() string`

GetOracleWalletSsoFileData returns the OracleWalletSsoFileData field if non-nil, zero value otherwise.

### GetOracleWalletSsoFileDataOk

`func (o *TargetUpdateDB) GetOracleWalletSsoFileDataOk() (*string, bool)`

GetOracleWalletSsoFileDataOk returns a tuple with the OracleWalletSsoFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleWalletSsoFileData

`func (o *TargetUpdateDB) SetOracleWalletSsoFileData(v string)`

SetOracleWalletSsoFileData sets OracleWalletSsoFileData field to given value.

### HasOracleWalletSsoFileData

`func (o *TargetUpdateDB) HasOracleWalletSsoFileData() bool`

HasOracleWalletSsoFileData returns a boolean if a field has been set.

### GetParentTargetName

`func (o *TargetUpdateDB) GetParentTargetName() string`

GetParentTargetName returns the ParentTargetName field if non-nil, zero value otherwise.

### GetParentTargetNameOk

`func (o *TargetUpdateDB) GetParentTargetNameOk() (*string, bool)`

GetParentTargetNameOk returns a tuple with the ParentTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTargetName

`func (o *TargetUpdateDB) SetParentTargetName(v string)`

SetParentTargetName sets ParentTargetName field to given value.

### HasParentTargetName

`func (o *TargetUpdateDB) HasParentTargetName() bool`

HasParentTargetName returns a boolean if a field has been set.

### GetPort

`func (o *TargetUpdateDB) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetUpdateDB) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetUpdateDB) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetUpdateDB) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPwd

`func (o *TargetUpdateDB) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *TargetUpdateDB) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *TargetUpdateDB) SetPwd(v string)`

SetPwd sets Pwd field to given value.

### HasPwd

`func (o *TargetUpdateDB) HasPwd() bool`

HasPwd returns a boolean if a field has been set.

### GetSnowflakeAccount

`func (o *TargetUpdateDB) GetSnowflakeAccount() string`

GetSnowflakeAccount returns the SnowflakeAccount field if non-nil, zero value otherwise.

### GetSnowflakeAccountOk

`func (o *TargetUpdateDB) GetSnowflakeAccountOk() (*string, bool)`

GetSnowflakeAccountOk returns a tuple with the SnowflakeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeAccount

`func (o *TargetUpdateDB) SetSnowflakeAccount(v string)`

SetSnowflakeAccount sets SnowflakeAccount field to given value.

### HasSnowflakeAccount

`func (o *TargetUpdateDB) HasSnowflakeAccount() bool`

HasSnowflakeAccount returns a boolean if a field has been set.

### GetSnowflakeApiPrivateKey

`func (o *TargetUpdateDB) GetSnowflakeApiPrivateKey() string`

GetSnowflakeApiPrivateKey returns the SnowflakeApiPrivateKey field if non-nil, zero value otherwise.

### GetSnowflakeApiPrivateKeyOk

`func (o *TargetUpdateDB) GetSnowflakeApiPrivateKeyOk() (*string, bool)`

GetSnowflakeApiPrivateKeyOk returns a tuple with the SnowflakeApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeApiPrivateKey

`func (o *TargetUpdateDB) SetSnowflakeApiPrivateKey(v string)`

SetSnowflakeApiPrivateKey sets SnowflakeApiPrivateKey field to given value.

### HasSnowflakeApiPrivateKey

`func (o *TargetUpdateDB) HasSnowflakeApiPrivateKey() bool`

HasSnowflakeApiPrivateKey returns a boolean if a field has been set.

### GetSnowflakeApiPrivateKeyPassword

`func (o *TargetUpdateDB) GetSnowflakeApiPrivateKeyPassword() string`

GetSnowflakeApiPrivateKeyPassword returns the SnowflakeApiPrivateKeyPassword field if non-nil, zero value otherwise.

### GetSnowflakeApiPrivateKeyPasswordOk

`func (o *TargetUpdateDB) GetSnowflakeApiPrivateKeyPasswordOk() (*string, bool)`

GetSnowflakeApiPrivateKeyPasswordOk returns a tuple with the SnowflakeApiPrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeApiPrivateKeyPassword

`func (o *TargetUpdateDB) SetSnowflakeApiPrivateKeyPassword(v string)`

SetSnowflakeApiPrivateKeyPassword sets SnowflakeApiPrivateKeyPassword field to given value.

### HasSnowflakeApiPrivateKeyPassword

`func (o *TargetUpdateDB) HasSnowflakeApiPrivateKeyPassword() bool`

HasSnowflakeApiPrivateKeyPassword returns a boolean if a field has been set.

### GetSsl

`func (o *TargetUpdateDB) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *TargetUpdateDB) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *TargetUpdateDB) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *TargetUpdateDB) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *TargetUpdateDB) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *TargetUpdateDB) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *TargetUpdateDB) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *TargetUpdateDB) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateDB) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateDB) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateDB) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateDB) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateDB) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateDB) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateDB) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateDB) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserName

`func (o *TargetUpdateDB) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *TargetUpdateDB) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *TargetUpdateDB) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *TargetUpdateDB) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


