# BastionsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]BastionListEntry**](BastionListEntry.md) |  | [optional] 

## Methods

### NewBastionsList

`func NewBastionsList() *BastionsList`

NewBastionsList instantiates a new BastionsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBastionsListWithDefaults

`func NewBastionsListWithDefaults() *BastionsList`

NewBastionsListWithDefaults instantiates a new BastionsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *BastionsList) GetClusters() []BastionListEntry`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *BastionsList) GetClustersOk() (*[]BastionListEntry, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *BastionsList) SetClusters(v []BastionListEntry)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *BastionsList) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


