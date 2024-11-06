# KubeConfigValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Clusters** | Pointer to [**[]KubeconfigNamedCluster**](KubeconfigNamedCluster.md) |  | [optional] 
**Contexts** | Pointer to [**[]KubeconfigNamedContext**](KubeconfigNamedContext.md) |  | [optional] 
**CurrentContext** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]KubeconfigUser**](KubeconfigUser.md) |  | [optional] 

## Methods

### NewKubeConfigValue

`func NewKubeConfigValue() *KubeConfigValue`

NewKubeConfigValue instantiates a new KubeConfigValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigValueWithDefaults

`func NewKubeConfigValueWithDefaults() *KubeConfigValue`

NewKubeConfigValueWithDefaults instantiates a new KubeConfigValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubeConfigValue) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubeConfigValue) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubeConfigValue) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KubeConfigValue) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetClusters

`func (o *KubeConfigValue) GetClusters() []KubeconfigNamedCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *KubeConfigValue) GetClustersOk() (*[]KubeconfigNamedCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *KubeConfigValue) SetClusters(v []KubeconfigNamedCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *KubeConfigValue) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetContexts

`func (o *KubeConfigValue) GetContexts() []KubeconfigNamedContext`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *KubeConfigValue) GetContextsOk() (*[]KubeconfigNamedContext, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *KubeConfigValue) SetContexts(v []KubeconfigNamedContext)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *KubeConfigValue) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetCurrentContext

`func (o *KubeConfigValue) GetCurrentContext() string`

GetCurrentContext returns the CurrentContext field if non-nil, zero value otherwise.

### GetCurrentContextOk

`func (o *KubeConfigValue) GetCurrentContextOk() (*string, bool)`

GetCurrentContextOk returns a tuple with the CurrentContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentContext

`func (o *KubeConfigValue) SetCurrentContext(v string)`

SetCurrentContext sets CurrentContext field to given value.

### HasCurrentContext

`func (o *KubeConfigValue) HasCurrentContext() bool`

HasCurrentContext returns a boolean if a field has been set.

### GetKind

`func (o *KubeConfigValue) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubeConfigValue) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubeConfigValue) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KubeConfigValue) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUsers

`func (o *KubeConfigValue) GetUsers() []KubeconfigUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *KubeConfigValue) GetUsersOk() (*[]KubeconfigUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *KubeConfigValue) SetUsers(v []KubeconfigUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *KubeConfigValue) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


