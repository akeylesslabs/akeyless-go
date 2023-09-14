# EKSTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EksAccessKeyId** | Pointer to **string** |  | [optional] 
**EksClusterCaCertificate** | Pointer to **string** |  | [optional] 
**EksClusterEndpoint** | Pointer to **string** |  | [optional] 
**EksClusterName** | Pointer to **string** |  | [optional] 
**EksRegion** | Pointer to **string** |  | [optional] 
**EksSecretAccessKey** | Pointer to **string** |  | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** |  | [optional] 

## Methods

### NewEKSTargetDetails

`func NewEKSTargetDetails() *EKSTargetDetails`

NewEKSTargetDetails instantiates a new EKSTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEKSTargetDetailsWithDefaults

`func NewEKSTargetDetailsWithDefaults() *EKSTargetDetails`

NewEKSTargetDetailsWithDefaults instantiates a new EKSTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEksAccessKeyId

`func (o *EKSTargetDetails) GetEksAccessKeyId() string`

GetEksAccessKeyId returns the EksAccessKeyId field if non-nil, zero value otherwise.

### GetEksAccessKeyIdOk

`func (o *EKSTargetDetails) GetEksAccessKeyIdOk() (*string, bool)`

GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAccessKeyId

`func (o *EKSTargetDetails) SetEksAccessKeyId(v string)`

SetEksAccessKeyId sets EksAccessKeyId field to given value.

### HasEksAccessKeyId

`func (o *EKSTargetDetails) HasEksAccessKeyId() bool`

HasEksAccessKeyId returns a boolean if a field has been set.

### GetEksClusterCaCertificate

`func (o *EKSTargetDetails) GetEksClusterCaCertificate() string`

GetEksClusterCaCertificate returns the EksClusterCaCertificate field if non-nil, zero value otherwise.

### GetEksClusterCaCertificateOk

`func (o *EKSTargetDetails) GetEksClusterCaCertificateOk() (*string, bool)`

GetEksClusterCaCertificateOk returns a tuple with the EksClusterCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterCaCertificate

`func (o *EKSTargetDetails) SetEksClusterCaCertificate(v string)`

SetEksClusterCaCertificate sets EksClusterCaCertificate field to given value.

### HasEksClusterCaCertificate

`func (o *EKSTargetDetails) HasEksClusterCaCertificate() bool`

HasEksClusterCaCertificate returns a boolean if a field has been set.

### GetEksClusterEndpoint

`func (o *EKSTargetDetails) GetEksClusterEndpoint() string`

GetEksClusterEndpoint returns the EksClusterEndpoint field if non-nil, zero value otherwise.

### GetEksClusterEndpointOk

`func (o *EKSTargetDetails) GetEksClusterEndpointOk() (*string, bool)`

GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterEndpoint

`func (o *EKSTargetDetails) SetEksClusterEndpoint(v string)`

SetEksClusterEndpoint sets EksClusterEndpoint field to given value.

### HasEksClusterEndpoint

`func (o *EKSTargetDetails) HasEksClusterEndpoint() bool`

HasEksClusterEndpoint returns a boolean if a field has been set.

### GetEksClusterName

`func (o *EKSTargetDetails) GetEksClusterName() string`

GetEksClusterName returns the EksClusterName field if non-nil, zero value otherwise.

### GetEksClusterNameOk

`func (o *EKSTargetDetails) GetEksClusterNameOk() (*string, bool)`

GetEksClusterNameOk returns a tuple with the EksClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterName

`func (o *EKSTargetDetails) SetEksClusterName(v string)`

SetEksClusterName sets EksClusterName field to given value.

### HasEksClusterName

`func (o *EKSTargetDetails) HasEksClusterName() bool`

HasEksClusterName returns a boolean if a field has been set.

### GetEksRegion

`func (o *EKSTargetDetails) GetEksRegion() string`

GetEksRegion returns the EksRegion field if non-nil, zero value otherwise.

### GetEksRegionOk

`func (o *EKSTargetDetails) GetEksRegionOk() (*string, bool)`

GetEksRegionOk returns a tuple with the EksRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksRegion

`func (o *EKSTargetDetails) SetEksRegion(v string)`

SetEksRegion sets EksRegion field to given value.

### HasEksRegion

`func (o *EKSTargetDetails) HasEksRegion() bool`

HasEksRegion returns a boolean if a field has been set.

### GetEksSecretAccessKey

`func (o *EKSTargetDetails) GetEksSecretAccessKey() string`

GetEksSecretAccessKey returns the EksSecretAccessKey field if non-nil, zero value otherwise.

### GetEksSecretAccessKeyOk

`func (o *EKSTargetDetails) GetEksSecretAccessKeyOk() (*string, bool)`

GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSecretAccessKey

`func (o *EKSTargetDetails) SetEksSecretAccessKey(v string)`

SetEksSecretAccessKey sets EksSecretAccessKey field to given value.

### HasEksSecretAccessKey

`func (o *EKSTargetDetails) HasEksSecretAccessKey() bool`

HasEksSecretAccessKey returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *EKSTargetDetails) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *EKSTargetDetails) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *EKSTargetDetails) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *EKSTargetDetails) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


