# OIDCAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRedirectURIs** | Pointer to **[]string** | Allowed redirect URIs after the authentication | [optional] 
**BoundClaims** | Pointer to [**[]OIDCCustomClaim**](OIDCCustomClaim.md) | The claims that login is restricted to. | [optional] 
**ClientId** | Pointer to **string** | Client ID | [optional] 
**ClientSecret** | Pointer to **string** | Client Secret | [optional] 
**IsInternal** | Pointer to **bool** | IsInternal indicates whether this is an internal Auth Method where the client has no control over it, or it was created by the client e.g - Sign In with Google will create an OIDC Auth Method with IsInternal&#x3D;true | [optional] 
**Issuer** | Pointer to **string** | Issuer URL | [optional] 
**UniqueIdentifier** | Pointer to **string** | A unique identifier to distinguish different users | [optional] 

## Methods

### NewOIDCAccessRules

`func NewOIDCAccessRules() *OIDCAccessRules`

NewOIDCAccessRules instantiates a new OIDCAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCAccessRulesWithDefaults

`func NewOIDCAccessRulesWithDefaults() *OIDCAccessRules`

NewOIDCAccessRulesWithDefaults instantiates a new OIDCAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedRedirectURIs

`func (o *OIDCAccessRules) GetAllowedRedirectURIs() []string`

GetAllowedRedirectURIs returns the AllowedRedirectURIs field if non-nil, zero value otherwise.

### GetAllowedRedirectURIsOk

`func (o *OIDCAccessRules) GetAllowedRedirectURIsOk() (*[]string, bool)`

GetAllowedRedirectURIsOk returns a tuple with the AllowedRedirectURIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRedirectURIs

`func (o *OIDCAccessRules) SetAllowedRedirectURIs(v []string)`

SetAllowedRedirectURIs sets AllowedRedirectURIs field to given value.

### HasAllowedRedirectURIs

`func (o *OIDCAccessRules) HasAllowedRedirectURIs() bool`

HasAllowedRedirectURIs returns a boolean if a field has been set.

### GetBoundClaims

`func (o *OIDCAccessRules) GetBoundClaims() []OIDCCustomClaim`

GetBoundClaims returns the BoundClaims field if non-nil, zero value otherwise.

### GetBoundClaimsOk

`func (o *OIDCAccessRules) GetBoundClaimsOk() (*[]OIDCCustomClaim, bool)`

GetBoundClaimsOk returns a tuple with the BoundClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundClaims

`func (o *OIDCAccessRules) SetBoundClaims(v []OIDCCustomClaim)`

SetBoundClaims sets BoundClaims field to given value.

### HasBoundClaims

`func (o *OIDCAccessRules) HasBoundClaims() bool`

HasBoundClaims returns a boolean if a field has been set.

### GetClientId

`func (o *OIDCAccessRules) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OIDCAccessRules) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OIDCAccessRules) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OIDCAccessRules) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OIDCAccessRules) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OIDCAccessRules) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OIDCAccessRules) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OIDCAccessRules) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetIsInternal

`func (o *OIDCAccessRules) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *OIDCAccessRules) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *OIDCAccessRules) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.

### HasIsInternal

`func (o *OIDCAccessRules) HasIsInternal() bool`

HasIsInternal returns a boolean if a field has been set.

### GetIssuer

`func (o *OIDCAccessRules) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OIDCAccessRules) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OIDCAccessRules) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OIDCAccessRules) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *OIDCAccessRules) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *OIDCAccessRules) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *OIDCAccessRules) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *OIDCAccessRules) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


