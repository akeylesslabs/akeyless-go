# CreateAuthMethodUniversalIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**DenyInheritance** | Pointer to **bool** | Deny from root to create children | [optional] 
**DenyRotate** | Pointer to **bool** | Deny from the token to rotate | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] 
**Name** | **string** | Auth Method name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | Pointer to **int32** | Token ttl | [optional] [default to 60]
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateAuthMethodUniversalIdentity

`func NewCreateAuthMethodUniversalIdentity(name string, ) *CreateAuthMethodUniversalIdentity`

NewCreateAuthMethodUniversalIdentity instantiates a new CreateAuthMethodUniversalIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodUniversalIdentityWithDefaults

`func NewCreateAuthMethodUniversalIdentityWithDefaults() *CreateAuthMethodUniversalIdentity`

NewCreateAuthMethodUniversalIdentityWithDefaults instantiates a new CreateAuthMethodUniversalIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodUniversalIdentity) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodUniversalIdentity) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodUniversalIdentity) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodUniversalIdentity) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundIps

`func (o *CreateAuthMethodUniversalIdentity) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodUniversalIdentity) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodUniversalIdentity) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodUniversalIdentity) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetDenyInheritance

`func (o *CreateAuthMethodUniversalIdentity) GetDenyInheritance() bool`

GetDenyInheritance returns the DenyInheritance field if non-nil, zero value otherwise.

### GetDenyInheritanceOk

`func (o *CreateAuthMethodUniversalIdentity) GetDenyInheritanceOk() (*bool, bool)`

GetDenyInheritanceOk returns a tuple with the DenyInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyInheritance

`func (o *CreateAuthMethodUniversalIdentity) SetDenyInheritance(v bool)`

SetDenyInheritance sets DenyInheritance field to given value.

### HasDenyInheritance

`func (o *CreateAuthMethodUniversalIdentity) HasDenyInheritance() bool`

HasDenyInheritance returns a boolean if a field has been set.

### GetDenyRotate

`func (o *CreateAuthMethodUniversalIdentity) GetDenyRotate() bool`

GetDenyRotate returns the DenyRotate field if non-nil, zero value otherwise.

### GetDenyRotateOk

`func (o *CreateAuthMethodUniversalIdentity) GetDenyRotateOk() (*bool, bool)`

GetDenyRotateOk returns a tuple with the DenyRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyRotate

`func (o *CreateAuthMethodUniversalIdentity) SetDenyRotate(v bool)`

SetDenyRotate sets DenyRotate field to given value.

### HasDenyRotate

`func (o *CreateAuthMethodUniversalIdentity) HasDenyRotate() bool`

HasDenyRotate returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *CreateAuthMethodUniversalIdentity) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodUniversalIdentity) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodUniversalIdentity) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodUniversalIdentity) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *CreateAuthMethodUniversalIdentity) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *CreateAuthMethodUniversalIdentity) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *CreateAuthMethodUniversalIdentity) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *CreateAuthMethodUniversalIdentity) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJwtTtl

`func (o *CreateAuthMethodUniversalIdentity) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *CreateAuthMethodUniversalIdentity) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *CreateAuthMethodUniversalIdentity) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *CreateAuthMethodUniversalIdentity) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodUniversalIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodUniversalIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodUniversalIdentity) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *CreateAuthMethodUniversalIdentity) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodUniversalIdentity) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodUniversalIdentity) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodUniversalIdentity) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *CreateAuthMethodUniversalIdentity) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateAuthMethodUniversalIdentity) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateAuthMethodUniversalIdentity) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateAuthMethodUniversalIdentity) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodUniversalIdentity) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodUniversalIdentity) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodUniversalIdentity) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodUniversalIdentity) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


