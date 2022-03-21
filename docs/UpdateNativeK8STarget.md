# UpdateNativeK8STarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**K8sClusterCaCert** | **string** | K8S cluster CA certificate | 
**K8sClusterEndpoint** | **string** | K8S cluster URL endpoint | 
**K8sClusterToken** | **string** | K8S cluster Bearer token | 
**KeepPrevVersion** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 

## Methods

### NewUpdateNativeK8STarget

`func NewUpdateNativeK8STarget(k8sClusterCaCert string, k8sClusterEndpoint string, k8sClusterToken string, name string, ) *UpdateNativeK8STarget`

NewUpdateNativeK8STarget instantiates a new UpdateNativeK8STarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNativeK8STargetWithDefaults

`func NewUpdateNativeK8STargetWithDefaults() *UpdateNativeK8STarget`

NewUpdateNativeK8STargetWithDefaults instantiates a new UpdateNativeK8STarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateNativeK8STarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateNativeK8STarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateNativeK8STarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateNativeK8STarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetK8sClusterCaCert

`func (o *UpdateNativeK8STarget) GetK8sClusterCaCert() string`

GetK8sClusterCaCert returns the K8sClusterCaCert field if non-nil, zero value otherwise.

### GetK8sClusterCaCertOk

`func (o *UpdateNativeK8STarget) GetK8sClusterCaCertOk() (*string, bool)`

GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterCaCert

`func (o *UpdateNativeK8STarget) SetK8sClusterCaCert(v string)`

SetK8sClusterCaCert sets K8sClusterCaCert field to given value.


### GetK8sClusterEndpoint

`func (o *UpdateNativeK8STarget) GetK8sClusterEndpoint() string`

GetK8sClusterEndpoint returns the K8sClusterEndpoint field if non-nil, zero value otherwise.

### GetK8sClusterEndpointOk

`func (o *UpdateNativeK8STarget) GetK8sClusterEndpointOk() (*string, bool)`

GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterEndpoint

`func (o *UpdateNativeK8STarget) SetK8sClusterEndpoint(v string)`

SetK8sClusterEndpoint sets K8sClusterEndpoint field to given value.


### GetK8sClusterToken

`func (o *UpdateNativeK8STarget) GetK8sClusterToken() string`

GetK8sClusterToken returns the K8sClusterToken field if non-nil, zero value otherwise.

### GetK8sClusterTokenOk

`func (o *UpdateNativeK8STarget) GetK8sClusterTokenOk() (*string, bool)`

GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterToken

`func (o *UpdateNativeK8STarget) SetK8sClusterToken(v string)`

SetK8sClusterToken sets K8sClusterToken field to given value.


### GetKeepPrevVersion

`func (o *UpdateNativeK8STarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateNativeK8STarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateNativeK8STarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateNativeK8STarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateNativeK8STarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateNativeK8STarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateNativeK8STarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateNativeK8STarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateNativeK8STarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNativeK8STarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNativeK8STarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateNativeK8STarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateNativeK8STarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateNativeK8STarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateNativeK8STarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateNativeK8STarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateNativeK8STarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateNativeK8STarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateNativeK8STarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateNativeK8STarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateNativeK8STarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateNativeK8STarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateNativeK8STarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateNativeK8STarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateNativeK8STarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateNativeK8STarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateNativeK8STarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


