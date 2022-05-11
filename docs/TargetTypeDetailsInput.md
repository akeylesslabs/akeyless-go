# TargetTypeDetailsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactoryAdminApikey** | Pointer to **string** |  | [optional] 
**ArtifactoryAdminUsername** | Pointer to **string** |  | [optional] 
**ArtifactoryBaseUrl** | Pointer to **string** |  | [optional] 
**AwsAccessKeyId** | Pointer to **string** |  | [optional] 
**AwsRegion** | Pointer to **string** |  | [optional] 
**AwsSecretAccessKey** | Pointer to **string** |  | [optional] 
**AwsSessionToken** | Pointer to **string** |  | [optional] 
**AzureClientId** | Pointer to **string** |  | [optional] 
**AzureClientSecret** | Pointer to **string** |  | [optional] 
**AzureTenantId** | Pointer to **string** |  | [optional] 
**CaCertData** | Pointer to **[]int32** | CACertData is the rsa 4096 certificate data in PEM format | [optional] 
**CaCertName** | Pointer to **string** | CACertName is the name of the certificate in SalesForce tenant | [optional] 
**ChefServerHostName** | Pointer to **string** |  | [optional] 
**ChefServerKey** | Pointer to **string** |  | [optional] 
**ChefServerPort** | Pointer to **string** |  | [optional] 
**ChefServerUrl** | Pointer to **string** |  | [optional] 
**ChefServerUsername** | Pointer to **string** |  | [optional] 
**ChefSkipSsl** | Pointer to **bool** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**DbHostName** | Pointer to **string** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**DbPort** | Pointer to **string** |  | [optional] 
**DbPwd** | Pointer to **string** |  | [optional] 
**DbServerCertificates** | Pointer to **string** | (Optional) DBServerCertificates defines the set of root certificate authorities that clients use when verifying server certificates. If DBServerCertificates is empty, TLS uses the host&#39;s root CA set. | [optional] 
**DbServerName** | Pointer to **string** | (Optional) ServerName is used to verify the hostname on the returned certificates unless InsecureSkipVerify is given. It is also included in the client&#39;s handshake to support virtual hosting unless it is an IP address. | [optional] 
**DbUserName** | Pointer to **string** |  | [optional] 
**EksAccessKeyId** | Pointer to **string** |  | [optional] 
**EksClusterCaCertificate** | Pointer to **string** |  | [optional] 
**EksClusterEndpoint** | Pointer to **string** |  | [optional] 
**EksClusterName** | Pointer to **string** |  | [optional] 
**EksRegion** | Pointer to **string** |  | [optional] 
**EksSecretAccessKey** | Pointer to **string** |  | [optional] 
**GcpServiceAccountEmail** | Pointer to **string** |  | [optional] 
**GcpServiceAccountKey** | Pointer to **string** |  | [optional] 
**GithubAppId** | Pointer to **int64** |  | [optional] 
**GithubAppPrivateKey** | Pointer to **string** |  | [optional] 
**GithubBaseUrl** | Pointer to **string** |  | [optional] 
**GkeClusterCaCertificate** | Pointer to **string** |  | [optional] 
**GkeClusterEndpoint** | Pointer to **string** |  | [optional] 
**GkeClusterName** | Pointer to **string** |  | [optional] 
**GkeServiceAccountKey** | Pointer to **string** |  | [optional] 
**GkeServiceAccountName** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**K8sBearerToken** | Pointer to **string** |  | [optional] 
**K8sClusterCaCertificate** | Pointer to **string** |  | [optional] 
**K8sClusterEndpoint** | Pointer to **string** |  | [optional] 
**LdapAudience** | Pointer to **string** |  | [optional] 
**LdapBindDn** | Pointer to **string** |  | [optional] 
**LdapBindPassword** | Pointer to **string** |  | [optional] 
**LdapCertificate** | Pointer to **string** |  | [optional] 
**LdapTokenExpiration** | Pointer to **string** |  | [optional] 
**LdapUrl** | Pointer to **string** |  | [optional] 
**MongodbAtlasApiPrivateKey** | Pointer to **string** |  | [optional] 
**MongodbAtlasApiPublicKey** | Pointer to **string** |  | [optional] 
**MongodbAtlasProjectId** | Pointer to **string** | mongodb atlas fields | [optional] 
**MongodbDbName** | Pointer to **string** | common fields | [optional] 
**MongodbDefaultAuthDb** | Pointer to **string** |  | [optional] 
**MongodbHostPort** | Pointer to **string** |  | [optional] 
**MongodbIsAtlas** | Pointer to **bool** |  | [optional] 
**MongodbPassword** | Pointer to **string** |  | [optional] 
**MongodbUriConnection** | Pointer to **string** | mongodb fields | [optional] 
**MongodbUriOptions** | Pointer to **string** |  | [optional] 
**MongodbUsername** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**PrivateKeyPassword** | Pointer to **string** |  | [optional] 
**RabbitmqServerPassword** | Pointer to **string** |  | [optional] 
**RabbitmqServerUri** | Pointer to **string** |  | [optional] 
**RabbitmqServerUser** | Pointer to **string** |  | [optional] 
**SecurityToken** | Pointer to **string** |  | [optional] 
**SfAccount** | Pointer to **string** |  | [optional] 
**SslConnectionCertificate** | Pointer to **string** | (Optional) SSLConnectionCertificate defines the certificate for SSL connection. Must be base64 certificate loaded by UI using file loader field | [optional] 
**SslConnectionMode** | Pointer to **bool** | (Optional) SSLConnectionMode defines if SSL mode will be used to connect to DB | [optional] 
**TenantUrl** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**VenafiApiKey** | Pointer to **string** |  | [optional] 
**VenafiBaseUrl** | Pointer to **string** |  | [optional] 
**VenafiTppPassword** | Pointer to **string** |  | [optional] 
**VenafiTppUsername** | Pointer to **string** |  | [optional] 
**VenafiUseTpp** | Pointer to **bool** |  | [optional] 
**VenafiZone** | Pointer to **string** |  | [optional] 

## Methods

### NewTargetTypeDetailsInput

`func NewTargetTypeDetailsInput() *TargetTypeDetailsInput`

NewTargetTypeDetailsInput instantiates a new TargetTypeDetailsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetTypeDetailsInputWithDefaults

`func NewTargetTypeDetailsInputWithDefaults() *TargetTypeDetailsInput`

NewTargetTypeDetailsInputWithDefaults instantiates a new TargetTypeDetailsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactoryAdminApikey

`func (o *TargetTypeDetailsInput) GetArtifactoryAdminApikey() string`

GetArtifactoryAdminApikey returns the ArtifactoryAdminApikey field if non-nil, zero value otherwise.

### GetArtifactoryAdminApikeyOk

