# PKICertificateIssueDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAnyName** | Pointer to **bool** |  | [optional] 
**AllowCopyExtFromCsr** | Pointer to **bool** |  | [optional] 
**AllowSubdomains** | Pointer to **bool** |  | [optional] 
**AllowedDomainsList** | Pointer to **[]string** |  | [optional] 
**AllowedExtraExtensions** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**AllowedUriSans** | Pointer to **[]string** |  | [optional] 
**BasicConstraintsValidForNonCa** | Pointer to **bool** |  | [optional] 
**CertificateAuthorityMode** | Pointer to **string** |  | [optional] 
**ClientFlag** | Pointer to **bool** |  | [optional] 
**CodeSigningFlag** | Pointer to **bool** |  | [optional] 
**Country** | Pointer to **[]string** |  | [optional] 
**CreatePrivateCrl** | Pointer to **bool** |  | [optional] 
**CreatePublicCrl** | Pointer to **bool** |  | [optional] 
**DestinationPath** | Pointer to **string** | DestinationPath is the destination to save generated certificates | [optional] 
**EnforceHostnames** | Pointer to **bool** |  | [optional] 
**ExpirationEvents** | Pointer to [**[]CertificateExpirationEvent**](CertificateExpirationEvent.md) | ExpirationNotification holds a list of expiration notices that should be sent in case a certificate is about to expire, this value is being propagated to the Certificate resources that are created | [optional] 
**GwClusterId** | Pointer to **int64** |  | [optional] 
**GwClusterUrl** | Pointer to **string** | GWClusterURL is required when CAMode is \&quot;public\&quot; and it defines the cluster URL the PKI should be issued from. The GW cluster must have permissions to read associated target&#39;s details | [optional] 
**IsCa** | Pointer to **bool** |  | [optional] 
**KeyBits** | Pointer to **int64** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**KeyUsageList** | Pointer to **[]string** |  | [optional] 
**Locality** | Pointer to **[]string** |  | [optional] 
**NotBeforeDuration** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 
**OrganizationList** | Pointer to **[]string** |  | [optional] 
**OrganizationUnitList** | Pointer to **[]string** |  | [optional] 
**PostalCode** | Pointer to **[]string** |  | [optional] 
**ProtectGeneratedCertificates** | Pointer to **bool** | ProtectGeneratedCertificates dictates whether the created certificates should be protected from deletion | [optional] 
**Province** | Pointer to **[]string** |  | [optional] 
**RequireCn** | Pointer to **bool** |  | [optional] 
**ServerFlag** | Pointer to **bool** |  | [optional] 
**StreetAddress** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPKICertificateIssueDetails

`func NewPKICertificateIssueDetails() *PKICertificateIssueDetails`

NewPKICertificateIssueDetails instantiates a new PKICertificateIssueDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKICertificateIssueDetailsWithDefaults

`func NewPKICertificateIssueDetailsWithDefaults() *PKICertificateIssueDetails`

NewPKICertificateIssueDetailsWithDefaults instantiates a new PKICertificateIssueDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAnyName

`func (o *PKICertificateIssueDetails) GetAllowAnyName() bool`

GetAllowAnyName returns the AllowAnyName field if non-nil, zero value otherwise.

### GetAllowAnyNameOk

`func (o *PKICertificateIssueDetails) GetAllowAnyNameOk() (*bool, bool)`

GetAllowAnyNameOk returns a tuple with the AllowAnyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnyName

`func (o *PKICertificateIssueDetails) SetAllowAnyName(v bool)`

SetAllowAnyName sets AllowAnyName field to given value.

### HasAllowAnyName

`func (o *PKICertificateIssueDetails) HasAllowAnyName() bool`

HasAllowAnyName returns a boolean if a field has been set.

### GetAllowCopyExtFromCsr

`func (o *PKICertificateIssueDetails) GetAllowCopyExtFromCsr() bool`

GetAllowCopyExtFromCsr returns the AllowCopyExtFromCsr field if non-nil, zero value otherwise.

