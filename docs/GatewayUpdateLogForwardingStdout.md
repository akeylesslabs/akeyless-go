# GatewayUpdateLogForwardingStdout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateLogForwardingStdout

`func NewGatewayUpdateLogForwardingStdout() *GatewayUpdateLogForwardingStdout`

NewGatewayUpdateLogForwardingStdout instantiates a new GatewayUpdateLogForwardingStdout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLogForwardingStdoutWithDefaults

`func NewGatewayUpdateLogForwardingStdoutWithDefaults() *GatewayUpdateLogForwardingStdout`

NewGatewayUpdateLogForwardingStdoutWithDefaults instantiates a new GatewayUpdateLogForwardingStdout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GatewayUpdateLogForwardingStdout) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayUpdateLogForwardingStdout) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayUpdateLogForwardingStdout) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayUpdateLogForwardingStdout) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateLogForwardingStdout) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateLogForwardingStdout) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateLogForwardingStdout) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateLogForwardingStdout) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GatewayUpdateLogForwardingStdout) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GatewayUpdateLogForwardingStdout) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GatewayUpdateLogForwardingStdout) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GatewayUpdateLogForwardingStdout) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GatewayUpdateLogForwardingStdout) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GatewayUpdateLogForwardingStdout) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GatewayUpdateLogForwardingStdout) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GatewayUpdateLogForwardingStdout) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLogForwardingStdout) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLogForwardingStdout) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLogForwardingStdout) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLogForwardingStdout) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLogForwardingStdout) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLogForwardingStdout) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLogForwardingStdout) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLogForwardingStdout) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