`func (o *TargetTypeDetailsInput) GetArtifactoryAdminApikeyOk() (*string, bool)`

GetArtifactoryAdminApikeyOk returns a tuple with the ArtifactoryAdminApikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminApikey

`func (o *TargetTypeDetailsInput) SetArtifactoryAdminApikey(v string)`

SetArtifactoryAdminApikey sets ArtifactoryAdminApikey field to given value.

### HasArtifactoryAdminApikey

`func (o *TargetTypeDetailsInput) HasArtifactoryAdminApikey() bool`

HasArtifactoryAdminApikey returns a boolean if a field has been set.

### GetArtifactoryAdminUsername

`func (o *TargetTypeDetailsInput) GetArtifactoryAdminUsername() string`

GetArtifactoryAdminUsername returns the ArtifactoryAdminUsername field if non-nil, zero value otherwise.

### GetArtifactoryAdminUsernameOk

`func (o *TargetTypeDetailsInput) GetArtifactoryAdminUsernameOk() (*string, bool)`

GetArtifactoryAdminUsernameOk returns a tuple with the ArtifactoryAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminUsername

`func (o *TargetTypeDetailsInput) SetArtifactoryAdminUsername(v string)`

SetArtifactoryAdminUsername sets ArtifactoryAdminUsername field to given value.

### HasArtifactoryAdminUsername

`func (o *TargetTypeDetailsInput) HasArtifactoryAdminUsername() bool`

HasArtifactoryAdminUsername returns a boolean if a field has been set.

### GetArtifactoryBaseUrl

`func (o *TargetTypeDetailsInput) GetArtifactoryBaseUrl() string`

GetArtifactoryBaseUrl returns the ArtifactoryBaseUrl field if non-nil, zero value otherwise.

### GetArtifactoryBaseUrlOk

`func (o *TargetTypeDetailsInput) GetArtifactoryBaseUrlOk() (*string, bool)`

GetArtifactoryBaseUrlOk returns a tuple with the ArtifactoryBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryBaseUrl

`func (o *TargetTypeDetailsInput) SetArtifactoryBaseUrl(v string)`

SetArtifactoryBaseUrl sets ArtifactoryBaseUrl field to given value.

### HasArtifactoryBaseUrl

`func (o *TargetTypeDetailsInput) HasArtifactoryBaseUrl() bool`

HasArtifactoryBaseUrl returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *TargetTypeDetailsInput) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *TargetTypeDetailsInput) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *TargetTypeDetailsInput) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *TargetTypeDetailsInput) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *TargetTypeDetailsInput) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *TargetTypeDetailsInput) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *TargetTypeDetailsInput) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *TargetTypeDetailsInput) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *TargetTypeDetailsInput) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *TargetTypeDetailsInput) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *TargetTypeDetailsInput) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *TargetTypeDetailsInput) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsSessionToken

`func (o *TargetTypeDetailsInput) GetAwsSessionToken() string`

GetAwsSessionToken returns the AwsSessionToken field if non-nil, zero value otherwise.

### GetAwsSessionTokenOk

`func (o *TargetTypeDetailsInput) GetAwsSessionTokenOk() (*string, bool)`

GetAwsSessionTokenOk returns a tuple with the AwsSessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSessionToken

`func (o *TargetTypeDetailsInput) SetAwsSessionToken(v string)`

SetAwsSessionToken sets AwsSessionToken field to given value.

### HasAwsSessionToken

`func (o *TargetTypeDetailsInput) HasAwsSessionToken() bool`

HasAwsSessionToken returns a boolean if a field has been set.

### GetAzureClientId

`func (o *TargetTypeDetailsInput) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *TargetTypeDetailsInput) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *TargetTypeDetailsInput) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *TargetTypeDetailsInput) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureClientSecret

`func (o *TargetTypeDetailsInput) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *TargetTypeDetailsInput) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *TargetTypeDetailsInput) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### HasAzureClientSecret

`func (o *TargetTypeDetailsInput) HasAzureClientSecret() bool`

HasAzureClientSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *TargetTypeDetailsInput) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *TargetTypeDetailsInput) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *TargetTypeDetailsInput) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *TargetTypeDetailsInput) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetCaCertData

`func (o *TargetTypeDetailsInput) GetCaCertData() []int32`

GetCaCertData returns the CaCertData field if non-nil, zero value otherwise.

### GetCaCertDataOk

`func (o *TargetTypeDetailsInput) GetCaCertDataOk() (*[]int32, bool)`

GetCaCertDataOk returns a tuple with the CaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertData

`func (o *TargetTypeDetailsInput) SetCaCertData(v []int32)`

SetCaCertData sets CaCertData field to given value.

### HasCaCertData

`func (o *TargetTypeDetailsInput) HasCaCertData() bool`

HasCaCertData returns a boolean if a field has been set.

### GetCaCertName

`func (o *TargetTypeDetailsInput) GetCaCertName() string`

GetCaCertName returns the CaCertName field if non-nil, zero value otherwise.

### GetCaCertNameOk

`func (o *TargetTypeDetailsInput) GetCaCertNameOk() (*string, bool)`

GetCaCertNameOk returns a tuple with the CaCertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertName

`func (o *TargetTypeDetailsInput) SetCaCertName(v string)`

SetCaCertName sets CaCertName field to given value.

### HasCaCertName

`func (o *TargetTypeDetailsInput) HasCaCertName() bool`

HasCaCertName returns a boolean if a field has been set.

### GetChefServerHostName

`func (o *TargetTypeDetailsInput) GetChefServerHostName() string`

GetChefServerHostName returns the ChefServerHostName field if non-nil, zero value otherwise.

### GetChefServerHostNameOk

`func (o *TargetTypeDetailsInput) GetChefServerHostNameOk() (*string, bool)`

GetChefServerHostNameOk returns a tuple with the ChefServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerHostName

`func (o *TargetTypeDetailsInput) SetChefServerHostName(v string)`

SetChefServerHostName sets ChefServerHostName field to given value.

### HasChefServerHostName

`func (o *TargetTypeDetailsInput) HasChefServerHostName() bool`

HasChefServerHostName returns a boolean if a field has been set.

### GetChefServerKey

`func (o *TargetTypeDetailsInput) GetChefServerKey() string`

GetChefServerKey returns the ChefServerKey field if non-nil, zero value otherwise.

### GetChefServerKeyOk

`func (o *TargetTypeDetailsInput) GetChefServerKeyOk() (*string, bool)`

GetChefServerKeyOk returns a tuple with the ChefServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerKey

`func (o *TargetTypeDetailsInput) SetChefServerKey(v string)`

SetChefServerKey sets ChefServerKey field to given value.

### HasChefServerKey

`func (o *TargetTypeDetailsInput) HasChefServerKey() bool`

