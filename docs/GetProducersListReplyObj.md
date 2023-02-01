# GetProducersListReplyObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Producers** | Pointer to [**[]Producer**](Producer.md) |  | [optional] 
**ProducersErrors** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetProducersListReplyObj

`func NewGetProducersListReplyObj() *GetProducersListReplyObj`

NewGetProducersListReplyObj instantiates a new GetProducersListReplyObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProducersListReplyObjWithDefaults

`func NewGetProducersListReplyObjWithDefaults() *GetProducersListReplyObj`

NewGetProducersListReplyObjWithDefaults instantiates a new GetProducersListReplyObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducers

`func (o *GetProducersListReplyObj) GetProducers() []Producer`

GetProducers returns the Producers field if non-nil, zero value otherwise.

### GetProducersOk

`func (o *GetProducersListReplyObj) GetProducersOk() (*[]Producer, bool)`

GetProducersOk returns a tuple with the Producers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducers

`func (o *GetProducersListReplyObj) SetProducers(v []Producer)`

SetProducers sets Producers field to given value.

### HasProducers

`func (o *GetProducersListReplyObj) HasProducers() bool`

HasProducers returns a boolean if a field has been set.

### GetProducersErrors

`func (o *GetProducersListReplyObj) GetProducersErrors() map[string]interface{}`

GetProducersErrors returns the ProducersErrors field if non-nil, zero value otherwise.

### GetProducersErrorsOk

`func (o *GetProducersListReplyObj) GetProducersErrorsOk() (*map[string]interface{}, bool)`

GetProducersErrorsOk returns a tuple with the ProducersErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducersErrors

`func (o *GetProducersListReplyObj) SetProducersErrors(v map[string]interface{})`

SetProducersErrors sets ProducersErrors field to given value.

### HasProducersErrors

`func (o *GetProducersListReplyObj) HasProducersErrors() bool`

HasProducersErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


