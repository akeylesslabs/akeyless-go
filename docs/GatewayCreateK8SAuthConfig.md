# GatewayCreateK8SAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | The access ID of the Kubernetes auth method | 
**ConfigEncryptionKeyName** | Pointer to **string** | Config encryption key | [optional] 
**K8sCaCert** | Pointer to **string** | The CA Cert (in PEM format) to use to call into the kubernetes API server | [optional] 
**K8sHost** | **string** | The URL of the kubernetes API server | 
**K8sIssuer** | Pointer to **string** | The Kubernetes JWT issuer name. If not set, kubernetes/serviceaccount will use as an issuer. | [optional] 
**Name** | **string** | K8S Auth config name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**SigningKey** | **string** | The private key (in base64 encoded of the PEM format) associated with the public key defined in the Kubernetes auth | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExp** | Pointer to **int64** | Time in seconds of expiration of the Akeyless Kube Auth Method token | [optional] 
**TokenReviewerJwt** | Pointer to **string** | A Kubernetes service account JWT used to access the TokenReview API to validate other JWTs. If not set, the JWT submitted in the authentication process will be used to access the Kubernetes TokenReview API. | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

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


### GetPassword

`func (o *GatewayCreateK8SAuthConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateK8SAuthConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateK8SAuthConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateK8SAuthConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

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

### GetUsername

`func (o *GatewayCreateK8SAuthConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateK8SAuthConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateK8SAuthConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateK8SAuthConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


