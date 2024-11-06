# KubeconfigCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateAuthority** | Pointer to **string** | CertificateAuthority is optional and can be omitted if not used. | [optional] 
**CertificateAuthorityData** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 

## Methods

### NewKubeconfigCluster

`func NewKubeconfigCluster() *KubeconfigCluster`

NewKubeconfigCluster instantiates a new KubeconfigCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeconfigClusterWithDefaults

`func NewKubeconfigClusterWithDefaults() *KubeconfigCluster`

NewKubeconfigClusterWithDefaults instantiates a new KubeconfigCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateAuthority

`func (o *KubeconfigCluster) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *KubeconfigCluster) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *KubeconfigCluster) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *KubeconfigCluster) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetCertificateAuthorityData

`func (o *KubeconfigCluster) GetCertificateAuthorityData() string`

GetCertificateAuthorityData returns the CertificateAuthorityData field if non-nil, zero value otherwise.

### GetCertificateAuthorityDataOk

`func (o *KubeconfigCluster) GetCertificateAuthorityDataOk() (*string, bool)`

GetCertificateAuthorityDataOk returns a tuple with the CertificateAuthorityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityData

`func (o *KubeconfigCluster) SetCertificateAuthorityData(v string)`

SetCertificateAuthorityData sets CertificateAuthorityData field to given value.

### HasCertificateAuthorityData

`func (o *KubeconfigCluster) HasCertificateAuthorityData() bool`

HasCertificateAuthorityData returns a boolean if a field has been set.

### GetServer

`func (o *KubeconfigCluster) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KubeconfigCluster) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KubeconfigCluster) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *KubeconfigCluster) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


