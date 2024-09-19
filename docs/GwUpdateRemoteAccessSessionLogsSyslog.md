# GwUpdateRemoteAccessSessionLogsSyslog

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

### NewGwUpdateRemoteAccessSessionLogsSyslog

`func NewGwUpdateRemoteAccessSessionLogsSyslog() *GwUpdateRemoteAccessSessionLogsSyslog`

NewGwUpdateRemoteAccessSessionLogsSyslog instantiates a new GwUpdateRemoteAccessSessionLogsSyslog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsSyslogWithDefaults

`func NewGwUpdateRemoteAccessSessionLogsSyslogWithDefaults() *GwUpdateRemoteAccessSessionLogsSyslog`

NewGwUpdateRemoteAccessSessionLogsSyslogWithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsSyslog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetFormatter

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetFormatter() string`

GetFormatter returns the Formatter field if non-nil, zero value otherwise.

### GetFormatterOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetFormatterOk() (*string, bool)`

GetFormatterOk returns a tuple with the Formatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatter

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetFormatter(v string)`

SetFormatter sets Formatter field to given value.

### HasFormatter

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasFormatter() bool`

HasFormatter returns a boolean if a field has been set.

### GetHost

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetNetwork

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetTargetTag

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetTargetTag() string`

GetTargetTag returns the TargetTag field if non-nil, zero value otherwise.

### GetTargetTagOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetTargetTagOk() (*string, bool)`

GetTargetTagOk returns a tuple with the TargetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTag

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetTargetTag(v string)`

SetTargetTag sets TargetTag field to given value.

### HasTargetTag

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasTargetTag() bool`

HasTargetTag returns a boolean if a field has been set.

### GetTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsSyslog) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


