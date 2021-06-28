# CreateNativeK8STarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**K8sClusterCaCert** | **string** | K8S cluster CA certificate | 
**K8sClusterEndpoint** | **string** | K8S cluster URL endpoint | 
**K8sClusterToken** | **string** | K8S cluster Bearer token | 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreateNativeK8STarget

`func NewCreateNativeK8STarget(k8sClusterCaCert string, k8sClusterEndpoint string, k8sClusterToken string, name string, ) *CreateNativeK8STarget`

NewCreateNativeK8STarget instantiates a new CreateNativeK8STarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNativeK8STargetWithDefaults

`func NewCreateNativeK8STargetWithDefaults() *CreateNativeK8STarget`

NewCreateNativeK8STargetWithDefaults instantiates a new CreateNativeK8STarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateNativeK8STarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateNativeK8STarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateNativeK8STarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateNativeK8STarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetK8sClusterCaCert

`func (o *CreateNativeK8STarget) GetK8sClusterCaCert() string`

GetK8sClusterCaCert returns the K8sClusterCaCert field if non-nil, zero value otherwise.

### GetK8sClusterCaCertOk

`func (o *CreateNativeK8STarget) GetK8sClusterCaCertOk() (*string, bool)`

GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterCaCert

`func (o *CreateNativeK8STarget) SetK8sClusterCaCert(v string)`

SetK8sClusterCaCert sets K8sClusterCaCert field to given value.


### GetK8sClusterEndpoint

`func (o *CreateNativeK8STarget) GetK8sClusterEndpoint() string`

GetK8sClusterEndpoint returns the K8sClusterEndpoint field if non-nil, zero value otherwise.

### GetK8sClusterEndpointOk

`func (o *CreateNativeK8STarget) GetK8sClusterEndpointOk() (*string, bool)`

GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterEndpoint

`func (o *CreateNativeK8STarget) SetK8sClusterEndpoint(v string)`

SetK8sClusterEndpoint sets K8sClusterEndpoint field to given value.


### GetK8sClusterToken

`func (o *CreateNativeK8STarget) GetK8sClusterToken() string`

GetK8sClusterToken returns the K8sClusterToken field if non-nil, zero value otherwise.

### GetK8sClusterTokenOk

`func (o *CreateNativeK8STarget) GetK8sClusterTokenOk() (*string, bool)`

GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterToken

`func (o *CreateNativeK8STarget) SetK8sClusterToken(v string)`

SetK8sClusterToken sets K8sClusterToken field to given value.


### GetKey

`func (o *CreateNativeK8STarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateNativeK8STarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateNativeK8STarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateNativeK8STarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateNativeK8STarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNativeK8STarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNativeK8STarget) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateNativeK8STarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateNativeK8STarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateNativeK8STarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateNativeK8STarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *CreateNativeK8STarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateNativeK8STarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateNativeK8STarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateNativeK8STarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateNativeK8STarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateNativeK8STarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateNativeK8STarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateNativeK8STarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateNativeK8STarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateNativeK8STarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateNativeK8STarget) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateNativeK8STarget) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


