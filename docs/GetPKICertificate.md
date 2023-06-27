# GetPKICertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltNames** | Pointer to **string** | The Subject Alternative Names to be included in the PKI certificate (in a comma-separated list) (if CSR is supplied this flag is ignored and any DNS.* names are taken from it) | [optional] 
**CertIssuerName** | **string** | The name of the PKI certificate issuer | 
**CommonName** | Pointer to **string** | The common name to be included in the PKI certificate (if CSR is supplied this flag is ignored and the CSR subject CN is taken) | [optional] 
**CsrDataBase64** | Pointer to **string** | Certificate Signing Request contents encoded in base64 to generate the certificate with | [optional] 
**ExtendedKeyUsage** | Pointer to **string** | A comma-separated list of extended key usage requests which will be used for certificate issuance. Supported values: &#39;clientauth&#39;, &#39;serverauth&#39;. | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyDataBase64** | Pointer to **string** | PKI key file contents. If this option is used, the certificate will be printed to stdout | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | Pointer to **int64** | Updated certificate lifetime in seconds (must be less than the Certificate Issuer default TTL) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UriSans** | Pointer to **string** | The URI Subject Alternative Names to be included in the PKI certificate (in a comma-separated list) (if CSR is supplied this flag is ignored and any URI.* names are taken from it) | [optional] 

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

### GetCsrDataBase64

`func (o *GetPKICertificate) GetCsrDataBase64() string`

GetCsrDataBase64 returns the CsrDataBase64 field if non-nil, zero value otherwise.

### GetCsrDataBase64Ok

`func (o *GetPKICertificate) GetCsrDataBase64Ok() (*string, bool)`

GetCsrDataBase64Ok returns a tuple with the CsrDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataBase64

`func (o *GetPKICertificate) SetCsrDataBase64(v string)`

SetCsrDataBase64 sets CsrDataBase64 field to given value.

### HasCsrDataBase64

`func (o *GetPKICertificate) HasCsrDataBase64() bool`

HasCsrDataBase64 returns a boolean if a field has been set.

### GetExtendedKeyUsage

`func (o *GetPKICertificate) GetExtendedKeyUsage() string`

GetExtendedKeyUsage returns the ExtendedKeyUsage field if non-nil, zero value otherwise.

### GetExtendedKeyUsageOk

`func (o *GetPKICertificate) GetExtendedKeyUsageOk() (*string, bool)`

GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsage

`func (o *GetPKICertificate) SetExtendedKeyUsage(v string)`

SetExtendedKeyUsage sets ExtendedKeyUsage field to given value.

### HasExtendedKeyUsage

`func (o *GetPKICertificate) HasExtendedKeyUsage() bool`

HasExtendedKeyUsage returns a boolean if a field has been set.

### GetJson

`func (o *GetPKICertificate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GetPKICertificate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GetPKICertificate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GetPKICertificate) HasJson() bool`

HasJson returns a boolean if a field has been set.

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

### GetTtl

`func (o *GetPKICertificate) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetPKICertificate) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetPKICertificate) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetPKICertificate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

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


