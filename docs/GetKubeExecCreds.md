# GetKubeExecCreds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltNames** | Pointer to **string** | The Subject Alternative Names to be included in the PKI certificate (in a comma-delimited list) | [optional] 
**CertIssuerName** | **string** | The name of the PKI certificate issuer | 
**CommonName** | Pointer to **string** | The common name to be included in the PKI certificate | [optional] 
**KeyFilePath** | **string** | The client public or private key file path (in case of a private key, it will be use to extract the public key) | 
**Outfile** | Pointer to **string** | Output file path with the certificate. If not provided, the file with the certificate will be created in the same location of the provided public key with the -cert extension | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UriSans** | Pointer to **string** | The URI Subject Alternative Names to be included in the PKI certificate (in a comma-delimited list) | [optional] 

## Methods

### NewGetKubeExecCreds

`func NewGetKubeExecCreds(certIssuerName string, keyFilePath string, ) *GetKubeExecCreds`

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

### GetKeyFilePath

`func (o *GetKubeExecCreds) GetKeyFilePath() string`

GetKeyFilePath returns the KeyFilePath field if non-nil, zero value otherwise.

### GetKeyFilePathOk

`func (o *GetKubeExecCreds) GetKeyFilePathOk() (*string, bool)`

GetKeyFilePathOk returns a tuple with the KeyFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFilePath

`func (o *GetKubeExecCreds) SetKeyFilePath(v string)`

SetKeyFilePath sets KeyFilePath field to given value.


### GetOutfile

`func (o *GetKubeExecCreds) GetOutfile() string`

GetOutfile returns the Outfile field if non-nil, zero value otherwise.

### GetOutfileOk

`func (o *GetKubeExecCreds) GetOutfileOk() (*string, bool)`

GetOutfileOk returns a tuple with the Outfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutfile

`func (o *GetKubeExecCreds) SetOutfile(v string)`

SetOutfile sets Outfile field to given value.

### HasOutfile

`func (o *GetKubeExecCreds) HasOutfile() bool`

HasOutfile returns a boolean if a field has been set.

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


