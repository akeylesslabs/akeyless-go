# K8SAuth

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

### NewK8SAuth

`func NewK8SAuth() *K8SAuth`

NewK8SAuth instantiates a new K8SAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8SAuthWithDefaults

`func NewK8SAuthWithDefaults() *K8SAuth`

NewK8SAuthWithDefaults instantiates a new K8SAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmTokenExpiration

`func (o *K8SAuth) GetAmTokenExpiration() int64`

GetAmTokenExpiration returns the AmTokenExpiration field if non-nil, zero value otherwise.

### GetAmTokenExpirationOk

`func (o *K8SAuth) GetAmTokenExpirationOk() (*int64, bool)`

GetAmTokenExpirationOk returns a tuple with the AmTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmTokenExpiration

`func (o *K8SAuth) SetAmTokenExpiration(v int64)`

SetAmTokenExpiration sets AmTokenExpiration field to given value.

### HasAmTokenExpiration

`func (o *K8SAuth) HasAmTokenExpiration() bool`

HasAmTokenExpiration returns a boolean if a field has been set.

### GetAuthMethodAccessId

`func (o *K8SAuth) GetAuthMethodAccessId() string`

GetAuthMethodAccessId returns the AuthMethodAccessId field if non-nil, zero value otherwise.

### GetAuthMethodAccessIdOk

`func (o *K8SAuth) GetAuthMethodAccessIdOk() (*string, bool)`

GetAuthMethodAccessIdOk returns a tuple with the AuthMethodAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodAccessId

`func (o *K8SAuth) SetAuthMethodAccessId(v string)`

SetAuthMethodAccessId sets AuthMethodAccessId field to given value.

### HasAuthMethodAccessId

`func (o *K8SAuth) HasAuthMethodAccessId() bool`

HasAuthMethodAccessId returns a boolean if a field has been set.

### GetAuthMethodPrvKeyPem

`func (o *K8SAuth) GetAuthMethodPrvKeyPem() string`

GetAuthMethodPrvKeyPem returns the AuthMethodPrvKeyPem field if non-nil, zero value otherwise.

### GetAuthMethodPrvKeyPemOk

`func (o *K8SAuth) GetAuthMethodPrvKeyPemOk() (*string, bool)`

GetAuthMethodPrvKeyPemOk returns a tuple with the AuthMethodPrvKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodPrvKeyPem

`func (o *K8SAuth) SetAuthMethodPrvKeyPem(v string)`

SetAuthMethodPrvKeyPem sets AuthMethodPrvKeyPem field to given value.

### HasAuthMethodPrvKeyPem

`func (o *K8SAuth) HasAuthMethodPrvKeyPem() bool`

HasAuthMethodPrvKeyPem returns a boolean if a field has been set.

### GetDisableIssValidation

`func (o *K8SAuth) GetDisableIssValidation() bool`

GetDisableIssValidation returns the DisableIssValidation field if non-nil, zero value otherwise.

### GetDisableIssValidationOk

`func (o *K8SAuth) GetDisableIssValidationOk() (*bool, bool)`

GetDisableIssValidationOk returns a tuple with the DisableIssValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIssValidation

`func (o *K8SAuth) SetDisableIssValidation(v bool)`

SetDisableIssValidation sets DisableIssValidation field to given value.

### HasDisableIssValidation

`func (o *K8SAuth) HasDisableIssValidation() bool`

HasDisableIssValidation returns a boolean if a field has been set.

### GetId

`func (o *K8SAuth) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *K8SAuth) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *K8SAuth) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *K8SAuth) HasId() bool`

HasId returns a boolean if a field has been set.

### GetK8sCaCert

`func (o *K8SAuth) GetK8sCaCert() string`

GetK8sCaCert returns the K8sCaCert field if non-nil, zero value otherwise.

### GetK8sCaCertOk

`func (o *K8SAuth) GetK8sCaCertOk() (*string, bool)`

GetK8sCaCertOk returns a tuple with the K8sCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sCaCert

`func (o *K8SAuth) SetK8sCaCert(v string)`

