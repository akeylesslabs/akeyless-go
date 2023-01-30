# TargetTypeDetailesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKeyId** | Pointer to **string** |  | [optional] 
**AwsRegion** | Pointer to **string** |  | [optional] 
**AwsSecretAccessKey** | Pointer to **string** |  | [optional] 
**AwsSessionToken** | Pointer to **string** |  | [optional] 
**DbHostName** | Pointer to **string** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**DbPort** | Pointer to **string** |  | [optional] 
**DbPwd** | Pointer to **string** |  | [optional] 
**DbServerCertificates** | Pointer to **string** | (Optional) DBServerCertificates defines the set of root certificate authorities that clients use when verifying server certificates. If DBServerCertificates is empty, TLS uses the host&#39;s root CA set. | [optional] 
**DbServerName** | Pointer to **string** | (Optional) ServerName is used to verify the hostname on the returned certificates unless InsecureSkipVerify is given. It is also included in the client&#39;s handshake to support virtual hosting unless it is an IP address. | [optional] 
**DbUserName** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**MongodbDbName** | Pointer to **string** |  | [optional] 
**MongodbUriConnection** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**PrivateKeyPassword** | Pointer to **string** |  | [optional] 
**RabbitmqServerPassword** | Pointer to **string** |  | [optional] 
**RabbitmqServerUri** | Pointer to **string** |  | [optional] 
**RabbitmqServerUser** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewTargetTypeDetailesInput

`func NewTargetTypeDetailesInput() *TargetTypeDetailesInput`

NewTargetTypeDetailesInput instantiates a new TargetTypeDetailesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetTypeDetailesInputWithDefaults

`func NewTargetTypeDetailesInputWithDefaults() *TargetTypeDetailesInput`

NewTargetTypeDetailesInputWithDefaults instantiates a new TargetTypeDetailesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKeyId

`func (o *TargetTypeDetailesInput) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *TargetTypeDetailesInput) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *TargetTypeDetailesInput) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *TargetTypeDetailesInput) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *TargetTypeDetailesInput) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *TargetTypeDetailesInput) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *TargetTypeDetailesInput) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *TargetTypeDetailesInput) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *TargetTypeDetailesInput) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *TargetTypeDetailesInput) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *TargetTypeDetailesInput) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *TargetTypeDetailesInput) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsSessionToken

`func (o *TargetTypeDetailesInput) GetAwsSessionToken() string`

GetAwsSessionToken returns the AwsSessionToken field if non-nil, zero value otherwise.

### GetAwsSessionTokenOk

`func (o *TargetTypeDetailesInput) GetAwsSessionTokenOk() (*string, bool)`

GetAwsSessionTokenOk returns a tuple with the AwsSessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSessionToken

`func (o *TargetTypeDetailesInput) SetAwsSessionToken(v string)`

SetAwsSessionToken sets AwsSessionToken field to given value.

### HasAwsSessionToken

`func (o *TargetTypeDetailesInput) HasAwsSessionToken() bool`

HasAwsSessionToken returns a boolean if a field has been set.

### GetDbHostName

`func (o *TargetTypeDetailesInput) GetDbHostName() string`

GetDbHostName returns the DbHostName field if non-nil, zero value otherwise.

### GetDbHostNameOk

`func (o *TargetTypeDetailesInput) GetDbHostNameOk() (*string, bool)`

GetDbHostNameOk returns a tuple with the DbHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHostName

`func (o *TargetTypeDetailesInput) SetDbHostName(v string)`

SetDbHostName sets DbHostName field to given value.

### HasDbHostName

`func (o *TargetTypeDetailesInput) HasDbHostName() bool`

HasDbHostName returns a boolean if a field has been set.

### GetDbName

`func (o *TargetTypeDetailesInput) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TargetTypeDetailesInput) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TargetTypeDetailesInput) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *TargetTypeDetailesInput) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbPort

`func (o *TargetTypeDetailesInput) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *TargetTypeDetailesInput) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *TargetTypeDetailesInput) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.

### HasDbPort

`func (o *TargetTypeDetailesInput) HasDbPort() bool`

HasDbPort returns a boolean if a field has been set.

### GetDbPwd

`func (o *TargetTypeDetailesInput) GetDbPwd() string`

GetDbPwd returns the DbPwd field if non-nil, zero value otherwise.

### GetDbPwdOk

`func (o *TargetTypeDetailesInput) GetDbPwdOk() (*string, bool)`

GetDbPwdOk returns a tuple with the DbPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPwd

`func (o *TargetTypeDetailesInput) SetDbPwd(v string)`

SetDbPwd sets DbPwd field to given value.

### HasDbPwd

`func (o *TargetTypeDetailesInput) HasDbPwd() bool`

HasDbPwd returns a boolean if a field has been set.

### GetDbServerCertificates

`func (o *TargetTypeDetailesInput) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *TargetTypeDetailesInput) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *TargetTypeDetailesInput) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *TargetTypeDetailesInput) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *TargetTypeDetailesInput) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *TargetTypeDetailesInput) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *TargetTypeDetailesInput) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *TargetTypeDetailesInput) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDbUserName

`func (o *TargetTypeDetailesInput) GetDbUserName() string`

GetDbUserName returns the DbUserName field if non-nil, zero value otherwise.

### GetDbUserNameOk

`func (o *TargetTypeDetailesInput) GetDbUserNameOk() (*string, bool)`

GetDbUserNameOk returns a tuple with the DbUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUserName

`func (o *TargetTypeDetailesInput) SetDbUserName(v string)`

SetDbUserName sets DbUserName field to given value.

### HasDbUserName

