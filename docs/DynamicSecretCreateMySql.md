# DynamicSecretCreateMySql

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
**MysqlRevocationStatements** | Pointer to **string** | MySQL Revocation statements | [optional] 
**MysqlScreationStatements** | Pointer to **string** | MySQL Creation statements | [optional] 
**MysqlUsername** | Pointer to **string** | MySQL Username | [optional] 
**Name** | **string** | Dynamic secret name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts) | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Ssl** | Pointer to **bool** | Enable/Disable SSL [true/false] | [optional] [default to false]
**SslCertificate** | Pointer to **string** | SSL connection certificate | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateMySql

`func NewDynamicSecretCreateMySql(name string, ) *DynamicSecretCreateMySql`

NewDynamicSecretCreateMySql instantiates a new DynamicSecretCreateMySql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateMySqlWithDefaults

`func NewDynamicSecretCreateMySqlWithDefaults() *DynamicSecretCreateMySql`

NewDynamicSecretCreateMySqlWithDefaults instantiates a new DynamicSecretCreateMySql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbServerCertificates

`func (o *DynamicSecretCreateMySql) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *DynamicSecretCreateMySql) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *DynamicSecretCreateMySql) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *DynamicSecretCreateMySql) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *DynamicSecretCreateMySql) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *DynamicSecretCreateMySql) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *DynamicSecretCreateMySql) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *DynamicSecretCreateMySql) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateMySql) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateMySql) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateMySql) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateMySql) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateMySql) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateMySql) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateMySql) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateMySql) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMysqlDbname

`func (o *DynamicSecretCreateMySql) GetMysqlDbname() string`

GetMysqlDbname returns the MysqlDbname field if non-nil, zero value otherwise.

### GetMysqlDbnameOk

`func (o *DynamicSecretCreateMySql) GetMysqlDbnameOk() (*string, bool)`

GetMysqlDbnameOk returns a tuple with the MysqlDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlDbname

`func (o *DynamicSecretCreateMySql) SetMysqlDbname(v string)`

SetMysqlDbname sets MysqlDbname field to given value.

### HasMysqlDbname

`func (o *DynamicSecretCreateMySql) HasMysqlDbname() bool`

HasMysqlDbname returns a boolean if a field has been set.

### GetMysqlHost

`func (o *DynamicSecretCreateMySql) GetMysqlHost() string`

GetMysqlHost returns the MysqlHost field if non-nil, zero value otherwise.

### GetMysqlHostOk

`func (o *DynamicSecretCreateMySql) GetMysqlHostOk() (*string, bool)`

GetMysqlHostOk returns a tuple with the MysqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlHost

`func (o *DynamicSecretCreateMySql) SetMysqlHost(v string)`

SetMysqlHost sets MysqlHost field to given value.

### HasMysqlHost

`func (o *DynamicSecretCreateMySql) HasMysqlHost() bool`

HasMysqlHost returns a boolean if a field has been set.

### GetMysqlPassword

`func (o *DynamicSecretCreateMySql) GetMysqlPassword() string`

GetMysqlPassword returns the MysqlPassword field if non-nil, zero value otherwise.

### GetMysqlPasswordOk

`func (o *DynamicSecretCreateMySql) GetMysqlPasswordOk() (*string, bool)`

GetMysqlPasswordOk returns a tuple with the MysqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlPassword

`func (o *DynamicSecretCreateMySql) SetMysqlPassword(v string)`

SetMysqlPassword sets MysqlPassword field to given value.

### HasMysqlPassword

`func (o *DynamicSecretCreateMySql) HasMysqlPassword() bool`

HasMysqlPassword returns a boolean if a field has been set.

### GetMysqlPort

`func (o *DynamicSecretCreateMySql) GetMysqlPort() string`

GetMysqlPort returns the MysqlPort field if non-nil, zero value otherwise.

### GetMysqlPortOk

`func (o *DynamicSecretCreateMySql) GetMysqlPortOk() (*string, bool)`

GetMysqlPortOk returns a tuple with the MysqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlPort

`func (o *DynamicSecretCreateMySql) SetMysqlPort(v string)`

SetMysqlPort sets MysqlPort field to given value.

### HasMysqlPort

`func (o *DynamicSecretCreateMySql) HasMysqlPort() bool`

HasMysqlPort returns a boolean if a field has been set.

### GetMysqlRevocationStatements

`func (o *DynamicSecretCreateMySql) GetMysqlRevocationStatements() string`

GetMysqlRevocationStatements returns the MysqlRevocationStatements field if non-nil, zero value otherwise.

### GetMysqlRevocationStatementsOk

`func (o *DynamicSecretCreateMySql) GetMysqlRevocationStatementsOk() (*string, bool)`

GetMysqlRevocationStatementsOk returns a tuple with the MysqlRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlRevocationStatements

`func (o *DynamicSecretCreateMySql) SetMysqlRevocationStatements(v string)`

SetMysqlRevocationStatements sets MysqlRevocationStatements field to given value.

### HasMysqlRevocationStatements

`func (o *DynamicSecretCreateMySql) HasMysqlRevocationStatements() bool`

HasMysqlRevocationStatements returns a boolean if a field has been set.

### GetMysqlScreationStatements

`func (o *DynamicSecretCreateMySql) GetMysqlScreationStatements() string`

GetMysqlScreationStatements returns the MysqlScreationStatements field if non-nil, zero value otherwise.

### GetMysqlScreationStatementsOk

`func (o *DynamicSecretCreateMySql) GetMysqlScreationStatementsOk() (*string, bool)`

GetMysqlScreationStatementsOk returns a tuple with the MysqlScreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlScreationStatements

`func (o *DynamicSecretCreateMySql) SetMysqlScreationStatements(v string)`

SetMysqlScreationStatements sets MysqlScreationStatements field to given value.

### HasMysqlScreationStatements

`func (o *DynamicSecretCreateMySql) HasMysqlScreationStatements() bool`

HasMysqlScreationStatements returns a boolean if a field has been set.

### GetMysqlUsername

`func (o *DynamicSecretCreateMySql) GetMysqlUsername() string`

GetMysqlUsername returns the MysqlUsername field if non-nil, zero value otherwise.

### GetMysqlUsernameOk

`func (o *DynamicSecretCreateMySql) GetMysqlUsernameOk() (*string, bool)`

GetMysqlUsernameOk returns a tuple with the MysqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlUsername

`func (o *DynamicSecretCreateMySql) SetMysqlUsername(v string)`

SetMysqlUsername sets MysqlUsername field to given value.

### HasMysqlUsername

`func (o *DynamicSecretCreateMySql) HasMysqlUsername() bool`

HasMysqlUsername returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateMySql) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateMySql) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateMySql) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateMySql) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateMySql) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateMySql) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateMySql) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateMySql) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretCreateMySql) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateMySql) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretCreateMySql) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateMySql) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateMySql) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateMySql) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateMySql) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretCreateMySql) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretCreateMySql) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretCreateMySql) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretCreateMySql) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretCreateMySql) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretCreateMySql) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretCreateMySql) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretCreateMySql) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSsl

`func (o *DynamicSecretCreateMySql) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *DynamicSecretCreateMySql) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *DynamicSecretCreateMySql) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *DynamicSecretCreateMySql) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *DynamicSecretCreateMySql) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *DynamicSecretCreateMySql) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *DynamicSecretCreateMySql) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *DynamicSecretCreateMySql) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateMySql) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateMySql) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateMySql) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateMySql) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateMySql) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateMySql) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateMySql) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateMySql) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateMySql) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateMySql) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateMySql) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateMySql) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateMySql) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateMySql) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateMySql) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateMySql) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateMySql) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateMySql) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateMySql) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateMySql) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


