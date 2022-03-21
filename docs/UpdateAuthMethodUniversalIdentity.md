# UpdateAuthMethodUniversalIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**DenyInheritance** | Pointer to **bool** | Deny from root to create children | [optional] 
**DenyRotate** | Pointer to **bool** | Deny from the token to rotate | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | Pointer to **int32** | Token ttl | [optional] [default to 60]
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateAuthMethodUniversalIdentity

`func NewUpdateAuthMethodUniversalIdentity(name string, ) *UpdateAuthMethodUniversalIdentity`

NewUpdateAuthMethodUniversalIdentity instantiates a new UpdateAuthMethodUniversalIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthMethodUniversalIdentityWithDefaults

`func NewUpdateAuthMethodUniversalIdentityWithDefaults() *UpdateAuthMethodUniversalIdentity`

NewUpdateAuthMethodUniversalIdentityWithDefaults instantiates a new UpdateAuthMethodUniversalIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *UpdateAuthMethodUniversalIdentity) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *UpdateAuthMethodUniversalIdentity) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *UpdateAuthMethodUniversalIdentity) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *UpdateAuthMethodUniversalIdentity) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundIps

`func (o *UpdateAuthMethodUniversalIdentity) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAuthMethodUniversalIdentity) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAuthMethodUniversalIdentity) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAuthMethodUniversalIdentity) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetDenyInheritance

`func (o *UpdateAuthMethodUniversalIdentity) GetDenyInheritance() bool`

GetDenyInheritance returns the DenyInheritance field if non-nil, zero value otherwise.

### GetDenyInheritanceOk

`func (o *UpdateAuthMethodUniversalIdentity) GetDenyInheritanceOk() (*bool, bool)`

GetDenyInheritanceOk returns a tuple with the DenyInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyInheritance

`func (o *UpdateAuthMethodUniversalIdentity) SetDenyInheritance(v bool)`

SetDenyInheritance sets DenyInheritance field to given value.

### HasDenyInheritance

`func (o *UpdateAuthMethodUniversalIdentity) HasDenyInheritance() bool`

HasDenyInheritance returns a boolean if a field has been set.

### GetDenyRotate

`func (o *UpdateAuthMethodUniversalIdentity) GetDenyRotate() bool`

GetDenyRotate returns the DenyRotate field if non-nil, zero value otherwise.

### GetDenyRotateOk

`func (o *UpdateAuthMethodUniversalIdentity) GetDenyRotateOk() (*bool, bool)`

GetDenyRotateOk returns a tuple with the DenyRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyRotate

`func (o *UpdateAuthMethodUniversalIdentity) SetDenyRotate(v bool)`

SetDenyRotate sets DenyRotate field to given value.

### HasDenyRotate

`func (o *UpdateAuthMethodUniversalIdentity) HasDenyRotate() bool`

HasDenyRotate returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *UpdateAuthMethodUniversalIdentity) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *UpdateAuthMethodUniversalIdentity) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *UpdateAuthMethodUniversalIdentity) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *UpdateAuthMethodUniversalIdentity) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetJwtTtl

`func (o *UpdateAuthMethodUniversalIdentity) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *UpdateAuthMethodUniversalIdentity) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *UpdateAuthMethodUniversalIdentity) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *UpdateAuthMethodUniversalIdentity) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *UpdateAuthMethodUniversalIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthMethodUniversalIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthMethodUniversalIdentity) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateAuthMethodUniversalIdentity) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateAuthMethodUniversalIdentity) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateAuthMethodUniversalIdentity) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateAuthMethodUniversalIdentity) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAuthMethodUniversalIdentity) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAuthMethodUniversalIdentity) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAuthMethodUniversalIdentity) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAuthMethodUniversalIdentity) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *UpdateAuthMethodUniversalIdentity) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *UpdateAuthMethodUniversalIdentity) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *UpdateAuthMethodUniversalIdentity) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *UpdateAuthMethodUniversalIdentity) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAuthMethodUniversalIdentity) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAuthMethodUniversalIdentity) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAuthMethodUniversalIdentity) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAuthMethodUniversalIdentity) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


