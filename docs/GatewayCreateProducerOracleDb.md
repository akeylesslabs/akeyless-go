# GatewayCreateProducerOracleDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbServerCertificates** | Pointer to **string** | (Optional) DB server certificates | [optional] 
**DbServerName** | Pointer to **string** | (Optional) Server name for certificate verification | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | **string** | Producer name | 
**OracleHost** | Pointer to **string** | Oracle Host | [optional] [default to "127.0.0.1"]
**OraclePassword** | Pointer to **string** | Oracle Password | [optional] 
**OraclePort** | Pointer to **string** | Oracle Port | [optional] [default to "1521"]
**OracleScreationStatements** | Pointer to **string** | Oracle Creation statements | [optional] 
**OracleServiceName** | Pointer to **string** | Oracle DB Name | [optional] 
**OracleUsername** | Pointer to **string** | Oracle Username | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** |  | [optional] 
**SecureAccessEnable** | Pointer to **string** |  | [optional] 
**SecureAccessHost** | Pointer to **[]string** |  | [optional] 
**SecureAccessWeb** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerOracleDb

`func NewGatewayCreateProducerOracleDb(name string, ) *GatewayCreateProducerOracleDb`

NewGatewayCreateProducerOracleDb instantiates a new GatewayCreateProducerOracleDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerOracleDbWithDefaults

`func NewGatewayCreateProducerOracleDbWithDefaults() *GatewayCreateProducerOracleDb`

NewGatewayCreateProducerOracleDbWithDefaults instantiates a new GatewayCreateProducerOracleDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbServerCertificates

`func (o *GatewayCreateProducerOracleDb) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *GatewayCreateProducerOracleDb) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *GatewayCreateProducerOracleDb) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *GatewayCreateProducerOracleDb) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *GatewayCreateProducerOracleDb) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *GatewayCreateProducerOracleDb) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *GatewayCreateProducerOracleDb) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *GatewayCreateProducerOracleDb) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayCreateProducerOracleDb) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayCreateProducerOracleDb) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayCreateProducerOracleDb) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayCreateProducerOracleDb) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateProducerOracleDb) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateProducerOracleDb) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateProducerOracleDb) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateProducerOracleDb) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerOracleDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerOracleDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerOracleDb) SetName(v string)`

SetName sets Name field to given value.


### GetOracleHost

`func (o *GatewayCreateProducerOracleDb) GetOracleHost() string`

GetOracleHost returns the OracleHost field if non-nil, zero value otherwise.

### GetOracleHostOk

`func (o *GatewayCreateProducerOracleDb) GetOracleHostOk() (*string, bool)`

GetOracleHostOk returns a tuple with the OracleHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleHost

`func (o *GatewayCreateProducerOracleDb) SetOracleHost(v string)`

SetOracleHost sets OracleHost field to given value.

### HasOracleHost

`func (o *GatewayCreateProducerOracleDb) HasOracleHost() bool`

HasOracleHost returns a boolean if a field has been set.

### GetOraclePassword

`func (o *GatewayCreateProducerOracleDb) GetOraclePassword() string`

GetOraclePassword returns the OraclePassword field if non-nil, zero value otherwise.

### GetOraclePasswordOk

`func (o *GatewayCreateProducerOracleDb) GetOraclePasswordOk() (*string, bool)`

GetOraclePasswordOk returns a tuple with the OraclePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOraclePassword

`func (o *GatewayCreateProducerOracleDb) SetOraclePassword(v string)`

SetOraclePassword sets OraclePassword field to given value.

### HasOraclePassword

`func (o *GatewayCreateProducerOracleDb) HasOraclePassword() bool`

HasOraclePassword returns a boolean if a field has been set.

### GetOraclePort

`func (o *GatewayCreateProducerOracleDb) GetOraclePort() string`

GetOraclePort returns the OraclePort field if non-nil, zero value otherwise.

### GetOraclePortOk

`func (o *GatewayCreateProducerOracleDb) GetOraclePortOk() (*string, bool)`

GetOraclePortOk returns a tuple with the OraclePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOraclePort

`func (o *GatewayCreateProducerOracleDb) SetOraclePort(v string)`

SetOraclePort sets OraclePort field to given value.

### HasOraclePort

`func (o *GatewayCreateProducerOracleDb) HasOraclePort() bool`

HasOraclePort returns a boolean if a field has been set.

### GetOracleScreationStatements

`func (o *GatewayCreateProducerOracleDb) GetOracleScreationStatements() string`

GetOracleScreationStatements returns the OracleScreationStatements field if non-nil, zero value otherwise.

### GetOracleScreationStatementsOk

`func (o *GatewayCreateProducerOracleDb) GetOracleScreationStatementsOk() (*string, bool)`

GetOracleScreationStatementsOk returns a tuple with the OracleScreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleScreationStatements

`func (o *GatewayCreateProducerOracleDb) SetOracleScreationStatements(v string)`

SetOracleScreationStatements sets OracleScreationStatements field to given value.

### HasOracleScreationStatements

`func (o *GatewayCreateProducerOracleDb) HasOracleScreationStatements() bool`

HasOracleScreationStatements returns a boolean if a field has been set.

### GetOracleServiceName

`func (o *GatewayCreateProducerOracleDb) GetOracleServiceName() string`

GetOracleServiceName returns the OracleServiceName field if non-nil, zero value otherwise.

### GetOracleServiceNameOk

`func (o *GatewayCreateProducerOracleDb) GetOracleServiceNameOk() (*string, bool)`

GetOracleServiceNameOk returns a tuple with the OracleServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleServiceName

`func (o *GatewayCreateProducerOracleDb) SetOracleServiceName(v string)`

SetOracleServiceName sets OracleServiceName field to given value.

### HasOracleServiceName

`func (o *GatewayCreateProducerOracleDb) HasOracleServiceName() bool`

HasOracleServiceName returns a boolean if a field has been set.

### GetOracleUsername

`func (o *GatewayCreateProducerOracleDb) GetOracleUsername() string`

GetOracleUsername returns the OracleUsername field if non-nil, zero value otherwise.

### GetOracleUsernameOk

`func (o *GatewayCreateProducerOracleDb) GetOracleUsernameOk() (*string, bool)`

GetOracleUsernameOk returns a tuple with the OracleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleUsername

`func (o *GatewayCreateProducerOracleDb) SetOracleUsername(v string)`

SetOracleUsername sets OracleUsername field to given value.

### HasOracleUsername

`func (o *GatewayCreateProducerOracleDb) HasOracleUsername() bool`

HasOracleUsername returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerOracleDb) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerOracleDb) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerOracleDb) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerOracleDb) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerOracleDb) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayCreateProducerOracleDb) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerOracleDb) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayCreateProducerOracleDb) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayCreateProducerOracleDb) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayCreateProducerOracleDb) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayCreateProducerOracleDb) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayCreateProducerOracleDb) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *GatewayCreateProducerOracleDb) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *GatewayCreateProducerOracleDb) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *GatewayCreateProducerOracleDb) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *GatewayCreateProducerOracleDb) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayCreateProducerOracleDb) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayCreateProducerOracleDb) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayCreateProducerOracleDb) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayCreateProducerOracleDb) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerOracleDb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerOracleDb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerOracleDb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerOracleDb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerOracleDb) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerOracleDb) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerOracleDb) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerOracleDb) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerOracleDb) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerOracleDb) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerOracleDb) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerOracleDb) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerOracleDb) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerOracleDb) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerOracleDb) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerOracleDb) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerOracleDb) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerOracleDb) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerOracleDb) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerOracleDb) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


