# CreateAuthMethodSAML

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedRedirectUri** | Pointer to **[]string** | Allowed redirect URIs after the authentication | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
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

### NewCreateAuthMethodSAML

`func NewCreateAuthMethodSAML(name string, uniqueIdentifier string, ) *CreateAuthMethodSAML`

NewCreateAuthMethodSAML instantiates a new CreateAuthMethodSAML object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodSAMLWithDefaults

`func NewCreateAuthMethodSAMLWithDefaults() *CreateAuthMethodSAML`

NewCreateAuthMethodSAMLWithDefaults instantiates a new CreateAuthMethodSAML object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodSAML) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodSAML) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodSAML) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodSAML) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedRedirectUri

`func (o *CreateAuthMethodSAML) GetAllowedRedirectUri() []string`

GetAllowedRedirectUri returns the AllowedRedirectUri field if non-nil, zero value otherwise.

### GetAllowedRedirectUriOk

`func (o *CreateAuthMethodSAML) GetAllowedRedirectUriOk() (*[]string, bool)`

GetAllowedRedirectUriOk returns a tuple with the AllowedRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRedirectUri

`func (o *CreateAuthMethodSAML) SetAllowedRedirectUri(v []string)`

SetAllowedRedirectUri sets AllowedRedirectUri field to given value.

### HasAllowedRedirectUri

`func (o *CreateAuthMethodSAML) HasAllowedRedirectUri() bool`

HasAllowedRedirectUri returns a boolean if a field has been set.

### GetBoundIps

`func (o *CreateAuthMethodSAML) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodSAML) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodSAML) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodSAML) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAuthMethodSAML) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAuthMethodSAML) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAuthMethodSAML) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAuthMethodSAML) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *CreateAuthMethodSAML) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodSAML) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodSAML) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodSAML) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *CreateAuthMethodSAML) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *CreateAuthMethodSAML) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *CreateAuthMethodSAML) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *CreateAuthMethodSAML) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetIdpMetadataUrl

`func (o *CreateAuthMethodSAML) GetIdpMetadataUrl() string`

GetIdpMetadataUrl returns the IdpMetadataUrl field if non-nil, zero value otherwise.

### GetIdpMetadataUrlOk

`func (o *CreateAuthMethodSAML) GetIdpMetadataUrlOk() (*string, bool)`

GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataUrl

`func (o *CreateAuthMethodSAML) SetIdpMetadataUrl(v string)`

SetIdpMetadataUrl sets IdpMetadataUrl field to given value.

### HasIdpMetadataUrl

`func (o *CreateAuthMethodSAML) HasIdpMetadataUrl() bool`

HasIdpMetadataUrl returns a boolean if a field has been set.

### GetIdpMetadataXmlData

`func (o *CreateAuthMethodSAML) GetIdpMetadataXmlData() string`

GetIdpMetadataXmlData returns the IdpMetadataXmlData field if non-nil, zero value otherwise.

### GetIdpMetadataXmlDataOk

`func (o *CreateAuthMethodSAML) GetIdpMetadataXmlDataOk() (*string, bool)`

GetIdpMetadataXmlDataOk returns a tuple with the IdpMetadataXmlData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXmlData

`func (o *CreateAuthMethodSAML) SetIdpMetadataXmlData(v string)`

SetIdpMetadataXmlData sets IdpMetadataXmlData field to given value.

### HasIdpMetadataXmlData

`func (o *CreateAuthMethodSAML) HasIdpMetadataXmlData() bool`

HasIdpMetadataXmlData returns a boolean if a field has been set.

### GetJson

`func (o *CreateAuthMethodSAML) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateAuthMethodSAML) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateAuthMethodSAML) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateAuthMethodSAML) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *CreateAuthMethodSAML) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *CreateAuthMethodSAML) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *CreateAuthMethodSAML) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *CreateAuthMethodSAML) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodSAML) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodSAML) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodSAML) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *CreateAuthMethodSAML) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *CreateAuthMethodSAML) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *CreateAuthMethodSAML) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *CreateAuthMethodSAML) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSubclaimsDelimiters

`func (o *CreateAuthMethodSAML) GetSubclaimsDelimiters() []string`

GetSubclaimsDelimiters returns the SubclaimsDelimiters field if non-nil, zero value otherwise.

### GetSubclaimsDelimitersOk

`func (o *CreateAuthMethodSAML) GetSubclaimsDelimitersOk() (*[]string, bool)`

GetSubclaimsDelimitersOk returns a tuple with the SubclaimsDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclaimsDelimiters

`func (o *CreateAuthMethodSAML) SetSubclaimsDelimiters(v []string)`

SetSubclaimsDelimiters sets SubclaimsDelimiters field to given value.

### HasSubclaimsDelimiters

`func (o *CreateAuthMethodSAML) HasSubclaimsDelimiters() bool`

HasSubclaimsDelimiters returns a boolean if a field has been set.

### GetToken

`func (o *CreateAuthMethodSAML) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodSAML) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodSAML) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodSAML) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodSAML) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodSAML) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodSAML) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodSAML) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *CreateAuthMethodSAML) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *CreateAuthMethodSAML) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *CreateAuthMethodSAML) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


