# AuthExpirationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecondsBefore** | Pointer to **int64** |  | [optional] 

## Methods

### NewAuthExpirationEvent

`func NewAuthExpirationEvent() *AuthExpirationEvent`

NewAuthExpirationEvent instantiates a new AuthExpirationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthExpirationEventWithDefaults

`func NewAuthExpirationEventWithDefaults() *AuthExpirationEvent`

NewAuthExpirationEventWithDefaults instantiates a new AuthExpirationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecondsBefore

`func (o *AuthExpirationEvent) GetSecondsBefore() int64`

GetSecondsBefore returns the SecondsBefore field if non-nil, zero value otherwise.

### GetSecondsBeforeOk

`func (o *AuthExpirationEvent) GetSecondsBeforeOk() (*int64, bool)`

GetSecondsBeforeOk returns a tuple with the SecondsBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsBefore

`func (o *AuthExpirationEvent) SetSecondsBefore(v int64)`

SetSecondsBefore sets SecondsBefore field to given value.

### HasSecondsBefore

`func (o *AuthExpirationEvent) HasSecondsBefore() bool`

HasSecondsBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