HasChefServerKey returns a boolean if a field has been set.

### GetChefServerPort

`func (o *TargetTypeDetailsInput) GetChefServerPort() string`

GetChefServerPort returns the ChefServerPort field if non-nil, zero value otherwise.

### GetChefServerPortOk

`func (o *TargetTypeDetailsInput) GetChefServerPortOk() (*string, bool)`

GetChefServerPortOk returns a tuple with the ChefServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerPort

`func (o *TargetTypeDetailsInput) SetChefServerPort(v string)`

SetChefServerPort sets ChefServerPort field to given value.

### HasChefServerPort

`func (o *TargetTypeDetailsInput) HasChefServerPort() bool`

HasChefServerPort returns a boolean if a field has been set.

### GetChefServerUrl

`func (o *TargetTypeDetailsInput) GetChefServerUrl() string`

GetChefServerUrl returns the ChefServerUrl field if non-nil, zero value otherwise.

### GetChefServerUrlOk

`func (o *TargetTypeDetailsInput) GetChefServerUrlOk() (*string, bool)`

GetChefServerUrlOk returns a tuple with the ChefServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUrl

`func (o *TargetTypeDetailsInput) SetChefServerUrl(v string)`

SetChefServerUrl sets ChefServerUrl field to given value.

### HasChefServerUrl

`func (o *TargetTypeDetailsInput) HasChefServerUrl() bool`

HasChefServerUrl returns a boolean if a field has been set.

### GetChefServerUsername

`func (o *TargetTypeDetailsInput) GetChefServerUsername() string`

GetChefServerUsername returns the ChefServerUsername field if non-nil, zero value otherwise.

### GetChefServerUsernameOk

`func (o *TargetTypeDetailsInput) GetChefServerUsernameOk() (*string, bool)`

GetChefServerUsernameOk returns a tuple with the ChefServerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUsername

`func (o *TargetTypeDetailsInput) SetChefServerUsername(v string)`

SetChefServerUsername sets ChefServerUsername field to given value.

### HasChefServerUsername

`func (o *TargetTypeDetailsInput) HasChefServerUsername() bool`

HasChefServerUsername returns a boolean if a field has been set.

### GetChefSkipSsl

`func (o *TargetTypeDetailsInput) GetChefSkipSsl() bool`

GetChefSkipSsl returns the ChefSkipSsl field if non-nil, zero value otherwise.

### GetChefSkipSslOk

`func (o *TargetTypeDetailsInput) GetChefSkipSslOk() (*bool, bool)`

GetChefSkipSslOk returns a tuple with the ChefSkipSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefSkipSsl

`func (o *TargetTypeDetailsInput) SetChefSkipSsl(v bool)`

SetChefSkipSsl sets ChefSkipSsl field to given value.

### HasChefSkipSsl

`func (o *TargetTypeDetailsInput) HasChefSkipSsl() bool`

HasChefSkipSsl returns a boolean if a field has been set.

### GetClientId

`func (o *TargetTypeDetailsInput) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TargetTypeDetailsInput) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TargetTypeDetailsInput) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TargetTypeDetailsInput) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *TargetTypeDetailsInput) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TargetTypeDetailsInput) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TargetTypeDetailsInput) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *TargetTypeDetailsInput) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDbHostName

`func (o *TargetTypeDetailsInput) GetDbHostName() string`

GetDbHostName returns the DbHostName field if non-nil, zero value otherwise.

### GetDbHostNameOk

`func (o *TargetTypeDetailsInput) GetDbHostNameOk() (*string, bool)`

GetDbHostNameOk returns a tuple with the DbHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHostName

`func (o *TargetTypeDetailsInput) SetDbHostName(v string)`

SetDbHostName sets DbHostName field to given value.

### HasDbHostName

`func (o *TargetTypeDetailsInput) HasDbHostName() bool`

HasDbHostName returns a boolean if a field has been set.

### GetDbName

`func (o *TargetTypeDetailsInput) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TargetTypeDetailsInput) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TargetTypeDetailsInput) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *TargetTypeDetailsInput) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbPort

`func (o *TargetTypeDetailsInput) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *TargetTypeDetailsInput) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *TargetTypeDetailsInput) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.

### HasDbPort

`func (o *TargetTypeDetailsInput) HasDbPort() bool`

HasDbPort returns a boolean if a field has been set.

### GetDbPwd

`func (o *TargetTypeDetailsInput) GetDbPwd() string`

GetDbPwd returns the DbPwd field if non-nil, zero value otherwise.

### GetDbPwdOk

`func (o *TargetTypeDetailsInput) GetDbPwdOk() (*string, bool)`

GetDbPwdOk returns a tuple with the DbPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPwd

`func (o *TargetTypeDetailsInput) SetDbPwd(v string)`

SetDbPwd sets DbPwd field to given value.

### HasDbPwd

`func (o *TargetTypeDetailsInput) HasDbPwd() bool`

HasDbPwd returns a boolean if a field has been set.

### GetDbServerCertificates

`func (o *TargetTypeDetailsInput) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *TargetTypeDetailsInput) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *TargetTypeDetailsInput) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *TargetTypeDetailsInput) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *TargetTypeDetailsInput) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *TargetTypeDetailsInput) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *TargetTypeDetailsInput) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *TargetTypeDetailsInput) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDbUserName

`func (o *TargetTypeDetailsInput) GetDbUserName() string`

GetDbUserName returns the DbUserName field if non-nil, zero value otherwise.

### GetDbUserNameOk

`func (o *TargetTypeDetailsInput) GetDbUserNameOk() (*string, bool)`

GetDbUserNameOk returns a tuple with the DbUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUserName

`func (o *TargetTypeDetailsInput) SetDbUserName(v string)`

SetDbUserName sets DbUserName field to given value.

### HasDbUserName

`func (o *TargetTypeDetailsInput) HasDbUserName() bool`

HasDbUserName returns a boolean if a field has been set.

### GetEksAccessKeyId

`func (o *TargetTypeDetailsInput) GetEksAccessKeyId() string`

GetEksAccessKeyId returns the EksAccessKeyId field if non-nil, zero value otherwise.

### GetEksAccessKeyIdOk

`func (o *TargetTypeDetailsInput) GetEksAccessKeyIdOk() (*string, bool)`

GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAccessKeyId

`func (o *TargetTypeDetailsInput) SetEksAccessKeyId(v string)`

SetEksAccessKeyId sets EksAccessKeyId field to given value.

### HasEksAccessKeyId

`func (o *TargetTypeDetailsInput) HasEksAccessKeyId() bool`

HasEksAccessKeyId returns a boolean if a field has been set.

### GetEksClusterCaCertificate

`func (o *TargetTypeDetailsInput) GetEksClusterCaCertificate() string`

