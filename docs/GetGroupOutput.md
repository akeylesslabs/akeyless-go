# GetGroupOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GroupAlias** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**ModificationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**UserAssignments** | Pointer to [**[]AccessPermissionAssignment**](AccessPermissionAssignment.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


