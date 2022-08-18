# GatewayUpdateK8SAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | The access ID of the Kubernetes auth method | 
**ClusterApiType** | Pointer to **string** | Cluster access type. options: [native_k8s, rancher] | [optional] [default to "native_k8s"]
**ConfigEncryptionKeyName** | Pointer to **string** | Config encryption key | [optional] 
**K8sCaCert** | Pointer to **string** | The CA Certificate (base64 encoded) to use to call into the kubernetes API server | [optional] 
**K8sHost** | **string** | The URL of the kubernetes API server | 
**K8sIssuer** | Pointer to **string** | The Kubernetes JWT issuer name. If not set, kubernetes/serviceaccount will use as an issuer. | [optional] 
**Name** | **string** | K8S Auth config name | 
**NewName** | **string** | K8S Auth config new name | 
**RancherApiKey** | Pointer to **string** | The api key used to access the TokenReview API to validate other JWTs (relevant for \&quot;rancher\&quot; only) | [optional] 
**RancherClusterId** | Pointer to **string** | The cluster id as define in rancher (relevant for \&quot;rancher\&quot; only) | [optional] 
**SigningKey** | **string** | The private key (base64 encoded) associated with the public key defined in the Kubernetes auth | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExp** | Pointer to **int64** | Time in seconds of expiration of the Akeyless Kube Auth Method token | [optional] [default to 300]
**TokenReviewerJwt** | Pointer to **string** | A Kubernetes service account JWT used to access the TokenReview API to validate other JWTs (relevant for \&quot;native_k8s\&quot; only). If not set, the JWT submitted in the authentication process will be used to access the Kubernetes TokenReview API. | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateK8SAuthConfig

`func NewGatewayUpdateK8SAuthConfig(accessId string, k8sHost string, name string, newName string, signingKey string, ) *GatewayUpdateK8SAuthConfig`

NewGatewayUpdateK8SAuthConfig instantiates a new GatewayUpdateK8SAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateK8SAuthConfigWithDefaults

`func NewGatewayUpdateK8SAuthConfigWithDefaults() *GatewayUpdateK8SAuthConfig`

NewGatewayUpdateK8SAuthConfigWithDefaults instantiates a new GatewayUpdateK8SAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *GatewayUpdateK8SAuthConfig) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *GatewayUpdateK8SAuthConfig) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *GatewayUpdateK8SAuthConfig) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetClusterApiType

`func (o *GatewayUpdateK8SAuthConfig) GetClusterApiType() string`

GetClusterApiType returns the ClusterApiType field if non-nil, zero value otherwise.

### GetClusterApiTypeOk

`func (o *GatewayUpdateK8SAuthConfig) GetClusterApiTypeOk() (*string, bool)`

GetClusterApiTypeOk returns a tuple with the ClusterApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterApiType

`func (o *GatewayUpdateK8SAuthConfig) SetClusterApiType(v string)`

SetClusterApiType sets ClusterApiType field to given value.

### HasClusterApiType

`func (o *GatewayUpdateK8SAuthConfig) HasClusterApiType() bool`

HasClusterApiType returns a boolean if a field has been set.

### GetConfigEncryptionKeyName

`func (o *GatewayUpdateK8SAuthConfig) GetConfigEncryptionKeyName() string`

GetConfigEncryptionKeyName returns the ConfigEncryptionKeyName field if non-nil, zero value otherwise.

### GetConfigEncryptionKeyNameOk

`func (o *GatewayUpdateK8SAuthConfig) GetConfigEncryptionKeyNameOk() (*string, bool)`

GetConfigEncryptionKeyNameOk returns a tuple with the ConfigEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigEncryptionKeyName

`func (o *GatewayUpdateK8SAuthConfig) SetConfigEncryptionKeyName(v string)`

SetConfigEncryptionKeyName sets ConfigEncryptionKeyName field to given value.

### HasConfigEncryptionKeyName

`func (o *GatewayUpdateK8SAuthConfig) HasConfigEncryptionKeyName() bool`

HasConfigEncryptionKeyName returns a boolean if a field has been set.

### GetK8sCaCert

`func (o *GatewayUpdateK8SAuthConfig) GetK8sCaCert() string`

GetK8sCaCert returns the K8sCaCert field if non-nil, zero value otherwise.

### GetK8sCaCertOk

`func (o *GatewayUpdateK8SAuthConfig) GetK8sCaCertOk() (*string, bool)`

GetK8sCaCertOk returns a tuple with the K8sCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sCaCert

`func (o *GatewayUpdateK8SAuthConfig) SetK8sCaCert(v string)`

SetK8sCaCert sets K8sCaCert field to given value.

### HasK8sCaCert

`func (o *GatewayUpdateK8SAuthConfig) HasK8sCaCert() bool`

HasK8sCaCert returns a boolean if a field has been set.

### GetK8sHost

`func (o *GatewayUpdateK8SAuthConfig) GetK8sHost() string`

