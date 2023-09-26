# DynamicSecretCreateEks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**EksAccessKeyId** | Pointer to **string** | Access Key ID | [optional] 
**EksAssumeRole** | Pointer to **string** | IAM assume role | [optional] 
**EksClusterCaCert** | Pointer to **string** | EKS cluster CA certificate | [optional] 
**EksClusterEndpoint** | Pointer to **string** | EKS cluster URL endpoint | [optional] 
**EksClusterName** | Pointer to **string** | EKS cluster name | [optional] 
**EksRegion** | Pointer to **string** | Region | [optional] [default to "us-east-2"]
**EksSecretAccessKey** | Pointer to **string** | Secret Access Key | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessAllowPortForwading** | Pointer to **bool** | Enable Port forwarding while using CLI access | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessClusterEndpoint** | Pointer to **string** | The K8s cluster endpoint URL | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "15m"]

## Methods

### NewDynamicSecretCreateEks

`func NewDynamicSecretCreateEks(name string, ) *DynamicSecretCreateEks`

NewDynamicSecretCreateEks instantiates a new DynamicSecretCreateEks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateEksWithDefaults

`func NewDynamicSecretCreateEksWithDefaults() *DynamicSecretCreateEks`

NewDynamicSecretCreateEksWithDefaults instantiates a new DynamicSecretCreateEks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretCreateEks) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateEks) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateEks) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateEks) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetEksAccessKeyId

`func (o *DynamicSecretCreateEks) GetEksAccessKeyId() string`

GetEksAccessKeyId returns the EksAccessKeyId field if non-nil, zero value otherwise.

### GetEksAccessKeyIdOk

`func (o *DynamicSecretCreateEks) GetEksAccessKeyIdOk() (*string, bool)`

GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAccessKeyId

`func (o *DynamicSecretCreateEks) SetEksAccessKeyId(v string)`

SetEksAccessKeyId sets EksAccessKeyId field to given value.

### HasEksAccessKeyId

`func (o *DynamicSecretCreateEks) HasEksAccessKeyId() bool`

HasEksAccessKeyId returns a boolean if a field has been set.

### GetEksAssumeRole

`func (o *DynamicSecretCreateEks) GetEksAssumeRole() string`

GetEksAssumeRole returns the EksAssumeRole field if non-nil, zero value otherwise.

### GetEksAssumeRoleOk

`func (o *DynamicSecretCreateEks) GetEksAssumeRoleOk() (*string, bool)`

GetEksAssumeRoleOk returns a tuple with the EksAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAssumeRole

`func (o *DynamicSecretCreateEks) SetEksAssumeRole(v string)`

SetEksAssumeRole sets EksAssumeRole field to given value.

### HasEksAssumeRole

`func (o *DynamicSecretCreateEks) HasEksAssumeRole() bool`

HasEksAssumeRole returns a boolean if a field has been set.

### GetEksClusterCaCert

`func (o *DynamicSecretCreateEks) GetEksClusterCaCert() string`

GetEksClusterCaCert returns the EksClusterCaCert field if non-nil, zero value otherwise.

### GetEksClusterCaCertOk

`func (o *DynamicSecretCreateEks) GetEksClusterCaCertOk() (*string, bool)`

GetEksClusterCaCertOk returns a tuple with the EksClusterCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterCaCert

`func (o *DynamicSecretCreateEks) SetEksClusterCaCert(v string)`

SetEksClusterCaCert sets EksClusterCaCert field to given value.

### HasEksClusterCaCert

`func (o *DynamicSecretCreateEks) HasEksClusterCaCert() bool`

HasEksClusterCaCert returns a boolean if a field has been set.

### GetEksClusterEndpoint

`func (o *DynamicSecretCreateEks) GetEksClusterEndpoint() string`

GetEksClusterEndpoint returns the EksClusterEndpoint field if non-nil, zero value otherwise.

### GetEksClusterEndpointOk

`func (o *DynamicSecretCreateEks) GetEksClusterEndpointOk() (*string, bool)`

GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterEndpoint

`func (o *DynamicSecretCreateEks) SetEksClusterEndpoint(v string)`

SetEksClusterEndpoint sets EksClusterEndpoint field to given value.

### HasEksClusterEndpoint

`func (o *DynamicSecretCreateEks) HasEksClusterEndpoint() bool`

