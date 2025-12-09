# UpdatePKICertIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**AllowAnyName** | Pointer to **bool** | If set, clients can request certificates for any CN | [optional] 
**AllowCopyExtFromCsr** | Pointer to **bool** | If set, will allow copying the extra extensions from the csr file (if given) | [optional] 
**AllowSubdomains** | Pointer to **bool** | If set, clients can request certificates for subdomains of the allowed domains | [optional] 
**AllowedDomains** | Pointer to **string** | A list of the allowed domains that clients can request to be included in the certificate (in a comma-delimited list) | [optional] 
**AllowedExtraExtensions** | Pointer to **string** | A json string containing the allowed extra extensions for the pki cert issuer | [optional] 
**AllowedIpSans** | Pointer to **string** | A list of the allowed CIDRs for ips that clients can request to be included in the certificate as part of the IP Subject Alternative Names (in a comma-delimited list) | [optional] 
**AllowedUriSans** | Pointer to **string** | A list of the allowed URIs that clients can request to be included in the certificate as part of the URI Subject Alternative Names (in a comma-delimited list) | [optional] 
**AutoRenew** | Pointer to **bool** | Automatically renew certificates before expiration | [optional] 
**ClientFlag** | Pointer to **bool** | If set, certificates will be flagged for client auth use | [optional] 
**CodeSigningFlag** | Pointer to **bool** | If set, certificates will be flagged for code signing use | [optional] 
**Country** | Pointer to **string** | A comma-separated list of countries that will be set in the issued certificate | [optional] 
**CreatePrivateCrl** | Pointer to **bool** | Set this to allow the issuer will expose a CRL endpoint in the Gateway | [optional] 
**CreatePrivateOcsp** | Pointer to **bool** | Set this to enable an OCSP endpoint in the Gateway and include its URL in AIA | [optional] 
**CreatePublicCrl** | Pointer to **bool** | Set this to allow the cert issuer will expose a public CRL endpoint | [optional] 
**CreatePublicOcsp** | Pointer to **bool** | Set this to enable a public OCSP endpoint and include its URL in AIA (served by UAM and includes account id) | [optional] 
**CriticalKeyUsage** | Pointer to **string** | Mark key usage as critical [true/false] | [optional] [default to "true"]
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**DestinationPath** | Pointer to **string** | A path in which to save generated certificates | [optional] 
**DisableWildcards** | Pointer to **bool** | If set, generation of wildcard certificates will be disabled. | [optional] 
**EnableAcme** | Pointer to **bool** | If set, the cert issuer will support the acme protocol | [optional] 
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the certificate would you like to be notified. | [optional] 
**GwClusterUrl** | Pointer to **string** | The GW cluster URL to issue the certificate from. Required in Public CA mode, to allow CRLs on private CA, or to enable ACME | [optional] 
**IsCa** | Pointer to **bool** | If set, the basic constraints extension will be added to certificate | [optional] 
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyUsage** | Pointer to **string** | key-usage | [optional] [default to "DigitalSignature,KeyAgreement,KeyEncipherment"]
**Locality** | Pointer to **string** | A comma-separated list of localities that will be set in the issued certificate | [optional] 
**MaxPathLen** | Pointer to **int64** | The maximum path length for the generated certificate. -1, means unlimited | [optional] [default to -1]
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | PKI certificate issuer name | 
**NewName** | Pointer to **string** | New item name | [optional] 
**NotEnforceHostnames** | Pointer to **bool** | If set, any names are allowed for CN and SANs in the certificate and not only a valid host name | [optional] 
**NotRequireCn** | Pointer to **bool** | If set, clients can request certificates without a CN | [optional] 
**OcspTtl** | Pointer to **string** | OCSP NextUpdate window for OCSP responses (min 10m). Supports s,m,h,d suffix. | [optional] 
**OrganizationalUnits** | Pointer to **string** | A comma-separated list of organizational units (OU) that will be set in the issued certificate | [optional] 
**Organizations** | Pointer to **string** | A comma-separated list of organizations (O) that will be set in the issued certificate | [optional] 
**PostalCode** | Pointer to **string** | A comma-separated list of postal codes that will be set in the issued certificate | [optional] 
**ProtectCertificates** | Pointer to **bool** | Whether to protect generated certificates from deletion | [optional] 
**Province** | Pointer to **string** | A comma-separated list of provinces that will be set in the issued certificate | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**ScheduledRenew** | Pointer to **int64** | Number of days before expiration to renew certificates | [optional] 
**ServerFlag** | Pointer to **bool** | If set, certificates will be flagged for server auth use | [optional] 
**SignerKeyName** | Pointer to **string** | A key to sign the certificate with, required in Private CA mode | [optional] 
**StreetAddress** | Pointer to **string** | A comma-separated list of street addresses that will be set in the issued certificate | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | **string** | The maximum requested Time To Live for issued certificates, in seconds. In case of Public CA, this is based on the CA target&#39;s supported maximum TTLs | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdatePKICertIssuer

