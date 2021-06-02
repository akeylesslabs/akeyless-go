# GetSSHCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIssuerName** | **string** | The name of the SSH certificate issuer | 
**CertUsername** | **string** | The username to sign in the SSH certificate | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**PublicKeyData** | Pointer to **string** | SSH public key file contents. If this option is used, the certificate will be printed to stdout | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

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


### GetPassword

`func (o *GetSSHCertificate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetSSHCertificate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetSSHCertificate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetSSHCertificate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

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

### GetUsername

`func (o *GetSSHCertificate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetSSHCertificate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetSSHCertificate) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetSSHCertificate) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