### GetAllowCopyExtFromCsrOk

`func (o *PKICertificateIssueDetails) GetAllowCopyExtFromCsrOk() (*bool, bool)`

GetAllowCopyExtFromCsrOk returns a tuple with the AllowCopyExtFromCsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCopyExtFromCsr

`func (o *PKICertificateIssueDetails) SetAllowCopyExtFromCsr(v bool)`

SetAllowCopyExtFromCsr sets AllowCopyExtFromCsr field to given value.

### HasAllowCopyExtFromCsr

`func (o *PKICertificateIssueDetails) HasAllowCopyExtFromCsr() bool`

HasAllowCopyExtFromCsr returns a boolean if a field has been set.

### GetAllowSubdomains

`func (o *PKICertificateIssueDetails) GetAllowSubdomains() bool`

GetAllowSubdomains returns the AllowSubdomains field if non-nil, zero value otherwise.

### GetAllowSubdomainsOk

`func (o *PKICertificateIssueDetails) GetAllowSubdomainsOk() (*bool, bool)`

GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubdomains

`func (o *PKICertificateIssueDetails) SetAllowSubdomains(v bool)`

SetAllowSubdomains sets AllowSubdomains field to given value.

### HasAllowSubdomains

`func (o *PKICertificateIssueDetails) HasAllowSubdomains() bool`

HasAllowSubdomains returns a boolean if a field has been set.

### GetAllowedDomainsList

`func (o *PKICertificateIssueDetails) GetAllowedDomainsList() []string`

GetAllowedDomainsList returns the AllowedDomainsList field if non-nil, zero value otherwise.

### GetAllowedDomainsListOk

`func (o *PKICertificateIssueDetails) GetAllowedDomainsListOk() (*[]string, bool)`

GetAllowedDomainsListOk returns a tuple with the AllowedDomainsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomainsList

`func (o *PKICertificateIssueDetails) SetAllowedDomainsList(v []string)`

SetAllowedDomainsList sets AllowedDomainsList field to given value.

### HasAllowedDomainsList

`func (o *PKICertificateIssueDetails) HasAllowedDomainsList() bool`

HasAllowedDomainsList returns a boolean if a field has been set.

### GetAllowedExtraExtensions

`func (o *PKICertificateIssueDetails) GetAllowedExtraExtensions() map[string][]string`

GetAllowedExtraExtensions returns the AllowedExtraExtensions field if non-nil, zero value otherwise.

### GetAllowedExtraExtensionsOk

`func (o *PKICertificateIssueDetails) GetAllowedExtraExtensionsOk() (*map[string][]string, bool)`

GetAllowedExtraExtensionsOk returns a tuple with the AllowedExtraExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExtraExtensions

`func (o *PKICertificateIssueDetails) SetAllowedExtraExtensions(v map[string][]string)`

SetAllowedExtraExtensions sets AllowedExtraExtensions field to given value.

### HasAllowedExtraExtensions

`func (o *PKICertificateIssueDetails) HasAllowedExtraExtensions() bool`

HasAllowedExtraExtensions returns a boolean if a field has been set.

### GetAllowedUriSans

`func (o *PKICertificateIssueDetails) GetAllowedUriSans() []string`

GetAllowedUriSans returns the AllowedUriSans field if non-nil, zero value otherwise.

### GetAllowedUriSansOk

`func (o *PKICertificateIssueDetails) GetAllowedUriSansOk() (*[]string, bool)`

GetAllowedUriSansOk returns a tuple with the AllowedUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUriSans

`func (o *PKICertificateIssueDetails) SetAllowedUriSans(v []string)`

SetAllowedUriSans sets AllowedUriSans field to given value.

### HasAllowedUriSans

`func (o *PKICertificateIssueDetails) HasAllowedUriSans() bool`

HasAllowedUriSans returns a boolean if a field has been set.

### GetBasicConstraintsValidForNonCa

