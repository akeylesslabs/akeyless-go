# CreateGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**GroupAlias** | **string** | A short group alias | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Group name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserAssignment** | Pointer to **string** | A json string defining the access permission assignment for this client | [optional] 

## Methods

### NewCreateGroup

`func NewCreateGroup(groupAlias string, name string, ) *CreateGroup`

NewCreateGroup instantiates a new CreateGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupWithDefaults

`func NewCreateGroupWithDefaults() *CreateGroup`

NewCreateGroupWithDefaults instantiates a new CreateGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupAlias

`func (o *CreateGroup) GetGroupAlias() string`

GetGroupAlias returns the GroupAlias field if non-nil, zero value otherwise.

### GetGroupAliasOk

`func (o *CreateGroup) GetGroupAliasOk() (*string, bool)`

GetGroupAliasOk returns a tuple with the GroupAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAlias

`func (o *CreateGroup) SetGroupAlias(v string)`

SetGroupAlias sets GroupAlias field to given value.


### GetJson

`func (o *CreateGroup) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateGroup) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateGroup) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateGroup) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *CreateGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroup) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *CreateGroup) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateGroup) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateGroup) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateGroup) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateGroup) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateGroup) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateGroup) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateGroup) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserAssignment

`func (o *CreateGroup) GetUserAssignment() string`

GetUserAssignment returns the UserAssignment field if non-nil, zero value otherwise.

### GetUserAssignmentOk

`func (o *CreateGroup) GetUserAssignmentOk() (*string, bool)`

GetUserAssignmentOk returns a tuple with the UserAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignment

`func (o *CreateGroup) SetUserAssignment(v string)`

SetUserAssignment sets UserAssignment field to given value.

### HasUserAssignment

`func (o *CreateGroup) HasUserAssignment() bool`

HasUserAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


