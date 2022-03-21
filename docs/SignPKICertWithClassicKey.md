# SignPKICertWithClassicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** | The common name to be included in the PKI certificate | [optional] 
**Country** | Pointer to **string** | A comma-separated list of the country that will be set in the issued certificate | [optional] 
**DisplayId** | **string** | The name of the key to use in the sign PKI Cert process | 
**DnsNames** | Pointer to **string** | DNS Names to be included in the PKI certificate (in a comma-delimited list) | [optional] 
**KeyUsage** | Pointer to **string** | key-usage | [optional] [default to "DigitalSignature,KeyAgreement,KeyEncipherment"]
**Locality** | Pointer to **string** | A comma-separated list of the locality that will be set in the issued certificate | [optional] 
**OrganizationalUnits** | Pointer to **string** | A comma-separated list of organizational units (OU) that will be set in the issued certificate | [optional] 
**Organizations** | Pointer to **string** | A comma-separated list of organizations (O) that will be set in the issued certificate | [optional] 
**PostalCode** | Pointer to **string** | A comma-separated list of the postal code that will be set in the issued certificate | [optional] 
**Province** | Pointer to **string** | A comma-separated list of the province that will be set in the issued certificate | [optional] 
**PublicKeyPemData** | Pointer to **string** | PublicKey using for signing in a PEM format. | [optional] 
**SigningMethod** | **string** | SigningMethod | 
**StreetAddress** | Pointer to **string** | A comma-separated list of the street address that will be set in the issued certificate | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | **int64** | he requested Time To Live for the certificate, in seconds | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UriSans** | Pointer to **string** | The URI Subject Alternative Names to be included in the PKI certificate (in a comma-delimited list) | [optional] 
**Version** | **int32** | classic key version | 

## Methods

### NewSignPKICertWithClassicKey

`func NewSignPKICertWithClassicKey(displayId string, signingMethod string, ttl int64, version int32, ) *SignPKICertWithClassicKey`

NewSignPKICertWithClassicKey instantiates a new SignPKICertWithClassicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignPKICertWithClassicKeyWithDefaults

`func NewSignPKICertWithClassicKeyWithDefaults() *SignPKICertWithClassicKey`

NewSignPKICertWithClassicKeyWithDefaults instantiates a new SignPKICertWithClassicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *SignPKICertWithClassicKey) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *SignPKICertWithClassicKey) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *SignPKICertWithClassicKey) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *SignPKICertWithClassicKey) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountry

`func (o *SignPKICertWithClassicKey) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SignPKICertWithClassicKey) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SignPKICertWithClassicKey) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SignPKICertWithClassicKey) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayId

`func (o *SignPKICertWithClassicKey) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *SignPKICertWithClassicKey) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *SignPKICertWithClassicKey) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetDnsNames

`func (o *SignPKICertWithClassicKey) GetDnsNames() string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *SignPKICertWithClassicKey) GetDnsNamesOk() (*string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *SignPKICertWithClassicKey) SetDnsNames(v string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *SignPKICertWithClassicKey) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### GetKeyUsage

`func (o *SignPKICertWithClassicKey) GetKeyUsage() string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *SignPKICertWithClassicKey) GetKeyUsageOk() (*string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *SignPKICertWithClassicKey) SetKeyUsage(v string)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *SignPKICertWithClassicKey) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetLocality

`func (o *SignPKICertWithClassicKey) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *SignPKICertWithClassicKey) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *SignPKICertWithClassicKey) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *SignPKICertWithClassicKey) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetOrganizationalUnits

`func (o *SignPKICertWithClassicKey) GetOrganizationalUnits() string`

GetOrganizationalUnits returns the OrganizationalUnits field if non-nil, zero value otherwise.

### GetOrganizationalUnitsOk

`func (o *SignPKICertWithClassicKey) GetOrganizationalUnitsOk() (*string, bool)`

GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnits

`func (o *SignPKICertWithClassicKey) SetOrganizationalUnits(v string)`

SetOrganizationalUnits sets OrganizationalUnits field to given value.

### HasOrganizationalUnits

`func (o *SignPKICertWithClassicKey) HasOrganizationalUnits() bool`

HasOrganizationalUnits returns a boolean if a field has been set.

### GetOrganizations

`func (o *SignPKICertWithClassicKey) GetOrganizations() string`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *SignPKICertWithClassicKey) GetOrganizationsOk() (*string, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *SignPKICertWithClassicKey) SetOrganizations(v string)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *SignPKICertWithClassicKey) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetPostalCode

`func (o *SignPKICertWithClassicKey) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *SignPKICertWithClassicKey) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *SignPKICertWithClassicKey) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *SignPKICertWithClassicKey) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetProvince

`func (o *SignPKICertWithClassicKey) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *SignPKICertWithClassicKey) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *SignPKICertWithClassicKey) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *SignPKICertWithClassicKey) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetPublicKeyPemData

`func (o *SignPKICertWithClassicKey) GetPublicKeyPemData() string`

GetPublicKeyPemData returns the PublicKeyPemData field if non-nil, zero value otherwise.

### GetPublicKeyPemDataOk

`func (o *SignPKICertWithClassicKey) GetPublicKeyPemDataOk() (*string, bool)`

GetPublicKeyPemDataOk returns a tuple with the PublicKeyPemData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyPemData

`func (o *SignPKICertWithClassicKey) SetPublicKeyPemData(v string)`

SetPublicKeyPemData sets PublicKeyPemData field to given value.

### HasPublicKeyPemData

`func (o *SignPKICertWithClassicKey) HasPublicKeyPemData() bool`

HasPublicKeyPemData returns a boolean if a field has been set.

### GetSigningMethod

`func (o *SignPKICertWithClassicKey) GetSigningMethod() string`

GetSigningMethod returns the SigningMethod field if non-nil, zero value otherwise.

### GetSigningMethodOk

`func (o *SignPKICertWithClassicKey) GetSigningMethodOk() (*string, bool)`

GetSigningMethodOk returns a tuple with the SigningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningMethod

`func (o *SignPKICertWithClassicKey) SetSigningMethod(v string)`

SetSigningMethod sets SigningMethod field to given value.


### GetStreetAddress

`func (o *SignPKICertWithClassicKey) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *SignPKICertWithClassicKey) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *SignPKICertWithClassicKey) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *SignPKICertWithClassicKey) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetToken

`func (o *SignPKICertWithClassicKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SignPKICertWithClassicKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SignPKICertWithClassicKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SignPKICertWithClassicKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *SignPKICertWithClassicKey) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SignPKICertWithClassicKey) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SignPKICertWithClassicKey) SetTtl(v int64)`

SetTtl sets Ttl field to given value.


### GetUidToken

`func (o *SignPKICertWithClassicKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *SignPKICertWithClassicKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *SignPKICertWithClassicKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *SignPKICertWithClassicKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUriSans

`func (o *SignPKICertWithClassicKey) GetUriSans() string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *SignPKICertWithClassicKey) GetUriSansOk() (*string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *SignPKICertWithClassicKey) SetUriSans(v string)`

SetUriSans sets UriSans field to given value.

### HasUriSans

`func (o *SignPKICertWithClassicKey) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.

### GetVersion

`func (o *SignPKICertWithClassicKey) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SignPKICertWithClassicKey) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SignPKICertWithClassicKey) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


