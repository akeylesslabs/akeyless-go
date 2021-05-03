# GetPKICertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltNames** | Pointer to **string** | The Subject Alternative Names to be included in the PKI certificate (in a comma-delimited list) | [optional] 
**CertIssuerName** | **string** | The name of the PKI certificate issuer | 
**CommonName** | Pointer to **string** | The common name to be included in the PKI certificate | [optional] 
**KeyDataBase64** | Pointer to **string** | PKI key file contents. If this option is used, the certificate will be printed to stdout | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UriSans** | Pointer to **string** | The URI Subject Alternative Names to be included in the PKI certificate (in a comma-delimited list) | [optional] 

## Methods

### NewGetPKICertificate

`func NewGetPKICertificate(certIssuerName string, ) *GetPKICertificate`

NewGetPKICertificate instantiates a new GetPKICertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPKICertificateWithDefaults

`func NewGetPKICertificateWithDefaults() *GetPKICertificate`

NewGetPKICertificateWithDefaults instantiates a new GetPKICertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAltNames

`func (o *GetPKICertificate) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *GetPKICertificate) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *GetPKICertificate) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.

### HasAltNames

`func (o *GetPKICertificate) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.

### GetCertIssuerName

`func (o *GetPKICertificate) GetCertIssuerName() string`

GetCertIssuerName returns the CertIssuerName field if non-nil, zero value otherwise.

### GetCertIssuerNameOk

`func (o *GetPKICertificate) GetCertIssuerNameOk() (*string, bool)`

GetCertIssuerNameOk returns a tuple with the CertIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerName

`func (o *GetPKICertificate) SetCertIssuerName(v string)`

SetCertIssuerName sets CertIssuerName field to given value.


### GetCommonName

`func (o *GetPKICertificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *GetPKICertificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *GetPKICertificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *GetPKICertificate) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetKeyDataBase64

`func (o *GetPKICertificate) GetKeyDataBase64() string`

GetKeyDataBase64 returns the KeyDataBase64 field if non-nil, zero value otherwise.

### GetKeyDataBase64Ok

`func (o *GetPKICertificate) GetKeyDataBase64Ok() (*string, bool)`

GetKeyDataBase64Ok returns a tuple with the KeyDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDataBase64

`func (o *GetPKICertificate) SetKeyDataBase64(v string)`

SetKeyDataBase64 sets KeyDataBase64 field to given value.

### HasKeyDataBase64

`func (o *GetPKICertificate) HasKeyDataBase64() bool`

HasKeyDataBase64 returns a boolean if a field has been set.

### GetToken

`func (o *GetPKICertificate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetPKICertificate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetPKICertificate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetPKICertificate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GetPKICertificate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GetPKICertificate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GetPKICertificate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GetPKICertificate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUriSans

`func (o *GetPKICertificate) GetUriSans() string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *GetPKICertificate) GetUriSansOk() (*string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *GetPKICertificate) SetUriSans(v string)`

SetUriSans sets UriSans field to given value.

### HasUriSans

`func (o *GetPKICertificate) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