`func (o *PKICertificateIssueDetails) GetBasicConstraintsValidForNonCa() bool`

GetBasicConstraintsValidForNonCa returns the BasicConstraintsValidForNonCa field if non-nil, zero value otherwise.

### GetBasicConstraintsValidForNonCaOk

`func (o *PKICertificateIssueDetails) GetBasicConstraintsValidForNonCaOk() (*bool, bool)`

GetBasicConstraintsValidForNonCaOk returns a tuple with the BasicConstraintsValidForNonCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraintsValidForNonCa

`func (o *PKICertificateIssueDetails) SetBasicConstraintsValidForNonCa(v bool)`

SetBasicConstraintsValidForNonCa sets BasicConstraintsValidForNonCa field to given value.

### HasBasicConstraintsValidForNonCa

`func (o *PKICertificateIssueDetails) HasBasicConstraintsValidForNonCa() bool`

HasBasicConstraintsValidForNonCa returns a boolean if a field has been set.

### GetCertificateAuthorityMode

`func (o *PKICertificateIssueDetails) GetCertificateAuthorityMode() string`

GetCertificateAuthorityMode returns the CertificateAuthorityMode field if non-nil, zero value otherwise.

### GetCertificateAuthorityModeOk

`func (o *PKICertificateIssueDetails) GetCertificateAuthorityModeOk() (*string, bool)`

GetCertificateAuthorityModeOk returns a tuple with the CertificateAuthorityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityMode

`func (o *PKICertificateIssueDetails) SetCertificateAuthorityMode(v string)`

SetCertificateAuthorityMode sets CertificateAuthorityMode field to given value.

### HasCertificateAuthorityMode

`func (o *PKICertificateIssueDetails) HasCertificateAuthorityMode() bool`

HasCertificateAuthorityMode returns a boolean if a field has been set.

### GetClientFlag

`func (o *PKICertificateIssueDetails) GetClientFlag() bool`

GetClientFlag returns the ClientFlag field if non-nil, zero value otherwise.

### GetClientFlagOk

`func (o *PKICertificateIssueDetails) GetClientFlagOk() (*bool, bool)`

GetClientFlagOk returns a tuple with the ClientFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlag

`func (o *PKICertificateIssueDetails) SetClientFlag(v bool)`

SetClientFlag sets ClientFlag field to given value.

### HasClientFlag

`func (o *PKICertificateIssueDetails) HasClientFlag() bool`

HasClientFlag returns a boolean if a field has been set.

### GetCodeSigningFlag

`func (o *PKICertificateIssueDetails) GetCodeSigningFlag() bool`

GetCodeSigningFlag returns the CodeSigningFlag field if non-nil, zero value otherwise.

### GetCodeSigningFlagOk

`func (o *PKICertificateIssueDetails) GetCodeSigningFlagOk() (*bool, bool)`

GetCodeSigningFlagOk returns a tuple with the CodeSigningFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSigningFlag

`func (o *PKICertificateIssueDetails) SetCodeSigningFlag(v bool)`

SetCodeSigningFlag sets CodeSigningFlag field to given value.

### HasCodeSigningFlag

`func (o *PKICertificateIssueDetails) HasCodeSigningFlag() bool`

HasCodeSigningFlag returns a boolean if a field has been set.

### GetCountry

