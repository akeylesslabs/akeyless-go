# DbTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudServiceProvider** | Pointer to **string** |  | [optional] 
**ClusterMode** | Pointer to **bool** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**DbClientId** | Pointer to **string** |  | [optional] 
**DbClientSecret** | Pointer to **string** |  | [optional] 
**DbHostName** | Pointer to **string** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**DbPort** | Pointer to **string** |  | [optional] 
**DbPrivateKey** | Pointer to **string** | (Optional) Private Key in PEM format | [optional] 
**DbPrivateKeyPassphrase** | Pointer to **string** |  | [optional] 
**DbPwd** | Pointer to **string** |  | [optional] 
**DbServerCertificates** | Pointer to **string** | (Optional) DBServerCertificates defines the set of root certificate authorities that clients use when verifying server certificates. If DBServerCertificates is empty, TLS uses the host&#39;s root CA set. | [optional] 
**DbServerName** | Pointer to **string** | (Optional) ServerName is used to verify the hostname on the returned certificates unless InsecureSkipVerify is given. It is also included in the client&#39;s handshake to support virtual hosting unless it is an IP address. | [optional] 
**DbTenantId** | Pointer to **string** |  | [optional] 
**DbUserName** | Pointer to **string** |  | [optional] 
**SfAccount** | Pointer to **string** |  | [optional] 
**SslConnectionCertificate** | Pointer to **string** | (Optional) SSLConnectionCertificate defines the certificate for SSL connection. Must be base64 certificate loaded by UI using file loader field | [optional] 
**SslConnectionMode** | Pointer to **bool** | (Optional) SSLConnectionMode defines if SSL mode will be used to connect to DB | [optional] 

## Methods

### NewDbTargetDetails

`func NewDbTargetDetails() *DbTargetDetails`

NewDbTargetDetails instantiates a new DbTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbTargetDetailsWithDefaults

`func NewDbTargetDetailsWithDefaults() *DbTargetDetails`

NewDbTargetDetailsWithDefaults instantiates a new DbTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudServiceProvider

`func (o *DbTargetDetails) GetCloudServiceProvider() string`

GetCloudServiceProvider returns the CloudServiceProvider field if non-nil, zero value otherwise.

### GetCloudServiceProviderOk

`func (o *DbTargetDetails) GetCloudServiceProviderOk() (*string, bool)`

GetCloudServiceProviderOk returns a tuple with the CloudServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudServiceProvider

`func (o *DbTargetDetails) SetCloudServiceProvider(v string)`

SetCloudServiceProvider sets CloudServiceProvider field to given value.

### HasCloudServiceProvider

`func (o *DbTargetDetails) HasCloudServiceProvider() bool`

HasCloudServiceProvider returns a boolean if a field has been set.

### GetClusterMode

`func (o *DbTargetDetails) GetClusterMode() bool`

GetClusterMode returns the ClusterMode field if non-nil, zero value otherwise.

### GetClusterModeOk

`func (o *DbTargetDetails) GetClusterModeOk() (*bool, bool)`

GetClusterModeOk returns a tuple with the ClusterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMode

`func (o *DbTargetDetails) SetClusterMode(v bool)`

SetClusterMode sets ClusterMode field to given value.

### HasClusterMode

`func (o *DbTargetDetails) HasClusterMode() bool`

HasClusterMode returns a boolean if a field has been set.

### GetConnectionType

`func (o *DbTargetDetails) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *DbTargetDetails) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *DbTargetDetails) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *DbTargetDetails) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetDbClientId

`func (o *DbTargetDetails) GetDbClientId() string`

GetDbClientId returns the DbClientId field if non-nil, zero value otherwise.

### GetDbClientIdOk

`func (o *DbTargetDetails) GetDbClientIdOk() (*string, bool)`

GetDbClientIdOk returns a tuple with the DbClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbClientId

`func (o *DbTargetDetails) SetDbClientId(v string)`

SetDbClientId sets DbClientId field to given value.

### HasDbClientId

`func (o *DbTargetDetails) HasDbClientId() bool`

HasDbClientId returns a boolean if a field has been set.

### GetDbClientSecret

`func (o *DbTargetDetails) GetDbClientSecret() string`

GetDbClientSecret returns the DbClientSecret field if non-nil, zero value otherwise.

### GetDbClientSecretOk

`func (o *DbTargetDetails) GetDbClientSecretOk() (*string, bool)`

GetDbClientSecretOk returns a tuple with the DbClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbClientSecret

`func (o *DbTargetDetails) SetDbClientSecret(v string)`

SetDbClientSecret sets DbClientSecret field to given value.

### HasDbClientSecret

`func (o *DbTargetDetails) HasDbClientSecret() bool`

HasDbClientSecret returns a boolean if a field has been set.

### GetDbHostName

`func (o *DbTargetDetails) GetDbHostName() string`

GetDbHostName returns the DbHostName field if non-nil, zero value otherwise.

### GetDbHostNameOk

`func (o *DbTargetDetails) GetDbHostNameOk() (*string, bool)`

GetDbHostNameOk returns a tuple with the DbHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHostName

`func (o *DbTargetDetails) SetDbHostName(v string)`

SetDbHostName sets DbHostName field to given value.

### HasDbHostName

`func (o *DbTargetDetails) HasDbHostName() bool`

HasDbHostName returns a boolean if a field has been set.

### GetDbName

`func (o *DbTargetDetails) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DbTargetDetails) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DbTargetDetails) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DbTargetDetails) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbPort

`func (o *DbTargetDetails) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *DbTargetDetails) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *DbTargetDetails) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.

### HasDbPort

`func (o *DbTargetDetails) HasDbPort() bool`

HasDbPort returns a boolean if a field has been set.

### GetDbPrivateKey

`func (o *DbTargetDetails) GetDbPrivateKey() string`

GetDbPrivateKey returns the DbPrivateKey field if non-nil, zero value otherwise.

### GetDbPrivateKeyOk

`func (o *DbTargetDetails) GetDbPrivateKeyOk() (*string, bool)`

GetDbPrivateKeyOk returns a tuple with the DbPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPrivateKey

`func (o *DbTargetDetails) SetDbPrivateKey(v string)`

SetDbPrivateKey sets DbPrivateKey field to given value.

### HasDbPrivateKey

`func (o *DbTargetDetails) HasDbPrivateKey() bool`

HasDbPrivateKey returns a boolean if a field has been set.

### GetDbPrivateKeyPassphrase

`func (o *DbTargetDetails) GetDbPrivateKeyPassphrase() string`

GetDbPrivateKeyPassphrase returns the DbPrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetDbPrivateKeyPassphraseOk

`func (o *DbTargetDetails) GetDbPrivateKeyPassphraseOk() (*string, bool)`

GetDbPrivateKeyPassphraseOk returns a tuple with the DbPrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPrivateKeyPassphrase

`func (o *DbTargetDetails) SetDbPrivateKeyPassphrase(v string)`

SetDbPrivateKeyPassphrase sets DbPrivateKeyPassphrase field to given value.

### HasDbPrivateKeyPassphrase

`func (o *DbTargetDetails) HasDbPrivateKeyPassphrase() bool`

HasDbPrivateKeyPassphrase returns a boolean if a field has been set.

### GetDbPwd

`func (o *DbTargetDetails) GetDbPwd() string`

GetDbPwd returns the DbPwd field if non-nil, zero value otherwise.

### GetDbPwdOk

`func (o *DbTargetDetails) GetDbPwdOk() (*string, bool)`

GetDbPwdOk returns a tuple with the DbPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPwd

`func (o *DbTargetDetails) SetDbPwd(v string)`

SetDbPwd sets DbPwd field to given value.

### HasDbPwd

`func (o *DbTargetDetails) HasDbPwd() bool`

HasDbPwd returns a boolean if a field has been set.

### GetDbServerCertificates

`func (o *DbTargetDetails) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *DbTargetDetails) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *DbTargetDetails) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *DbTargetDetails) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *DbTargetDetails) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *DbTargetDetails) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *DbTargetDetails) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *DbTargetDetails) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDbTenantId

`func (o *DbTargetDetails) GetDbTenantId() string`

GetDbTenantId returns the DbTenantId field if non-nil, zero value otherwise.

### GetDbTenantIdOk

`func (o *DbTargetDetails) GetDbTenantIdOk() (*string, bool)`

GetDbTenantIdOk returns a tuple with the DbTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTenantId

`func (o *DbTargetDetails) SetDbTenantId(v string)`

SetDbTenantId sets DbTenantId field to given value.

### HasDbTenantId

`func (o *DbTargetDetails) HasDbTenantId() bool`

HasDbTenantId returns a boolean if a field has been set.

### GetDbUserName

`func (o *DbTargetDetails) GetDbUserName() string`

GetDbUserName returns the DbUserName field if non-nil, zero value otherwise.

### GetDbUserNameOk

`func (o *DbTargetDetails) GetDbUserNameOk() (*string, bool)`

GetDbUserNameOk returns a tuple with the DbUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUserName

`func (o *DbTargetDetails) SetDbUserName(v string)`

SetDbUserName sets DbUserName field to given value.

### HasDbUserName

`func (o *DbTargetDetails) HasDbUserName() bool`

HasDbUserName returns a boolean if a field has been set.

### GetSfAccount

`func (o *DbTargetDetails) GetSfAccount() string`

GetSfAccount returns the SfAccount field if non-nil, zero value otherwise.

### GetSfAccountOk

`func (o *DbTargetDetails) GetSfAccountOk() (*string, bool)`

GetSfAccountOk returns a tuple with the SfAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfAccount

`func (o *DbTargetDetails) SetSfAccount(v string)`

SetSfAccount sets SfAccount field to given value.

### HasSfAccount

`func (o *DbTargetDetails) HasSfAccount() bool`

HasSfAccount returns a boolean if a field has been set.

### GetSslConnectionCertificate

`func (o *DbTargetDetails) GetSslConnectionCertificate() string`

GetSslConnectionCertificate returns the SslConnectionCertificate field if non-nil, zero value otherwise.

### GetSslConnectionCertificateOk

`func (o *DbTargetDetails) GetSslConnectionCertificateOk() (*string, bool)`

GetSslConnectionCertificateOk returns a tuple with the SslConnectionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnectionCertificate

`func (o *DbTargetDetails) SetSslConnectionCertificate(v string)`

SetSslConnectionCertificate sets SslConnectionCertificate field to given value.

### HasSslConnectionCertificate

`func (o *DbTargetDetails) HasSslConnectionCertificate() bool`

HasSslConnectionCertificate returns a boolean if a field has been set.

### GetSslConnectionMode

`func (o *DbTargetDetails) GetSslConnectionMode() bool`

GetSslConnectionMode returns the SslConnectionMode field if non-nil, zero value otherwise.

### GetSslConnectionModeOk

`func (o *DbTargetDetails) GetSslConnectionModeOk() (*bool, bool)`

GetSslConnectionModeOk returns a tuple with the SslConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnectionMode

`func (o *DbTargetDetails) SetSslConnectionMode(v bool)`

SetSslConnectionMode sets SslConnectionMode field to given value.

### HasSslConnectionMode

`func (o *DbTargetDetails) HasSslConnectionMode() bool`

HasSslConnectionMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


