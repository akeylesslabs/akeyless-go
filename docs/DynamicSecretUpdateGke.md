# DynamicSecretUpdateGke

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
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretUpdateGke

`func NewDynamicSecretUpdateGke(name string, ) *DynamicSecretUpdateGke`

NewDynamicSecretUpdateGke instantiates a new DynamicSecretUpdateGke object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateGkeWithDefaults

`func NewDynamicSecretUpdateGkeWithDefaults() *DynamicSecretUpdateGke`

NewDynamicSecretUpdateGkeWithDefaults instantiates a new DynamicSecretUpdateGke object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretUpdateGke) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateGke) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateGke) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateGke) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateGke) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateGke) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateGke) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateGke) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGkeAccountKey

`func (o *DynamicSecretUpdateGke) GetGkeAccountKey() string`

GetGkeAccountKey returns the GkeAccountKey field if non-nil, zero value otherwise.

### GetGkeAccountKeyOk

`func (o *DynamicSecretUpdateGke) GetGkeAccountKeyOk() (*string, bool)`

GetGkeAccountKeyOk returns a tuple with the GkeAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeAccountKey

`func (o *DynamicSecretUpdateGke) SetGkeAccountKey(v string)`

SetGkeAccountKey sets GkeAccountKey field to given value.

### HasGkeAccountKey

`func (o *DynamicSecretUpdateGke) HasGkeAccountKey() bool`

HasGkeAccountKey returns a boolean if a field has been set.

### GetGkeClusterCert

`func (o *DynamicSecretUpdateGke) GetGkeClusterCert() string`

GetGkeClusterCert returns the GkeClusterCert field if non-nil, zero value otherwise.

### GetGkeClusterCertOk

`func (o *DynamicSecretUpdateGke) GetGkeClusterCertOk() (*string, bool)`

GetGkeClusterCertOk returns a tuple with the GkeClusterCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterCert

`func (o *DynamicSecretUpdateGke) SetGkeClusterCert(v string)`

SetGkeClusterCert sets GkeClusterCert field to given value.

### HasGkeClusterCert

`func (o *DynamicSecretUpdateGke) HasGkeClusterCert() bool`

HasGkeClusterCert returns a boolean if a field has been set.

### GetGkeClusterEndpoint

`func (o *DynamicSecretUpdateGke) GetGkeClusterEndpoint() string`

GetGkeClusterEndpoint returns the GkeClusterEndpoint field if non-nil, zero value otherwise.

### GetGkeClusterEndpointOk

`func (o *DynamicSecretUpdateGke) GetGkeClusterEndpointOk() (*string, bool)`

GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterEndpoint

`func (o *DynamicSecretUpdateGke) SetGkeClusterEndpoint(v string)`

SetGkeClusterEndpoint sets GkeClusterEndpoint field to given value.

### HasGkeClusterEndpoint

`func (o *DynamicSecretUpdateGke) HasGkeClusterEndpoint() bool`

HasGkeClusterEndpoint returns a boolean if a field has been set.

### GetGkeClusterName

`func (o *DynamicSecretUpdateGke) GetGkeClusterName() string`

GetGkeClusterName returns the GkeClusterName field if non-nil, zero value otherwise.

### GetGkeClusterNameOk

`func (o *DynamicSecretUpdateGke) GetGkeClusterNameOk() (*string, bool)`

GetGkeClusterNameOk returns a tuple with the GkeClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterName

`func (o *DynamicSecretUpdateGke) SetGkeClusterName(v string)`

SetGkeClusterName sets GkeClusterName field to given value.

### HasGkeClusterName

`func (o *DynamicSecretUpdateGke) HasGkeClusterName() bool`

HasGkeClusterName returns a boolean if a field has been set.

### GetGkeServiceAccountEmail

`func (o *DynamicSecretUpdateGke) GetGkeServiceAccountEmail() string`

GetGkeServiceAccountEmail returns the GkeServiceAccountEmail field if non-nil, zero value otherwise.

### GetGkeServiceAccountEmailOk

`func (o *DynamicSecretUpdateGke) GetGkeServiceAccountEmailOk() (*string, bool)`

GetGkeServiceAccountEmailOk returns a tuple with the GkeServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeServiceAccountEmail

`func (o *DynamicSecretUpdateGke) SetGkeServiceAccountEmail(v string)`

SetGkeServiceAccountEmail sets GkeServiceAccountEmail field to given value.

### HasGkeServiceAccountEmail

`func (o *DynamicSecretUpdateGke) HasGkeServiceAccountEmail() bool`

HasGkeServiceAccountEmail returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateGke) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateGke) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateGke) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateGke) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateGke) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateGke) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateGke) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateGke) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateGke) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateGke) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateGke) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateGke) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateGke) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateGke) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateGke) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessAllowPortForwading

`func (o *DynamicSecretUpdateGke) GetSecureAccessAllowPortForwading() bool`

GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field if non-nil, zero value otherwise.

### GetSecureAccessAllowPortForwadingOk

`func (o *DynamicSecretUpdateGke) GetSecureAccessAllowPortForwadingOk() (*bool, bool)`

GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowPortForwading

`func (o *DynamicSecretUpdateGke) SetSecureAccessAllowPortForwading(v bool)`

SetSecureAccessAllowPortForwading sets SecureAccessAllowPortForwading field to given value.

### HasSecureAccessAllowPortForwading

`func (o *DynamicSecretUpdateGke) HasSecureAccessAllowPortForwading() bool`

HasSecureAccessAllowPortForwading returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateGke) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretUpdateGke) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateGke) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretUpdateGke) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessClusterEndpoint

`func (o *DynamicSecretUpdateGke) GetSecureAccessClusterEndpoint() string`

GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field if non-nil, zero value otherwise.

### GetSecureAccessClusterEndpointOk

`func (o *DynamicSecretUpdateGke) GetSecureAccessClusterEndpointOk() (*string, bool)`

GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessClusterEndpoint

`func (o *DynamicSecretUpdateGke) SetSecureAccessClusterEndpoint(v string)`

SetSecureAccessClusterEndpoint sets SecureAccessClusterEndpoint field to given value.

### HasSecureAccessClusterEndpoint

`func (o *DynamicSecretUpdateGke) HasSecureAccessClusterEndpoint() bool`

HasSecureAccessClusterEndpoint returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretUpdateGke) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretUpdateGke) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretUpdateGke) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretUpdateGke) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretUpdateGke) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretUpdateGke) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretUpdateGke) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretUpdateGke) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateGke) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateGke) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateGke) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateGke) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateGke) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateGke) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateGke) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateGke) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateGke) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateGke) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateGke) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateGke) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateGke) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateGke) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateGke) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateGke) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateGke) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateGke) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateGke) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateGke) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


