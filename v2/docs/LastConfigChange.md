# LastConfigChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastK8sAuthsChange** | Pointer to [**K8SAuthsConfigLastChange**](K8SAuthsConfigLastChange.md) |  | [optional] 
**LastMigrationsChange** | Pointer to [**MigrationsConfigLastChange**](MigrationsConfigLastChange.md) |  | [optional] 

## Methods

### NewLastConfigChange

`func NewLastConfigChange() *LastConfigChange`

NewLastConfigChange instantiates a new LastConfigChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastConfigChangeWithDefaults

`func NewLastConfigChangeWithDefaults() *LastConfigChange`

NewLastConfigChangeWithDefaults instantiates a new LastConfigChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastK8sAuthsChange

`func (o *LastConfigChange) GetLastK8sAuthsChange() K8SAuthsConfigLastChange`

GetLastK8sAuthsChange returns the LastK8sAuthsChange field if non-nil, zero value otherwise.

### GetLastK8sAuthsChangeOk

`func (o *LastConfigChange) GetLastK8sAuthsChangeOk() (*K8SAuthsConfigLastChange, bool)`

GetLastK8sAuthsChangeOk returns a tuple with the LastK8sAuthsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastK8sAuthsChange

`func (o *LastConfigChange) SetLastK8sAuthsChange(v K8SAuthsConfigLastChange)`

SetLastK8sAuthsChange sets LastK8sAuthsChange field to given value.

### HasLastK8sAuthsChange

`func (o *LastConfigChange) HasLastK8sAuthsChange() bool`

HasLastK8sAuthsChange returns a boolean if a field has been set.

### GetLastMigrationsChange

`func (o *LastConfigChange) GetLastMigrationsChange() MigrationsConfigLastChange`

GetLastMigrationsChange returns the LastMigrationsChange field if non-nil, zero value otherwise.

### GetLastMigrationsChangeOk

`func (o *LastConfigChange) GetLastMigrationsChangeOk() (*MigrationsConfigLastChange, bool)`

GetLastMigrationsChangeOk returns a tuple with the LastMigrationsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMigrationsChange

`func (o *LastConfigChange) SetLastMigrationsChange(v MigrationsConfigLastChange)`

SetLastMigrationsChange sets LastMigrationsChange field to given value.

### HasLastMigrationsChange

`func (o *LastConfigChange) HasLastMigrationsChange() bool`

HasLastMigrationsChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


