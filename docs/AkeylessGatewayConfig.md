# AkeylessGatewayConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admins** | Pointer to [**AdminsConfigPart**](AdminsConfigPart.md) |  | [optional] 
**Cache** | Pointer to [**CacheConfigPart**](CacheConfigPart.md) |  | [optional] 
**Cf** | Pointer to [**CFConfigPart**](CFConfigPart.md) |  | [optional] 
**ConfigProtectionKeyName** | Pointer to **string** |  | [optional] 
**General** | Pointer to [**GeneralConfigPart**](GeneralConfigPart.md) |  | [optional] 
**K8sAuths** | Pointer to [**K8SAuthsConfigPart**](K8SAuthsConfigPart.md) |  | [optional] 
**KmipClients** | Pointer to [**KMIPConfigPart**](KMIPConfigPart.md) |  | [optional] 
**Ldap** | Pointer to [**LdapConfigPart**](LdapConfigPart.md) |  | [optional] 
**Leadership** | Pointer to [**LeadershipConfigPart**](LeadershipConfigPart.md) |  | [optional] 
**LogForwarding** | Pointer to [**LogForwardingConfigPart**](LogForwardingConfigPart.md) |  | [optional] 
**Migrations** | Pointer to [**MigrationsConfigPart**](MigrationsConfigPart.md) |  | [optional] 
**Producers** | Pointer to [**ProducersConfigPart**](ProducersConfigPart.md) |  | [optional] 
**Rotators** | Pointer to [**RotatorsConfigPart**](RotatorsConfigPart.md) |  | [optional] 
**Saml** | Pointer to [**DefaultConfigPart**](DefaultConfigPart.md) |  | [optional] 
**Uidentity** | Pointer to [**UIdentityConfigPart**](UIdentityConfigPart.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewAkeylessGatewayConfig

`func NewAkeylessGatewayConfig() *AkeylessGatewayConfig`

NewAkeylessGatewayConfig instantiates a new AkeylessGatewayConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAkeylessGatewayConfigWithDefaults

`func NewAkeylessGatewayConfigWithDefaults() *AkeylessGatewayConfig`

NewAkeylessGatewayConfigWithDefaults instantiates a new AkeylessGatewayConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmins

`func (o *AkeylessGatewayConfig) GetAdmins() AdminsConfigPart`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *AkeylessGatewayConfig) GetAdminsOk() (*AdminsConfigPart, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *AkeylessGatewayConfig) SetAdmins(v AdminsConfigPart)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *AkeylessGatewayConfig) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetCache

`func (o *AkeylessGatewayConfig) GetCache() CacheConfigPart`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *AkeylessGatewayConfig) GetCacheOk() (*CacheConfigPart, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *AkeylessGatewayConfig) SetCache(v CacheConfigPart)`

SetCache sets Cache field to given value.

### HasCache

`func (o *AkeylessGatewayConfig) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCf

`func (o *AkeylessGatewayConfig) GetCf() CFConfigPart`

GetCf returns the Cf field if non-nil, zero value otherwise.

### GetCfOk

`func (o *AkeylessGatewayConfig) GetCfOk() (*CFConfigPart, bool)`

GetCfOk returns a tuple with the Cf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCf

`func (o *AkeylessGatewayConfig) SetCf(v CFConfigPart)`

SetCf sets Cf field to given value.

### HasCf

`func (o *AkeylessGatewayConfig) HasCf() bool`

HasCf returns a boolean if a field has been set.

### GetConfigProtectionKeyName

`func (o *AkeylessGatewayConfig) GetConfigProtectionKeyName() string`

GetConfigProtectionKeyName returns the ConfigProtectionKeyName field if non-nil, zero value otherwise.

### GetConfigProtectionKeyNameOk

`func (o *AkeylessGatewayConfig) GetConfigProtectionKeyNameOk() (*string, bool)`

GetConfigProtectionKeyNameOk returns a tuple with the ConfigProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigProtectionKeyName

`func (o *AkeylessGatewayConfig) SetConfigProtectionKeyName(v string)`

SetConfigProtectionKeyName sets ConfigProtectionKeyName field to given value.

### HasConfigProtectionKeyName

`func (o *AkeylessGatewayConfig) HasConfigProtectionKeyName() bool`

HasConfigProtectionKeyName returns a boolean if a field has been set.

### GetGeneral

`func (o *AkeylessGatewayConfig) GetGeneral() GeneralConfigPart`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *AkeylessGatewayConfig) GetGeneralOk() (*GeneralConfigPart, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *AkeylessGatewayConfig) SetGeneral(v GeneralConfigPart)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *AkeylessGatewayConfig) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetK8sAuths

`func (o *AkeylessGatewayConfig) GetK8sAuths() K8SAuthsConfigPart`

GetK8sAuths returns the K8sAuths field if non-nil, zero value otherwise.

### GetK8sAuthsOk

`func (o *AkeylessGatewayConfig) GetK8sAuthsOk() (*K8SAuthsConfigPart, bool)`

GetK8sAuthsOk returns a tuple with the K8sAuths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAuths

`func (o *AkeylessGatewayConfig) SetK8sAuths(v K8SAuthsConfigPart)`

SetK8sAuths sets K8sAuths field to given value.

### HasK8sAuths

`func (o *AkeylessGatewayConfig) HasK8sAuths() bool`

HasK8sAuths returns a boolean if a field has been set.

### GetKmipClients

`func (o *AkeylessGatewayConfig) GetKmipClients() KMIPConfigPart`

