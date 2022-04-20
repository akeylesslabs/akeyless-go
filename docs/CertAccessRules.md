# CertAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundCommonNames** | Pointer to **[]string** | A list of names. At least one must exist in the Common Name. Supports globbing. | [optional] 
**BoundDnsSans** | Pointer to **[]string** | A list of DNS names. At least one must exist in the SANs. Supports globbing. | [optional] 
**BoundEmailSans** | Pointer to **[]string** | A list of Email Addresses. At least one must exist in the SANs. Supports globbing. | [optional] 
**BoundExtensions** | Pointer to **[]string** | A list of extensions formatted as \&quot;oid:value\&quot;. Expects the extension value to be some type of ASN1 encoded string. All values must match. Supports globbing on \&quot;value\&quot;. | [optional] 
**BoundOrganizationalUnits** | Pointer to **[]string** | A list of Organizational Units names. At least one must exist in the OU field. | [optional] 
**BoundUriSans** | Pointer to **[]string** | A list of URIs. At least one must exist in the SANs. Supports globbing. | [optional] 
**Certificate** | Pointer to **[]int32** | Base64 encdoed PEM certificate | [optional] 
**RevokedCertIds** | Pointer to **[]string** | A list of revoked cert ids | [optional] 
**UniqueIdentifier** | Pointer to **string** | A unique identifier to distinguish different users | [optional] 

## Methods

### NewCertAccessRules

`func NewCertAccessRules() *CertAccessRules`

NewCertAccessRules instantiates a new CertAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertAccessRulesWithDefaults

`func NewCertAccessRulesWithDefaults() *CertAccessRules`

NewCertAccessRulesWithDefaults instantiates a new CertAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundCommonNames

`func (o *CertAccessRules) GetBoundCommonNames() []string`

GetBoundCommonNames returns the BoundCommonNames field if non-nil, zero value otherwise.

### GetBoundCommonNamesOk

`func (o *CertAccessRules) GetBoundCommonNamesOk() (*[]string, bool)`

GetBoundCommonNamesOk returns a tuple with the BoundCommonNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCommonNames

`func (o *CertAccessRules) SetBoundCommonNames(v []string)`

SetBoundCommonNames sets BoundCommonNames field to given value.

### HasBoundCommonNames

`func (o *CertAccessRules) HasBoundCommonNames() bool`

HasBoundCommonNames returns a boolean if a field has been set.

### GetBoundDnsSans

`func (o *CertAccessRules) GetBoundDnsSans() []string`

GetBoundDnsSans returns the BoundDnsSans field if non-nil, zero value otherwise.

### GetBoundDnsSansOk

`func (o *CertAccessRules) GetBoundDnsSansOk() (*[]string, bool)`

GetBoundDnsSansOk returns a tuple with the BoundDnsSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDnsSans

`func (o *CertAccessRules) SetBoundDnsSans(v []string)`

SetBoundDnsSans sets BoundDnsSans field to given value.

### HasBoundDnsSans

`func (o *CertAccessRules) HasBoundDnsSans() bool`

HasBoundDnsSans returns a boolean if a field has been set.

### GetBoundEmailSans

`func (o *CertAccessRules) GetBoundEmailSans() []string`

GetBoundEmailSans returns the BoundEmailSans field if non-nil, zero value otherwise.

### GetBoundEmailSansOk

`func (o *CertAccessRules) GetBoundEmailSansOk() (*[]string, bool)`

GetBoundEmailSansOk returns a tuple with the BoundEmailSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEmailSans

`func (o *CertAccessRules) SetBoundEmailSans(v []string)`

SetBoundEmailSans sets BoundEmailSans field to given value.

### HasBoundEmailSans

`func (o *CertAccessRules) HasBoundEmailSans() bool`

HasBoundEmailSans returns a boolean if a field has been set.

### GetBoundExtensions

`func (o *CertAccessRules) GetBoundExtensions() []string`

GetBoundExtensions returns the BoundExtensions field if non-nil, zero value otherwise.

### GetBoundExtensionsOk

`func (o *CertAccessRules) GetBoundExtensionsOk() (*[]string, bool)`

GetBoundExtensionsOk returns a tuple with the BoundExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundExtensions

`func (o *CertAccessRules) SetBoundExtensions(v []string)`

SetBoundExtensions sets BoundExtensions field to given value.

### HasBoundExtensions

`func (o *CertAccessRules) HasBoundExtensions() bool`

HasBoundExtensions returns a boolean if a field has been set.

### GetBoundOrganizationalUnits

`func (o *CertAccessRules) GetBoundOrganizationalUnits() []string`

GetBoundOrganizationalUnits returns the BoundOrganizationalUnits field if non-nil, zero value otherwise.

### GetBoundOrganizationalUnitsOk

`func (o *CertAccessRules) GetBoundOrganizationalUnitsOk() (*[]string, bool)`

GetBoundOrganizationalUnitsOk returns a tuple with the BoundOrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundOrganizationalUnits

`func (o *CertAccessRules) SetBoundOrganizationalUnits(v []string)`

SetBoundOrganizationalUnits sets BoundOrganizationalUnits field to given value.

### HasBoundOrganizationalUnits

`func (o *CertAccessRules) HasBoundOrganizationalUnits() bool`

HasBoundOrganizationalUnits returns a boolean if a field has been set.

### GetBoundUriSans

`func (o *CertAccessRules) GetBoundUriSans() []string`

GetBoundUriSans returns the BoundUriSans field if non-nil, zero value otherwise.

### GetBoundUriSansOk

`func (o *CertAccessRules) GetBoundUriSansOk() (*[]string, bool)`

GetBoundUriSansOk returns a tuple with the BoundUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUriSans

`func (o *CertAccessRules) SetBoundUriSans(v []string)`

SetBoundUriSans sets BoundUriSans field to given value.

### HasBoundUriSans

`func (o *CertAccessRules) HasBoundUriSans() bool`

HasBoundUriSans returns a boolean if a field has been set.

### GetCertificate

`func (o *CertAccessRules) GetCertificate() []int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertAccessRules) GetCertificateOk() (*[]int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertAccessRules) SetCertificate(v []int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertAccessRules) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetRevokedCertIds

`func (o *CertAccessRules) GetRevokedCertIds() []string`

GetRevokedCertIds returns the RevokedCertIds field if non-nil, zero value otherwise.

### GetRevokedCertIdsOk

`func (o *CertAccessRules) GetRevokedCertIdsOk() (*[]string, bool)`

GetRevokedCertIdsOk returns a tuple with the RevokedCertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCertIds

`func (o *CertAccessRules) SetRevokedCertIds(v []string)`

SetRevokedCertIds sets RevokedCertIds field to given value.

### HasRevokedCertIds

`func (o *CertAccessRules) HasRevokedCertIds() bool`

HasRevokedCertIds returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *CertAccessRules) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *CertAccessRules) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *CertAccessRules) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *CertAccessRules) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


