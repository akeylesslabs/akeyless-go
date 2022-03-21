# SecureRemoteAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AllowPortForwarding** | Pointer to **bool** |  | [optional] 
**AllowProvidingExternalUsername** | Pointer to **bool** |  | [optional] 
**BastionApi** | Pointer to **string** |  | [optional] 
**BastionIssuer** | Pointer to **string** |  | [optional] 
**BastionIssuerId** | Pointer to **int64** |  | [optional] 
**BastionSsh** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**DashboardUrl** | Pointer to **string** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **[]string** |  | [optional] 
**IsCli** | Pointer to **bool** |  | [optional] 
**IsWeb** | Pointer to **bool** |  | [optional] 
**Isolated** | Pointer to **bool** |  | [optional] 
**Native** | Pointer to **bool** |  | [optional] 
**RdpUser** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **bool** |  | [optional] 
**SshPrivateKey** | Pointer to **bool** |  | [optional] 
**SshUser** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UseInternalBastion** | Pointer to **bool** |  | [optional] 
**WebProxy** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecureRemoteAccess

`func NewSecureRemoteAccess() *SecureRemoteAccess`

NewSecureRemoteAccess instantiates a new SecureRemoteAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecureRemoteAccessWithDefaults

`func NewSecureRemoteAccessWithDefaults() *SecureRemoteAccess`

NewSecureRemoteAccessWithDefaults instantiates a new SecureRemoteAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SecureRemoteAccess) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SecureRemoteAccess) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SecureRemoteAccess) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SecureRemoteAccess) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAllowPortForwarding

`func (o *SecureRemoteAccess) GetAllowPortForwarding() bool`

GetAllowPortForwarding returns the AllowPortForwarding field if non-nil, zero value otherwise.

### GetAllowPortForwardingOk

`func (o *SecureRemoteAccess) GetAllowPortForwardingOk() (*bool, bool)`

GetAllowPortForwardingOk returns a tuple with the AllowPortForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPortForwarding

`func (o *SecureRemoteAccess) SetAllowPortForwarding(v bool)`

SetAllowPortForwarding sets AllowPortForwarding field to given value.

### HasAllowPortForwarding

`func (o *SecureRemoteAccess) HasAllowPortForwarding() bool`

HasAllowPortForwarding returns a boolean if a field has been set.

### GetAllowProvidingExternalUsername

`func (o *SecureRemoteAccess) GetAllowProvidingExternalUsername() bool`

GetAllowProvidingExternalUsername returns the AllowProvidingExternalUsername field if non-nil, zero value otherwise.

### GetAllowProvidingExternalUsernameOk

`func (o *SecureRemoteAccess) GetAllowProvidingExternalUsernameOk() (*bool, bool)`

GetAllowProvidingExternalUsernameOk returns a tuple with the AllowProvidingExternalUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvidingExternalUsername

`func (o *SecureRemoteAccess) SetAllowProvidingExternalUsername(v bool)`

SetAllowProvidingExternalUsername sets AllowProvidingExternalUsername field to given value.

### HasAllowProvidingExternalUsername

`func (o *SecureRemoteAccess) HasAllowProvidingExternalUsername() bool`

HasAllowProvidingExternalUsername returns a boolean if a field has been set.

### GetBastionApi

`func (o *SecureRemoteAccess) GetBastionApi() string`

GetBastionApi returns the BastionApi field if non-nil, zero value otherwise.

### GetBastionApiOk

`func (o *SecureRemoteAccess) GetBastionApiOk() (*string, bool)`

GetBastionApiOk returns a tuple with the BastionApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionApi

`func (o *SecureRemoteAccess) SetBastionApi(v string)`

SetBastionApi sets BastionApi field to given value.

### HasBastionApi

`func (o *SecureRemoteAccess) HasBastionApi() bool`

HasBastionApi returns a boolean if a field has been set.

### GetBastionIssuer

`func (o *SecureRemoteAccess) GetBastionIssuer() string`

GetBastionIssuer returns the BastionIssuer field if non-nil, zero value otherwise.

### GetBastionIssuerOk