GetEksClusterCaCertificate returns the EksClusterCaCertificate field if non-nil, zero value otherwise.

### GetEksClusterCaCertificateOk

`func (o *TargetTypeDetailsInput) GetEksClusterCaCertificateOk() (*string, bool)`

GetEksClusterCaCertificateOk returns a tuple with the EksClusterCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterCaCertificate

`func (o *TargetTypeDetailsInput) SetEksClusterCaCertificate(v string)`

SetEksClusterCaCertificate sets EksClusterCaCertificate field to given value.

### HasEksClusterCaCertificate

`func (o *TargetTypeDetailsInput) HasEksClusterCaCertificate() bool`

HasEksClusterCaCertificate returns a boolean if a field has been set.

### GetEksClusterEndpoint

`func (o *TargetTypeDetailsInput) GetEksClusterEndpoint() string`

GetEksClusterEndpoint returns the EksClusterEndpoint field if non-nil, zero value otherwise.

### GetEksClusterEndpointOk

`func (o *TargetTypeDetailsInput) GetEksClusterEndpointOk() (*string, bool)`

GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterEndpoint

`func (o *TargetTypeDetailsInput) SetEksClusterEndpoint(v string)`

SetEksClusterEndpoint sets EksClusterEndpoint field to given value.

### HasEksClusterEndpoint

`func (o *TargetTypeDetailsInput) HasEksClusterEndpoint() bool`

HasEksClusterEndpoint returns a boolean if a field has been set.

### GetEksClusterName

`func (o *TargetTypeDetailsInput) GetEksClusterName() string`

GetEksClusterName returns the EksClusterName field if non-nil, zero value otherwise.

### GetEksClusterNameOk

`func (o *TargetTypeDetailsInput) GetEksClusterNameOk() (*string, bool)`

GetEksClusterNameOk returns a tuple with the EksClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterName

`func (o *TargetTypeDetailsInput) SetEksClusterName(v string)`

SetEksClusterName sets EksClusterName field to given value.

### HasEksClusterName

`func (o *TargetTypeDetailsInput) HasEksClusterName() bool`

HasEksClusterName returns a boolean if a field has been set.

### GetEksRegion

`func (o *TargetTypeDetailsInput) GetEksRegion() string`

GetEksRegion returns the EksRegion field if non-nil, zero value otherwise.

### GetEksRegionOk

`func (o *TargetTypeDetailsInput) GetEksRegionOk() (*string, bool)`

GetEksRegionOk returns a tuple with the EksRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksRegion

`func (o *TargetTypeDetailsInput) SetEksRegion(v string)`

SetEksRegion sets EksRegion field to given value.

### HasEksRegion

`func (o *TargetTypeDetailsInput) HasEksRegion() bool`

HasEksRegion returns a boolean if a field has been set.

### GetEksSecretAccessKey

`func (o *TargetTypeDetailsInput) GetEksSecretAccessKey() string`

GetEksSecretAccessKey returns the EksSecretAccessKey field if non-nil, zero value otherwise.

### GetEksSecretAccessKeyOk

`func (o *TargetTypeDetailsInput) GetEksSecretAccessKeyOk() (*string, bool)`

GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSecretAccessKey

`func (o *TargetTypeDetailsInput) SetEksSecretAccessKey(v string)`

SetEksSecretAccessKey sets EksSecretAccessKey field to given value.

### HasEksSecretAccessKey

`func (o *TargetTypeDetailsInput) HasEksSecretAccessKey() bool`

HasEksSecretAccessKey returns a boolean if a field has been set.

### GetGcpServiceAccountEmail

