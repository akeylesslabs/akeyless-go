# GatewayCreateProducerPostgreSQL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationStatements** | Pointer to **string** | PostgreSQL Creation statements | [optional] 
**Name** | **string** | Producer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**PostgresqlDbName** | **string** | PostgreSQL DB Name | 
**PostgresqlHost** | Pointer to **string** | PostgreSQL Host | [optional] [default to "127.0.0.1"]
**PostgresqlPassword** | **string** | PostgreSQL Password | 
**PostgresqlPort** | Pointer to **string** | PostgreSQL Port | [optional] [default to "5432"]
**PostgresqlUsername** | **string** | PostgreSQL Username | 
**ProducerEncryptionKey** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayCreateProducerPostgreSQL

`func NewGatewayCreateProducerPostgreSQL(name string, postgresqlDbName string, postgresqlPassword string, postgresqlUsername string, ) *GatewayCreateProducerPostgreSQL`

NewGatewayCreateProducerPostgreSQL instantiates a new GatewayCreateProducerPostgreSQL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerPostgreSQLWithDefaults

`func NewGatewayCreateProducerPostgreSQLWithDefaults() *GatewayCreateProducerPostgreSQL`

NewGatewayCreateProducerPostgreSQLWithDefaults instantiates a new GatewayCreateProducerPostgreSQL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationStatements

`func (o *GatewayCreateProducerPostgreSQL) GetCreationStatements() string`

GetCreationStatements returns the CreationStatements field if non-nil, zero value otherwise.

### GetCreationStatementsOk

`func (o *GatewayCreateProducerPostgreSQL) GetCreationStatementsOk() (*string, bool)`

GetCreationStatementsOk returns a tuple with the CreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStatements

`func (o *GatewayCreateProducerPostgreSQL) SetCreationStatements(v string)`

SetCreationStatements sets CreationStatements field to given value.

### HasCreationStatements

`func (o *GatewayCreateProducerPostgreSQL) HasCreationStatements() bool`

HasCreationStatements returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerPostgreSQL) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerPostgreSQL) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerPostgreSQL) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *GatewayCreateProducerPostgreSQL) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateProducerPostgreSQL) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateProducerPostgreSQL) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateProducerPostgreSQL) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPostgresqlDbName

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlDbName() string`

GetPostgresqlDbName returns the PostgresqlDbName field if non-nil, zero value otherwise.

### GetPostgresqlDbNameOk

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlDbNameOk() (*string, bool)`

GetPostgresqlDbNameOk returns a tuple with the PostgresqlDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlDbName

`func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlDbName(v string)`

SetPostgresqlDbName sets PostgresqlDbName field to given value.


### GetPostgresqlHost

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlHost() string`

GetPostgresqlHost returns the PostgresqlHost field if non-nil, zero value otherwise.

### GetPostgresqlHostOk

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlHostOk() (*string, bool)`

GetPostgresqlHostOk returns a tuple with the PostgresqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlHost

`func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlHost(v string)`

SetPostgresqlHost sets PostgresqlHost field to given value.

### HasPostgresqlHost

`func (o *GatewayCreateProducerPostgreSQL) HasPostgresqlHost() bool`

HasPostgresqlHost returns a boolean if a field has been set.

### GetPostgresqlPassword

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlPassword() string`

GetPostgresqlPassword returns the PostgresqlPassword field if non-nil, zero value otherwise.

### GetPostgresqlPasswordOk

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlPasswordOk() (*string, bool)`

GetPostgresqlPasswordOk returns a tuple with the PostgresqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlPassword

`func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlPassword(v string)`

SetPostgresqlPassword sets PostgresqlPassword field to given value.


### GetPostgresqlPort

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlPort() string`

GetPostgresqlPort returns the PostgresqlPort field if non-nil, zero value otherwise.

### GetPostgresqlPortOk

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlPortOk() (*string, bool)`

GetPostgresqlPortOk returns a tuple with the PostgresqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlPort

`func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlPort(v string)`

SetPostgresqlPort sets PostgresqlPort field to given value.

### HasPostgresqlPort

`func (o *GatewayCreateProducerPostgreSQL) HasPostgresqlPort() bool`

HasPostgresqlPort returns a boolean if a field has been set.

### GetPostgresqlUsername

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlUsername() string`

GetPostgresqlUsername returns the PostgresqlUsername field if non-nil, zero value otherwise.

### GetPostgresqlUsernameOk

`func (o *GatewayCreateProducerPostgreSQL) GetPostgresqlUsernameOk() (*string, bool)`

GetPostgresqlUsernameOk returns a tuple with the PostgresqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlUsername

`func (o *GatewayCreateProducerPostgreSQL) SetPostgresqlUsername(v string)`

SetPostgresqlUsername sets PostgresqlUsername field to given value.


### GetProducerEncryptionKey

`func (o *GatewayCreateProducerPostgreSQL) GetProducerEncryptionKey() string`

GetProducerEncryptionKey returns the ProducerEncryptionKey field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyOk

`func (o *GatewayCreateProducerPostgreSQL) GetProducerEncryptionKeyOk() (*string, bool)`

GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKey

`func (o *GatewayCreateProducerPostgreSQL) SetProducerEncryptionKey(v string)`

SetProducerEncryptionKey sets ProducerEncryptionKey field to given value.

### HasProducerEncryptionKey

`func (o *GatewayCreateProducerPostgreSQL) HasProducerEncryptionKey() bool`

HasProducerEncryptionKey returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerPostgreSQL) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerPostgreSQL) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerPostgreSQL) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerPostgreSQL) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerPostgreSQL) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerPostgreSQL) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerPostgreSQL) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerPostgreSQL) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerPostgreSQL) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerPostgreSQL) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerPostgreSQL) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerPostgreSQL) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayCreateProducerPostgreSQL) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateProducerPostgreSQL) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateProducerPostgreSQL) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateProducerPostgreSQL) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