`func (o *SecureRemoteAccess) GetBastionIssuerOk() (*string, bool)`

GetBastionIssuerOk returns a tuple with the BastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionIssuer

`func (o *SecureRemoteAccess) SetBastionIssuer(v string)`

SetBastionIssuer sets BastionIssuer field to given value.

### HasBastionIssuer

`func (o *SecureRemoteAccess) HasBastionIssuer() bool`

HasBastionIssuer returns a boolean if a field has been set.

### GetBastionIssuerId

`func (o *SecureRemoteAccess) GetBastionIssuerId() int64`

GetBastionIssuerId returns the BastionIssuerId field if non-nil, zero value otherwise.

### GetBastionIssuerIdOk

`func (o *SecureRemoteAccess) GetBastionIssuerIdOk() (*int64, bool)`

GetBastionIssuerIdOk returns a tuple with the BastionIssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionIssuerId

`func (o *SecureRemoteAccess) SetBastionIssuerId(v int64)`

SetBastionIssuerId sets BastionIssuerId field to given value.

### HasBastionIssuerId

`func (o *SecureRemoteAccess) HasBastionIssuerId() bool`

HasBastionIssuerId returns a boolean if a field has been set.

### GetBastionSsh

`func (o *SecureRemoteAccess) GetBastionSsh() string`

GetBastionSsh returns the BastionSsh field if non-nil, zero value otherwise.

### GetBastionSshOk

`func (o *SecureRemoteAccess) GetBastionSshOk() (*string, bool)`

GetBastionSshOk returns a tuple with the BastionSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionSsh

`func (o *SecureRemoteAccess) SetBastionSsh(v string)`

SetBastionSsh sets BastionSsh field to given value.

### HasBastionSsh

`func (o *SecureRemoteAccess) HasBastionSsh() bool`

HasBastionSsh returns a boolean if a field has been set.

### GetCategory

`func (o *SecureRemoteAccess) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SecureRemoteAccess) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SecureRemoteAccess) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SecureRemoteAccess) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *SecureRemoteAccess) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *SecureRemoteAccess) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *SecureRemoteAccess) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *SecureRemoteAccess) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetDbName

`func (o *SecureRemoteAccess) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SecureRemoteAccess) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SecureRemoteAccess) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *SecureRemoteAccess) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDomain

`func (o *SecureRemoteAccess) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SecureRemoteAccess) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SecureRemoteAccess) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SecureRemoteAccess) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEnable

`func (o *SecureRemoteAccess) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SecureRemoteAccess) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SecureRemoteAccess) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SecureRemoteAccess) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEndpoint

`func (o *SecureRemoteAccess) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SecureRemoteAccess) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SecureRemoteAccess) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *SecureRemoteAccess) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHost

`func (o *SecureRemoteAccess) GetHost() []string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SecureRemoteAccess) GetHostOk() (*[]string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SecureRemoteAccess) SetHost(v []string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SecureRemoteAccess) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIsCli

`func (o *SecureRemoteAccess) GetIsCli() bool`

GetIsCli returns the IsCli field if non-nil, zero value otherwise.

### GetIsCliOk

`func (o *SecureRemoteAccess) GetIsCliOk() (*bool, bool)`

GetIsCliOk returns a tuple with the IsCli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCli

`func (o *SecureRemoteAccess) SetIsCli(v bool)`

SetIsCli sets IsCli field to given value.

### HasIsCli

`func (o *SecureRemoteAccess) HasIsCli() bool`

HasIsCli returns a boolean if a field has been set.

### GetIsWeb

`func (o *SecureRemoteAccess) GetIsWeb() bool`

GetIsWeb returns the IsWeb field if non-nil, zero value otherwise.

### GetIsWebOk

`func (o *SecureRemoteAccess) GetIsWebOk() (*bool, bool)`

GetIsWebOk returns a tuple with the IsWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeb

`func (o *SecureRemoteAccess) SetIsWeb(v bool)`

SetIsWeb sets IsWeb field to given value.

### HasIsWeb

`func (o *SecureRemoteAccess) HasIsWeb() bool`

