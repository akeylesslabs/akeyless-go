# DynamicSecretUpdateMsSql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MssqlCreateStatements** | Pointer to **string** | MSSQL Creation statements | [optional] 
**MssqlDbname** | Pointer to **string** | MSSQL Name | [optional] 
**MssqlHost** | Pointer to **string** | MSSQL Host | [optional] [default to "127.0.0.1"]
**MssqlPassword** | Pointer to **string** | MSSQL Password | [optional] 
**MssqlPort** | Pointer to **string** | MSSQL Port | [optional] [default to "1433"]
**MssqlRevocationStatements** | Pointer to **string** | MSSQL Revocation statements | [optional] 
**MssqlUsername** | Pointer to **string** | MSSQL Username | [optional] 
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessDbSchema** | Pointer to **string** | The DB schema | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts) | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretUpdateMsSql

`func NewDynamicSecretUpdateMsSql(name string, ) *DynamicSecretUpdateMsSql`

NewDynamicSecretUpdateMsSql instantiates a new DynamicSecretUpdateMsSql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateMsSqlWithDefaults

`func NewDynamicSecretUpdateMsSqlWithDefaults() *DynamicSecretUpdateMsSql`

NewDynamicSecretUpdateMsSqlWithDefaults instantiates a new DynamicSecretUpdateMsSql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretUpdateMsSql) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateMsSql) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateMsSql) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateMsSql) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateMsSql) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateMsSql) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateMsSql) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateMsSql) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateMsSql) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateMsSql) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateMsSql) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateMsSql) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMssqlCreateStatements

`func (o *DynamicSecretUpdateMsSql) GetMssqlCreateStatements() string`

GetMssqlCreateStatements returns the MssqlCreateStatements field if non-nil, zero value otherwise.

### GetMssqlCreateStatementsOk

`func (o *DynamicSecretUpdateMsSql) GetMssqlCreateStatementsOk() (*string, bool)`

GetMssqlCreateStatementsOk returns a tuple with the MssqlCreateStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlCreateStatements

`func (o *DynamicSecretUpdateMsSql) SetMssqlCreateStatements(v string)`

SetMssqlCreateStatements sets MssqlCreateStatements field to given value.

### HasMssqlCreateStatements

`func (o *DynamicSecretUpdateMsSql) HasMssqlCreateStatements() bool`

HasMssqlCreateStatements returns a boolean if a field has been set.

### GetMssqlDbname

`func (o *DynamicSecretUpdateMsSql) GetMssqlDbname() string`

GetMssqlDbname returns the MssqlDbname field if non-nil, zero value otherwise.

### GetMssqlDbnameOk

`func (o *DynamicSecretUpdateMsSql) GetMssqlDbnameOk() (*string, bool)`

GetMssqlDbnameOk returns a tuple with the MssqlDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlDbname

`func (o *DynamicSecretUpdateMsSql) SetMssqlDbname(v string)`

SetMssqlDbname sets MssqlDbname field to given value.

### HasMssqlDbname

`func (o *DynamicSecretUpdateMsSql) HasMssqlDbname() bool`

HasMssqlDbname returns a boolean if a field has been set.

### GetMssqlHost

`func (o *DynamicSecretUpdateMsSql) GetMssqlHost() string`

GetMssqlHost returns the MssqlHost field if non-nil, zero value otherwise.

### GetMssqlHostOk

`func (o *DynamicSecretUpdateMsSql) GetMssqlHostOk() (*string, bool)`

GetMssqlHostOk returns a tuple with the MssqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlHost

`func (o *DynamicSecretUpdateMsSql) SetMssqlHost(v string)`

SetMssqlHost sets MssqlHost field to given value.

### HasMssqlHost

`func (o *DynamicSecretUpdateMsSql) HasMssqlHost() bool`

HasMssqlHost returns a boolean if a field has been set.

### GetMssqlPassword

`func (o *DynamicSecretUpdateMsSql) GetMssqlPassword() string`

GetMssqlPassword returns the MssqlPassword field if non-nil, zero value otherwise.

### GetMssqlPasswordOk

`func (o *DynamicSecretUpdateMsSql) GetMssqlPasswordOk() (*string, bool)`

GetMssqlPasswordOk returns a tuple with the MssqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlPassword

`func (o *DynamicSecretUpdateMsSql) SetMssqlPassword(v string)`

SetMssqlPassword sets MssqlPassword field to given value.

### HasMssqlPassword

`func (o *DynamicSecretUpdateMsSql) HasMssqlPassword() bool`

HasMssqlPassword returns a boolean if a field has been set.

### GetMssqlPort

`func (o *DynamicSecretUpdateMsSql) GetMssqlPort() string`

GetMssqlPort returns the MssqlPort field if non-nil, zero value otherwise.

### GetMssqlPortOk

`func (o *DynamicSecretUpdateMsSql) GetMssqlPortOk() (*string, bool)`

GetMssqlPortOk returns a tuple with the MssqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlPort

`func (o *DynamicSecretUpdateMsSql) SetMssqlPort(v string)`

SetMssqlPort sets MssqlPort field to given value.

### HasMssqlPort

`func (o *DynamicSecretUpdateMsSql) HasMssqlPort() bool`

HasMssqlPort returns a boolean if a field has been set.

### GetMssqlRevocationStatements

`func (o *DynamicSecretUpdateMsSql) GetMssqlRevocationStatements() string`

GetMssqlRevocationStatements returns the MssqlRevocationStatements field if non-nil, zero value otherwise.

### GetMssqlRevocationStatementsOk

`func (o *DynamicSecretUpdateMsSql) GetMssqlRevocationStatementsOk() (*string, bool)`

GetMssqlRevocationStatementsOk returns a tuple with the MssqlRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlRevocationStatements

`func (o *DynamicSecretUpdateMsSql) SetMssqlRevocationStatements(v string)`

SetMssqlRevocationStatements sets MssqlRevocationStatements field to given value.

### HasMssqlRevocationStatements

`func (o *DynamicSecretUpdateMsSql) HasMssqlRevocationStatements() bool`

HasMssqlRevocationStatements returns a boolean if a field has been set.

### GetMssqlUsername

`func (o *DynamicSecretUpdateMsSql) GetMssqlUsername() string`

GetMssqlUsername returns the MssqlUsername field if non-nil, zero value otherwise.

### GetMssqlUsernameOk

`func (o *DynamicSecretUpdateMsSql) GetMssqlUsernameOk() (*string, bool)`

GetMssqlUsernameOk returns a tuple with the MssqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlUsername

`func (o *DynamicSecretUpdateMsSql) SetMssqlUsername(v string)`

SetMssqlUsername sets MssqlUsername field to given value.

### HasMssqlUsername

`func (o *DynamicSecretUpdateMsSql) HasMssqlUsername() bool`

HasMssqlUsername returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateMsSql) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateMsSql) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateMsSql) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateMsSql) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateMsSql) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateMsSql) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateMsSql) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DynamicSecretUpdateMsSql) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretUpdateMsSql) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretUpdateMsSql) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretUpdateMsSql) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateMsSql) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateMsSql) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateMsSql) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateMsSql) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateMsSql) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateMsSql) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessDbSchema

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessDbSchema() string`

GetSecureAccessDbSchema returns the SecureAccessDbSchema field if non-nil, zero value otherwise.

### GetSecureAccessDbSchemaOk

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessDbSchemaOk() (*string, bool)`

GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbSchema

`func (o *DynamicSecretUpdateMsSql) SetSecureAccessDbSchema(v string)`

SetSecureAccessDbSchema sets SecureAccessDbSchema field to given value.

### HasSecureAccessDbSchema

`func (o *DynamicSecretUpdateMsSql) HasSecureAccessDbSchema() bool`

HasSecureAccessDbSchema returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateMsSql) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateMsSql) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretUpdateMsSql) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretUpdateMsSql) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretUpdateMsSql) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretUpdateMsSql) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretUpdateMsSql) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateMsSql) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateMsSql) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateMsSql) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateMsSql) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateMsSql) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateMsSql) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateMsSql) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateMsSql) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateMsSql) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateMsSql) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateMsSql) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateMsSql) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateMsSql) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateMsSql) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateMsSql) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateMsSql) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateMsSql) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateMsSql) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateMsSql) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateMsSql) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


