# ListRolesOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPage** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) |  | [optional] 

## Methods

### NewListRolesOutput

`func NewListRolesOutput() *ListRolesOutput`

NewListRolesOutput instantiates a new ListRolesOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRolesOutputWithDefaults

`func NewListRolesOutputWithDefaults() *ListRolesOutput`

NewListRolesOutputWithDefaults instantiates a new ListRolesOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPage

`func (o *ListRolesOutput) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ListRolesOutput) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ListRolesOutput) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ListRolesOutput) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetRoles

`func (o *ListRolesOutput) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListRolesOutput) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListRolesOutput) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ListRolesOutput) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


