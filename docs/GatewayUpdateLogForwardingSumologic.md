# GatewayUpdateLogForwardingSumologic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**Endpoint** | Pointer to **string** | Sumologic endpoint URL | [optional] 
**Host** | Pointer to **string** | Sumologic host | [optional] [default to "use-existing"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**SumologicTags** | Pointer to **string** | A comma-separated list of Sumologic tags | [optional] [default to "use-existing"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateLogForwardingSumologic

`func NewGatewayUpdateLogForwardingSumologic() *GatewayUpdateLogForwardingSumologic`

NewGatewayUpdateLogForwardingSumologic instantiates a new GatewayUpdateLogForwardingSumologic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLogForwardingSumologicWithDefaults

`func NewGatewayUpdateLogForwardingSumologicWithDefaults() *GatewayUpdateLogForwardingSumologic`

NewGatewayUpdateLogForwardingSumologicWithDefaults instantiates a new GatewayUpdateLogForwardingSumologic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GatewayUpdateLogForwardingSumologic) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayUpdateLogForwardingSumologic) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayUpdateLogForwardingSumologic) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayUpdateLogForwardingSumologic) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEndpoint

`func (o *GatewayUpdateLogForwardingSumologic) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GatewayUpdateLogForwardingSumologic) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GatewayUpdateLogForwardingSumologic) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *GatewayUpdateLogForwardingSumologic) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHost

`func (o *GatewayUpdateLogForwardingSumologic) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GatewayUpdateLogForwardingSumologic) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GatewayUpdateLogForwardingSumologic) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GatewayUpdateLogForwardingSumologic) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateLogForwardingSumologic) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateLogForwardingSumologic) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateLogForwardingSumologic) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateLogForwardingSumologic) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GatewayUpdateLogForwardingSumologic) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GatewayUpdateLogForwardingSumologic) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GatewayUpdateLogForwardingSumologic) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GatewayUpdateLogForwardingSumologic) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GatewayUpdateLogForwardingSumologic) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GatewayUpdateLogForwardingSumologic) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GatewayUpdateLogForwardingSumologic) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GatewayUpdateLogForwardingSumologic) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetSumologicTags

`func (o *GatewayUpdateLogForwardingSumologic) GetSumologicTags() string`

GetSumologicTags returns the SumologicTags field if non-nil, zero value otherwise.

### GetSumologicTagsOk

`func (o *GatewayUpdateLogForwardingSumologic) GetSumologicTagsOk() (*string, bool)`

GetSumologicTagsOk returns a tuple with the SumologicTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumologicTags

`func (o *GatewayUpdateLogForwardingSumologic) SetSumologicTags(v string)`

SetSumologicTags sets SumologicTags field to given value.

### HasSumologicTags

`func (o *GatewayUpdateLogForwardingSumologic) HasSumologicTags() bool`

HasSumologicTags returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLogForwardingSumologic) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLogForwardingSumologic) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLogForwardingSumologic) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLogForwardingSumologic) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLogForwardingSumologic) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLogForwardingSumologic) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLogForwardingSumologic) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLogForwardingSumologic) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


