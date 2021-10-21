# GatewayGetK8SAuthConfigOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmTokenExpiration** | Pointer to **int64** | AuthMethodTokenExpiration is time in seconds of expiration of the Akeyless Kube Auth Method token | [optional] 
**AuthMethodAccessId** | Pointer to **string** | AuthMethodAccessId of the Kubernetes auth method | [optional] 
**AuthMethodPrvKeyPem** | Pointer to **string** | AuthMethodSigningKey is the private key (in base64 of the PEM format) associated with the public key defined in the Kubernetes auth method, that used to sign the internal token for the Akeyless Kubernetes Auth Method | [optional] 
**DisableIssValidation** | Pointer to **bool** | DisableISSValidation is optional parameter to disable ISS validation | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**K8sCaCert** | Pointer to **string** | K8SCACert is the CA Cert to use to call into the kubernetes API | [optional] 
**K8sHost** | Pointer to **string** | K8SHost is the url string for the kubernetes API | [optional] 
**K8sIssuer** | Pointer to **string** | K8SIssuer is the claim that specifies who issued the Kubernetes token | [optional] 
**K8sPubKeysPem** | Pointer to **[]string** | K8SPublicKeysPEM is the list of public key in PEM format | [optional] 
**K8sTokenReviewerJwt** | Pointer to **string** | K8STokenReviewerJW\&quot;K8S Auth config %v successfully created\\n\&quot;, clictx.Color().Bold(c.K8SAuthConfigName)T is the bearer to use during the TokenReview API call | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProtectionKey** | Pointer to **string** |  | [optional] 
**UseLocalCaJwt** | Pointer to **bool** | UseLocalCAJwt is an optional parameter to set defaulting to using the local CA cert and service account jwt when running in a Kubernetes pod | [optional] 

## Methods

### NewGatewayGetK8SAuthConfigOutput

`func NewGatewayGetK8SAuthConfigOutput() *GatewayGetK8SAuthConfigOutput`

NewGatewayGetK8SAuthConfigOutput instantiates a new GatewayGetK8SAuthConfigOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayGetK8SAuthConfigOutputWithDefaults

`func NewGatewayGetK8SAuthConfigOutputWithDefaults() *GatewayGetK8SAuthConfigOutput`

NewGatewayGetK8SAuthConfigOutputWithDefaults instantiates a new GatewayGetK8SAuthConfigOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmTokenExpiration

`func (o *GatewayGetK8SAuthConfigOutput) GetAmTokenExpiration() int64`

GetAmTokenExpiration returns the AmTokenExpiration field if non-nil, zero value otherwise.

### GetAmTokenExpirationOk

`func (o *GatewayGetK8SAuthConfigOutput) GetAmTokenExpirationOk() (*int64, bool)`

GetAmTokenExpirationOk returns a tuple with the AmTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmTokenExpiration

`func (o *GatewayGetK8SAuthConfigOutput) SetAmTokenExpiration(v int64)`

SetAmTokenExpiration sets AmTokenExpiration field to given value.

### HasAmTokenExpiration

`func (o *GatewayGetK8SAuthConfigOutput) HasAmTokenExpiration() bool`

HasAmTokenExpiration returns a boolean if a field has been set.

### GetAuthMethodAccessId

`func (o *GatewayGetK8SAuthConfigOutput) GetAuthMethodAccessId() string`

GetAuthMethodAccessId returns the AuthMethodAccessId field if non-nil, zero value otherwise.

### GetAuthMethodAccessIdOk

`func (o *GatewayGetK8SAuthConfigOutput) GetAuthMethodAccessIdOk() (*string, bool)`

GetAuthMethodAccessIdOk returns a tuple with the AuthMethodAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodAccessId

`func (o *GatewayGetK8SAuthConfigOutput) SetAuthMethodAccessId(v string)`

SetAuthMethodAccessId sets AuthMethodAccessId field to given value.

### HasAuthMethodAccessId

`func (o *GatewayGetK8SAuthConfigOutput) HasAuthMethodAccessId() bool`

HasAuthMethodAccessId returns a boolean if a field has been set.

### GetAuthMethodPrvKeyPem

`func (o *GatewayGetK8SAuthConfigOutput) GetAuthMethodPrvKeyPem() string`

GetAuthMethodPrvKeyPem returns the AuthMethodPrvKeyPem field if non-nil, zero value otherwise.

### GetAuthMethodPrvKeyPemOk

`func (o *GatewayGetK8SAuthConfigOutput) GetAuthMethodPrvKeyPemOk() (*string, bool)`

GetAuthMethodPrvKeyPemOk returns a tuple with the AuthMethodPrvKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodPrvKeyPem

`func (o *GatewayGetK8SAuthConfigOutput) SetAuthMethodPrvKeyPem(v string)`

SetAuthMethodPrvKeyPem sets AuthMethodPrvKeyPem field to given value.

### HasAuthMethodPrvKeyPem

`func (o *GatewayGetK8SAuthConfigOutput) HasAuthMethodPrvKeyPem() bool`

HasAuthMethodPrvKeyPem returns a boolean if a field has been set.

### GetDisableIssValidation

`func (o *GatewayGetK8SAuthConfigOutput) GetDisableIssValidation() bool`

GetDisableIssValidation returns the DisableIssValidation field if non-nil, zero value otherwise.

### GetDisableIssValidationOk

`func (o *GatewayGetK8SAuthConfigOutput) GetDisableIssValidationOk() (*bool, bool)`