`func (o *TargetTypeDetailsInput) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *TargetTypeDetailsInput) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *TargetTypeDetailsInput) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.

### HasGcpServiceAccountEmail

`func (o *TargetTypeDetailsInput) HasGcpServiceAccountEmail() bool`

HasGcpServiceAccountEmail returns a boolean if a field has been set.

### GetGcpServiceAccountKey

`func (o *TargetTypeDetailsInput) GetGcpServiceAccountKey() string`

GetGcpServiceAccountKey returns the GcpServiceAccountKey field if non-nil, zero value otherwise.

### GetGcpServiceAccountKeyOk

`func (o *TargetTypeDetailsInput) GetGcpServiceAccountKeyOk() (*string, bool)`

GetGcpServiceAccountKeyOk returns a tuple with the GcpServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountKey

`func (o *TargetTypeDetailsInput) SetGcpServiceAccountKey(v string)`

SetGcpServiceAccountKey sets GcpServiceAccountKey field to given value.

### HasGcpServiceAccountKey

`func (o *TargetTypeDetailsInput) HasGcpServiceAccountKey() bool`

HasGcpServiceAccountKey returns a boolean if a field has been set.

### GetGithubAppId

`func (o *TargetTypeDetailsInput) GetGithubAppId() int64`

GetGithubAppId returns the GithubAppId field if non-nil, zero value otherwise.

### GetGithubAppIdOk

`func (o *TargetTypeDetailsInput) GetGithubAppIdOk() (*int64, bool)`

GetGithubAppIdOk returns a tuple with the GithubAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppId

`func (o *TargetTypeDetailsInput) SetGithubAppId(v int64)`

SetGithubAppId sets GithubAppId field to given value.

### HasGithubAppId

`func (o *TargetTypeDetailsInput) HasGithubAppId() bool`

HasGithubAppId returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *TargetTypeDetailsInput) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *TargetTypeDetailsInput) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *TargetTypeDetailsInput) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *TargetTypeDetailsInput) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetGithubBaseUrl

`func (o *TargetTypeDetailsInput) GetGithubBaseUrl() string`

GetGithubBaseUrl returns the GithubBaseUrl field if non-nil, zero value otherwise.

### GetGithubBaseUrlOk

`func (o *TargetTypeDetailsInput) GetGithubBaseUrlOk() (*string, bool)`

GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubBaseUrl

`func (o *TargetTypeDetailsInput) SetGithubBaseUrl(v string)`

SetGithubBaseUrl sets GithubBaseUrl field to given value.

### HasGithubBaseUrl

`func (o *TargetTypeDetailsInput) HasGithubBaseUrl() bool`

HasGithubBaseUrl returns a boolean if a field has been set.

### GetGkeClusterCaCertificate

`func (o *TargetTypeDetailsInput) GetGkeClusterCaCertificate() string`

GetGkeClusterCaCertificate returns the GkeClusterCaCertificate field if non-nil, zero value otherwise.

### GetGkeClusterCaCertificateOk

`func (o *TargetTypeDetailsInput) GetGkeClusterCaCertificateOk() (*string, bool)`

GetGkeClusterCaCertificateOk returns a tuple with the GkeClusterCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterCaCertificate

`func (o *TargetTypeDetailsInput) SetGkeClusterCaCertificate(v string)`

SetGkeClusterCaCertificate sets GkeClusterCaCertificate field to given value.

### HasGkeClusterCaCertificate

`func (o *TargetTypeDetailsInput) HasGkeClusterCaCertificate() bool`

HasGkeClusterCaCertificate returns a boolean if a field has been set.

### GetGkeClusterEndpoint

`func (o *TargetTypeDetailsInput) GetGkeClusterEndpoint() string`

GetGkeClusterEndpoint returns the GkeClusterEndpoint field if non-nil, zero value otherwise.

### GetGkeClusterEndpointOk

`func (o *TargetTypeDetailsInput) GetGkeClusterEndpointOk() (*string, bool)`

GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterEndpoint

`func (o *TargetTypeDetailsInput) SetGkeClusterEndpoint(v string)`

SetGkeClusterEndpoint sets GkeClusterEndpoint field to given value.

### HasGkeClusterEndpoint

`func (o *TargetTypeDetailsInput) HasGkeClusterEndpoint() bool`

HasGkeClusterEndpoint returns a boolean if a field has been set.

### GetGkeClusterName

`func (o *TargetTypeDetailsInput) GetGkeClusterName() string`

GetGkeClusterName returns the GkeClusterName field if non-nil, zero value otherwise.

### GetGkeClusterNameOk

`func (o *TargetTypeDetailsInput) GetGkeClusterNameOk() (*string, bool)`

GetGkeClusterNameOk returns a tuple with the GkeClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterName

`func (o *TargetTypeDetailsInput) SetGkeClusterName(v string)`

SetGkeClusterName sets GkeClusterName field to given value.

### HasGkeClusterName

`func (o *TargetTypeDetailsInput) HasGkeClusterName() bool`

HasGkeClusterName returns a boolean if a field has been set.

### GetGkeServiceAccountKey

`func (o *TargetTypeDetailsInput) GetGkeServiceAccountKey() string`

GetGkeServiceAccountKey returns the GkeServiceAccountKey field if non-nil, zero value otherwise.

### GetGkeServiceAccountKeyOk

`func (o *TargetTypeDetailsInput) GetGkeServiceAccountKeyOk() (*string, bool)`

GetGkeServiceAccountKeyOk returns a tuple with the GkeServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeServiceAccountKey

`func (o *TargetTypeDetailsInput) SetGkeServiceAccountKey(v string)`

SetGkeServiceAccountKey sets GkeServiceAccountKey field to given value.

### HasGkeServiceAccountKey

`func (o *TargetTypeDetailsInput) HasGkeServiceAccountKey() bool`

HasGkeServiceAccountKey returns a boolean if a field has been set.

### GetGkeServiceAccountName

`func (o *TargetTypeDetailsInput) GetGkeServiceAccountName() string`

GetGkeServiceAccountName returns the GkeServiceAccountName field if non-nil, zero value otherwise.

### GetGkeServiceAccountNameOk

`func (o *TargetTypeDetailsInput) GetGkeServiceAccountNameOk() (*string, bool)`

GetGkeServiceAccountNameOk returns a tuple with the GkeServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeServiceAccountName

`func (o *TargetTypeDetailsInput) SetGkeServiceAccountName(v string)`

SetGkeServiceAccountName sets GkeServiceAccountName field to given value.

### HasGkeServiceAccountName

`func (o *TargetTypeDetailsInput) HasGkeServiceAccountName() bool`

HasGkeServiceAccountName returns a boolean if a field has been set.

### GetHost

`func (o *TargetTypeDetailsInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetTypeDetailsInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetTypeDetailsInput) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TargetTypeDetailsInput) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetK8sBearerToken

`func (o *TargetTypeDetailsInput) GetK8sBearerToken() string`

GetK8sBearerToken returns the K8sBearerToken field if non-nil, zero value otherwise.

### GetK8sBearerTokenOk

`func (o *TargetTypeDetailsInput) GetK8sBearerTokenOk() (*string, bool)`

GetK8sBearerTokenOk returns a tuple with the K8sBearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sBearerToken

`func (o *TargetTypeDetailsInput) SetK8sBearerToken(v string)`

SetK8sBearerToken sets K8sBearerToken field to given value.

### HasK8sBearerToken

`func (o *TargetTypeDetailsInput) HasK8sBearerToken() bool`

HasK8sBearerToken returns a boolean if a field has been set.

### GetK8sClusterCaCertificate

`func (o *TargetTypeDetailsInput) GetK8sClusterCaCertificate() string`

GetK8sClusterCaCertificate returns the K8sClusterCaCertificate field if non-nil, zero value otherwise.

### GetK8sClusterCaCertificateOk

`func (o *TargetTypeDetailsInput) GetK8sClusterCaCertificateOk() (*string, bool)`

GetK8sClusterCaCertificateOk returns a tuple with the K8sClusterCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterCaCertificate

`func (o *TargetTypeDetailsInput) SetK8sClusterCaCertificate(v string)`

SetK8sClusterCaCertificate sets K8sClusterCaCertificate field to given value.

### HasK8sClusterCaCertificate

`func (o *TargetTypeDetailsInput) HasK8sClusterCaCertificate() bool`

HasK8sClusterCaCertificate returns a boolean if a field has been set.

### GetK8sClusterEndpoint

`func (o *TargetTypeDetailsInput) GetK8sClusterEndpoint() string`

GetK8sClusterEndpoint returns the K8sClusterEndpoint field if non-nil, zero value otherwise.

### GetK8sClusterEndpointOk

`func (o *TargetTypeDetailsInput) GetK8sClusterEndpointOk() (*string, bool)`

GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterEndpoint

`func (o *TargetTypeDetailsInput) SetK8sClusterEndpoint(v string)`

SetK8sClusterEndpoint sets K8sClusterEndpoint field to given value.

### HasK8sClusterEndpoint

`func (o *TargetTypeDetailsInput) HasK8sClusterEndpoint() bool`

HasK8sClusterEndpoint returns a boolean if a field has been set.

### GetLdapAudience

`func (o *TargetTypeDetailsInput) GetLdapAudience() string`

GetLdapAudience returns the LdapAudience field if non-nil, zero value otherwise.

### GetLdapAudienceOk

`func (o *TargetTypeDetailsInput) GetLdapAudienceOk() (*string, bool)`

GetLdapAudienceOk returns a tuple with the LdapAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAudience

`func (o *TargetTypeDetailsInput) SetLdapAudience(v string)`

SetLdapAudience sets LdapAudience field to given value.

### HasLdapAudience

`func (o *TargetTypeDetailsInput) HasLdapAudience() bool`

HasLdapAudience returns a boolean if a field has been set.

### GetLdapBindDn

`func (o *TargetTypeDetailsInput) GetLdapBindDn() string`

