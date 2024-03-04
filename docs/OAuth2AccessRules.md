# OAuth2AccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | The audience in the JWT. | [optional] 
**AuthorizedGwClusterName** | Pointer to **string** | The gateway cluster name that is authorized to access JWKeySetURL | [optional] 
**BoundClaims** | Pointer to [**[]OAuth2CustomClaim**](OAuth2CustomClaim.md) | The claims that login is restricted to. | [optional] 
**BoundClientsId** | Pointer to **[]string** | The clients ids that login is restricted to. | [optional] 
**Certificate** | Pointer to **string** | Certificate to use when calling jwks_uri from the gateway. in PEM format | [optional] 
**Issuer** | Pointer to **string** | Issuer URL | [optional] 
**JwksJsonData** | Pointer to **string** | The JSON Web Key Set (JWKS) that containing the public keys that should be used to verify any JSON Web Token (JWT) issued by the authorization server. base64 encoded string | [optional] 
**JwksUri** | Pointer to **string** | The URL to the JSON Web Key Set (JWKS) that containing the public keys that should be used to verify any JSON Web Token (JWT) issued by the authorization server. | [optional] 
**UniqueIdentifier** | Pointer to **string** | A unique identifier to distinguish different users | [optional] 

## Methods

### NewOAuth2AccessRules

`func NewOAuth2AccessRules() *OAuth2AccessRules`

NewOAuth2AccessRules instantiates a new OAuth2AccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2AccessRulesWithDefaults

`func NewOAuth2AccessRulesWithDefaults() *OAuth2AccessRules`

NewOAuth2AccessRulesWithDefaults instantiates a new OAuth2AccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *OAuth2AccessRules) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *OAuth2AccessRules) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *OAuth2AccessRules) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *OAuth2AccessRules) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetAuthorizedGwClusterName

`func (o *OAuth2AccessRules) GetAuthorizedGwClusterName() string`

GetAuthorizedGwClusterName returns the AuthorizedGwClusterName field if non-nil, zero value otherwise.

### GetAuthorizedGwClusterNameOk

`func (o *OAuth2AccessRules) GetAuthorizedGwClusterNameOk() (*string, bool)`

GetAuthorizedGwClusterNameOk returns a tuple with the AuthorizedGwClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedGwClusterName

`func (o *OAuth2AccessRules) SetAuthorizedGwClusterName(v string)`

SetAuthorizedGwClusterName sets AuthorizedGwClusterName field to given value.

### HasAuthorizedGwClusterName

`func (o *OAuth2AccessRules) HasAuthorizedGwClusterName() bool`

HasAuthorizedGwClusterName returns a boolean if a field has been set.

### GetBoundClaims

`func (o *OAuth2AccessRules) GetBoundClaims() []OAuth2CustomClaim`

GetBoundClaims returns the BoundClaims field if non-nil, zero value otherwise.

### GetBoundClaimsOk

`func (o *OAuth2AccessRules) GetBoundClaimsOk() (*[]OAuth2CustomClaim, bool)`

GetBoundClaimsOk returns a tuple with the BoundClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundClaims

`func (o *OAuth2AccessRules) SetBoundClaims(v []OAuth2CustomClaim)`

SetBoundClaims sets BoundClaims field to given value.

### HasBoundClaims

`func (o *OAuth2AccessRules) HasBoundClaims() bool`

HasBoundClaims returns a boolean if a field has been set.

### GetBoundClientsId

`func (o *OAuth2AccessRules) GetBoundClientsId() []string`

GetBoundClientsId returns the BoundClientsId field if non-nil, zero value otherwise.

### GetBoundClientsIdOk

`func (o *OAuth2AccessRules) GetBoundClientsIdOk() (*[]string, bool)`

GetBoundClientsIdOk returns a tuple with the BoundClientsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundClientsId

`func (o *OAuth2AccessRules) SetBoundClientsId(v []string)`

SetBoundClientsId sets BoundClientsId field to given value.

### HasBoundClientsId

`func (o *OAuth2AccessRules) HasBoundClientsId() bool`

HasBoundClientsId returns a boolean if a field has been set.

### GetCertificate

`func (o *OAuth2AccessRules) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *OAuth2AccessRules) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *OAuth2AccessRules) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *OAuth2AccessRules) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetIssuer

`func (o *OAuth2AccessRules) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OAuth2AccessRules) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OAuth2AccessRules) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OAuth2AccessRules) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJwksJsonData

`func (o *OAuth2AccessRules) GetJwksJsonData() string`

GetJwksJsonData returns the JwksJsonData field if non-nil, zero value otherwise.

### GetJwksJsonDataOk

`func (o *OAuth2AccessRules) GetJwksJsonDataOk() (*string, bool)`

GetJwksJsonDataOk returns a tuple with the JwksJsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksJsonData

`func (o *OAuth2AccessRules) SetJwksJsonData(v string)`

SetJwksJsonData sets JwksJsonData field to given value.

### HasJwksJsonData

`func (o *OAuth2AccessRules) HasJwksJsonData() bool`

HasJwksJsonData returns a boolean if a field has been set.

### GetJwksUri

`func (o *OAuth2AccessRules) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *OAuth2AccessRules) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *OAuth2AccessRules) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *OAuth2AccessRules) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *OAuth2AccessRules) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *OAuth2AccessRules) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *OAuth2AccessRules) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *OAuth2AccessRules) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


