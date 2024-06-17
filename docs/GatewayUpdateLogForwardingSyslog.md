# GatewayUpdateLogForwardingSyslog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**EnableTls** | Pointer to **bool** | Enable tls relevant only for network type TCP | [optional] 
**Formatter** | Pointer to **string** | Syslog formatter [text/cef] | [optional] [default to "text"]
**Host** | Pointer to **string** | Syslog host | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Network** | Pointer to **string** | Syslog network [tcp/udp] | [optional] [default to "tcp"]
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**TargetTag** | Pointer to **string** | Syslog target tag | [optional] [default to "use-existing"]
**TlsCertificate** | Pointer to **string** | Syslog tls certificate | [optional] [default to "use-existing"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateLogForwardingSyslog

`func NewGatewayUpdateLogForwardingSyslog() *GatewayUpdateLogForwardingSyslog`

NewGatewayUpdateLogForwardingSyslog instantiates a new GatewayUpdateLogForwardingSyslog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLogForwardingSyslogWithDefaults

`func NewGatewayUpdateLogForwardingSyslogWithDefaults() *GatewayUpdateLogForwardingSyslog`

NewGatewayUpdateLogForwardingSyslogWithDefaults instantiates a new GatewayUpdateLogForwardingSyslog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GatewayUpdateLogForwardingSyslog) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayUpdateLogForwardingSyslog) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayUpdateLogForwardingSyslog) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayUpdateLogForwardingSyslog) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTls

`func (o *GatewayUpdateLogForwardingSyslog) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *GatewayUpdateLogForwardingSyslog) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *GatewayUpdateLogForwardingSyslog) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *GatewayUpdateLogForwardingSyslog) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetFormatter

`func (o *GatewayUpdateLogForwardingSyslog) GetFormatter() string`

GetFormatter returns the Formatter field if non-nil, zero value otherwise.

### GetFormatterOk

`func (o *GatewayUpdateLogForwardingSyslog) GetFormatterOk() (*string, bool)`

GetFormatterOk returns a tuple with the Formatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatter

`func (o *GatewayUpdateLogForwardingSyslog) SetFormatter(v string)`

SetFormatter sets Formatter field to given value.

### HasFormatter

`func (o *GatewayUpdateLogForwardingSyslog) HasFormatter() bool`

HasFormatter returns a boolean if a field has been set.

### GetHost

`func (o *GatewayUpdateLogForwardingSyslog) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GatewayUpdateLogForwardingSyslog) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GatewayUpdateLogForwardingSyslog) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GatewayUpdateLogForwardingSyslog) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateLogForwardingSyslog) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateLogForwardingSyslog) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateLogForwardingSyslog) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateLogForwardingSyslog) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetNetwork

`func (o *GatewayUpdateLogForwardingSyslog) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayUpdateLogForwardingSyslog) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayUpdateLogForwardingSyslog) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GatewayUpdateLogForwardingSyslog) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GatewayUpdateLogForwardingSyslog) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GatewayUpdateLogForwardingSyslog) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GatewayUpdateLogForwardingSyslog) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GatewayUpdateLogForwardingSyslog) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GatewayUpdateLogForwardingSyslog) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GatewayUpdateLogForwardingSyslog) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GatewayUpdateLogForwardingSyslog) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GatewayUpdateLogForwardingSyslog) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetTargetTag

`func (o *GatewayUpdateLogForwardingSyslog) GetTargetTag() string`

GetTargetTag returns the TargetTag field if non-nil, zero value otherwise.

### GetTargetTagOk

`func (o *GatewayUpdateLogForwardingSyslog) GetTargetTagOk() (*string, bool)`

GetTargetTagOk returns a tuple with the TargetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTag

`func (o *GatewayUpdateLogForwardingSyslog) SetTargetTag(v string)`

SetTargetTag sets TargetTag field to given value.

### HasTargetTag

`func (o *GatewayUpdateLogForwardingSyslog) HasTargetTag() bool`

HasTargetTag returns a boolean if a field has been set.

### GetTlsCertificate

`func (o *GatewayUpdateLogForwardingSyslog) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *GatewayUpdateLogForwardingSyslog) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *GatewayUpdateLogForwardingSyslog) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *GatewayUpdateLogForwardingSyslog) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLogForwardingSyslog) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLogForwardingSyslog) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLogForwardingSyslog) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLogForwardingSyslog) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLogForwardingSyslog) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLogForwardingSyslog) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLogForwardingSyslog) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLogForwardingSyslog) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


