# GatewayUpdateProducerPostgreSQL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationStatements** | Pointer to **string** | PostgreSQL Creation statements | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**PostgresqlDbName** | Pointer to **string** | PostgreSQL DB Name | [optional] 
**PostgresqlHost** | Pointer to **string** | PostgreSQL Host | [optional] [default to "127.0.0.1"]
**PostgresqlPassword** | Pointer to **string** | PostgreSQL Password | [optional] 
**PostgresqlPort** | Pointer to **string** | PostgreSQL Port | [optional] [default to "5432"]
**PostgresqlUsername** | Pointer to **string** | PostgreSQL Username | [optional] 
**ProducerEncryptionKey** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RevocationStatement** | Pointer to **string** | PostgreSQL Revocation statements | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessDbSchema** | Pointer to **string** | The DB schema | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Ssl** | Pointer to **bool** | Enable/Disable SSL [true/false] | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayUpdateProducerPostgreSQL

`func NewGatewayUpdateProducerPostgreSQL(name string, ) *GatewayUpdateProducerPostgreSQL`

NewGatewayUpdateProducerPostgreSQL instantiates a new GatewayUpdateProducerPostgreSQL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerPostgreSQLWithDefaults

`func NewGatewayUpdateProducerPostgreSQLWithDefaults() *GatewayUpdateProducerPostgreSQL`

NewGatewayUpdateProducerPostgreSQLWithDefaults instantiates a new GatewayUpdateProducerPostgreSQL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationStatements

`func (o *GatewayUpdateProducerPostgreSQL) GetCreationStatements() string`

GetCreationStatements returns the CreationStatements field if non-nil, zero value otherwise.

### GetCreationStatementsOk

`func (o *GatewayUpdateProducerPostgreSQL) GetCreationStatementsOk() (*string, bool)`

GetCreationStatementsOk returns a tuple with the CreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStatements

`func (o *GatewayUpdateProducerPostgreSQL) SetCreationStatements(v string)`

SetCreationStatements sets CreationStatements field to given value.

### HasCreationStatements

`func (o *GatewayUpdateProducerPostgreSQL) HasCreationStatements() bool`

HasCreationStatements returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayUpdateProducerPostgreSQL) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerPostgreSQL) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerPostgreSQL) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerPostgreSQL) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateProducerPostgreSQL) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateProducerPostgreSQL) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateProducerPostgreSQL) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateProducerPostgreSQL) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerPostgreSQL) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerPostgreSQL) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerPostgreSQL) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerPostgreSQL) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerPostgreSQL) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerPostgreSQL) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerPostgreSQL) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPostgresqlDbName

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlDbName() string`

GetPostgresqlDbName returns the PostgresqlDbName field if non-nil, zero value otherwise.

### GetPostgresqlDbNameOk

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlDbNameOk() (*string, bool)`

GetPostgresqlDbNameOk returns a tuple with the PostgresqlDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlDbName

`func (o *GatewayUpdateProducerPostgreSQL) SetPostgresqlDbName(v string)`

SetPostgresqlDbName sets PostgresqlDbName field to given value.

### HasPostgresqlDbName

`func (o *GatewayUpdateProducerPostgreSQL) HasPostgresqlDbName() bool`

HasPostgresqlDbName returns a boolean if a field has been set.

### GetPostgresqlHost

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlHost() string`

GetPostgresqlHost returns the PostgresqlHost field if non-nil, zero value otherwise.

### GetPostgresqlHostOk

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlHostOk() (*string, bool)`

GetPostgresqlHostOk returns a tuple with the PostgresqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlHost

`func (o *GatewayUpdateProducerPostgreSQL) SetPostgresqlHost(v string)`

SetPostgresqlHost sets PostgresqlHost field to given value.

### HasPostgresqlHost

`func (o *GatewayUpdateProducerPostgreSQL) HasPostgresqlHost() bool`

HasPostgresqlHost returns a boolean if a field has been set.

### GetPostgresqlPassword

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlPassword() string`

GetPostgresqlPassword returns the PostgresqlPassword field if non-nil, zero value otherwise.

### GetPostgresqlPasswordOk

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlPasswordOk() (*string, bool)`

GetPostgresqlPasswordOk returns a tuple with the PostgresqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlPassword

`func (o *GatewayUpdateProducerPostgreSQL) SetPostgresqlPassword(v string)`

SetPostgresqlPassword sets PostgresqlPassword field to given value.

### HasPostgresqlPassword

`func (o *GatewayUpdateProducerPostgreSQL) HasPostgresqlPassword() bool`

HasPostgresqlPassword returns a boolean if a field has been set.