`func NewUpdatePKICertIssuer(name string, ttl string, ) *UpdatePKICertIssuer`

NewUpdatePKICertIssuer instantiates a new UpdatePKICertIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePKICertIssuerWithDefaults

`func NewUpdatePKICertIssuerWithDefaults() *UpdatePKICertIssuer`

NewUpdatePKICertIssuerWithDefaults instantiates a new UpdatePKICertIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *UpdatePKICertIssuer) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *UpdatePKICertIssuer) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *UpdatePKICertIssuer) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *UpdatePKICertIssuer) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetAllowAnyName

`func (o *UpdatePKICertIssuer) GetAllowAnyName() bool`

GetAllowAnyName returns the AllowAnyName field if non-nil, zero value otherwise.

### GetAllowAnyNameOk

`func (o *UpdatePKICertIssuer) GetAllowAnyNameOk() (*bool, bool)`

GetAllowAnyNameOk returns a tuple with the AllowAnyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnyName

`func (o *UpdatePKICertIssuer) SetAllowAnyName(v bool)`

SetAllowAnyName sets AllowAnyName field to given value.

### HasAllowAnyName

`func (o *UpdatePKICertIssuer) HasAllowAnyName() bool`

HasAllowAnyName returns a boolean if a field has been set.

### GetAllowCopyExtFromCsr

`func (o *UpdatePKICertIssuer) GetAllowCopyExtFromCsr() bool`

GetAllowCopyExtFromCsr returns the AllowCopyExtFromCsr field if non-nil, zero value otherwise.

### GetAllowCopyExtFromCsrOk

`func (o *UpdatePKICertIssuer) GetAllowCopyExtFromCsrOk() (*bool, bool)`

GetAllowCopyExtFromCsrOk returns a tuple with the AllowCopyExtFromCsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCopyExtFromCsr

`func (o *UpdatePKICertIssuer) SetAllowCopyExtFromCsr(v bool)`

SetAllowCopyExtFromCsr sets AllowCopyExtFromCsr field to given value.

### HasAllowCopyExtFromCsr

`func (o *UpdatePKICertIssuer) HasAllowCopyExtFromCsr() bool`

HasAllowCopyExtFromCsr returns a boolean if a field has been set.

### GetAllowSubdomains

`func (o *UpdatePKICertIssuer) GetAllowSubdomains() bool`

GetAllowSubdomains returns the AllowSubdomains field if non-nil, zero value otherwise.

### GetAllowSubdomainsOk

`func (o *UpdatePKICertIssuer) GetAllowSubdomainsOk() (*bool, bool)`

GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubdomains

`func (o *UpdatePKICertIssuer) SetAllowSubdomains(v bool)`

SetAllowSubdomains sets AllowSubdomains field to given value.

### HasAllowSubdomains

`func (o *UpdatePKICertIssuer) HasAllowSubdomains() bool`

HasAllowSubdomains returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *UpdatePKICertIssuer) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *UpdatePKICertIssuer) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *UpdatePKICertIssuer) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *UpdatePKICertIssuer) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAllowedExtraExtensions

`func (o *UpdatePKICertIssuer) GetAllowedExtraExtensions() string`

GetAllowedExtraExtensions returns the AllowedExtraExtensions field if non-nil, zero value otherwise.

### GetAllowedExtraExtensionsOk