HasEksClusterEndpoint returns a boolean if a field has been set.

### GetEksClusterName

`func (o *DynamicSecretCreateEks) GetEksClusterName() string`

GetEksClusterName returns the EksClusterName field if non-nil, zero value otherwise.

### GetEksClusterNameOk

`func (o *DynamicSecretCreateEks) GetEksClusterNameOk() (*string, bool)`

GetEksClusterNameOk returns a tuple with the EksClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterName

`func (o *DynamicSecretCreateEks) SetEksClusterName(v string)`

SetEksClusterName sets EksClusterName field to given value.

### HasEksClusterName

`func (o *DynamicSecretCreateEks) HasEksClusterName() bool`

HasEksClusterName returns a boolean if a field has been set.

### GetEksRegion

`func (o *DynamicSecretCreateEks) GetEksRegion() string`

GetEksRegion returns the EksRegion field if non-nil, zero value otherwise.

### GetEksRegionOk

`func (o *DynamicSecretCreateEks) GetEksRegionOk() (*string, bool)`

GetEksRegionOk returns a tuple with the EksRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksRegion

`func (o *DynamicSecretCreateEks) SetEksRegion(v string)`

SetEksRegion sets EksRegion field to given value.

### HasEksRegion

`func (o *DynamicSecretCreateEks) HasEksRegion() bool`

HasEksRegion returns a boolean if a field has been set.

### GetEksSecretAccessKey

`func (o *DynamicSecretCreateEks) GetEksSecretAccessKey() string`

GetEksSecretAccessKey returns the EksSecretAccessKey field if non-nil, zero value otherwise.

### GetEksSecretAccessKeyOk

`func (o *DynamicSecretCreateEks) GetEksSecretAccessKeyOk() (*string, bool)`

GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSecretAccessKey

`func (o *DynamicSecretCreateEks) SetEksSecretAccessKey(v string)`

SetEksSecretAccessKey sets EksSecretAccessKey field to given value.

### HasEksSecretAccessKey

`func (o *DynamicSecretCreateEks) HasEksSecretAccessKey() bool`

HasEksSecretAccessKey returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateEks) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateEks) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateEks) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateEks) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateEks) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateEks) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateEks) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateEks) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateEks) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateEks) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateEks) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessAllowPortForwading

`func (o *DynamicSecretCreateEks) GetSecureAccessAllowPortForwading() bool`

GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field if non-nil, zero value otherwise.

### GetSecureAccessAllowPortForwadingOk

`func (o *DynamicSecretCreateEks) GetSecureAccessAllowPortForwadingOk() (*bool, bool)`

GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowPortForwading

`func (o *DynamicSecretCreateEks) SetSecureAccessAllowPortForwading(v bool)`

SetSecureAccessAllowPortForwading sets SecureAccessAllowPortForwading field to given value.

### HasSecureAccessAllowPortForwading

`func (o *DynamicSecretCreateEks) HasSecureAccessAllowPortForwading() bool`

HasSecureAccessAllowPortForwading returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateEks) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretCreateEks) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateEks) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretCreateEks) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessClusterEndpoint

`func (o *DynamicSecretCreateEks) GetSecureAccessClusterEndpoint() string`

GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field if non-nil, zero value otherwise.

### GetSecureAccessClusterEndpointOk

`func (o *DynamicSecretCreateEks) GetSecureAccessClusterEndpointOk() (*string, bool)`

GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessClusterEndpoint

`func (o *DynamicSecretCreateEks) SetSecureAccessClusterEndpoint(v string)`

SetSecureAccessClusterEndpoint sets SecureAccessClusterEndpoint field to given value.

### HasSecureAccessClusterEndpoint

`func (o *DynamicSecretCreateEks) HasSecureAccessClusterEndpoint() bool`

HasSecureAccessClusterEndpoint returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateEks) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateEks) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateEks) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateEks) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretCreateEks) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretCreateEks) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretCreateEks) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretCreateEks) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateEks) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateEks) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateEks) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateEks) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateEks) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateEks) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateEks) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateEks) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateEks) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateEks) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateEks) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateEks) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateEks) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateEks) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateEks) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateEks) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateEks) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateEks) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateEks) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateEks) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


