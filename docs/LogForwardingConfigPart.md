# LogForwardingConfigPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElasticsearchConfig** | Pointer to [**ElasticsearchLogForwardingConfig**](ElasticsearchLogForwardingConfig.md) |  | [optional] 
**LoganEnable** | Pointer to **bool** |  | [optional] 
**LoganUrl** | Pointer to **string** |  | [optional] 
**LogstashConfig** | Pointer to [**LogstashLogForwardingConfig**](LogstashLogForwardingConfig.md) |  | [optional] 
**LogzIoConfig** | Pointer to [**LogzIoLogForwardingConfig**](LogzIoLogForwardingConfig.md) |  | [optional] 
**PullIntervalSec** | Pointer to **string** |  | [optional] 
**SplunkConfig** | Pointer to [**SplunkLogForwardingConfig**](SplunkLogForwardingConfig.md) |  | [optional] 
**SyslogConfig** | Pointer to [**SyslogLogForwardingConfig**](SyslogLogForwardingConfig.md) |  | [optional] 
**TargetLogType** | Pointer to **string** |  | [optional] 

## Methods

### NewLogForwardingConfigPart

`func NewLogForwardingConfigPart() *LogForwardingConfigPart`

NewLogForwardingConfigPart instantiates a new LogForwardingConfigPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogForwardingConfigPartWithDefaults

`func NewLogForwardingConfigPartWithDefaults() *LogForwardingConfigPart`

NewLogForwardingConfigPartWithDefaults instantiates a new LogForwardingConfigPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElasticsearchConfig

`func (o *LogForwardingConfigPart) GetElasticsearchConfig() ElasticsearchLogForwardingConfig`

GetElasticsearchConfig returns the ElasticsearchConfig field if non-nil, zero value otherwise.

### GetElasticsearchConfigOk

`func (o *LogForwardingConfigPart) GetElasticsearchConfigOk() (*ElasticsearchLogForwardingConfig, bool)`

GetElasticsearchConfigOk returns a tuple with the ElasticsearchConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticsearchConfig

`func (o *LogForwardingConfigPart) SetElasticsearchConfig(v ElasticsearchLogForwardingConfig)`

SetElasticsearchConfig sets ElasticsearchConfig field to given value.

### HasElasticsearchConfig

`func (o *LogForwardingConfigPart) HasElasticsearchConfig() bool`

HasElasticsearchConfig returns a boolean if a field has been set.

### GetLoganEnable

`func (o *LogForwardingConfigPart) GetLoganEnable() bool`

GetLoganEnable returns the LoganEnable field if non-nil, zero value otherwise.

### GetLoganEnableOk

`func (o *LogForwardingConfigPart) GetLoganEnableOk() (*bool, bool)`

GetLoganEnableOk returns a tuple with the LoganEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoganEnable

`func (o *LogForwardingConfigPart) SetLoganEnable(v bool)`

SetLoganEnable sets LoganEnable field to given value.

### HasLoganEnable

`func (o *LogForwardingConfigPart) HasLoganEnable() bool`

HasLoganEnable returns a boolean if a field has been set.

### GetLoganUrl

`func (o *LogForwardingConfigPart) GetLoganUrl() string`

GetLoganUrl returns the LoganUrl field if non-nil, zero value otherwise.

### GetLoganUrlOk

`func (o *LogForwardingConfigPart) GetLoganUrlOk() (*string, bool)`

GetLoganUrlOk returns a tuple with the LoganUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoganUrl

`func (o *LogForwardingConfigPart) SetLoganUrl(v string)`

SetLoganUrl sets LoganUrl field to given value.

### HasLoganUrl

`func (o *LogForwardingConfigPart) HasLoganUrl() bool`

HasLoganUrl returns a boolean if a field has been set.

### GetLogstashConfig

`func (o *LogForwardingConfigPart) GetLogstashConfig() LogstashLogForwardingConfig`

GetLogstashConfig returns the LogstashConfig field if non-nil, zero value otherwise.

### GetLogstashConfigOk

`func (o *LogForwardingConfigPart) GetLogstashConfigOk() (*LogstashLogForwardingConfig, bool)`

GetLogstashConfigOk returns a tuple with the LogstashConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogstashConfig

`func (o *LogForwardingConfigPart) SetLogstashConfig(v LogstashLogForwardingConfig)`

SetLogstashConfig sets LogstashConfig field to given value.

### HasLogstashConfig

`func (o *LogForwardingConfigPart) HasLogstashConfig() bool`

HasLogstashConfig returns a boolean if a field has been set.

### GetLogzIoConfig

`func (o *LogForwardingConfigPart) GetLogzIoConfig() LogzIoLogForwardingConfig`

GetLogzIoConfig returns the LogzIoConfig field if non-nil, zero value otherwise.

### GetLogzIoConfigOk

`func (o *LogForwardingConfigPart) GetLogzIoConfigOk() (*LogzIoLogForwardingConfig, bool)`

GetLogzIoConfigOk returns a tuple with the LogzIoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogzIoConfig

`func (o *LogForwardingConfigPart) SetLogzIoConfig(v LogzIoLogForwardingConfig)`

SetLogzIoConfig sets LogzIoConfig field to given value.

### HasLogzIoConfig

`func (o *LogForwardingConfigPart) HasLogzIoConfig() bool`

HasLogzIoConfig returns a boolean if a field has been set.

### GetPullIntervalSec

`func (o *LogForwardingConfigPart) GetPullIntervalSec() string`

GetPullIntervalSec returns the PullIntervalSec field if non-nil, zero value otherwise.

### GetPullIntervalSecOk

`func (o *LogForwardingConfigPart) GetPullIntervalSecOk() (*string, bool)`

GetPullIntervalSecOk returns a tuple with the PullIntervalSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullIntervalSec

`func (o *LogForwardingConfigPart) SetPullIntervalSec(v string)`

SetPullIntervalSec sets PullIntervalSec field to given value.

### HasPullIntervalSec

`func (o *LogForwardingConfigPart) HasPullIntervalSec() bool`

HasPullIntervalSec returns a boolean if a field has been set.

### GetSplunkConfig

`func (o *LogForwardingConfigPart) GetSplunkConfig() SplunkLogForwardingConfig`

GetSplunkConfig returns the SplunkConfig field if non-nil, zero value otherwise.

### GetSplunkConfigOk

`func (o *LogForwardingConfigPart) GetSplunkConfigOk() (*SplunkLogForwardingConfig, bool)`

GetSplunkConfigOk returns a tuple with the SplunkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkConfig

`func (o *LogForwardingConfigPart) SetSplunkConfig(v SplunkLogForwardingConfig)`

SetSplunkConfig sets SplunkConfig field to given value.

### HasSplunkConfig

`func (o *LogForwardingConfigPart) HasSplunkConfig() bool`

HasSplunkConfig returns a boolean if a field has been set.

### GetSyslogConfig

`func (o *LogForwardingConfigPart) GetSyslogConfig() SyslogLogForwardingConfig`

GetSyslogConfig returns the SyslogConfig field if non-nil, zero value otherwise.

### GetSyslogConfigOk

`func (o *LogForwardingConfigPart) GetSyslogConfigOk() (*SyslogLogForwardingConfig, bool)`

GetSyslogConfigOk returns a tuple with the SyslogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogConfig

`func (o *LogForwardingConfigPart) SetSyslogConfig(v SyslogLogForwardingConfig)`

SetSyslogConfig sets SyslogConfig field to given value.

### HasSyslogConfig

`func (o *LogForwardingConfigPart) HasSyslogConfig() bool`

HasSyslogConfig returns a boolean if a field has been set.

### GetTargetLogType

`func (o *LogForwardingConfigPart) GetTargetLogType() string`

GetTargetLogType returns the TargetLogType field if non-nil, zero value otherwise.

### GetTargetLogTypeOk

`func (o *LogForwardingConfigPart) GetTargetLogTypeOk() (*string, bool)`

GetTargetLogTypeOk returns a tuple with the TargetLogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLogType

`func (o *LogForwardingConfigPart) SetTargetLogType(v string)`

SetTargetLogType sets TargetLogType field to given value.

### HasTargetLogType

`func (o *LogForwardingConfigPart) HasTargetLogType() bool`

HasTargetLogType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


