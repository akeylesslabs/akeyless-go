# NullTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewNullTime

`func NewNullTime() *NullTime`

NewNullTime instantiates a new NullTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullTimeWithDefaults

`func NewNullTimeWithDefaults() *NullTime`

NewNullTimeWithDefaults instantiates a new NullTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *NullTime) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NullTime) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NullTime) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *NullTime) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValid

`func (o *NullTime) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *NullTime) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *NullTime) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *NullTime) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


