# GetGroupOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GrantedRoles** | Pointer to [**[]GrantedRole**](GrantedRole.md) | GrantedRoles are the roles a member of this group gets at login, through the group&#39;s akeyless_groups claim. Read-only, and filled only for the SCIM bearer that owns the group: a group name on its own says nothing about the access it carries, which is what a reviewer certifying the entitlement has to judge. | [optional] 
**GrantedRolesKnown** | Pointer to **bool** | GrantedRolesKnown is set when GrantedRoles was computed, so an empty list means \&quot;grants no role\&quot; rather than \&quot;not computed\&quot; (another caller, or an older UAM). | [optional] 
**GroupAlias** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**IsSubclaimsWithOperator** | Pointer to **bool** |  | [optional] 
**ModificationDate** | Pointer to **time.Time** |  | [optional] 
**ScimExternalId** | Pointer to **string** |  | [optional] 
**ScimManaged** | Pointer to **bool** |  | [optional] 
**ScimProvisionedAccessId** | Pointer to **string** |  | [optional] 
**UserAssignments** | Pointer to [**[]AccessPermissionAssignment**](AccessPermissionAssignment.md) |  | [optional] 
**Version** | Pointer to **string** | Version identifies the member assignments this group was read with. Pass it back as UpdateGroupInput.ExpectedVersion to make the update conditional. | [optional] 

## Methods

### NewGetGroupOutput

`func NewGetGroupOutput() *GetGroupOutput`

NewGetGroupOutput instantiates a new GetGroupOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupOutputWithDefaults

`func NewGetGroupOutputWithDefaults() *GetGroupOutput`

NewGetGroupOutputWithDefaults instantiates a new GetGroupOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetGroupOutput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGroupOutput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGroupOutput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGroupOutput) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreationDate

`func (o *GetGroupOutput) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *GetGroupOutput) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *GetGroupOutput) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *GetGroupOutput) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *GetGroupOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGroupOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGroupOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGroupOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGrantedRoles

`func (o *GetGroupOutput) GetGrantedRoles() []GrantedRole`

GetGrantedRoles returns the GrantedRoles field if non-nil, zero value otherwise.

### GetGrantedRolesOk

`func (o *GetGroupOutput) GetGrantedRolesOk() (*[]GrantedRole, bool)`

GetGrantedRolesOk returns a tuple with the GrantedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedRoles

`func (o *GetGroupOutput) SetGrantedRoles(v []GrantedRole)`

SetGrantedRoles sets GrantedRoles field to given value.

### HasGrantedRoles

`func (o *GetGroupOutput) HasGrantedRoles() bool`

HasGrantedRoles returns a boolean if a field has been set.

### GetGrantedRolesKnown

`func (o *GetGroupOutput) GetGrantedRolesKnown() bool`

GetGrantedRolesKnown returns the GrantedRolesKnown field if non-nil, zero value otherwise.

### GetGrantedRolesKnownOk

`func (o *GetGroupOutput) GetGrantedRolesKnownOk() (*bool, bool)`

GetGrantedRolesKnownOk returns a tuple with the GrantedRolesKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedRolesKnown

`func (o *GetGroupOutput) SetGrantedRolesKnown(v bool)`

SetGrantedRolesKnown sets GrantedRolesKnown field to given value.

### HasGrantedRolesKnown

`func (o *GetGroupOutput) HasGrantedRolesKnown() bool`

HasGrantedRolesKnown returns a boolean if a field has been set.

### GetGroupAlias

`func (o *GetGroupOutput) GetGroupAlias() string`

GetGroupAlias returns the GroupAlias field if non-nil, zero value otherwise.

### GetGroupAliasOk

`func (o *GetGroupOutput) GetGroupAliasOk() (*string, bool)`

GetGroupAliasOk returns a tuple with the GroupAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAlias

`func (o *GetGroupOutput) SetGroupAlias(v string)`

SetGroupAlias sets GroupAlias field to given value.

### HasGroupAlias

`func (o *GetGroupOutput) HasGroupAlias() bool`

HasGroupAlias returns a boolean if a field has been set.

### GetGroupId

`func (o *GetGroupOutput) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GetGroupOutput) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GetGroupOutput) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GetGroupOutput) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *GetGroupOutput) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetGroupOutput) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetGroupOutput) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetGroupOutput) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetIsSubclaimsWithOperator

`func (o *GetGroupOutput) GetIsSubclaimsWithOperator() bool`

GetIsSubclaimsWithOperator returns the IsSubclaimsWithOperator field if non-nil, zero value otherwise.

### GetIsSubclaimsWithOperatorOk

`func (o *GetGroupOutput) GetIsSubclaimsWithOperatorOk() (*bool, bool)`

GetIsSubclaimsWithOperatorOk returns a tuple with the IsSubclaimsWithOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubclaimsWithOperator

`func (o *GetGroupOutput) SetIsSubclaimsWithOperator(v bool)`

SetIsSubclaimsWithOperator sets IsSubclaimsWithOperator field to given value.

### HasIsSubclaimsWithOperator

`func (o *GetGroupOutput) HasIsSubclaimsWithOperator() bool`

HasIsSubclaimsWithOperator returns a boolean if a field has been set.

### GetModificationDate

`func (o *GetGroupOutput) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *GetGroupOutput) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *GetGroupOutput) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *GetGroupOutput) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetScimExternalId

`func (o *GetGroupOutput) GetScimExternalId() string`

GetScimExternalId returns the ScimExternalId field if non-nil, zero value otherwise.

### GetScimExternalIdOk

`func (o *GetGroupOutput) GetScimExternalIdOk() (*string, bool)`

GetScimExternalIdOk returns a tuple with the ScimExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimExternalId

`func (o *GetGroupOutput) SetScimExternalId(v string)`

SetScimExternalId sets ScimExternalId field to given value.

### HasScimExternalId

`func (o *GetGroupOutput) HasScimExternalId() bool`

HasScimExternalId returns a boolean if a field has been set.

### GetScimManaged

`func (o *GetGroupOutput) GetScimManaged() bool`

GetScimManaged returns the ScimManaged field if non-nil, zero value otherwise.

### GetScimManagedOk

`func (o *GetGroupOutput) GetScimManagedOk() (*bool, bool)`

GetScimManagedOk returns a tuple with the ScimManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimManaged

`func (o *GetGroupOutput) SetScimManaged(v bool)`

SetScimManaged sets ScimManaged field to given value.

### HasScimManaged

`func (o *GetGroupOutput) HasScimManaged() bool`

HasScimManaged returns a boolean if a field has been set.

### GetScimProvisionedAccessId

`func (o *GetGroupOutput) GetScimProvisionedAccessId() string`

GetScimProvisionedAccessId returns the ScimProvisionedAccessId field if non-nil, zero value otherwise.

### GetScimProvisionedAccessIdOk

`func (o *GetGroupOutput) GetScimProvisionedAccessIdOk() (*string, bool)`

GetScimProvisionedAccessIdOk returns a tuple with the ScimProvisionedAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimProvisionedAccessId

`func (o *GetGroupOutput) SetScimProvisionedAccessId(v string)`

SetScimProvisionedAccessId sets ScimProvisionedAccessId field to given value.

### HasScimProvisionedAccessId

`func (o *GetGroupOutput) HasScimProvisionedAccessId() bool`

HasScimProvisionedAccessId returns a boolean if a field has been set.

### GetUserAssignments

`func (o *GetGroupOutput) GetUserAssignments() []AccessPermissionAssignment`

GetUserAssignments returns the UserAssignments field if non-nil, zero value otherwise.

### GetUserAssignmentsOk

`func (o *GetGroupOutput) GetUserAssignmentsOk() (*[]AccessPermissionAssignment, bool)`

GetUserAssignmentsOk returns a tuple with the UserAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignments

`func (o *GetGroupOutput) SetUserAssignments(v []AccessPermissionAssignment)`

SetUserAssignments sets UserAssignments field to given value.

### HasUserAssignments

`func (o *GetGroupOutput) HasUserAssignments() bool`

HasUserAssignments returns a boolean if a field has been set.

### GetVersion

`func (o *GetGroupOutput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetGroupOutput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetGroupOutput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetGroupOutput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


