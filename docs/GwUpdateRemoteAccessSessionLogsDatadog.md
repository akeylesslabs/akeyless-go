# GwUpdateRemoteAccessSessionLogsDatadog

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

### NewGwUpdateRemoteAccessSessionLogsDatadog

`func NewGwUpdateRemoteAccessSessionLogsDatadog() *GwUpdateRemoteAccessSessionLogsDatadog`

NewGwUpdateRemoteAccessSessionLogsDatadog instantiates a new GwUpdateRemoteAccessSessionLogsDatadog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsDatadogWithDefaults

`func NewGwUpdateRemoteAccessSessionLogsDatadogWithDefaults() *GwUpdateRemoteAccessSessionLogsDatadog`

NewGwUpdateRemoteAccessSessionLogsDatadogWithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsDatadog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetHost

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLogService

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetLogService() string`

GetLogService returns the LogService field if non-nil, zero value otherwise.

### GetLogServiceOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetLogServiceOk() (*string, bool)`

GetLogServiceOk returns a tuple with the LogService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogService

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetLogService(v string)`

SetLogService sets LogService field to given value.

### HasLogService

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasLogService() bool`

HasLogService returns a boolean if a field has been set.

### GetLogSource

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetLogSource() string`

GetLogSource returns the LogSource field if non-nil, zero value otherwise.

### GetLogSourceOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetLogSourceOk() (*string, bool)`

GetLogSourceOk returns a tuple with the LogSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSource

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetLogSource(v string)`

SetLogSource sets LogSource field to given value.

### HasLogSource

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasLogSource() bool`

HasLogSource returns a boolean if a field has been set.

### GetLogTags

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetLogTags() string`

GetLogTags returns the LogTags field if non-nil, zero value otherwise.

### GetLogTagsOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetLogTagsOk() (*string, bool)`

GetLogTagsOk returns a tuple with the LogTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTags

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetLogTags(v string)`

SetLogTags sets LogTags field to given value.

### HasLogTags

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasLogTags() bool`

HasLogTags returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsDatadog) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