`func (o *UpdatePKICertIssuer) GetAllowedExtraExtensionsOk() (*string, bool)`

GetAllowedExtraExtensionsOk returns a tuple with the AllowedExtraExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExtraExtensions

`func (o *UpdatePKICertIssuer) SetAllowedExtraExtensions(v string)`

SetAllowedExtraExtensions sets AllowedExtraExtensions field to given value.

### HasAllowedExtraExtensions

`func (o *UpdatePKICertIssuer) HasAllowedExtraExtensions() bool`

HasAllowedExtraExtensions returns a boolean if a field has been set.

### GetAllowedIpSans

`func (o *UpdatePKICertIssuer) GetAllowedIpSans() string`

GetAllowedIpSans returns the AllowedIpSans field if non-nil, zero value otherwise.

### GetAllowedIpSansOk

`func (o *UpdatePKICertIssuer) GetAllowedIpSansOk() (*string, bool)`

GetAllowedIpSansOk returns a tuple with the AllowedIpSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpSans

`func (o *UpdatePKICertIssuer) SetAllowedIpSans(v string)`

SetAllowedIpSans sets AllowedIpSans field to given value.

### HasAllowedIpSans

`func (o *UpdatePKICertIssuer) HasAllowedIpSans() bool`

HasAllowedIpSans returns a boolean if a field has been set.

### GetAllowedUriSans

`func (o *UpdatePKICertIssuer) GetAllowedUriSans() string`

GetAllowedUriSans returns the AllowedUriSans field if non-nil, zero value otherwise.

### GetAllowedUriSansOk

`func (o *UpdatePKICertIssuer) GetAllowedUriSansOk() (*string, bool)`

GetAllowedUriSansOk returns a tuple with the AllowedUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUriSans

`func (o *UpdatePKICertIssuer) SetAllowedUriSans(v string)`

SetAllowedUriSans sets AllowedUriSans field to given value.

### HasAllowedUriSans

`func (o *UpdatePKICertIssuer) HasAllowedUriSans() bool`

HasAllowedUriSans returns a boolean if a field has been set.

### GetAutoRenew

`func (o *UpdatePKICertIssuer) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *UpdatePKICertIssuer) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *UpdatePKICertIssuer) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.

### HasAutoRenew

`func (o *UpdatePKICertIssuer) HasAutoRenew() bool`

HasAutoRenew returns a boolean if a field has been set.

### GetClientFlag

`func (o *UpdatePKICertIssuer) GetClientFlag() bool`

GetClientFlag returns the ClientFlag field if non-nil, zero value otherwise.

### GetClientFlagOk

`func (o *UpdatePKICertIssuer) GetClientFlagOk() (*bool, bool)`

GetClientFlagOk returns a tuple with the ClientFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlag

`func (o *UpdatePKICertIssuer) SetClientFlag(v bool)`

SetClientFlag sets ClientFlag field to given value.

### HasClientFlag

`func (o *UpdatePKICertIssuer) HasClientFlag() bool`

HasClientFlag returns a boolean if a field has been set.

### GetCodeSigningFlag

`func (o *UpdatePKICertIssuer) GetCodeSigningFlag() bool`

GetCodeSigningFlag returns the CodeSigningFlag field if non-nil, zero value otherwise.

### GetCodeSigningFlagOk

`func (o *UpdatePKICertIssuer) GetCodeSigningFlagOk() (*bool, bool)`

GetCodeSigningFlagOk returns a tuple with the CodeSigningFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSigningFlag

`func (o *UpdatePKICertIssuer) SetCodeSigningFlag(v bool)`

SetCodeSigningFlag sets CodeSigningFlag field to given value.

### HasCodeSigningFlag

`func (o *UpdatePKICertIssuer) HasCodeSigningFlag() bool`

HasCodeSigningFlag returns a boolean if a field has been set.

### GetCountry

`func (o *UpdatePKICertIssuer) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdatePKICertIssuer) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdatePKICertIssuer) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdatePKICertIssuer) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatePrivateCrl

`func (o *UpdatePKICertIssuer) GetCreatePrivateCrl() bool`

