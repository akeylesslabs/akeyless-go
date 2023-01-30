# ListTargetsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPage** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]Target**](Target.md) |  | [optional] 

## Methods

### NewListTargetsOutput

`func NewListTargetsOutput() *ListTargetsOutput`

NewListTargetsOutput instantiates a new ListTargetsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTargetsOutputWithDefaults

`func NewListTargetsOutputWithDefaults() *ListTargetsOutput`

NewListTargetsOutputWithDefaults instantiates a new ListTargetsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPage

`func (o *ListTargetsOutput) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ListTargetsOutput) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ListTargetsOutput) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ListTargetsOutput) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetTargets

`func (o *ListTargetsOutput) GetTargets() []Target`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ListTargetsOutput) GetTargetsOk() (*[]Target, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ListTargetsOutput) SetTargets(v []Target)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ListTargetsOutput) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


