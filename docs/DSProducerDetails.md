# DSProducerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**AdminName** | Pointer to **string** |  | [optional] 
**AdminPwd** | Pointer to **string** |  | [optional] 
**AdminRotationIntervalDays** | Pointer to **int64** |  | [optional] 
**ArtifactoryAdminApikey** | Pointer to **string** |  | [optional] 
**ArtifactoryAdminUsername** | Pointer to **string** |  | [optional] 
**ArtifactoryBaseUrl** | Pointer to **string** |  | [optional] 
**ArtifactoryTokenAudience** | Pointer to **string** |  | [optional] 
**ArtifactoryTokenScope** | Pointer to **string** |  | [optional] 
**AwsAccessKeyId** | Pointer to **string** |  | [optional] 
**AwsAccessMode** | Pointer to **string** |  | [optional] 
**AwsRegion** | Pointer to **string** |  | [optional] 
**AwsRoleArns** | Pointer to **string** |  | [optional] 
**AwsSecretAccessKey** | Pointer to **string** |  | [optional] 
**AwsSessionToken** | Pointer to **string** |  | [optional] 
**AwsUserConsoleAccess** | Pointer to **bool** |  | [optional] 
**AwsUserGroups** | Pointer to **string** |  | [optional] 
**AwsUserPolicies** | Pointer to **string** |  | [optional] 
**AwsUserProgrammaticAccess** | Pointer to **bool** |  | [optional] 
**AzureAppObjectId** | Pointer to **string** |  | [optional] 
**AzureClientId** | Pointer to **string** |  | [optional] 
**AzureClientSecret** | Pointer to **string** |  | [optional] 
**AzureFixedUserNameSubClaimKey** | Pointer to **string** |  | [optional] 
**AzureFixedUserOnly** | Pointer to **bool** |  | [optional] 
**AzureResourceGroupName** | Pointer to **string** |  | [optional] 
**AzureResourceName** | Pointer to **string** |  | [optional] 
**AzureSubscriptionId** | Pointer to **string** |  | [optional] 
**AzureTenantId** | Pointer to **string** |  | [optional] 
**AzureUserGroupsObjId** | Pointer to **string** |  | [optional] 
**AzureUserPortalAccess** | Pointer to **bool** |  | [optional] 
**AzureUserProgrammaticAccess** | Pointer to **bool** |  | [optional] 
**AzureUserRolesTemplateId** | Pointer to **string** |  | [optional] 
**CassandraCreationStatements** | Pointer to **string** |  | [optional] 
**ChefOrganizations** | Pointer to **string** |  | [optional] 
**ChefServerAccessMode** | Pointer to **string** |  | [optional] 
**ChefServerHostName** | Pointer to **string** |  | [optional] 
**ChefServerKey** | Pointer to **string** |  | [optional] 
**ChefServerPort** | Pointer to **string** |  | [optional] 
**ChefServerUrl** | Pointer to **string** |  | [optional] 
**ChefServerUsername** | Pointer to **string** |  | [optional] 
**ChefSkipSsl** | Pointer to **bool** |  | [optional] 
**CreateSyncUrl** | Pointer to **string** |  | [optional] 
**DbHostName** | Pointer to **string** |  | [optional] 
**DbIsolationLevel** | Pointer to **string** |  | [optional] 
**DbMaxIdleConns** | Pointer to **string** |  | [optional] 
**DbMaxOpenConns** | Pointer to **string** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**DbPort** | Pointer to **string** |  | [optional] 
**DbPwd** | Pointer to **string** |  | [optional] 
**DbServerCertificates** | Pointer to **string** | (Optional) DBServerCertificates defines the set of root certificate authorities that clients use when verifying server certificates. If DBServerCertificates is empty, TLS uses the host&#39;s root CA set. | [optional] 
**DbServerName** | Pointer to **string** | (Optional) ServerName is used to verify the hostname on the returned certificates unless InsecureSkipVerify is given. It is also included in the client&#39;s handshake to support virtual hosting unless it is an IP address. | [optional] 
**DbUserName** | Pointer to **string** |  | [optional] 
**DynamicSecretId** | Pointer to **int64** |  | [optional] 
**DynamicSecretKey** | Pointer to **string** |  | [optional] 
**DynamicSecretName** | Pointer to **string** |  | [optional] 
**DynamicSecretType** | Pointer to **string** |  | [optional] 
**EksAccessKeyId** | Pointer to **string** |  | [optional] 
**EksAssumeRole** | Pointer to **string** |  | [optional] 
**EksClusterCaCertificate** | Pointer to **string** |  | [optional] 
**EksClusterEndpoint** | Pointer to **string** |  | [optional] 
**EksClusterName** | Pointer to **string** |  | [optional] 
**EksRegion** | Pointer to **string** |  | [optional] 
**EksSecretAccessKey** | Pointer to **string** |  | [optional] 
**EnableAdminRotation** | Pointer to **bool** |  | [optional] 
**ExternallyProvidedUser** | Pointer to **string** |  | [optional] 
**FailureMessage** | Pointer to **string** |  | [optional] 
**FixedUserOnly** | Pointer to **string** |  | [optional] 
**GcpKeyAlgo** | Pointer to **string** |  | [optional] 
**GcpServiceAccountEmail** | Pointer to **string** |  | [optional] 
**GcpServiceAccountKey** | Pointer to **string** |  | [optional] 
**GcpTokenLifetime** | Pointer to **string** |  | [optional] 
**GcpTokenScope** | Pointer to **string** |  | [optional] 
**GcpTokenType** | Pointer to **string** |  | [optional] 
**GithubAppId** | Pointer to **int64** |  | [optional] 
**GithubAppPrivateKey** | Pointer to **string** |  | [optional] 
**GithubBaseUrl** | Pointer to **string** |  | [optional] 
**GithubInstallationId** | Pointer to **int64** |  | [optional] 
**GithubInstallationTokenPermissions** | Pointer to **map[string]string** |  | [optional] 
**GithubInstallationTokenRepositories** | Pointer to **[]string** |  | [optional] 
**GithubInstallationTokenRepositoriesIds** | Pointer to **[]int64** |  | [optional] 
**GithubRepositoryPath** | Pointer to **string** |  | [optional] 
**GkeClusterCaCertificate** | Pointer to **string** |  | [optional] 
**GkeClusterEndpoint** | Pointer to **string** |  | [optional] 
**GkeClusterName** | Pointer to **string** |  | [optional] 
**GkeServiceAccountKey** | Pointer to **string** |  | [optional] 
**GkeServiceAccountName** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **string** |  | [optional] 
**HanadbCreationStatements** | Pointer to **string** |  | [optional] 
**HanadbRevocationStatements** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**HostPort** | Pointer to **string** |  | [optional] 
**IsFixedUser** | Pointer to **string** |  | [optional] 
**ItemTargetsAssoc** | Pointer to [**[]ItemTargetAssociation**](ItemTargetAssociation.md) |  | [optional] 
**K8sBearerToken** | Pointer to **string** |  | [optional] 
**K8sClusterCaCertificate** | Pointer to **string** |  | [optional] 
**K8sClusterEndpoint** | Pointer to **string** |  | [optional] 
**K8sNamespace** | Pointer to **string** |  | [optional] 
**K8sServiceAccount** | Pointer to **string** |  | [optional] 
**LastAdminRotation** | Pointer to **int64** |  | [optional] 
**LdapAudience** | Pointer to **string** |  | [optional] 
**LdapBindDn** | Pointer to **string** |  | [optional] 
**LdapBindPassword** | Pointer to **string** |  | [optional] 
**LdapCertificate** | Pointer to **string** |  | [optional] 
**LdapTokenExpiration** | Pointer to **string** |  | [optional] 
**LdapUrl** | Pointer to **string** |  | [optional] 
**LdapUserAttr** | Pointer to **string** |  | [optional] 
**LdapUserDn** | Pointer to **string** |  | [optional] 
**MongodbAtlasApiPrivateKey** | Pointer to **string** |  | [optional] 
**MongodbAtlasApiPublicKey** | Pointer to **string** |  | [optional] 
**MongodbAtlasProjectId** | Pointer to **string** | mongodb atlas fields | [optional] 
**MongodbCustomData** | Pointer to **string** |  | [optional] 
**MongodbDbName** | Pointer to **string** | common fields | [optional] 
**MongodbDefaultAuthDb** | Pointer to **string** |  | [optional] 
**MongodbHostPort** | Pointer to **string** |  | [optional] 
**MongodbIsAtlas** | Pointer to **bool** |  | [optional] 
**MongodbPassword** | Pointer to **string** |  | [optional] 
**MongodbRoles** | Pointer to **string** | common fields | [optional] 
**MongodbUriConnection** | Pointer to **string** | mongodb fields | [optional] 
**MongodbUriOptions** | Pointer to **string** |  | [optional] 
**MongodbUsername** | Pointer to **string** |  | [optional] 
**MssqlCreationStatements** | Pointer to **string** |  | [optional] 
**MssqlRevocationStatements** | Pointer to **string** |  | [optional] 
**MysqlCreationStatements** | Pointer to **string** |  | [optional] 
**OracleCreationStatements** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordLength** | Pointer to **int64** |  | [optional] 
**PasswordPolicy** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**PostgresCreationStatements** | Pointer to **string** |  | [optional] 
**PostgresRevocationStatements** | Pointer to **string** |  | [optional] 
**RabbitmqServerPassword** | Pointer to **string** |  | [optional] 
**RabbitmqServerUri** | Pointer to **string** |  | [optional] 
**RabbitmqServerUser** | Pointer to **string** |  | [optional] 
**RabbitmqUserConfPermission** | Pointer to **string** |  | [optional] 
**RabbitmqUserReadPermission** | Pointer to **string** |  | [optional] 
**RabbitmqUserTags** | Pointer to **string** |  | [optional] 
**RabbitmqUserVhost** | Pointer to **string** |  | [optional] 
**RabbitmqUserWritePermission** | Pointer to **string** |  | [optional] 
**RedshiftCreationStatements** | Pointer to **string** |  | [optional] 
**RevokeSyncUrl** | Pointer to **string** |  | [optional] 
**RotateSyncUrl** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**SecureRemoteAccessDetails** | Pointer to [**SecureRemoteAccess**](SecureRemoteAccess.md) |  | [optional] 
**SessionExtensionWarnIntervalMin** | Pointer to **int64** |  | [optional] 
**SfAccount** | Pointer to **string** |  | [optional] 
**SfUserRole** | Pointer to **string** | generated  users info | [optional] 
**SfWarehouseName** | Pointer to **string** |  | [optional] 
**ShouldStop** | Pointer to **string** | TODO delete this after migration | [optional] 
**SslConnectionCertificate** | Pointer to **string** | (Optional) SSLConnectionCertificate defines the certificate for SSL connection. Must be base64 certificate loaded by UI using file loader field | [optional] 
**SslConnectionMode** | Pointer to **bool** | (Optional) SSLConnectionMode defines if SSL mode will be used to connect to DB | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TimeoutSeconds** | Pointer to **int64** |  | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**UserPrincipalName** | Pointer to **string** |  | [optional] 
**UserTtl** | Pointer to **string** |  | [optional] 
**UsernameLength** | Pointer to **int64** |  | [optional] 
**UsernamePolicy** | Pointer to **string** |  | [optional] 
**VenafiAllowSubdomains** | Pointer to **bool** |  | [optional] 
**VenafiAllowedDomains** | Pointer to **[]string** |  | [optional] 
**VenafiApiKey** | Pointer to **string** |  | [optional] 
**VenafiAutoGeneratedFolder** | Pointer to **string** |  | [optional] 
**VenafiBaseUrl** | Pointer to **string** |  | [optional] 
**VenafiRootFirstInChain** | Pointer to **bool** |  | [optional] 
**VenafiSignUsingAkeylessPki** | Pointer to **bool** |  | [optional] 
**VenafiSignerKeyName** | Pointer to **string** |  | [optional] 
**VenafiStorePrivateKey** | Pointer to **bool** |  | [optional] 
**VenafiTppPassword** | Pointer to **string** |  | [optional] 
**VenafiTppUsername** | Pointer to **string** |  | [optional] 
**VenafiUseTpp** | Pointer to **bool** |  | [optional] 
**VenafiZone** | Pointer to **string** |  | [optional] 
**WarnBeforeUserExpirationMin** | Pointer to **int64** |  | [optional] 

## Methods

### NewDSProducerDetails

`func NewDSProducerDetails() *DSProducerDetails`

NewDSProducerDetails instantiates a new DSProducerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSProducerDetailsWithDefaults

`func NewDSProducerDetailsWithDefaults() *DSProducerDetails`

NewDSProducerDetailsWithDefaults instantiates a new DSProducerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DSProducerDetails) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DSProducerDetails) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DSProducerDetails) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DSProducerDetails) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAdminName

`func (o *DSProducerDetails) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *DSProducerDetails) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *DSProducerDetails) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *DSProducerDetails) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAdminPwd

`func (o *DSProducerDetails) GetAdminPwd() string`

GetAdminPwd returns the AdminPwd field if non-nil, zero value otherwise.

### GetAdminPwdOk

`func (o *DSProducerDetails) GetAdminPwdOk() (*string, bool)`

GetAdminPwdOk returns a tuple with the AdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPwd

`func (o *DSProducerDetails) SetAdminPwd(v string)`

SetAdminPwd sets AdminPwd field to given value.

### HasAdminPwd

`func (o *DSProducerDetails) HasAdminPwd() bool`

HasAdminPwd returns a boolean if a field has been set.

### GetAdminRotationIntervalDays

`func (o *DSProducerDetails) GetAdminRotationIntervalDays() int64`

GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field if non-nil, zero value otherwise.

### GetAdminRotationIntervalDaysOk

`func (o *DSProducerDetails) GetAdminRotationIntervalDaysOk() (*int64, bool)`

GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminRotationIntervalDays

`func (o *DSProducerDetails) SetAdminRotationIntervalDays(v int64)`

SetAdminRotationIntervalDays sets AdminRotationIntervalDays field to given value.

### HasAdminRotationIntervalDays

`func (o *DSProducerDetails) HasAdminRotationIntervalDays() bool`

HasAdminRotationIntervalDays returns a boolean if a field has been set.

### GetArtifactoryAdminApikey

`func (o *DSProducerDetails) GetArtifactoryAdminApikey() string`

GetArtifactoryAdminApikey returns the ArtifactoryAdminApikey field if non-nil, zero value otherwise.

### GetArtifactoryAdminApikeyOk

`func (o *DSProducerDetails) GetArtifactoryAdminApikeyOk() (*string, bool)`

GetArtifactoryAdminApikeyOk returns a tuple with the ArtifactoryAdminApikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminApikey

`func (o *DSProducerDetails) SetArtifactoryAdminApikey(v string)`

SetArtifactoryAdminApikey sets ArtifactoryAdminApikey field to given value.

### HasArtifactoryAdminApikey

`func (o *DSProducerDetails) HasArtifactoryAdminApikey() bool`

HasArtifactoryAdminApikey returns a boolean if a field has been set.

### GetArtifactoryAdminUsername

`func (o *DSProducerDetails) GetArtifactoryAdminUsername() string`

GetArtifactoryAdminUsername returns the ArtifactoryAdminUsername field if non-nil, zero value otherwise.

### GetArtifactoryAdminUsernameOk

`func (o *DSProducerDetails) GetArtifactoryAdminUsernameOk() (*string, bool)`

GetArtifactoryAdminUsernameOk returns a tuple with the ArtifactoryAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminUsername

`func (o *DSProducerDetails) SetArtifactoryAdminUsername(v string)`

SetArtifactoryAdminUsername sets ArtifactoryAdminUsername field to given value.

### HasArtifactoryAdminUsername

`func (o *DSProducerDetails) HasArtifactoryAdminUsername() bool`

HasArtifactoryAdminUsername returns a boolean if a field has been set.

### GetArtifactoryBaseUrl

`func (o *DSProducerDetails) GetArtifactoryBaseUrl() string`

GetArtifactoryBaseUrl returns the ArtifactoryBaseUrl field if non-nil, zero value otherwise.

### GetArtifactoryBaseUrlOk

`func (o *DSProducerDetails) GetArtifactoryBaseUrlOk() (*string, bool)`

GetArtifactoryBaseUrlOk returns a tuple with the ArtifactoryBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryBaseUrl

`func (o *DSProducerDetails) SetArtifactoryBaseUrl(v string)`

SetArtifactoryBaseUrl sets ArtifactoryBaseUrl field to given value.

### HasArtifactoryBaseUrl

`func (o *DSProducerDetails) HasArtifactoryBaseUrl() bool`

HasArtifactoryBaseUrl returns a boolean if a field has been set.

### GetArtifactoryTokenAudience

`func (o *DSProducerDetails) GetArtifactoryTokenAudience() string`

GetArtifactoryTokenAudience returns the ArtifactoryTokenAudience field if non-nil, zero value otherwise.

### GetArtifactoryTokenAudienceOk

`func (o *DSProducerDetails) GetArtifactoryTokenAudienceOk() (*string, bool)`

GetArtifactoryTokenAudienceOk returns a tuple with the ArtifactoryTokenAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryTokenAudience

`func (o *DSProducerDetails) SetArtifactoryTokenAudience(v string)`

SetArtifactoryTokenAudience sets ArtifactoryTokenAudience field to given value.

### HasArtifactoryTokenAudience

`func (o *DSProducerDetails) HasArtifactoryTokenAudience() bool`

HasArtifactoryTokenAudience returns a boolean if a field has been set.

### GetArtifactoryTokenScope

`func (o *DSProducerDetails) GetArtifactoryTokenScope() string`

GetArtifactoryTokenScope returns the ArtifactoryTokenScope field if non-nil, zero value otherwise.

### GetArtifactoryTokenScopeOk

`func (o *DSProducerDetails) GetArtifactoryTokenScopeOk() (*string, bool)`

GetArtifactoryTokenScopeOk returns a tuple with the ArtifactoryTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryTokenScope

`func (o *DSProducerDetails) SetArtifactoryTokenScope(v string)`

SetArtifactoryTokenScope sets ArtifactoryTokenScope field to given value.

### HasArtifactoryTokenScope

`func (o *DSProducerDetails) HasArtifactoryTokenScope() bool`

HasArtifactoryTokenScope returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *DSProducerDetails) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *DSProducerDetails) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *DSProducerDetails) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *DSProducerDetails) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsAccessMode

`func (o *DSProducerDetails) GetAwsAccessMode() string`

GetAwsAccessMode returns the AwsAccessMode field if non-nil, zero value otherwise.

### GetAwsAccessModeOk

`func (o *DSProducerDetails) GetAwsAccessModeOk() (*string, bool)`

GetAwsAccessModeOk returns a tuple with the AwsAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessMode

`func (o *DSProducerDetails) SetAwsAccessMode(v string)`

SetAwsAccessMode sets AwsAccessMode field to given value.

### HasAwsAccessMode

`func (o *DSProducerDetails) HasAwsAccessMode() bool`

HasAwsAccessMode returns a boolean if a field has been set.

### GetAwsRegion

`func (o *DSProducerDetails) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *DSProducerDetails) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *DSProducerDetails) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *DSProducerDetails) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAwsRoleArns

`func (o *DSProducerDetails) GetAwsRoleArns() string`

GetAwsRoleArns returns the AwsRoleArns field if non-nil, zero value otherwise.

### GetAwsRoleArnsOk

`func (o *DSProducerDetails) GetAwsRoleArnsOk() (*string, bool)`

GetAwsRoleArnsOk returns a tuple with the AwsRoleArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArns

`func (o *DSProducerDetails) SetAwsRoleArns(v string)`

SetAwsRoleArns sets AwsRoleArns field to given value.

### HasAwsRoleArns

`func (o *DSProducerDetails) HasAwsRoleArns() bool`

HasAwsRoleArns returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *DSProducerDetails) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *DSProducerDetails) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *DSProducerDetails) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *DSProducerDetails) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsSessionToken

`func (o *DSProducerDetails) GetAwsSessionToken() string`

GetAwsSessionToken returns the AwsSessionToken field if non-nil, zero value otherwise.

### GetAwsSessionTokenOk

`func (o *DSProducerDetails) GetAwsSessionTokenOk() (*string, bool)`

GetAwsSessionTokenOk returns a tuple with the AwsSessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSessionToken

`func (o *DSProducerDetails) SetAwsSessionToken(v string)`

SetAwsSessionToken sets AwsSessionToken field to given value.

### HasAwsSessionToken

`func (o *DSProducerDetails) HasAwsSessionToken() bool`

HasAwsSessionToken returns a boolean if a field has been set.

### GetAwsUserConsoleAccess

`func (o *DSProducerDetails) GetAwsUserConsoleAccess() bool`

GetAwsUserConsoleAccess returns the AwsUserConsoleAccess field if non-nil, zero value otherwise.

### GetAwsUserConsoleAccessOk

`func (o *DSProducerDetails) GetAwsUserConsoleAccessOk() (*bool, bool)`

GetAwsUserConsoleAccessOk returns a tuple with the AwsUserConsoleAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserConsoleAccess

`func (o *DSProducerDetails) SetAwsUserConsoleAccess(v bool)`

SetAwsUserConsoleAccess sets AwsUserConsoleAccess field to given value.

### HasAwsUserConsoleAccess

`func (o *DSProducerDetails) HasAwsUserConsoleAccess() bool`

HasAwsUserConsoleAccess returns a boolean if a field has been set.

### GetAwsUserGroups

`func (o *DSProducerDetails) GetAwsUserGroups() string`

GetAwsUserGroups returns the AwsUserGroups field if non-nil, zero value otherwise.

### GetAwsUserGroupsOk

`func (o *DSProducerDetails) GetAwsUserGroupsOk() (*string, bool)`

GetAwsUserGroupsOk returns a tuple with the AwsUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserGroups

`func (o *DSProducerDetails) SetAwsUserGroups(v string)`

SetAwsUserGroups sets AwsUserGroups field to given value.

### HasAwsUserGroups

`func (o *DSProducerDetails) HasAwsUserGroups() bool`

HasAwsUserGroups returns a boolean if a field has been set.

### GetAwsUserPolicies

`func (o *DSProducerDetails) GetAwsUserPolicies() string`

GetAwsUserPolicies returns the AwsUserPolicies field if non-nil, zero value otherwise.

### GetAwsUserPoliciesOk

`func (o *DSProducerDetails) GetAwsUserPoliciesOk() (*string, bool)`

GetAwsUserPoliciesOk returns a tuple with the AwsUserPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserPolicies

`func (o *DSProducerDetails) SetAwsUserPolicies(v string)`

SetAwsUserPolicies sets AwsUserPolicies field to given value.

### HasAwsUserPolicies

`func (o *DSProducerDetails) HasAwsUserPolicies() bool`

HasAwsUserPolicies returns a boolean if a field has been set.

### GetAwsUserProgrammaticAccess

`func (o *DSProducerDetails) GetAwsUserProgrammaticAccess() bool`

GetAwsUserProgrammaticAccess returns the AwsUserProgrammaticAccess field if non-nil, zero value otherwise.

### GetAwsUserProgrammaticAccessOk

`func (o *DSProducerDetails) GetAwsUserProgrammaticAccessOk() (*bool, bool)`

GetAwsUserProgrammaticAccessOk returns a tuple with the AwsUserProgrammaticAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserProgrammaticAccess

`func (o *DSProducerDetails) SetAwsUserProgrammaticAccess(v bool)`

SetAwsUserProgrammaticAccess sets AwsUserProgrammaticAccess field to given value.

### HasAwsUserProgrammaticAccess

`func (o *DSProducerDetails) HasAwsUserProgrammaticAccess() bool`

HasAwsUserProgrammaticAccess returns a boolean if a field has been set.

### GetAzureAppObjectId

`func (o *DSProducerDetails) GetAzureAppObjectId() string`

GetAzureAppObjectId returns the AzureAppObjectId field if non-nil, zero value otherwise.

### GetAzureAppObjectIdOk

`func (o *DSProducerDetails) GetAzureAppObjectIdOk() (*string, bool)`

GetAzureAppObjectIdOk returns a tuple with the AzureAppObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppObjectId

`func (o *DSProducerDetails) SetAzureAppObjectId(v string)`

SetAzureAppObjectId sets AzureAppObjectId field to given value.

### HasAzureAppObjectId

`func (o *DSProducerDetails) HasAzureAppObjectId() bool`

HasAzureAppObjectId returns a boolean if a field has been set.

### GetAzureClientId

`func (o *DSProducerDetails) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *DSProducerDetails) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *DSProducerDetails) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *DSProducerDetails) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureClientSecret

`func (o *DSProducerDetails) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *DSProducerDetails) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *DSProducerDetails) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### HasAzureClientSecret

`func (o *DSProducerDetails) HasAzureClientSecret() bool`

HasAzureClientSecret returns a boolean if a field has been set.

### GetAzureFixedUserNameSubClaimKey

`func (o *DSProducerDetails) GetAzureFixedUserNameSubClaimKey() string`

GetAzureFixedUserNameSubClaimKey returns the AzureFixedUserNameSubClaimKey field if non-nil, zero value otherwise.

### GetAzureFixedUserNameSubClaimKeyOk

`func (o *DSProducerDetails) GetAzureFixedUserNameSubClaimKeyOk() (*string, bool)`

GetAzureFixedUserNameSubClaimKeyOk returns a tuple with the AzureFixedUserNameSubClaimKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureFixedUserNameSubClaimKey

`func (o *DSProducerDetails) SetAzureFixedUserNameSubClaimKey(v string)`

SetAzureFixedUserNameSubClaimKey sets AzureFixedUserNameSubClaimKey field to given value.

### HasAzureFixedUserNameSubClaimKey

`func (o *DSProducerDetails) HasAzureFixedUserNameSubClaimKey() bool`

HasAzureFixedUserNameSubClaimKey returns a boolean if a field has been set.

### GetAzureFixedUserOnly

`func (o *DSProducerDetails) GetAzureFixedUserOnly() bool`

GetAzureFixedUserOnly returns the AzureFixedUserOnly field if non-nil, zero value otherwise.

### GetAzureFixedUserOnlyOk

`func (o *DSProducerDetails) GetAzureFixedUserOnlyOk() (*bool, bool)`

GetAzureFixedUserOnlyOk returns a tuple with the AzureFixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureFixedUserOnly

`func (o *DSProducerDetails) SetAzureFixedUserOnly(v bool)`

SetAzureFixedUserOnly sets AzureFixedUserOnly field to given value.

### HasAzureFixedUserOnly

`func (o *DSProducerDetails) HasAzureFixedUserOnly() bool`

HasAzureFixedUserOnly returns a boolean if a field has been set.

### GetAzureResourceGroupName

`func (o *DSProducerDetails) GetAzureResourceGroupName() string`

GetAzureResourceGroupName returns the AzureResourceGroupName field if non-nil, zero value otherwise.

### GetAzureResourceGroupNameOk

`func (o *DSProducerDetails) GetAzureResourceGroupNameOk() (*string, bool)`

GetAzureResourceGroupNameOk returns a tuple with the AzureResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroupName

`func (o *DSProducerDetails) SetAzureResourceGroupName(v string)`

SetAzureResourceGroupName sets AzureResourceGroupName field to given value.

### HasAzureResourceGroupName

`func (o *DSProducerDetails) HasAzureResourceGroupName() bool`

HasAzureResourceGroupName returns a boolean if a field has been set.

### GetAzureResourceName

`func (o *DSProducerDetails) GetAzureResourceName() string`

GetAzureResourceName returns the AzureResourceName field if non-nil, zero value otherwise.

### GetAzureResourceNameOk

`func (o *DSProducerDetails) GetAzureResourceNameOk() (*string, bool)`

GetAzureResourceNameOk returns a tuple with the AzureResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceName

`func (o *DSProducerDetails) SetAzureResourceName(v string)`

SetAzureResourceName sets AzureResourceName field to given value.

### HasAzureResourceName

`func (o *DSProducerDetails) HasAzureResourceName() bool`

HasAzureResourceName returns a boolean if a field has been set.

### GetAzureSubscriptionId

