# GatewayCreateK8SAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | The access ID of the Kubernetes auth method | 
**ClusterApiType** | Pointer to **string** | Cluster access type. options: [native_k8s, rancher] | [optional] [default to "native_k8s"]
**ConfigEncryptionKeyName** | Pointer to **string** | Config encryption key | [optional] 
**DisableIssuerValidation** | Pointer to **string** | Disable issuer validation [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**K8sAuthType** | Pointer to **string** | K8S auth type [token/certificate]. (relevant for \&quot;native_k8s\&quot; only) | [optional] [default to "token"]
**K8sCaCert** | Pointer to **string** | The CA Certificate (base64 encoded) to use to call into the kubernetes API server | [optional] 
**K8sClientCertificate** | Pointer to **string** | Content of the k8 client certificate (PEM format) in a Base64 format (relevant for \&quot;native_k8s\&quot; only) | [optional] 
**K8sClientKey** | Pointer to **string** | Content of the k8 client private key (PEM format) in a Base64 format (relevant for \&quot;native_k8s\&quot; only) | [optional] 
**K8sHost** | **string** | The URL of the kubernetes API server | 
**K8sIssuer** | Pointer to **string** | The Kubernetes JWT issuer name. K8SIssuer is the claim that specifies who issued the Kubernetes token | [optional] [default to "kubernetes/serviceaccount"]
**Name** | **string** | K8S Auth config name | 
**RancherApiKey** | Pointer to **string** | The api key used to access the TokenReview API to validate other JWTs (relevant for \&quot;rancher\&quot; only) | [optional] 
**RancherClusterId** | Pointer to **string** | The cluster id as define in rancher (relevant for \&quot;rancher\&quot; only) | [optional] 
**SigningKey** | **string** | The private key (base64 encoded) associated with the public key defined in the Kubernetes auth | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExp** | Pointer to **int64** | Time in seconds of expiration of the Akeyless Kube Auth Method token | [optional] [default to 300]
**TokenReviewerJwt** | Pointer to **string** | A Kubernetes service account JWT used to access the TokenReview API to validate other JWTs (relevant for \&quot;native_k8s\&quot; only). If not set, the JWT submitted in the authentication process will be used to access the Kubernetes TokenReview API. | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseGwServiceAccount** | Pointer to **bool** | Use the GW&#39;s service account | [optional] 

## Methods

### NewGatewayCreateK8SAuthConfig

`func NewGatewayCreateK8SAuthConfig(accessId string, k8sHost string, name string, signingKey string, ) *GatewayCreateK8SAuthConfig`

NewGatewayCreateK8SAuthConfig instantiates a new GatewayCreateK8SAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateK8SAuthConfigWithDefaults

`func NewGatewayCreateK8SAuthConfigWithDefaults() *GatewayCreateK8SAuthConfig`

NewGatewayCreateK8SAuthConfigWithDefaults instantiates a new GatewayCreateK8SAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *GatewayCreateK8SAuthConfig) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *GatewayCreateK8SAuthConfig) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *GatewayCreateK8SAuthConfig) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetClusterApiType

`func (o *GatewayCreateK8SAuthConfig) GetClusterApiType() string`

GetClusterApiType returns the ClusterApiType field if non-nil, zero value otherwise.

### GetClusterApiTypeOk

`func (o *GatewayCreateK8SAuthConfig) GetClusterApiTypeOk() (*string, bool)`

GetClusterApiTypeOk returns a tuple with the ClusterApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterApiType

`func (o *GatewayCreateK8SAuthConfig) SetClusterApiType(v string)`

SetClusterApiType sets ClusterApiType field to given value.

### HasClusterApiType

`func (o *GatewayCreateK8SAuthConfig) HasClusterApiType() bool`

HasClusterApiType returns a boolean if a field has been set.

### GetConfigEncryptionKeyName

`func (o *GatewayCreateK8SAuthConfig) GetConfigEncryptionKeyName() string`

GetConfigEncryptionKeyName returns the ConfigEncryptionKeyName field if non-nil, zero value otherwise.

### GetConfigEncryptionKeyNameOk

`func (o *GatewayCreateK8SAuthConfig) GetConfigEncryptionKeyNameOk() (*string, bool)`

GetConfigEncryptionKeyNameOk returns a tuple with the ConfigEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigEncryptionKeyName

`func (o *GatewayCreateK8SAuthConfig) SetConfigEncryptionKeyName(v string)`

SetConfigEncryptionKeyName sets ConfigEncryptionKeyName field to given value.

### HasConfigEncryptionKeyName

`func (o *GatewayCreateK8SAuthConfig) HasConfigEncryptionKeyName() bool`

HasConfigEncryptionKeyName returns a boolean if a field has been set.

### GetDisableIssuerValidation

`func (o *GatewayCreateK8SAuthConfig) GetDisableIssuerValidation() string`

