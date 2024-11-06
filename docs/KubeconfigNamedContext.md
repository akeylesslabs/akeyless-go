# KubeconfigNamedContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**KubeconfigContext**](KubeconfigContext.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewKubeconfigNamedContext

`func NewKubeconfigNamedContext() *KubeconfigNamedContext`

NewKubeconfigNamedContext instantiates a new KubeconfigNamedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeconfigNamedContextWithDefaults

`func NewKubeconfigNamedContextWithDefaults() *KubeconfigNamedContext`

NewKubeconfigNamedContextWithDefaults instantiates a new KubeconfigNamedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *KubeconfigNamedContext) GetContext() KubeconfigContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *KubeconfigNamedContext) GetContextOk() (*KubeconfigContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *KubeconfigNamedContext) SetContext(v KubeconfigContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *KubeconfigNamedContext) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetName

`func (o *KubeconfigNamedContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubeconfigNamedContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubeconfigNamedContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubeconfigNamedContext) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


