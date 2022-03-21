# CreateAuthMethodLDAP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**PublicKeyData** | Pointer to **string** | A public key generated for LDAP authentication method on Akeyless in base64 format [RSA2048] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateAuthMethodLDAP

`func NewCreateAuthMethodLDAP(name string, ) *CreateAuthMethodLDAP`

NewCreateAuthMethodLDAP instantiates a new CreateAuthMethodLDAP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodLDAPWithDefaults

`func NewCreateAuthMethodLDAPWithDefaults() *CreateAuthMethodLDAP`

NewCreateAuthMethodLDAPWithDefaults instantiates a new CreateAuthMethodLDAP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodLDAP) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodLDAP) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodLDAP) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodLDAP) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundIps

`func (o *CreateAuthMethodLDAP) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodLDAP) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodLDAP) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodLDAP) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *CreateAuthMethodLDAP) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodLDAP) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodLDAP) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodLDAP) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetJwtTtl

`func (o *CreateAuthMethodLDAP) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *CreateAuthMethodLDAP) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *CreateAuthMethodLDAP) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *CreateAuthMethodLDAP) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodLDAP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodLDAP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodLDAP) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKeyData

`func (o *CreateAuthMethodLDAP) GetPublicKeyData() string`

GetPublicKeyData returns the PublicKeyData field if non-nil, zero value otherwise.

### GetPublicKeyDataOk

`func (o *CreateAuthMethodLDAP) GetPublicKeyDataOk() (*string, bool)`

GetPublicKeyDataOk returns a tuple with the PublicKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyData

`func (o *CreateAuthMethodLDAP) SetPublicKeyData(v string)`

SetPublicKeyData sets PublicKeyData field to given value.

### HasPublicKeyData

`func (o *CreateAuthMethodLDAP) HasPublicKeyData() bool`

HasPublicKeyData returns a boolean if a field has been set.

### GetToken

`func (o *CreateAuthMethodLDAP) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodLDAP) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodLDAP) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodLDAP) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodLDAP) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodLDAP) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodLDAP) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodLDAP) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


