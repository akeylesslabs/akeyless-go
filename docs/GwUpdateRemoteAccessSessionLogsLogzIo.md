# GwUpdateRemoteAccessSessionLogsLogzIo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**LogzIoToken** | Pointer to **string** | Logz-io token | [optional] 
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**Protocol** | Pointer to **string** | LogzIo protocol [tcp/https] | [optional] 
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGwUpdateRemoteAccessSessionLogsLogzIo

`func NewGwUpdateRemoteAccessSessionLogsLogzIo() *GwUpdateRemoteAccessSessionLogsLogzIo`

NewGwUpdateRemoteAccessSessionLogsLogzIo instantiates a new GwUpdateRemoteAccessSessionLogsLogzIo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsLogzIoWithDefaults

`func NewGwUpdateRemoteAccessSessionLogsLogzIoWithDefaults() *GwUpdateRemoteAccessSessionLogsLogzIo`

NewGwUpdateRemoteAccessSessionLogsLogzIoWithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsLogzIo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLogzIoToken

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetLogzIoToken() string`

GetLogzIoToken returns the LogzIoToken field if non-nil, zero value otherwise.

### GetLogzIoTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetLogzIoTokenOk() (*string, bool)`

GetLogzIoTokenOk returns a tuple with the LogzIoToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogzIoToken

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) SetLogzIoToken(v string)`

SetLogzIoToken sets LogzIoToken field to given value.

### HasLogzIoToken

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) HasLogzIoToken() bool`

HasLogzIoToken returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetProtocol

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsLogzIo) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