GetCreatePrivateCrl returns the CreatePrivateCrl field if non-nil, zero value otherwise.

### GetCreatePrivateCrlOk

`func (o *UpdatePKICertIssuer) GetCreatePrivateCrlOk() (*bool, bool)`

GetCreatePrivateCrlOk returns a tuple with the CreatePrivateCrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePrivateCrl

`func (o *UpdatePKICertIssuer) SetCreatePrivateCrl(v bool)`

SetCreatePrivateCrl sets CreatePrivateCrl field to given value.

### HasCreatePrivateCrl

`func (o *UpdatePKICertIssuer) HasCreatePrivateCrl() bool`

HasCreatePrivateCrl returns a boolean if a field has been set.

### GetCreatePrivateOcsp

`func (o *UpdatePKICertIssuer) GetCreatePrivateOcsp() bool`

GetCreatePrivateOcsp returns the CreatePrivateOcsp field if non-nil, zero value otherwise.

### GetCreatePrivateOcspOk

`func (o *UpdatePKICertIssuer) GetCreatePrivateOcspOk() (*bool, bool)`

GetCreatePrivateOcspOk returns a tuple with the CreatePrivateOcsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePrivateOcsp

`func (o *UpdatePKICertIssuer) SetCreatePrivateOcsp(v bool)`

SetCreatePrivateOcsp sets CreatePrivateOcsp field to given value.

### HasCreatePrivateOcsp

`func (o *UpdatePKICertIssuer) HasCreatePrivateOcsp() bool`

HasCreatePrivateOcsp returns a boolean if a field has been set.

### GetCreatePublicCrl

`func (o *UpdatePKICertIssuer) GetCreatePublicCrl() bool`

GetCreatePublicCrl returns the CreatePublicCrl field if non-nil, zero value otherwise.

### GetCreatePublicCrlOk

`func (o *UpdatePKICertIssuer) GetCreatePublicCrlOk() (*bool, bool)`

GetCreatePublicCrlOk returns a tuple with the CreatePublicCrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePublicCrl

`func (o *UpdatePKICertIssuer) SetCreatePublicCrl(v bool)`

SetCreatePublicCrl sets CreatePublicCrl field to given value.

### HasCreatePublicCrl

`func (o *UpdatePKICertIssuer) HasCreatePublicCrl() bool`

HasCreatePublicCrl returns a boolean if a field has been set.

### GetCreatePublicOcsp

`func (o *UpdatePKICertIssuer) GetCreatePublicOcsp() bool`

GetCreatePublicOcsp returns the CreatePublicOcsp field if non-nil, zero value otherwise.

### GetCreatePublicOcspOk

`func (o *UpdatePKICertIssuer) GetCreatePublicOcspOk() (*bool, bool)`

GetCreatePublicOcspOk returns a tuple with the CreatePublicOcsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePublicOcsp

`func (o *UpdatePKICertIssuer) SetCreatePublicOcsp(v bool)`

SetCreatePublicOcsp sets CreatePublicOcsp field to given value.

### HasCreatePublicOcsp

`func (o *UpdatePKICertIssuer) HasCreatePublicOcsp() bool`

HasCreatePublicOcsp returns a boolean if a field has been set.

### GetCriticalKeyUsage

`func (o *UpdatePKICertIssuer) GetCriticalKeyUsage() string`

GetCriticalKeyUsage returns the CriticalKeyUsage field if non-nil, zero value otherwise.

### GetCriticalKeyUsageOk

`func (o *UpdatePKICertIssuer) GetCriticalKeyUsageOk() (*string, bool)`

GetCriticalKeyUsageOk returns a tuple with the CriticalKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalKeyUsage

`func (o *UpdatePKICertIssuer) SetCriticalKeyUsage(v string)`

SetCriticalKeyUsage sets CriticalKeyUsage field to given value.

### HasCriticalKeyUsage

`func (o *UpdatePKICertIssuer) HasCriticalKeyUsage() bool`

HasCriticalKeyUsage returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *UpdatePKICertIssuer) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *UpdatePKICertIssuer) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *UpdatePKICertIssuer) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *UpdatePKICertIssuer) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePKICertIssuer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePKICertIssuer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePKICertIssuer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePKICertIssuer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinationPath

`func (o *UpdatePKICertIssuer) GetDestinationPath() string`

GetDestinationPath returns the DestinationPath field if non-nil, zero value otherwise.

### GetDestinationPathOk

`func (o *UpdatePKICertIssuer) GetDestinationPathOk() (*string, bool)`

GetDestinationPathOk returns a tuple with the DestinationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPath

`func (o *UpdatePKICertIssuer) SetDestinationPath(v string)`

SetDestinationPath sets DestinationPath field to given value.

### HasDestinationPath

`func (o *UpdatePKICertIssuer) HasDestinationPath() bool`

HasDestinationPath returns a boolean if a field has been set.

### GetDisableWildcards

`func (o *UpdatePKICertIssuer) GetDisableWildcards() bool`

GetDisableWildcards returns the DisableWildcards field if non-nil, zero value otherwise.

### GetDisableWildcardsOk

`func (o *UpdatePKICertIssuer) GetDisableWildcardsOk() (*bool, bool)`

GetDisableWildcardsOk returns a tuple with the DisableWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableWildcards

`func (o *UpdatePKICertIssuer) SetDisableWildcards(v bool)`

SetDisableWildcards sets DisableWildcards field to given value.

### HasDisableWildcards

`func (o *UpdatePKICertIssuer) HasDisableWildcards() bool`

HasDisableWildcards returns a boolean if a field has been set.

### GetEnableAcme

`func (o *UpdatePKICertIssuer) GetEnableAcme() bool`

GetEnableAcme returns the EnableAcme field if non-nil, zero value otherwise.

### GetEnableAcmeOk

`func (o *UpdatePKICertIssuer) GetEnableAcmeOk() (*bool, bool)`

GetEnableAcmeOk returns a tuple with the EnableAcme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAcme

`func (o *UpdatePKICertIssuer) SetEnableAcme(v bool)`

SetEnableAcme sets EnableAcme field to given value.

### HasEnableAcme

`func (o *UpdatePKICertIssuer) HasEnableAcme() bool`

HasEnableAcme returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *UpdatePKICertIssuer) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *UpdatePKICertIssuer) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *UpdatePKICertIssuer) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *UpdatePKICertIssuer) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetGwClusterUrl

`func (o *UpdatePKICertIssuer) GetGwClusterUrl() string`

GetGwClusterUrl returns the GwClusterUrl field if non-nil, zero value otherwise.

### GetGwClusterUrlOk

`func (o *UpdatePKICertIssuer) GetGwClusterUrlOk() (*string, bool)`

GetGwClusterUrlOk returns a tuple with the GwClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwClusterUrl

`func (o *UpdatePKICertIssuer) SetGwClusterUrl(v string)`

SetGwClusterUrl sets GwClusterUrl field to given value.

### HasGwClusterUrl

`func (o *UpdatePKICertIssuer) HasGwClusterUrl() bool`

HasGwClusterUrl returns a boolean if a field has been set.

### GetIsCa

`func (o *UpdatePKICertIssuer) GetIsCa() bool`

GetIsCa returns the IsCa field if non-nil, zero value otherwise.

### GetIsCaOk

`func (o *UpdatePKICertIssuer) GetIsCaOk() (*bool, bool)`

GetIsCaOk returns a tuple with the IsCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCa

`func (o *UpdatePKICertIssuer) SetIsCa(v bool)`

SetIsCa sets IsCa field to given value.

### HasIsCa

`func (o *UpdatePKICertIssuer) HasIsCa() bool`

HasIsCa returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *UpdatePKICertIssuer) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *UpdatePKICertIssuer) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *UpdatePKICertIssuer) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *UpdatePKICertIssuer) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *UpdatePKICertIssuer) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdatePKICertIssuer) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdatePKICertIssuer) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdatePKICertIssuer) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyUsage

`func (o *UpdatePKICertIssuer) GetKeyUsage() string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *UpdatePKICertIssuer) GetKeyUsageOk() (*string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *UpdatePKICertIssuer) SetKeyUsage(v string)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *UpdatePKICertIssuer) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetLocality

`func (o *UpdatePKICertIssuer) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *UpdatePKICertIssuer) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *UpdatePKICertIssuer) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *UpdatePKICertIssuer) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetMaxPathLen

`func (o *UpdatePKICertIssuer) GetMaxPathLen() int64`

GetMaxPathLen returns the MaxPathLen field if non-nil, zero value otherwise.

### GetMaxPathLenOk

`func (o *UpdatePKICertIssuer) GetMaxPathLenOk() (*int64, bool)`

GetMaxPathLenOk returns a tuple with the MaxPathLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPathLen

`func (o *UpdatePKICertIssuer) SetMaxPathLen(v int64)`

SetMaxPathLen sets MaxPathLen field to given value.

### HasMaxPathLen

`func (o *UpdatePKICertIssuer) HasMaxPathLen() bool`

HasMaxPathLen returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdatePKICertIssuer) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdatePKICertIssuer) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdatePKICertIssuer) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdatePKICertIssuer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UpdatePKICertIssuer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePKICertIssuer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePKICertIssuer) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdatePKICertIssuer) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdatePKICertIssuer) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdatePKICertIssuer) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdatePKICertIssuer) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetNotEnforceHostnames

`func (o *UpdatePKICertIssuer) GetNotEnforceHostnames() bool`

GetNotEnforceHostnames returns the NotEnforceHostnames field if non-nil, zero value otherwise.

### GetNotEnforceHostnamesOk

`func (o *UpdatePKICertIssuer) GetNotEnforceHostnamesOk() (*bool, bool)`

GetNotEnforceHostnamesOk returns a tuple with the NotEnforceHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEnforceHostnames

`func (o *UpdatePKICertIssuer) SetNotEnforceHostnames(v bool)`

SetNotEnforceHostnames sets NotEnforceHostnames field to given value.

### HasNotEnforceHostnames

`func (o *UpdatePKICertIssuer) HasNotEnforceHostnames() bool`

HasNotEnforceHostnames returns a boolean if a field has been set.

### GetNotRequireCn

`func (o *UpdatePKICertIssuer) GetNotRequireCn() bool`

GetNotRequireCn returns the NotRequireCn field if non-nil, zero value otherwise.

### GetNotRequireCnOk

`func (o *UpdatePKICertIssuer) GetNotRequireCnOk() (*bool, bool)`

GetNotRequireCnOk returns a tuple with the NotRequireCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRequireCn

`func (o *UpdatePKICertIssuer) SetNotRequireCn(v bool)`

SetNotRequireCn sets NotRequireCn field to given value.

### HasNotRequireCn

`func (o *UpdatePKICertIssuer) HasNotRequireCn() bool`

HasNotRequireCn returns a boolean if a field has been set.

### GetOcspTtl

`func (o *UpdatePKICertIssuer) GetOcspTtl() string`

GetOcspTtl returns the OcspTtl field if non-nil, zero value otherwise.

### GetOcspTtlOk

`func (o *UpdatePKICertIssuer) GetOcspTtlOk() (*string, bool)`

GetOcspTtlOk returns a tuple with the OcspTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspTtl

`func (o *UpdatePKICertIssuer) SetOcspTtl(v string)`

SetOcspTtl sets OcspTtl field to given value.

### HasOcspTtl

`func (o *UpdatePKICertIssuer) HasOcspTtl() bool`

HasOcspTtl returns a boolean if a field has been set.

### GetOrganizationalUnits

`func (o *UpdatePKICertIssuer) GetOrganizationalUnits() string`

GetOrganizationalUnits returns the OrganizationalUnits field if non-nil, zero value otherwise.

### GetOrganizationalUnitsOk

`func (o *UpdatePKICertIssuer) GetOrganizationalUnitsOk() (*string, bool)`

GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnits

`func (o *UpdatePKICertIssuer) SetOrganizationalUnits(v string)`

SetOrganizationalUnits sets OrganizationalUnits field to given value.

### HasOrganizationalUnits

`func (o *UpdatePKICertIssuer) HasOrganizationalUnits() bool`

HasOrganizationalUnits returns a boolean if a field has been set.

### GetOrganizations