GetKmipClients returns the KmipClients field if non-nil, zero value otherwise.

### GetKmipClientsOk

`func (o *AkeylessGatewayConfig) GetKmipClientsOk() (*KMIPConfigPart, bool)`

GetKmipClientsOk returns a tuple with the KmipClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipClients

`func (o *AkeylessGatewayConfig) SetKmipClients(v KMIPConfigPart)`

SetKmipClients sets KmipClients field to given value.

### HasKmipClients

`func (o *AkeylessGatewayConfig) HasKmipClients() bool`

HasKmipClients returns a boolean if a field has been set.

### GetLdap

`func (o *AkeylessGatewayConfig) GetLdap() LdapConfigPart`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *AkeylessGatewayConfig) GetLdapOk() (*LdapConfigPart, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *AkeylessGatewayConfig) SetLdap(v LdapConfigPart)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *AkeylessGatewayConfig) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetLeadership

`func (o *AkeylessGatewayConfig) GetLeadership() LeadershipConfigPart`

GetLeadership returns the Leadership field if non-nil, zero value otherwise.

### GetLeadershipOk

`func (o *AkeylessGatewayConfig) GetLeadershipOk() (*LeadershipConfigPart, bool)`

GetLeadershipOk returns a tuple with the Leadership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadership

`func (o *AkeylessGatewayConfig) SetLeadership(v LeadershipConfigPart)`

SetLeadership sets Leadership field to given value.

### HasLeadership

`func (o *AkeylessGatewayConfig) HasLeadership() bool`

HasLeadership returns a boolean if a field has been set.

### GetLogForwarding

`func (o *AkeylessGatewayConfig) GetLogForwarding() LogForwardingConfigPart`

GetLogForwarding returns the LogForwarding field if non-nil, zero value otherwise.

### GetLogForwardingOk

`func (o *AkeylessGatewayConfig) GetLogForwardingOk() (*LogForwardingConfigPart, bool)`

GetLogForwardingOk returns a tuple with the LogForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogForwarding

`func (o *AkeylessGatewayConfig) SetLogForwarding(v LogForwardingConfigPart)`

SetLogForwarding sets LogForwarding field to given value.

### HasLogForwarding

`func (o *AkeylessGatewayConfig) HasLogForwarding() bool`

HasLogForwarding returns a boolean if a field has been set.

### GetMigrations

`func (o *AkeylessGatewayConfig) GetMigrations() MigrationsConfigPart`

GetMigrations returns the Migrations field if non-nil, zero value otherwise.

### GetMigrationsOk

`func (o *AkeylessGatewayConfig) GetMigrationsOk() (*MigrationsConfigPart, bool)`

GetMigrationsOk returns a tuple with the Migrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrations

`func (o *AkeylessGatewayConfig) SetMigrations(v MigrationsConfigPart)`

SetMigrations sets Migrations field to given value.

### HasMigrations

`func (o *AkeylessGatewayConfig) HasMigrations() bool`

HasMigrations returns a boolean if a field has been set.

### GetProducers

`func (o *AkeylessGatewayConfig) GetProducers() ProducersConfigPart`

GetProducers returns the Producers field if non-nil, zero value otherwise.

### GetProducersOk

`func (o *AkeylessGatewayConfig) GetProducersOk() (*ProducersConfigPart, bool)`

GetProducersOk returns a tuple with the Producers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducers

`func (o *AkeylessGatewayConfig) SetProducers(v ProducersConfigPart)`

SetProducers sets Producers field to given value.

### HasProducers

`func (o *AkeylessGatewayConfig) HasProducers() bool`

HasProducers returns a boolean if a field has been set.

### GetRotators

`func (o *AkeylessGatewayConfig) GetRotators() RotatorsConfigPart`

GetRotators returns the Rotators field if non-nil, zero value otherwise.

### GetRotatorsOk

`func (o *AkeylessGatewayConfig) GetRotatorsOk() (*RotatorsConfigPart, bool)`

GetRotatorsOk returns a tuple with the Rotators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotators

`func (o *AkeylessGatewayConfig) SetRotators(v RotatorsConfigPart)`

SetRotators sets Rotators field to given value.

### HasRotators

`func (o *AkeylessGatewayConfig) HasRotators() bool`

HasRotators returns a boolean if a field has been set.

### GetSaml

`func (o *AkeylessGatewayConfig) GetSaml() DefaultConfigPart`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *AkeylessGatewayConfig) GetSamlOk() (*DefaultConfigPart, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *AkeylessGatewayConfig) SetSaml(v DefaultConfigPart)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *AkeylessGatewayConfig) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetUidentity

`func (o *AkeylessGatewayConfig) GetUidentity() UIdentityConfigPart`

GetUidentity returns the Uidentity field if non-nil, zero value otherwise.

### GetUidentityOk

`func (o *AkeylessGatewayConfig) GetUidentityOk() (*UIdentityConfigPart, bool)`

GetUidentityOk returns a tuple with the Uidentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidentity

`func (o *AkeylessGatewayConfig) SetUidentity(v UIdentityConfigPart)`

SetUidentity sets Uidentity field to given value.

### HasUidentity

`func (o *AkeylessGatewayConfig) HasUidentity() bool`

HasUidentity returns a boolean if a field has been set.

### GetVersion

`func (o *AkeylessGatewayConfig) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AkeylessGatewayConfig) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AkeylessGatewayConfig) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AkeylessGatewayConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


