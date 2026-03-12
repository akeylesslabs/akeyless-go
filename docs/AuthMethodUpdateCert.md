# AuthMethodUpdateCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedClientType** | Pointer to **[]string** | limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension] | [optional] 
**AllowedCors** | Pointer to **string** | Comma separated list of allowed CORS domains to be validated as part of the authentication flow. | [optional] 
**AuditLogsClaims** | Pointer to **[]string** | Subclaims to include in audit logs, e.g \&quot;--audit-logs-claims email --audit-logs-claims username\&quot; | [optional] 
**BoundCommonNames** | Pointer to **[]string** | A list of names. At least one must exist in the Common Name. Supports globbing. | [optional] 
**BoundDnsSans** | Pointer to **[]string** | A list of DNS names. At least one must exist in the SANs. Supports globbing. | [optional] 
**BoundEmailSans** | Pointer to **[]string** | A list of Email Addresses. At least one must exist in the SANs. Supports globbing. | [optional] 
**BoundExtensions** | Pointer to **[]string** | A list of extensions formatted as \&quot;oid:value\&quot;. Expects the extension value to be some type of ASN1 encoded string. All values much match. Supports globbing on \&quot;value\&quot;. | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**BoundOrganizationalUnits** | Pointer to **[]string** | A list of Organizational Units names. At least one must exist in the OU field. | [optional] 
**BoundUriSans** | Pointer to **[]string** | A list of URIs. At least one must exist in the SANs. Supports globbing. | [optional] 
**CertificateData** | Pointer to **string** | The certificate data in base64, if no file was provided | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the auth method would you like to be notified. | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca] | [optional] 
**RequireCrlDp** | Pointer to **bool** | Require certificate CRL distribution points (CDP) and enforce CRL validation during authentication. | [optional] 
**RevokedCertIds** | Pointer to **[]string** | A list of revoked cert ids | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | **string** | A unique identifier (ID) value should be configured, such as common_name or organizational_unit Whenever a user logs in with a token, these authentication types issue a \&quot;sub claim\&quot; that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization. | 

## Methods

### NewAuthMethodUpdateCert

`func NewAuthMethodUpdateCert(name string, uniqueIdentifier string, ) *AuthMethodUpdateCert`

NewAuthMethodUpdateCert instantiates a new AuthMethodUpdateCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodUpdateCertWithDefaults

`func NewAuthMethodUpdateCertWithDefaults() *AuthMethodUpdateCert`

NewAuthMethodUpdateCertWithDefaults instantiates a new AuthMethodUpdateCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *AuthMethodUpdateCert) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *AuthMethodUpdateCert) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *AuthMethodUpdateCert) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *AuthMethodUpdateCert) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedClientType

`func (o *AuthMethodUpdateCert) GetAllowedClientType() []string`

GetAllowedClientType returns the AllowedClientType field if non-nil, zero value otherwise.

### GetAllowedClientTypeOk

`func (o *AuthMethodUpdateCert) GetAllowedClientTypeOk() (*[]string, bool)`

GetAllowedClientTypeOk returns a tuple with the AllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientType

`func (o *AuthMethodUpdateCert) SetAllowedClientType(v []string)`

SetAllowedClientType sets AllowedClientType field to given value.

### HasAllowedClientType

`func (o *AuthMethodUpdateCert) HasAllowedClientType() bool`

HasAllowedClientType returns a boolean if a field has been set.

### GetAllowedCors

`func (o *AuthMethodUpdateCert) GetAllowedCors() string`

GetAllowedCors returns the AllowedCors field if non-nil, zero value otherwise.

### GetAllowedCorsOk

`func (o *AuthMethodUpdateCert) GetAllowedCorsOk() (*string, bool)`

GetAllowedCorsOk returns a tuple with the AllowedCors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCors

`func (o *AuthMethodUpdateCert) SetAllowedCors(v string)`

SetAllowedCors sets AllowedCors field to given value.

### HasAllowedCors

`func (o *AuthMethodUpdateCert) HasAllowedCors() bool`

HasAllowedCors returns a boolean if a field has been set.

### GetAuditLogsClaims

`func (o *AuthMethodUpdateCert) GetAuditLogsClaims() []string`

GetAuditLogsClaims returns the AuditLogsClaims field if non-nil, zero value otherwise.

### GetAuditLogsClaimsOk

`func (o *AuthMethodUpdateCert) GetAuditLogsClaimsOk() (*[]string, bool)`

