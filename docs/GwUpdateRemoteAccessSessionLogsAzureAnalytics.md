# GwUpdateRemoteAccessSessionLogsAzureAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**EnableBatch** | Pointer to **string** | Enable batch forwarding [true/false] | [optional] [default to "true"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**WorkspaceId** | Pointer to **string** | Azure workspace id | [optional] 
**WorkspaceKey** | Pointer to **string** | Azure workspace key | [optional] 

## Methods

### NewGwUpdateRemoteAccessSessionLogsAzureAnalytics

`func NewGwUpdateRemoteAccessSessionLogsAzureAnalytics() *GwUpdateRemoteAccessSessionLogsAzureAnalytics`

NewGwUpdateRemoteAccessSessionLogsAzureAnalytics instantiates a new GwUpdateRemoteAccessSessionLogsAzureAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsAzureAnalyticsWithDefaults

`func NewGwUpdateRemoteAccessSessionLogsAzureAnalyticsWithDefaults() *GwUpdateRemoteAccessSessionLogsAzureAnalytics`

NewGwUpdateRemoteAccessSessionLogsAzureAnalyticsWithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsAzureAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableBatch

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetEnableBatch() string`

GetEnableBatch returns the EnableBatch field if non-nil, zero value otherwise.

### GetEnableBatchOk

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetEnableBatchOk() (*string, bool)`

GetEnableBatchOk returns a tuple with the EnableBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBatch

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) SetEnableBatch(v string)`

SetEnableBatch sets EnableBatch field to given value.

### HasEnableBatch

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) HasEnableBatch() bool`

HasEnableBatch returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetWorkspaceKey

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetWorkspaceKey() string`

GetWorkspaceKey returns the WorkspaceKey field if non-nil, zero value otherwise.

### GetWorkspaceKeyOk

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) GetWorkspaceKeyOk() (*string, bool)`

GetWorkspaceKeyOk returns a tuple with the WorkspaceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceKey

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) SetWorkspaceKey(v string)`

SetWorkspaceKey sets WorkspaceKey field to given value.

### HasWorkspaceKey

`func (o *GwUpdateRemoteAccessSessionLogsAzureAnalytics) HasWorkspaceKey() bool`

HasWorkspaceKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


