# GatewayCreateProducerMSSQL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MssqlCreateStatements** | Pointer to **string** | MSSQL Creation statements | [optional] 
**MssqlDbname** | **string** | MSSQL Name | 
**MssqlHost** | Pointer to **string** | MSSQL Host | [optional] [default to "127.0.0.1"]
**MssqlPassword** | **string** | MSSQL Password | 
**MssqlPort** | Pointer to **string** | MSSQL Port | [optional] [default to "1433"]
**MssqlRevocationStatements** | Pointer to **string** | MSSQL Revocation statements | [optional] 
**MssqlUsername** | **string** | MSSQL Username | 
**Name** | **string** | Producer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayCreateProducerMSSQL

`func NewGatewayCreateProducerMSSQL(mssqlDbname string, mssqlPassword string, mssqlUsername string, name string, ) *GatewayCreateProducerMSSQL`

NewGatewayCreateProducerMSSQL instantiates a new GatewayCreateProducerMSSQL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerMSSQLWithDefaults

`func NewGatewayCreateProducerMSSQLWithDefaults() *GatewayCreateProducerMSSQL`

NewGatewayCreateProducerMSSQLWithDefaults instantiates a new GatewayCreateProducerMSSQL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMssqlCreateStatements

`func (o *GatewayCreateProducerMSSQL) GetMssqlCreateStatements() string`

GetMssqlCreateStatements returns the MssqlCreateStatements field if non-nil, zero value otherwise.

### GetMssqlCreateStatementsOk

`func (o *GatewayCreateProducerMSSQL) GetMssqlCreateStatementsOk() (*string, bool)`

GetMssqlCreateStatementsOk returns a tuple with the MssqlCreateStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlCreateStatements

`func (o *GatewayCreateProducerMSSQL) SetMssqlCreateStatements(v string)`

SetMssqlCreateStatements sets MssqlCreateStatements field to given value.

### HasMssqlCreateStatements

`func (o *GatewayCreateProducerMSSQL) HasMssqlCreateStatements() bool`

HasMssqlCreateStatements returns a boolean if a field has been set.

### GetMssqlDbname

`func (o *GatewayCreateProducerMSSQL) GetMssqlDbname() string`

GetMssqlDbname returns the MssqlDbname field if non-nil, zero value otherwise.

### GetMssqlDbnameOk

`func (o *GatewayCreateProducerMSSQL) GetMssqlDbnameOk() (*string, bool)`

GetMssqlDbnameOk returns a tuple with the MssqlDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlDbname

`func (o *GatewayCreateProducerMSSQL) SetMssqlDbname(v string)`

SetMssqlDbname sets MssqlDbname field to given value.


### GetMssqlHost

`func (o *GatewayCreateProducerMSSQL) GetMssqlHost() string`

GetMssqlHost returns the MssqlHost field if non-nil, zero value otherwise.

### GetMssqlHostOk

`func (o *GatewayCreateProducerMSSQL) GetMssqlHostOk() (*string, bool)`

GetMssqlHostOk returns a tuple with the MssqlHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlHost

`func (o *GatewayCreateProducerMSSQL) SetMssqlHost(v string)`

SetMssqlHost sets MssqlHost field to given value.

### HasMssqlHost

`func (o *GatewayCreateProducerMSSQL) HasMssqlHost() bool`

HasMssqlHost returns a boolean if a field has been set.

### GetMssqlPassword

`func (o *GatewayCreateProducerMSSQL) GetMssqlPassword() string`

GetMssqlPassword returns the MssqlPassword field if non-nil, zero value otherwise.

### GetMssqlPasswordOk

`func (o *GatewayCreateProducerMSSQL) GetMssqlPasswordOk() (*string, bool)`

GetMssqlPasswordOk returns a tuple with the MssqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlPassword

`func (o *GatewayCreateProducerMSSQL) SetMssqlPassword(v string)`

SetMssqlPassword sets MssqlPassword field to given value.


### GetMssqlPort

`func (o *GatewayCreateProducerMSSQL) GetMssqlPort() string`

GetMssqlPort returns the MssqlPort field if non-nil, zero value otherwise.

### GetMssqlPortOk

`func (o *GatewayCreateProducerMSSQL) GetMssqlPortOk() (*string, bool)`

GetMssqlPortOk returns a tuple with the MssqlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlPort

`func (o *GatewayCreateProducerMSSQL) SetMssqlPort(v string)`

SetMssqlPort sets MssqlPort field to given value.

### HasMssqlPort

`func (o *GatewayCreateProducerMSSQL) HasMssqlPort() bool`

HasMssqlPort returns a boolean if a field has been set.

### GetMssqlRevocationStatements

`func (o *GatewayCreateProducerMSSQL) GetMssqlRevocationStatements() string`

GetMssqlRevocationStatements returns the MssqlRevocationStatements field if non-nil, zero value otherwise.

### GetMssqlRevocationStatementsOk

`func (o *GatewayCreateProducerMSSQL) GetMssqlRevocationStatementsOk() (*string, bool)`

GetMssqlRevocationStatementsOk returns a tuple with the MssqlRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlRevocationStatements

`func (o *GatewayCreateProducerMSSQL) SetMssqlRevocationStatements(v string)`

SetMssqlRevocationStatements sets MssqlRevocationStatements field to given value.

### HasMssqlRevocationStatements

`func (o *GatewayCreateProducerMSSQL) HasMssqlRevocationStatements() bool`

HasMssqlRevocationStatements returns a boolean if a field has been set.

### GetMssqlUsername

`func (o *GatewayCreateProducerMSSQL) GetMssqlUsername() string`

GetMssqlUsername returns the MssqlUsername field if non-nil, zero value otherwise.

### GetMssqlUsernameOk

`func (o *GatewayCreateProducerMSSQL) GetMssqlUsernameOk() (*string, bool)`

GetMssqlUsernameOk returns a tuple with the MssqlUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlUsername

`func (o *GatewayCreateProducerMSSQL) SetMssqlUsername(v string)`

SetMssqlUsername sets MssqlUsername field to given value.


### GetName

`func (o *GatewayCreateProducerMSSQL) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerMSSQL) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerMSSQL) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *GatewayCreateProducerMSSQL) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateProducerMSSQL) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateProducerMSSQL) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateProducerMSSQL) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerMSSQL) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerMSSQL) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerMSSQL) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerMSSQL) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerMSSQL) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerMSSQL) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerMSSQL) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerMSSQL) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerMSSQL) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerMSSQL) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerMSSQL) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerMSSQL) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerMSSQL) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerMSSQL) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerMSSQL) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerMSSQL) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayCreateProducerMSSQL) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateProducerMSSQL) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateProducerMSSQL) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateProducerMSSQL) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


