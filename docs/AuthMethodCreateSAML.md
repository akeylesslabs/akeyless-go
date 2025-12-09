# AuthMethodCreateSAML

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedClientType** | Pointer to **[]string** | limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension] | [optional] 
**AllowedRedirectUri** | Pointer to **[]string** | Allowed redirect URIs after the authentication | [optional] 
**AuditLogsClaims** | Pointer to **[]string** | Subclaims to include in audit logs, e.g \&quot;--audit-logs-claims email --audit-logs-claims username\&quot; | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the auth method would you like to be notified. | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**IdpMetadataUrl** | Pointer to **string** | IDP metadata url | [optional] 
**IdpMetadataXmlData** | Pointer to **string** | IDP metadata xml data | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca] | [optional] 
**SubclaimsDelimiters** | Pointer to **[]string** | A list of additional sub claims delimiters (relevant only for SAML, OIDC, OAuth2/JWT) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | **string** | A unique identifier (ID) value should be configured for OAuth2, LDAP and SAML authentication method types and is usually a value such as the email, username, or upn for example. Whenever a user logs in with a token, these authentication types issue a \&quot;sub claim\&quot; that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization. | 

## Methods

### NewAuthMethodCreateSAML

`func NewAuthMethodCreateSAML(name string, uniqueIdentifier string, ) *AuthMethodCreateSAML`

NewAuthMethodCreateSAML instantiates a new AuthMethodCreateSAML object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodCreateSAMLWithDefaults

`func NewAuthMethodCreateSAMLWithDefaults() *AuthMethodCreateSAML`

NewAuthMethodCreateSAMLWithDefaults instantiates a new AuthMethodCreateSAML object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *AuthMethodCreateSAML) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *AuthMethodCreateSAML) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *AuthMethodCreateSAML) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *AuthMethodCreateSAML) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedClientType

`func (o *AuthMethodCreateSAML) GetAllowedClientType() []string`

GetAllowedClientType returns the AllowedClientType field if non-nil, zero value otherwise.

### GetAllowedClientTypeOk

`func (o *AuthMethodCreateSAML) GetAllowedClientTypeOk() (*[]string, bool)`

GetAllowedClientTypeOk returns a tuple with the AllowedClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClientType

`func (o *AuthMethodCreateSAML) SetAllowedClientType(v []string)`

SetAllowedClientType sets AllowedClientType field to given value.

### HasAllowedClientType

`func (o *AuthMethodCreateSAML) HasAllowedClientType() bool`

HasAllowedClientType returns a boolean if a field has been set.

### GetAllowedRedirectUri

`func (o *AuthMethodCreateSAML) GetAllowedRedirectUri() []string`

GetAllowedRedirectUri returns the AllowedRedirectUri field if non-nil, zero value otherwise.

### GetAllowedRedirectUriOk

`func (o *AuthMethodCreateSAML) GetAllowedRedirectUriOk() (*[]string, bool)`

GetAllowedRedirectUriOk returns a tuple with the AllowedRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRedirectUri

`func (o *AuthMethodCreateSAML) SetAllowedRedirectUri(v []string)`

SetAllowedRedirectUri sets AllowedRedirectUri field to given value.

### HasAllowedRedirectUri

`func (o *AuthMethodCreateSAML) HasAllowedRedirectUri() bool`

HasAllowedRedirectUri returns a boolean if a field has been set.

### GetAuditLogsClaims

`func (o *AuthMethodCreateSAML) GetAuditLogsClaims() []string`

GetAuditLogsClaims returns the AuditLogsClaims field if non-nil, zero value otherwise.

### GetAuditLogsClaimsOk

`func (o *AuthMethodCreateSAML) GetAuditLogsClaimsOk() (*[]string, bool)`

GetAuditLogsClaimsOk returns a tuple with the AuditLogsClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogsClaims

`func (o *AuthMethodCreateSAML) SetAuditLogsClaims(v []string)`

SetAuditLogsClaims sets AuditLogsClaims field to given value.

### HasAuditLogsClaims

`func (o *AuthMethodCreateSAML) HasAuditLogsClaims() bool`

HasAuditLogsClaims returns a boolean if a field has been set.

### GetBoundIps

