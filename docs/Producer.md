# Producer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewProducer

`func NewProducer() *Producer`

NewProducer instantiates a new Producer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProducerWithDefaults

`func NewProducerWithDefaults() *Producer`

NewProducerWithDefaults instantiates a new Producer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Producer) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Producer) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Producer) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Producer) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetId

`func (o *Producer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Producer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Producer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Producer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Producer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Producer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Producer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Producer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Producer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Producer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Producer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Producer) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


