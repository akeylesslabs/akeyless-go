# AuthMethodCreateKerberos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedClientType** | Pointer to **[]string** | limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension] | [optional] 
**AuditLogsClaims** | Pointer to **[]string** | Subclaims to include in audit logs, e.g \&quot;--audit-logs-claims email --audit-logs-claims username\&quot; | [optional] 
**BindDn** | Pointer to **string** |  | [optional] 
**BindDnPassword** | Pointer to **string** |  | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the auth method would you like to be notified. | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GroupAttr** | Pointer to **string** |  | [optional] 
**GroupDn** | Pointer to **string** |  | [optional] 
**GroupFilter** | Pointer to **string** |  | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**KeytabFileData** | Pointer to **string** |  | [optional] 
**KeytabFilePath** | Pointer to **string** |  | [optional] 
**Krb5ConfData** | Pointer to **string** |  | [optional] 
**Krb5ConfPath** | Pointer to **string** |  | [optional] 
**LdapAnonymousSearch** | Pointer to **bool** |  | [optional] 
**LdapCaCert** | Pointer to **string** |  | [optional] 
**LdapUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** | Auth Method name | 
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca] | [optional] 
**SubclaimsDelimiters** | Pointer to **[]string** | A list of additional sub claims delimiters (relevant only for SAML, OIDC, OAuth2/JWT) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | Pointer to **string** | A unique identifier (ID) value which is a \&quot;sub claim\&quot; name that contains details uniquely identifying that resource. This \&quot;sub claim\&quot; is used to distinguish between different identities. | [optional] 
**UserAttribute** | Pointer to **string** |  | [optional] 
**UserDn** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthMethodCreateKerberos

`func NewAuthMethodCreateKerberos(name string, ) *AuthMethodCreateKerberos`

NewAuthMethodCreateKerberos instantiates a new AuthMethodCreateKerberos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodCreateKerberosWithDefaults

`func NewAuthMethodCreateKerberosWithDefaults() *AuthMethodCreateKerberos`

NewAuthMethodCreateKerberosWithDefaults instantiates a new AuthMethodCreateKerberos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *AuthMethodCreateKerberos) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *AuthMethodCreateKerberos) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *AuthMethodCreateKerberos) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *AuthMethodCreateKerberos) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedClientType

`func (o *AuthMethodCreateKerberos) GetAllowedClientType() []string`

GetAllowedClientType returns the AllowedClientType field if non-nil, zero value otherwise.

### GetAllowedClientTypeOk

`func (o *AuthMethodCreateKerberos) GetAllowedClientTypeOk() (*[]string, bool)`

GetAllowedClientTypeOk returns a tuple with the AllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientType

`func (o *AuthMethodCreateKerberos) SetAllowedClientType(v []string)`

SetAllowedClientType sets AllowedClientType field to given value.

### HasAllowedClientType

`func (o *AuthMethodCreateKerberos) HasAllowedClientType() bool`

HasAllowedClientType returns a boolean if a field has been set.

### GetAuditLogsClaims

`func (o *AuthMethodCreateKerberos) GetAuditLogsClaims() []string`

GetAuditLogsClaims returns the AuditLogsClaims field if non-nil, zero value otherwise.

### GetAuditLogsClaimsOk

`func (o *AuthMethodCreateKerberos) GetAuditLogsClaimsOk() (*[]string, bool)`

GetAuditLogsClaimsOk returns a tuple with the AuditLogsClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsClaims

`func (o *AuthMethodCreateKerberos) SetAuditLogsClaims(v []string)`

SetAuditLogsClaims sets AuditLogsClaims field to given value.

### HasAuditLogsClaims

`func (o *AuthMethodCreateKerberos) HasAuditLogsClaims() bool`

HasAuditLogsClaims returns a boolean if a field has been set.

### GetBindDn

`func (o *AuthMethodCreateKerberos) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *AuthMethodCreateKerberos) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *AuthMethodCreateKerberos) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *AuthMethodCreateKerberos) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindDnPassword

`func (o *AuthMethodCreateKerberos) GetBindDnPassword() string`

GetBindDnPassword returns the BindDnPassword field if non-nil, zero value otherwise.

### GetBindDnPasswordOk

`func (o *AuthMethodCreateKerberos) GetBindDnPasswordOk() (*string, bool)`

GetBindDnPasswordOk returns a tuple with the BindDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDnPassword

`func (o *AuthMethodCreateKerberos) SetBindDnPassword(v string)`

SetBindDnPassword sets BindDnPassword field to given value.

### HasBindDnPassword

`func (o *AuthMethodCreateKerberos) HasBindDnPassword() bool`

HasBindDnPassword returns a boolean if a field has been set.

### GetBoundIps

`func (o *AuthMethodCreateKerberos) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *AuthMethodCreateKerberos) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *AuthMethodCreateKerberos) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *AuthMethodCreateKerberos) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *AuthMethodCreateKerberos) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *AuthMethodCreateKerberos) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *AuthMethodCreateKerberos) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *AuthMethodCreateKerberos) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *AuthMethodCreateKerberos) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthMethodCreateKerberos) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthMethodCreateKerberos) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthMethodCreateKerberos) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *AuthMethodCreateKerberos) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *AuthMethodCreateKerberos) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *AuthMethodCreateKerberos) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *AuthMethodCreateKerberos) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *AuthMethodCreateKerberos) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *AuthMethodCreateKerberos) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *AuthMethodCreateKerberos) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *AuthMethodCreateKerberos) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGroupAttr

`func (o *AuthMethodCreateKerberos) GetGroupAttr() string`

GetGroupAttr returns the GroupAttr field if non-nil, zero value otherwise.

### GetGroupAttrOk

`func (o *AuthMethodCreateKerberos) GetGroupAttrOk() (*string, bool)`

GetGroupAttrOk returns a tuple with the GroupAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttr

`func (o *AuthMethodCreateKerberos) SetGroupAttr(v string)`

SetGroupAttr sets GroupAttr field to given value.

### HasGroupAttr

`func (o *AuthMethodCreateKerberos) HasGroupAttr() bool`

HasGroupAttr returns a boolean if a field has been set.

### GetGroupDn

`func (o *AuthMethodCreateKerberos) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *AuthMethodCreateKerberos) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *AuthMethodCreateKerberos) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *AuthMethodCreateKerberos) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetGroupFilter

`func (o *AuthMethodCreateKerberos) GetGroupFilter() string`

GetGroupFilter returns the GroupFilter field if non-nil, zero value otherwise.

### GetGroupFilterOk

`func (o *AuthMethodCreateKerberos) GetGroupFilterOk() (*string, bool)`

GetGroupFilterOk returns a tuple with the GroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFilter

`func (o *AuthMethodCreateKerberos) SetGroupFilter(v string)`

SetGroupFilter sets GroupFilter field to given value.

### HasGroupFilter

`func (o *AuthMethodCreateKerberos) HasGroupFilter() bool`

HasGroupFilter returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *AuthMethodCreateKerberos) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *AuthMethodCreateKerberos) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *AuthMethodCreateKerberos) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *AuthMethodCreateKerberos) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *AuthMethodCreateKerberos) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AuthMethodCreateKerberos) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AuthMethodCreateKerberos) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AuthMethodCreateKerberos) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *AuthMethodCreateKerberos) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *AuthMethodCreateKerberos) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *AuthMethodCreateKerberos) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *AuthMethodCreateKerberos) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetKeytabFileData

`func (o *AuthMethodCreateKerberos) GetKeytabFileData() string`

GetKeytabFileData returns the KeytabFileData field if non-nil, zero value otherwise.

### GetKeytabFileDataOk

`func (o *AuthMethodCreateKerberos) GetKeytabFileDataOk() (*string, bool)`

GetKeytabFileDataOk returns a tuple with the KeytabFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytabFileData

`func (o *AuthMethodCreateKerberos) SetKeytabFileData(v string)`

SetKeytabFileData sets KeytabFileData field to given value.

### HasKeytabFileData

`func (o *AuthMethodCreateKerberos) HasKeytabFileData() bool`

HasKeytabFileData returns a boolean if a field has been set.

### GetKeytabFilePath

`func (o *AuthMethodCreateKerberos) GetKeytabFilePath() string`

GetKeytabFilePath returns the KeytabFilePath field if non-nil, zero value otherwise.

### GetKeytabFilePathOk

`func (o *AuthMethodCreateKerberos) GetKeytabFilePathOk() (*string, bool)`

GetKeytabFilePathOk returns a tuple with the KeytabFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytabFilePath

`func (o *AuthMethodCreateKerberos) SetKeytabFilePath(v string)`

SetKeytabFilePath sets KeytabFilePath field to given value.

### HasKeytabFilePath

`func (o *AuthMethodCreateKerberos) HasKeytabFilePath() bool`

HasKeytabFilePath returns a boolean if a field has been set.

### GetKrb5ConfData

`func (o *AuthMethodCreateKerberos) GetKrb5ConfData() string`

GetKrb5ConfData returns the Krb5ConfData field if non-nil, zero value otherwise.

### GetKrb5ConfDataOk

`func (o *AuthMethodCreateKerberos) GetKrb5ConfDataOk() (*string, bool)`

GetKrb5ConfDataOk returns a tuple with the Krb5ConfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrb5ConfData

`func (o *AuthMethodCreateKerberos) SetKrb5ConfData(v string)`

SetKrb5ConfData sets Krb5ConfData field to given value.

### HasKrb5ConfData

`func (o *AuthMethodCreateKerberos) HasKrb5ConfData() bool`

HasKrb5ConfData returns a boolean if a field has been set.

### GetKrb5ConfPath

`func (o *AuthMethodCreateKerberos) GetKrb5ConfPath() string`

GetKrb5ConfPath returns the Krb5ConfPath field if non-nil, zero value otherwise.

### GetKrb5ConfPathOk

`func (o *AuthMethodCreateKerberos) GetKrb5ConfPathOk() (*string, bool)`

GetKrb5ConfPathOk returns a tuple with the Krb5ConfPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrb5ConfPath

`func (o *AuthMethodCreateKerberos) SetKrb5ConfPath(v string)`

SetKrb5ConfPath sets Krb5ConfPath field to given value.

### HasKrb5ConfPath

`func (o *AuthMethodCreateKerberos) HasKrb5ConfPath() bool`

HasKrb5ConfPath returns a boolean if a field has been set.

### GetLdapAnonymousSearch

`func (o *AuthMethodCreateKerberos) GetLdapAnonymousSearch() bool`

GetLdapAnonymousSearch returns the LdapAnonymousSearch field if non-nil, zero value otherwise.

### GetLdapAnonymousSearchOk

`func (o *AuthMethodCreateKerberos) GetLdapAnonymousSearchOk() (*bool, bool)`

GetLdapAnonymousSearchOk returns a tuple with the LdapAnonymousSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAnonymousSearch

`func (o *AuthMethodCreateKerberos) SetLdapAnonymousSearch(v bool)`

SetLdapAnonymousSearch sets LdapAnonymousSearch field to given value.

### HasLdapAnonymousSearch

`func (o *AuthMethodCreateKerberos) HasLdapAnonymousSearch() bool`

HasLdapAnonymousSearch returns a boolean if a field has been set.

### GetLdapCaCert

`func (o *AuthMethodCreateKerberos) GetLdapCaCert() string`

GetLdapCaCert returns the LdapCaCert field if non-nil, zero value otherwise.

### GetLdapCaCertOk

`func (o *AuthMethodCreateKerberos) GetLdapCaCertOk() (*string, bool)`

GetLdapCaCertOk returns a tuple with the LdapCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCaCert

`func (o *AuthMethodCreateKerberos) SetLdapCaCert(v string)`

SetLdapCaCert sets LdapCaCert field to given value.

### HasLdapCaCert

`func (o *AuthMethodCreateKerberos) HasLdapCaCert() bool`

HasLdapCaCert returns a boolean if a field has been set.

### GetLdapUrl

`func (o *AuthMethodCreateKerberos) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *AuthMethodCreateKerberos) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *AuthMethodCreateKerberos) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *AuthMethodCreateKerberos) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetName

`func (o *AuthMethodCreateKerberos) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthMethodCreateKerberos) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthMethodCreateKerberos) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *AuthMethodCreateKerberos) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuthMethodCreateKerberos) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuthMethodCreateKerberos) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuthMethodCreateKerberos) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSubclaimsDelimiters

`func (o *AuthMethodCreateKerberos) GetSubclaimsDelimiters() []string`

GetSubclaimsDelimiters returns the SubclaimsDelimiters field if non-nil, zero value otherwise.

### GetSubclaimsDelimitersOk

`func (o *AuthMethodCreateKerberos) GetSubclaimsDelimitersOk() (*[]string, bool)`

GetSubclaimsDelimitersOk returns a tuple with the SubclaimsDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclaimsDelimiters

`func (o *AuthMethodCreateKerberos) SetSubclaimsDelimiters(v []string)`

SetSubclaimsDelimiters sets SubclaimsDelimiters field to given value.

### HasSubclaimsDelimiters

`func (o *AuthMethodCreateKerberos) HasSubclaimsDelimiters() bool`

HasSubclaimsDelimiters returns a boolean if a field has been set.

### GetToken

`func (o *AuthMethodCreateKerberos) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthMethodCreateKerberos) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthMethodCreateKerberos) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthMethodCreateKerberos) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AuthMethodCreateKerberos) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AuthMethodCreateKerberos) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AuthMethodCreateKerberos) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AuthMethodCreateKerberos) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *AuthMethodCreateKerberos) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *AuthMethodCreateKerberos) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *AuthMethodCreateKerberos) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *AuthMethodCreateKerberos) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.

### GetUserAttribute

`func (o *AuthMethodCreateKerberos) GetUserAttribute() string`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *AuthMethodCreateKerberos) GetUserAttributeOk() (*string, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *AuthMethodCreateKerberos) SetUserAttribute(v string)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *AuthMethodCreateKerberos) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### GetUserDn

`func (o *AuthMethodCreateKerberos) GetUserDn() string`

GetUserDn returns the UserDn field if non-nil, zero value otherwise.

### GetUserDnOk

`func (o *AuthMethodCreateKerberos) GetUserDnOk() (*string, bool)`

GetUserDnOk returns a tuple with the UserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDn

`func (o *AuthMethodCreateKerberos) SetUserDn(v string)`

SetUserDn sets UserDn field to given value.

### HasUserDn

`func (o *AuthMethodCreateKerberos) HasUserDn() bool`

HasUserDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


