# GatewayUpdateLogForwardingGoogleChronicle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | Google chronicle customer id | [optional] 
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**GcpKey** | Pointer to **string** | Base64-encoded service account private key text | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**LogType** | Pointer to **string** | Google chronicle log type | [optional] 
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**Region** | Pointer to **string** | Google chronicle region [eu_multi_region/london/us_multi_region/singapore/tel_aviv] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateLogForwardingGoogleChronicle

`func NewGatewayUpdateLogForwardingGoogleChronicle() *GatewayUpdateLogForwardingGoogleChronicle`

NewGatewayUpdateLogForwardingGoogleChronicle instantiates a new GatewayUpdateLogForwardingGoogleChronicle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLogForwardingGoogleChronicleWithDefaults

`func NewGatewayUpdateLogForwardingGoogleChronicleWithDefaults() *GatewayUpdateLogForwardingGoogleChronicle`

NewGatewayUpdateLogForwardingGoogleChronicleWithDefaults instantiates a new GatewayUpdateLogForwardingGoogleChronicle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEnable

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGcpKey

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLogType

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetRegion

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLogForwardingGoogleChronicle) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLogForwardingGoogleChronicle) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLogForwardingGoogleChronicle) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


