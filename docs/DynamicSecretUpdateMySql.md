# DynamicSecretUpdateMySql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbServerCertificates** | Pointer to **string** | (Optional) DB server certificates | [optional] 
**DbServerName** | Pointer to **string** | (Optional) Server name for certificate verification | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MysqlDbname** | Pointer to **string** | MySQL DB Name | [optional] 
**MysqlHost** | Pointer to **string** | MySQL Host | [optional] [default to "127.0.0.1"]
**MysqlPassword** | Pointer to **string** | MySQL Password | [optional] 
**MysqlPort** | Pointer to **string** | MySQL Port | [optional] [default to "3306"]
**MysqlRevocationStatements** | Pointer to **string** | MySQL Revocation statements | [optional] 
**MysqlScreationStatements** | Pointer to **string** | MySQL Creation statements | [optional] 
**MysqlUsername** | Pointer to **string** | MySQL Username | [optional] 
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
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

### NewDynamicSecretUpdateMySql

`func NewDynamicSecretUpdateMySql(name string, ) *DynamicSecretUpdateMySql`

NewDynamicSecretUpdateMySql instantiates a new DynamicSecretUpdateMySql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateMySqlWithDefaults

`func NewDynamicSecretUpdateMySqlWithDefaults() *DynamicSecretUpdateMySql`

NewDynamicSecretUpdateMySqlWithDefaults instantiates a new DynamicSecretUpdateMySql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbServerCertificates

`func (o *DynamicSecretUpdateMySql) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *DynamicSecretUpdateMySql) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *DynamicSecretUpdateMySql) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *DynamicSecretUpdateMySql) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *DynamicSecretUpdateMySql) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *DynamicSecretUpdateMySql) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *DynamicSecretUpdateMySql) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *DynamicSecretUpdateMySql) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretUpdateMySql) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateMySql) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateMySql) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateMySql) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateMySql) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateMySql) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateMySql) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateMySql) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateMySql) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateMySql) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateMySql) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateMySql) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMysqlDbname

`func (o *DynamicSecretUpdateMySql) GetMysqlDbname() string`

GetMysqlDbname returns the MysqlDbname field if non-nil, zero value otherwise.

### GetMysqlDbnameOk

`func (o *DynamicSecretUpdateMySql) GetMysqlDbnameOk() (*string, bool)`

GetMysqlDbnameOk returns a tuple with the MysqlDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlDbname

`func (o *DynamicSecretUpdateMySql) SetMysqlDbname(v string)`

SetMysqlDbname sets MysqlDbname field to given value.

### HasMysqlDbname

`func (o *DynamicSecretUpdateMySql) HasMysqlDbname() bool`

HasMysqlDbname returns a boolean if a field has been set.

### GetMysqlHost

`func (o *DynamicSecretUpdateMySql) GetMysqlHost() string`

GetMysqlHost returns the MysqlHost field if non-nil, zero value otherwise.

### GetMysqlHostOk

`func (o *DynamicSecretUpdateMySql) GetMysqlHostOk() (*string, bool)`

GetMysqlHostOk returns a tuple with the MysqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlHost

`func (o *DynamicSecretUpdateMySql) SetMysqlHost(v string)`

SetMysqlHost sets MysqlHost field to given value.

### HasMysqlHost

`func (o *DynamicSecretUpdateMySql) HasMysqlHost() bool`

HasMysqlHost returns a boolean if a field has been set.

### GetMysqlPassword

`func (o *DynamicSecretUpdateMySql) GetMysqlPassword() string`

GetMysqlPassword returns the MysqlPassword field if non-nil, zero value otherwise.

### GetMysqlPasswordOk

`func (o *DynamicSecretUpdateMySql) GetMysqlPasswordOk() (*string, bool)`

GetMysqlPasswordOk returns a tuple with the MysqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlPassword

`func (o *DynamicSecretUpdateMySql) SetMysqlPassword(v string)`

SetMysqlPassword sets MysqlPassword field to given value.

### HasMysqlPassword

`func (o *DynamicSecretUpdateMySql) HasMysqlPassword() bool`

HasMysqlPassword returns a boolean if a field has been set.

### GetMysqlPort

`func (o *DynamicSecretUpdateMySql) GetMysqlPort() string`

GetMysqlPort returns the MysqlPort field if non-nil, zero value otherwise.

### GetMysqlPortOk

`func (o *DynamicSecretUpdateMySql) GetMysqlPortOk() (*string, bool)`

GetMysqlPortOk returns a tuple with the MysqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlPort

`func (o *DynamicSecretUpdateMySql) SetMysqlPort(v string)`

SetMysqlPort sets MysqlPort field to given value.

### HasMysqlPort

`func (o *DynamicSecretUpdateMySql) HasMysqlPort() bool`

HasMysqlPort returns a boolean if a field has been set.

### GetMysqlRevocationStatements

`func (o *DynamicSecretUpdateMySql) GetMysqlRevocationStatements() string`

GetMysqlRevocationStatements returns the MysqlRevocationStatements field if non-nil, zero value otherwise.

### GetMysqlRevocationStatementsOk

`func (o *DynamicSecretUpdateMySql) GetMysqlRevocationStatementsOk() (*string, bool)`

GetMysqlRevocationStatementsOk returns a tuple with the MysqlRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlRevocationStatements

`func (o *DynamicSecretUpdateMySql) SetMysqlRevocationStatements(v string)`

SetMysqlRevocationStatements sets MysqlRevocationStatements field to given value.

### HasMysqlRevocationStatements

`func (o *DynamicSecretUpdateMySql) HasMysqlRevocationStatements() bool`

HasMysqlRevocationStatements returns a boolean if a field has been set.

### GetMysqlScreationStatements

`func (o *DynamicSecretUpdateMySql) GetMysqlScreationStatements() string`

GetMysqlScreationStatements returns the MysqlScreationStatements field if non-nil, zero value otherwise.

### GetMysqlScreationStatementsOk

`func (o *DynamicSecretUpdateMySql) GetMysqlScreationStatementsOk() (*string, bool)`

GetMysqlScreationStatementsOk returns a tuple with the MysqlScreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlScreationStatements

`func (o *DynamicSecretUpdateMySql) SetMysqlScreationStatements(v string)`

SetMysqlScreationStatements sets MysqlScreationStatements field to given value.

### HasMysqlScreationStatements

`func (o *DynamicSecretUpdateMySql) HasMysqlScreationStatements() bool`

HasMysqlScreationStatements returns a boolean if a field has been set.

### GetMysqlUsername

`func (o *DynamicSecretUpdateMySql) GetMysqlUsername() string`

GetMysqlUsername returns the MysqlUsername field if non-nil, zero value otherwise.

### GetMysqlUsernameOk

`func (o *DynamicSecretUpdateMySql) GetMysqlUsernameOk() (*string, bool)`

GetMysqlUsernameOk returns a tuple with the MysqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlUsername

`func (o *DynamicSecretUpdateMySql) SetMysqlUsername(v string)`

SetMysqlUsername sets MysqlUsername field to given value.

### HasMysqlUsername

`func (o *DynamicSecretUpdateMySql) HasMysqlUsername() bool`

HasMysqlUsername returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateMySql) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateMySql) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateMySql) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateMySql) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateMySql) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateMySql) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateMySql) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretUpdateMySql) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretUpdateMySql) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretUpdateMySql) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretUpdateMySql) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateMySql) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateMySql) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateMySql) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateMySql) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateMySql) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretUpdateMySql) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateMySql) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateMySql) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateMySql) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateMySql) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateMySql) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateMySql) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretUpdateMySql) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretUpdateMySql) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretUpdateMySql) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretUpdateMySql) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretUpdateMySql) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretUpdateMySql) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretUpdateMySql) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretUpdateMySql) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSsl

`func (o *DynamicSecretUpdateMySql) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *DynamicSecretUpdateMySql) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *DynamicSecretUpdateMySql) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *DynamicSecretUpdateMySql) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetSslCertificate

`func (o *DynamicSecretUpdateMySql) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *DynamicSecretUpdateMySql) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *DynamicSecretUpdateMySql) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *DynamicSecretUpdateMySql) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateMySql) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateMySql) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateMySql) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateMySql) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateMySql) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateMySql) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateMySql) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateMySql) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateMySql) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateMySql) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateMySql) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateMySql) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateMySql) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateMySql) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateMySql) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateMySql) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateMySql) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateMySql) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateMySql) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateMySql) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


