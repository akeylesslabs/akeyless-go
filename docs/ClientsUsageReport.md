# ClientsUsageReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Clients** | Pointer to [**[]ClientUsageInfo**](ClientUsageInfo.md) |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Time** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**TotalClients** | Pointer to **int64** |  | [optional] 

## Methods

### NewClientsUsageReport

`func NewClientsUsageReport() *ClientsUsageReport`

NewClientsUsageReport instantiates a new ClientsUsageReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientsUsageReportWithDefaults

`func NewClientsUsageReportWithDefaults() *ClientsUsageReport`

NewClientsUsageReportWithDefaults instantiates a new ClientsUsageReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ClientsUsageReport) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ClientsUsageReport) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ClientsUsageReport) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ClientsUsageReport) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetClients

`func (o *ClientsUsageReport) GetClients() []ClientUsageInfo`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ClientsUsageReport) GetClientsOk() (*[]ClientUsageInfo, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ClientsUsageReport) SetClients(v []ClientUsageInfo)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ClientsUsageReport) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetProduct

`func (o *ClientsUsageReport) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ClientsUsageReport) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ClientsUsageReport) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ClientsUsageReport) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTime

`func (o *ClientsUsageReport) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ClientsUsageReport) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ClientsUsageReport) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ClientsUsageReport) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTotalClients

`func (o *ClientsUsageReport) GetTotalClients() int64`

GetTotalClients returns the TotalClients field if non-nil, zero value otherwise.

### GetTotalClientsOk

`func (o *ClientsUsageReport) GetTotalClientsOk() (*int64, bool)`

GetTotalClientsOk returns a tuple with the TotalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClients

`func (o *ClientsUsageReport) SetTotalClients(v int64)`

SetTotalClients sets TotalClients field to given value.

### HasTotalClients

`func (o *ClientsUsageReport) HasTotalClients() bool`

HasTotalClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