GetAuditLogsClaimsOk returns a tuple with the AuditLogsClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsClaims

`func (o *AuthMethodUpdateCert) SetAuditLogsClaims(v []string)`

SetAuditLogsClaims sets AuditLogsClaims field to given value.

### HasAuditLogsClaims

`func (o *AuthMethodUpdateCert) HasAuditLogsClaims() bool`

HasAuditLogsClaims returns a boolean if a field has been set.

### GetBoundCommonNames

`func (o *AuthMethodUpdateCert) GetBoundCommonNames() []string`

GetBoundCommonNames returns the BoundCommonNames field if non-nil, zero value otherwise.

### GetBoundCommonNamesOk

`func (o *AuthMethodUpdateCert) GetBoundCommonNamesOk() (*[]string, bool)`

GetBoundCommonNamesOk returns a tuple with the BoundCommonNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCommonNames

`func (o *AuthMethodUpdateCert) SetBoundCommonNames(v []string)`

SetBoundCommonNames sets BoundCommonNames field to given value.

### HasBoundCommonNames

`func (o *AuthMethodUpdateCert) HasBoundCommonNames() bool`

HasBoundCommonNames returns a boolean if a field has been set.

### GetBoundDnsSans

`func (o *AuthMethodUpdateCert) GetBoundDnsSans() []string`

GetBoundDnsSans returns the BoundDnsSans field if non-nil, zero value otherwise.

### GetBoundDnsSansOk

`func (o *AuthMethodUpdateCert) GetBoundDnsSansOk() (*[]string, bool)`

GetBoundDnsSansOk returns a tuple with the BoundDnsSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDnsSans

`func (o *AuthMethodUpdateCert) SetBoundDnsSans(v []string)`

SetBoundDnsSans sets BoundDnsSans field to given value.

### HasBoundDnsSans

`func (o *AuthMethodUpdateCert) HasBoundDnsSans() bool`

HasBoundDnsSans returns a boolean if a field has been set.

### GetBoundEmailSans

`func (o *AuthMethodUpdateCert) GetBoundEmailSans() []string`

GetBoundEmailSans returns the BoundEmailSans field if non-nil, zero value otherwise.

### GetBoundEmailSansOk

`func (o *AuthMethodUpdateCert) GetBoundEmailSansOk() (*[]string, bool)`

GetBoundEmailSansOk returns a tuple with the BoundEmailSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEmailSans

`func (o *AuthMethodUpdateCert) SetBoundEmailSans(v []string)`

SetBoundEmailSans sets BoundEmailSans field to given value.

### HasBoundEmailSans

`func (o *AuthMethodUpdateCert) HasBoundEmailSans() bool`

HasBoundEmailSans returns a boolean if a field has been set.

### GetBoundExtensions

`func (o *AuthMethodUpdateCert) GetBoundExtensions() []string`

GetBoundExtensions returns the BoundExtensions field if non-nil, zero value otherwise.

### GetBoundExtensionsOk

`func (o *AuthMethodUpdateCert) GetBoundExtensionsOk() (*[]string, bool)`

GetBoundExtensionsOk returns a tuple with the BoundExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundExtensions

`func (o *AuthMethodUpdateCert) SetBoundExtensions(v []string)`

SetBoundExtensions sets BoundExtensions field to given value.

### HasBoundExtensions

`func (o *AuthMethodUpdateCert) HasBoundExtensions() bool`

HasBoundExtensions returns a boolean if a field has been set.

### GetBoundIps

