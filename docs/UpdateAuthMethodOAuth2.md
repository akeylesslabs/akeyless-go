# UpdateAuthMethodOAuth2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**Audience** | Pointer to **string** | The audience in the JWT | [optional] 
**BoundClientIds** | Pointer to **[]string** | The clients ids that the access is restricted to | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GatewayUrl** | Pointer to **string** | Akeyless Gateway URL (Configuration Management port). Relevant only when the jwks-uri is accessible only from the gateway. | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Issuer** | Pointer to **string** | Issuer URL | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwksJsonData** | Pointer to **string** | The JSON Web Key Set (JWKS) that containing the public keys that should be used to verify any JSON Web Token (JWT) issued by the authorization server. base64 encoded string | [optional] 
**JwksUri** | **string** | The URL to the JSON Web Key Set (JWKS) that containing the public keys that should be used to verify any JSON Web Token (JWT) issued by the authorization server. | [default to "default_jwks_url"]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**SubclaimsDelimiters** | Pointer to **[]string** | A list of additional sub claims delimiters (relevant only for SAML, OIDC, OAuth2/JWT) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | **string** | A unique identifier (ID) value should be configured for OAuth2, LDAP and SAML authentication method types and is usually a value such as the email, username, or upn for example. Whenever a user logs in with a token, these authentication types issue a \&quot;sub claim\&quot; that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization. | 

## Methods

### NewUpdateAuthMethodOAuth2

`func NewUpdateAuthMethodOAuth2(jwksUri string, name string, uniqueIdentifier string, ) *UpdateAuthMethodOAuth2`

NewUpdateAuthMethodOAuth2 instantiates a new UpdateAuthMethodOAuth2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthMethodOAuth2WithDefaults

`func NewUpdateAuthMethodOAuth2WithDefaults() *UpdateAuthMethodOAuth2`

NewUpdateAuthMethodOAuth2WithDefaults instantiates a new UpdateAuthMethodOAuth2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *UpdateAuthMethodOAuth2) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *UpdateAuthMethodOAuth2) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *UpdateAuthMethodOAuth2) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *UpdateAuthMethodOAuth2) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAudience

`func (o *UpdateAuthMethodOAuth2) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *UpdateAuthMethodOAuth2) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *UpdateAuthMethodOAuth2) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *UpdateAuthMethodOAuth2) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBoundClientIds

`func (o *UpdateAuthMethodOAuth2) GetBoundClientIds() []string`

GetBoundClientIds returns the BoundClientIds field if non-nil, zero value otherwise.

### GetBoundClientIdsOk

`func (o *UpdateAuthMethodOAuth2) GetBoundClientIdsOk() (*[]string, bool)`

GetBoundClientIdsOk returns a tuple with the BoundClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundClientIds

`func (o *UpdateAuthMethodOAuth2) SetBoundClientIds(v []string)`

SetBoundClientIds sets BoundClientIds field to given value.

### HasBoundClientIds

`func (o *UpdateAuthMethodOAuth2) HasBoundClientIds() bool`

HasBoundClientIds returns a boolean if a field has been set.

### GetBoundIps

`func (o *UpdateAuthMethodOAuth2) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAuthMethodOAuth2) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAuthMethodOAuth2) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAuthMethodOAuth2) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *UpdateAuthMethodOAuth2) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *UpdateAuthMethodOAuth2) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *UpdateAuthMethodOAuth2) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *UpdateAuthMethodOAuth2) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGatewayUrl

`func (o *UpdateAuthMethodOAuth2) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *UpdateAuthMethodOAuth2) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *UpdateAuthMethodOAuth2) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *UpdateAuthMethodOAuth2) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *UpdateAuthMethodOAuth2) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *UpdateAuthMethodOAuth2) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *UpdateAuthMethodOAuth2) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *UpdateAuthMethodOAuth2) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetIssuer

`func (o *UpdateAuthMethodOAuth2) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *UpdateAuthMethodOAuth2) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *UpdateAuthMethodOAuth2) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *UpdateAuthMethodOAuth2) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJson

`func (o *UpdateAuthMethodOAuth2) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateAuthMethodOAuth2) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateAuthMethodOAuth2) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateAuthMethodOAuth2) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwksJsonData

`func (o *UpdateAuthMethodOAuth2) GetJwksJsonData() string`

GetJwksJsonData returns the JwksJsonData field if non-nil, zero value otherwise.

### GetJwksJsonDataOk

`func (o *UpdateAuthMethodOAuth2) GetJwksJsonDataOk() (*string, bool)`

GetJwksJsonDataOk returns a tuple with the JwksJsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksJsonData

`func (o *UpdateAuthMethodOAuth2) SetJwksJsonData(v string)`

SetJwksJsonData sets JwksJsonData field to given value.

### HasJwksJsonData

`func (o *UpdateAuthMethodOAuth2) HasJwksJsonData() bool`

HasJwksJsonData returns a boolean if a field has been set.

### GetJwksUri

`func (o *UpdateAuthMethodOAuth2) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *UpdateAuthMethodOAuth2) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *UpdateAuthMethodOAuth2) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.


### GetJwtTtl

`func (o *UpdateAuthMethodOAuth2) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *UpdateAuthMethodOAuth2) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *UpdateAuthMethodOAuth2) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *UpdateAuthMethodOAuth2) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *UpdateAuthMethodOAuth2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthMethodOAuth2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthMethodOAuth2) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateAuthMethodOAuth2) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateAuthMethodOAuth2) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateAuthMethodOAuth2) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateAuthMethodOAuth2) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetSubclaimsDelimiters

`func (o *UpdateAuthMethodOAuth2) GetSubclaimsDelimiters() []string`

GetSubclaimsDelimiters returns the SubclaimsDelimiters field if non-nil, zero value otherwise.

### GetSubclaimsDelimitersOk

`func (o *UpdateAuthMethodOAuth2) GetSubclaimsDelimitersOk() (*[]string, bool)`

GetSubclaimsDelimitersOk returns a tuple with the SubclaimsDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclaimsDelimiters

`func (o *UpdateAuthMethodOAuth2) SetSubclaimsDelimiters(v []string)`

SetSubclaimsDelimiters sets SubclaimsDelimiters field to given value.

### HasSubclaimsDelimiters

`func (o *UpdateAuthMethodOAuth2) HasSubclaimsDelimiters() bool`

HasSubclaimsDelimiters returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAuthMethodOAuth2) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAuthMethodOAuth2) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAuthMethodOAuth2) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAuthMethodOAuth2) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAuthMethodOAuth2) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAuthMethodOAuth2) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAuthMethodOAuth2) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAuthMethodOAuth2) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *UpdateAuthMethodOAuth2) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *UpdateAuthMethodOAuth2) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *UpdateAuthMethodOAuth2) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