GetLdapBindDn returns the LdapBindDn field if non-nil, zero value otherwise.

### GetLdapBindDnOk

`func (o *TargetTypeDetailsInput) GetLdapBindDnOk() (*string, bool)`

GetLdapBindDnOk returns a tuple with the LdapBindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindDn

`func (o *TargetTypeDetailsInput) SetLdapBindDn(v string)`

SetLdapBindDn sets LdapBindDn field to given value.

### HasLdapBindDn

`func (o *TargetTypeDetailsInput) HasLdapBindDn() bool`

HasLdapBindDn returns a boolean if a field has been set.

### GetLdapBindPassword

`func (o *TargetTypeDetailsInput) GetLdapBindPassword() string`

GetLdapBindPassword returns the LdapBindPassword field if non-nil, zero value otherwise.

### GetLdapBindPasswordOk

`func (o *TargetTypeDetailsInput) GetLdapBindPasswordOk() (*string, bool)`

GetLdapBindPasswordOk returns a tuple with the LdapBindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindPassword

`func (o *TargetTypeDetailsInput) SetLdapBindPassword(v string)`

SetLdapBindPassword sets LdapBindPassword field to given value.

### HasLdapBindPassword

`func (o *TargetTypeDetailsInput) HasLdapBindPassword() bool`

HasLdapBindPassword returns a boolean if a field has been set.

### GetLdapCertificate

`func (o *TargetTypeDetailsInput) GetLdapCertificate() string`

GetLdapCertificate returns the LdapCertificate field if non-nil, zero value otherwise.

### GetLdapCertificateOk

`func (o *TargetTypeDetailsInput) GetLdapCertificateOk() (*string, bool)`

GetLdapCertificateOk returns a tuple with the LdapCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCertificate

`func (o *TargetTypeDetailsInput) SetLdapCertificate(v string)`

SetLdapCertificate sets LdapCertificate field to given value.

### HasLdapCertificate

`func (o *TargetTypeDetailsInput) HasLdapCertificate() bool`

HasLdapCertificate returns a boolean if a field has been set.

### GetLdapTokenExpiration

`func (o *TargetTypeDetailsInput) GetLdapTokenExpiration() string`

GetLdapTokenExpiration returns the LdapTokenExpiration field if non-nil, zero value otherwise.

### GetLdapTokenExpirationOk

`func (o *TargetTypeDetailsInput) GetLdapTokenExpirationOk() (*string, bool)`

GetLdapTokenExpirationOk returns a tuple with the LdapTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapTokenExpiration

`func (o *TargetTypeDetailsInput) SetLdapTokenExpiration(v string)`

SetLdapTokenExpiration sets LdapTokenExpiration field to given value.

### HasLdapTokenExpiration

`func (o *TargetTypeDetailsInput) HasLdapTokenExpiration() bool`

HasLdapTokenExpiration returns a boolean if a field has been set.

### GetLdapUrl

`func (o *TargetTypeDetailsInput) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *TargetTypeDetailsInput) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *TargetTypeDetailsInput) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *TargetTypeDetailsInput) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetMongodbAtlasApiPrivateKey

`func (o *TargetTypeDetailsInput) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *TargetTypeDetailsInput) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *TargetTypeDetailsInput) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *TargetTypeDetailsInput) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *TargetTypeDetailsInput) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *TargetTypeDetailsInput) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *TargetTypeDetailsInput) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *TargetTypeDetailsInput) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *TargetTypeDetailsInput) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *TargetTypeDetailsInput) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *TargetTypeDetailsInput) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *TargetTypeDetailsInput) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbDbName

`func (o *TargetTypeDetailsInput) GetMongodbDbName() string`

GetMongodbDbName returns the MongodbDbName field if non-nil, zero value otherwise.

### GetMongodbDbNameOk

`func (o *TargetTypeDetailsInput) GetMongodbDbNameOk() (*string, bool)`

GetMongodbDbNameOk returns a tuple with the MongodbDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDbName

`func (o *TargetTypeDetailsInput) SetMongodbDbName(v string)`

SetMongodbDbName sets MongodbDbName field to given value.

### HasMongodbDbName

`func (o *TargetTypeDetailsInput) HasMongodbDbName() bool`

HasMongodbDbName returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *TargetTypeDetailsInput) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *TargetTypeDetailsInput) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *TargetTypeDetailsInput) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *TargetTypeDetailsInput) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbHostPort

`func (o *TargetTypeDetailsInput) GetMongodbHostPort() string`

GetMongodbHostPort returns the MongodbHostPort field if non-nil, zero value otherwise.

### GetMongodbHostPortOk

`func (o *TargetTypeDetailsInput) GetMongodbHostPortOk() (*string, bool)`

GetMongodbHostPortOk returns a tuple with the MongodbHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbHostPort

`func (o *TargetTypeDetailsInput) SetMongodbHostPort(v string)`

SetMongodbHostPort sets MongodbHostPort field to given value.

### HasMongodbHostPort

`func (o *TargetTypeDetailsInput) HasMongodbHostPort() bool`

HasMongodbHostPort returns a boolean if a field has been set.

### GetMongodbIsAtlas

`func (o *TargetTypeDetailsInput) GetMongodbIsAtlas() bool`

GetMongodbIsAtlas returns the MongodbIsAtlas field if non-nil, zero value otherwise.

### GetMongodbIsAtlasOk

`func (o *TargetTypeDetailsInput) GetMongodbIsAtlasOk() (*bool, bool)`

GetMongodbIsAtlasOk returns a tuple with the MongodbIsAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbIsAtlas

`func (o *TargetTypeDetailsInput) SetMongodbIsAtlas(v bool)`

SetMongodbIsAtlas sets MongodbIsAtlas field to given value.

### HasMongodbIsAtlas

`func (o *TargetTypeDetailsInput) HasMongodbIsAtlas() bool`

HasMongodbIsAtlas returns a boolean if a field has been set.

### GetMongodbPassword

`func (o *TargetTypeDetailsInput) GetMongodbPassword() string`

GetMongodbPassword returns the MongodbPassword field if non-nil, zero value otherwise.

### GetMongodbPasswordOk

`func (o *TargetTypeDetailsInput) GetMongodbPasswordOk() (*string, bool)`

GetMongodbPasswordOk returns a tuple with the MongodbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbPassword

`func (o *TargetTypeDetailsInput) SetMongodbPassword(v string)`

SetMongodbPassword sets MongodbPassword field to given value.

### HasMongodbPassword

`func (o *TargetTypeDetailsInput) HasMongodbPassword() bool`

HasMongodbPassword returns a boolean if a field has been set.

### GetMongodbUriConnection

`func (o *TargetTypeDetailsInput) GetMongodbUriConnection() string`

