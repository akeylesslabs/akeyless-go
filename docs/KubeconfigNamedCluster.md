# KubeconfigNamedCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**KubeconfigCluster**](KubeconfigCluster.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewKubeconfigNamedCluster

`func NewKubeconfigNamedCluster() *KubeconfigNamedCluster`

NewKubeconfigNamedCluster instantiates a new KubeconfigNamedCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeconfigNamedClusterWithDefaults

`func NewKubeconfigNamedClusterWithDefaults() *KubeconfigNamedCluster`

NewKubeconfigNamedClusterWithDefaults instantiates a new KubeconfigNamedCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *KubeconfigNamedCluster) GetCluster() KubeconfigCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *KubeconfigNamedCluster) GetClusterOk() (*KubeconfigCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *KubeconfigNamedCluster) SetCluster(v KubeconfigCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *KubeconfigNamedCluster) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetName

`func (o *KubeconfigNamedCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubeconfigNamedCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubeconfigNamedCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubeconfigNamedCluster) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


