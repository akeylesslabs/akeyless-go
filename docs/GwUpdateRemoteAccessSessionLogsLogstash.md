# GwUpdateRemoteAccessSessionLogsLogstash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **string** | Logstash dns | [optional] 
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**EnableTls** | Pointer to **bool** | Enable tls | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**Protocol** | Pointer to **string** | Logstash protocol [tcp/udp] | [optional] 
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**TlsCertificate** | Pointer to **string** | Logstash tls certificate | [optional] [default to "use-existing"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGwUpdateRemoteAccessSessionLogsLogstash

`func NewGwUpdateRemoteAccessSessionLogsLogstash() *GwUpdateRemoteAccessSessionLogsLogstash`

NewGwUpdateRemoteAccessSessionLogsLogstash instantiates a new GwUpdateRemoteAccessSessionLogsLogstash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsLogstashWithDefaults

`func NewGwUpdateRemoteAccessSessionLogsLogstashWithDefaults() *GwUpdateRemoteAccessSessionLogsLogstash`

NewGwUpdateRemoteAccessSessionLogsLogstashWithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsLogstash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetProtocol

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsLogstash) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