GetMongodbUriConnection returns the MongodbUriConnection field if non-nil, zero value otherwise.

### GetMongodbUriConnectionOk

`func (o *TargetTypeDetailsInput) GetMongodbUriConnectionOk() (*string, bool)`

GetMongodbUriConnectionOk returns a tuple with the MongodbUriConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriConnection

`func (o *TargetTypeDetailsInput) SetMongodbUriConnection(v string)`

SetMongodbUriConnection sets MongodbUriConnection field to given value.

### HasMongodbUriConnection

`func (o *TargetTypeDetailsInput) HasMongodbUriConnection() bool`

HasMongodbUriConnection returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *TargetTypeDetailsInput) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *TargetTypeDetailsInput) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *TargetTypeDetailsInput) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *TargetTypeDetailsInput) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetMongodbUsername

`func (o *TargetTypeDetailsInput) GetMongodbUsername() string`

GetMongodbUsername returns the MongodbUsername field if non-nil, zero value otherwise.

### GetMongodbUsernameOk

`func (o *TargetTypeDetailsInput) GetMongodbUsernameOk() (*string, bool)`

GetMongodbUsernameOk returns a tuple with the MongodbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUsername

`func (o *TargetTypeDetailsInput) SetMongodbUsername(v string)`

SetMongodbUsername sets MongodbUsername field to given value.

### HasMongodbUsername

`func (o *TargetTypeDetailsInput) HasMongodbUsername() bool`

HasMongodbUsername returns a boolean if a field has been set.

### GetPassword

`func (o *TargetTypeDetailsInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetTypeDetailsInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetTypeDetailsInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TargetTypeDetailsInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPayload

`func (o *TargetTypeDetailsInput) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *TargetTypeDetailsInput) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *TargetTypeDetailsInput) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *TargetTypeDetailsInput) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPort

`func (o *TargetTypeDetailsInput) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetTypeDetailsInput) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetTypeDetailsInput) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetTypeDetailsInput) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateKey

`func (o *TargetTypeDetailsInput) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *TargetTypeDetailsInput) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *TargetTypeDetailsInput) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *TargetTypeDetailsInput) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyPassword

`func (o *TargetTypeDetailsInput) GetPrivateKeyPassword() string`

GetPrivateKeyPassword returns the PrivateKeyPassword field if non-nil, zero value otherwise.

### GetPrivateKeyPasswordOk

`func (o *TargetTypeDetailsInput) GetPrivateKeyPasswordOk() (*string, bool)`

GetPrivateKeyPasswordOk returns a tuple with the PrivateKeyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassword

`func (o *TargetTypeDetailsInput) SetPrivateKeyPassword(v string)`

SetPrivateKeyPassword sets PrivateKeyPassword field to given value.

### HasPrivateKeyPassword

`func (o *TargetTypeDetailsInput) HasPrivateKeyPassword() bool`

HasPrivateKeyPassword returns a boolean if a field has been set.

### GetRabbitmqServerPassword

`func (o *TargetTypeDetailsInput) GetRabbitmqServerPassword() string`

GetRabbitmqServerPassword returns the RabbitmqServerPassword field if non-nil, zero value otherwise.

### GetRabbitmqServerPasswordOk

`func (o *TargetTypeDetailsInput) GetRabbitmqServerPasswordOk() (*string, bool)`

GetRabbitmqServerPasswordOk returns a tuple with the RabbitmqServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerPassword

`func (o *TargetTypeDetailsInput) SetRabbitmqServerPassword(v string)`

SetRabbitmqServerPassword sets RabbitmqServerPassword field to given value.

### HasRabbitmqServerPassword

`func (o *TargetTypeDetailsInput) HasRabbitmqServerPassword() bool`

HasRabbitmqServerPassword returns a boolean if a field has been set.

### GetRabbitmqServerUri

`func (o *TargetTypeDetailsInput) GetRabbitmqServerUri() string`

GetRabbitmqServerUri returns the RabbitmqServerUri field if non-nil, zero value otherwise.

### GetRabbitmqServerUriOk

`func (o *TargetTypeDetailsInput) GetRabbitmqServerUriOk() (*string, bool)`

GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUri

`func (o *TargetTypeDetailsInput) SetRabbitmqServerUri(v string)`

SetRabbitmqServerUri sets RabbitmqServerUri field to given value.

### HasRabbitmqServerUri

`func (o *TargetTypeDetailsInput) HasRabbitmqServerUri() bool`

HasRabbitmqServerUri returns a boolean if a field has been set.

### GetRabbitmqServerUser

`func (o *TargetTypeDetailsInput) GetRabbitmqServerUser() string`

GetRabbitmqServerUser returns the RabbitmqServerUser field if non-nil, zero value otherwise.

### GetRabbitmqServerUserOk

`func (o *TargetTypeDetailsInput) GetRabbitmqServerUserOk() (*string, bool)`

GetRabbitmqServerUserOk returns a tuple with the RabbitmqServerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUser

`func (o *TargetTypeDetailsInput) SetRabbitmqServerUser(v string)`

SetRabbitmqServerUser sets RabbitmqServerUser field to given value.

### HasRabbitmqServerUser

`func (o *TargetTypeDetailsInput) HasRabbitmqServerUser() bool`

HasRabbitmqServerUser returns a boolean if a field has been set.

### GetSecurityToken