`func (o *AuthMethodCreateSAML) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *AuthMethodCreateSAML) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *AuthMethodCreateSAML) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *AuthMethodCreateSAML) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *AuthMethodCreateSAML) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *AuthMethodCreateSAML) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *AuthMethodCreateSAML) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *AuthMethodCreateSAML) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *AuthMethodCreateSAML) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthMethodCreateSAML) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthMethodCreateSAML) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthMethodCreateSAML) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *AuthMethodCreateSAML) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *AuthMethodCreateSAML) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *AuthMethodCreateSAML) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *AuthMethodCreateSAML) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *AuthMethodCreateSAML) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *AuthMethodCreateSAML) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *AuthMethodCreateSAML) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *AuthMethodCreateSAML) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *AuthMethodCreateSAML) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *AuthMethodCreateSAML) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *AuthMethodCreateSAML) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *AuthMethodCreateSAML) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetIdpMetadataUrl

`func (o *AuthMethodCreateSAML) GetIdpMetadataUrl() string`

GetIdpMetadataUrl returns the IdpMetadataUrl field if non-nil, zero value otherwise.

### GetIdpMetadataUrlOk

`func (o *AuthMethodCreateSAML) GetIdpMetadataUrlOk() (*string, bool)`

GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataUrl

`func (o *AuthMethodCreateSAML) SetIdpMetadataUrl(v string)`

SetIdpMetadataUrl sets IdpMetadataUrl field to given value.

### HasIdpMetadataUrl

`func (o *AuthMethodCreateSAML) HasIdpMetadataUrl() bool`

HasIdpMetadataUrl returns a boolean if a field has been set.

### GetIdpMetadataXmlData

`func (o *AuthMethodCreateSAML) GetIdpMetadataXmlData() string`

GetIdpMetadataXmlData returns the IdpMetadataXmlData field if non-nil, zero value otherwise.

### GetIdpMetadataXmlDataOk

`func (o *AuthMethodCreateSAML) GetIdpMetadataXmlDataOk() (*string, bool)`

GetIdpMetadataXmlDataOk returns a tuple with the IdpMetadataXmlData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXmlData

`func (o *AuthMethodCreateSAML) SetIdpMetadataXmlData(v string)`

SetIdpMetadataXmlData sets IdpMetadataXmlData field to given value.

### HasIdpMetadataXmlData

`func (o *AuthMethodCreateSAML) HasIdpMetadataXmlData() bool`

HasIdpMetadataXmlData returns a boolean if a field has been set.

### GetJson

`func (o *AuthMethodCreateSAML) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AuthMethodCreateSAML) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AuthMethodCreateSAML) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AuthMethodCreateSAML) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *AuthMethodCreateSAML) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *AuthMethodCreateSAML) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *AuthMethodCreateSAML) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *AuthMethodCreateSAML) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *AuthMethodCreateSAML) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthMethodCreateSAML) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthMethodCreateSAML) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *AuthMethodCreateSAML) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuthMethodCreateSAML) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuthMethodCreateSAML) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuthMethodCreateSAML) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSubclaimsDelimiters

`func (o *AuthMethodCreateSAML) GetSubclaimsDelimiters() []string`

GetSubclaimsDelimiters returns the SubclaimsDelimiters field if non-nil, zero value otherwise.

### GetSubclaimsDelimitersOk

`func (o *AuthMethodCreateSAML) GetSubclaimsDelimitersOk() (*[]string, bool)`

GetSubclaimsDelimitersOk returns a tuple with the SubclaimsDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclaimsDelimiters

`func (o *AuthMethodCreateSAML) SetSubclaimsDelimiters(v []string)`

SetSubclaimsDelimiters sets SubclaimsDelimiters field to given value.

### HasSubclaimsDelimiters

`func (o *AuthMethodCreateSAML) HasSubclaimsDelimiters() bool`

HasSubclaimsDelimiters returns a boolean if a field has been set.

### GetToken

`func (o *AuthMethodCreateSAML) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthMethodCreateSAML) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthMethodCreateSAML) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthMethodCreateSAML) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *AuthMethodCreateSAML) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *AuthMethodCreateSAML) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *AuthMethodCreateSAML) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *AuthMethodCreateSAML) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *AuthMethodCreateSAML) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *AuthMethodCreateSAML) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *AuthMethodCreateSAML) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


