# GatewayUpdateLogForwardingDatadog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | Datadog api key | [optional] 
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**Host** | Pointer to **string** | Datadog host | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**LogService** | Pointer to **string** | Datadog log service | [optional] [default to "use-existing"]
**LogSource** | Pointer to **string** | Datadog log source | [optional] [default to "use-existing"]
**LogTags** | Pointer to **string** | A comma-separated list of Datadog log tags formatted as \&quot;key:value\&quot; strings | [optional] [default to "use-existing"]
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateLogForwardingDatadog

`func NewGatewayUpdateLogForwardingDatadog() *GatewayUpdateLogForwardingDatadog`

NewGatewayUpdateLogForwardingDatadog instantiates a new GatewayUpdateLogForwardingDatadog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLogForwardingDatadogWithDefaults

`func NewGatewayUpdateLogForwardingDatadogWithDefaults() *GatewayUpdateLogForwardingDatadog`

NewGatewayUpdateLogForwardingDatadogWithDefaults instantiates a new GatewayUpdateLogForwardingDatadog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *GatewayUpdateLogForwardingDatadog) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GatewayUpdateLogForwardingDatadog) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GatewayUpdateLogForwardingDatadog) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GatewayUpdateLogForwardingDatadog) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetEnable

`func (o *GatewayUpdateLogForwardingDatadog) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayUpdateLogForwardingDatadog) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayUpdateLogForwardingDatadog) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayUpdateLogForwardingDatadog) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetHost

`func (o *GatewayUpdateLogForwardingDatadog) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GatewayUpdateLogForwardingDatadog) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GatewayUpdateLogForwardingDatadog) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GatewayUpdateLogForwardingDatadog) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateLogForwardingDatadog) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateLogForwardingDatadog) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateLogForwardingDatadog) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateLogForwardingDatadog) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLogService

`func (o *GatewayUpdateLogForwardingDatadog) GetLogService() string`

GetLogService returns the LogService field if non-nil, zero value otherwise.

### GetLogServiceOk

`func (o *GatewayUpdateLogForwardingDatadog) GetLogServiceOk() (*string, bool)`

GetLogServiceOk returns a tuple with the LogService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogService

`func (o *GatewayUpdateLogForwardingDatadog) SetLogService(v string)`

SetLogService sets LogService field to given value.

### HasLogService

`func (o *GatewayUpdateLogForwardingDatadog) HasLogService() bool`

HasLogService returns a boolean if a field has been set.

### GetLogSource

`func (o *GatewayUpdateLogForwardingDatadog) GetLogSource() string`

GetLogSource returns the LogSource field if non-nil, zero value otherwise.

### GetLogSourceOk

`func (o *GatewayUpdateLogForwardingDatadog) GetLogSourceOk() (*string, bool)`

GetLogSourceOk returns a tuple with the LogSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSource

`func (o *GatewayUpdateLogForwardingDatadog) SetLogSource(v string)`

SetLogSource sets LogSource field to given value.

### HasLogSource

`func (o *GatewayUpdateLogForwardingDatadog) HasLogSource() bool`

HasLogSource returns a boolean if a field has been set.

### GetLogTags

`func (o *GatewayUpdateLogForwardingDatadog) GetLogTags() string`

GetLogTags returns the LogTags field if non-nil, zero value otherwise.

### GetLogTagsOk

`func (o *GatewayUpdateLogForwardingDatadog) GetLogTagsOk() (*string, bool)`

GetLogTagsOk returns a tuple with the LogTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTags

`func (o *GatewayUpdateLogForwardingDatadog) SetLogTags(v string)`

SetLogTags sets LogTags field to given value.

### HasLogTags

`func (o *GatewayUpdateLogForwardingDatadog) HasLogTags() bool`

HasLogTags returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GatewayUpdateLogForwardingDatadog) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GatewayUpdateLogForwardingDatadog) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GatewayUpdateLogForwardingDatadog) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GatewayUpdateLogForwardingDatadog) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GatewayUpdateLogForwardingDatadog) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GatewayUpdateLogForwardingDatadog) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GatewayUpdateLogForwardingDatadog) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GatewayUpdateLogForwardingDatadog) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLogForwardingDatadog) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLogForwardingDatadog) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLogForwardingDatadog) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLogForwardingDatadog) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLogForwardingDatadog) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLogForwardingDatadog) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLogForwardingDatadog) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLogForwardingDatadog) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