`func (o *DSProducerDetails) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *DSProducerDetails) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *DSProducerDetails) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *DSProducerDetails) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *DSProducerDetails) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *DSProducerDetails) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *DSProducerDetails) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *DSProducerDetails) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetAzureUserGroupsObjId

`func (o *DSProducerDetails) GetAzureUserGroupsObjId() string`

GetAzureUserGroupsObjId returns the AzureUserGroupsObjId field if non-nil, zero value otherwise.

### GetAzureUserGroupsObjIdOk

`func (o *DSProducerDetails) GetAzureUserGroupsObjIdOk() (*string, bool)`

GetAzureUserGroupsObjIdOk returns a tuple with the AzureUserGroupsObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureUserGroupsObjId

`func (o *DSProducerDetails) SetAzureUserGroupsObjId(v string)`

SetAzureUserGroupsObjId sets AzureUserGroupsObjId field to given value.

### HasAzureUserGroupsObjId

`func (o *DSProducerDetails) HasAzureUserGroupsObjId() bool`

HasAzureUserGroupsObjId returns a boolean if a field has been set.

### GetAzureUserPortalAccess

`func (o *DSProducerDetails) GetAzureUserPortalAccess() bool`

GetAzureUserPortalAccess returns the AzureUserPortalAccess field if non-nil, zero value otherwise.

### GetAzureUserPortalAccessOk

`func (o *DSProducerDetails) GetAzureUserPortalAccessOk() (*bool, bool)`

GetAzureUserPortalAccessOk returns a tuple with the AzureUserPortalAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureUserPortalAccess

`func (o *DSProducerDetails) SetAzureUserPortalAccess(v bool)`

SetAzureUserPortalAccess sets AzureUserPortalAccess field to given value.

### HasAzureUserPortalAccess

`func (o *DSProducerDetails) HasAzureUserPortalAccess() bool`

HasAzureUserPortalAccess returns a boolean if a field has been set.

### GetAzureUserProgrammaticAccess

`func (o *DSProducerDetails) GetAzureUserProgrammaticAccess() bool`

GetAzureUserProgrammaticAccess returns the AzureUserProgrammaticAccess field if non-nil, zero value otherwise.

### GetAzureUserProgrammaticAccessOk

`func (o *DSProducerDetails) GetAzureUserProgrammaticAccessOk() (*bool, bool)`

GetAzureUserProgrammaticAccessOk returns a tuple with the AzureUserProgrammaticAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureUserProgrammaticAccess

`func (o *DSProducerDetails) SetAzureUserProgrammaticAccess(v bool)`

SetAzureUserProgrammaticAccess sets AzureUserProgrammaticAccess field to given value.

### HasAzureUserProgrammaticAccess

`func (o *DSProducerDetails) HasAzureUserProgrammaticAccess() bool`

HasAzureUserProgrammaticAccess returns a boolean if a field has been set.

### GetAzureUserRolesTemplateId

`func (o *DSProducerDetails) GetAzureUserRolesTemplateId() string`

GetAzureUserRolesTemplateId returns the AzureUserRolesTemplateId field if non-nil, zero value otherwise.

### GetAzureUserRolesTemplateIdOk

`func (o *DSProducerDetails) GetAzureUserRolesTemplateIdOk() (*string, bool)`

GetAzureUserRolesTemplateIdOk returns a tuple with the AzureUserRolesTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureUserRolesTemplateId

`func (o *DSProducerDetails) SetAzureUserRolesTemplateId(v string)`

SetAzureUserRolesTemplateId sets AzureUserRolesTemplateId field to given value.

### HasAzureUserRolesTemplateId

`func (o *DSProducerDetails) HasAzureUserRolesTemplateId() bool`

HasAzureUserRolesTemplateId returns a boolean if a field has been set.

### GetCassandraCreationStatements

`func (o *DSProducerDetails) GetCassandraCreationStatements() string`

GetCassandraCreationStatements returns the CassandraCreationStatements field if non-nil, zero value otherwise.

### GetCassandraCreationStatementsOk

`func (o *DSProducerDetails) GetCassandraCreationStatementsOk() (*string, bool)`

GetCassandraCreationStatementsOk returns a tuple with the CassandraCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCreationStatements

`func (o *DSProducerDetails) SetCassandraCreationStatements(v string)`

SetCassandraCreationStatements sets CassandraCreationStatements field to given value.

### HasCassandraCreationStatements

`func (o *DSProducerDetails) HasCassandraCreationStatements() bool`

HasCassandraCreationStatements returns a boolean if a field has been set.

### GetChefOrganizations

`func (o *DSProducerDetails) GetChefOrganizations() string`

GetChefOrganizations returns the ChefOrganizations field if non-nil, zero value otherwise.

### GetChefOrganizationsOk

`func (o *DSProducerDetails) GetChefOrganizationsOk() (*string, bool)`

GetChefOrganizationsOk returns a tuple with the ChefOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefOrganizations

`func (o *DSProducerDetails) SetChefOrganizations(v string)`

SetChefOrganizations sets ChefOrganizations field to given value.

### HasChefOrganizations

`func (o *DSProducerDetails) HasChefOrganizations() bool`

HasChefOrganizations returns a boolean if a field has been set.

### GetChefServerAccessMode

`func (o *DSProducerDetails) GetChefServerAccessMode() string`

GetChefServerAccessMode returns the ChefServerAccessMode field if non-nil, zero value otherwise.

### GetChefServerAccessModeOk

`func (o *DSProducerDetails) GetChefServerAccessModeOk() (*string, bool)`

GetChefServerAccessModeOk returns a tuple with the ChefServerAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerAccessMode

`func (o *DSProducerDetails) SetChefServerAccessMode(v string)`

SetChefServerAccessMode sets ChefServerAccessMode field to given value.

### HasChefServerAccessMode

`func (o *DSProducerDetails) HasChefServerAccessMode() bool`

HasChefServerAccessMode returns a boolean if a field has been set.

### GetChefServerHostName

`func (o *DSProducerDetails) GetChefServerHostName() string`

GetChefServerHostName returns the ChefServerHostName field if non-nil, zero value otherwise.

### GetChefServerHostNameOk

`func (o *DSProducerDetails) GetChefServerHostNameOk() (*string, bool)`

GetChefServerHostNameOk returns a tuple with the ChefServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerHostName

`func (o *DSProducerDetails) SetChefServerHostName(v string)`

SetChefServerHostName sets ChefServerHostName field to given value.

### HasChefServerHostName

`func (o *DSProducerDetails) HasChefServerHostName() bool`

HasChefServerHostName returns a boolean if a field has been set.

### GetChefServerKey

`func (o *DSProducerDetails) GetChefServerKey() string`

GetChefServerKey returns the ChefServerKey field if non-nil, zero value otherwise.

### GetChefServerKeyOk

`func (o *DSProducerDetails) GetChefServerKeyOk() (*string, bool)`

GetChefServerKeyOk returns a tuple with the ChefServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerKey

`func (o *DSProducerDetails) SetChefServerKey(v string)`

SetChefServerKey sets ChefServerKey field to given value.

### HasChefServerKey

`func (o *DSProducerDetails) HasChefServerKey() bool`

HasChefServerKey returns a boolean if a field has been set.

### GetChefServerPort

`func (o *DSProducerDetails) GetChefServerPort() string`

GetChefServerPort returns the ChefServerPort field if non-nil, zero value otherwise.

### GetChefServerPortOk

`func (o *DSProducerDetails) GetChefServerPortOk() (*string, bool)`

GetChefServerPortOk returns a tuple with the ChefServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerPort

`func (o *DSProducerDetails) SetChefServerPort(v string)`

SetChefServerPort sets ChefServerPort field to given value.

### HasChefServerPort

`func (o *DSProducerDetails) HasChefServerPort() bool`

HasChefServerPort returns a boolean if a field has been set.

### GetChefServerUrl

`func (o *DSProducerDetails) GetChefServerUrl() string`

GetChefServerUrl returns the ChefServerUrl field if non-nil, zero value otherwise.

### GetChefServerUrlOk

`func (o *DSProducerDetails) GetChefServerUrlOk() (*string, bool)`

GetChefServerUrlOk returns a tuple with the ChefServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUrl

`func (o *DSProducerDetails) SetChefServerUrl(v string)`

SetChefServerUrl sets ChefServerUrl field to given value.

### HasChefServerUrl

`func (o *DSProducerDetails) HasChefServerUrl() bool`

HasChefServerUrl returns a boolean if a field has been set.

### GetChefServerUsername

`func (o *DSProducerDetails) GetChefServerUsername() string`

GetChefServerUsername returns the ChefServerUsername field if non-nil, zero value otherwise.

### GetChefServerUsernameOk

`func (o *DSProducerDetails) GetChefServerUsernameOk() (*string, bool)`

GetChefServerUsernameOk returns a tuple with the ChefServerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUsername

`func (o *DSProducerDetails) SetChefServerUsername(v string)`

SetChefServerUsername sets ChefServerUsername field to given value.

### HasChefServerUsername

`func (o *DSProducerDetails) HasChefServerUsername() bool`

HasChefServerUsername returns a boolean if a field has been set.

### GetChefSkipSsl

`func (o *DSProducerDetails) GetChefSkipSsl() bool`

GetChefSkipSsl returns the ChefSkipSsl field if non-nil, zero value otherwise.

### GetChefSkipSslOk

`func (o *DSProducerDetails) GetChefSkipSslOk() (*bool, bool)`

GetChefSkipSslOk returns a tuple with the ChefSkipSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefSkipSsl

`func (o *DSProducerDetails) SetChefSkipSsl(v bool)`

SetChefSkipSsl sets ChefSkipSsl field to given value.

### HasChefSkipSsl

`func (o *DSProducerDetails) HasChefSkipSsl() bool`

HasChefSkipSsl returns a boolean if a field has been set.

### GetCreateSyncUrl

`func (o *DSProducerDetails) GetCreateSyncUrl() string`

GetCreateSyncUrl returns the CreateSyncUrl field if non-nil, zero value otherwise.

### GetCreateSyncUrlOk

`func (o *DSProducerDetails) GetCreateSyncUrlOk() (*string, bool)`

GetCreateSyncUrlOk returns a tuple with the CreateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSyncUrl

`func (o *DSProducerDetails) SetCreateSyncUrl(v string)`

SetCreateSyncUrl sets CreateSyncUrl field to given value.

### HasCreateSyncUrl

`func (o *DSProducerDetails) HasCreateSyncUrl() bool`

HasCreateSyncUrl returns a boolean if a field has been set.

### GetDbHostName

`func (o *DSProducerDetails) GetDbHostName() string`

GetDbHostName returns the DbHostName field if non-nil, zero value otherwise.

### GetDbHostNameOk

`func (o *DSProducerDetails) GetDbHostNameOk() (*string, bool)`

GetDbHostNameOk returns a tuple with the DbHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHostName

`func (o *DSProducerDetails) SetDbHostName(v string)`

SetDbHostName sets DbHostName field to given value.

### HasDbHostName

`func (o *DSProducerDetails) HasDbHostName() bool`

HasDbHostName returns a boolean if a field has been set.

### GetDbIsolationLevel

`func (o *DSProducerDetails) GetDbIsolationLevel() string`

GetDbIsolationLevel returns the DbIsolationLevel field if non-nil, zero value otherwise.

### GetDbIsolationLevelOk

`func (o *DSProducerDetails) GetDbIsolationLevelOk() (*string, bool)`

GetDbIsolationLevelOk returns a tuple with the DbIsolationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbIsolationLevel

`func (o *DSProducerDetails) SetDbIsolationLevel(v string)`

SetDbIsolationLevel sets DbIsolationLevel field to given value.

### HasDbIsolationLevel

`func (o *DSProducerDetails) HasDbIsolationLevel() bool`

HasDbIsolationLevel returns a boolean if a field has been set.

### GetDbMaxIdleConns

`func (o *DSProducerDetails) GetDbMaxIdleConns() string`

GetDbMaxIdleConns returns the DbMaxIdleConns field if non-nil, zero value otherwise.

### GetDbMaxIdleConnsOk

`func (o *DSProducerDetails) GetDbMaxIdleConnsOk() (*string, bool)`

GetDbMaxIdleConnsOk returns a tuple with the DbMaxIdleConns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbMaxIdleConns

`func (o *DSProducerDetails) SetDbMaxIdleConns(v string)`

SetDbMaxIdleConns sets DbMaxIdleConns field to given value.

### HasDbMaxIdleConns

`func (o *DSProducerDetails) HasDbMaxIdleConns() bool`

HasDbMaxIdleConns returns a boolean if a field has been set.

### GetDbMaxOpenConns

`func (o *DSProducerDetails) GetDbMaxOpenConns() string`

GetDbMaxOpenConns returns the DbMaxOpenConns field if non-nil, zero value otherwise.

### GetDbMaxOpenConnsOk

`func (o *DSProducerDetails) GetDbMaxOpenConnsOk() (*string, bool)`

GetDbMaxOpenConnsOk returns a tuple with the DbMaxOpenConns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbMaxOpenConns

`func (o *DSProducerDetails) SetDbMaxOpenConns(v string)`

SetDbMaxOpenConns sets DbMaxOpenConns field to given value.

### HasDbMaxOpenConns

`func (o *DSProducerDetails) HasDbMaxOpenConns() bool`

HasDbMaxOpenConns returns a boolean if a field has been set.

### GetDbName

`func (o *DSProducerDetails) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DSProducerDetails) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DSProducerDetails) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DSProducerDetails) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDbPort

`func (o *DSProducerDetails) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *DSProducerDetails) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *DSProducerDetails) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.

### HasDbPort

`func (o *DSProducerDetails) HasDbPort() bool`

HasDbPort returns a boolean if a field has been set.

### GetDbPwd

`func (o *DSProducerDetails) GetDbPwd() string`

GetDbPwd returns the DbPwd field if non-nil, zero value otherwise.

### GetDbPwdOk

`func (o *DSProducerDetails) GetDbPwdOk() (*string, bool)`

GetDbPwdOk returns a tuple with the DbPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPwd

`func (o *DSProducerDetails) SetDbPwd(v string)`

SetDbPwd sets DbPwd field to given value.

### HasDbPwd

`func (o *DSProducerDetails) HasDbPwd() bool`

HasDbPwd returns a boolean if a field has been set.

### GetDbServerCertificates

`func (o *DSProducerDetails) GetDbServerCertificates() string`

GetDbServerCertificates returns the DbServerCertificates field if non-nil, zero value otherwise.

### GetDbServerCertificatesOk

`func (o *DSProducerDetails) GetDbServerCertificatesOk() (*string, bool)`

GetDbServerCertificatesOk returns a tuple with the DbServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerCertificates

`func (o *DSProducerDetails) SetDbServerCertificates(v string)`

SetDbServerCertificates sets DbServerCertificates field to given value.

### HasDbServerCertificates

`func (o *DSProducerDetails) HasDbServerCertificates() bool`

HasDbServerCertificates returns a boolean if a field has been set.

### GetDbServerName

`func (o *DSProducerDetails) GetDbServerName() string`

GetDbServerName returns the DbServerName field if non-nil, zero value otherwise.

### GetDbServerNameOk

`func (o *DSProducerDetails) GetDbServerNameOk() (*string, bool)`

GetDbServerNameOk returns a tuple with the DbServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbServerName

`func (o *DSProducerDetails) SetDbServerName(v string)`

SetDbServerName sets DbServerName field to given value.

### HasDbServerName

`func (o *DSProducerDetails) HasDbServerName() bool`

HasDbServerName returns a boolean if a field has been set.

### GetDbUserName

`func (o *DSProducerDetails) GetDbUserName() string`

GetDbUserName returns the DbUserName field if non-nil, zero value otherwise.

### GetDbUserNameOk

`func (o *DSProducerDetails) GetDbUserNameOk() (*string, bool)`

GetDbUserNameOk returns a tuple with the DbUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUserName

`func (o *DSProducerDetails) SetDbUserName(v string)`

SetDbUserName sets DbUserName field to given value.

### HasDbUserName

`func (o *DSProducerDetails) HasDbUserName() bool`

HasDbUserName returns a boolean if a field has been set.

### GetDynamicSecretId

`func (o *DSProducerDetails) GetDynamicSecretId() int64`

GetDynamicSecretId returns the DynamicSecretId field if non-nil, zero value otherwise.

### GetDynamicSecretIdOk

`func (o *DSProducerDetails) GetDynamicSecretIdOk() (*int64, bool)`

GetDynamicSecretIdOk returns a tuple with the DynamicSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretId

`func (o *DSProducerDetails) SetDynamicSecretId(v int64)`

SetDynamicSecretId sets DynamicSecretId field to given value.

### HasDynamicSecretId

`func (o *DSProducerDetails) HasDynamicSecretId() bool`

HasDynamicSecretId returns a boolean if a field has been set.

### GetDynamicSecretKey

`func (o *DSProducerDetails) GetDynamicSecretKey() string`

GetDynamicSecretKey returns the DynamicSecretKey field if non-nil, zero value otherwise.

### GetDynamicSecretKeyOk

`func (o *DSProducerDetails) GetDynamicSecretKeyOk() (*string, bool)`

GetDynamicSecretKeyOk returns a tuple with the DynamicSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretKey

`func (o *DSProducerDetails) SetDynamicSecretKey(v string)`

SetDynamicSecretKey sets DynamicSecretKey field to given value.

### HasDynamicSecretKey

`func (o *DSProducerDetails) HasDynamicSecretKey() bool`

HasDynamicSecretKey returns a boolean if a field has been set.

### GetDynamicSecretName

`func (o *DSProducerDetails) GetDynamicSecretName() string`

GetDynamicSecretName returns the DynamicSecretName field if non-nil, zero value otherwise.

### GetDynamicSecretNameOk

`func (o *DSProducerDetails) GetDynamicSecretNameOk() (*string, bool)`

GetDynamicSecretNameOk returns a tuple with the DynamicSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretName

`func (o *DSProducerDetails) SetDynamicSecretName(v string)`

SetDynamicSecretName sets DynamicSecretName field to given value.

### HasDynamicSecretName

`func (o *DSProducerDetails) HasDynamicSecretName() bool`

HasDynamicSecretName returns a boolean if a field has been set.

### GetDynamicSecretType

`func (o *DSProducerDetails) GetDynamicSecretType() string`

GetDynamicSecretType returns the DynamicSecretType field if non-nil, zero value otherwise.

### GetDynamicSecretTypeOk

`func (o *DSProducerDetails) GetDynamicSecretTypeOk() (*string, bool)`

GetDynamicSecretTypeOk returns a tuple with the DynamicSecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretType

`func (o *DSProducerDetails) SetDynamicSecretType(v string)`

SetDynamicSecretType sets DynamicSecretType field to given value.

### HasDynamicSecretType

`func (o *DSProducerDetails) HasDynamicSecretType() bool`

HasDynamicSecretType returns a boolean if a field has been set.

### GetEksAccessKeyId

`func (o *DSProducerDetails) GetEksAccessKeyId() string`

GetEksAccessKeyId returns the EksAccessKeyId field if non-nil, zero value otherwise.

### GetEksAccessKeyIdOk

`func (o *DSProducerDetails) GetEksAccessKeyIdOk() (*string, bool)`

GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAccessKeyId

`func (o *DSProducerDetails) SetEksAccessKeyId(v string)`

SetEksAccessKeyId sets EksAccessKeyId field to given value.

### HasEksAccessKeyId

`func (o *DSProducerDetails) HasEksAccessKeyId() bool`

HasEksAccessKeyId returns a boolean if a field has been set.

### GetEksAssumeRole

`func (o *DSProducerDetails) GetEksAssumeRole() string`

GetEksAssumeRole returns the EksAssumeRole field if non-nil, zero value otherwise.

### GetEksAssumeRoleOk

`func (o *DSProducerDetails) GetEksAssumeRoleOk() (*string, bool)`

GetEksAssumeRoleOk returns a tuple with the EksAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAssumeRole

`func (o *DSProducerDetails) SetEksAssumeRole(v string)`

SetEksAssumeRole sets EksAssumeRole field to given value.

### HasEksAssumeRole

`func (o *DSProducerDetails) HasEksAssumeRole() bool`

HasEksAssumeRole returns a boolean if a field has been set.

### GetEksClusterCaCertificate

`func (o *DSProducerDetails) GetEksClusterCaCertificate() string`

GetEksClusterCaCertificate returns the EksClusterCaCertificate field if non-nil, zero value otherwise.

### GetEksClusterCaCertificateOk

`func (o *DSProducerDetails) GetEksClusterCaCertificateOk() (*string, bool)`

GetEksClusterCaCertificateOk returns a tuple with the EksClusterCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterCaCertificate

`func (o *DSProducerDetails) SetEksClusterCaCertificate(v string)`

SetEksClusterCaCertificate sets EksClusterCaCertificate field to given value.

### HasEksClusterCaCertificate

`func (o *DSProducerDetails) HasEksClusterCaCertificate() bool`

HasEksClusterCaCertificate returns a boolean if a field has been set.

### GetEksClusterEndpoint

`func (o *DSProducerDetails) GetEksClusterEndpoint() string`

GetEksClusterEndpoint returns the EksClusterEndpoint field if non-nil, zero value otherwise.

### GetEksClusterEndpointOk

`func (o *DSProducerDetails) GetEksClusterEndpointOk() (*string, bool)`

GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterEndpoint

`func (o *DSProducerDetails) SetEksClusterEndpoint(v string)`

SetEksClusterEndpoint sets EksClusterEndpoint field to given value.

### HasEksClusterEndpoint

`func (o *DSProducerDetails) HasEksClusterEndpoint() bool`

HasEksClusterEndpoint returns a boolean if a field has been set.

### GetEksClusterName

`func (o *DSProducerDetails) GetEksClusterName() string`

GetEksClusterName returns the EksClusterName field if non-nil, zero value otherwise.

### GetEksClusterNameOk

`func (o *DSProducerDetails) GetEksClusterNameOk() (*string, bool)`

GetEksClusterNameOk returns a tuple with the EksClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterName

`func (o *DSProducerDetails) SetEksClusterName(v string)`

SetEksClusterName sets EksClusterName field to given value.

### HasEksClusterName

`func (o *DSProducerDetails) HasEksClusterName() bool`

HasEksClusterName returns a boolean if a field has been set.

### GetEksRegion

`func (o *DSProducerDetails) GetEksRegion() string`

GetEksRegion returns the EksRegion field if non-nil, zero value otherwise.

### GetEksRegionOk

`func (o *DSProducerDetails) GetEksRegionOk() (*string, bool)`

GetEksRegionOk returns a tuple with the EksRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksRegion

`func (o *DSProducerDetails) SetEksRegion(v string)`

SetEksRegion sets EksRegion field to given value.

### HasEksRegion

`func (o *DSProducerDetails) HasEksRegion() bool`

HasEksRegion returns a boolean if a field has been set.

### GetEksSecretAccessKey

`func (o *DSProducerDetails) GetEksSecretAccessKey() string`

GetEksSecretAccessKey returns the EksSecretAccessKey field if non-nil, zero value otherwise.

### GetEksSecretAccessKeyOk

`func (o *DSProducerDetails) GetEksSecretAccessKeyOk() (*string, bool)`

GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSecretAccessKey

`func (o *DSProducerDetails) SetEksSecretAccessKey(v string)`

SetEksSecretAccessKey sets EksSecretAccessKey field to given value.

### HasEksSecretAccessKey

`func (o *DSProducerDetails) HasEksSecretAccessKey() bool`

HasEksSecretAccessKey returns a boolean if a field has been set.

### GetEnableAdminRotation

`func (o *DSProducerDetails) GetEnableAdminRotation() bool`

GetEnableAdminRotation returns the EnableAdminRotation field if non-nil, zero value otherwise.

### GetEnableAdminRotationOk

`func (o *DSProducerDetails) GetEnableAdminRotationOk() (*bool, bool)`

GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdminRotation

`func (o *DSProducerDetails) SetEnableAdminRotation(v bool)`

SetEnableAdminRotation sets EnableAdminRotation field to given value.

### HasEnableAdminRotation

`func (o *DSProducerDetails) HasEnableAdminRotation() bool`

HasEnableAdminRotation returns a boolean if a field has been set.

### GetExternallyProvidedUser

`func (o *DSProducerDetails) GetExternallyProvidedUser() string`

GetExternallyProvidedUser returns the ExternallyProvidedUser field if non-nil, zero value otherwise.

### GetExternallyProvidedUserOk

`func (o *DSProducerDetails) GetExternallyProvidedUserOk() (*string, bool)`

GetExternallyProvidedUserOk returns a tuple with the ExternallyProvidedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyProvidedUser

`func (o *DSProducerDetails) SetExternallyProvidedUser(v string)`

SetExternallyProvidedUser sets ExternallyProvidedUser field to given value.

### HasExternallyProvidedUser

`func (o *DSProducerDetails) HasExternallyProvidedUser() bool`

HasExternallyProvidedUser returns a boolean if a field has been set.

### GetFailureMessage

`func (o *DSProducerDetails) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *DSProducerDetails) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *DSProducerDetails) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *DSProducerDetails) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetFixedUserOnly

`func (o *DSProducerDetails) GetFixedUserOnly() string`

GetFixedUserOnly returns the FixedUserOnly field if non-nil, zero value otherwise.

### GetFixedUserOnlyOk

`func (o *DSProducerDetails) GetFixedUserOnlyOk() (*string, bool)`

GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserOnly

`func (o *DSProducerDetails) SetFixedUserOnly(v string)`

SetFixedUserOnly sets FixedUserOnly field to given value.

### HasFixedUserOnly

`func (o *DSProducerDetails) HasFixedUserOnly() bool`

HasFixedUserOnly returns a boolean if a field has been set.

### GetGcpKeyAlgo

`func (o *DSProducerDetails) GetGcpKeyAlgo() string`

GetGcpKeyAlgo returns the GcpKeyAlgo field if non-nil, zero value otherwise.

### GetGcpKeyAlgoOk

`func (o *DSProducerDetails) GetGcpKeyAlgoOk() (*string, bool)`

GetGcpKeyAlgoOk returns a tuple with the GcpKeyAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKeyAlgo

`func (o *DSProducerDetails) SetGcpKeyAlgo(v string)`

SetGcpKeyAlgo sets GcpKeyAlgo field to given value.

### HasGcpKeyAlgo

`func (o *DSProducerDetails) HasGcpKeyAlgo() bool`

HasGcpKeyAlgo returns a boolean if a field has been set.

### GetGcpServiceAccountEmail

