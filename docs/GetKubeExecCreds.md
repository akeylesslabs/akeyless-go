# GetKubeExecCreds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltNames** | Pointer to **string** | The Subject Alternative Names to be included in the PKI certificate (in a comma-separated list) (if CSR is supplied this flag is ignored and any DNS.* names are taken from it) | [optional] 
**ApiVersion** | Pointer to **string** | Client authentication API version | [optional] [default to "v1"]
**CertIssuerName** | **string** | The name of the PKI certificate issuer | 
**CommonName** | Pointer to **string** | The common name to be included in the PKI certificate (if CSR is supplied this flag is ignored and the CSR subject CN is taken) | [optional] 
**CsrDataBase64** | Pointer to **string** | Certificate Signing Request contents encoded in base64 to generate the certificate with | [optional] 
**ExtendedKeyUsage** | Pointer to **string** | A comma-separated list of extended key usage requests which will be used for certificate issuance. Supported values: &#39;clientauth&#39;, &#39;serverauth&#39;. | [optional] 
**ExtraExtensions** | Pointer to **string** | A json string that defines the requested extra extensions for the certificate | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyDataBase64** | Pointer to **string** | PKI key file contents. If this option is used, the certificate will be printed to stdout | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | Pointer to **string** | Updated certificate lifetime in seconds (must be less than the Certificate Issuer default TTL) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UriSans** | Pointer to **string** | The URI Subject Alternative Names to be included in the PKI certificate (in a comma-separated list) (if CSR is supplied this flag is ignored and any URI.* names are taken from it) | [optional] 

## Methods

### NewGetKubeExecCreds

`func NewGetKubeExecCreds(certIssuerName string, ) *GetKubeExecCreds`

NewGetKubeExecCreds instantiates a new GetKubeExecCreds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubeExecCredsWithDefaults

`func NewGetKubeExecCredsWithDefaults() *GetKubeExecCreds`

NewGetKubeExecCredsWithDefaults instantiates a new GetKubeExecCreds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAltNames

`func (o *GetKubeExecCreds) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *GetKubeExecCreds) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *GetKubeExecCreds) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.

### HasAltNames

`func (o *GetKubeExecCreds) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.

### GetApiVersion

`func (o *GetKubeExecCreds) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetKubeExecCreds) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetKubeExecCreds) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *GetKubeExecCreds) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCertIssuerName

`func (o *GetKubeExecCreds) GetCertIssuerName() string`

GetCertIssuerName returns the CertIssuerName field if non-nil, zero value otherwise.

### GetCertIssuerNameOk

`func (o *GetKubeExecCreds) GetCertIssuerNameOk() (*string, bool)`

GetCertIssuerNameOk returns a tuple with the CertIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerName

`func (o *GetKubeExecCreds) SetCertIssuerName(v string)`

SetCertIssuerName sets CertIssuerName field to given value.


### GetCommonName

`func (o *GetKubeExecCreds) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *GetKubeExecCreds) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *GetKubeExecCreds) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *GetKubeExecCreds) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCsrDataBase64

`func (o *GetKubeExecCreds) GetCsrDataBase64() string`

GetCsrDataBase64 returns the CsrDataBase64 field if non-nil, zero value otherwise.

### GetCsrDataBase64Ok

`func (o *GetKubeExecCreds) GetCsrDataBase64Ok() (*string, bool)`

GetCsrDataBase64Ok returns a tuple with the CsrDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrDataBase64

`func (o *GetKubeExecCreds) SetCsrDataBase64(v string)`

SetCsrDataBase64 sets CsrDataBase64 field to given value.

### HasCsrDataBase64

`func (o *GetKubeExecCreds) HasCsrDataBase64() bool`

HasCsrDataBase64 returns a boolean if a field has been set.

### GetExtendedKeyUsage

`func (o *GetKubeExecCreds) GetExtendedKeyUsage() string`

GetExtendedKeyUsage returns the ExtendedKeyUsage field if non-nil, zero value otherwise.

### GetExtendedKeyUsageOk

`func (o *GetKubeExecCreds) GetExtendedKeyUsageOk() (*string, bool)`

GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsage

`func (o *GetKubeExecCreds) SetExtendedKeyUsage(v string)`

SetExtendedKeyUsage sets ExtendedKeyUsage field to given value.

### HasExtendedKeyUsage

`func (o *GetKubeExecCreds) HasExtendedKeyUsage() bool`

HasExtendedKeyUsage returns a boolean if a field has been set.

### GetExtraExtensions

`func (o *GetKubeExecCreds) GetExtraExtensions() string`

GetExtraExtensions returns the ExtraExtensions field if non-nil, zero value otherwise.

### GetExtraExtensionsOk

`func (o *GetKubeExecCreds) GetExtraExtensionsOk() (*string, bool)`

GetExtraExtensionsOk returns a tuple with the ExtraExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraExtensions

`func (o *GetKubeExecCreds) SetExtraExtensions(v string)`

SetExtraExtensions sets ExtraExtensions field to given value.

### HasExtraExtensions

`func (o *GetKubeExecCreds) HasExtraExtensions() bool`

HasExtraExtensions returns a boolean if a field has been set.

### GetJson

`func (o *GetKubeExecCreds) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GetKubeExecCreds) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GetKubeExecCreds) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GetKubeExecCreds) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyDataBase64

`func (o *GetKubeExecCreds) GetKeyDataBase64() string`

GetKeyDataBase64 returns the KeyDataBase64 field if non-nil, zero value otherwise.

### GetKeyDataBase64Ok

`func (o *GetKubeExecCreds) GetKeyDataBase64Ok() (*string, bool)`

GetKeyDataBase64Ok returns a tuple with the KeyDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDataBase64

`func (o *GetKubeExecCreds) SetKeyDataBase64(v string)`

SetKeyDataBase64 sets KeyDataBase64 field to given value.

### HasKeyDataBase64

`func (o *GetKubeExecCreds) HasKeyDataBase64() bool`

HasKeyDataBase64 returns a boolean if a field has been set.

### GetToken

`func (o *GetKubeExecCreds) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetKubeExecCreds) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetKubeExecCreds) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetKubeExecCreds) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *GetKubeExecCreds) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetKubeExecCreds) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetKubeExecCreds) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetKubeExecCreds) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUidToken

`func (o *GetKubeExecCreds) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GetKubeExecCreds) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GetKubeExecCreds) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GetKubeExecCreds) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUriSans

`func (o *GetKubeExecCreds) GetUriSans() string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *GetKubeExecCreds) GetUriSansOk() (*string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *GetKubeExecCreds) SetUriSans(v string)`

SetUriSans sets UriSans field to given value.

### HasUriSans

`func (o *GetKubeExecCreds) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


