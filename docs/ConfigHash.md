# ConfigHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admins** | Pointer to **string** |  | [optional] 
**Cache** | Pointer to **string** |  | [optional] 
**CustomerFragements** | Pointer to **string** |  | [optional] 
**General** | Pointer to **string** |  | [optional] 
**K8sAuths** | Pointer to **string** |  | [optional] 
**Kmip** | Pointer to **string** |  | [optional] 
**Ldap** | Pointer to **string** |  | [optional] 
**Leadership** | Pointer to **string** |  | [optional] 
**LogForwarding** | Pointer to **string** |  | [optional] 
**MQueue** | Pointer to **string** |  | [optional] 
**Migrations** | Pointer to **string** |  | [optional] 
**Producers** | Pointer to **map[string]interface{}** |  | [optional] 
**ProducersStatus** | Pointer to **string** |  | [optional] 
**Rotators** | Pointer to **map[string]interface{}** |  | [optional] 
**Saml** | Pointer to **string** |  | [optional] 
**UniversalIdentity** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigHash

`func NewConfigHash() *ConfigHash`

NewConfigHash instantiates a new ConfigHash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigHashWithDefaults

`func NewConfigHashWithDefaults() *ConfigHash`

NewConfigHashWithDefaults instantiates a new ConfigHash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmins

`func (o *ConfigHash) GetAdmins() string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *ConfigHash) GetAdminsOk() (*string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *ConfigHash) SetAdmins(v string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *ConfigHash) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetCache

`func (o *ConfigHash) GetCache() string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *ConfigHash) GetCacheOk() (*string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *ConfigHash) SetCache(v string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *ConfigHash) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCustomerFragements

`func (o *ConfigHash) GetCustomerFragements() string`

GetCustomerFragements returns the CustomerFragements field if non-nil, zero value otherwise.

### GetCustomerFragementsOk

`func (o *ConfigHash) GetCustomerFragementsOk() (*string, bool)`

GetCustomerFragementsOk returns a tuple with the CustomerFragements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFragements

`func (o *ConfigHash) SetCustomerFragements(v string)`

SetCustomerFragements sets CustomerFragements field to given value.

### HasCustomerFragements

`func (o *ConfigHash) HasCustomerFragements() bool`

HasCustomerFragements returns a boolean if a field has been set.

### GetGeneral

`func (o *ConfigHash) GetGeneral() string`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *ConfigHash) GetGeneralOk() (*string, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *ConfigHash) SetGeneral(v string)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *ConfigHash) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetK8sAuths

`func (o *ConfigHash) GetK8sAuths() string`

GetK8sAuths returns the K8sAuths field if non-nil, zero value otherwise.

### GetK8sAuthsOk

`func (o *ConfigHash) GetK8sAuthsOk() (*string, bool)`

GetK8sAuthsOk returns a tuple with the K8sAuths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAuths

`func (o *ConfigHash) SetK8sAuths(v string)`

SetK8sAuths sets K8sAuths field to given value.

### HasK8sAuths

`func (o *ConfigHash) HasK8sAuths() bool`

HasK8sAuths returns a boolean if a field has been set.

### GetKmip

`func (o *ConfigHash) GetKmip() string`

GetKmip returns the Kmip field if non-nil, zero value otherwise.

### GetKmipOk

`func (o *ConfigHash) GetKmipOk() (*string, bool)`

GetKmipOk returns a tuple with the Kmip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmip

`func (o *ConfigHash) SetKmip(v string)`

SetKmip sets Kmip field to given value.

### HasKmip

`func (o *ConfigHash) HasKmip() bool`

HasKmip returns a boolean if a field has been set.

### GetLdap

`func (o *ConfigHash) GetLdap() string`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *ConfigHash) GetLdapOk() (*string, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *ConfigHash) SetLdap(v string)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *ConfigHash) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetLeadership

`func (o *ConfigHash) GetLeadership() string`

GetLeadership returns the Leadership field if non-nil, zero value otherwise.

### GetLeadershipOk

`func (o *ConfigHash) GetLeadershipOk() (*string, bool)`

GetLeadershipOk returns a tuple with the Leadership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadership

`func (o *ConfigHash) SetLeadership(v string)`

SetLeadership sets Leadership field to given value.

### HasLeadership

`func (o *ConfigHash) HasLeadership() bool`

