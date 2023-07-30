# AuthMethodRoleAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOps** | Pointer to **[]string** |  | [optional] 
**AssocId** | Pointer to **string** |  | [optional] 
**AuthMethodSubClaims** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**IsSubClaimsCaseSensitive** | Pointer to **bool** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**Rules**](Rules.md) |  | [optional] 

## Methods

### NewAuthMethodRoleAssociation

`func NewAuthMethodRoleAssociation() *AuthMethodRoleAssociation`

NewAuthMethodRoleAssociation instantiates a new AuthMethodRoleAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodRoleAssociationWithDefaults

`func NewAuthMethodRoleAssociationWithDefaults() *AuthMethodRoleAssociation`

NewAuthMethodRoleAssociationWithDefaults instantiates a new AuthMethodRoleAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOps

`func (o *AuthMethodRoleAssociation) GetAllowedOps() []string`

GetAllowedOps returns the AllowedOps field if non-nil, zero value otherwise.

### GetAllowedOpsOk

`func (o *AuthMethodRoleAssociation) GetAllowedOpsOk() (*[]string, bool)`

GetAllowedOpsOk returns a tuple with the AllowedOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOps

`func (o *AuthMethodRoleAssociation) SetAllowedOps(v []string)`

SetAllowedOps sets AllowedOps field to given value.

### HasAllowedOps

`func (o *AuthMethodRoleAssociation) HasAllowedOps() bool`

HasAllowedOps returns a boolean if a field has been set.

### GetAssocId

`func (o *AuthMethodRoleAssociation) GetAssocId() string`

GetAssocId returns the AssocId field if non-nil, zero value otherwise.

### GetAssocIdOk

`func (o *AuthMethodRoleAssociation) GetAssocIdOk() (*string, bool)`

GetAssocIdOk returns a tuple with the AssocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocId

`func (o *AuthMethodRoleAssociation) SetAssocId(v string)`

SetAssocId sets AssocId field to given value.

### HasAssocId

`func (o *AuthMethodRoleAssociation) HasAssocId() bool`

HasAssocId returns a boolean if a field has been set.

### GetAuthMethodSubClaims

`func (o *AuthMethodRoleAssociation) GetAuthMethodSubClaims() map[string][]string`

GetAuthMethodSubClaims returns the AuthMethodSubClaims field if non-nil, zero value otherwise.

### GetAuthMethodSubClaimsOk

`func (o *AuthMethodRoleAssociation) GetAuthMethodSubClaimsOk() (*map[string][]string, bool)`

GetAuthMethodSubClaimsOk returns a tuple with the AuthMethodSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodSubClaims

`func (o *AuthMethodRoleAssociation) SetAuthMethodSubClaims(v map[string][]string)`

SetAuthMethodSubClaims sets AuthMethodSubClaims field to given value.

### HasAuthMethodSubClaims

`func (o *AuthMethodRoleAssociation) HasAuthMethodSubClaims() bool`

HasAuthMethodSubClaims returns a boolean if a field has been set.

### GetIsSubClaimsCaseSensitive

`func (o *AuthMethodRoleAssociation) GetIsSubClaimsCaseSensitive() bool`

GetIsSubClaimsCaseSensitive returns the IsSubClaimsCaseSensitive field if non-nil, zero value otherwise.

### GetIsSubClaimsCaseSensitiveOk

`func (o *AuthMethodRoleAssociation) GetIsSubClaimsCaseSensitiveOk() (*bool, bool)`

GetIsSubClaimsCaseSensitiveOk returns a tuple with the IsSubClaimsCaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubClaimsCaseSensitive

`func (o *AuthMethodRoleAssociation) SetIsSubClaimsCaseSensitive(v bool)`

SetIsSubClaimsCaseSensitive sets IsSubClaimsCaseSensitive field to given value.

### HasIsSubClaimsCaseSensitive

`func (o *AuthMethodRoleAssociation) HasIsSubClaimsCaseSensitive() bool`

HasIsSubClaimsCaseSensitive returns a boolean if a field has been set.

### GetRoleName

`func (o *AuthMethodRoleAssociation) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AuthMethodRoleAssociation) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AuthMethodRoleAssociation) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AuthMethodRoleAssociation) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRules

`func (o *AuthMethodRoleAssociation) GetRules() Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AuthMethodRoleAssociation) GetRulesOk() (*Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AuthMethodRoleAssociation) SetRules(v Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AuthMethodRoleAssociation) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


