# CertificateTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CsrCnfBase64** | Pointer to **string** |  | [optional] 
**DigestAlgo** | Pointer to **string** |  | [optional] 
**Locality** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Province** | Pointer to **string** |  | [optional] 
**SelfSignedEnabled** | Pointer to **bool** |  | [optional] 
**Ttl** | Pointer to **int64** |  | [optional] 

## Methods

### NewCertificateTemplateInfo

`func NewCertificateTemplateInfo() *CertificateTemplateInfo`

NewCertificateTemplateInfo instantiates a new CertificateTemplateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateTemplateInfoWithDefaults

`func NewCertificateTemplateInfoWithDefaults() *CertificateTemplateInfo`

NewCertificateTemplateInfoWithDefaults instantiates a new CertificateTemplateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *CertificateTemplateInfo) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateTemplateInfo) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateTemplateInfo) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CertificateTemplateInfo) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountry

`func (o *CertificateTemplateInfo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CertificateTemplateInfo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CertificateTemplateInfo) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CertificateTemplateInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCsrCnfBase64

`func (o *CertificateTemplateInfo) GetCsrCnfBase64() string`

GetCsrCnfBase64 returns the CsrCnfBase64 field if non-nil, zero value otherwise.

### GetCsrCnfBase64Ok

`func (o *CertificateTemplateInfo) GetCsrCnfBase64Ok() (*string, bool)`

GetCsrCnfBase64Ok returns a tuple with the CsrCnfBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrCnfBase64

`func (o *CertificateTemplateInfo) SetCsrCnfBase64(v string)`

SetCsrCnfBase64 sets CsrCnfBase64 field to given value.

### HasCsrCnfBase64

`func (o *CertificateTemplateInfo) HasCsrCnfBase64() bool`

HasCsrCnfBase64 returns a boolean if a field has been set.

### GetDigestAlgo

`func (o *CertificateTemplateInfo) GetDigestAlgo() string`

GetDigestAlgo returns the DigestAlgo field if non-nil, zero value otherwise.

### GetDigestAlgoOk

`func (o *CertificateTemplateInfo) GetDigestAlgoOk() (*string, bool)`

GetDigestAlgoOk returns a tuple with the DigestAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgo

`func (o *CertificateTemplateInfo) SetDigestAlgo(v string)`

SetDigestAlgo sets DigestAlgo field to given value.

### HasDigestAlgo

`func (o *CertificateTemplateInfo) HasDigestAlgo() bool`

HasDigestAlgo returns a boolean if a field has been set.

### GetLocality

`func (o *CertificateTemplateInfo) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *CertificateTemplateInfo) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *CertificateTemplateInfo) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *CertificateTemplateInfo) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetOrganization

`func (o *CertificateTemplateInfo) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificateTemplateInfo) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificateTemplateInfo) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CertificateTemplateInfo) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProvince

`func (o *CertificateTemplateInfo) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *CertificateTemplateInfo) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *CertificateTemplateInfo) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *CertificateTemplateInfo) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetSelfSignedEnabled

`func (o *CertificateTemplateInfo) GetSelfSignedEnabled() bool`

GetSelfSignedEnabled returns the SelfSignedEnabled field if non-nil, zero value otherwise.

### GetSelfSignedEnabledOk

`func (o *CertificateTemplateInfo) GetSelfSignedEnabledOk() (*bool, bool)`

GetSelfSignedEnabledOk returns a tuple with the SelfSignedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSignedEnabled

`func (o *CertificateTemplateInfo) SetSelfSignedEnabled(v bool)`

SetSelfSignedEnabled sets SelfSignedEnabled field to given value.

### HasSelfSignedEnabled

`func (o *CertificateTemplateInfo) HasSelfSignedEnabled() bool`

HasSelfSignedEnabled returns a boolean if a field has been set.

### GetTtl

`func (o *CertificateTemplateInfo) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CertificateTemplateInfo) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CertificateTemplateInfo) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CertificateTemplateInfo) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