`func (o *AuthMethodUpdateCert) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *AuthMethodUpdateCert) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *AuthMethodUpdateCert) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *AuthMethodUpdateCert) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundOrganizationalUnits

`func (o *AuthMethodUpdateCert) GetBoundOrganizationalUnits() []string`

GetBoundOrganizationalUnits returns the BoundOrganizationalUnits field if non-nil, zero value otherwise.

### GetBoundOrganizationalUnitsOk

`func (o *AuthMethodUpdateCert) GetBoundOrganizationalUnitsOk() (*[]string, bool)`

GetBoundOrganizationalUnitsOk returns a tuple with the BoundOrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundOrganizationalUnits

`func (o *AuthMethodUpdateCert) SetBoundOrganizationalUnits(v []string)`

SetBoundOrganizationalUnits sets BoundOrganizationalUnits field to given value.

### HasBoundOrganizationalUnits

`func (o *AuthMethodUpdateCert) HasBoundOrganizationalUnits() bool`

HasBoundOrganizationalUnits returns a boolean if a field has been set.

### GetBoundUriSans

`func (o *AuthMethodUpdateCert) GetBoundUriSans() []string`

GetBoundUriSans returns the BoundUriSans field if non-nil, zero value otherwise.

### GetBoundUriSansOk

`func (o *AuthMethodUpdateCert) GetBoundUriSansOk() (*[]string, bool)`

GetBoundUriSansOk returns a tuple with the BoundUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUriSans

`func (o *AuthMethodUpdateCert) SetBoundUriSans(v []string)`

SetBoundUriSans sets BoundUriSans field to given value.

### HasBoundUriSans

`func (o *AuthMethodUpdateCert) HasBoundUriSans() bool`

HasBoundUriSans returns a boolean if a field has been set.

### GetCertificateData

`func (o *AuthMethodUpdateCert) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *AuthMethodUpdateCert) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *AuthMethodUpdateCert) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.

### HasCertificateData

`func (o *AuthMethodUpdateCert) HasCertificateData() bool`

HasCertificateData returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *AuthMethodUpdateCert) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *AuthMethodUpdateCert) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *AuthMethodUpdateCert) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *AuthMethodUpdateCert) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *AuthMethodUpdateCert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthMethodUpdateCert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthMethodUpdateCert) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthMethodUpdateCert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *AuthMethodUpdateCert) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *AuthMethodUpdateCert) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *AuthMethodUpdateCert) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *AuthMethodUpdateCert) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *AuthMethodUpdateCert) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *AuthMethodUpdateCert) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *AuthMethodUpdateCert) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *AuthMethodUpdateCert) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *AuthMethodUpdateCert) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *AuthMethodUpdateCert) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *AuthMethodUpdateCert) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *AuthMethodUpdateCert) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *AuthMethodUpdateCert) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AuthMethodUpdateCert) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AuthMethodUpdateCert) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AuthMethodUpdateCert) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *AuthMethodUpdateCert) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *AuthMethodUpdateCert) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *AuthMethodUpdateCert) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *AuthMethodUpdateCert) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *AuthMethodUpdateCert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthMethodUpdateCert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthMethodUpdateCert) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *AuthMethodUpdateCert) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *AuthMethodUpdateCert) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *AuthMethodUpdateCert) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *AuthMethodUpdateCert) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProductType

`func (o *AuthMethodUpdateCert) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuthMethodUpdateCert) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuthMethodUpdateCert) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuthMethodUpdateCert) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetRequireCrlDp

`func (o *AuthMethodUpdateCert) GetRequireCrlDp() bool`

GetRequireCrlDp returns the RequireCrlDp field if non-nil, zero value otherwise.

### GetRequireCrlDpOk

`func (o *AuthMethodUpdateCert) GetRequireCrlDpOk() (*bool, bool)`

GetRequireCrlDpOk returns a tuple with the RequireCrlDp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCrlDp

`func (o *AuthMethodUpdateCert) SetRequireCrlDp(v bool)`

SetRequireCrlDp sets RequireCrlDp field to given value.

### HasRequireCrlDp

`func (o *AuthMethodUpdateCert) HasRequireCrlDp() bool`

HasRequireCrlDp returns a boolean if a field has been set.

### GetRevokedCertIds

`func (o *AuthMethodUpdateCert) GetRevokedCertIds() []string`

GetRevokedCertIds returns the RevokedCertIds field if non-nil, zero value otherwise.

### GetRevokedCertIdsOk

`func (o *AuthMethodUpdateCert) GetRevokedCertIdsOk() (*[]string, bool)`

GetRevokedCertIdsOk returns a tuple with the RevokedCertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCertIds

`func (o *AuthMethodUpdateCert) SetRevokedCertIds(v []string)`

SetRevokedCertIds sets RevokedCertIds field to given value.

### HasRevokedCertIds

`func (o *AuthMethodUpdateCert) HasRevokedCertIds() bool`

HasRevokedCertIds returns a boolean if a field has been set.

### GetToken

`func (o *AuthMethodUpdateCert) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthMethodUpdateCert) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthMethodUpdateCert) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthMethodUpdateCert) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AuthMethodUpdateCert) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AuthMethodUpdateCert) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AuthMethodUpdateCert) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AuthMethodUpdateCert) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *AuthMethodUpdateCert) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *AuthMethodUpdateCert) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *AuthMethodUpdateCert) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


