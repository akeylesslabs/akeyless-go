# AuthMethodCreateCert

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
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca] | [optional] 
**RequireCrlDp** | Pointer to **bool** | Require certificate CRL distribution points (CDP) and enforce CRL validation during authentication. | [optional] 
**RevokedCertIds** | Pointer to **[]string** | A list of revoked cert ids | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | **string** | A unique identifier (ID) value should be configured, such as common_name or organizational_unit Whenever a user logs in with a token, these authentication types issue a \&quot;sub claim\&quot; that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization. | 

## Methods

### NewAuthMethodCreateCert

`func NewAuthMethodCreateCert(name string, uniqueIdentifier string, ) *AuthMethodCreateCert`

NewAuthMethodCreateCert instantiates a new AuthMethodCreateCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodCreateCertWithDefaults

`func NewAuthMethodCreateCertWithDefaults() *AuthMethodCreateCert`

NewAuthMethodCreateCertWithDefaults instantiates a new AuthMethodCreateCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *AuthMethodCreateCert) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *AuthMethodCreateCert) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *AuthMethodCreateCert) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *AuthMethodCreateCert) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedClientType

`func (o *AuthMethodCreateCert) GetAllowedClientType() []string`

GetAllowedClientType returns the AllowedClientType field if non-nil, zero value otherwise.

### GetAllowedClientTypeOk

`func (o *AuthMethodCreateCert) GetAllowedClientTypeOk() (*[]string, bool)`

GetAllowedClientTypeOk returns a tuple with the AllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientType

`func (o *AuthMethodCreateCert) SetAllowedClientType(v []string)`

SetAllowedClientType sets AllowedClientType field to given value.

### HasAllowedClientType

`func (o *AuthMethodCreateCert) HasAllowedClientType() bool`

HasAllowedClientType returns a boolean if a field has been set.

### GetAllowedCors

`func (o *AuthMethodCreateCert) GetAllowedCors() string`

GetAllowedCors returns the AllowedCors field if non-nil, zero value otherwise.

### GetAllowedCorsOk

`func (o *AuthMethodCreateCert) GetAllowedCorsOk() (*string, bool)`

GetAllowedCorsOk returns a tuple with the AllowedCors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCors

`func (o *AuthMethodCreateCert) SetAllowedCors(v string)`

SetAllowedCors sets AllowedCors field to given value.

### HasAllowedCors

`func (o *AuthMethodCreateCert) HasAllowedCors() bool`

HasAllowedCors returns a boolean if a field has been set.

### GetAuditLogsClaims

`func (o *AuthMethodCreateCert) GetAuditLogsClaims() []string`

GetAuditLogsClaims returns the AuditLogsClaims field if non-nil, zero value otherwise.

### GetAuditLogsClaimsOk

`func (o *AuthMethodCreateCert) GetAuditLogsClaimsOk() (*[]string, bool)`

GetAuditLogsClaimsOk returns a tuple with the AuditLogsClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsClaims

`func (o *AuthMethodCreateCert) SetAuditLogsClaims(v []string)`

SetAuditLogsClaims sets AuditLogsClaims field to given value.

### HasAuditLogsClaims

`func (o *AuthMethodCreateCert) HasAuditLogsClaims() bool`

HasAuditLogsClaims returns a boolean if a field has been set.

### GetBoundCommonNames

`func (o *AuthMethodCreateCert) GetBoundCommonNames() []string`

GetBoundCommonNames returns the BoundCommonNames field if non-nil, zero value otherwise.

### GetBoundCommonNamesOk

`func (o *AuthMethodCreateCert) GetBoundCommonNamesOk() (*[]string, bool)`

GetBoundCommonNamesOk returns a tuple with the BoundCommonNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCommonNames

`func (o *AuthMethodCreateCert) SetBoundCommonNames(v []string)`

SetBoundCommonNames sets BoundCommonNames field to given value.

### HasBoundCommonNames

`func (o *AuthMethodCreateCert) HasBoundCommonNames() bool`

HasBoundCommonNames returns a boolean if a field has been set.

### GetBoundDnsSans

`func (o *AuthMethodCreateCert) GetBoundDnsSans() []string`

GetBoundDnsSans returns the BoundDnsSans field if non-nil, zero value otherwise.

### GetBoundDnsSansOk

`func (o *AuthMethodCreateCert) GetBoundDnsSansOk() (*[]string, bool)`

GetBoundDnsSansOk returns a tuple with the BoundDnsSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDnsSans

`func (o *AuthMethodCreateCert) SetBoundDnsSans(v []string)`

SetBoundDnsSans sets BoundDnsSans field to given value.

### HasBoundDnsSans

`func (o *AuthMethodCreateCert) HasBoundDnsSans() bool`

HasBoundDnsSans returns a boolean if a field has been set.

### GetBoundEmailSans

`func (o *AuthMethodCreateCert) GetBoundEmailSans() []string`

GetBoundEmailSans returns the BoundEmailSans field if non-nil, zero value otherwise.

### GetBoundEmailSansOk

`func (o *AuthMethodCreateCert) GetBoundEmailSansOk() (*[]string, bool)`

GetBoundEmailSansOk returns a tuple with the BoundEmailSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEmailSans

`func (o *AuthMethodCreateCert) SetBoundEmailSans(v []string)`

SetBoundEmailSans sets BoundEmailSans field to given value.

### HasBoundEmailSans

`func (o *AuthMethodCreateCert) HasBoundEmailSans() bool`

HasBoundEmailSans returns a boolean if a field has been set.

### GetBoundExtensions

`func (o *AuthMethodCreateCert) GetBoundExtensions() []string`

