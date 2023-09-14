# DynamicSecretProducerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureMessage** | Pointer to **string** |  | [optional] 
**GwClusterId** | Pointer to **int64** |  | [optional] 
**K8sAllowedNamespaces** | Pointer to **string** | Relevant only for generic k8s producer | [optional] 
**K8sDynamicMode** | Pointer to **bool** | Relevant only for generic k8s producer | [optional] 
**ProducerLastKeepAlive** | Pointer to **string** |  | [optional] 
**ProducerMetadata** | Pointer to **string** |  | [optional] 
**ProducerStatus** | Pointer to **string** | RotationStatus defines types of rotation Status | [optional] 
**ProducerType** | Pointer to **string** |  | [optional] 

## Methods

### NewDynamicSecretProducerInfo

`func NewDynamicSecretProducerInfo() *DynamicSecretProducerInfo`

NewDynamicSecretProducerInfo instantiates a new DynamicSecretProducerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretProducerInfoWithDefaults

`func NewDynamicSecretProducerInfoWithDefaults() *DynamicSecretProducerInfo`

NewDynamicSecretProducerInfoWithDefaults instantiates a new DynamicSecretProducerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureMessage

`func (o *DynamicSecretProducerInfo) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *DynamicSecretProducerInfo) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *DynamicSecretProducerInfo) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *DynamicSecretProducerInfo) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetGwClusterId

`func (o *DynamicSecretProducerInfo) GetGwClusterId() int64`

GetGwClusterId returns the GwClusterId field if non-nil, zero value otherwise.

### GetGwClusterIdOk

`func (o *DynamicSecretProducerInfo) GetGwClusterIdOk() (*int64, bool)`

GetGwClusterIdOk returns a tuple with the GwClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwClusterId

`func (o *DynamicSecretProducerInfo) SetGwClusterId(v int64)`

SetGwClusterId sets GwClusterId field to given value.

### HasGwClusterId

`func (o *DynamicSecretProducerInfo) HasGwClusterId() bool`

HasGwClusterId returns a boolean if a field has been set.

### GetK8sAllowedNamespaces

`func (o *DynamicSecretProducerInfo) GetK8sAllowedNamespaces() string`

GetK8sAllowedNamespaces returns the K8sAllowedNamespaces field if non-nil, zero value otherwise.

### GetK8sAllowedNamespacesOk

`func (o *DynamicSecretProducerInfo) GetK8sAllowedNamespacesOk() (*string, bool)`

GetK8sAllowedNamespacesOk returns a tuple with the K8sAllowedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAllowedNamespaces

`func (o *DynamicSecretProducerInfo) SetK8sAllowedNamespaces(v string)`

SetK8sAllowedNamespaces sets K8sAllowedNamespaces field to given value.

### HasK8sAllowedNamespaces

`func (o *DynamicSecretProducerInfo) HasK8sAllowedNamespaces() bool`

HasK8sAllowedNamespaces returns a boolean if a field has been set.

### GetK8sDynamicMode

`func (o *DynamicSecretProducerInfo) GetK8sDynamicMode() bool`

GetK8sDynamicMode returns the K8sDynamicMode field if non-nil, zero value otherwise.

### GetK8sDynamicModeOk

`func (o *DynamicSecretProducerInfo) GetK8sDynamicModeOk() (*bool, bool)`

GetK8sDynamicModeOk returns a tuple with the K8sDynamicMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sDynamicMode

`func (o *DynamicSecretProducerInfo) SetK8sDynamicMode(v bool)`

SetK8sDynamicMode sets K8sDynamicMode field to given value.

### HasK8sDynamicMode

`func (o *DynamicSecretProducerInfo) HasK8sDynamicMode() bool`

HasK8sDynamicMode returns a boolean if a field has been set.

### GetProducerLastKeepAlive

`func (o *DynamicSecretProducerInfo) GetProducerLastKeepAlive() string`

GetProducerLastKeepAlive returns the ProducerLastKeepAlive field if non-nil, zero value otherwise.

### GetProducerLastKeepAliveOk

`func (o *DynamicSecretProducerInfo) GetProducerLastKeepAliveOk() (*string, bool)`

GetProducerLastKeepAliveOk returns a tuple with the ProducerLastKeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerLastKeepAlive

`func (o *DynamicSecretProducerInfo) SetProducerLastKeepAlive(v string)`

SetProducerLastKeepAlive sets ProducerLastKeepAlive field to given value.

### HasProducerLastKeepAlive

`func (o *DynamicSecretProducerInfo) HasProducerLastKeepAlive() bool`

HasProducerLastKeepAlive returns a boolean if a field has been set.

### GetProducerMetadata

`func (o *DynamicSecretProducerInfo) GetProducerMetadata() string`

GetProducerMetadata returns the ProducerMetadata field if non-nil, zero value otherwise.

### GetProducerMetadataOk

`func (o *DynamicSecretProducerInfo) GetProducerMetadataOk() (*string, bool)`

GetProducerMetadataOk returns a tuple with the ProducerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerMetadata

`func (o *DynamicSecretProducerInfo) SetProducerMetadata(v string)`

SetProducerMetadata sets ProducerMetadata field to given value.

### HasProducerMetadata

`func (o *DynamicSecretProducerInfo) HasProducerMetadata() bool`

HasProducerMetadata returns a boolean if a field has been set.

### GetProducerStatus

`func (o *DynamicSecretProducerInfo) GetProducerStatus() string`

GetProducerStatus returns the ProducerStatus field if non-nil, zero value otherwise.

### GetProducerStatusOk

`func (o *DynamicSecretProducerInfo) GetProducerStatusOk() (*string, bool)`

GetProducerStatusOk returns a tuple with the ProducerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerStatus

`func (o *DynamicSecretProducerInfo) SetProducerStatus(v string)`

SetProducerStatus sets ProducerStatus field to given value.

### HasProducerStatus

`func (o *DynamicSecretProducerInfo) HasProducerStatus() bool`

HasProducerStatus returns a boolean if a field has been set.

### GetProducerType

`func (o *DynamicSecretProducerInfo) GetProducerType() string`

GetProducerType returns the ProducerType field if non-nil, zero value otherwise.

### GetProducerTypeOk

`func (o *DynamicSecretProducerInfo) GetProducerTypeOk() (*string, bool)`

GetProducerTypeOk returns a tuple with the ProducerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerType

`func (o *DynamicSecretProducerInfo) SetProducerType(v string)`

SetProducerType sets ProducerType field to given value.

### HasProducerType

`func (o *DynamicSecretProducerInfo) HasProducerType() bool`

HasProducerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


