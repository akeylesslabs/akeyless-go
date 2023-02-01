# CertificateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtKeyUsage** | Pointer to **[]int64** |  | [optional] 
**KeyUsage** | Pointer to **int64** | KeyUsage represents the set of actions that are valid for a given key. It&#39;s a bitmap of the KeyUsage* constants. | [optional] 
**DnsNames** | Pointer to **[]string** |  | [optional] 
**EmailAddresses** | Pointer to **[]string** |  | [optional] 
**Extensions** | Pointer to [**[]Extension**](Extension.md) |  | [optional] 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**IsCa** | Pointer to **bool** |  | [optional] 
**Issuer** | Pointer to [**Name**](Name.md) |  | [optional] 
**IssuingCertificateUrl** | Pointer to **[]string** |  | [optional] 
**KeySize** | Pointer to **int64** |  | [optional] 
**NotAfter** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**NotBefore** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**OcspServer** | Pointer to **[]string** |  | [optional] 
**PublicKeyAlgorithmName** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Sha1Fingerprint** | Pointer to **string** |  | [optional] 
**Sha256Fingerprint** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**SignatureAlgorithmName** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to [**Name**](Name.md) |  | [optional] 
**SubjectPublicKey** | Pointer to **string** |  | [optional] 
**Uris** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewCertificateInfo

`func NewCertificateInfo() *CertificateInfo`

NewCertificateInfo instantiates a new CertificateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateInfoWithDefaults

`func NewCertificateInfoWithDefaults() *CertificateInfo`

NewCertificateInfoWithDefaults instantiates a new CertificateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtKeyUsage

`func (o *CertificateInfo) GetExtKeyUsage() []int64`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *CertificateInfo) GetExtKeyUsageOk() (*[]int64, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsage

`func (o *CertificateInfo) SetExtKeyUsage(v []int64)`

SetExtKeyUsage sets ExtKeyUsage field to given value.

### HasExtKeyUsage

`func (o *CertificateInfo) HasExtKeyUsage() bool`

HasExtKeyUsage returns a boolean if a field has been set.

### GetKeyUsage

`func (o *CertificateInfo) GetKeyUsage() int64`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *CertificateInfo) GetKeyUsageOk() (*int64, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *CertificateInfo) SetKeyUsage(v int64)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *CertificateInfo) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetDnsNames

`func (o *CertificateInfo) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *CertificateInfo) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *CertificateInfo) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *CertificateInfo) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *CertificateInfo) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *CertificateInfo) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *CertificateInfo) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *CertificateInfo) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetExtensions

`func (o *CertificateInfo) GetExtensions() []Extension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CertificateInfo) GetExtensionsOk() (*[]Extension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CertificateInfo) SetExtensions(v []Extension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CertificateInfo) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetIpAddresses

`func (o *CertificateInfo) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CertificateInfo) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CertificateInfo) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *CertificateInfo) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetIsCa

`func (o *CertificateInfo) GetIsCa() bool`

GetIsCa returns the IsCa field if non-nil, zero value otherwise.

### GetIsCaOk

`func (o *CertificateInfo) GetIsCaOk() (*bool, bool)`

GetIsCaOk returns a tuple with the IsCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCa

`func (o *CertificateInfo) SetIsCa(v bool)`

SetIsCa sets IsCa field to given value.

### HasIsCa

`func (o *CertificateInfo) HasIsCa() bool`

HasIsCa returns a boolean if a field has been set.

### GetIssuer

`func (o *CertificateInfo) GetIssuer() Name`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateInfo) GetIssuerOk() (*Name, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateInfo) SetIssuer(v Name)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertificateInfo) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetIssuingCertificateUrl

`func (o *CertificateInfo) GetIssuingCertificateUrl() []string`

GetIssuingCertificateUrl returns the IssuingCertificateUrl field if non-nil, zero value otherwise.

### GetIssuingCertificateUrlOk

`func (o *CertificateInfo) GetIssuingCertificateUrlOk() (*[]string, bool)`

GetIssuingCertificateUrlOk returns a tuple with the IssuingCertificateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificateUrl

`func (o *CertificateInfo) SetIssuingCertificateUrl(v []string)`

SetIssuingCertificateUrl sets IssuingCertificateUrl field to given value.

### HasIssuingCertificateUrl

`func (o *CertificateInfo) HasIssuingCertificateUrl() bool`

HasIssuingCertificateUrl returns a boolean if a field has been set.

### GetKeySize

`func (o *CertificateInfo) GetKeySize() int64`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *CertificateInfo) GetKeySizeOk() (*int64, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *CertificateInfo) SetKeySize(v int64)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *CertificateInfo) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetNotAfter

`func (o *CertificateInfo) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateInfo) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateInfo) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateInfo) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificateInfo) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateInfo) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateInfo) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateInfo) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetOcspServer

`func (o *CertificateInfo) GetOcspServer() []string`

GetOcspServer returns the OcspServer field if non-nil, zero value otherwise.

### GetOcspServerOk

`func (o *CertificateInfo) GetOcspServerOk() (*[]string, bool)`

