# GatewayPartialUpdateK8SAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseDefaultIdentityBool** | Pointer to **bool** |  | [optional] 
**AccessId** | Pointer to **string** | The access ID of the Kubernetes auth method | [optional] 
**ConfigEncryptionKeyName** | Pointer to **string** | Config encryption key | [optional] 
**DisableIssuerValidation** | Pointer to **string** | Disable issuer validation [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**K8sAuthType** | Pointer to **string** | K8S auth type [token/certificate]. (relevant for \&quot;native_k8s\&quot; only) | [optional] [default to "token"]
**K8sCaCert** | Pointer to **string** | The CA Certificate (base64 encoded) to use to call into the kubernetes API server | [optional] 
**K8sClientCertificate** | Pointer to **string** | Content of the k8 client certificate (PEM format) in a Base64 format (relevant for \&quot;native_k8s\&quot; only) | [optional] 
**K8sClientKey** | Pointer to **string** | Content of the k8 client private key (PEM format) in a Base64 format (relevant for \&quot;native_k8s\&quot; only) | [optional] 
**K8sHost** | Pointer to **string** | The URL of the kubernetes API server | [optional] 
**K8sIssuer** | Pointer to **string** | The Kubernetes JWT issuer name. K8SIssuer is the claim that specifies who issued the Kubernetes token | [optional] 
**Name** | Pointer to **string** | K8S Auth config name | [optional] 
**NewName** | Pointer to **string** | K8S Auth config new name | [optional] 
**RancherApiKey** | Pointer to **string** | The api key used to access the TokenReview API to validate other JWTs (relevant for \&quot;rancher\&quot; only) | [optional] 
**RancherClusterId** | Pointer to **string** | The cluster id as define in rancher (relevant for \&quot;rancher\&quot; only) | [optional] 
**SigningKey** | Pointer to **string** | The private key (base64 encoded) associated with the public key defined in the Kubernetes auth | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExp** | Pointer to **int64** | Time in seconds of expiration of the Akeyless Kube Auth Method token | [optional] 
**TokenReviewerJwt** | Pointer to **string** | A Kubernetes service account JWT used to access the TokenReview API to validate other JWTs (relevant for \&quot;native_k8s\&quot; only). If not set, the JWT submitted in the authentication process will be used to access the Kubernetes TokenReview API. | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseGwServiceAccount** | Pointer to **string** | Use the GW&#39;s service account | [optional] 

## Methods

### NewGatewayPartialUpdateK8SAuthConfig

`func NewGatewayPartialUpdateK8SAuthConfig() *GatewayPartialUpdateK8SAuthConfig`

NewGatewayPartialUpdateK8SAuthConfig instantiates a new GatewayPartialUpdateK8SAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPartialUpdateK8SAuthConfigWithDefaults

`func NewGatewayPartialUpdateK8SAuthConfigWithDefaults() *GatewayPartialUpdateK8SAuthConfig`

NewGatewayPartialUpdateK8SAuthConfigWithDefaults instantiates a new GatewayPartialUpdateK8SAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseDefaultIdentityBool

`func (o *GatewayPartialUpdateK8SAuthConfig) GetUseDefaultIdentityBool() bool`

GetUseDefaultIdentityBool returns the UseDefaultIdentityBool field if non-nil, zero value otherwise.

### GetUseDefaultIdentityBoolOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetUseDefaultIdentityBoolOk() (*bool, bool)`

GetUseDefaultIdentityBoolOk returns a tuple with the UseDefaultIdentityBool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultIdentityBool

`func (o *GatewayPartialUpdateK8SAuthConfig) SetUseDefaultIdentityBool(v bool)`

SetUseDefaultIdentityBool sets UseDefaultIdentityBool field to given value.

### HasUseDefaultIdentityBool

`func (o *GatewayPartialUpdateK8SAuthConfig) HasUseDefaultIdentityBool() bool`

HasUseDefaultIdentityBool returns a boolean if a field has been set.

### GetAccessId

`func (o *GatewayPartialUpdateK8SAuthConfig) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *GatewayPartialUpdateK8SAuthConfig) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *GatewayPartialUpdateK8SAuthConfig) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetConfigEncryptionKeyName

`func (o *GatewayPartialUpdateK8SAuthConfig) GetConfigEncryptionKeyName() string`

GetConfigEncryptionKeyName returns the ConfigEncryptionKeyName field if non-nil, zero value otherwise.

### GetConfigEncryptionKeyNameOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetConfigEncryptionKeyNameOk() (*string, bool)`

GetConfigEncryptionKeyNameOk returns a tuple with the ConfigEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigEncryptionKeyName

`func (o *GatewayPartialUpdateK8SAuthConfig) SetConfigEncryptionKeyName(v string)`

SetConfigEncryptionKeyName sets ConfigEncryptionKeyName field to given value.

### HasConfigEncryptionKeyName

`func (o *GatewayPartialUpdateK8SAuthConfig) HasConfigEncryptionKeyName() bool`

HasConfigEncryptionKeyName returns a boolean if a field has been set.

### GetDisableIssuerValidation

`func (o *GatewayPartialUpdateK8SAuthConfig) GetDisableIssuerValidation() string`

GetDisableIssuerValidation returns the DisableIssuerValidation field if non-nil, zero value otherwise.

### GetDisableIssuerValidationOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetDisableIssuerValidationOk() (*string, bool)`

GetDisableIssuerValidationOk returns a tuple with the DisableIssuerValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIssuerValidation

`func (o *GatewayPartialUpdateK8SAuthConfig) SetDisableIssuerValidation(v string)`

SetDisableIssuerValidation sets DisableIssuerValidation field to given value.

### HasDisableIssuerValidation

`func (o *GatewayPartialUpdateK8SAuthConfig) HasDisableIssuerValidation() bool`

HasDisableIssuerValidation returns a boolean if a field has been set.

### GetJson

`func (o *GatewayPartialUpdateK8SAuthConfig) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayPartialUpdateK8SAuthConfig) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayPartialUpdateK8SAuthConfig) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sAuthType

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sAuthType() string`

GetK8sAuthType returns the K8sAuthType field if non-nil, zero value otherwise.

### GetK8sAuthTypeOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sAuthTypeOk() (*string, bool)`

GetK8sAuthTypeOk returns a tuple with the K8sAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAuthType

`func (o *GatewayPartialUpdateK8SAuthConfig) SetK8sAuthType(v string)`

SetK8sAuthType sets K8sAuthType field to given value.

### HasK8sAuthType

`func (o *GatewayPartialUpdateK8SAuthConfig) HasK8sAuthType() bool`

HasK8sAuthType returns a boolean if a field has been set.

### GetK8sCaCert

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sCaCert() string`

GetK8sCaCert returns the K8sCaCert field if non-nil, zero value otherwise.

### GetK8sCaCertOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sCaCertOk() (*string, bool)`

GetK8sCaCertOk returns a tuple with the K8sCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sCaCert

`func (o *GatewayPartialUpdateK8SAuthConfig) SetK8sCaCert(v string)`

SetK8sCaCert sets K8sCaCert field to given value.

### HasK8sCaCert

`func (o *GatewayPartialUpdateK8SAuthConfig) HasK8sCaCert() bool`

HasK8sCaCert returns a boolean if a field has been set.

### GetK8sClientCertificate

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sClientCertificate() string`

GetK8sClientCertificate returns the K8sClientCertificate field if non-nil, zero value otherwise.

### GetK8sClientCertificateOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sClientCertificateOk() (*string, bool)`

GetK8sClientCertificateOk returns a tuple with the K8sClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientCertificate

`func (o *GatewayPartialUpdateK8SAuthConfig) SetK8sClientCertificate(v string)`

SetK8sClientCertificate sets K8sClientCertificate field to given value.

### HasK8sClientCertificate

`func (o *GatewayPartialUpdateK8SAuthConfig) HasK8sClientCertificate() bool`

HasK8sClientCertificate returns a boolean if a field has been set.

### GetK8sClientKey

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sClientKey() string`

GetK8sClientKey returns the K8sClientKey field if non-nil, zero value otherwise.

### GetK8sClientKeyOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sClientKeyOk() (*string, bool)`

GetK8sClientKeyOk returns a tuple with the K8sClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientKey

`func (o *GatewayPartialUpdateK8SAuthConfig) SetK8sClientKey(v string)`

SetK8sClientKey sets K8sClientKey field to given value.

### HasK8sClientKey

`func (o *GatewayPartialUpdateK8SAuthConfig) HasK8sClientKey() bool`

HasK8sClientKey returns a boolean if a field has been set.

### GetK8sHost

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sHost() string`

GetK8sHost returns the K8sHost field if non-nil, zero value otherwise.

### GetK8sHostOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sHostOk() (*string, bool)`

GetK8sHostOk returns a tuple with the K8sHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sHost

`func (o *GatewayPartialUpdateK8SAuthConfig) SetK8sHost(v string)`

SetK8sHost sets K8sHost field to given value.

### HasK8sHost

`func (o *GatewayPartialUpdateK8SAuthConfig) HasK8sHost() bool`

HasK8sHost returns a boolean if a field has been set.

### GetK8sIssuer

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sIssuer() string`

GetK8sIssuer returns the K8sIssuer field if non-nil, zero value otherwise.

### GetK8sIssuerOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetK8sIssuerOk() (*string, bool)`

GetK8sIssuerOk returns a tuple with the K8sIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sIssuer

`func (o *GatewayPartialUpdateK8SAuthConfig) SetK8sIssuer(v string)`

SetK8sIssuer sets K8sIssuer field to given value.

### HasK8sIssuer

`func (o *GatewayPartialUpdateK8SAuthConfig) HasK8sIssuer() bool`

HasK8sIssuer returns a boolean if a field has been set.

### GetName

`func (o *GatewayPartialUpdateK8SAuthConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayPartialUpdateK8SAuthConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayPartialUpdateK8SAuthConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewName

`func (o *GatewayPartialUpdateK8SAuthConfig) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayPartialUpdateK8SAuthConfig) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayPartialUpdateK8SAuthConfig) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRancherApiKey

`func (o *GatewayPartialUpdateK8SAuthConfig) GetRancherApiKey() string`

GetRancherApiKey returns the RancherApiKey field if non-nil, zero value otherwise.

### GetRancherApiKeyOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetRancherApiKeyOk() (*string, bool)`

GetRancherApiKeyOk returns a tuple with the RancherApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRancherApiKey

`func (o *GatewayPartialUpdateK8SAuthConfig) SetRancherApiKey(v string)`

SetRancherApiKey sets RancherApiKey field to given value.

### HasRancherApiKey

`func (o *GatewayPartialUpdateK8SAuthConfig) HasRancherApiKey() bool`

HasRancherApiKey returns a boolean if a field has been set.

### GetRancherClusterId

`func (o *GatewayPartialUpdateK8SAuthConfig) GetRancherClusterId() string`

GetRancherClusterId returns the RancherClusterId field if non-nil, zero value otherwise.

### GetRancherClusterIdOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetRancherClusterIdOk() (*string, bool)`

GetRancherClusterIdOk returns a tuple with the RancherClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRancherClusterId

`func (o *GatewayPartialUpdateK8SAuthConfig) SetRancherClusterId(v string)`

SetRancherClusterId sets RancherClusterId field to given value.

### HasRancherClusterId

`func (o *GatewayPartialUpdateK8SAuthConfig) HasRancherClusterId() bool`