HasIsWeb returns a boolean if a field has been set.

### GetIsolated

`func (o *SecureRemoteAccess) GetIsolated() bool`

GetIsolated returns the Isolated field if non-nil, zero value otherwise.

### GetIsolatedOk

`func (o *SecureRemoteAccess) GetIsolatedOk() (*bool, bool)`

GetIsolatedOk returns a tuple with the Isolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolated

`func (o *SecureRemoteAccess) SetIsolated(v bool)`

SetIsolated sets Isolated field to given value.

### HasIsolated

`func (o *SecureRemoteAccess) HasIsolated() bool`

HasIsolated returns a boolean if a field has been set.

### GetNative

`func (o *SecureRemoteAccess) GetNative() bool`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *SecureRemoteAccess) GetNativeOk() (*bool, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *SecureRemoteAccess) SetNative(v bool)`

SetNative sets Native field to given value.

### HasNative

`func (o *SecureRemoteAccess) HasNative() bool`

HasNative returns a boolean if a field has been set.

### GetRdpUser

`func (o *SecureRemoteAccess) GetRdpUser() string`

GetRdpUser returns the RdpUser field if non-nil, zero value otherwise.

### GetRdpUserOk

`func (o *SecureRemoteAccess) GetRdpUserOk() (*string, bool)`

GetRdpUserOk returns a tuple with the RdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpUser

`func (o *SecureRemoteAccess) SetRdpUser(v string)`

SetRdpUser sets RdpUser field to given value.

### HasRdpUser

`func (o *SecureRemoteAccess) HasRdpUser() bool`

HasRdpUser returns a boolean if a field has been set.

### GetRegion

`func (o *SecureRemoteAccess) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SecureRemoteAccess) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SecureRemoteAccess) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SecureRemoteAccess) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSchema

`func (o *SecureRemoteAccess) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SecureRemoteAccess) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SecureRemoteAccess) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SecureRemoteAccess) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSshPassword

`func (o *SecureRemoteAccess) GetSshPassword() bool`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *SecureRemoteAccess) GetSshPasswordOk() (*bool, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *SecureRemoteAccess) SetSshPassword(v bool)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *SecureRemoteAccess) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPrivateKey

`func (o *SecureRemoteAccess) GetSshPrivateKey() bool`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *SecureRemoteAccess) GetSshPrivateKeyOk() (*bool, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *SecureRemoteAccess) SetSshPrivateKey(v bool)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *SecureRemoteAccess) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### GetSshUser

`func (o *SecureRemoteAccess) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *SecureRemoteAccess) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *SecureRemoteAccess) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *SecureRemoteAccess) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetUrl

`func (o *SecureRemoteAccess) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SecureRemoteAccess) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SecureRemoteAccess) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SecureRemoteAccess) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUseInternalBastion

`func (o *SecureRemoteAccess) GetUseInternalBastion() bool`

GetUseInternalBastion returns the UseInternalBastion field if non-nil, zero value otherwise.

### GetUseInternalBastionOk

`func (o *SecureRemoteAccess) GetUseInternalBastionOk() (*bool, bool)`

GetUseInternalBastionOk returns a tuple with the UseInternalBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseInternalBastion

`func (o *SecureRemoteAccess) SetUseInternalBastion(v bool)`

SetUseInternalBastion sets UseInternalBastion field to given value.

### HasUseInternalBastion

`func (o *SecureRemoteAccess) HasUseInternalBastion() bool`

HasUseInternalBastion returns a boolean if a field has been set.

### GetWebProxy

`func (o *SecureRemoteAccess) GetWebProxy() bool`

GetWebProxy returns the WebProxy field if non-nil, zero value otherwise.

### GetWebProxyOk

`func (o *SecureRemoteAccess) GetWebProxyOk() (*bool, bool)`

GetWebProxyOk returns a tuple with the WebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebProxy

`func (o *SecureRemoteAccess) SetWebProxy(v bool)`

SetWebProxy sets WebProxy field to given value.

### HasWebProxy

`func (o *SecureRemoteAccess) HasWebProxy() bool`

HasWebProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


