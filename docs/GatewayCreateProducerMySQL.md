# GatewayCreateProducerMySQL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbServerCertificates** | Pointer to **string** | (Optional) DB server certificates | [optional] 
**DbServerName** | Pointer to **string** | (Optional) Server name for certificate verification | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MysqlDbname** | Pointer to **string** | MySQL DB Name | [optional] 
**MysqlHost** | Pointer to **string** | MySQL Host | [optional] [default to "127.0.0.1"]
**MysqlPassword** | Pointer to **string** | MySQL Password | [optional] 
**MysqlPort** | Pointer to **string** | MySQL Port | [optional] [default to "3306"]
**MysqlScreationStatements** | Pointer to **string** | MySQL Creation statements | [optional] 
**MysqlUsername** | Pointer to **string** | MySQL Username | [optional] 
**Name** | **string** | Producer name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Ssl** | Pointer to **bool** | Enable/Disable SSL [true/false] | [optional] [default to false]
**SslCertificate** | Pointer to **string** | SSL connection certificate | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerMySQL

`func NewGatewayCreateProducerMySQL(name string, ) *GatewayCreateProducerMySQL`

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

### GetDeleteProtection

`func (o *GatewayCreateProducerMySQL) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayCreateProducerMySQL) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayCreateProducerMySQL) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayCreateProducerMySQL) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateProducerMySQL) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateProducerMySQL) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateProducerMySQL) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateProducerMySQL) HasJson() bool`

HasJson returns a boolean if a field has been set.

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

### HasMysqlDbname

`func (o *GatewayCreateProducerMySQL) HasMysqlDbname() bool`

HasMysqlDbname returns a boolean if a field has been set.

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

### HasMysqlPassword

`func (o *GatewayCreateProducerMySQL) HasMysqlPassword() bool`

HasMysqlPassword returns a boolean if a field has been set.

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

### HasMysqlUsername

`func (o *GatewayCreateProducerMySQL) HasMysqlUsername() bool`

HasMysqlUsername returns a boolean if a field has been set.

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

### GetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerMySQL) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayCreateProducerMySQL) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerMySQL) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayCreateProducerMySQL) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayCreateProducerMySQL) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayCreateProducerMySQL) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayCreateProducerMySQL) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayCreateProducerMySQL) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayCreateProducerMySQL) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayCreateProducerMySQL) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayCreateProducerMySQL) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayCreateProducerMySQL) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayCreateProducerMySQL) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayCreateProducerMySQL) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayCreateProducerMySQL) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayCreateProducerMySQL) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSsl

`func (o *GatewayCreateProducerMySQL) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *GatewayCreateProducerMySQL) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *GatewayCreateProducerMySQL) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *GatewayCreateProducerMySQL) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *GatewayCreateProducerMySQL) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *GatewayCreateProducerMySQL) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *GatewayCreateProducerMySQL) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *GatewayCreateProducerMySQL) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerMySQL) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerMySQL) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerMySQL) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerMySQL) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerMySQL) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerMySQL) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerMySQL) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerMySQL) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


