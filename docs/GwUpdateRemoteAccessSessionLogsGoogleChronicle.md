# GwUpdateRemoteAccessSessionLogsGoogleChronicle

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

### NewGwUpdateRemoteAccessSessionLogsGoogleChronicle

`func NewGwUpdateRemoteAccessSessionLogsGoogleChronicle() *GwUpdateRemoteAccessSessionLogsGoogleChronicle`

NewGwUpdateRemoteAccessSessionLogsGoogleChronicle instantiates a new GwUpdateRemoteAccessSessionLogsGoogleChronicle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsGoogleChronicleWithDefaults

`func NewGwUpdateRemoteAccessSessionLogsGoogleChronicleWithDefaults() *GwUpdateRemoteAccessSessionLogsGoogleChronicle`

NewGwUpdateRemoteAccessSessionLogsGoogleChronicleWithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsGoogleChronicle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGcpKey

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLogType

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetRegion

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsGoogleChronicle) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


