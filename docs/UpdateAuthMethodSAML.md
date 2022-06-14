# UpdateAuthMethodSAML

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AllowedRedirectUri** | Pointer to **[]string** | Allowed redirect URIs after the authentication | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**IdpMetadataUrl** | Pointer to **string** | IDP metadata url | [optional] 
**IdpMetadataXmlData** | Pointer to **string** | IDP metadata xml data | [optional] 
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] 
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | **string** | A unique identifier (ID) value should be configured for OAuth2, LDAP and SAML authentication method types and is usually a value such as the email, username, or upn for example. Whenever a user logs in with a token, these authentication types issue a \&quot;sub claim\&quot; that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization. | 

## Methods

### NewUpdateAuthMethodSAML

`func NewUpdateAuthMethodSAML(name string, uniqueIdentifier string, ) *UpdateAuthMethodSAML`

NewUpdateAuthMethodSAML instantiates a new UpdateAuthMethodSAML object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthMethodSAMLWithDefaults

`func NewUpdateAuthMethodSAMLWithDefaults() *UpdateAuthMethodSAML`

NewUpdateAuthMethodSAMLWithDefaults instantiates a new UpdateAuthMethodSAML object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *UpdateAuthMethodSAML) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *UpdateAuthMethodSAML) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *UpdateAuthMethodSAML) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *UpdateAuthMethodSAML) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAllowedRedirectUri

`func (o *UpdateAuthMethodSAML) GetAllowedRedirectUri() []string`

GetAllowedRedirectUri returns the AllowedRedirectUri field if non-nil, zero value otherwise.

### GetAllowedRedirectUriOk

`func (o *UpdateAuthMethodSAML) GetAllowedRedirectUriOk() (*[]string, bool)`

GetAllowedRedirectUriOk returns a tuple with the AllowedRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRedirectUri

`func (o *UpdateAuthMethodSAML) SetAllowedRedirectUri(v []string)`

SetAllowedRedirectUri sets AllowedRedirectUri field to given value.

### HasAllowedRedirectUri

`func (o *UpdateAuthMethodSAML) HasAllowedRedirectUri() bool`

HasAllowedRedirectUri returns a boolean if a field has been set.

### GetBoundIps

`func (o *UpdateAuthMethodSAML) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAuthMethodSAML) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAuthMethodSAML) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAuthMethodSAML) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *UpdateAuthMethodSAML) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *UpdateAuthMethodSAML) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *UpdateAuthMethodSAML) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *UpdateAuthMethodSAML) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *UpdateAuthMethodSAML) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *UpdateAuthMethodSAML) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *UpdateAuthMethodSAML) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *UpdateAuthMethodSAML) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetIdpMetadataUrl

`func (o *UpdateAuthMethodSAML) GetIdpMetadataUrl() string`

GetIdpMetadataUrl returns the IdpMetadataUrl field if non-nil, zero value otherwise.

### GetIdpMetadataUrlOk

`func (o *UpdateAuthMethodSAML) GetIdpMetadataUrlOk() (*string, bool)`

GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataUrl

`func (o *UpdateAuthMethodSAML) SetIdpMetadataUrl(v string)`

SetIdpMetadataUrl sets IdpMetadataUrl field to given value.

### HasIdpMetadataUrl

`func (o *UpdateAuthMethodSAML) HasIdpMetadataUrl() bool`

HasIdpMetadataUrl returns a boolean if a field has been set.

### GetIdpMetadataXmlData

`func (o *UpdateAuthMethodSAML) GetIdpMetadataXmlData() string`

GetIdpMetadataXmlData returns the IdpMetadataXmlData field if non-nil, zero value otherwise.

### GetIdpMetadataXmlDataOk

`func (o *UpdateAuthMethodSAML) GetIdpMetadataXmlDataOk() (*string, bool)`

GetIdpMetadataXmlDataOk returns a tuple with the IdpMetadataXmlData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXmlData

`func (o *UpdateAuthMethodSAML) SetIdpMetadataXmlData(v string)`

SetIdpMetadataXmlData sets IdpMetadataXmlData field to given value.

### HasIdpMetadataXmlData

`func (o *UpdateAuthMethodSAML) HasIdpMetadataXmlData() bool`

HasIdpMetadataXmlData returns a boolean if a field has been set.

### GetJwtTtl

`func (o *UpdateAuthMethodSAML) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *UpdateAuthMethodSAML) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *UpdateAuthMethodSAML) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *UpdateAuthMethodSAML) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *UpdateAuthMethodSAML) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthMethodSAML) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthMethodSAML) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateAuthMethodSAML) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateAuthMethodSAML) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateAuthMethodSAML) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateAuthMethodSAML) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAuthMethodSAML) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAuthMethodSAML) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAuthMethodSAML) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAuthMethodSAML) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAuthMethodSAML) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAuthMethodSAML) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAuthMethodSAML) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAuthMethodSAML) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *UpdateAuthMethodSAML) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *UpdateAuthMethodSAML) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *UpdateAuthMethodSAML) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


