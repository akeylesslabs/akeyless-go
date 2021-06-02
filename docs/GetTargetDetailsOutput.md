# GetTargetDetailsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to [**Target**](Target.md) |  | [optional] 
**Value** | Pointer to [**TargetTypeDetailsInput**](TargetTypeDetailsInput.md) |  | [optional] 

## Methods

### NewGetTargetDetailsOutput

`func NewGetTargetDetailsOutput() *GetTargetDetailsOutput`

NewGetTargetDetailsOutput instantiates a new GetTargetDetailsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTargetDetailsOutputWithDefaults

`func NewGetTargetDetailsOutputWithDefaults() *GetTargetDetailsOutput`

NewGetTargetDetailsOutputWithDefaults instantiates a new GetTargetDetailsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *GetTargetDetailsOutput) GetTarget() Target`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetTargetDetailsOutput) GetTargetOk() (*Target, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetTargetDetailsOutput) SetTarget(v Target)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GetTargetDetailsOutput) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetValue

`func (o *GetTargetDetailsOutput) GetValue() TargetTypeDetailsInput`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTargetDetailsOutput) GetValueOk() (*TargetTypeDetailsInput, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTargetDetailsOutput) SetValue(v TargetTypeDetailsInput)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetTargetDetailsOutput) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


