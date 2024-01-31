# UpdateAuthMethodLDAP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GenKey** | Pointer to **string** | Automatically generate key-pair for LDAP configuration. If set to false, a public key needs to be provided [true/false] | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**PublicKeyData** | Pointer to **string** | A public key generated for LDAP authentication method on Akeyless in base64 or PEM format [RSA2048] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UniqueIdentifier** | Pointer to **string** | A unique identifier (ID) value should be configured for OAuth2, LDAP and SAML authentication method types and is usually a value such as the email, username, or upn for example. Whenever a user logs in with a token, these authentication types issue a \&quot;sub claim\&quot; that contains details uniquely identifying that user. This sub claim includes a key containing the ID value that you configured, and is used to distinguish between different users from within the same organization. | [optional] [default to "users"]

## Methods

### NewUpdateAuthMethodLDAP

`func NewUpdateAuthMethodLDAP(name string, ) *UpdateAuthMethodLDAP`

NewUpdateAuthMethodLDAP instantiates a new UpdateAuthMethodLDAP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthMethodLDAPWithDefaults

`func NewUpdateAuthMethodLDAPWithDefaults() *UpdateAuthMethodLDAP`

NewUpdateAuthMethodLDAPWithDefaults instantiates a new UpdateAuthMethodLDAP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *UpdateAuthMethodLDAP) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *UpdateAuthMethodLDAP) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *UpdateAuthMethodLDAP) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *UpdateAuthMethodLDAP) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundIps

`func (o *UpdateAuthMethodLDAP) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAuthMethodLDAP) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAuthMethodLDAP) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAuthMethodLDAP) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAuthMethodLDAP) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAuthMethodLDAP) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAuthMethodLDAP) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAuthMethodLDAP) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *UpdateAuthMethodLDAP) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *UpdateAuthMethodLDAP) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *UpdateAuthMethodLDAP) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *UpdateAuthMethodLDAP) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGenKey

`func (o *UpdateAuthMethodLDAP) GetGenKey() string`

GetGenKey returns the GenKey field if non-nil, zero value otherwise.

### GetGenKeyOk

`func (o *UpdateAuthMethodLDAP) GetGenKeyOk() (*string, bool)`

GetGenKeyOk returns a tuple with the GenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenKey

`func (o *UpdateAuthMethodLDAP) SetGenKey(v string)`

SetGenKey sets GenKey field to given value.

### HasGenKey

`func (o *UpdateAuthMethodLDAP) HasGenKey() bool`

HasGenKey returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *UpdateAuthMethodLDAP) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *UpdateAuthMethodLDAP) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *UpdateAuthMethodLDAP) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *UpdateAuthMethodLDAP) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *UpdateAuthMethodLDAP) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateAuthMethodLDAP) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateAuthMethodLDAP) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateAuthMethodLDAP) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *UpdateAuthMethodLDAP) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *UpdateAuthMethodLDAP) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *UpdateAuthMethodLDAP) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *UpdateAuthMethodLDAP) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *UpdateAuthMethodLDAP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthMethodLDAP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthMethodLDAP) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateAuthMethodLDAP) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateAuthMethodLDAP) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateAuthMethodLDAP) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateAuthMethodLDAP) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPublicKeyData

`func (o *UpdateAuthMethodLDAP) GetPublicKeyData() string`

GetPublicKeyData returns the PublicKeyData field if non-nil, zero value otherwise.

### GetPublicKeyDataOk

`func (o *UpdateAuthMethodLDAP) GetPublicKeyDataOk() (*string, bool)`

GetPublicKeyDataOk returns a tuple with the PublicKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyData

`func (o *UpdateAuthMethodLDAP) SetPublicKeyData(v string)`

SetPublicKeyData sets PublicKeyData field to given value.

### HasPublicKeyData

`func (o *UpdateAuthMethodLDAP) HasPublicKeyData() bool`

HasPublicKeyData returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAuthMethodLDAP) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAuthMethodLDAP) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAuthMethodLDAP) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAuthMethodLDAP) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAuthMethodLDAP) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAuthMethodLDAP) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAuthMethodLDAP) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAuthMethodLDAP) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *UpdateAuthMethodLDAP) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *UpdateAuthMethodLDAP) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *UpdateAuthMethodLDAP) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *UpdateAuthMethodLDAP) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


