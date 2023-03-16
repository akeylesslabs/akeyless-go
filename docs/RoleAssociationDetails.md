# RoleAssociationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssocId** | Pointer to **string** |  | [optional] 
**AuthMethodName** | Pointer to **string** |  | [optional] 
**AuthMethodSubClaims** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**SubClaimsCaseSensitive** | Pointer to **bool** |  | [optional] 

## Methods

### NewRoleAssociationDetails

`func NewRoleAssociationDetails() *RoleAssociationDetails`

NewRoleAssociationDetails instantiates a new RoleAssociationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssociationDetailsWithDefaults

`func NewRoleAssociationDetailsWithDefaults() *RoleAssociationDetails`

NewRoleAssociationDetailsWithDefaults instantiates a new RoleAssociationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssocId

`func (o *RoleAssociationDetails) GetAssocId() string`

GetAssocId returns the AssocId field if non-nil, zero value otherwise.

### GetAssocIdOk

`func (o *RoleAssociationDetails) GetAssocIdOk() (*string, bool)`

GetAssocIdOk returns a tuple with the AssocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocId

`func (o *RoleAssociationDetails) SetAssocId(v string)`

SetAssocId sets AssocId field to given value.

### HasAssocId

`func (o *RoleAssociationDetails) HasAssocId() bool`

HasAssocId returns a boolean if a field has been set.

### GetAuthMethodName

`func (o *RoleAssociationDetails) GetAuthMethodName() string`

GetAuthMethodName returns the AuthMethodName field if non-nil, zero value otherwise.

### GetAuthMethodNameOk

`func (o *RoleAssociationDetails) GetAuthMethodNameOk() (*string, bool)`

GetAuthMethodNameOk returns a tuple with the AuthMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodName

`func (o *RoleAssociationDetails) SetAuthMethodName(v string)`

SetAuthMethodName sets AuthMethodName field to given value.

### HasAuthMethodName

`func (o *RoleAssociationDetails) HasAuthMethodName() bool`

HasAuthMethodName returns a boolean if a field has been set.

### GetAuthMethodSubClaims

`func (o *RoleAssociationDetails) GetAuthMethodSubClaims() map[string][]string`

GetAuthMethodSubClaims returns the AuthMethodSubClaims field if non-nil, zero value otherwise.

### GetAuthMethodSubClaimsOk

`func (o *RoleAssociationDetails) GetAuthMethodSubClaimsOk() (*map[string][]string, bool)`

GetAuthMethodSubClaimsOk returns a tuple with the AuthMethodSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodSubClaims

`func (o *RoleAssociationDetails) SetAuthMethodSubClaims(v map[string][]string)`

SetAuthMethodSubClaims sets AuthMethodSubClaims field to given value.

### HasAuthMethodSubClaims

`func (o *RoleAssociationDetails) HasAuthMethodSubClaims() bool`

HasAuthMethodSubClaims returns a boolean if a field has been set.

### GetRoleName

`func (o *RoleAssociationDetails) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *RoleAssociationDetails) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *RoleAssociationDetails) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *RoleAssociationDetails) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetSubClaimsCaseSensitive

`func (o *RoleAssociationDetails) GetSubClaimsCaseSensitive() bool`

GetSubClaimsCaseSensitive returns the SubClaimsCaseSensitive field if non-nil, zero value otherwise.

### GetSubClaimsCaseSensitiveOk

`func (o *RoleAssociationDetails) GetSubClaimsCaseSensitiveOk() (*bool, bool)`

GetSubClaimsCaseSensitiveOk returns a tuple with the SubClaimsCaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaimsCaseSensitive

`func (o *RoleAssociationDetails) SetSubClaimsCaseSensitive(v bool)`

SetSubClaimsCaseSensitive sets SubClaimsCaseSensitive field to given value.

### HasSubClaimsCaseSensitive

`func (o *RoleAssociationDetails) HasSubClaimsCaseSensitive() bool`

HasSubClaimsCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


