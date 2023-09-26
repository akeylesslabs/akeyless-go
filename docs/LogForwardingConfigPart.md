# LogForwardingConfigPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsS3Config** | Pointer to [**AwsS3LogForwardingConfig**](AwsS3LogForwardingConfig.md) |  | [optional] 
**AzureAnalyticsConfig** | Pointer to [**AzureLogAnalyticsForwardingConfig**](AzureLogAnalyticsForwardingConfig.md) |  | [optional] 
**DatadogConfig** | Pointer to [**DatadogForwardingConfig**](DatadogForwardingConfig.md) |  | [optional] 
**ElasticsearchConfig** | Pointer to [**ElasticsearchLogForwardingConfig**](ElasticsearchLogForwardingConfig.md) |  | [optional] 
**GoogleChronicleConfig** | Pointer to [**GoogleChronicleForwardingConfig**](GoogleChronicleForwardingConfig.md) |  | [optional] 
**JsonOutput** | Pointer to **bool** |  | [optional] 
**LoganEnable** | Pointer to **bool** |  | [optional] 
**LoganUrl** | Pointer to **string** |  | [optional] 
**LogstashConfig** | Pointer to [**LogstashLogForwardingConfig**](LogstashLogForwardingConfig.md) |  | [optional] 
**LogzIoConfig** | Pointer to [**LogzIoLogForwardingConfig**](LogzIoLogForwardingConfig.md) |  | [optional] 
**PullIntervalSec** | Pointer to **string** |  | [optional] 
**SplunkConfig** | Pointer to [**SplunkLogForwardingConfig**](SplunkLogForwardingConfig.md) |  | [optional] 
**SumoLogicConfig** | Pointer to [**SumologicLogForwardingConfig**](SumologicLogForwardingConfig.md) |  | [optional] 
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

### GetAwsS3Config

`func (o *LogForwardingConfigPart) GetAwsS3Config() AwsS3LogForwardingConfig`

GetAwsS3Config returns the AwsS3Config field if non-nil, zero value otherwise.

### GetAwsS3ConfigOk

`func (o *LogForwardingConfigPart) GetAwsS3ConfigOk() (*AwsS3LogForwardingConfig, bool)`

GetAwsS3ConfigOk returns a tuple with the AwsS3Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3Config

`func (o *LogForwardingConfigPart) SetAwsS3Config(v AwsS3LogForwardingConfig)`

SetAwsS3Config sets AwsS3Config field to given value.

### HasAwsS3Config

`func (o *LogForwardingConfigPart) HasAwsS3Config() bool`

HasAwsS3Config returns a boolean if a field has been set.

### GetAzureAnalyticsConfig

`func (o *LogForwardingConfigPart) GetAzureAnalyticsConfig() AzureLogAnalyticsForwardingConfig`

GetAzureAnalyticsConfig returns the AzureAnalyticsConfig field if non-nil, zero value otherwise.

### GetAzureAnalyticsConfigOk

`func (o *LogForwardingConfigPart) GetAzureAnalyticsConfigOk() (*AzureLogAnalyticsForwardingConfig, bool)`

GetAzureAnalyticsConfigOk returns a tuple with the AzureAnalyticsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAnalyticsConfig

`func (o *LogForwardingConfigPart) SetAzureAnalyticsConfig(v AzureLogAnalyticsForwardingConfig)`

SetAzureAnalyticsConfig sets AzureAnalyticsConfig field to given value.

### HasAzureAnalyticsConfig

`func (o *LogForwardingConfigPart) HasAzureAnalyticsConfig() bool`

HasAzureAnalyticsConfig returns a boolean if a field has been set.

### GetDatadogConfig

`func (o *LogForwardingConfigPart) GetDatadogConfig() DatadogForwardingConfig`

GetDatadogConfig returns the DatadogConfig field if non-nil, zero value otherwise.

### GetDatadogConfigOk

`func (o *LogForwardingConfigPart) GetDatadogConfigOk() (*DatadogForwardingConfig, bool)`

GetDatadogConfigOk returns a tuple with the DatadogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadogConfig

`func (o *LogForwardingConfigPart) SetDatadogConfig(v DatadogForwardingConfig)`

SetDatadogConfig sets DatadogConfig field to given value.

### HasDatadogConfig

`func (o *LogForwardingConfigPart) HasDatadogConfig() bool`

HasDatadogConfig returns a boolean if a field has been set.

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

### GetGoogleChronicleConfig

`func (o *LogForwardingConfigPart) GetGoogleChronicleConfig() GoogleChronicleForwardingConfig`

GetGoogleChronicleConfig returns the GoogleChronicleConfig field if non-nil, zero value otherwise.

### GetGoogleChronicleConfigOk

`func (o *LogForwardingConfigPart) GetGoogleChronicleConfigOk() (*GoogleChronicleForwardingConfig, bool)`

GetGoogleChronicleConfigOk returns a tuple with the GoogleChronicleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleChronicleConfig

`func (o *LogForwardingConfigPart) SetGoogleChronicleConfig(v GoogleChronicleForwardingConfig)`

SetGoogleChronicleConfig sets GoogleChronicleConfig field to given value.

### HasGoogleChronicleConfig

`func (o *LogForwardingConfigPart) HasGoogleChronicleConfig() bool`

HasGoogleChronicleConfig returns a boolean if a field has been set.

### GetJsonOutput

`func (o *LogForwardingConfigPart) GetJsonOutput() bool`

GetJsonOutput returns the JsonOutput field if non-nil, zero value otherwise.

### GetJsonOutputOk

`func (o *LogForwardingConfigPart) GetJsonOutputOk() (*bool, bool)`

GetJsonOutputOk returns a tuple with the JsonOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonOutput

`func (o *LogForwardingConfigPart) SetJsonOutput(v bool)`

SetJsonOutput sets JsonOutput field to given value.

### HasJsonOutput

`func (o *LogForwardingConfigPart) HasJsonOutput() bool`

HasJsonOutput returns a boolean if a field has been set.

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

### GetSumoLogicConfig

`func (o *LogForwardingConfigPart) GetSumoLogicConfig() SumologicLogForwardingConfig`

GetSumoLogicConfig returns the SumoLogicConfig field if non-nil, zero value otherwise.

### GetSumoLogicConfigOk

`func (o *LogForwardingConfigPart) GetSumoLogicConfigOk() (*SumologicLogForwardingConfig, bool)`

GetSumoLogicConfigOk returns a tuple with the SumoLogicConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumoLogicConfig

`func (o *LogForwardingConfigPart) SetSumoLogicConfig(v SumologicLogForwardingConfig)`

SetSumoLogicConfig sets SumoLogicConfig field to given value.

### HasSumoLogicConfig

`func (o *LogForwardingConfigPart) HasSumoLogicConfig() bool`

HasSumoLogicConfig returns a boolean if a field has been set.

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


