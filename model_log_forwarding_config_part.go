/*
Akeyless API

The purpose of this application is to provide access to Akeyless API.

API version: 2.0
Contact: support@akeyless.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// LogForwardingConfigPart struct for LogForwardingConfigPart
type LogForwardingConfigPart struct {
	AwsS3Config *AwsS3LogForwardingConfig `json:"aws_s3_config,omitempty"`
	AzureAnalyticsConfig *AzureLogAnalyticsForwardingConfig `json:"azure_analytics_config,omitempty"`
	DatadogConfig *DatadogForwardingConfig `json:"datadog_config,omitempty"`
	ElasticsearchConfig *ElasticsearchLogForwardingConfig `json:"elasticsearch_config,omitempty"`
	JsonOutput *bool `json:"json_output,omitempty"`
	LoganEnable *bool `json:"logan_enable,omitempty"`
	LoganUrl *string `json:"logan_url,omitempty"`
	LogstashConfig *LogstashLogForwardingConfig `json:"logstash_config,omitempty"`
	LogzIoConfig *LogzIoLogForwardingConfig `json:"logz_io_config,omitempty"`
	PullIntervalSec *string `json:"pull_interval_sec,omitempty"`
	SplunkConfig *SplunkLogForwardingConfig `json:"splunk_config,omitempty"`
	SyslogConfig *SyslogLogForwardingConfig `json:"syslog_config,omitempty"`
	TargetLogType *string `json:"target_log_type,omitempty"`
}

// NewLogForwardingConfigPart instantiates a new LogForwardingConfigPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwardingConfigPart() *LogForwardingConfigPart {
	this := LogForwardingConfigPart{}
	return &this
}

// NewLogForwardingConfigPartWithDefaults instantiates a new LogForwardingConfigPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwardingConfigPartWithDefaults() *LogForwardingConfigPart {
	this := LogForwardingConfigPart{}
	return &this
}

// GetAwsS3Config returns the AwsS3Config field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetAwsS3Config() AwsS3LogForwardingConfig {
	if o == nil || o.AwsS3Config == nil {
		var ret AwsS3LogForwardingConfig
		return ret
	}
	return *o.AwsS3Config
}

// GetAwsS3ConfigOk returns a tuple with the AwsS3Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetAwsS3ConfigOk() (*AwsS3LogForwardingConfig, bool) {
	if o == nil || o.AwsS3Config == nil {
		return nil, false
	}
	return o.AwsS3Config, true
}

// HasAwsS3Config returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasAwsS3Config() bool {
	if o != nil && o.AwsS3Config != nil {
		return true
	}

	return false
}

// SetAwsS3Config gets a reference to the given AwsS3LogForwardingConfig and assigns it to the AwsS3Config field.
func (o *LogForwardingConfigPart) SetAwsS3Config(v AwsS3LogForwardingConfig) {
	o.AwsS3Config = &v
}

// GetAzureAnalyticsConfig returns the AzureAnalyticsConfig field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetAzureAnalyticsConfig() AzureLogAnalyticsForwardingConfig {
	if o == nil || o.AzureAnalyticsConfig == nil {
		var ret AzureLogAnalyticsForwardingConfig
		return ret
	}
	return *o.AzureAnalyticsConfig
}

// GetAzureAnalyticsConfigOk returns a tuple with the AzureAnalyticsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetAzureAnalyticsConfigOk() (*AzureLogAnalyticsForwardingConfig, bool) {
	if o == nil || o.AzureAnalyticsConfig == nil {
		return nil, false
	}
	return o.AzureAnalyticsConfig, true
}

// HasAzureAnalyticsConfig returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasAzureAnalyticsConfig() bool {
	if o != nil && o.AzureAnalyticsConfig != nil {
		return true
	}

	return false
}

// SetAzureAnalyticsConfig gets a reference to the given AzureLogAnalyticsForwardingConfig and assigns it to the AzureAnalyticsConfig field.
func (o *LogForwardingConfigPart) SetAzureAnalyticsConfig(v AzureLogAnalyticsForwardingConfig) {
	o.AzureAnalyticsConfig = &v
}

// GetDatadogConfig returns the DatadogConfig field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetDatadogConfig() DatadogForwardingConfig {
	if o == nil || o.DatadogConfig == nil {
		var ret DatadogForwardingConfig
		return ret
	}
	return *o.DatadogConfig
}

// GetDatadogConfigOk returns a tuple with the DatadogConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetDatadogConfigOk() (*DatadogForwardingConfig, bool) {
	if o == nil || o.DatadogConfig == nil {
		return nil, false
	}
	return o.DatadogConfig, true
}

// HasDatadogConfig returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasDatadogConfig() bool {
	if o != nil && o.DatadogConfig != nil {
		return true
	}

	return false
}

// SetDatadogConfig gets a reference to the given DatadogForwardingConfig and assigns it to the DatadogConfig field.
func (o *LogForwardingConfigPart) SetDatadogConfig(v DatadogForwardingConfig) {
	o.DatadogConfig = &v
}

// GetElasticsearchConfig returns the ElasticsearchConfig field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetElasticsearchConfig() ElasticsearchLogForwardingConfig {
	if o == nil || o.ElasticsearchConfig == nil {
		var ret ElasticsearchLogForwardingConfig
		return ret
	}
	return *o.ElasticsearchConfig
}

// GetElasticsearchConfigOk returns a tuple with the ElasticsearchConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetElasticsearchConfigOk() (*ElasticsearchLogForwardingConfig, bool) {
	if o == nil || o.ElasticsearchConfig == nil {
		return nil, false
	}
	return o.ElasticsearchConfig, true
}

// HasElasticsearchConfig returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasElasticsearchConfig() bool {
	if o != nil && o.ElasticsearchConfig != nil {
		return true
	}

	return false
}

// SetElasticsearchConfig gets a reference to the given ElasticsearchLogForwardingConfig and assigns it to the ElasticsearchConfig field.
func (o *LogForwardingConfigPart) SetElasticsearchConfig(v ElasticsearchLogForwardingConfig) {
	o.ElasticsearchConfig = &v
}

// GetJsonOutput returns the JsonOutput field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetJsonOutput() bool {
	if o == nil || o.JsonOutput == nil {
		var ret bool
		return ret
	}
	return *o.JsonOutput
}

// GetJsonOutputOk returns a tuple with the JsonOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetJsonOutputOk() (*bool, bool) {
	if o == nil || o.JsonOutput == nil {
		return nil, false
	}
	return o.JsonOutput, true
}

// HasJsonOutput returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasJsonOutput() bool {
	if o != nil && o.JsonOutput != nil {
		return true
	}

	return false
}

// SetJsonOutput gets a reference to the given bool and assigns it to the JsonOutput field.
func (o *LogForwardingConfigPart) SetJsonOutput(v bool) {
	o.JsonOutput = &v
}

// GetLoganEnable returns the LoganEnable field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetLoganEnable() bool {
	if o == nil || o.LoganEnable == nil {
		var ret bool
		return ret
	}
	return *o.LoganEnable
}

// GetLoganEnableOk returns a tuple with the LoganEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetLoganEnableOk() (*bool, bool) {
	if o == nil || o.LoganEnable == nil {
		return nil, false
	}
	return o.LoganEnable, true
}

// HasLoganEnable returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasLoganEnable() bool {
	if o != nil && o.LoganEnable != nil {
		return true
	}

	return false
}

// SetLoganEnable gets a reference to the given bool and assigns it to the LoganEnable field.
func (o *LogForwardingConfigPart) SetLoganEnable(v bool) {
	o.LoganEnable = &v
}

// GetLoganUrl returns the LoganUrl field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetLoganUrl() string {
	if o == nil || o.LoganUrl == nil {
		var ret string
		return ret
	}
	return *o.LoganUrl
}

// GetLoganUrlOk returns a tuple with the LoganUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetLoganUrlOk() (*string, bool) {
	if o == nil || o.LoganUrl == nil {
		return nil, false
	}
	return o.LoganUrl, true
}

// HasLoganUrl returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasLoganUrl() bool {
	if o != nil && o.LoganUrl != nil {
		return true
	}

	return false
}

// SetLoganUrl gets a reference to the given string and assigns it to the LoganUrl field.
func (o *LogForwardingConfigPart) SetLoganUrl(v string) {
	o.LoganUrl = &v
}

// GetLogstashConfig returns the LogstashConfig field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetLogstashConfig() LogstashLogForwardingConfig {
	if o == nil || o.LogstashConfig == nil {
		var ret LogstashLogForwardingConfig
		return ret
	}
	return *o.LogstashConfig
}

// GetLogstashConfigOk returns a tuple with the LogstashConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetLogstashConfigOk() (*LogstashLogForwardingConfig, bool) {
	if o == nil || o.LogstashConfig == nil {
		return nil, false
	}
	return o.LogstashConfig, true
}

// HasLogstashConfig returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasLogstashConfig() bool {
	if o != nil && o.LogstashConfig != nil {
		return true
	}

	return false
}

// SetLogstashConfig gets a reference to the given LogstashLogForwardingConfig and assigns it to the LogstashConfig field.
func (o *LogForwardingConfigPart) SetLogstashConfig(v LogstashLogForwardingConfig) {
	o.LogstashConfig = &v
}

// GetLogzIoConfig returns the LogzIoConfig field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetLogzIoConfig() LogzIoLogForwardingConfig {
	if o == nil || o.LogzIoConfig == nil {
		var ret LogzIoLogForwardingConfig
		return ret
	}
	return *o.LogzIoConfig
}

// GetLogzIoConfigOk returns a tuple with the LogzIoConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetLogzIoConfigOk() (*LogzIoLogForwardingConfig, bool) {
	if o == nil || o.LogzIoConfig == nil {
		return nil, false
	}
	return o.LogzIoConfig, true
}

// HasLogzIoConfig returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasLogzIoConfig() bool {
	if o != nil && o.LogzIoConfig != nil {
		return true
	}

	return false
}

// SetLogzIoConfig gets a reference to the given LogzIoLogForwardingConfig and assigns it to the LogzIoConfig field.
func (o *LogForwardingConfigPart) SetLogzIoConfig(v LogzIoLogForwardingConfig) {
	o.LogzIoConfig = &v
}

// GetPullIntervalSec returns the PullIntervalSec field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetPullIntervalSec() string {
	if o == nil || o.PullIntervalSec == nil {
		var ret string
		return ret
	}
	return *o.PullIntervalSec
}

// GetPullIntervalSecOk returns a tuple with the PullIntervalSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetPullIntervalSecOk() (*string, bool) {
	if o == nil || o.PullIntervalSec == nil {
		return nil, false
	}
	return o.PullIntervalSec, true
}

// HasPullIntervalSec returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasPullIntervalSec() bool {
	if o != nil && o.PullIntervalSec != nil {
		return true
	}

	return false
}

// SetPullIntervalSec gets a reference to the given string and assigns it to the PullIntervalSec field.
func (o *LogForwardingConfigPart) SetPullIntervalSec(v string) {
	o.PullIntervalSec = &v
}

// GetSplunkConfig returns the SplunkConfig field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetSplunkConfig() SplunkLogForwardingConfig {
	if o == nil || o.SplunkConfig == nil {
		var ret SplunkLogForwardingConfig
		return ret
	}
	return *o.SplunkConfig
}

// GetSplunkConfigOk returns a tuple with the SplunkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetSplunkConfigOk() (*SplunkLogForwardingConfig, bool) {
	if o == nil || o.SplunkConfig == nil {
		return nil, false
	}
	return o.SplunkConfig, true
}

// HasSplunkConfig returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasSplunkConfig() bool {
	if o != nil && o.SplunkConfig != nil {
		return true
	}

	return false
}

// SetSplunkConfig gets a reference to the given SplunkLogForwardingConfig and assigns it to the SplunkConfig field.
func (o *LogForwardingConfigPart) SetSplunkConfig(v SplunkLogForwardingConfig) {
	o.SplunkConfig = &v
}

// GetSyslogConfig returns the SyslogConfig field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetSyslogConfig() SyslogLogForwardingConfig {
	if o == nil || o.SyslogConfig == nil {
		var ret SyslogLogForwardingConfig
		return ret
	}
	return *o.SyslogConfig
}

// GetSyslogConfigOk returns a tuple with the SyslogConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetSyslogConfigOk() (*SyslogLogForwardingConfig, bool) {
	if o == nil || o.SyslogConfig == nil {
		return nil, false
	}
	return o.SyslogConfig, true
}

// HasSyslogConfig returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasSyslogConfig() bool {
	if o != nil && o.SyslogConfig != nil {
		return true
	}

	return false
}

// SetSyslogConfig gets a reference to the given SyslogLogForwardingConfig and assigns it to the SyslogConfig field.
func (o *LogForwardingConfigPart) SetSyslogConfig(v SyslogLogForwardingConfig) {
	o.SyslogConfig = &v
}

// GetTargetLogType returns the TargetLogType field value if set, zero value otherwise.
func (o *LogForwardingConfigPart) GetTargetLogType() string {
	if o == nil || o.TargetLogType == nil {
		var ret string
		return ret
	}
	return *o.TargetLogType
}

// GetTargetLogTypeOk returns a tuple with the TargetLogType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwardingConfigPart) GetTargetLogTypeOk() (*string, bool) {
	if o == nil || o.TargetLogType == nil {
		return nil, false
	}
	return o.TargetLogType, true
}

// HasTargetLogType returns a boolean if a field has been set.
func (o *LogForwardingConfigPart) HasTargetLogType() bool {
	if o != nil && o.TargetLogType != nil {
		return true
	}

	return false
}

// SetTargetLogType gets a reference to the given string and assigns it to the TargetLogType field.
func (o *LogForwardingConfigPart) SetTargetLogType(v string) {
	o.TargetLogType = &v
}

func (o LogForwardingConfigPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsS3Config != nil {
		toSerialize["aws_s3_config"] = o.AwsS3Config
	}
	if o.AzureAnalyticsConfig != nil {
		toSerialize["azure_analytics_config"] = o.AzureAnalyticsConfig
	}
	if o.DatadogConfig != nil {
		toSerialize["datadog_config"] = o.DatadogConfig
	}
	if o.ElasticsearchConfig != nil {
		toSerialize["elasticsearch_config"] = o.ElasticsearchConfig
	}
	if o.JsonOutput != nil {
		toSerialize["json_output"] = o.JsonOutput
	}
	if o.LoganEnable != nil {
		toSerialize["logan_enable"] = o.LoganEnable
	}
	if o.LoganUrl != nil {
		toSerialize["logan_url"] = o.LoganUrl
	}
	if o.LogstashConfig != nil {
		toSerialize["logstash_config"] = o.LogstashConfig
	}
	if o.LogzIoConfig != nil {
		toSerialize["logz_io_config"] = o.LogzIoConfig
	}
	if o.PullIntervalSec != nil {
		toSerialize["pull_interval_sec"] = o.PullIntervalSec
	}
	if o.SplunkConfig != nil {
		toSerialize["splunk_config"] = o.SplunkConfig
	}
	if o.SyslogConfig != nil {
		toSerialize["syslog_config"] = o.SyslogConfig
	}
	if o.TargetLogType != nil {
		toSerialize["target_log_type"] = o.TargetLogType
	}
	return json.Marshal(toSerialize)
}

type NullableLogForwardingConfigPart struct {
	value *LogForwardingConfigPart
	isSet bool
}

func (v NullableLogForwardingConfigPart) Get() *LogForwardingConfigPart {
	return v.value
}

func (v *NullableLogForwardingConfigPart) Set(val *LogForwardingConfigPart) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwardingConfigPart) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwardingConfigPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwardingConfigPart(val *LogForwardingConfigPart) *NullableLogForwardingConfigPart {
	return &NullableLogForwardingConfigPart{value: val, isSet: true}
}

func (v NullableLogForwardingConfigPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwardingConfigPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


