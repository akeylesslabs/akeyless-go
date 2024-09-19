# GwUpdateRemoteAccessSessionLogsSumologic

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

### NewGwUpdateRemoteAccessSessionLogsSumologic

`func NewGwUpdateRemoteAccessSessionLogsSumologic() *GwUpdateRemoteAccessSessionLogsSumologic`

NewGwUpdateRemoteAccessSessionLogsSumologic instantiates a new GwUpdateRemoteAccessSessionLogsSumologic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsSumologicWithDefaults

`func NewGwUpdateRemoteAccessSessionLogsSumologicWithDefaults() *GwUpdateRemoteAccessSessionLogsSumologic`

NewGwUpdateRemoteAccessSessionLogsSumologicWithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsSumologic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEndpoint

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHost

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetSumologicTags

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetSumologicTags() string`

GetSumologicTags returns the SumologicTags field if non-nil, zero value otherwise.

### GetSumologicTagsOk

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetSumologicTagsOk() (*string, bool)`

GetSumologicTagsOk returns a tuple with the SumologicTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumologicTags

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) SetSumologicTags(v string)`

SetSumologicTags sets SumologicTags field to given value.

### HasSumologicTags

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) HasSumologicTags() bool`

HasSumologicTags returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsSumologic) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


