# UpdateGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**GroupAlias** | **string** | A short group alias | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Group name | 
**NewName** | Pointer to **string** | Group new name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserAssignment** | Pointer to **string** | A json string defining the access permission assignment for this client | [optional] 

## Methods

### NewUpdateGroup

`func NewUpdateGroup(groupAlias string, name string, ) *UpdateGroup`

NewUpdateGroup instantiates a new UpdateGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupWithDefaults

`func NewUpdateGroupWithDefaults() *UpdateGroup`

NewUpdateGroupWithDefaults instantiates a new UpdateGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupAlias

`func (o *UpdateGroup) GetGroupAlias() string`

GetGroupAlias returns the GroupAlias field if non-nil, zero value otherwise.

### GetGroupAliasOk

`func (o *UpdateGroup) GetGroupAliasOk() (*string, bool)`

GetGroupAliasOk returns a tuple with the GroupAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAlias

`func (o *UpdateGroup) SetGroupAlias(v string)`

SetGroupAlias sets GroupAlias field to given value.


### GetJson

`func (o *UpdateGroup) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateGroup) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateGroup) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateGroup) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *UpdateGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGroup) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateGroup) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateGroup) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateGroup) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateGroup) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateGroup) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateGroup) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateGroup) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateGroup) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateGroup) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateGroup) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateGroup) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateGroup) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserAssignment

`func (o *UpdateGroup) GetUserAssignment() string`

GetUserAssignment returns the UserAssignment field if non-nil, zero value otherwise.

### GetUserAssignmentOk

`func (o *UpdateGroup) GetUserAssignmentOk() (*string, bool)`

GetUserAssignmentOk returns a tuple with the UserAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignment

`func (o *UpdateGroup) SetUserAssignment(v string)`

SetUserAssignment sets UserAssignment field to given value.

### HasUserAssignment

`func (o *UpdateGroup) HasUserAssignment() bool`

HasUserAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