HasLeadership returns a boolean if a field has been set.

### GetLogForwarding

`func (o *ConfigHash) GetLogForwarding() string`

GetLogForwarding returns the LogForwarding field if non-nil, zero value otherwise.

### GetLogForwardingOk

`func (o *ConfigHash) GetLogForwardingOk() (*string, bool)`

GetLogForwardingOk returns a tuple with the LogForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogForwarding

`func (o *ConfigHash) SetLogForwarding(v string)`

SetLogForwarding sets LogForwarding field to given value.

### HasLogForwarding

`func (o *ConfigHash) HasLogForwarding() bool`

HasLogForwarding returns a boolean if a field has been set.

### GetMQueue

`func (o *ConfigHash) GetMQueue() string`

GetMQueue returns the MQueue field if non-nil, zero value otherwise.

### GetMQueueOk

`func (o *ConfigHash) GetMQueueOk() (*string, bool)`

GetMQueueOk returns a tuple with the MQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMQueue

`func (o *ConfigHash) SetMQueue(v string)`

SetMQueue sets MQueue field to given value.

### HasMQueue

`func (o *ConfigHash) HasMQueue() bool`

HasMQueue returns a boolean if a field has been set.

### GetMigrations

`func (o *ConfigHash) GetMigrations() string`

GetMigrations returns the Migrations field if non-nil, zero value otherwise.

### GetMigrationsOk

`func (o *ConfigHash) GetMigrationsOk() (*string, bool)`

GetMigrationsOk returns a tuple with the Migrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrations

`func (o *ConfigHash) SetMigrations(v string)`

SetMigrations sets Migrations field to given value.

### HasMigrations

`func (o *ConfigHash) HasMigrations() bool`

HasMigrations returns a boolean if a field has been set.

### GetProducers

`func (o *ConfigHash) GetProducers() map[string]interface{}`

GetProducers returns the Producers field if non-nil, zero value otherwise.

### GetProducersOk

`func (o *ConfigHash) GetProducersOk() (*map[string]interface{}, bool)`

GetProducersOk returns a tuple with the Producers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducers

`func (o *ConfigHash) SetProducers(v map[string]interface{})`

SetProducers sets Producers field to given value.

### HasProducers

`func (o *ConfigHash) HasProducers() bool`

HasProducers returns a boolean if a field has been set.

### GetProducersStatus

`func (o *ConfigHash) GetProducersStatus() string`

GetProducersStatus returns the ProducersStatus field if non-nil, zero value otherwise.

### GetProducersStatusOk

`func (o *ConfigHash) GetProducersStatusOk() (*string, bool)`

GetProducersStatusOk returns a tuple with the ProducersStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducersStatus

`func (o *ConfigHash) SetProducersStatus(v string)`

SetProducersStatus sets ProducersStatus field to given value.

### HasProducersStatus

`func (o *ConfigHash) HasProducersStatus() bool`

HasProducersStatus returns a boolean if a field has been set.

### GetRotators

`func (o *ConfigHash) GetRotators() map[string]interface{}`

GetRotators returns the Rotators field if non-nil, zero value otherwise.

### GetRotatorsOk

`func (o *ConfigHash) GetRotatorsOk() (*map[string]interface{}, bool)`

GetRotatorsOk returns a tuple with the Rotators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotators

`func (o *ConfigHash) SetRotators(v map[string]interface{})`

SetRotators sets Rotators field to given value.

### HasRotators

`func (o *ConfigHash) HasRotators() bool`

HasRotators returns a boolean if a field has been set.

### GetSaml

`func (o *ConfigHash) GetSaml() string`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *ConfigHash) GetSamlOk() (*string, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *ConfigHash) SetSaml(v string)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *ConfigHash) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetUniversalIdentity

`func (o *ConfigHash) GetUniversalIdentity() string`

GetUniversalIdentity returns the UniversalIdentity field if non-nil, zero value otherwise.

### GetUniversalIdentityOk

`func (o *ConfigHash) GetUniversalIdentityOk() (*string, bool)`

GetUniversalIdentityOk returns a tuple with the UniversalIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalIdentity

`func (o *ConfigHash) SetUniversalIdentity(v string)`

SetUniversalIdentity sets UniversalIdentity field to given value.

### HasUniversalIdentity

`func (o *ConfigHash) HasUniversalIdentity() bool`

HasUniversalIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


