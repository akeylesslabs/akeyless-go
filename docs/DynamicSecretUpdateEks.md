# DynamicSecretUpdateEks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EksAccessKeyId** | Pointer to **string** | Access Key ID | [optional] 
**EksAssumeRole** | Pointer to **string** | IAM assume role | [optional] 
**EksClusterCaCert** | Pointer to **string** | EKS cluster CA certificate | [optional] 
**EksClusterEndpoint** | Pointer to **string** | EKS cluster URL endpoint | [optional] 
**EksClusterName** | Pointer to **string** | EKS cluster name | [optional] 
**EksRegion** | Pointer to **string** | Region | [optional] [default to "us-east-2"]
**EksSecretAccessKey** | Pointer to **string** | Secret Access Key | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
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

### NewDynamicSecretUpdateEks

`func NewDynamicSecretUpdateEks(name string, ) *DynamicSecretUpdateEks`

NewDynamicSecretUpdateEks instantiates a new DynamicSecretUpdateEks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateEksWithDefaults

`func NewDynamicSecretUpdateEksWithDefaults() *DynamicSecretUpdateEks`

NewDynamicSecretUpdateEksWithDefaults instantiates a new DynamicSecretUpdateEks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretUpdateEks) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateEks) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateEks) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateEks) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateEks) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateEks) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateEks) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateEks) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEksAccessKeyId

`func (o *DynamicSecretUpdateEks) GetEksAccessKeyId() string`

GetEksAccessKeyId returns the EksAccessKeyId field if non-nil, zero value otherwise.

### GetEksAccessKeyIdOk

`func (o *DynamicSecretUpdateEks) GetEksAccessKeyIdOk() (*string, bool)`

GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAccessKeyId

`func (o *DynamicSecretUpdateEks) SetEksAccessKeyId(v string)`

SetEksAccessKeyId sets EksAccessKeyId field to given value.

### HasEksAccessKeyId

`func (o *DynamicSecretUpdateEks) HasEksAccessKeyId() bool`

HasEksAccessKeyId returns a boolean if a field has been set.

### GetEksAssumeRole

`func (o *DynamicSecretUpdateEks) GetEksAssumeRole() string`

GetEksAssumeRole returns the EksAssumeRole field if non-nil, zero value otherwise.

### GetEksAssumeRoleOk

`func (o *DynamicSecretUpdateEks) GetEksAssumeRoleOk() (*string, bool)`

GetEksAssumeRoleOk returns a tuple with the EksAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAssumeRole

`func (o *DynamicSecretUpdateEks) SetEksAssumeRole(v string)`

SetEksAssumeRole sets EksAssumeRole field to given value.

### HasEksAssumeRole

`func (o *DynamicSecretUpdateEks) HasEksAssumeRole() bool`

HasEksAssumeRole returns a boolean if a field has been set.

### GetEksClusterCaCert

`func (o *DynamicSecretUpdateEks) GetEksClusterCaCert() string`

GetEksClusterCaCert returns the EksClusterCaCert field if non-nil, zero value otherwise.

### GetEksClusterCaCertOk

`func (o *DynamicSecretUpdateEks) GetEksClusterCaCertOk() (*string, bool)`

GetEksClusterCaCertOk returns a tuple with the EksClusterCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterCaCert

`func (o *DynamicSecretUpdateEks) SetEksClusterCaCert(v string)`

SetEksClusterCaCert sets EksClusterCaCert field to given value.

### HasEksClusterCaCert

`func (o *DynamicSecretUpdateEks) HasEksClusterCaCert() bool`

HasEksClusterCaCert returns a boolean if a field has been set.

### GetEksClusterEndpoint

`func (o *DynamicSecretUpdateEks) GetEksClusterEndpoint() string`

GetEksClusterEndpoint returns the EksClusterEndpoint field if non-nil, zero value otherwise.

### GetEksClusterEndpointOk

`func (o *DynamicSecretUpdateEks) GetEksClusterEndpointOk() (*string, bool)`

GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterEndpoint

`func (o *DynamicSecretUpdateEks) SetEksClusterEndpoint(v string)`

SetEksClusterEndpoint sets EksClusterEndpoint field to given value.

### HasEksClusterEndpoint

`func (o *DynamicSecretUpdateEks) HasEksClusterEndpoint() bool`

HasEksClusterEndpoint returns a boolean if a field has been set.

