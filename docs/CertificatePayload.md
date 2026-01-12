# CertificatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationEvents** | Pointer to [**[]CertificateExpirationEvent**](CertificateExpirationEvent.md) |  | [optional] 
**Folder** | Pointer to **string** |  | [optional] 
**MaxDialTimeout** | Pointer to **int64** |  | [optional] 
**MaxScanDuration** | Pointer to **int64** |  | [optional] 
**MaxWorkers** | Pointer to **int64** |  | [optional] 
**PortRanges** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]CertificateScanTarget**](CertificateScanTarget.md) |  | [optional] 

## Methods

### NewCertificatePayload

`func NewCertificatePayload() *CertificatePayload`

NewCertificatePayload instantiates a new CertificatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatePayloadWithDefaults

`func NewCertificatePayloadWithDefaults() *CertificatePayload`

NewCertificatePayloadWithDefaults instantiates a new CertificatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationEvents

`func (o *CertificatePayload) GetExpirationEvents() []CertificateExpirationEvent`

GetExpirationEvents returns the ExpirationEvents field if non-nil, zero value otherwise.

### GetExpirationEventsOk

`func (o *CertificatePayload) GetExpirationEventsOk() (*[]CertificateExpirationEvent, bool)`

GetExpirationEventsOk returns a tuple with the ExpirationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEvents

`func (o *CertificatePayload) SetExpirationEvents(v []CertificateExpirationEvent)`

SetExpirationEvents sets ExpirationEvents field to given value.

### HasExpirationEvents

`func (o *CertificatePayload) HasExpirationEvents() bool`

HasExpirationEvents returns a boolean if a field has been set.

### GetFolder

`func (o *CertificatePayload) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *CertificatePayload) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *CertificatePayload) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *CertificatePayload) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetMaxDialTimeout

`func (o *CertificatePayload) GetMaxDialTimeout() int64`

GetMaxDialTimeout returns the MaxDialTimeout field if non-nil, zero value otherwise.

### GetMaxDialTimeoutOk

`func (o *CertificatePayload) GetMaxDialTimeoutOk() (*int64, bool)`

GetMaxDialTimeoutOk returns a tuple with the MaxDialTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDialTimeout

`func (o *CertificatePayload) SetMaxDialTimeout(v int64)`

SetMaxDialTimeout sets MaxDialTimeout field to given value.

### HasMaxDialTimeout

`func (o *CertificatePayload) HasMaxDialTimeout() bool`

HasMaxDialTimeout returns a boolean if a field has been set.

### GetMaxScanDuration

`func (o *CertificatePayload) GetMaxScanDuration() int64`

GetMaxScanDuration returns the MaxScanDuration field if non-nil, zero value otherwise.

### GetMaxScanDurationOk

`func (o *CertificatePayload) GetMaxScanDurationOk() (*int64, bool)`

GetMaxScanDurationOk returns a tuple with the MaxScanDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScanDuration

`func (o *CertificatePayload) SetMaxScanDuration(v int64)`

SetMaxScanDuration sets MaxScanDuration field to given value.

### HasMaxScanDuration

`func (o *CertificatePayload) HasMaxScanDuration() bool`

HasMaxScanDuration returns a boolean if a field has been set.

### GetMaxWorkers

`func (o *CertificatePayload) GetMaxWorkers() int64`

GetMaxWorkers returns the MaxWorkers field if non-nil, zero value otherwise.

### GetMaxWorkersOk

`func (o *CertificatePayload) GetMaxWorkersOk() (*int64, bool)`

GetMaxWorkersOk returns a tuple with the MaxWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkers

`func (o *CertificatePayload) SetMaxWorkers(v int64)`

SetMaxWorkers sets MaxWorkers field to given value.

### HasMaxWorkers

`func (o *CertificatePayload) HasMaxWorkers() bool`

HasMaxWorkers returns a boolean if a field has been set.

### GetPortRanges

`func (o *CertificatePayload) GetPortRanges() string`

GetPortRanges returns the PortRanges field if non-nil, zero value otherwise.

### GetPortRangesOk

`func (o *CertificatePayload) GetPortRangesOk() (*string, bool)`

GetPortRangesOk returns a tuple with the PortRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRanges

`func (o *CertificatePayload) SetPortRanges(v string)`

SetPortRanges sets PortRanges field to given value.

### HasPortRanges

`func (o *CertificatePayload) HasPortRanges() bool`

HasPortRanges returns a boolean if a field has been set.

### GetTargets

`func (o *CertificatePayload) GetTargets() []CertificateScanTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CertificatePayload) GetTargetsOk() (*[]CertificateScanTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CertificatePayload) SetTargets(v []CertificateScanTarget)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *CertificatePayload) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


