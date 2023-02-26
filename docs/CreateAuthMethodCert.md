# CreateAuthMethodCert

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
**RevokedCertIds** | Pointer to **[]string** | A list of revoked cert ids | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | **string** | A unique identifier (ID) value should be configured, such as common_name or organizational_unit Whenever a user logs in with a token, these authentication types issue a \&quot;sub claim\&quot; that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization. | 

## Methods

### NewCreateAuthMethodCert

`func NewCreateAuthMethodCert(name string, uniqueIdentifier string, ) *CreateAuthMethodCert`

NewCreateAuthMethodCert instantiates a new CreateAuthMethodCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodCertWithDefaults

`func NewCreateAuthMethodCertWithDefaults() *CreateAuthMethodCert`

NewCreateAuthMethodCertWithDefaults instantiates a new CreateAuthMethodCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodCert) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodCert) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodCert) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodCert) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundCommonNames

`func (o *CreateAuthMethodCert) GetBoundCommonNames() []string`

GetBoundCommonNames returns the BoundCommonNames field if non-nil, zero value otherwise.

### GetBoundCommonNamesOk

`func (o *CreateAuthMethodCert) GetBoundCommonNamesOk() (*[]string, bool)`

GetBoundCommonNamesOk returns a tuple with the BoundCommonNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCommonNames

`func (o *CreateAuthMethodCert) SetBoundCommonNames(v []string)`

SetBoundCommonNames sets BoundCommonNames field to given value.

### HasBoundCommonNames

`func (o *CreateAuthMethodCert) HasBoundCommonNames() bool`

HasBoundCommonNames returns a boolean if a field has been set.

### GetBoundDnsSans

`func (o *CreateAuthMethodCert) GetBoundDnsSans() []string`

GetBoundDnsSans returns the BoundDnsSans field if non-nil, zero value otherwise.

### GetBoundDnsSansOk

`func (o *CreateAuthMethodCert) GetBoundDnsSansOk() (*[]string, bool)`

GetBoundDnsSansOk returns a tuple with the BoundDnsSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDnsSans

`func (o *CreateAuthMethodCert) SetBoundDnsSans(v []string)`

SetBoundDnsSans sets BoundDnsSans field to given value.

### HasBoundDnsSans

`func (o *CreateAuthMethodCert) HasBoundDnsSans() bool`

HasBoundDnsSans returns a boolean if a field has been set.

### GetBoundEmailSans

`func (o *CreateAuthMethodCert) GetBoundEmailSans() []string`

GetBoundEmailSans returns the BoundEmailSans field if non-nil, zero value otherwise.

### GetBoundEmailSansOk

`func (o *CreateAuthMethodCert) GetBoundEmailSansOk() (*[]string, bool)`

GetBoundEmailSansOk returns a tuple with the BoundEmailSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEmailSans

`func (o *CreateAuthMethodCert) SetBoundEmailSans(v []string)`

SetBoundEmailSans sets BoundEmailSans field to given value.

### HasBoundEmailSans

`func (o *CreateAuthMethodCert) HasBoundEmailSans() bool`

HasBoundEmailSans returns a boolean if a field has been set.

### GetBoundExtensions

`func (o *CreateAuthMethodCert) GetBoundExtensions() []string`

GetBoundExtensions returns the BoundExtensions field if non-nil, zero value otherwise.

### GetBoundExtensionsOk

`func (o *CreateAuthMethodCert) GetBoundExtensionsOk() (*[]string, bool)`

GetBoundExtensionsOk returns a tuple with the BoundExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundExtensions

`func (o *CreateAuthMethodCert) SetBoundExtensions(v []string)`

SetBoundExtensions sets BoundExtensions field to given value.

### HasBoundExtensions

`func (o *CreateAuthMethodCert) HasBoundExtensions() bool`

HasBoundExtensions returns a boolean if a field has been set.

### GetBoundIps

`func (o *CreateAuthMethodCert) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodCert) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodCert) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodCert) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundOrganizationalUnits

`func (o *CreateAuthMethodCert) GetBoundOrganizationalUnits() []string`

GetBoundOrganizationalUnits returns the BoundOrganizationalUnits field if non-nil, zero value otherwise.

### GetBoundOrganizationalUnitsOk

`func (o *CreateAuthMethodCert) GetBoundOrganizationalUnitsOk() (*[]string, bool)`

GetBoundOrganizationalUnitsOk returns a tuple with the BoundOrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundOrganizationalUnits

`func (o *CreateAuthMethodCert) SetBoundOrganizationalUnits(v []string)`

SetBoundOrganizationalUnits sets BoundOrganizationalUnits field to given value.

### HasBoundOrganizationalUnits

`func (o *CreateAuthMethodCert) HasBoundOrganizationalUnits() bool`

HasBoundOrganizationalUnits returns a boolean if a field has been set.

### GetBoundUriSans

`func (o *CreateAuthMethodCert) GetBoundUriSans() []string`

GetBoundUriSans returns the BoundUriSans field if non-nil, zero value otherwise.

### GetBoundUriSansOk

`func (o *CreateAuthMethodCert) GetBoundUriSansOk() (*[]string, bool)`

GetBoundUriSansOk returns a tuple with the BoundUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUriSans

`func (o *CreateAuthMethodCert) SetBoundUriSans(v []string)`

SetBoundUriSans sets BoundUriSans field to given value.

### HasBoundUriSans

`func (o *CreateAuthMethodCert) HasBoundUriSans() bool`

HasBoundUriSans returns a boolean if a field has been set.

### GetCertificateData

`func (o *CreateAuthMethodCert) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *CreateAuthMethodCert) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *CreateAuthMethodCert) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.

### HasCertificateData

`func (o *CreateAuthMethodCert) HasCertificateData() bool`

HasCertificateData returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *CreateAuthMethodCert) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodCert) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodCert) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodCert) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *CreateAuthMethodCert) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *CreateAuthMethodCert) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *CreateAuthMethodCert) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *CreateAuthMethodCert) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *CreateAuthMethodCert) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateAuthMethodCert) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateAuthMethodCert) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateAuthMethodCert) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *CreateAuthMethodCert) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *CreateAuthMethodCert) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *CreateAuthMethodCert) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *CreateAuthMethodCert) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodCert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodCert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodCert) SetName(v string)`

SetName sets Name field to given value.


### GetRevokedCertIds

`func (o *CreateAuthMethodCert) GetRevokedCertIds() []string`

GetRevokedCertIds returns the RevokedCertIds field if non-nil, zero value otherwise.

### GetRevokedCertIdsOk

`func (o *CreateAuthMethodCert) GetRevokedCertIdsOk() (*[]string, bool)`

GetRevokedCertIdsOk returns a tuple with the RevokedCertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCertIds

`func (o *CreateAuthMethodCert) SetRevokedCertIds(v []string)`

SetRevokedCertIds sets RevokedCertIds field to given value.

### HasRevokedCertIds

`func (o *CreateAuthMethodCert) HasRevokedCertIds() bool`

HasRevokedCertIds returns a boolean if a field has been set.

### GetToken

`func (o *CreateAuthMethodCert) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodCert) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodCert) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodCert) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodCert) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodCert) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodCert) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodCert) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *CreateAuthMethodCert) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *CreateAuthMethodCert) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *CreateAuthMethodCert) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


