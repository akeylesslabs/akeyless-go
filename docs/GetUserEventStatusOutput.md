# GetUserEventStatusOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessStatus** | Pointer to **string** |  | [optional] 
**EventCreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUserEventStatusOutput

`func NewGetUserEventStatusOutput() *GetUserEventStatusOutput`

NewGetUserEventStatusOutput instantiates a new GetUserEventStatusOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserEventStatusOutputWithDefaults

`func NewGetUserEventStatusOutputWithDefaults() *GetUserEventStatusOutput`

NewGetUserEventStatusOutputWithDefaults instantiates a new GetUserEventStatusOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessStatus

`func (o *GetUserEventStatusOutput) GetAccessStatus() string`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *GetUserEventStatusOutput) GetAccessStatusOk() (*string, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *GetUserEventStatusOutput) SetAccessStatus(v string)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *GetUserEventStatusOutput) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetEventCreatedAt

`func (o *GetUserEventStatusOutput) GetEventCreatedAt() time.Time`

GetEventCreatedAt returns the EventCreatedAt field if non-nil, zero value otherwise.

### GetEventCreatedAtOk

`func (o *GetUserEventStatusOutput) GetEventCreatedAtOk() (*time.Time, bool)`

GetEventCreatedAtOk returns a tuple with the EventCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedAt

`func (o *GetUserEventStatusOutput) SetEventCreatedAt(v time.Time)`

SetEventCreatedAt sets EventCreatedAt field to given value.

### HasEventCreatedAt

`func (o *GetUserEventStatusOutput) HasEventCreatedAt() bool`

HasEventCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *GetUserEventStatusOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserEventStatusOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserEventStatusOutput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetUserEventStatusOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