`func (o *DSProducerDetails) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *DSProducerDetails) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *DSProducerDetails) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.

### HasGcpServiceAccountEmail

`func (o *DSProducerDetails) HasGcpServiceAccountEmail() bool`

HasGcpServiceAccountEmail returns a boolean if a field has been set.

### GetGcpServiceAccountKey

`func (o *DSProducerDetails) GetGcpServiceAccountKey() string`

GetGcpServiceAccountKey returns the GcpServiceAccountKey field if non-nil, zero value otherwise.

### GetGcpServiceAccountKeyOk

`func (o *DSProducerDetails) GetGcpServiceAccountKeyOk() (*string, bool)`

GetGcpServiceAccountKeyOk returns a tuple with the GcpServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountKey

`func (o *DSProducerDetails) SetGcpServiceAccountKey(v string)`

SetGcpServiceAccountKey sets GcpServiceAccountKey field to given value.

### HasGcpServiceAccountKey

`func (o *DSProducerDetails) HasGcpServiceAccountKey() bool`

HasGcpServiceAccountKey returns a boolean if a field has been set.

### GetGcpTokenLifetime

`func (o *DSProducerDetails) GetGcpTokenLifetime() string`

GetGcpTokenLifetime returns the GcpTokenLifetime field if non-nil, zero value otherwise.

### GetGcpTokenLifetimeOk

`func (o *DSProducerDetails) GetGcpTokenLifetimeOk() (*string, bool)`

GetGcpTokenLifetimeOk returns a tuple with the GcpTokenLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTokenLifetime

`func (o *DSProducerDetails) SetGcpTokenLifetime(v string)`

SetGcpTokenLifetime sets GcpTokenLifetime field to given value.

### HasGcpTokenLifetime

`func (o *DSProducerDetails) HasGcpTokenLifetime() bool`

HasGcpTokenLifetime returns a boolean if a field has been set.

### GetGcpTokenScope

`func (o *DSProducerDetails) GetGcpTokenScope() string`

GetGcpTokenScope returns the GcpTokenScope field if non-nil, zero value otherwise.

### GetGcpTokenScopeOk

`func (o *DSProducerDetails) GetGcpTokenScopeOk() (*string, bool)`

GetGcpTokenScopeOk returns a tuple with the GcpTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTokenScope

`func (o *DSProducerDetails) SetGcpTokenScope(v string)`

SetGcpTokenScope sets GcpTokenScope field to given value.

### HasGcpTokenScope

`func (o *DSProducerDetails) HasGcpTokenScope() bool`

HasGcpTokenScope returns a boolean if a field has been set.

### GetGcpTokenType

`func (o *DSProducerDetails) GetGcpTokenType() string`

GetGcpTokenType returns the GcpTokenType field if non-nil, zero value otherwise.

### GetGcpTokenTypeOk

`func (o *DSProducerDetails) GetGcpTokenTypeOk() (*string, bool)`

GetGcpTokenTypeOk returns a tuple with the GcpTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTokenType

`func (o *DSProducerDetails) SetGcpTokenType(v string)`

SetGcpTokenType sets GcpTokenType field to given value.

### HasGcpTokenType

`func (o *DSProducerDetails) HasGcpTokenType() bool`

HasGcpTokenType returns a boolean if a field has been set.

### GetGithubAppId

`func (o *DSProducerDetails) GetGithubAppId() int64`

GetGithubAppId returns the GithubAppId field if non-nil, zero value otherwise.

### GetGithubAppIdOk

`func (o *DSProducerDetails) GetGithubAppIdOk() (*int64, bool)`

GetGithubAppIdOk returns a tuple with the GithubAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppId

`func (o *DSProducerDetails) SetGithubAppId(v int64)`

SetGithubAppId sets GithubAppId field to given value.

### HasGithubAppId

`func (o *DSProducerDetails) HasGithubAppId() bool`

HasGithubAppId returns a boolean if a field has been set.

### GetGithubAppPrivateKey

`func (o *DSProducerDetails) GetGithubAppPrivateKey() string`

GetGithubAppPrivateKey returns the GithubAppPrivateKey field if non-nil, zero value otherwise.

### GetGithubAppPrivateKeyOk

`func (o *DSProducerDetails) GetGithubAppPrivateKeyOk() (*string, bool)`

GetGithubAppPrivateKeyOk returns a tuple with the GithubAppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubAppPrivateKey

`func (o *DSProducerDetails) SetGithubAppPrivateKey(v string)`

SetGithubAppPrivateKey sets GithubAppPrivateKey field to given value.

### HasGithubAppPrivateKey

`func (o *DSProducerDetails) HasGithubAppPrivateKey() bool`

HasGithubAppPrivateKey returns a boolean if a field has been set.

### GetGithubBaseUrl

`func (o *DSProducerDetails) GetGithubBaseUrl() string`

GetGithubBaseUrl returns the GithubBaseUrl field if non-nil, zero value otherwise.

### GetGithubBaseUrlOk

`func (o *DSProducerDetails) GetGithubBaseUrlOk() (*string, bool)`

GetGithubBaseUrlOk returns a tuple with the GithubBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubBaseUrl

`func (o *DSProducerDetails) SetGithubBaseUrl(v string)`

SetGithubBaseUrl sets GithubBaseUrl field to given value.

### HasGithubBaseUrl

`func (o *DSProducerDetails) HasGithubBaseUrl() bool`

HasGithubBaseUrl returns a boolean if a field has been set.

### GetGithubInstallationId

`func (o *DSProducerDetails) GetGithubInstallationId() int64`

GetGithubInstallationId returns the GithubInstallationId field if non-nil, zero value otherwise.

### GetGithubInstallationIdOk

`func (o *DSProducerDetails) GetGithubInstallationIdOk() (*int64, bool)`

GetGithubInstallationIdOk returns a tuple with the GithubInstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubInstallationId

`func (o *DSProducerDetails) SetGithubInstallationId(v int64)`

SetGithubInstallationId sets GithubInstallationId field to given value.

### HasGithubInstallationId

`func (o *DSProducerDetails) HasGithubInstallationId() bool`

HasGithubInstallationId returns a boolean if a field has been set.

### GetGithubInstallationTokenPermissions

`func (o *DSProducerDetails) GetGithubInstallationTokenPermissions() map[string]string`

GetGithubInstallationTokenPermissions returns the GithubInstallationTokenPermissions field if non-nil, zero value otherwise.

### GetGithubInstallationTokenPermissionsOk

`func (o *DSProducerDetails) GetGithubInstallationTokenPermissionsOk() (*map[string]string, bool)`

GetGithubInstallationTokenPermissionsOk returns a tuple with the GithubInstallationTokenPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubInstallationTokenPermissions

`func (o *DSProducerDetails) SetGithubInstallationTokenPermissions(v map[string]string)`

SetGithubInstallationTokenPermissions sets GithubInstallationTokenPermissions field to given value.

### HasGithubInstallationTokenPermissions

`func (o *DSProducerDetails) HasGithubInstallationTokenPermissions() bool`

HasGithubInstallationTokenPermissions returns a boolean if a field has been set.

### GetGithubInstallationTokenRepositories

`func (o *DSProducerDetails) GetGithubInstallationTokenRepositories() []string`

GetGithubInstallationTokenRepositories returns the GithubInstallationTokenRepositories field if non-nil, zero value otherwise.

### GetGithubInstallationTokenRepositoriesOk

`func (o *DSProducerDetails) GetGithubInstallationTokenRepositoriesOk() (*[]string, bool)`

GetGithubInstallationTokenRepositoriesOk returns a tuple with the GithubInstallationTokenRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubInstallationTokenRepositories

`func (o *DSProducerDetails) SetGithubInstallationTokenRepositories(v []string)`

SetGithubInstallationTokenRepositories sets GithubInstallationTokenRepositories field to given value.

### HasGithubInstallationTokenRepositories

`func (o *DSProducerDetails) HasGithubInstallationTokenRepositories() bool`

HasGithubInstallationTokenRepositories returns a boolean if a field has been set.

### GetGithubInstallationTokenRepositoriesIds

`func (o *DSProducerDetails) GetGithubInstallationTokenRepositoriesIds() []int64`

GetGithubInstallationTokenRepositoriesIds returns the GithubInstallationTokenRepositoriesIds field if non-nil, zero value otherwise.

### GetGithubInstallationTokenRepositoriesIdsOk

`func (o *DSProducerDetails) GetGithubInstallationTokenRepositoriesIdsOk() (*[]int64, bool)`

GetGithubInstallationTokenRepositoriesIdsOk returns a tuple with the GithubInstallationTokenRepositoriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubInstallationTokenRepositoriesIds

`func (o *DSProducerDetails) SetGithubInstallationTokenRepositoriesIds(v []int64)`

SetGithubInstallationTokenRepositoriesIds sets GithubInstallationTokenRepositoriesIds field to given value.

### HasGithubInstallationTokenRepositoriesIds

`func (o *DSProducerDetails) HasGithubInstallationTokenRepositoriesIds() bool`

HasGithubInstallationTokenRepositoriesIds returns a boolean if a field has been set.

### GetGithubRepositoryPath

`func (o *DSProducerDetails) GetGithubRepositoryPath() string`

GetGithubRepositoryPath returns the GithubRepositoryPath field if non-nil, zero value otherwise.

### GetGithubRepositoryPathOk

`func (o *DSProducerDetails) GetGithubRepositoryPathOk() (*string, bool)`

GetGithubRepositoryPathOk returns a tuple with the GithubRepositoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubRepositoryPath

`func (o *DSProducerDetails) SetGithubRepositoryPath(v string)`

SetGithubRepositoryPath sets GithubRepositoryPath field to given value.

### HasGithubRepositoryPath

`func (o *DSProducerDetails) HasGithubRepositoryPath() bool`

HasGithubRepositoryPath returns a boolean if a field has been set.

### GetGkeClusterCaCertificate

`func (o *DSProducerDetails) GetGkeClusterCaCertificate() string`

GetGkeClusterCaCertificate returns the GkeClusterCaCertificate field if non-nil, zero value otherwise.

### GetGkeClusterCaCertificateOk

`func (o *DSProducerDetails) GetGkeClusterCaCertificateOk() (*string, bool)`

GetGkeClusterCaCertificateOk returns a tuple with the GkeClusterCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterCaCertificate

`func (o *DSProducerDetails) SetGkeClusterCaCertificate(v string)`

SetGkeClusterCaCertificate sets GkeClusterCaCertificate field to given value.

### HasGkeClusterCaCertificate

`func (o *DSProducerDetails) HasGkeClusterCaCertificate() bool`

HasGkeClusterCaCertificate returns a boolean if a field has been set.

### GetGkeClusterEndpoint

`func (o *DSProducerDetails) GetGkeClusterEndpoint() string`

GetGkeClusterEndpoint returns the GkeClusterEndpoint field if non-nil, zero value otherwise.

### GetGkeClusterEndpointOk

`func (o *DSProducerDetails) GetGkeClusterEndpointOk() (*string, bool)`

GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterEndpoint

`func (o *DSProducerDetails) SetGkeClusterEndpoint(v string)`

SetGkeClusterEndpoint sets GkeClusterEndpoint field to given value.

### HasGkeClusterEndpoint

`func (o *DSProducerDetails) HasGkeClusterEndpoint() bool`

HasGkeClusterEndpoint returns a boolean if a field has been set.

### GetGkeClusterName

`func (o *DSProducerDetails) GetGkeClusterName() string`

GetGkeClusterName returns the GkeClusterName field if non-nil, zero value otherwise.

### GetGkeClusterNameOk

`func (o *DSProducerDetails) GetGkeClusterNameOk() (*string, bool)`

GetGkeClusterNameOk returns a tuple with the GkeClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterName

`func (o *DSProducerDetails) SetGkeClusterName(v string)`

SetGkeClusterName sets GkeClusterName field to given value.

### HasGkeClusterName

`func (o *DSProducerDetails) HasGkeClusterName() bool`

HasGkeClusterName returns a boolean if a field has been set.

### GetGkeServiceAccountKey

`func (o *DSProducerDetails) GetGkeServiceAccountKey() string`

GetGkeServiceAccountKey returns the GkeServiceAccountKey field if non-nil, zero value otherwise.

### GetGkeServiceAccountKeyOk

`func (o *DSProducerDetails) GetGkeServiceAccountKeyOk() (*string, bool)`

GetGkeServiceAccountKeyOk returns a tuple with the GkeServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeServiceAccountKey

`func (o *DSProducerDetails) SetGkeServiceAccountKey(v string)`

SetGkeServiceAccountKey sets GkeServiceAccountKey field to given value.

### HasGkeServiceAccountKey

`func (o *DSProducerDetails) HasGkeServiceAccountKey() bool`

HasGkeServiceAccountKey returns a boolean if a field has been set.

### GetGkeServiceAccountName

`func (o *DSProducerDetails) GetGkeServiceAccountName() string`

GetGkeServiceAccountName returns the GkeServiceAccountName field if non-nil, zero value otherwise.

### GetGkeServiceAccountNameOk

`func (o *DSProducerDetails) GetGkeServiceAccountNameOk() (*string, bool)`

GetGkeServiceAccountNameOk returns a tuple with the GkeServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeServiceAccountName

`func (o *DSProducerDetails) SetGkeServiceAccountName(v string)`

SetGkeServiceAccountName sets GkeServiceAccountName field to given value.

### HasGkeServiceAccountName

`func (o *DSProducerDetails) HasGkeServiceAccountName() bool`

HasGkeServiceAccountName returns a boolean if a field has been set.

### GetGroups

`func (o *DSProducerDetails) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *DSProducerDetails) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *DSProducerDetails) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *DSProducerDetails) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHanadbCreationStatements

`func (o *DSProducerDetails) GetHanadbCreationStatements() string`

GetHanadbCreationStatements returns the HanadbCreationStatements field if non-nil, zero value otherwise.

### GetHanadbCreationStatementsOk

`func (o *DSProducerDetails) GetHanadbCreationStatementsOk() (*string, bool)`

GetHanadbCreationStatementsOk returns a tuple with the HanadbCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbCreationStatements

`func (o *DSProducerDetails) SetHanadbCreationStatements(v string)`

SetHanadbCreationStatements sets HanadbCreationStatements field to given value.

### HasHanadbCreationStatements

`func (o *DSProducerDetails) HasHanadbCreationStatements() bool`

HasHanadbCreationStatements returns a boolean if a field has been set.

### GetHanadbRevocationStatements

`func (o *DSProducerDetails) GetHanadbRevocationStatements() string`

GetHanadbRevocationStatements returns the HanadbRevocationStatements field if non-nil, zero value otherwise.

### GetHanadbRevocationStatementsOk

`func (o *DSProducerDetails) GetHanadbRevocationStatementsOk() (*string, bool)`

GetHanadbRevocationStatementsOk returns a tuple with the HanadbRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanadbRevocationStatements

`func (o *DSProducerDetails) SetHanadbRevocationStatements(v string)`

SetHanadbRevocationStatements sets HanadbRevocationStatements field to given value.

### HasHanadbRevocationStatements

`func (o *DSProducerDetails) HasHanadbRevocationStatements() bool`

HasHanadbRevocationStatements returns a boolean if a field has been set.

### GetHostName

`func (o *DSProducerDetails) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *DSProducerDetails) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *DSProducerDetails) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *DSProducerDetails) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHostPort

`func (o *DSProducerDetails) GetHostPort() string`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *DSProducerDetails) GetHostPortOk() (*string, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *DSProducerDetails) SetHostPort(v string)`

SetHostPort sets HostPort field to given value.

### HasHostPort

`func (o *DSProducerDetails) HasHostPort() bool`

HasHostPort returns a boolean if a field has been set.

### GetIsFixedUser

`func (o *DSProducerDetails) GetIsFixedUser() string`

GetIsFixedUser returns the IsFixedUser field if non-nil, zero value otherwise.

### GetIsFixedUserOk

`func (o *DSProducerDetails) GetIsFixedUserOk() (*string, bool)`

GetIsFixedUserOk returns a tuple with the IsFixedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixedUser

`func (o *DSProducerDetails) SetIsFixedUser(v string)`

SetIsFixedUser sets IsFixedUser field to given value.

### HasIsFixedUser

`func (o *DSProducerDetails) HasIsFixedUser() bool`

HasIsFixedUser returns a boolean if a field has been set.

### GetItemTargetsAssoc

`func (o *DSProducerDetails) GetItemTargetsAssoc() []ItemTargetAssociation`

GetItemTargetsAssoc returns the ItemTargetsAssoc field if non-nil, zero value otherwise.

### GetItemTargetsAssocOk

`func (o *DSProducerDetails) GetItemTargetsAssocOk() (*[]ItemTargetAssociation, bool)`

GetItemTargetsAssocOk returns a tuple with the ItemTargetsAssoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTargetsAssoc

`func (o *DSProducerDetails) SetItemTargetsAssoc(v []ItemTargetAssociation)`

SetItemTargetsAssoc sets ItemTargetsAssoc field to given value.

### HasItemTargetsAssoc

`func (o *DSProducerDetails) HasItemTargetsAssoc() bool`

HasItemTargetsAssoc returns a boolean if a field has been set.

### GetK8sBearerToken

`func (o *DSProducerDetails) GetK8sBearerToken() string`

GetK8sBearerToken returns the K8sBearerToken field if non-nil, zero value otherwise.

### GetK8sBearerTokenOk

`func (o *DSProducerDetails) GetK8sBearerTokenOk() (*string, bool)`

GetK8sBearerTokenOk returns a tuple with the K8sBearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sBearerToken

`func (o *DSProducerDetails) SetK8sBearerToken(v string)`

SetK8sBearerToken sets K8sBearerToken field to given value.

### HasK8sBearerToken

`func (o *DSProducerDetails) HasK8sBearerToken() bool`

HasK8sBearerToken returns a boolean if a field has been set.

### GetK8sClusterCaCertificate

`func (o *DSProducerDetails) GetK8sClusterCaCertificate() string`

GetK8sClusterCaCertificate returns the K8sClusterCaCertificate field if non-nil, zero value otherwise.

### GetK8sClusterCaCertificateOk

`func (o *DSProducerDetails) GetK8sClusterCaCertificateOk() (*string, bool)`

GetK8sClusterCaCertificateOk returns a tuple with the K8sClusterCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterCaCertificate

`func (o *DSProducerDetails) SetK8sClusterCaCertificate(v string)`

SetK8sClusterCaCertificate sets K8sClusterCaCertificate field to given value.

### HasK8sClusterCaCertificate

`func (o *DSProducerDetails) HasK8sClusterCaCertificate() bool`

HasK8sClusterCaCertificate returns a boolean if a field has been set.

### GetK8sClusterEndpoint

`func (o *DSProducerDetails) GetK8sClusterEndpoint() string`

GetK8sClusterEndpoint returns the K8sClusterEndpoint field if non-nil, zero value otherwise.

### GetK8sClusterEndpointOk

`func (o *DSProducerDetails) GetK8sClusterEndpointOk() (*string, bool)`

GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterEndpoint

`func (o *DSProducerDetails) SetK8sClusterEndpoint(v string)`

SetK8sClusterEndpoint sets K8sClusterEndpoint field to given value.

### HasK8sClusterEndpoint

`func (o *DSProducerDetails) HasK8sClusterEndpoint() bool`

HasK8sClusterEndpoint returns a boolean if a field has been set.

### GetK8sNamespace

