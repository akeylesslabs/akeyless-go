# CreatePKICertIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAnyName** | Pointer to **bool** | If set, clients can request certificates for any CN | [optional] 
**AllowSubdomains** | Pointer to **bool** | If set, clients can request certificates for subdomains and wildcard subdomains of the allowed domains | [optional] 
**AllowedDomains** | Pointer to **string** | A list of the allowed domains that clients can request to be included in the certificate (in a comma-delimited list) | [optional] 
**AllowedUriSans** | Pointer to **string** | A list of the allowed URIs that clients can request to be included in the certificate as part of the URI Subject Alternative Names (in a comma-delimited list) | [optional] 
**ClientFlag** | Pointer to **bool** | If set, certificates will be flagged for client auth use | [optional] 
**CodeSigningFlag** | Pointer to **bool** | If set, certificates will be flagged for code signing use | [optional] 
**Country** | Pointer to **string** | A comma-separated list of the country that will be set in the issued certificate | [optional] 
**KeyUsage** | Pointer to **string** | key-usage | [optional] [default to "DigitalSignature,KeyAgreement,KeyEncipherment"]
**Locality** | Pointer to **string** | A comma-separated list of the locality that will be set in the issued certificate | [optional] 
**Metadata** | Pointer to **string** | A metadata about the issuer | [optional] 
**Name** | **string** | PKI certificate issuer name | 
**NotEnforceHostnames** | Pointer to **bool** | If set, any names are allowed for CN and SANs in the certificate and not only a valid host name | [optional] 
**NotRequireCn** | Pointer to **bool** | If set, clients can request certificates without a CN | [optional] 
**OrganizationalUnits** | Pointer to **string** | A comma-separated list of organizational units (OU) that will be set in the issued certificate | [optional] 
**Organizations** | Pointer to **string** | A comma-separated list of organizations (O) that will be set in the issued certificate | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**PostalCode** | Pointer to **string** | A comma-separated list of the postal code that will be set in the issued certificate | [optional] 
**Province** | Pointer to **string** | A comma-separated list of the province that will be set in the issued certificate | [optional] 
**ServerFlag** | Pointer to **bool** | If set, certificates will be flagged for server auth use | [optional] 
**SignerKeyName** | **string** | A key to sign the certificate with | 
**StreetAddress** | Pointer to **string** | A comma-separated list of the street address that will be set in the issued certificate | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | **int64** | The requested Time To Live for the certificate, use second units | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreatePKICertIssuer

`func NewCreatePKICertIssuer(name string, signerKeyName string, ttl int64, ) *CreatePKICertIssuer`

NewCreatePKICertIssuer instantiates a new CreatePKICertIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePKICertIssuerWithDefaults

`func NewCreatePKICertIssuerWithDefaults() *CreatePKICertIssuer`

NewCreatePKICertIssuerWithDefaults instantiates a new CreatePKICertIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAnyName

`func (o *CreatePKICertIssuer) GetAllowAnyName() bool`

GetAllowAnyName returns the AllowAnyName field if non-nil, zero value otherwise.

### GetAllowAnyNameOk

`func (o *CreatePKICertIssuer) GetAllowAnyNameOk() (*bool, bool)`

GetAllowAnyNameOk returns a tuple with the AllowAnyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnyName

`func (o *CreatePKICertIssuer) SetAllowAnyName(v bool)`

SetAllowAnyName sets AllowAnyName field to given value.

### HasAllowAnyName

`func (o *CreatePKICertIssuer) HasAllowAnyName() bool`

HasAllowAnyName returns a boolean if a field has been set.

### GetAllowSubdomains

`func (o *CreatePKICertIssuer) GetAllowSubdomains() bool`

GetAllowSubdomains returns the AllowSubdomains field if non-nil, zero value otherwise.

### GetAllowSubdomainsOk

`func (o *CreatePKICertIssuer) GetAllowSubdomainsOk() (*bool, bool)`

GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubdomains

`func (o *CreatePKICertIssuer) SetAllowSubdomains(v bool)`

SetAllowSubdomains sets AllowSubdomains field to given value.

