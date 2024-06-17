# GatewayUpdateLogForwardingLogstash

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

### NewGatewayUpdateLogForwardingLogstash

`func NewGatewayUpdateLogForwardingLogstash() *GatewayUpdateLogForwardingLogstash`

NewGatewayUpdateLogForwardingLogstash instantiates a new GatewayUpdateLogForwardingLogstash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLogForwardingLogstashWithDefaults

`func NewGatewayUpdateLogForwardingLogstashWithDefaults() *GatewayUpdateLogForwardingLogstash`

NewGatewayUpdateLogForwardingLogstashWithDefaults instantiates a new GatewayUpdateLogForwardingLogstash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *GatewayUpdateLogForwardingLogstash) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *GatewayUpdateLogForwardingLogstash) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *GatewayUpdateLogForwardingLogstash) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *GatewayUpdateLogForwardingLogstash) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetEnable

`func (o *GatewayUpdateLogForwardingLogstash) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayUpdateLogForwardingLogstash) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayUpdateLogForwardingLogstash) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayUpdateLogForwardingLogstash) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTls

`func (o *GatewayUpdateLogForwardingLogstash) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *GatewayUpdateLogForwardingLogstash) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *GatewayUpdateLogForwardingLogstash) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *GatewayUpdateLogForwardingLogstash) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateLogForwardingLogstash) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateLogForwardingLogstash) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateLogForwardingLogstash) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateLogForwardingLogstash) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GatewayUpdateLogForwardingLogstash) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GatewayUpdateLogForwardingLogstash) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GatewayUpdateLogForwardingLogstash) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GatewayUpdateLogForwardingLogstash) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetProtocol

`func (o *GatewayUpdateLogForwardingLogstash) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayUpdateLogForwardingLogstash) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayUpdateLogForwardingLogstash) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GatewayUpdateLogForwardingLogstash) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPullInterval

`func (o *GatewayUpdateLogForwardingLogstash) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GatewayUpdateLogForwardingLogstash) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GatewayUpdateLogForwardingLogstash) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GatewayUpdateLogForwardingLogstash) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetTlsCertificate

`func (o *GatewayUpdateLogForwardingLogstash) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *GatewayUpdateLogForwardingLogstash) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *GatewayUpdateLogForwardingLogstash) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *GatewayUpdateLogForwardingLogstash) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLogForwardingLogstash) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLogForwardingLogstash) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLogForwardingLogstash) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLogForwardingLogstash) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLogForwardingLogstash) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLogForwardingLogstash) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLogForwardingLogstash) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLogForwardingLogstash) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