`func (o *DSProducerDetails) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *DSProducerDetails) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *DSProducerDetails) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *DSProducerDetails) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### GetK8sServiceAccount

`func (o *DSProducerDetails) GetK8sServiceAccount() string`

GetK8sServiceAccount returns the K8sServiceAccount field if non-nil, zero value otherwise.

### GetK8sServiceAccountOk

`func (o *DSProducerDetails) GetK8sServiceAccountOk() (*string, bool)`

GetK8sServiceAccountOk returns a tuple with the K8sServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sServiceAccount

`func (o *DSProducerDetails) SetK8sServiceAccount(v string)`

SetK8sServiceAccount sets K8sServiceAccount field to given value.

### HasK8sServiceAccount

`func (o *DSProducerDetails) HasK8sServiceAccount() bool`

HasK8sServiceAccount returns a boolean if a field has been set.

### GetLastAdminRotation

`func (o *DSProducerDetails) GetLastAdminRotation() int64`

GetLastAdminRotation returns the LastAdminRotation field if non-nil, zero value otherwise.

### GetLastAdminRotationOk

`func (o *DSProducerDetails) GetLastAdminRotationOk() (*int64, bool)`

GetLastAdminRotationOk returns a tuple with the LastAdminRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAdminRotation

`func (o *DSProducerDetails) SetLastAdminRotation(v int64)`

SetLastAdminRotation sets LastAdminRotation field to given value.

### HasLastAdminRotation

`func (o *DSProducerDetails) HasLastAdminRotation() bool`

HasLastAdminRotation returns a boolean if a field has been set.

### GetLdapAudience

`func (o *DSProducerDetails) GetLdapAudience() string`

GetLdapAudience returns the LdapAudience field if non-nil, zero value otherwise.

### GetLdapAudienceOk

`func (o *DSProducerDetails) GetLdapAudienceOk() (*string, bool)`

GetLdapAudienceOk returns a tuple with the LdapAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAudience

`func (o *DSProducerDetails) SetLdapAudience(v string)`

SetLdapAudience sets LdapAudience field to given value.

### HasLdapAudience

`func (o *DSProducerDetails) HasLdapAudience() bool`

HasLdapAudience returns a boolean if a field has been set.

### GetLdapBindDn

`func (o *DSProducerDetails) GetLdapBindDn() string`

GetLdapBindDn returns the LdapBindDn field if non-nil, zero value otherwise.

### GetLdapBindDnOk

`func (o *DSProducerDetails) GetLdapBindDnOk() (*string, bool)`

GetLdapBindDnOk returns a tuple with the LdapBindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindDn

`func (o *DSProducerDetails) SetLdapBindDn(v string)`

SetLdapBindDn sets LdapBindDn field to given value.

### HasLdapBindDn

`func (o *DSProducerDetails) HasLdapBindDn() bool`

HasLdapBindDn returns a boolean if a field has been set.

### GetLdapBindPassword

`func (o *DSProducerDetails) GetLdapBindPassword() string`

GetLdapBindPassword returns the LdapBindPassword field if non-nil, zero value otherwise.

### GetLdapBindPasswordOk

`func (o *DSProducerDetails) GetLdapBindPasswordOk() (*string, bool)`

GetLdapBindPasswordOk returns a tuple with the LdapBindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindPassword

`func (o *DSProducerDetails) SetLdapBindPassword(v string)`

SetLdapBindPassword sets LdapBindPassword field to given value.

### HasLdapBindPassword

`func (o *DSProducerDetails) HasLdapBindPassword() bool`

HasLdapBindPassword returns a boolean if a field has been set.

### GetLdapCertificate

`func (o *DSProducerDetails) GetLdapCertificate() string`

GetLdapCertificate returns the LdapCertificate field if non-nil, zero value otherwise.

### GetLdapCertificateOk

`func (o *DSProducerDetails) GetLdapCertificateOk() (*string, bool)`

GetLdapCertificateOk returns a tuple with the LdapCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCertificate

`func (o *DSProducerDetails) SetLdapCertificate(v string)`

SetLdapCertificate sets LdapCertificate field to given value.

### HasLdapCertificate

`func (o *DSProducerDetails) HasLdapCertificate() bool`

HasLdapCertificate returns a boolean if a field has been set.

### GetLdapTokenExpiration

`func (o *DSProducerDetails) GetLdapTokenExpiration() string`

GetLdapTokenExpiration returns the LdapTokenExpiration field if non-nil, zero value otherwise.

### GetLdapTokenExpirationOk

`func (o *DSProducerDetails) GetLdapTokenExpirationOk() (*string, bool)`

GetLdapTokenExpirationOk returns a tuple with the LdapTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapTokenExpiration

`func (o *DSProducerDetails) SetLdapTokenExpiration(v string)`

SetLdapTokenExpiration sets LdapTokenExpiration field to given value.

### HasLdapTokenExpiration

`func (o *DSProducerDetails) HasLdapTokenExpiration() bool`

HasLdapTokenExpiration returns a boolean if a field has been set.

### GetLdapUrl

`func (o *DSProducerDetails) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *DSProducerDetails) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *DSProducerDetails) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *DSProducerDetails) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetLdapUserAttr

`func (o *DSProducerDetails) GetLdapUserAttr() string`

GetLdapUserAttr returns the LdapUserAttr field if non-nil, zero value otherwise.

### GetLdapUserAttrOk

`func (o *DSProducerDetails) GetLdapUserAttrOk() (*string, bool)`

GetLdapUserAttrOk returns a tuple with the LdapUserAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserAttr

`func (o *DSProducerDetails) SetLdapUserAttr(v string)`

SetLdapUserAttr sets LdapUserAttr field to given value.

### HasLdapUserAttr

`func (o *DSProducerDetails) HasLdapUserAttr() bool`

HasLdapUserAttr returns a boolean if a field has been set.

### GetLdapUserDn

`func (o *DSProducerDetails) GetLdapUserDn() string`

GetLdapUserDn returns the LdapUserDn field if non-nil, zero value otherwise.

### GetLdapUserDnOk

`func (o *DSProducerDetails) GetLdapUserDnOk() (*string, bool)`

GetLdapUserDnOk returns a tuple with the LdapUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDn

`func (o *DSProducerDetails) SetLdapUserDn(v string)`

SetLdapUserDn sets LdapUserDn field to given value.

### HasLdapUserDn

`func (o *DSProducerDetails) HasLdapUserDn() bool`

HasLdapUserDn returns a boolean if a field has been set.

### GetMongodbAtlasApiPrivateKey

`func (o *DSProducerDetails) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *DSProducerDetails) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *DSProducerDetails) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *DSProducerDetails) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *DSProducerDetails) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *DSProducerDetails) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *DSProducerDetails) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *DSProducerDetails) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *DSProducerDetails) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *DSProducerDetails) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *DSProducerDetails) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *DSProducerDetails) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbCustomData

`func (o *DSProducerDetails) GetMongodbCustomData() string`

GetMongodbCustomData returns the MongodbCustomData field if non-nil, zero value otherwise.

### GetMongodbCustomDataOk

`func (o *DSProducerDetails) GetMongodbCustomDataOk() (*string, bool)`

GetMongodbCustomDataOk returns a tuple with the MongodbCustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbCustomData

`func (o *DSProducerDetails) SetMongodbCustomData(v string)`

SetMongodbCustomData sets MongodbCustomData field to given value.

### HasMongodbCustomData

`func (o *DSProducerDetails) HasMongodbCustomData() bool`

HasMongodbCustomData returns a boolean if a field has been set.

### GetMongodbDbName

`func (o *DSProducerDetails) GetMongodbDbName() string`

GetMongodbDbName returns the MongodbDbName field if non-nil, zero value otherwise.

### GetMongodbDbNameOk

`func (o *DSProducerDetails) GetMongodbDbNameOk() (*string, bool)`

GetMongodbDbNameOk returns a tuple with the MongodbDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDbName

`func (o *DSProducerDetails) SetMongodbDbName(v string)`

SetMongodbDbName sets MongodbDbName field to given value.

### HasMongodbDbName

`func (o *DSProducerDetails) HasMongodbDbName() bool`

HasMongodbDbName returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *DSProducerDetails) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *DSProducerDetails) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *DSProducerDetails) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *DSProducerDetails) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbHostPort

`func (o *DSProducerDetails) GetMongodbHostPort() string`

GetMongodbHostPort returns the MongodbHostPort field if non-nil, zero value otherwise.

### GetMongodbHostPortOk

`func (o *DSProducerDetails) GetMongodbHostPortOk() (*string, bool)`

GetMongodbHostPortOk returns a tuple with the MongodbHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbHostPort

`func (o *DSProducerDetails) SetMongodbHostPort(v string)`

SetMongodbHostPort sets MongodbHostPort field to given value.

### HasMongodbHostPort

`func (o *DSProducerDetails) HasMongodbHostPort() bool`

HasMongodbHostPort returns a boolean if a field has been set.

### GetMongodbIsAtlas

`func (o *DSProducerDetails) GetMongodbIsAtlas() bool`

GetMongodbIsAtlas returns the MongodbIsAtlas field if non-nil, zero value otherwise.

### GetMongodbIsAtlasOk

`func (o *DSProducerDetails) GetMongodbIsAtlasOk() (*bool, bool)`

GetMongodbIsAtlasOk returns a tuple with the MongodbIsAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbIsAtlas

`func (o *DSProducerDetails) SetMongodbIsAtlas(v bool)`

SetMongodbIsAtlas sets MongodbIsAtlas field to given value.

### HasMongodbIsAtlas

`func (o *DSProducerDetails) HasMongodbIsAtlas() bool`

HasMongodbIsAtlas returns a boolean if a field has been set.

### GetMongodbPassword

`func (o *DSProducerDetails) GetMongodbPassword() string`

GetMongodbPassword returns the MongodbPassword field if non-nil, zero value otherwise.

### GetMongodbPasswordOk

`func (o *DSProducerDetails) GetMongodbPasswordOk() (*string, bool)`

GetMongodbPasswordOk returns a tuple with the MongodbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbPassword

`func (o *DSProducerDetails) SetMongodbPassword(v string)`

SetMongodbPassword sets MongodbPassword field to given value.

### HasMongodbPassword

`func (o *DSProducerDetails) HasMongodbPassword() bool`

HasMongodbPassword returns a boolean if a field has been set.

### GetMongodbRoles

`func (o *DSProducerDetails) GetMongodbRoles() string`

GetMongodbRoles returns the MongodbRoles field if non-nil, zero value otherwise.

### GetMongodbRolesOk

`func (o *DSProducerDetails) GetMongodbRolesOk() (*string, bool)`

GetMongodbRolesOk returns a tuple with the MongodbRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbRoles

`func (o *DSProducerDetails) SetMongodbRoles(v string)`

SetMongodbRoles sets MongodbRoles field to given value.

### HasMongodbRoles

`func (o *DSProducerDetails) HasMongodbRoles() bool`

HasMongodbRoles returns a boolean if a field has been set.

### GetMongodbUriConnection

`func (o *DSProducerDetails) GetMongodbUriConnection() string`

GetMongodbUriConnection returns the MongodbUriConnection field if non-nil, zero value otherwise.

### GetMongodbUriConnectionOk

`func (o *DSProducerDetails) GetMongodbUriConnectionOk() (*string, bool)`

GetMongodbUriConnectionOk returns a tuple with the MongodbUriConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriConnection

`func (o *DSProducerDetails) SetMongodbUriConnection(v string)`

SetMongodbUriConnection sets MongodbUriConnection field to given value.

### HasMongodbUriConnection

`func (o *DSProducerDetails) HasMongodbUriConnection() bool`

HasMongodbUriConnection returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *DSProducerDetails) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *DSProducerDetails) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *DSProducerDetails) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *DSProducerDetails) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetMongodbUsername

`func (o *DSProducerDetails) GetMongodbUsername() string`

GetMongodbUsername returns the MongodbUsername field if non-nil, zero value otherwise.

### GetMongodbUsernameOk

`func (o *DSProducerDetails) GetMongodbUsernameOk() (*string, bool)`

GetMongodbUsernameOk returns a tuple with the MongodbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUsername

`func (o *DSProducerDetails) SetMongodbUsername(v string)`

SetMongodbUsername sets MongodbUsername field to given value.

### HasMongodbUsername

`func (o *DSProducerDetails) HasMongodbUsername() bool`

HasMongodbUsername returns a boolean if a field has been set.

### GetMssqlCreationStatements

`func (o *DSProducerDetails) GetMssqlCreationStatements() string`

GetMssqlCreationStatements returns the MssqlCreationStatements field if non-nil, zero value otherwise.

### GetMssqlCreationStatementsOk

`func (o *DSProducerDetails) GetMssqlCreationStatementsOk() (*string, bool)`

GetMssqlCreationStatementsOk returns a tuple with the MssqlCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlCreationStatements

`func (o *DSProducerDetails) SetMssqlCreationStatements(v string)`

SetMssqlCreationStatements sets MssqlCreationStatements field to given value.

### HasMssqlCreationStatements

`func (o *DSProducerDetails) HasMssqlCreationStatements() bool`

HasMssqlCreationStatements returns a boolean if a field has been set.

### GetMssqlRevocationStatements

`func (o *DSProducerDetails) GetMssqlRevocationStatements() string`

GetMssqlRevocationStatements returns the MssqlRevocationStatements field if non-nil, zero value otherwise.

### GetMssqlRevocationStatementsOk

`func (o *DSProducerDetails) GetMssqlRevocationStatementsOk() (*string, bool)`

GetMssqlRevocationStatementsOk returns a tuple with the MssqlRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlRevocationStatements

`func (o *DSProducerDetails) SetMssqlRevocationStatements(v string)`

SetMssqlRevocationStatements sets MssqlRevocationStatements field to given value.

### HasMssqlRevocationStatements

`func (o *DSProducerDetails) HasMssqlRevocationStatements() bool`

HasMssqlRevocationStatements returns a boolean if a field has been set.

### GetMysqlCreationStatements

`func (o *DSProducerDetails) GetMysqlCreationStatements() string`

GetMysqlCreationStatements returns the MysqlCreationStatements field if non-nil, zero value otherwise.

### GetMysqlCreationStatementsOk

`func (o *DSProducerDetails) GetMysqlCreationStatementsOk() (*string, bool)`

GetMysqlCreationStatementsOk returns a tuple with the MysqlCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlCreationStatements

`func (o *DSProducerDetails) SetMysqlCreationStatements(v string)`

SetMysqlCreationStatements sets MysqlCreationStatements field to given value.

### HasMysqlCreationStatements

`func (o *DSProducerDetails) HasMysqlCreationStatements() bool`

HasMysqlCreationStatements returns a boolean if a field has been set.

### GetOracleCreationStatements

`func (o *DSProducerDetails) GetOracleCreationStatements() string`

GetOracleCreationStatements returns the OracleCreationStatements field if non-nil, zero value otherwise.

### GetOracleCreationStatementsOk

`func (o *DSProducerDetails) GetOracleCreationStatementsOk() (*string, bool)`

GetOracleCreationStatementsOk returns a tuple with the OracleCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleCreationStatements

`func (o *DSProducerDetails) SetOracleCreationStatements(v string)`

SetOracleCreationStatements sets OracleCreationStatements field to given value.

### HasOracleCreationStatements

`func (o *DSProducerDetails) HasOracleCreationStatements() bool`

HasOracleCreationStatements returns a boolean if a field has been set.

### GetPassword

`func (o *DSProducerDetails) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DSProducerDetails) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DSProducerDetails) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DSProducerDetails) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordLength

`func (o *DSProducerDetails) GetPasswordLength() int64`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DSProducerDetails) GetPasswordLengthOk() (*int64, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DSProducerDetails) SetPasswordLength(v int64)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DSProducerDetails) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *DSProducerDetails) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *DSProducerDetails) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *DSProducerDetails) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *DSProducerDetails) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetPayload

`func (o *DSProducerDetails) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DSProducerDetails) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DSProducerDetails) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DSProducerDetails) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPostgresCreationStatements

`func (o *DSProducerDetails) GetPostgresCreationStatements() string`

GetPostgresCreationStatements returns the PostgresCreationStatements field if non-nil, zero value otherwise.

### GetPostgresCreationStatementsOk

`func (o *DSProducerDetails) GetPostgresCreationStatementsOk() (*string, bool)`

GetPostgresCreationStatementsOk returns a tuple with the PostgresCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresCreationStatements

`func (o *DSProducerDetails) SetPostgresCreationStatements(v string)`

SetPostgresCreationStatements sets PostgresCreationStatements field to given value.

### HasPostgresCreationStatements

`func (o *DSProducerDetails) HasPostgresCreationStatements() bool`

HasPostgresCreationStatements returns a boolean if a field has been set.

### GetPostgresRevocationStatements

`func (o *DSProducerDetails) GetPostgresRevocationStatements() string`

GetPostgresRevocationStatements returns the PostgresRevocationStatements field if non-nil, zero value otherwise.

### GetPostgresRevocationStatementsOk

`func (o *DSProducerDetails) GetPostgresRevocationStatementsOk() (*string, bool)`

GetPostgresRevocationStatementsOk returns a tuple with the PostgresRevocationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresRevocationStatements

`func (o *DSProducerDetails) SetPostgresRevocationStatements(v string)`

SetPostgresRevocationStatements sets PostgresRevocationStatements field to given value.

### HasPostgresRevocationStatements

`func (o *DSProducerDetails) HasPostgresRevocationStatements() bool`

HasPostgresRevocationStatements returns a boolean if a field has been set.

### GetRabbitmqServerPassword

`func (o *DSProducerDetails) GetRabbitmqServerPassword() string`

GetRabbitmqServerPassword returns the RabbitmqServerPassword field if non-nil, zero value otherwise.

### GetRabbitmqServerPasswordOk

`func (o *DSProducerDetails) GetRabbitmqServerPasswordOk() (*string, bool)`

GetRabbitmqServerPasswordOk returns a tuple with the RabbitmqServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerPassword

`func (o *DSProducerDetails) SetRabbitmqServerPassword(v string)`

SetRabbitmqServerPassword sets RabbitmqServerPassword field to given value.

### HasRabbitmqServerPassword

`func (o *DSProducerDetails) HasRabbitmqServerPassword() bool`

HasRabbitmqServerPassword returns a boolean if a field has been set.

### GetRabbitmqServerUri

`func (o *DSProducerDetails) GetRabbitmqServerUri() string`

GetRabbitmqServerUri returns the RabbitmqServerUri field if non-nil, zero value otherwise.

### GetRabbitmqServerUriOk

`func (o *DSProducerDetails) GetRabbitmqServerUriOk() (*string, bool)`

GetRabbitmqServerUriOk returns a tuple with the RabbitmqServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUri

`func (o *DSProducerDetails) SetRabbitmqServerUri(v string)`

SetRabbitmqServerUri sets RabbitmqServerUri field to given value.

### HasRabbitmqServerUri

`func (o *DSProducerDetails) HasRabbitmqServerUri() bool`

HasRabbitmqServerUri returns a boolean if a field has been set.

### GetRabbitmqServerUser

`func (o *DSProducerDetails) GetRabbitmqServerUser() string`

GetRabbitmqServerUser returns the RabbitmqServerUser field if non-nil, zero value otherwise.

### GetRabbitmqServerUserOk

`func (o *DSProducerDetails) GetRabbitmqServerUserOk() (*string, bool)`

GetRabbitmqServerUserOk returns a tuple with the RabbitmqServerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqServerUser

`func (o *DSProducerDetails) SetRabbitmqServerUser(v string)`

SetRabbitmqServerUser sets RabbitmqServerUser field to given value.

### HasRabbitmqServerUser

`func (o *DSProducerDetails) HasRabbitmqServerUser() bool`

HasRabbitmqServerUser returns a boolean if a field has been set.

### GetRabbitmqUserConfPermission

`func (o *DSProducerDetails) GetRabbitmqUserConfPermission() string`

GetRabbitmqUserConfPermission returns the RabbitmqUserConfPermission field if non-nil, zero value otherwise.

### GetRabbitmqUserConfPermissionOk

`func (o *DSProducerDetails) GetRabbitmqUserConfPermissionOk() (*string, bool)`

GetRabbitmqUserConfPermissionOk returns a tuple with the RabbitmqUserConfPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserConfPermission

`func (o *DSProducerDetails) SetRabbitmqUserConfPermission(v string)`

SetRabbitmqUserConfPermission sets RabbitmqUserConfPermission field to given value.

### HasRabbitmqUserConfPermission

`func (o *DSProducerDetails) HasRabbitmqUserConfPermission() bool`

HasRabbitmqUserConfPermission returns a boolean if a field has been set.

### GetRabbitmqUserReadPermission

`func (o *DSProducerDetails) GetRabbitmqUserReadPermission() string`

GetRabbitmqUserReadPermission returns the RabbitmqUserReadPermission field if non-nil, zero value otherwise.

### GetRabbitmqUserReadPermissionOk

`func (o *DSProducerDetails) GetRabbitmqUserReadPermissionOk() (*string, bool)`

GetRabbitmqUserReadPermissionOk returns a tuple with the RabbitmqUserReadPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserReadPermission

`func (o *DSProducerDetails) SetRabbitmqUserReadPermission(v string)`

SetRabbitmqUserReadPermission sets RabbitmqUserReadPermission field to given value.

### HasRabbitmqUserReadPermission

`func (o *DSProducerDetails) HasRabbitmqUserReadPermission() bool`

HasRabbitmqUserReadPermission returns a boolean if a field has been set.

### GetRabbitmqUserTags

`func (o *DSProducerDetails) GetRabbitmqUserTags() string`

GetRabbitmqUserTags returns the RabbitmqUserTags field if non-nil, zero value otherwise.

### GetRabbitmqUserTagsOk

`func (o *DSProducerDetails) GetRabbitmqUserTagsOk() (*string, bool)`

GetRabbitmqUserTagsOk returns a tuple with the RabbitmqUserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserTags

`func (o *DSProducerDetails) SetRabbitmqUserTags(v string)`

SetRabbitmqUserTags sets RabbitmqUserTags field to given value.

### HasRabbitmqUserTags

`func (o *DSProducerDetails) HasRabbitmqUserTags() bool`

HasRabbitmqUserTags returns a boolean if a field has been set.

### GetRabbitmqUserVhost

`func (o *DSProducerDetails) GetRabbitmqUserVhost() string`

GetRabbitmqUserVhost returns the RabbitmqUserVhost field if non-nil, zero value otherwise.

### GetRabbitmqUserVhostOk

`func (o *DSProducerDetails) GetRabbitmqUserVhostOk() (*string, bool)`

GetRabbitmqUserVhostOk returns a tuple with the RabbitmqUserVhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserVhost

`func (o *DSProducerDetails) SetRabbitmqUserVhost(v string)`

SetRabbitmqUserVhost sets RabbitmqUserVhost field to given value.

### HasRabbitmqUserVhost

`func (o *DSProducerDetails) HasRabbitmqUserVhost() bool`

HasRabbitmqUserVhost returns a boolean if a field has been set.

### GetRabbitmqUserWritePermission

`func (o *DSProducerDetails) GetRabbitmqUserWritePermission() string`

GetRabbitmqUserWritePermission returns the RabbitmqUserWritePermission field if non-nil, zero value otherwise.

### GetRabbitmqUserWritePermissionOk

`func (o *DSProducerDetails) GetRabbitmqUserWritePermissionOk() (*string, bool)`

GetRabbitmqUserWritePermissionOk returns a tuple with the RabbitmqUserWritePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmqUserWritePermission

`func (o *DSProducerDetails) SetRabbitmqUserWritePermission(v string)`

SetRabbitmqUserWritePermission sets RabbitmqUserWritePermission field to given value.

### HasRabbitmqUserWritePermission

`func (o *DSProducerDetails) HasRabbitmqUserWritePermission() bool`

HasRabbitmqUserWritePermission returns a boolean if a field has been set.

### GetRedshiftCreationStatements

`func (o *DSProducerDetails) GetRedshiftCreationStatements() string`

GetRedshiftCreationStatements returns the RedshiftCreationStatements field if non-nil, zero value otherwise.

### GetRedshiftCreationStatementsOk

`func (o *DSProducerDetails) GetRedshiftCreationStatementsOk() (*string, bool)`

GetRedshiftCreationStatementsOk returns a tuple with the RedshiftCreationStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftCreationStatements

`func (o *DSProducerDetails) SetRedshiftCreationStatements(v string)`

SetRedshiftCreationStatements sets RedshiftCreationStatements field to given value.

### HasRedshiftCreationStatements

`func (o *DSProducerDetails) HasRedshiftCreationStatements() bool`

HasRedshiftCreationStatements returns a boolean if a field has been set.

### GetRevokeSyncUrl

`func (o *DSProducerDetails) GetRevokeSyncUrl() string`

GetRevokeSyncUrl returns the RevokeSyncUrl field if non-nil, zero value otherwise.

### GetRevokeSyncUrlOk

`func (o *DSProducerDetails) GetRevokeSyncUrlOk() (*string, bool)`

GetRevokeSyncUrlOk returns a tuple with the RevokeSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSyncUrl

`func (o *DSProducerDetails) SetRevokeSyncUrl(v string)`

SetRevokeSyncUrl sets RevokeSyncUrl field to given value.

### HasRevokeSyncUrl

`func (o *DSProducerDetails) HasRevokeSyncUrl() bool`

HasRevokeSyncUrl returns a boolean if a field has been set.

### GetRotateSyncUrl

`func (o *DSProducerDetails) GetRotateSyncUrl() string`

GetRotateSyncUrl returns the RotateSyncUrl field if non-nil, zero value otherwise.

### GetRotateSyncUrlOk

`func (o *DSProducerDetails) GetRotateSyncUrlOk() (*string, bool)`

GetRotateSyncUrlOk returns a tuple with the RotateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateSyncUrl

`func (o *DSProducerDetails) SetRotateSyncUrl(v string)`

SetRotateSyncUrl sets RotateSyncUrl field to given value.

### HasRotateSyncUrl

`func (o *DSProducerDetails) HasRotateSyncUrl() bool`

HasRotateSyncUrl returns a boolean if a field has been set.

### GetScopes

`func (o *DSProducerDetails) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DSProducerDetails) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DSProducerDetails) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DSProducerDetails) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSecureRemoteAccessDetails

`func (o *DSProducerDetails) GetSecureRemoteAccessDetails() SecureRemoteAccess`

GetSecureRemoteAccessDetails returns the SecureRemoteAccessDetails field if non-nil, zero value otherwise.

### GetSecureRemoteAccessDetailsOk

`func (o *DSProducerDetails) GetSecureRemoteAccessDetailsOk() (*SecureRemoteAccess, bool)`

GetSecureRemoteAccessDetailsOk returns a tuple with the SecureRemoteAccessDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureRemoteAccessDetails

`func (o *DSProducerDetails) SetSecureRemoteAccessDetails(v SecureRemoteAccess)`

SetSecureRemoteAccessDetails sets SecureRemoteAccessDetails field to given value.

### HasSecureRemoteAccessDetails

`func (o *DSProducerDetails) HasSecureRemoteAccessDetails() bool`

HasSecureRemoteAccessDetails returns a boolean if a field has been set.

### GetSessionExtensionWarnIntervalMin

`func (o *DSProducerDetails) GetSessionExtensionWarnIntervalMin() int64`

GetSessionExtensionWarnIntervalMin returns the SessionExtensionWarnIntervalMin field if non-nil, zero value otherwise.

### GetSessionExtensionWarnIntervalMinOk

`func (o *DSProducerDetails) GetSessionExtensionWarnIntervalMinOk() (*int64, bool)`

GetSessionExtensionWarnIntervalMinOk returns a tuple with the SessionExtensionWarnIntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExtensionWarnIntervalMin

`func (o *DSProducerDetails) SetSessionExtensionWarnIntervalMin(v int64)`

SetSessionExtensionWarnIntervalMin sets SessionExtensionWarnIntervalMin field to given value.

### HasSessionExtensionWarnIntervalMin

`func (o *DSProducerDetails) HasSessionExtensionWarnIntervalMin() bool`

HasSessionExtensionWarnIntervalMin returns a boolean if a field has been set.

### GetSfAccount

`func (o *DSProducerDetails) GetSfAccount() string`

GetSfAccount returns the SfAccount field if non-nil, zero value otherwise.

### GetSfAccountOk

`func (o *DSProducerDetails) GetSfAccountOk() (*string, bool)`

GetSfAccountOk returns a tuple with the SfAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfAccount

`func (o *DSProducerDetails) SetSfAccount(v string)`

SetSfAccount sets SfAccount field to given value.

### HasSfAccount

`func (o *DSProducerDetails) HasSfAccount() bool`

HasSfAccount returns a boolean if a field has been set.

### GetSfUserRole

`func (o *DSProducerDetails) GetSfUserRole() string`

GetSfUserRole returns the SfUserRole field if non-nil, zero value otherwise.

### GetSfUserRoleOk

`func (o *DSProducerDetails) GetSfUserRoleOk() (*string, bool)`

GetSfUserRoleOk returns a tuple with the SfUserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfUserRole

`func (o *DSProducerDetails) SetSfUserRole(v string)`

SetSfUserRole sets SfUserRole field to given value.

### HasSfUserRole

`func (o *DSProducerDetails) HasSfUserRole() bool`

HasSfUserRole returns a boolean if a field has been set.

### GetSfWarehouseName

`func (o *DSProducerDetails) GetSfWarehouseName() string`

GetSfWarehouseName returns the SfWarehouseName field if non-nil, zero value otherwise.

### GetSfWarehouseNameOk

`func (o *DSProducerDetails) GetSfWarehouseNameOk() (*string, bool)`

GetSfWarehouseNameOk returns a tuple with the SfWarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfWarehouseName

`func (o *DSProducerDetails) SetSfWarehouseName(v string)`

SetSfWarehouseName sets SfWarehouseName field to given value.

### HasSfWarehouseName

`func (o *DSProducerDetails) HasSfWarehouseName() bool`

HasSfWarehouseName returns a boolean if a field has been set.

### GetShouldStop

`func (o *DSProducerDetails) GetShouldStop() string`

GetShouldStop returns the ShouldStop field if non-nil, zero value otherwise.

### GetShouldStopOk

`func (o *DSProducerDetails) GetShouldStopOk() (*string, bool)`

GetShouldStopOk returns a tuple with the ShouldStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldStop

`func (o *DSProducerDetails) SetShouldStop(v string)`

SetShouldStop sets ShouldStop field to given value.

### HasShouldStop

`func (o *DSProducerDetails) HasShouldStop() bool`

HasShouldStop returns a boolean if a field has been set.

### GetSslConnectionCertificate

`func (o *DSProducerDetails) GetSslConnectionCertificate() string`

GetSslConnectionCertificate returns the SslConnectionCertificate field if non-nil, zero value otherwise.

### GetSslConnectionCertificateOk

`func (o *DSProducerDetails) GetSslConnectionCertificateOk() (*string, bool)`

GetSslConnectionCertificateOk returns a tuple with the SslConnectionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnectionCertificate

`func (o *DSProducerDetails) SetSslConnectionCertificate(v string)`

SetSslConnectionCertificate sets SslConnectionCertificate field to given value.

### HasSslConnectionCertificate

`func (o *DSProducerDetails) HasSslConnectionCertificate() bool`

HasSslConnectionCertificate returns a boolean if a field has been set.

### GetSslConnectionMode

`func (o *DSProducerDetails) GetSslConnectionMode() bool`

GetSslConnectionMode returns the SslConnectionMode field if non-nil, zero value otherwise.

### GetSslConnectionModeOk

`func (o *DSProducerDetails) GetSslConnectionModeOk() (*bool, bool)`

GetSslConnectionModeOk returns a tuple with the SslConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnectionMode

`func (o *DSProducerDetails) SetSslConnectionMode(v bool)`

SetSslConnectionMode sets SslConnectionMode field to given value.

### HasSslConnectionMode

`func (o *DSProducerDetails) HasSslConnectionMode() bool`

HasSslConnectionMode returns a boolean if a field has been set.

### GetTags

`func (o *DSProducerDetails) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DSProducerDetails) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DSProducerDetails) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DSProducerDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *DSProducerDetails) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *DSProducerDetails) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *DSProducerDetails) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *DSProducerDetails) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *DSProducerDetails) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *DSProducerDetails) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *DSProducerDetails) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *DSProducerDetails) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.

### GetUserName

`func (o *DSProducerDetails) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DSProducerDetails) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DSProducerDetails) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DSProducerDetails) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *DSProducerDetails) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *DSProducerDetails) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *DSProducerDetails) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *DSProducerDetails) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetUserTtl

`func (o *DSProducerDetails) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DSProducerDetails) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DSProducerDetails) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DSProducerDetails) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsernameLength

`func (o *DSProducerDetails) GetUsernameLength() int64`

GetUsernameLength returns the UsernameLength field if non-nil, zero value otherwise.

### GetUsernameLengthOk

`func (o *DSProducerDetails) GetUsernameLengthOk() (*int64, bool)`

GetUsernameLengthOk returns a tuple with the UsernameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameLength

`func (o *DSProducerDetails) SetUsernameLength(v int64)`

SetUsernameLength sets UsernameLength field to given value.

### HasUsernameLength

`func (o *DSProducerDetails) HasUsernameLength() bool`

HasUsernameLength returns a boolean if a field has been set.

### GetUsernamePolicy

`func (o *DSProducerDetails) GetUsernamePolicy() string`

GetUsernamePolicy returns the UsernamePolicy field if non-nil, zero value otherwise.

### GetUsernamePolicyOk

`func (o *DSProducerDetails) GetUsernamePolicyOk() (*string, bool)`

GetUsernamePolicyOk returns a tuple with the UsernamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernamePolicy

`func (o *DSProducerDetails) SetUsernamePolicy(v string)`

SetUsernamePolicy sets UsernamePolicy field to given value.

### HasUsernamePolicy

`func (o *DSProducerDetails) HasUsernamePolicy() bool`

HasUsernamePolicy returns a boolean if a field has been set.

### GetVenafiAllowSubdomains

`func (o *DSProducerDetails) GetVenafiAllowSubdomains() bool`

GetVenafiAllowSubdomains returns the VenafiAllowSubdomains field if non-nil, zero value otherwise.

### GetVenafiAllowSubdomainsOk

`func (o *DSProducerDetails) GetVenafiAllowSubdomainsOk() (*bool, bool)`

GetVenafiAllowSubdomainsOk returns a tuple with the VenafiAllowSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiAllowSubdomains

`func (o *DSProducerDetails) SetVenafiAllowSubdomains(v bool)`

SetVenafiAllowSubdomains sets VenafiAllowSubdomains field to given value.

### HasVenafiAllowSubdomains

`func (o *DSProducerDetails) HasVenafiAllowSubdomains() bool`

HasVenafiAllowSubdomains returns a boolean if a field has been set.

### GetVenafiAllowedDomains

`func (o *DSProducerDetails) GetVenafiAllowedDomains() []string`

GetVenafiAllowedDomains returns the VenafiAllowedDomains field if non-nil, zero value otherwise.

### GetVenafiAllowedDomainsOk

`func (o *DSProducerDetails) GetVenafiAllowedDomainsOk() (*[]string, bool)`

GetVenafiAllowedDomainsOk returns a tuple with the VenafiAllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiAllowedDomains

`func (o *DSProducerDetails) SetVenafiAllowedDomains(v []string)`

SetVenafiAllowedDomains sets VenafiAllowedDomains field to given value.

### HasVenafiAllowedDomains

`func (o *DSProducerDetails) HasVenafiAllowedDomains() bool`

HasVenafiAllowedDomains returns a boolean if a field has been set.

### GetVenafiApiKey

`func (o *DSProducerDetails) GetVenafiApiKey() string`

GetVenafiApiKey returns the VenafiApiKey field if non-nil, zero value otherwise.

### GetVenafiApiKeyOk

`func (o *DSProducerDetails) GetVenafiApiKeyOk() (*string, bool)`

GetVenafiApiKeyOk returns a tuple with the VenafiApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiApiKey

`func (o *DSProducerDetails) SetVenafiApiKey(v string)`

SetVenafiApiKey sets VenafiApiKey field to given value.

### HasVenafiApiKey

`func (o *DSProducerDetails) HasVenafiApiKey() bool`

HasVenafiApiKey returns a boolean if a field has been set.

### GetVenafiAutoGeneratedFolder

`func (o *DSProducerDetails) GetVenafiAutoGeneratedFolder() string`

GetVenafiAutoGeneratedFolder returns the VenafiAutoGeneratedFolder field if non-nil, zero value otherwise.

### GetVenafiAutoGeneratedFolderOk

`func (o *DSProducerDetails) GetVenafiAutoGeneratedFolderOk() (*string, bool)`

GetVenafiAutoGeneratedFolderOk returns a tuple with the VenafiAutoGeneratedFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiAutoGeneratedFolder

`func (o *DSProducerDetails) SetVenafiAutoGeneratedFolder(v string)`

SetVenafiAutoGeneratedFolder sets VenafiAutoGeneratedFolder field to given value.

### HasVenafiAutoGeneratedFolder

`func (o *DSProducerDetails) HasVenafiAutoGeneratedFolder() bool`

HasVenafiAutoGeneratedFolder returns a boolean if a field has been set.

### GetVenafiBaseUrl

`func (o *DSProducerDetails) GetVenafiBaseUrl() string`

GetVenafiBaseUrl returns the VenafiBaseUrl field if non-nil, zero value otherwise.

### GetVenafiBaseUrlOk

`func (o *DSProducerDetails) GetVenafiBaseUrlOk() (*string, bool)`

GetVenafiBaseUrlOk returns a tuple with the VenafiBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiBaseUrl

`func (o *DSProducerDetails) SetVenafiBaseUrl(v string)`

SetVenafiBaseUrl sets VenafiBaseUrl field to given value.

### HasVenafiBaseUrl

`func (o *DSProducerDetails) HasVenafiBaseUrl() bool`

HasVenafiBaseUrl returns a boolean if a field has been set.

### GetVenafiRootFirstInChain

`func (o *DSProducerDetails) GetVenafiRootFirstInChain() bool`

GetVenafiRootFirstInChain returns the VenafiRootFirstInChain field if non-nil, zero value otherwise.

### GetVenafiRootFirstInChainOk

`func (o *DSProducerDetails) GetVenafiRootFirstInChainOk() (*bool, bool)`

GetVenafiRootFirstInChainOk returns a tuple with the VenafiRootFirstInChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiRootFirstInChain

`func (o *DSProducerDetails) SetVenafiRootFirstInChain(v bool)`

SetVenafiRootFirstInChain sets VenafiRootFirstInChain field to given value.

### HasVenafiRootFirstInChain

`func (o *DSProducerDetails) HasVenafiRootFirstInChain() bool`

HasVenafiRootFirstInChain returns a boolean if a field has been set.

### GetVenafiSignUsingAkeylessPki

`func (o *DSProducerDetails) GetVenafiSignUsingAkeylessPki() bool`

GetVenafiSignUsingAkeylessPki returns the VenafiSignUsingAkeylessPki field if non-nil, zero value otherwise.

### GetVenafiSignUsingAkeylessPkiOk

`func (o *DSProducerDetails) GetVenafiSignUsingAkeylessPkiOk() (*bool, bool)`

GetVenafiSignUsingAkeylessPkiOk returns a tuple with the VenafiSignUsingAkeylessPki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiSignUsingAkeylessPki

`func (o *DSProducerDetails) SetVenafiSignUsingAkeylessPki(v bool)`

SetVenafiSignUsingAkeylessPki sets VenafiSignUsingAkeylessPki field to given value.

### HasVenafiSignUsingAkeylessPki

`func (o *DSProducerDetails) HasVenafiSignUsingAkeylessPki() bool`

HasVenafiSignUsingAkeylessPki returns a boolean if a field has been set.

### GetVenafiSignerKeyName

`func (o *DSProducerDetails) GetVenafiSignerKeyName() string`

GetVenafiSignerKeyName returns the VenafiSignerKeyName field if non-nil, zero value otherwise.

### GetVenafiSignerKeyNameOk

`func (o *DSProducerDetails) GetVenafiSignerKeyNameOk() (*string, bool)`

GetVenafiSignerKeyNameOk returns a tuple with the VenafiSignerKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiSignerKeyName

`func (o *DSProducerDetails) SetVenafiSignerKeyName(v string)`

SetVenafiSignerKeyName sets VenafiSignerKeyName field to given value.

### HasVenafiSignerKeyName

`func (o *DSProducerDetails) HasVenafiSignerKeyName() bool`

HasVenafiSignerKeyName returns a boolean if a field has been set.

### GetVenafiStorePrivateKey

`func (o *DSProducerDetails) GetVenafiStorePrivateKey() bool`

GetVenafiStorePrivateKey returns the VenafiStorePrivateKey field if non-nil, zero value otherwise.

### GetVenafiStorePrivateKeyOk

`func (o *DSProducerDetails) GetVenafiStorePrivateKeyOk() (*bool, bool)`

GetVenafiStorePrivateKeyOk returns a tuple with the VenafiStorePrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiStorePrivateKey

`func (o *DSProducerDetails) SetVenafiStorePrivateKey(v bool)`

SetVenafiStorePrivateKey sets VenafiStorePrivateKey field to given value.

### HasVenafiStorePrivateKey

`func (o *DSProducerDetails) HasVenafiStorePrivateKey() bool`

HasVenafiStorePrivateKey returns a boolean if a field has been set.

### GetVenafiTppPassword

`func (o *DSProducerDetails) GetVenafiTppPassword() string`

GetVenafiTppPassword returns the VenafiTppPassword field if non-nil, zero value otherwise.

### GetVenafiTppPasswordOk

`func (o *DSProducerDetails) GetVenafiTppPasswordOk() (*string, bool)`

GetVenafiTppPasswordOk returns a tuple with the VenafiTppPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiTppPassword

`func (o *DSProducerDetails) SetVenafiTppPassword(v string)`

SetVenafiTppPassword sets VenafiTppPassword field to given value.

### HasVenafiTppPassword

`func (o *DSProducerDetails) HasVenafiTppPassword() bool`

HasVenafiTppPassword returns a boolean if a field has been set.

### GetVenafiTppUsername

`func (o *DSProducerDetails) GetVenafiTppUsername() string`

GetVenafiTppUsername returns the VenafiTppUsername field if non-nil, zero value otherwise.

### GetVenafiTppUsernameOk

`func (o *DSProducerDetails) GetVenafiTppUsernameOk() (*string, bool)`

GetVenafiTppUsernameOk returns a tuple with the VenafiTppUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiTppUsername

`func (o *DSProducerDetails) SetVenafiTppUsername(v string)`

SetVenafiTppUsername sets VenafiTppUsername field to given value.

### HasVenafiTppUsername

`func (o *DSProducerDetails) HasVenafiTppUsername() bool`

HasVenafiTppUsername returns a boolean if a field has been set.

### GetVenafiUseTpp

`func (o *DSProducerDetails) GetVenafiUseTpp() bool`

GetVenafiUseTpp returns the VenafiUseTpp field if non-nil, zero value otherwise.

### GetVenafiUseTppOk

`func (o *DSProducerDetails) GetVenafiUseTppOk() (*bool, bool)`

GetVenafiUseTppOk returns a tuple with the VenafiUseTpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiUseTpp

`func (o *DSProducerDetails) SetVenafiUseTpp(v bool)`

SetVenafiUseTpp sets VenafiUseTpp field to given value.

### HasVenafiUseTpp

`func (o *DSProducerDetails) HasVenafiUseTpp() bool`

HasVenafiUseTpp returns a boolean if a field has been set.

### GetVenafiZone

`func (o *DSProducerDetails) GetVenafiZone() string`

GetVenafiZone returns the VenafiZone field if non-nil, zero value otherwise.

### GetVenafiZoneOk

`func (o *DSProducerDetails) GetVenafiZoneOk() (*string, bool)`

GetVenafiZoneOk returns a tuple with the VenafiZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiZone

`func (o *DSProducerDetails) SetVenafiZone(v string)`

SetVenafiZone sets VenafiZone field to given value.

### HasVenafiZone

`func (o *DSProducerDetails) HasVenafiZone() bool`

HasVenafiZone returns a boolean if a field has been set.

### GetWarnBeforeUserExpirationMin

`func (o *DSProducerDetails) GetWarnBeforeUserExpirationMin() int64`

GetWarnBeforeUserExpirationMin returns the WarnBeforeUserExpirationMin field if non-nil, zero value otherwise.

### GetWarnBeforeUserExpirationMinOk

`func (o *DSProducerDetails) GetWarnBeforeUserExpirationMinOk() (*int64, bool)`

GetWarnBeforeUserExpirationMinOk returns a tuple with the WarnBeforeUserExpirationMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnBeforeUserExpirationMin

`func (o *DSProducerDetails) SetWarnBeforeUserExpirationMin(v int64)`

SetWarnBeforeUserExpirationMin sets WarnBeforeUserExpirationMin field to given value.

### HasWarnBeforeUserExpirationMin

`func (o *DSProducerDetails) HasWarnBeforeUserExpirationMin() bool`

HasWarnBeforeUserExpirationMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


