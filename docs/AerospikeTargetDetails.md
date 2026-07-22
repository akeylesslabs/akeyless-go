# AerospikeTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AerospikeAdminUsername** | Pointer to **string** |  | [optional] 
**AerospikeClientId** | Pointer to **string** |  | [optional] 
**AerospikeClientSecret** | Pointer to **string** |  | [optional] 
**AerospikeCloud** | Pointer to **bool** |  | [optional] 
**AerospikeClusterId** | Pointer to **string** |  | [optional] 
**AerospikeHostname** | Pointer to **string** |  | [optional] 
**AerospikeNamespace** | Pointer to **string** |  | [optional] 
**AerospikePassword** | Pointer to **string** |  | [optional] 
**AerospikePort** | Pointer to **string** |  | [optional] 
**ClientCertificate** | Pointer to **string** | (Optional) ClientCertificate defines the client certificate for mutual TLS. Must be base64 certificate loaded by UI using file loader field | [optional] 
**ClientKeyPassphrase** | Pointer to **string** | (Optional) ClientKeyPassphrase defines the passphrase for the client private key | [optional] 
**ClientPrivateKey** | Pointer to **string** | (Optional) ClientPrivateKey defines the client private key for mutual TLS. Must be base64 private key loaded by UI using file loader field | [optional] 
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
**EnableMtls** | Pointer to **bool** | (Optional) EnableMTLS defines if mutual TLS will be used to connect to DB | [optional] 
**OracleWalletDetails** | Pointer to [**WalletDetails**](WalletDetails.md) |  | [optional] 
**SfAccount** | Pointer to **string** |  | [optional] 
**SkipServerNameValidation** | Pointer to **string** | (Optional) SkipServerNameValidation disables server name verification while still validating the certificate chain. Postgres treats empty as legacy \&quot;skip hostname validation\&quot;; MySQL treats empty as false. | [optional] 
**SslConnectionCertificate** | Pointer to **string** | (Optional) SSLConnectionCertificate defines the certificate for SSL connection. Must be base64 certificate loaded by UI using file loader field | [optional] 
**SslConnectionMode** | Pointer to **bool** | (Optional) SSLConnectionMode defines if SSL mode will be used to connect to DB | [optional] 

## Methods

### NewAerospikeTargetDetails

`func NewAerospikeTargetDetails() *AerospikeTargetDetails`

NewAerospikeTargetDetails instantiates a new AerospikeTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAerospikeTargetDetailsWithDefaults

`func NewAerospikeTargetDetailsWithDefaults() *AerospikeTargetDetails`

NewAerospikeTargetDetailsWithDefaults instantiates a new AerospikeTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAerospikeAdminUsername

`func (o *AerospikeTargetDetails) GetAerospikeAdminUsername() string`

GetAerospikeAdminUsername returns the AerospikeAdminUsername field if non-nil, zero value otherwise.

### GetAerospikeAdminUsernameOk

`func (o *AerospikeTargetDetails) GetAerospikeAdminUsernameOk() (*string, bool)`

GetAerospikeAdminUsernameOk returns a tuple with the AerospikeAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeAdminUsername

`func (o *AerospikeTargetDetails) SetAerospikeAdminUsername(v string)`

SetAerospikeAdminUsername sets AerospikeAdminUsername field to given value.

### HasAerospikeAdminUsername

`func (o *AerospikeTargetDetails) HasAerospikeAdminUsername() bool`

HasAerospikeAdminUsername returns a boolean if a field has been set.

### GetAerospikeClientId

`func (o *AerospikeTargetDetails) GetAerospikeClientId() string`

GetAerospikeClientId returns the AerospikeClientId field if non-nil, zero value otherwise.

### GetAerospikeClientIdOk

`func (o *AerospikeTargetDetails) GetAerospikeClientIdOk() (*string, bool)`

GetAerospikeClientIdOk returns a tuple with the AerospikeClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientId

`func (o *AerospikeTargetDetails) SetAerospikeClientId(v string)`

SetAerospikeClientId sets AerospikeClientId field to given value.

### HasAerospikeClientId

`func (o *AerospikeTargetDetails) HasAerospikeClientId() bool`

HasAerospikeClientId returns a boolean if a field has been set.

### GetAerospikeClientSecret

`func (o *AerospikeTargetDetails) GetAerospikeClientSecret() string`

GetAerospikeClientSecret returns the AerospikeClientSecret field if non-nil, zero value otherwise.

### GetAerospikeClientSecretOk

`func (o *AerospikeTargetDetails) GetAerospikeClientSecretOk() (*string, bool)`

