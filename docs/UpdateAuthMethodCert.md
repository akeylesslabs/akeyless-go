# UpdateAuthMethodCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**BoundCommonNames** | Pointer to **[]string** | A list of names. At least one must exist in the Common Name. Supports globbing. | [optional] 
**BoundDnsSans** | Pointer to **[]string** | A list of DNS names. At least one must exist in the SANs. Supports globbing. | [optional] 
**BoundEmailSans** | Pointer to **[]string** | A list of Email Addresses. At least one must exist in the SANs. Supports globbing. | [optional] 
**BoundExtensions** | Pointer to **[]string** | A list of extensions formatted as \&quot;oid:value\&quot;. Expects the extension value to be some type of ASN1 encoded string. All values much match. Supports globbing on \&quot;value\&quot;. | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**BoundOrganizationalUnits** | Pointer to **[]string** | A list of Organizational Units names. At least one must exist in the OU field. | [optional] 
**BoundUriSans** | Pointer to **[]string** | A list of URIs. At least one must exist in the SANs. Supports globbing. | [optional] 
**CertificateData** | Pointer to **string** | The certificate data in base64, if no file was provided | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**RevokedCertIds** | Pointer to **[]string** | A list of revoked cert ids | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | **string** | A unique identifier (ID) value should be configured, such as common_name or organizational_unit Whenever a user logs in with a token, these authentication types issue a \&quot;sub claim\&quot; that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization. | 

## Methods

### NewUpdateAuthMethodCert

`func NewUpdateAuthMethodCert(name string, uniqueIdentifier string, ) *UpdateAuthMethodCert`

NewUpdateAuthMethodCert instantiates a new UpdateAuthMethodCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthMethodCertWithDefaults

`func NewUpdateAuthMethodCertWithDefaults() *UpdateAuthMethodCert`

NewUpdateAuthMethodCertWithDefaults instantiates a new UpdateAuthMethodCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *UpdateAuthMethodCert) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *UpdateAuthMethodCert) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *UpdateAuthMethodCert) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *UpdateAuthMethodCert) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundCommonNames

`func (o *UpdateAuthMethodCert) GetBoundCommonNames() []string`

GetBoundCommonNames returns the BoundCommonNames field if non-nil, zero value otherwise.

### GetBoundCommonNamesOk

`func (o *UpdateAuthMethodCert) GetBoundCommonNamesOk() (*[]string, bool)`

GetBoundCommonNamesOk returns a tuple with the BoundCommonNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCommonNames

`func (o *UpdateAuthMethodCert) SetBoundCommonNames(v []string)`

SetBoundCommonNames sets BoundCommonNames field to given value.

### HasBoundCommonNames

`func (o *UpdateAuthMethodCert) HasBoundCommonNames() bool`

HasBoundCommonNames returns a boolean if a field has been set.

### GetBoundDnsSans

`func (o *UpdateAuthMethodCert) GetBoundDnsSans() []string`

GetBoundDnsSans returns the BoundDnsSans field if non-nil, zero value otherwise.

### GetBoundDnsSansOk

`func (o *UpdateAuthMethodCert) GetBoundDnsSansOk() (*[]string, bool)`

GetBoundDnsSansOk returns a tuple with the BoundDnsSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDnsSans

`func (o *UpdateAuthMethodCert) SetBoundDnsSans(v []string)`

SetBoundDnsSans sets BoundDnsSans field to given value.

### HasBoundDnsSans

`func (o *UpdateAuthMethodCert) HasBoundDnsSans() bool`

HasBoundDnsSans returns a boolean if a field has been set.

### GetBoundEmailSans

`func (o *UpdateAuthMethodCert) GetBoundEmailSans() []string`

GetBoundEmailSans returns the BoundEmailSans field if non-nil, zero value otherwise.

### GetBoundEmailSansOk

`func (o *UpdateAuthMethodCert) GetBoundEmailSansOk() (*[]string, bool)`

GetBoundEmailSansOk returns a tuple with the BoundEmailSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEmailSans

`func (o *UpdateAuthMethodCert) SetBoundEmailSans(v []string)`

SetBoundEmailSans sets BoundEmailSans field to given value.

### HasBoundEmailSans

`func (o *UpdateAuthMethodCert) HasBoundEmailSans() bool`

HasBoundEmailSans returns a boolean if a field has been set.

### GetBoundExtensions

`func (o *UpdateAuthMethodCert) GetBoundExtensions() []string`

