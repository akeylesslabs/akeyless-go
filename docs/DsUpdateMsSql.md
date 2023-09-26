# DsUpdateMsSql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
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

### NewDsUpdateMsSql

`func NewDsUpdateMsSql(name string, ) *DsUpdateMsSql`

NewDsUpdateMsSql instantiates a new DsUpdateMsSql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsUpdateMsSqlWithDefaults

`func NewDsUpdateMsSqlWithDefaults() *DsUpdateMsSql`

NewDsUpdateMsSqlWithDefaults instantiates a new DsUpdateMsSql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DsUpdateMsSql) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DsUpdateMsSql) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DsUpdateMsSql) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DsUpdateMsSql) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DsUpdateMsSql) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DsUpdateMsSql) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DsUpdateMsSql) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DsUpdateMsSql) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMssqlCreateStatements

`func (o *DsUpdateMsSql) GetMssqlCreateStatements() string`

GetMssqlCreateStatements returns the MssqlCreateStatements field if non-nil, zero value otherwise.

### GetMssqlCreateStatementsOk

`func (o *DsUpdateMsSql) GetMssqlCreateStatementsOk() (*string, bool)`

GetMssqlCreateStatementsOk returns a tuple with the MssqlCreateStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlCreateStatements

`func (o *DsUpdateMsSql) SetMssqlCreateStatements(v string)`

SetMssqlCreateStatements sets MssqlCreateStatements field to given value.

### HasMssqlCreateStatements

`func (o *DsUpdateMsSql) HasMssqlCreateStatements() bool`

HasMssqlCreateStatements returns a boolean if a field has been set.

### GetMssqlDbname

`func (o *DsUpdateMsSql) GetMssqlDbname() string`

GetMssqlDbname returns the MssqlDbname field if non-nil, zero value otherwise.

### GetMssqlDbnameOk

`func (o *DsUpdateMsSql) GetMssqlDbnameOk() (*string, bool)`

GetMssqlDbnameOk returns a tuple with the MssqlDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlDbname

`func (o *DsUpdateMsSql) SetMssqlDbname(v string)`

SetMssqlDbname sets MssqlDbname field to given value.

### HasMssqlDbname

`func (o *DsUpdateMsSql) HasMssqlDbname() bool`

HasMssqlDbname returns a boolean if a field has been set.

### GetMssqlHost

`func (o *DsUpdateMsSql) GetMssqlHost() string`

GetMssqlHost returns the MssqlHost field if non-nil, zero value otherwise.

### GetMssqlHostOk

`func (o *DsUpdateMsSql) GetMssqlHostOk() (*string, bool)`

GetMssqlHostOk returns a tuple with the MssqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlHost

`func (o *DsUpdateMsSql) SetMssqlHost(v string)`

SetMssqlHost sets MssqlHost field to given value.

### HasMssqlHost

`func (o *DsUpdateMsSql) HasMssqlHost() bool`

HasMssqlHost returns a boolean if a field has been set.

### GetMssqlPassword

`func (o *DsUpdateMsSql) GetMssqlPassword() string`

GetMssqlPassword returns the MssqlPassword field if non-nil, zero value otherwise.

### GetMssqlPasswordOk

`func (o *DsUpdateMsSql) GetMssqlPasswordOk() (*string, bool)`

GetMssqlPasswordOk returns a tuple with the MssqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlPassword

`func (o *DsUpdateMsSql) SetMssqlPassword(v string)`

SetMssqlPassword sets MssqlPassword field to given value.

### HasMssqlPassword

`func (o *DsUpdateMsSql) HasMssqlPassword() bool`

HasMssqlPassword returns a boolean if a field has been set.

### GetMssqlPort

`func (o *DsUpdateMsSql) GetMssqlPort() string`

GetMssqlPort returns the MssqlPort field if non-nil, zero value otherwise.

### GetMssqlPortOk

`func (o *DsUpdateMsSql) GetMssqlPortOk() (*string, bool)`

GetMssqlPortOk returns a tuple with the MssqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlPort

`func (o *DsUpdateMsSql) SetMssqlPort(v string)`

SetMssqlPort sets MssqlPort field to given value.

### HasMssqlPort

`func (o *DsUpdateMsSql) HasMssqlPort() bool`

HasMssqlPort returns a boolean if a field has been set.

### GetMssqlRevocationStatements

`func (o *DsUpdateMsSql) GetMssqlRevocationStatements() string`

GetMssqlRevocationStatements returns the MssqlRevocationStatements field if non-nil, zero value otherwise.

### GetMssqlRevocationStatementsOk

`func (o *DsUpdateMsSql) GetMssqlRevocationStatementsOk() (*string, bool)`

GetMssqlRevocationStatementsOk returns a tuple with the MssqlRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlRevocationStatements

`func (o *DsUpdateMsSql) SetMssqlRevocationStatements(v string)`

SetMssqlRevocationStatements sets MssqlRevocationStatements field to given value.

### HasMssqlRevocationStatements

`func (o *DsUpdateMsSql) HasMssqlRevocationStatements() bool`

HasMssqlRevocationStatements returns a boolean if a field has been set.

### GetMssqlUsername

`func (o *DsUpdateMsSql) GetMssqlUsername() string`

GetMssqlUsername returns the MssqlUsername field if non-nil, zero value otherwise.

### GetMssqlUsernameOk

`func (o *DsUpdateMsSql) GetMssqlUsernameOk() (*string, bool)`

GetMssqlUsernameOk returns a tuple with the MssqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlUsername

`func (o *DsUpdateMsSql) SetMssqlUsername(v string)`

SetMssqlUsername sets MssqlUsername field to given value.

### HasMssqlUsername

`func (o *DsUpdateMsSql) HasMssqlUsername() bool`

HasMssqlUsername returns a boolean if a field has been set.

### GetName

`func (o *DsUpdateMsSql) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DsUpdateMsSql) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DsUpdateMsSql) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DsUpdateMsSql) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DsUpdateMsSql) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DsUpdateMsSql) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DsUpdateMsSql) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DsUpdateMsSql) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DsUpdateMsSql) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DsUpdateMsSql) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DsUpdateMsSql) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DsUpdateMsSql) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DsUpdateMsSql) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DsUpdateMsSql) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DsUpdateMsSql) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessDbSchema

`func (o *DsUpdateMsSql) GetSecureAccessDbSchema() string`

GetSecureAccessDbSchema returns the SecureAccessDbSchema field if non-nil, zero value otherwise.

### GetSecureAccessDbSchemaOk

`func (o *DsUpdateMsSql) GetSecureAccessDbSchemaOk() (*string, bool)`

GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbSchema

`func (o *DsUpdateMsSql) SetSecureAccessDbSchema(v string)`

SetSecureAccessDbSchema sets SecureAccessDbSchema field to given value.

### HasSecureAccessDbSchema

`func (o *DsUpdateMsSql) HasSecureAccessDbSchema() bool`

HasSecureAccessDbSchema returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DsUpdateMsSql) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DsUpdateMsSql) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DsUpdateMsSql) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DsUpdateMsSql) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DsUpdateMsSql) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DsUpdateMsSql) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DsUpdateMsSql) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DsUpdateMsSql) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DsUpdateMsSql) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DsUpdateMsSql) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DsUpdateMsSql) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DsUpdateMsSql) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DsUpdateMsSql) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DsUpdateMsSql) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DsUpdateMsSql) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DsUpdateMsSql) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DsUpdateMsSql) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DsUpdateMsSql) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DsUpdateMsSql) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DsUpdateMsSql) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DsUpdateMsSql) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DsUpdateMsSql) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DsUpdateMsSql) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DsUpdateMsSql) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DsUpdateMsSql) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DsUpdateMsSql) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DsUpdateMsSql) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DsUpdateMsSql) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DsUpdateMsSql) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DsUpdateMsSql) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DsUpdateMsSql) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DsUpdateMsSql) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


