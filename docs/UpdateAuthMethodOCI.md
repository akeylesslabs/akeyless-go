# UpdateAuthMethodOCI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**Description** | Pointer to **string** | Auth Method description | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GroupOcid** | **[]string** | A list of required groups ocids | 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**TenantOcid** | **string** | The Oracle Cloud tenant ID | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateAuthMethodOCI

`func NewUpdateAuthMethodOCI(groupOcid []string, name string, tenantOcid string, ) *UpdateAuthMethodOCI`

NewUpdateAuthMethodOCI instantiates a new UpdateAuthMethodOCI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthMethodOCIWithDefaults

`func NewUpdateAuthMethodOCIWithDefaults() *UpdateAuthMethodOCI`

NewUpdateAuthMethodOCIWithDefaults instantiates a new UpdateAuthMethodOCI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *UpdateAuthMethodOCI) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *UpdateAuthMethodOCI) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *UpdateAuthMethodOCI) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *UpdateAuthMethodOCI) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundIps

`func (o *UpdateAuthMethodOCI) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAuthMethodOCI) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAuthMethodOCI) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAuthMethodOCI) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAuthMethodOCI) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAuthMethodOCI) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAuthMethodOCI) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAuthMethodOCI) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *UpdateAuthMethodOCI) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *UpdateAuthMethodOCI) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *UpdateAuthMethodOCI) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *UpdateAuthMethodOCI) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGroupOcid

`func (o *UpdateAuthMethodOCI) GetGroupOcid() []string`

GetGroupOcid returns the GroupOcid field if non-nil, zero value otherwise.

### GetGroupOcidOk

`func (o *UpdateAuthMethodOCI) GetGroupOcidOk() (*[]string, bool)`

GetGroupOcidOk returns a tuple with the GroupOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOcid

`func (o *UpdateAuthMethodOCI) SetGroupOcid(v []string)`

SetGroupOcid sets GroupOcid field to given value.


### GetGwBoundIps

`func (o *UpdateAuthMethodOCI) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *UpdateAuthMethodOCI) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *UpdateAuthMethodOCI) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *UpdateAuthMethodOCI) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *UpdateAuthMethodOCI) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateAuthMethodOCI) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateAuthMethodOCI) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateAuthMethodOCI) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *UpdateAuthMethodOCI) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *UpdateAuthMethodOCI) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *UpdateAuthMethodOCI) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *UpdateAuthMethodOCI) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *UpdateAuthMethodOCI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthMethodOCI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthMethodOCI) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateAuthMethodOCI) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateAuthMethodOCI) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateAuthMethodOCI) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateAuthMethodOCI) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetTenantOcid

`func (o *UpdateAuthMethodOCI) GetTenantOcid() string`

GetTenantOcid returns the TenantOcid field if non-nil, zero value otherwise.

### GetTenantOcidOk

`func (o *UpdateAuthMethodOCI) GetTenantOcidOk() (*string, bool)`

GetTenantOcidOk returns a tuple with the TenantOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOcid

`func (o *UpdateAuthMethodOCI) SetTenantOcid(v string)`

SetTenantOcid sets TenantOcid field to given value.


### GetToken

`func (o *UpdateAuthMethodOCI) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAuthMethodOCI) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAuthMethodOCI) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAuthMethodOCI) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAuthMethodOCI) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAuthMethodOCI) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAuthMethodOCI) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAuthMethodOCI) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


