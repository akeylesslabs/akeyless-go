# DynamicSecretProducerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GwClusterId** | Pointer to **int64** |  | [optional] 
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