`func (o *UpdatePKICertIssuer) GetOrganizations() string`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *UpdatePKICertIssuer) GetOrganizationsOk() (*string, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *UpdatePKICertIssuer) SetOrganizations(v string)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *UpdatePKICertIssuer) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetPostalCode

`func (o *UpdatePKICertIssuer) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UpdatePKICertIssuer) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UpdatePKICertIssuer) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UpdatePKICertIssuer) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetProtectCertificates

`func (o *UpdatePKICertIssuer) GetProtectCertificates() bool`

GetProtectCertificates returns the ProtectCertificates field if non-nil, zero value otherwise.

### GetProtectCertificatesOk

`func (o *UpdatePKICertIssuer) GetProtectCertificatesOk() (*bool, bool)`

GetProtectCertificatesOk returns a tuple with the ProtectCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectCertificates

`func (o *UpdatePKICertIssuer) SetProtectCertificates(v bool)`

SetProtectCertificates sets ProtectCertificates field to given value.

### HasProtectCertificates

`func (o *UpdatePKICertIssuer) HasProtectCertificates() bool`

HasProtectCertificates returns a boolean if a field has been set.

### GetProvince

`func (o *UpdatePKICertIssuer) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *UpdatePKICertIssuer) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *UpdatePKICertIssuer) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *UpdatePKICertIssuer) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetRmTag

`func (o *UpdatePKICertIssuer) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *UpdatePKICertIssuer) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *UpdatePKICertIssuer) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *UpdatePKICertIssuer) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetScheduledRenew

`func (o *UpdatePKICertIssuer) GetScheduledRenew() int64`

GetScheduledRenew returns the ScheduledRenew field if non-nil, zero value otherwise.

### GetScheduledRenewOk

`func (o *UpdatePKICertIssuer) GetScheduledRenewOk() (*int64, bool)`

GetScheduledRenewOk returns a tuple with the ScheduledRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRenew

`func (o *UpdatePKICertIssuer) SetScheduledRenew(v int64)`

SetScheduledRenew sets ScheduledRenew field to given value.

### HasScheduledRenew

`func (o *UpdatePKICertIssuer) HasScheduledRenew() bool`

HasScheduledRenew returns a boolean if a field has been set.

### GetServerFlag

`func (o *UpdatePKICertIssuer) GetServerFlag() bool`

GetServerFlag returns the ServerFlag field if non-nil, zero value otherwise.

### GetServerFlagOk

`func (o *UpdatePKICertIssuer) GetServerFlagOk() (*bool, bool)`

GetServerFlagOk returns a tuple with the ServerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFlag

`func (o *UpdatePKICertIssuer) SetServerFlag(v bool)`

SetServerFlag sets ServerFlag field to given value.

### HasServerFlag

`func (o *UpdatePKICertIssuer) HasServerFlag() bool`

HasServerFlag returns a boolean if a field has been set.

### GetSignerKeyName

`func (o *UpdatePKICertIssuer) GetSignerKeyName() string`

GetSignerKeyName returns the SignerKeyName field if non-nil, zero value otherwise.

### GetSignerKeyNameOk

`func (o *UpdatePKICertIssuer) GetSignerKeyNameOk() (*string, bool)`

GetSignerKeyNameOk returns a tuple with the SignerKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerKeyName

`func (o *UpdatePKICertIssuer) SetSignerKeyName(v string)`

SetSignerKeyName sets SignerKeyName field to given value.

### HasSignerKeyName

`func (o *UpdatePKICertIssuer) HasSignerKeyName() bool`

HasSignerKeyName returns a boolean if a field has been set.

### GetStreetAddress

`func (o *UpdatePKICertIssuer) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *UpdatePKICertIssuer) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *UpdatePKICertIssuer) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *UpdatePKICertIssuer) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetToken

`func (o *UpdatePKICertIssuer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdatePKICertIssuer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdatePKICertIssuer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdatePKICertIssuer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *UpdatePKICertIssuer) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *UpdatePKICertIssuer) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *UpdatePKICertIssuer) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### GetUidToken

`func (o *UpdatePKICertIssuer) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdatePKICertIssuer) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdatePKICertIssuer) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdatePKICertIssuer) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