### GetEksClusterName

`func (o *DynamicSecretUpdateEks) GetEksClusterName() string`

GetEksClusterName returns the EksClusterName field if non-nil, zero value otherwise.

### GetEksClusterNameOk

`func (o *DynamicSecretUpdateEks) GetEksClusterNameOk() (*string, bool)`

GetEksClusterNameOk returns a tuple with the EksClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterName

`func (o *DynamicSecretUpdateEks) SetEksClusterName(v string)`

SetEksClusterName sets EksClusterName field to given value.

### HasEksClusterName

`func (o *DynamicSecretUpdateEks) HasEksClusterName() bool`

HasEksClusterName returns a boolean if a field has been set.

### GetEksRegion

`func (o *DynamicSecretUpdateEks) GetEksRegion() string`

GetEksRegion returns the EksRegion field if non-nil, zero value otherwise.

### GetEksRegionOk

`func (o *DynamicSecretUpdateEks) GetEksRegionOk() (*string, bool)`

GetEksRegionOk returns a tuple with the EksRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksRegion

`func (o *DynamicSecretUpdateEks) SetEksRegion(v string)`

SetEksRegion sets EksRegion field to given value.

### HasEksRegion

`func (o *DynamicSecretUpdateEks) HasEksRegion() bool`

HasEksRegion returns a boolean if a field has been set.

### GetEksSecretAccessKey

`func (o *DynamicSecretUpdateEks) GetEksSecretAccessKey() string`

GetEksSecretAccessKey returns the EksSecretAccessKey field if non-nil, zero value otherwise.

### GetEksSecretAccessKeyOk

`func (o *DynamicSecretUpdateEks) GetEksSecretAccessKeyOk() (*string, bool)`

GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSecretAccessKey

`func (o *DynamicSecretUpdateEks) SetEksSecretAccessKey(v string)`

SetEksSecretAccessKey sets EksSecretAccessKey field to given value.

### HasEksSecretAccessKey

`func (o *DynamicSecretUpdateEks) HasEksSecretAccessKey() bool`

HasEksSecretAccessKey returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateEks) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateEks) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateEks) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateEks) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateEks) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateEks) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateEks) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateEks) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateEks) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateEks) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateEks) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateEks) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateEks) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateEks) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateEks) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessAllowPortForwading

`func (o *DynamicSecretUpdateEks) GetSecureAccessAllowPortForwading() bool`

GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field if non-nil, zero value otherwise.

### GetSecureAccessAllowPortForwadingOk

`func (o *DynamicSecretUpdateEks) GetSecureAccessAllowPortForwadingOk() (*bool, bool)`

GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowPortForwading

`func (o *DynamicSecretUpdateEks) SetSecureAccessAllowPortForwading(v bool)`

SetSecureAccessAllowPortForwading sets SecureAccessAllowPortForwading field to given value.

### HasSecureAccessAllowPortForwading

`func (o *DynamicSecretUpdateEks) HasSecureAccessAllowPortForwading() bool`

HasSecureAccessAllowPortForwading returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateEks) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretUpdateEks) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateEks) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateEks) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessClusterEndpoint

`func (o *DynamicSecretUpdateEks) GetSecureAccessClusterEndpoint() string`

GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field if non-nil, zero value otherwise.

### GetSecureAccessClusterEndpointOk

`func (o *DynamicSecretUpdateEks) GetSecureAccessClusterEndpointOk() (*string, bool)`

GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessClusterEndpoint

`func (o *DynamicSecretUpdateEks) SetSecureAccessClusterEndpoint(v string)`

SetSecureAccessClusterEndpoint sets SecureAccessClusterEndpoint field to given value.

### HasSecureAccessClusterEndpoint

`func (o *DynamicSecretUpdateEks) HasSecureAccessClusterEndpoint() bool`

HasSecureAccessClusterEndpoint returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateEks) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateEks) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateEks) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateEks) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretUpdateEks) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretUpdateEks) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretUpdateEks) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretUpdateEks) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateEks) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateEks) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateEks) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateEks) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateEks) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateEks) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateEks) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateEks) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateEks) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateEks) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateEks) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateEks) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateEks) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateEks) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateEks) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateEks) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateEks) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateEks) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateEks) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateEks) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