### HasAllowSubdomains

`func (o *CreatePKICertIssuer) HasAllowSubdomains() bool`

HasAllowSubdomains returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *CreatePKICertIssuer) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *CreatePKICertIssuer) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *CreatePKICertIssuer) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *CreatePKICertIssuer) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAllowedUriSans

`func (o *CreatePKICertIssuer) GetAllowedUriSans() string`

GetAllowedUriSans returns the AllowedUriSans field if non-nil, zero value otherwise.

### GetAllowedUriSansOk

`func (o *CreatePKICertIssuer) GetAllowedUriSansOk() (*string, bool)`

GetAllowedUriSansOk returns a tuple with the AllowedUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUriSans

`func (o *CreatePKICertIssuer) SetAllowedUriSans(v string)`

SetAllowedUriSans sets AllowedUriSans field to given value.

### HasAllowedUriSans

`func (o *CreatePKICertIssuer) HasAllowedUriSans() bool`

HasAllowedUriSans returns a boolean if a field has been set.

### GetClientFlag

`func (o *CreatePKICertIssuer) GetClientFlag() bool`

GetClientFlag returns the ClientFlag field if non-nil, zero value otherwise.

### GetClientFlagOk

`func (o *CreatePKICertIssuer) GetClientFlagOk() (*bool, bool)`

GetClientFlagOk returns a tuple with the ClientFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlag

`func (o *CreatePKICertIssuer) SetClientFlag(v bool)`

SetClientFlag sets ClientFlag field to given value.

### HasClientFlag

`func (o *CreatePKICertIssuer) HasClientFlag() bool`

HasClientFlag returns a boolean if a field has been set.

### GetCodeSigningFlag

`func (o *CreatePKICertIssuer) GetCodeSigningFlag() bool`

GetCodeSigningFlag returns the CodeSigningFlag field if non-nil, zero value otherwise.

### GetCodeSigningFlagOk

`func (o *CreatePKICertIssuer) GetCodeSigningFlagOk() (*bool, bool)`

GetCodeSigningFlagOk returns a tuple with the CodeSigningFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSigningFlag

`func (o *CreatePKICertIssuer) SetCodeSigningFlag(v bool)`

SetCodeSigningFlag sets CodeSigningFlag field to given value.

### HasCodeSigningFlag

`func (o *CreatePKICertIssuer) HasCodeSigningFlag() bool`

HasCodeSigningFlag returns a boolean if a field has been set.

### GetCountry

`func (o *CreatePKICertIssuer) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreatePKICertIssuer) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreatePKICertIssuer) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreatePKICertIssuer) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetKeyUsage

`func (o *CreatePKICertIssuer) GetKeyUsage() string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *CreatePKICertIssuer) GetKeyUsageOk() (*string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *CreatePKICertIssuer) SetKeyUsage(v string)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *CreatePKICertIssuer) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetLocality

