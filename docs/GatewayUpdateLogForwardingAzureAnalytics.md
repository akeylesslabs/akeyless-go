# GatewayUpdateLogForwardingAzureAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**WorkspaceId** | Pointer to **string** | Azure workspace id | [optional] 
**WorkspaceKey** | Pointer to **string** | Azure workspace key | [optional] 

## Methods

### NewGatewayUpdateLogForwardingAzureAnalytics

`func NewGatewayUpdateLogForwardingAzureAnalytics() *GatewayUpdateLogForwardingAzureAnalytics`

NewGatewayUpdateLogForwardingAzureAnalytics instantiates a new GatewayUpdateLogForwardingAzureAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLogForwardingAzureAnalyticsWithDefaults

`func NewGatewayUpdateLogForwardingAzureAnalyticsWithDefaults() *GatewayUpdateLogForwardingAzureAnalytics`

NewGatewayUpdateLogForwardingAzureAnalyticsWithDefaults instantiates a new GatewayUpdateLogForwardingAzureAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayUpdateLogForwardingAzureAnalytics) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayUpdateLogForwardingAzureAnalytics) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateLogForwardingAzureAnalytics) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateLogForwardingAzureAnalytics) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GatewayUpdateLogForwardingAzureAnalytics) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GatewayUpdateLogForwardingAzureAnalytics) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GatewayUpdateLogForwardingAzureAnalytics) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GatewayUpdateLogForwardingAzureAnalytics) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLogForwardingAzureAnalytics) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLogForwardingAzureAnalytics) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLogForwardingAzureAnalytics) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLogForwardingAzureAnalytics) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GatewayUpdateLogForwardingAzureAnalytics) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *GatewayUpdateLogForwardingAzureAnalytics) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetWorkspaceKey

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetWorkspaceKey() string`

GetWorkspaceKey returns the WorkspaceKey field if non-nil, zero value otherwise.

### GetWorkspaceKeyOk

`func (o *GatewayUpdateLogForwardingAzureAnalytics) GetWorkspaceKeyOk() (*string, bool)`

GetWorkspaceKeyOk returns a tuple with the WorkspaceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceKey

`func (o *GatewayUpdateLogForwardingAzureAnalytics) SetWorkspaceKey(v string)`

SetWorkspaceKey sets WorkspaceKey field to given value.

### HasWorkspaceKey

`func (o *GatewayUpdateLogForwardingAzureAnalytics) HasWorkspaceKey() bool`

HasWorkspaceKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


