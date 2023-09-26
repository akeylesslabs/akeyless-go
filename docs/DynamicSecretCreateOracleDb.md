# DynamicSecretCreateOracleDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbServerCertificates** | Pointer to **string** | (Optional) DB server certificates | [optional] 
**DbServerName** | Pointer to **string** | (Optional) Server name for certificate verification | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**OracleHost** | Pointer to **string** | Oracle Host | [optional] [default to "127.0.0.1"]
**OraclePassword** | Pointer to **string** | Oracle Password | [optional] 
**OraclePort** | Pointer to **string** | Oracle Port | [optional] [default to "1521"]
**OracleScreationStatements** | Pointer to **string** | Oracle Creation statements | [optional] 
**OracleServiceName** | Pointer to **string** | Oracle DB Name | [optional] 
**OracleUsername** | Pointer to **string** | Oracle Username | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] [default to "false"]
**SecureAccessHost** | Pointer to **[]string** | Target DB servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts) | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateOracleDb

`func NewDynamicSecretCreateOracleDb(name string, ) *DynamicSecretCreateOracleDb`

NewDynamicSecretCreateOracleDb instantiates a new DynamicSecretCreateOracleDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateOracleDbWithDefaults

`func NewDynamicSecretCreateOracleDbWithDefaults() *DynamicSecretCreateOracleDb`

NewDynamicSecretCreateOracleDbWithDefaults instantiates a new DynamicSecretCreateOracleDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbServerCertificates

`func (o *DynamicSecretCreateOracleDb) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *DynamicSecretCreateOracleDb) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *DynamicSecretCreateOracleDb) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *DynamicSecretCreateOracleDb) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *DynamicSecretCreateOracleDb) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *DynamicSecretCreateOracleDb) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *DynamicSecretCreateOracleDb) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *DynamicSecretCreateOracleDb) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateOracleDb) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateOracleDb) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateOracleDb) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateOracleDb) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateOracleDb) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateOracleDb) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateOracleDb) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateOracleDb) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateOracleDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateOracleDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateOracleDb) SetName(v string)`

SetName sets Name field to given value.


### GetOracleHost

`func (o *DynamicSecretCreateOracleDb) GetOracleHost() string`

GetOracleHost returns the OracleHost field if non-nil, zero value otherwise.

### GetOracleHostOk

`func (o *DynamicSecretCreateOracleDb) GetOracleHostOk() (*string, bool)`

GetOracleHostOk returns a tuple with the OracleHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleHost

`func (o *DynamicSecretCreateOracleDb) SetOracleHost(v string)`

SetOracleHost sets OracleHost field to given value.

### HasOracleHost

`func (o *DynamicSecretCreateOracleDb) HasOracleHost() bool`

HasOracleHost returns a boolean if a field has been set.

### GetOraclePassword

`func (o *DynamicSecretCreateOracleDb) GetOraclePassword() string`

GetOraclePassword returns the OraclePassword field if non-nil, zero value otherwise.

### GetOraclePasswordOk

`func (o *DynamicSecretCreateOracleDb) GetOraclePasswordOk() (*string, bool)`

GetOraclePasswordOk returns a tuple with the OraclePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOraclePassword

`func (o *DynamicSecretCreateOracleDb) SetOraclePassword(v string)`

SetOraclePassword sets OraclePassword field to given value.

### HasOraclePassword

`func (o *DynamicSecretCreateOracleDb) HasOraclePassword() bool`

HasOraclePassword returns a boolean if a field has been set.

### GetOraclePort

`func (o *DynamicSecretCreateOracleDb) GetOraclePort() string`

GetOraclePort returns the OraclePort field if non-nil, zero value otherwise.

### GetOraclePortOk

`func (o *DynamicSecretCreateOracleDb) GetOraclePortOk() (*string, bool)`

GetOraclePortOk returns a tuple with the OraclePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOraclePort

`func (o *DynamicSecretCreateOracleDb) SetOraclePort(v string)`

SetOraclePort sets OraclePort field to given value.

### HasOraclePort

`func (o *DynamicSecretCreateOracleDb) HasOraclePort() bool`

HasOraclePort returns a boolean if a field has been set.

### GetOracleScreationStatements

`func (o *DynamicSecretCreateOracleDb) GetOracleScreationStatements() string`

GetOracleScreationStatements returns the OracleScreationStatements field if non-nil, zero value otherwise.

### GetOracleScreationStatementsOk

`func (o *DynamicSecretCreateOracleDb) GetOracleScreationStatementsOk() (*string, bool)`

GetOracleScreationStatementsOk returns a tuple with the OracleScreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleScreationStatements

`func (o *DynamicSecretCreateOracleDb) SetOracleScreationStatements(v string)`

SetOracleScreationStatements sets OracleScreationStatements field to given value.

### HasOracleScreationStatements

`func (o *DynamicSecretCreateOracleDb) HasOracleScreationStatements() bool`

HasOracleScreationStatements returns a boolean if a field has been set.

### GetOracleServiceName

`func (o *DynamicSecretCreateOracleDb) GetOracleServiceName() string`

GetOracleServiceName returns the OracleServiceName field if non-nil, zero value otherwise.

### GetOracleServiceNameOk

`func (o *DynamicSecretCreateOracleDb) GetOracleServiceNameOk() (*string, bool)`

GetOracleServiceNameOk returns a tuple with the OracleServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleServiceName

`func (o *DynamicSecretCreateOracleDb) SetOracleServiceName(v string)`

SetOracleServiceName sets OracleServiceName field to given value.

### HasOracleServiceName

`func (o *DynamicSecretCreateOracleDb) HasOracleServiceName() bool`

HasOracleServiceName returns a boolean if a field has been set.

### GetOracleUsername

`func (o *DynamicSecretCreateOracleDb) GetOracleUsername() string`

GetOracleUsername returns the OracleUsername field if non-nil, zero value otherwise.

### GetOracleUsernameOk

`func (o *DynamicSecretCreateOracleDb) GetOracleUsernameOk() (*string, bool)`

GetOracleUsernameOk returns a tuple with the OracleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleUsername

`func (o *DynamicSecretCreateOracleDb) SetOracleUsername(v string)`

SetOracleUsername sets OracleUsername field to given value.

### HasOracleUsername

`func (o *DynamicSecretCreateOracleDb) HasOracleUsername() bool`

HasOracleUsername returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateOracleDb) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateOracleDb) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateOracleDb) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateOracleDb) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateOracleDb) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretCreateOracleDb) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateOracleDb) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretCreateOracleDb) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateOracleDb) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateOracleDb) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateOracleDb) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateOracleDb) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *DynamicSecretCreateOracleDb) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *DynamicSecretCreateOracleDb) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *DynamicSecretCreateOracleDb) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *DynamicSecretCreateOracleDb) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretCreateOracleDb) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretCreateOracleDb) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretCreateOracleDb) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretCreateOracleDb) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateOracleDb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateOracleDb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateOracleDb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateOracleDb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateOracleDb) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateOracleDb) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateOracleDb) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateOracleDb) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateOracleDb) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateOracleDb) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateOracleDb) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateOracleDb) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateOracleDb) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateOracleDb) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateOracleDb) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateOracleDb) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateOracleDb) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateOracleDb) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateOracleDb) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateOracleDb) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


