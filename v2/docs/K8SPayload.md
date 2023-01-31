# K8SPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ca** | Pointer to **[]int32** |  | [optional] 
**ClientCert** | Pointer to **[]int32** |  | [optional] 
**ClientKey** | Pointer to **[]int32** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**SkipSystem** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewK8SPayload

`func NewK8SPayload() *K8SPayload`

NewK8SPayload instantiates a new K8SPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8SPayloadWithDefaults

`func NewK8SPayloadWithDefaults() *K8SPayload`

NewK8SPayloadWithDefaults instantiates a new K8SPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *K8SPayload) GetCa() []int32`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *K8SPayload) GetCaOk() (*[]int32, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *K8SPayload) SetCa(v []int32)`

SetCa sets Ca field to given value.

### HasCa

`func (o *K8SPayload) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetClientCert

`func (o *K8SPayload) GetClientCert() []int32`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *K8SPayload) GetClientCertOk() (*[]int32, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *K8SPayload) SetClientCert(v []int32)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *K8SPayload) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetClientKey

`func (o *K8SPayload) GetClientKey() []int32`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *K8SPayload) GetClientKeyOk() (*[]int32, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *K8SPayload) SetClientKey(v []int32)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *K8SPayload) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetNamespace

`func (o *K8SPayload) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *K8SPayload) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *K8SPayload) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *K8SPayload) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPassword

`func (o *K8SPayload) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *K8SPayload) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *K8SPayload) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *K8SPayload) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServer

`func (o *K8SPayload) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *K8SPayload) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *K8SPayload) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *K8SPayload) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSkipSystem

`func (o *K8SPayload) GetSkipSystem() bool`

GetSkipSystem returns the SkipSystem field if non-nil, zero value otherwise.

### GetSkipSystemOk

`func (o *K8SPayload) GetSkipSystemOk() (*bool, bool)`

GetSkipSystemOk returns a tuple with the SkipSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSystem

`func (o *K8SPayload) SetSkipSystem(v bool)`

SetSkipSystem sets SkipSystem field to given value.

### HasSkipSystem

`func (o *K8SPayload) HasSkipSystem() bool`

HasSkipSystem returns a boolean if a field has been set.

### GetToken

`func (o *K8SPayload) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *K8SPayload) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *K8SPayload) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *K8SPayload) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsername

`func (o *K8SPayload) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *K8SPayload) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *K8SPayload) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *K8SPayload) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