GetBoundExtensions returns the BoundExtensions field if non-nil, zero value otherwise.

### GetBoundExtensionsOk

`func (o *UpdateAuthMethodCert) GetBoundExtensionsOk() (*[]string, bool)`

GetBoundExtensionsOk returns a tuple with the BoundExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundExtensions

`func (o *UpdateAuthMethodCert) SetBoundExtensions(v []string)`

SetBoundExtensions sets BoundExtensions field to given value.

### HasBoundExtensions

`func (o *UpdateAuthMethodCert) HasBoundExtensions() bool`

HasBoundExtensions returns a boolean if a field has been set.

### GetBoundIps

`func (o *UpdateAuthMethodCert) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAuthMethodCert) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAuthMethodCert) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAuthMethodCert) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundOrganizationalUnits

`func (o *UpdateAuthMethodCert) GetBoundOrganizationalUnits() []string`

GetBoundOrganizationalUnits returns the BoundOrganizationalUnits field if non-nil, zero value otherwise.

### GetBoundOrganizationalUnitsOk

`func (o *UpdateAuthMethodCert) GetBoundOrganizationalUnitsOk() (*[]string, bool)`

GetBoundOrganizationalUnitsOk returns a tuple with the BoundOrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundOrganizationalUnits

`func (o *UpdateAuthMethodCert) SetBoundOrganizationalUnits(v []string)`

SetBoundOrganizationalUnits sets BoundOrganizationalUnits field to given value.

### HasBoundOrganizationalUnits

`func (o *UpdateAuthMethodCert) HasBoundOrganizationalUnits() bool`

HasBoundOrganizationalUnits returns a boolean if a field has been set.

### GetBoundUriSans

`func (o *UpdateAuthMethodCert) GetBoundUriSans() []string`

GetBoundUriSans returns the BoundUriSans field if non-nil, zero value otherwise.

### GetBoundUriSansOk

`func (o *UpdateAuthMethodCert) GetBoundUriSansOk() (*[]string, bool)`

GetBoundUriSansOk returns a tuple with the BoundUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUriSans

`func (o *UpdateAuthMethodCert) SetBoundUriSans(v []string)`

SetBoundUriSans sets BoundUriSans field to given value.

### HasBoundUriSans

`func (o *UpdateAuthMethodCert) HasBoundUriSans() bool`

HasBoundUriSans returns a boolean if a field has been set.

### GetCertificateData

`func (o *UpdateAuthMethodCert) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *UpdateAuthMethodCert) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *UpdateAuthMethodCert) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.

### HasCertificateData

`func (o *UpdateAuthMethodCert) HasCertificateData() bool`

HasCertificateData returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *UpdateAuthMethodCert) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *UpdateAuthMethodCert) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *UpdateAuthMethodCert) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *UpdateAuthMethodCert) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *UpdateAuthMethodCert) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *UpdateAuthMethodCert) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *UpdateAuthMethodCert) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *UpdateAuthMethodCert) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *UpdateAuthMethodCert) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateAuthMethodCert) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateAuthMethodCert) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateAuthMethodCert) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *UpdateAuthMethodCert) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *UpdateAuthMethodCert) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *UpdateAuthMethodCert) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *UpdateAuthMethodCert) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *UpdateAuthMethodCert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthMethodCert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthMethodCert) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateAuthMethodCert) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateAuthMethodCert) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateAuthMethodCert) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateAuthMethodCert) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRevokedCertIds

`func (o *UpdateAuthMethodCert) GetRevokedCertIds() []string`

GetRevokedCertIds returns the RevokedCertIds field if non-nil, zero value otherwise.

### GetRevokedCertIdsOk

`func (o *UpdateAuthMethodCert) GetRevokedCertIdsOk() (*[]string, bool)`

GetRevokedCertIdsOk returns a tuple with the RevokedCertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCertIds

`func (o *UpdateAuthMethodCert) SetRevokedCertIds(v []string)`

SetRevokedCertIds sets RevokedCertIds field to given value.

### HasRevokedCertIds

`func (o *UpdateAuthMethodCert) HasRevokedCertIds() bool`

HasRevokedCertIds returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAuthMethodCert) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAuthMethodCert) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAuthMethodCert) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAuthMethodCert) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAuthMethodCert) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAuthMethodCert) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAuthMethodCert) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAuthMethodCert) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *UpdateAuthMethodCert) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *UpdateAuthMethodCert) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *UpdateAuthMethodCert) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


