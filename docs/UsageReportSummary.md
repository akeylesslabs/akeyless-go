# UsageReportSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientsByAuthMethodTypes** | Pointer to **map[string]int64** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**SecretsByTypes** | Pointer to **map[string]int64** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TotalClients** | Pointer to **int64** |  | [optional] 
**TotalSecrets** | Pointer to **int64** |  | [optional] 

## Methods

### NewUsageReportSummary

`func NewUsageReportSummary() *UsageReportSummary`

NewUsageReportSummary instantiates a new UsageReportSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageReportSummaryWithDefaults

`func NewUsageReportSummaryWithDefaults() *UsageReportSummary`

NewUsageReportSummaryWithDefaults instantiates a new UsageReportSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientsByAuthMethodTypes

`func (o *UsageReportSummary) GetClientsByAuthMethodTypes() map[string]int64`

GetClientsByAuthMethodTypes returns the ClientsByAuthMethodTypes field if non-nil, zero value otherwise.

### GetClientsByAuthMethodTypesOk

`func (o *UsageReportSummary) GetClientsByAuthMethodTypesOk() (*map[string]int64, bool)`

GetClientsByAuthMethodTypesOk returns a tuple with the ClientsByAuthMethodTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsByAuthMethodTypes

`func (o *UsageReportSummary) SetClientsByAuthMethodTypes(v map[string]int64)`

SetClientsByAuthMethodTypes sets ClientsByAuthMethodTypes field to given value.

### HasClientsByAuthMethodTypes

`func (o *UsageReportSummary) HasClientsByAuthMethodTypes() bool`

HasClientsByAuthMethodTypes returns a boolean if a field has been set.

### GetProduct

`func (o *UsageReportSummary) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *UsageReportSummary) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *UsageReportSummary) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *UsageReportSummary) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSecretsByTypes

`func (o *UsageReportSummary) GetSecretsByTypes() map[string]int64`

GetSecretsByTypes returns the SecretsByTypes field if non-nil, zero value otherwise.

### GetSecretsByTypesOk

`func (o *UsageReportSummary) GetSecretsByTypesOk() (*map[string]int64, bool)`

GetSecretsByTypesOk returns a tuple with the SecretsByTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsByTypes

`func (o *UsageReportSummary) SetSecretsByTypes(v map[string]int64)`

SetSecretsByTypes sets SecretsByTypes field to given value.

### HasSecretsByTypes

`func (o *UsageReportSummary) HasSecretsByTypes() bool`

HasSecretsByTypes returns a boolean if a field has been set.

### GetTime

`func (o *UsageReportSummary) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UsageReportSummary) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UsageReportSummary) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *UsageReportSummary) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTotalClients

`func (o *UsageReportSummary) GetTotalClients() int64`

GetTotalClients returns the TotalClients field if non-nil, zero value otherwise.

### GetTotalClientsOk

`func (o *UsageReportSummary) GetTotalClientsOk() (*int64, bool)`

GetTotalClientsOk returns a tuple with the TotalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClients

`func (o *UsageReportSummary) SetTotalClients(v int64)`

SetTotalClients sets TotalClients field to given value.

### HasTotalClients

`func (o *UsageReportSummary) HasTotalClients() bool`

HasTotalClients returns a boolean if a field has been set.

### GetTotalSecrets

`func (o *UsageReportSummary) GetTotalSecrets() int64`

GetTotalSecrets returns the TotalSecrets field if non-nil, zero value otherwise.

### GetTotalSecretsOk

`func (o *UsageReportSummary) GetTotalSecretsOk() (*int64, bool)`

GetTotalSecretsOk returns a tuple with the TotalSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSecrets

`func (o *UsageReportSummary) SetTotalSecrets(v int64)`

SetTotalSecrets sets TotalSecrets field to given value.

### HasTotalSecrets

`func (o *UsageReportSummary) HasTotalSecrets() bool`

HasTotalSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