GetDisableIssuerValidation returns the DisableIssuerValidation field if non-nil, zero value otherwise.

### GetDisableIssuerValidationOk

`func (o *GatewayCreateK8SAuthConfig) GetDisableIssuerValidationOk() (*string, bool)`

GetDisableIssuerValidationOk returns a tuple with the DisableIssuerValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIssuerValidation

`func (o *GatewayCreateK8SAuthConfig) SetDisableIssuerValidation(v string)`

SetDisableIssuerValidation sets DisableIssuerValidation field to given value.

### HasDisableIssuerValidation

`func (o *GatewayCreateK8SAuthConfig) HasDisableIssuerValidation() bool`

HasDisableIssuerValidation returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateK8SAuthConfig) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateK8SAuthConfig) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateK8SAuthConfig) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateK8SAuthConfig) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sAuthType

`func (o *GatewayCreateK8SAuthConfig) GetK8sAuthType() string`

GetK8sAuthType returns the K8sAuthType field if non-nil, zero value otherwise.

### GetK8sAuthTypeOk

`func (o *GatewayCreateK8SAuthConfig) GetK8sAuthTypeOk() (*string, bool)`

GetK8sAuthTypeOk returns a tuple with the K8sAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAuthType

`func (o *GatewayCreateK8SAuthConfig) SetK8sAuthType(v string)`

SetK8sAuthType sets K8sAuthType field to given value.

### HasK8sAuthType

`func (o *GatewayCreateK8SAuthConfig) HasK8sAuthType() bool`

HasK8sAuthType returns a boolean if a field has been set.

### GetK8sCaCert

`func (o *GatewayCreateK8SAuthConfig) GetK8sCaCert() string`

GetK8sCaCert returns the K8sCaCert field if non-nil, zero value otherwise.

### GetK8sCaCertOk

`func (o *GatewayCreateK8SAuthConfig) GetK8sCaCertOk() (*string, bool)`

GetK8sCaCertOk returns a tuple with the K8sCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sCaCert

`func (o *GatewayCreateK8SAuthConfig) SetK8sCaCert(v string)`

SetK8sCaCert sets K8sCaCert field to given value.

### HasK8sCaCert

`func (o *GatewayCreateK8SAuthConfig) HasK8sCaCert() bool`

HasK8sCaCert returns a boolean if a field has been set.

### GetK8sClientCertificate

`func (o *GatewayCreateK8SAuthConfig) GetK8sClientCertificate() string`

GetK8sClientCertificate returns the K8sClientCertificate field if non-nil, zero value otherwise.

### GetK8sClientCertificateOk

`func (o *GatewayCreateK8SAuthConfig) GetK8sClientCertificateOk() (*string, bool)`

GetK8sClientCertificateOk returns a tuple with the K8sClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientCertificate

`func (o *GatewayCreateK8SAuthConfig) SetK8sClientCertificate(v string)`

SetK8sClientCertificate sets K8sClientCertificate field to given value.

### HasK8sClientCertificate

`func (o *GatewayCreateK8SAuthConfig) HasK8sClientCertificate() bool`

HasK8sClientCertificate returns a boolean if a field has been set.

### GetK8sClientKey

`func (o *GatewayCreateK8SAuthConfig) GetK8sClientKey() string`

GetK8sClientKey returns the K8sClientKey field if non-nil, zero value otherwise.

### GetK8sClientKeyOk

`func (o *GatewayCreateK8SAuthConfig) GetK8sClientKeyOk() (*string, bool)`

GetK8sClientKeyOk returns a tuple with the K8sClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientKey

`func (o *GatewayCreateK8SAuthConfig) SetK8sClientKey(v string)`

SetK8sClientKey sets K8sClientKey field to given value.

### HasK8sClientKey

`func (o *GatewayCreateK8SAuthConfig) HasK8sClientKey() bool`

HasK8sClientKey returns a boolean if a field has been set.

### GetK8sHost

`func (o *GatewayCreateK8SAuthConfig) GetK8sHost() string`

GetK8sHost returns the K8sHost field if non-nil, zero value otherwise.

### GetK8sHostOk

`func (o *GatewayCreateK8SAuthConfig) GetK8sHostOk() (*string, bool)`

GetK8sHostOk returns a tuple with the K8sHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sHost

`func (o *GatewayCreateK8SAuthConfig) SetK8sHost(v string)`

SetK8sHost sets K8sHost field to given value.


### GetK8sIssuer

`func (o *GatewayCreateK8SAuthConfig) GetK8sIssuer() string`

GetK8sIssuer returns the K8sIssuer field if non-nil, zero value otherwise.

### GetK8sIssuerOk

`func (o *GatewayCreateK8SAuthConfig) GetK8sIssuerOk() (*string, bool)`