SetK8sCaCert sets K8sCaCert field to given value.

### HasK8sCaCert

`func (o *K8SAuth) HasK8sCaCert() bool`

HasK8sCaCert returns a boolean if a field has been set.

### GetK8sHost

`func (o *K8SAuth) GetK8sHost() string`

GetK8sHost returns the K8sHost field if non-nil, zero value otherwise.

### GetK8sHostOk

`func (o *K8SAuth) GetK8sHostOk() (*string, bool)`

GetK8sHostOk returns a tuple with the K8sHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sHost

`func (o *K8SAuth) SetK8sHost(v string)`

SetK8sHost sets K8sHost field to given value.

### HasK8sHost

`func (o *K8SAuth) HasK8sHost() bool`

HasK8sHost returns a boolean if a field has been set.

### GetK8sIssuer

`func (o *K8SAuth) GetK8sIssuer() string`

GetK8sIssuer returns the K8sIssuer field if non-nil, zero value otherwise.

### GetK8sIssuerOk

`func (o *K8SAuth) GetK8sIssuerOk() (*string, bool)`

GetK8sIssuerOk returns a tuple with the K8sIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sIssuer

`func (o *K8SAuth) SetK8sIssuer(v string)`

SetK8sIssuer sets K8sIssuer field to given value.

### HasK8sIssuer

`func (o *K8SAuth) HasK8sIssuer() bool`

HasK8sIssuer returns a boolean if a field has been set.

### GetK8sPubKeysPem

`func (o *K8SAuth) GetK8sPubKeysPem() []string`

GetK8sPubKeysPem returns the K8sPubKeysPem field if non-nil, zero value otherwise.

### GetK8sPubKeysPemOk

`func (o *K8SAuth) GetK8sPubKeysPemOk() (*[]string, bool)`

GetK8sPubKeysPemOk returns a tuple with the K8sPubKeysPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPubKeysPem

`func (o *K8SAuth) SetK8sPubKeysPem(v []string)`

SetK8sPubKeysPem sets K8sPubKeysPem field to given value.

### HasK8sPubKeysPem

`func (o *K8SAuth) HasK8sPubKeysPem() bool`

HasK8sPubKeysPem returns a boolean if a field has been set.

### GetK8sTokenReviewerJwt

`func (o *K8SAuth) GetK8sTokenReviewerJwt() string`

GetK8sTokenReviewerJwt returns the K8sTokenReviewerJwt field if non-nil, zero value otherwise.

### GetK8sTokenReviewerJwtOk

`func (o *K8SAuth) GetK8sTokenReviewerJwtOk() (*string, bool)`

GetK8sTokenReviewerJwtOk returns a tuple with the K8sTokenReviewerJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sTokenReviewerJwt

`func (o *K8SAuth) SetK8sTokenReviewerJwt(v string)`

SetK8sTokenReviewerJwt sets K8sTokenReviewerJwt field to given value.

### HasK8sTokenReviewerJwt

`func (o *K8SAuth) HasK8sTokenReviewerJwt() bool`

HasK8sTokenReviewerJwt returns a boolean if a field has been set.

### GetName

`func (o *K8SAuth) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *K8SAuth) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *K8SAuth) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *K8SAuth) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtectionKey

`func (o *K8SAuth) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *K8SAuth) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *K8SAuth) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *K8SAuth) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetUseLocalCaJwt

`func (o *K8SAuth) GetUseLocalCaJwt() bool`

GetUseLocalCaJwt returns the UseLocalCaJwt field if non-nil, zero value otherwise.

### GetUseLocalCaJwtOk

`func (o *K8SAuth) GetUseLocalCaJwtOk() (*bool, bool)`

GetUseLocalCaJwtOk returns a tuple with the UseLocalCaJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalCaJwt

`func (o *K8SAuth) SetUseLocalCaJwt(v bool)`

SetUseLocalCaJwt sets UseLocalCaJwt field to given value.

### HasUseLocalCaJwt

`func (o *K8SAuth) HasUseLocalCaJwt() bool`

HasUseLocalCaJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


