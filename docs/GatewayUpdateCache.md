# GatewayUpdateCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupInterval** | Pointer to **string** | Secure backup interval in minutes. To ensure service continuity in case of power cycle and network outage secrets will be backed up periodically per backup interval | [optional] [default to "1"]
**EnableCache** | Pointer to **string** | Enable cache [true/false] | [optional] 
**EnableProactive** | Pointer to **string** | Enable proactive caching [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MinimumFetchInterval** | Pointer to **string** | When using Cache or/and Proactive Cache, additional secrets will be fetched upon requesting a secret, based on the requestor&#39;s access policy. Define minimum fetching interval to avoid over fetching in a given time frame | [optional] [default to "5"]
**StaleTimeout** | Pointer to **string** | Stale timeout in minutes, cache entries which are not accessed within timeout will be removed from cache | [optional] [default to "60"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateCache

`func NewGatewayUpdateCache() *GatewayUpdateCache`

NewGatewayUpdateCache instantiates a new GatewayUpdateCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateCacheWithDefaults

`func NewGatewayUpdateCacheWithDefaults() *GatewayUpdateCache`

NewGatewayUpdateCacheWithDefaults instantiates a new GatewayUpdateCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupInterval

`func (o *GatewayUpdateCache) GetBackupInterval() string`

GetBackupInterval returns the BackupInterval field if non-nil, zero value otherwise.

### GetBackupIntervalOk

`func (o *GatewayUpdateCache) GetBackupIntervalOk() (*string, bool)`

GetBackupIntervalOk returns a tuple with the BackupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInterval

`func (o *GatewayUpdateCache) SetBackupInterval(v string)`

SetBackupInterval sets BackupInterval field to given value.

### HasBackupInterval

`func (o *GatewayUpdateCache) HasBackupInterval() bool`

HasBackupInterval returns a boolean if a field has been set.

### GetEnableCache

`func (o *GatewayUpdateCache) GetEnableCache() string`

GetEnableCache returns the EnableCache field if non-nil, zero value otherwise.

### GetEnableCacheOk

`func (o *GatewayUpdateCache) GetEnableCacheOk() (*string, bool)`

GetEnableCacheOk returns a tuple with the EnableCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCache

`func (o *GatewayUpdateCache) SetEnableCache(v string)`

SetEnableCache sets EnableCache field to given value.

### HasEnableCache

`func (o *GatewayUpdateCache) HasEnableCache() bool`

HasEnableCache returns a boolean if a field has been set.

### GetEnableProactive

`func (o *GatewayUpdateCache) GetEnableProactive() string`

GetEnableProactive returns the EnableProactive field if non-nil, zero value otherwise.

### GetEnableProactiveOk

`func (o *GatewayUpdateCache) GetEnableProactiveOk() (*string, bool)`

GetEnableProactiveOk returns a tuple with the EnableProactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProactive

`func (o *GatewayUpdateCache) SetEnableProactive(v string)`

SetEnableProactive sets EnableProactive field to given value.

### HasEnableProactive

`func (o *GatewayUpdateCache) HasEnableProactive() bool`

HasEnableProactive returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateCache) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateCache) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateCache) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateCache) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMinimumFetchInterval

`func (o *GatewayUpdateCache) GetMinimumFetchInterval() string`

GetMinimumFetchInterval returns the MinimumFetchInterval field if non-nil, zero value otherwise.

### GetMinimumFetchIntervalOk

`func (o *GatewayUpdateCache) GetMinimumFetchIntervalOk() (*string, bool)`

GetMinimumFetchIntervalOk returns a tuple with the MinimumFetchInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFetchInterval

`func (o *GatewayUpdateCache) SetMinimumFetchInterval(v string)`

SetMinimumFetchInterval sets MinimumFetchInterval field to given value.

### HasMinimumFetchInterval

`func (o *GatewayUpdateCache) HasMinimumFetchInterval() bool`

HasMinimumFetchInterval returns a boolean if a field has been set.

### GetStaleTimeout

`func (o *GatewayUpdateCache) GetStaleTimeout() string`

GetStaleTimeout returns the StaleTimeout field if non-nil, zero value otherwise.

### GetStaleTimeoutOk

`func (o *GatewayUpdateCache) GetStaleTimeoutOk() (*string, bool)`

GetStaleTimeoutOk returns a tuple with the StaleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTimeout

`func (o *GatewayUpdateCache) SetStaleTimeout(v string)`

SetStaleTimeout sets StaleTimeout field to given value.

### HasStaleTimeout

`func (o *GatewayUpdateCache) HasStaleTimeout() bool`

HasStaleTimeout returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateCache) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateCache) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateCache) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateCache) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateCache) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateCache) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateCache) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateCache) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