`func (o *PKICertificateIssueDetails) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PKICertificateIssueDetails) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PKICertificateIssueDetails) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PKICertificateIssueDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatePrivateCrl

`func (o *PKICertificateIssueDetails) GetCreatePrivateCrl() bool`

GetCreatePrivateCrl returns the CreatePrivateCrl field if non-nil, zero value otherwise.

### GetCreatePrivateCrlOk

`func (o *PKICertificateIssueDetails) GetCreatePrivateCrlOk() (*bool, bool)`

GetCreatePrivateCrlOk returns a tuple with the CreatePrivateCrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePrivateCrl

`func (o *PKICertificateIssueDetails) SetCreatePrivateCrl(v bool)`

SetCreatePrivateCrl sets CreatePrivateCrl field to given value.

### HasCreatePrivateCrl

`func (o *PKICertificateIssueDetails) HasCreatePrivateCrl() bool`

HasCreatePrivateCrl returns a boolean if a field has been set.

### GetCreatePublicCrl

`func (o *PKICertificateIssueDetails) GetCreatePublicCrl() bool`

GetCreatePublicCrl returns the CreatePublicCrl field if non-nil, zero value otherwise.

### GetCreatePublicCrlOk

`func (o *PKICertificateIssueDetails) GetCreatePublicCrlOk() (*bool, bool)`

GetCreatePublicCrlOk returns a tuple with the CreatePublicCrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePublicCrl

`func (o *PKICertificateIssueDetails) SetCreatePublicCrl(v bool)`

SetCreatePublicCrl sets CreatePublicCrl field to given value.

### HasCreatePublicCrl

`func (o *PKICertificateIssueDetails) HasCreatePublicCrl() bool`

HasCreatePublicCrl returns a boolean if a field has been set.

### GetDestinationPath

`func (o *PKICertificateIssueDetails) GetDestinationPath() string`

GetDestinationPath returns the DestinationPath field if non-nil, zero value otherwise.

### GetDestinationPathOk

`func (o *PKICertificateIssueDetails) GetDestinationPathOk() (*string, bool)`

GetDestinationPathOk returns a tuple with the DestinationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPath

`func (o *PKICertificateIssueDetails) SetDestinationPath(v string)`

SetDestinationPath sets DestinationPath field to given value.

### HasDestinationPath

`func (o *PKICertificateIssueDetails) HasDestinationPath() bool`

HasDestinationPath returns a boolean if a field has been set.

### GetEnforceHostnames

`func (o *PKICertificateIssueDetails) GetEnforceHostnames() bool`

GetEnforceHostnames returns the EnforceHostnames field if non-nil, zero value otherwise.

### GetEnforceHostnamesOk

`func (o *PKICertificateIssueDetails) GetEnforceHostnamesOk() (*bool, bool)`

GetEnforceHostnamesOk returns a tuple with the EnforceHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceHostnames

`func (o *PKICertificateIssueDetails) SetEnforceHostnames(v bool)`

SetEnforceHostnames sets EnforceHostnames field to given value.

### HasEnforceHostnames

`func (o *PKICertificateIssueDetails) HasEnforceHostnames() bool`

HasEnforceHostnames returns a boolean if a field has been set.

### GetExpirationEvents

`func (o *PKICertificateIssueDetails) GetExpirationEvents() []CertificateExpirationEvent`

GetExpirationEvents returns the ExpirationEvents field if non-nil, zero value otherwise.

### GetExpirationEventsOk

`func (o *PKICertificateIssueDetails) GetExpirationEventsOk() (*[]CertificateExpirationEvent, bool)`

GetExpirationEventsOk returns a tuple with the ExpirationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEvents

`func (o *PKICertificateIssueDetails) SetExpirationEvents(v []CertificateExpirationEvent)`

SetExpirationEvents sets ExpirationEvents field to given value.

### HasExpirationEvents

`func (o *PKICertificateIssueDetails) HasExpirationEvents() bool`

HasExpirationEvents returns a boolean if a field has been set.

### GetGwClusterId

`func (o *PKICertificateIssueDetails) GetGwClusterId() int64`

GetGwClusterId returns the GwClusterId field if non-nil, zero value otherwise.

### GetGwClusterIdOk

`func (o *PKICertificateIssueDetails) GetGwClusterIdOk() (*int64, bool)`

GetGwClusterIdOk returns a tuple with the GwClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwClusterId

`func (o *PKICertificateIssueDetails) SetGwClusterId(v int64)`

SetGwClusterId sets GwClusterId field to given value.

### HasGwClusterId

`func (o *PKICertificateIssueDetails) HasGwClusterId() bool`

HasGwClusterId returns a boolean if a field has been set.

### GetGwClusterUrl

`func (o *PKICertificateIssueDetails) GetGwClusterUrl() string`

GetGwClusterUrl returns the GwClusterUrl field if non-nil, zero value otherwise.

### GetGwClusterUrlOk

`func (o *PKICertificateIssueDetails) GetGwClusterUrlOk() (*string, bool)`

GetGwClusterUrlOk returns a tuple with the GwClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwClusterUrl

`func (o *PKICertificateIssueDetails) SetGwClusterUrl(v string)`

SetGwClusterUrl sets GwClusterUrl field to given value.

### HasGwClusterUrl

`func (o *PKICertificateIssueDetails) HasGwClusterUrl() bool`

HasGwClusterUrl returns a boolean if a field has been set.

### GetIsCa

`func (o *PKICertificateIssueDetails) GetIsCa() bool`

GetIsCa returns the IsCa field if non-nil, zero value otherwise.

### GetIsCaOk

`func (o *PKICertificateIssueDetails) GetIsCaOk() (*bool, bool)`

GetIsCaOk returns a tuple with the IsCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCa

`func (o *PKICertificateIssueDetails) SetIsCa(v bool)`

SetIsCa sets IsCa field to given value.

### HasIsCa

`func (o *PKICertificateIssueDetails) HasIsCa() bool`

HasIsCa returns a boolean if a field has been set.

### GetKeyBits

`func (o *PKICertificateIssueDetails) GetKeyBits() int64`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *PKICertificateIssueDetails) GetKeyBitsOk() (*int64, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *PKICertificateIssueDetails) SetKeyBits(v int64)`

