# TargetUpdateK8s

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**K8sAuthType** | Pointer to **string** | K8S auth type [token/certificate] | [optional] [default to "token"]
**K8sClientCertificate** | Pointer to **string** | Content of the k8 client certificate (PEM format) in a Base64 format | [optional] 
**K8sClientKey** | Pointer to **string** | Content of the k8 client private key (PEM format) in a Base64 format | [optional] 
**K8sClusterCaCert** | Pointer to **string** | K8S cluster CA certificate | [optional] 
**K8sClusterEndpoint** | Pointer to **string** | K8S cluster URL endpoint | [optional] 
**K8sClusterName** | Pointer to **string** | K8S cluster name | [optional] 
**K8sClusterToken** | Pointer to **string** | K8S cluster Bearer token | [optional] 
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseGwServiceAccount** | Pointer to **bool** | Use the GW&#39;s service account | [optional] 

## Methods

### NewTargetUpdateK8s

`func NewTargetUpdateK8s(name string, ) *TargetUpdateK8s`

NewTargetUpdateK8s instantiates a new TargetUpdateK8s object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateK8sWithDefaults

`func NewTargetUpdateK8sWithDefaults() *TargetUpdateK8s`

NewTargetUpdateK8sWithDefaults instantiates a new TargetUpdateK8s object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TargetUpdateK8s) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateK8s) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateK8s) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateK8s) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateK8s) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateK8s) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateK8s) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateK8s) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sAuthType

`func (o *TargetUpdateK8s) GetK8sAuthType() string`

GetK8sAuthType returns the K8sAuthType field if non-nil, zero value otherwise.

### GetK8sAuthTypeOk

`func (o *TargetUpdateK8s) GetK8sAuthTypeOk() (*string, bool)`

GetK8sAuthTypeOk returns a tuple with the K8sAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAuthType

`func (o *TargetUpdateK8s) SetK8sAuthType(v string)`

SetK8sAuthType sets K8sAuthType field to given value.

### HasK8sAuthType

`func (o *TargetUpdateK8s) HasK8sAuthType() bool`

HasK8sAuthType returns a boolean if a field has been set.

### GetK8sClientCertificate

`func (o *TargetUpdateK8s) GetK8sClientCertificate() string`

GetK8sClientCertificate returns the K8sClientCertificate field if non-nil, zero value otherwise.

### GetK8sClientCertificateOk

`func (o *TargetUpdateK8s) GetK8sClientCertificateOk() (*string, bool)`

GetK8sClientCertificateOk returns a tuple with the K8sClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientCertificate

`func (o *TargetUpdateK8s) SetK8sClientCertificate(v string)`

SetK8sClientCertificate sets K8sClientCertificate field to given value.

### HasK8sClientCertificate

`func (o *TargetUpdateK8s) HasK8sClientCertificate() bool`

HasK8sClientCertificate returns a boolean if a field has been set.

### GetK8sClientKey

`func (o *TargetUpdateK8s) GetK8sClientKey() string`

GetK8sClientKey returns the K8sClientKey field if non-nil, zero value otherwise.

### GetK8sClientKeyOk

`func (o *TargetUpdateK8s) GetK8sClientKeyOk() (*string, bool)`

GetK8sClientKeyOk returns a tuple with the K8sClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientKey

`func (o *TargetUpdateK8s) SetK8sClientKey(v string)`

SetK8sClientKey sets K8sClientKey field to given value.

### HasK8sClientKey

`func (o *TargetUpdateK8s) HasK8sClientKey() bool`

HasK8sClientKey returns a boolean if a field has been set.

### GetK8sClusterCaCert

`func (o *TargetUpdateK8s) GetK8sClusterCaCert() string`

GetK8sClusterCaCert returns the K8sClusterCaCert field if non-nil, zero value otherwise.

### GetK8sClusterCaCertOk

`func (o *TargetUpdateK8s) GetK8sClusterCaCertOk() (*string, bool)`

GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterCaCert

`func (o *TargetUpdateK8s) SetK8sClusterCaCert(v string)`

SetK8sClusterCaCert sets K8sClusterCaCert field to given value.

### HasK8sClusterCaCert

`func (o *TargetUpdateK8s) HasK8sClusterCaCert() bool`

HasK8sClusterCaCert returns a boolean if a field has been set.

### GetK8sClusterEndpoint

`func (o *TargetUpdateK8s) GetK8sClusterEndpoint() string`

GetK8sClusterEndpoint returns the K8sClusterEndpoint field if non-nil, zero value otherwise.

### GetK8sClusterEndpointOk

`func (o *TargetUpdateK8s) GetK8sClusterEndpointOk() (*string, bool)`

GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterEndpoint

`func (o *TargetUpdateK8s) SetK8sClusterEndpoint(v string)`

SetK8sClusterEndpoint sets K8sClusterEndpoint field to given value.

### HasK8sClusterEndpoint

`func (o *TargetUpdateK8s) HasK8sClusterEndpoint() bool`

HasK8sClusterEndpoint returns a boolean if a field has been set.

### GetK8sClusterName

`func (o *TargetUpdateK8s) GetK8sClusterName() string`

GetK8sClusterName returns the K8sClusterName field if non-nil, zero value otherwise.

### GetK8sClusterNameOk

`func (o *TargetUpdateK8s) GetK8sClusterNameOk() (*string, bool)`

GetK8sClusterNameOk returns a tuple with the K8sClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterName

`func (o *TargetUpdateK8s) SetK8sClusterName(v string)`

SetK8sClusterName sets K8sClusterName field to given value.

### HasK8sClusterName

`func (o *TargetUpdateK8s) HasK8sClusterName() bool`

HasK8sClusterName returns a boolean if a field has been set.

### GetK8sClusterToken

`func (o *TargetUpdateK8s) GetK8sClusterToken() string`

GetK8sClusterToken returns the K8sClusterToken field if non-nil, zero value otherwise.

### GetK8sClusterTokenOk

`func (o *TargetUpdateK8s) GetK8sClusterTokenOk() (*string, bool)`

GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterToken

`func (o *TargetUpdateK8s) SetK8sClusterToken(v string)`

SetK8sClusterToken sets K8sClusterToken field to given value.

### HasK8sClusterToken

`func (o *TargetUpdateK8s) HasK8sClusterToken() bool`

HasK8sClusterToken returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateK8s) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateK8s) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateK8s) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateK8s) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateK8s) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateK8s) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateK8s) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateK8s) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateK8s) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateK8s) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateK8s) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateK8s) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateK8s) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateK8s) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateK8s) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateK8s) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateK8s) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateK8s) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateK8s) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateK8s) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateK8s) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateK8s) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateK8s) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateK8s) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateK8s) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateK8s) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateK8s) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseGwServiceAccount

`func (o *TargetUpdateK8s) GetUseGwServiceAccount() bool`

GetUseGwServiceAccount returns the UseGwServiceAccount field if non-nil, zero value otherwise.

### GetUseGwServiceAccountOk

`func (o *TargetUpdateK8s) GetUseGwServiceAccountOk() (*bool, bool)`

GetUseGwServiceAccountOk returns a tuple with the UseGwServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwServiceAccount

`func (o *TargetUpdateK8s) SetUseGwServiceAccount(v bool)`

SetUseGwServiceAccount sets UseGwServiceAccount field to given value.

### HasUseGwServiceAccount

`func (o *TargetUpdateK8s) HasUseGwServiceAccount() bool`

HasUseGwServiceAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


