# UpdateDBTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**DbServerCertificates** | Pointer to **string** | (Optional) DB server certificates | [optional] 
**DbServerName** | Pointer to **string** | (Optional) Server name for certificate verification | [optional] 
**DbType** | **string** |  | 
**Host** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MongodbAtlas** | Pointer to **bool** |  | [optional] 
**MongodbAtlasApiPrivateKey** | Pointer to **string** | MongoDB Atlas private key | [optional] 
**MongodbAtlasApiPublicKey** | Pointer to **string** | MongoDB Atlas public key | [optional] 
**MongodbAtlasProjectId** | Pointer to **string** | MongoDB Atlas project ID | [optional] 
**MongodbDefaultAuthDb** | Pointer to **string** | MongoDB server default authentication database | [optional] 
**MongodbUriOptions** | Pointer to **string** | MongoDB server URI options | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**OracleServiceName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Pwd** | Pointer to **string** |  | [optional] 
**SnowflakeAccount** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Create new version for the target | [optional] [default to false]
**UserName** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewUpdateDBTarget

`func NewUpdateDBTarget(dbType string, name string, ) *UpdateDBTarget`

NewUpdateDBTarget instantiates a new UpdateDBTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDBTargetWithDefaults

`func NewUpdateDBTargetWithDefaults() *UpdateDBTarget`

NewUpdateDBTargetWithDefaults instantiates a new UpdateDBTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateDBTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateDBTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateDBTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateDBTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDbName

`func (o *UpdateDBTarget) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *UpdateDBTarget) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *UpdateDBTarget) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *UpdateDBTarget) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbServerCertificates

`func (o *UpdateDBTarget) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *UpdateDBTarget) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *UpdateDBTarget) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *UpdateDBTarget) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *UpdateDBTarget) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *UpdateDBTarget) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *UpdateDBTarget) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *UpdateDBTarget) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDbType

`func (o *UpdateDBTarget) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *UpdateDBTarget) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *UpdateDBTarget) SetDbType(v string)`

SetDbType sets DbType field to given value.


### GetHost

`func (o *UpdateDBTarget) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateDBTarget) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateDBTarget) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateDBTarget) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKey

`func (o *UpdateDBTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateDBTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateDBTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateDBTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMongodbAtlas

`func (o *UpdateDBTarget) GetMongodbAtlas() bool`

GetMongodbAtlas returns the MongodbAtlas field if non-nil, zero value otherwise.

### GetMongodbAtlasOk

`func (o *UpdateDBTarget) GetMongodbAtlasOk() (*bool, bool)`

GetMongodbAtlasOk returns a tuple with the MongodbAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlas

`func (o *UpdateDBTarget) SetMongodbAtlas(v bool)`

SetMongodbAtlas sets MongodbAtlas field to given value.

### HasMongodbAtlas

`func (o *UpdateDBTarget) HasMongodbAtlas() bool`

HasMongodbAtlas returns a boolean if a field has been set.

### GetMongodbAtlasApiPrivateKey

`func (o *UpdateDBTarget) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *UpdateDBTarget) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *UpdateDBTarget) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *UpdateDBTarget) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *UpdateDBTarget) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *UpdateDBTarget) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *UpdateDBTarget) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *UpdateDBTarget) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *UpdateDBTarget) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *UpdateDBTarget) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *UpdateDBTarget) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *UpdateDBTarget) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *UpdateDBTarget) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *UpdateDBTarget) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *UpdateDBTarget) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *UpdateDBTarget) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *UpdateDBTarget) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *UpdateDBTarget) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *UpdateDBTarget) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *UpdateDBTarget) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetName

`func (o *UpdateDBTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDBTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDBTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateDBTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateDBTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateDBTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateDBTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetOracleServiceName

`func (o *UpdateDBTarget) GetOracleServiceName() string`

GetOracleServiceName returns the OracleServiceName field if non-nil, zero value otherwise.

### GetOracleServiceNameOk

`func (o *UpdateDBTarget) GetOracleServiceNameOk() (*string, bool)`

GetOracleServiceNameOk returns a tuple with the OracleServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleServiceName

`func (o *UpdateDBTarget) SetOracleServiceName(v string)`

SetOracleServiceName sets OracleServiceName field to given value.

### HasOracleServiceName

`func (o *UpdateDBTarget) HasOracleServiceName() bool`

HasOracleServiceName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateDBTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateDBTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateDBTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateDBTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *UpdateDBTarget) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateDBTarget) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateDBTarget) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateDBTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPwd

`func (o *UpdateDBTarget) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *UpdateDBTarget) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *UpdateDBTarget) SetPwd(v string)`

SetPwd sets Pwd field to given value.

### HasPwd

`func (o *UpdateDBTarget) HasPwd() bool`

HasPwd returns a boolean if a field has been set.

### GetSnowflakeAccount

`func (o *UpdateDBTarget) GetSnowflakeAccount() string`

GetSnowflakeAccount returns the SnowflakeAccount field if non-nil, zero value otherwise.

### GetSnowflakeAccountOk

`func (o *UpdateDBTarget) GetSnowflakeAccountOk() (*string, bool)`

GetSnowflakeAccountOk returns a tuple with the SnowflakeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeAccount

`func (o *UpdateDBTarget) SetSnowflakeAccount(v string)`

SetSnowflakeAccount sets SnowflakeAccount field to given value.

### HasSnowflakeAccount

`func (o *UpdateDBTarget) HasSnowflakeAccount() bool`

HasSnowflakeAccount returns a boolean if a field has been set.

### GetToken

`func (o *UpdateDBTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateDBTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateDBTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateDBTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateDBTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateDBTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateDBTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateDBTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateDBTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateDBTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateDBTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateDBTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.

### GetUserName

`func (o *UpdateDBTarget) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UpdateDBTarget) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UpdateDBTarget) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UpdateDBTarget) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateDBTarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateDBTarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateDBTarget) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateDBTarget) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


