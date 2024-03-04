# CreateAuthMethodOIDC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedRedirectUri** | Pointer to **[]string** | Allowed redirect URIs after the authentication | [optional] 
**Audience** | Pointer to **string** | Audience claim to be used as part of the authentication flow. In case set, it must match the one configured on the Identity Provider&#39;s Application | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**ClientId** | Pointer to **string** | Client ID | [optional] 
**ClientSecret** | Pointer to **string** | Client Secret | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Issuer** | Pointer to **string** | Issuer URL | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca] | [optional] 
**RequiredScopes** | Pointer to **[]string** | RequiredScopes is a list of required scopes that the oidc method will request from the oidc provider and the user must approve | [optional] 
**RequiredScopesPrefix** | Pointer to **string** | RequiredScopesPrefix is a a prefix to add to all required-scopes when requesting them from the oidc server (for example, azures&#39; Application ID URI) | [optional] 
**SubclaimsDelimiters** | Pointer to **[]string** | A list of additional sub claims delimiters (relevant only for SAML, OIDC, OAuth2/JWT) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | **string** | A unique identifier (ID) value should be configured for OIDC, OAuth2, LDAP and SAML authentication method types and is usually a value such as the email, username, or upn for example. Whenever a user logs in with a token, these authentication types issue a \&quot;sub claim\&quot; that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization. | 

## Methods

### NewCreateAuthMethodOIDC

`func NewCreateAuthMethodOIDC(name string, uniqueIdentifier string, ) *CreateAuthMethodOIDC`

NewCreateAuthMethodOIDC instantiates a new CreateAuthMethodOIDC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodOIDCWithDefaults

`func NewCreateAuthMethodOIDCWithDefaults() *CreateAuthMethodOIDC`

NewCreateAuthMethodOIDCWithDefaults instantiates a new CreateAuthMethodOIDC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodOIDC) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodOIDC) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodOIDC) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodOIDC) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedRedirectUri

`func (o *CreateAuthMethodOIDC) GetAllowedRedirectUri() []string`

GetAllowedRedirectUri returns the AllowedRedirectUri field if non-nil, zero value otherwise.

### GetAllowedRedirectUriOk

`func (o *CreateAuthMethodOIDC) GetAllowedRedirectUriOk() (*[]string, bool)`

GetAllowedRedirectUriOk returns a tuple with the AllowedRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRedirectUri

`func (o *CreateAuthMethodOIDC) SetAllowedRedirectUri(v []string)`

SetAllowedRedirectUri sets AllowedRedirectUri field to given value.

### HasAllowedRedirectUri

`func (o *CreateAuthMethodOIDC) HasAllowedRedirectUri() bool`

HasAllowedRedirectUri returns a boolean if a field has been set.

### GetAudience

`func (o *CreateAuthMethodOIDC) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CreateAuthMethodOIDC) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CreateAuthMethodOIDC) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CreateAuthMethodOIDC) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBoundIps

`func (o *CreateAuthMethodOIDC) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodOIDC) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodOIDC) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodOIDC) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetClientId

`func (o *CreateAuthMethodOIDC) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateAuthMethodOIDC) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateAuthMethodOIDC) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateAuthMethodOIDC) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *CreateAuthMethodOIDC) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateAuthMethodOIDC) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateAuthMethodOIDC) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CreateAuthMethodOIDC) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAuthMethodOIDC) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAuthMethodOIDC) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAuthMethodOIDC) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAuthMethodOIDC) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *CreateAuthMethodOIDC) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodOIDC) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodOIDC) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodOIDC) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *CreateAuthMethodOIDC) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *CreateAuthMethodOIDC) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *CreateAuthMethodOIDC) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *CreateAuthMethodOIDC) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetIssuer

`func (o *CreateAuthMethodOIDC) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CreateAuthMethodOIDC) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CreateAuthMethodOIDC) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CreateAuthMethodOIDC) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJson

`func (o *CreateAuthMethodOIDC) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateAuthMethodOIDC) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateAuthMethodOIDC) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateAuthMethodOIDC) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *CreateAuthMethodOIDC) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *CreateAuthMethodOIDC) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *CreateAuthMethodOIDC) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *CreateAuthMethodOIDC) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodOIDC) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodOIDC) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodOIDC) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *CreateAuthMethodOIDC) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *CreateAuthMethodOIDC) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *CreateAuthMethodOIDC) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *CreateAuthMethodOIDC) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetRequiredScopes

`func (o *CreateAuthMethodOIDC) GetRequiredScopes() []string`

GetRequiredScopes returns the RequiredScopes field if non-nil, zero value otherwise.

### GetRequiredScopesOk

`func (o *CreateAuthMethodOIDC) GetRequiredScopesOk() (*[]string, bool)`

GetRequiredScopesOk returns a tuple with the RequiredScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredScopes

`func (o *CreateAuthMethodOIDC) SetRequiredScopes(v []string)`

SetRequiredScopes sets RequiredScopes field to given value.

### HasRequiredScopes

`func (o *CreateAuthMethodOIDC) HasRequiredScopes() bool`

HasRequiredScopes returns a boolean if a field has been set.

### GetRequiredScopesPrefix

`func (o *CreateAuthMethodOIDC) GetRequiredScopesPrefix() string`

GetRequiredScopesPrefix returns the RequiredScopesPrefix field if non-nil, zero value otherwise.

### GetRequiredScopesPrefixOk

`func (o *CreateAuthMethodOIDC) GetRequiredScopesPrefixOk() (*string, bool)`

GetRequiredScopesPrefixOk returns a tuple with the RequiredScopesPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredScopesPrefix

`func (o *CreateAuthMethodOIDC) SetRequiredScopesPrefix(v string)`

SetRequiredScopesPrefix sets RequiredScopesPrefix field to given value.

### HasRequiredScopesPrefix

`func (o *CreateAuthMethodOIDC) HasRequiredScopesPrefix() bool`

HasRequiredScopesPrefix returns a boolean if a field has been set.

### GetSubclaimsDelimiters

`func (o *CreateAuthMethodOIDC) GetSubclaimsDelimiters() []string`

GetSubclaimsDelimiters returns the SubclaimsDelimiters field if non-nil, zero value otherwise.

### GetSubclaimsDelimitersOk

`func (o *CreateAuthMethodOIDC) GetSubclaimsDelimitersOk() (*[]string, bool)`

GetSubclaimsDelimitersOk returns a tuple with the SubclaimsDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclaimsDelimiters

`func (o *CreateAuthMethodOIDC) SetSubclaimsDelimiters(v []string)`

SetSubclaimsDelimiters sets SubclaimsDelimiters field to given value.

### HasSubclaimsDelimiters

`func (o *CreateAuthMethodOIDC) HasSubclaimsDelimiters() bool`

HasSubclaimsDelimiters returns a boolean if a field has been set.

### GetToken

`func (o *CreateAuthMethodOIDC) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodOIDC) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodOIDC) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodOIDC) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodOIDC) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodOIDC) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodOIDC) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodOIDC) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *CreateAuthMethodOIDC) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *CreateAuthMethodOIDC) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *CreateAuthMethodOIDC) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