SetKeyBits sets KeyBits field to given value.

### HasKeyBits

`func (o *PKICertificateIssueDetails) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.

### GetKeyType

`func (o *PKICertificateIssueDetails) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PKICertificateIssueDetails) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PKICertificateIssueDetails) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *PKICertificateIssueDetails) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetKeyUsageList

`func (o *PKICertificateIssueDetails) GetKeyUsageList() []string`

GetKeyUsageList returns the KeyUsageList field if non-nil, zero value otherwise.

### GetKeyUsageListOk

`func (o *PKICertificateIssueDetails) GetKeyUsageListOk() (*[]string, bool)`

GetKeyUsageListOk returns a tuple with the KeyUsageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsageList

`func (o *PKICertificateIssueDetails) SetKeyUsageList(v []string)`

SetKeyUsageList sets KeyUsageList field to given value.

### HasKeyUsageList

`func (o *PKICertificateIssueDetails) HasKeyUsageList() bool`

HasKeyUsageList returns a boolean if a field has been set.

### GetLocality

`func (o *PKICertificateIssueDetails) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PKICertificateIssueDetails) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PKICertificateIssueDetails) SetLocality(v []string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *PKICertificateIssueDetails) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetNotBeforeDuration

`func (o *PKICertificateIssueDetails) GetNotBeforeDuration() int64`

GetNotBeforeDuration returns the NotBeforeDuration field if non-nil, zero value otherwise.

### GetNotBeforeDurationOk

`func (o *PKICertificateIssueDetails) GetNotBeforeDurationOk() (*int64, bool)`

GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDuration

`func (o *PKICertificateIssueDetails) SetNotBeforeDuration(v int64)`

SetNotBeforeDuration sets NotBeforeDuration field to given value.

### HasNotBeforeDuration

`func (o *PKICertificateIssueDetails) HasNotBeforeDuration() bool`

HasNotBeforeDuration returns a boolean if a field has been set.

### GetOrganizationList

`func (o *PKICertificateIssueDetails) GetOrganizationList() []string`

GetOrganizationList returns the OrganizationList field if non-nil, zero value otherwise.

### GetOrganizationListOk

`func (o *PKICertificateIssueDetails) GetOrganizationListOk() (*[]string, bool)`

GetOrganizationListOk returns a tuple with the OrganizationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationList

`func (o *PKICertificateIssueDetails) SetOrganizationList(v []string)`

SetOrganizationList sets OrganizationList field to given value.

### HasOrganizationList

`func (o *PKICertificateIssueDetails) HasOrganizationList() bool`

HasOrganizationList returns a boolean if a field has been set.

### GetOrganizationUnitList

`func (o *PKICertificateIssueDetails) GetOrganizationUnitList() []string`

GetOrganizationUnitList returns the OrganizationUnitList field if non-nil, zero value otherwise.

### GetOrganizationUnitListOk

`func (o *PKICertificateIssueDetails) GetOrganizationUnitListOk() (*[]string, bool)`

GetOrganizationUnitListOk returns a tuple with the OrganizationUnitList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnitList

`func (o *PKICertificateIssueDetails) SetOrganizationUnitList(v []string)`

SetOrganizationUnitList sets OrganizationUnitList field to given value.

### HasOrganizationUnitList

`func (o *PKICertificateIssueDetails) HasOrganizationUnitList() bool`

HasOrganizationUnitList returns a boolean if a field has been set.

### GetPostalCode

`func (o *PKICertificateIssueDetails) GetPostalCode() []string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PKICertificateIssueDetails) GetPostalCodeOk() (*[]string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PKICertificateIssueDetails) SetPostalCode(v []string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PKICertificateIssueDetails) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetProtectGeneratedCertificates

`func (o *PKICertificateIssueDetails) GetProtectGeneratedCertificates() bool`

GetProtectGeneratedCertificates returns the ProtectGeneratedCertificates field if non-nil, zero value otherwise.

### GetProtectGeneratedCertificatesOk

`func (o *PKICertificateIssueDetails) GetProtectGeneratedCertificatesOk() (*bool, bool)`

GetProtectGeneratedCertificatesOk returns a tuple with the ProtectGeneratedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectGeneratedCertificates

`func (o *PKICertificateIssueDetails) SetProtectGeneratedCertificates(v bool)`

SetProtectGeneratedCertificates sets ProtectGeneratedCertificates field to given value.

### HasProtectGeneratedCertificates

`func (o *PKICertificateIssueDetails) HasProtectGeneratedCertificates() bool`

HasProtectGeneratedCertificates returns a boolean if a field has been set.

### GetProvince

`func (o *PKICertificateIssueDetails) GetProvince() []string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PKICertificateIssueDetails) GetProvinceOk() (*[]string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PKICertificateIssueDetails) SetProvince(v []string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *PKICertificateIssueDetails) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetRequireCn

`func (o *PKICertificateIssueDetails) GetRequireCn() bool`

GetRequireCn returns the RequireCn field if non-nil, zero value otherwise.

### GetRequireCnOk

`func (o *PKICertificateIssueDetails) GetRequireCnOk() (*bool, bool)`

GetRequireCnOk returns a tuple with the RequireCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCn

`func (o *PKICertificateIssueDetails) SetRequireCn(v bool)`

SetRequireCn sets RequireCn field to given value.

### HasRequireCn

`func (o *PKICertificateIssueDetails) HasRequireCn() bool`

HasRequireCn returns a boolean if a field has been set.

### GetServerFlag

`func (o *PKICertificateIssueDetails) GetServerFlag() bool`

GetServerFlag returns the ServerFlag field if non-nil, zero value otherwise.

### GetServerFlagOk

`func (o *PKICertificateIssueDetails) GetServerFlagOk() (*bool, bool)`

GetServerFlagOk returns a tuple with the ServerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFlag

`func (o *PKICertificateIssueDetails) SetServerFlag(v bool)`

SetServerFlag sets ServerFlag field to given value.

### HasServerFlag

`func (o *PKICertificateIssueDetails) HasServerFlag() bool`

HasServerFlag returns a boolean if a field has been set.

### GetStreetAddress

`func (o *PKICertificateIssueDetails) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *PKICertificateIssueDetails) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *PKICertificateIssueDetails) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *PKICertificateIssueDetails) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