`func (o *TargetTypeDetailesInput) HasDbUserName() bool`

HasDbUserName returns a boolean if a field has been set.

### GetHost

`func (o *TargetTypeDetailesInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetTypeDetailesInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetTypeDetailesInput) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TargetTypeDetailesInput) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMongodbDbName

`func (o *TargetTypeDetailesInput) GetMongodbDbName() string`

GetMongodbDbName returns the MongodbDbName field if non-nil, zero value otherwise.

### GetMongodbDbNameOk

`func (o *TargetTypeDetailesInput) GetMongodbDbNameOk() (*string, bool)`

GetMongodbDbNameOk returns a tuple with the MongodbDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDbName

`func (o *TargetTypeDetailesInput) SetMongodbDbName(v string)`

SetMongodbDbName sets MongodbDbName field to given value.

### HasMongodbDbName

`func (o *TargetTypeDetailesInput) HasMongodbDbName() bool`

HasMongodbDbName returns a boolean if a field has been set.

### GetMongodbUriConnection

`func (o *TargetTypeDetailesInput) GetMongodbUriConnection() string`

GetMongodbUriConnection returns the MongodbUriConnection field if non-nil, zero value otherwise.

### GetMongodbUriConnectionOk

`func (o *TargetTypeDetailesInput) GetMongodbUriConnectionOk() (*string, bool)`

GetMongodbUriConnectionOk returns a tuple with the MongodbUriConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriConnection

`func (o *TargetTypeDetailesInput) SetMongodbUriConnection(v string)`

SetMongodbUriConnection sets MongodbUriConnection field to given value.

### HasMongodbUriConnection

`func (o *TargetTypeDetailesInput) HasMongodbUriConnection() bool`

HasMongodbUriConnection returns a boolean if a field has been set.

### GetPassword

`func (o *TargetTypeDetailesInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetTypeDetailesInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetTypeDetailesInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TargetTypeDetailesInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *TargetTypeDetailesInput) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetTypeDetailesInput) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetTypeDetailesInput) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetTypeDetailesInput) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateKey

`func (o *TargetTypeDetailesInput) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *TargetTypeDetailesInput) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *TargetTypeDetailesInput) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *TargetTypeDetailesInput) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyPassword

`func (o *TargetTypeDetailesInput) GetPrivateKeyPassword() string`

GetPrivateKeyPassword returns the PrivateKeyPassword field if non-nil, zero value otherwise.

### GetPrivateKeyPasswordOk

`func (o *TargetTypeDetailesInput) GetPrivateKeyPasswordOk() (*string, bool)`

GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassword

`func (o *TargetTypeDetailesInput) SetPrivateKeyPassword(v string)`

SetPrivateKeyPassword sets PrivateKeyPassword field to given value.

### HasPrivateKeyPassword

`func (o *TargetTypeDetailesInput) HasPrivateKeyPassword() bool`

HasPrivateKeyPassword returns a boolean if a field has been set.

### GetRabbitmqServerPassword

`func (o *TargetTypeDetailesInput) GetRabbitmqServerPassword() string`

GetRabbitmqServerPassword returns the RabbitmqServerPassword field if non-nil, zero value otherwise.

### GetRabbitmqServerPasswordOk

`func (o *TargetTypeDetailesInput) GetRabbitmqServerPasswordOk() (*string, bool)`

GetRabbitmqServerPasswordOk returns a tuple with the RabbitmqServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerPassword

`func (o *TargetTypeDetailesInput) SetRabbitmqServerPassword(v string)`

SetRabbitmqServerPassword sets RabbitmqServerPassword field to given value.

### HasRabbitmqServerPassword

`func (o *TargetTypeDetailesInput) HasRabbitmqServerPassword() bool`

HasRabbitmqServerPassword returns a boolean if a field has been set.

### GetRabbitmqServerUri

`func (o *TargetTypeDetailesInput) GetRabbitmqServerUri() string`

GetRabbitmqServerUri returns the RabbitmqServerUri field if non-nil, zero value otherwise.

### GetRabbitmqServerUriOk

`func (o *TargetTypeDetailesInput) GetRabbitmqServerUriOk() (*string, bool)`

GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUri

`func (o *TargetTypeDetailesInput) SetRabbitmqServerUri(v string)`

SetRabbitmqServerUri sets RabbitmqServerUri field to given value.

### HasRabbitmqServerUri

`func (o *TargetTypeDetailesInput) HasRabbitmqServerUri() bool`

HasRabbitmqServerUri returns a boolean if a field has been set.

### GetRabbitmqServerUser

`func (o *TargetTypeDetailesInput) GetRabbitmqServerUser() string`

GetRabbitmqServerUser returns the RabbitmqServerUser field if non-nil, zero value otherwise.

### GetRabbitmqServerUserOk

`func (o *TargetTypeDetailesInput) GetRabbitmqServerUserOk() (*string, bool)`

GetRabbitmqServerUserOk returns a tuple with the RabbitmqServerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUser

`func (o *TargetTypeDetailesInput) SetRabbitmqServerUser(v string)`

SetRabbitmqServerUser sets RabbitmqServerUser field to given value.

### HasRabbitmqServerUser

`func (o *TargetTypeDetailesInput) HasRabbitmqServerUser() bool`

HasRabbitmqServerUser returns a boolean if a field has been set.

### GetUrl

`func (o *TargetTypeDetailesInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TargetTypeDetailesInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TargetTypeDetailesInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TargetTypeDetailesInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *TargetTypeDetailesInput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TargetTypeDetailesInput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TargetTypeDetailesInput) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TargetTypeDetailesInput) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