GetAerospikeClientSecretOk returns a tuple with the AerospikeClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClientSecret

`func (o *AerospikeTargetDetails) SetAerospikeClientSecret(v string)`

SetAerospikeClientSecret sets AerospikeClientSecret field to given value.

### HasAerospikeClientSecret

`func (o *AerospikeTargetDetails) HasAerospikeClientSecret() bool`

HasAerospikeClientSecret returns a boolean if a field has been set.

### GetAerospikeCloud

`func (o *AerospikeTargetDetails) GetAerospikeCloud() bool`

GetAerospikeCloud returns the AerospikeCloud field if non-nil, zero value otherwise.

### GetAerospikeCloudOk

`func (o *AerospikeTargetDetails) GetAerospikeCloudOk() (*bool, bool)`

GetAerospikeCloudOk returns a tuple with the AerospikeCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeCloud

`func (o *AerospikeTargetDetails) SetAerospikeCloud(v bool)`

SetAerospikeCloud sets AerospikeCloud field to given value.

### HasAerospikeCloud

`func (o *AerospikeTargetDetails) HasAerospikeCloud() bool`

HasAerospikeCloud returns a boolean if a field has been set.

### GetAerospikeClusterId

`func (o *AerospikeTargetDetails) GetAerospikeClusterId() string`

GetAerospikeClusterId returns the AerospikeClusterId field if non-nil, zero value otherwise.

### GetAerospikeClusterIdOk

`func (o *AerospikeTargetDetails) GetAerospikeClusterIdOk() (*string, bool)`

GetAerospikeClusterIdOk returns a tuple with the AerospikeClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeClusterId

`func (o *AerospikeTargetDetails) SetAerospikeClusterId(v string)`

SetAerospikeClusterId sets AerospikeClusterId field to given value.

### HasAerospikeClusterId

`func (o *AerospikeTargetDetails) HasAerospikeClusterId() bool`

HasAerospikeClusterId returns a boolean if a field has been set.

### GetAerospikeHostname

`func (o *AerospikeTargetDetails) GetAerospikeHostname() string`

GetAerospikeHostname returns the AerospikeHostname field if non-nil, zero value otherwise.

### GetAerospikeHostnameOk

`func (o *AerospikeTargetDetails) GetAerospikeHostnameOk() (*string, bool)`

GetAerospikeHostnameOk returns a tuple with the AerospikeHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeHostname

`func (o *AerospikeTargetDetails) SetAerospikeHostname(v string)`

SetAerospikeHostname sets AerospikeHostname field to given value.

### HasAerospikeHostname

`func (o *AerospikeTargetDetails) HasAerospikeHostname() bool`

HasAerospikeHostname returns a boolean if a field has been set.

### GetAerospikeNamespace

`func (o *AerospikeTargetDetails) GetAerospikeNamespace() string`

GetAerospikeNamespace returns the AerospikeNamespace field if non-nil, zero value otherwise.

### GetAerospikeNamespaceOk

`func (o *AerospikeTargetDetails) GetAerospikeNamespaceOk() (*string, bool)`

GetAerospikeNamespaceOk returns a tuple with the AerospikeNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikeNamespace

`func (o *AerospikeTargetDetails) SetAerospikeNamespace(v string)`

SetAerospikeNamespace sets AerospikeNamespace field to given value.

### HasAerospikeNamespace

`func (o *AerospikeTargetDetails) HasAerospikeNamespace() bool`

HasAerospikeNamespace returns a boolean if a field has been set.

### GetAerospikePassword

`func (o *AerospikeTargetDetails) GetAerospikePassword() string`

GetAerospikePassword returns the AerospikePassword field if non-nil, zero value otherwise.

### GetAerospikePasswordOk

`func (o *AerospikeTargetDetails) GetAerospikePasswordOk() (*string, bool)`

GetAerospikePasswordOk returns a tuple with the AerospikePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikePassword

`func (o *AerospikeTargetDetails) SetAerospikePassword(v string)`

SetAerospikePassword sets AerospikePassword field to given value.

### HasAerospikePassword

`func (o *AerospikeTargetDetails) HasAerospikePassword() bool`

HasAerospikePassword returns a boolean if a field has been set.

### GetAerospikePort

`func (o *AerospikeTargetDetails) GetAerospikePort() string`

GetAerospikePort returns the AerospikePort field if non-nil, zero value otherwise.

### GetAerospikePortOk

`func (o *AerospikeTargetDetails) GetAerospikePortOk() (*string, bool)`

GetAerospikePortOk returns a tuple with the AerospikePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAerospikePort

`func (o *AerospikeTargetDetails) SetAerospikePort(v string)`

SetAerospikePort sets AerospikePort field to given value.

### HasAerospikePort

`func (o *AerospikeTargetDetails) HasAerospikePort() bool`

HasAerospikePort returns a boolean if a field has been set.

### GetClientCertificate

`func (o *AerospikeTargetDetails) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *AerospikeTargetDetails) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *AerospikeTargetDetails) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *AerospikeTargetDetails) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetClientKeyPassphrase

`func (o *AerospikeTargetDetails) GetClientKeyPassphrase() string`

GetClientKeyPassphrase returns the ClientKeyPassphrase field if non-nil, zero value otherwise.

### GetClientKeyPassphraseOk

`func (o *AerospikeTargetDetails) GetClientKeyPassphraseOk() (*string, bool)`

GetClientKeyPassphraseOk returns a tuple with the ClientKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKeyPassphrase

`func (o *AerospikeTargetDetails) SetClientKeyPassphrase(v string)`

SetClientKeyPassphrase sets ClientKeyPassphrase field to given value.

### HasClientKeyPassphrase

`func (o *AerospikeTargetDetails) HasClientKeyPassphrase() bool`

HasClientKeyPassphrase returns a boolean if a field has been set.

### GetClientPrivateKey

`func (o *AerospikeTargetDetails) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *AerospikeTargetDetails) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *AerospikeTargetDetails) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *AerospikeTargetDetails) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### GetCloudServiceProvider

`func (o *AerospikeTargetDetails) GetCloudServiceProvider() string`

GetCloudServiceProvider returns the CloudServiceProvider field if non-nil, zero value otherwise.

### GetCloudServiceProviderOk

`func (o *AerospikeTargetDetails) GetCloudServiceProviderOk() (*string, bool)`

GetCloudServiceProviderOk returns a tuple with the CloudServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudServiceProvider

`func (o *AerospikeTargetDetails) SetCloudServiceProvider(v string)`

SetCloudServiceProvider sets CloudServiceProvider field to given value.

### HasCloudServiceProvider

`func (o *AerospikeTargetDetails) HasCloudServiceProvider() bool`

HasCloudServiceProvider returns a boolean if a field has been set.

### GetClusterMode

`func (o *AerospikeTargetDetails) GetClusterMode() bool`

GetClusterMode returns the ClusterMode field if non-nil, zero value otherwise.

### GetClusterModeOk

`func (o *AerospikeTargetDetails) GetClusterModeOk() (*bool, bool)`

GetClusterModeOk returns a tuple with the ClusterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMode

`func (o *AerospikeTargetDetails) SetClusterMode(v bool)`

SetClusterMode sets ClusterMode field to given value.

### HasClusterMode

`func (o *AerospikeTargetDetails) HasClusterMode() bool`

HasClusterMode returns a boolean if a field has been set.

### GetConnectionType

`func (o *AerospikeTargetDetails) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AerospikeTargetDetails) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AerospikeTargetDetails) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *AerospikeTargetDetails) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetDbClientId

`func (o *AerospikeTargetDetails) GetDbClientId() string`

GetDbClientId returns the DbClientId field if non-nil, zero value otherwise.

### GetDbClientIdOk

`func (o *AerospikeTargetDetails) GetDbClientIdOk() (*string, bool)`

GetDbClientIdOk returns a tuple with the DbClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbClientId

`func (o *AerospikeTargetDetails) SetDbClientId(v string)`

SetDbClientId sets DbClientId field to given value.

### HasDbClientId

`func (o *AerospikeTargetDetails) HasDbClientId() bool`

HasDbClientId returns a boolean if a field has been set.

### GetDbClientSecret

`func (o *AerospikeTargetDetails) GetDbClientSecret() string`

GetDbClientSecret returns the DbClientSecret field if non-nil, zero value otherwise.

### GetDbClientSecretOk

`func (o *AerospikeTargetDetails) GetDbClientSecretOk() (*string, bool)`

GetDbClientSecretOk returns a tuple with the DbClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbClientSecret

`func (o *AerospikeTargetDetails) SetDbClientSecret(v string)`

SetDbClientSecret sets DbClientSecret field to given value.

### HasDbClientSecret

`func (o *AerospikeTargetDetails) HasDbClientSecret() bool`

HasDbClientSecret returns a boolean if a field has been set.

### GetDbHostName

`func (o *AerospikeTargetDetails) GetDbHostName() string`

GetDbHostName returns the DbHostName field if non-nil, zero value otherwise.

### GetDbHostNameOk

`func (o *AerospikeTargetDetails) GetDbHostNameOk() (*string, bool)`

GetDbHostNameOk returns a tuple with the DbHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHostName

`func (o *AerospikeTargetDetails) SetDbHostName(v string)`

SetDbHostName sets DbHostName field to given value.

### HasDbHostName

`func (o *AerospikeTargetDetails) HasDbHostName() bool`

HasDbHostName returns a boolean if a field has been set.

### GetDbName

`func (o *AerospikeTargetDetails) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *AerospikeTargetDetails) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *AerospikeTargetDetails) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *AerospikeTargetDetails) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbPort

`func (o *AerospikeTargetDetails) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *AerospikeTargetDetails) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *AerospikeTargetDetails) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.

### HasDbPort

`func (o *AerospikeTargetDetails) HasDbPort() bool`

HasDbPort returns a boolean if a field has been set.

### GetDbPrivateKey

`func (o *AerospikeTargetDetails) GetDbPrivateKey() string`

GetDbPrivateKey returns the DbPrivateKey field if non-nil, zero value otherwise.

### GetDbPrivateKeyOk

`func (o *AerospikeTargetDetails) GetDbPrivateKeyOk() (*string, bool)`

GetDbPrivateKeyOk returns a tuple with the DbPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPrivateKey

`func (o *AerospikeTargetDetails) SetDbPrivateKey(v string)`

SetDbPrivateKey sets DbPrivateKey field to given value.

### HasDbPrivateKey

`func (o *AerospikeTargetDetails) HasDbPrivateKey() bool`

HasDbPrivateKey returns a boolean if a field has been set.

### GetDbPrivateKeyPassphrase

`func (o *AerospikeTargetDetails) GetDbPrivateKeyPassphrase() string`

GetDbPrivateKeyPassphrase returns the DbPrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetDbPrivateKeyPassphraseOk

`func (o *AerospikeTargetDetails) GetDbPrivateKeyPassphraseOk() (*string, bool)`

GetDbPrivateKeyPassphraseOk returns a tuple with the DbPrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPrivateKeyPassphrase

`func (o *AerospikeTargetDetails) SetDbPrivateKeyPassphrase(v string)`

SetDbPrivateKeyPassphrase sets DbPrivateKeyPassphrase field to given value.

### HasDbPrivateKeyPassphrase

`func (o *AerospikeTargetDetails) HasDbPrivateKeyPassphrase() bool`

HasDbPrivateKeyPassphrase returns a boolean if a field has been set.

### GetDbPwd

`func (o *AerospikeTargetDetails) GetDbPwd() string`

GetDbPwd returns the DbPwd field if non-nil, zero value otherwise.

### GetDbPwdOk

`func (o *AerospikeTargetDetails) GetDbPwdOk() (*string, bool)`

GetDbPwdOk returns a tuple with the DbPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPwd

`func (o *AerospikeTargetDetails) SetDbPwd(v string)`

SetDbPwd sets DbPwd field to given value.

### HasDbPwd

`func (o *AerospikeTargetDetails) HasDbPwd() bool`

HasDbPwd returns a boolean if a field has been set.

### GetDbServerCertificates

`func (o *AerospikeTargetDetails) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *AerospikeTargetDetails) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *AerospikeTargetDetails) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *AerospikeTargetDetails) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *AerospikeTargetDetails) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *AerospikeTargetDetails) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *AerospikeTargetDetails) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *AerospikeTargetDetails) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDbTenantId

`func (o *AerospikeTargetDetails) GetDbTenantId() string`

GetDbTenantId returns the DbTenantId field if non-nil, zero value otherwise.

### GetDbTenantIdOk

`func (o *AerospikeTargetDetails) GetDbTenantIdOk() (*string, bool)`

GetDbTenantIdOk returns a tuple with the DbTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTenantId

`func (o *AerospikeTargetDetails) SetDbTenantId(v string)`

SetDbTenantId sets DbTenantId field to given value.

### HasDbTenantId

`func (o *AerospikeTargetDetails) HasDbTenantId() bool`

HasDbTenantId returns a boolean if a field has been set.

### GetDbUserName

`func (o *AerospikeTargetDetails) GetDbUserName() string`

GetDbUserName returns the DbUserName field if non-nil, zero value otherwise.

### GetDbUserNameOk

`func (o *AerospikeTargetDetails) GetDbUserNameOk() (*string, bool)`

GetDbUserNameOk returns a tuple with the DbUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUserName

`func (o *AerospikeTargetDetails) SetDbUserName(v string)`

SetDbUserName sets DbUserName field to given value.

### HasDbUserName

