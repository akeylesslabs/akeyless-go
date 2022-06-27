# GetSSHCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIssuerName** | **string** | The name of the SSH certificate issuer | 
**CertUsername** | **string** | The username to sign in the SSH certificate | 
**LegacySigningAlgName** | Pointer to **bool** | Set this option to output legacy (&#39;ssh-rsa-cert-v01@openssh.com&#39;) signing algorithm name in the certificate. | [optional] 
**PublicKeyData** | Pointer to **string** | SSH public key file contents. If this option is used, the certificate will be printed to stdout | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | Pointer to **int64** | Updated certificate lifetime in seconds (must be less than the Certificate Issuer default TTL) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGetSSHCertificate

`func NewGetSSHCertificate(certIssuerName string, certUsername string, ) *GetSSHCertificate`

NewGetSSHCertificate instantiates a new GetSSHCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSSHCertificateWithDefaults

`func NewGetSSHCertificateWithDefaults() *GetSSHCertificate`

NewGetSSHCertificateWithDefaults instantiates a new GetSSHCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIssuerName

`func (o *GetSSHCertificate) GetCertIssuerName() string`

GetCertIssuerName returns the CertIssuerName field if non-nil, zero value otherwise.

### GetCertIssuerNameOk

`func (o *GetSSHCertificate) GetCertIssuerNameOk() (*string, bool)`

GetCertIssuerNameOk returns a tuple with the CertIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerName

`func (o *GetSSHCertificate) SetCertIssuerName(v string)`

SetCertIssuerName sets CertIssuerName field to given value.


### GetCertUsername

`func (o *GetSSHCertificate) GetCertUsername() string`

GetCertUsername returns the CertUsername field if non-nil, zero value otherwise.

### GetCertUsernameOk

`func (o *GetSSHCertificate) GetCertUsernameOk() (*string, bool)`

GetCertUsernameOk returns a tuple with the CertUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertUsername

`func (o *GetSSHCertificate) SetCertUsername(v string)`

SetCertUsername sets CertUsername field to given value.


### GetLegacySigningAlgName

`func (o *GetSSHCertificate) GetLegacySigningAlgName() bool`

GetLegacySigningAlgName returns the LegacySigningAlgName field if non-nil, zero value otherwise.

### GetLegacySigningAlgNameOk

`func (o *GetSSHCertificate) GetLegacySigningAlgNameOk() (*bool, bool)`

GetLegacySigningAlgNameOk returns a tuple with the LegacySigningAlgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacySigningAlgName

`func (o *GetSSHCertificate) SetLegacySigningAlgName(v bool)`

SetLegacySigningAlgName sets LegacySigningAlgName field to given value.

### HasLegacySigningAlgName

`func (o *GetSSHCertificate) HasLegacySigningAlgName() bool`

HasLegacySigningAlgName returns a boolean if a field has been set.

### GetPublicKeyData

`func (o *GetSSHCertificate) GetPublicKeyData() string`

GetPublicKeyData returns the PublicKeyData field if non-nil, zero value otherwise.

### GetPublicKeyDataOk

`func (o *GetSSHCertificate) GetPublicKeyDataOk() (*string, bool)`

GetPublicKeyDataOk returns a tuple with the PublicKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyData

`func (o *GetSSHCertificate) SetPublicKeyData(v string)`

SetPublicKeyData sets PublicKeyData field to given value.

### HasPublicKeyData

`func (o *GetSSHCertificate) HasPublicKeyData() bool`

HasPublicKeyData returns a boolean if a field has been set.

### GetToken

`func (o *GetSSHCertificate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetSSHCertificate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetSSHCertificate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetSSHCertificate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *GetSSHCertificate) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetSSHCertificate) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetSSHCertificate) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetSSHCertificate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUidToken

`func (o *GetSSHCertificate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GetSSHCertificate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GetSSHCertificate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GetSSHCertificate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


