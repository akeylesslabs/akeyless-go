# KubeconfigGenerateOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConflictedClustersNames** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**KubeConfigValue**](KubeConfigValue.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKubeconfigGenerateOutput

`func NewKubeconfigGenerateOutput() *KubeconfigGenerateOutput`

NewKubeconfigGenerateOutput instantiates a new KubeconfigGenerateOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeconfigGenerateOutputWithDefaults

`func NewKubeconfigGenerateOutputWithDefaults() *KubeconfigGenerateOutput`

NewKubeconfigGenerateOutputWithDefaults instantiates a new KubeconfigGenerateOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConflictedClustersNames

`func (o *KubeconfigGenerateOutput) GetConflictedClustersNames() []string`

GetConflictedClustersNames returns the ConflictedClustersNames field if non-nil, zero value otherwise.

### GetConflictedClustersNamesOk

`func (o *KubeconfigGenerateOutput) GetConflictedClustersNamesOk() (*[]string, bool)`

GetConflictedClustersNamesOk returns a tuple with the ConflictedClustersNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictedClustersNames

`func (o *KubeconfigGenerateOutput) SetConflictedClustersNames(v []string)`

SetConflictedClustersNames sets ConflictedClustersNames field to given value.

### HasConflictedClustersNames

`func (o *KubeconfigGenerateOutput) HasConflictedClustersNames() bool`

HasConflictedClustersNames returns a boolean if a field has been set.

### GetData

`func (o *KubeconfigGenerateOutput) GetData() KubeConfigValue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KubeconfigGenerateOutput) GetDataOk() (*KubeConfigValue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KubeconfigGenerateOutput) SetData(v KubeConfigValue)`

SetData sets Data field to given value.

### HasData

`func (o *KubeconfigGenerateOutput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *KubeconfigGenerateOutput) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *KubeconfigGenerateOutput) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *KubeconfigGenerateOutput) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *KubeconfigGenerateOutput) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