`func (o *AerospikeTargetDetails) HasDbUserName() bool`

HasDbUserName returns a boolean if a field has been set.

### GetEnableMtls

`func (o *AerospikeTargetDetails) GetEnableMtls() bool`

GetEnableMtls returns the EnableMtls field if non-nil, zero value otherwise.

### GetEnableMtlsOk

`func (o *AerospikeTargetDetails) GetEnableMtlsOk() (*bool, bool)`

GetEnableMtlsOk returns a tuple with the EnableMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMtls

`func (o *AerospikeTargetDetails) SetEnableMtls(v bool)`

SetEnableMtls sets EnableMtls field to given value.

### HasEnableMtls

`func (o *AerospikeTargetDetails) HasEnableMtls() bool`

HasEnableMtls returns a boolean if a field has been set.

### GetOracleWalletDetails

`func (o *AerospikeTargetDetails) GetOracleWalletDetails() WalletDetails`

GetOracleWalletDetails returns the OracleWalletDetails field if non-nil, zero value otherwise.

### GetOracleWalletDetailsOk

`func (o *AerospikeTargetDetails) GetOracleWalletDetailsOk() (*WalletDetails, bool)`

GetOracleWalletDetailsOk returns a tuple with the OracleWalletDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleWalletDetails

`func (o *AerospikeTargetDetails) SetOracleWalletDetails(v WalletDetails)`

SetOracleWalletDetails sets OracleWalletDetails field to given value.

### HasOracleWalletDetails

`func (o *AerospikeTargetDetails) HasOracleWalletDetails() bool`

HasOracleWalletDetails returns a boolean if a field has been set.

### GetSfAccount

`func (o *AerospikeTargetDetails) GetSfAccount() string`

GetSfAccount returns the SfAccount field if non-nil, zero value otherwise.

### GetSfAccountOk

`func (o *AerospikeTargetDetails) GetSfAccountOk() (*string, bool)`

GetSfAccountOk returns a tuple with the SfAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfAccount

`func (o *AerospikeTargetDetails) SetSfAccount(v string)`

SetSfAccount sets SfAccount field to given value.

### HasSfAccount

`func (o *AerospikeTargetDetails) HasSfAccount() bool`

HasSfAccount returns a boolean if a field has been set.

### GetSkipServerNameValidation

`func (o *AerospikeTargetDetails) GetSkipServerNameValidation() string`

GetSkipServerNameValidation returns the SkipServerNameValidation field if non-nil, zero value otherwise.

### GetSkipServerNameValidationOk

`func (o *AerospikeTargetDetails) GetSkipServerNameValidationOk() (*string, bool)`

GetSkipServerNameValidationOk returns a tuple with the SkipServerNameValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipServerNameValidation

`func (o *AerospikeTargetDetails) SetSkipServerNameValidation(v string)`

SetSkipServerNameValidation sets SkipServerNameValidation field to given value.

### HasSkipServerNameValidation

`func (o *AerospikeTargetDetails) HasSkipServerNameValidation() bool`

HasSkipServerNameValidation returns a boolean if a field has been set.

### GetSslConnectionCertificate

`func (o *AerospikeTargetDetails) GetSslConnectionCertificate() string`

GetSslConnectionCertificate returns the SslConnectionCertificate field if non-nil, zero value otherwise.

### GetSslConnectionCertificateOk

`func (o *AerospikeTargetDetails) GetSslConnectionCertificateOk() (*string, bool)`

GetSslConnectionCertificateOk returns a tuple with the SslConnectionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnectionCertificate

`func (o *AerospikeTargetDetails) SetSslConnectionCertificate(v string)`

SetSslConnectionCertificate sets SslConnectionCertificate field to given value.

### HasSslConnectionCertificate

`func (o *AerospikeTargetDetails) HasSslConnectionCertificate() bool`

HasSslConnectionCertificate returns a boolean if a field has been set.

### GetSslConnectionMode

`func (o *AerospikeTargetDetails) GetSslConnectionMode() bool`

GetSslConnectionMode returns the SslConnectionMode field if non-nil, zero value otherwise.

### GetSslConnectionModeOk

`func (o *AerospikeTargetDetails) GetSslConnectionModeOk() (*bool, bool)`

GetSslConnectionModeOk returns a tuple with the SslConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnectionMode

`func (o *AerospikeTargetDetails) SetSslConnectionMode(v bool)`

SetSslConnectionMode sets SslConnectionMode field to given value.

### HasSslConnectionMode

`func (o *AerospikeTargetDetails) HasSslConnectionMode() bool`

HasSslConnectionMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