GetK8sHost returns the K8sHost field if non-nil, zero value otherwise.

### GetK8sHostOk

`func (o *GatewayUpdateK8SAuthConfig) GetK8sHostOk() (*string, bool)`

GetK8sHostOk returns a tuple with the K8sHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sHost

`func (o *GatewayUpdateK8SAuthConfig) SetK8sHost(v string)`

SetK8sHost sets K8sHost field to given value.


### GetK8sIssuer

`func (o *GatewayUpdateK8SAuthConfig) GetK8sIssuer() string`

GetK8sIssuer returns the K8sIssuer field if non-nil, zero value otherwise.

### GetK8sIssuerOk

`func (o *GatewayUpdateK8SAuthConfig) GetK8sIssuerOk() (*string, bool)`

GetK8sIssuerOk returns a tuple with the K8sIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sIssuer

`func (o *GatewayUpdateK8SAuthConfig) SetK8sIssuer(v string)`

SetK8sIssuer sets K8sIssuer field to given value.

### HasK8sIssuer

`func (o *GatewayUpdateK8SAuthConfig) HasK8sIssuer() bool`

HasK8sIssuer returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateK8SAuthConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateK8SAuthConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateK8SAuthConfig) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateK8SAuthConfig) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateK8SAuthConfig) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateK8SAuthConfig) SetNewName(v string)`

SetNewName sets NewName field to given value.


### GetRancherApiKey

`func (o *GatewayUpdateK8SAuthConfig) GetRancherApiKey() string`

GetRancherApiKey returns the RancherApiKey field if non-nil, zero value otherwise.

### GetRancherApiKeyOk

`func (o *GatewayUpdateK8SAuthConfig) GetRancherApiKeyOk() (*string, bool)`

GetRancherApiKeyOk returns a tuple with the RancherApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRancherApiKey

`func (o *GatewayUpdateK8SAuthConfig) SetRancherApiKey(v string)`

SetRancherApiKey sets RancherApiKey field to given value.

### HasRancherApiKey

`func (o *GatewayUpdateK8SAuthConfig) HasRancherApiKey() bool`

HasRancherApiKey returns a boolean if a field has been set.

### GetRancherClusterId

`func (o *GatewayUpdateK8SAuthConfig) GetRancherClusterId() string`

GetRancherClusterId returns the RancherClusterId field if non-nil, zero value otherwise.

### GetRancherClusterIdOk

`func (o *GatewayUpdateK8SAuthConfig) GetRancherClusterIdOk() (*string, bool)`

GetRancherClusterIdOk returns a tuple with the RancherClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRancherClusterId

`func (o *GatewayUpdateK8SAuthConfig) SetRancherClusterId(v string)`

SetRancherClusterId sets RancherClusterId field to given value.

### HasRancherClusterId

`func (o *GatewayUpdateK8SAuthConfig) HasRancherClusterId() bool`

HasRancherClusterId returns a boolean if a field has been set.

### GetSigningKey

`func (o *GatewayUpdateK8SAuthConfig) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *GatewayUpdateK8SAuthConfig) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *GatewayUpdateK8SAuthConfig) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.


### GetToken

`func (o *GatewayUpdateK8SAuthConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateK8SAuthConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateK8SAuthConfig) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateK8SAuthConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenExp

`func (o *GatewayUpdateK8SAuthConfig) GetTokenExp() int64`

GetTokenExp returns the TokenExp field if non-nil, zero value otherwise.

### GetTokenExpOk

`func (o *GatewayUpdateK8SAuthConfig) GetTokenExpOk() (*int64, bool)`

GetTokenExpOk returns a tuple with the TokenExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExp

`func (o *GatewayUpdateK8SAuthConfig) SetTokenExp(v int64)`

SetTokenExp sets TokenExp field to given value.

### HasTokenExp

`func (o *GatewayUpdateK8SAuthConfig) HasTokenExp() bool`

HasTokenExp returns a boolean if a field has been set.

### GetTokenReviewerJwt

`func (o *GatewayUpdateK8SAuthConfig) GetTokenReviewerJwt() string`

GetTokenReviewerJwt returns the TokenReviewerJwt field if non-nil, zero value otherwise.

### GetTokenReviewerJwtOk

`func (o *GatewayUpdateK8SAuthConfig) GetTokenReviewerJwtOk() (*string, bool)`

GetTokenReviewerJwtOk returns a tuple with the TokenReviewerJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReviewerJwt

`func (o *GatewayUpdateK8SAuthConfig) SetTokenReviewerJwt(v string)`

SetTokenReviewerJwt sets TokenReviewerJwt field to given value.

### HasTokenReviewerJwt

`func (o *GatewayUpdateK8SAuthConfig) HasTokenReviewerJwt() bool`

HasTokenReviewerJwt returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateK8SAuthConfig) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateK8SAuthConfig) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateK8SAuthConfig) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateK8SAuthConfig) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