GetK8sIssuerOk returns a tuple with the K8sIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sIssuer

`func (o *GatewayCreateK8SAuthConfig) SetK8sIssuer(v string)`

SetK8sIssuer sets K8sIssuer field to given value.

### HasK8sIssuer

`func (o *GatewayCreateK8SAuthConfig) HasK8sIssuer() bool`

HasK8sIssuer returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateK8SAuthConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateK8SAuthConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateK8SAuthConfig) SetName(v string)`

SetName sets Name field to given value.


### GetRancherApiKey

`func (o *GatewayCreateK8SAuthConfig) GetRancherApiKey() string`

GetRancherApiKey returns the RancherApiKey field if non-nil, zero value otherwise.

### GetRancherApiKeyOk

`func (o *GatewayCreateK8SAuthConfig) GetRancherApiKeyOk() (*string, bool)`

GetRancherApiKeyOk returns a tuple with the RancherApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRancherApiKey

`func (o *GatewayCreateK8SAuthConfig) SetRancherApiKey(v string)`

SetRancherApiKey sets RancherApiKey field to given value.

### HasRancherApiKey

`func (o *GatewayCreateK8SAuthConfig) HasRancherApiKey() bool`

HasRancherApiKey returns a boolean if a field has been set.

### GetRancherClusterId

`func (o *GatewayCreateK8SAuthConfig) GetRancherClusterId() string`

GetRancherClusterId returns the RancherClusterId field if non-nil, zero value otherwise.

### GetRancherClusterIdOk

`func (o *GatewayCreateK8SAuthConfig) GetRancherClusterIdOk() (*string, bool)`

GetRancherClusterIdOk returns a tuple with the RancherClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRancherClusterId

`func (o *GatewayCreateK8SAuthConfig) SetRancherClusterId(v string)`

SetRancherClusterId sets RancherClusterId field to given value.

### HasRancherClusterId

`func (o *GatewayCreateK8SAuthConfig) HasRancherClusterId() bool`

HasRancherClusterId returns a boolean if a field has been set.

### GetSigningKey

`func (o *GatewayCreateK8SAuthConfig) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *GatewayCreateK8SAuthConfig) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *GatewayCreateK8SAuthConfig) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.


### GetToken

`func (o *GatewayCreateK8SAuthConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateK8SAuthConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateK8SAuthConfig) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateK8SAuthConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenExp

`func (o *GatewayCreateK8SAuthConfig) GetTokenExp() int64`

GetTokenExp returns the TokenExp field if non-nil, zero value otherwise.

### GetTokenExpOk

`func (o *GatewayCreateK8SAuthConfig) GetTokenExpOk() (*int64, bool)`

GetTokenExpOk returns a tuple with the TokenExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExp

`func (o *GatewayCreateK8SAuthConfig) SetTokenExp(v int64)`

SetTokenExp sets TokenExp field to given value.

### HasTokenExp

`func (o *GatewayCreateK8SAuthConfig) HasTokenExp() bool`

HasTokenExp returns a boolean if a field has been set.

### GetTokenReviewerJwt

`func (o *GatewayCreateK8SAuthConfig) GetTokenReviewerJwt() string`

GetTokenReviewerJwt returns the TokenReviewerJwt field if non-nil, zero value otherwise.

### GetTokenReviewerJwtOk

`func (o *GatewayCreateK8SAuthConfig) GetTokenReviewerJwtOk() (*string, bool)`

GetTokenReviewerJwtOk returns a tuple with the TokenReviewerJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReviewerJwt

`func (o *GatewayCreateK8SAuthConfig) SetTokenReviewerJwt(v string)`

SetTokenReviewerJwt sets TokenReviewerJwt field to given value.

### HasTokenReviewerJwt

`func (o *GatewayCreateK8SAuthConfig) HasTokenReviewerJwt() bool`

HasTokenReviewerJwt returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateK8SAuthConfig) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateK8SAuthConfig) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateK8SAuthConfig) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateK8SAuthConfig) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseGwServiceAccount

`func (o *GatewayCreateK8SAuthConfig) GetUseGwServiceAccount() bool`

GetUseGwServiceAccount returns the UseGwServiceAccount field if non-nil, zero value otherwise.

### GetUseGwServiceAccountOk

`func (o *GatewayCreateK8SAuthConfig) GetUseGwServiceAccountOk() (*bool, bool)`

GetUseGwServiceAccountOk returns a tuple with the UseGwServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwServiceAccount

`func (o *GatewayCreateK8SAuthConfig) SetUseGwServiceAccount(v bool)`

SetUseGwServiceAccount sets UseGwServiceAccount field to given value.

### HasUseGwServiceAccount

`func (o *GatewayCreateK8SAuthConfig) HasUseGwServiceAccount() bool`

HasUseGwServiceAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


