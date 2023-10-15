# NativeK8sTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**K8sAuthType** | Pointer to **string** |  | [optional] 
**K8sBearerToken** | Pointer to **string** |  | [optional] 
**K8sClientCertData** | Pointer to **string** | For K8s Client certificates authentication | [optional] 
**K8sClientKeyData** | Pointer to **string** |  | [optional] 
**K8sClusterCaCertificate** | Pointer to **string** |  | [optional] 
**K8sClusterEndpoint** | Pointer to **string** |  | [optional] 
**UseGwServiceAccount** | Pointer to **bool** |  | [optional] 

## Methods

### NewNativeK8sTargetDetails

`func NewNativeK8sTargetDetails() *NativeK8sTargetDetails`

NewNativeK8sTargetDetails instantiates a new NativeK8sTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNativeK8sTargetDetailsWithDefaults

`func NewNativeK8sTargetDetailsWithDefaults() *NativeK8sTargetDetails`

NewNativeK8sTargetDetailsWithDefaults instantiates a new NativeK8sTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetK8sAuthType

`func (o *NativeK8sTargetDetails) GetK8sAuthType() string`

GetK8sAuthType returns the K8sAuthType field if non-nil, zero value otherwise.

### GetK8sAuthTypeOk

`func (o *NativeK8sTargetDetails) GetK8sAuthTypeOk() (*string, bool)`

GetK8sAuthTypeOk returns a tuple with the K8sAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAuthType

`func (o *NativeK8sTargetDetails) SetK8sAuthType(v string)`

SetK8sAuthType sets K8sAuthType field to given value.

### HasK8sAuthType

`func (o *NativeK8sTargetDetails) HasK8sAuthType() bool`

HasK8sAuthType returns a boolean if a field has been set.

### GetK8sBearerToken

`func (o *NativeK8sTargetDetails) GetK8sBearerToken() string`

GetK8sBearerToken returns the K8sBearerToken field if non-nil, zero value otherwise.

### GetK8sBearerTokenOk

`func (o *NativeK8sTargetDetails) GetK8sBearerTokenOk() (*string, bool)`

GetK8sBearerTokenOk returns a tuple with the K8sBearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sBearerToken

`func (o *NativeK8sTargetDetails) SetK8sBearerToken(v string)`

SetK8sBearerToken sets K8sBearerToken field to given value.

### HasK8sBearerToken

`func (o *NativeK8sTargetDetails) HasK8sBearerToken() bool`

HasK8sBearerToken returns a boolean if a field has been set.

### GetK8sClientCertData

`func (o *NativeK8sTargetDetails) GetK8sClientCertData() string`

GetK8sClientCertData returns the K8sClientCertData field if non-nil, zero value otherwise.

### GetK8sClientCertDataOk

`func (o *NativeK8sTargetDetails) GetK8sClientCertDataOk() (*string, bool)`

GetK8sClientCertDataOk returns a tuple with the K8sClientCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientCertData

`func (o *NativeK8sTargetDetails) SetK8sClientCertData(v string)`

SetK8sClientCertData sets K8sClientCertData field to given value.

### HasK8sClientCertData

`func (o *NativeK8sTargetDetails) HasK8sClientCertData() bool`

HasK8sClientCertData returns a boolean if a field has been set.

### GetK8sClientKeyData

`func (o *NativeK8sTargetDetails) GetK8sClientKeyData() string`

GetK8sClientKeyData returns the K8sClientKeyData field if non-nil, zero value otherwise.

### GetK8sClientKeyDataOk

`func (o *NativeK8sTargetDetails) GetK8sClientKeyDataOk() (*string, bool)`

GetK8sClientKeyDataOk returns a tuple with the K8sClientKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientKeyData

`func (o *NativeK8sTargetDetails) SetK8sClientKeyData(v string)`

SetK8sClientKeyData sets K8sClientKeyData field to given value.

### HasK8sClientKeyData

`func (o *NativeK8sTargetDetails) HasK8sClientKeyData() bool`

HasK8sClientKeyData returns a boolean if a field has been set.

### GetK8sClusterCaCertificate

`func (o *NativeK8sTargetDetails) GetK8sClusterCaCertificate() string`

GetK8sClusterCaCertificate returns the K8sClusterCaCertificate field if non-nil, zero value otherwise.

### GetK8sClusterCaCertificateOk

`func (o *NativeK8sTargetDetails) GetK8sClusterCaCertificateOk() (*string, bool)`

GetK8sClusterCaCertificateOk returns a tuple with the K8sClusterCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterCaCertificate

`func (o *NativeK8sTargetDetails) SetK8sClusterCaCertificate(v string)`

SetK8sClusterCaCertificate sets K8sClusterCaCertificate field to given value.

### HasK8sClusterCaCertificate

`func (o *NativeK8sTargetDetails) HasK8sClusterCaCertificate() bool`

HasK8sClusterCaCertificate returns a boolean if a field has been set.

### GetK8sClusterEndpoint

`func (o *NativeK8sTargetDetails) GetK8sClusterEndpoint() string`

GetK8sClusterEndpoint returns the K8sClusterEndpoint field if non-nil, zero value otherwise.

### GetK8sClusterEndpointOk

`func (o *NativeK8sTargetDetails) GetK8sClusterEndpointOk() (*string, bool)`

GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterEndpoint

`func (o *NativeK8sTargetDetails) SetK8sClusterEndpoint(v string)`

SetK8sClusterEndpoint sets K8sClusterEndpoint field to given value.

### HasK8sClusterEndpoint

`func (o *NativeK8sTargetDetails) HasK8sClusterEndpoint() bool`

HasK8sClusterEndpoint returns a boolean if a field has been set.

### GetUseGwServiceAccount

`func (o *NativeK8sTargetDetails) GetUseGwServiceAccount() bool`

GetUseGwServiceAccount returns the UseGwServiceAccount field if non-nil, zero value otherwise.

### GetUseGwServiceAccountOk

`func (o *NativeK8sTargetDetails) GetUseGwServiceAccountOk() (*bool, bool)`

GetUseGwServiceAccountOk returns a tuple with the UseGwServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwServiceAccount

`func (o *NativeK8sTargetDetails) SetUseGwServiceAccount(v bool)`

SetUseGwServiceAccount sets UseGwServiceAccount field to given value.

### HasUseGwServiceAccount

`func (o *NativeK8sTargetDetails) HasUseGwServiceAccount() bool`

HasUseGwServiceAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


