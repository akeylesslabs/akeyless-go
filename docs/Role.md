# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientPermissions** | Pointer to **[]string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**RoleAuthMethodsAssoc** | Pointer to [**[]RoleAuthMethodAssociation**](RoleAuthMethodAssociation.md) |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**Rules**](Rules.md) |  | [optional] 

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientPermissions

`func (o *Role) GetClientPermissions() []string`

GetClientPermissions returns the ClientPermissions field if non-nil, zero value otherwise.

### GetClientPermissionsOk

`func (o *Role) GetClientPermissionsOk() (*[]string, bool)`

GetClientPermissionsOk returns a tuple with the ClientPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPermissions

`func (o *Role) SetClientPermissions(v []string)`

SetClientPermissions sets ClientPermissions field to given value.

### HasClientPermissions

`func (o *Role) HasClientPermissions() bool`

HasClientPermissions returns a boolean if a field has been set.

### GetComment

`func (o *Role) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Role) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Role) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Role) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRoleAuthMethodsAssoc

`func (o *Role) GetRoleAuthMethodsAssoc() []RoleAuthMethodAssociation`

GetRoleAuthMethodsAssoc returns the RoleAuthMethodsAssoc field if non-nil, zero value otherwise.

### GetRoleAuthMethodsAssocOk

`func (o *Role) GetRoleAuthMethodsAssocOk() (*[]RoleAuthMethodAssociation, bool)`

GetRoleAuthMethodsAssocOk returns a tuple with the RoleAuthMethodsAssoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAuthMethodsAssoc

`func (o *Role) SetRoleAuthMethodsAssoc(v []RoleAuthMethodAssociation)`

SetRoleAuthMethodsAssoc sets RoleAuthMethodsAssoc field to given value.

### HasRoleAuthMethodsAssoc

`func (o *Role) HasRoleAuthMethodsAssoc() bool`

HasRoleAuthMethodsAssoc returns a boolean if a field has been set.

### GetRoleName

`func (o *Role) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *Role) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *Role) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *Role) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRules

`func (o *Role) GetRules() Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Role) GetRulesOk() (*Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Role) SetRules(v Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Role) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