`func (o *TargetTypeDetailsInput) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *TargetTypeDetailsInput) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *TargetTypeDetailsInput) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *TargetTypeDetailsInput) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetSfAccount

`func (o *TargetTypeDetailsInput) GetSfAccount() string`

GetSfAccount returns the SfAccount field if non-nil, zero value otherwise.

### GetSfAccountOk

`func (o *TargetTypeDetailsInput) GetSfAccountOk() (*string, bool)`

GetSfAccountOk returns a tuple with the SfAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfAccount

`func (o *TargetTypeDetailsInput) SetSfAccount(v string)`

SetSfAccount sets SfAccount field to given value.

### HasSfAccount

`func (o *TargetTypeDetailsInput) HasSfAccount() bool`

HasSfAccount returns a boolean if a field has been set.

### GetSslConnectionCertificate

`func (o *TargetTypeDetailsInput) GetSslConnectionCertificate() string`

GetSslConnectionCertificate returns the SslConnectionCertificate field if non-nil, zero value otherwise.

### GetSslConnectionCertificateOk

`func (o *TargetTypeDetailsInput) GetSslConnectionCertificateOk() (*string, bool)`

GetSslConnectionCertificateOk returns a tuple with the SslConnectionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnectionCertificate

`func (o *TargetTypeDetailsInput) SetSslConnectionCertificate(v string)`

SetSslConnectionCertificate sets SslConnectionCertificate field to given value.

### HasSslConnectionCertificate

`func (o *TargetTypeDetailsInput) HasSslConnectionCertificate() bool`

HasSslConnectionCertificate returns a boolean if a field has been set.

### GetSslConnectionMode

`func (o *TargetTypeDetailsInput) GetSslConnectionMode() bool`

GetSslConnectionMode returns the SslConnectionMode field if non-nil, zero value otherwise.

### GetSslConnectionModeOk

`func (o *TargetTypeDetailsInput) GetSslConnectionModeOk() (*bool, bool)`

GetSslConnectionModeOk returns a tuple with the SslConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnectionMode

`func (o *TargetTypeDetailsInput) SetSslConnectionMode(v bool)`

SetSslConnectionMode sets SslConnectionMode field to given value.

### HasSslConnectionMode

`func (o *TargetTypeDetailsInput) HasSslConnectionMode() bool`

HasSslConnectionMode returns a boolean if a field has been set.

### GetTenantUrl

`func (o *TargetTypeDetailsInput) GetTenantUrl() string`

GetTenantUrl returns the TenantUrl field if non-nil, zero value otherwise.

### GetTenantUrlOk

`func (o *TargetTypeDetailsInput) GetTenantUrlOk() (*string, bool)`

GetTenantUrlOk returns a tuple with the TenantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUrl

`func (o *TargetTypeDetailsInput) SetTenantUrl(v string)`

SetTenantUrl sets TenantUrl field to given value.

### HasTenantUrl

`func (o *TargetTypeDetailsInput) HasTenantUrl() bool`

HasTenantUrl returns a boolean if a field has been set.

### GetUrl

`func (o *TargetTypeDetailsInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TargetTypeDetailsInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TargetTypeDetailsInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TargetTypeDetailsInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *TargetTypeDetailsInput) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *TargetTypeDetailsInput) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *TargetTypeDetailsInput) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *TargetTypeDetailsInput) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.

### GetUserName

`func (o *TargetTypeDetailsInput) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *TargetTypeDetailsInput) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *TargetTypeDetailsInput) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *TargetTypeDetailsInput) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUsername

`func (o *TargetTypeDetailsInput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TargetTypeDetailsInput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TargetTypeDetailsInput) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TargetTypeDetailsInput) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVenafiApiKey

`func (o *TargetTypeDetailsInput) GetVenafiApiKey() string`

GetVenafiApiKey returns the VenafiApiKey field if non-nil, zero value otherwise.

### GetVenafiApiKeyOk

`func (o *TargetTypeDetailsInput) GetVenafiApiKeyOk() (*string, bool)`

GetVenafiApiKeyOk returns a tuple with the VenafiApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiApiKey

`func (o *TargetTypeDetailsInput) SetVenafiApiKey(v string)`

SetVenafiApiKey sets VenafiApiKey field to given value.

### HasVenafiApiKey

`func (o *TargetTypeDetailsInput) HasVenafiApiKey() bool`

HasVenafiApiKey returns a boolean if a field has been set.

### GetVenafiBaseUrl

`func (o *TargetTypeDetailsInput) GetVenafiBaseUrl() string`

GetVenafiBaseUrl returns the VenafiBaseUrl field if non-nil, zero value otherwise.

### GetVenafiBaseUrlOk

`func (o *TargetTypeDetailsInput) GetVenafiBaseUrlOk() (*string, bool)`

GetVenafiBaseUrlOk returns a tuple with the VenafiBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiBaseUrl

`func (o *TargetTypeDetailsInput) SetVenafiBaseUrl(v string)`

SetVenafiBaseUrl sets VenafiBaseUrl field to given value.

### HasVenafiBaseUrl

`func (o *TargetTypeDetailsInput) HasVenafiBaseUrl() bool`

HasVenafiBaseUrl returns a boolean if a field has been set.

### GetVenafiTppPassword

`func (o *TargetTypeDetailsInput) GetVenafiTppPassword() string`

GetVenafiTppPassword returns the VenafiTppPassword field if non-nil, zero value otherwise.

### GetVenafiTppPasswordOk

`func (o *TargetTypeDetailsInput) GetVenafiTppPasswordOk() (*string, bool)`

GetVenafiTppPasswordOk returns a tuple with the VenafiTppPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiTppPassword

`func (o *TargetTypeDetailsInput) SetVenafiTppPassword(v string)`

SetVenafiTppPassword sets VenafiTppPassword field to given value.

### HasVenafiTppPassword

`func (o *TargetTypeDetailsInput) HasVenafiTppPassword() bool`

HasVenafiTppPassword returns a boolean if a field has been set.

### GetVenafiTppUsername

`func (o *TargetTypeDetailsInput) GetVenafiTppUsername() string`

GetVenafiTppUsername returns the VenafiTppUsername field if non-nil, zero value otherwise.

### GetVenafiTppUsernameOk

`func (o *TargetTypeDetailsInput) GetVenafiTppUsernameOk() (*string, bool)`

GetVenafiTppUsernameOk returns a tuple with the VenafiTppUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiTppUsername

`func (o *TargetTypeDetailsInput) SetVenafiTppUsername(v string)`

SetVenafiTppUsername sets VenafiTppUsername field to given value.

### HasVenafiTppUsername

`func (o *TargetTypeDetailsInput) HasVenafiTppUsername() bool`

HasVenafiTppUsername returns a boolean if a field has been set.

### GetVenafiUseTpp

`func (o *TargetTypeDetailsInput) GetVenafiUseTpp() bool`

GetVenafiUseTpp returns the VenafiUseTpp field if non-nil, zero value otherwise.

### GetVenafiUseTppOk

`func (o *TargetTypeDetailsInput) GetVenafiUseTppOk() (*bool, bool)`

GetVenafiUseTppOk returns a tuple with the VenafiUseTpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiUseTpp

`func (o *TargetTypeDetailsInput) SetVenafiUseTpp(v bool)`

SetVenafiUseTpp sets VenafiUseTpp field to given value.

### HasVenafiUseTpp

`func (o *TargetTypeDetailsInput) HasVenafiUseTpp() bool`

HasVenafiUseTpp returns a boolean if a field has been set.

### GetVenafiZone

`func (o *TargetTypeDetailsInput) GetVenafiZone() string`

GetVenafiZone returns the VenafiZone field if non-nil, zero value otherwise.

### GetVenafiZoneOk

`func (o *TargetTypeDetailsInput) GetVenafiZoneOk() (*string, bool)`

GetVenafiZoneOk returns a tuple with the VenafiZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiZone

`func (o *TargetTypeDetailsInput) SetVenafiZone(v string)`

SetVenafiZone sets VenafiZone field to given value.

### HasVenafiZone

`func (o *TargetTypeDetailsInput) HasVenafiZone() bool`

HasVenafiZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


