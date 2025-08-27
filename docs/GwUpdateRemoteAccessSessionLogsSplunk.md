# GwUpdateRemoteAccessSessionLogsSplunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**EnableBatch** | Pointer to **string** | Enable batch forwarding [true/false] | [optional] [default to "true"]
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

### NewGwUpdateRemoteAccessSessionLogsSplunk

`func NewGwUpdateRemoteAccessSessionLogsSplunk() *GwUpdateRemoteAccessSessionLogsSplunk`

NewGwUpdateRemoteAccessSessionLogsSplunk instantiates a new GwUpdateRemoteAccessSessionLogsSplunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsSplunkWithDefaults

`func NewGwUpdateRemoteAccessSessionLogsSplunkWithDefaults() *GwUpdateRemoteAccessSessionLogsSplunk`

NewGwUpdateRemoteAccessSessionLogsSplunkWithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsSplunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableBatch

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetEnableBatch() string`

GetEnableBatch returns the EnableBatch field if non-nil, zero value otherwise.

### GetEnableBatchOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetEnableBatchOk() (*string, bool)`

GetEnableBatchOk returns a tuple with the EnableBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBatch

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetEnableBatch(v string)`

SetEnableBatch sets EnableBatch field to given value.

### HasEnableBatch

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasEnableBatch() bool`

HasEnableBatch returns a boolean if a field has been set.

### GetEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetIndex

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetSource

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSplunkToken

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetSplunkToken() string`

GetSplunkToken returns the SplunkToken field if non-nil, zero value otherwise.

### GetSplunkTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetSplunkTokenOk() (*string, bool)`

GetSplunkTokenOk returns a tuple with the SplunkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkToken

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetSplunkToken(v string)`

SetSplunkToken sets SplunkToken field to given value.

### HasSplunkToken

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasSplunkToken() bool`

HasSplunkToken returns a boolean if a field has been set.

### GetSplunkUrl

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetSplunkUrl() string`

GetSplunkUrl returns the SplunkUrl field if non-nil, zero value otherwise.

### GetSplunkUrlOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetSplunkUrlOk() (*string, bool)`

GetSplunkUrlOk returns a tuple with the SplunkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkUrl

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetSplunkUrl(v string)`

SetSplunkUrl sets SplunkUrl field to given value.

### HasSplunkUrl

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasSplunkUrl() bool`

HasSplunkUrl returns a boolean if a field has been set.

### GetTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsSplunk) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


