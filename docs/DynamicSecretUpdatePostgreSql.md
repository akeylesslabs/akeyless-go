# DynamicSecretUpdatePostgreSql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationStatements** | Pointer to **string** | PostgreSQL Creation statements | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
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
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts) | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Ssl** | Pointer to **bool** | Enable/Disable SSL [true/false] | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretUpdatePostgreSql

`func NewDynamicSecretUpdatePostgreSql(name string, ) *DynamicSecretUpdatePostgreSql`

NewDynamicSecretUpdatePostgreSql instantiates a new DynamicSecretUpdatePostgreSql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdatePostgreSqlWithDefaults

`func NewDynamicSecretUpdatePostgreSqlWithDefaults() *DynamicSecretUpdatePostgreSql`

NewDynamicSecretUpdatePostgreSqlWithDefaults instantiates a new DynamicSecretUpdatePostgreSql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationStatements

`func (o *DynamicSecretUpdatePostgreSql) GetCreationStatements() string`

GetCreationStatements returns the CreationStatements field if non-nil, zero value otherwise.

### GetCreationStatementsOk

`func (o *DynamicSecretUpdatePostgreSql) GetCreationStatementsOk() (*string, bool)`

GetCreationStatementsOk returns a tuple with the CreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStatements

`func (o *DynamicSecretUpdatePostgreSql) SetCreationStatements(v string)`

SetCreationStatements sets CreationStatements field to given value.

### HasCreationStatements

`func (o *DynamicSecretUpdatePostgreSql) HasCreationStatements() bool`

HasCreationStatements returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretUpdatePostgreSql) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdatePostgreSql) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdatePostgreSql) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdatePostgreSql) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdatePostgreSql) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdatePostgreSql) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdatePostgreSql) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdatePostgreSql) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdatePostgreSql) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdatePostgreSql) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdatePostgreSql) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdatePostgreSql) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdatePostgreSql) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdatePostgreSql) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdatePostgreSql) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdatePostgreSql) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdatePostgreSql) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdatePostgreSql) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdatePostgreSql) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretUpdatePostgreSql) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretUpdatePostgreSql) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretUpdatePostgreSql) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretUpdatePostgreSql) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetPostgresqlDbName

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlDbName() string`

GetPostgresqlDbName returns the PostgresqlDbName field if non-nil, zero value otherwise.

### GetPostgresqlDbNameOk

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlDbNameOk() (*string, bool)`

GetPostgresqlDbNameOk returns a tuple with the PostgresqlDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlDbName

`func (o *DynamicSecretUpdatePostgreSql) SetPostgresqlDbName(v string)`

SetPostgresqlDbName sets PostgresqlDbName field to given value.

### HasPostgresqlDbName

`func (o *DynamicSecretUpdatePostgreSql) HasPostgresqlDbName() bool`

HasPostgresqlDbName returns a boolean if a field has been set.

### GetPostgresqlHost

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlHost() string`

GetPostgresqlHost returns the PostgresqlHost field if non-nil, zero value otherwise.

### GetPostgresqlHostOk

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlHostOk() (*string, bool)`

GetPostgresqlHostOk returns a tuple with the PostgresqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlHost

`func (o *DynamicSecretUpdatePostgreSql) SetPostgresqlHost(v string)`

SetPostgresqlHost sets PostgresqlHost field to given value.

### HasPostgresqlHost

`func (o *DynamicSecretUpdatePostgreSql) HasPostgresqlHost() bool`

HasPostgresqlHost returns a boolean if a field has been set.

### GetPostgresqlPassword

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlPassword() string`

GetPostgresqlPassword returns the PostgresqlPassword field if non-nil, zero value otherwise.

### GetPostgresqlPasswordOk

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlPasswordOk() (*string, bool)`

GetPostgresqlPasswordOk returns a tuple with the PostgresqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlPassword

`func (o *DynamicSecretUpdatePostgreSql) SetPostgresqlPassword(v string)`

SetPostgresqlPassword sets PostgresqlPassword field to given value.

### HasPostgresqlPassword

`func (o *DynamicSecretUpdatePostgreSql) HasPostgresqlPassword() bool`

HasPostgresqlPassword returns a boolean if a field has been set.

### GetPostgresqlPort

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlPort() string`

GetPostgresqlPort returns the PostgresqlPort field if non-nil, zero value otherwise.

### GetPostgresqlPortOk

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlPortOk() (*string, bool)`

