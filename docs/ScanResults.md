# ScanResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountExisting** | Pointer to **int64** |  | [optional] 
**CountFailed** | Pointer to **int64** |  | [optional] 
**CountHosts** | Pointer to **int64** |  | [optional] 
**CountNew** | Pointer to **int64** |  | [optional] 
**CountSubdomains** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**FailedTargets** | Pointer to [**[]TargetError**](TargetError.md) |  | [optional] 
**ItemNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewScanResults

`func NewScanResults() *ScanResults`

NewScanResults instantiates a new ScanResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanResultsWithDefaults

`func NewScanResultsWithDefaults() *ScanResults`

NewScanResultsWithDefaults instantiates a new ScanResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountExisting

`func (o *ScanResults) GetCountExisting() int64`

GetCountExisting returns the CountExisting field if non-nil, zero value otherwise.

### GetCountExistingOk

`func (o *ScanResults) GetCountExistingOk() (*int64, bool)`

GetCountExistingOk returns a tuple with the CountExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountExisting

`func (o *ScanResults) SetCountExisting(v int64)`

SetCountExisting sets CountExisting field to given value.

### HasCountExisting

`func (o *ScanResults) HasCountExisting() bool`

HasCountExisting returns a boolean if a field has been set.

### GetCountFailed

`func (o *ScanResults) GetCountFailed() int64`

GetCountFailed returns the CountFailed field if non-nil, zero value otherwise.

### GetCountFailedOk

`func (o *ScanResults) GetCountFailedOk() (*int64, bool)`

GetCountFailedOk returns a tuple with the CountFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountFailed

`func (o *ScanResults) SetCountFailed(v int64)`

SetCountFailed sets CountFailed field to given value.

### HasCountFailed

`func (o *ScanResults) HasCountFailed() bool`

HasCountFailed returns a boolean if a field has been set.

### GetCountHosts

`func (o *ScanResults) GetCountHosts() int64`

GetCountHosts returns the CountHosts field if non-nil, zero value otherwise.

### GetCountHostsOk

`func (o *ScanResults) GetCountHostsOk() (*int64, bool)`

GetCountHostsOk returns a tuple with the CountHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountHosts

`func (o *ScanResults) SetCountHosts(v int64)`

SetCountHosts sets CountHosts field to given value.

### HasCountHosts

`func (o *ScanResults) HasCountHosts() bool`

HasCountHosts returns a boolean if a field has been set.

### GetCountNew

`func (o *ScanResults) GetCountNew() int64`

GetCountNew returns the CountNew field if non-nil, zero value otherwise.

### GetCountNewOk

`func (o *ScanResults) GetCountNewOk() (*int64, bool)`

GetCountNewOk returns a tuple with the CountNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountNew

`func (o *ScanResults) SetCountNew(v int64)`

SetCountNew sets CountNew field to given value.

### HasCountNew

`func (o *ScanResults) HasCountNew() bool`

HasCountNew returns a boolean if a field has been set.

### GetCountSubdomains

`func (o *ScanResults) GetCountSubdomains() int64`

GetCountSubdomains returns the CountSubdomains field if non-nil, zero value otherwise.

### GetCountSubdomainsOk

`func (o *ScanResults) GetCountSubdomainsOk() (*int64, bool)`

GetCountSubdomainsOk returns a tuple with the CountSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountSubdomains

`func (o *ScanResults) SetCountSubdomains(v int64)`

SetCountSubdomains sets CountSubdomains field to given value.

### HasCountSubdomains

`func (o *ScanResults) HasCountSubdomains() bool`

HasCountSubdomains returns a boolean if a field has been set.

### GetError

`func (o *ScanResults) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScanResults) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScanResults) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ScanResults) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFailedTargets

`func (o *ScanResults) GetFailedTargets() []TargetError`

GetFailedTargets returns the FailedTargets field if non-nil, zero value otherwise.

### GetFailedTargetsOk

`func (o *ScanResults) GetFailedTargetsOk() (*[]TargetError, bool)`

GetFailedTargetsOk returns a tuple with the FailedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTargets

`func (o *ScanResults) SetFailedTargets(v []TargetError)`

SetFailedTargets sets FailedTargets field to given value.

### HasFailedTargets

`func (o *ScanResults) HasFailedTargets() bool`

HasFailedTargets returns a boolean if a field has been set.

### GetItemNames

`func (o *ScanResults) GetItemNames() []string`

GetItemNames returns the ItemNames field if non-nil, zero value otherwise.

### GetItemNamesOk

`func (o *ScanResults) GetItemNamesOk() (*[]string, bool)`

GetItemNamesOk returns a tuple with the ItemNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemNames

`func (o *ScanResults) SetItemNames(v []string)`

SetItemNames sets ItemNames field to given value.

### HasItemNames

`func (o *ScanResults) HasItemNames() bool`

HasItemNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


