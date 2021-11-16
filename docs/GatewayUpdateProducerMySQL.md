# GatewayUpdateProducerMySQL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbServerCertificates** | Pointer to **string** | (Optional) DB server certificates | [optional] 
**DbServerName** | Pointer to **string** | (Optional) Server name for certificate verification | [optional] 
**MysqlDbname** | Pointer to **string** | MySQL DB Name | [optional] 
**MysqlHost** | Pointer to **string** | MySQL Host | [optional] [default to "127.0.0.1"]
**MysqlPassword** | Pointer to **string** | MySQL Password | [optional] 
**MysqlPort** | Pointer to **string** | MySQL Port | [optional] [default to "3306"]
**MysqlScreationStatements** | Pointer to **string** | MySQL Creation statements | [optional] 
**MysqlUsername** | Pointer to **string** | MySQL Username | [optional] 
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** |  | [optional] 
**SecureAccessEnable** | Pointer to **string** |  | [optional] 
**SecureAccessHost** | Pointer to **[]string** |  | [optional] 
**SecureAccessWeb** | Pointer to **bool** |  | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayUpdateProducerMySQL

`func NewGatewayUpdateProducerMySQL(name string, ) *GatewayUpdateProducerMySQL`

NewGatewayUpdateProducerMySQL instantiates a new GatewayUpdateProducerMySQL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerMySQLWithDefaults

`func NewGatewayUpdateProducerMySQLWithDefaults() *GatewayUpdateProducerMySQL`

NewGatewayUpdateProducerMySQLWithDefaults instantiates a new GatewayUpdateProducerMySQL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbServerCertificates

`func (o *GatewayUpdateProducerMySQL) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *GatewayUpdateProducerMySQL) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *GatewayUpdateProducerMySQL) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *GatewayUpdateProducerMySQL) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *GatewayUpdateProducerMySQL) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *GatewayUpdateProducerMySQL) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *GatewayUpdateProducerMySQL) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *GatewayUpdateProducerMySQL) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetMysqlDbname

`func (o *GatewayUpdateProducerMySQL) GetMysqlDbname() string`

GetMysqlDbname returns the MysqlDbname field if non-nil, zero value otherwise.

### GetMysqlDbnameOk

`func (o *GatewayUpdateProducerMySQL) GetMysqlDbnameOk() (*string, bool)`

GetMysqlDbnameOk returns a tuple with the MysqlDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlDbname

`func (o *GatewayUpdateProducerMySQL) SetMysqlDbname(v string)`

SetMysqlDbname sets MysqlDbname field to given value.

### HasMysqlDbname

`func (o *GatewayUpdateProducerMySQL) HasMysqlDbname() bool`

HasMysqlDbname returns a boolean if a field has been set.

### GetMysqlHost

`func (o *GatewayUpdateProducerMySQL) GetMysqlHost() string`

GetMysqlHost returns the MysqlHost field if non-nil, zero value otherwise.

### GetMysqlHostOk

`func (o *GatewayUpdateProducerMySQL) GetMysqlHostOk() (*string, bool)`

GetMysqlHostOk returns a tuple with the MysqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlHost

`func (o *GatewayUpdateProducerMySQL) SetMysqlHost(v string)`

SetMysqlHost sets MysqlHost field to given value.

### HasMysqlHost

`func (o *GatewayUpdateProducerMySQL) HasMysqlHost() bool`

HasMysqlHost returns a boolean if a field has been set.

### GetMysqlPassword

`func (o *GatewayUpdateProducerMySQL) GetMysqlPassword() string`

GetMysqlPassword returns the MysqlPassword field if non-nil, zero value otherwise.

### GetMysqlPasswordOk

`func (o *GatewayUpdateProducerMySQL) GetMysqlPasswordOk() (*string, bool)`

GetMysqlPasswordOk returns a tuple with the MysqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlPassword

`func (o *GatewayUpdateProducerMySQL) SetMysqlPassword(v string)`

SetMysqlPassword sets MysqlPassword field to given value.

### HasMysqlPassword

`func (o *GatewayUpdateProducerMySQL) HasMysqlPassword() bool`

HasMysqlPassword returns a boolean if a field has been set.

### GetMysqlPort

`func (o *GatewayUpdateProducerMySQL) GetMysqlPort() string`

GetMysqlPort returns the MysqlPort field if non-nil, zero value otherwise.

### GetMysqlPortOk

`func (o *GatewayUpdateProducerMySQL) GetMysqlPortOk() (*string, bool)`

GetMysqlPortOk returns a tuple with the MysqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlPort

`func (o *GatewayUpdateProducerMySQL) SetMysqlPort(v string)`

SetMysqlPort sets MysqlPort field to given value.

### HasMysqlPort

`func (o *GatewayUpdateProducerMySQL) HasMysqlPort() bool`

HasMysqlPort returns a boolean if a field has been set.

### GetMysqlScreationStatements

`func (o *GatewayUpdateProducerMySQL) GetMysqlScreationStatements() string`

GetMysqlScreationStatements returns the MysqlScreationStatements field if non-nil, zero value otherwise.

### GetMysqlScreationStatementsOk

`func (o *GatewayUpdateProducerMySQL) GetMysqlScreationStatementsOk() (*string, bool)`

GetMysqlScreationStatementsOk returns a tuple with the MysqlScreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlScreationStatements

`func (o *GatewayUpdateProducerMySQL) SetMysqlScreationStatements(v string)`

SetMysqlScreationStatements sets MysqlScreationStatements field to given value.

### HasMysqlScreationStatements

`func (o *GatewayUpdateProducerMySQL) HasMysqlScreationStatements() bool`

HasMysqlScreationStatements returns a boolean if a field has been set.

### GetMysqlUsername

`func (o *GatewayUpdateProducerMySQL) GetMysqlUsername() string`

GetMysqlUsername returns the MysqlUsername field if non-nil, zero value otherwise.

### GetMysqlUsernameOk

`func (o *GatewayUpdateProducerMySQL) GetMysqlUsernameOk() (*string, bool)`

GetMysqlUsernameOk returns a tuple with the MysqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlUsername

`func (o *GatewayUpdateProducerMySQL) SetMysqlUsername(v string)`

SetMysqlUsername sets MysqlUsername field to given value.

### HasMysqlUsername

`func (o *GatewayUpdateProducerMySQL) HasMysqlUsername() bool`

HasMysqlUsername returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerMySQL) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerMySQL) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerMySQL) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerMySQL) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerMySQL) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerMySQL) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerMySQL) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *GatewayUpdateProducerMySQL) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayUpdateProducerMySQL) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayUpdateProducerMySQL) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayUpdateProducerMySQL) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerMySQL) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerMySQL) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerMySQL) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerMySQL) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerMySQL) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayUpdateProducerMySQL) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerMySQL) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerMySQL) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayUpdateProducerMySQL) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayUpdateProducerMySQL) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayUpdateProducerMySQL) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayUpdateProducerMySQL) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayUpdateProducerMySQL) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayUpdateProducerMySQL) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayUpdateProducerMySQL) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayUpdateProducerMySQL) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayUpdateProducerMySQL) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayUpdateProducerMySQL) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayUpdateProducerMySQL) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayUpdateProducerMySQL) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerMySQL) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerMySQL) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerMySQL) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerMySQL) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerMySQL) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerMySQL) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerMySQL) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerMySQL) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerMySQL) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerMySQL) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerMySQL) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerMySQL) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerMySQL) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerMySQL) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerMySQL) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerMySQL) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayUpdateProducerMySQL) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayUpdateProducerMySQL) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayUpdateProducerMySQL) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayUpdateProducerMySQL) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