### GetPostgresqlPort

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlPort() string`

GetPostgresqlPort returns the PostgresqlPort field if non-nil, zero value otherwise.

### GetPostgresqlPortOk

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlPortOk() (*string, bool)`

GetPostgresqlPortOk returns a tuple with the PostgresqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlPort

`func (o *GatewayUpdateProducerPostgreSQL) SetPostgresqlPort(v string)`

SetPostgresqlPort sets PostgresqlPort field to given value.

### HasPostgresqlPort

`func (o *GatewayUpdateProducerPostgreSQL) HasPostgresqlPort() bool`

HasPostgresqlPort returns a boolean if a field has been set.

### GetPostgresqlUsername

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlUsername() string`

GetPostgresqlUsername returns the PostgresqlUsername field if non-nil, zero value otherwise.

### GetPostgresqlUsernameOk

`func (o *GatewayUpdateProducerPostgreSQL) GetPostgresqlUsernameOk() (*string, bool)`

GetPostgresqlUsernameOk returns a tuple with the PostgresqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlUsername

`func (o *GatewayUpdateProducerPostgreSQL) SetPostgresqlUsername(v string)`

SetPostgresqlUsername sets PostgresqlUsername field to given value.

### HasPostgresqlUsername

`func (o *GatewayUpdateProducerPostgreSQL) HasPostgresqlUsername() bool`

HasPostgresqlUsername returns a boolean if a field has been set.

### GetProducerEncryptionKey

`func (o *GatewayUpdateProducerPostgreSQL) GetProducerEncryptionKey() string`

GetProducerEncryptionKey returns the ProducerEncryptionKey field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyOk

`func (o *GatewayUpdateProducerPostgreSQL) GetProducerEncryptionKeyOk() (*string, bool)`

GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKey

`func (o *GatewayUpdateProducerPostgreSQL) SetProducerEncryptionKey(v string)`

SetProducerEncryptionKey sets ProducerEncryptionKey field to given value.

### HasProducerEncryptionKey

`func (o *GatewayUpdateProducerPostgreSQL) HasProducerEncryptionKey() bool`

HasProducerEncryptionKey returns a boolean if a field has been set.

### GetRevocationStatement

`func (o *GatewayUpdateProducerPostgreSQL) GetRevocationStatement() string`

GetRevocationStatement returns the RevocationStatement field if non-nil, zero value otherwise.

### GetRevocationStatementOk

`func (o *GatewayUpdateProducerPostgreSQL) GetRevocationStatementOk() (*string, bool)`

GetRevocationStatementOk returns a tuple with the RevocationStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationStatement

`func (o *GatewayUpdateProducerPostgreSQL) SetRevocationStatement(v string)`

SetRevocationStatement sets RevocationStatement field to given value.

### HasRevocationStatement

`func (o *GatewayUpdateProducerPostgreSQL) HasRevocationStatement() bool`

HasRevocationStatement returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerPostgreSQL) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerPostgreSQL) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessDbSchema

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessDbSchema() string`

GetSecureAccessDbSchema returns the SecureAccessDbSchema field if non-nil, zero value otherwise.

### GetSecureAccessDbSchemaOk

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessDbSchemaOk() (*string, bool)`

GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbSchema

`func (o *GatewayUpdateProducerPostgreSQL) SetSecureAccessDbSchema(v string)`

SetSecureAccessDbSchema sets SecureAccessDbSchema field to given value.

### HasSecureAccessDbSchema

`func (o *GatewayUpdateProducerPostgreSQL) HasSecureAccessDbSchema() bool`

HasSecureAccessDbSchema returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayUpdateProducerPostgreSQL) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayUpdateProducerPostgreSQL) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayUpdateProducerPostgreSQL) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayUpdateProducerPostgreSQL) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayUpdateProducerPostgreSQL) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayUpdateProducerPostgreSQL) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayUpdateProducerPostgreSQL) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSsl

`func (o *GatewayUpdateProducerPostgreSQL) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *GatewayUpdateProducerPostgreSQL) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *GatewayUpdateProducerPostgreSQL) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *GatewayUpdateProducerPostgreSQL) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerPostgreSQL) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerPostgreSQL) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerPostgreSQL) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerPostgreSQL) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerPostgreSQL) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerPostgreSQL) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerPostgreSQL) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerPostgreSQL) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerPostgreSQL) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerPostgreSQL) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerPostgreSQL) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerPostgreSQL) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerPostgreSQL) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerPostgreSQL) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerPostgreSQL) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerPostgreSQL) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerPostgreSQL) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerPostgreSQL) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerPostgreSQL) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerPostgreSQL) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


