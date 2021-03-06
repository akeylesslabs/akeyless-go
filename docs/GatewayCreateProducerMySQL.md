# GatewayCreateProducerMySQL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbServerCertificates** | Pointer to **string** | (Optional) DB server certificates | [optional] 
**DbServerName** | Pointer to **string** | (Optional) Server name for certificate verification | [optional] 
**MysqlDbname** | **string** | MySQL DB Name | 
**MysqlHost** | Pointer to **string** | MySQL Host | [optional] [default to "127.0.0.1"]
**MysqlPassword** | **string** | MySQL Password | 
**MysqlPort** | Pointer to **string** | MySQL Port | [optional] [default to "3306"]
**MysqlScreationStatements** | Pointer to **string** | MySQL Creation statements | [optional] 
**MysqlUsername** | **string** | MySQL Username | 
**Name** | **string** | Producer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayCreateProducerMySQL

`func NewGatewayCreateProducerMySQL(mysqlDbname string, mysqlPassword string, mysqlUsername string, name string, ) *GatewayCreateProducerMySQL`

NewGatewayCreateProducerMySQL instantiates a new GatewayCreateProducerMySQL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerMySQLWithDefaults

`func NewGatewayCreateProducerMySQLWithDefaults() *GatewayCreateProducerMySQL`

NewGatewayCreateProducerMySQLWithDefaults instantiates a new GatewayCreateProducerMySQL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbServerCertificates

`func (o *GatewayCreateProducerMySQL) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *GatewayCreateProducerMySQL) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *GatewayCreateProducerMySQL) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *GatewayCreateProducerMySQL) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *GatewayCreateProducerMySQL) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *GatewayCreateProducerMySQL) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *GatewayCreateProducerMySQL) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *GatewayCreateProducerMySQL) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetMysqlDbname

`func (o *GatewayCreateProducerMySQL) GetMysqlDbname() string`

GetMysqlDbname returns the MysqlDbname field if non-nil, zero value otherwise.

### GetMysqlDbnameOk

`func (o *GatewayCreateProducerMySQL) GetMysqlDbnameOk() (*string, bool)`

GetMysqlDbnameOk returns a tuple with the MysqlDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlDbname

`func (o *GatewayCreateProducerMySQL) SetMysqlDbname(v string)`

SetMysqlDbname sets MysqlDbname field to given value.


### GetMysqlHost

`func (o *GatewayCreateProducerMySQL) GetMysqlHost() string`

GetMysqlHost returns the MysqlHost field if non-nil, zero value otherwise.

### GetMysqlHostOk

`func (o *GatewayCreateProducerMySQL) GetMysqlHostOk() (*string, bool)`

GetMysqlHostOk returns a tuple with the MysqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlHost

`func (o *GatewayCreateProducerMySQL) SetMysqlHost(v string)`

SetMysqlHost sets MysqlHost field to given value.

### HasMysqlHost

`func (o *GatewayCreateProducerMySQL) HasMysqlHost() bool`

HasMysqlHost returns a boolean if a field has been set.

### GetMysqlPassword

`func (o *GatewayCreateProducerMySQL) GetMysqlPassword() string`

GetMysqlPassword returns the MysqlPassword field if non-nil, zero value otherwise.

### GetMysqlPasswordOk

`func (o *GatewayCreateProducerMySQL) GetMysqlPasswordOk() (*string, bool)`

GetMysqlPasswordOk returns a tuple with the MysqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlPassword

`func (o *GatewayCreateProducerMySQL) SetMysqlPassword(v string)`

SetMysqlPassword sets MysqlPassword field to given value.


### GetMysqlPort

`func (o *GatewayCreateProducerMySQL) GetMysqlPort() string`

GetMysqlPort returns the MysqlPort field if non-nil, zero value otherwise.

### GetMysqlPortOk

`func (o *GatewayCreateProducerMySQL) GetMysqlPortOk() (*string, bool)`

GetMysqlPortOk returns a tuple with the MysqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlPort

`func (o *GatewayCreateProducerMySQL) SetMysqlPort(v string)`

SetMysqlPort sets MysqlPort field to given value.

### HasMysqlPort

`func (o *GatewayCreateProducerMySQL) HasMysqlPort() bool`

HasMysqlPort returns a boolean if a field has been set.

### GetMysqlScreationStatements

`func (o *GatewayCreateProducerMySQL) GetMysqlScreationStatements() string`

GetMysqlScreationStatements returns the MysqlScreationStatements field if non-nil, zero value otherwise.

### GetMysqlScreationStatementsOk

`func (o *GatewayCreateProducerMySQL) GetMysqlScreationStatementsOk() (*string, bool)`

GetMysqlScreationStatementsOk returns a tuple with the MysqlScreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlScreationStatements

`func (o *GatewayCreateProducerMySQL) SetMysqlScreationStatements(v string)`

SetMysqlScreationStatements sets MysqlScreationStatements field to given value.

### HasMysqlScreationStatements

`func (o *GatewayCreateProducerMySQL) HasMysqlScreationStatements() bool`

HasMysqlScreationStatements returns a boolean if a field has been set.

### GetMysqlUsername

`func (o *GatewayCreateProducerMySQL) GetMysqlUsername() string`

GetMysqlUsername returns the MysqlUsername field if non-nil, zero value otherwise.

### GetMysqlUsernameOk

`func (o *GatewayCreateProducerMySQL) GetMysqlUsernameOk() (*string, bool)`

GetMysqlUsernameOk returns a tuple with the MysqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlUsername

`func (o *GatewayCreateProducerMySQL) SetMysqlUsername(v string)`

SetMysqlUsername sets MysqlUsername field to given value.


### GetName

`func (o *GatewayCreateProducerMySQL) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerMySQL) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerMySQL) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *GatewayCreateProducerMySQL) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateProducerMySQL) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateProducerMySQL) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateProducerMySQL) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerMySQL) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerMySQL) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerMySQL) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerMySQL) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerMySQL) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerMySQL) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerMySQL) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerMySQL) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerMySQL) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerMySQL) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerMySQL) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerMySQL) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerMySQL) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerMySQL) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerMySQL) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerMySQL) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayCreateProducerMySQL) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateProducerMySQL) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateProducerMySQL) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateProducerMySQL) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


