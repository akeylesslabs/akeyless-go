# KubeconfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewKubeconfigContext

`func NewKubeconfigContext() *KubeconfigContext`

NewKubeconfigContext instantiates a new KubeconfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeconfigContextWithDefaults

`func NewKubeconfigContextWithDefaults() *KubeconfigContext`

NewKubeconfigContextWithDefaults instantiates a new KubeconfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *KubeconfigContext) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *KubeconfigContext) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *KubeconfigContext) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *KubeconfigContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetNamespace

`func (o *KubeconfigContext) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubeconfigContext) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubeconfigContext) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubeconfigContext) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetUser

`func (o *KubeconfigContext) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *KubeconfigContext) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *KubeconfigContext) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *KubeconfigContext) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


