# DynamicSecretCreateGke

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GkeAccountKey** | Pointer to **string** | GKE Service Account key file path | [optional] 
**GkeClusterCert** | Pointer to **string** | GKE cluster CA certificate | [optional] 
**GkeClusterEndpoint** | Pointer to **string** | GKE cluster URL endpoint | [optional] 
**GkeClusterName** | Pointer to **string** | GKE cluster name | [optional] 
**GkeServiceAccountEmail** | Pointer to **string** | GKE service account email | [optional] 
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
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateGke

`func NewDynamicSecretCreateGke(name string, ) *DynamicSecretCreateGke`

NewDynamicSecretCreateGke instantiates a new DynamicSecretCreateGke object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateGkeWithDefaults

`func NewDynamicSecretCreateGkeWithDefaults() *DynamicSecretCreateGke`

NewDynamicSecretCreateGkeWithDefaults instantiates a new DynamicSecretCreateGke object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretCreateGke) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateGke) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateGke) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateGke) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateGke) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateGke) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateGke) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateGke) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGkeAccountKey

`func (o *DynamicSecretCreateGke) GetGkeAccountKey() string`

GetGkeAccountKey returns the GkeAccountKey field if non-nil, zero value otherwise.

### GetGkeAccountKeyOk

`func (o *DynamicSecretCreateGke) GetGkeAccountKeyOk() (*string, bool)`

GetGkeAccountKeyOk returns a tuple with the GkeAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeAccountKey

`func (o *DynamicSecretCreateGke) SetGkeAccountKey(v string)`

SetGkeAccountKey sets GkeAccountKey field to given value.

### HasGkeAccountKey

`func (o *DynamicSecretCreateGke) HasGkeAccountKey() bool`

HasGkeAccountKey returns a boolean if a field has been set.

### GetGkeClusterCert

`func (o *DynamicSecretCreateGke) GetGkeClusterCert() string`

GetGkeClusterCert returns the GkeClusterCert field if non-nil, zero value otherwise.

### GetGkeClusterCertOk

`func (o *DynamicSecretCreateGke) GetGkeClusterCertOk() (*string, bool)`

GetGkeClusterCertOk returns a tuple with the GkeClusterCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterCert

`func (o *DynamicSecretCreateGke) SetGkeClusterCert(v string)`

SetGkeClusterCert sets GkeClusterCert field to given value.

### HasGkeClusterCert

`func (o *DynamicSecretCreateGke) HasGkeClusterCert() bool`

HasGkeClusterCert returns a boolean if a field has been set.

### GetGkeClusterEndpoint

`func (o *DynamicSecretCreateGke) GetGkeClusterEndpoint() string`

GetGkeClusterEndpoint returns the GkeClusterEndpoint field if non-nil, zero value otherwise.

### GetGkeClusterEndpointOk

`func (o *DynamicSecretCreateGke) GetGkeClusterEndpointOk() (*string, bool)`

GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterEndpoint

`func (o *DynamicSecretCreateGke) SetGkeClusterEndpoint(v string)`

SetGkeClusterEndpoint sets GkeClusterEndpoint field to given value.

### HasGkeClusterEndpoint

`func (o *DynamicSecretCreateGke) HasGkeClusterEndpoint() bool`

HasGkeClusterEndpoint returns a boolean if a field has been set.

### GetGkeClusterName

`func (o *DynamicSecretCreateGke) GetGkeClusterName() string`

GetGkeClusterName returns the GkeClusterName field if non-nil, zero value otherwise.

### GetGkeClusterNameOk

`func (o *DynamicSecretCreateGke) GetGkeClusterNameOk() (*string, bool)`

GetGkeClusterNameOk returns a tuple with the GkeClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterName

`func (o *DynamicSecretCreateGke) SetGkeClusterName(v string)`

SetGkeClusterName sets GkeClusterName field to given value.

### HasGkeClusterName

`func (o *DynamicSecretCreateGke) HasGkeClusterName() bool`

HasGkeClusterName returns a boolean if a field has been set.

### GetGkeServiceAccountEmail

`func (o *DynamicSecretCreateGke) GetGkeServiceAccountEmail() string`

GetGkeServiceAccountEmail returns the GkeServiceAccountEmail field if non-nil, zero value otherwise.

### GetGkeServiceAccountEmailOk

`func (o *DynamicSecretCreateGke) GetGkeServiceAccountEmailOk() (*string, bool)`

GetGkeServiceAccountEmailOk returns a tuple with the GkeServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeServiceAccountEmail

`func (o *DynamicSecretCreateGke) SetGkeServiceAccountEmail(v string)`

SetGkeServiceAccountEmail sets GkeServiceAccountEmail field to given value.

### HasGkeServiceAccountEmail

`func (o *DynamicSecretCreateGke) HasGkeServiceAccountEmail() bool`

HasGkeServiceAccountEmail returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateGke) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateGke) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateGke) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateGke) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateGke) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateGke) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateGke) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateGke) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateGke) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateGke) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateGke) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessAllowPortForwading

`func (o *DynamicSecretCreateGke) GetSecureAccessAllowPortForwading() bool`

GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field if non-nil, zero value otherwise.

### GetSecureAccessAllowPortForwadingOk

`func (o *DynamicSecretCreateGke) GetSecureAccessAllowPortForwadingOk() (*bool, bool)`

GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowPortForwading

`func (o *DynamicSecretCreateGke) SetSecureAccessAllowPortForwading(v bool)`

SetSecureAccessAllowPortForwading sets SecureAccessAllowPortForwading field to given value.

### HasSecureAccessAllowPortForwading

`func (o *DynamicSecretCreateGke) HasSecureAccessAllowPortForwading() bool`

HasSecureAccessAllowPortForwading returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateGke) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretCreateGke) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateGke) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretCreateGke) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessClusterEndpoint

`func (o *DynamicSecretCreateGke) GetSecureAccessClusterEndpoint() string`

GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field if non-nil, zero value otherwise.

### GetSecureAccessClusterEndpointOk

`func (o *DynamicSecretCreateGke) GetSecureAccessClusterEndpointOk() (*string, bool)`

GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessClusterEndpoint

`func (o *DynamicSecretCreateGke) SetSecureAccessClusterEndpoint(v string)`

SetSecureAccessClusterEndpoint sets SecureAccessClusterEndpoint field to given value.

### HasSecureAccessClusterEndpoint

`func (o *DynamicSecretCreateGke) HasSecureAccessClusterEndpoint() bool`

HasSecureAccessClusterEndpoint returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateGke) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateGke) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateGke) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateGke) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretCreateGke) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretCreateGke) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretCreateGke) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretCreateGke) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateGke) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateGke) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateGke) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateGke) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateGke) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateGke) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateGke) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateGke) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateGke) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateGke) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateGke) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateGke) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateGke) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateGke) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateGke) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateGke) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateGke) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateGke) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateGke) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateGke) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