HasRancherClusterId returns a boolean if a field has been set.

### GetSigningKey

`func (o *GatewayPartialUpdateK8SAuthConfig) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *GatewayPartialUpdateK8SAuthConfig) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.

### HasSigningKey

`func (o *GatewayPartialUpdateK8SAuthConfig) HasSigningKey() bool`

HasSigningKey returns a boolean if a field has been set.

### GetToken

`func (o *GatewayPartialUpdateK8SAuthConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayPartialUpdateK8SAuthConfig) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayPartialUpdateK8SAuthConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenExp

`func (o *GatewayPartialUpdateK8SAuthConfig) GetTokenExp() int64`

GetTokenExp returns the TokenExp field if non-nil, zero value otherwise.

### GetTokenExpOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetTokenExpOk() (*int64, bool)`

GetTokenExpOk returns a tuple with the TokenExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExp

`func (o *GatewayPartialUpdateK8SAuthConfig) SetTokenExp(v int64)`

SetTokenExp sets TokenExp field to given value.

### HasTokenExp

`func (o *GatewayPartialUpdateK8SAuthConfig) HasTokenExp() bool`

HasTokenExp returns a boolean if a field has been set.

### GetTokenReviewerJwt

`func (o *GatewayPartialUpdateK8SAuthConfig) GetTokenReviewerJwt() string`

GetTokenReviewerJwt returns the TokenReviewerJwt field if non-nil, zero value otherwise.

### GetTokenReviewerJwtOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetTokenReviewerJwtOk() (*string, bool)`

GetTokenReviewerJwtOk returns a tuple with the TokenReviewerJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReviewerJwt

`func (o *GatewayPartialUpdateK8SAuthConfig) SetTokenReviewerJwt(v string)`

SetTokenReviewerJwt sets TokenReviewerJwt field to given value.

### HasTokenReviewerJwt

`func (o *GatewayPartialUpdateK8SAuthConfig) HasTokenReviewerJwt() bool`

HasTokenReviewerJwt returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayPartialUpdateK8SAuthConfig) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayPartialUpdateK8SAuthConfig) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayPartialUpdateK8SAuthConfig) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseGwServiceAccount

`func (o *GatewayPartialUpdateK8SAuthConfig) GetUseGwServiceAccount() string`

GetUseGwServiceAccount returns the UseGwServiceAccount field if non-nil, zero value otherwise.

### GetUseGwServiceAccountOk

`func (o *GatewayPartialUpdateK8SAuthConfig) GetUseGwServiceAccountOk() (*string, bool)`

GetUseGwServiceAccountOk returns a tuple with the UseGwServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwServiceAccount

`func (o *GatewayPartialUpdateK8SAuthConfig) SetUseGwServiceAccount(v string)`

SetUseGwServiceAccount sets UseGwServiceAccount field to given value.

### HasUseGwServiceAccount

`func (o *GatewayPartialUpdateK8SAuthConfig) HasUseGwServiceAccount() bool`

HasUseGwServiceAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