GetPostgresqlPortOk returns a tuple with the PostgresqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlPort

`func (o *DynamicSecretUpdatePostgreSql) SetPostgresqlPort(v string)`

SetPostgresqlPort sets PostgresqlPort field to given value.

### HasPostgresqlPort

`func (o *DynamicSecretUpdatePostgreSql) HasPostgresqlPort() bool`

HasPostgresqlPort returns a boolean if a field has been set.

### GetPostgresqlUsername

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlUsername() string`

GetPostgresqlUsername returns the PostgresqlUsername field if non-nil, zero value otherwise.

### GetPostgresqlUsernameOk

`func (o *DynamicSecretUpdatePostgreSql) GetPostgresqlUsernameOk() (*string, bool)`

GetPostgresqlUsernameOk returns a tuple with the PostgresqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlUsername

`func (o *DynamicSecretUpdatePostgreSql) SetPostgresqlUsername(v string)`

SetPostgresqlUsername sets PostgresqlUsername field to given value.

### HasPostgresqlUsername

`func (o *DynamicSecretUpdatePostgreSql) HasPostgresqlUsername() bool`

HasPostgresqlUsername returns a boolean if a field has been set.

### GetProducerEncryptionKey

`func (o *DynamicSecretUpdatePostgreSql) GetProducerEncryptionKey() string`

GetProducerEncryptionKey returns the ProducerEncryptionKey field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyOk

`func (o *DynamicSecretUpdatePostgreSql) GetProducerEncryptionKeyOk() (*string, bool)`

GetProducerEncryptionKeyOk returns a tuple with the ProducerEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKey

`func (o *DynamicSecretUpdatePostgreSql) SetProducerEncryptionKey(v string)`

SetProducerEncryptionKey sets ProducerEncryptionKey field to given value.

### HasProducerEncryptionKey

`func (o *DynamicSecretUpdatePostgreSql) HasProducerEncryptionKey() bool`

HasProducerEncryptionKey returns a boolean if a field has been set.

### GetRevocationStatement

`func (o *DynamicSecretUpdatePostgreSql) GetRevocationStatement() string`

GetRevocationStatement returns the RevocationStatement field if non-nil, zero value otherwise.

### GetRevocationStatementOk

`func (o *DynamicSecretUpdatePostgreSql) GetRevocationStatementOk() (*string, bool)`

GetRevocationStatementOk returns a tuple with the RevocationStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationStatement

`func (o *DynamicSecretUpdatePostgreSql) SetRevocationStatement(v string)`

SetRevocationStatement sets RevocationStatement field to given value.

### HasRevocationStatement

`func (o *DynamicSecretUpdatePostgreSql) HasRevocationStatement() bool`

HasRevocationStatement returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdatePostgreSql) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretUpdatePostgreSql) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessDbSchema

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessDbSchema() string`

GetSecureAccessDbSchema returns the SecureAccessDbSchema field if non-nil, zero value otherwise.

### GetSecureAccessDbSchemaOk

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessDbSchemaOk() (*string, bool)`

GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbSchema

`func (o *DynamicSecretUpdatePostgreSql) SetSecureAccessDbSchema(v string)`

SetSecureAccessDbSchema sets SecureAccessDbSchema field to given value.

### HasSecureAccessDbSchema

`func (o *DynamicSecretUpdatePostgreSql) HasSecureAccessDbSchema() bool`

HasSecureAccessDbSchema returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdatePostgreSql) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdatePostgreSql) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretUpdatePostgreSql) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretUpdatePostgreSql) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretUpdatePostgreSql) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretUpdatePostgreSql) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretUpdatePostgreSql) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSsl

`func (o *DynamicSecretUpdatePostgreSql) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *DynamicSecretUpdatePostgreSql) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *DynamicSecretUpdatePostgreSql) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *DynamicSecretUpdatePostgreSql) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdatePostgreSql) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdatePostgreSql) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdatePostgreSql) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdatePostgreSql) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdatePostgreSql) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdatePostgreSql) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdatePostgreSql) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdatePostgreSql) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdatePostgreSql) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdatePostgreSql) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdatePostgreSql) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdatePostgreSql) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdatePostgreSql) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdatePostgreSql) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdatePostgreSql) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdatePostgreSql) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdatePostgreSql) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdatePostgreSql) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdatePostgreSql) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdatePostgreSql) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


