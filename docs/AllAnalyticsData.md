# AllAnalyticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsData** | Pointer to [**map[string][][]string**](array.md) |  | [optional] 
**CertificatesExpiryData** | Pointer to [**CertificateAnalyticAggregation**](CertificateAnalyticAggregation.md) |  | [optional] 
**ClientsUsageReports** | Pointer to [**map[string]ClientsUsageReport**](ClientsUsageReport.md) |  | [optional] 
**DateUpdated** | Pointer to **int64** |  | [optional] 
**UsageReports** | Pointer to [**map[string]UsageReportSummary**](UsageReportSummary.md) |  | [optional] 

## Methods

### NewAllAnalyticsData

`func NewAllAnalyticsData() *AllAnalyticsData`

NewAllAnalyticsData instantiates a new AllAnalyticsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllAnalyticsDataWithDefaults

`func NewAllAnalyticsDataWithDefaults() *AllAnalyticsData`

NewAllAnalyticsDataWithDefaults instantiates a new AllAnalyticsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsData

`func (o *AllAnalyticsData) GetAnalyticsData() map[string][][]string`

GetAnalyticsData returns the AnalyticsData field if non-nil, zero value otherwise.

### GetAnalyticsDataOk

`func (o *AllAnalyticsData) GetAnalyticsDataOk() (*map[string][][]string, bool)`

GetAnalyticsDataOk returns a tuple with the AnalyticsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsData

`func (o *AllAnalyticsData) SetAnalyticsData(v map[string][][]string)`

SetAnalyticsData sets AnalyticsData field to given value.

### HasAnalyticsData

`func (o *AllAnalyticsData) HasAnalyticsData() bool`

HasAnalyticsData returns a boolean if a field has been set.

### GetCertificatesExpiryData

`func (o *AllAnalyticsData) GetCertificatesExpiryData() CertificateAnalyticAggregation`

GetCertificatesExpiryData returns the CertificatesExpiryData field if non-nil, zero value otherwise.

### GetCertificatesExpiryDataOk

`func (o *AllAnalyticsData) GetCertificatesExpiryDataOk() (*CertificateAnalyticAggregation, bool)`

GetCertificatesExpiryDataOk returns a tuple with the CertificatesExpiryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatesExpiryData

`func (o *AllAnalyticsData) SetCertificatesExpiryData(v CertificateAnalyticAggregation)`

SetCertificatesExpiryData sets CertificatesExpiryData field to given value.

### HasCertificatesExpiryData

`func (o *AllAnalyticsData) HasCertificatesExpiryData() bool`

HasCertificatesExpiryData returns a boolean if a field has been set.

### GetClientsUsageReports

`func (o *AllAnalyticsData) GetClientsUsageReports() map[string]ClientsUsageReport`

GetClientsUsageReports returns the ClientsUsageReports field if non-nil, zero value otherwise.

### GetClientsUsageReportsOk

`func (o *AllAnalyticsData) GetClientsUsageReportsOk() (*map[string]ClientsUsageReport, bool)`

GetClientsUsageReportsOk returns a tuple with the ClientsUsageReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsUsageReports

`func (o *AllAnalyticsData) SetClientsUsageReports(v map[string]ClientsUsageReport)`

SetClientsUsageReports sets ClientsUsageReports field to given value.

### HasClientsUsageReports

`func (o *AllAnalyticsData) HasClientsUsageReports() bool`

HasClientsUsageReports returns a boolean if a field has been set.

### GetDateUpdated

`func (o *AllAnalyticsData) GetDateUpdated() int64`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AllAnalyticsData) GetDateUpdatedOk() (*int64, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AllAnalyticsData) SetDateUpdated(v int64)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AllAnalyticsData) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetUsageReports

`func (o *AllAnalyticsData) GetUsageReports() map[string]UsageReportSummary`

GetUsageReports returns the UsageReports field if non-nil, zero value otherwise.

### GetUsageReportsOk

`func (o *AllAnalyticsData) GetUsageReportsOk() (*map[string]UsageReportSummary, bool)`

GetUsageReportsOk returns a tuple with the UsageReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageReports

`func (o *AllAnalyticsData) SetUsageReports(v map[string]UsageReportSummary)`

SetUsageReports sets UsageReports field to given value.

### HasUsageReports

`func (o *AllAnalyticsData) HasUsageReports() bool`

HasUsageReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