`func (o *CreatePKICertIssuer) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *CreatePKICertIssuer) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *CreatePKICertIssuer) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *CreatePKICertIssuer) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetMetadata

`func (o *CreatePKICertIssuer) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePKICertIssuer) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePKICertIssuer) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePKICertIssuer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreatePKICertIssuer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePKICertIssuer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePKICertIssuer) SetName(v string)`

SetName sets Name field to given value.


### GetNotEnforceHostnames

`func (o *CreatePKICertIssuer) GetNotEnforceHostnames() bool`

GetNotEnforceHostnames returns the NotEnforceHostnames field if non-nil, zero value otherwise.

### GetNotEnforceHostnamesOk

`func (o *CreatePKICertIssuer) GetNotEnforceHostnamesOk() (*bool, bool)`

GetNotEnforceHostnamesOk returns a tuple with the NotEnforceHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEnforceHostnames

`func (o *CreatePKICertIssuer) SetNotEnforceHostnames(v bool)`

SetNotEnforceHostnames sets NotEnforceHostnames field to given value.

### HasNotEnforceHostnames

`func (o *CreatePKICertIssuer) HasNotEnforceHostnames() bool`

HasNotEnforceHostnames returns a boolean if a field has been set.

### GetNotRequireCn

`func (o *CreatePKICertIssuer) GetNotRequireCn() bool`

GetNotRequireCn returns the NotRequireCn field if non-nil, zero value otherwise.

### GetNotRequireCnOk

`func (o *CreatePKICertIssuer) GetNotRequireCnOk() (*bool, bool)`

GetNotRequireCnOk returns a tuple with the NotRequireCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRequireCn

`func (o *CreatePKICertIssuer) SetNotRequireCn(v bool)`

SetNotRequireCn sets NotRequireCn field to given value.

### HasNotRequireCn

`func (o *CreatePKICertIssuer) HasNotRequireCn() bool`

HasNotRequireCn returns a boolean if a field has been set.

### GetOrganizationalUnits

`func (o *CreatePKICertIssuer) GetOrganizationalUnits() string`

GetOrganizationalUnits returns the OrganizationalUnits field if non-nil, zero value otherwise.

### GetOrganizationalUnitsOk

`func (o *CreatePKICertIssuer) GetOrganizationalUnitsOk() (*string, bool)`

GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnits

`func (o *CreatePKICertIssuer) SetOrganizationalUnits(v string)`

SetOrganizationalUnits sets OrganizationalUnits field to given value.

### HasOrganizationalUnits

`func (o *CreatePKICertIssuer) HasOrganizationalUnits() bool`

HasOrganizationalUnits returns a boolean if a field has been set.

### GetOrganizations

`func (o *CreatePKICertIssuer) GetOrganizations() string`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *CreatePKICertIssuer) GetOrganizationsOk() (*string, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *CreatePKICertIssuer) SetOrganizations(v string)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *CreatePKICertIssuer) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetPassword

`func (o *CreatePKICertIssuer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreatePKICertIssuer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreatePKICertIssuer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreatePKICertIssuer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPostalCode

`func (o *CreatePKICertIssuer) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CreatePKICertIssuer) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CreatePKICertIssuer) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CreatePKICertIssuer) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetProvince

`func (o *CreatePKICertIssuer) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *CreatePKICertIssuer) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *CreatePKICertIssuer) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *CreatePKICertIssuer) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetServerFlag

`func (o *CreatePKICertIssuer) GetServerFlag() bool`

GetServerFlag returns the ServerFlag field if non-nil, zero value otherwise.

### GetServerFlagOk

`func (o *CreatePKICertIssuer) GetServerFlagOk() (*bool, bool)`

GetServerFlagOk returns a tuple with the ServerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFlag

`func (o *CreatePKICertIssuer) SetServerFlag(v bool)`

SetServerFlag sets ServerFlag field to given value.

### HasServerFlag

`func (o *CreatePKICertIssuer) HasServerFlag() bool`

HasServerFlag returns a boolean if a field has been set.

### GetSignerKeyName

`func (o *CreatePKICertIssuer) GetSignerKeyName() string`

GetSignerKeyName returns the SignerKeyName field if non-nil, zero value otherwise.

### GetSignerKeyNameOk

`func (o *CreatePKICertIssuer) GetSignerKeyNameOk() (*string, bool)`

GetSignerKeyNameOk returns a tuple with the SignerKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerKeyName

`func (o *CreatePKICertIssuer) SetSignerKeyName(v string)`

SetSignerKeyName sets SignerKeyName field to given value.


### GetStreetAddress

`func (o *CreatePKICertIssuer) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *CreatePKICertIssuer) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *CreatePKICertIssuer) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *CreatePKICertIssuer) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetToken

`func (o *CreatePKICertIssuer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreatePKICertIssuer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreatePKICertIssuer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreatePKICertIssuer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *CreatePKICertIssuer) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreatePKICertIssuer) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreatePKICertIssuer) SetTtl(v int64)`

SetTtl sets Ttl field to given value.


### GetUidToken

`func (o *CreatePKICertIssuer) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreatePKICertIssuer) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreatePKICertIssuer) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreatePKICertIssuer) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreatePKICertIssuer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreatePKICertIssuer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreatePKICertIssuer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreatePKICertIssuer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