GetBoundExtensions returns the BoundExtensions field if non-nil, zero value otherwise.

### GetBoundExtensionsOk

`func (o *AuthMethodCreateCert) GetBoundExtensionsOk() (*[]string, bool)`

GetBoundExtensionsOk returns a tuple with the BoundExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundExtensions

`func (o *AuthMethodCreateCert) SetBoundExtensions(v []string)`

SetBoundExtensions sets BoundExtensions field to given value.

### HasBoundExtensions

`func (o *AuthMethodCreateCert) HasBoundExtensions() bool`

HasBoundExtensions returns a boolean if a field has been set.

### GetBoundIps

`func (o *AuthMethodCreateCert) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *AuthMethodCreateCert) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *AuthMethodCreateCert) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *AuthMethodCreateCert) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundOrganizationalUnits

`func (o *AuthMethodCreateCert) GetBoundOrganizationalUnits() []string`

GetBoundOrganizationalUnits returns the BoundOrganizationalUnits field if non-nil, zero value otherwise.

### GetBoundOrganizationalUnitsOk

`func (o *AuthMethodCreateCert) GetBoundOrganizationalUnitsOk() (*[]string, bool)`

GetBoundOrganizationalUnitsOk returns a tuple with the BoundOrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundOrganizationalUnits

`func (o *AuthMethodCreateCert) SetBoundOrganizationalUnits(v []string)`

SetBoundOrganizationalUnits sets BoundOrganizationalUnits field to given value.

### HasBoundOrganizationalUnits

`func (o *AuthMethodCreateCert) HasBoundOrganizationalUnits() bool`

HasBoundOrganizationalUnits returns a boolean if a field has been set.

### GetBoundUriSans

`func (o *AuthMethodCreateCert) GetBoundUriSans() []string`

GetBoundUriSans returns the BoundUriSans field if non-nil, zero value otherwise.

### GetBoundUriSansOk

`func (o *AuthMethodCreateCert) GetBoundUriSansOk() (*[]string, bool)`

GetBoundUriSansOk returns a tuple with the BoundUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUriSans

`func (o *AuthMethodCreateCert) SetBoundUriSans(v []string)`

SetBoundUriSans sets BoundUriSans field to given value.

### HasBoundUriSans

`func (o *AuthMethodCreateCert) HasBoundUriSans() bool`

HasBoundUriSans returns a boolean if a field has been set.

### GetCertificateData

`func (o *AuthMethodCreateCert) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *AuthMethodCreateCert) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *AuthMethodCreateCert) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.

### HasCertificateData

`func (o *AuthMethodCreateCert) HasCertificateData() bool`

HasCertificateData returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *AuthMethodCreateCert) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *AuthMethodCreateCert) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *AuthMethodCreateCert) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *AuthMethodCreateCert) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *AuthMethodCreateCert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthMethodCreateCert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthMethodCreateCert) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthMethodCreateCert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *AuthMethodCreateCert) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *AuthMethodCreateCert) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *AuthMethodCreateCert) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *AuthMethodCreateCert) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *AuthMethodCreateCert) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *AuthMethodCreateCert) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *AuthMethodCreateCert) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *AuthMethodCreateCert) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *AuthMethodCreateCert) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *AuthMethodCreateCert) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *AuthMethodCreateCert) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *AuthMethodCreateCert) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *AuthMethodCreateCert) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AuthMethodCreateCert) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AuthMethodCreateCert) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AuthMethodCreateCert) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *AuthMethodCreateCert) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *AuthMethodCreateCert) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *AuthMethodCreateCert) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *AuthMethodCreateCert) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *AuthMethodCreateCert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthMethodCreateCert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthMethodCreateCert) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *AuthMethodCreateCert) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuthMethodCreateCert) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuthMethodCreateCert) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuthMethodCreateCert) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetRequireCrlDp

`func (o *AuthMethodCreateCert) GetRequireCrlDp() bool`

GetRequireCrlDp returns the RequireCrlDp field if non-nil, zero value otherwise.

### GetRequireCrlDpOk

`func (o *AuthMethodCreateCert) GetRequireCrlDpOk() (*bool, bool)`

GetRequireCrlDpOk returns a tuple with the RequireCrlDp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCrlDp

`func (o *AuthMethodCreateCert) SetRequireCrlDp(v bool)`

SetRequireCrlDp sets RequireCrlDp field to given value.

### HasRequireCrlDp

`func (o *AuthMethodCreateCert) HasRequireCrlDp() bool`

HasRequireCrlDp returns a boolean if a field has been set.

### GetRevokedCertIds

`func (o *AuthMethodCreateCert) GetRevokedCertIds() []string`

GetRevokedCertIds returns the RevokedCertIds field if non-nil, zero value otherwise.

### GetRevokedCertIdsOk

`func (o *AuthMethodCreateCert) GetRevokedCertIdsOk() (*[]string, bool)`

GetRevokedCertIdsOk returns a tuple with the RevokedCertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCertIds

`func (o *AuthMethodCreateCert) SetRevokedCertIds(v []string)`

SetRevokedCertIds sets RevokedCertIds field to given value.

### HasRevokedCertIds

`func (o *AuthMethodCreateCert) HasRevokedCertIds() bool`

HasRevokedCertIds returns a boolean if a field has been set.

### GetToken

`func (o *AuthMethodCreateCert) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthMethodCreateCert) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthMethodCreateCert) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthMethodCreateCert) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AuthMethodCreateCert) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AuthMethodCreateCert) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AuthMethodCreateCert) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AuthMethodCreateCert) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *AuthMethodCreateCert) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *AuthMethodCreateCert) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *AuthMethodCreateCert) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


