# CreateAuthMethodEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**Email** | **string** | An email address to be invited to have access | 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**ProductType** | Pointer to **[]string** | Choose the relevant product type for the auth method [sm, sra, pm, dp, ca] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateAuthMethodEmail

`func NewCreateAuthMethodEmail(email string, name string, ) *CreateAuthMethodEmail`

NewCreateAuthMethodEmail instantiates a new CreateAuthMethodEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodEmailWithDefaults

`func NewCreateAuthMethodEmailWithDefaults() *CreateAuthMethodEmail`

NewCreateAuthMethodEmailWithDefaults instantiates a new CreateAuthMethodEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodEmail) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodEmail) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodEmail) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodEmail) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundIps

`func (o *CreateAuthMethodEmail) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodEmail) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodEmail) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodEmail) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAuthMethodEmail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAuthMethodEmail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAuthMethodEmail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAuthMethodEmail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *CreateAuthMethodEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateAuthMethodEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateAuthMethodEmail) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetForceSubClaims

`func (o *CreateAuthMethodEmail) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodEmail) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodEmail) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodEmail) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *CreateAuthMethodEmail) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *CreateAuthMethodEmail) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *CreateAuthMethodEmail) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *CreateAuthMethodEmail) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *CreateAuthMethodEmail) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateAuthMethodEmail) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateAuthMethodEmail) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateAuthMethodEmail) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *CreateAuthMethodEmail) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *CreateAuthMethodEmail) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *CreateAuthMethodEmail) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *CreateAuthMethodEmail) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodEmail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodEmail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodEmail) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *CreateAuthMethodEmail) GetProductType() []string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *CreateAuthMethodEmail) GetProductTypeOk() (*[]string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *CreateAuthMethodEmail) SetProductType(v []string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *CreateAuthMethodEmail) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetToken

`func (o *CreateAuthMethodEmail) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodEmail) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodEmail) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodEmail) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodEmail) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodEmail) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodEmail) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodEmail) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