GetDisableIssValidationOk returns a tuple with the DisableIssValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIssValidation

`func (o *GatewayGetK8SAuthConfigOutput) SetDisableIssValidation(v bool)`

SetDisableIssValidation sets DisableIssValidation field to given value.

### HasDisableIssValidation

`func (o *GatewayGetK8SAuthConfigOutput) HasDisableIssValidation() bool`

HasDisableIssValidation returns a boolean if a field has been set.

### GetId

`func (o *GatewayGetK8SAuthConfigOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayGetK8SAuthConfigOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayGetK8SAuthConfigOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayGetK8SAuthConfigOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetK8sCaCert

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sCaCert() string`

GetK8sCaCert returns the K8sCaCert field if non-nil, zero value otherwise.

### GetK8sCaCertOk

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sCaCertOk() (*string, bool)`

GetK8sCaCertOk returns a tuple with the K8sCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sCaCert

`func (o *GatewayGetK8SAuthConfigOutput) SetK8sCaCert(v string)`

SetK8sCaCert sets K8sCaCert field to given value.

### HasK8sCaCert

`func (o *GatewayGetK8SAuthConfigOutput) HasK8sCaCert() bool`

HasK8sCaCert returns a boolean if a field has been set.

### GetK8sHost

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sHost() string`

GetK8sHost returns the K8sHost field if non-nil, zero value otherwise.

### GetK8sHostOk

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sHostOk() (*string, bool)`

GetK8sHostOk returns a tuple with the K8sHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sHost

`func (o *GatewayGetK8SAuthConfigOutput) SetK8sHost(v string)`

SetK8sHost sets K8sHost field to given value.

### HasK8sHost

`func (o *GatewayGetK8SAuthConfigOutput) HasK8sHost() bool`

HasK8sHost returns a boolean if a field has been set.

### GetK8sIssuer

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sIssuer() string`

GetK8sIssuer returns the K8sIssuer field if non-nil, zero value otherwise.

### GetK8sIssuerOk

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sIssuerOk() (*string, bool)`

GetK8sIssuerOk returns a tuple with the K8sIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sIssuer

`func (o *GatewayGetK8SAuthConfigOutput) SetK8sIssuer(v string)`

SetK8sIssuer sets K8sIssuer field to given value.

### HasK8sIssuer

`func (o *GatewayGetK8SAuthConfigOutput) HasK8sIssuer() bool`

HasK8sIssuer returns a boolean if a field has been set.

### GetK8sPubKeysPem

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sPubKeysPem() []string`

GetK8sPubKeysPem returns the K8sPubKeysPem field if non-nil, zero value otherwise.

### GetK8sPubKeysPemOk

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sPubKeysPemOk() (*[]string, bool)`

GetK8sPubKeysPemOk returns a tuple with the K8sPubKeysPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPubKeysPem

`func (o *GatewayGetK8SAuthConfigOutput) SetK8sPubKeysPem(v []string)`

SetK8sPubKeysPem sets K8sPubKeysPem field to given value.

### HasK8sPubKeysPem

`func (o *GatewayGetK8SAuthConfigOutput) HasK8sPubKeysPem() bool`

HasK8sPubKeysPem returns a boolean if a field has been set.

### GetK8sTokenReviewerJwt

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sTokenReviewerJwt() string`

GetK8sTokenReviewerJwt returns the K8sTokenReviewerJwt field if non-nil, zero value otherwise.

### GetK8sTokenReviewerJwtOk

`func (o *GatewayGetK8SAuthConfigOutput) GetK8sTokenReviewerJwtOk() (*string, bool)`

GetK8sTokenReviewerJwtOk returns a tuple with the K8sTokenReviewerJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sTokenReviewerJwt

`func (o *GatewayGetK8SAuthConfigOutput) SetK8sTokenReviewerJwt(v string)`

SetK8sTokenReviewerJwt sets K8sTokenReviewerJwt field to given value.

### HasK8sTokenReviewerJwt

`func (o *GatewayGetK8SAuthConfigOutput) HasK8sTokenReviewerJwt() bool`

HasK8sTokenReviewerJwt returns a boolean if a field has been set.

### GetName

`func (o *GatewayGetK8SAuthConfigOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayGetK8SAuthConfigOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayGetK8SAuthConfigOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayGetK8SAuthConfigOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtectionKey

`func (o *GatewayGetK8SAuthConfigOutput) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *GatewayGetK8SAuthConfigOutput) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *GatewayGetK8SAuthConfigOutput) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *GatewayGetK8SAuthConfigOutput) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetUseLocalCaJwt

`func (o *GatewayGetK8SAuthConfigOutput) GetUseLocalCaJwt() bool`

GetUseLocalCaJwt returns the UseLocalCaJwt field if non-nil, zero value otherwise.

### GetUseLocalCaJwtOk

`func (o *GatewayGetK8SAuthConfigOutput) GetUseLocalCaJwtOk() (*bool, bool)`

GetUseLocalCaJwtOk returns a tuple with the UseLocalCaJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalCaJwt

`func (o *GatewayGetK8SAuthConfigOutput) SetUseLocalCaJwt(v bool)`

SetUseLocalCaJwt sets UseLocalCaJwt field to given value.

### HasUseLocalCaJwt

`func (o *GatewayGetK8SAuthConfigOutput) HasUseLocalCaJwt() bool`

HasUseLocalCaJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


