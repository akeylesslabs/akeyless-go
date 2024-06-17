# GatewayUpdateLogForwardingSplunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**EnableTls** | Pointer to **bool** | Enable tls | [optional] 
**Index** | Pointer to **string** | Splunk index | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**Source** | Pointer to **string** | Splunk source | [optional] [default to "use-existing"]
**SourceType** | Pointer to **string** | Splunk source type | [optional] [default to "use-existing"]
**SplunkToken** | Pointer to **string** | Splunk token | [optional] 
**SplunkUrl** | Pointer to **string** | Splunk server URL | [optional] 
**TlsCertificate** | Pointer to **string** | Splunk tls certificate | [optional] [default to "use-existing"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateLogForwardingSplunk

`func NewGatewayUpdateLogForwardingSplunk() *GatewayUpdateLogForwardingSplunk`

NewGatewayUpdateLogForwardingSplunk instantiates a new GatewayUpdateLogForwardingSplunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLogForwardingSplunkWithDefaults

`func NewGatewayUpdateLogForwardingSplunkWithDefaults() *GatewayUpdateLogForwardingSplunk`

NewGatewayUpdateLogForwardingSplunkWithDefaults instantiates a new GatewayUpdateLogForwardingSplunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GatewayUpdateLogForwardingSplunk) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayUpdateLogForwardingSplunk) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayUpdateLogForwardingSplunk) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayUpdateLogForwardingSplunk) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTls

`func (o *GatewayUpdateLogForwardingSplunk) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *GatewayUpdateLogForwardingSplunk) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *GatewayUpdateLogForwardingSplunk) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *GatewayUpdateLogForwardingSplunk) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetIndex

`func (o *GatewayUpdateLogForwardingSplunk) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GatewayUpdateLogForwardingSplunk) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GatewayUpdateLogForwardingSplunk) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *GatewayUpdateLogForwardingSplunk) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateLogForwardingSplunk) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateLogForwardingSplunk) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateLogForwardingSplunk) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateLogForwardingSplunk) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GatewayUpdateLogForwardingSplunk) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GatewayUpdateLogForwardingSplunk) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GatewayUpdateLogForwardingSplunk) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GatewayUpdateLogForwardingSplunk) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GatewayUpdateLogForwardingSplunk) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GatewayUpdateLogForwardingSplunk) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GatewayUpdateLogForwardingSplunk) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GatewayUpdateLogForwardingSplunk) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetSource

`func (o *GatewayUpdateLogForwardingSplunk) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GatewayUpdateLogForwardingSplunk) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GatewayUpdateLogForwardingSplunk) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GatewayUpdateLogForwardingSplunk) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *GatewayUpdateLogForwardingSplunk) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GatewayUpdateLogForwardingSplunk) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GatewayUpdateLogForwardingSplunk) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GatewayUpdateLogForwardingSplunk) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSplunkToken

`func (o *GatewayUpdateLogForwardingSplunk) GetSplunkToken() string`

GetSplunkToken returns the SplunkToken field if non-nil, zero value otherwise.

### GetSplunkTokenOk

`func (o *GatewayUpdateLogForwardingSplunk) GetSplunkTokenOk() (*string, bool)`

GetSplunkTokenOk returns a tuple with the SplunkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkToken

`func (o *GatewayUpdateLogForwardingSplunk) SetSplunkToken(v string)`

SetSplunkToken sets SplunkToken field to given value.

### HasSplunkToken

`func (o *GatewayUpdateLogForwardingSplunk) HasSplunkToken() bool`

HasSplunkToken returns a boolean if a field has been set.

### GetSplunkUrl

`func (o *GatewayUpdateLogForwardingSplunk) GetSplunkUrl() string`

GetSplunkUrl returns the SplunkUrl field if non-nil, zero value otherwise.

### GetSplunkUrlOk

`func (o *GatewayUpdateLogForwardingSplunk) GetSplunkUrlOk() (*string, bool)`

GetSplunkUrlOk returns a tuple with the SplunkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkUrl

`func (o *GatewayUpdateLogForwardingSplunk) SetSplunkUrl(v string)`

SetSplunkUrl sets SplunkUrl field to given value.

### HasSplunkUrl

`func (o *GatewayUpdateLogForwardingSplunk) HasSplunkUrl() bool`

HasSplunkUrl returns a boolean if a field has been set.

### GetTlsCertificate

`func (o *GatewayUpdateLogForwardingSplunk) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *GatewayUpdateLogForwardingSplunk) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *GatewayUpdateLogForwardingSplunk) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *GatewayUpdateLogForwardingSplunk) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLogForwardingSplunk) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLogForwardingSplunk) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLogForwardingSplunk) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLogForwardingSplunk) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLogForwardingSplunk) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLogForwardingSplunk) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLogForwardingSplunk) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLogForwardingSplunk) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