GetOcspServerOk returns a tuple with the OcspServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServer

`func (o *CertificateInfo) SetOcspServer(v []string)`

SetOcspServer sets OcspServer field to given value.

### HasOcspServer

`func (o *CertificateInfo) HasOcspServer() bool`

HasOcspServer returns a boolean if a field has been set.

### GetPublicKeyAlgorithmName

`func (o *CertificateInfo) GetPublicKeyAlgorithmName() string`

GetPublicKeyAlgorithmName returns the PublicKeyAlgorithmName field if non-nil, zero value otherwise.

### GetPublicKeyAlgorithmNameOk

`func (o *CertificateInfo) GetPublicKeyAlgorithmNameOk() (*string, bool)`

GetPublicKeyAlgorithmNameOk returns a tuple with the PublicKeyAlgorithmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyAlgorithmName

`func (o *CertificateInfo) SetPublicKeyAlgorithmName(v string)`

SetPublicKeyAlgorithmName sets PublicKeyAlgorithmName field to given value.

### HasPublicKeyAlgorithmName

`func (o *CertificateInfo) HasPublicKeyAlgorithmName() bool`

HasPublicKeyAlgorithmName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CertificateInfo) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateInfo) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateInfo) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificateInfo) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSha1Fingerprint

`func (o *CertificateInfo) GetSha1Fingerprint() string`

GetSha1Fingerprint returns the Sha1Fingerprint field if non-nil, zero value otherwise.

### GetSha1FingerprintOk

`func (o *CertificateInfo) GetSha1FingerprintOk() (*string, bool)`

GetSha1FingerprintOk returns a tuple with the Sha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Fingerprint

`func (o *CertificateInfo) SetSha1Fingerprint(v string)`

SetSha1Fingerprint sets Sha1Fingerprint field to given value.

### HasSha1Fingerprint

`func (o *CertificateInfo) HasSha1Fingerprint() bool`

HasSha1Fingerprint returns a boolean if a field has been set.

### GetSha256Fingerprint

`func (o *CertificateInfo) GetSha256Fingerprint() string`

GetSha256Fingerprint returns the Sha256Fingerprint field if non-nil, zero value otherwise.

### GetSha256FingerprintOk

`func (o *CertificateInfo) GetSha256FingerprintOk() (*string, bool)`

GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Fingerprint

`func (o *CertificateInfo) SetSha256Fingerprint(v string)`

SetSha256Fingerprint sets Sha256Fingerprint field to given value.

### HasSha256Fingerprint

`func (o *CertificateInfo) HasSha256Fingerprint() bool`

HasSha256Fingerprint returns a boolean if a field has been set.

### GetSignature

`func (o *CertificateInfo) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CertificateInfo) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CertificateInfo) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CertificateInfo) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignatureAlgorithmName

`func (o *CertificateInfo) GetSignatureAlgorithmName() string`

GetSignatureAlgorithmName returns the SignatureAlgorithmName field if non-nil, zero value otherwise.

### GetSignatureAlgorithmNameOk

`func (o *CertificateInfo) GetSignatureAlgorithmNameOk() (*string, bool)`

GetSignatureAlgorithmNameOk returns a tuple with the SignatureAlgorithmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithmName

`func (o *CertificateInfo) SetSignatureAlgorithmName(v string)`

SetSignatureAlgorithmName sets SignatureAlgorithmName field to given value.

### HasSignatureAlgorithmName

`func (o *CertificateInfo) HasSignatureAlgorithmName() bool`

HasSignatureAlgorithmName returns a boolean if a field has been set.

### GetSubject

`func (o *CertificateInfo) GetSubject() Name`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateInfo) GetSubjectOk() (*Name, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateInfo) SetSubject(v Name)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertificateInfo) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectPublicKey

`func (o *CertificateInfo) GetSubjectPublicKey() string`

GetSubjectPublicKey returns the SubjectPublicKey field if non-nil, zero value otherwise.

### GetSubjectPublicKeyOk

`func (o *CertificateInfo) GetSubjectPublicKeyOk() (*string, bool)`

GetSubjectPublicKeyOk returns a tuple with the SubjectPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPublicKey

`func (o *CertificateInfo) SetSubjectPublicKey(v string)`

SetSubjectPublicKey sets SubjectPublicKey field to given value.

### HasSubjectPublicKey

`func (o *CertificateInfo) HasSubjectPublicKey() bool`

HasSubjectPublicKey returns a boolean if a field has been set.

### GetUris

`func (o *CertificateInfo) GetUris() []string`

GetUris returns the Uris field if non-nil, zero value otherwise.

### GetUrisOk

`func (o *CertificateInfo) GetUrisOk() (*[]string, bool)`

GetUrisOk returns a tuple with the Uris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUris

`func (o *CertificateInfo) SetUris(v []string)`

SetUris sets Uris field to given value.

### HasUris

`func (o *CertificateInfo) HasUris() bool`

HasUris returns a boolean if a field has been set.

### GetVersion

`func (o *CertificateInfo) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CertificateInfo) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CertificateInfo) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CertificateInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


