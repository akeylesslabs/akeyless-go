# ListGroupsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]Group**](Group.md) |  | [optional] 
**NextPage** | Pointer to **string** |  | [optional] 

## Methods

### NewListGroupsOutput

`func NewListGroupsOutput() *ListGroupsOutput`

NewListGroupsOutput instantiates a new ListGroupsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGroupsOutputWithDefaults

`func NewListGroupsOutputWithDefaults() *ListGroupsOutput`

NewListGroupsOutputWithDefaults instantiates a new ListGroupsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ListGroupsOutput) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ListGroupsOutput) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ListGroupsOutput) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ListGroupsOutput) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetNextPage

`func (o *ListGroupsOutput) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ListGroupsOutput) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ListGroupsOutput) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ListGroupsOutput) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


